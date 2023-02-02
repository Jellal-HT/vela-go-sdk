/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export2secret

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the DockerRegistry type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DockerRegistry{}

// DockerRegistry Specify the docker data
type DockerRegistry struct {
	// Specify the password of the docker registry
	Password *string `json:"password,omitempty"`
	// Specify the server of the docker registry
	Server *string `json:"server,omitempty"`
	// Specify the username of the docker registry
	Username *string `json:"username,omitempty"`
}

// NewDockerRegistryWith instantiates a new DockerRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerRegistryWith() *DockerRegistry {
	this := DockerRegistry{}
	var server string = "https://index.docker.io/v1/"
	this.Server = &server
	return &this
}

// NewDockerRegistry instantiates a new DockerRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerRegistry() *DockerRegistry {
	this := DockerRegistry{}
	var server string = "https://index.docker.io/v1/"
	this.Server = &server
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DockerRegistry) GetPassword() string {
	if o == nil || utils.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerRegistry) GetPasswordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DockerRegistry) HasPassword() bool {
	if o != nil && !utils.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the password field.
// Password:  Specify the password of the docker registry
func (o *DockerRegistry) SetPassword(v string) *DockerRegistry {
	o.Password = &v
	return o
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *DockerRegistry) GetServer() string {
	if o == nil || utils.IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerRegistry) GetServerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *DockerRegistry) HasServer() bool {
	if o != nil && !utils.IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the server field.
// Server:  Specify the server of the docker registry
func (o *DockerRegistry) SetServer(v string) *DockerRegistry {
	o.Server = &v
	return o
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DockerRegistry) GetUsername() string {
	if o == nil || utils.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerRegistry) GetUsernameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DockerRegistry) HasUsername() bool {
	if o != nil && !utils.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the username field.
// Username:  Specify the username of the docker registry
func (o *DockerRegistry) SetUsername(v string) *DockerRegistry {
	o.Username = &v
	return o
}

func (o DockerRegistry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerRegistry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !utils.IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !utils.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableDockerRegistry struct {
	value *DockerRegistry
	isSet bool
}

func (v NullableDockerRegistry) Get() *DockerRegistry {
	return v.value
}

func (v *NullableDockerRegistry) Set(val *DockerRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerRegistry(val *DockerRegistry) *NullableDockerRegistry {
	return &NullableDockerRegistry{value: val, isSet: true}
}

func (v NullableDockerRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
