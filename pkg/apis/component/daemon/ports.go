/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
	"encoding/json"
	"vela-go-sdk/pkg/apis/utils"
)

// checks if the Ports type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Ports{}

// Ports struct for Ports
type Ports struct {
	// Specify if the port should be exposed
	expose bool `json:"expose"`
	// Name of the port
	name *string `json:"name,omitempty"`
	// Number of port to expose on the pod's IP address
	port int32 `json:"port"`
	// Protocol for port. Must be UDP, TCP, or SCTP
	protocol string `json:"protocol"`
}

// NewPortsWith instantiates a new Ports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortsWith(expose bool, port int32, protocol string) *Ports {
	this := Ports{}
	this.expose = expose
	this.port = port
	this.protocol = protocol
	return &this
}

// NewPorts instantiates a new Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPorts() *Ports {
	this := Ports{}
	var expose bool = false
	this.expose = expose
	var protocol string = "TCP"
	this.protocol = protocol
	return &this
}

// GetExpose returns the Expose field value
func (o *Ports) GetExpose() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.expose
}

// GetExposeOk returns a tuple with the Expose field value
// and a boolean to check if the value has been set.
func (o *Ports) GetExposeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.expose, true
}

// Expose sets field value
func (o *Ports) Expose(v bool) *Ports {
	o.expose = v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ports) GetName() string {
	if o == nil || utils.IsNil(o.name) {
		var ret string
		return ret
	}
	return *o.name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ports) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.name) {
		return nil, false
	}
	return o.name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ports) HasName() bool {
	if o != nil && !utils.IsNil(o.name) {
		return true
	}

	return false
}

// Name gets a reference to the given string and assigns it to the name field.
// name:  Name of the port
func (o *Ports) Name(v string) *Ports {
	o.name = &v
	return o
}

// GetPort returns the Port field value
func (o *Ports) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Ports) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.port, true
}

// Port sets field value
func (o *Ports) Port(v int32) *Ports {
	o.port = v
	return o
}

// GetProtocol returns the Protocol field value
func (o *Ports) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *Ports) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.protocol, true
}

// Protocol sets field value
func (o *Ports) Protocol(v string) *Ports {
	o.protocol = v
	return o
}

func (o Ports) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ports) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expose"] = o.expose
	if !utils.IsNil(o.name) {
		toSerialize["name"] = o.name
	}
	toSerialize["port"] = o.port
	toSerialize["protocol"] = o.protocol
	return toSerialize, nil
}

type NullablePorts struct {
	value *Ports
	isSet bool
}

func (v NullablePorts) Get() *Ports {
	return v.value
}

func (v *NullablePorts) Set(val *Ports) {
	v.value = val
	v.isSet = true
}

func (v NullablePorts) IsSet() bool {
	return v.isSet
}

func (v *NullablePorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePorts(val *Ports) *NullablePorts {
	return &NullablePorts{value: val, isSet: true}
}

func (v NullablePorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}