/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grafana_dashboard

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the GrafanaDashboardSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GrafanaDashboardSpec{}

// GrafanaDashboardSpec struct for GrafanaDashboardSpec
type GrafanaDashboardSpec struct {
	// The json model of the grafana dashboard
	Data *string `json:"data,omitempty"`
	// The name of the grafana instance.
	Grafana *string `json:"grafana,omitempty"`
	// The uid of the grafana dashboard.
	Uid *string `json:"uid,omitempty"`
}

// NewGrafanaDashboardSpecWith instantiates a new GrafanaDashboardSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrafanaDashboardSpecWith() *GrafanaDashboardSpec {
	this := GrafanaDashboardSpec{}
	var grafana string = "default"
	this.Grafana = &grafana
	return &this
}

// NewGrafanaDashboardSpec instantiates a new GrafanaDashboardSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaDashboardSpec() *GrafanaDashboardSpec {
	this := GrafanaDashboardSpec{}
	var grafana string = "default"
	this.Grafana = &grafana
	return &this
}

// NewGrafanaDashboardSpecs converts a list GrafanaDashboardSpec pointers to objects.
// This is helpful when the SetGrafanaDashboardSpec requires a list of objects
func NewGrafanaDashboardSpecs(ps ...*GrafanaDashboardSpec) []GrafanaDashboardSpec {
	objs := []GrafanaDashboardSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GrafanaDashboardComponent) GetData() string {
	if o == nil || utils.IsNil(o.Properties.Data) {
		var ret string
		return ret
	}
	return *o.Properties.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaDashboardComponent) GetDataOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Data) {
		return nil, false
	}
	return o.Properties.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GrafanaDashboardComponent) HasData() bool {
	if o != nil && !utils.IsNil(o.Properties.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the data field.
// Data:  The json model of the grafana dashboard
func (o *GrafanaDashboardComponent) SetData(v string) *GrafanaDashboardComponent {
	o.Properties.Data = &v
	return o
}

// GetGrafana returns the Grafana field value if set, zero value otherwise.
func (o *GrafanaDashboardComponent) GetGrafana() string {
	if o == nil || utils.IsNil(o.Properties.Grafana) {
		var ret string
		return ret
	}
	return *o.Properties.Grafana
}

// GetGrafanaOk returns a tuple with the Grafana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaDashboardComponent) GetGrafanaOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Grafana) {
		return nil, false
	}
	return o.Properties.Grafana, true
}

// HasGrafana returns a boolean if a field has been set.
func (o *GrafanaDashboardComponent) HasGrafana() bool {
	if o != nil && !utils.IsNil(o.Properties.Grafana) {
		return true
	}

	return false
}

// SetGrafana gets a reference to the given string and assigns it to the grafana field.
// Grafana:  The name of the grafana instance.
func (o *GrafanaDashboardComponent) SetGrafana(v string) *GrafanaDashboardComponent {
	o.Properties.Grafana = &v
	return o
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *GrafanaDashboardComponent) GetUid() string {
	if o == nil || utils.IsNil(o.Properties.Uid) {
		var ret string
		return ret
	}
	return *o.Properties.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaDashboardComponent) GetUidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Uid) {
		return nil, false
	}
	return o.Properties.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *GrafanaDashboardComponent) HasUid() bool {
	if o != nil && !utils.IsNil(o.Properties.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the uid field.
// Uid:  The uid of the grafana dashboard.
func (o *GrafanaDashboardComponent) SetUid(v string) *GrafanaDashboardComponent {
	o.Properties.Uid = &v
	return o
}

func (o GrafanaDashboardSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrafanaDashboardSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !utils.IsNil(o.Grafana) {
		toSerialize["grafana"] = o.Grafana
	}
	if !utils.IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableGrafanaDashboardSpec struct {
	value *GrafanaDashboardSpec
	isSet bool
}

func (v NullableGrafanaDashboardSpec) Get() *GrafanaDashboardSpec {
	return v.value
}

func (v *NullableGrafanaDashboardSpec) Set(val *GrafanaDashboardSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGrafanaDashboardSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGrafanaDashboardSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrafanaDashboardSpec(val *GrafanaDashboardSpec) *NullableGrafanaDashboardSpec {
	return &NullableGrafanaDashboardSpec{value: val, isSet: true}
}

func (v NullableGrafanaDashboardSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrafanaDashboardSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const GrafanaDashboardType = "grafana-dashboard"

func init() {
	sdkcommon.RegisterComponent(GrafanaDashboardType, FromComponent)
}

type GrafanaDashboardComponent struct {
	Base       apis.ComponentBase
	Properties GrafanaDashboardSpec
}

func GrafanaDashboard(name string) *GrafanaDashboardComponent {
	g := &GrafanaDashboardComponent{Base: apis.ComponentBase{
		Name: name,
		Type: GrafanaDashboardType,
	}}
	return g
}

func (g *GrafanaDashboardComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range g.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  g.Base.DependsOn,
		Inputs:     g.Base.Inputs,
		Name:       g.Base.Name,
		Outputs:    g.Base.Outputs,
		Properties: util.Object2RawExtension(g.Properties),
		Traits:     traits,
		Type:       GrafanaDashboardType,
	}
	return res
}

func (g *GrafanaDashboardComponent) FromComponent(from common.ApplicationComponent) (*GrafanaDashboardComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		g.Base.Traits = append(g.Base.Traits, _t)
	}
	var properties GrafanaDashboardSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	g.Base.Name = from.Name
	g.Base.DependsOn = from.DependsOn
	g.Base.Inputs = from.Inputs
	g.Base.Outputs = from.Outputs
	g.Base.Type = GrafanaDashboardType
	g.Properties = properties
	return g, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	g := &GrafanaDashboardComponent{}
	return g.FromComponent(from)
}

func (g *GrafanaDashboardComponent) AddTrait(traits ...apis.Trait) *GrafanaDashboardComponent {
	g.Base.Traits = append(g.Base.Traits, traits...)
	return g
}

func (g *GrafanaDashboardComponent) GetTrait(_type string) apis.Trait {
	for _, _t := range g.Base.Traits {
		if _t.DefType() == _type {
			return _t
		}
	}
	return nil
}

func (g *GrafanaDashboardComponent) ComponentName() string {
	return g.Base.Name
}

func (g *GrafanaDashboardComponent) DefType() string {
	return GrafanaDashboardType
}

func (g *GrafanaDashboardComponent) DependsOn(dependsOn []string) *GrafanaDashboardComponent {
	g.Base.DependsOn = dependsOn
	return g
}

func (g *GrafanaDashboardComponent) Inputs(input common.StepInputs) *GrafanaDashboardComponent {
	g.Base.Inputs = input
	return g
}

func (g *GrafanaDashboardComponent) Outputs(output common.StepOutputs) *GrafanaDashboardComponent {
	g.Base.Outputs = output
	return g
}

func (g *GrafanaDashboardComponent) AddDependsOn(dependsOn string) *GrafanaDashboardComponent {
	g.Base.DependsOn = append(g.Base.DependsOn, dependsOn)
	return g
}
