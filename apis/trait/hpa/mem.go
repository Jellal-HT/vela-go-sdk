/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpa

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the Mem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Mem{}

// Mem struct for Mem
type Mem struct {
	// Specify resource metrics in terms of percentage(\"Utilization\") or direct value(\"AverageValue\")
	type_ string `json:"type"`
	// Specify  the value of MEM utilization or averageValue
	value int32 `json:"value"`
}

// NewMemWith instantiates a new Mem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemWith(type_ string, value int32) *Mem {
	this := Mem{}
	this.type_ = type_
	this.value = value
	return &this
}

// NewMem instantiates a new Mem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMem() *Mem {
	this := Mem{}
	var type_ string = "Utilization"
	this.type_ = type_
	var value int32 = 50
	this.value = value
	return &this
}

// GetType returns the Type field value
func (o *Mem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Mem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *Mem) Type(v string) *Mem {
	o.type_ = v
	return o
}

// GetValue returns the Value field value
func (o *Mem) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Mem) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.value, true
}

// Value sets field value
func (o *Mem) Value(v int32) *Mem {
	o.value = v
	return o
}

func (o Mem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.type_
	toSerialize["value"] = o.value
	return toSerialize, nil
}

type NullableMem struct {
	value *Mem
	isSet bool
}

func (v NullableMem) Get() *Mem {
	return v.value
}

func (v *NullableMem) Set(val *Mem) {
	v.value = val
	v.isSet = true
}

func (v NullableMem) IsSet() bool {
	return v.isSet
}

func (v *NullableMem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMem(val *Mem) *NullableMem {
	return &NullableMem{value: val, isSet: true}
}

func (v NullableMem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}