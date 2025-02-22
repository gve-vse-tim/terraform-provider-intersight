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

// HyperflexSnapshotFilesAllOf Definition of the list of properties defined in 'hyperflex.SnapshotFiles', excluding properties defined in parent classes.
type HyperflexSnapshotFilesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                          `json:"ObjectType"`
	NameTrackedFiles     []HyperflexFilePath             `json:"NameTrackedFiles,omitempty"`
	UuidTrackedDisksMap  []HyperflexMapUuidToTrackedDisk `json:"UuidTrackedDisksMap,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexSnapshotFilesAllOf HyperflexSnapshotFilesAllOf

// NewHyperflexSnapshotFilesAllOf instantiates a new HyperflexSnapshotFilesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSnapshotFilesAllOf(classId string, objectType string) *HyperflexSnapshotFilesAllOf {
	this := HyperflexSnapshotFilesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSnapshotFilesAllOfWithDefaults instantiates a new HyperflexSnapshotFilesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSnapshotFilesAllOfWithDefaults() *HyperflexSnapshotFilesAllOf {
	this := HyperflexSnapshotFilesAllOf{}
	var classId string = "hyperflex.SnapshotFiles"
	this.ClassId = classId
	var objectType string = "hyperflex.SnapshotFiles"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSnapshotFilesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotFilesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSnapshotFilesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSnapshotFilesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotFilesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSnapshotFilesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNameTrackedFiles returns the NameTrackedFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotFilesAllOf) GetNameTrackedFiles() []HyperflexFilePath {
	if o == nil {
		var ret []HyperflexFilePath
		return ret
	}
	return o.NameTrackedFiles
}

// GetNameTrackedFilesOk returns a tuple with the NameTrackedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotFilesAllOf) GetNameTrackedFilesOk() ([]HyperflexFilePath, bool) {
	if o == nil || o.NameTrackedFiles == nil {
		return nil, false
	}
	return o.NameTrackedFiles, true
}

// HasNameTrackedFiles returns a boolean if a field has been set.
func (o *HyperflexSnapshotFilesAllOf) HasNameTrackedFiles() bool {
	if o != nil && o.NameTrackedFiles != nil {
		return true
	}

	return false
}

// SetNameTrackedFiles gets a reference to the given []HyperflexFilePath and assigns it to the NameTrackedFiles field.
func (o *HyperflexSnapshotFilesAllOf) SetNameTrackedFiles(v []HyperflexFilePath) {
	o.NameTrackedFiles = v
}

// GetUuidTrackedDisksMap returns the UuidTrackedDisksMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotFilesAllOf) GetUuidTrackedDisksMap() []HyperflexMapUuidToTrackedDisk {
	if o == nil {
		var ret []HyperflexMapUuidToTrackedDisk
		return ret
	}
	return o.UuidTrackedDisksMap
}

// GetUuidTrackedDisksMapOk returns a tuple with the UuidTrackedDisksMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotFilesAllOf) GetUuidTrackedDisksMapOk() ([]HyperflexMapUuidToTrackedDisk, bool) {
	if o == nil || o.UuidTrackedDisksMap == nil {
		return nil, false
	}
	return o.UuidTrackedDisksMap, true
}

// HasUuidTrackedDisksMap returns a boolean if a field has been set.
func (o *HyperflexSnapshotFilesAllOf) HasUuidTrackedDisksMap() bool {
	if o != nil && o.UuidTrackedDisksMap != nil {
		return true
	}

	return false
}

// SetUuidTrackedDisksMap gets a reference to the given []HyperflexMapUuidToTrackedDisk and assigns it to the UuidTrackedDisksMap field.
func (o *HyperflexSnapshotFilesAllOf) SetUuidTrackedDisksMap(v []HyperflexMapUuidToTrackedDisk) {
	o.UuidTrackedDisksMap = v
}

func (o HyperflexSnapshotFilesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NameTrackedFiles != nil {
		toSerialize["NameTrackedFiles"] = o.NameTrackedFiles
	}
	if o.UuidTrackedDisksMap != nil {
		toSerialize["UuidTrackedDisksMap"] = o.UuidTrackedDisksMap
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSnapshotFilesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexSnapshotFilesAllOf := _HyperflexSnapshotFilesAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexSnapshotFilesAllOf); err == nil {
		*o = HyperflexSnapshotFilesAllOf(varHyperflexSnapshotFilesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NameTrackedFiles")
		delete(additionalProperties, "UuidTrackedDisksMap")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexSnapshotFilesAllOf struct {
	value *HyperflexSnapshotFilesAllOf
	isSet bool
}

func (v NullableHyperflexSnapshotFilesAllOf) Get() *HyperflexSnapshotFilesAllOf {
	return v.value
}

func (v *NullableHyperflexSnapshotFilesAllOf) Set(val *HyperflexSnapshotFilesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSnapshotFilesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSnapshotFilesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSnapshotFilesAllOf(val *HyperflexSnapshotFilesAllOf) *NullableHyperflexSnapshotFilesAllOf {
	return &NullableHyperflexSnapshotFilesAllOf{value: val, isSet: true}
}

func (v NullableHyperflexSnapshotFilesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSnapshotFilesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
