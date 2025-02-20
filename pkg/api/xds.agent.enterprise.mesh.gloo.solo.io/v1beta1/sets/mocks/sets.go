// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1beta1sets is a generated GoMock package.
package mock_v1beta1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1"
	v1beta1sets "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockXdsConfigSet is a mock of XdsConfigSet interface
type MockXdsConfigSet struct {
	ctrl     *gomock.Controller
	recorder *MockXdsConfigSetMockRecorder
}

// MockXdsConfigSetMockRecorder is the mock recorder for MockXdsConfigSet
type MockXdsConfigSetMockRecorder struct {
	mock *MockXdsConfigSet
}

// NewMockXdsConfigSet creates a new mock instance
func NewMockXdsConfigSet(ctrl *gomock.Controller) *MockXdsConfigSet {
	mock := &MockXdsConfigSet{ctrl: ctrl}
	mock.recorder = &MockXdsConfigSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockXdsConfigSet) EXPECT() *MockXdsConfigSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockXdsConfigSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockXdsConfigSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockXdsConfigSet)(nil).Keys))
}

// List mocks base method
func (m *MockXdsConfigSet) List(filterResource ...func(*v1beta1.XdsConfig) bool) []*v1beta1.XdsConfig {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1beta1.XdsConfig)
	return ret0
}

// List indicates an expected call of List
func (mr *MockXdsConfigSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockXdsConfigSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockXdsConfigSet) Map() map[string]*v1beta1.XdsConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1beta1.XdsConfig)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockXdsConfigSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockXdsConfigSet)(nil).Map))
}

// Insert mocks base method
func (m *MockXdsConfigSet) Insert(xdsConfig ...*v1beta1.XdsConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range xdsConfig {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockXdsConfigSetMockRecorder) Insert(xdsConfig ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockXdsConfigSet)(nil).Insert), xdsConfig...)
}

// Equal mocks base method
func (m *MockXdsConfigSet) Equal(xdsConfigSet v1beta1sets.XdsConfigSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", xdsConfigSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockXdsConfigSetMockRecorder) Equal(xdsConfigSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockXdsConfigSet)(nil).Equal), xdsConfigSet)
}

// Has mocks base method
func (m *MockXdsConfigSet) Has(xdsConfig ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", xdsConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockXdsConfigSetMockRecorder) Has(xdsConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockXdsConfigSet)(nil).Has), xdsConfig)
}

// Delete mocks base method
func (m *MockXdsConfigSet) Delete(xdsConfig ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", xdsConfig)
}

// Delete indicates an expected call of Delete
func (mr *MockXdsConfigSetMockRecorder) Delete(xdsConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockXdsConfigSet)(nil).Delete), xdsConfig)
}

// Union mocks base method
func (m *MockXdsConfigSet) Union(set v1beta1sets.XdsConfigSet) v1beta1sets.XdsConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1beta1sets.XdsConfigSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockXdsConfigSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockXdsConfigSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockXdsConfigSet) Difference(set v1beta1sets.XdsConfigSet) v1beta1sets.XdsConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1beta1sets.XdsConfigSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockXdsConfigSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockXdsConfigSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockXdsConfigSet) Intersection(set v1beta1sets.XdsConfigSet) v1beta1sets.XdsConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1beta1sets.XdsConfigSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockXdsConfigSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockXdsConfigSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockXdsConfigSet) Find(id ezkube.ResourceId) (*v1beta1.XdsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1beta1.XdsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockXdsConfigSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockXdsConfigSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockXdsConfigSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockXdsConfigSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockXdsConfigSet)(nil).Length))
}

// Generic mocks base method
func (m *MockXdsConfigSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockXdsConfigSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockXdsConfigSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockXdsConfigSet) Delta(newSet v1beta1sets.XdsConfigSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockXdsConfigSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockXdsConfigSet)(nil).Delta), newSet)
}
