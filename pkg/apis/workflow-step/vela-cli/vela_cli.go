/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vela_cli

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the VelaCliSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VelaCliSpec{}

// VelaCliSpec struct for VelaCliSpec
type VelaCliSpec struct {
	// Specify the name of the addon.
	addonName *string `json:"addonName,omitempty"`
	// Specify the vela command
	command []string `json:"command,omitempty"`
	// Specify the image
	image *string `json:"image,omitempty"`
	// specify serviceAccountName want to use
	serviceAccountName *string  `json:"serviceAccountName,omitempty"`
	storage            *Storage `json:"storage,omitempty"`
}

// NewVelaCliSpecWith instantiates a new VelaCliSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVelaCliSpecWith() *VelaCliSpec {
	this := VelaCliSpec{}
	var image string = "oamdev/vela-cli:v1.6.4"
	this.image = &image
	var serviceAccountName string = "kubevela-vela-core"
	this.serviceAccountName = &serviceAccountName
	return &this
}

// NewVelaCliSpec instantiates a new VelaCliSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelaCliSpec() *VelaCliSpec {
	this := VelaCliSpec{}
	var image string = "oamdev/vela-cli:v1.6.4"
	this.image = &image
	var serviceAccountName string = "kubevela-vela-core"
	this.serviceAccountName = &serviceAccountName
	return &this
}

// GetAddonName returns the AddonName field value if set, zero value otherwise.
func (o *VelaCliWorkflowStep) GetAddonName() string {
	if o == nil || utils.IsNil(o.Properties.addonName) {
		var ret string
		return ret
	}
	return *o.Properties.addonName
}

// GetAddonNameOk returns a tuple with the AddonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetAddonNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.addonName) {
		return nil, false
	}
	return o.Properties.addonName, true
}

// HasAddonName returns a boolean if a field has been set.
func (o *VelaCliWorkflowStep) HasAddonName() bool {
	if o != nil && !utils.IsNil(o.Properties.addonName) {
		return true
	}

	return false
}

// AddonName gets a reference to the given string and assigns it to the addonName field.
// addonName:  Specify the name of the addon.
func (o *VelaCliWorkflowStep) AddonName(v string) *VelaCliWorkflowStep {
	o.Properties.addonName = &v
	return o
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *VelaCliWorkflowStep) GetCommand() []string {
	if o == nil || utils.IsNil(o.Properties.command) {
		var ret []string
		return ret
	}
	return o.Properties.command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetCommandOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.command) {
		return nil, false
	}
	return o.Properties.command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *VelaCliWorkflowStep) HasCommand() bool {
	if o != nil && !utils.IsNil(o.Properties.command) {
		return true
	}

	return false
}

// Command gets a reference to the given []string and assigns it to the command field.
// command:  Specify the vela command
func (o *VelaCliWorkflowStep) Command(v []string) *VelaCliWorkflowStep {
	o.Properties.command = v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *VelaCliWorkflowStep) GetImage() string {
	if o == nil || utils.IsNil(o.Properties.image) {
		var ret string
		return ret
	}
	return *o.Properties.image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetImageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.image) {
		return nil, false
	}
	return o.Properties.image, true
}

// HasImage returns a boolean if a field has been set.
func (o *VelaCliWorkflowStep) HasImage() bool {
	if o != nil && !utils.IsNil(o.Properties.image) {
		return true
	}

	return false
}

// Image gets a reference to the given string and assigns it to the image field.
// image:  Specify the image
func (o *VelaCliWorkflowStep) Image(v string) *VelaCliWorkflowStep {
	o.Properties.image = &v
	return o
}

// GetServiceAccountName returns the ServiceAccountName field value if set, zero value otherwise.
func (o *VelaCliWorkflowStep) GetServiceAccountName() string {
	if o == nil || utils.IsNil(o.Properties.serviceAccountName) {
		var ret string
		return ret
	}
	return *o.Properties.serviceAccountName
}

// GetServiceAccountNameOk returns a tuple with the ServiceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetServiceAccountNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.serviceAccountName) {
		return nil, false
	}
	return o.Properties.serviceAccountName, true
}

// HasServiceAccountName returns a boolean if a field has been set.
func (o *VelaCliWorkflowStep) HasServiceAccountName() bool {
	if o != nil && !utils.IsNil(o.Properties.serviceAccountName) {
		return true
	}

	return false
}

