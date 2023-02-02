/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package worker

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Secret type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Secret{}

// Secret struct for Secret
type Secret struct {
	DefaultMode *int32  `json:"defaultMode,omitempty"`
	Items       []Items `json:"items,omitempty"`
	MountPath   *string `json:"mountPath,omitempty"`
	Name        *string `json:"name,omitempty"`
	SecretName  *string `json:"secretName,omitempty"`
}

// NewSecretWith instantiates a new Secret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretWith() *Secret {
	this := Secret{}
	var defaultMode int32 = 420
	this.DefaultMode = &defaultMode
	return &this
}

// NewSecret instantiates a new Secret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecret() *Secret {
	this := Secret{}
	var defaultMode int32 = 420
	this.DefaultMode = &defaultMode
	return &this
}

// GetDefaultMode returns the DefaultMode field value if set, zero value otherwise.
func (o *Secret) GetDefaultMode() int32 {
	if o == nil || utils.IsNil(o.DefaultMode) {
		var ret int32
		return ret
	}
	return *o.DefaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetDefaultModeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.DefaultMode) {
		return nil, false
	}
	return o.DefaultMode, true
}

// HasDefaultMode returns a boolean if a field has been set.
func (o *Secret) HasDefaultMode() bool {
	if o != nil && !utils.IsNil(o.DefaultMode) {
		return true
	}

	return false
}

// SetDefaultMode gets a reference to the given int32 and assigns it to the defaultMode field.
// DefaultMode:
func (o *Secret) SetDefaultMode(v int32) *Secret {
	o.DefaultMode = &v
	return o
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Secret) GetItems() []Items {
	if o == nil || utils.IsNil(o.Items) {
		var ret []Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetItemsOk() ([]Items, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Secret) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Items and assigns it to the items field.
// Items:
func (o *Secret) SetItems(v []Items) *Secret {
	o.Items = v
	return o
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *Secret) GetMountPath() string {
	if o == nil || utils.IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetMountPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *Secret) HasMountPath() bool {
	if o != nil && !utils.IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the mountPath field.
// MountPath:
func (o *Secret) SetMountPath(v string) *Secret {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Secret) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Secret) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *Secret) SetName(v string) *Secret {
	o.Name = &v
	return o
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *Secret) GetSecretName() string {
	if o == nil || utils.IsNil(o.SecretName) {
		var ret string
		return ret
	}
	return *o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetSecretNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SecretName) {
		return nil, false
	}
	return o.SecretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *Secret) HasSecretName() bool {
	if o != nil && !utils.IsNil(o.SecretName) {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the secretName field.
// SecretName:
func (o *Secret) SetSecretName(v string) *Secret {
	o.SecretName = &v
	return o
}

func (o Secret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Secret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DefaultMode) {
		toSerialize["defaultMode"] = o.DefaultMode
	}
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !utils.IsNil(o.MountPath) {
		toSerialize["mountPath"] = o.MountPath
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.SecretName) {
		toSerialize["secretName"] = o.SecretName
	}
	return toSerialize, nil
}

type NullableSecret struct {
	value *Secret
	isSet bool
}

func (v NullableSecret) Get() *Secret {
	return v.value
}

func (v *NullableSecret) Set(val *Secret) {
	v.value = val
	v.isSet = true
}

func (v NullableSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecret(val *Secret) *NullableSecret {
	return &NullableSecret{value: val, isSet: true}
}

func (v NullableSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
