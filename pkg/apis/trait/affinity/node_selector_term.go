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

// checks if the NodeSelectorTerm type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NodeSelectorTerm{}

// NodeSelectorTerm struct for NodeSelectorTerm
type NodeSelectorTerm struct {
	MatchExpressions []NodeSelecor `json:"matchExpressions,omitempty"`
	MatchFields      []NodeSelecor `json:"matchFields,omitempty"`
}

// NewNodeSelectorTermWith instantiates a new NodeSelectorTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeSelectorTermWith() *NodeSelectorTerm {
	this := NodeSelectorTerm{}
	return &this
}

// NewNodeSelectorTerm instantiates a new NodeSelectorTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeSelectorTerm() *NodeSelectorTerm {
	this := NodeSelectorTerm{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *NodeSelectorTerm) GetMatchExpressions() []NodeSelecor {
	if o == nil || utils.IsNil(o.MatchExpressions) {
		var ret []NodeSelecor
		return ret
	}
	return o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeSelectorTerm) GetMatchExpressionsOk() ([]NodeSelecor, bool) {
	if o == nil || utils.IsNil(o.MatchExpressions) {
		return nil, false
	}
	return o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *NodeSelectorTerm) HasMatchExpressions() bool {
	if o != nil && !utils.IsNil(o.MatchExpressions) {
		return true
	}

	return false
}

// SetMatchExpressions gets a reference to the given []NodeSelecor and assigns it to the matchExpressions field.
// MatchExpressions:
func (o *NodeSelectorTerm) SetMatchExpressions(v []NodeSelecor) *NodeSelectorTerm {
	o.MatchExpressions = v
	return o
}

// GetMatchFields returns the MatchFields field value if set, zero value otherwise.
func (o *NodeSelectorTerm) GetMatchFields() []NodeSelecor {
	if o == nil || utils.IsNil(o.MatchFields) {
		var ret []NodeSelecor
		return ret
	}
	return o.MatchFields
}

// GetMatchFieldsOk returns a tuple with the MatchFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeSelectorTerm) GetMatchFieldsOk() ([]NodeSelecor, bool) {
	if o == nil || utils.IsNil(o.MatchFields) {
		return nil, false
	}
	return o.MatchFields, true
}

// HasMatchFields returns a boolean if a field has been set.
func (o *NodeSelectorTerm) HasMatchFields() bool {
	if o != nil && !utils.IsNil(o.MatchFields) {
		return true
	}

	return false
}

// SetMatchFields gets a reference to the given []NodeSelecor and assigns it to the matchFields field.
// MatchFields:
func (o *NodeSelectorTerm) SetMatchFields(v []NodeSelecor) *NodeSelectorTerm {
	o.MatchFields = v
	return o
}

func (o NodeSelectorTerm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeSelectorTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MatchExpressions) {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if !utils.IsNil(o.MatchFields) {
		toSerialize["matchFields"] = o.MatchFields
	}
	return toSerialize, nil
}

type NullableNodeSelectorTerm struct {
	value *NodeSelectorTerm
	isSet bool
}

func (v NullableNodeSelectorTerm) Get() *NodeSelectorTerm {
	return v.value
}

func (v *NullableNodeSelectorTerm) Set(val *NodeSelectorTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeSelectorTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeSelectorTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeSelectorTerm(val *NodeSelectorTerm) *NullableNodeSelectorTerm {
	return &NullableNodeSelectorTerm{value: val, isSet: true}
}

func (v NullableNodeSelectorTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeSelectorTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
