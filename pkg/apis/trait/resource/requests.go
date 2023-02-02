/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Requests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Requests{}

// Requests Specify the resources in requests
type Requests struct {
	// Specify the amount of cpu for requests
	Cpu *float32 `json:"cpu,omitempty"`
	// Specify the amount of memory for requests
	Memory *string `json:"memory,omitempty"`
}

// NewRequestsWith instantiates a new Requests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsWith() *Requests {
	this := Requests{}
	var cpu float32 = 1
	this.Cpu = &cpu
	var memory string = "2048Mi"
	this.Memory = &memory
	return &this
}

// NewRequests instantiates a new Requests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequests() *Requests {
	this := Requests{}
	var cpu float32 = 1
	this.Cpu = &cpu
	var memory string = "2048Mi"
	this.Memory = &memory
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Requests) GetCpu() float32 {
	if o == nil || utils.IsNil(o.Cpu) {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requests) GetCpuOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Requests) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the cpu field.
// Cpu:  Specify the amount of cpu for requests
func (o *Requests) SetCpu(v float32) *Requests {
	o.Cpu = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Requests) GetMemory() string {
	if o == nil || utils.IsNil(o.Memory) {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requests) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Requests) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// Memory:  Specify the amount of memory for requests
func (o *Requests) SetMemory(v string) *Requests {
	o.Memory = &v
	return o
}

func (o Requests) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Requests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !utils.IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableRequests struct {
	value *Requests
	isSet bool
}

func (v NullableRequests) Get() *Requests {
	return v.value
}

func (v *NullableRequests) Set(val *Requests) {
	v.value = val
	v.isSet = true
}

func (v NullableRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequests(val *Requests) *NullableRequests {
	return &NullableRequests{value: val, isSet: true}
}

func (v NullableRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
