// Code generated by MockGen. DO NOT EDIT.
// Source: overlap.go

// Package overlap is a generated GoMock package.
package overlap

import (
	reflect "reflect"

	gomock "github.com/neutrinocorp/mock/gomock"
)

// MockReadWriteCloser is a mock of ReadWriteCloser interface.
type MockReadWriteCloser struct {
	ctrl     *gomock.Controller
	recorder *MockReadWriteCloserMockRecorder
}

// MockReadWriteCloserMockRecorder is the mock recorder for MockReadWriteCloser.
type MockReadWriteCloserMockRecorder struct {
	mock *MockReadWriteCloser
}

// NewMockReadWriteCloser creates a new mock instance.
func NewMockReadWriteCloser(ctrl *gomock.Controller) *MockReadWriteCloser {
	mock := &MockReadWriteCloser{ctrl: ctrl}
	mock.recorder = &MockReadWriteCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadWriteCloser) EXPECT() *MockReadWriteCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReadWriteCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReadWriteCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadWriteCloser)(nil).Close))
}

// Read mocks base method.
func (m *MockReadWriteCloser) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockReadWriteCloserMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadWriteCloser)(nil).Read), arg0)
}

// Write mocks base method.
func (m *MockReadWriteCloser) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockReadWriteCloserMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockReadWriteCloser)(nil).Write), arg0)
}
