/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the HostPath type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostPath{}

// HostPath struct for HostPath
type HostPath struct {
	mountPath        string  `json:"mountPath"`
	mountPropagation *string `json:"mountPropagation,omitempty"`
	name             string  `json:"name"`
	path             string  `json:"path"`
	readOnly         *bool   `json:"readOnly,omitempty"`
}

// NewHostPathWith instantiates a new HostPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostPathWith(mountPath string, name string, path string) *HostPath {
	this := HostPath{}
	this.mountPath = mountPath
	this.name = name
	this.path = path
	return &this
}

// NewHostPath instantiates a new HostPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostPath() *HostPath {
	this := HostPath{}
	return &this
}

// GetMountPath returns the MountPath field value
func (o *HostPath) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.mountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *HostPath) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mountPath, true
}

// MountPath sets field value
func (o *HostPath) MountPath(v string) *HostPath {
	o.mountPath = v
	return o
}

// GetMountPropagation returns the MountPropagation field value if set, zero value otherwise.
func (o *HostPath) GetMountPropagation() string {
	if o == nil || utils.IsNil(o.mountPropagation) {
		var ret string
		return ret
	}
	return *o.mountPropagation
}

// GetMountPropagationOk returns a tuple with the MountPropagation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetMountPropagationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.mountPropagation) {
		return nil, false
	}
	return o.mountPropagation, true
}

// HasMountPropagation returns a boolean if a field has been set.
func (o *HostPath) HasMountPropagation() bool {
	if o != nil && !utils.IsNil(o.mountPropagation) {
		return true
	}

	return false
}

// MountPropagation gets a reference to the given string and assigns it to the mountPropagation field.
// mountPropagation:
func (o *HostPath) MountPropagation(v string) *HostPath {
	o.mountPropagation = &v
	return o
}

// GetName returns the Name field value
func (o *HostPath) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HostPath) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *HostPath) Name(v string) *HostPath {
	o.name = v
	return o
}

// GetPath returns the Path field value
func (o *HostPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *HostPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.path, true
}

// Path sets field value
func (o *HostPath) Path(v string) *HostPath {
	o.path = v
	return o
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *HostPath) GetReadOnly() bool {
	if o == nil || utils.IsNil(o.readOnly) {
		var ret bool
		return ret
	}
	return *o.readOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostPath) GetReadOnlyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.readOnly) {
		return nil, false
	}
	return o.readOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *HostPath) HasReadOnly() bool {
	if o != nil && !utils.IsNil(o.readOnly) {
		return true
	}

	return false
}

// ReadOnly gets a reference to the given bool and assigns it to the readOnly field.
// readOnly:
func (o *HostPath) ReadOnly(v bool) *HostPath {
	o.readOnly = &v
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
	toSerialize["mountPath"] = o.mountPath
	if !utils.IsNil(o.mountPropagation) {
		toSerialize["mountPropagation"] = o.mountPropagation
	}
	toSerialize["name"] = o.name
	toSerialize["path"] = o.path
	if !utils.IsNil(o.readOnly) {
		toSerialize["readOnly"] = o.readOnly
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