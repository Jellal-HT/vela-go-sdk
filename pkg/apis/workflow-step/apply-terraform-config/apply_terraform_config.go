/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_config

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyTerraformConfigSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyTerraformConfigSpec{}

// ApplyTerraformConfigSpec struct for ApplyTerraformConfigSpec
type ApplyTerraformConfigSpec struct {
	// whether to delete resource
	deleteResource *bool `json:"deleteResource,omitempty"`
	forceDelete    *bool `json:"forceDelete,omitempty"`
	// the envs for job
	jobEnv      map[string]interface{} `json:"jobEnv,omitempty"`
	providerRef *ProviderRef           `json:"providerRef,omitempty"`
	// region is cloud provider's region. It will override the region in the region field of providerRef
	region *string `json:"region,omitempty"`
	source *Source `json:"source,omitempty"`
	// the variable in the configuration
	variable                   map[string]interface{}      `json:"variable,omitempty"`
	writeConnectionSecretToRef *WriteConnectionSecretToRef `json:"writeConnectionSecretToRef,omitempty"`
}

// NewApplyTerraformConfigSpecWith instantiates a new ApplyTerraformConfigSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyTerraformConfigSpecWith() *ApplyTerraformConfigSpec {
	this := ApplyTerraformConfigSpec{}
	var deleteResource bool = true
	this.deleteResource = &deleteResource
	var forceDelete bool = false
	this.forceDelete = &forceDelete
	return &this
}

// NewApplyTerraformConfigSpec instantiates a new ApplyTerraformConfigSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyTerraformConfigSpec() *ApplyTerraformConfigSpec {
	this := ApplyTerraformConfigSpec{}
	var deleteResource bool = true
	this.deleteResource = &deleteResource
	var forceDelete bool = false
	this.forceDelete = &forceDelete
	return &this
}

// GetDeleteResource returns the DeleteResource field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetDeleteResource() bool {
	if o == nil || utils.IsNil(o.Properties.deleteResource) {
		var ret bool
		return ret
	}
	return *o.Properties.deleteResource
}

// GetDeleteResourceOk returns a tuple with the DeleteResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetDeleteResourceOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.deleteResource) {
		return nil, false
	}
	return o.Properties.deleteResource, true
}

// HasDeleteResource returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasDeleteResource() bool {
	if o != nil && !utils.IsNil(o.Properties.deleteResource) {
		return true
	}

	return false
}

// DeleteResource gets a reference to the given bool and assigns it to the deleteResource field.
// deleteResource:  whether to delete resource
func (o *ApplyTerraformConfigWorkflowStep) DeleteResource(v bool) *ApplyTerraformConfigWorkflowStep {
	o.Properties.deleteResource = &v
	return o
}

// GetForceDelete returns the ForceDelete field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetForceDelete() bool {
	if o == nil || utils.IsNil(o.Properties.forceDelete) {
		var ret bool
		return ret
	}
	return *o.Properties.forceDelete
}

// GetForceDeleteOk returns a tuple with the ForceDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetForceDeleteOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.forceDelete) {
		return nil, false
	}
	return o.Properties.forceDelete, true
}

// HasForceDelete returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasForceDelete() bool {
	if o != nil && !utils.IsNil(o.Properties.forceDelete) {
		return true
	}

	return false
}

// ForceDelete gets a reference to the given bool and assigns it to the forceDelete field.
// forceDelete:
func (o *ApplyTerraformConfigWorkflowStep) ForceDelete(v bool) *ApplyTerraformConfigWorkflowStep {
	o.Properties.forceDelete = &v
	return o
}

// GetJobEnv returns the JobEnv field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetJobEnv() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.jobEnv) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.jobEnv
}

// GetJobEnvOk returns a tuple with the JobEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetJobEnvOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.jobEnv) {
		return map[string]interface{}{}, false
	}
	return o.Properties.jobEnv, true
}

// HasJobEnv returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasJobEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.jobEnv) {
		return true
	}

	return false
}

// JobEnv gets a reference to the given map[string]interface{} and assigns it to the jobEnv field.
// jobEnv:  the envs for job
func (o *ApplyTerraformConfigWorkflowStep) JobEnv(v map[string]interface{}) *ApplyTerraformConfigWorkflowStep {
	o.Properties.jobEnv = v
	return o
}

// GetProviderRef returns the ProviderRef field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetProviderRef() ProviderRef {
	if o == nil || utils.IsNil(o.Properties.providerRef) {
		var ret ProviderRef
		return ret
	}
	return *o.Properties.providerRef
}

