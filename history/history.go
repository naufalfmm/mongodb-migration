package history

import (
	"context"
	"errors"

	"github.com/naufalfmm/mongodb-migration/history_data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MigrationRecord struct {
	DB *mongo.Database

	CollectionName string
}

func (mh *MigrationRecord) InitializeHistory(ctx context.Context) error {
	err := mh.DB.CreateCollection(ctx, mh.CollectionName)

	return err
}

func (mh *MigrationRecord) DropHistory(ctx context.Context) error {
	err := mh.DB.Collection(mh.CollectionName).Drop(ctx)

	return err
}

func (mh *MigrationRecord) SaveHistory(ctx context.Context, migrationData interface{}) error {
	migrHistoryCollection := mh.DB.Collection(mh.CollectionName)

	_, err := migrHistoryCollection.InsertOne(ctx, migrationData, options.InsertOne())

	return err
}

func (mh *MigrationRecord) DeleteHistory(ctx context.Context, migrationData interface{}) error {
	migrHistoryCollection := mh.DB.Collection(mh.CollectionName)

	filter := bson.M{"migrationName": migrationData.(*history_data.MigrationRecordData).MigrationName}

	_, err := migrHistoryCollection.DeleteOne(ctx, filter)

	return err
}

func (mh *MigrationRecord) GetHistory(ctx context.Context, migrationName string) (interface{}, error) {
	var migrationHistory history_data.MigrationRecordData

	migrHistoryCollection := mh.DB.Collection(mh.CollectionName)

	filter := bson.M{"migrationName": migrationName}

	cur := migrHistoryCollection.FindOne(ctx, filter)
	if cur.Err() != nil {
		if errors.Is(cur.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, cur.Err()
	}

	err := cur.Decode(&migrationHistory)
	if err != nil {
		return nil, err
	}

	return &migrationHistory, nil
}

func (mh *MigrationRecord) GetAllHistories(ctx context.Context) (*[]history_data.MigrationRecordData, error) {
	var migrationHistories []history_data.MigrationRecordData

	migrHistoryCollection := mh.DB.Collection(mh.CollectionName)

	cur, err := migrHistoryCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cur.Decode(&migrationHistories)
	if err != nil {
		return nil, err
	}

	return &migrationHistories, nil
}

func (mh *MigrationRecord) GetLatestHistory(ctx context.Context) (interface{}, error) {
	var latestMigrHistory history_data.MigrationRecordData

	migrHistoryCollection := mh.DB.Collection(mh.CollectionName)

	sort := options.FindOne().SetSort(bson.M{"createdAt": -1})

	cur := migrHistoryCollection.FindOne(ctx, nil, sort)
	if cur.Err() != nil {
		return nil, cur.Err()
	}

	err := cur.Decode(&latestMigrHistory)
	if err != nil {
		return nil, err
	}

	return &latestMigrHistory, nil
}
