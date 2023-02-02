/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package init_container

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the InitContainerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InitContainerSpec{}

// InitContainerSpec struct for InitContainerSpec
type InitContainerSpec struct {
	// Specify the mount path of app container
	appMountPath string `json:"appMountPath"`
	// Specify the args run in the init container
	args []string `json:"args,omitempty"`
	// Specify the commands run in the init container
	cmd []string `json:"cmd,omitempty"`
	// Specify the env run in the init container
	env []Env `json:"env,omitempty"`
	// Specify the extra volume mounts for the init container
	extraVolumeMounts []ExtraVolumeMounts `json:"extraVolumeMounts"`
	// Specify the image of init container
	image string `json:"image"`
	// Specify image pull policy for your service
	imagePullPolicy string `json:"imagePullPolicy"`
	// Specify the mount path of init container
	initMountPath string `json:"initMountPath"`
	// Specify the mount name of shared volume
	mountName string `json:"mountName"`
	// Specify the name of init container
	name string `json:"name"`
}

// NewInitContainerSpecWith instantiates a new InitContainerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitContainerSpecWith(appMountPath string, extraVolumeMounts []ExtraVolumeMounts, image string, imagePullPolicy string, initMountPath string, mountName string, name string) *InitContainerSpec {
	this := InitContainerSpec{}
	this.appMountPath = appMountPath
	this.extraVolumeMounts = extraVolumeMounts
	this.image = image
	this.imagePullPolicy = imagePullPolicy
	this.initMountPath = initMountPath
	this.mountName = mountName
	this.name = name
	return &this
}

// NewInitContainerSpec instantiates a new InitContainerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitContainerSpec() *InitContainerSpec {
	this := InitContainerSpec{}
	var imagePullPolicy string = "IfNotPresent"
	this.imagePullPolicy = imagePullPolicy
	var mountName string = "workdir"
	this.mountName = mountName
	return &this
}

// GetAppMountPath returns the AppMountPath field value
func (o *InitContainerTrait) GetAppMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.appMountPath
}

// GetAppMountPathOk returns a tuple with the AppMountPath field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetAppMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.appMountPath, true
}

// AppMountPath sets field value
func (o *InitContainerTrait) AppMountPath(v string) *InitContainerTrait {
	o.Properties.appMountPath = v
	return o
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *InitContainerTrait) GetArgs() []string {
	if o == nil || utils.IsNil(o.Properties.args) {
		var ret []string
		return ret
	}
	return o.Properties.args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetArgsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.args) {
		return nil, false
	}
	return o.Properties.args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *InitContainerTrait) HasArgs() bool {
	if o != nil && !utils.IsNil(o.Properties.args) {
		return true
	}

	return false
}

// Args gets a reference to the given []string and assigns it to the args field.
// args:  Specify the args run in the init container
func (o *InitContainerTrait) Args(v []string) *InitContainerTrait {
	o.Properties.args = v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *InitContainerTrait) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		var ret []string
		return ret
	}
	return o.Properties.cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.cmd) {
		return nil, false
	}
	return o.Properties.cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *InitContainerTrait) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.cmd) {
		return true
	}

	return false
}

// Cmd gets a reference to the given []string and assigns it to the cmd field.
// cmd:  Specify the commands run in the init container
func (o *InitContainerTrait) Cmd(v []string) *InitContainerTrait {
	o.Properties.cmd = v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *InitContainerTrait) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.env) {
		var ret []Env
		return ret
	}
	return o.Properties.env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.env) {
		return nil, false
	}
	return o.Properties.env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *InitContainerTrait) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.env) {
		return true
	}

	return false
}

// Env gets a reference to the given []Env and assigns it to the env field.
// env:  Specify the env run in the init container
func (o *InitContainerTrait) Env(v []Env) *InitContainerTrait {
	o.Properties.env = v
	return o
}

// GetExtraVolumeMounts returns the ExtraVolumeMounts field value
func (o *InitContainerTrait) GetExtraVolumeMounts() []ExtraVolumeMounts {
	if o == nil {
		var ret []ExtraVolumeMounts
		return ret
	}

	return o.Properties.extraVolumeMounts
}

