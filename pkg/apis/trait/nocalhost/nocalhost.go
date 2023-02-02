/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nocalhost

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the NocalhostSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NocalhostSpec{}

// NocalhostSpec struct for NocalhostSpec
type NocalhostSpec struct {
	Command              *Command               `json:"command,omitempty"`
	Debug                *Debug                 `json:"debug,omitempty"`
	Env                  []Env                  `json:"env,omitempty"`
	GitUrl               *string                `json:"gitUrl,omitempty"`
	HotReload            *bool                  `json:"hotReload,omitempty"`
	Image                *Image                 `json:"image,omitempty"`
	PersistentVolumeDirs []PersistentVolumeDirs `json:"persistentVolumeDirs,omitempty"`
	Port                 *int32                 `json:"port,omitempty"`
	PortForward          []string               `json:"portForward,omitempty"`
	Resources            *Resources             `json:"resources,omitempty"`
	ServiceType          *string                `json:"serviceType,omitempty"`
	Shell                *string                `json:"shell,omitempty"`
	StorageClass         *string                `json:"storageClass,omitempty"`
	Sync                 *Sync                  `json:"sync,omitempty"`
	WorkDir              *string                `json:"workDir,omitempty"`
}

// NewNocalhostSpecWith instantiates a new NocalhostSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNocalhostSpecWith() *NocalhostSpec {
	this := NocalhostSpec{}
	var hotReload bool = true
	this.HotReload = &hotReload
	var serviceType string = "deployment"
	this.ServiceType = &serviceType
	var shell string = "bash"
	this.Shell = &shell
	var workDir string = "/home/nocalhost-dev"
	this.WorkDir = &workDir
	return &this
}

// NewNocalhostSpec instantiates a new NocalhostSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNocalhostSpec() *NocalhostSpec {
	this := NocalhostSpec{}
	var hotReload bool = true
	this.HotReload = &hotReload
	var serviceType string = "deployment"
	this.ServiceType = &serviceType
	var shell string = "bash"
	this.Shell = &shell
	var workDir string = "/home/nocalhost-dev"
	this.WorkDir = &workDir
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *NocalhostTrait) GetCommand() Command {
	if o == nil || utils.IsNil(o.Properties.Command) {
		var ret Command
		return ret
	}
	return *o.Properties.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetCommandOk() (*Command, bool) {
	if o == nil || utils.IsNil(o.Properties.Command) {
		return nil, false
	}
	return o.Properties.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *NocalhostTrait) HasCommand() bool {
	if o != nil && !utils.IsNil(o.Properties.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given Command and assigns it to the command field.
// Command:
func (o *NocalhostTrait) SetCommand(v Command) *NocalhostTrait {
	o.Properties.Command = &v
	return o
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *NocalhostTrait) GetDebug() Debug {
	if o == nil || utils.IsNil(o.Properties.Debug) {
		var ret Debug
		return ret
	}
	return *o.Properties.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetDebugOk() (*Debug, bool) {
	if o == nil || utils.IsNil(o.Properties.Debug) {
		return nil, false
	}
	return o.Properties.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *NocalhostTrait) HasDebug() bool {
	if o != nil && !utils.IsNil(o.Properties.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given Debug and assigns it to the debug field.
// Debug:
func (o *NocalhostTrait) SetDebug(v Debug) *NocalhostTrait {
	o.Properties.Debug = &v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *NocalhostTrait) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret []Env
		return ret
	}
	return o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *NocalhostTrait) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// Env:
func (o *NocalhostTrait) SetEnv(v []Env) *NocalhostTrait {
	o.Properties.Env = v
	return o
}

// GetGitUrl returns the GitUrl field value if set, zero value otherwise.
func (o *NocalhostTrait) GetGitUrl() string {
	if o == nil || utils.IsNil(o.Properties.GitUrl) {
		var ret string
		return ret
	}
	return *o.Properties.GitUrl
}

// GetGitUrlOk returns a tuple with the GitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetGitUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.GitUrl) {
		return nil, false
	}
	return o.Properties.GitUrl, true
}

// HasGitUrl returns a boolean if a field has been set.
func (o *NocalhostTrait) HasGitUrl() bool {
	if o != nil && !utils.IsNil(o.Properties.GitUrl) {
		return true
	}

	return false
}

// SetGitUrl gets a reference to the given string and assigns it to the gitUrl field.
// GitUrl:
func (o *NocalhostTrait) SetGitUrl(v string) *NocalhostTrait {
	o.Properties.GitUrl = &v
	return o
}

// GetHotReload returns the HotReload field value if set, zero value otherwise.
func (o *NocalhostTrait) GetHotReload() bool {
	if o == nil || utils.IsNil(o.Properties.HotReload) {
		var ret bool
		return ret
	}
	return *o.Properties.HotReload
}

// GetHotReloadOk returns a tuple with the HotReload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetHotReloadOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Properties.HotReload) {
		return nil, false
	}
	return o.Properties.HotReload, true
}

