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

// SnmpPolicyInventoryAllOf Definition of the list of properties defined in 'snmp.PolicyInventory', excluding properties defined in parent classes.
type SnmpPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long.
	AccessCommunityString *string `json:"AccessCommunityString,omitempty"`
	// Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users. * `Disabled` - Blocks access to the information in the inventory tables. * `Limited` - Partial access to read the information in the inventory tables. * `Full` - Full access to read the information in the inventory tables.
	CommunityAccess *string `json:"CommunityAccess,omitempty"`
	// State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host.
	Enabled *bool `json:"Enabled,omitempty"`
	// User-defined unique identification of the static engine.
	EngineId *string `json:"EngineId,omitempty"`
	// Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269).
	SnmpPort  *int64     `json:"SnmpPort,omitempty"`
	SnmpTraps []SnmpTrap `json:"SnmpTraps,omitempty"`
	SnmpUsers []SnmpUser `json:"SnmpUsers,omitempty"`
	// Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number.
	SysContact *string `json:"SysContact,omitempty"`
	// Location of host on which the SNMP agent (server) runs.
	SysLocation *string `json:"SysLocation,omitempty"`
	// SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users.
	TrapCommunity *string `json:"TrapCommunity,omitempty"`
	// State of the SNMP v2c on the endpoint. If enabled, the endpoint sends SNMP v2c properties to the designated host.
	V2Enabled *bool `json:"V2Enabled,omitempty"`
	// State of the SNMP v3 on the endpoint. If enabled, the endpoint sends SNMP v3 properties to the designated host.
	V3Enabled            *bool                 `json:"V3Enabled,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpPolicyInventoryAllOf SnmpPolicyInventoryAllOf

// NewSnmpPolicyInventoryAllOf instantiates a new SnmpPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpPolicyInventoryAllOf(classId string, objectType string) *SnmpPolicyInventoryAllOf {
	this := SnmpPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSnmpPolicyInventoryAllOfWithDefaults instantiates a new SnmpPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpPolicyInventoryAllOfWithDefaults() *SnmpPolicyInventoryAllOf {
	this := SnmpPolicyInventoryAllOf{}
	var classId string = "snmp.PolicyInventory"
	this.ClassId = classId
	var objectType string = "snmp.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SnmpPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SnmpPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SnmpPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SnmpPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessCommunityString returns the AccessCommunityString field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetAccessCommunityString() string {
	if o == nil || o.AccessCommunityString == nil {
		var ret string
		return ret
	}
	return *o.AccessCommunityString
}

// GetAccessCommunityStringOk returns a tuple with the AccessCommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetAccessCommunityStringOk() (*string, bool) {
	if o == nil || o.AccessCommunityString == nil {
		return nil, false
	}
	return o.AccessCommunityString, true
}

// HasAccessCommunityString returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasAccessCommunityString() bool {
	if o != nil && o.AccessCommunityString != nil {
		return true
	}

	return false
}

// SetAccessCommunityString gets a reference to the given string and assigns it to the AccessCommunityString field.
func (o *SnmpPolicyInventoryAllOf) SetAccessCommunityString(v string) {
	o.AccessCommunityString = &v
}

// GetCommunityAccess returns the CommunityAccess field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetCommunityAccess() string {
	if o == nil || o.CommunityAccess == nil {
		var ret string
		return ret
	}
	return *o.CommunityAccess
}

// GetCommunityAccessOk returns a tuple with the CommunityAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetCommunityAccessOk() (*string, bool) {
	if o == nil || o.CommunityAccess == nil {
		return nil, false
	}
	return o.CommunityAccess, true
}

// HasCommunityAccess returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasCommunityAccess() bool {
	if o != nil && o.CommunityAccess != nil {
		return true
	}

	return false
}

// SetCommunityAccess gets a reference to the given string and assigns it to the CommunityAccess field.
func (o *SnmpPolicyInventoryAllOf) SetCommunityAccess(v string) {
	o.CommunityAccess = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SnmpPolicyInventoryAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetEngineId() string {
	if o == nil || o.EngineId == nil {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetEngineIdOk() (*string, bool) {
	if o == nil || o.EngineId == nil {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasEngineId() bool {
	if o != nil && o.EngineId != nil {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *SnmpPolicyInventoryAllOf) SetEngineId(v string) {
	o.EngineId = &v
}

// GetSnmpPort returns the SnmpPort field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetSnmpPort() int64 {
	if o == nil || o.SnmpPort == nil {
		var ret int64
		return ret
	}
	return *o.SnmpPort
}

// GetSnmpPortOk returns a tuple with the SnmpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetSnmpPortOk() (*int64, bool) {
	if o == nil || o.SnmpPort == nil {
		return nil, false
	}
	return o.SnmpPort, true
}

// HasSnmpPort returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasSnmpPort() bool {
	if o != nil && o.SnmpPort != nil {
		return true
	}

	return false
}

// SetSnmpPort gets a reference to the given int64 and assigns it to the SnmpPort field.
func (o *SnmpPolicyInventoryAllOf) SetSnmpPort(v int64) {
	o.SnmpPort = &v
}

// GetSnmpTraps returns the SnmpTraps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnmpPolicyInventoryAllOf) GetSnmpTraps() []SnmpTrap {
	if o == nil {
		var ret []SnmpTrap
		return ret
	}
	return o.SnmpTraps
}

// GetSnmpTrapsOk returns a tuple with the SnmpTraps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnmpPolicyInventoryAllOf) GetSnmpTrapsOk() ([]SnmpTrap, bool) {
	if o == nil || o.SnmpTraps == nil {
		return nil, false
	}
	return o.SnmpTraps, true
}

// HasSnmpTraps returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasSnmpTraps() bool {
	if o != nil && o.SnmpTraps != nil {
		return true
	}

	return false
}

// SetSnmpTraps gets a reference to the given []SnmpTrap and assigns it to the SnmpTraps field.
func (o *SnmpPolicyInventoryAllOf) SetSnmpTraps(v []SnmpTrap) {
	o.SnmpTraps = v
}

// GetSnmpUsers returns the SnmpUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnmpPolicyInventoryAllOf) GetSnmpUsers() []SnmpUser {
	if o == nil {
		var ret []SnmpUser
		return ret
	}
	return o.SnmpUsers
}

// GetSnmpUsersOk returns a tuple with the SnmpUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnmpPolicyInventoryAllOf) GetSnmpUsersOk() ([]SnmpUser, bool) {
	if o == nil || o.SnmpUsers == nil {
		return nil, false
	}
	return o.SnmpUsers, true
}

// HasSnmpUsers returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasSnmpUsers() bool {
	if o != nil && o.SnmpUsers != nil {
		return true
	}

	return false
}

// SetSnmpUsers gets a reference to the given []SnmpUser and assigns it to the SnmpUsers field.
func (o *SnmpPolicyInventoryAllOf) SetSnmpUsers(v []SnmpUser) {
	o.SnmpUsers = v
}

// GetSysContact returns the SysContact field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetSysContact() string {
	if o == nil || o.SysContact == nil {
		var ret string
		return ret
	}
	return *o.SysContact
}

// GetSysContactOk returns a tuple with the SysContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetSysContactOk() (*string, bool) {
	if o == nil || o.SysContact == nil {
		return nil, false
	}
	return o.SysContact, true
}

// HasSysContact returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasSysContact() bool {
	if o != nil && o.SysContact != nil {
		return true
	}

	return false
}

// SetSysContact gets a reference to the given string and assigns it to the SysContact field.
func (o *SnmpPolicyInventoryAllOf) SetSysContact(v string) {
	o.SysContact = &v
}

// GetSysLocation returns the SysLocation field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetSysLocation() string {
	if o == nil || o.SysLocation == nil {
		var ret string
		return ret
	}
	return *o.SysLocation
}

// GetSysLocationOk returns a tuple with the SysLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetSysLocationOk() (*string, bool) {
	if o == nil || o.SysLocation == nil {
		return nil, false
	}
	return o.SysLocation, true
}

// HasSysLocation returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasSysLocation() bool {
	if o != nil && o.SysLocation != nil {
		return true
	}

	return false
}

// SetSysLocation gets a reference to the given string and assigns it to the SysLocation field.
func (o *SnmpPolicyInventoryAllOf) SetSysLocation(v string) {
	o.SysLocation = &v
}

// GetTrapCommunity returns the TrapCommunity field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetTrapCommunity() string {
	if o == nil || o.TrapCommunity == nil {
		var ret string
		return ret
	}
	return *o.TrapCommunity
}

// GetTrapCommunityOk returns a tuple with the TrapCommunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetTrapCommunityOk() (*string, bool) {
	if o == nil || o.TrapCommunity == nil {
		return nil, false
	}
	return o.TrapCommunity, true
}

// HasTrapCommunity returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasTrapCommunity() bool {
	if o != nil && o.TrapCommunity != nil {
		return true
	}

	return false
}

// SetTrapCommunity gets a reference to the given string and assigns it to the TrapCommunity field.
func (o *SnmpPolicyInventoryAllOf) SetTrapCommunity(v string) {
	o.TrapCommunity = &v
}

// GetV2Enabled returns the V2Enabled field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetV2Enabled() bool {
	if o == nil || o.V2Enabled == nil {
		var ret bool
		return ret
	}
	return *o.V2Enabled
}

// GetV2EnabledOk returns a tuple with the V2Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetV2EnabledOk() (*bool, bool) {
	if o == nil || o.V2Enabled == nil {
		return nil, false
	}
	return o.V2Enabled, true
}

// HasV2Enabled returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasV2Enabled() bool {
	if o != nil && o.V2Enabled != nil {
		return true
	}

	return false
}

// SetV2Enabled gets a reference to the given bool and assigns it to the V2Enabled field.
func (o *SnmpPolicyInventoryAllOf) SetV2Enabled(v bool) {
	o.V2Enabled = &v
}

// GetV3Enabled returns the V3Enabled field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetV3Enabled() bool {
	if o == nil || o.V3Enabled == nil {
		var ret bool
		return ret
	}
	return *o.V3Enabled
}

// GetV3EnabledOk returns a tuple with the V3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetV3EnabledOk() (*bool, bool) {
	if o == nil || o.V3Enabled == nil {
		return nil, false
	}
	return o.V3Enabled, true
}

// HasV3Enabled returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasV3Enabled() bool {
	if o != nil && o.V3Enabled != nil {
		return true
	}

	return false
}

// SetV3Enabled gets a reference to the given bool and assigns it to the V3Enabled field.
func (o *SnmpPolicyInventoryAllOf) SetV3Enabled(v bool) {
	o.V3Enabled = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *SnmpPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *SnmpPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *SnmpPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o SnmpPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessCommunityString != nil {
		toSerialize["AccessCommunityString"] = o.AccessCommunityString
	}
	if o.CommunityAccess != nil {
		toSerialize["CommunityAccess"] = o.CommunityAccess
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.EngineId != nil {
		toSerialize["EngineId"] = o.EngineId
	}
	if o.SnmpPort != nil {
		toSerialize["SnmpPort"] = o.SnmpPort
	}
	if o.SnmpTraps != nil {
		toSerialize["SnmpTraps"] = o.SnmpTraps
	}
	if o.SnmpUsers != nil {
		toSerialize["SnmpUsers"] = o.SnmpUsers
	}
	if o.SysContact != nil {
		toSerialize["SysContact"] = o.SysContact
	}
	if o.SysLocation != nil {
		toSerialize["SysLocation"] = o.SysLocation
	}
	if o.TrapCommunity != nil {
		toSerialize["TrapCommunity"] = o.TrapCommunity
	}
	if o.V2Enabled != nil {
		toSerialize["V2Enabled"] = o.V2Enabled
	}
	if o.V3Enabled != nil {
		toSerialize["V3Enabled"] = o.V3Enabled
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SnmpPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSnmpPolicyInventoryAllOf := _SnmpPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varSnmpPolicyInventoryAllOf); err == nil {
		*o = SnmpPolicyInventoryAllOf(varSnmpPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessCommunityString")
		delete(additionalProperties, "CommunityAccess")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "EngineId")
		delete(additionalProperties, "SnmpPort")
		delete(additionalProperties, "SnmpTraps")
		delete(additionalProperties, "SnmpUsers")
		delete(additionalProperties, "SysContact")
		delete(additionalProperties, "SysLocation")
		delete(additionalProperties, "TrapCommunity")
		delete(additionalProperties, "V2Enabled")
		delete(additionalProperties, "V3Enabled")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpPolicyInventoryAllOf struct {
	value *SnmpPolicyInventoryAllOf
	isSet bool
}

func (v NullableSnmpPolicyInventoryAllOf) Get() *SnmpPolicyInventoryAllOf {
	return v.value
}

func (v *NullableSnmpPolicyInventoryAllOf) Set(val *SnmpPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpPolicyInventoryAllOf(val *SnmpPolicyInventoryAllOf) *NullableSnmpPolicyInventoryAllOf {
	return &NullableSnmpPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableSnmpPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
