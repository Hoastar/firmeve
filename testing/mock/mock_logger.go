// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/firmeve/firmeve/kernel/contract (interfaces: Loggable)

// Package mock is a generated GoMock package.
package mock

import (
	contract "github.com/firmeve/firmeve/kernel/contract"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockLoggable is a mock of Loggable interface
type MockLoggable struct {
	ctrl     *gomock.Controller
	recorder *MockLoggableMockRecorder
}

// MockLoggableMockRecorder is the mock recorder for MockLoggable
type MockLoggableMockRecorder struct {
	mock *MockLoggable
}

// NewMockLoggable creates a new mock instance
func NewMockLoggable(ctrl *gomock.Controller) *MockLoggable {
	mock := &MockLoggable{ctrl: ctrl}
	mock.recorder = &MockLoggableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoggable) EXPECT() *MockLoggableMockRecorder {
	return m.recorder
}

// Debug mocks base method
func (m *MockLoggable) Debug(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockLoggableMockRecorder) Debug(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLoggable)(nil).Debug), varargs...)
}

// Error mocks base method
func (m *MockLoggable) Error(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockLoggableMockRecorder) Error(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLoggable)(nil).Error), varargs...)
}

// Fatal mocks base method
func (m *MockLoggable) Fatal(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockLoggableMockRecorder) Fatal(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLoggable)(nil).Fatal), varargs...)
}

// Info mocks base method
func (m *MockLoggable) Info(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockLoggableMockRecorder) Info(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLoggable)(nil).Info), varargs...)
}

// Panic mocks base method
func (m *MockLoggable) Panic(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panic", varargs...)
}

// Panic indicates an expected call of Panic
func (mr *MockLoggableMockRecorder) Panic(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panic", reflect.TypeOf((*MockLoggable)(nil).Panic), varargs...)
}

// Warn mocks base method
func (m *MockLoggable) Warn(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn
func (mr *MockLoggableMockRecorder) Warn(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLoggable)(nil).Warn), varargs...)
}

// With mocks base method
func (m *MockLoggable) With(arg0 ...interface{}) contract.Loggable {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "With", varargs...)
	ret0, _ := ret[0].(contract.Loggable)
	return ret0
}

// With indicates an expected call of With
func (mr *MockLoggableMockRecorder) With(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockLoggable)(nil).With), arg0...)
}

// Writer mocks base method
func (m *MockLoggable) Writer(arg0 string) io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer", arg0)
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// Writer indicates an expected call of Writer
func (mr *MockLoggableMockRecorder) Writer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockLoggable)(nil).Writer), arg0)
}