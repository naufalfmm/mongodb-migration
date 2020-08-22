// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	config "github.com/naufalfmm/mongodb-migration/config"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// Connect provides a mock function with given fields: ctx
func (_m *Driver) Connect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields: ctx
func (_m *Driver) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDB provides a mock function with given fields:
func (_m *Driver) GetDB() *mongo.Database {
	ret := _m.Called()

	var r0 *mongo.Database
	if rf, ok := ret.Get(0).(func() *mongo.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Database)
		}
	}

	return r0
}

// SetClient provides a mock function with given fields: cfg
func (_m *Driver) SetClient(cfg config.DatabaseConfig) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(config.DatabaseConfig) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetClientWithContext provides a mock function with given fields: ctx, cfg
func (_m *Driver) SetClientWithContext(ctx context.Context, cfg config.DatabaseConfig) error {
	ret := _m.Called(ctx, cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.DatabaseConfig) error); ok {
		r0 = rf(ctx, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
