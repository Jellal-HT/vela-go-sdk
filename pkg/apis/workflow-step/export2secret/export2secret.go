/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export2secret

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Export2secretSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Export2secretSpec{}

// Export2secretSpec struct for Export2secretSpec
type Export2secretSpec struct {
	// Specify the cluster of the secret
	Cluster string `json:"cluster"`
	// Specify the data of secret
	Data           map[string]interface{} `json:"data"`
	DockerRegistry *DockerRegistry        `json:"dockerRegistry,omitempty"`
	// Specify the kind of the secret
	Kind string `json:"kind"`
	// Specify the namespace of the secret
	Namespace *string `json:"namespace,omitempty"`
	// Specify the name of the secret
	SecretName string `json:"secretName"`
	// Specify the type of the secret
	Type *string `json:"type,omitempty"`
}

// NewExport2secretSpecWith instantiates a new Export2secretSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExport2secretSpecWith(cluster string, data map[string]interface{}, kind string, secretName string) *Export2secretSpec {
	this := Export2secretSpec{}
	this.Cluster = cluster
	this.Data = data
	this.Kind = kind
	this.SecretName = secretName
	return &this
}

// NewExport2secretSpec instantiates a new Export2secretSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExport2secretSpec() *Export2secretSpec {
	this := Export2secretSpec{}
	var cluster string = ""
	this.Cluster = cluster
	var kind string = "generic"
	this.Kind = kind
	return &this
}

// NewExport2secretSpecs converts a list Export2secretSpec pointers to objects.
// This is helpful when the SetExport2secretSpec requires a list of objects
func NewExport2secretSpecs(ps ...*Export2secretSpec) []Export2secretSpec {
	objs := []Export2secretSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetCluster returns the Cluster field value
func (o *Export2secretWorkflowStep) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Export2secretWorkflowStep) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.Cluster, true
}

// SetCluster sets field value
func (o *Export2secretWorkflowStep) SetCluster(v string) *Export2secretWorkflowStep {
	o.Properties.Cluster = v
	return o
}

// GetData returns the Data field value
func (o *Export2secretWorkflowStep) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Export2secretWorkflowStep) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties.Data, true
}

// SetData sets field value
func (o *Export2secretWorkflowStep) SetData(v map[string]interface{}) *Export2secretWorkflowStep {
	o.Properties.Data = v
	return o
}

// GetDockerRegistry returns the DockerRegistry field value if set, zero value otherwise.
func (o *Export2secretWorkflowStep) GetDockerRegistry() DockerRegistry {
	if o == nil || utils.IsNil(o.Properties.DockerRegistry) {
		var ret DockerRegistry
		return ret
	}
	return *o.Properties.DockerRegistry
}

// GetDockerRegistryOk returns a tuple with the DockerRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export2secretWorkflowStep) GetDockerRegistryOk() (*DockerRegistry, bool) {
	if o == nil || utils.IsNil(o.Properties.DockerRegistry) {
		return nil, false
	}
	return o.Properties.DockerRegistry, true
}

// HasDockerRegistry returns a boolean if a field has been set.
func (o *Export2secretWorkflowStep) HasDockerRegistry() bool {
	if o != nil && !utils.IsNil(o.Properties.DockerRegistry) {
		return true
	}

	return false
}

// SetDockerRegistry gets a reference to the given DockerRegistry and assigns it to the dockerRegistry field.
// DockerRegistry:
func (o *Export2secretWorkflowStep) SetDockerRegistry(v DockerRegistry) *Export2secretWorkflowStep {
	o.Properties.DockerRegistry = &v
	return o
}

// GetKind returns the Kind field value
func (o *Export2secretWorkflowStep) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Export2secretWorkflowStep) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.Kind, true
}

// SetKind sets field value
func (o *Export2secretWorkflowStep) SetKind(v string) *Export2secretWorkflowStep {
	o.Properties.Kind = v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Export2secretWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export2secretWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Export2secretWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace of the secret
func (o *Export2secretWorkflowStep) SetNamespace(v string) *Export2secretWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

// GetSecretName returns the SecretName field value
func (o *Export2secretWorkflowStep) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *Export2secretWorkflowStep) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.SecretName, true
}

