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

// HyperflexSoftwareVersionPolicyAllOf Definition of the list of properties defined in 'hyperflex.SoftwareVersionPolicy', excluding properties defined in parent classes.
type HyperflexSoftwareVersionPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster.
	HxdpVersion *string `json:"HxdpVersion,omitempty"`
	// Desired  hypervisor version to apply for all the nodes on the HyperFlex cluster.
	HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
	// Desired server firmware version to apply on the HyperFlex Cluster.
	ServerFirmwareVersion  *string                              `json:"ServerFirmwareVersion,omitempty"`
	ServerFirmwareVersions []HyperflexServerFirmwareVersionInfo `json:"ServerFirmwareVersions,omitempty"`
	UpgradeTypes           []string                             `json:"UpgradeTypes,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles           []HyperflexClusterProfileRelationship       `json:"ClusterProfiles,omitempty"`
	HxdpVersionInfo           *SoftwareHyperflexDistributableRelationship `json:"HxdpVersionInfo,omitempty"`
	HypervisorVersionInfo     *SoftwareHyperflexDistributableRelationship `json:"HypervisorVersionInfo,omitempty"`
	Organization              *OrganizationOrganizationRelationship       `json:"Organization,omitempty"`
	ServerFirmwareVersionInfo *FirmwareDistributableRelationship          `json:"ServerFirmwareVersionInfo,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _HyperflexSoftwareVersionPolicyAllOf HyperflexSoftwareVersionPolicyAllOf

// NewHyperflexSoftwareVersionPolicyAllOf instantiates a new HyperflexSoftwareVersionPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSoftwareVersionPolicyAllOf(classId string, objectType string) *HyperflexSoftwareVersionPolicyAllOf {
	this := HyperflexSoftwareVersionPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSoftwareVersionPolicyAllOfWithDefaults instantiates a new HyperflexSoftwareVersionPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSoftwareVersionPolicyAllOfWithDefaults() *HyperflexSoftwareVersionPolicyAllOf {
	this := HyperflexSoftwareVersionPolicyAllOf{}
	var classId string = "hyperflex.SoftwareVersionPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.SoftwareVersionPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSoftwareVersionPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSoftwareVersionPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSoftwareVersionPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSoftwareVersionPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHxdpVersion returns the HxdpVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersion() string {
	if o == nil || o.HxdpVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpVersion
}

// GetHxdpVersionOk returns a tuple with the HxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersionOk() (*string, bool) {
	if o == nil || o.HxdpVersion == nil {
		return nil, false
	}
	return o.HxdpVersion, true
}

// HasHxdpVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasHxdpVersion() bool {
	if o != nil && o.HxdpVersion != nil {
		return true
	}

	return false
}

// SetHxdpVersion gets a reference to the given string and assigns it to the HxdpVersion field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetHxdpVersion(v string) {
	o.HxdpVersion = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersion() string {
	if o == nil || o.HypervisorVersion == nil {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || o.HypervisorVersion == nil {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasHypervisorVersion() bool {
	if o != nil && o.HypervisorVersion != nil {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersion() string {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersionOk() (*string, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given string and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetServerFirmwareVersion(v string) {
	o.ServerFirmwareVersion = &v
}

// GetServerFirmwareVersions returns the ServerFirmwareVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersions() []HyperflexServerFirmwareVersionInfo {
	if o == nil {
		var ret []HyperflexServerFirmwareVersionInfo
		return ret
	}
	return o.ServerFirmwareVersions
}

// GetServerFirmwareVersionsOk returns a tuple with the ServerFirmwareVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersionsOk() ([]HyperflexServerFirmwareVersionInfo, bool) {
	if o == nil || o.ServerFirmwareVersions == nil {
		return nil, false
	}
	return o.ServerFirmwareVersions, true
}

// HasServerFirmwareVersions returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasServerFirmwareVersions() bool {
	if o != nil && o.ServerFirmwareVersions != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersions gets a reference to the given []HyperflexServerFirmwareVersionInfo and assigns it to the ServerFirmwareVersions field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetServerFirmwareVersions(v []HyperflexServerFirmwareVersionInfo) {
	o.ServerFirmwareVersions = v
}

// GetUpgradeTypes returns the UpgradeTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareVersionPolicyAllOf) GetUpgradeTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UpgradeTypes
}

// GetUpgradeTypesOk returns a tuple with the UpgradeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareVersionPolicyAllOf) GetUpgradeTypesOk() ([]string, bool) {
	if o == nil || o.UpgradeTypes == nil {
		return nil, false
	}
	return o.UpgradeTypes, true
}

// HasUpgradeTypes returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasUpgradeTypes() bool {
	if o != nil && o.UpgradeTypes != nil {
		return true
	}

	return false
}

// SetUpgradeTypes gets a reference to the given []string and assigns it to the UpgradeTypes field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetUpgradeTypes(v []string) {
	o.UpgradeTypes = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareVersionPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareVersionPolicyAllOf) GetClusterProfilesOk() ([]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetHxdpVersionInfo returns the HxdpVersionInfo field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersionInfo() SoftwareHyperflexDistributableRelationship {
	if o == nil || o.HxdpVersionInfo == nil {
		var ret SoftwareHyperflexDistributableRelationship
		return ret
	}
	return *o.HxdpVersionInfo
}

// GetHxdpVersionInfoOk returns a tuple with the HxdpVersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool) {
	if o == nil || o.HxdpVersionInfo == nil {
		return nil, false
	}
	return o.HxdpVersionInfo, true
}

// HasHxdpVersionInfo returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasHxdpVersionInfo() bool {
	if o != nil && o.HxdpVersionInfo != nil {
		return true
	}

	return false
}

