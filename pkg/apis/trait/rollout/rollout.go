/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rollout

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the RolloutSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RolloutSpec{}

// RolloutSpec struct for RolloutSpec
type RolloutSpec struct {
	// Specify the batch partition in current deploying. It is used to control batch processes.
	BatchPartition *int32 `json:"batchPartition,omitempty"`
	// Specify the rollout batches, The total number of replicas of all batches needs to be equal to number of targetSize.
	RolloutBatches []RolloutBatch `json:"rolloutBatches,omitempty"`
	// Specify the target revision, it should be set if you want to rollback. such as: componentname-v1
	TargetRevision *string `json:"targetRevision,omitempty"`
	// Specify the count of replicas.
	TargetSize int32 `json:"targetSize"`
}

// NewRolloutSpecWith instantiates a new RolloutSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolloutSpecWith(targetSize int32) *RolloutSpec {
	this := RolloutSpec{}
	this.TargetSize = targetSize
	return &this
}

// NewRolloutSpec instantiates a new RolloutSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolloutSpec() *RolloutSpec {
	this := RolloutSpec{}
	return &this
}

// NewRolloutSpecs converts a list RolloutSpec pointers to objects.
// This is helpful when the SetRolloutSpec requires a list of objects
func NewRolloutSpecs(ps ...*RolloutSpec) []RolloutSpec {
	objs := []RolloutSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetBatchPartition returns the BatchPartition field value if set, zero value otherwise.
func (o *RolloutTrait) GetBatchPartition() int32 {
	if o == nil || utils.IsNil(o.Properties.BatchPartition) {
		var ret int32
		return ret
	}
	return *o.Properties.BatchPartition
}

// GetBatchPartitionOk returns a tuple with the BatchPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutTrait) GetBatchPartitionOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.BatchPartition) {
		return nil, false
	}
	return o.Properties.BatchPartition, true
}

// HasBatchPartition returns a boolean if a field has been set.
func (o *RolloutTrait) HasBatchPartition() bool {
	if o != nil && !utils.IsNil(o.Properties.BatchPartition) {
		return true
	}

	return false
}

// SetBatchPartition gets a reference to the given int32 and assigns it to the batchPartition field.
// BatchPartition:  Specify the batch partition in current deploying. It is used to control batch processes.
func (o *RolloutTrait) SetBatchPartition(v int32) *RolloutTrait {
	o.Properties.BatchPartition = &v
	return o
}

// GetRolloutBatches returns the RolloutBatches field value if set, zero value otherwise.
func (o *RolloutTrait) GetRolloutBatches() []RolloutBatch {
	if o == nil || utils.IsNil(o.Properties.RolloutBatches) {
		var ret []RolloutBatch
		return ret
	}
	return o.Properties.RolloutBatches
}

// GetRolloutBatchesOk returns a tuple with the RolloutBatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutTrait) GetRolloutBatchesOk() ([]RolloutBatch, bool) {
	if o == nil || utils.IsNil(o.Properties.RolloutBatches) {
		return nil, false
	}
	return o.Properties.RolloutBatches, true
}

// HasRolloutBatches returns a boolean if a field has been set.
func (o *RolloutTrait) HasRolloutBatches() bool {
	if o != nil && !utils.IsNil(o.Properties.RolloutBatches) {
		return true
	}

	return false
}

// SetRolloutBatches gets a reference to the given []RolloutBatch and assigns it to the rolloutBatches field.
// RolloutBatches:  Specify the rollout batches, The total number of replicas of all batches needs to be equal to number of targetSize.
func (o *RolloutTrait) SetRolloutBatches(v []RolloutBatch) *RolloutTrait {
	o.Properties.RolloutBatches = v
	return o
}

// GetTargetRevision returns the TargetRevision field value if set, zero value otherwise.
func (o *RolloutTrait) GetTargetRevision() string {
	if o == nil || utils.IsNil(o.Properties.TargetRevision) {
		var ret string
		return ret
	}
	return *o.Properties.TargetRevision
}

// GetTargetRevisionOk returns a tuple with the TargetRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutTrait) GetTargetRevisionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.TargetRevision) {
		return nil, false
	}
	return o.Properties.TargetRevision, true
}

// HasTargetRevision returns a boolean if a field has been set.
func (o *RolloutTrait) HasTargetRevision() bool {
	if o != nil && !utils.IsNil(o.Properties.TargetRevision) {
		return true
	}

	return false
}

// SetTargetRevision gets a reference to the given string and assigns it to the targetRevision field.
// TargetRevision:  Specify the target revision, it should be set if you want to rollback. such as: componentname-v1
func (o *RolloutTrait) SetTargetRevision(v string) *RolloutTrait {
	o.Properties.TargetRevision = &v
	return o
}

// GetTargetSize returns the TargetSize field value
func (o *RolloutTrait) GetTargetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Properties.TargetSize
}

// GetTargetSizeOk returns a tuple with the TargetSize field value
// and a boolean to check if the value has been set.
func (o *RolloutTrait) GetTargetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.TargetSize, true
}

// SetTargetSize sets field value
func (o *RolloutTrait) SetTargetSize(v int32) *RolloutTrait {
	o.Properties.TargetSize = v
	return o
}

func (o RolloutSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolloutSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BatchPartition) {
		toSerialize["batchPartition"] = o.BatchPartition
	}
	if !utils.IsNil(o.RolloutBatches) {
		toSerialize["rolloutBatches"] = o.RolloutBatches
	}
	if !utils.IsNil(o.TargetRevision) {
		toSerialize["targetRevision"] = o.TargetRevision
	}
	toSerialize["targetSize"] = o.TargetSize
	return toSerialize, nil
}

type NullableRolloutSpec struct {
	value *RolloutSpec
	isSet bool
}

func (v NullableRolloutSpec) Get() *RolloutSpec {
	return v.value
}

func (v *NullableRolloutSpec) Set(val *RolloutSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRolloutSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRolloutSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolloutSpec(val *RolloutSpec) *NullableRolloutSpec {
	return &NullableRolloutSpec{value: val, isSet: true}
}

func (v NullableRolloutSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolloutSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const RolloutType = "rollout"

func init() {
	sdkcommon.RegisterTrait(RolloutType, FromTrait)
}

type RolloutTrait struct {
	Base       apis.TraitBase
	Properties RolloutSpec
}

func Rollout() *RolloutTrait {
	r := &RolloutTrait{Base: apis.TraitBase{}}
	return r
}

func (r *RolloutTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(r.Properties),
		Type:       RolloutType,
	}
	return res
}

func (r *RolloutTrait) FromTrait(from common.ApplicationTrait) (*RolloutTrait, error) {
	var properties RolloutSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Type = RolloutType
	r.Properties = properties
	return r, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	r := &RolloutTrait{}
	return r.FromTrait(from)
}

func (r *RolloutTrait) DefType() string {
	return RolloutType
}