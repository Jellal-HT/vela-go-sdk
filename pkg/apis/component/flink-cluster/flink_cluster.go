/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flink_cluster

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the FlinkClusterSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlinkClusterSpec{}

// FlinkClusterSpec struct for FlinkClusterSpec
type FlinkClusterSpec struct {
	// Specify the flink cluster version, e.g \"v1_14\"
	FlinkVersion *string `json:"flinkVersion,omitempty"`
	// Specify the image for flink cluster to run, e.g \"flink:latest\"
	Image *string `json:"image,omitempty"`
	// Specify the uri for the jar of the flink cluster job, e.g \"local:///opt/flink/examples/streaming/StateMachineExample.jar\"
	JarURI *string `json:"jarURI,omitempty"`
	// Specify the cpu of the flink cluster jobManager, e.g 1
	Jmcpu *int32 `json:"jmcpu,omitempty"`
	// Specify the memory of the flink cluster jobManager, e.g \"1024m\"
	Jmmem *string `json:"jmmem,omitempty"`
	// Specify the flink cluster name
	Name *string `json:"name,omitempty"`
	// Specify the namespace for flink cluster to install
	Namespace *string `json:"namespace,omitempty"`
	// Specify the taskmanager.numberOfTaskSlots, e.g \"2\"
	Nots *string `json:"nots,omitempty"`
	// Specify the parallelism of the flink cluster job, e.g 2
	Parallelism *int32 `json:"parallelism,omitempty"`
	// Specify the replicas of the flink cluster jobManager, e.g 1
	Replicas *int32 `json:"replicas,omitempty"`
	// Specify the cpu of the flink cluster taskManager, e.g 1
	Tmcpu *int32 `json:"tmcpu,omitempty"`
	// Specify the memory of the flink cluster taskManager, e.g \"1024m\"
	Tmmem *string `json:"tmmem,omitempty"`
	// Specify the upgradeMode of the flink cluster job, e.g \"stateless\"
	UpgradeMode *string `json:"upgradeMode,omitempty"`
}

// NewFlinkClusterSpecWith instantiates a new FlinkClusterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlinkClusterSpecWith() *FlinkClusterSpec {
	this := FlinkClusterSpec{}
	return &this
}

// NewFlinkClusterSpec instantiates a new FlinkClusterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlinkClusterSpec() *FlinkClusterSpec {
	this := FlinkClusterSpec{}
	return &this
}

// NewFlinkClusterSpecs converts a list FlinkClusterSpec pointers to objects.
// This is helpful when the SetFlinkClusterSpec requires a list of objects
func NewFlinkClusterSpecs(ps ...*FlinkClusterSpec) []FlinkClusterSpec {
	objs := []FlinkClusterSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetFlinkVersion returns the FlinkVersion field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetFlinkVersion() string {
	if o == nil || utils.IsNil(o.Properties.FlinkVersion) {
		var ret string
		return ret
	}
	return *o.Properties.FlinkVersion
}

// GetFlinkVersionOk returns a tuple with the FlinkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetFlinkVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.FlinkVersion) {
		return nil, false
	}
	return o.Properties.FlinkVersion, true
}

// HasFlinkVersion returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasFlinkVersion() bool {
	if o != nil && !utils.IsNil(o.Properties.FlinkVersion) {
		return true
	}

	return false
}

// SetFlinkVersion gets a reference to the given string and assigns it to the flinkVersion field.
// FlinkVersion:  Specify the flink cluster version, e.g \"v1_14\"
func (o *FlinkClusterComponent) SetFlinkVersion(v string) *FlinkClusterComponent {
	o.Properties.FlinkVersion = &v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetImage() string {
	if o == nil || utils.IsNil(o.Properties.Image) {
		var ret string
		return ret
	}
	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetImageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Image) {
		return nil, false
	}
	return o.Properties.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasImage() bool {
	if o != nil && !utils.IsNil(o.Properties.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the image field.
// Image:  Specify the image for flink cluster to run, e.g \"flink:latest\"
func (o *FlinkClusterComponent) SetImage(v string) *FlinkClusterComponent {
	o.Properties.Image = &v
	return o
}

// GetJarURI returns the JarURI field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetJarURI() string {
	if o == nil || utils.IsNil(o.Properties.JarURI) {
		var ret string
		return ret
	}
	return *o.Properties.JarURI
}

// GetJarURIOk returns a tuple with the JarURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetJarURIOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.JarURI) {
		return nil, false
	}
	return o.Properties.JarURI, true
}

