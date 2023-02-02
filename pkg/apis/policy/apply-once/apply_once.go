/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_once

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyOnceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyOnceSpec{}

// ApplyOnceSpec struct for ApplyOnceSpec
type ApplyOnceSpec struct {
	// Whether to enable apply-once for the whole application
	enable *bool `json:"enable,omitempty"`
	// Specify the rules for configuring apply-once policy in resource level
	rules []ApplyOncePolicyRule `json:"rules,omitempty"`
}

// NewApplyOnceSpecWith instantiates a new ApplyOnceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyOnceSpecWith() *ApplyOnceSpec {
	this := ApplyOnceSpec{}
	var enable bool = false
	this.enable = &enable
	return &this
}

// NewApplyOnceSpec instantiates a new ApplyOnceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyOnceSpec() *ApplyOnceSpec {
	this := ApplyOnceSpec{}
	var enable bool = false
	this.enable = &enable
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ApplyOncePolicy) GetEnable() bool {
	if o == nil || utils.IsNil(o.Properties.enable) {
		var ret bool
		return ret
	}
	return *o.Properties.enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOncePolicy) GetEnableOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.enable) {
		return nil, false
	}
	return o.Properties.enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ApplyOncePolicy) HasEnable() bool {
	if o != nil && !utils.IsNil(o.Properties.enable) {
		return true
	}

	return false
}

// Enable gets a reference to the given bool and assigns it to the enable field.
// enable:  Whether to enable apply-once for the whole application
func (o *ApplyOncePolicy) Enable(v bool) *ApplyOncePolicy {
	o.Properties.enable = &v
	return o
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ApplyOncePolicy) GetRules() []ApplyOncePolicyRule {
	if o == nil || utils.IsNil(o.Properties.rules) {
		var ret []ApplyOncePolicyRule
		return ret
	}
	return o.Properties.rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOncePolicy) GetRulesOk() ([]ApplyOncePolicyRule, bool) {
	if o == nil || utils.IsNil(o.Properties.rules) {
		return nil, false
	}
	return o.Properties.rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ApplyOncePolicy) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.rules) {
		return true
	}

	return false
}

// Rules gets a reference to the given []ApplyOncePolicyRule and assigns it to the rules field.
// rules:  Specify the rules for configuring apply-once policy in resource level
func (o *ApplyOncePolicy) Rules(v []ApplyOncePolicyRule) *ApplyOncePolicy {
	o.Properties.rules = v
	return o
}

func (o ApplyOnceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyOnceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.enable) {
		toSerialize["enable"] = o.enable
	}
	if !utils.IsNil(o.rules) {
		toSerialize["rules"] = o.rules
	}
	return toSerialize, nil
}

type NullableApplyOnceSpec struct {
	value *ApplyOnceSpec
	isSet bool
}

func (v NullableApplyOnceSpec) Get() *ApplyOnceSpec {
	return v.value
}

func (v *NullableApplyOnceSpec) Set(val *ApplyOnceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyOnceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyOnceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyOnceSpec(val *ApplyOnceSpec) *NullableApplyOnceSpec {
	return &NullableApplyOnceSpec{value: val, isSet: true}
}

func (v NullableApplyOnceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyOnceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyOnceType = "apply-once"

func init() {
	sdkcommon.RegisterPolicy(ApplyOnceType, FromPolicy)
}

type ApplyOncePolicy struct {
	Base       apis.PolicyBase
	Properties ApplyOnceSpec
}

func ApplyOnce(name string) *ApplyOncePolicy {
	a := &ApplyOncePolicy{Base: apis.PolicyBase{
		Name: name,
		Type: ApplyOnceType,
	}}
	return a
}

func (a *ApplyOncePolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       a.Base.Name,
		Properties: util.Object2RawExtension(a.Properties),
		Type:       ApplyOnceType,
	}
	return res
}

func (a *ApplyOncePolicy) FromPolicy(from v1beta1.AppPolicy) (*ApplyOncePolicy, error) {
	var properties ApplyOnceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	a.Base.Name = from.Name
	a.Properties = properties
	return a, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	a := &ApplyOncePolicy{}
	return a.FromPolicy(from)
}

func (a *ApplyOncePolicy) DefType() string {
	return ApplyOnceType
}
