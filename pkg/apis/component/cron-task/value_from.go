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

// checks if the ValueFrom type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ValueFrom{}

// ValueFrom Specifies a source the value of this var should come from
type ValueFrom struct {
	ConfigMapKeyRef *ConfigMapKeyRef `json:"configMapKeyRef,omitempty"`
	SecretKeyRef    *SecretKeyRef    `json:"secretKeyRef,omitempty"`
}

// NewValueFromWith instantiates a new ValueFrom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueFromWith() *ValueFrom {
	this := ValueFrom{}
	return &this
}

// NewValueFrom instantiates a new ValueFrom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueFrom() *ValueFrom {
	this := ValueFrom{}
	return &this
}

// GetConfigMapKeyRef returns the ConfigMapKeyRef field value if set, zero value otherwise.
func (o *ValueFrom) GetConfigMapKeyRef() ConfigMapKeyRef {
	if o == nil || utils.IsNil(o.ConfigMapKeyRef) {
		var ret ConfigMapKeyRef
		return ret
	}
	return *o.ConfigMapKeyRef
}

// GetConfigMapKeyRefOk returns a tuple with the ConfigMapKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueFrom) GetConfigMapKeyRefOk() (*ConfigMapKeyRef, bool) {
	if o == nil || utils.IsNil(o.ConfigMapKeyRef) {
		return nil, false
	}
	return o.ConfigMapKeyRef, true
}

// HasConfigMapKeyRef returns a boolean if a field has been set.
func (o *ValueFrom) HasConfigMapKeyRef() bool {
	if o != nil && !utils.IsNil(o.ConfigMapKeyRef) {
		return true
	}

	return false
}

// SetConfigMapKeyRef gets a reference to the given ConfigMapKeyRef and assigns it to the configMapKeyRef field.
// ConfigMapKeyRef:
func (o *ValueFrom) SetConfigMapKeyRef(v ConfigMapKeyRef) *ValueFrom {
	o.ConfigMapKeyRef = &v
	return o
}

// GetSecretKeyRef returns the SecretKeyRef field value if set, zero value otherwise.
func (o *ValueFrom) GetSecretKeyRef() SecretKeyRef {
	if o == nil || utils.IsNil(o.SecretKeyRef) {
		var ret SecretKeyRef
		return ret
	}
	return *o.SecretKeyRef
}

// GetSecretKeyRefOk returns a tuple with the SecretKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueFrom) GetSecretKeyRefOk() (*SecretKeyRef, bool) {
	if o == nil || utils.IsNil(o.SecretKeyRef) {
		return nil, false
	}
	return o.SecretKeyRef, true
}

// HasSecretKeyRef returns a boolean if a field has been set.
func (o *ValueFrom) HasSecretKeyRef() bool {
	if o != nil && !utils.IsNil(o.SecretKeyRef) {
		return true
	}

	return false
}

// SetSecretKeyRef gets a reference to the given SecretKeyRef and assigns it to the secretKeyRef field.
// SecretKeyRef:
func (o *ValueFrom) SetSecretKeyRef(v SecretKeyRef) *ValueFrom {
	o.SecretKeyRef = &v
	return o
}

func (o ValueFrom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueFrom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ConfigMapKeyRef) {
		toSerialize["configMapKeyRef"] = o.ConfigMapKeyRef
	}
	if !utils.IsNil(o.SecretKeyRef) {
		toSerialize["secretKeyRef"] = o.SecretKeyRef
	}
	return toSerialize, nil
}

type NullableValueFrom struct {
	value *ValueFrom
	isSet bool
}

func (v NullableValueFrom) Get() *ValueFrom {
	return v.value
}

func (v *NullableValueFrom) Set(val *ValueFrom) {
	v.value = val
	v.isSet = true
}

func (v NullableValueFrom) IsSet() bool {
	return v.isSet
}

func (v *NullableValueFrom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueFrom(val *ValueFrom) *NullableValueFrom {
	return &NullableValueFrom{value: val, isSet: true}
}

func (v NullableValueFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueFrom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
