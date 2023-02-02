/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webservice

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the WebserviceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebserviceSpec{}

// WebserviceSpec struct for WebserviceSpec
type WebserviceSpec struct {
	addRevisionLabel *bool `json:"addRevisionLabel,omitempty"`
	// Specify the annotations in the workload
	annotations *map[string]string `json:"annotations,omitempty"`
	// Arguments to the entrypoint
	args []string `json:"args,omitempty"`
	// Commands to run in the container
	cmd []string `json:"cmd,omitempty"`
	// Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	cpu *string `json:"cpu,omitempty"`
	// Define arguments by using environment variables
	env        []Env   `json:"env,omitempty"`
	exposeType *string `json:"exposeType,omitempty"`
	// Specify the hostAliases to add
	hostAliases []HostAliases `json:"hostAliases,omitempty"`
	// Which image would you like to use for your service +short=i
	image *string `json:"image,omitempty"`
	// Specify image pull policy for your service
	imagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Specify image pull secrets for your service
	imagePullSecrets []string `json:"imagePullSecrets,omitempty"`
	// Specify the labels in the workload
	labels        *map[string]string `json:"labels,omitempty"`
	livenessProbe *HealthProbe       `json:"livenessProbe,omitempty"`
	// Specifies the attributes of the memory resource required for the container.
	memory *string `json:"memory,omitempty"`
	port   *int32  `json:"port,omitempty"`
	// Which ports do you want customer traffic sent to, defaults to 80
	ports          []Ports       `json:"ports,omitempty"`
	readinessProbe *HealthProbe  `json:"readinessProbe,omitempty"`
	volumeMounts   *VolumeMounts `json:"volumeMounts,omitempty"`
	// Deprecated field, use volumeMounts instead.
	volumes []Volumes `json:"volumes,omitempty"`
}

// NewWebserviceSpecWith instantiates a new WebserviceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebserviceSpecWith() *WebserviceSpec {
	this := WebserviceSpec{}
	var addRevisionLabel bool = false
	this.addRevisionLabel = &addRevisionLabel
	var exposeType string = "ClusterIP"
	this.exposeType = &exposeType
	return &this
}

// NewWebserviceSpec instantiates a new WebserviceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebserviceSpec() *WebserviceSpec {
	this := WebserviceSpec{}
	var addRevisionLabel bool = false
	this.addRevisionLabel = &addRevisionLabel
	var exposeType string = "ClusterIP"
	this.exposeType = &exposeType
	return &this
}

// GetAddRevisionLabel returns the AddRevisionLabel field value if set, zero value otherwise.
func (o *WebserviceComponent) GetAddRevisionLabel() bool {
	if o == nil || utils.IsNil(o.Properties.addRevisionLabel) {
		var ret bool
		return ret
	}
	return *o.Properties.addRevisionLabel
}

// GetAddRevisionLabelOk returns a tuple with the AddRevisionLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetAddRevisionLabelOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.addRevisionLabel) {
		return nil, false
	}
	return o.Properties.addRevisionLabel, true
}

// HasAddRevisionLabel returns a boolean if a field has been set.
func (o *WebserviceComponent) HasAddRevisionLabel() bool {
	if o != nil && !utils.IsNil(o.Properties.addRevisionLabel) {
		return true
	}

	return false
}

// AddRevisionLabel gets a reference to the given bool and assigns it to the addRevisionLabel field.
// addRevisionLabel:
func (o *WebserviceComponent) AddRevisionLabel(v bool) *WebserviceComponent {
	o.Properties.addRevisionLabel = &v
	return o
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *WebserviceComponent) GetAnnotations() map[string]string {
	if o == nil || utils.IsNil(o.Properties.annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.annotations) {
		return nil, false
	}
	return o.Properties.annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *WebserviceComponent) HasAnnotations() bool {
	if o != nil && !utils.IsNil(o.Properties.annotations) {
		return true
	}

	return false
}

// Annotations gets a reference to the given map[string]string and assigns it to the annotations field.
// annotations:  Specify the annotations in the workload
func (o *WebserviceComponent) Annotations(v map[string]string) *WebserviceComponent {
	o.Properties.annotations = &v
	return o
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *WebserviceComponent) GetArgs() []string {
	if o == nil || utils.IsNil(o.Properties.args) {
		var ret []string
		return ret
	}
	return o.Properties.args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetArgsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.args) {
		return nil, false
	}
	return o.Properties.args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *WebserviceComponent) HasArgs() bool {
	if o != nil && !utils.IsNil(o.Properties.args) {
		return true
	}

	return false
}

