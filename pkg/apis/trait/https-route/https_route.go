/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package https_route

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/vela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
)

// checks if the HttpsRouteSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpsRouteSpec{}

// HttpsRouteSpec struct for HttpsRouteSpec
type HttpsRouteSpec struct {
	TLSPort *int32 `json:"TLSPort,omitempty"`
	// Specify some domains, the domain may be prefixed with a wildcard label (*.)
	Domains []string `json:"domains,omitempty"`
	// Specify some HTTP matchers, filters and actions.
	Rules []Rules `json:"rules,omitempty"`
	// Specify the TLS secrets
	Secrets []Secrets `json:"secrets,omitempty"`
}

// NewHttpsRouteSpecWith instantiates a new HttpsRouteSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpsRouteSpecWith() *HttpsRouteSpec {
	this := HttpsRouteSpec{}
	var tLSPort int32 = 8443
	this.TLSPort = &tLSPort
	return &this
}

// NewHttpsRouteSpec instantiates a new HttpsRouteSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpsRouteSpec() *HttpsRouteSpec {
	this := HttpsRouteSpec{}
	var tLSPort int32 = 8443
	this.TLSPort = &tLSPort
	return &this
}

// NewHttpsRouteSpecs converts a list HttpsRouteSpec pointers to objects.
// This is helpful when the SetHttpsRouteSpec requires a list of objects
func NewHttpsRouteSpecs(ps ...*HttpsRouteSpec) []HttpsRouteSpec {
	objs := []HttpsRouteSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// GetTLSPort returns the TLSPort field value if set, zero value otherwise.
func (o *HTTPSRouteTrait) GetTLSPort() int32 {
	if o == nil || utils.IsNil(o.Properties.TLSPort) {
		var ret int32
		return ret
	}
	return *o.Properties.TLSPort
}

// GetTLSPortOk returns a tuple with the TLSPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPSRouteTrait) GetTLSPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.TLSPort) {
		return nil, false
	}
	return o.Properties.TLSPort, true
}

// HasTLSPort returns a boolean if a field has been set.
func (o *HTTPSRouteTrait) HasTLSPort() bool {
	if o != nil && !utils.IsNil(o.Properties.TLSPort) {
		return true
	}

	return false
}

// SetTLSPort gets a reference to the given int32 and assigns it to the tLSPort field.
// TLSPort:
func (o *HTTPSRouteTrait) SetTLSPort(v int32) *HTTPSRouteTrait {
	o.Properties.TLSPort = &v
	return o
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *HTTPSRouteTrait) GetDomains() []string {
	if o == nil || utils.IsNil(o.Properties.Domains) {
		var ret []string
		return ret
	}
	return o.Properties.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPSRouteTrait) GetDomainsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Domains) {
		return nil, false
	}
	return o.Properties.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *HTTPSRouteTrait) HasDomains() bool {
	if o != nil && !utils.IsNil(o.Properties.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the domains field.
// Domains:  Specify some domains, the domain may be prefixed with a wildcard label (*.)
func (o *HTTPSRouteTrait) SetDomains(v []string) *HTTPSRouteTrait {
	o.Properties.Domains = v
	return o
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *HTTPSRouteTrait) GetRules() []Rules {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		var ret []Rules
		return ret
	}
	return o.Properties.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPSRouteTrait) GetRulesOk() ([]Rules, bool) {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		return nil, false
	}
	return o.Properties.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *HTTPSRouteTrait) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []Rules and assigns it to the rules field.
// Rules:  Specify some HTTP matchers, filters and actions.
func (o *HTTPSRouteTrait) SetRules(v []Rules) *HTTPSRouteTrait {
	o.Properties.Rules = v
	return o
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *HTTPSRouteTrait) GetSecrets() []Secrets {
	if o == nil || utils.IsNil(o.Properties.Secrets) {
		var ret []Secrets
		return ret
	}
	return o.Properties.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPSRouteTrait) GetSecretsOk() ([]Secrets, bool) {
	if o == nil || utils.IsNil(o.Properties.Secrets) {
		return nil, false
	}
	return o.Properties.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *HTTPSRouteTrait) HasSecrets() bool {
	if o != nil && !utils.IsNil(o.Properties.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []Secrets and assigns it to the secrets field.
// Secrets:  Specify the TLS secrets
func (o *HTTPSRouteTrait) SetSecrets(v []Secrets) *HTTPSRouteTrait {
	o.Properties.Secrets = v
	return o
}

func (o HttpsRouteSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpsRouteSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TLSPort) {
		toSerialize["TLSPort"] = o.TLSPort
	}
	if !utils.IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !utils.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !utils.IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}

type NullableHttpsRouteSpec struct {
	value *HttpsRouteSpec
	isSet bool
}

func (v NullableHttpsRouteSpec) Get() *HttpsRouteSpec {
	return v.value
}

func (v *NullableHttpsRouteSpec) Set(val *HttpsRouteSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpsRouteSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpsRouteSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpsRouteSpec(val *HttpsRouteSpec) *NullableHttpsRouteSpec {
	return &NullableHttpsRouteSpec{value: val, isSet: true}
}

func (v NullableHttpsRouteSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpsRouteSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const HttpsRouteType = "https-route"

func init() {
	sdkcommon.RegisterTrait(HttpsRouteType, FromTrait)
}

type HTTPSRouteTrait struct {
	Base       apis.TraitBase
	Properties HttpsRouteSpec
}

func HttpsRoute() *HTTPSRouteTrait {
	h := &HTTPSRouteTrait{Base: apis.TraitBase{}}
	return h
}

func (h *HTTPSRouteTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(h.Properties),
		Type:       HttpsRouteType,
	}
	return res
}

func (h *HTTPSRouteTrait) FromTrait(from common.ApplicationTrait) (*HTTPSRouteTrait, error) {
	var properties HttpsRouteSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	h.Base.Type = HttpsRouteType
	h.Properties = properties
	return h, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	h := &HTTPSRouteTrait{}
	return h.FromTrait(from)
}

func (h *HTTPSRouteTrait) DefType() string {
	return HttpsRouteType
}
