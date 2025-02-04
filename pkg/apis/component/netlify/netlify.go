/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlify

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the NetlifySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetlifySpec{}

// NetlifySpec struct for NetlifySpec
type NetlifySpec struct {
	Backofflimit *int32 `json:"backofflimit,omitempty"`
	Completions  *int32 `json:"completions,omitempty"`
	// Specify the repo address your resources of the frontend service  store
	Gitrepo *string `json:"gitrepo,omitempty"`
	// Specify the direcotry name for saving the resources from git or other repo supporting git-cli
	Reponame      *string `json:"reponame,omitempty"`
	RestartPolicy *string `json:"restartPolicy,omitempty"`
	// Specify the website name on app.netlify.com
	Sitname *string `json:"sitname,omitempty"`
	// Specify the token got from app.netlify.com
	Token *string `json:"token,omitempty"`
}

// NewNetlifySpecWith instantiates a new NetlifySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetlifySpecWith() *NetlifySpec {
	this := NetlifySpec{}
	var backofflimit int32 = 5
	this.Backofflimit = &backofflimit
	var completions int32 = 1
	this.Completions = &completions
	var gitrepo string = "https://github.com/oeular/oeular.github.io.Properties.git"
	this.Gitrepo = &gitrepo
	var reponame string = "oeular.github.io"
	this.Reponame = &reponame
	var restartPolicy string = "OnFailure"
	this.RestartPolicy = &restartPolicy
	var sitname string = "deploy-from-vela"
	this.Sitname = &sitname
	var token string = "your-own-token-after-base64"
	this.Token = &token
	return &this
}

// NewNetlifySpec instantiates a new NetlifySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetlifySpec() *NetlifySpec {
	this := NetlifySpec{}
	var backofflimit int32 = 5
	this.Backofflimit = &backofflimit
	var completions int32 = 1
	this.Completions = &completions
	var gitrepo string = "https://github.com/oeular/oeular.github.io.Properties.git"
	this.Gitrepo = &gitrepo
	var reponame string = "oeular.github.io"
	this.Reponame = &reponame
	var restartPolicy string = "OnFailure"
	this.RestartPolicy = &restartPolicy
	var sitname string = "deploy-from-vela"
	this.Sitname = &sitname
	var token string = "your-own-token-after-base64"
	this.Token = &token
	return &this
}

// NewNetlifySpecs converts a list NetlifySpec pointers to objects.
// This is helpful when the SetNetlifySpec requires a list of objects
func NewNetlifySpecs(ps ...*NetlifySpec) []NetlifySpec {
	objs := []NetlifySpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetBackofflimit returns the Backofflimit field value if set, zero value otherwise.
func (o *NetlifyComponent) GetBackofflimit() int32 {
	if o == nil || utils.IsNil(o.Properties.Backofflimit) {
		var ret int32
		return ret
	}
	return *o.Properties.Backofflimit
}

// GetBackofflimitOk returns a tuple with the Backofflimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetBackofflimitOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Backofflimit) {
		return nil, false
	}
	return o.Properties.Backofflimit, true
}

// HasBackofflimit returns a boolean if a field has been set.
func (o *NetlifyComponent) HasBackofflimit() bool {
	if o != nil && !utils.IsNil(o.Properties.Backofflimit) {
		return true
	}

	return false
}

// SetBackofflimit gets a reference to the given int32 and assigns it to the backofflimit field.
// Backofflimit:
func (o *NetlifyComponent) SetBackofflimit(v int32) *NetlifyComponent {
	o.Properties.Backofflimit = &v
	return o
}

// GetCompletions returns the Completions field value if set, zero value otherwise.
func (o *NetlifyComponent) GetCompletions() int32 {
	if o == nil || utils.IsNil(o.Properties.Completions) {
		var ret int32
		return ret
	}
	return *o.Properties.Completions
}

// GetCompletionsOk returns a tuple with the Completions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetCompletionsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.Completions) {
		return nil, false
	}
	return o.Properties.Completions, true
}

// HasCompletions returns a boolean if a field has been set.
func (o *NetlifyComponent) HasCompletions() bool {
	if o != nil && !utils.IsNil(o.Properties.Completions) {
		return true
	}

	return false
}

