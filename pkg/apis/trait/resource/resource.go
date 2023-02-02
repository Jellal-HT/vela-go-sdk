/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ResourceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ResourceSpec{}

// ResourceSpec struct for ResourceSpec
type ResourceSpec struct {
	// Specify the amount of cpu for requests and limits
	cpu    *float32 `json:"cpu,omitempty"`
	limits *Limits  `json:"limits,omitempty"`
	// Specify the amount of memory for requests and limits
	memory   *string   `json:"memory,omitempty"`
	requests *Requests `json:"requests,omitempty"`
}

// NewResourceSpecWith instantiates a new ResourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSpecWith() *ResourceSpec {
	this := ResourceSpec{}
	var cpu float32 = 1
	this.cpu = &cpu
	var memory string = "2048Mi"
	this.memory = &memory
	return &this
}

// NewResourceSpec instantiates a new ResourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSpec() *ResourceSpec {
	this := ResourceSpec{}
	var cpu float32 = 1
	this.cpu = &cpu
	var memory string = "2048Mi"
	this.memory = &memory
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ResourceTrait) GetCpu() float32 {
	if o == nil || utils.IsNil(o.Properties.cpu) {
		var ret float32
		return ret
	}
	return *o.Properties.cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetCpuOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Properties.cpu) {
		return nil, false
	}
	return o.Properties.cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ResourceTrait) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.cpu) {
		return true
	}

	return false
}

// Cpu gets a reference to the given float32 and assigns it to the cpu field.
// cpu:  Specify the amount of cpu for requests and limits
func (o *ResourceTrait) Cpu(v float32) *ResourceTrait {
	o.Properties.cpu = &v
	return o
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *ResourceTrait) GetLimits() Limits {
	if o == nil || utils.IsNil(o.Properties.limits) {
		var ret Limits
		return ret
	}
	return *o.Properties.limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetLimitsOk() (*Limits, bool) {
	if o == nil || utils.IsNil(o.Properties.limits) {
		return nil, false
	}
	return o.Properties.limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *ResourceTrait) HasLimits() bool {
	if o != nil && !utils.IsNil(o.Properties.limits) {
		return true
	}

	return false
}

// Limits gets a reference to the given Limits and assigns it to the limits field.
// limits:
func (o *ResourceTrait) Limits(v Limits) *ResourceTrait {
	o.Properties.limits = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ResourceTrait) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.memory) {
		var ret string
		return ret
	}
	return *o.Properties.memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.memory) {
		return nil, false
	}
	return o.Properties.memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ResourceTrait) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.memory) {
		return true
	}

	return false
}

// Memory gets a reference to the given string and assigns it to the memory field.
// memory:  Specify the amount of memory for requests and limits
func (o *ResourceTrait) Memory(v string) *ResourceTrait {
	o.Properties.memory = &v
	return o
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *ResourceTrait) GetRequests() Requests {
	if o == nil || utils.IsNil(o.Properties.requests) {
		var ret Requests
		return ret
	}
	return *o.Properties.requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetRequestsOk() (*Requests, bool) {
	if o == nil || utils.IsNil(o.Properties.requests) {
		return nil, false
	}
	return o.Properties.requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *ResourceTrait) HasRequests() bool {
	if o != nil && !utils.IsNil(o.Properties.requests) {
		return true
	}

	return false
}

// Requests gets a reference to the given Requests and assigns it to the requests field.
// requests:
func (o *ResourceTrait) Requests(v Requests) *ResourceTrait {
	o.Properties.requests = &v
	return o
}

func (o ResourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.cpu) {
		toSerialize["cpu"] = o.cpu
	}
	if !utils.IsNil(o.limits) {
		toSerialize["limits"] = o.limits
	}
	if !utils.IsNil(o.memory) {
		toSerialize["memory"] = o.memory
	}
	if !utils.IsNil(o.requests) {
		toSerialize["requests"] = o.requests
	}
	return toSerialize, nil
}

type NullableResourceSpec struct {
	value *ResourceSpec
	isSet bool
}

func (v NullableResourceSpec) Get() *ResourceSpec {
	return v.value
}

func (v *NullableResourceSpec) Set(val *ResourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSpec(val *ResourceSpec) *NullableResourceSpec {
	return &NullableResourceSpec{value: val, isSet: true}
}

func (v NullableResourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ResourceType = "resource"

func init() {
	sdkcommon.RegisterTrait(ResourceType, FromTrait)
}

type ResourceTrait struct {
	Base       apis.TraitBase
	Properties ResourceSpec
}

func Resource() *ResourceTrait {
	r := &ResourceTrait{Base: apis.TraitBase{}}
	return r
}

func (r *ResourceTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(r.Properties),
		Type:       ResourceType,
	}
	return res
}

func (r *ResourceTrait) FromTrait(from common.ApplicationTrait) (*ResourceTrait, error) {
	var properties ResourceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Properties = properties
	return r, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	r := &ResourceTrait{}
	return r.FromTrait(from)
}

func (r *ResourceTrait) Type() string {
	return ResourceType
}
