// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/discovery/v1alpha2/workload.proto

package v1alpha2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *WorkloadSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadSpec)
	if !ok {
		that2, ok := that.(WorkloadSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMesh()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMesh()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMesh(), target.GetMesh()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAppMesh()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAppMesh()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAppMesh(), target.GetAppMesh()) {
			return false
		}
	}

	switch m.WorkloadType.(type) {

	case *WorkloadSpec_Kubernetes:
		if _, ok := target.WorkloadType.(*WorkloadSpec_Kubernetes); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubernetes()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubernetes()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubernetes(), target.GetKubernetes()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.WorkloadType != target.WorkloadType {
			return false
		}
	}

	return true
}

// Equal function
func (m *WorkloadStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadStatus)
	if !ok {
		that2, ok := that.(WorkloadStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if len(m.GetAppliedAccessLogRecords()) != len(target.GetAppliedAccessLogRecords()) {
		return false
	}
	for idx, v := range m.GetAppliedAccessLogRecords() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedAccessLogRecords()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedAccessLogRecords()[idx]) {
				return false
			}
		}

	}

	if len(m.GetAppliedWasmDeployments()) != len(target.GetAppliedWasmDeployments()) {
		return false
	}
	for idx, v := range m.GetAppliedWasmDeployments() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedWasmDeployments()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedWasmDeployments()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkloadSpec_KubernetesWorkload) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadSpec_KubernetesWorkload)
	if !ok {
		that2, ok := that.(WorkloadSpec_KubernetesWorkload)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetController()).(equality.Equalizer); ok {
		if !h.Equal(target.GetController()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetController(), target.GetController()) {
			return false
		}
	}

	if len(m.GetPodLabels()) != len(target.GetPodLabels()) {
		return false
	}
	for k, v := range m.GetPodLabels() {

		if strings.Compare(v, target.GetPodLabels()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetServiceAccountName(), target.GetServiceAccountName()) != 0 {
		return false
	}

	if len(m.GetEndpoints()) != len(target.GetEndpoints()) {
		return false
	}
	for idx, v := range m.GetEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEndpoints()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkloadSpec_AppMesh) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadSpec_AppMesh)
	if !ok {
		that2, ok := that.(WorkloadSpec_AppMesh)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetVirtualNodeName(), target.GetVirtualNodeName()) != 0 {
		return false
	}

	if len(m.GetPorts()) != len(target.GetPorts()) {
		return false
	}
	for idx, v := range m.GetPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPorts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkloadSpec_KubernetesWorkload_EndpointsSubset) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadSpec_KubernetesWorkload_EndpointsSubset)
	if !ok {
		that2, ok := that.(WorkloadSpec_KubernetesWorkload_EndpointsSubset)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetIpAddresses()) != len(target.GetIpAddresses()) {
		return false
	}
	for idx, v := range m.GetIpAddresses() {

		if strings.Compare(v, target.GetIpAddresses()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetPorts()) != len(target.GetPorts()) {
		return false
	}
	for idx, v := range m.GetPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPorts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkloadSpec_AppMesh_ContainerPort) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadSpec_AppMesh_ContainerPort)
	if !ok {
		that2, ok := that.(WorkloadSpec_AppMesh_ContainerPort)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	if strings.Compare(m.GetProtocol(), target.GetProtocol()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *WorkloadStatus_AppliedAccessLogRecord) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadStatus_AppliedAccessLogRecord)
	if !ok {
		that2, ok := that.(WorkloadStatus_AppliedAccessLogRecord)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *WorkloadStatus_AppliedWasmDeployment) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadStatus_AppliedWasmDeployment)
	if !ok {
		that2, ok := that.(WorkloadStatus_AppliedWasmDeployment)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}
