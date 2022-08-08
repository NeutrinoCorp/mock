// Code generated by MockGen. DO NOT EDIT.
// Source: external.go

// Package source is a generated GoMock package.
package source

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	generics "github.com/golang/mock/mockgen/internal/tests/generics"
	other "github.com/golang/mock/mockgen/internal/tests/generics/other"
	constraints "golang.org/x/exp/constraints"
)

// MockExternalConstraint is a mock of ExternalConstraint interface.
type MockExternalConstraint[I constraints.Integer, F constraints.Float] struct {
	ctrl     *gomock.Controller
	recorder *MockExternalConstraintMockRecorder[I, F]
}

// MockExternalConstraintMockRecorder is the mock recorder for MockExternalConstraint.
type MockExternalConstraintMockRecorder[I constraints.Integer, F constraints.Float] struct {
	mock *MockExternalConstraint[I, F]
}

// NewMockExternalConstraint creates a new mock instance.
func NewMockExternalConstraint[I constraints.Integer, F constraints.Float](ctrl *gomock.Controller) *MockExternalConstraint[I, F] {
	mock := &MockExternalConstraint[I, F]{ctrl: ctrl}
	mock.recorder = &MockExternalConstraintMockRecorder[I, F]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalConstraint[I, F]) EXPECT() *MockExternalConstraintMockRecorder[I, F] {
	return m.recorder
}

// Eight mocks base method.
func (m *MockExternalConstraint[I, F]) Eight(arg0 F) other.Two[I, F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[I, F])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockExternalConstraintMockRecorder[I, F]) Eight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Eight), arg0)
}

// Five mocks base method.
func (m *MockExternalConstraint[I, F]) Five(arg0 I) generics.Baz[F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(generics.Baz[F])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockExternalConstraintMockRecorder[I, F]) Five(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Five), arg0)
}

// Four mocks base method.
func (m *MockExternalConstraint[I, F]) Four(arg0 I) generics.Foo[I, F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(generics.Foo[I, F])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockExternalConstraintMockRecorder[I, F]) Four(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Four), arg0)
}

// Nine mocks base method.
func (m *MockExternalConstraint[I, F]) Nine(arg0 generics.Iface[I]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockExternalConstraintMockRecorder[I, F]) Nine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Nine), arg0)
}

// One mocks base method.
func (m *MockExternalConstraint[I, F]) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockExternalConstraintMockRecorder[I, F]) One(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).One), arg0)
}

// Seven mocks base method.
func (m *MockExternalConstraint[I, F]) Seven(arg0 I) other.One[I] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[I])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockExternalConstraintMockRecorder[I, F]) Seven(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Seven), arg0)
}

// Six mocks base method.
func (m *MockExternalConstraint[I, F]) Six(arg0 I) *generics.Baz[F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*generics.Baz[F])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockExternalConstraintMockRecorder[I, F]) Six(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Six), arg0)
}

// Ten mocks base method.
func (m *MockExternalConstraint[I, F]) Ten(arg0 *I) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockExternalConstraintMockRecorder[I, F]) Ten(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Ten), arg0)
}

// Three mocks base method.
func (m *MockExternalConstraint[I, F]) Three(arg0 I) F {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(F)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockExternalConstraintMockRecorder[I, F]) Three(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Three), arg0)
}

// Two mocks base method.
func (m *MockExternalConstraint[I, F]) Two(arg0 I) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockExternalConstraintMockRecorder[I, F]) Two(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Two), arg0)
}

// MockTwentyTwo is a mock of TwentyTwo interface.
type MockTwentyTwo[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockTwentyTwoMockRecorder[T]
}

// MockTwentyTwoMockRecorder is the mock recorder for MockTwentyTwo.
type MockTwentyTwoMockRecorder[T any] struct {
	mock *MockTwentyTwo[T]
}

// NewMockTwentyTwo creates a new mock instance.
func NewMockTwentyTwo[T any](ctrl *gomock.Controller) *MockTwentyTwo[T] {
	mock := &MockTwentyTwo[T]{ctrl: ctrl}
	mock.recorder = &MockTwentyTwoMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTwentyTwo[T]) EXPECT() *MockTwentyTwoMockRecorder[T] {
	return m.recorder
}

// TwentyTwo mocks base method.
func (m *MockTwentyTwo[T]) TwentyTwo() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyTwo")
	ret0, _ := ret[0].(T)
	return ret0
}

// TwentyTwo indicates an expected call of TwentyTwo.
func (mr *MockTwentyTwoMockRecorder[T]) TwentyTwo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyTwo", reflect.TypeOf((*MockTwentyTwo[T])(nil).TwentyTwo))
}
