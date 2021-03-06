package driver

import (
	"context"
	"time"

	"github.com/naufalfmm/mongodb-migration/client"
	"github.com/naufalfmm/mongodb-migration/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDriver struct {
	Client client.Client
	db     *mongo.Database

	Config config.DatabaseConfig
}

func (md *MongoDriver) GetDB() *mongo.Database {
	return md.db
}

func (md *MongoDriver) SetClient(cfg config.DatabaseConfig) error {
	ctx := context.TODO()

	err := md.SetClientWithContext(ctx, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (md *MongoDriver) SetClientWithContext(ctx context.Context, cfg config.DatabaseConfig) error {
	var (
		cancel context.CancelFunc
	)

	if cfg.DBURI() == nil {
		cfg.SetURI()
	}

	md.Client.ApplyURI(*cfg.DBURI())

	err := md.Client.NewClient()
	if err != nil {
		return err
	}

	if cfg.DBTimeout() != nil {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(*cfg.DBTimeout())*time.Second)
		defer cancel()
	}

	err = md.Client.Connect(ctx)
	if err != nil {
		return err
	}

	err = md.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db := md.Client.Database(cfg.DBName())

	md.db = db

	md.Config = cfg

	return nil
}

func (md *MongoDriver) Connect(ctx context.Context) error {
	return md.Client.Connect(ctx)
}

func (md *MongoDriver) Disconnect(ctx context.Context) error {
	return md.Client.Disconnect(ctx)
}
