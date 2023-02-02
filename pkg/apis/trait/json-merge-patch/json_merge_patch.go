/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package json_merge_patch

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the JsonMergePatchSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &JsonMergePatchSpec{}

// JsonMergePatchSpec struct for JsonMergePatchSpec
type JsonMergePatchSpec struct {
}

// NewJsonMergePatchSpecWith instantiates a new JsonMergePatchSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMergePatchSpecWith() *JsonMergePatchSpec {
	this := JsonMergePatchSpec{}
	return &this
}

// NewJsonMergePatchSpec instantiates a new JsonMergePatchSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMergePatchSpec() *JsonMergePatchSpec {
	this := JsonMergePatchSpec{}
	return &this
}

func (o JsonMergePatchSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonMergePatchSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableJsonMergePatchSpec struct {
	value *JsonMergePatchSpec
	isSet bool
}

func (v NullableJsonMergePatchSpec) Get() *JsonMergePatchSpec {
	return v.value
}

func (v *NullableJsonMergePatchSpec) Set(val *JsonMergePatchSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMergePatchSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMergePatchSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMergePatchSpec(val *JsonMergePatchSpec) *NullableJsonMergePatchSpec {
	return &NullableJsonMergePatchSpec{value: val, isSet: true}
}

func (v NullableJsonMergePatchSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMergePatchSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const JsonMergePatchType = "json-merge-patch"

func init() {
	sdkcommon.RegisterTrait(JsonMergePatchType, FromTrait)
}

type JSONMergePatchTrait struct {
	Base       apis.TraitBase
	Properties JsonMergePatchSpec
}

func JsonMergePatch() *JSONMergePatchTrait {
	j := &JSONMergePatchTrait{Base: apis.TraitBase{}}
	return j
}

func (j *JSONMergePatchTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(j.Properties),
		Type:       JsonMergePatchType,
	}
	return res
}

func (j *JSONMergePatchTrait) FromTrait(from common.ApplicationTrait) (*JSONMergePatchTrait, error) {
	var properties JsonMergePatchSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	j.Base.Type = JsonMergePatchType
	j.Properties = properties
	return j, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	j := &JSONMergePatchTrait{}
	return j.FromTrait(from)
}

func (j *JSONMergePatchTrait) DefType() string {
	return JsonMergePatchType
}
