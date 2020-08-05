package history_data

import "time"

type MigrationRecordData struct {
	ID            string    `bson:"_id"`
	MigrationName string    `bson:"migrationName"`
	CreatedAt     time.Time `bson:"createdAt"`
}

func (mrd *MigrationRecordData) GetID() interface{} {
	return mrd.ID
}

func (mrd *MigrationRecordData) SetMigrationName(name string) error {
	mrd.MigrationName = name

	return nil
}

func (mrd *MigrationRecordData) GetMigrationName() string {
	return mrd.MigrationName
}

func (mrd *MigrationRecordData) SetCreatedAt(createdAt time.Time) error {
	mrd.CreatedAt = createdAt

	return nil
}

func (mrd *MigrationRecordData) SetNowAsCreatedAt() error {
	now := time.Now()
	mrd.CreatedAt = now

	return nil
}

func (mrd *MigrationRecordData) GetCreatedAt() time.Time {
	return mrd.CreatedAt
}
