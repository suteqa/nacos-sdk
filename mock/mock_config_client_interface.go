// Code generated by MockGen. DO NOT EDIT.
// Source: clients/config_client/config_client_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	vo "github.com/suteqa/nacos-sdk/vo"
	reflect "reflect"
)

// MockIConfigClient is a mock of IConfigClient interface
type MockIConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockIConfigClientMockRecorder
}

// MockIConfigClientMockRecorder is the mock recorder for MockIConfigClient
type MockIConfigClientMockRecorder struct {
	mock *MockIConfigClient
}

// NewMockIConfigClient creates a new mock instance
func NewMockIConfigClient(ctrl *gomock.Controller) *MockIConfigClient {
	mock := &MockIConfigClient{ctrl: ctrl}
	mock.recorder = &MockIConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConfigClient) EXPECT() *MockIConfigClientMockRecorder {
	return m.recorder
}

// GetConfig mocks base method
func (m *MockIConfigClient) GetConfig(param vo.ConfigParam) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", param)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockIConfigClientMockRecorder) GetConfig(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockIConfigClient)(nil).GetConfig), param)
}

// PublishConfig mocks base method
func (m *MockIConfigClient) PublishConfig(param vo.ConfigParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishConfig", param)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishConfig indicates an expected call of PublishConfig
func (mr *MockIConfigClientMockRecorder) PublishConfig(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishConfig", reflect.TypeOf((*MockIConfigClient)(nil).PublishConfig), param)
}

// DeleteConfig mocks base method
func (m *MockIConfigClient) DeleteConfig(param vo.ConfigParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConfig", param)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteConfig indicates an expected call of DeleteConfig
func (mr *MockIConfigClientMockRecorder) DeleteConfig(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfig", reflect.TypeOf((*MockIConfigClient)(nil).DeleteConfig), param)
}

// ListenConfig mocks base method
func (m *MockIConfigClient) ListenConfig(params []vo.ConfigParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenConfig", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenConfig indicates an expected call of ListenConfig
func (mr *MockIConfigClientMockRecorder) ListenConfig(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenConfig", reflect.TypeOf((*MockIConfigClient)(nil).ListenConfig), params)
}

// AddConfigToListen mocks base method
func (m *MockIConfigClient) AddConfigToListen(params []vo.ConfigParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddConfigToListen", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddConfigToListen indicates an expected call of AddConfigToListen
func (mr *MockIConfigClientMockRecorder) AddConfigToListen(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConfigToListen", reflect.TypeOf((*MockIConfigClient)(nil).AddConfigToListen), params)
}

// StopListenConfig mocks base method
func (m *MockIConfigClient) StopListenConfig() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopListenConfig")
}

// StopListenConfig indicates an expected call of StopListenConfig
func (mr *MockIConfigClientMockRecorder) StopListenConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopListenConfig", reflect.TypeOf((*MockIConfigClient)(nil).StopListenConfig))
}

// GetConfigContent mocks base method
func (m *MockIConfigClient) GetConfigContent(dataId, groupId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigContent", dataId, groupId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigContent indicates an expected call of GetConfigContent
func (mr *MockIConfigClientMockRecorder) GetConfigContent(dataId, groupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigContent", reflect.TypeOf((*MockIConfigClient)(nil).GetConfigContent), dataId, groupId)
}
