/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// NetworkconfigPolicyAllOf Definition of the list of properties defined in 'networkconfig.Policy', excluding properties defined in parent classes.
type NetworkconfigPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of the secondary DNS server.
	AlternateIpv4dnsServer *string `json:"AlternateIpv4dnsServer,omitempty"`
	// IP address of the secondary DNS server.
	AlternateIpv6dnsServer *string `json:"AlternateIpv6dnsServer,omitempty"`
	// The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.
	DynamicDnsDomain *string `json:"DynamicDnsDomain,omitempty"`
	// If enabled, updates the resource records to the DNS from Cisco IMC.
	EnableDynamicDns *bool `json:"EnableDynamicDns,omitempty"`
	// If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.
	EnableIpv4dnsFromDhcp *bool `json:"EnableIpv4dnsFromDhcp,omitempty"`
	// If enabled, allows to configure IPv6 properties.
	EnableIpv6 *bool `json:"EnableIpv6,omitempty"`
	// If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.
	EnableIpv6dnsFromDhcp *bool `json:"EnableIpv6dnsFromDhcp,omitempty"`
	// IP address of the primary DNS server.
	PreferredIpv4dnsServer *string `json:"PreferredIpv4dnsServer,omitempty"`
	// IP address of the primary DNS server.
	PreferredIpv6dnsServer *string                               `json:"PreferredIpv6dnsServer,omitempty"`
	ApplianceAccount       *IamAccountRelationship               `json:"ApplianceAccount,omitempty"`
	Organization           *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkconfigPolicyAllOf NetworkconfigPolicyAllOf

// NewNetworkconfigPolicyAllOf instantiates a new NetworkconfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkconfigPolicyAllOf(classId string, objectType string) *NetworkconfigPolicyAllOf {
	this := NetworkconfigPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkconfigPolicyAllOfWithDefaults instantiates a new NetworkconfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkconfigPolicyAllOfWithDefaults() *NetworkconfigPolicyAllOf {
	this := NetworkconfigPolicyAllOf{}
	var classId string = "networkconfig.Policy"
	this.ClassId = classId
	var objectType string = "networkconfig.Policy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkconfigPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkconfigPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkconfigPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkconfigPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlternateIpv4dnsServer returns the AlternateIpv4dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetAlternateIpv4dnsServer() string {
	if o == nil || o.AlternateIpv4dnsServer == nil {
		var ret string
		return ret
	}
	return *o.AlternateIpv4dnsServer
}

// GetAlternateIpv4dnsServerOk returns a tuple with the AlternateIpv4dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetAlternateIpv4dnsServerOk() (*string, bool) {
	if o == nil || o.AlternateIpv4dnsServer == nil {
		return nil, false
	}
	return o.AlternateIpv4dnsServer, true
}

// HasAlternateIpv4dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasAlternateIpv4dnsServer() bool {
	if o != nil && o.AlternateIpv4dnsServer != nil {
		return true
	}

	return false
}

// SetAlternateIpv4dnsServer gets a reference to the given string and assigns it to the AlternateIpv4dnsServer field.
func (o *NetworkconfigPolicyAllOf) SetAlternateIpv4dnsServer(v string) {
	o.AlternateIpv4dnsServer = &v
}

// GetAlternateIpv6dnsServer returns the AlternateIpv6dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetAlternateIpv6dnsServer() string {
	if o == nil || o.AlternateIpv6dnsServer == nil {
		var ret string
		return ret
	}
	return *o.AlternateIpv6dnsServer
}

// GetAlternateIpv6dnsServerOk returns a tuple with the AlternateIpv6dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetAlternateIpv6dnsServerOk() (*string, bool) {
	if o == nil || o.AlternateIpv6dnsServer == nil {
		return nil, false
	}
	return o.AlternateIpv6dnsServer, true
}

// HasAlternateIpv6dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasAlternateIpv6dnsServer() bool {
	if o != nil && o.AlternateIpv6dnsServer != nil {
		return true
	}

	return false
}

// SetAlternateIpv6dnsServer gets a reference to the given string and assigns it to the AlternateIpv6dnsServer field.
func (o *NetworkconfigPolicyAllOf) SetAlternateIpv6dnsServer(v string) {
	o.AlternateIpv6dnsServer = &v
}

// GetDynamicDnsDomain returns the DynamicDnsDomain field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetDynamicDnsDomain() string {
	if o == nil || o.DynamicDnsDomain == nil {
		var ret string
		return ret
	}
	return *o.DynamicDnsDomain
}

// GetDynamicDnsDomainOk returns a tuple with the DynamicDnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetDynamicDnsDomainOk() (*string, bool) {
	if o == nil || o.DynamicDnsDomain == nil {
		return nil, false
	}
	return o.DynamicDnsDomain, true
}

