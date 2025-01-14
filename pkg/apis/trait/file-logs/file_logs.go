/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package file_logs

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the FileLogsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FileLogsSpec{}

// FileLogsSpec struct for FileLogsSpec
type FileLogsSpec struct {
	Path string `json:"path"`
}

// NewFileLogsSpecWith instantiates a new FileLogsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLogsSpecWith(path string) *FileLogsSpec {
	this := FileLogsSpec{}
	this.Path = path
	return &this
}

// NewFileLogsSpec instantiates a new FileLogsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLogsSpec() *FileLogsSpec {
	this := FileLogsSpec{}
	return &this
}

// NewFileLogsSpecs converts a list FileLogsSpec pointers to objects.
// This is helpful when the SetFileLogsSpec requires a list of objects
func NewFileLogsSpecs(ps ...*FileLogsSpec) []FileLogsSpec {
	objs := []FileLogsSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetPath returns the Path field value
func (o *FileLogsTrait) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FileLogsTrait) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties.Path, true
}

// SetPath sets field value
func (o *FileLogsTrait) SetPath(v string) *FileLogsTrait {
	o.Properties.Path = v
	return o
}

func (o FileLogsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLogsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullableFileLogsSpec struct {
	value *FileLogsSpec
	isSet bool
}

func (v NullableFileLogsSpec) Get() *FileLogsSpec {
	return v.value
}

func (v *NullableFileLogsSpec) Set(val *FileLogsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLogsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLogsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLogsSpec(val *FileLogsSpec) *NullableFileLogsSpec {
	return &NullableFileLogsSpec{value: val, isSet: true}
}

func (v NullableFileLogsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLogsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const FileLogsType = "file-logs"

func init() {
	sdkcommon.RegisterTrait(FileLogsType, FromTrait)
}

type FileLogsTrait struct {
	Base       apis.TraitBase
	Properties FileLogsSpec
}

func FileLogs() *FileLogsTrait {
	f := &FileLogsTrait{Base: apis.TraitBase{}}
	return f
}

func (f *FileLogsTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(f.Properties),
		Type:       FileLogsType,
	}
	return res
}

func (f *FileLogsTrait) FromTrait(from common.ApplicationTrait) (*FileLogsTrait, error) {
	var properties FileLogsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	f.Base.Type = FileLogsType
	f.Properties = properties
	return f, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	f := &FileLogsTrait{}
	return f.FromTrait(from)
}

func (f *FileLogsTrait) DefType() string {
	return FileLogsType
}
