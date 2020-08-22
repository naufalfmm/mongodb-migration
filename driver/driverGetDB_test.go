package driver

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func ItShouldReturnDBValue(t *testing.T) {
	emptyDB := mongo.Database{}

	md := MongoDriver{
		db: &emptyDB,
	}

	val := md.GetDB()

	if !reflect.DeepEqual(val, md.db) {
		t.Errorf("It should return %+v but get %+v", md.db, val)
	}
}

func TestGetDB(t *testing.T) {
	t.Run("It should return DB value", ItShouldReturnDBValue)
}
