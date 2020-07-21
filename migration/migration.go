package lib

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/naufalfmm/mongodb-migration/config"
	"github.com/naufalfmm/mongodb-migration/constants"
	"github.com/naufalfmm/mongodb-migration/driver"
	"github.com/naufalfmm/mongodb-migration/history"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMigration struct {
	Source string
	Driver *driver.MongoDriver

	HistoryCollection *history.MigrationHistory
}

type JSONCommand struct {
	UpCommand   bson.M `bson:"up"`
	DownCommand bson.M `bson:"down"`
}

func (mm *MongoMigration) StartMigration(source string, client *mongo.Client, cfg config.MongoConfig, migrationHistoryCollectionName string) error {
	var mongoDriver driver.MongoDriver

	mongoDriver.SetClient(cfg)

	mm.StartMigrationWithDriver(source, &mongoDriver, migrationHistoryCollectionName)

	return nil
}

func (mm *MongoMigration) StartMigrationWithDriver(source string, driver *driver.MongoDriver, migrationHistoryCollectionName string) error {
	mm.Driver = driver
	mm.Source = source

	historyColl := history.MigrationHistory{
		DB:             mm.Driver.DB,
		CollectionName: migrationHistoryCollectionName,
	}

	mm.HistoryCollection = &historyColl

	ctx := context.TODO()

	err := mm.HistoryCollection.InitializeHistoryCollection(ctx)
	if err != nil {
		cmdErr := err.(mongo.CommandError)
		if cmdErr.Code != 48 {
			return cmdErr
		}
	}

	return nil
}

func (mm *MongoMigration) Run(direction int, steps int) error {
	step := 0

	dir := http.Dir(mm.Source)
	file, err := dir.Open("/")
	if err != nil {
		panic(err)
	}

	files, err := file.Readdir(0)
	if err != nil {
		panic(err)
	}

	for _, info := range files {
		if strings.HasSuffix(info.Name(), ".json") {
			_, err = mm.RunSpecificFile(info.Name(), direction)
			if err != nil {
				return err
			}
			step++
		}

		if step == steps {
			break
		}
	}

	return nil
}

func (mm *MongoMigration) RunSpecificFile(migrationFileName string, direction int) (interface{}, error) {
	var command JSONCommand

	migrationFile, err := os.Open(mm.Source + migrationFileName)
	if err != nil {
		return nil, err
	}
	defer migrationFile.Close()

	migrBytes, err := ioutil.ReadAll(migrationFile)
	if err != nil {
		return nil, err
	}

	err = bson.UnmarshalExtJSON(migrBytes, true, &command)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	directionFunc := mm.executeUp
	usedCommand := command.UpCommand
	if direction == constants.DOWN {
		directionFunc = mm.executeDown
		usedCommand = command.DownCommand
	}

	res, err := directionFunc(ctx, usedCommand, migrationFileName)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (mm *MongoMigration) executeUp(ctx context.Context, upCmd bson.M, migrFileName string) (interface{}, error) {
	mm.HistoryCollection.SaveHistory(ctx, migrFileName)

	res, err := mm.executeCommand(ctx, upCmd)
	if err != nil {
		mm.HistoryCollection.DeleteHistory(ctx, migrFileName)
		return nil, err
	}

	return res, nil
}

func (mm *MongoMigration) executeDown(ctx context.Context, downCmd bson.M, migrFileName string) (interface{}, error) {
	res, err := mm.executeCommand(ctx, downCmd)
	if err != nil {
		return nil, err
	}

	mm.HistoryCollection.DeleteHistory(ctx, migrFileName)

	return res, nil
}

func (mm *MongoMigration) executeCommand(ctx context.Context, cmd bson.M) (interface{}, error) {
	res := mm.Driver.DB.RunCommand(ctx, cmd)

	return res, res.Err()
}
