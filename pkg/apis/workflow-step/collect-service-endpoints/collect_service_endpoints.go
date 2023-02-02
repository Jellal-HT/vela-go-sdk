/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package collect_service_endpoints

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the CollectServiceEndpointsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CollectServiceEndpointsSpec{}

// CollectServiceEndpointsSpec struct for CollectServiceEndpointsSpec
type CollectServiceEndpointsSpec struct {
	// Filter the component of the endpoints
	Components []string `json:"components,omitempty"`
	// Specify the name of the application
	Name *string `json:"name,omitempty"`
	// Specify the namespace of the application
	Namespace *string `json:"namespace,omitempty"`
	// Filter the endpoint that are only outer
	Outer *bool `json:"outer,omitempty"`
	// Filter the port of the endpoints
	Port *int32 `json:"port,omitempty"`
	// Filter the port name of the endpoints
	PortName *string `json:"portName,omitempty"`
	// The protocal of endpoint url
	Protocal *string `json:"protocal,omitempty"`
}

// NewCollectServiceEndpointsSpecWith instantiates a new CollectServiceEndpointsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectServiceEndpointsSpecWith() *CollectServiceEndpointsSpec {
	this := CollectServiceEndpointsSpec{}
	var protocal string = "http"
	this.Protocal = &protocal
	return &this
}

// NewCollectServiceEndpointsSpec instantiates a new CollectServiceEndpointsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectServiceEndpointsSpec() *CollectServiceEndpointsSpec {
	this := CollectServiceEndpointsSpec{}
	var protocal string = "http"
	this.Protocal = &protocal
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetComponents() []string {
	if o == nil || utils.IsNil(o.Properties.Components) {
		var ret []string
		return ret
	}
	return o.Properties.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetComponentsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Components) {
		return nil, false
	}
	return o.Properties.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasComponents() bool {
	if o != nil && !utils.IsNil(o.Properties.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []string and assigns it to the components field.
// Components:  Filter the component of the endpoints
func (o *CollectServiceEndpointsWorkflowStep) SetComponents(v []string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.Components = v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of the application
func (o *CollectServiceEndpointsWorkflowStep) SetName(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace of the application
func (o *CollectServiceEndpointsWorkflowStep) SetNamespace(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

// GetOuter returns the Outer field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetOuter() bool {
	if o == nil || utils.IsNil(o.Properties.Outer) {
		var ret bool
		return ret
	}
	return *o.Properties.Outer
}

// GetOuterOk returns a tuple with the Outer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetOuterOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.Outer) {
		return nil, false
	}
	return o.Properties.Outer, true
}

// HasOuter returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasOuter() bool {
	if o != nil && !utils.IsNil(o.Properties.Outer) {
		return true
	}

	return false
}

// SetOuter gets a reference to the given bool and assigns it to the outer field.
// Outer:  Filter the endpoint that are only outer
func (o *CollectServiceEndpointsWorkflowStep) SetOuter(v bool) *CollectServiceEndpointsWorkflowStep {
	o.Properties.Outer = &v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetPort() int32 {
	if o == nil || utils.IsNil(o.Properties.Port) {
		var ret int32
		return ret
	}
	return *o.Properties.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Port) {
		return nil, false
	}
	return o.Properties.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasPort() bool {
	if o != nil && !utils.IsNil(o.Properties.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the port field.
// Port:  Filter the port of the endpoints
func (o *CollectServiceEndpointsWorkflowStep) SetPort(v int32) *CollectServiceEndpointsWorkflowStep {
	o.Properties.Port = &v
	return o
}

// GetPortName returns the PortName field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetPortName() string {
	if o == nil || utils.IsNil(o.Properties.PortName) {
		var ret string
		return ret
	}
	return *o.Properties.PortName
}

// GetPortNameOk returns a tuple with the PortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetPortNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.PortName) {
		return nil, false
	}
	return o.Properties.PortName, true
}

// HasPortName returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasPortName() bool {
	if o != nil && !utils.IsNil(o.Properties.PortName) {
		return true
	}

	return false
}

// SetPortName gets a reference to the given string and assigns it to the portName field.
// PortName:  Filter the port name of the endpoints
func (o *CollectServiceEndpointsWorkflowStep) SetPortName(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.PortName = &v
	return o
}

// GetProtocal returns the Protocal field value if set, zero value otherwise.
func (o *CollectServiceEndpointsWorkflowStep) GetProtocal() string {
	if o == nil || utils.IsNil(o.Properties.Protocal) {
		var ret string
		return ret
	}
	return *o.Properties.Protocal
}

// GetProtocalOk returns a tuple with the Protocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectServiceEndpointsWorkflowStep) GetProtocalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Protocal) {
		return nil, false
	}
	return o.Properties.Protocal, true
}

