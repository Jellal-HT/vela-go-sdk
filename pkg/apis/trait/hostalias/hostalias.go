/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hostalias

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HostaliasSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostaliasSpec{}

// HostaliasSpec struct for HostaliasSpec
type HostaliasSpec struct {
	// Specify the hostAliases to add
	hostAliases []HostAliases `json:"hostAliases"`
}

// NewHostaliasSpecWith instantiates a new HostaliasSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostaliasSpecWith(hostAliases []HostAliases) *HostaliasSpec {
	this := HostaliasSpec{}
	this.hostAliases = hostAliases
	return &this
}

// NewHostaliasSpec instantiates a new HostaliasSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostaliasSpec() *HostaliasSpec {
	this := HostaliasSpec{}
	return &this
}

// GetHostAliases returns the HostAliases field value
func (o *HostaliasTrait) GetHostAliases() []HostAliases {
	if o == nil {
		var ret []HostAliases
		return ret
	}

	return o.Properties.hostAliases
}

// GetHostAliasesOk returns a tuple with the HostAliases field value
// and a boolean to check if the value has been set.
func (o *HostaliasTrait) GetHostAliasesOk() ([]HostAliases, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.hostAliases, true
}

// HostAliases sets field value
func (o *HostaliasTrait) HostAliases(v []HostAliases) *HostaliasTrait {
	o.Properties.hostAliases = v
	return o
}

func (o HostaliasSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostaliasSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostAliases"] = o.hostAliases
	return toSerialize, nil
}

type NullableHostaliasSpec struct {
	value *HostaliasSpec
	isSet bool
}

func (v NullableHostaliasSpec) Get() *HostaliasSpec {
	return v.value
}

func (v *NullableHostaliasSpec) Set(val *HostaliasSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHostaliasSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHostaliasSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostaliasSpec(val *HostaliasSpec) *NullableHostaliasSpec {
	return &NullableHostaliasSpec{value: val, isSet: true}
}

func (v NullableHostaliasSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostaliasSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const HostaliasType = "hostalias"

func init() {
	sdkcommon.RegisterTrait(HostaliasType, FromTrait)
}

type HostaliasTrait struct {
	Base       apis.TraitBase
	Properties HostaliasSpec
}

func Hostalias() *HostaliasTrait {
	h := &HostaliasTrait{Base: apis.TraitBase{}}
	return h
}

func (h *HostaliasTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(h.Properties),
		Type:       HostaliasType,
	}
	return res
}

func (h *HostaliasTrait) FromTrait(from common.ApplicationTrait) (*HostaliasTrait, error) {
	var properties HostaliasSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	h.Properties = properties
	return h, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	h := &HostaliasTrait{}
	return h.FromTrait(from)
}

func (h *HostaliasTrait) Type() string {
	return HostaliasType
}
