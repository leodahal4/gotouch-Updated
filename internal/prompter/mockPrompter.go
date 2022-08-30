// Code generated by MockGen. DO NOT EDIT.
// Source: ./prompter.go

// Package prompter is a generated GoMock package.
package prompter

import (
	fmt "fmt"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPrompter is a mock of Prompter interface.
type MockPrompter struct {
	ctrl     *gomock.Controller
	recorder *MockPrompterMockRecorder
}

// MockPrompterMockRecorder is the mock recorder for MockPrompter.
type MockPrompterMockRecorder struct {
	mock *MockPrompter
}

// NewMockPrompter creates a new mock instance.
func NewMockPrompter(ctrl *gomock.Controller) *MockPrompter {
	mock := &MockPrompter{ctrl: ctrl}
	mock.recorder = &MockPrompterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrompter) EXPECT() *MockPrompterMockRecorder {
	return m.recorder
}

// AskForSelectionFromList mocks base method.
func (m *MockPrompter) AskForSelectionFromList(direction string, list []fmt.Stringer) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskForSelectionFromList", direction, list)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskForSelectionFromList indicates an expected call of AskForSelectionFromList.
func (mr *MockPrompterMockRecorder) AskForSelectionFromList(direction, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskForSelectionFromList", reflect.TypeOf((*MockPrompter)(nil).AskForSelectionFromList), direction, list)
}

// AskForString mocks base method.
func (m *MockPrompter) AskForString(direction string, validator StringValidator) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskForString", direction, validator)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskForString indicates an expected call of AskForString.
func (mr *MockPrompterMockRecorder) AskForString(direction, validator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskForString", reflect.TypeOf((*MockPrompter)(nil).AskForString), direction, validator)
}

// AskForYesOrNo mocks base method.
func (m *MockPrompter) AskForYesOrNo(direction string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskForYesOrNo", direction)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskForYesOrNo indicates an expected call of AskForYesOrNo.
func (mr *MockPrompterMockRecorder) AskForYesOrNo(direction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskForYesOrNo", reflect.TypeOf((*MockPrompter)(nil).AskForYesOrNo), direction)
}
