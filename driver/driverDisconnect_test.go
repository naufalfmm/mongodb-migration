package driver

import (
	"context"
	"errors"
	"testing"

	clientMocks "github.com/naufalfmm/mongodb-migration/mocks/client"
)

func ItShouldCallDisconnect(t *testing.T) {
	emptyCtx := context.TODO()

	mockClient := &clientMocks.Client{}

	mockClient.On("Disconnect", emptyCtx).Return(nil).Once()

	md := MongoDriver{
		Client: mockClient,
	}

	err := md.Disconnect(emptyCtx)

	if err != nil {
		t.Errorf("It should not return error")
	}
}

func ItShouldErrorWhenDisconnectClientError(t *testing.T) {
	emptyCtx := context.TODO()

	mockClient := &clientMocks.Client{}

	mockClient.On("Disconnect", emptyCtx).Return(errors.New("Any Error")).Once()

	md := MongoDriver{
		Client: mockClient,
	}

	err := md.Disconnect(emptyCtx)

	if err == nil {
		t.Errorf("It should return error")
		return
	}

	if err.Error() != "Any Error" {
		t.Errorf("It should return error %+v but get %+v", errors.New("Any Error"), err)

	}
}

func TestDisconnect(t *testing.T) {
	t.Run("It should call Disconnect", ItShouldCallDisconnect)
	t.Run("It should error when Disconnect error", ItShouldErrorWhenDisconnectClientError)
}
