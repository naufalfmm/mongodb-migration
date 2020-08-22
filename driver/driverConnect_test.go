package driver

import (
	"context"
	"errors"
	"testing"

	clientMocks "github.com/naufalfmm/mongodb-migration/mocks/client"
)

func ItShouldCallConnect(t *testing.T) {
	emptyCtx := context.TODO()

	mockClient := &clientMocks.Client{}

	mockClient.On("Connect", emptyCtx).Return(nil).Once()

	md := MongoDriver{
		Client: mockClient,
	}

	err := md.Connect(emptyCtx)

	if err != nil {
		t.Errorf("It should not return error")
	}
}

func ItShouldErrorWhenConnectClientError(t *testing.T) {
	emptyCtx := context.TODO()

	mockClient := &clientMocks.Client{}

	mockClient.On("Connect", emptyCtx).Return(errors.New("Any Error")).Once()

	md := MongoDriver{
		Client: mockClient,
	}

	err := md.Connect(emptyCtx)

	if err == nil {
		t.Errorf("It should return error")
		return
	}

	if err.Error() != "Any Error" {
		t.Errorf("It should return error %+v but get %+v", errors.New("Any Error"), err)

	}
}

func TestConnect(t *testing.T) {
	t.Run("It should call Connect", ItShouldCallConnect)
	t.Run("It should error when Connect error", ItShouldErrorWhenConnectClientError)
}
