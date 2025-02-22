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

// WorkflowTemplateFunctionMetaAllOf Definition of the list of properties defined in 'workflow.TemplateFunctionMeta', excluding properties defined in parent classes.
type WorkflowTemplateFunctionMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                   `json:"ObjectType"`
	Comments   NullableWorkflowComments `json:"Comments,omitempty"`
	Inputs     []WorkflowBaseDataType   `json:"Inputs,omitempty"`
	// The flag indicates whether a guided mode template is supported for it or not.
	IsGuidedModeSupported *bool `json:"IsGuidedModeSupported,omitempty"`
	// The template function name.
	Name                 *string                `json:"Name,omitempty"`
	Outputs              []WorkflowBaseDataType `json:"Outputs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTemplateFunctionMetaAllOf WorkflowTemplateFunctionMetaAllOf

// NewWorkflowTemplateFunctionMetaAllOf instantiates a new WorkflowTemplateFunctionMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTemplateFunctionMetaAllOf(classId string, objectType string) *WorkflowTemplateFunctionMetaAllOf {
	this := WorkflowTemplateFunctionMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTemplateFunctionMetaAllOfWithDefaults instantiates a new WorkflowTemplateFunctionMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTemplateFunctionMetaAllOfWithDefaults() *WorkflowTemplateFunctionMetaAllOf {
	this := WorkflowTemplateFunctionMetaAllOf{}
	var classId string = "workflow.TemplateFunctionMeta"
	this.ClassId = classId
	var objectType string = "workflow.TemplateFunctionMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTemplateFunctionMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTemplateFunctionMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTemplateFunctionMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTemplateFunctionMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateFunctionMetaAllOf) GetComments() WorkflowComments {
	if o == nil || o.Comments.Get() == nil {
		var ret WorkflowComments
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateFunctionMetaAllOf) GetCommentsOk() (*WorkflowComments, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableWorkflowComments and assigns it to the Comments field.
func (o *WorkflowTemplateFunctionMetaAllOf) SetComments(v WorkflowComments) {
	o.Comments.Set(&v)
}

// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *WorkflowTemplateFunctionMetaAllOf) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *WorkflowTemplateFunctionMetaAllOf) UnsetComments() {
	o.Comments.Unset()
}

// GetInputs returns the Inputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateFunctionMetaAllOf) GetInputs() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateFunctionMetaAllOf) GetInputsOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []WorkflowBaseDataType and assigns it to the Inputs field.
func (o *WorkflowTemplateFunctionMetaAllOf) SetInputs(v []WorkflowBaseDataType) {
	o.Inputs = v
}

// GetIsGuidedModeSupported returns the IsGuidedModeSupported field value if set, zero value otherwise.
func (o *WorkflowTemplateFunctionMetaAllOf) GetIsGuidedModeSupported() bool {
	if o == nil || o.IsGuidedModeSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsGuidedModeSupported
}

// GetIsGuidedModeSupportedOk returns a tuple with the IsGuidedModeSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) GetIsGuidedModeSupportedOk() (*bool, bool) {
	if o == nil || o.IsGuidedModeSupported == nil {
		return nil, false
	}
	return o.IsGuidedModeSupported, true
}

// HasIsGuidedModeSupported returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) HasIsGuidedModeSupported() bool {
	if o != nil && o.IsGuidedModeSupported != nil {
		return true
	}

	return false
}

// SetIsGuidedModeSupported gets a reference to the given bool and assigns it to the IsGuidedModeSupported field.
func (o *WorkflowTemplateFunctionMetaAllOf) SetIsGuidedModeSupported(v bool) {
	o.IsGuidedModeSupported = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTemplateFunctionMetaAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTemplateFunctionMetaAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateFunctionMetaAllOf) GetOutputs() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateFunctionMetaAllOf) GetOutputsOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMetaAllOf) HasOutputs() bool {
	if o != nil && o.Outputs != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []WorkflowBaseDataType and assigns it to the Outputs field.
func (o *WorkflowTemplateFunctionMetaAllOf) SetOutputs(v []WorkflowBaseDataType) {
	o.Outputs = v
}

func (o WorkflowTemplateFunctionMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Comments.IsSet() {
		toSerialize["Comments"] = o.Comments.Get()
	}
	if o.Inputs != nil {
		toSerialize["Inputs"] = o.Inputs
	}
	if o.IsGuidedModeSupported != nil {
		toSerialize["IsGuidedModeSupported"] = o.IsGuidedModeSupported
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Outputs != nil {
		toSerialize["Outputs"] = o.Outputs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTemplateFunctionMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTemplateFunctionMetaAllOf := _WorkflowTemplateFunctionMetaAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTemplateFunctionMetaAllOf); err == nil {
		*o = WorkflowTemplateFunctionMetaAllOf(varWorkflowTemplateFunctionMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Comments")
		delete(additionalProperties, "Inputs")
		delete(additionalProperties, "IsGuidedModeSupported")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Outputs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTemplateFunctionMetaAllOf struct {
	value *WorkflowTemplateFunctionMetaAllOf
	isSet bool
}

func (v NullableWorkflowTemplateFunctionMetaAllOf) Get() *WorkflowTemplateFunctionMetaAllOf {
	return v.value
}

func (v *NullableWorkflowTemplateFunctionMetaAllOf) Set(val *WorkflowTemplateFunctionMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTemplateFunctionMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTemplateFunctionMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTemplateFunctionMetaAllOf(val *WorkflowTemplateFunctionMetaAllOf) *NullableWorkflowTemplateFunctionMetaAllOf {
	return &NullableWorkflowTemplateFunctionMetaAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTemplateFunctionMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTemplateFunctionMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
