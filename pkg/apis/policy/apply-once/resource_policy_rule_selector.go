/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_once

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the ResourcePolicyRuleSelector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ResourcePolicyRuleSelector{}

// ResourcePolicyRuleSelector struct for ResourcePolicyRuleSelector
type ResourcePolicyRuleSelector struct {
	// Select resources by component names
	ComponentNames []string `json:"componentNames,omitempty"`
	// Select resources by component types
	ComponentTypes []string `json:"componentTypes,omitempty"`
	// Select resources by oamTypes (COMPONENT or TRAIT)
	OamTypes []string `json:"oamTypes,omitempty"`
	// Select resources by their names
	ResourceNames []string `json:"resourceNames,omitempty"`
	// Select resources by resource types (like Deployment)
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	// Select resources by trait types
	TraitTypes []string `json:"traitTypes,omitempty"`
}

// NewResourcePolicyRuleSelectorWith instantiates a new ResourcePolicyRuleSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePolicyRuleSelectorWith() *ResourcePolicyRuleSelector {
	this := ResourcePolicyRuleSelector{}
	return &this
}

// NewResourcePolicyRuleSelector instantiates a new ResourcePolicyRuleSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePolicyRuleSelector() *ResourcePolicyRuleSelector {
	this := ResourcePolicyRuleSelector{}
	return &this
}

// GetComponentNames returns the ComponentNames field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetComponentNames() []string {
	if o == nil || utils.IsNil(o.ComponentNames) {
		var ret []string
		return ret
	}
	return o.ComponentNames
}

// GetComponentNamesOk returns a tuple with the ComponentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetComponentNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ComponentNames) {
		return nil, false
	}
	return o.ComponentNames, true
}

// HasComponentNames returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasComponentNames() bool {
	if o != nil && !utils.IsNil(o.ComponentNames) {
		return true
	}

	return false
}

// SetComponentNames gets a reference to the given []string and assigns it to the componentNames field.
// ComponentNames:  Select resources by component names
func (o *ResourcePolicyRuleSelector) SetComponentNames(v []string) *ResourcePolicyRuleSelector {
	o.ComponentNames = v
	return o
}

// GetComponentTypes returns the ComponentTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetComponentTypes() []string {
	if o == nil || utils.IsNil(o.ComponentTypes) {
		var ret []string
		return ret
	}
	return o.ComponentTypes
}

// GetComponentTypesOk returns a tuple with the ComponentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetComponentTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ComponentTypes) {
		return nil, false
	}
	return o.ComponentTypes, true
}

// HasComponentTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasComponentTypes() bool {
	if o != nil && !utils.IsNil(o.ComponentTypes) {
		return true
	}

	return false
}

// SetComponentTypes gets a reference to the given []string and assigns it to the componentTypes field.
// ComponentTypes:  Select resources by component types
func (o *ResourcePolicyRuleSelector) SetComponentTypes(v []string) *ResourcePolicyRuleSelector {
	o.ComponentTypes = v
	return o
}

// GetOamTypes returns the OamTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetOamTypes() []string {
	if o == nil || utils.IsNil(o.OamTypes) {
		var ret []string
		return ret
	}
	return o.OamTypes
}

// GetOamTypesOk returns a tuple with the OamTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetOamTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OamTypes) {
		return nil, false
	}
	return o.OamTypes, true
}

// HasOamTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasOamTypes() bool {
	if o != nil && !utils.IsNil(o.OamTypes) {
		return true
	}

	return false
}

// SetOamTypes gets a reference to the given []string and assigns it to the oamTypes field.
// OamTypes:  Select resources by oamTypes (COMPONENT or TRAIT)
func (o *ResourcePolicyRuleSelector) SetOamTypes(v []string) *ResourcePolicyRuleSelector {
	o.OamTypes = v
	return o
}

// GetResourceNames returns the ResourceNames field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetResourceNames() []string {
	if o == nil || utils.IsNil(o.ResourceNames) {
		var ret []string
		return ret
	}
	return o.ResourceNames
}

