/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s_update_strategy

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the Strategy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Strategy{}

// Strategy Specify the strategy of update
type Strategy struct {
	rollingStrategy *RollingStrategy `json:"rollingStrategy,omitempty"`
	// Specify the strategy type
	type_ string `json:"type"`
}

// NewStrategyWith instantiates a new Strategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStrategyWith(type_ string) *Strategy {
	this := Strategy{}
	this.type_ = type_
	return &this
}

// NewStrategy instantiates a new Strategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategy() *Strategy {
	this := Strategy{}
	var type_ string = "RollingUpdate"
	this.type_ = type_
	return &this
}

// GetRollingStrategy returns the RollingStrategy field value if set, zero value otherwise.
func (o *Strategy) GetRollingStrategy() RollingStrategy {
	if o == nil || utils.IsNil(o.rollingStrategy) {
		var ret RollingStrategy
		return ret
	}
	return *o.rollingStrategy
}

// GetRollingStrategyOk returns a tuple with the RollingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Strategy) GetRollingStrategyOk() (*RollingStrategy, bool) {
	if o == nil || utils.IsNil(o.rollingStrategy) {
		return nil, false
	}
	return o.rollingStrategy, true
}

// HasRollingStrategy returns a boolean if a field has been set.
func (o *Strategy) HasRollingStrategy() bool {
	if o != nil && !utils.IsNil(o.rollingStrategy) {
		return true
	}

	return false
}

// RollingStrategy gets a reference to the given RollingStrategy and assigns it to the rollingStrategy field.
// rollingStrategy:
func (o *Strategy) RollingStrategy(v RollingStrategy) *Strategy {
	o.rollingStrategy = &v
	return o
}

// GetType returns the Type field value
func (o *Strategy) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.type_
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Strategy) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.type_, true
}

// Type sets field value
func (o *Strategy) Type(v string) *Strategy {
	o.type_ = v
	return o
}

func (o Strategy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Strategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.rollingStrategy) {
		toSerialize["rollingStrategy"] = o.rollingStrategy
	}
	toSerialize["type"] = o.type_
	return toSerialize, nil
}

type NullableStrategy struct {
	value *Strategy
	isSet bool
}

func (v NullableStrategy) Get() *Strategy {
	return v.value
}

func (v *NullableStrategy) Set(val *Strategy) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategy(val *Strategy) *NullableStrategy {
	return &NullableStrategy{value: val, isSet: true}
}

func (v NullableStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}