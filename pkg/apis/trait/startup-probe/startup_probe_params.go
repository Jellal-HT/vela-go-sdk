/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package startup_probe

import (
	"encoding/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the StartupProbeParams type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StartupProbeParams{}

// StartupProbeParams struct for StartupProbeParams
type StartupProbeParams struct {
	// Specify the name of the target container, if not set, use the component name
	containerName *string `json:"containerName,omitempty"`
	exec          *Exec   `json:"exec,omitempty"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded. Minimum value is 1.
	failureThreshold *int32   `json:"failureThreshold,omitempty"`
	grpc             *Grpc    `json:"grpc,omitempty"`
	httpGet          *HttpGet `json:"httpGet,omitempty"`
	// Number of seconds after the container has started before liveness probes are initiated. Minimum value is 0.
	initialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`
	// How often, in seconds, to execute the probe. Minimum value is 1.
	periodSeconds *int32 `json:"periodSeconds,omitempty"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.  Minimum value is 1.
	successThreshold *int32     `json:"successThreshold,omitempty"`
	tcpSocket        *TcpSocket `json:"tcpSocket,omitempty"`
	// Optional duration in seconds the pod needs to terminate gracefully upon probe failure. Set this value longer than the expected cleanup time for your process.
	terminationGracePeriodSeconds *int32 `json:"terminationGracePeriodSeconds,omitempty"`
	// Number of seconds after which the probe times out. Minimum value is 1.
	timeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// NewStartupProbeParamsWith instantiates a new StartupProbeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartupProbeParamsWith() *StartupProbeParams {
	this := StartupProbeParams{}
	var containerName string = ""
	this.containerName = &containerName
	var failureThreshold int32 = 3
	this.failureThreshold = &failureThreshold
	var initialDelaySeconds int32 = 0
	this.initialDelaySeconds = &initialDelaySeconds
	var periodSeconds int32 = 10
	this.periodSeconds = &periodSeconds
	var successThreshold int32 = 1
	this.successThreshold = &successThreshold
	var timeoutSeconds int32 = 1
	this.timeoutSeconds = &timeoutSeconds
	return &this
}

// NewStartupProbeParams instantiates a new StartupProbeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartupProbeParams() *StartupProbeParams {
	this := StartupProbeParams{}
	var containerName string = ""
	this.containerName = &containerName
	var failureThreshold int32 = 3
	this.failureThreshold = &failureThreshold
	var initialDelaySeconds int32 = 0
	this.initialDelaySeconds = &initialDelaySeconds
	var periodSeconds int32 = 10
	this.periodSeconds = &periodSeconds
	var successThreshold int32 = 1
	this.successThreshold = &successThreshold
	var timeoutSeconds int32 = 1
	this.timeoutSeconds = &timeoutSeconds
	return &this
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *StartupProbeParams) GetContainerName() string {
	if o == nil || utils.IsNil(o.containerName) {
		var ret string
		return ret
	}
	return *o.containerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetContainerNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.containerName) {
		return nil, false
	}
	return o.containerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *StartupProbeParams) HasContainerName() bool {
	if o != nil && !utils.IsNil(o.containerName) {
		return true
	}

	return false
}

// ContainerName gets a reference to the given string and assigns it to the containerName field.
// containerName:  Specify the name of the target container, if not set, use the component name
func (o *StartupProbeParams) ContainerName(v string) *StartupProbeParams {
	o.containerName = &v
	return o
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *StartupProbeParams) GetExec() Exec {
	if o == nil || utils.IsNil(o.exec) {
		var ret Exec
		return ret
	}
	return *o.exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetExecOk() (*Exec, bool) {
	if o == nil || utils.IsNil(o.exec) {
		return nil, false
	}
	return o.exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *StartupProbeParams) HasExec() bool {
	if o != nil && !utils.IsNil(o.exec) {
		return true
	}

	return false
}

// Exec gets a reference to the given Exec and assigns it to the exec field.
// exec:
func (o *StartupProbeParams) Exec(v Exec) *StartupProbeParams {
	o.exec = &v
	return o
}

// GetFailureThreshold returns the FailureThreshold field value if set, zero value otherwise.
func (o *StartupProbeParams) GetFailureThreshold() int32 {
	if o == nil || utils.IsNil(o.failureThreshold) {
		var ret int32
		return ret
	}
	return *o.failureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetFailureThresholdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.failureThreshold) {
		return nil, false
	}
	return o.failureThreshold, true
}

// HasFailureThreshold returns a boolean if a field has been set.
func (o *StartupProbeParams) HasFailureThreshold() bool {
	if o != nil && !utils.IsNil(o.failureThreshold) {
		return true
	}

	return false
}

// FailureThreshold gets a reference to the given int32 and assigns it to the failureThreshold field.
// failureThreshold:  Minimum consecutive failures for the probe to be considered failed after having succeeded. Minimum value is 1.
func (o *StartupProbeParams) FailureThreshold(v int32) *StartupProbeParams {
	o.failureThreshold = &v
	return o
}

// GetGrpc returns the Grpc field value if set, zero value otherwise.
func (o *StartupProbeParams) GetGrpc() Grpc {
	if o == nil || utils.IsNil(o.grpc) {
		var ret Grpc
		return ret
	}
	return *o.grpc
}

// GetGrpcOk returns a tuple with the Grpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetGrpcOk() (*Grpc, bool) {
	if o == nil || utils.IsNil(o.grpc) {
		return nil, false
	}
	return o.grpc, true
}

// HasGrpc returns a boolean if a field has been set.
func (o *StartupProbeParams) HasGrpc() bool {
	if o != nil && !utils.IsNil(o.grpc) {
		return true
	}

	return false
}

// Grpc gets a reference to the given Grpc and assigns it to the grpc field.
// grpc:
func (o *StartupProbeParams) Grpc(v Grpc) *StartupProbeParams {
	o.grpc = &v
	return o
}

// GetHttpGet returns the HttpGet field value if set, zero value otherwise.
func (o *StartupProbeParams) GetHttpGet() HttpGet {
	if o == nil || utils.IsNil(o.httpGet) {
		var ret HttpGet
		return ret
	}
	return *o.httpGet
}

// GetHttpGetOk returns a tuple with the HttpGet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetHttpGetOk() (*HttpGet, bool) {
	if o == nil || utils.IsNil(o.httpGet) {
		return nil, false
	}
	return o.httpGet, true
}

// HasHttpGet returns a boolean if a field has been set.
func (o *StartupProbeParams) HasHttpGet() bool {
	if o != nil && !utils.IsNil(o.httpGet) {
		return true
	}

	return false
}

// HttpGet gets a reference to the given HttpGet and assigns it to the httpGet field.
// httpGet:
func (o *StartupProbeParams) HttpGet(v HttpGet) *StartupProbeParams {
	o.httpGet = &v
	return o
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetInitialDelaySeconds() int32 {
	if o == nil || utils.IsNil(o.initialDelaySeconds) {
		var ret int32
		return ret
	}
	return *o.initialDelaySeconds
}

// GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetInitialDelaySecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.initialDelaySeconds) {
		return nil, false
	}
	return o.initialDelaySeconds, true
}

// HasInitialDelaySeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasInitialDelaySeconds() bool {
	if o != nil && !utils.IsNil(o.initialDelaySeconds) {
		return true
	}

	return false
}

// InitialDelaySeconds gets a reference to the given int32 and assigns it to the initialDelaySeconds field.
// initialDelaySeconds:  Number of seconds after the container has started before liveness probes are initiated. Minimum value is 0.
func (o *StartupProbeParams) InitialDelaySeconds(v int32) *StartupProbeParams {
	o.initialDelaySeconds = &v
	return o
}

// GetPeriodSeconds returns the PeriodSeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetPeriodSeconds() int32 {
	if o == nil || utils.IsNil(o.periodSeconds) {
		var ret int32
		return ret
	}
	return *o.periodSeconds
}

// GetPeriodSecondsOk returns a tuple with the PeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetPeriodSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.periodSeconds) {
		return nil, false
	}
	return o.periodSeconds, true
}

// HasPeriodSeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasPeriodSeconds() bool {
	if o != nil && !utils.IsNil(o.periodSeconds) {
		return true
	}

	return false
}

// PeriodSeconds gets a reference to the given int32 and assigns it to the periodSeconds field.
// periodSeconds:  How often, in seconds, to execute the probe. Minimum value is 1.
func (o *StartupProbeParams) PeriodSeconds(v int32) *StartupProbeParams {
	o.periodSeconds = &v
	return o
}

// GetSuccessThreshold returns the SuccessThreshold field value if set, zero value otherwise.
func (o *StartupProbeParams) GetSuccessThreshold() int32 {
	if o == nil || utils.IsNil(o.successThreshold) {
		var ret int32
		return ret
	}
	return *o.successThreshold
}

// GetSuccessThresholdOk returns a tuple with the SuccessThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetSuccessThresholdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.successThreshold) {
		return nil, false
	}
	return o.successThreshold, true
}

// HasSuccessThreshold returns a boolean if a field has been set.
func (o *StartupProbeParams) HasSuccessThreshold() bool {
	if o != nil && !utils.IsNil(o.successThreshold) {
		return true
	}

	return false
}

// SuccessThreshold gets a reference to the given int32 and assigns it to the successThreshold field.
// successThreshold:  Minimum consecutive successes for the probe to be considered successful after having failed.  Minimum value is 1.
func (o *StartupProbeParams) SuccessThreshold(v int32) *StartupProbeParams {
	o.successThreshold = &v
	return o
}

// GetTcpSocket returns the TcpSocket field value if set, zero value otherwise.
func (o *StartupProbeParams) GetTcpSocket() TcpSocket {
	if o == nil || utils.IsNil(o.tcpSocket) {
		var ret TcpSocket
		return ret
	}
	return *o.tcpSocket
}

// GetTcpSocketOk returns a tuple with the TcpSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetTcpSocketOk() (*TcpSocket, bool) {
	if o == nil || utils.IsNil(o.tcpSocket) {
		return nil, false
	}
	return o.tcpSocket, true
}

// HasTcpSocket returns a boolean if a field has been set.
func (o *StartupProbeParams) HasTcpSocket() bool {
	if o != nil && !utils.IsNil(o.tcpSocket) {
		return true
	}

	return false
}

// TcpSocket gets a reference to the given TcpSocket and assigns it to the tcpSocket field.
// tcpSocket:
func (o *StartupProbeParams) TcpSocket(v TcpSocket) *StartupProbeParams {
	o.tcpSocket = &v
	return o
}

// GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetTerminationGracePeriodSeconds() int32 {
	if o == nil || utils.IsNil(o.terminationGracePeriodSeconds) {
		var ret int32
		return ret
	}
	return *o.terminationGracePeriodSeconds
}

// GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetTerminationGracePeriodSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.terminationGracePeriodSeconds) {
		return nil, false
	}
	return o.terminationGracePeriodSeconds, true
}

// HasTerminationGracePeriodSeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasTerminationGracePeriodSeconds() bool {
	if o != nil && !utils.IsNil(o.terminationGracePeriodSeconds) {
		return true
	}

	return false
}

// TerminationGracePeriodSeconds gets a reference to the given int32 and assigns it to the terminationGracePeriodSeconds field.
// terminationGracePeriodSeconds:  Optional duration in seconds the pod needs to terminate gracefully upon probe failure. Set this value longer than the expected cleanup time for your process.
func (o *StartupProbeParams) TerminationGracePeriodSeconds(v int32) *StartupProbeParams {
	o.terminationGracePeriodSeconds = &v
	return o
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *StartupProbeParams) GetTimeoutSeconds() int32 {
	if o == nil || utils.IsNil(o.timeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.timeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupProbeParams) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.timeoutSeconds) {
		return nil, false
	}
	return o.timeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *StartupProbeParams) HasTimeoutSeconds() bool {
	if o != nil && !utils.IsNil(o.timeoutSeconds) {
		return true
	}

	return false
}

// TimeoutSeconds gets a reference to the given int32 and assigns it to the timeoutSeconds field.
// timeoutSeconds:  Number of seconds after which the probe times out. Minimum value is 1.
func (o *StartupProbeParams) TimeoutSeconds(v int32) *StartupProbeParams {
	o.timeoutSeconds = &v
	return o
}

func (o StartupProbeParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartupProbeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.containerName) {
		toSerialize["containerName"] = o.containerName
	}
	if !utils.IsNil(o.exec) {
		toSerialize["exec"] = o.exec
	}
	if !utils.IsNil(o.failureThreshold) {
		toSerialize["failureThreshold"] = o.failureThreshold
	}
	if !utils.IsNil(o.grpc) {
		toSerialize["grpc"] = o.grpc
	}
	if !utils.IsNil(o.httpGet) {
		toSerialize["httpGet"] = o.httpGet
	}
	if !utils.IsNil(o.initialDelaySeconds) {
		toSerialize["initialDelaySeconds"] = o.initialDelaySeconds
	}
	if !utils.IsNil(o.periodSeconds) {
		toSerialize["periodSeconds"] = o.periodSeconds
	}
	if !utils.IsNil(o.successThreshold) {
		toSerialize["successThreshold"] = o.successThreshold
	}
	if !utils.IsNil(o.tcpSocket) {
		toSerialize["tcpSocket"] = o.tcpSocket
	}
	if !utils.IsNil(o.terminationGracePeriodSeconds) {
		toSerialize["terminationGracePeriodSeconds"] = o.terminationGracePeriodSeconds
	}
	if !utils.IsNil(o.timeoutSeconds) {
		toSerialize["timeoutSeconds"] = o.timeoutSeconds
	}
	return toSerialize, nil
}

type NullableStartupProbeParams struct {
	value *StartupProbeParams
	isSet bool
}

func (v NullableStartupProbeParams) Get() *StartupProbeParams {
	return v.value
}

func (v *NullableStartupProbeParams) Set(val *StartupProbeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableStartupProbeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableStartupProbeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartupProbeParams(val *StartupProbeParams) *NullableStartupProbeParams {
	return &NullableStartupProbeParams{value: val, isSet: true}
}

func (v NullableStartupProbeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartupProbeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
