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

// TamS3DataSourceAllOf Definition of the list of properties defined in 'tam.S3DataSource', excluding properties defined in parent classes.
type TamS3DataSourceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string          `json:"ObjectType"`
	Queries    []TamQueryEntry `json:"Queries,omitempty"`
	// Path used to access file in s3 containing data.
	S3Path               *string `json:"S3Path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamS3DataSourceAllOf TamS3DataSourceAllOf

// NewTamS3DataSourceAllOf instantiates a new TamS3DataSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamS3DataSourceAllOf(classId string, objectType string) *TamS3DataSourceAllOf {
	this := TamS3DataSourceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamS3DataSourceAllOfWithDefaults instantiates a new TamS3DataSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamS3DataSourceAllOfWithDefaults() *TamS3DataSourceAllOf {
	this := TamS3DataSourceAllOf{}
	var classId string = "tam.S3DataSource"
	this.ClassId = classId
	var objectType string = "tam.S3DataSource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamS3DataSourceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamS3DataSourceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamS3DataSourceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamS3DataSourceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamS3DataSourceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamS3DataSourceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetQueries returns the Queries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamS3DataSourceAllOf) GetQueries() []TamQueryEntry {
	if o == nil {
		var ret []TamQueryEntry
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamS3DataSourceAllOf) GetQueriesOk() ([]TamQueryEntry, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *TamS3DataSourceAllOf) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []TamQueryEntry and assigns it to the Queries field.
func (o *TamS3DataSourceAllOf) SetQueries(v []TamQueryEntry) {
	o.Queries = v
}

// GetS3Path returns the S3Path field value if set, zero value otherwise.
func (o *TamS3DataSourceAllOf) GetS3Path() string {
	if o == nil || o.S3Path == nil {
		var ret string
		return ret
	}
	return *o.S3Path
}

// GetS3PathOk returns a tuple with the S3Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamS3DataSourceAllOf) GetS3PathOk() (*string, bool) {
	if o == nil || o.S3Path == nil {
		return nil, false
	}
	return o.S3Path, true
}

// HasS3Path returns a boolean if a field has been set.
func (o *TamS3DataSourceAllOf) HasS3Path() bool {
	if o != nil && o.S3Path != nil {
		return true
	}

	return false
}

// SetS3Path gets a reference to the given string and assigns it to the S3Path field.
func (o *TamS3DataSourceAllOf) SetS3Path(v string) {
	o.S3Path = &v
}

func (o TamS3DataSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Queries != nil {
		toSerialize["Queries"] = o.Queries
	}
	if o.S3Path != nil {
		toSerialize["S3Path"] = o.S3Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamS3DataSourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamS3DataSourceAllOf := _TamS3DataSourceAllOf{}

	if err = json.Unmarshal(bytes, &varTamS3DataSourceAllOf); err == nil {
		*o = TamS3DataSourceAllOf(varTamS3DataSourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Queries")
		delete(additionalProperties, "S3Path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamS3DataSourceAllOf struct {
	value *TamS3DataSourceAllOf
	isSet bool
}

func (v NullableTamS3DataSourceAllOf) Get() *TamS3DataSourceAllOf {
	return v.value
}

func (v *NullableTamS3DataSourceAllOf) Set(val *TamS3DataSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamS3DataSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamS3DataSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamS3DataSourceAllOf(val *TamS3DataSourceAllOf) *NullableTamS3DataSourceAllOf {
	return &NullableTamS3DataSourceAllOf{value: val, isSet: true}
}

func (v NullableTamS3DataSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamS3DataSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