// GetResourceNamesOk returns a tuple with the ResourceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetResourceNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ResourceNames) {
		return nil, false
	}
	return o.ResourceNames, true
}

// HasResourceNames returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasResourceNames() bool {
	if o != nil && !utils.IsNil(o.ResourceNames) {
		return true
	}

	return false
}

// SetResourceNames gets a reference to the given []string and assigns it to the resourceNames field.
// ResourceNames:  Select resources by their names
func (o *ResourcePolicyRuleSelector) SetResourceNames(v []string) *ResourcePolicyRuleSelector {
	o.ResourceNames = v
	return o
}

// GetResourceTypes returns the ResourceTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetResourceTypes() []string {
	if o == nil || utils.IsNil(o.ResourceTypes) {
		var ret []string
		return ret
	}
	return o.ResourceTypes
}

// GetResourceTypesOk returns a tuple with the ResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetResourceTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ResourceTypes) {
		return nil, false
	}
	return o.ResourceTypes, true
}

// HasResourceTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasResourceTypes() bool {
	if o != nil && !utils.IsNil(o.ResourceTypes) {
		return true
	}

	return false
}

// SetResourceTypes gets a reference to the given []string and assigns it to the resourceTypes field.
// ResourceTypes:  Select resources by resource types (like Deployment)
func (o *ResourcePolicyRuleSelector) SetResourceTypes(v []string) *ResourcePolicyRuleSelector {
	o.ResourceTypes = v
	return o
}

// GetTraitTypes returns the TraitTypes field value if set, zero value otherwise.
func (o *ResourcePolicyRuleSelector) GetTraitTypes() []string {
	if o == nil || utils.IsNil(o.TraitTypes) {
		var ret []string
		return ret
	}
	return o.TraitTypes
}

// GetTraitTypesOk returns a tuple with the TraitTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicyRuleSelector) GetTraitTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.TraitTypes) {
		return nil, false
	}
	return o.TraitTypes, true
}

// HasTraitTypes returns a boolean if a field has been set.
func (o *ResourcePolicyRuleSelector) HasTraitTypes() bool {
	if o != nil && !utils.IsNil(o.TraitTypes) {
		return true
	}

	return false
}

// SetTraitTypes gets a reference to the given []string and assigns it to the traitTypes field.
// TraitTypes:  Select resources by trait types
func (o *ResourcePolicyRuleSelector) SetTraitTypes(v []string) *ResourcePolicyRuleSelector {
	o.TraitTypes = v
	return o
}

func (o ResourcePolicyRuleSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcePolicyRuleSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ComponentNames) {
		toSerialize["componentNames"] = o.ComponentNames
	}
	if !utils.IsNil(o.ComponentTypes) {
		toSerialize["componentTypes"] = o.ComponentTypes
	}
	if !utils.IsNil(o.OamTypes) {
		toSerialize["oamTypes"] = o.OamTypes
	}
	if !utils.IsNil(o.ResourceNames) {
		toSerialize["resourceNames"] = o.ResourceNames
	}
	if !utils.IsNil(o.ResourceTypes) {
		toSerialize["resourceTypes"] = o.ResourceTypes
	}
	if !utils.IsNil(o.TraitTypes) {
		toSerialize["traitTypes"] = o.TraitTypes
	}
	return toSerialize, nil
}

type NullableResourcePolicyRuleSelector struct {
	value *ResourcePolicyRuleSelector
	isSet bool
}

func (v NullableResourcePolicyRuleSelector) Get() *ResourcePolicyRuleSelector {
	return v.value
}

func (v *NullableResourcePolicyRuleSelector) Set(val *ResourcePolicyRuleSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePolicyRuleSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePolicyRuleSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePolicyRuleSelector(val *ResourcePolicyRuleSelector) *NullableResourcePolicyRuleSelector {
	return &NullableResourcePolicyRuleSelector{value: val, isSet: true}
}

func (v NullableResourcePolicyRuleSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePolicyRuleSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
