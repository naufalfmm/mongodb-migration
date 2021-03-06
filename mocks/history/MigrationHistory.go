// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MigrationHistory is an autogenerated mock type for the MigrationHistory type
type MigrationHistory struct {
	mock.Mock
}

// DeleteHistory provides a mock function with given fields: ctx, migrationData
func (_m *MigrationHistory) DeleteHistory(ctx context.Context, migrationData interface{}) error {
	ret := _m.Called(ctx, migrationData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, migrationData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropHistory provides a mock function with given fields: ctx
func (_m *MigrationHistory) DropHistory(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHistory provides a mock function with given fields: ctx, migrationName
func (_m *MigrationHistory) GetHistory(ctx context.Context, migrationName string) (interface{}, error) {
	ret := _m.Called(ctx, migrationName)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string) interface{}); ok {
		r0 = rf(ctx, migrationName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, migrationName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestHistory provides a mock function with given fields: ctx
func (_m *MigrationHistory) GetLatestHistory(ctx context.Context) (interface{}, error) {
	ret := _m.Called(ctx)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context) interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitializeHistory provides a mock function with given fields: ctx
func (_m *MigrationHistory) InitializeHistory(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveHistory provides a mock function with given fields: ctx, migrationData
func (_m *MigrationHistory) SaveHistory(ctx context.Context, migrationData interface{}) error {
	ret := _m.Called(ctx, migrationData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, migrationData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
