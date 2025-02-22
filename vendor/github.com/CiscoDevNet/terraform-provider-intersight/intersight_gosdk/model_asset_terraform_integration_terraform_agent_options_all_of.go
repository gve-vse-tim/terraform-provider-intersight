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

// AssetTerraformIntegrationTerraformAgentOptionsAllOf Definition of the list of properties defined in 'asset.TerraformIntegrationTerraformAgentOptions', excluding properties defined in parent classes.
type AssetTerraformIntegrationTerraformAgentOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string   `json:"ObjectType"`
	ManagedHosts []string `json:"ManagedHosts,omitempty"`
	// Agent pool name for Terraform Agent platform type.
	TerraformAgentPoolName *string `json:"TerraformAgentPoolName,omitempty"`
	// Organization for Terraform Agent platform type.
	TerraformOrganization *string `json:"TerraformOrganization,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _AssetTerraformIntegrationTerraformAgentOptionsAllOf AssetTerraformIntegrationTerraformAgentOptionsAllOf

// NewAssetTerraformIntegrationTerraformAgentOptionsAllOf instantiates a new AssetTerraformIntegrationTerraformAgentOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTerraformIntegrationTerraformAgentOptionsAllOf(classId string, objectType string) *AssetTerraformIntegrationTerraformAgentOptionsAllOf {
	this := AssetTerraformIntegrationTerraformAgentOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetTerraformIntegrationTerraformAgentOptionsAllOfWithDefaults instantiates a new AssetTerraformIntegrationTerraformAgentOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTerraformIntegrationTerraformAgentOptionsAllOfWithDefaults() *AssetTerraformIntegrationTerraformAgentOptionsAllOf {
	this := AssetTerraformIntegrationTerraformAgentOptionsAllOf{}
	var classId string = "asset.TerraformIntegrationTerraformAgentOptions"
	this.ClassId = classId
	var objectType string = "asset.TerraformIntegrationTerraformAgentOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetManagedHosts returns the ManagedHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetManagedHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ManagedHosts
}

// GetManagedHostsOk returns a tuple with the ManagedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetManagedHostsOk() ([]string, bool) {
	if o == nil || o.ManagedHosts == nil {
		return nil, false
	}
	return o.ManagedHosts, true
}

// HasManagedHosts returns a boolean if a field has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) HasManagedHosts() bool {
	if o != nil && o.ManagedHosts != nil {
		return true
	}

	return false
}

// SetManagedHosts gets a reference to the given []string and assigns it to the ManagedHosts field.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetManagedHosts(v []string) {
	o.ManagedHosts = v
}

// GetTerraformAgentPoolName returns the TerraformAgentPoolName field value if set, zero value otherwise.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformAgentPoolName() string {
	if o == nil || o.TerraformAgentPoolName == nil {
		var ret string
		return ret
	}
	return *o.TerraformAgentPoolName
}

// GetTerraformAgentPoolNameOk returns a tuple with the TerraformAgentPoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformAgentPoolNameOk() (*string, bool) {
	if o == nil || o.TerraformAgentPoolName == nil {
		return nil, false
	}
	return o.TerraformAgentPoolName, true
}

// HasTerraformAgentPoolName returns a boolean if a field has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) HasTerraformAgentPoolName() bool {
	if o != nil && o.TerraformAgentPoolName != nil {
		return true
	}

	return false
}

// SetTerraformAgentPoolName gets a reference to the given string and assigns it to the TerraformAgentPoolName field.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetTerraformAgentPoolName(v string) {
	o.TerraformAgentPoolName = &v
}

// GetTerraformOrganization returns the TerraformOrganization field value if set, zero value otherwise.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformOrganization() string {
	if o == nil || o.TerraformOrganization == nil {
		var ret string
		return ret
	}
	return *o.TerraformOrganization
}

// GetTerraformOrganizationOk returns a tuple with the TerraformOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformOrganizationOk() (*string, bool) {
	if o == nil || o.TerraformOrganization == nil {
		return nil, false
	}
	return o.TerraformOrganization, true
}

// HasTerraformOrganization returns a boolean if a field has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) HasTerraformOrganization() bool {
	if o != nil && o.TerraformOrganization != nil {
		return true
	}

	return false
}

// SetTerraformOrganization gets a reference to the given string and assigns it to the TerraformOrganization field.
func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetTerraformOrganization(v string) {
	o.TerraformOrganization = &v
}

func (o AssetTerraformIntegrationTerraformAgentOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ManagedHosts != nil {
		toSerialize["ManagedHosts"] = o.ManagedHosts
	}
	if o.TerraformAgentPoolName != nil {
		toSerialize["TerraformAgentPoolName"] = o.TerraformAgentPoolName
	}
	if o.TerraformOrganization != nil {
		toSerialize["TerraformOrganization"] = o.TerraformOrganization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTerraformIntegrationTerraformAgentOptionsAllOf := _AssetTerraformIntegrationTerraformAgentOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetTerraformIntegrationTerraformAgentOptionsAllOf); err == nil {
		*o = AssetTerraformIntegrationTerraformAgentOptionsAllOf(varAssetTerraformIntegrationTerraformAgentOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ManagedHosts")
		delete(additionalProperties, "TerraformAgentPoolName")
		delete(additionalProperties, "TerraformOrganization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf struct {
	value *AssetTerraformIntegrationTerraformAgentOptionsAllOf
	isSet bool
}

func (v NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf) Get() *AssetTerraformIntegrationTerraformAgentOptionsAllOf {
	return v.value
}

func (v *NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf) Set(val *AssetTerraformIntegrationTerraformAgentOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTerraformIntegrationTerraformAgentOptionsAllOf(val *AssetTerraformIntegrationTerraformAgentOptionsAllOf) *NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf {
	return &NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTerraformIntegrationTerraformAgentOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