// SetCompletions gets a reference to the given int32 and assigns it to the completions field.
// Completions:
func (o *NetlifyComponent) SetCompletions(v int32) *NetlifyComponent {
	o.Properties.Completions = &v
	return o
}

// GetGitrepo returns the Gitrepo field value if set, zero value otherwise.
func (o *NetlifyComponent) GetGitrepo() string {
	if o == nil || utils.IsNil(o.Properties.Gitrepo) {
		var ret string
		return ret
	}
	return *o.Properties.Gitrepo
}

// GetGitrepoOk returns a tuple with the Gitrepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetGitrepoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Gitrepo) {
		return nil, false
	}
	return o.Properties.Gitrepo, true
}

// HasGitrepo returns a boolean if a field has been set.
func (o *NetlifyComponent) HasGitrepo() bool {
	if o != nil && !utils.IsNil(o.Properties.Gitrepo) {
		return true
	}

	return false
}

// SetGitrepo gets a reference to the given string and assigns it to the gitrepo field.
// Gitrepo:  Specify the repo address your resources of the frontend service  store
func (o *NetlifyComponent) SetGitrepo(v string) *NetlifyComponent {
	o.Properties.Gitrepo = &v
	return o
}

// GetReponame returns the Reponame field value if set, zero value otherwise.
func (o *NetlifyComponent) GetReponame() string {
	if o == nil || utils.IsNil(o.Properties.Reponame) {
		var ret string
		return ret
	}
	return *o.Properties.Reponame
}

// GetReponameOk returns a tuple with the Reponame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetReponameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Reponame) {
		return nil, false
	}
	return o.Properties.Reponame, true
}

// HasReponame returns a boolean if a field has been set.
func (o *NetlifyComponent) HasReponame() bool {
	if o != nil && !utils.IsNil(o.Properties.Reponame) {
		return true
	}

	return false
}

// SetReponame gets a reference to the given string and assigns it to the reponame field.
// Reponame:  Specify the direcotry name for saving the resources from git or other repo supporting git-cli
func (o *NetlifyComponent) SetReponame(v string) *NetlifyComponent {
	o.Properties.Reponame = &v
	return o
}

// GetRestartPolicy returns the RestartPolicy field value if set, zero value otherwise.
func (o *NetlifyComponent) GetRestartPolicy() string {
	if o == nil || utils.IsNil(o.Properties.RestartPolicy) {
		var ret string
		return ret
	}
	return *o.Properties.RestartPolicy
}

// GetRestartPolicyOk returns a tuple with the RestartPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetRestartPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.RestartPolicy) {
		return nil, false
	}
	return o.Properties.RestartPolicy, true
}

// HasRestartPolicy returns a boolean if a field has been set.
func (o *NetlifyComponent) HasRestartPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.RestartPolicy) {
		return true
	}

	return false
}

// SetRestartPolicy gets a reference to the given string and assigns it to the restartPolicy field.
// RestartPolicy:
func (o *NetlifyComponent) SetRestartPolicy(v string) *NetlifyComponent {
	o.Properties.RestartPolicy = &v
	return o
}

// GetSitname returns the Sitname field value if set, zero value otherwise.
func (o *NetlifyComponent) GetSitname() string {
	if o == nil || utils.IsNil(o.Properties.Sitname) {
		var ret string
		return ret
	}
	return *o.Properties.Sitname
}

// GetSitnameOk returns a tuple with the Sitname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetSitnameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Sitname) {
		return nil, false
	}
	return o.Properties.Sitname, true
}

// HasSitname returns a boolean if a field has been set.
func (o *NetlifyComponent) HasSitname() bool {
	if o != nil && !utils.IsNil(o.Properties.Sitname) {
		return true
	}

	return false
}

