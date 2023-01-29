/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the Selector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Selector{}

// Selector struct for Selector
type Selector struct {
	matchExpressions *MatchExpressions  `json:"matchExpressions,omitempty"`
	matchLabels      *map[string]string `json:"matchLabels,omitempty"`
}

// NewSelectorWith instantiates a new Selector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectorWith() *Selector {
	this := Selector{}
	return &this
}

// NewSelector instantiates a new Selector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelector() *Selector {
	this := Selector{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *Selector) GetMatchExpressions() MatchExpressions {
	if o == nil || utils.IsNil(o.matchExpressions) {
		var ret MatchExpressions
		return ret
	}
	return *o.matchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selector) GetMatchExpressionsOk() (*MatchExpressions, bool) {
	if o == nil || utils.IsNil(o.matchExpressions) {
		return nil, false
	}
	return o.matchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *Selector) HasMatchExpressions() bool {
	if o != nil && !utils.IsNil(o.matchExpressions) {
		return true
	}

	return false
}

// MatchExpressions gets a reference to the given MatchExpressions and assigns it to the matchExpressions field.
// matchExpressions:
func (o *Selector) MatchExpressions(v MatchExpressions) *Selector {
	o.matchExpressions = &v
	return o
}

// GetMatchLabels returns the MatchLabels field value if set, zero value otherwise.
func (o *Selector) GetMatchLabels() map[string]string {
	if o == nil || utils.IsNil(o.matchLabels) {
		var ret map[string]string
		return ret
	}
	return *o.matchLabels
}

// GetMatchLabelsOk returns a tuple with the MatchLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selector) GetMatchLabelsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.matchLabels) {
		return nil, false
	}
	return o.matchLabels, true
}

// HasMatchLabels returns a boolean if a field has been set.
func (o *Selector) HasMatchLabels() bool {
	if o != nil && !utils.IsNil(o.matchLabels) {
		return true
	}

	return false
}

// MatchLabels gets a reference to the given map[string]string and assigns it to the matchLabels field.
// matchLabels:
func (o *Selector) MatchLabels(v map[string]string) *Selector {
	o.matchLabels = &v
	return o
}

func (o Selector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Selector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.matchExpressions) {
		toSerialize["matchExpressions"] = o.matchExpressions
	}
	if !utils.IsNil(o.matchLabels) {
		toSerialize["matchLabels"] = o.matchLabels
	}
	return toSerialize, nil
}

type NullableSelector struct {
	value *Selector
	isSet bool
}

func (v NullableSelector) Get() *Selector {
	return v.value
}

func (v *NullableSelector) Set(val *Selector) {
	v.value = val
	v.isSet = true
}

func (v NullableSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelector(val *Selector) *NullableSelector {
	return &NullableSelector{value: val, isSet: true}
}

func (v NullableSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}