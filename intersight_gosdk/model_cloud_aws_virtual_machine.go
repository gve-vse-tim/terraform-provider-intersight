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

// CloudAwsVirtualMachine Configuration details of the virtual machine.
type CloudAwsVirtualMachine struct {
	CloudBaseVirtualMachine
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string                           `json:"ObjectType"`
	AwsBillingUnit *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	KeyPair        *CloudAwsKeyPairRelationship     `json:"KeyPair,omitempty"`
	Location       *CloudAwsVpcRelationship         `json:"Location,omitempty"`
	// An array of relationships to cloudAwsSecurityGroup resources.
	SecurityGroups       []CloudAwsSecurityGroupRelationship `json:"SecurityGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsVirtualMachine CloudAwsVirtualMachine

// NewCloudAwsVirtualMachine instantiates a new CloudAwsVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsVirtualMachine(classId string, objectType string) *CloudAwsVirtualMachine {
	this := CloudAwsVirtualMachine{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var powerState string = "Unknown"
	this.PowerState = &powerState
	var provider string = "Unknown"
	this.Provider = &provider
	var state string = "None"
	this.State = &state
	return &this
}

// NewCloudAwsVirtualMachineWithDefaults instantiates a new CloudAwsVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsVirtualMachineWithDefaults() *CloudAwsVirtualMachine {
	this := CloudAwsVirtualMachine{}
	var classId string = "cloud.AwsVirtualMachine"
	this.ClassId = classId
	var objectType string = "cloud.AwsVirtualMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsVirtualMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsVirtualMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsVirtualMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsVirtualMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAwsBillingUnit returns the AwsBillingUnit field value if set, zero value otherwise.
func (o *CloudAwsVirtualMachine) GetAwsBillingUnit() CloudAwsBillingUnitRelationship {
	if o == nil || o.AwsBillingUnit == nil {
		var ret CloudAwsBillingUnitRelationship
		return ret
	}
	return *o.AwsBillingUnit
}

// GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachine) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool) {
	if o == nil || o.AwsBillingUnit == nil {
		return nil, false
	}
	return o.AwsBillingUnit, true
}

// HasAwsBillingUnit returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachine) HasAwsBillingUnit() bool {
	if o != nil && o.AwsBillingUnit != nil {
		return true
	}

	return false
}

// SetAwsBillingUnit gets a reference to the given CloudAwsBillingUnitRelationship and assigns it to the AwsBillingUnit field.
func (o *CloudAwsVirtualMachine) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship) {
	o.AwsBillingUnit = &v
}

// GetKeyPair returns the KeyPair field value if set, zero value otherwise.
func (o *CloudAwsVirtualMachine) GetKeyPair() CloudAwsKeyPairRelationship {
	if o == nil || o.KeyPair == nil {
		var ret CloudAwsKeyPairRelationship
		return ret
	}
	return *o.KeyPair
}

// GetKeyPairOk returns a tuple with the KeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachine) GetKeyPairOk() (*CloudAwsKeyPairRelationship, bool) {
	if o == nil || o.KeyPair == nil {
		return nil, false
	}
	return o.KeyPair, true
}

// HasKeyPair returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachine) HasKeyPair() bool {
	if o != nil && o.KeyPair != nil {
		return true
	}

	return false
}

// SetKeyPair gets a reference to the given CloudAwsKeyPairRelationship and assigns it to the KeyPair field.
func (o *CloudAwsVirtualMachine) SetKeyPair(v CloudAwsKeyPairRelationship) {
	o.KeyPair = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CloudAwsVirtualMachine) GetLocation() CloudAwsVpcRelationship {
	if o == nil || o.Location == nil {
		var ret CloudAwsVpcRelationship
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachine) GetLocationOk() (*CloudAwsVpcRelationship, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachine) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given CloudAwsVpcRelationship and assigns it to the Location field.
func (o *CloudAwsVirtualMachine) SetLocation(v CloudAwsVpcRelationship) {
	o.Location = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVirtualMachine) GetSecurityGroups() []CloudAwsSecurityGroupRelationship {
	if o == nil {
		var ret []CloudAwsSecurityGroupRelationship
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVirtualMachine) GetSecurityGroupsOk() ([]CloudAwsSecurityGroupRelationship, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachine) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []CloudAwsSecurityGroupRelationship and assigns it to the SecurityGroups field.
func (o *CloudAwsVirtualMachine) SetSecurityGroups(v []CloudAwsSecurityGroupRelationship) {
	o.SecurityGroups = v
}

func (o CloudAwsVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVirtualMachine, errCloudBaseVirtualMachine := json.Marshal(o.CloudBaseVirtualMachine)
	if errCloudBaseVirtualMachine != nil {
		return []byte{}, errCloudBaseVirtualMachine
	}
	errCloudBaseVirtualMachine = json.Unmarshal([]byte(serializedCloudBaseVirtualMachine), &toSerialize)
	if errCloudBaseVirtualMachine != nil {
		return []byte{}, errCloudBaseVirtualMachine
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AwsBillingUnit != nil {
		toSerialize["AwsBillingUnit"] = o.AwsBillingUnit
	}
	if o.KeyPair != nil {
		toSerialize["KeyPair"] = o.KeyPair
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsVirtualMachine) UnmarshalJSON(bytes []byte) (err error) {
	type CloudAwsVirtualMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string                           `json:"ObjectType"`
		AwsBillingUnit *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
		KeyPair        *CloudAwsKeyPairRelationship     `json:"KeyPair,omitempty"`
		Location       *CloudAwsVpcRelationship         `json:"Location,omitempty"`
		// An array of relationships to cloudAwsSecurityGroup resources.
		SecurityGroups []CloudAwsSecurityGroupRelationship `json:"SecurityGroups,omitempty"`
	}

	varCloudAwsVirtualMachineWithoutEmbeddedStruct := CloudAwsVirtualMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudAwsVirtualMachineWithoutEmbeddedStruct)
	if err == nil {
		varCloudAwsVirtualMachine := _CloudAwsVirtualMachine{}
		varCloudAwsVirtualMachine.ClassId = varCloudAwsVirtualMachineWithoutEmbeddedStruct.ClassId
		varCloudAwsVirtualMachine.ObjectType = varCloudAwsVirtualMachineWithoutEmbeddedStruct.ObjectType
		varCloudAwsVirtualMachine.AwsBillingUnit = varCloudAwsVirtualMachineWithoutEmbeddedStruct.AwsBillingUnit
		varCloudAwsVirtualMachine.KeyPair = varCloudAwsVirtualMachineWithoutEmbeddedStruct.KeyPair
		varCloudAwsVirtualMachine.Location = varCloudAwsVirtualMachineWithoutEmbeddedStruct.Location
		varCloudAwsVirtualMachine.SecurityGroups = varCloudAwsVirtualMachineWithoutEmbeddedStruct.SecurityGroups
		*o = CloudAwsVirtualMachine(varCloudAwsVirtualMachine)
	} else {
		return err
	}

	varCloudAwsVirtualMachine := _CloudAwsVirtualMachine{}

	err = json.Unmarshal(bytes, &varCloudAwsVirtualMachine)
	if err == nil {
		o.CloudBaseVirtualMachine = varCloudAwsVirtualMachine.CloudBaseVirtualMachine
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AwsBillingUnit")
		delete(additionalProperties, "KeyPair")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "SecurityGroups")

		// remove fields from embedded structs
		reflectCloudBaseVirtualMachine := reflect.ValueOf(o.CloudBaseVirtualMachine)
		for i := 0; i < reflectCloudBaseVirtualMachine.Type().NumField(); i++ {
			t := reflectCloudBaseVirtualMachine.Type().Field(i)

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

type NullableCloudAwsVirtualMachine struct {
	value *CloudAwsVirtualMachine
	isSet bool
}

func (v NullableCloudAwsVirtualMachine) Get() *CloudAwsVirtualMachine {
	return v.value
}

func (v *NullableCloudAwsVirtualMachine) Set(val *CloudAwsVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsVirtualMachine(val *CloudAwsVirtualMachine) *NullableCloudAwsVirtualMachine {
	return &NullableCloudAwsVirtualMachine{value: val, isSet: true}
}

func (v NullableCloudAwsVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
