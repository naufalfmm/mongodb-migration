package migration

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/naufalfmm/mongodb-migration/config"
	"github.com/naufalfmm/mongodb-migration/constants"
	"github.com/naufalfmm/mongodb-migration/driver"
	"github.com/naufalfmm/mongodb-migration/history"
	"github.com/naufalfmm/mongodb-migration/history_data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMigration struct {
	Source string
	Driver driver.Driver

	HistoryCollection history.MigrationHistory
}

type JSONCommand struct {
	UpCommand   bson.M `bson:"up"`
	DownCommand bson.M `bson:"down"`
}

func (mm *MongoMigration) StartMigration(source string, client *mongo.Client, cfg config.DatabaseConfig, historyCollection history.MigrationHistory) error {
	var mongoDriver driver.Driver

	mongoDriver.SetClient(cfg)

	mm.StartMigrationWithDriver(source, mongoDriver, historyCollection)

	return nil
}

func (mm *MongoMigration) StartMigrationWithDriver(source string, driver driver.Driver, historyCollection history.MigrationHistory) error {
	mm.Driver = driver
	mm.Source = source

	ctx := context.TODO()

	err := historyCollection.InitializeHistory(ctx)
	if err != nil {
		cmdErr := err.(mongo.CommandError)
		if cmdErr.Code != 48 {
			return cmdErr
		}
	}

	mm.HistoryCollection = historyCollection

	return nil
}

func (mm *MongoMigration) runUp(ctx context.Context, files []os.FileInfo, steps int) error {
	step := 0
	fileExt := ".json"

	for _, info := range files {
		if strings.HasSuffix(info.Name(), fileExt) {
			migrName := strings.TrimRight(info.Name(), fileExt)
			hd, err := mm.HistoryCollection.GetHistory(ctx, migrName)
			if err != nil {
				return err
			}

			if hd != nil {
				continue
			}

			_, err = mm.runSpecificFile(ctx, migrName, constants.UP)
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

func (mm *MongoMigration) runDown(ctx context.Context, files []os.FileInfo, steps int) error {
	step := 0
	fileExt := ".json"

	for _, info := range files {
		if strings.HasSuffix(info.Name(), fileExt) {
			migrName := strings.TrimRight(info.Name(), fileExt)
			hd, err := mm.HistoryCollection.GetHistory(ctx, migrName)
			if err != nil {
				return err
			}

			if hd == nil {
				continue
			}

			_, err = mm.runSpecificFile(ctx, migrName, constants.DOWN)
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

func (mm *MongoMigration) Run(ctx context.Context, direction int, steps int) error {
	var err error = nil

	dir := http.Dir(mm.Source)
	file, err := dir.Open("/")
	if err != nil {
		panic(err)
	}

	files, err := file.Readdir(0)
	if err != nil {
		panic(err)
	}

	directionRun := mm.runUp
	if direction == constants.DOWN {
		directionRun = mm.runDown
	}

	err = directionRun(ctx, files, steps)

	return err
}

func (mm *MongoMigration) runSpecificFile(ctx context.Context, migrationName string, direction int) (interface{}, error) {
	var command JSONCommand

	migrationFile, err := os.Open(mm.Source + migrationName + ".json")
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

	directionFunc := mm.executeUp
	usedCommand := command.UpCommand
	if direction == constants.DOWN {
		directionFunc = mm.executeDown
		usedCommand = command.DownCommand
	}

	res, err := directionFunc(ctx, usedCommand, migrationName)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (mm *MongoMigration) RunSpecificFile(ctx context.Context, migrationFileName string, direction int) (interface{}, error) {
	fileExt := ".json"
	migrName := strings.TrimRight(migrationFileName, fileExt)

	res, err := mm.runSpecificFile(ctx, migrName, direction)

	return res, err
}

func (mm *MongoMigration) executeUp(ctx context.Context, upCmd bson.M, migrName string) (interface{}, error) {
	migrRecord := history_data.MigrationRecordData{
		MigrationName: migrName,
		CreatedAt:     time.Now(),
	}

	mm.HistoryCollection.SaveHistory(ctx, &migrRecord)

	res, err := mm.executeCommand(ctx, upCmd)
	if err != nil {
		mm.HistoryCollection.DeleteHistory(ctx, &migrRecord)
		return nil, err
	}

	return res, nil
}

func (mm *MongoMigration) executeDown(ctx context.Context, downCmd bson.M, migrName string) (interface{}, error) {
	migrRecord := history_data.MigrationRecordData{
		MigrationName: migrName,
		CreatedAt:     time.Now(),
	}

	res, err := mm.executeCommand(ctx, downCmd)
	if err != nil {
		return nil, err
	}

	mm.HistoryCollection.DeleteHistory(ctx, &migrRecord)

	return res, nil
}

func (mm *MongoMigration) executeCommand(ctx context.Context, cmd bson.M) (interface{}, error) {
	res := mm.Driver.GetDB().RunCommand(ctx, cmd)

	return res, res.Err()
}
