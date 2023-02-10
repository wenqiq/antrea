// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/multicluster/controllers/multicluster/commonarea (interfaces: CommonArea,RemoteCommonArea,ImportReconciler,RemoteCommonAreaGetter)

// Package commonarea is a generated GoMock package.
package commonarea

import (
	v1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	common "antrea.io/antrea/multicluster/controllers/multicluster/common"
	context "context"
	gomock "github.com/golang/mock/gomock"
	meta "k8s.io/apimachinery/pkg/api/meta"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	manager "sigs.k8s.io/controller-runtime/pkg/manager"
)

// MockCommonArea is a mock of CommonArea interface
type MockCommonArea struct {
	ctrl     *gomock.Controller
	recorder *MockCommonAreaMockRecorder
}

// MockCommonAreaMockRecorder is the mock recorder for MockCommonArea
type MockCommonAreaMockRecorder struct {
	mock *MockCommonArea
}

// NewMockCommonArea creates a new mock instance
func NewMockCommonArea(ctrl *gomock.Controller) *MockCommonArea {
	mock := &MockCommonArea{ctrl: ctrl}
	mock.recorder = &MockCommonAreaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommonArea) EXPECT() *MockCommonAreaMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCommonArea) Create(arg0 context.Context, arg1 client.Object, arg2 ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCommonAreaMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommonArea)(nil).Create), varargs...)
}

// Delete mocks base method
func (m *MockCommonArea) Delete(arg0 context.Context, arg1 client.Object, arg2 ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCommonAreaMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommonArea)(nil).Delete), varargs...)
}

// DeleteAllOf mocks base method
func (m *MockCommonArea) DeleteAllOf(arg0 context.Context, arg1 client.Object, arg2 ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOf", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOf indicates an expected call of DeleteAllOf
func (mr *MockCommonAreaMockRecorder) DeleteAllOf(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOf", reflect.TypeOf((*MockCommonArea)(nil).DeleteAllOf), varargs...)
}

// Get mocks base method
func (m *MockCommonArea) Get(arg0 context.Context, arg1 types.NamespacedName, arg2 client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockCommonAreaMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCommonArea)(nil).Get), arg0, arg1, arg2)
}

// GetClusterID mocks base method
func (m *MockCommonArea) GetClusterID() common.ClusterID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterID")
	ret0, _ := ret[0].(common.ClusterID)
	return ret0
}

// GetClusterID indicates an expected call of GetClusterID
func (mr *MockCommonAreaMockRecorder) GetClusterID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterID", reflect.TypeOf((*MockCommonArea)(nil).GetClusterID))
}

// GetNamespace mocks base method
func (m *MockCommonArea) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockCommonAreaMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockCommonArea)(nil).GetNamespace))
}

// List mocks base method
func (m *MockCommonArea) List(arg0 context.Context, arg1 client.ObjectList, arg2 ...client.ListOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockCommonAreaMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCommonArea)(nil).List), varargs...)
}

// Patch mocks base method
func (m *MockCommonArea) Patch(arg0 context.Context, arg1 client.Object, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch
func (mr *MockCommonAreaMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockCommonArea)(nil).Patch), varargs...)
}

// RESTMapper mocks base method
func (m *MockCommonArea) RESTMapper() meta.RESTMapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	return ret0
}

// RESTMapper indicates an expected call of RESTMapper
func (mr *MockCommonAreaMockRecorder) RESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTMapper", reflect.TypeOf((*MockCommonArea)(nil).RESTMapper))
}

// Scheme mocks base method
func (m *MockCommonArea) Scheme() *runtime.Scheme {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheme")
	ret0, _ := ret[0].(*runtime.Scheme)
	return ret0
}

// Scheme indicates an expected call of Scheme
func (mr *MockCommonAreaMockRecorder) Scheme() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheme", reflect.TypeOf((*MockCommonArea)(nil).Scheme))
}

// Status mocks base method
func (m *MockCommonArea) Status() client.StatusWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(client.StatusWriter)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockCommonAreaMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCommonArea)(nil).Status))
}

// Update mocks base method
func (m *MockCommonArea) Update(arg0 context.Context, arg1 client.Object, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCommonAreaMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommonArea)(nil).Update), varargs...)
}

// MockRemoteCommonArea is a mock of RemoteCommonArea interface
type MockRemoteCommonArea struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteCommonAreaMockRecorder
}

