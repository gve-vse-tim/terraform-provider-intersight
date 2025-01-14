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

// StorageNetAppNonDataIpInterfaceAllOf Definition of the list of properties defined in 'storage.NetAppNonDataIpInterface', excluding properties defined in parent classes.
type StorageNetAppNonDataIpInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                            `json:"ObjectType"`
	Array           *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	ArrayController *StorageNetAppNodeRelationship    `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppNonDataIpInterfaceEvent resources.
	Events               []StorageNetAppNonDataIpInterfaceEventRelationship `json:"Events,omitempty"`
	NetAppEthernetPort   *StorageNetAppEthernetPortRelationship             `json:"NetAppEthernetPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNonDataIpInterfaceAllOf StorageNetAppNonDataIpInterfaceAllOf

// NewStorageNetAppNonDataIpInterfaceAllOf instantiates a new StorageNetAppNonDataIpInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNonDataIpInterfaceAllOf(classId string, objectType string) *StorageNetAppNonDataIpInterfaceAllOf {
	this := StorageNetAppNonDataIpInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNonDataIpInterfaceAllOfWithDefaults instantiates a new StorageNetAppNonDataIpInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNonDataIpInterfaceAllOfWithDefaults() *StorageNetAppNonDataIpInterfaceAllOf {
	this := StorageNetAppNonDataIpInterfaceAllOf{}
	var classId string = "storage.NetAppNonDataIpInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppNonDataIpInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNonDataIpInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNonDataIpInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppNonDataIpInterfaceAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppNonDataIpInterfaceAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetEvents() []StorageNetAppNonDataIpInterfaceEventRelationship {
	if o == nil {
		var ret []StorageNetAppNonDataIpInterfaceEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetEventsOk() ([]StorageNetAppNonDataIpInterfaceEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppNonDataIpInterfaceEventRelationship and assigns it to the Events field.
func (o *StorageNetAppNonDataIpInterfaceAllOf) SetEvents(v []StorageNetAppNonDataIpInterfaceEventRelationship) {
	o.Events = v
}

// GetNetAppEthernetPort returns the NetAppEthernetPort field value if set, zero value otherwise.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship {
	if o == nil || o.NetAppEthernetPort == nil {
		var ret StorageNetAppEthernetPortRelationship
		return ret
	}
	return *o.NetAppEthernetPort
}

// GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool) {
	if o == nil || o.NetAppEthernetPort == nil {
		return nil, false
	}
	return o.NetAppEthernetPort, true
}

// HasNetAppEthernetPort returns a boolean if a field has been set.
func (o *StorageNetAppNonDataIpInterfaceAllOf) HasNetAppEthernetPort() bool {
	if o != nil && o.NetAppEthernetPort != nil {
		return true
	}

	return false
}

// SetNetAppEthernetPort gets a reference to the given StorageNetAppEthernetPortRelationship and assigns it to the NetAppEthernetPort field.
func (o *StorageNetAppNonDataIpInterfaceAllOf) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship) {
	o.NetAppEthernetPort = &v
}

func (o StorageNetAppNonDataIpInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.NetAppEthernetPort != nil {
		toSerialize["NetAppEthernetPort"] = o.NetAppEthernetPort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNonDataIpInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppNonDataIpInterfaceAllOf := _StorageNetAppNonDataIpInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppNonDataIpInterfaceAllOf); err == nil {
		*o = StorageNetAppNonDataIpInterfaceAllOf(varStorageNetAppNonDataIpInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "NetAppEthernetPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppNonDataIpInterfaceAllOf struct {
	value *StorageNetAppNonDataIpInterfaceAllOf
	isSet bool
}

func (v NullableStorageNetAppNonDataIpInterfaceAllOf) Get() *StorageNetAppNonDataIpInterfaceAllOf {
	return v.value
}

func (v *NullableStorageNetAppNonDataIpInterfaceAllOf) Set(val *StorageNetAppNonDataIpInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNonDataIpInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNonDataIpInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNonDataIpInterfaceAllOf(val *StorageNetAppNonDataIpInterfaceAllOf) *NullableStorageNetAppNonDataIpInterfaceAllOf {
	return &NullableStorageNetAppNonDataIpInterfaceAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppNonDataIpInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNonDataIpInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
