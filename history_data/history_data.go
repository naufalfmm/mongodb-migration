package history_data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MigrationRecordData struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	MigrationName string             `bson:"migrationName"`
	CreatedAt     time.Time          `bson:"createdAt"`
}
