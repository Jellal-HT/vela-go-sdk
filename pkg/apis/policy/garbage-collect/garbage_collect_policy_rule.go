/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garbage_collect

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the GarbageCollectPolicyRule type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GarbageCollectPolicyRule{}

// GarbageCollectPolicyRule struct for GarbageCollectPolicyRule
type GarbageCollectPolicyRule struct {
	// Specify how to select the targets of the rule
	Selector []ResourcePolicyRuleSelector `json:"selector,omitempty"`
	// Specify the strategy for target resource to recycle
	Strategy *string `json:"strategy,omitempty"`
}

// NewGarbageCollectPolicyRuleWith instantiates a new GarbageCollectPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGarbageCollectPolicyRuleWith() *GarbageCollectPolicyRule {
	this := GarbageCollectPolicyRule{}
	var strategy string = "onAppUpdate"
	this.Strategy = &strategy
	return &this
}

// NewGarbageCollectPolicyRule instantiates a new GarbageCollectPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGarbageCollectPolicyRule() *GarbageCollectPolicyRule {
	this := GarbageCollectPolicyRule{}
	var strategy string = "onAppUpdate"
	this.Strategy = &strategy
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *GarbageCollectPolicyRule) GetSelector() []ResourcePolicyRuleSelector {
	if o == nil || utils.IsNil(o.Selector) {
		var ret []ResourcePolicyRuleSelector
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicyRule) GetSelectorOk() ([]ResourcePolicyRuleSelector, bool) {
	if o == nil || utils.IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *GarbageCollectPolicyRule) HasSelector() bool {
	if o != nil && !utils.IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given []ResourcePolicyRuleSelector and assigns it to the selector field.
// Selector:  Specify how to select the targets of the rule
func (o *GarbageCollectPolicyRule) SetSelector(v []ResourcePolicyRuleSelector) *GarbageCollectPolicyRule {
	o.Selector = v
	return o
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *GarbageCollectPolicyRule) GetStrategy() string {
	if o == nil || utils.IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GarbageCollectPolicyRule) GetStrategyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *GarbageCollectPolicyRule) HasStrategy() bool {
	if o != nil && !utils.IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the strategy field.
// Strategy:  Specify the strategy for target resource to recycle
func (o *GarbageCollectPolicyRule) SetStrategy(v string) *GarbageCollectPolicyRule {
	o.Strategy = &v
	return o
}

func (o GarbageCollectPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GarbageCollectPolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !utils.IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	return toSerialize, nil
}

type NullableGarbageCollectPolicyRule struct {
	value *GarbageCollectPolicyRule
	isSet bool
}

func (v NullableGarbageCollectPolicyRule) Get() *GarbageCollectPolicyRule {
	return v.value
}

func (v *NullableGarbageCollectPolicyRule) Set(val *GarbageCollectPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGarbageCollectPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGarbageCollectPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGarbageCollectPolicyRule(val *GarbageCollectPolicyRule) *NullableGarbageCollectPolicyRule {
	return &NullableGarbageCollectPolicyRule{value: val, isSet: true}
}

func (v NullableGarbageCollectPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGarbageCollectPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
