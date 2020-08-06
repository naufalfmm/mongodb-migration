package driver

import (
	"context"

	"github.com/naufalfmm/mongodb-migration/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Driver interface {
	GetDB() *mongo.Database
	SetClient(cfg config.DatabaseConfig) error
	SetClientWithContext(ctx context.Context, cfg config.DatabaseConfig) error
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
