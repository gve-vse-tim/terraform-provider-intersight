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

// AccessPolicyInventory Policy to configure server or chassis management options.
type AccessPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                          `json:"ObjectType"`
	AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
	ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
	// VLAN to be used for server access over Inband network.
	InbandVlan           *int64                  `json:"InbandVlan,omitempty"`
	InbandIpPool         *IppoolPoolRelationship `json:"InbandIpPool,omitempty"`
	InbandVrf            *VrfVrfRelationship     `json:"InbandVrf,omitempty"`
	OutOfBandIpPool      *IppoolPoolRelationship `json:"OutOfBandIpPool,omitempty"`
	OutOfBandVrf         *VrfVrfRelationship     `json:"OutOfBandVrf,omitempty"`
	TargetMo             *MoBaseMoRelationship   `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyInventory AccessPolicyInventory

// NewAccessPolicyInventory instantiates a new AccessPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyInventory(classId string, objectType string) *AccessPolicyInventory {
	this := AccessPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAccessPolicyInventoryWithDefaults instantiates a new AccessPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyInventoryWithDefaults() *AccessPolicyInventory {
	this := AccessPolicyInventory{}
	var classId string = "access.PolicyInventory"
	this.ClassId = classId
	var objectType string = "access.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AccessPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddressType returns the AddressType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetAddressType() AccessAddressType {
	if o == nil || o.AddressType.Get() == nil {
		var ret AccessAddressType
		return ret
	}
	return *o.AddressType.Get()
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetAddressTypeOk() (*AccessAddressType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressType.Get(), o.AddressType.IsSet()
}

// HasAddressType returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasAddressType() bool {
	if o != nil && o.AddressType.IsSet() {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given NullableAccessAddressType and assigns it to the AddressType field.
func (o *AccessPolicyInventory) SetAddressType(v AccessAddressType) {
	o.AddressType.Set(&v)
}

// SetAddressTypeNil sets the value for AddressType to be an explicit nil
func (o *AccessPolicyInventory) SetAddressTypeNil() {
	o.AddressType.Set(nil)
}

// UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
func (o *AccessPolicyInventory) UnsetAddressType() {
	o.AddressType.Unset()
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetConfigurationType() AccessConfigurationType {
	if o == nil || o.ConfigurationType.Get() == nil {
		var ret AccessConfigurationType
		return ret
	}
	return *o.ConfigurationType.Get()
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetConfigurationTypeOk() (*AccessConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationType.Get(), o.ConfigurationType.IsSet()
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasConfigurationType() bool {
	if o != nil && o.ConfigurationType.IsSet() {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given NullableAccessConfigurationType and assigns it to the ConfigurationType field.
func (o *AccessPolicyInventory) SetConfigurationType(v AccessConfigurationType) {
	o.ConfigurationType.Set(&v)
}

// SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil
func (o *AccessPolicyInventory) SetConfigurationTypeNil() {
	o.ConfigurationType.Set(nil)
}

// UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
func (o *AccessPolicyInventory) UnsetConfigurationType() {
	o.ConfigurationType.Unset()
}

// GetInbandVlan returns the InbandVlan field value if set, zero value otherwise.
func (o *AccessPolicyInventory) GetInbandVlan() int64 {
	if o == nil || o.InbandVlan == nil {
		var ret int64
		return ret
	}
	return *o.InbandVlan
}

// GetInbandVlanOk returns a tuple with the InbandVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetInbandVlanOk() (*int64, bool) {
	if o == nil || o.InbandVlan == nil {
		return nil, false
	}
	return o.InbandVlan, true
}

// HasInbandVlan returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasInbandVlan() bool {
	if o != nil && o.InbandVlan != nil {
		return true
	}

	return false
}

// SetInbandVlan gets a reference to the given int64 and assigns it to the InbandVlan field.
func (o *AccessPolicyInventory) SetInbandVlan(v int64) {
	o.InbandVlan = &v
}

// GetInbandIpPool returns the InbandIpPool field value if set, zero value otherwise.
func (o *AccessPolicyInventory) GetInbandIpPool() IppoolPoolRelationship {
	if o == nil || o.InbandIpPool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.InbandIpPool
}

// GetInbandIpPoolOk returns a tuple with the InbandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.InbandIpPool == nil {
		return nil, false
	}
	return o.InbandIpPool, true
}

// HasInbandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasInbandIpPool() bool {
	if o != nil && o.InbandIpPool != nil {
		return true
	}

	return false
}

// SetInbandIpPool gets a reference to the given IppoolPoolRelationship and assigns it to the InbandIpPool field.
func (o *AccessPolicyInventory) SetInbandIpPool(v IppoolPoolRelationship) {
	o.InbandIpPool = &v
}

// GetInbandVrf returns the InbandVrf field value if set, zero value otherwise.
func (o *AccessPolicyInventory) GetInbandVrf() VrfVrfRelationship {
	if o == nil || o.InbandVrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.InbandVrf
}

// GetInbandVrfOk returns a tuple with the InbandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetInbandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.InbandVrf == nil {
		return nil, false
	}
	return o.InbandVrf, true
}

// HasInbandVrf returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasInbandVrf() bool {
	if o != nil && o.InbandVrf != nil {
		return true
	}

	return false
}

// SetInbandVrf gets a reference to the given VrfVrfRelationship and assigns it to the InbandVrf field.
func (o *AccessPolicyInventory) SetInbandVrf(v VrfVrfRelationship) {
	o.InbandVrf = &v
}

// GetOutOfBandIpPool returns the OutOfBandIpPool field value if set, zero value otherwise.
func (o *AccessPolicyInventory) GetOutOfBandIpPool() IppoolPoolRelationship {
	if o == nil || o.OutOfBandIpPool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.OutOfBandIpPool
}

// GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.OutOfBandIpPool == nil {
		return nil, false
	}
	return o.OutOfBandIpPool, true
}

// HasOutOfBandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasOutOfBandIpPool() bool {
	if o != nil && o.OutOfBandIpPool != nil {
		return true
	}

	return false
}

// SetOutOfBandIpPool gets a reference to the given IppoolPoolRelationship and assigns it to the OutOfBandIpPool field.
func (o *AccessPolicyInventory) SetOutOfBandIpPool(v IppoolPoolRelationship) {
	o.OutOfBandIpPool = &v
}

// GetOutOfBandVrf returns the OutOfBandVrf field value if set, zero value otherwise.
func (o *AccessPolicyInventory) GetOutOfBandVrf() VrfVrfRelationship {
	if o == nil || o.OutOfBandVrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.OutOfBandVrf
}

// GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.OutOfBandVrf == nil {
		return nil, false
	}
	return o.OutOfBandVrf, true
}

// HasOutOfBandVrf returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasOutOfBandVrf() bool {
	if o != nil && o.OutOfBandVrf != nil {
		return true
	}

	return false
}

// SetOutOfBandVrf gets a reference to the given VrfVrfRelationship and assigns it to the OutOfBandVrf field.
func (o *AccessPolicyInventory) SetOutOfBandVrf(v VrfVrfRelationship) {
	o.OutOfBandVrf = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *AccessPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *AccessPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o AccessPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AddressType.IsSet() {
		toSerialize["AddressType"] = o.AddressType.Get()
	}
	if o.ConfigurationType.IsSet() {
		toSerialize["ConfigurationType"] = o.ConfigurationType.Get()
	}
	if o.InbandVlan != nil {
		toSerialize["InbandVlan"] = o.InbandVlan
	}
	if o.InbandIpPool != nil {
		toSerialize["InbandIpPool"] = o.InbandIpPool
	}
	if o.InbandVrf != nil {
		toSerialize["InbandVrf"] = o.InbandVrf
	}
	if o.OutOfBandIpPool != nil {
		toSerialize["OutOfBandIpPool"] = o.OutOfBandIpPool
	}
	if o.OutOfBandVrf != nil {
		toSerialize["OutOfBandVrf"] = o.OutOfBandVrf
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type AccessPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                          `json:"ObjectType"`
		AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
		ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
		// VLAN to be used for server access over Inband network.
		InbandVlan      *int64                  `json:"InbandVlan,omitempty"`
		InbandIpPool    *IppoolPoolRelationship `json:"InbandIpPool,omitempty"`
		InbandVrf       *VrfVrfRelationship     `json:"InbandVrf,omitempty"`
		OutOfBandIpPool *IppoolPoolRelationship `json:"OutOfBandIpPool,omitempty"`
		OutOfBandVrf    *VrfVrfRelationship     `json:"OutOfBandVrf,omitempty"`
		TargetMo        *MoBaseMoRelationship   `json:"TargetMo,omitempty"`
	}

	varAccessPolicyInventoryWithoutEmbeddedStruct := AccessPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAccessPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varAccessPolicyInventory := _AccessPolicyInventory{}
		varAccessPolicyInventory.ClassId = varAccessPolicyInventoryWithoutEmbeddedStruct.ClassId
		varAccessPolicyInventory.ObjectType = varAccessPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varAccessPolicyInventory.AddressType = varAccessPolicyInventoryWithoutEmbeddedStruct.AddressType
		varAccessPolicyInventory.ConfigurationType = varAccessPolicyInventoryWithoutEmbeddedStruct.ConfigurationType
		varAccessPolicyInventory.InbandVlan = varAccessPolicyInventoryWithoutEmbeddedStruct.InbandVlan
		varAccessPolicyInventory.InbandIpPool = varAccessPolicyInventoryWithoutEmbeddedStruct.InbandIpPool
		varAccessPolicyInventory.InbandVrf = varAccessPolicyInventoryWithoutEmbeddedStruct.InbandVrf
		varAccessPolicyInventory.OutOfBandIpPool = varAccessPolicyInventoryWithoutEmbeddedStruct.OutOfBandIpPool
		varAccessPolicyInventory.OutOfBandVrf = varAccessPolicyInventoryWithoutEmbeddedStruct.OutOfBandVrf
		varAccessPolicyInventory.TargetMo = varAccessPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = AccessPolicyInventory(varAccessPolicyInventory)
	} else {
		return err
	}

	varAccessPolicyInventory := _AccessPolicyInventory{}

	err = json.Unmarshal(bytes, &varAccessPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varAccessPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddressType")
		delete(additionalProperties, "ConfigurationType")
		delete(additionalProperties, "InbandVlan")
		delete(additionalProperties, "InbandIpPool")
		delete(additionalProperties, "InbandVrf")
		delete(additionalProperties, "OutOfBandIpPool")
		delete(additionalProperties, "OutOfBandVrf")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableAccessPolicyInventory struct {
	value *AccessPolicyInventory
	isSet bool
}

func (v NullableAccessPolicyInventory) Get() *AccessPolicyInventory {
	return v.value
}

func (v *NullableAccessPolicyInventory) Set(val *AccessPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyInventory(val *AccessPolicyInventory) *NullableAccessPolicyInventory {
	return &NullableAccessPolicyInventory{value: val, isSet: true}
}

func (v NullableAccessPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}