// HasJarURI returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasJarURI() bool {
	if o != nil && !utils.IsNil(o.Properties.JarURI) {
		return true
	}

	return false
}

// SetJarURI gets a reference to the given string and assigns it to the jarURI field.
// JarURI:  Specify the uri for the jar of the flink cluster job, e.g \"local:///opt/flink/examples/streaming/StateMachineExample.jar\"
func (o *FlinkClusterComponent) SetJarURI(v string) *FlinkClusterComponent {
	o.Properties.JarURI = &v
	return o
}

// GetJmcpu returns the Jmcpu field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetJmcpu() int32 {
	if o == nil || utils.IsNil(o.Properties.Jmcpu) {
		var ret int32
		return ret
	}
	return *o.Properties.Jmcpu
}

// GetJmcpuOk returns a tuple with the Jmcpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetJmcpuOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Jmcpu) {
		return nil, false
	}
	return o.Properties.Jmcpu, true
}

// HasJmcpu returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasJmcpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Jmcpu) {
		return true
	}

	return false
}

// SetJmcpu gets a reference to the given int32 and assigns it to the jmcpu field.
// Jmcpu:  Specify the cpu of the flink cluster jobManager, e.g 1
func (o *FlinkClusterComponent) SetJmcpu(v int32) *FlinkClusterComponent {
	o.Properties.Jmcpu = &v
	return o
}

// GetJmmem returns the Jmmem field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetJmmem() string {
	if o == nil || utils.IsNil(o.Properties.Jmmem) {
		var ret string
		return ret
	}
	return *o.Properties.Jmmem
}

// GetJmmemOk returns a tuple with the Jmmem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetJmmemOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Jmmem) {
		return nil, false
	}
	return o.Properties.Jmmem, true
}

// HasJmmem returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasJmmem() bool {
	if o != nil && !utils.IsNil(o.Properties.Jmmem) {
		return true
	}

	return false
}

// SetJmmem gets a reference to the given string and assigns it to the jmmem field.
// Jmmem:  Specify the memory of the flink cluster jobManager, e.g \"1024m\"
func (o *FlinkClusterComponent) SetJmmem(v string) *FlinkClusterComponent {
	o.Properties.Jmmem = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the flink cluster name
func (o *FlinkClusterComponent) SetName(v string) *FlinkClusterComponent {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace for flink cluster to install
func (o *FlinkClusterComponent) SetNamespace(v string) *FlinkClusterComponent {
	o.Properties.Namespace = &v
	return o
}

// GetNots returns the Nots field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetNots() string {
	if o == nil || utils.IsNil(o.Properties.Nots) {
		var ret string
		return ret
	}
	return *o.Properties.Nots
}

// GetNotsOk returns a tuple with the Nots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetNotsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Nots) {
		return nil, false
	}
	return o.Properties.Nots, true
}

// HasNots returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasNots() bool {
	if o != nil && !utils.IsNil(o.Properties.Nots) {
		return true
	}

	return false
}

// SetNots gets a reference to the given string and assigns it to the nots field.
// Nots:  Specify the taskmanager.numberOfTaskSlots, e.g \"2\"
func (o *FlinkClusterComponent) SetNots(v string) *FlinkClusterComponent {
	o.Properties.Nots = &v
	return o
}

// GetParallelism returns the Parallelism field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetParallelism() int32 {
	if o == nil || utils.IsNil(o.Properties.Parallelism) {
		var ret int32
		return ret
	}
	return *o.Properties.Parallelism
}

// GetParallelismOk returns a tuple with the Parallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetParallelismOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Parallelism) {
		return nil, false
	}
	return o.Properties.Parallelism, true
}