// GetExtraVolumeMountsOk returns a tuple with the ExtraVolumeMounts field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetExtraVolumeMountsOk() ([]ExtraVolumeMounts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.extraVolumeMounts, true
}

// ExtraVolumeMounts sets field value
func (o *InitContainerTrait) ExtraVolumeMounts(v []ExtraVolumeMounts) *InitContainerTrait {
	o.Properties.extraVolumeMounts = v
	return o
}

// GetImage returns the Image field value
func (o *InitContainerTrait) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.image, true
}

// Image sets field value
func (o *InitContainerTrait) Image(v string) *InitContainerTrait {
	o.Properties.image = v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value
func (o *InitContainerTrait) GetImagePullPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.imagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetImagePullPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.imagePullPolicy, true
}

// ImagePullPolicy sets field value
func (o *InitContainerTrait) ImagePullPolicy(v string) *InitContainerTrait {
	o.Properties.imagePullPolicy = v
	return o
}

// GetInitMountPath returns the InitMountPath field value
func (o *InitContainerTrait) GetInitMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.initMountPath
}

// GetInitMountPathOk returns a tuple with the InitMountPath field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetInitMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.initMountPath, true
}

// InitMountPath sets field value
func (o *InitContainerTrait) InitMountPath(v string) *InitContainerTrait {
	o.Properties.initMountPath = v
	return o
}

// GetMountName returns the MountName field value
func (o *InitContainerTrait) GetMountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.mountName
}

// GetMountNameOk returns a tuple with the MountName field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetMountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.mountName, true
}

// MountName sets field value
func (o *InitContainerTrait) MountName(v string) *InitContainerTrait {
	o.Properties.mountName = v
	return o
}

// GetName returns the Name field value
func (o *InitContainerTrait) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InitContainerTrait) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.name, true
}

// Name sets field value
func (o *InitContainerTrait) Name(v string) *InitContainerTrait {
	o.Properties.name = v
	return o
}

func (o InitContainerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitContainerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appMountPath"] = o.appMountPath
	if !utils.IsNil(o.args) {
		toSerialize["args"] = o.args
	}
	if !utils.IsNil(o.cmd) {
		toSerialize["cmd"] = o.cmd
	}
	if !utils.IsNil(o.env) {
		toSerialize["env"] = o.env
	}
	toSerialize["extraVolumeMounts"] = o.extraVolumeMounts
	toSerialize["image"] = o.image
	toSerialize["imagePullPolicy"] = o.imagePullPolicy
	toSerialize["initMountPath"] = o.initMountPath
	toSerialize["mountName"] = o.mountName
	toSerialize["name"] = o.name
	return toSerialize, nil
}

type NullableInitContainerSpec struct {
	value *InitContainerSpec
	isSet bool
}

func (v NullableInitContainerSpec) Get() *InitContainerSpec {
	return v.value
}

func (v *NullableInitContainerSpec) Set(val *InitContainerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableInitContainerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableInitContainerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitContainerSpec(val *InitContainerSpec) *NullableInitContainerSpec {
	return &NullableInitContainerSpec{value: val, isSet: true}
}

func (v NullableInitContainerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitContainerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const InitContainerType = "init-container"

func init() {
	sdkcommon.RegisterTrait(InitContainerType, FromTrait)
}

type InitContainerTrait struct {
	Base       apis.TraitBase
	Properties InitContainerSpec
}

func InitContainer() *InitContainerTrait {
	i := &InitContainerTrait{Base: apis.TraitBase{}}
	return i
}

func (i *InitContainerTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(i.Properties),
		Type:       InitContainerType,
	}
	return res
}

func (i *InitContainerTrait) FromTrait(from common.ApplicationTrait) (*InitContainerTrait, error) {
	var properties InitContainerSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	i.Properties = properties
	return i, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	i := &InitContainerTrait{}
	return i.FromTrait(from)
}

func (i *InitContainerTrait) Type() string {
	return InitContainerType
}
