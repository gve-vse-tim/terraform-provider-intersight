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

// FirmwareUpgradeAllOf Definition of the list of properties defined in 'firmware.Upgrade', excluding properties defined in parent classes.
type FirmwareUpgradeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType              string                                      `json:"ObjectType"`
	ExcludeComponentList    []string                                    `json:"ExcludeComponentList,omitempty"`
	ExcludeComponentPidList NullableFirmwareExcludeComponentPidListType `json:"ExcludeComponentPidList,omitempty"`
	Device                  *AssetDeviceRegistrationRelationship        `json:"Device,omitempty"`
	Server                  *ComputePhysicalRelationship                `json:"Server,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _FirmwareUpgradeAllOf FirmwareUpgradeAllOf

// NewFirmwareUpgradeAllOf instantiates a new FirmwareUpgradeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeAllOf(classId string, objectType string) *FirmwareUpgradeAllOf {
	this := FirmwareUpgradeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareUpgradeAllOfWithDefaults instantiates a new FirmwareUpgradeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeAllOfWithDefaults() *FirmwareUpgradeAllOf {
	this := FirmwareUpgradeAllOf{}
	var classId string = "firmware.Upgrade"
	this.ClassId = classId
	var objectType string = "firmware.Upgrade"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgradeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgradeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgradeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgradeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExcludeComponentList returns the ExcludeComponentList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeAllOf) GetExcludeComponentList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeComponentList
}

// GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeAllOf) GetExcludeComponentListOk() ([]string, bool) {
	if o == nil || o.ExcludeComponentList == nil {
		return nil, false
	}
	return o.ExcludeComponentList, true
}

// HasExcludeComponentList returns a boolean if a field has been set.
func (o *FirmwareUpgradeAllOf) HasExcludeComponentList() bool {
	if o != nil && o.ExcludeComponentList != nil {
		return true
	}

	return false
}

// SetExcludeComponentList gets a reference to the given []string and assigns it to the ExcludeComponentList field.
func (o *FirmwareUpgradeAllOf) SetExcludeComponentList(v []string) {
	o.ExcludeComponentList = v
}

// GetExcludeComponentPidList returns the ExcludeComponentPidList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeAllOf) GetExcludeComponentPidList() FirmwareExcludeComponentPidListType {
	if o == nil || o.ExcludeComponentPidList.Get() == nil {
		var ret FirmwareExcludeComponentPidListType
		return ret
	}
	return *o.ExcludeComponentPidList.Get()
}

// GetExcludeComponentPidListOk returns a tuple with the ExcludeComponentPidList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeAllOf) GetExcludeComponentPidListOk() (*FirmwareExcludeComponentPidListType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExcludeComponentPidList.Get(), o.ExcludeComponentPidList.IsSet()
}

// HasExcludeComponentPidList returns a boolean if a field has been set.
func (o *FirmwareUpgradeAllOf) HasExcludeComponentPidList() bool {
	if o != nil && o.ExcludeComponentPidList.IsSet() {
		return true
	}

	return false
}

// SetExcludeComponentPidList gets a reference to the given NullableFirmwareExcludeComponentPidListType and assigns it to the ExcludeComponentPidList field.
func (o *FirmwareUpgradeAllOf) SetExcludeComponentPidList(v FirmwareExcludeComponentPidListType) {
	o.ExcludeComponentPidList.Set(&v)
}

// SetExcludeComponentPidListNil sets the value for ExcludeComponentPidList to be an explicit nil
func (o *FirmwareUpgradeAllOf) SetExcludeComponentPidListNil() {
	o.ExcludeComponentPidList.Set(nil)
}

// UnsetExcludeComponentPidList ensures that no value is present for ExcludeComponentPidList, not even an explicit nil
func (o *FirmwareUpgradeAllOf) UnsetExcludeComponentPidList() {
	o.ExcludeComponentPidList.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FirmwareUpgradeAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareUpgradeAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareUpgradeAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *FirmwareUpgradeAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *FirmwareUpgradeAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *FirmwareUpgradeAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o FirmwareUpgradeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExcludeComponentList != nil {
		toSerialize["ExcludeComponentList"] = o.ExcludeComponentList
	}
	if o.ExcludeComponentPidList.IsSet() {
		toSerialize["ExcludeComponentPidList"] = o.ExcludeComponentPidList.Get()
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareUpgradeAllOf := _FirmwareUpgradeAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareUpgradeAllOf); err == nil {
		*o = FirmwareUpgradeAllOf(varFirmwareUpgradeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeComponentList")
		delete(additionalProperties, "ExcludeComponentPidList")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareUpgradeAllOf struct {
	value *FirmwareUpgradeAllOf
	isSet bool
}

func (v NullableFirmwareUpgradeAllOf) Get() *FirmwareUpgradeAllOf {
	return v.value
}

func (v *NullableFirmwareUpgradeAllOf) Set(val *FirmwareUpgradeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeAllOf(val *FirmwareUpgradeAllOf) *NullableFirmwareUpgradeAllOf {
	return &NullableFirmwareUpgradeAllOf{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
