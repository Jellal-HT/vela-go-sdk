/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vela_cli

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HostPath type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostPath{}

// HostPath struct for HostPath
type HostPath struct {
	mountPath *string `json:"mountPath,omitempty"`
	name      *string `json:"name,omitempty"`
	path      *string `json:"path,omitempty"`
	type_     *string `json:"type,omitempty"`
}

// NewHostPathWith instantiates a new HostPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostPathWith() *HostPath {
	this := HostPath{}
	var type_ string = "Directory"
	this.type_ = &type_
	return &this
}

// NewHostPath instantiates a new HostPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostPath() *HostPath {
	this := HostPath{}
	var type_ string = "Directory"
	this.type_ = &type_
	return &this
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *HostPath) GetMountPath() string {
	if o == nil || utils.IsNil(o.mountPath) {
		var ret string
		return ret
	}
	return *o.mountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetMountPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.mountPath) {
		return nil, false
	}
	return o.mountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *HostPath) HasMountPath() bool {
	if o != nil && !utils.IsNil(o.mountPath) {
		return true
	}

	return false
}

// MountPath gets a reference to the given string and assigns it to the mountPath field.
// mountPath:
func (o *HostPath) MountPath(v string) *HostPath {
	o.mountPath = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostPath) GetName() string {
	if o == nil || utils.IsNil(o.name) {
		var ret string
		return ret
	}
	return *o.name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.name) {
		return nil, false
	}
	return o.name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostPath) HasName() bool {
	if o != nil && !utils.IsNil(o.name) {
		return true
	}

	return false
}

// Name gets a reference to the given string and assigns it to the name field.
// name:
func (o *HostPath) Name(v string) *HostPath {
	o.name = &v
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HostPath) GetPath() string {
	if o == nil || utils.IsNil(o.path) {
		var ret string
		return ret
	}
	return *o.path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.path) {
		return nil, false
	}
	return o.path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HostPath) HasPath() bool {
	if o != nil && !utils.IsNil(o.path) {
		return true
	}

	return false
}

// Path gets a reference to the given string and assigns it to the path field.
// path:
func (o *HostPath) Path(v string) *HostPath {
	o.path = &v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HostPath) GetType() string {
	if o == nil || utils.IsNil(o.type_) {
		var ret string
		return ret
	}
	return *o.type_
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.type_) {
		return nil, false
	}
	return o.type_, true
}

// HasType returns a boolean if a field has been set.
func (o *HostPath) HasType() bool {
	if o != nil && !utils.IsNil(o.type_) {
		return true
	}

	return false
}

// Type gets a reference to the given string and assigns it to the type_ field.
// type_:
func (o *HostPath) Type(v string) *HostPath {
	o.type_ = &v
	return o
}

func (o HostPath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.mountPath) {
		toSerialize["mountPath"] = o.mountPath
	}
	if !utils.IsNil(o.name) {
		toSerialize["name"] = o.name
	}
	if !utils.IsNil(o.path) {
		toSerialize["path"] = o.path
	}
	if !utils.IsNil(o.type_) {
		toSerialize["type"] = o.type_
	}
	return toSerialize, nil
}

type NullableHostPath struct {
	value *HostPath
	isSet bool
}

func (v NullableHostPath) Get() *HostPath {
	return v.value
}

func (v *NullableHostPath) Set(val *HostPath) {
	v.value = val
	v.isSet = true
}

func (v NullableHostPath) IsSet() bool {
	return v.isSet
}

func (v *NullableHostPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostPath(val *HostPath) *NullableHostPath {
	return &NullableHostPath{value: val, isSet: true}
}

func (v NullableHostPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