// HasParallelism returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasParallelism() bool {
	if o != nil && !utils.IsNil(o.Properties.Parallelism) {
		return true
	}

	return false
}

// SetParallelism gets a reference to the given int32 and assigns it to the parallelism field.
// Parallelism:  Specify the parallelism of the flink cluster job, e.g 2
func (o *FlinkClusterComponent) SetParallelism(v int32) *FlinkClusterComponent {
	o.Properties.Parallelism = &v
	return o
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetReplicas() int32 {
	if o == nil || utils.IsNil(o.Properties.Replicas) {
		var ret int32
		return ret
	}
	return *o.Properties.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetReplicasOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Replicas) {
		return nil, false
	}
	return o.Properties.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasReplicas() bool {
	if o != nil && !utils.IsNil(o.Properties.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the replicas field.
// Replicas:  Specify the replicas of the flink cluster jobManager, e.g 1
func (o *FlinkClusterComponent) SetReplicas(v int32) *FlinkClusterComponent {
	o.Properties.Replicas = &v
	return o
}

// GetTmcpu returns the Tmcpu field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetTmcpu() int32 {
	if o == nil || utils.IsNil(o.Properties.Tmcpu) {
		var ret int32
		return ret
	}
	return *o.Properties.Tmcpu
}

// GetTmcpuOk returns a tuple with the Tmcpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetTmcpuOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Tmcpu) {
		return nil, false
	}
	return o.Properties.Tmcpu, true
}

// HasTmcpu returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasTmcpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Tmcpu) {
		return true
	}

	return false
}

// SetTmcpu gets a reference to the given int32 and assigns it to the tmcpu field.
// Tmcpu:  Specify the cpu of the flink cluster taskManager, e.g 1
func (o *FlinkClusterComponent) SetTmcpu(v int32) *FlinkClusterComponent {
	o.Properties.Tmcpu = &v
	return o
}

// GetTmmem returns the Tmmem field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetTmmem() string {
	if o == nil || utils.IsNil(o.Properties.Tmmem) {
		var ret string
		return ret
	}
	return *o.Properties.Tmmem
}

// GetTmmemOk returns a tuple with the Tmmem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetTmmemOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Tmmem) {
		return nil, false
	}
	return o.Properties.Tmmem, true
}

// HasTmmem returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasTmmem() bool {
	if o != nil && !utils.IsNil(o.Properties.Tmmem) {
		return true
	}

	return false
}

// SetTmmem gets a reference to the given string and assigns it to the tmmem field.
// Tmmem:  Specify the memory of the flink cluster taskManager, e.g \"1024m\"
func (o *FlinkClusterComponent) SetTmmem(v string) *FlinkClusterComponent {
	o.Properties.Tmmem = &v
	return o
}

// GetUpgradeMode returns the UpgradeMode field value if set, zero value otherwise.
func (o *FlinkClusterComponent) GetUpgradeMode() string {
	if o == nil || utils.IsNil(o.Properties.UpgradeMode) {
		var ret string
		return ret
	}
	return *o.Properties.UpgradeMode
}

// GetUpgradeModeOk returns a tuple with the UpgradeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkClusterComponent) GetUpgradeModeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.UpgradeMode) {
		return nil, false
	}
	return o.Properties.UpgradeMode, true
}

// HasUpgradeMode returns a boolean if a field has been set.
func (o *FlinkClusterComponent) HasUpgradeMode() bool {
	if o != nil && !utils.IsNil(o.Properties.UpgradeMode) {
		return true
	}

	return false
}

// SetUpgradeMode gets a reference to the given string and assigns it to the upgradeMode field.
// UpgradeMode:  Specify the upgradeMode of the flink cluster job, e.g \"stateless\"
func (o *FlinkClusterComponent) SetUpgradeMode(v string) *FlinkClusterComponent {
	o.Properties.UpgradeMode = &v
	return o
}

