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

// ForecastModelAllOf Definition of the list of properties defined in 'forecast.Model', excluding properties defined in parent classes.
type ForecastModelAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The standard error of the estimate is a measure of the accuracy of predictions from predective modeling.
	Accuracy  *float32  `json:"Accuracy,omitempty"`
	ModelData []float32 `json:"ModelData,omitempty"`
	// Model type indicating type of predictive model used for computing forecast. * `Linear` - The Enum value Linear represents that the predictive model type used for forecast computation is linear regression.
	ModelType            *string `json:"ModelType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForecastModelAllOf ForecastModelAllOf

// NewForecastModelAllOf instantiates a new ForecastModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastModelAllOf(classId string, objectType string) *ForecastModelAllOf {
	this := ForecastModelAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var modelType string = "Linear"
	this.ModelType = &modelType
	return &this
}

// NewForecastModelAllOfWithDefaults instantiates a new ForecastModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastModelAllOfWithDefaults() *ForecastModelAllOf {
	this := ForecastModelAllOf{}
	var classId string = "forecast.Model"
	this.ClassId = classId
	var objectType string = "forecast.Model"
	this.ObjectType = objectType
	var modelType string = "Linear"
	this.ModelType = &modelType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ForecastModelAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ForecastModelAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ForecastModelAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ForecastModelAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ForecastModelAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ForecastModelAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *ForecastModelAllOf) GetAccuracy() float32 {
	if o == nil || o.Accuracy == nil {
		var ret float32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastModelAllOf) GetAccuracyOk() (*float32, bool) {
	if o == nil || o.Accuracy == nil {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *ForecastModelAllOf) HasAccuracy() bool {
	if o != nil && o.Accuracy != nil {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given float32 and assigns it to the Accuracy field.
func (o *ForecastModelAllOf) SetAccuracy(v float32) {
	o.Accuracy = &v
}

// GetModelData returns the ModelData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastModelAllOf) GetModelData() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}
	return o.ModelData
}

// GetModelDataOk returns a tuple with the ModelData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastModelAllOf) GetModelDataOk() ([]float32, bool) {
	if o == nil || o.ModelData == nil {
		return nil, false
	}
	return o.ModelData, true
}

// HasModelData returns a boolean if a field has been set.
func (o *ForecastModelAllOf) HasModelData() bool {
	if o != nil && o.ModelData != nil {
		return true
	}

	return false
}

// SetModelData gets a reference to the given []float32 and assigns it to the ModelData field.
func (o *ForecastModelAllOf) SetModelData(v []float32) {
	o.ModelData = v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *ForecastModelAllOf) GetModelType() string {
	if o == nil || o.ModelType == nil {
		var ret string
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastModelAllOf) GetModelTypeOk() (*string, bool) {
	if o == nil || o.ModelType == nil {
		return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *ForecastModelAllOf) HasModelType() bool {
	if o != nil && o.ModelType != nil {
		return true
	}

	return false
}

// SetModelType gets a reference to the given string and assigns it to the ModelType field.
func (o *ForecastModelAllOf) SetModelType(v string) {
	o.ModelType = &v
}

func (o ForecastModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Accuracy != nil {
		toSerialize["Accuracy"] = o.Accuracy
	}
	if o.ModelData != nil {
		toSerialize["ModelData"] = o.ModelData
	}
	if o.ModelType != nil {
		toSerialize["ModelType"] = o.ModelType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ForecastModelAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varForecastModelAllOf := _ForecastModelAllOf{}

	if err = json.Unmarshal(bytes, &varForecastModelAllOf); err == nil {
		*o = ForecastModelAllOf(varForecastModelAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Accuracy")
		delete(additionalProperties, "ModelData")
		delete(additionalProperties, "ModelType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForecastModelAllOf struct {
	value *ForecastModelAllOf
	isSet bool
}

func (v NullableForecastModelAllOf) Get() *ForecastModelAllOf {
	return v.value
}

func (v *NullableForecastModelAllOf) Set(val *ForecastModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastModelAllOf(val *ForecastModelAllOf) *NullableForecastModelAllOf {
	return &NullableForecastModelAllOf{value: val, isSet: true}
}

func (v NullableForecastModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
