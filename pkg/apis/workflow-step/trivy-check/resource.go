/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trivy_check

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Resource{}

// Resource struct for Resource
type Resource struct {
	Kind *Kind   `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewResourceWith instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceWith() *Resource {
	this := Resource{}
	var kind Kind = "Deployment"
	this.Kind = &kind
	return &this
}

// NewResource instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResource() *Resource {
	this := Resource{}
	var kind Kind = "Deployment"
	this.Kind = &kind
	return &this
}

// NewResources converts a list Resource pointers to objects.
// This is helpful when the SetResource requires a list of objects
func NewResources(ps ...*Resource) []Resource {
	objs := []Resource{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Resource) GetKind() Kind {
	if o == nil || utils.IsNil(o.Kind) {
		var ret Kind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetKindOk() (*Kind, bool) {
	if o == nil || utils.IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Resource) HasKind() bool {
	if o != nil && !utils.IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given Kind and assigns it to the kind field.
// Kind:
func (o *Resource) SetKind(v Kind) *Resource {
	o.Kind = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Resource) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Resource) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:
func (o *Resource) SetName(v string) *Resource {
	o.Name = &v
	return o
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}