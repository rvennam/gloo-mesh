// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockWasmDeploymentEventHandler is a mock of WasmDeploymentEventHandler interface
type MockWasmDeploymentEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentEventHandlerMockRecorder
}

// MockWasmDeploymentEventHandlerMockRecorder is the mock recorder for MockWasmDeploymentEventHandler
type MockWasmDeploymentEventHandlerMockRecorder struct {
	mock *MockWasmDeploymentEventHandler
}

// NewMockWasmDeploymentEventHandler creates a new mock instance
func NewMockWasmDeploymentEventHandler(ctrl *gomock.Controller) *MockWasmDeploymentEventHandler {
	mock := &MockWasmDeploymentEventHandler{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWasmDeploymentEventHandler) EXPECT() *MockWasmDeploymentEventHandlerMockRecorder {
	return m.recorder
}

// CreateWasmDeployment mocks base method
func (m *MockWasmDeploymentEventHandler) CreateWasmDeployment(obj *v1beta1.WasmDeployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWasmDeployment", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWasmDeployment indicates an expected call of CreateWasmDeployment
func (mr *MockWasmDeploymentEventHandlerMockRecorder) CreateWasmDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWasmDeployment", reflect.TypeOf((*MockWasmDeploymentEventHandler)(nil).CreateWasmDeployment), obj)
}

// UpdateWasmDeployment mocks base method
func (m *MockWasmDeploymentEventHandler) UpdateWasmDeployment(old, new *v1beta1.WasmDeployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWasmDeployment", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWasmDeployment indicates an expected call of UpdateWasmDeployment
func (mr *MockWasmDeploymentEventHandlerMockRecorder) UpdateWasmDeployment(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWasmDeployment", reflect.TypeOf((*MockWasmDeploymentEventHandler)(nil).UpdateWasmDeployment), old, new)
}

// DeleteWasmDeployment mocks base method
func (m *MockWasmDeploymentEventHandler) DeleteWasmDeployment(obj *v1beta1.WasmDeployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWasmDeployment", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWasmDeployment indicates an expected call of DeleteWasmDeployment
func (mr *MockWasmDeploymentEventHandlerMockRecorder) DeleteWasmDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWasmDeployment", reflect.TypeOf((*MockWasmDeploymentEventHandler)(nil).DeleteWasmDeployment), obj)
}

// GenericWasmDeployment mocks base method
func (m *MockWasmDeploymentEventHandler) GenericWasmDeployment(obj *v1beta1.WasmDeployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericWasmDeployment", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericWasmDeployment indicates an expected call of GenericWasmDeployment
func (mr *MockWasmDeploymentEventHandlerMockRecorder) GenericWasmDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericWasmDeployment", reflect.TypeOf((*MockWasmDeploymentEventHandler)(nil).GenericWasmDeployment), obj)
}

// MockWasmDeploymentEventWatcher is a mock of WasmDeploymentEventWatcher interface
type MockWasmDeploymentEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentEventWatcherMockRecorder
}

// MockWasmDeploymentEventWatcherMockRecorder is the mock recorder for MockWasmDeploymentEventWatcher
type MockWasmDeploymentEventWatcherMockRecorder struct {
	mock *MockWasmDeploymentEventWatcher
}

// NewMockWasmDeploymentEventWatcher creates a new mock instance
func NewMockWasmDeploymentEventWatcher(ctrl *gomock.Controller) *MockWasmDeploymentEventWatcher {
	mock := &MockWasmDeploymentEventWatcher{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWasmDeploymentEventWatcher) EXPECT() *MockWasmDeploymentEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockWasmDeploymentEventWatcher) AddEventHandler(ctx context.Context, h controller.WasmDeploymentEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockWasmDeploymentEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockWasmDeploymentEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockVirtualDestinationEventHandler is a mock of VirtualDestinationEventHandler interface
type MockVirtualDestinationEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualDestinationEventHandlerMockRecorder
}

// MockVirtualDestinationEventHandlerMockRecorder is the mock recorder for MockVirtualDestinationEventHandler
type MockVirtualDestinationEventHandlerMockRecorder struct {
	mock *MockVirtualDestinationEventHandler
}

// NewMockVirtualDestinationEventHandler creates a new mock instance
func NewMockVirtualDestinationEventHandler(ctrl *gomock.Controller) *MockVirtualDestinationEventHandler {
	mock := &MockVirtualDestinationEventHandler{ctrl: ctrl}
	mock.recorder = &MockVirtualDestinationEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualDestinationEventHandler) EXPECT() *MockVirtualDestinationEventHandlerMockRecorder {
	return m.recorder
}

// CreateVirtualDestination mocks base method
func (m *MockVirtualDestinationEventHandler) CreateVirtualDestination(obj *v1beta1.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualDestination", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVirtualDestination indicates an expected call of CreateVirtualDestination
func (mr *MockVirtualDestinationEventHandlerMockRecorder) CreateVirtualDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).CreateVirtualDestination), obj)
}

// UpdateVirtualDestination mocks base method
func (m *MockVirtualDestinationEventHandler) UpdateVirtualDestination(old, new *v1beta1.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVirtualDestination", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualDestination indicates an expected call of UpdateVirtualDestination
func (mr *MockVirtualDestinationEventHandlerMockRecorder) UpdateVirtualDestination(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).UpdateVirtualDestination), old, new)
}

// DeleteVirtualDestination mocks base method
func (m *MockVirtualDestinationEventHandler) DeleteVirtualDestination(obj *v1beta1.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualDestination", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualDestination indicates an expected call of DeleteVirtualDestination
func (mr *MockVirtualDestinationEventHandlerMockRecorder) DeleteVirtualDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).DeleteVirtualDestination), obj)
}

// GenericVirtualDestination mocks base method
func (m *MockVirtualDestinationEventHandler) GenericVirtualDestination(obj *v1beta1.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericVirtualDestination", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericVirtualDestination indicates an expected call of GenericVirtualDestination
func (mr *MockVirtualDestinationEventHandlerMockRecorder) GenericVirtualDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).GenericVirtualDestination), obj)
}

// MockVirtualDestinationEventWatcher is a mock of VirtualDestinationEventWatcher interface
type MockVirtualDestinationEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualDestinationEventWatcherMockRecorder
}

// MockVirtualDestinationEventWatcherMockRecorder is the mock recorder for MockVirtualDestinationEventWatcher
type MockVirtualDestinationEventWatcherMockRecorder struct {
	mock *MockVirtualDestinationEventWatcher
}

// NewMockVirtualDestinationEventWatcher creates a new mock instance
func NewMockVirtualDestinationEventWatcher(ctrl *gomock.Controller) *MockVirtualDestinationEventWatcher {
	mock := &MockVirtualDestinationEventWatcher{ctrl: ctrl}
	mock.recorder = &MockVirtualDestinationEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualDestinationEventWatcher) EXPECT() *MockVirtualDestinationEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockVirtualDestinationEventWatcher) AddEventHandler(ctx context.Context, h controller.VirtualDestinationEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockVirtualDestinationEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockVirtualDestinationEventWatcher)(nil).AddEventHandler), varargs...)
}
