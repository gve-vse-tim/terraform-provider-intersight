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

// FabricSystemQosPolicy Configuration object sent by user to setup Quality of Service (QoS) for this switch.
type FabricSystemQosPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                                `json:"ObjectType"`
	Classes      []FabricQosClass                      `json:"Classes,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	Profiles             []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSystemQosPolicy FabricSystemQosPolicy

// NewFabricSystemQosPolicy instantiates a new FabricSystemQosPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSystemQosPolicy(classId string, objectType string) *FabricSystemQosPolicy {
	this := FabricSystemQosPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricSystemQosPolicyWithDefaults instantiates a new FabricSystemQosPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSystemQosPolicyWithDefaults() *FabricSystemQosPolicy {
	this := FabricSystemQosPolicy{}
	var classId string = "fabric.SystemQosPolicy"
	this.ClassId = classId
	var objectType string = "fabric.SystemQosPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSystemQosPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSystemQosPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSystemQosPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSystemQosPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSystemQosPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSystemQosPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClasses returns the Classes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSystemQosPolicy) GetClasses() []FabricQosClass {
	if o == nil {
		var ret []FabricQosClass
		return ret
	}
	return o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSystemQosPolicy) GetClassesOk() ([]FabricQosClass, bool) {
	if o == nil || o.Classes == nil {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *FabricSystemQosPolicy) HasClasses() bool {
	if o != nil && o.Classes != nil {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []FabricQosClass and assigns it to the Classes field.
func (o *FabricSystemQosPolicy) SetClasses(v []FabricQosClass) {
	o.Classes = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricSystemQosPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSystemQosPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricSystemQosPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricSystemQosPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSystemQosPolicy) GetProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSystemQosPolicy) GetProfilesOk() ([]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricSystemQosPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricSystemQosPolicy) SetProfiles(v []FabricSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricSystemQosPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Classes != nil {
		toSerialize["Classes"] = o.Classes
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricSystemQosPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type FabricSystemQosPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                `json:"ObjectType"`
		Classes      []FabricQosClass                      `json:"Classes,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fabricSwitchProfile resources.
		Profiles []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	}

	varFabricSystemQosPolicyWithoutEmbeddedStruct := FabricSystemQosPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricSystemQosPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricSystemQosPolicy := _FabricSystemQosPolicy{}
		varFabricSystemQosPolicy.ClassId = varFabricSystemQosPolicyWithoutEmbeddedStruct.ClassId
		varFabricSystemQosPolicy.ObjectType = varFabricSystemQosPolicyWithoutEmbeddedStruct.ObjectType
		varFabricSystemQosPolicy.Classes = varFabricSystemQosPolicyWithoutEmbeddedStruct.Classes
		varFabricSystemQosPolicy.Organization = varFabricSystemQosPolicyWithoutEmbeddedStruct.Organization
		varFabricSystemQosPolicy.Profiles = varFabricSystemQosPolicyWithoutEmbeddedStruct.Profiles
		*o = FabricSystemQosPolicy(varFabricSystemQosPolicy)
	} else {
		return err
	}

	varFabricSystemQosPolicy := _FabricSystemQosPolicy{}

	err = json.Unmarshal(bytes, &varFabricSystemQosPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricSystemQosPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Classes")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableFabricSystemQosPolicy struct {
	value *FabricSystemQosPolicy
	isSet bool
}

func (v NullableFabricSystemQosPolicy) Get() *FabricSystemQosPolicy {
	return v.value
}

func (v *NullableFabricSystemQosPolicy) Set(val *FabricSystemQosPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSystemQosPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSystemQosPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSystemQosPolicy(val *FabricSystemQosPolicy) *NullableFabricSystemQosPolicy {
	return &NullableFabricSystemQosPolicy{value: val, isSet: true}
}

func (v NullableFabricSystemQosPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSystemQosPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