// SetSitname gets a reference to the given string and assigns it to the sitname field.
// Sitname:  Specify the website name on app.netlify.com
func (o *NetlifyComponent) SetSitname(v string) *NetlifyComponent {
	o.Properties.Sitname = &v
	return o
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *NetlifyComponent) GetToken() string {
	if o == nil || utils.IsNil(o.Properties.Token) {
		var ret string
		return ret
	}
	return *o.Properties.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetlifyComponent) GetTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Token) {
		return nil, false
	}
	return o.Properties.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *NetlifyComponent) HasToken() bool {
	if o != nil && !utils.IsNil(o.Properties.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the token field.
// Token:  Specify the token got from app.netlify.com
func (o *NetlifyComponent) SetToken(v string) *NetlifyComponent {
	o.Properties.Token = &v
	return o
}

func (o NetlifySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetlifySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Backofflimit) {
		toSerialize["backofflimit"] = o.Backofflimit
	}
	if !utils.IsNil(o.Completions) {
		toSerialize["completions"] = o.Completions
	}
	if !utils.IsNil(o.Gitrepo) {
		toSerialize["gitrepo"] = o.Gitrepo
	}
	if !utils.IsNil(o.Reponame) {
		toSerialize["reponame"] = o.Reponame
	}
	if !utils.IsNil(o.RestartPolicy) {
		toSerialize["restartPolicy"] = o.RestartPolicy
	}
	if !utils.IsNil(o.Sitname) {
		toSerialize["sitname"] = o.Sitname
	}
	if !utils.IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableNetlifySpec struct {
	value *NetlifySpec
	isSet bool
}

func (v NullableNetlifySpec) Get() *NetlifySpec {
	return v.value
}

func (v *NullableNetlifySpec) Set(val *NetlifySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetlifySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetlifySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetlifySpec(val *NetlifySpec) *NullableNetlifySpec {
	return &NullableNetlifySpec{value: val, isSet: true}
}

func (v NullableNetlifySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetlifySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const NetlifyType = "netlify"

func init() {
	sdkcommon.RegisterComponent(NetlifyType, FromComponent)
}

type NetlifyComponent struct {
	Base       apis.ComponentBase
	Properties NetlifySpec
}

func Netlify(name string) *NetlifyComponent {
	n := &NetlifyComponent{Base: apis.ComponentBase{
		Name: name,
		Type: NetlifyType,
	}}
	return n
}

func (n *NetlifyComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range n.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  n.Base.DependsOn,
		Inputs:     n.Base.Inputs,
		Name:       n.Base.Name,
		Outputs:    n.Base.Outputs,
		Properties: util.Object2RawExtension(n.Properties),
		Traits:     traits,
		Type:       NetlifyType,
	}
	return res
}

func (n *NetlifyComponent) FromComponent(from common.ApplicationComponent) (*NetlifyComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		n.Base.Traits = append(n.Base.Traits, _t)
	}
	var properties NetlifySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	n.Base.Name = from.Name
	n.Base.DependsOn = from.DependsOn
	n.Base.Inputs = from.Inputs
	n.Base.Outputs = from.Outputs
	n.Base.Type = NetlifyType
	n.Properties = properties
	return n, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	n := &NetlifyComponent{}
	return n.FromComponent(from)
}

func (n *NetlifyComponent) AddTrait(traits ...apis.Trait) *NetlifyComponent {
	n.Base.Traits = append(n.Base.Traits, traits...)
	return n
}

func (n *NetlifyComponent) GetTrait(_type string) apis.Trait {
	for _, _t := range n.Base.Traits {
		if _t.DefType() == _type {
			return _t
		}
	}
	return nil
}

func (n *NetlifyComponent) ComponentName() string {
	return n.Base.Name
}

func (n *NetlifyComponent) DefType() string {
	return NetlifyType
}

func (n *NetlifyComponent) DependsOn(dependsOn []string) *NetlifyComponent {
	n.Base.DependsOn = dependsOn
	return n
}

func (n *NetlifyComponent) Inputs(input common.StepInputs) *NetlifyComponent {
	n.Base.Inputs = input
	return n
}

func (n *NetlifyComponent) Outputs(output common.StepOutputs) *NetlifyComponent {
	n.Base.Outputs = output
	return n
}

func (n *NetlifyComponent) AddDependsOn(dependsOn string) *NetlifyComponent {
	n.Base.DependsOn = append(n.Base.DependsOn, dependsOn)
	return n
}
