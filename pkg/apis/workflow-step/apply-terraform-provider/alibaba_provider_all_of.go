/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_provider

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the AlibabaProviderAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AlibabaProviderAllOf{}

// AlibabaProviderAllOf struct for AlibabaProviderAllOf
type AlibabaProviderAllOf struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewAlibabaProviderAllOfWith instantiates a new AlibabaProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlibabaProviderAllOfWith(name string, type_ string) *AlibabaProviderAllOf {
	this := AlibabaProviderAllOf{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewAlibabaProviderAllOf instantiates a new AlibabaProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlibabaProviderAllOf() *AlibabaProviderAllOf {
	this := AlibabaProviderAllOf{}
	var name string = "alibaba-provider"
	this.Name = name
	return &this
}

// GetName returns the Name field value
func (o *AlibabaProviderAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlibabaProviderAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlibabaProviderAllOf) SetName(v string) *AlibabaProviderAllOf {
	o.Name = v
	return o
}

// GetType returns the Type field value
func (o *AlibabaProviderAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlibabaProviderAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AlibabaProviderAllOf) SetType(v string) *AlibabaProviderAllOf {
	o.Type = v
	return o
}

func (o AlibabaProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlibabaProviderAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAlibabaProviderAllOf struct {
	value *AlibabaProviderAllOf
	isSet bool
}

func (v NullableAlibabaProviderAllOf) Get() *AlibabaProviderAllOf {
	return v.value
}

func (v *NullableAlibabaProviderAllOf) Set(val *AlibabaProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlibabaProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlibabaProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlibabaProviderAllOf(val *AlibabaProviderAllOf) *NullableAlibabaProviderAllOf {
	return &NullableAlibabaProviderAllOf{value: val, isSet: true}
}

func (v NullableAlibabaProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlibabaProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
