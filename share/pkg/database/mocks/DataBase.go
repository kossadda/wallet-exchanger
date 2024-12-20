// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	sqlx "github.com/jmoiron/sqlx"
	mock "github.com/stretchr/testify/mock"
)

// DataBase is an autogenerated mock type for the DataBase type
type DataBase struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *DataBase) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transaction provides a mock function with given fields: fn
func (_m *DataBase) Transaction(fn func(*sqlx.Tx) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*sqlx.Tx) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDataBase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataBase creates a new instance of DataBase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataBase(t mockConstructorTestingTNewDataBase) *DataBase {
	mock := &DataBase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
