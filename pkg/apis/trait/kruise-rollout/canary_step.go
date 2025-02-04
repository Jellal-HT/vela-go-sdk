/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kruise_rollout

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the CanaryStep type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CanaryStep{}

// CanaryStep struct for CanaryStep
type CanaryStep struct {
	// Define the behavior after release each step, if not filled, the default requires manual determination. If filled, it indicates the time to wait in seconds, e.g., 60
	Duration *int32    `json:"duration,omitempty"`
	Replicas *Replicas `json:"replicas,omitempty"`
	// Define the percentage of traffic routing to the new version in each step, e.g., 20%, 40%...
	Weight *int32 `json:"weight,omitempty"`
}

// NewCanaryStepWith instantiates a new CanaryStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanaryStepWith() *CanaryStep {
	this := CanaryStep{}
	return &this
}

// NewCanaryStep instantiates a new CanaryStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanaryStep() *CanaryStep {
	this := CanaryStep{}
	return &this
}

// NewCanarySteps converts a list CanaryStep pointers to objects.
// This is helpful when the SetCanaryStep requires a list of objects
func NewCanarySteps(ps ...*CanaryStep) []CanaryStep {
	objs := []CanaryStep{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *CanaryStep) GetDuration() int32 {
	if o == nil || utils.IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanaryStep) GetDurationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *CanaryStep) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the duration field.
// Duration:  Define the behavior after release each step, if not filled, the default requires manual determination. If filled, it indicates the time to wait in seconds, e.g., 60
func (o *CanaryStep) SetDuration(v int32) *CanaryStep {
	o.Duration = &v
	return o
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *CanaryStep) GetReplicas() Replicas {
	if o == nil || utils.IsNil(o.Replicas) {
		var ret Replicas
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanaryStep) GetReplicasOk() (*Replicas, bool) {
	if o == nil || utils.IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *CanaryStep) HasReplicas() bool {
	if o != nil && !utils.IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given Replicas and assigns it to the replicas field.
// Replicas:
func (o *CanaryStep) SetReplicas(v Replicas) *CanaryStep {
	o.Replicas = &v
	return o
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *CanaryStep) GetWeight() int32 {
	if o == nil || utils.IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanaryStep) GetWeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *CanaryStep) HasWeight() bool {
	if o != nil && !utils.IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the weight field.
// Weight:  Define the percentage of traffic routing to the new version in each step, e.g., 20%, 40%...
func (o *CanaryStep) SetWeight(v int32) *CanaryStep {
	o.Weight = &v
	return o
}

func (o CanaryStep) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CanaryStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !utils.IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	if !utils.IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableCanaryStep struct {
	value *CanaryStep
	isSet bool
}

func (v NullableCanaryStep) Get() *CanaryStep {
	return v.value
}

func (v *NullableCanaryStep) Set(val *CanaryStep) {
	v.value = val
	v.isSet = true
}

func (v NullableCanaryStep) IsSet() bool {
	return v.isSet
}

func (v *NullableCanaryStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanaryStep(val *CanaryStep) *NullableCanaryStep {
	return &NullableCanaryStep{value: val, isSet: true}
}

func (v NullableCanaryStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanaryStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
