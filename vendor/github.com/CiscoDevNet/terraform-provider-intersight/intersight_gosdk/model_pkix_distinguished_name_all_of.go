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

// PkixDistinguishedNameAllOf Definition of the list of properties defined in 'pkix.DistinguishedName', excluding properties defined in parent classes.
type PkixDistinguishedNameAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A required component that identifies a person or an object.
	CommonName           *string  `json:"CommonName,omitempty"`
	Country              []string `json:"Country,omitempty"`
	Locality             []string `json:"Locality,omitempty"`
	Organization         []string `json:"Organization,omitempty"`
	OrganizationalUnit   []string `json:"OrganizationalUnit,omitempty"`
	State                []string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixDistinguishedNameAllOf PkixDistinguishedNameAllOf

// NewPkixDistinguishedNameAllOf instantiates a new PkixDistinguishedNameAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixDistinguishedNameAllOf(classId string, objectType string) *PkixDistinguishedNameAllOf {
	this := PkixDistinguishedNameAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPkixDistinguishedNameAllOfWithDefaults instantiates a new PkixDistinguishedNameAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixDistinguishedNameAllOfWithDefaults() *PkixDistinguishedNameAllOf {
	this := PkixDistinguishedNameAllOf{}
	var classId string = "pkix.DistinguishedName"
	this.ClassId = classId
	var objectType string = "pkix.DistinguishedName"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixDistinguishedNameAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedNameAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixDistinguishedNameAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PkixDistinguishedNameAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedNameAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixDistinguishedNameAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *PkixDistinguishedNameAllOf) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedNameAllOf) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *PkixDistinguishedNameAllOf) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *PkixDistinguishedNameAllOf) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedNameAllOf) GetCountry() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedNameAllOf) GetCountryOk() ([]string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PkixDistinguishedNameAllOf) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []string and assigns it to the Country field.
func (o *PkixDistinguishedNameAllOf) SetCountry(v []string) {
	o.Country = v
}

// GetLocality returns the Locality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedNameAllOf) GetLocality() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedNameAllOf) GetLocalityOk() ([]string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *PkixDistinguishedNameAllOf) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given []string and assigns it to the Locality field.
func (o *PkixDistinguishedNameAllOf) SetLocality(v []string) {
	o.Locality = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedNameAllOf) GetOrganization() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedNameAllOf) GetOrganizationOk() ([]string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PkixDistinguishedNameAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given []string and assigns it to the Organization field.
func (o *PkixDistinguishedNameAllOf) SetOrganization(v []string) {
	o.Organization = v
}

// GetOrganizationalUnit returns the OrganizationalUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedNameAllOf) GetOrganizationalUnit() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OrganizationalUnit
}

// GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedNameAllOf) GetOrganizationalUnitOk() ([]string, bool) {
	if o == nil || o.OrganizationalUnit == nil {
		return nil, false
	}
	return o.OrganizationalUnit, true
}

// HasOrganizationalUnit returns a boolean if a field has been set.
func (o *PkixDistinguishedNameAllOf) HasOrganizationalUnit() bool {
	if o != nil && o.OrganizationalUnit != nil {
		return true
	}

	return false
}

// SetOrganizationalUnit gets a reference to the given []string and assigns it to the OrganizationalUnit field.
func (o *PkixDistinguishedNameAllOf) SetOrganizationalUnit(v []string) {
	o.OrganizationalUnit = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedNameAllOf) GetState() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedNameAllOf) GetStateOk() ([]string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PkixDistinguishedNameAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given []string and assigns it to the State field.
func (o *PkixDistinguishedNameAllOf) SetState(v []string) {
	o.State = v
}

func (o PkixDistinguishedNameAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CommonName != nil {
		toSerialize["CommonName"] = o.CommonName
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.Locality != nil {
		toSerialize["Locality"] = o.Locality
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OrganizationalUnit != nil {
		toSerialize["OrganizationalUnit"] = o.OrganizationalUnit
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixDistinguishedNameAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPkixDistinguishedNameAllOf := _PkixDistinguishedNameAllOf{}

	if err = json.Unmarshal(bytes, &varPkixDistinguishedNameAllOf); err == nil {
		*o = PkixDistinguishedNameAllOf(varPkixDistinguishedNameAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommonName")
		delete(additionalProperties, "Country")
		delete(additionalProperties, "Locality")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OrganizationalUnit")
		delete(additionalProperties, "State")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePkixDistinguishedNameAllOf struct {
	value *PkixDistinguishedNameAllOf
	isSet bool
}

func (v NullablePkixDistinguishedNameAllOf) Get() *PkixDistinguishedNameAllOf {
	return v.value
}

func (v *NullablePkixDistinguishedNameAllOf) Set(val *PkixDistinguishedNameAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixDistinguishedNameAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixDistinguishedNameAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixDistinguishedNameAllOf(val *PkixDistinguishedNameAllOf) *NullablePkixDistinguishedNameAllOf {
	return &NullablePkixDistinguishedNameAllOf{value: val, isSet: true}
}

func (v NullablePkixDistinguishedNameAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixDistinguishedNameAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
