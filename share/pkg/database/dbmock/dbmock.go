// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/database.go

// Package dbmock is a generated GoMock package.
package dbmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockDataBase is a mock of DataBase interface.
type MockDataBase struct {
	ctrl     *gomock.Controller
	recorder *MockDataBaseMockRecorder
}

// MockDataBaseMockRecorder is the mock recorder for MockDataBase.
type MockDataBaseMockRecorder struct {
	mock *MockDataBase
}

// NewMockDataBase creates a new mock instance.
func NewMockDataBase(ctrl *gomock.Controller) *MockDataBase {
	mock := &MockDataBase{ctrl: ctrl}
	mock.recorder = &MockDataBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBase) EXPECT() *MockDataBaseMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDataBase) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDataBaseMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDataBase)(nil).Close))
}

// Transaction mocks base method.
func (m *MockDataBase) Transaction(fn func(*sqlx.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockDataBaseMockRecorder) Transaction(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockDataBase)(nil).Transaction), fn)
}
