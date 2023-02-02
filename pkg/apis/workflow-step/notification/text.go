/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Text type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Text{}

// Text Specify the message content of dingtalk notification
type Text struct {
	Content *string `json:"content,omitempty"`
}

// NewTextWith instantiates a new Text object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextWith() *Text {
	this := Text{}
	return &this
}

// NewText instantiates a new Text object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewText() *Text {
	this := Text{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Text) GetContent() string {
	if o == nil || utils.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Text) GetContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Text) HasContent() bool {
	if o != nil && !utils.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the content field.
// Content:
func (o *Text) SetContent(v string) *Text {
	o.Content = &v
	return o
}

func (o Text) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Text) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableText struct {
	value *Text
	isSet bool
}

func (v NullableText) Get() *Text {
	return v.value
}

func (v *NullableText) Set(val *Text) {
	v.value = val
	v.isSet = true
}

func (v NullableText) IsSet() bool {
	return v.isSet
}

func (v *NullableText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableText(val *Text) *NullableText {
	return &NullableText{value: val, isSet: true}
}

func (v NullableText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