// MockRemoteCommonAreaMockRecorder is the mock recorder for MockRemoteCommonArea
type MockRemoteCommonAreaMockRecorder struct {
	mock *MockRemoteCommonArea
}

// NewMockRemoteCommonArea creates a new mock instance
func NewMockRemoteCommonArea(ctrl *gomock.Controller) *MockRemoteCommonArea {
	mock := &MockRemoteCommonArea{ctrl: ctrl}
	mock.recorder = &MockRemoteCommonAreaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteCommonArea) EXPECT() *MockRemoteCommonAreaMockRecorder {
	return m.recorder
}

// AddImportReconciler mocks base method
func (m *MockRemoteCommonArea) AddImportReconciler(arg0 ImportReconciler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddImportReconciler", arg0)
}

// AddImportReconciler indicates an expected call of AddImportReconciler
func (mr *MockRemoteCommonAreaMockRecorder) AddImportReconciler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImportReconciler", reflect.TypeOf((*MockRemoteCommonArea)(nil).AddImportReconciler), arg0)
}

// Create mocks base method
func (m *MockRemoteCommonArea) Create(arg0 context.Context, arg1 client.Object, arg2 ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRemoteCommonAreaMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRemoteCommonArea)(nil).Create), varargs...)
}

// Delete mocks base method
func (m *MockRemoteCommonArea) Delete(arg0 context.Context, arg1 client.Object, arg2 ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRemoteCommonAreaMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRemoteCommonArea)(nil).Delete), varargs...)
}

// DeleteAllOf mocks base method
func (m *MockRemoteCommonArea) DeleteAllOf(arg0 context.Context, arg1 client.Object, arg2 ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOf", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOf indicates an expected call of DeleteAllOf
func (mr *MockRemoteCommonAreaMockRecorder) DeleteAllOf(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOf", reflect.TypeOf((*MockRemoteCommonArea)(nil).DeleteAllOf), varargs...)
}

// Get mocks base method
func (m *MockRemoteCommonArea) Get(arg0 context.Context, arg1 types.NamespacedName, arg2 client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockRemoteCommonAreaMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRemoteCommonArea)(nil).Get), arg0, arg1, arg2)
}

// GetClusterID mocks base method
func (m *MockRemoteCommonArea) GetClusterID() common.ClusterID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterID")
	ret0, _ := ret[0].(common.ClusterID)
	return ret0
}

// GetClusterID indicates an expected call of GetClusterID
func (mr *MockRemoteCommonAreaMockRecorder) GetClusterID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterID", reflect.TypeOf((*MockRemoteCommonArea)(nil).GetClusterID))
}

// GetLocalClusterID mocks base method
func (m *MockRemoteCommonArea) GetLocalClusterID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalClusterID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLocalClusterID indicates an expected call of GetLocalClusterID
func (mr *MockRemoteCommonAreaMockRecorder) GetLocalClusterID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalClusterID", reflect.TypeOf((*MockRemoteCommonArea)(nil).GetLocalClusterID))
}

// GetNamespace mocks base method
func (m *MockRemoteCommonArea) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockRemoteCommonAreaMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockRemoteCommonArea)(nil).GetNamespace))
}

// GetStatus mocks base method
func (m *MockRemoteCommonArea) GetStatus() []v1alpha1.ClusterCondition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].([]v1alpha1.ClusterCondition)
	return ret0
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockRemoteCommonAreaMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockRemoteCommonArea)(nil).GetStatus))
}

// IsConnected mocks base method
func (m *MockRemoteCommonArea) IsConnected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected
func (mr *MockRemoteCommonAreaMockRecorder) IsConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockRemoteCommonArea)(nil).IsConnected))
}

// List mocks base method
func (m *MockRemoteCommonArea) List(arg0 context.Context, arg1 client.ObjectList, arg2 ...client.ListOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockRemoteCommonAreaMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRemoteCommonArea)(nil).List), varargs...)
}

// Patch mocks base method
func (m *MockRemoteCommonArea) Patch(arg0 context.Context, arg1 client.Object, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch
func (mr *MockRemoteCommonAreaMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRemoteCommonArea)(nil).Patch), varargs...)
}

// RESTMapper mocks base method
func (m *MockRemoteCommonArea) RESTMapper() meta.RESTMapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	return ret0
}

// RESTMapper indicates an expected call of RESTMapper
func (mr *MockRemoteCommonAreaMockRecorder) RESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTMapper", reflect.TypeOf((*MockRemoteCommonArea)(nil).RESTMapper))
}

