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

// TelemetryDruidHavingNumericFilterAllOf struct for TelemetryDruidHavingNumericFilterAllOf
type TelemetryDruidHavingNumericFilterAllOf struct {
	// aggregate metric
	Aggregation          string  `json:"aggregation"`
	Value                float64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidHavingNumericFilterAllOf TelemetryDruidHavingNumericFilterAllOf

// NewTelemetryDruidHavingNumericFilterAllOf instantiates a new TelemetryDruidHavingNumericFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidHavingNumericFilterAllOf(aggregation string, value float64) *TelemetryDruidHavingNumericFilterAllOf {
	this := TelemetryDruidHavingNumericFilterAllOf{}
	this.Aggregation = aggregation
	this.Value = value
	return &this
}

// NewTelemetryDruidHavingNumericFilterAllOfWithDefaults instantiates a new TelemetryDruidHavingNumericFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidHavingNumericFilterAllOfWithDefaults() *TelemetryDruidHavingNumericFilterAllOf {
	this := TelemetryDruidHavingNumericFilterAllOf{}
	return &this
}

// GetAggregation returns the Aggregation field value
func (o *TelemetryDruidHavingNumericFilterAllOf) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingNumericFilterAllOf) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *TelemetryDruidHavingNumericFilterAllOf) SetAggregation(v string) {
	o.Aggregation = v
}

// GetValue returns the Value field value
func (o *TelemetryDruidHavingNumericFilterAllOf) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingNumericFilterAllOf) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TelemetryDruidHavingNumericFilterAllOf) SetValue(v float64) {
	o.Value = v
}

func (o TelemetryDruidHavingNumericFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aggregation"] = o.Aggregation
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidHavingNumericFilterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidHavingNumericFilterAllOf := _TelemetryDruidHavingNumericFilterAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidHavingNumericFilterAllOf); err == nil {
		*o = TelemetryDruidHavingNumericFilterAllOf(varTelemetryDruidHavingNumericFilterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "aggregation")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidHavingNumericFilterAllOf struct {
	value *TelemetryDruidHavingNumericFilterAllOf
	isSet bool
}

func (v NullableTelemetryDruidHavingNumericFilterAllOf) Get() *TelemetryDruidHavingNumericFilterAllOf {
	return v.value
}

func (v *NullableTelemetryDruidHavingNumericFilterAllOf) Set(val *TelemetryDruidHavingNumericFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidHavingNumericFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidHavingNumericFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidHavingNumericFilterAllOf(val *TelemetryDruidHavingNumericFilterAllOf) *NullableTelemetryDruidHavingNumericFilterAllOf {
	return &NullableTelemetryDruidHavingNumericFilterAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidHavingNumericFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidHavingNumericFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
