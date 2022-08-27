// Code generated by MockGen. DO NOT EDIT.
// Source: ./manager.go

// Package manager is a generated GoMock package.
package manager

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CreateDirectoryIfNotExists mocks base method.
func (m *MockManager) CreateDirectoryIfNotExists(directoryName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDirectoryIfNotExists", directoryName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDirectoryIfNotExists indicates an expected call of CreateDirectoryIfNotExists.
func (mr *MockManagerMockRecorder) CreateDirectoryIfNotExists(directoryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDirectoryIfNotExists", reflect.TypeOf((*MockManager)(nil).CreateDirectoryIfNotExists), directoryName)
}

// EditGoModule mocks base method.
func (m *MockManager) EditGoModule(projectName, folderName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditGoModule", projectName, folderName)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditGoModule indicates an expected call of EditGoModule.
func (mr *MockManagerMockRecorder) EditGoModule(projectName, folderName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditGoModule", reflect.TypeOf((*MockManager)(nil).EditGoModule), projectName, folderName)
}

// GetExtractLocation mocks base method.
func (m *MockManager) GetExtractLocation() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtractLocation")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetExtractLocation indicates an expected call of GetExtractLocation.
func (mr *MockManagerMockRecorder) GetExtractLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtractLocation", reflect.TypeOf((*MockManager)(nil).GetExtractLocation))
}

// GetStream mocks base method.
func (m *MockManager) GetStream() io.ReadCloser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStream")
	ret0, _ := ret[0].(io.ReadCloser)
	return ret0
}

// GetStream indicates an expected call of GetStream.
func (mr *MockManagerMockRecorder) GetStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStream", reflect.TypeOf((*MockManager)(nil).GetStream))
}

// GetWd mocks base method.
func (m *MockManager) GetWd() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWd")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWd indicates an expected call of GetWd.
func (mr *MockManagerMockRecorder) GetWd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWd", reflect.TypeOf((*MockManager)(nil).GetWd))
}

// IsTest mocks base method.
func (m *MockManager) IsTest() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTest")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTest indicates an expected call of IsTest.
func (mr *MockManagerMockRecorder) IsTest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTest", reflect.TypeOf((*MockManager)(nil).IsTest))
}
