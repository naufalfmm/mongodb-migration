package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	client  *mongo.Client
	options []*options.ClientOptions
}

func (mc *MongoClient) NewClient() error {
	client, err := mongo.NewClient(mc.options...)
	if err != nil {
		return err
	}

	mc.client = client

	return nil
}

func (mc *MongoClient) ApplyURI(uri string) {
	newOpts := options.Client().ApplyURI(uri)

	mc.options = append(mc.options, newOpts)
}

func (mc *MongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	err := mc.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}

func (mc *MongoClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return mc.client.Database(name, opts...)
}

func (mc *MongoClient) Connect(ctx context.Context) error {
	return mc.client.Connect(ctx)
}

func (mc *MongoClient) Disconnect(ctx context.Context) error {
	return mc.client.Disconnect(ctx)
}
