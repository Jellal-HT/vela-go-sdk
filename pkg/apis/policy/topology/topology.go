/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package topology

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the TopologySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TopologySpec{}

// TopologySpec struct for TopologySpec
type TopologySpec struct {
	// Ignore empty cluster error
	allowEmpty *bool `json:"allowEmpty,omitempty"`
	// Specify the label selector for clusters
	clusterLabelSelector *map[string]string `json:"clusterLabelSelector,omitempty"`
	// Deprecated: Use clusterLabelSelector instead.
	clusterSelector *map[string]string `json:"clusterSelector,omitempty"`
	// Specify the names of the clusters to select.
	clusters []string `json:"clusters,omitempty"`
	// Specify the target namespace to deploy in the selected clusters, default inherit the original namespace.
	namespace *string `json:"namespace,omitempty"`
}

// NewTopologySpecWith instantiates a new TopologySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologySpecWith() *TopologySpec {
	this := TopologySpec{}
	return &this
}

// NewTopologySpec instantiates a new TopologySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologySpec() *TopologySpec {
	this := TopologySpec{}
	return &this
}

// GetAllowEmpty returns the AllowEmpty field value if set, zero value otherwise.
func (o *TopologyPolicy) GetAllowEmpty() bool {
	if o == nil || utils.IsNil(o.Properties.allowEmpty) {
		var ret bool
		return ret
	}
	return *o.Properties.allowEmpty
}

// GetAllowEmptyOk returns a tuple with the AllowEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetAllowEmptyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.allowEmpty) {
		return nil, false
	}
	return o.Properties.allowEmpty, true
}

// HasAllowEmpty returns a boolean if a field has been set.
func (o *TopologyPolicy) HasAllowEmpty() bool {
	if o != nil && !utils.IsNil(o.Properties.allowEmpty) {
		return true
	}

	return false
}

// AllowEmpty gets a reference to the given bool and assigns it to the allowEmpty field.
// allowEmpty:  Ignore empty cluster error
func (o *TopologyPolicy) AllowEmpty(v bool) *TopologyPolicy {
	o.Properties.allowEmpty = &v
	return o
}

// GetClusterLabelSelector returns the ClusterLabelSelector field value if set, zero value otherwise.
func (o *TopologyPolicy) GetClusterLabelSelector() map[string]string {
	if o == nil || utils.IsNil(o.Properties.clusterLabelSelector) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.clusterLabelSelector
}

// GetClusterLabelSelectorOk returns a tuple with the ClusterLabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetClusterLabelSelectorOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.clusterLabelSelector) {
		return nil, false
	}
	return o.Properties.clusterLabelSelector, true
}

// HasClusterLabelSelector returns a boolean if a field has been set.
func (o *TopologyPolicy) HasClusterLabelSelector() bool {
	if o != nil && !utils.IsNil(o.Properties.clusterLabelSelector) {
		return true
	}

	return false
}

// ClusterLabelSelector gets a reference to the given map[string]string and assigns it to the clusterLabelSelector field.
// clusterLabelSelector:  Specify the label selector for clusters
func (o *TopologyPolicy) ClusterLabelSelector(v map[string]string) *TopologyPolicy {
	o.Properties.clusterLabelSelector = &v
	return o
}

// GetClusterSelector returns the ClusterSelector field value if set, zero value otherwise.
func (o *TopologyPolicy) GetClusterSelector() map[string]string {
	if o == nil || utils.IsNil(o.Properties.clusterSelector) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.clusterSelector
}

// GetClusterSelectorOk returns a tuple with the ClusterSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetClusterSelectorOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.clusterSelector) {
		return nil, false
	}
	return o.Properties.clusterSelector, true
}

// HasClusterSelector returns a boolean if a field has been set.
func (o *TopologyPolicy) HasClusterSelector() bool {
	if o != nil && !utils.IsNil(o.Properties.clusterSelector) {
		return true
	}

	return false
}

// ClusterSelector gets a reference to the given map[string]string and assigns it to the clusterSelector field.
// clusterSelector:  Deprecated: Use clusterLabelSelector instead.
func (o *TopologyPolicy) ClusterSelector(v map[string]string) *TopologyPolicy {
	o.Properties.clusterSelector = &v
	return o
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *TopologyPolicy) GetClusters() []string {
	if o == nil || utils.IsNil(o.Properties.clusters) {
		var ret []string
		return ret
	}
	return o.Properties.clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetClustersOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.clusters) {
		return nil, false
	}
	return o.Properties.clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *TopologyPolicy) HasClusters() bool {
	if o != nil && !utils.IsNil(o.Properties.clusters) {
		return true
	}

	return false
}

// Clusters gets a reference to the given []string and assigns it to the clusters field.
// clusters:  Specify the names of the clusters to select.
func (o *TopologyPolicy) Clusters(v []string) *TopologyPolicy {
	o.Properties.clusters = v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *TopologyPolicy) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.namespace) {
		var ret string
		return ret
	}
	return *o.Properties.namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyPolicy) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.namespace) {
		return nil, false
	}
	return o.Properties.namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *TopologyPolicy) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.namespace) {
		return true
	}

	return false
}

// Namespace gets a reference to the given string and assigns it to the namespace field.
// namespace:  Specify the target namespace to deploy in the selected clusters, default inherit the original namespace.
func (o *TopologyPolicy) Namespace(v string) *TopologyPolicy {
	o.Properties.namespace = &v
	return o
}

func (o TopologySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.allowEmpty) {
		toSerialize["allowEmpty"] = o.allowEmpty
	}
	if !utils.IsNil(o.clusterLabelSelector) {
		toSerialize["clusterLabelSelector"] = o.clusterLabelSelector
	}
	if !utils.IsNil(o.clusterSelector) {
		toSerialize["clusterSelector"] = o.clusterSelector
	}
	if !utils.IsNil(o.clusters) {
		toSerialize["clusters"] = o.clusters
	}
	if !utils.IsNil(o.namespace) {
		toSerialize["namespace"] = o.namespace
	}
	return toSerialize, nil
}

type NullableTopologySpec struct {
	value *TopologySpec
	isSet bool
}

func (v NullableTopologySpec) Get() *TopologySpec {
	return v.value
}

func (v *NullableTopologySpec) Set(val *TopologySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologySpec(val *TopologySpec) *NullableTopologySpec {
	return &NullableTopologySpec{value: val, isSet: true}
}

func (v NullableTopologySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const TopologyType = "topology"

func init() {
	sdkcommon.RegisterPolicy(TopologyType, FromPolicy)
}

type TopologyPolicy struct {
	Base       apis.PolicyBase
	Properties TopologySpec
}

func Topology(name string) *TopologyPolicy {
	t := &TopologyPolicy{Base: apis.PolicyBase{
		Name: name,
	}}
	return t
}

func (t *TopologyPolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       t.Base.Name,
		Properties: util.Object2RawExtension(t.Properties),
		Type:       TopologyType,
	}
	return res
}

func (t *TopologyPolicy) FromPolicy(from v1beta1.AppPolicy) (*TopologyPolicy, error) {
	var properties TopologySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	t.Base.Name = from.Name
	t.Properties = properties
	return t, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	t := &TopologyPolicy{}
	return t.FromPolicy(from)
}

func (t *TopologyPolicy) Type() string {
	return TopologyType
}
