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

// StorageNetAppLunMap NetApp LUN mapping is the process of associating a LUN with an igroup. When a LUN is mapped to an igroup, initiators in the igroup are granted access to the LUN.
type StorageNetAppLunMap struct {
	StorageBaseHostLun
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UUID of the initiator group.
	IgroupUuid *string `json:"IgroupUuid,omitempty"`
	// Universally unique identifier of the LUN.
	LunUuid *string `json:"LunUuid,omitempty"`
	// An array of relationships to storageNetAppInitiatorGroup resources.
	Host                 []StorageNetAppInitiatorGroupRelationship `json:"Host,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship       `json:"Tenant,omitempty"`
	Volume               *StorageNetAppLunRelationship             `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppLunMap StorageNetAppLunMap

// NewStorageNetAppLunMap instantiates a new StorageNetAppLunMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppLunMap(classId string, objectType string) *StorageNetAppLunMap {
	this := StorageNetAppLunMap{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppLunMapWithDefaults instantiates a new StorageNetAppLunMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppLunMapWithDefaults() *StorageNetAppLunMap {
	this := StorageNetAppLunMap{}
	var classId string = "storage.NetAppLunMap"
	this.ClassId = classId
	var objectType string = "storage.NetAppLunMap"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppLunMap) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMap) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppLunMap) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppLunMap) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMap) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppLunMap) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIgroupUuid returns the IgroupUuid field value if set, zero value otherwise.
func (o *StorageNetAppLunMap) GetIgroupUuid() string {
	if o == nil || o.IgroupUuid == nil {
		var ret string
		return ret
	}
	return *o.IgroupUuid
}

// GetIgroupUuidOk returns a tuple with the IgroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMap) GetIgroupUuidOk() (*string, bool) {
	if o == nil || o.IgroupUuid == nil {
		return nil, false
	}
	return o.IgroupUuid, true
}

// HasIgroupUuid returns a boolean if a field has been set.
func (o *StorageNetAppLunMap) HasIgroupUuid() bool {
	if o != nil && o.IgroupUuid != nil {
		return true
	}

	return false
}

// SetIgroupUuid gets a reference to the given string and assigns it to the IgroupUuid field.
func (o *StorageNetAppLunMap) SetIgroupUuid(v string) {
	o.IgroupUuid = &v
}

// GetLunUuid returns the LunUuid field value if set, zero value otherwise.
func (o *StorageNetAppLunMap) GetLunUuid() string {
	if o == nil || o.LunUuid == nil {
		var ret string
		return ret
	}
	return *o.LunUuid
}

// GetLunUuidOk returns a tuple with the LunUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMap) GetLunUuidOk() (*string, bool) {
	if o == nil || o.LunUuid == nil {
		return nil, false
	}
	return o.LunUuid, true
}

// HasLunUuid returns a boolean if a field has been set.
func (o *StorageNetAppLunMap) HasLunUuid() bool {
	if o != nil && o.LunUuid != nil {
		return true
	}

	return false
}