// GetProviderRefOk returns a tuple with the ProviderRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetProviderRefOk() (*ProviderRef, bool) {
	if o == nil || utils.IsNil(o.Properties.providerRef) {
		return nil, false
	}
	return o.Properties.providerRef, true
}

// HasProviderRef returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasProviderRef() bool {
	if o != nil && !utils.IsNil(o.Properties.providerRef) {
		return true
	}

	return false
}

// ProviderRef gets a reference to the given ProviderRef and assigns it to the providerRef field.
// providerRef:
func (o *ApplyTerraformConfigWorkflowStep) ProviderRef(v ProviderRef) *ApplyTerraformConfigWorkflowStep {
	o.Properties.providerRef = &v
	return o
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetRegion() string {
	if o == nil || utils.IsNil(o.Properties.region) {
		var ret string
		return ret
	}
	return *o.Properties.region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetRegionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.region) {
		return nil, false
	}
	return o.Properties.region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasRegion() bool {
	if o != nil && !utils.IsNil(o.Properties.region) {
		return true
	}

	return false
}

// Region gets a reference to the given string and assigns it to the region field.
// region:  region is cloud provider's region. It will override the region in the region field of providerRef
func (o *ApplyTerraformConfigWorkflowStep) Region(v string) *ApplyTerraformConfigWorkflowStep {
	o.Properties.region = &v
	return o
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetSource() Source {
	if o == nil || utils.IsNil(o.Properties.source) {
		var ret Source
		return ret
	}
	return *o.Properties.source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetSourceOk() (*Source, bool) {
	if o == nil || utils.IsNil(o.Properties.source) {
		return nil, false
	}
	return o.Properties.source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasSource() bool {
	if o != nil && !utils.IsNil(o.Properties.source) {
		return true
	}

	return false
}

// Source gets a reference to the given Source and assigns it to the source field.
// source:
func (o *ApplyTerraformConfigWorkflowStep) Source(v Source) *ApplyTerraformConfigWorkflowStep {
	o.Properties.source = &v
	return o
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetVariable() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.variable) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetVariableOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.variable) {
		return map[string]interface{}{}, false
	}
	return o.Properties.variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasVariable() bool {
	if o != nil && !utils.IsNil(o.Properties.variable) {
		return true
	}

	return false
}

// Variable gets a reference to the given map[string]interface{} and assigns it to the variable field.
// variable:  the variable in the configuration
func (o *ApplyTerraformConfigWorkflowStep) Variable(v map[string]interface{}) *ApplyTerraformConfigWorkflowStep {
	o.Properties.variable = v
	return o
}

// GetWriteConnectionSecretToRef returns the WriteConnectionSecretToRef field value if set, zero value otherwise.
func (o *ApplyTerraformConfigWorkflowStep) GetWriteConnectionSecretToRef() WriteConnectionSecretToRef {
	if o == nil || utils.IsNil(o.Properties.writeConnectionSecretToRef) {
		var ret WriteConnectionSecretToRef
		return ret
	}
	return *o.Properties.writeConnectionSecretToRef
}

// GetWriteConnectionSecretToRefOk returns a tuple with the WriteConnectionSecretToRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyTerraformConfigWorkflowStep) GetWriteConnectionSecretToRefOk() (*WriteConnectionSecretToRef, bool) {
	if o == nil || utils.IsNil(o.Properties.writeConnectionSecretToRef) {
		return nil, false
	}
	return o.Properties.writeConnectionSecretToRef, true
}

// HasWriteConnectionSecretToRef returns a boolean if a field has been set.
func (o *ApplyTerraformConfigWorkflowStep) HasWriteConnectionSecretToRef() bool {
	if o != nil && !utils.IsNil(o.Properties.writeConnectionSecretToRef) {
		return true
	}

	return false
}

// WriteConnectionSecretToRef gets a reference to the given WriteConnectionSecretToRef and assigns it to the writeConnectionSecretToRef field.
// writeConnectionSecretToRef:
func (o *ApplyTerraformConfigWorkflowStep) WriteConnectionSecretToRef(v WriteConnectionSecretToRef) *ApplyTerraformConfigWorkflowStep {
	o.Properties.writeConnectionSecretToRef = &v
	return o
}

func (o ApplyTerraformConfigSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyTerraformConfigSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.deleteResource) {
		toSerialize["deleteResource"] = o.deleteResource
	}
	if !utils.IsNil(o.forceDelete) {
		toSerialize["forceDelete"] = o.forceDelete
	}
	if !utils.IsNil(o.jobEnv) {
		toSerialize["jobEnv"] = o.jobEnv
	}
	if !utils.IsNil(o.providerRef) {
		toSerialize["providerRef"] = o.providerRef
	}
	if !utils.IsNil(o.region) {
		toSerialize["region"] = o.region
	}
	if !utils.IsNil(o.source) {
		toSerialize["source"] = o.source
	}
	if !utils.IsNil(o.variable) {
		toSerialize["variable"] = o.variable
	}
	if !utils.IsNil(o.writeConnectionSecretToRef) {
		toSerialize["writeConnectionSecretToRef"] = o.writeConnectionSecretToRef
	}
	return toSerialize, nil
}

type NullableApplyTerraformConfigSpec struct {
	value *ApplyTerraformConfigSpec
	isSet bool
}

func (v NullableApplyTerraformConfigSpec) Get() *ApplyTerraformConfigSpec {
	return v.value
}

func (v *NullableApplyTerraformConfigSpec) Set(val *ApplyTerraformConfigSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyTerraformConfigSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyTerraformConfigSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyTerraformConfigSpec(val *ApplyTerraformConfigSpec) *NullableApplyTerraformConfigSpec {
	return &NullableApplyTerraformConfigSpec{value: val, isSet: true}
}

func (v NullableApplyTerraformConfigSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyTerraformConfigSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyTerraformConfigType = "apply-terraform-config"

func init() {
	sdkcommon.RegisterWorkflowStep(ApplyTerraformConfigType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ApplyTerraformConfigType, FromWorkflowSubStep)
}

type ApplyTerraformConfigWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ApplyTerraformConfigSpec
}

func ApplyTerraformConfig(name string) *ApplyTerraformConfigWorkflowStep {
	a := &ApplyTerraformConfigWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ApplyTerraformConfigType,
	}}
	return a
}

