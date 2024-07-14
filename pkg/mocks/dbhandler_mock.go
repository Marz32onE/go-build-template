// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marz32one/go-build-template/pkg/storage (interfaces: DBHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDBHandler is a mock of DBHandler interface.
type MockDBHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDBHandlerMockRecorder
}

// MockDBHandlerMockRecorder is the mock recorder for MockDBHandler.
type MockDBHandlerMockRecorder struct {
	mock *MockDBHandler
}

// NewMockDBHandler creates a new mock instance.
func NewMockDBHandler(ctrl *gomock.Controller) *MockDBHandler {
	mock := &MockDBHandler{ctrl: ctrl}
	mock.recorder = &MockDBHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBHandler) EXPECT() *MockDBHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDBHandler) Create(arg0 interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDBHandlerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDBHandler)(nil).Create), arg0)
}

// Find mocks base method.
func (m *MockDBHandler) Find(arg0 interface{}, arg1 ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockDBHandlerMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDBHandler)(nil).Find), varargs...)
}
