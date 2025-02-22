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

// VirtualizationEsxiVmStorageConfigurationAllOf Definition of the list of properties defined in 'virtualization.EsxiVmStorageConfiguration', excluding properties defined in parent classes.
type VirtualizationEsxiVmStorageConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Datastore where virtual machine is deployed.
	Datastore            *string                    `json:"Datastore,omitempty"`
	Disks                []VirtualizationVmEsxiDisk `json:"Disks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationEsxiVmStorageConfigurationAllOf VirtualizationEsxiVmStorageConfigurationAllOf

// NewVirtualizationEsxiVmStorageConfigurationAllOf instantiates a new VirtualizationEsxiVmStorageConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationEsxiVmStorageConfigurationAllOf(classId string, objectType string) *VirtualizationEsxiVmStorageConfigurationAllOf {
	this := VirtualizationEsxiVmStorageConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationEsxiVmStorageConfigurationAllOfWithDefaults instantiates a new VirtualizationEsxiVmStorageConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationEsxiVmStorageConfigurationAllOfWithDefaults() *VirtualizationEsxiVmStorageConfigurationAllOf {
	this := VirtualizationEsxiVmStorageConfigurationAllOf{}
	var classId string = "virtualization.EsxiVmStorageConfiguration"
	this.ClassId = classId
	var objectType string = "virtualization.EsxiVmStorageConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetDatastore(v string) {
	o.Datastore = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDisks() []VirtualizationVmEsxiDisk {
	if o == nil {
		var ret []VirtualizationVmEsxiDisk
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDisksOk() ([]VirtualizationVmEsxiDisk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []VirtualizationVmEsxiDisk and assigns it to the Disks field.
func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetDisks(v []VirtualizationVmEsxiDisk) {
	o.Disks = v
}

func (o VirtualizationEsxiVmStorageConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Datastore != nil {
		toSerialize["Datastore"] = o.Datastore
	}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationEsxiVmStorageConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationEsxiVmStorageConfigurationAllOf := _VirtualizationEsxiVmStorageConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationEsxiVmStorageConfigurationAllOf); err == nil {
		*o = VirtualizationEsxiVmStorageConfigurationAllOf(varVirtualizationEsxiVmStorageConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Datastore")
		delete(additionalProperties, "Disks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationEsxiVmStorageConfigurationAllOf struct {
	value *VirtualizationEsxiVmStorageConfigurationAllOf
	isSet bool
}

func (v NullableVirtualizationEsxiVmStorageConfigurationAllOf) Get() *VirtualizationEsxiVmStorageConfigurationAllOf {
	return v.value
}

func (v *NullableVirtualizationEsxiVmStorageConfigurationAllOf) Set(val *VirtualizationEsxiVmStorageConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationEsxiVmStorageConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationEsxiVmStorageConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationEsxiVmStorageConfigurationAllOf(val *VirtualizationEsxiVmStorageConfigurationAllOf) *NullableVirtualizationEsxiVmStorageConfigurationAllOf {
	return &NullableVirtualizationEsxiVmStorageConfigurationAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationEsxiVmStorageConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationEsxiVmStorageConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
