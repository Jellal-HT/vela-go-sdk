/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package override

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the OverrideSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OverrideSpec{}

// OverrideSpec struct for OverrideSpec
type OverrideSpec struct {
	// Specify the overridden component configuration.
	Components []PatchParams `json:"components,omitempty"`
	// Specify a list of component names to use, if empty, all components will be selected.
	Selector []string `json:"selector,omitempty"`
}

// NewOverrideSpecWith instantiates a new OverrideSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideSpecWith() *OverrideSpec {
	this := OverrideSpec{}
	return &this
}

// NewOverrideSpec instantiates a new OverrideSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideSpec() *OverrideSpec {
	this := OverrideSpec{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *OverridePolicy) GetComponents() []PatchParams {
	if o == nil || utils.IsNil(o.Properties.Components) {
		var ret []PatchParams
		return ret
	}
	return o.Properties.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverridePolicy) GetComponentsOk() ([]PatchParams, bool) {
	if o == nil || utils.IsNil(o.Properties.Components) {
		return nil, false
	}
	return o.Properties.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *OverridePolicy) HasComponents() bool {
	if o != nil && !utils.IsNil(o.Properties.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []PatchParams and assigns it to the components field.
// Components:  Specify the overridden component configuration.
func (o *OverridePolicy) SetComponents(v []PatchParams) *OverridePolicy {
	o.Properties.Components = v
	return o
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *OverridePolicy) GetSelector() []string {
	if o == nil || utils.IsNil(o.Properties.Selector) {
		var ret []string
		return ret
	}
	return o.Properties.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverridePolicy) GetSelectorOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Selector) {
		return nil, false
	}
	return o.Properties.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *OverridePolicy) HasSelector() bool {
	if o != nil && !utils.IsNil(o.Properties.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given []string and assigns it to the selector field.
// Selector:  Specify a list of component names to use, if empty, all components will be selected.
func (o *OverridePolicy) SetSelector(v []string) *OverridePolicy {
	o.Properties.Selector = v
	return o
}

func (o OverrideSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !utils.IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	return toSerialize, nil
}

type NullableOverrideSpec struct {
	value *OverrideSpec
	isSet bool
}

func (v NullableOverrideSpec) Get() *OverrideSpec {
	return v.value
}

func (v *NullableOverrideSpec) Set(val *OverrideSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideSpec(val *OverrideSpec) *NullableOverrideSpec {
	return &NullableOverrideSpec{value: val, isSet: true}
}

func (v NullableOverrideSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const OverrideType = "override"

func init() {
	sdkcommon.RegisterPolicy(OverrideType, FromPolicy)
}

type OverridePolicy struct {
	Base       apis.PolicyBase
	Properties OverrideSpec
}

func Override(name string) *OverridePolicy {
	o := &OverridePolicy{Base: apis.PolicyBase{
		Name: name,
		Type: OverrideType,
	}}
	return o
}

func (o *OverridePolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       o.Base.Name,
		Properties: util.Object2RawExtension(o.Properties),
		Type:       OverrideType,
	}
	return res
}

func (o *OverridePolicy) FromPolicy(from v1beta1.AppPolicy) (*OverridePolicy, error) {
	var properties OverrideSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	o.Base.Name = from.Name
	o.Base.Type = OverrideType
	o.Properties = properties
	return o, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	o := &OverridePolicy{}
	return o.FromPolicy(from)
}

func (o *OverridePolicy) DefType() string {
	return OverrideType
}