// SetHxdpVersionInfo gets a reference to the given SoftwareHyperflexDistributableRelationship and assigns it to the HxdpVersionInfo field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetHxdpVersionInfo(v SoftwareHyperflexDistributableRelationship) {
	o.HxdpVersionInfo = &v
}

// GetHypervisorVersionInfo returns the HypervisorVersionInfo field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersionInfo() SoftwareHyperflexDistributableRelationship {
	if o == nil || o.HypervisorVersionInfo == nil {
		var ret SoftwareHyperflexDistributableRelationship
		return ret
	}
	return *o.HypervisorVersionInfo
}

// GetHypervisorVersionInfoOk returns a tuple with the HypervisorVersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool) {
	if o == nil || o.HypervisorVersionInfo == nil {
		return nil, false
	}
	return o.HypervisorVersionInfo, true
}

// HasHypervisorVersionInfo returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasHypervisorVersionInfo() bool {
	if o != nil && o.HypervisorVersionInfo != nil {
		return true
	}

	return false
}

// SetHypervisorVersionInfo gets a reference to the given SoftwareHyperflexDistributableRelationship and assigns it to the HypervisorVersionInfo field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetHypervisorVersionInfo(v SoftwareHyperflexDistributableRelationship) {
	o.HypervisorVersionInfo = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetServerFirmwareVersionInfo returns the ServerFirmwareVersionInfo field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersionInfo() FirmwareDistributableRelationship {
	if o == nil || o.ServerFirmwareVersionInfo == nil {
		var ret FirmwareDistributableRelationship
		return ret
	}
	return *o.ServerFirmwareVersionInfo
}

// GetServerFirmwareVersionInfoOk returns a tuple with the ServerFirmwareVersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersionInfoOk() (*FirmwareDistributableRelationship, bool) {
	if o == nil || o.ServerFirmwareVersionInfo == nil {
		return nil, false
	}
	return o.ServerFirmwareVersionInfo, true
}

// HasServerFirmwareVersionInfo returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicyAllOf) HasServerFirmwareVersionInfo() bool {
	if o != nil && o.ServerFirmwareVersionInfo != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersionInfo gets a reference to the given FirmwareDistributableRelationship and assigns it to the ServerFirmwareVersionInfo field.
func (o *HyperflexSoftwareVersionPolicyAllOf) SetServerFirmwareVersionInfo(v FirmwareDistributableRelationship) {
	o.ServerFirmwareVersionInfo = &v
}

func (o HyperflexSoftwareVersionPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HxdpVersion != nil {
		toSerialize["HxdpVersion"] = o.HxdpVersion
	}
	if o.HypervisorVersion != nil {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}
	if o.ServerFirmwareVersions != nil {
		toSerialize["ServerFirmwareVersions"] = o.ServerFirmwareVersions
	}
	if o.UpgradeTypes != nil {
		toSerialize["UpgradeTypes"] = o.UpgradeTypes
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.HxdpVersionInfo != nil {
		toSerialize["HxdpVersionInfo"] = o.HxdpVersionInfo
	}
	if o.HypervisorVersionInfo != nil {
		toSerialize["HypervisorVersionInfo"] = o.HypervisorVersionInfo
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ServerFirmwareVersionInfo != nil {
		toSerialize["ServerFirmwareVersionInfo"] = o.ServerFirmwareVersionInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSoftwareVersionPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexSoftwareVersionPolicyAllOf := _HyperflexSoftwareVersionPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexSoftwareVersionPolicyAllOf); err == nil {
		*o = HyperflexSoftwareVersionPolicyAllOf(varHyperflexSoftwareVersionPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HxdpVersion")
		delete(additionalProperties, "HypervisorVersion")
		delete(additionalProperties, "ServerFirmwareVersion")
		delete(additionalProperties, "ServerFirmwareVersions")
		delete(additionalProperties, "UpgradeTypes")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "HxdpVersionInfo")
		delete(additionalProperties, "HypervisorVersionInfo")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "ServerFirmwareVersionInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexSoftwareVersionPolicyAllOf struct {
	value *HyperflexSoftwareVersionPolicyAllOf
	isSet bool
}

func (v NullableHyperflexSoftwareVersionPolicyAllOf) Get() *HyperflexSoftwareVersionPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexSoftwareVersionPolicyAllOf) Set(val *HyperflexSoftwareVersionPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSoftwareVersionPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSoftwareVersionPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSoftwareVersionPolicyAllOf(val *HyperflexSoftwareVersionPolicyAllOf) *NullableHyperflexSoftwareVersionPolicyAllOf {
	return &NullableHyperflexSoftwareVersionPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexSoftwareVersionPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSoftwareVersionPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
