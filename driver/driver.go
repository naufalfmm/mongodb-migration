package driver

import (
	"context"
	"log"
	"time"

	"github.com/naufalfmm/mongodb-migration/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDriver struct {
	Client *mongo.Client
	DB     *mongo.Database

	Config *config.MongoConfig
}

func (md *MongoDriver) SetClient(cfg config.MongoConfig) error {
	ctx := context.Background()

	err := md.SetClientWithContext(ctx, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (md *MongoDriver) SetClientWithContext(ctx context.Context, cfg config.MongoConfig) error {
	var cancel context.CancelFunc

	if cfg.DBURI == nil {
		cfg.SetURL()
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(*cfg.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	md.Client = client

	if cfg.DBTimeout != nil {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(*cfg.DBTimeout)*time.Second)
		defer cancel()
	}

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db := client.Database(cfg.DBName)

	md.DB = db

	md.Config = &cfg

	return nil
}