// SetLunUuid gets a reference to the given string and assigns it to the LunUuid field.
func (o *StorageNetAppLunMap) SetLunUuid(v string) {
	o.LunUuid = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppLunMap) GetHost() []StorageNetAppInitiatorGroupRelationship {
	if o == nil {
		var ret []StorageNetAppInitiatorGroupRelationship
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppLunMap) GetHostOk() ([]StorageNetAppInitiatorGroupRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *StorageNetAppLunMap) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given []StorageNetAppInitiatorGroupRelationship and assigns it to the Host field.
func (o *StorageNetAppLunMap) SetHost(v []StorageNetAppInitiatorGroupRelationship) {
	o.Host = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppLunMap) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMap) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppLunMap) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppLunMap) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StorageNetAppLunMap) GetVolume() StorageNetAppLunRelationship {
	if o == nil || o.Volume == nil {
		var ret StorageNetAppLunRelationship
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMap) GetVolumeOk() (*StorageNetAppLunRelationship, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageNetAppLunMap) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given StorageNetAppLunRelationship and assigns it to the Volume field.
func (o *StorageNetAppLunMap) SetVolume(v StorageNetAppLunRelationship) {
	o.Volume = &v
}

func (o StorageNetAppLunMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHostLun, errStorageBaseHostLun := json.Marshal(o.StorageBaseHostLun)
	if errStorageBaseHostLun != nil {
		return []byte{}, errStorageBaseHostLun
	}
	errStorageBaseHostLun = json.Unmarshal([]byte(serializedStorageBaseHostLun), &toSerialize)
	if errStorageBaseHostLun != nil {
		return []byte{}, errStorageBaseHostLun
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IgroupUuid != nil {
		toSerialize["IgroupUuid"] = o.IgroupUuid
	}
	if o.LunUuid != nil {
		toSerialize["LunUuid"] = o.LunUuid
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppLunMap) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppLunMapWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// UUID of the initiator group.
		IgroupUuid *string `json:"IgroupUuid,omitempty"`
		// Universally unique identifier of the LUN.
		LunUuid *string `json:"LunUuid,omitempty"`
		// An array of relationships to storageNetAppInitiatorGroup resources.
		Host   []StorageNetAppInitiatorGroupRelationship `json:"Host,omitempty"`
		Tenant *StorageNetAppStorageVmRelationship       `json:"Tenant,omitempty"`
		Volume *StorageNetAppLunRelationship             `json:"Volume,omitempty"`
	}

	varStorageNetAppLunMapWithoutEmbeddedStruct := StorageNetAppLunMapWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppLunMapWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppLunMap := _StorageNetAppLunMap{}
		varStorageNetAppLunMap.ClassId = varStorageNetAppLunMapWithoutEmbeddedStruct.ClassId
		varStorageNetAppLunMap.ObjectType = varStorageNetAppLunMapWithoutEmbeddedStruct.ObjectType
		varStorageNetAppLunMap.IgroupUuid = varStorageNetAppLunMapWithoutEmbeddedStruct.IgroupUuid
		varStorageNetAppLunMap.LunUuid = varStorageNetAppLunMapWithoutEmbeddedStruct.LunUuid
		varStorageNetAppLunMap.Host = varStorageNetAppLunMapWithoutEmbeddedStruct.Host
		varStorageNetAppLunMap.Tenant = varStorageNetAppLunMapWithoutEmbeddedStruct.Tenant
		varStorageNetAppLunMap.Volume = varStorageNetAppLunMapWithoutEmbeddedStruct.Volume
		*o = StorageNetAppLunMap(varStorageNetAppLunMap)
	} else {
		return err
	}

	varStorageNetAppLunMap := _StorageNetAppLunMap{}

	err = json.Unmarshal(bytes, &varStorageNetAppLunMap)
	if err == nil {
		o.StorageBaseHostLun = varStorageNetAppLunMap.StorageBaseHostLun
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IgroupUuid")
		delete(additionalProperties, "LunUuid")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "Tenant")
		delete(additionalProperties, "Volume")

		// remove fields from embedded structs
		reflectStorageBaseHostLun := reflect.ValueOf(o.StorageBaseHostLun)
		for i := 0; i < reflectStorageBaseHostLun.Type().NumField(); i++ {
			t := reflectStorageBaseHostLun.Type().Field(i)

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

type NullableStorageNetAppLunMap struct {
	value *StorageNetAppLunMap
	isSet bool
}

func (v NullableStorageNetAppLunMap) Get() *StorageNetAppLunMap {
	return v.value
}

func (v *NullableStorageNetAppLunMap) Set(val *StorageNetAppLunMap) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppLunMap) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppLunMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppLunMap(val *StorageNetAppLunMap) *NullableStorageNetAppLunMap {
	return &NullableStorageNetAppLunMap{value: val, isSet: true}
}

func (v NullableStorageNetAppLunMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppLunMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