func (a *ApplyTerraformConfigWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range a.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  a.Base.DependsOn,
		If:         a.Base.If,
		Inputs:     a.Base.Inputs,
		Meta:       a.Base.Meta,
		Name:       a.Base.Name,
		Outputs:    a.Base.Outputs,
		Properties: util.Object2RawExtension(a.Properties),
		SubSteps:   subSteps,
		Timeout:    a.Base.Timeout,
		Type:       ApplyTerraformConfigType,
	}
	return res
}

func (a *ApplyTerraformConfigWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ApplyTerraformConfigWorkflowStep, error) {
	var properties ApplyTerraformConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := a.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	a.Base.Name = from.Name
	a.Base.DependsOn = from.DependsOn
	a.Base.Inputs = from.Inputs
	a.Base.Outputs = from.Outputs
	a.Base.If = from.If
	a.Base.Timeout = from.Timeout
	a.Base.Meta = from.Meta
	a.Properties = properties
	a.Base.SubSteps = subSteps
	return a, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	a := &ApplyTerraformConfigWorkflowStep{}
	return a.FromWorkflowStep(from)
}

func (a *ApplyTerraformConfigWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ApplyTerraformConfigWorkflowStep, error) {
	var properties ApplyTerraformConfigSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	a.Base.Name = from.Name
	a.Base.DependsOn = from.DependsOn
	a.Base.Inputs = from.Inputs
	a.Base.Outputs = from.Outputs
	a.Base.If = from.If
	a.Base.Timeout = from.Timeout
	a.Base.Meta = from.Meta
	a.Properties = properties
	return a, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	a := &ApplyTerraformConfigWorkflowStep{}
	return a.FromWorkflowSubStep(from)
}

func (a *ApplyTerraformConfigWorkflowStep) If(_if string) *ApplyTerraformConfigWorkflowStep {
	a.Base.If = _if
	return a
}

func (a *ApplyTerraformConfigWorkflowStep) Alias(alias string) *ApplyTerraformConfigWorkflowStep {
	a.Base.Meta.Alias = alias
	return a
}

func (a *ApplyTerraformConfigWorkflowStep) Timeout(timeout string) *ApplyTerraformConfigWorkflowStep {
	a.Base.Timeout = timeout
	return a
}

func (a *ApplyTerraformConfigWorkflowStep) DependsOn(dependsOn []string) *ApplyTerraformConfigWorkflowStep {
	a.Base.DependsOn = dependsOn
	return a
}

func (a *ApplyTerraformConfigWorkflowStep) Inputs(input common.StepInputs) *ApplyTerraformConfigWorkflowStep {
	a.Base.Inputs = input
	return a
}

func (a *ApplyTerraformConfigWorkflowStep) Outputs(output common.StepOutputs) *ApplyTerraformConfigWorkflowStep {
	a.Base.Outputs = output
	return a
}

func (a *ApplyTerraformConfigWorkflowStep) DefName() string {
	return a.Base.Name
}

func (a *ApplyTerraformConfigWorkflowStep) DefType() string {
	return ApplyTerraformConfigType
}