// HasProtocal returns a boolean if a field has been set.
func (o *CollectServiceEndpointsWorkflowStep) HasProtocal() bool {
	if o != nil && !utils.IsNil(o.Properties.Protocal) {
		return true
	}

	return false
}

// SetProtocal gets a reference to the given string and assigns it to the protocal field.
// Protocal:  The protocal of endpoint url
func (o *CollectServiceEndpointsWorkflowStep) SetProtocal(v string) *CollectServiceEndpointsWorkflowStep {
	o.Properties.Protocal = &v
	return o
}

func (o CollectServiceEndpointsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectServiceEndpointsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !utils.IsNil(o.Outer) {
		toSerialize["outer"] = o.Outer
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.PortName) {
		toSerialize["portName"] = o.PortName
	}
	if !utils.IsNil(o.Protocal) {
		toSerialize["protocal"] = o.Protocal
	}
	return toSerialize, nil
}

type NullableCollectServiceEndpointsSpec struct {
	value *CollectServiceEndpointsSpec
	isSet bool
}

func (v NullableCollectServiceEndpointsSpec) Get() *CollectServiceEndpointsSpec {
	return v.value
}

func (v *NullableCollectServiceEndpointsSpec) Set(val *CollectServiceEndpointsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectServiceEndpointsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectServiceEndpointsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectServiceEndpointsSpec(val *CollectServiceEndpointsSpec) *NullableCollectServiceEndpointsSpec {
	return &NullableCollectServiceEndpointsSpec{value: val, isSet: true}
}

func (v NullableCollectServiceEndpointsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectServiceEndpointsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const CollectServiceEndpointsType = "collect-service-endpoints"

func init() {
	sdkcommon.RegisterWorkflowStep(CollectServiceEndpointsType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(CollectServiceEndpointsType, FromWorkflowSubStep)
}

type CollectServiceEndpointsWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties CollectServiceEndpointsSpec
}

func CollectServiceEndpoints(name string) *CollectServiceEndpointsWorkflowStep {
	c := &CollectServiceEndpointsWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: CollectServiceEndpointsType,
	}}
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range c.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  c.Base.DependsOn,
		If:         c.Base.If,
		Inputs:     c.Base.Inputs,
		Meta:       c.Base.Meta,
		Name:       c.Base.Name,
		Outputs:    c.Base.Outputs,
		Properties: util.Object2RawExtension(c.Properties),
		SubSteps:   subSteps,
		Timeout:    c.Base.Timeout,
		Type:       CollectServiceEndpointsType,
	}
	return res
}

func (c *CollectServiceEndpointsWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*CollectServiceEndpointsWorkflowStep, error) {
	var properties CollectServiceEndpointsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := c.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	c.Base.Name = from.Name
	c.Base.DependsOn = from.DependsOn
	c.Base.Inputs = from.Inputs
	c.Base.Outputs = from.Outputs
	c.Base.If = from.If
	c.Base.Timeout = from.Timeout
	c.Base.Meta = from.Meta
	c.Base.Type = CollectServiceEndpointsType
	c.Properties = properties
	c.Base.SubSteps = subSteps
	return c, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	c := &CollectServiceEndpointsWorkflowStep{}
	return c.FromWorkflowStep(from)
}

func (c *CollectServiceEndpointsWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*CollectServiceEndpointsWorkflowStep, error) {
	var properties CollectServiceEndpointsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	c.Base.Name = from.Name
	c.Base.DependsOn = from.DependsOn
	c.Base.Inputs = from.Inputs
	c.Base.Outputs = from.Outputs
	c.Base.If = from.If
	c.Base.Timeout = from.Timeout
	c.Base.Meta = from.Meta
	c.Base.Type = CollectServiceEndpointsType
	c.Properties = properties
	return c, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	c := &CollectServiceEndpointsWorkflowStep{}
	return c.FromWorkflowSubStep(from)
}

func (c *CollectServiceEndpointsWorkflowStep) WorkflowStepName() string {
	return c.Base.Name
}

func (c *CollectServiceEndpointsWorkflowStep) DefType() string {
	return CollectServiceEndpointsType
}

func (c *CollectServiceEndpointsWorkflowStep) If(_if string) *CollectServiceEndpointsWorkflowStep {
	c.Base.If = _if
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Alias(alias string) *CollectServiceEndpointsWorkflowStep {
	c.Base.Meta.Alias = alias
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Timeout(timeout string) *CollectServiceEndpointsWorkflowStep {
	c.Base.Timeout = timeout
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) DependsOn(dependsOn []string) *CollectServiceEndpointsWorkflowStep {
	c.Base.DependsOn = dependsOn
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Inputs(input common.StepInputs) *CollectServiceEndpointsWorkflowStep {
	c.Base.Inputs = input
	return c
}

func (c *CollectServiceEndpointsWorkflowStep) Outputs(output common.StepOutputs) *CollectServiceEndpointsWorkflowStep {
	c.Base.Outputs = output
	return c
}
