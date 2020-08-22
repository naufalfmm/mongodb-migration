package driver

import (
	"context"
	"errors"
	"reflect"
	"testing"

	clientMocks "github.com/naufalfmm/mongodb-migration/mocks/client"
	configMocks "github.com/naufalfmm/mongodb-migration/mocks/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ItShouldCallSetClientWithContext(t *testing.T) {
	mongoURI := "mongodb://user:password@localhost:1111/db-coba"
	dbName := "db-coba"
	emptyCtx := context.TODO()
	emptyDB := mongo.Database{}

	// mock of SetClientWithContext
	// --- start line ---
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
	// --- end line ---

	md := MongoDriver{
		Client: mockClient,
	}
	err := md.SetClient(mockConf)

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

func ItShouldErrorWhenSetClientWithContextError(t *testing.T) {
	mongoURI := "mongodb://user:password@localhost:1111/db-coba"
	dbName := "db-coba"
	emptyCtx := context.TODO()

	// mock of SetClientWithContext
	// --- start line ---
	mockConf := &configMocks.DatabaseConfig{}
	mockClient := &clientMocks.Client{}

	mockConf.On("DBURI").Return(&mongoURI).Twice()

	mockClient.On("ApplyURI", mongoURI).Once()
	mockClient.On("NewClient").Return(errors.New("Any Error")).Once()
	// --- end line ---

	md := MongoDriver{
		Client: mockClient,
	}
	err := md.SetClient(mockConf)

	mockConf.AssertExpectations(t)
	mockConf.AssertNotCalled(t, "SetURI")
	mockConf.AssertNotCalled(t, "DBTimeout")
	mockConf.AssertNotCalled(t, "DBName")
	mockClient.AssertExpectations(t)
	mockClient.AssertNotCalled(t, "Connect", emptyCtx)
	mockClient.AssertNotCalled(t, "Ping", emptyCtx, readpref.Primary())
	mockClient.AssertNotCalled(t, "Database", dbName)

	if err == nil {
		t.Errorf("It should return error")
		return
	}

	if err.Error() != "Any Error" {
		t.Errorf("It should return error %+v but get %+v", errors.New("Any Error"), err)
	}
}

func TestSetClient(t *testing.T) {
	t.Run("It should call SetClientWithContext", ItShouldCallSetClientWithContext)
	t.Run("It should error when SetClientWithContext error", ItShouldErrorWhenSetClientWithContextError)
}