// Scheme mocks base method
func (m *MockRemoteCommonArea) Scheme() *runtime.Scheme {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheme")
	ret0, _ := ret[0].(*runtime.Scheme)
	return ret0
}

// Scheme indicates an expected call of Scheme
func (mr *MockRemoteCommonAreaMockRecorder) Scheme() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheme", reflect.TypeOf((*MockRemoteCommonArea)(nil).Scheme))
}

// Start mocks base method
func (m *MockRemoteCommonArea) Start() context.CancelFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockRemoteCommonAreaMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRemoteCommonArea)(nil).Start))
}

// StartWatching mocks base method
func (m *MockRemoteCommonArea) StartWatching() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWatching")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWatching indicates an expected call of StartWatching
func (mr *MockRemoteCommonAreaMockRecorder) StartWatching() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWatching", reflect.TypeOf((*MockRemoteCommonArea)(nil).StartWatching))
}

// Status mocks base method
func (m *MockRemoteCommonArea) Status() client.StatusWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(client.StatusWriter)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockRemoteCommonAreaMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockRemoteCommonArea)(nil).Status))
}

// Stop mocks base method
func (m *MockRemoteCommonArea) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockRemoteCommonAreaMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRemoteCommonArea)(nil).Stop))
}

// StopWatching mocks base method
func (m *MockRemoteCommonArea) StopWatching() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopWatching")
}

// StopWatching indicates an expected call of StopWatching
func (mr *MockRemoteCommonAreaMockRecorder) StopWatching() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopWatching", reflect.TypeOf((*MockRemoteCommonArea)(nil).StopWatching))
}

// Update mocks base method
func (m *MockRemoteCommonArea) Update(arg0 context.Context, arg1 client.Object, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRemoteCommonAreaMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRemoteCommonArea)(nil).Update), varargs...)
}

// MockImportReconciler is a mock of ImportReconciler interface
type MockImportReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockImportReconcilerMockRecorder
}

// MockImportReconcilerMockRecorder is the mock recorder for MockImportReconciler
type MockImportReconcilerMockRecorder struct {
	mock *MockImportReconciler
}

// NewMockImportReconciler creates a new mock instance
func NewMockImportReconciler(ctrl *gomock.Controller) *MockImportReconciler {
	mock := &MockImportReconciler{ctrl: ctrl}
	mock.recorder = &MockImportReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImportReconciler) EXPECT() *MockImportReconcilerMockRecorder {
	return m.recorder
}

// SetupWithManager mocks base method
func (m *MockImportReconciler) SetupWithManager(arg0 manager.Manager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupWithManager", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupWithManager indicates an expected call of SetupWithManager
func (mr *MockImportReconcilerMockRecorder) SetupWithManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupWithManager", reflect.TypeOf((*MockImportReconciler)(nil).SetupWithManager), arg0)
}

// MockRemoteCommonAreaGetter is a mock of RemoteCommonAreaGetter interface
type MockRemoteCommonAreaGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteCommonAreaGetterMockRecorder
}

// MockRemoteCommonAreaGetterMockRecorder is the mock recorder for MockRemoteCommonAreaGetter
type MockRemoteCommonAreaGetterMockRecorder struct {
	mock *MockRemoteCommonAreaGetter
}

// NewMockRemoteCommonAreaGetter creates a new mock instance
func NewMockRemoteCommonAreaGetter(ctrl *gomock.Controller) *MockRemoteCommonAreaGetter {
	mock := &MockRemoteCommonAreaGetter{ctrl: ctrl}
	mock.recorder = &MockRemoteCommonAreaGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteCommonAreaGetter) EXPECT() *MockRemoteCommonAreaGetterMockRecorder {
	return m.recorder
}

// GetRemoteCommonAreaAndLocalID mocks base method
func (m *MockRemoteCommonAreaGetter) GetRemoteCommonAreaAndLocalID() (RemoteCommonArea, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteCommonAreaAndLocalID")
	ret0, _ := ret[0].(RemoteCommonArea)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRemoteCommonAreaAndLocalID indicates an expected call of GetRemoteCommonAreaAndLocalID
func (mr *MockRemoteCommonAreaGetterMockRecorder) GetRemoteCommonAreaAndLocalID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteCommonAreaAndLocalID", reflect.TypeOf((*MockRemoteCommonAreaGetter)(nil).GetRemoteCommonAreaAndLocalID))
}
