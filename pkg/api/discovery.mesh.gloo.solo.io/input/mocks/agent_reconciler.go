// Code generated by MockGen. DO NOT EDIT.
// Source: ./agent_reconciler.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	reflect "reflect"

	v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v10 "k8s.io/api/apps/v1"
	v11 "k8s.io/api/core/v1"
)

// MockmultiClusterAgentReconciler is a mock of multiClusterAgentReconciler interface
type MockmultiClusterAgentReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockmultiClusterAgentReconcilerMockRecorder
}

// MockmultiClusterAgentReconcilerMockRecorder is the mock recorder for MockmultiClusterAgentReconciler
type MockmultiClusterAgentReconcilerMockRecorder struct {
	mock *MockmultiClusterAgentReconciler
}

// NewMockmultiClusterAgentReconciler creates a new mock instance
func NewMockmultiClusterAgentReconciler(ctrl *gomock.Controller) *MockmultiClusterAgentReconciler {
	mock := &MockmultiClusterAgentReconciler{ctrl: ctrl}
	mock.recorder = &MockmultiClusterAgentReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmultiClusterAgentReconciler) EXPECT() *MockmultiClusterAgentReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileSettings(clusterName string, obj *v1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileSettings(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileSettings), clusterName, obj)
}

// ReconcileMesh mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileMesh(clusterName string, obj *v1beta2.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileMesh(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileMesh), clusterName, obj)
}

// ReconcileConfigMap mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileConfigMap(clusterName string, obj *v11.ConfigMap) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileConfigMap", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileConfigMap indicates an expected call of ReconcileConfigMap
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileConfigMap(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileConfigMap", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileConfigMap), clusterName, obj)
}

// ReconcileService mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileService(clusterName string, obj *v11.Service) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileService indicates an expected call of ReconcileService
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileService", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileService), clusterName, obj)
}

// ReconcilePod mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcilePod(clusterName string, obj *v11.Pod) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePod", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePod indicates an expected call of ReconcilePod
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcilePod(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePod", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcilePod), clusterName, obj)
}

// ReconcileEndpoints mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileEndpoints(clusterName string, obj *v11.Endpoints) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileEndpoints", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileEndpoints indicates an expected call of ReconcileEndpoints
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileEndpoints(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileEndpoints", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileEndpoints), clusterName, obj)
}

// ReconcileNode mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileNode(clusterName string, obj *v11.Node) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNode", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileNode indicates an expected call of ReconcileNode
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileNode(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNode", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileNode), clusterName, obj)
}

// ReconcileDeployment mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileDeployment(clusterName string, obj *v10.Deployment) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDeployment", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDeployment indicates an expected call of ReconcileDeployment
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileDeployment(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDeployment", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileDeployment), clusterName, obj)
}

// ReconcileReplicaSet mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileReplicaSet(clusterName string, obj *v10.ReplicaSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileReplicaSet", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileReplicaSet indicates an expected call of ReconcileReplicaSet
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileReplicaSet(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReplicaSet", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileReplicaSet), clusterName, obj)
}

// ReconcileDaemonSet mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileDaemonSet(clusterName string, obj *v10.DaemonSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDaemonSet", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDaemonSet indicates an expected call of ReconcileDaemonSet
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileDaemonSet(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDaemonSet", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileDaemonSet), clusterName, obj)
}

// ReconcileStatefulSet mocks base method
func (m *MockmultiClusterAgentReconciler) ReconcileStatefulSet(clusterName string, obj *v10.StatefulSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileStatefulSet", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileStatefulSet indicates an expected call of ReconcileStatefulSet
func (mr *MockmultiClusterAgentReconcilerMockRecorder) ReconcileStatefulSet(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileStatefulSet", reflect.TypeOf((*MockmultiClusterAgentReconciler)(nil).ReconcileStatefulSet), clusterName, obj)
}

// MocksingleClusterAgentReconciler is a mock of singleClusterAgentReconciler interface
type MocksingleClusterAgentReconciler struct {
	ctrl     *gomock.Controller
	recorder *MocksingleClusterAgentReconcilerMockRecorder
}

// MocksingleClusterAgentReconcilerMockRecorder is the mock recorder for MocksingleClusterAgentReconciler
type MocksingleClusterAgentReconcilerMockRecorder struct {
	mock *MocksingleClusterAgentReconciler
}

// NewMocksingleClusterAgentReconciler creates a new mock instance
func NewMocksingleClusterAgentReconciler(ctrl *gomock.Controller) *MocksingleClusterAgentReconciler {
	mock := &MocksingleClusterAgentReconciler{ctrl: ctrl}
	mock.recorder = &MocksingleClusterAgentReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocksingleClusterAgentReconciler) EXPECT() *MocksingleClusterAgentReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileSettings(obj *v1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileSettings), obj)
}

// ReconcileMesh mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileMesh(obj *v1beta2.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileMesh), obj)
}

// ReconcileConfigMap mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileConfigMap(obj *v11.ConfigMap) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileConfigMap", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileConfigMap indicates an expected call of ReconcileConfigMap
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileConfigMap(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileConfigMap", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileConfigMap), obj)
}

// ReconcileService mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileService(obj *v11.Service) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileService", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileService indicates an expected call of ReconcileService
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileService", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileService), obj)
}

// ReconcilePod mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcilePod(obj *v11.Pod) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePod", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePod indicates an expected call of ReconcilePod
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcilePod(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePod", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcilePod), obj)
}

// ReconcileEndpoints mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileEndpoints(obj *v11.Endpoints) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileEndpoints", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileEndpoints indicates an expected call of ReconcileEndpoints
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileEndpoints(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileEndpoints", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileEndpoints), obj)
}

// ReconcileNode mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileNode(obj *v11.Node) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNode", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileNode indicates an expected call of ReconcileNode
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileNode(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNode", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileNode), obj)
}

// ReconcileDeployment mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileDeployment(obj *v10.Deployment) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDeployment", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDeployment indicates an expected call of ReconcileDeployment
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDeployment", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileDeployment), obj)
}

// ReconcileReplicaSet mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileReplicaSet(obj *v10.ReplicaSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileReplicaSet", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileReplicaSet indicates an expected call of ReconcileReplicaSet
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileReplicaSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReplicaSet", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileReplicaSet), obj)
}

// ReconcileDaemonSet mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileDaemonSet(obj *v10.DaemonSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDaemonSet", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDaemonSet indicates an expected call of ReconcileDaemonSet
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileDaemonSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDaemonSet", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileDaemonSet), obj)
}

// ReconcileStatefulSet mocks base method
func (m *MocksingleClusterAgentReconciler) ReconcileStatefulSet(obj *v10.StatefulSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileStatefulSet", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileStatefulSet indicates an expected call of ReconcileStatefulSet
func (mr *MocksingleClusterAgentReconcilerMockRecorder) ReconcileStatefulSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileStatefulSet", reflect.TypeOf((*MocksingleClusterAgentReconciler)(nil).ReconcileStatefulSet), obj)
}
