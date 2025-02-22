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

// WorkflowWorkflowCtxAllOf Definition of the list of properties defined in 'workflow.WorkflowCtx', excluding properties defined in parent classes.
type WorkflowWorkflowCtxAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                           `json:"ObjectType"`
	InitiatorCtx  NullableWorkflowInitiatorContext `json:"InitiatorCtx,omitempty"`
	TargetCtxList []WorkflowTargetContext          `json:"TargetCtxList,omitempty"`
	// The name of workflowMeta of the workflow running.
	WorkflowMetaName *string `json:"WorkflowMetaName,omitempty"`
	// The subtype of the workflow.
	WorkflowSubtype *string `json:"WorkflowSubtype,omitempty"`
	// Type of the workflow being started. This can be any string for client services to distinguish workflow by type.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowCtxAllOf WorkflowWorkflowCtxAllOf

// NewWorkflowWorkflowCtxAllOf instantiates a new WorkflowWorkflowCtxAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowCtxAllOf(classId string, objectType string) *WorkflowWorkflowCtxAllOf {
	this := WorkflowWorkflowCtxAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowWorkflowCtxAllOfWithDefaults instantiates a new WorkflowWorkflowCtxAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowCtxAllOfWithDefaults() *WorkflowWorkflowCtxAllOf {
	this := WorkflowWorkflowCtxAllOf{}
	var classId string = "workflow.WorkflowCtx"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowCtx"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowCtxAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtxAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowCtxAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowCtxAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtxAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowCtxAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInitiatorCtx returns the InitiatorCtx field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowCtxAllOf) GetInitiatorCtx() WorkflowInitiatorContext {
	if o == nil || o.InitiatorCtx.Get() == nil {
		var ret WorkflowInitiatorContext
		return ret
	}
	return *o.InitiatorCtx.Get()
}

// GetInitiatorCtxOk returns a tuple with the InitiatorCtx field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowCtxAllOf) GetInitiatorCtxOk() (*WorkflowInitiatorContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatorCtx.Get(), o.InitiatorCtx.IsSet()
}

// HasInitiatorCtx returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtxAllOf) HasInitiatorCtx() bool {
	if o != nil && o.InitiatorCtx.IsSet() {
		return true
	}

	return false
}

// SetInitiatorCtx gets a reference to the given NullableWorkflowInitiatorContext and assigns it to the InitiatorCtx field.
func (o *WorkflowWorkflowCtxAllOf) SetInitiatorCtx(v WorkflowInitiatorContext) {
	o.InitiatorCtx.Set(&v)
}

// SetInitiatorCtxNil sets the value for InitiatorCtx to be an explicit nil
func (o *WorkflowWorkflowCtxAllOf) SetInitiatorCtxNil() {
	o.InitiatorCtx.Set(nil)
}

// UnsetInitiatorCtx ensures that no value is present for InitiatorCtx, not even an explicit nil
func (o *WorkflowWorkflowCtxAllOf) UnsetInitiatorCtx() {
	o.InitiatorCtx.Unset()
}

// GetTargetCtxList returns the TargetCtxList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowCtxAllOf) GetTargetCtxList() []WorkflowTargetContext {
	if o == nil {
		var ret []WorkflowTargetContext
		return ret
	}
	return o.TargetCtxList
}

// GetTargetCtxListOk returns a tuple with the TargetCtxList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowCtxAllOf) GetTargetCtxListOk() ([]WorkflowTargetContext, bool) {
	if o == nil || o.TargetCtxList == nil {
		return nil, false
	}
	return o.TargetCtxList, true
}

// HasTargetCtxList returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtxAllOf) HasTargetCtxList() bool {
	if o != nil && o.TargetCtxList != nil {
		return true
	}

	return false
}

// SetTargetCtxList gets a reference to the given []WorkflowTargetContext and assigns it to the TargetCtxList field.
func (o *WorkflowWorkflowCtxAllOf) SetTargetCtxList(v []WorkflowTargetContext) {
	o.TargetCtxList = v
}

