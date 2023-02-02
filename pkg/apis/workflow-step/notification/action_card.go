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

// checks if the ActionCard type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ActionCard{}

// ActionCard struct for ActionCard
type ActionCard struct {
	BtnOrientation *string              `json:"btnOrientation,omitempty"`
	Btns           utils.NullableString `json:"btns,omitempty"`
	HideAvatar     *string              `json:"hideAvatar,omitempty"`
	SingleTitle    *string              `json:"singleTitle,omitempty"`
	SingleURL      *string              `json:"singleURL,omitempty"`
	Text           *string              `json:"text,omitempty"`
	Title          *string              `json:"title,omitempty"`
}

// NewActionCardWith instantiates a new ActionCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionCardWith() *ActionCard {
	this := ActionCard{}
	return &this
}

// NewActionCard instantiates a new ActionCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionCard() *ActionCard {
	this := ActionCard{}
	return &this
}

// GetBtnOrientation returns the BtnOrientation field value if set, zero value otherwise.
func (o *ActionCard) GetBtnOrientation() string {
	if o == nil || utils.IsNil(o.BtnOrientation) {
		var ret string
		return ret
	}
	return *o.BtnOrientation
}

// GetBtnOrientationOk returns a tuple with the BtnOrientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCard) GetBtnOrientationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BtnOrientation) {
		return nil, false
	}
	return o.BtnOrientation, true
}

// HasBtnOrientation returns a boolean if a field has been set.
func (o *ActionCard) HasBtnOrientation() bool {
	if o != nil && !utils.IsNil(o.BtnOrientation) {
		return true
	}

	return false
}

// SetBtnOrientation gets a reference to the given string and assigns it to the btnOrientation field.
// BtnOrientation:
func (o *ActionCard) SetBtnOrientation(v string) *ActionCard {
	o.BtnOrientation = &v
	return o
}

// GetBtns returns the Btns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCard) GetBtns() string {
	if o == nil || utils.IsNil(o.Btns.Get()) {
		var ret string
		return ret
	}
	return *o.Btns.Get()
}

// GetBtnsOk returns a tuple with the Btns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCard) GetBtnsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Btns.Get(), o.Btns.IsSet()
}

// HasBtns returns a boolean if a field has been set.
func (o *ActionCard) HasBtns() bool {
	if o != nil && o.Btns.IsSet() {
		return true
	}

	return false
}

// SetBtns gets a reference to the given utils.NullableString and assigns it to the btns field.
// Btns:
func (o *ActionCard) SetBtns(v string) *ActionCard {
	o.Btns.Set(&v)
	return o
}

// SetBtnsNil sets the value for Btns to be an explicit nil
func (o *ActionCard) SetBtnsNil() {
	o.Btns.Set(nil)
}

// UnsetBtns ensures that no value is present for Btns, not even an explicit nil
func (o *ActionCard) UnsetBtns() {
	o.Btns.Unset()
}

// GetHideAvatar returns the HideAvatar field value if set, zero value otherwise.
func (o *ActionCard) GetHideAvatar() string {
	if o == nil || utils.IsNil(o.HideAvatar) {
		var ret string
		return ret
	}
	return *o.HideAvatar
}

// GetHideAvatarOk returns a tuple with the HideAvatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCard) GetHideAvatarOk() (*string, bool) {
	if o == nil || utils.IsNil(o.HideAvatar) {
		return nil, false
	}
	return o.HideAvatar, true
}

// HasHideAvatar returns a boolean if a field has been set.
func (o *ActionCard) HasHideAvatar() bool {
	if o != nil && !utils.IsNil(o.HideAvatar) {
		return true
	}

	return false
}

// SetHideAvatar gets a reference to the given string and assigns it to the hideAvatar field.
// HideAvatar:
func (o *ActionCard) SetHideAvatar(v string) *ActionCard {
	o.HideAvatar = &v
	return o
}

// GetSingleTitle returns the SingleTitle field value if set, zero value otherwise.
func (o *ActionCard) GetSingleTitle() string {
	if o == nil || utils.IsNil(o.SingleTitle) {
		var ret string
		return ret
	}
	return *o.SingleTitle
}

// GetSingleTitleOk returns a tuple with the SingleTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCard) GetSingleTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SingleTitle) {
		return nil, false
	}
	return o.SingleTitle, true
}

// HasSingleTitle returns a boolean if a field has been set.
func (o *ActionCard) HasSingleTitle() bool {
	if o != nil && !utils.IsNil(o.SingleTitle) {
		return true
	}

	return false
}

// SetSingleTitle gets a reference to the given string and assigns it to the singleTitle field.
// SingleTitle:
func (o *ActionCard) SetSingleTitle(v string) *ActionCard {
	o.SingleTitle = &v
	return o
}

// GetSingleURL returns the SingleURL field value if set, zero value otherwise.
func (o *ActionCard) GetSingleURL() string {
	if o == nil || utils.IsNil(o.SingleURL) {
		var ret string
		return ret
	}
	return *o.SingleURL
}

// GetSingleURLOk returns a tuple with the SingleURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCard) GetSingleURLOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SingleURL) {
		return nil, false
	}
	return o.SingleURL, true
}

// HasSingleURL returns a boolean if a field has been set.
func (o *ActionCard) HasSingleURL() bool {
	if o != nil && !utils.IsNil(o.SingleURL) {
		return true
	}

	return false
}

// SetSingleURL gets a reference to the given string and assigns it to the singleURL field.
// SingleURL:
func (o *ActionCard) SetSingleURL(v string) *ActionCard {
	o.SingleURL = &v
	return o
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ActionCard) GetText() string {
	if o == nil || utils.IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCard) GetTextOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ActionCard) HasText() bool {
	if o != nil && !utils.IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the text field.
// Text:
func (o *ActionCard) SetText(v string) *ActionCard {
	o.Text = &v
	return o
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ActionCard) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCard) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ActionCard) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the title field.
// Title:
func (o *ActionCard) SetTitle(v string) *ActionCard {
	o.Title = &v
	return o
}

func (o ActionCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BtnOrientation) {
		toSerialize["btnOrientation"] = o.BtnOrientation
	}
	if o.Btns.IsSet() {
		toSerialize["btns"] = o.Btns.Get()
	}
	if !utils.IsNil(o.HideAvatar) {
		toSerialize["hideAvatar"] = o.HideAvatar
	}
	if !utils.IsNil(o.SingleTitle) {
		toSerialize["singleTitle"] = o.SingleTitle
	}
	if !utils.IsNil(o.SingleURL) {
		toSerialize["singleURL"] = o.SingleURL
	}
	if !utils.IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableActionCard struct {
	value *ActionCard
	isSet bool
}

func (v NullableActionCard) Get() *ActionCard {
	return v.value
}

func (v *NullableActionCard) Set(val *ActionCard) {
	v.value = val
	v.isSet = true
}

func (v NullableActionCard) IsSet() bool {
	return v.isSet
}

func (v *NullableActionCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionCard(val *ActionCard) *NullableActionCard {
	return &NullableActionCard{value: val, isSet: true}
}

func (v NullableActionCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
