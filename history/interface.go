package history

import (
	"context"

	"github.com/naufalfmm/mongodb-migration/history_data"
)

type MigrationHistory interface {
	InitializeHistory(ctx context.Context, historyCollectionName string) error
	DropHistory(ctx context.Context) error
	SaveHistory(ctx context.Context, migrationData history_data.MigrationHistoryData) error
	DeleteHistory(ctx context.Context, migrationData history_data.MigrationHistoryData) error
	GetHistory(ctx context.Context, migrationName string) (*history_data.MigrationHistoryData, error)
	GetAllHistories(ctx context.Context) (*[]history_data.MigrationHistoryData, error)
	GetLatestHistory(ctx context.Context) (*history_data.MigrationHistoryData, error)
}
