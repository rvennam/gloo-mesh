// Code generated by MockGen. DO NOT EDIT.
// Source: ./sidecar_detector.go

// Package mock_detector is a generated GoMock package.
package mock_detector

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	v1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	v10 "k8s.io/api/core/v1"
)

// MockSidecarDetector is a mock of SidecarDetector interface
type MockSidecarDetector struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarDetectorMockRecorder
}

// MockSidecarDetectorMockRecorder is the mock recorder for MockSidecarDetector
type MockSidecarDetectorMockRecorder struct {
	mock *MockSidecarDetector
}

// NewMockSidecarDetector creates a new mock instance
func NewMockSidecarDetector(ctrl *gomock.Controller) *MockSidecarDetector {
	mock := &MockSidecarDetector{ctrl: ctrl}
	mock.recorder = &MockSidecarDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSidecarDetector) EXPECT() *MockSidecarDetectorMockRecorder {
	return m.recorder
}

// DetectMeshSidecar mocks base method
func (m *MockSidecarDetector) DetectMeshSidecar(pod *v10.Pod, meshes v1sets.MeshSet) *v1.Mesh {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectMeshSidecar", pod, meshes)
	ret0, _ := ret[0].(*v1.Mesh)
	return ret0
}

// DetectMeshSidecar indicates an expected call of DetectMeshSidecar
func (mr *MockSidecarDetectorMockRecorder) DetectMeshSidecar(pod, meshes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectMeshSidecar", reflect.TypeOf((*MockSidecarDetector)(nil).DetectMeshSidecar), pod, meshes)
}
