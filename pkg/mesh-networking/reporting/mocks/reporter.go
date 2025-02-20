// Code generated by MockGen. DO NOT EDIT.
// Source: ./reporter.go

// Package mock_reporting is a generated GoMock package.
package mock_reporting

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
)

// MockReporter is a mock of Reporter interface
type MockReporter struct {
	ctrl     *gomock.Controller
	recorder *MockReporterMockRecorder
}

// MockReporterMockRecorder is the mock recorder for MockReporter
type MockReporterMockRecorder struct {
	mock *MockReporter
}

// NewMockReporter creates a new mock instance
func NewMockReporter(ctrl *gomock.Controller) *MockReporter {
	mock := &MockReporter{ctrl: ctrl}
	mock.recorder = &MockReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReporter) EXPECT() *MockReporterMockRecorder {
	return m.recorder
}

// ReportTrafficPolicyToDestination mocks base method
func (m *MockReporter) ReportTrafficPolicyToDestination(destination *v1.Destination, trafficPolicy ezkube.ResourceId, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportTrafficPolicyToDestination", destination, trafficPolicy, err)
}

// ReportTrafficPolicyToDestination indicates an expected call of ReportTrafficPolicyToDestination
func (mr *MockReporterMockRecorder) ReportTrafficPolicyToDestination(destination, trafficPolicy, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportTrafficPolicyToDestination", reflect.TypeOf((*MockReporter)(nil).ReportTrafficPolicyToDestination), destination, trafficPolicy, err)
}

// ReportAccessPolicyToDestination mocks base method
func (m *MockReporter) ReportAccessPolicyToDestination(destination *v1.Destination, accessPolicy ezkube.ResourceId, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportAccessPolicyToDestination", destination, accessPolicy, err)
}

// ReportAccessPolicyToDestination indicates an expected call of ReportAccessPolicyToDestination
func (mr *MockReporterMockRecorder) ReportAccessPolicyToDestination(destination, accessPolicy, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportAccessPolicyToDestination", reflect.TypeOf((*MockReporter)(nil).ReportAccessPolicyToDestination), destination, accessPolicy, err)
}

// ReportVirtualMeshToMesh mocks base method
func (m *MockReporter) ReportVirtualMeshToMesh(mesh *v1.Mesh, virtualMesh ezkube.ResourceId, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportVirtualMeshToMesh", mesh, virtualMesh, err)
}

// ReportVirtualMeshToMesh indicates an expected call of ReportVirtualMeshToMesh
func (mr *MockReporterMockRecorder) ReportVirtualMeshToMesh(mesh, virtualMesh, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportVirtualMeshToMesh", reflect.TypeOf((*MockReporter)(nil).ReportVirtualMeshToMesh), mesh, virtualMesh, err)
}
