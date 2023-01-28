/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_object

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/api"
	"vela-go-sdk/apis/utils"
)

// checks if the ApplyObjectSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyObjectSpec{}

// ApplyObjectSpec struct for ApplyObjectSpec
type ApplyObjectSpec struct {
	// The cluster you want to apply the resource to, default is the current control plane cluster
	cluster string `json:"cluster"`
	// Specify Kubernetes native resource object to be applied
	value map[string]interface{} `json:"value"`
}

// NewApplyObjectSpecWith instantiates a new ApplyObjectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyObjectSpecWith(cluster string, value map[string]interface{}) *ApplyObjectSpec {
	this := ApplyObjectSpec{}
	this.cluster = cluster
	this.value = value
	return &this
}

// NewApplyObjectSpec instantiates a new ApplyObjectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyObjectSpec() *ApplyObjectSpec {
	this := ApplyObjectSpec{}
	var cluster string = ""
	this.cluster = cluster
	return &this
}

// GetCluster returns the Cluster field value
func (o *ApplyObjectWorkflowStep) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ApplyObjectWorkflowStep) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.cluster, true
}

// Cluster sets field value
func (o *ApplyObjectWorkflowStep) Cluster(v string) *ApplyObjectWorkflowStep {
	o.Properties.cluster = v
	return o
}

// GetValue returns the Value field value
func (o *ApplyObjectWorkflowStep) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties.value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ApplyObjectWorkflowStep) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties.value, true
}

// Value sets field value
func (o *ApplyObjectWorkflowStep) Value(v map[string]interface{}) *ApplyObjectWorkflowStep {
	o.Properties.value = v
	return o
}

func (o ApplyObjectSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyObjectSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.cluster
	toSerialize["value"] = o.value
	return toSerialize, nil
}

type NullableApplyObjectSpec struct {
	value *ApplyObjectSpec
	isSet bool
}

func (v NullableApplyObjectSpec) Get() *ApplyObjectSpec {
	return v.value
}

func (v *NullableApplyObjectSpec) Set(val *ApplyObjectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyObjectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyObjectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyObjectSpec(val *ApplyObjectSpec) *NullableApplyObjectSpec {
	return &NullableApplyObjectSpec{value: val, isSet: true}
}

func (v NullableApplyObjectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyObjectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyObjectType = "apply-object"

type ApplyObjectWorkflowStep struct {
	Base       api.WorkflowStepBase
	Properties ApplyObjectSpec
}

func ApplyObject() *ApplyObjectWorkflowStep {
	a := &ApplyObjectWorkflowStep{Base: api.WorkflowStepBase{}}
	return a
}

func (a *ApplyObjectWorkflowStep) Build() common.WorkflowStep {
	res := common.WorkflowStep{
		Properties: util.Object2RawExtension(a.Properties),
		Type:       ApplyObjectType,
	}
	return res
}

func (a *ApplyObjectWorkflowStep) Props() *ApplyObjectSpec {
	return &a.Properties
}