// ServiceAccountName gets a reference to the given string and assigns it to the serviceAccountName field.
// serviceAccountName:  specify serviceAccountName want to use
func (o *VelaCliWorkflowStep) ServiceAccountName(v string) *VelaCliWorkflowStep {
	o.Properties.serviceAccountName = &v
	return o
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *VelaCliWorkflowStep) GetStorage() Storage {
	if o == nil || utils.IsNil(o.Properties.storage) {
		var ret Storage
		return ret
	}
	return *o.Properties.storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetStorageOk() (*Storage, bool) {
	if o == nil || utils.IsNil(o.Properties.storage) {
		return nil, false
	}
	return o.Properties.storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *VelaCliWorkflowStep) HasStorage() bool {
	if o != nil && !utils.IsNil(o.Properties.storage) {
		return true
	}

	return false
}

// Storage gets a reference to the given Storage and assigns it to the storage field.
// storage:
func (o *VelaCliWorkflowStep) Storage(v Storage) *VelaCliWorkflowStep {
	o.Properties.storage = &v
	return o
}

func (o VelaCliSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VelaCliSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.addonName) {
		toSerialize["addonName"] = o.addonName
	}
	if !utils.IsNil(o.command) {
		toSerialize["command"] = o.command
	}
	if !utils.IsNil(o.image) {
		toSerialize["image"] = o.image
	}
	if !utils.IsNil(o.serviceAccountName) {
		toSerialize["serviceAccountName"] = o.serviceAccountName
	}
	if !utils.IsNil(o.storage) {
		toSerialize["storage"] = o.storage
	}
	return toSerialize, nil
}

type NullableVelaCliSpec struct {
	value *VelaCliSpec
	isSet bool
}

func (v NullableVelaCliSpec) Get() *VelaCliSpec {
	return v.value
}

func (v *NullableVelaCliSpec) Set(val *VelaCliSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVelaCliSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVelaCliSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelaCliSpec(val *VelaCliSpec) *NullableVelaCliSpec {
	return &NullableVelaCliSpec{value: val, isSet: true}
}

func (v NullableVelaCliSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelaCliSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const VelaCliType = "vela-cli"

func init() {
	sdkcommon.RegisterWorkflowStep(VelaCliType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(VelaCliType, FromWorkflowSubStep)
}

type VelaCliWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties VelaCliSpec
}

func VelaCli(name string) *VelaCliWorkflowStep {
	v := &VelaCliWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: VelaCliType,
	}}
	return v
}

func (v *VelaCliWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range v.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  v.Base.DependsOn,
		If:         v.Base.If,
		Inputs:     v.Base.Inputs,
		Meta:       v.Base.Meta,
		Name:       v.Base.Name,
		Outputs:    v.Base.Outputs,
		Properties: util.Object2RawExtension(v.Properties),
		SubSteps:   subSteps,
		Timeout:    v.Base.Timeout,
		Type:       VelaCliType,
	}
	return res
}

func (v *VelaCliWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*VelaCliWorkflowStep, error) {
	var properties VelaCliSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := v.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	v.Base.Name = from.Name
	v.Base.DependsOn = from.DependsOn
	v.Base.Inputs = from.Inputs
	v.Base.Outputs = from.Outputs
	v.Base.If = from.If
	v.Base.Timeout = from.Timeout
	v.Base.Meta = from.Meta
	v.Properties = properties
	v.Base.SubSteps = subSteps
	return v, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	v := &VelaCliWorkflowStep{}
	return v.FromWorkflowStep(from)
}

func (v *VelaCliWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*VelaCliWorkflowStep, error) {
	var properties VelaCliSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	v.Base.Name = from.Name
	v.Base.DependsOn = from.DependsOn
	v.Base.Inputs = from.Inputs
	v.Base.Outputs = from.Outputs
	v.Base.If = from.If
	v.Base.Timeout = from.Timeout
	v.Base.Meta = from.Meta
	v.Properties = properties
	return v, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	v := &VelaCliWorkflowStep{}
	return v.FromWorkflowSubStep(from)
}

func (v *VelaCliWorkflowStep) If(_if string) *VelaCliWorkflowStep {
	v.Base.If = _if
	return v
}

func (v *VelaCliWorkflowStep) Alias(alias string) *VelaCliWorkflowStep {
	v.Base.Meta.Alias = alias
	return v
}

func (v *VelaCliWorkflowStep) Timeout(timeout string) *VelaCliWorkflowStep {
	v.Base.Timeout = timeout
	return v
}

func (v *VelaCliWorkflowStep) DependsOn(dependsOn []string) *VelaCliWorkflowStep {
	v.Base.DependsOn = dependsOn
	return v
}

func (v *VelaCliWorkflowStep) Inputs(input common.StepInputs) *VelaCliWorkflowStep {
	v.Base.Inputs = input
	return v
}

func (v *VelaCliWorkflowStep) Outputs(output common.StepOutputs) *VelaCliWorkflowStep {
	v.Base.Outputs = output
	return v
}

func (v *VelaCliWorkflowStep) DefName() string {
	return v.Base.Name
}

func (v *VelaCliWorkflowStep) DefType() string {
	return VelaCliType
}
