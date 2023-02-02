/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export_service

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ExportServiceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExportServiceSpec{}

// ExportServiceSpec struct for ExportServiceSpec
type ExportServiceSpec struct {
	// Specify the ip to be export
	Ip *string `json:"ip,omitempty"`
	// Specify the name of the export destination
	Name *string `json:"name,omitempty"`
	// Specify the namespace of the export destination
	Namespace *string `json:"namespace,omitempty"`
	// Specify the port to be used in service
	Port *int32 `json:"port,omitempty"`
	// Specify the port to be export
	TargetPort *int32 `json:"targetPort,omitempty"`
	// Specify the topology to export
	Topology *string `json:"topology,omitempty"`
}

// NewExportServiceSpecWith instantiates a new ExportServiceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportServiceSpecWith() *ExportServiceSpec {
	this := ExportServiceSpec{}
	return &this
}

// NewExportServiceSpec instantiates a new ExportServiceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportServiceSpec() *ExportServiceSpec {
	this := ExportServiceSpec{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ExportServiceWorkflowStep) GetIp() string {
	if o == nil || utils.IsNil(o.Properties.Ip) {
		var ret string
		return ret
	}
	return *o.Properties.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportServiceWorkflowStep) GetIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Ip) {
		return nil, false
	}
	return o.Properties.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ExportServiceWorkflowStep) HasIp() bool {
	if o != nil && !utils.IsNil(o.Properties.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the ip field.
// Ip:  Specify the ip to be export
func (o *ExportServiceWorkflowStep) SetIp(v string) *ExportServiceWorkflowStep {
	o.Properties.Ip = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExportServiceWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportServiceWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExportServiceWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of the export destination
func (o *ExportServiceWorkflowStep) SetName(v string) *ExportServiceWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ExportServiceWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportServiceWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ExportServiceWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace of the export destination
func (o *ExportServiceWorkflowStep) SetNamespace(v string) *ExportServiceWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ExportServiceWorkflowStep) GetPort() int32 {
	if o == nil || utils.IsNil(o.Properties.Port) {
		var ret int32
		return ret
	}
	return *o.Properties.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportServiceWorkflowStep) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Port) {
		return nil, false
	}
	return o.Properties.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ExportServiceWorkflowStep) HasPort() bool {
	if o != nil && !utils.IsNil(o.Properties.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the port field.
// Port:  Specify the port to be used in service
func (o *ExportServiceWorkflowStep) SetPort(v int32) *ExportServiceWorkflowStep {
	o.Properties.Port = &v
	return o
}

// GetTargetPort returns the TargetPort field value if set, zero value otherwise.
func (o *ExportServiceWorkflowStep) GetTargetPort() int32 {
	if o == nil || utils.IsNil(o.Properties.TargetPort) {
		var ret int32
		return ret
	}
	return *o.Properties.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportServiceWorkflowStep) GetTargetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.TargetPort) {
		return nil, false
	}
	return o.Properties.TargetPort, true
}

// HasTargetPort returns a boolean if a field has been set.
func (o *ExportServiceWorkflowStep) HasTargetPort() bool {
	if o != nil && !utils.IsNil(o.Properties.TargetPort) {
		return true
	}

	return false
}

// SetTargetPort gets a reference to the given int32 and assigns it to the targetPort field.
// TargetPort:  Specify the port to be export
func (o *ExportServiceWorkflowStep) SetTargetPort(v int32) *ExportServiceWorkflowStep {
	o.Properties.TargetPort = &v
	return o
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *ExportServiceWorkflowStep) GetTopology() string {
	if o == nil || utils.IsNil(o.Properties.Topology) {
		var ret string
		return ret
	}
	return *o.Properties.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportServiceWorkflowStep) GetTopologyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Topology) {
		return nil, false
	}
	return o.Properties.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *ExportServiceWorkflowStep) HasTopology() bool {
	if o != nil && !utils.IsNil(o.Properties.Topology) {
		return true
	}

	return false
}