func (o FlinkClusterSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlinkClusterSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.FlinkVersion) {
		toSerialize["flinkVersion"] = o.FlinkVersion
	}
	if !utils.IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !utils.IsNil(o.JarURI) {
		toSerialize["jarURI"] = o.JarURI
	}
	if !utils.IsNil(o.Jmcpu) {
		toSerialize["jmcpu"] = o.Jmcpu
	}
	if !utils.IsNil(o.Jmmem) {
		toSerialize["jmmem"] = o.Jmmem
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !utils.IsNil(o.Nots) {
		toSerialize["nots"] = o.Nots
	}
	if !utils.IsNil(o.Parallelism) {
		toSerialize["parallelism"] = o.Parallelism
	}
	if !utils.IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	if !utils.IsNil(o.Tmcpu) {
		toSerialize["tmcpu"] = o.Tmcpu
	}
	if !utils.IsNil(o.Tmmem) {
		toSerialize["tmmem"] = o.Tmmem
	}
	if !utils.IsNil(o.UpgradeMode) {
		toSerialize["upgradeMode"] = o.UpgradeMode
	}
	return toSerialize, nil
}

type NullableFlinkClusterSpec struct {
	value *FlinkClusterSpec
	isSet bool
}

func (v NullableFlinkClusterSpec) Get() *FlinkClusterSpec {
	return v.value
}

func (v *NullableFlinkClusterSpec) Set(val *FlinkClusterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFlinkClusterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFlinkClusterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlinkClusterSpec(val *FlinkClusterSpec) *NullableFlinkClusterSpec {
	return &NullableFlinkClusterSpec{value: val, isSet: true}
}

func (v NullableFlinkClusterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlinkClusterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const FlinkClusterType = "flink-cluster"

func init() {
	sdkcommon.RegisterComponent(FlinkClusterType, FromComponent)
}

type FlinkClusterComponent struct {
	Base       apis.ComponentBase
	Properties FlinkClusterSpec
}

func FlinkCluster(name string) *FlinkClusterComponent {
	f := &FlinkClusterComponent{Base: apis.ComponentBase{
		Name: name,
		Type: FlinkClusterType,
	}}
	return f
}

func (f *FlinkClusterComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range f.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  f.Base.DependsOn,
		Inputs:     f.Base.Inputs,
		Name:       f.Base.Name,
		Outputs:    f.Base.Outputs,
		Properties: util.Object2RawExtension(f.Properties),
		Traits:     traits,
		Type:       FlinkClusterType,
	}
	return res
}

func (f *FlinkClusterComponent) FromComponent(from common.ApplicationComponent) (*FlinkClusterComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		f.Base.Traits = append(f.Base.Traits, _t)
	}
	var properties FlinkClusterSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	f.Base.Name = from.Name
	f.Base.DependsOn = from.DependsOn
	f.Base.Inputs = from.Inputs
	f.Base.Outputs = from.Outputs
	f.Base.Type = FlinkClusterType
	f.Properties = properties
	return f, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	f := &FlinkClusterComponent{}
	return f.FromComponent(from)
}

func (f *FlinkClusterComponent) AddTrait(traits ...apis.Trait) *FlinkClusterComponent {
	f.Base.Traits = append(f.Base.Traits, traits...)
	return f
}

func (f *FlinkClusterComponent) GetTrait(_type string) apis.Trait {
	for _, _t := range f.Base.Traits {
		if _t.DefType() == _type {
			return _t
		}
	}
	return nil
}

func (f *FlinkClusterComponent) ComponentName() string {
	return f.Base.Name
}

func (f *FlinkClusterComponent) DefType() string {
	return FlinkClusterType
}

func (f *FlinkClusterComponent) DependsOn(dependsOn []string) *FlinkClusterComponent {
	f.Base.DependsOn = dependsOn
	return f
}

func (f *FlinkClusterComponent) Inputs(input common.StepInputs) *FlinkClusterComponent {
	f.Base.Inputs = input
	return f
}

func (f *FlinkClusterComponent) Outputs(output common.StepOutputs) *FlinkClusterComponent {
	f.Base.Outputs = output
	return f
}

func (f *FlinkClusterComponent) AddDependsOn(dependsOn string) *FlinkClusterComponent {
	f.Base.DependsOn = append(f.Base.DependsOn, dependsOn)
	return f
}
