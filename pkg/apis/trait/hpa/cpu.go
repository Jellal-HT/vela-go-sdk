/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpa

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Cpu type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Cpu{}

// Cpu struct for Cpu
type Cpu struct {
	// Specify resource metrics in terms of percentage(\"Utilization\") or direct value(\"AverageValue\")
	type_ *string `json:"type,omitempty"`
	// Specify the value of CPU utilization or averageValue
	value *int32 `json:"value,omitempty"`
}

// NewCpuWith instantiates a new Cpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuWith() *Cpu {
	this := Cpu{}
	var type_ string = "Utilization"
	this.type_ = &type_
	var value int32 = 50
	this.value = &value
	return &this
}

// NewCpu instantiates a new Cpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpu() *Cpu {
	this := Cpu{}
	var type_ string = "Utilization"
	this.type_ = &type_
	var value int32 = 50
	this.value = &value
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Cpu) GetType() string {
	if o == nil || utils.IsNil(o.type_) {
		var ret string
		return ret
	}
	return *o.type_
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.type_) {
		return nil, false
	}
	return o.type_, true
}

// HasType returns a boolean if a field has been set.
func (o *Cpu) HasType() bool {
	if o != nil && !utils.IsNil(o.type_) {
		return true
	}

	return false
}

// Type gets a reference to the given string and assigns it to the type_ field.
// type_:  Specify resource metrics in terms of percentage(\"Utilization\") or direct value(\"AverageValue\")
func (o *Cpu) Type(v string) *Cpu {
	o.type_ = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Cpu) GetValue() int32 {
	if o == nil || utils.IsNil(o.value) {
		var ret int32
		return ret
	}
	return *o.value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetValueOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.value) {
		return nil, false
	}
	return o.value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Cpu) HasValue() bool {
	if o != nil && !utils.IsNil(o.value) {
		return true
	}

	return false
}

// Value gets a reference to the given int32 and assigns it to the value field.
// value:  Specify the value of CPU utilization or averageValue
func (o *Cpu) Value(v int32) *Cpu {
	o.value = &v
	return o
}

func (o Cpu) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.type_) {
		toSerialize["type"] = o.type_
	}
	if !utils.IsNil(o.value) {
		toSerialize["value"] = o.value
	}
	return toSerialize, nil
}

type NullableCpu struct {
	value *Cpu
	isSet bool
}

func (v NullableCpu) Get() *Cpu {
	return v.value
}

func (v *NullableCpu) Set(val *Cpu) {
	v.value = val
	v.isSet = true
}

func (v NullableCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpu(val *Cpu) *NullableCpu {
	return &NullableCpu{value: val, isSet: true}
}

func (v NullableCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
