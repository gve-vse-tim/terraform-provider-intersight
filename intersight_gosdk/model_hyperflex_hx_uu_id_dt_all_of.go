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

// HyperflexHxUuIdDtAllOf Definition of the list of properties defined in 'hyperflex.HxUuIdDt', excluding properties defined in parent classes.
type HyperflexHxUuIdDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string              `json:"ObjectType"`
	Links      []HyperflexHxLinkDt `json:"Links,omitempty"`
	// The unique identifier string of an entity.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxUuIdDtAllOf HyperflexHxUuIdDtAllOf

// NewHyperflexHxUuIdDtAllOf instantiates a new HyperflexHxUuIdDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxUuIdDtAllOf(classId string, objectType string) *HyperflexHxUuIdDtAllOf {
	this := HyperflexHxUuIdDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxUuIdDtAllOfWithDefaults instantiates a new HyperflexHxUuIdDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxUuIdDtAllOfWithDefaults() *HyperflexHxUuIdDtAllOf {
	this := HyperflexHxUuIdDtAllOf{}
	var classId string = "hyperflex.HxUuIdDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxUuIdDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxUuIdDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxUuIdDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxUuIdDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxUuIdDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxUuIdDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxUuIdDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxUuIdDtAllOf) GetLinks() []HyperflexHxLinkDt {
	if o == nil {
		var ret []HyperflexHxLinkDt
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxUuIdDtAllOf) GetLinksOk() ([]HyperflexHxLinkDt, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HyperflexHxUuIdDtAllOf) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []HyperflexHxLinkDt and assigns it to the Links field.
func (o *HyperflexHxUuIdDtAllOf) SetLinks(v []HyperflexHxLinkDt) {
	o.Links = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexHxUuIdDtAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxUuIdDtAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexHxUuIdDtAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexHxUuIdDtAllOf) SetUuid(v string) {
	o.Uuid = &v
}

func (o HyperflexHxUuIdDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Links != nil {
		toSerialize["Links"] = o.Links
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxUuIdDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxUuIdDtAllOf := _HyperflexHxUuIdDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxUuIdDtAllOf); err == nil {
		*o = HyperflexHxUuIdDtAllOf(varHyperflexHxUuIdDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Links")
		delete(additionalProperties, "Uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxUuIdDtAllOf struct {
	value *HyperflexHxUuIdDtAllOf
	isSet bool
}

func (v NullableHyperflexHxUuIdDtAllOf) Get() *HyperflexHxUuIdDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxUuIdDtAllOf) Set(val *HyperflexHxUuIdDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxUuIdDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxUuIdDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxUuIdDtAllOf(val *HyperflexHxUuIdDtAllOf) *NullableHyperflexHxUuIdDtAllOf {
	return &NullableHyperflexHxUuIdDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxUuIdDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxUuIdDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
