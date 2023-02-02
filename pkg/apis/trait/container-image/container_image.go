/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_image

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"fmt"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// ContainerImageSpec - struct for ContainerImageSpec
type ContainerImageSpec struct {
	ContainerImageSpecOneOf *ContainerImageSpecOneOf
	PatchParams             *PatchParams
}

// ContainerImageSpecOneOfAsContainerImageSpec is a convenience function that returns ContainerImageSpecOneOf wrapped in ContainerImageSpec
func ContainerImageSpecOneOfAsContainerImageSpec(v *ContainerImageSpecOneOf) ContainerImageSpec {
	return ContainerImageSpec{
		ContainerImageSpecOneOf: v,
	}
}

// PatchParamsAsContainerImageSpec is a convenience function that returns PatchParams wrapped in ContainerImageSpec
func PatchParamsAsContainerImageSpec(v *PatchParams) ContainerImageSpec {
	return ContainerImageSpec{
		PatchParams: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContainerImageSpec) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContainerImageSpecOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.ContainerImageSpecOneOf)
	if err == nil {
		jsonContainerImageSpecOneOf, _ := json.Marshal(dst.ContainerImageSpecOneOf)
		if string(jsonContainerImageSpecOneOf) == "{}" { // empty struct
			dst.ContainerImageSpecOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ContainerImageSpecOneOf = nil
	}

	// try to unmarshal data into PatchParams
	err = utils.NewStrictDecoder(data).Decode(&dst.PatchParams)
	if err == nil {
		jsonPatchParams, _ := json.Marshal(dst.PatchParams)
		if string(jsonPatchParams) == "{}" { // empty struct
			dst.PatchParams = nil
		} else {
			match++
		}
	} else {
		dst.PatchParams = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContainerImageSpecOneOf = nil
		dst.PatchParams = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContainerImageSpec)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContainerImageSpec)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContainerImageSpec) MarshalJSON() ([]byte, error) {
	if src.ContainerImageSpecOneOf != nil {
		return json.Marshal(&src.ContainerImageSpecOneOf)
	}

	if src.PatchParams != nil {
		return json.Marshal(&src.PatchParams)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContainerImageSpec) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContainerImageSpecOneOf != nil {
		return obj.ContainerImageSpecOneOf
	}

	if obj.PatchParams != nil {
		return obj.PatchParams
	}

	// all schemas are nil
	return nil
}

type NullableContainerImageSpec struct {
	value *ContainerImageSpec
	isSet bool
}

func (v NullableContainerImageSpec) Get() *ContainerImageSpec {
	return v.value
}

func (v *NullableContainerImageSpec) Set(val *ContainerImageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerImageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerImageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerImageSpec(val *ContainerImageSpec) *NullableContainerImageSpec {
	return &NullableContainerImageSpec{value: val, isSet: true}
}

func (v NullableContainerImageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerImageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ContainerImageType = "container-image"

func init() {
	sdkcommon.RegisterTrait(ContainerImageType, FromTrait)
}

type ContainerImageTrait struct {
	Base       apis.TraitBase
	Properties ContainerImageSpec
}

func ContainerImage() *ContainerImageTrait {
	c := &ContainerImageTrait{Base: apis.TraitBase{}}
	return c
}

func (c *ContainerImageTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(c.Properties),
		Type:       ContainerImageType,
	}
	return res
}

func (c *ContainerImageTrait) FromTrait(from common.ApplicationTrait) (*ContainerImageTrait, error) {
	var properties ContainerImageSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	c.Properties = properties
	return c, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	c := &ContainerImageTrait{}
	return c.FromTrait(from)
}

func (c *ContainerImageTrait) Type() string {
	return ContainerImageType
}