// HasDynamicDnsDomain returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasDynamicDnsDomain() bool {
	if o != nil && o.DynamicDnsDomain != nil {
		return true
	}

	return false
}

// SetDynamicDnsDomain gets a reference to the given string and assigns it to the DynamicDnsDomain field.
func (o *NetworkconfigPolicyAllOf) SetDynamicDnsDomain(v string) {
	o.DynamicDnsDomain = &v
}

// GetEnableDynamicDns returns the EnableDynamicDns field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetEnableDynamicDns() bool {
	if o == nil || o.EnableDynamicDns == nil {
		var ret bool
		return ret
	}
	return *o.EnableDynamicDns
}

// GetEnableDynamicDnsOk returns a tuple with the EnableDynamicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetEnableDynamicDnsOk() (*bool, bool) {
	if o == nil || o.EnableDynamicDns == nil {
		return nil, false
	}
	return o.EnableDynamicDns, true
}

// HasEnableDynamicDns returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasEnableDynamicDns() bool {
	if o != nil && o.EnableDynamicDns != nil {
		return true
	}

	return false
}

// SetEnableDynamicDns gets a reference to the given bool and assigns it to the EnableDynamicDns field.
func (o *NetworkconfigPolicyAllOf) SetEnableDynamicDns(v bool) {
	o.EnableDynamicDns = &v
}

// GetEnableIpv4dnsFromDhcp returns the EnableIpv4dnsFromDhcp field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetEnableIpv4dnsFromDhcp() bool {
	if o == nil || o.EnableIpv4dnsFromDhcp == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpv4dnsFromDhcp
}

// GetEnableIpv4dnsFromDhcpOk returns a tuple with the EnableIpv4dnsFromDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetEnableIpv4dnsFromDhcpOk() (*bool, bool) {
	if o == nil || o.EnableIpv4dnsFromDhcp == nil {
		return nil, false
	}
	return o.EnableIpv4dnsFromDhcp, true
}

// HasEnableIpv4dnsFromDhcp returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasEnableIpv4dnsFromDhcp() bool {
	if o != nil && o.EnableIpv4dnsFromDhcp != nil {
		return true
	}

	return false
}

// SetEnableIpv4dnsFromDhcp gets a reference to the given bool and assigns it to the EnableIpv4dnsFromDhcp field.
func (o *NetworkconfigPolicyAllOf) SetEnableIpv4dnsFromDhcp(v bool) {
	o.EnableIpv4dnsFromDhcp = &v
}

// GetEnableIpv6 returns the EnableIpv6 field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetEnableIpv6() bool {
	if o == nil || o.EnableIpv6 == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpv6
}

// GetEnableIpv6Ok returns a tuple with the EnableIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetEnableIpv6Ok() (*bool, bool) {
	if o == nil || o.EnableIpv6 == nil {
		return nil, false
	}
	return o.EnableIpv6, true
}

// HasEnableIpv6 returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasEnableIpv6() bool {
	if o != nil && o.EnableIpv6 != nil {
		return true
	}

	return false
}

// SetEnableIpv6 gets a reference to the given bool and assigns it to the EnableIpv6 field.
func (o *NetworkconfigPolicyAllOf) SetEnableIpv6(v bool) {
	o.EnableIpv6 = &v
}

// GetEnableIpv6dnsFromDhcp returns the EnableIpv6dnsFromDhcp field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetEnableIpv6dnsFromDhcp() bool {
	if o == nil || o.EnableIpv6dnsFromDhcp == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpv6dnsFromDhcp
}

// GetEnableIpv6dnsFromDhcpOk returns a tuple with the EnableIpv6dnsFromDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetEnableIpv6dnsFromDhcpOk() (*bool, bool) {
	if o == nil || o.EnableIpv6dnsFromDhcp == nil {
		return nil, false
	}
	return o.EnableIpv6dnsFromDhcp, true
}

// HasEnableIpv6dnsFromDhcp returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasEnableIpv6dnsFromDhcp() bool {
	if o != nil && o.EnableIpv6dnsFromDhcp != nil {
		return true
	}

	return false
}

// SetEnableIpv6dnsFromDhcp gets a reference to the given bool and assigns it to the EnableIpv6dnsFromDhcp field.
func (o *NetworkconfigPolicyAllOf) SetEnableIpv6dnsFromDhcp(v bool) {
	o.EnableIpv6dnsFromDhcp = &v
}

// GetPreferredIpv4dnsServer returns the PreferredIpv4dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetPreferredIpv4dnsServer() string {
	if o == nil || o.PreferredIpv4dnsServer == nil {
		var ret string
		return ret
	}
	return *o.PreferredIpv4dnsServer
}

// GetPreferredIpv4dnsServerOk returns a tuple with the PreferredIpv4dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetPreferredIpv4dnsServerOk() (*string, bool) {
	if o == nil || o.PreferredIpv4dnsServer == nil {
		return nil, false
	}
	return o.PreferredIpv4dnsServer, true
}