// HasHotReload returns a boolean if a field has been set.
func (o *NocalhostTrait) HasHotReload() bool {
	if o != nil && !utils.IsNil(o.Properties.HotReload) {
		return true
	}

	return false
}

// SetHotReload gets a reference to the given bool and assigns it to the hotReload field.
// HotReload:
func (o *NocalhostTrait) SetHotReload(v bool) *NocalhostTrait {
	o.Properties.HotReload = &v
	return o
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *NocalhostTrait) GetImage() Image {
	if o == nil || utils.IsNil(o.Properties.Image) {
		var ret Image
		return ret
	}
	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetImageOk() (*Image, bool) {
	if o == nil || utils.IsNil(o.Properties.Image) {
		return nil, false
	}
	return o.Properties.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *NocalhostTrait) HasImage() bool {
	if o != nil && !utils.IsNil(o.Properties.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the image field.
// Image:
func (o *NocalhostTrait) SetImage(v Image) *NocalhostTrait {
	o.Properties.Image = &v
	return o
}

// GetPersistentVolumeDirs returns the PersistentVolumeDirs field value if set, zero value otherwise.
func (o *NocalhostTrait) GetPersistentVolumeDirs() []PersistentVolumeDirs {
	if o == nil || utils.IsNil(o.Properties.PersistentVolumeDirs) {
		var ret []PersistentVolumeDirs
		return ret
	}
	return o.Properties.PersistentVolumeDirs
}

// GetPersistentVolumeDirsOk returns a tuple with the PersistentVolumeDirs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetPersistentVolumeDirsOk() ([]PersistentVolumeDirs, bool) {
	if o == nil || utils.IsNil(o.Properties.PersistentVolumeDirs) {
		return nil, false
	}
	return o.Properties.PersistentVolumeDirs, true
}

// HasPersistentVolumeDirs returns a boolean if a field has been set.
func (o *NocalhostTrait) HasPersistentVolumeDirs() bool {
	if o != nil && !utils.IsNil(o.Properties.PersistentVolumeDirs) {
		return true
	}

	return false
}

// SetPersistentVolumeDirs gets a reference to the given []PersistentVolumeDirs and assigns it to the persistentVolumeDirs field.
// PersistentVolumeDirs:
func (o *NocalhostTrait) SetPersistentVolumeDirs(v []PersistentVolumeDirs) *NocalhostTrait {
	o.Properties.PersistentVolumeDirs = v
	return o
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NocalhostTrait) GetPort() int32 {
	if o == nil || utils.IsNil(o.Properties.Port) {
		var ret int32
		return ret
	}
	return *o.Properties.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Port) {
		return nil, false
	}
	return o.Properties.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NocalhostTrait) HasPort() bool {
	if o != nil && !utils.IsNil(o.Properties.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the port field.
// Port:
func (o *NocalhostTrait) SetPort(v int32) *NocalhostTrait {
	o.Properties.Port = &v
	return o
}

// GetPortForward returns the PortForward field value if set, zero value otherwise.
func (o *NocalhostTrait) GetPortForward() []string {
	if o == nil || utils.IsNil(o.Properties.PortForward) {
		var ret []string
		return ret
	}
	return o.Properties.PortForward
}

// GetPortForwardOk returns a tuple with the PortForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetPortForwardOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.PortForward) {
		return nil, false
	}
	return o.Properties.PortForward, true
}

// HasPortForward returns a boolean if a field has been set.
func (o *NocalhostTrait) HasPortForward() bool {
	if o != nil && !utils.IsNil(o.Properties.PortForward) {
		return true
	}

	return false
}

// SetPortForward gets a reference to the given []string and assigns it to the portForward field.
// PortForward:
func (o *NocalhostTrait) SetPortForward(v []string) *NocalhostTrait {
	o.Properties.PortForward = v
	return o
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *NocalhostTrait) GetResources() Resources {
	if o == nil || utils.IsNil(o.Properties.Resources) {
		var ret Resources
		return ret
	}
	return *o.Properties.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetResourcesOk() (*Resources, bool) {
	if o == nil || utils.IsNil(o.Properties.Resources) {
		return nil, false
	}
	return o.Properties.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *NocalhostTrait) HasResources() bool {
	if o != nil && !utils.IsNil(o.Properties.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given Resources and assigns it to the resources field.
// Resources:
func (o *NocalhostTrait) SetResources(v Resources) *NocalhostTrait {
	o.Properties.Resources = &v
	return o
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *NocalhostTrait) GetServiceType() string {
	if o == nil || utils.IsNil(o.Properties.ServiceType) {
		var ret string
		return ret
	}
	return *o.Properties.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetServiceTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ServiceType) {
		return nil, false
	}
	return o.Properties.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *NocalhostTrait) HasServiceType() bool {
	if o != nil && !utils.IsNil(o.Properties.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the serviceType field.
// ServiceType:
func (o *NocalhostTrait) SetServiceType(v string) *NocalhostTrait {
	o.Properties.ServiceType = &v
	return o
}

// GetShell returns the Shell field value if set, zero value otherwise.
func (o *NocalhostTrait) GetShell() string {
	if o == nil || utils.IsNil(o.Properties.Shell) {
		var ret string
		return ret
	}
	return *o.Properties.Shell
}

// GetShellOk returns a tuple with the Shell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetShellOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Shell) {
		return nil, false
	}
	return o.Properties.Shell, true
}

// HasShell returns a boolean if a field has been set.
func (o *NocalhostTrait) HasShell() bool {
	if o != nil && !utils.IsNil(o.Properties.Shell) {
		return true
	}

	return false
}

// SetShell gets a reference to the given string and assigns it to the shell field.
// Shell:
func (o *NocalhostTrait) SetShell(v string) *NocalhostTrait {
	o.Properties.Shell = &v
	return o
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *NocalhostTrait) GetStorageClass() string {
	if o == nil || utils.IsNil(o.Properties.StorageClass) {
		var ret string
		return ret
	}
	return *o.Properties.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetStorageClassOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.StorageClass) {
		return nil, false
	}
	return o.Properties.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *NocalhostTrait) HasStorageClass() bool {
	if o != nil && !utils.IsNil(o.Properties.StorageClass) {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given string and assigns it to the storageClass field.
// StorageClass:
func (o *NocalhostTrait) SetStorageClass(v string) *NocalhostTrait {
	o.Properties.StorageClass = &v
	return o
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *NocalhostTrait) GetSync() Sync {
	if o == nil || utils.IsNil(o.Properties.Sync) {
		var ret Sync
		return ret
	}
	return *o.Properties.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetSyncOk() (*Sync, bool) {
	if o == nil || utils.IsNil(o.Properties.Sync) {
		return nil, false
	}
	return o.Properties.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *NocalhostTrait) HasSync() bool {
	if o != nil && !utils.IsNil(o.Properties.Sync) {
		return true
	}

	return false
}

// SetSync gets a reference to the given Sync and assigns it to the sync field.
// Sync:
func (o *NocalhostTrait) SetSync(v Sync) *NocalhostTrait {
	o.Properties.Sync = &v
	return o
}

// GetWorkDir returns the WorkDir field value if set, zero value otherwise.
func (o *NocalhostTrait) GetWorkDir() string {
	if o == nil || utils.IsNil(o.Properties.WorkDir) {
		var ret string
		return ret
	}
	return *o.Properties.WorkDir
}

// GetWorkDirOk returns a tuple with the WorkDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NocalhostTrait) GetWorkDirOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.WorkDir) {
		return nil, false
	}
	return o.Properties.WorkDir, true
}

// HasWorkDir returns a boolean if a field has been set.
func (o *NocalhostTrait) HasWorkDir() bool {
	if o != nil && !utils.IsNil(o.Properties.WorkDir) {
		return true
	}

	return false
}

// SetWorkDir gets a reference to the given string and assigns it to the workDir field.
// WorkDir:
func (o *NocalhostTrait) SetWorkDir(v string) *NocalhostTrait {
	o.Properties.WorkDir = &v
	return o
}

func (o NocalhostSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NocalhostSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !utils.IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !utils.IsNil(o.GitUrl) {
		toSerialize["gitUrl"] = o.GitUrl
	}
	if !utils.IsNil(o.HotReload) {
		toSerialize["hotReload"] = o.HotReload
	}
	if !utils.IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !utils.IsNil(o.PersistentVolumeDirs) {
		toSerialize["persistentVolumeDirs"] = o.PersistentVolumeDirs
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.PortForward) {
		toSerialize["portForward"] = o.PortForward
	}
	if !utils.IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !utils.IsNil(o.ServiceType) {
		toSerialize["serviceType"] = o.ServiceType
	}
	if !utils.IsNil(o.Shell) {
		toSerialize["shell"] = o.Shell
	}
	if !utils.IsNil(o.StorageClass) {
		toSerialize["storageClass"] = o.StorageClass
	}
	if !utils.IsNil(o.Sync) {
		toSerialize["sync"] = o.Sync
	}
	if !utils.IsNil(o.WorkDir) {
		toSerialize["workDir"] = o.WorkDir
	}
	return toSerialize, nil
}

type NullableNocalhostSpec struct {
	value *NocalhostSpec
	isSet bool
}

func (v NullableNocalhostSpec) Get() *NocalhostSpec {
	return v.value
}

func (v *NullableNocalhostSpec) Set(val *NocalhostSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNocalhostSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNocalhostSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNocalhostSpec(val *NocalhostSpec) *NullableNocalhostSpec {
	return &NullableNocalhostSpec{value: val, isSet: true}
}

func (v NullableNocalhostSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNocalhostSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const NocalhostType = "nocalhost"

func init() {
	sdkcommon.RegisterTrait(NocalhostType, FromTrait)
}

type NocalhostTrait struct {
	Base       apis.TraitBase
	Properties NocalhostSpec
}

func Nocalhost() *NocalhostTrait {
	n := &NocalhostTrait{Base: apis.TraitBase{}}
	return n
}

func (n *NocalhostTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(n.Properties),
		Type:       NocalhostType,
	}
	return res
}

func (n *NocalhostTrait) FromTrait(from common.ApplicationTrait) (*NocalhostTrait, error) {
	var properties NocalhostSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	n.Properties = properties
	return n, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	n := &NocalhostTrait{}
	return n.FromTrait(from)
}

func (n *NocalhostTrait) DefType() string {
	return NocalhostType
}
