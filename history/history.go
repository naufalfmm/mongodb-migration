package history

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MigrationHistory struct {
	DB *mongo.Database

	CollectionName string
}

type MigrationHistoryModel struct {
	MigrationName string    `bson:"migrationName"`
	CreatedAt     time.Time `bson:"createdAt"`
}

func (mh *MigrationHistory) InitializeHistoryCollection(ctx context.Context) error {
	err := mh.DB.CreateCollection(ctx, mh.CollectionName)
	if err != nil {
		return err
	}

	return nil
}

func (mh *MigrationHistory) SaveHistory(ctx context.Context, migrName string) error {
	migrHistoryCollection := mh.DB.Collection(mh.CollectionName)

	migHist := MigrationHistoryModel{
		MigrationName: migrName,
		CreatedAt:     time.Now(),
	}

	_, err := migrHistoryCollection.InsertOne(ctx, migHist)
	if err != nil {
		return err
	}

	return nil
}

func (mh *MigrationHistory) DeleteHistory(ctx context.Context, migrName string) error {
	migrHistoryCollection := mh.DB.Collection(mh.CollectionName)

	filter := bson.M{"migrationName": migrName}

	_, err := migrHistoryCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (mh *MigrationHistory) GetLatestMigration(ctx context.Context) (*MigrationHistory, error) {
	var latestMigrHistory MigrationHistory

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
