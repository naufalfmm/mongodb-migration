package history_data

import "time"

type MigrationHistoryData interface {
	GetID() interface{}
	SetMigrationName(name string) error
	GetMigrationName() string
	SetCreatedAt(createdAt time.Time) error
	SetNowAsCreatedAt() error
	GetCreatedAt() time.Time
}
