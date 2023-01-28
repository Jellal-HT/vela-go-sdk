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
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"vela-go-sdk/api"
	"vela-go-sdk/apis/utils"
)

// checks if the RequestSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequestSpec{}

// RequestSpec struct for RequestSpec
type RequestSpec struct {
	body   map[string]interface{} `json:"body,omitempty"`
	header *map[string]string     `json:"header,omitempty"`
	method string                 `json:"method"`
	url    string                 `json:"url"`
}

// NewRequestSpecWith instantiates a new RequestSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSpecWith(method string, url string) *RequestSpec {
	this := RequestSpec{}
	this.method = method
	this.url = url
	return &this
}

// NewRequestSpec instantiates a new RequestSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSpec() *RequestSpec {
	this := RequestSpec{}
	var method string = "GET"
	this.method = method
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *RequestWorkflowStep) GetBody() map[string]interface{} {
	if o == nil || utils.IsNil(o.Properties.body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties.body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Properties.body) {
		return map[string]interface{}{}, false
	}
	return o.Properties.body, true
}

// HasBody returns a boolean if a field has been set.
func (o *RequestWorkflowStep) HasBody() bool {
	if o != nil && !utils.IsNil(o.Properties.body) {
		return true
	}

	return false
}

// Body gets a reference to the given map[string]interface{} and assigns it to the body field.
// body:
func (o *RequestWorkflowStep) Body(v map[string]interface{}) *RequestWorkflowStep {
	o.Properties.body = v
	return o
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *RequestWorkflowStep) GetHeader() map[string]string {
	if o == nil || utils.IsNil(o.Properties.header) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetHeaderOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.header) {
		return nil, false
	}
	return o.Properties.header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *RequestWorkflowStep) HasHeader() bool {
	if o != nil && !utils.IsNil(o.Properties.header) {
		return true
	}

	return false
}

// Header gets a reference to the given map[string]string and assigns it to the header field.
// header:
func (o *RequestWorkflowStep) Header(v map[string]string) *RequestWorkflowStep {
	o.Properties.header = &v
	return o
}

// GetMethod returns the Method field value
func (o *RequestWorkflowStep) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.method, true
}

// Method sets field value
func (o *RequestWorkflowStep) Method(v string) *RequestWorkflowStep {
	o.Properties.method = v
	return o
}

// GetUrl returns the Url field value
func (o *RequestWorkflowStep) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RequestWorkflowStep) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.url, true
}

// Url sets field value
func (o *RequestWorkflowStep) Url(v string) *RequestWorkflowStep {
	o.Properties.url = v
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
	if !utils.IsNil(o.body) {
		toSerialize["body"] = o.body
	}
	if !utils.IsNil(o.header) {
		toSerialize["header"] = o.header
	}
	toSerialize["method"] = o.method
	toSerialize["url"] = o.url
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

type RequestWorkflowStep struct {
	Base       api.WorkflowStepBase
	Properties RequestSpec
}

func Request() *RequestWorkflowStep {
	r := &RequestWorkflowStep{Base: api.WorkflowStepBase{}}
	return r
}

func (r *RequestWorkflowStep) Build() common.WorkflowStep {
	res := common.WorkflowStep{
		Properties: util.Object2RawExtension(r.Properties),
		Type:       RequestType,
	}
	return res
}

func (r *RequestWorkflowStep) Props() *RequestSpec {
	return &r.Properties
}