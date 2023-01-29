/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the PodAffinityTerm type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PodAffinityTerm{}

// PodAffinityTerm struct for PodAffinityTerm
type PodAffinityTerm struct {
	labelSelector     *LabelSelector `json:"labelSelector,omitempty"`
	namespaceSelector *LabelSelector `json:"namespaceSelector,omitempty"`
	namespaces        []string       `json:"namespaces,omitempty"`
	topologyKey       string         `json:"topologyKey"`
}

// NewPodAffinityTermWith instantiates a new PodAffinityTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodAffinityTermWith(topologyKey string) *PodAffinityTerm {
	this := PodAffinityTerm{}
	this.topologyKey = topologyKey
	return &this
}

// NewPodAffinityTerm instantiates a new PodAffinityTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodAffinityTerm() *PodAffinityTerm {
	this := PodAffinityTerm{}
	return &this
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *PodAffinityTerm) GetLabelSelector() LabelSelector {
	if o == nil || utils.IsNil(o.labelSelector) {
		var ret LabelSelector
		return ret
	}
	return *o.labelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAffinityTerm) GetLabelSelectorOk() (*LabelSelector, bool) {
	if o == nil || utils.IsNil(o.labelSelector) {
		return nil, false
	}
	return o.labelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *PodAffinityTerm) HasLabelSelector() bool {
	if o != nil && !utils.IsNil(o.labelSelector) {
		return true
	}

	return false
}

// LabelSelector gets a reference to the given LabelSelector and assigns it to the labelSelector field.
// labelSelector:
func (o *PodAffinityTerm) LabelSelector(v LabelSelector) *PodAffinityTerm {
	o.labelSelector = &v
	return o
}

// GetNamespaceSelector returns the NamespaceSelector field value if set, zero value otherwise.
func (o *PodAffinityTerm) GetNamespaceSelector() LabelSelector {
	if o == nil || utils.IsNil(o.namespaceSelector) {
		var ret LabelSelector
		return ret
	}
	return *o.namespaceSelector
}

// GetNamespaceSelectorOk returns a tuple with the NamespaceSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAffinityTerm) GetNamespaceSelectorOk() (*LabelSelector, bool) {
	if o == nil || utils.IsNil(o.namespaceSelector) {
		return nil, false
	}
	return o.namespaceSelector, true
}

// HasNamespaceSelector returns a boolean if a field has been set.
func (o *PodAffinityTerm) HasNamespaceSelector() bool {
	if o != nil && !utils.IsNil(o.namespaceSelector) {
		return true
	}

	return false
}

// NamespaceSelector gets a reference to the given LabelSelector and assigns it to the namespaceSelector field.
// namespaceSelector:
func (o *PodAffinityTerm) NamespaceSelector(v LabelSelector) *PodAffinityTerm {
	o.namespaceSelector = &v
	return o
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *PodAffinityTerm) GetNamespaces() []string {
	if o == nil || utils.IsNil(o.namespaces) {
		var ret []string
		return ret
	}
	return o.namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAffinityTerm) GetNamespacesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.namespaces) {
		return nil, false
	}
	return o.namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *PodAffinityTerm) HasNamespaces() bool {
	if o != nil && !utils.IsNil(o.namespaces) {
		return true
	}

	return false
}

// Namespaces gets a reference to the given []string and assigns it to the namespaces field.
// namespaces:
func (o *PodAffinityTerm) Namespaces(v []string) *PodAffinityTerm {
	o.namespaces = v
	return o
}

// GetTopologyKey returns the TopologyKey field value
func (o *PodAffinityTerm) GetTopologyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.topologyKey
}

// GetTopologyKeyOk returns a tuple with the TopologyKey field value
// and a boolean to check if the value has been set.
func (o *PodAffinityTerm) GetTopologyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.topologyKey, true
}

// TopologyKey sets field value
func (o *PodAffinityTerm) TopologyKey(v string) *PodAffinityTerm {
	o.topologyKey = v
	return o
}

func (o PodAffinityTerm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodAffinityTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.labelSelector) {
		toSerialize["labelSelector"] = o.labelSelector
	}
	if !utils.IsNil(o.namespaceSelector) {
		toSerialize["namespaceSelector"] = o.namespaceSelector
	}
	if !utils.IsNil(o.namespaces) {
		toSerialize["namespaces"] = o.namespaces
	}
	toSerialize["topologyKey"] = o.topologyKey
	return toSerialize, nil
}

type NullablePodAffinityTerm struct {
	value *PodAffinityTerm
	isSet bool
}

func (v NullablePodAffinityTerm) Get() *PodAffinityTerm {
	return v.value
}

func (v *NullablePodAffinityTerm) Set(val *PodAffinityTerm) {
	v.value = val
	v.isSet = true
}

func (v NullablePodAffinityTerm) IsSet() bool {
	return v.isSet
}

func (v *NullablePodAffinityTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodAffinityTerm(val *PodAffinityTerm) *NullablePodAffinityTerm {
	return &NullablePodAffinityTerm{value: val, isSet: true}
}

func (v NullablePodAffinityTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodAffinityTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}