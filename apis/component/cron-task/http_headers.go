/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the HttpHeaders type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpHeaders{}

// HttpHeaders struct for HttpHeaders
type HttpHeaders struct {
	name  string `json:"name"`
	value string `json:"value"`
}

// NewHttpHeadersWith instantiates a new HttpHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpHeadersWith(name string, value string) *HttpHeaders {
	this := HttpHeaders{}
	this.name = name
	this.value = value
	return &this
}

// NewHttpHeaders instantiates a new HttpHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHeaders() *HttpHeaders {
	this := HttpHeaders{}
	return &this
}

// GetName returns the Name field value
func (o *HttpHeaders) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HttpHeaders) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.name, true
}

// Name sets field value
func (o *HttpHeaders) Name(v string) *HttpHeaders {
	o.name = v
	return o
}

// GetValue returns the Value field value
func (o *HttpHeaders) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HttpHeaders) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.value, true
}

// Value sets field value
func (o *HttpHeaders) Value(v string) *HttpHeaders {
	o.value = v
	return o
}

func (o HttpHeaders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpHeaders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.name
	toSerialize["value"] = o.value
	return toSerialize, nil
}

type NullableHttpHeaders struct {
	value *HttpHeaders
	isSet bool
}

func (v NullableHttpHeaders) Get() *HttpHeaders {
	return v.value
}

func (v *NullableHttpHeaders) Set(val *HttpHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpHeaders(val *HttpHeaders) *NullableHttpHeaders {
	return &NullableHttpHeaders{value: val, isSet: true}
}

func (v NullableHttpHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}