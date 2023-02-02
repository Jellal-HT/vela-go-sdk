/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package override

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Traits type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Traits{}

// Traits struct for Traits
type Traits struct {
	// Specify if the trait should be remove, default false
	Disable *bool `json:"disable,omitempty"`
	// Specify the properties to override.
	Properties map[string]interface{} `json:"properties,omitempty"`
	// Specify the type of the trait to be patched.
	Type *string `json:"type,omitempty"`
}

// NewTraitsWith instantiates a new Traits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraitsWith() *Traits {
	this := Traits{}
	var disable bool = false
	this.Disable = &disable
	return &this
}

// NewTraits instantiates a new Traits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraits() *Traits {
	this := Traits{}
	var disable bool = false
	this.Disable = &disable
	return &this
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *Traits) GetDisable() bool {
	if o == nil || utils.IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Traits) GetDisableOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *Traits) HasDisable() bool {
	if o != nil && !utils.IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the disable field.
// Disable:  Specify if the trait should be remove, default false
func (o *Traits) SetDisable(v bool) *Traits {
	o.Disable = &v
	return o
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Traits) GetProperties() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Traits) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Traits) HasProperties() bool {
	if o != nil && !utils.IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the properties field.
// Properties:  Specify the properties to override.
func (o *Traits) SetProperties(v map[string]interface{}) *Traits {
	o.Properties = v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Traits) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Traits) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Traits) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the type_ field.
// Type:  Specify the type of the trait to be patched.
func (o *Traits) SetType(v string) *Traits {
	o.Type = &v
	return o
}

func (o Traits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Traits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !utils.IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTraits struct {
	value *Traits
	isSet bool
}

func (v NullableTraits) Get() *Traits {
	return v.value
}

func (v *NullableTraits) Set(val *Traits) {
	v.value = val
	v.isSet = true
}

func (v NullableTraits) IsSet() bool {
	return v.isSet
}

func (v *NullableTraits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraits(val *Traits) *NullableTraits {
	return &NullableTraits{value: val, isSet: true}
}

func (v NullableTraits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
