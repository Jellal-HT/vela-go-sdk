/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nocalhost

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the Env type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Env{}

// Env struct for Env
type Env struct {
	name  string `json:"name"`
	value string `json:"value"`
}

// NewEnvWith instantiates a new Env object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvWith(name string, value string) *Env {
	this := Env{}
	this.name = name
	this.value = value
	return &this
}

// NewEnv instantiates a new Env object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnv() *Env {
	this := Env{}
	return &this
}

// GetName returns the Name field value
func (o *Env) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Env) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *Env) Name(v string) *Env {
	o.name = v
	return o
}

// GetValue returns the Value field value
func (o *Env) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Env) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.value, true
}

// Value sets field value
func (o *Env) Value(v string) *Env {
	o.value = v
	return o
}

func (o Env) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Env) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.name
	toSerialize["value"] = o.value
	return toSerialize, nil
}

type NullableEnv struct {
	value *Env
	isSet bool
}

func (v NullableEnv) Get() *Env {
	return v.value
}

func (v *NullableEnv) Set(val *Env) {
	v.value = val
	v.isSet = true
}

func (v NullableEnv) IsSet() bool {
	return v.isSet
}

func (v *NullableEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnv(val *Env) *NullableEnv {
	return &NullableEnv{value: val, isSet: true}
}

func (v NullableEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}