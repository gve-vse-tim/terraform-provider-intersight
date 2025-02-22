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

// TelemetryDruidContainsSearchSpec The specification for a Druid 'insensitive_contains' search
type TelemetryDruidContainsSearchSpec struct {
	Type string `json:"type"`
	// The value to match.  If any part of a dimension value contains the value specified in this search query spec, regardless of case, a \"match\" occurs.
	Value string `json:"value"`
	// Whether or not search is case sensitive
	CaseSensitive        *bool `json:"case_sensitive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidContainsSearchSpec TelemetryDruidContainsSearchSpec

// NewTelemetryDruidContainsSearchSpec instantiates a new TelemetryDruidContainsSearchSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidContainsSearchSpec(type_ string, value string) *TelemetryDruidContainsSearchSpec {
	this := TelemetryDruidContainsSearchSpec{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewTelemetryDruidContainsSearchSpecWithDefaults instantiates a new TelemetryDruidContainsSearchSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidContainsSearchSpecWithDefaults() *TelemetryDruidContainsSearchSpec {
	this := TelemetryDruidContainsSearchSpec{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidContainsSearchSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidContainsSearchSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidContainsSearchSpec) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *TelemetryDruidContainsSearchSpec) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidContainsSearchSpec) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TelemetryDruidContainsSearchSpec) SetValue(v string) {
	o.Value = v
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *TelemetryDruidContainsSearchSpec) GetCaseSensitive() bool {
	if o == nil || o.CaseSensitive == nil {
		var ret bool
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidContainsSearchSpec) GetCaseSensitiveOk() (*bool, bool) {
	if o == nil || o.CaseSensitive == nil {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *TelemetryDruidContainsSearchSpec) HasCaseSensitive() bool {
	if o != nil && o.CaseSensitive != nil {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given bool and assigns it to the CaseSensitive field.
func (o *TelemetryDruidContainsSearchSpec) SetCaseSensitive(v bool) {
	o.CaseSensitive = &v
}

func (o TelemetryDruidContainsSearchSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.CaseSensitive != nil {
		toSerialize["case_sensitive"] = o.CaseSensitive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidContainsSearchSpec) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidContainsSearchSpec := _TelemetryDruidContainsSearchSpec{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidContainsSearchSpec); err == nil {
		*o = TelemetryDruidContainsSearchSpec(varTelemetryDruidContainsSearchSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "case_sensitive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidContainsSearchSpec struct {
	value *TelemetryDruidContainsSearchSpec
	isSet bool
}

func (v NullableTelemetryDruidContainsSearchSpec) Get() *TelemetryDruidContainsSearchSpec {
	return v.value
}

func (v *NullableTelemetryDruidContainsSearchSpec) Set(val *TelemetryDruidContainsSearchSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidContainsSearchSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidContainsSearchSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidContainsSearchSpec(val *TelemetryDruidContainsSearchSpec) *NullableTelemetryDruidContainsSearchSpec {
	return &NullableTelemetryDruidContainsSearchSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidContainsSearchSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidContainsSearchSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