// SetSecretName sets field value
func (o *Export2secretWorkflowStep) SetSecretName(v string) *Export2secretWorkflowStep {
	o.Properties.SecretName = v
	return o
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Export2secretWorkflowStep) GetType() string {
	if o == nil || utils.IsNil(o.Properties.Type) {
		var ret string
		return ret
	}
	return *o.Properties.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export2secretWorkflowStep) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Type) {
		return nil, false
	}
	return o.Properties.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Export2secretWorkflowStep) HasType() bool {
	if o != nil && !utils.IsNil(o.Properties.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the type_ field.
// Type:  Specify the type of the secret
func (o *Export2secretWorkflowStep) SetType(v string) *Export2secretWorkflowStep {
	o.Properties.Type = &v
	return o
}

func (o Export2secretSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Export2secretSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["data"] = o.Data
	if !utils.IsNil(o.DockerRegistry) {
		toSerialize["dockerRegistry"] = o.DockerRegistry
	}
	toSerialize["kind"] = o.Kind
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	toSerialize["secretName"] = o.SecretName
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableExport2secretSpec struct {
	value *Export2secretSpec
	isSet bool
}

func (v NullableExport2secretSpec) Get() *Export2secretSpec {
	return v.value
}

func (v *NullableExport2secretSpec) Set(val *Export2secretSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableExport2secretSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableExport2secretSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExport2secretSpec(val *Export2secretSpec) *NullableExport2secretSpec {
	return &NullableExport2secretSpec{value: val, isSet: true}
}

func (v NullableExport2secretSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExport2secretSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const Export2secretType = "export2secret"

func init() {
	sdkcommon.RegisterWorkflowStep(Export2secretType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(Export2secretType, FromWorkflowSubStep)
}

type Export2secretWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties Export2secretSpec
}

func Export2secret(name string) *Export2secretWorkflowStep {
	e := &Export2secretWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: Export2secretType,
	}}
	return e
}

func (e *Export2secretWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range e.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  e.Base.DependsOn,
		If:         e.Base.If,
		Inputs:     e.Base.Inputs,
		Meta:       e.Base.Meta,
		Name:       e.Base.Name,
		Outputs:    e.Base.Outputs,
		Properties: util.Object2RawExtension(e.Properties),
		SubSteps:   subSteps,
		Timeout:    e.Base.Timeout,
		Type:       Export2secretType,
	}
	return res
}

func (e *Export2secretWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*Export2secretWorkflowStep, error) {
	var properties Export2secretSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := e.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	e.Base.Name = from.Name
	e.Base.DependsOn = from.DependsOn
	e.Base.Inputs = from.Inputs
	e.Base.Outputs = from.Outputs
	e.Base.If = from.If
	e.Base.Timeout = from.Timeout
	e.Base.Meta = from.Meta
	e.Base.Type = Export2secretType
	e.Properties = properties
	e.Base.SubSteps = subSteps
	return e, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	e := &Export2secretWorkflowStep{}
	return e.FromWorkflowStep(from)
}

func (e *Export2secretWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*Export2secretWorkflowStep, error) {
	var properties Export2secretSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	e.Base.Name = from.Name
	e.Base.DependsOn = from.DependsOn
	e.Base.Inputs = from.Inputs
	e.Base.Outputs = from.Outputs
	e.Base.If = from.If
	e.Base.Timeout = from.Timeout
	e.Base.Meta = from.Meta
	e.Base.Type = Export2secretType
	e.Properties = properties
	return e, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	e := &Export2secretWorkflowStep{}
	return e.FromWorkflowSubStep(from)
}

func (e *Export2secretWorkflowStep) WorkflowStepName() string {
	return e.Base.Name
}

func (e *Export2secretWorkflowStep) DefType() string {
	return Export2secretType
}

func (e *Export2secretWorkflowStep) If(_if string) *Export2secretWorkflowStep {
	e.Base.If = _if
	return e
}

func (e *Export2secretWorkflowStep) Alias(alias string) *Export2secretWorkflowStep {
	e.Base.Meta.Alias = alias
	return e
}

func (e *Export2secretWorkflowStep) Timeout(timeout string) *Export2secretWorkflowStep {
	e.Base.Timeout = timeout
	return e
}

func (e *Export2secretWorkflowStep) DependsOn(dependsOn []string) *Export2secretWorkflowStep {
	e.Base.DependsOn = dependsOn
	return e
}

func (e *Export2secretWorkflowStep) Inputs(input common.StepInputs) *Export2secretWorkflowStep {
	e.Base.Inputs = input
	return e
}

func (e *Export2secretWorkflowStep) Outputs(output common.StepOutputs) *Export2secretWorkflowStep {
	e.Base.Outputs = output
	return e
}
