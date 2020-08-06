package history

import (
	"context"
)

type MigrationHistory interface {
	InitializeHistory(ctx context.Context) error
	DropHistory(ctx context.Context) error
	SaveHistory(ctx context.Context, migrationData interface{}) error
	DeleteHistory(ctx context.Context, migrationData interface{}) error
	GetHistory(ctx context.Context, migrationName string) (interface{}, error)
	GetLatestHistory(ctx context.Context) (interface{}, error)
}