// Args gets a reference to the given []string and assigns it to the args field.
// args:  Arguments to the entrypoint
func (o *WebserviceComponent) Args(v []string) *WebserviceComponent {
	o.Properties.args = v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *WebserviceComponent) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		var ret []string
		return ret
	}
	return o.Properties.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		return nil, false
	}
	return o.Properties.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *WebserviceComponent) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.cmd) {
		return true
	}

	return false
}

// Cmd gets a reference to the given []string and assigns it to the cmd field.
// cmd:  Commands to run in the container
func (o *WebserviceComponent) Cmd(v []string) *WebserviceComponent {
	o.Properties.cmd = v
	return o
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *WebserviceComponent) GetCpu() string {
	if o == nil || utils.IsNil(o.Properties.cpu) {
		var ret string
		return ret
	}
	return *o.Properties.cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.cpu) {
		return nil, false
	}
	return o.Properties.cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *WebserviceComponent) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.cpu) {
		return true
	}

	return false
}

// Cpu gets a reference to the given string and assigns it to the cpu field.
// cpu:  Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (o *WebserviceComponent) Cpu(v string) *WebserviceComponent {
	o.Properties.cpu = &v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *WebserviceComponent) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.env) {
		var ret []Env
		return ret
	}
	return o.Properties.env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.env) {
		return nil, false
	}
	return o.Properties.env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *WebserviceComponent) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.env) {
		return true
	}

	return false
}

// Env gets a reference to the given []Env and assigns it to the env field.
// env:  Define arguments by using environment variables
func (o *WebserviceComponent) Env(v []Env) *WebserviceComponent {
	o.Properties.env = v
	return o
}

// GetExposeType returns the ExposeType field value if set, zero value otherwise.
func (o *WebserviceComponent) GetExposeType() string {
	if o == nil || utils.IsNil(o.Properties.exposeType) {
		var ret string
		return ret
	}
	return *o.Properties.exposeType
}

// GetExposeTypeOk returns a tuple with the ExposeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetExposeTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.exposeType) {
		return nil, false
	}
	return o.Properties.exposeType, true
}

// HasExposeType returns a boolean if a field has been set.
func (o *WebserviceComponent) HasExposeType() bool {
	if o != nil && !utils.IsNil(o.Properties.exposeType) {
		return true
	}

	return false
}

// ExposeType gets a reference to the given string and assigns it to the exposeType field.
// exposeType:
func (o *WebserviceComponent) ExposeType(v string) *WebserviceComponent {
	o.Properties.exposeType = &v
	return o
}

// GetHostAliases returns the HostAliases field value if set, zero value otherwise.
func (o *WebserviceComponent) GetHostAliases() []HostAliases {
	if o == nil || utils.IsNil(o.Properties.hostAliases) {
		var ret []HostAliases
		return ret
	}
	return o.Properties.hostAliases
}

// GetHostAliasesOk returns a tuple with the HostAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetHostAliasesOk() ([]HostAliases, bool) {
	if o == nil || utils.IsNil(o.Properties.hostAliases) {
		return nil, false
	}
	return o.Properties.hostAliases, true
}

// HasHostAliases returns a boolean if a field has been set.
func (o *WebserviceComponent) HasHostAliases() bool {
	if o != nil && !utils.IsNil(o.Properties.hostAliases) {
		return true
	}

	return false
}

// HostAliases gets a reference to the given []HostAliases and assigns it to the hostAliases field.
// hostAliases:  Specify the hostAliases to add
func (o *WebserviceComponent) HostAliases(v []HostAliases) *WebserviceComponent {
	o.Properties.hostAliases = v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *WebserviceComponent) GetImage() string {
	if o == nil || utils.IsNil(o.Properties.image) {
		var ret string
		return ret
	}
	return *o.Properties.image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetImageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.image) {
		return nil, false
	}
	return o.Properties.image, true
}

// HasImage returns a boolean if a field has been set.
func (o *WebserviceComponent) HasImage() bool {
	if o != nil && !utils.IsNil(o.Properties.image) {
		return true
	}

	return false
}

// Image gets a reference to the given string and assigns it to the image field.
// image:  Which image would you like to use for your service +short=i
func (o *WebserviceComponent) Image(v string) *WebserviceComponent {
	o.Properties.image = &v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *WebserviceComponent) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.Properties.imagePullPolicy) {
		var ret string
		return ret
	}
	return *o.Properties.imagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.imagePullPolicy) {
		return nil, false
	}
	return o.Properties.imagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *WebserviceComponent) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.imagePullPolicy) {
		return true
	}

	return false
}

// ImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// imagePullPolicy:  Specify image pull policy for your service
func (o *WebserviceComponent) ImagePullPolicy(v string) *WebserviceComponent {
	o.Properties.imagePullPolicy = &v
	return o
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *WebserviceComponent) GetImagePullSecrets() []string {
	if o == nil || utils.IsNil(o.Properties.imagePullSecrets) {
		var ret []string
		return ret
	}
	return o.Properties.imagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetImagePullSecretsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.imagePullSecrets) {
		return nil, false
	}
	return o.Properties.imagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *WebserviceComponent) HasImagePullSecrets() bool {
	if o != nil && !utils.IsNil(o.Properties.imagePullSecrets) {
		return true
	}

	return false
}

// ImagePullSecrets gets a reference to the given []string and assigns it to the imagePullSecrets field.
// imagePullSecrets:  Specify image pull secrets for your service
func (o *WebserviceComponent) ImagePullSecrets(v []string) *WebserviceComponent {
	o.Properties.imagePullSecrets = v
	return o
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *WebserviceComponent) GetLabels() map[string]string {
	if o == nil || utils.IsNil(o.Properties.labels) {
		var ret map[string]string
		return ret
	}
	return *o.Properties.labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.labels) {
		return nil, false
	}
	return o.Properties.labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebserviceComponent) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Properties.labels) {
		return true
	}

	return false
}

// Labels gets a reference to the given map[string]string and assigns it to the labels field.
// labels:  Specify the labels in the workload
func (o *WebserviceComponent) Labels(v map[string]string) *WebserviceComponent {
	o.Properties.labels = &v
	return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *WebserviceComponent) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.livenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.livenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.livenessProbe) {
		return nil, false
	}
	return o.Properties.livenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *WebserviceComponent) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.livenessProbe) {
		return true
	}

	return false
}

// LivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// livenessProbe:
func (o *WebserviceComponent) LivenessProbe(v HealthProbe) *WebserviceComponent {
	o.Properties.livenessProbe = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *WebserviceComponent) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.memory) {
		var ret string
		return ret
	}
	return *o.Properties.memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.memory) {
		return nil, false
	}
	return o.Properties.memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *WebserviceComponent) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.memory) {
		return true
	}

	return false
}

// Memory gets a reference to the given string and assigns it to the memory field.
// memory:  Specifies the attributes of the memory resource required for the container.
func (o *WebserviceComponent) Memory(v string) *WebserviceComponent {
	o.Properties.memory = &v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *WebserviceComponent) GetPort() int32 {
	if o == nil || utils.IsNil(o.Properties.port) {
		var ret int32
		return ret
	}
	return *o.Properties.port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.port) {
		return nil, false
	}
	return o.Properties.port, true
}

// HasPort returns a boolean if a field has been set.
func (o *WebserviceComponent) HasPort() bool {
	if o != nil && !utils.IsNil(o.Properties.port) {
		return true
	}

	return false
}

// Port gets a reference to the given int32 and assigns it to the port field.
// port:
func (o *WebserviceComponent) Port(v int32) *WebserviceComponent {
	o.Properties.port = &v
	return o
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *WebserviceComponent) GetPorts() []Ports {
	if o == nil || utils.IsNil(o.Properties.ports) {
		var ret []Ports
		return ret
	}
	return o.Properties.ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetPortsOk() ([]Ports, bool) {
	if o == nil || utils.IsNil(o.Properties.ports) {
		return nil, false
	}
	return o.Properties.ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *WebserviceComponent) HasPorts() bool {
	if o != nil && !utils.IsNil(o.Properties.ports) {
		return true
	}

	return false
}

// Ports gets a reference to the given []Ports and assigns it to the ports field.
// ports:  Which ports do you want customer traffic sent to, defaults to 80
func (o *WebserviceComponent) Ports(v []Ports) *WebserviceComponent {
	o.Properties.ports = v
	return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *WebserviceComponent) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.readinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.readinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.readinessProbe) {
		return nil, false
	}
	return o.Properties.readinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *WebserviceComponent) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.readinessProbe) {
		return true
	}

	return false
}

// ReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// readinessProbe:
func (o *WebserviceComponent) ReadinessProbe(v HealthProbe) *WebserviceComponent {
	o.Properties.readinessProbe = &v
	return o
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *WebserviceComponent) GetVolumeMounts() VolumeMounts {
	if o == nil || utils.IsNil(o.Properties.volumeMounts) {
		var ret VolumeMounts
		return ret
	}
	return *o.Properties.volumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetVolumeMountsOk() (*VolumeMounts, bool) {
	if o == nil || utils.IsNil(o.Properties.volumeMounts) {
		return nil, false
	}
	return o.Properties.volumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *WebserviceComponent) HasVolumeMounts() bool {
	if o != nil && !utils.IsNil(o.Properties.volumeMounts) {
		return true
	}

	return false
}

// VolumeMounts gets a reference to the given VolumeMounts and assigns it to the volumeMounts field.
// volumeMounts:
func (o *WebserviceComponent) VolumeMounts(v VolumeMounts) *WebserviceComponent {
	o.Properties.volumeMounts = &v
	return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *WebserviceComponent) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.Properties.volumes) {
		var ret []Volumes
		return ret
	}
	return o.Properties.volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebserviceComponent) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.Properties.volumes) {
		return nil, false
	}
	return o.Properties.volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *WebserviceComponent) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.Properties.volumes) {
		return true
	}

	return false
}

