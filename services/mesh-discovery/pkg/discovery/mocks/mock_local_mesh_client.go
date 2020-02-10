// Code generated by MockGen. DO NOT EDIT.
// Source: ./local_mesh_client.go

// Package mock_discovery is a generated GoMock package.
package mock_discovery

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockLocalMeshClient is a mock of LocalMeshClient interface
type MockLocalMeshClient struct {
	ctrl     *gomock.Controller
	recorder *MockLocalMeshClientMockRecorder
}

// MockLocalMeshClientMockRecorder is the mock recorder for MockLocalMeshClient
type MockLocalMeshClientMockRecorder struct {
	mock *MockLocalMeshClient
}

// NewMockLocalMeshClient creates a new mock instance
func NewMockLocalMeshClient(ctrl *gomock.Controller) *MockLocalMeshClient {
	mock := &MockLocalMeshClient{ctrl: ctrl}
	mock.recorder = &MockLocalMeshClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalMeshClient) EXPECT() *MockLocalMeshClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockLocalMeshClient) Create(ctx context.Context, mesh *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockLocalMeshClientMockRecorder) Create(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLocalMeshClient)(nil).Create), ctx, mesh)
}

// Delete mocks base method
func (m *MockLocalMeshClient) Delete(ctx context.Context, mesh *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockLocalMeshClientMockRecorder) Delete(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLocalMeshClient)(nil).Delete), ctx, mesh)
}

// Get mocks base method
func (m *MockLocalMeshClient) Get(ctx context.Context, objKey client.ObjectKey) (*v1alpha1.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objKey)
	ret0, _ := ret[0].(*v1alpha1.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockLocalMeshClientMockRecorder) Get(ctx, objKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLocalMeshClient)(nil).Get), ctx, objKey)
}

// List mocks base method
func (m *MockLocalMeshClient) List(ctx context.Context, opts ...client.ListOption) (*v1alpha1.MeshList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1alpha1.MeshList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockLocalMeshClientMockRecorder) List(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLocalMeshClient)(nil).List), varargs...)
}