// GetWorkflowMetaName returns the WorkflowMetaName field value if set, zero value otherwise.
func (o *WorkflowWorkflowCtxAllOf) GetWorkflowMetaName() string {
	if o == nil || o.WorkflowMetaName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowMetaName
}

// GetWorkflowMetaNameOk returns a tuple with the WorkflowMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtxAllOf) GetWorkflowMetaNameOk() (*string, bool) {
	if o == nil || o.WorkflowMetaName == nil {
		return nil, false
	}
	return o.WorkflowMetaName, true
}

// HasWorkflowMetaName returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtxAllOf) HasWorkflowMetaName() bool {
	if o != nil && o.WorkflowMetaName != nil {
		return true
	}

	return false
}

// SetWorkflowMetaName gets a reference to the given string and assigns it to the WorkflowMetaName field.
func (o *WorkflowWorkflowCtxAllOf) SetWorkflowMetaName(v string) {
	o.WorkflowMetaName = &v
}

// GetWorkflowSubtype returns the WorkflowSubtype field value if set, zero value otherwise.
func (o *WorkflowWorkflowCtxAllOf) GetWorkflowSubtype() string {
	if o == nil || o.WorkflowSubtype == nil {
		var ret string
		return ret
	}
	return *o.WorkflowSubtype
}

// GetWorkflowSubtypeOk returns a tuple with the WorkflowSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtxAllOf) GetWorkflowSubtypeOk() (*string, bool) {
	if o == nil || o.WorkflowSubtype == nil {
		return nil, false
	}
	return o.WorkflowSubtype, true
}

// HasWorkflowSubtype returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtxAllOf) HasWorkflowSubtype() bool {
	if o != nil && o.WorkflowSubtype != nil {
		return true
	}

	return false
}

// SetWorkflowSubtype gets a reference to the given string and assigns it to the WorkflowSubtype field.
func (o *WorkflowWorkflowCtxAllOf) SetWorkflowSubtype(v string) {
	o.WorkflowSubtype = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *WorkflowWorkflowCtxAllOf) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtxAllOf) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtxAllOf) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *WorkflowWorkflowCtxAllOf) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o WorkflowWorkflowCtxAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InitiatorCtx.IsSet() {
		toSerialize["InitiatorCtx"] = o.InitiatorCtx.Get()
	}
	if o.TargetCtxList != nil {
		toSerialize["TargetCtxList"] = o.TargetCtxList
	}
	if o.WorkflowMetaName != nil {
		toSerialize["WorkflowMetaName"] = o.WorkflowMetaName
	}
	if o.WorkflowSubtype != nil {
		toSerialize["WorkflowSubtype"] = o.WorkflowSubtype
	}
	if o.WorkflowType != nil {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowCtxAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWorkflowCtxAllOf := _WorkflowWorkflowCtxAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWorkflowCtxAllOf); err == nil {
		*o = WorkflowWorkflowCtxAllOf(varWorkflowWorkflowCtxAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitiatorCtx")
		delete(additionalProperties, "TargetCtxList")
		delete(additionalProperties, "WorkflowMetaName")
		delete(additionalProperties, "WorkflowSubtype")
		delete(additionalProperties, "WorkflowType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWorkflowCtxAllOf struct {
	value *WorkflowWorkflowCtxAllOf
	isSet bool
}

func (v NullableWorkflowWorkflowCtxAllOf) Get() *WorkflowWorkflowCtxAllOf {
	return v.value
}

func (v *NullableWorkflowWorkflowCtxAllOf) Set(val *WorkflowWorkflowCtxAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowCtxAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowCtxAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowCtxAllOf(val *WorkflowWorkflowCtxAllOf) *NullableWorkflowWorkflowCtxAllOf {
	return &NullableWorkflowWorkflowCtxAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowCtxAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowCtxAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
