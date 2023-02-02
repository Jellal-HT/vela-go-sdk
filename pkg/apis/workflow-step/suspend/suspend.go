/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package suspend

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the SuspendSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SuspendSpec{}

// SuspendSpec struct for SuspendSpec
type SuspendSpec struct {
	// Specify the wait duration time to resume workflow such as \"30s\", \"1min\" or \"2m15s\"
	duration *string `json:"duration,omitempty"`
}

// NewSuspendSpecWith instantiates a new SuspendSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuspendSpecWith() *SuspendSpec {
	this := SuspendSpec{}
	return &this
}

// NewSuspendSpec instantiates a new SuspendSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuspendSpec() *SuspendSpec {
	this := SuspendSpec{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SuspendWorkflowStep) GetDuration() string {
	if o == nil || utils.IsNil(o.Properties.duration) {
		var ret string
		return ret
	}
	return *o.Properties.duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuspendWorkflowStep) GetDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.duration) {
		return nil, false
	}
	return o.Properties.duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SuspendWorkflowStep) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Properties.duration) {
		return true
	}

	return false
}

// Duration gets a reference to the given string and assigns it to the duration field.
// duration:  Specify the wait duration time to resume workflow such as \"30s\", \"1min\" or \"2m15s\"
func (o *SuspendWorkflowStep) Duration(v string) *SuspendWorkflowStep {
	o.Properties.duration = &v
	return o
}

func (o SuspendSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuspendSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.duration) {
		toSerialize["duration"] = o.duration
	}
	return toSerialize, nil
}

type NullableSuspendSpec struct {
	value *SuspendSpec
	isSet bool
}

func (v NullableSuspendSpec) Get() *SuspendSpec {
	return v.value
}

func (v *NullableSuspendSpec) Set(val *SuspendSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSuspendSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSuspendSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuspendSpec(val *SuspendSpec) *NullableSuspendSpec {
	return &NullableSuspendSpec{value: val, isSet: true}
}

func (v NullableSuspendSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuspendSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const SuspendType = "suspend"

func init() {
	sdkcommon.RegisterWorkflowStep(SuspendType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(SuspendType, FromWorkflowSubStep)
}

type SuspendWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties SuspendSpec
}

func Suspend(name string) *SuspendWorkflowStep {
	s := &SuspendWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: SuspendType,
	}}
	return s
}

func (s *SuspendWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range s.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  s.Base.DependsOn,
		If:         s.Base.If,
		Inputs:     s.Base.Inputs,
		Meta:       s.Base.Meta,
		Name:       s.Base.Name,
		Outputs:    s.Base.Outputs,
		Properties: util.Object2RawExtension(s.Properties),
		SubSteps:   subSteps,
		Timeout:    s.Base.Timeout,
		Type:       SuspendType,
	}
	return res
}

func (s *SuspendWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*SuspendWorkflowStep, error) {
	var properties SuspendSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := s.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	s.Base.Name = from.Name
	s.Base.DependsOn = from.DependsOn
	s.Base.Inputs = from.Inputs
	s.Base.Outputs = from.Outputs
	s.Base.If = from.If
	s.Base.Timeout = from.Timeout
	s.Base.Meta = from.Meta
	s.Properties = properties
	s.Base.SubSteps = subSteps
	return s, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	s := &SuspendWorkflowStep{}
	return s.FromWorkflowStep(from)
}

func (s *SuspendWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*SuspendWorkflowStep, error) {
	var properties SuspendSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	s.Base.Name = from.Name
	s.Base.DependsOn = from.DependsOn
	s.Base.Inputs = from.Inputs
	s.Base.Outputs = from.Outputs
	s.Base.If = from.If
	s.Base.Timeout = from.Timeout
	s.Base.Meta = from.Meta
	s.Properties = properties
	return s, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	s := &SuspendWorkflowStep{}
	return s.FromWorkflowSubStep(from)
}

func (s *SuspendWorkflowStep) If(_if string) *SuspendWorkflowStep {
	s.Base.If = _if
	return s
}

func (s *SuspendWorkflowStep) Alias(alias string) *SuspendWorkflowStep {
	s.Base.Meta.Alias = alias
	return s
}

func (s *SuspendWorkflowStep) Timeout(timeout string) *SuspendWorkflowStep {
	s.Base.Timeout = timeout
	return s
}

func (s *SuspendWorkflowStep) DependsOn(dependsOn []string) *SuspendWorkflowStep {
	s.Base.DependsOn = dependsOn
	return s
}

func (s *SuspendWorkflowStep) Inputs(input common.StepInputs) *SuspendWorkflowStep {
	s.Base.Inputs = input
	return s
}

func (s *SuspendWorkflowStep) Outputs(output common.StepOutputs) *SuspendWorkflowStep {
	s.Base.Outputs = output
	return s
}

func (s *SuspendWorkflowStep) DefName() string {
	return s.Base.Name
}

func (s *SuspendWorkflowStep) DefType() string {
	return SuspendType
}
