/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the NodeAffinity type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NodeAffinity{}

// NodeAffinity Specify the node affinity scheduling rules for the pod
type NodeAffinity struct {
	// Specify the preferred during scheduling ignored during execution
	Preferred []Preferred `json:"preferred,omitempty"`
	Required  *Required   `json:"required,omitempty"`
}

// NewNodeAffinityWith instantiates a new NodeAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeAffinityWith() *NodeAffinity {
	this := NodeAffinity{}
	return &this
}

// NewNodeAffinity instantiates a new NodeAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeAffinity() *NodeAffinity {
	this := NodeAffinity{}
	return &this
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *NodeAffinity) GetPreferred() []Preferred {
	if o == nil || utils.IsNil(o.Preferred) {
		var ret []Preferred
		return ret
	}
	return o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeAffinity) GetPreferredOk() ([]Preferred, bool) {
	if o == nil || utils.IsNil(o.Preferred) {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *NodeAffinity) HasPreferred() bool {
	if o != nil && !utils.IsNil(o.Preferred) {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given []Preferred and assigns it to the preferred field.
// Preferred:  Specify the preferred during scheduling ignored during execution
func (o *NodeAffinity) SetPreferred(v []Preferred) *NodeAffinity {
	o.Preferred = v
	return o
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *NodeAffinity) GetRequired() Required {
	if o == nil || utils.IsNil(o.Required) {
		var ret Required
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeAffinity) GetRequiredOk() (*Required, bool) {
	if o == nil || utils.IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *NodeAffinity) HasRequired() bool {
	if o != nil && !utils.IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given Required and assigns it to the required field.
// Required:
func (o *NodeAffinity) SetRequired(v Required) *NodeAffinity {
	o.Required = &v
	return o
}

func (o NodeAffinity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Preferred) {
		toSerialize["preferred"] = o.Preferred
	}
	if !utils.IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableNodeAffinity struct {
	value *NodeAffinity
	isSet bool
}

func (v NullableNodeAffinity) Get() *NodeAffinity {
	return v.value
}

func (v *NullableNodeAffinity) Set(val *NodeAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeAffinity(val *NodeAffinity) *NullableNodeAffinity {
	return &NullableNodeAffinity{value: val, isSet: true}
}

func (v NullableNodeAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
