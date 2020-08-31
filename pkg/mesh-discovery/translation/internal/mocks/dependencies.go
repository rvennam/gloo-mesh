// Code generated by MockGen. DO NOT EDIT.
// Source: ./dependencies.go

// Package mock_translation is a generated GoMock package.
package mock_translation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	input "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/input"
	mesh "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/mesh"
	traffictarget "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/traffictarget"
	workload "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/workload"
)

// MockDependencyFactory is a mock of DependencyFactory interface
type MockDependencyFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDependencyFactoryMockRecorder
}

// MockDependencyFactoryMockRecorder is the mock recorder for MockDependencyFactory
type MockDependencyFactoryMockRecorder struct {
	mock *MockDependencyFactory
}

// NewMockDependencyFactory creates a new mock instance
func NewMockDependencyFactory(ctrl *gomock.Controller) *MockDependencyFactory {
	mock := &MockDependencyFactory{ctrl: ctrl}
	mock.recorder = &MockDependencyFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDependencyFactory) EXPECT() *MockDependencyFactoryMockRecorder {
	return m.recorder
}

// MakeMeshTranslator mocks base method
func (m *MockDependencyFactory) MakeMeshTranslator(ctx context.Context, in input.Snapshot) mesh.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeMeshTranslator", ctx, in)
	ret0, _ := ret[0].(mesh.Translator)
	return ret0
}

// MakeMeshTranslator indicates an expected call of MakeMeshTranslator
func (mr *MockDependencyFactoryMockRecorder) MakeMeshTranslator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeMeshTranslator", reflect.TypeOf((*MockDependencyFactory)(nil).MakeMeshTranslator), ctx, in)
}

// MakeWorkloadTranslator mocks base method
func (m *MockDependencyFactory) MakeWorkloadTranslator(ctx context.Context, in input.Snapshot) workload.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeWorkloadTranslator", ctx, in)
	ret0, _ := ret[0].(workload.Translator)
	return ret0
}

// MakeWorkloadTranslator indicates an expected call of MakeWorkloadTranslator
func (mr *MockDependencyFactoryMockRecorder) MakeWorkloadTranslator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeWorkloadTranslator", reflect.TypeOf((*MockDependencyFactory)(nil).MakeWorkloadTranslator), ctx, in)
}

// MakeTrafficTargetTranslator mocks base method
func (m *MockDependencyFactory) MakeTrafficTargetTranslator(ctx context.Context) traffictarget.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeTrafficTargetTranslator", ctx)
	ret0, _ := ret[0].(traffictarget.Translator)
	return ret0
}

// MakeTrafficTargetTranslator indicates an expected call of MakeTrafficTargetTranslator
func (mr *MockDependencyFactoryMockRecorder) MakeTrafficTargetTranslator(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeTrafficTargetTranslator", reflect.TypeOf((*MockDependencyFactory)(nil).MakeTrafficTargetTranslator), ctx)
}
