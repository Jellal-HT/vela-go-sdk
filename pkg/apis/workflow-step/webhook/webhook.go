/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhook

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the WebhookSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebhookSpec{}

// WebhookSpec struct for WebhookSpec
type WebhookSpec struct {
	// Specify the data you want to send
	Data map[string]interface{} `json:"data,omitempty"`
	Url  *Url                   `json:"url,omitempty"`
}

// NewWebhookSpecWith instantiates a new WebhookSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSpecWith() *WebhookSpec {
	this := WebhookSpec{}
	return &this
}

// NewWebhookSpec instantiates a new WebhookSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSpec() *WebhookSpec {
	this := WebhookSpec{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WebhookWorkflowStep) GetData() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWorkflowStep) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.Data) {
		return map[string]interface{}{}, false
	}
	return o.Properties.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WebhookWorkflowStep) HasData() bool {
	if o != nil && !utils.IsNil(o.Properties.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the data field.
// Data:  Specify the data you want to send
func (o *WebhookWorkflowStep) SetData(v map[string]interface{}) *WebhookWorkflowStep {
	o.Properties.Data = v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebhookWorkflowStep) GetUrl() Url {
	if o == nil || utils.IsNil(o.Properties.Url) {
		var ret Url
		return ret
	}
	return *o.Properties.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWorkflowStep) GetUrlOk() (*Url, bool) {
	if o == nil || utils.IsNil(o.Properties.Url) {
		return nil, false
	}
	return o.Properties.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookWorkflowStep) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Properties.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given Url and assigns it to the url field.
// Url:
func (o *WebhookWorkflowStep) SetUrl(v Url) *WebhookWorkflowStep {
	o.Properties.Url = &v
	return o
}

func (o WebhookSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableWebhookSpec struct {
	value *WebhookSpec
	isSet bool
}

func (v NullableWebhookSpec) Get() *WebhookSpec {
	return v.value
}

func (v *NullableWebhookSpec) Set(val *WebhookSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSpec(val *WebhookSpec) *NullableWebhookSpec {
	return &NullableWebhookSpec{value: val, isSet: true}
}

func (v NullableWebhookSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const WebhookType = "webhook"

func init() {
	sdkcommon.RegisterWorkflowStep(WebhookType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(WebhookType, FromWorkflowSubStep)
}

type WebhookWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties WebhookSpec
}

func Webhook(name string) *WebhookWorkflowStep {
	w := &WebhookWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: WebhookType,
	}}
	return w
}

func (w *WebhookWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range w.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  w.Base.DependsOn,
		If:         w.Base.If,
		Inputs:     w.Base.Inputs,
		Meta:       w.Base.Meta,
		Name:       w.Base.Name,
		Outputs:    w.Base.Outputs,
		Properties: util.Object2RawExtension(w.Properties),
		SubSteps:   subSteps,
		Timeout:    w.Base.Timeout,
		Type:       WebhookType,
	}
	return res
}

func (w *WebhookWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*WebhookWorkflowStep, error) {
	var properties WebhookSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := w.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	w.Base.Name = from.Name
	w.Base.DependsOn = from.DependsOn
	w.Base.Inputs = from.Inputs
	w.Base.Outputs = from.Outputs
	w.Base.If = from.If
	w.Base.Timeout = from.Timeout
	w.Base.Meta = from.Meta
	w.Properties = properties
	w.Base.SubSteps = subSteps
	return w, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	w := &WebhookWorkflowStep{}
	return w.FromWorkflowStep(from)
}

func (w *WebhookWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*WebhookWorkflowStep, error) {
	var properties WebhookSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	w.Base.Name = from.Name
	w.Base.DependsOn = from.DependsOn
	w.Base.Inputs = from.Inputs
	w.Base.Outputs = from.Outputs
	w.Base.If = from.If
	w.Base.Timeout = from.Timeout
	w.Base.Meta = from.Meta
	w.Properties = properties
	return w, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	w := &WebhookWorkflowStep{}
	return w.FromWorkflowSubStep(from)
}

func (w *WebhookWorkflowStep) If(_if string) *WebhookWorkflowStep {
	w.Base.If = _if
	return w
}

func (w *WebhookWorkflowStep) Alias(alias string) *WebhookWorkflowStep {
	w.Base.Meta.Alias = alias
	return w
}

func (w *WebhookWorkflowStep) Timeout(timeout string) *WebhookWorkflowStep {
	w.Base.Timeout = timeout
	return w
}

func (w *WebhookWorkflowStep) DependsOn(dependsOn []string) *WebhookWorkflowStep {
	w.Base.DependsOn = dependsOn
	return w
}

func (w *WebhookWorkflowStep) Inputs(input common.StepInputs) *WebhookWorkflowStep {
	w.Base.Inputs = input
	return w
}

func (w *WebhookWorkflowStep) Outputs(output common.StepOutputs) *WebhookWorkflowStep {
	w.Base.Outputs = output
	return w
}

func (w *WebhookWorkflowStep) DefName() string {
	return w.Base.Name
}

func (w *WebhookWorkflowStep) DefType() string {
	return WebhookType
}