// SetTopology gets a reference to the given string and assigns it to the topology field.
// Topology:  Specify the topology to export
func (o *ExportServiceWorkflowStep) SetTopology(v string) *ExportServiceWorkflowStep {
	o.Properties.Topology = &v
	return o
}

func (o ExportServiceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportServiceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.TargetPort) {
		toSerialize["targetPort"] = o.TargetPort
	}
	if !utils.IsNil(o.Topology) {
		toSerialize["topology"] = o.Topology
	}
	return toSerialize, nil
}

type NullableExportServiceSpec struct {
	value *ExportServiceSpec
	isSet bool
}

func (v NullableExportServiceSpec) Get() *ExportServiceSpec {
	return v.value
}

func (v *NullableExportServiceSpec) Set(val *ExportServiceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableExportServiceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableExportServiceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportServiceSpec(val *ExportServiceSpec) *NullableExportServiceSpec {
	return &NullableExportServiceSpec{value: val, isSet: true}
}

func (v NullableExportServiceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportServiceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ExportServiceType = "export-service"

func init() {
	sdkcommon.RegisterWorkflowStep(ExportServiceType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ExportServiceType, FromWorkflowSubStep)
}

type ExportServiceWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ExportServiceSpec
}

func ExportService(name string) *ExportServiceWorkflowStep {
	e := &ExportServiceWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ExportServiceType,
	}}
	return e
}

func (e *ExportServiceWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range e.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  e.Base.DependsOn,
		If:         e.Base.If,
		Inputs:     e.Base.Inputs,
		Meta:       e.Base.Meta,
		Name:       e.Base.Name,
		Outputs:    e.Base.Outputs,
		Properties: util.Object2RawExtension(e.Properties),
		SubSteps:   subSteps,
		Timeout:    e.Base.Timeout,
		Type:       ExportServiceType,
	}
	return res
}

func (e *ExportServiceWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ExportServiceWorkflowStep, error) {
	var properties ExportServiceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := e.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	e.Base.Name = from.Name
	e.Base.DependsOn = from.DependsOn
	e.Base.Inputs = from.Inputs
	e.Base.Outputs = from.Outputs
	e.Base.If = from.If
	e.Base.Timeout = from.Timeout
	e.Base.Meta = from.Meta
	e.Properties = properties
	e.Base.SubSteps = subSteps
	return e, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	e := &ExportServiceWorkflowStep{}
	return e.FromWorkflowStep(from)
}

func (e *ExportServiceWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ExportServiceWorkflowStep, error) {
	var properties ExportServiceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	e.Base.Name = from.Name
	e.Base.DependsOn = from.DependsOn
	e.Base.Inputs = from.Inputs
	e.Base.Outputs = from.Outputs
	e.Base.If = from.If
	e.Base.Timeout = from.Timeout
	e.Base.Meta = from.Meta
	e.Properties = properties
	return e, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	e := &ExportServiceWorkflowStep{}
	return e.FromWorkflowSubStep(from)
}

func (e *ExportServiceWorkflowStep) If(_if string) *ExportServiceWorkflowStep {
	e.Base.If = _if
	return e
}

func (e *ExportServiceWorkflowStep) Alias(alias string) *ExportServiceWorkflowStep {
	e.Base.Meta.Alias = alias
	return e
}

func (e *ExportServiceWorkflowStep) Timeout(timeout string) *ExportServiceWorkflowStep {
	e.Base.Timeout = timeout
	return e
}

func (e *ExportServiceWorkflowStep) DependsOn(dependsOn []string) *ExportServiceWorkflowStep {
	e.Base.DependsOn = dependsOn
	return e
}

func (e *ExportServiceWorkflowStep) Inputs(input common.StepInputs) *ExportServiceWorkflowStep {
	e.Base.Inputs = input
	return e
}

func (e *ExportServiceWorkflowStep) Outputs(output common.StepOutputs) *ExportServiceWorkflowStep {
	e.Base.Outputs = output
	return e
}

func (e *ExportServiceWorkflowStep) DefName() string {
	return e.Base.Name
}

func (e *ExportServiceWorkflowStep) DefType() string {
	return ExportServiceType
}
