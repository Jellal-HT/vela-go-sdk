/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package init_container

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the ExtraVolumeMounts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExtraVolumeMounts{}

// ExtraVolumeMounts struct for ExtraVolumeMounts
type ExtraVolumeMounts struct {
	// The mountPath for mount in the init container
	mountPath string `json:"mountPath"`
	// The name of the volume to be mounted
	name string `json:"name"`
}

// NewExtraVolumeMountsWith instantiates a new ExtraVolumeMounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraVolumeMountsWith(mountPath string, name string) *ExtraVolumeMounts {
	this := ExtraVolumeMounts{}
	this.mountPath = mountPath
	this.name = name
	return &this
}

// NewExtraVolumeMounts instantiates a new ExtraVolumeMounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtraVolumeMounts() *ExtraVolumeMounts {
	this := ExtraVolumeMounts{}
	return &this
}

// GetMountPath returns the MountPath field value
func (o *ExtraVolumeMounts) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.mountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *ExtraVolumeMounts) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.mountPath, true
}

// MountPath sets field value
func (o *ExtraVolumeMounts) MountPath(v string) *ExtraVolumeMounts {
	o.mountPath = v
	return o
}

// GetName returns the Name field value
func (o *ExtraVolumeMounts) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtraVolumeMounts) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *ExtraVolumeMounts) Name(v string) *ExtraVolumeMounts {
	o.name = v
	return o
}

func (o ExtraVolumeMounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtraVolumeMounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mountPath"] = o.mountPath
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableExtraVolumeMounts struct {
	value *ExtraVolumeMounts
	isSet bool
}

func (v NullableExtraVolumeMounts) Get() *ExtraVolumeMounts {
	return v.value
}

func (v *NullableExtraVolumeMounts) Set(val *ExtraVolumeMounts) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraVolumeMounts) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraVolumeMounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraVolumeMounts(val *ExtraVolumeMounts) *NullableExtraVolumeMounts {
	return &NullableExtraVolumeMounts{value: val, isSet: true}
}

func (v NullableExtraVolumeMounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraVolumeMounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}