// HasPreferredIpv4dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasPreferredIpv4dnsServer() bool {
	if o != nil && o.PreferredIpv4dnsServer != nil {
		return true
	}

	return false
}

// SetPreferredIpv4dnsServer gets a reference to the given string and assigns it to the PreferredIpv4dnsServer field.
func (o *NetworkconfigPolicyAllOf) SetPreferredIpv4dnsServer(v string) {
	o.PreferredIpv4dnsServer = &v
}

// GetPreferredIpv6dnsServer returns the PreferredIpv6dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetPreferredIpv6dnsServer() string {
	if o == nil || o.PreferredIpv6dnsServer == nil {
		var ret string
		return ret
	}
	return *o.PreferredIpv6dnsServer
}

// GetPreferredIpv6dnsServerOk returns a tuple with the PreferredIpv6dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetPreferredIpv6dnsServerOk() (*string, bool) {
	if o == nil || o.PreferredIpv6dnsServer == nil {
		return nil, false
	}
	return o.PreferredIpv6dnsServer, true
}

// HasPreferredIpv6dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasPreferredIpv6dnsServer() bool {
	if o != nil && o.PreferredIpv6dnsServer != nil {
		return true
	}

	return false
}

// SetPreferredIpv6dnsServer gets a reference to the given string and assigns it to the PreferredIpv6dnsServer field.
func (o *NetworkconfigPolicyAllOf) SetPreferredIpv6dnsServer(v string) {
	o.PreferredIpv6dnsServer = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetApplianceAccount() IamAccountRelationship {
	if o == nil || o.ApplianceAccount == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.ApplianceAccount == nil {
		return nil, false
	}
	return o.ApplianceAccount, true
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount != nil {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given IamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *NetworkconfigPolicyAllOf) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *NetworkconfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *NetworkconfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkconfigPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkconfigPolicyAllOf) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *NetworkconfigPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *NetworkconfigPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o NetworkconfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlternateIpv4dnsServer != nil {
		toSerialize["AlternateIpv4dnsServer"] = o.AlternateIpv4dnsServer
	}
	if o.AlternateIpv6dnsServer != nil {
		toSerialize["AlternateIpv6dnsServer"] = o.AlternateIpv6dnsServer
	}
	if o.DynamicDnsDomain != nil {
		toSerialize["DynamicDnsDomain"] = o.DynamicDnsDomain
	}
	if o.EnableDynamicDns != nil {
		toSerialize["EnableDynamicDns"] = o.EnableDynamicDns
	}
	if o.EnableIpv4dnsFromDhcp != nil {
		toSerialize["EnableIpv4dnsFromDhcp"] = o.EnableIpv4dnsFromDhcp
	}
	if o.EnableIpv6 != nil {
		toSerialize["EnableIpv6"] = o.EnableIpv6
	}
	if o.EnableIpv6dnsFromDhcp != nil {
		toSerialize["EnableIpv6dnsFromDhcp"] = o.EnableIpv6dnsFromDhcp
	}
	if o.PreferredIpv4dnsServer != nil {
		toSerialize["PreferredIpv4dnsServer"] = o.PreferredIpv4dnsServer
	}
	if o.PreferredIpv6dnsServer != nil {
		toSerialize["PreferredIpv6dnsServer"] = o.PreferredIpv6dnsServer
	}
	if o.ApplianceAccount != nil {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkconfigPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkconfigPolicyAllOf := _NetworkconfigPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varNetworkconfigPolicyAllOf); err == nil {
		*o = NetworkconfigPolicyAllOf(varNetworkconfigPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlternateIpv4dnsServer")
		delete(additionalProperties, "AlternateIpv6dnsServer")
		delete(additionalProperties, "DynamicDnsDomain")
		delete(additionalProperties, "EnableDynamicDns")
		delete(additionalProperties, "EnableIpv4dnsFromDhcp")
		delete(additionalProperties, "EnableIpv6")
		delete(additionalProperties, "EnableIpv6dnsFromDhcp")
		delete(additionalProperties, "PreferredIpv4dnsServer")
		delete(additionalProperties, "PreferredIpv6dnsServer")
		delete(additionalProperties, "ApplianceAccount")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkconfigPolicyAllOf struct {
	value *NetworkconfigPolicyAllOf
	isSet bool
}

func (v NullableNetworkconfigPolicyAllOf) Get() *NetworkconfigPolicyAllOf {
	return v.value
}

func (v *NullableNetworkconfigPolicyAllOf) Set(val *NetworkconfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkconfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkconfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkconfigPolicyAllOf(val *NetworkconfigPolicyAllOf) *NullableNetworkconfigPolicyAllOf {
	return &NullableNetworkconfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableNetworkconfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkconfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
