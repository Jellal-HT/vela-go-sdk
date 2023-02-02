/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Resources type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Resources{}

// Resources struct for Resources
type Resources struct {
	limits   *Limits   `json:"limits,omitempty"`
	requests *Requests `json:"requests,omitempty"`
}

// NewResourcesWith instantiates a new Resources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcesWith() *Resources {
	this := Resources{}
	return &this
}

// NewResources instantiates a new Resources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResources() *Resources {
	this := Resources{}
	return &this
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *Resources) GetLimits() Limits {
	if o == nil || utils.IsNil(o.limits) {
		var ret Limits
		return ret
	}
	return *o.limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resources) GetLimitsOk() (*Limits, bool) {
	if o == nil || utils.IsNil(o.limits) {
		return nil, false
	}
	return o.limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *Resources) HasLimits() bool {
	if o != nil && !utils.IsNil(o.limits) {
		return true
	}

	return false
}

// Limits gets a reference to the given Limits and assigns it to the limits field.
// limits:
func (o *Resources) Limits(v Limits) *Resources {
	o.limits = &v
	return o
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *Resources) GetRequests() Requests {
	if o == nil || utils.IsNil(o.requests) {
		var ret Requests
		return ret
	}
	return *o.requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resources) GetRequestsOk() (*Requests, bool) {
	if o == nil || utils.IsNil(o.requests) {
		return nil, false
	}
	return o.requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *Resources) HasRequests() bool {
	if o != nil && !utils.IsNil(o.requests) {
		return true
	}

	return false
}

// Requests gets a reference to the given Requests and assigns it to the requests field.
// requests:
func (o *Resources) Requests(v Requests) *Resources {
	o.requests = &v
	return o
}

func (o Resources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.limits) {
		toSerialize["limits"] = o.limits
	}
	if !utils.IsNil(o.requests) {
		toSerialize["requests"] = o.requests
	}
	return toSerialize, nil
}

type NullableResources struct {
	value *Resources
	isSet bool
}

func (v NullableResources) Get() *Resources {
	return v.value
}

func (v *NullableResources) Set(val *Resources) {
	v.value = val
	v.isSet = true
}

func (v NullableResources) IsSet() bool {
	return v.isSet
}

func (v *NullableResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResources(val *Resources) *NullableResources {
	return &NullableResources{value: val, isSet: true}
}

func (v NullableResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
