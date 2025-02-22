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
	"reflect"
	"strings"
)

// StorageHitachiArray The details of the Hitachi storage array.
type StorageHitachiArray struct {
	StorageBaseArray
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of controller 1 of the storage system.
	Ctl1Ip *string `json:"Ctl1Ip,omitempty"`
	// GUM (Gateway for Unified Management) version of the controller 1.
	Ctl1MicroVersion *string `json:"Ctl1MicroVersion,omitempty"`
	// IP address of controller 2 of the storage system.
	Ctl2Ip *string `json:"Ctl2Ip,omitempty"`
	// GUM (Gateway for Unified Management) version of the controller 2.
	Ctl2MicroVersion *string `json:"Ctl2MicroVersion,omitempty"`
	// ID of the Storage device.
	DeviceId *string `json:"DeviceId,omitempty"`
	// IP address of the SVP (Service Processor). The SVP provides out‑of‑band configuration and management of the storage system, and collects performance data for key components to enable diagnostic testing and analysis.
	SvpIp *string `json:"SvpIp,omitempty"`
	// Controller operated by the REST API.
	TargetCtl            *string                              `json:"TargetCtl,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiArray StorageHitachiArray

// NewStorageHitachiArray instantiates a new StorageHitachiArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiArray(classId string, objectType string) *StorageHitachiArray {
	this := StorageHitachiArray{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiArrayWithDefaults instantiates a new StorageHitachiArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiArrayWithDefaults() *StorageHitachiArray {
	this := StorageHitachiArray{}
	var classId string = "storage.HitachiArray"
	this.ClassId = classId
	var objectType string = "storage.HitachiArray"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiArray) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiArray) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiArray) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiArray) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCtl1Ip returns the Ctl1Ip field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetCtl1Ip() string {
	if o == nil || o.Ctl1Ip == nil {
		var ret string
		return ret
	}
	return *o.Ctl1Ip
}

// GetCtl1IpOk returns a tuple with the Ctl1Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetCtl1IpOk() (*string, bool) {
	if o == nil || o.Ctl1Ip == nil {
		return nil, false
	}
	return o.Ctl1Ip, true
}

// HasCtl1Ip returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasCtl1Ip() bool {
	if o != nil && o.Ctl1Ip != nil {
		return true
	}

	return false
}

// SetCtl1Ip gets a reference to the given string and assigns it to the Ctl1Ip field.
func (o *StorageHitachiArray) SetCtl1Ip(v string) {
	o.Ctl1Ip = &v
}

// GetCtl1MicroVersion returns the Ctl1MicroVersion field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetCtl1MicroVersion() string {
	if o == nil || o.Ctl1MicroVersion == nil {
		var ret string
		return ret
	}
	return *o.Ctl1MicroVersion
}

// GetCtl1MicroVersionOk returns a tuple with the Ctl1MicroVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetCtl1MicroVersionOk() (*string, bool) {
	if o == nil || o.Ctl1MicroVersion == nil {
		return nil, false
	}
	return o.Ctl1MicroVersion, true
}

// HasCtl1MicroVersion returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasCtl1MicroVersion() bool {
	if o != nil && o.Ctl1MicroVersion != nil {
		return true
	}

	return false
}

// SetCtl1MicroVersion gets a reference to the given string and assigns it to the Ctl1MicroVersion field.
func (o *StorageHitachiArray) SetCtl1MicroVersion(v string) {
	o.Ctl1MicroVersion = &v
}

// GetCtl2Ip returns the Ctl2Ip field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetCtl2Ip() string {
	if o == nil || o.Ctl2Ip == nil {
		var ret string
		return ret
	}
	return *o.Ctl2Ip
}

// GetCtl2IpOk returns a tuple with the Ctl2Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetCtl2IpOk() (*string, bool) {
	if o == nil || o.Ctl2Ip == nil {
		return nil, false
	}
	return o.Ctl2Ip, true
}

// HasCtl2Ip returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasCtl2Ip() bool {
	if o != nil && o.Ctl2Ip != nil {
		return true
	}

	return false
}

// SetCtl2Ip gets a reference to the given string and assigns it to the Ctl2Ip field.
func (o *StorageHitachiArray) SetCtl2Ip(v string) {
	o.Ctl2Ip = &v
}

// GetCtl2MicroVersion returns the Ctl2MicroVersion field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetCtl2MicroVersion() string {
	if o == nil || o.Ctl2MicroVersion == nil {
		var ret string
		return ret
	}
	return *o.Ctl2MicroVersion
}

// GetCtl2MicroVersionOk returns a tuple with the Ctl2MicroVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetCtl2MicroVersionOk() (*string, bool) {
	if o == nil || o.Ctl2MicroVersion == nil {
		return nil, false
	}
	return o.Ctl2MicroVersion, true
}

// HasCtl2MicroVersion returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasCtl2MicroVersion() bool {
	if o != nil && o.Ctl2MicroVersion != nil {
		return true
	}

	return false
}

// SetCtl2MicroVersion gets a reference to the given string and assigns it to the Ctl2MicroVersion field.
func (o *StorageHitachiArray) SetCtl2MicroVersion(v string) {
	o.Ctl2MicroVersion = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *StorageHitachiArray) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetSvpIp returns the SvpIp field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetSvpIp() string {
	if o == nil || o.SvpIp == nil {
		var ret string
		return ret
	}
	return *o.SvpIp
}

// GetSvpIpOk returns a tuple with the SvpIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetSvpIpOk() (*string, bool) {
	if o == nil || o.SvpIp == nil {
		return nil, false
	}
	return o.SvpIp, true
}

// HasSvpIp returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasSvpIp() bool {
	if o != nil && o.SvpIp != nil {
		return true
	}

	return false
}

// SetSvpIp gets a reference to the given string and assigns it to the SvpIp field.
func (o *StorageHitachiArray) SetSvpIp(v string) {
	o.SvpIp = &v
}

// GetTargetCtl returns the TargetCtl field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetTargetCtl() string {
	if o == nil || o.TargetCtl == nil {
		var ret string
		return ret
	}
	return *o.TargetCtl
}

// GetTargetCtlOk returns a tuple with the TargetCtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetTargetCtlOk() (*string, bool) {
	if o == nil || o.TargetCtl == nil {
		return nil, false
	}
	return o.TargetCtl, true
}

// HasTargetCtl returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasTargetCtl() bool {
	if o != nil && o.TargetCtl != nil {
		return true
	}

	return false
}

// SetTargetCtl gets a reference to the given string and assigns it to the TargetCtl field.
func (o *StorageHitachiArray) SetTargetCtl(v string) {
	o.TargetCtl = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiArray) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArray) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiArray) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiArray) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseArray, errStorageBaseArray := json.Marshal(o.StorageBaseArray)
	if errStorageBaseArray != nil {
		return []byte{}, errStorageBaseArray
	}
	errStorageBaseArray = json.Unmarshal([]byte(serializedStorageBaseArray), &toSerialize)
	if errStorageBaseArray != nil {
		return []byte{}, errStorageBaseArray
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Ctl1Ip != nil {
		toSerialize["Ctl1Ip"] = o.Ctl1Ip
	}
	if o.Ctl1MicroVersion != nil {
		toSerialize["Ctl1MicroVersion"] = o.Ctl1MicroVersion
	}
	if o.Ctl2Ip != nil {
		toSerialize["Ctl2Ip"] = o.Ctl2Ip
	}
	if o.Ctl2MicroVersion != nil {
		toSerialize["Ctl2MicroVersion"] = o.Ctl2MicroVersion
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.SvpIp != nil {
		toSerialize["SvpIp"] = o.SvpIp
	}
	if o.TargetCtl != nil {
		toSerialize["TargetCtl"] = o.TargetCtl
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiArray) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiArrayWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address of controller 1 of the storage system.
		Ctl1Ip *string `json:"Ctl1Ip,omitempty"`
		// GUM (Gateway for Unified Management) version of the controller 1.
		Ctl1MicroVersion *string `json:"Ctl1MicroVersion,omitempty"`
		// IP address of controller 2 of the storage system.
		Ctl2Ip *string `json:"Ctl2Ip,omitempty"`
		// GUM (Gateway for Unified Management) version of the controller 2.
		Ctl2MicroVersion *string `json:"Ctl2MicroVersion,omitempty"`
		// ID of the Storage device.
		DeviceId *string `json:"DeviceId,omitempty"`
		// IP address of the SVP (Service Processor). The SVP provides out‑of‑band configuration and management of the storage system, and collects performance data for key components to enable diagnostic testing and analysis.
		SvpIp *string `json:"SvpIp,omitempty"`
		// Controller operated by the REST API.
		TargetCtl        *string                              `json:"TargetCtl,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiArrayWithoutEmbeddedStruct := StorageHitachiArrayWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiArrayWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiArray := _StorageHitachiArray{}
		varStorageHitachiArray.ClassId = varStorageHitachiArrayWithoutEmbeddedStruct.ClassId
		varStorageHitachiArray.ObjectType = varStorageHitachiArrayWithoutEmbeddedStruct.ObjectType
		varStorageHitachiArray.Ctl1Ip = varStorageHitachiArrayWithoutEmbeddedStruct.Ctl1Ip
		varStorageHitachiArray.Ctl1MicroVersion = varStorageHitachiArrayWithoutEmbeddedStruct.Ctl1MicroVersion
		varStorageHitachiArray.Ctl2Ip = varStorageHitachiArrayWithoutEmbeddedStruct.Ctl2Ip
		varStorageHitachiArray.Ctl2MicroVersion = varStorageHitachiArrayWithoutEmbeddedStruct.Ctl2MicroVersion
		varStorageHitachiArray.DeviceId = varStorageHitachiArrayWithoutEmbeddedStruct.DeviceId
		varStorageHitachiArray.SvpIp = varStorageHitachiArrayWithoutEmbeddedStruct.SvpIp
		varStorageHitachiArray.TargetCtl = varStorageHitachiArrayWithoutEmbeddedStruct.TargetCtl
		varStorageHitachiArray.RegisteredDevice = varStorageHitachiArrayWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiArray(varStorageHitachiArray)
	} else {
		return err
	}

	varStorageHitachiArray := _StorageHitachiArray{}

	err = json.Unmarshal(bytes, &varStorageHitachiArray)
	if err == nil {
		o.StorageBaseArray = varStorageHitachiArray.StorageBaseArray
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ctl1Ip")
		delete(additionalProperties, "Ctl1MicroVersion")
		delete(additionalProperties, "Ctl2Ip")
		delete(additionalProperties, "Ctl2MicroVersion")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "SvpIp")
		delete(additionalProperties, "TargetCtl")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseArray := reflect.ValueOf(o.StorageBaseArray)
		for i := 0; i < reflectStorageBaseArray.Type().NumField(); i++ {
			t := reflectStorageBaseArray.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiArray struct {
	value *StorageHitachiArray
	isSet bool
}

func (v NullableStorageHitachiArray) Get() *StorageHitachiArray {
	return v.value
}

func (v *NullableStorageHitachiArray) Set(val *StorageHitachiArray) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiArray) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiArray(val *StorageHitachiArray) *NullableStorageHitachiArray {
	return &NullableStorageHitachiArray{value: val, isSet: true}
}

func (v NullableStorageHitachiArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
