/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go

// Package mock_tags is a generated GoMock package.
package mock_tags

import (
	context "context"
	reflect "reflect"

	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-10-01/resources"
	gomock "go.uber.org/mock/gomock"
)

// Mockclient is a mock of client interface.
type Mockclient struct {
	ctrl     *gomock.Controller
	recorder *MockclientMockRecorder
}

// MockclientMockRecorder is the mock recorder for Mockclient.
type MockclientMockRecorder struct {
	mock *Mockclient
}

// NewMockclient creates a new mock instance.
func NewMockclient(ctrl *gomock.Controller) *Mockclient {
	mock := &Mockclient{ctrl: ctrl}
	mock.recorder = &MockclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockclient) EXPECT() *MockclientMockRecorder {
	return m.recorder
}

// GetAtScope mocks base method.
func (m *Mockclient) GetAtScope(arg0 context.Context, arg1 string) (resources.TagsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAtScope", arg0, arg1)
	ret0, _ := ret[0].(resources.TagsResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAtScope indicates an expected call of GetAtScope.
func (mr *MockclientMockRecorder) GetAtScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAtScope", reflect.TypeOf((*Mockclient)(nil).GetAtScope), arg0, arg1)
}

// UpdateAtScope mocks base method.
func (m *Mockclient) UpdateAtScope(arg0 context.Context, arg1 string, arg2 resources.TagsPatchResource) (resources.TagsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAtScope", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.TagsResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAtScope indicates an expected call of UpdateAtScope.
func (mr *MockclientMockRecorder) UpdateAtScope(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAtScope", reflect.TypeOf((*Mockclient)(nil).UpdateAtScope), arg0, arg1, arg2)
}