// Volumes gets a reference to the given []Volumes and assigns it to the volumes field.
// volumes:  Deprecated field, use volumeMounts instead.
func (o *WebserviceComponent) Volumes(v []Volumes) *WebserviceComponent {
	o.Properties.volumes = v
	return o
}

func (o WebserviceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebserviceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.addRevisionLabel) {
		toSerialize["addRevisionLabel"] = o.addRevisionLabel
	}
	if !utils.IsNil(o.annotations) {
		toSerialize["annotations"] = o.annotations
	}
	if !utils.IsNil(o.args) {
		toSerialize["args"] = o.args
	}
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	if !utils.IsNil(o.cpu) {
		toSerialize["cpu"] = o.cpu
	}
	if !utils.IsNil(o.env) {
		toSerialize["env"] = o.env
	}
	if !utils.IsNil(o.exposeType) {
		toSerialize["exposeType"] = o.exposeType
	}
	if !utils.IsNil(o.hostAliases) {
		toSerialize["hostAliases"] = o.hostAliases
	}
	if !utils.IsNil(o.image) {
		toSerialize["image"] = o.image
	}
	if !utils.IsNil(o.imagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.imagePullPolicy
	}
	if !utils.IsNil(o.imagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.imagePullSecrets
	}
	if !utils.IsNil(o.labels) {
		toSerialize["labels"] = o.labels
	}
	if !utils.IsNil(o.livenessProbe) {
		toSerialize["livenessProbe"] = o.livenessProbe
	}
	if !utils.IsNil(o.memory) {
		toSerialize["memory"] = o.memory
	}
	if !utils.IsNil(o.port) {
		toSerialize["port"] = o.port
	}
	if !utils.IsNil(o.ports) {
		toSerialize["ports"] = o.ports
	}
	if !utils.IsNil(o.readinessProbe) {
		toSerialize["readinessProbe"] = o.readinessProbe
	}
	if !utils.IsNil(o.volumeMounts) {
		toSerialize["volumeMounts"] = o.volumeMounts
	}
	if !utils.IsNil(o.volumes) {
		toSerialize["volumes"] = o.volumes
	}
	return toSerialize, nil
}

type NullableWebserviceSpec struct {
	value *WebserviceSpec
	isSet bool
}

func (v NullableWebserviceSpec) Get() *WebserviceSpec {
	return v.value
}

func (v *NullableWebserviceSpec) Set(val *WebserviceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWebserviceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWebserviceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebserviceSpec(val *WebserviceSpec) *NullableWebserviceSpec {
	return &NullableWebserviceSpec{value: val, isSet: true}
}

func (v NullableWebserviceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebserviceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const WebserviceType = "webservice"

func init() {
	sdkcommon.RegisterComponent(WebserviceType, FromComponent)
}

type WebserviceComponent struct {
	Base       apis.ComponentBase
	Properties WebserviceSpec
}

func Webservice(name string) *WebserviceComponent {
	w := &WebserviceComponent{Base: apis.ComponentBase{
		Name: name,
		Type: WebserviceType,
	}}
	return w
}

func (w *WebserviceComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range w.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  w.Base.DependsOn,
		Inputs:     w.Base.Inputs,
		Name:       w.Base.Name,
		Outputs:    w.Base.Outputs,
		Properties: util.Object2RawExtension(w.Properties),
		Traits:     traits,
		Type:       WebserviceType,
	}
	return res
}

func (w *WebserviceComponent) FromComponent(from common.ApplicationComponent) (*WebserviceComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		w.Base.Traits = append(w.Base.Traits, _t)
	}
	var properties WebserviceSpec
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
	w.Properties = properties
	return w, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	w := &WebserviceComponent{}
	return w.FromComponent(from)
}

func (w *WebserviceComponent) AddTrait(traits ...apis.Trait) *WebserviceComponent {
	w.Base.Traits = append(w.Base.Traits, traits...)
	return w
}

func (w *WebserviceComponent) DefName() string {
	return w.Base.Name
}

func (w *WebserviceComponent) DefType() string {
	return WebserviceType
}
