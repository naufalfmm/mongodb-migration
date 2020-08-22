package driver

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	clientMocks "github.com/naufalfmm/mongodb-migration/mocks/client"
	configMocks "github.com/naufalfmm/mongodb-migration/mocks/config"
)

func ItShouldCreateDriver(t *testing.T) {
	mongoURI := "mongodb://user:password@localhost:1111/db-coba"
	dbName := "db-coba"
	emptyCtx := context.TODO()
	emptyDB := mongo.Database{}

	mockConf := &configMocks.DatabaseConfig{}
	mockClient := &clientMocks.Client{}

	mockConf.On("DBURI").Return(&mongoURI).Twice()
	mockConf.On("DBTimeout").Return(nil).Once()
	mockConf.On("DBName").Return(dbName).Once()

	mockClient.On("ApplyURI", mongoURI).Once()
	mockClient.On("NewClient").Return(nil).Once()
	mockClient.On("Connect", emptyCtx).Return(nil).Once()
	mockClient.On("Ping", emptyCtx, readpref.Primary()).Return(nil).Once()
	mockClient.On("Database", dbName).Return(&emptyDB).Once()

	md := MongoDriver{
		Client: mockClient,
	}
	err := md.SetClientWithContext(emptyCtx, mockConf)

	mockConf.AssertExpectations(t)
	mockConf.AssertNotCalled(t, "SetURI")
	mockClient.AssertExpectations(t)

	mdShould := MongoDriver{
		Client: mockClient,
		db:     &emptyDB,
		Config: mockConf,
	}

	if err != nil {
		t.Errorf("It should not return error")
	}

	if !reflect.DeepEqual(md, mdShould) {
		t.Errorf("It should create %#v, but create %#v", mdShould, md)
	}
}

func ItShouldCalledSetURI(t *testing.T) {
	mongoURI := "mongodb://user:password@localhost:1111/db-coba"
	dbName := "db-coba"
	emptyCtx := context.TODO()
	emptyDB := mongo.Database{}

	mockConf := &configMocks.DatabaseConfig{}
	mockClient := &clientMocks.Client{}

	mockConf.On("DBURI").Return(nil).Once()
	mockConf.On("SetURI").Return(nil).Once()
	mockConf.On("DBURI").Return(&mongoURI).Once()
	mockConf.On("DBTimeout").Return(nil).Once()
	mockConf.On("DBName").Return(dbName).Once()

	mockClient.On("ApplyURI", mongoURI).Once()
	mockClient.On("NewClient").Return(nil).Once()
	mockClient.On("Connect", emptyCtx).Return(nil).Once()
	mockClient.On("Ping", emptyCtx, readpref.Primary()).Return(nil).Once()
	mockClient.On("Database", dbName).Return(&emptyDB).Once()

	md := MongoDriver{
		Client: mockClient,
	}
	err := md.SetClientWithContext(emptyCtx, mockConf)

	mockConf.AssertExpectations(t)
	mockClient.AssertExpectations(t)

	mdShould := MongoDriver{
		Client: mockClient,
		db:     &emptyDB,
		Config: mockConf,
	}

	if err != nil {
		t.Errorf("It should not return error")
	}

	if !reflect.DeepEqual(md, mdShould) {
		t.Errorf("It should create %#v, but create %#v", mdShould, md)
	}
}

func TestSetClientWithContext(t *testing.T) {
	t.Run("It should create driver", ItShouldCreateDriver)
	t.Run("It should called SetURI", ItShouldCalledSetURI)
}
