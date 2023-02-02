/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package read_only

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the RuleSelector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RuleSelector{}

// RuleSelector struct for RuleSelector
type RuleSelector struct {
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

// NewRuleSelectorWith instantiates a new RuleSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSelectorWith() *RuleSelector {
	this := RuleSelector{}
	return &this
}

// NewRuleSelector instantiates a new RuleSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSelector() *RuleSelector {
	this := RuleSelector{}
	return &this
}

// GetComponentNames returns the ComponentNames field value if set, zero value otherwise.
func (o *RuleSelector) GetComponentNames() []string {
	if o == nil || utils.IsNil(o.ComponentNames) {
		var ret []string
		return ret
	}
	return o.ComponentNames
}

// GetComponentNamesOk returns a tuple with the ComponentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSelector) GetComponentNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ComponentNames) {
		return nil, false
	}
	return o.ComponentNames, true
}

// HasComponentNames returns a boolean if a field has been set.
func (o *RuleSelector) HasComponentNames() bool {
	if o != nil && !utils.IsNil(o.ComponentNames) {
		return true
	}

	return false
}

// SetComponentNames gets a reference to the given []string and assigns it to the componentNames field.
// ComponentNames:  Select resources by component names
func (o *RuleSelector) SetComponentNames(v []string) *RuleSelector {
	o.ComponentNames = v
	return o
}

// GetComponentTypes returns the ComponentTypes field value if set, zero value otherwise.
func (o *RuleSelector) GetComponentTypes() []string {
	if o == nil || utils.IsNil(o.ComponentTypes) {
		var ret []string
		return ret
	}
	return o.ComponentTypes
}

// GetComponentTypesOk returns a tuple with the ComponentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSelector) GetComponentTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ComponentTypes) {
		return nil, false
	}
	return o.ComponentTypes, true
}

// HasComponentTypes returns a boolean if a field has been set.
func (o *RuleSelector) HasComponentTypes() bool {
	if o != nil && !utils.IsNil(o.ComponentTypes) {
		return true
	}

	return false
}

// SetComponentTypes gets a reference to the given []string and assigns it to the componentTypes field.
// ComponentTypes:  Select resources by component types
func (o *RuleSelector) SetComponentTypes(v []string) *RuleSelector {
	o.ComponentTypes = v
	return o
}

// GetOamTypes returns the OamTypes field value if set, zero value otherwise.
func (o *RuleSelector) GetOamTypes() []string {
	if o == nil || utils.IsNil(o.OamTypes) {
		var ret []string
		return ret
	}
	return o.OamTypes
}

// GetOamTypesOk returns a tuple with the OamTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSelector) GetOamTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OamTypes) {
		return nil, false
	}
	return o.OamTypes, true
}

// HasOamTypes returns a boolean if a field has been set.
func (o *RuleSelector) HasOamTypes() bool {
	if o != nil && !utils.IsNil(o.OamTypes) {
		return true
	}

	return false
}

// SetOamTypes gets a reference to the given []string and assigns it to the oamTypes field.
// OamTypes:  Select resources by oamTypes (COMPONENT or TRAIT)
func (o *RuleSelector) SetOamTypes(v []string) *RuleSelector {
	o.OamTypes = v
	return o
}

// GetResourceNames returns the ResourceNames field value if set, zero value otherwise.
func (o *RuleSelector) GetResourceNames() []string {
	if o == nil || utils.IsNil(o.ResourceNames) {
		var ret []string
		return ret
	}
	return o.ResourceNames
}

// GetResourceNamesOk returns a tuple with the ResourceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSelector) GetResourceNamesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ResourceNames) {
		return nil, false
	}
	return o.ResourceNames, true
}

// HasResourceNames returns a boolean if a field has been set.
func (o *RuleSelector) HasResourceNames() bool {
	if o != nil && !utils.IsNil(o.ResourceNames) {
		return true
	}

	return false
}

// SetResourceNames gets a reference to the given []string and assigns it to the resourceNames field.
// ResourceNames:  Select resources by their names
func (o *RuleSelector) SetResourceNames(v []string) *RuleSelector {
	o.ResourceNames = v
	return o
}

// GetResourceTypes returns the ResourceTypes field value if set, zero value otherwise.
func (o *RuleSelector) GetResourceTypes() []string {
	if o == nil || utils.IsNil(o.ResourceTypes) {
		var ret []string
		return ret
	}
	return o.ResourceTypes
}

// GetResourceTypesOk returns a tuple with the ResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSelector) GetResourceTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ResourceTypes) {
		return nil, false
	}
	return o.ResourceTypes, true
}

// HasResourceTypes returns a boolean if a field has been set.
func (o *RuleSelector) HasResourceTypes() bool {
	if o != nil && !utils.IsNil(o.ResourceTypes) {
		return true
	}

	return false
}

// SetResourceTypes gets a reference to the given []string and assigns it to the resourceTypes field.
// ResourceTypes:  Select resources by resource types (like Deployment)
func (o *RuleSelector) SetResourceTypes(v []string) *RuleSelector {
	o.ResourceTypes = v
	return o
}

// GetTraitTypes returns the TraitTypes field value if set, zero value otherwise.
func (o *RuleSelector) GetTraitTypes() []string {
	if o == nil || utils.IsNil(o.TraitTypes) {
		var ret []string
		return ret
	}
	return o.TraitTypes
}

// GetTraitTypesOk returns a tuple with the TraitTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSelector) GetTraitTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.TraitTypes) {
		return nil, false
	}
	return o.TraitTypes, true
}

// HasTraitTypes returns a boolean if a field has been set.
func (o *RuleSelector) HasTraitTypes() bool {
	if o != nil && !utils.IsNil(o.TraitTypes) {
		return true
	}

	return false
}

// SetTraitTypes gets a reference to the given []string and assigns it to the traitTypes field.
// TraitTypes:  Select resources by trait types
func (o *RuleSelector) SetTraitTypes(v []string) *RuleSelector {
	o.TraitTypes = v
	return o
}

func (o RuleSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleSelector) ToMap() (map[string]interface{}, error) {
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

type NullableRuleSelector struct {
	value *RuleSelector
	isSet bool
}

func (v NullableRuleSelector) Get() *RuleSelector {
	return v.value
}

func (v *NullableRuleSelector) Set(val *RuleSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSelector(val *RuleSelector) *NullableRuleSelector {
	return &NullableRuleSelector{value: val, isSet: true}
}

func (v NullableRuleSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
