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

// NotificationAbstractMoConditionAllOf Definition of the list of properties defined in 'notification.AbstractMoCondition', excluding properties defined in parent classes.
type NotificationAbstractMoConditionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The condition can be switched on/off with out necessity to change the subscription settings: actions, conditions, etc. Ex.: Subscription MO can be configured, but switched off.
	Enabled *bool `json:"Enabled,omitempty"`
	// MoType for which the condition is created.
	MoType               *string  `json:"MoType,omitempty"`
	Operations           []string `json:"Operations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationAbstractMoConditionAllOf NotificationAbstractMoConditionAllOf

// NewNotificationAbstractMoConditionAllOf instantiates a new NotificationAbstractMoConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAbstractMoConditionAllOf(classId string, objectType string) *NotificationAbstractMoConditionAllOf {
	this := NotificationAbstractMoConditionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationAbstractMoConditionAllOfWithDefaults instantiates a new NotificationAbstractMoConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAbstractMoConditionAllOfWithDefaults() *NotificationAbstractMoConditionAllOf {
	this := NotificationAbstractMoConditionAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationAbstractMoConditionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationAbstractMoConditionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationAbstractMoConditionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationAbstractMoConditionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationAbstractMoConditionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationAbstractMoConditionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NotificationAbstractMoConditionAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAbstractMoConditionAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NotificationAbstractMoConditionAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NotificationAbstractMoConditionAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *NotificationAbstractMoConditionAllOf) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAbstractMoConditionAllOf) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *NotificationAbstractMoConditionAllOf) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *NotificationAbstractMoConditionAllOf) SetMoType(v string) {
	o.MoType = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationAbstractMoConditionAllOf) GetOperations() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationAbstractMoConditionAllOf) GetOperationsOk() ([]string, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *NotificationAbstractMoConditionAllOf) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []string and assigns it to the Operations field.
func (o *NotificationAbstractMoConditionAllOf) SetOperations(v []string) {
	o.Operations = v
}

func (o NotificationAbstractMoConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.Operations != nil {
		toSerialize["Operations"] = o.Operations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationAbstractMoConditionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationAbstractMoConditionAllOf := _NotificationAbstractMoConditionAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationAbstractMoConditionAllOf); err == nil {
		*o = NotificationAbstractMoConditionAllOf(varNotificationAbstractMoConditionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "Operations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationAbstractMoConditionAllOf struct {
	value *NotificationAbstractMoConditionAllOf
	isSet bool
}

func (v NullableNotificationAbstractMoConditionAllOf) Get() *NotificationAbstractMoConditionAllOf {
	return v.value
}

func (v *NullableNotificationAbstractMoConditionAllOf) Set(val *NotificationAbstractMoConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAbstractMoConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAbstractMoConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAbstractMoConditionAllOf(val *NotificationAbstractMoConditionAllOf) *NullableNotificationAbstractMoConditionAllOf {
	return &NullableNotificationAbstractMoConditionAllOf{value: val, isSet: true}
}

func (v NullableNotificationAbstractMoConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAbstractMoConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
