/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ConfigMapKeyRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ConfigMapKeyRef{}

// ConfigMapKeyRef Selects a key of a config map in the pod's namespace
type ConfigMapKeyRef struct {
	// The key of the config map to select from. Must be a valid secret key
	key *string `json:"key,omitempty"`
	// The name of the config map in the pod's namespace to select from
	name *string `json:"name,omitempty"`
}

// NewConfigMapKeyRefWith instantiates a new ConfigMapKeyRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigMapKeyRefWith() *ConfigMapKeyRef {
	this := ConfigMapKeyRef{}
	return &this
}

// NewConfigMapKeyRef instantiates a new ConfigMapKeyRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMapKeyRef() *ConfigMapKeyRef {
	this := ConfigMapKeyRef{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ConfigMapKeyRef) GetKey() string {
	if o == nil || utils.IsNil(o.key) {
		var ret string
		return ret
	}
	return *o.key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMapKeyRef) GetKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.key) {
		return nil, false
	}
	return o.key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ConfigMapKeyRef) HasKey() bool {
	if o != nil && !utils.IsNil(o.key) {
		return true
	}

	return false
}

// Key gets a reference to the given string and assigns it to the key field.
// key:  The key of the config map to select from. Must be a valid secret key
func (o *ConfigMapKeyRef) Key(v string) *ConfigMapKeyRef {
	o.key = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigMapKeyRef) GetName() string {
	if o == nil || utils.IsNil(o.name) {
		var ret string
		return ret
	}
	return *o.name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMapKeyRef) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.name) {
		return nil, false
	}
	return o.name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigMapKeyRef) HasName() bool {
	if o != nil && !utils.IsNil(o.name) {
		return true
	}

	return false
}

// Name gets a reference to the given string and assigns it to the name field.
// name:  The name of the config map in the pod's namespace to select from
func (o *ConfigMapKeyRef) Name(v string) *ConfigMapKeyRef {
	o.name = &v
	return o
}

func (o ConfigMapKeyRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigMapKeyRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.key) {
		toSerialize["key"] = o.key
	}
	if !utils.IsNil(o.name) {
		toSerialize["name"] = o.name
	}
	return toSerialize, nil
}

type NullableConfigMapKeyRef struct {
	value *ConfigMapKeyRef
	isSet bool
}

func (v NullableConfigMapKeyRef) Get() *ConfigMapKeyRef {
	return v.value
}

func (v *NullableConfigMapKeyRef) Set(val *ConfigMapKeyRef) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigMapKeyRef) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigMapKeyRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigMapKeyRef(val *ConfigMapKeyRef) *NullableConfigMapKeyRef {
	return &NullableConfigMapKeyRef{value: val, isSet: true}
}

func (v NullableConfigMapKeyRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigMapKeyRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
