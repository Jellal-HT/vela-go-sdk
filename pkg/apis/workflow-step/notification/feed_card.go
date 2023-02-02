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

// checks if the FeedCard type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FeedCard{}

// FeedCard struct for FeedCard
type FeedCard struct {
	Links []Link `json:"links,omitempty"`
}

// NewFeedCardWith instantiates a new FeedCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedCardWith() *FeedCard {
	this := FeedCard{}
	return &this
}

// NewFeedCard instantiates a new FeedCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedCard() *FeedCard {
	this := FeedCard{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FeedCard) GetLinks() []Link {
	if o == nil || utils.IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedCard) GetLinksOk() ([]Link, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FeedCard) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the links field.
// Links:
func (o *FeedCard) SetLinks(v []Link) *FeedCard {
	o.Links = v
	return o
}

func (o FeedCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableFeedCard struct {
	value *FeedCard
	isSet bool
}

func (v NullableFeedCard) Get() *FeedCard {
	return v.value
}

func (v *NullableFeedCard) Set(val *FeedCard) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedCard) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedCard(val *FeedCard) *NullableFeedCard {
	return &NullableFeedCard{value: val, isSet: true}
}

func (v NullableFeedCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
