/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package request

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the RequestSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequestSpec{}

// RequestSpec struct for RequestSpec
type RequestSpec struct {
	Body   map[string]interface{} `json:"body,omitempty"`
	Header *map[string]string     `json:"header,omitempty"`
	Method *string                `json:"method,omitempty"`
	Url    *string                `json:"url,omitempty"`
}

// NewRequestSpecWith instantiates a new RequestSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSpecWith() *RequestSpec {
	this := RequestSpec{}
	var method string = "GET"
	this.Method = &method
	return &this
}

// NewRequestSpec instantiates a new RequestSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSpec() *RequestSpec {
	this := RequestSpec{}
	var method string = "GET"
	this.Method = &method
	return &this
}

// NewRequestSpecs converts a list RequestSpec pointers to objects.
// This is helpful when the SetRequestSpec requires a list of objects
func NewRequestSpecs(ps ...*RequestSpec) []RequestSpec {
	objs := []RequestSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *RequestWorkflowStep) GetBody() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.Body) {
		return map[string]interface{}{}, false
	}
	return o.Properties.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *RequestWorkflowStep) HasBody() bool {
	if o != nil && !utils.IsNil(o.Properties.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the body field.
// Body:
func (o *RequestWorkflowStep) SetBody(v map[string]interface{}) *RequestWorkflowStep {
	o.Properties.Body = v
	return o
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *RequestWorkflowStep) GetHeader() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Header) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetHeaderOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Header) {
		return nil, false
	}
	return o.Properties.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *RequestWorkflowStep) HasHeader() bool {
	if o != nil && !utils.IsNil(o.Properties.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given map[string]string and assigns it to the header field.
// Header:
func (o *RequestWorkflowStep) SetHeader(v map[string]string) *RequestWorkflowStep {
	o.Properties.Header = &v
	return o
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RequestWorkflowStep) GetMethod() string {
	if o == nil || utils.IsNil(o.Properties.Method) {
		var ret string
		return ret
	}
	return *o.Properties.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Method) {
		return nil, false
	}
	return o.Properties.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RequestWorkflowStep) HasMethod() bool {
	if o != nil && !utils.IsNil(o.Properties.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the method field.
// Method:
func (o *RequestWorkflowStep) SetMethod(v string) *RequestWorkflowStep {
	o.Properties.Method = &v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RequestWorkflowStep) GetUrl() string {
	if o == nil || utils.IsNil(o.Properties.Url) {
		var ret string
		return ret
	}
	return *o.Properties.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Url) {
		return nil, false
	}
	return o.Properties.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RequestWorkflowStep) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Properties.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the url field.
// Url:
func (o *RequestWorkflowStep) SetUrl(v string) *RequestWorkflowStep {
	o.Properties.Url = &v
	return o
}

func (o RequestSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !utils.IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !utils.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableRequestSpec struct {
	value *RequestSpec
	isSet bool
}

func (v NullableRequestSpec) Get() *RequestSpec {
	return v.value
}

func (v *NullableRequestSpec) Set(val *RequestSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSpec(val *RequestSpec) *NullableRequestSpec {
	return &NullableRequestSpec{value: val, isSet: true}
}

func (v NullableRequestSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const RequestType = "request"

func init() {
	sdkcommon.RegisterWorkflowStep(RequestType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(RequestType, FromWorkflowSubStep)
}

type RequestWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties RequestSpec
}

func Request(name string) *RequestWorkflowStep {
	r := &RequestWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: RequestType,
	}}
	return r
}

func (r *RequestWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range r.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  r.Base.DependsOn,
		If:         r.Base.If,
		Inputs:     r.Base.Inputs,
		Meta:       r.Base.Meta,
		Name:       r.Base.Name,
		Outputs:    r.Base.Outputs,
		Properties: util.Object2RawExtension(r.Properties),
		SubSteps:   subSteps,
		Timeout:    r.Base.Timeout,
		Type:       RequestType,
	}
	return res
}

func (r *RequestWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*RequestWorkflowStep, error) {
	var properties RequestSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := r.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	r.Base.Name = from.Name
	r.Base.DependsOn = from.DependsOn
	r.Base.Inputs = from.Inputs
	r.Base.Outputs = from.Outputs
	r.Base.If = from.If
	r.Base.Timeout = from.Timeout
	r.Base.Meta = from.Meta
	r.Base.Type = RequestType
	r.Properties = properties
	r.Base.SubSteps = subSteps
	return r, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	r := &RequestWorkflowStep{}
	return r.FromWorkflowStep(from)
}

func (r *RequestWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*RequestWorkflowStep, error) {
	var properties RequestSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Name = from.Name
	r.Base.DependsOn = from.DependsOn
	r.Base.Inputs = from.Inputs
	r.Base.Outputs = from.Outputs
	r.Base.If = from.If
	r.Base.Timeout = from.Timeout
	r.Base.Meta = from.Meta
	r.Base.Type = RequestType
	r.Properties = properties
	return r, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	r := &RequestWorkflowStep{}
	return r.FromWorkflowSubStep(from)
}

func (r *RequestWorkflowStep) WorkflowStepName() string {
	return r.Base.Name
}

func (r *RequestWorkflowStep) DefType() string {
	return RequestType
}

func (r *RequestWorkflowStep) If(_if string) *RequestWorkflowStep {
	r.Base.If = _if
	return r
}

func (r *RequestWorkflowStep) Alias(alias string) *RequestWorkflowStep {
	r.Base.Meta.Alias = alias
	return r
}

func (r *RequestWorkflowStep) Timeout(timeout string) *RequestWorkflowStep {
	r.Base.Timeout = timeout
	return r
}

func (r *RequestWorkflowStep) DependsOn(dependsOn []string) *RequestWorkflowStep {
	r.Base.DependsOn = dependsOn
	return r
}

func (r *RequestWorkflowStep) Inputs(input common.StepInputs) *RequestWorkflowStep {
	r.Base.Inputs = input
	return r
}

func (r *RequestWorkflowStep) Outputs(output common.StepOutputs) *RequestWorkflowStep {
	r.Base.Outputs = output
	return r
}
