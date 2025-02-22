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

// WorkflowPropertiesAllOf Definition of the list of properties defined in 'workflow.Properties', excluding properties defined in parent classes.
type WorkflowPropertiesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When set to false task is not cloneable. It is set to true only if task is of ApiTask type and it is not system defined.
	Cloneable *bool `json:"Cloneable,omitempty"`
	// When set to false the task definition can only be used by internal system workflows. When set to true then the task can be included in user defined workflows.
	ExternalMeta     *bool                  `json:"ExternalMeta,omitempty"`
	InputDefinition  []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
	OutputDefinition []WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
	// The number of times a task should be tried before marking as failed.
	RetryCount *int64 `json:"RetryCount,omitempty"`
	// The delay in seconds after which the the task is re-tried.
	RetryDelay *int64 `json:"RetryDelay,omitempty"`
	// The retry policy for the task. * `Fixed` - The enum specifies the option as Fixed where the task retry happens after fixed time specified by RetryDelay.
	RetryPolicy *string `json:"RetryPolicy,omitempty"`
	// Set to true if the task implementation starts another workfow as part of the execution.
	StartsWorkflow *bool `json:"StartsWorkflow,omitempty"`
	// Supported status of the definition. * `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported. * `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added.
	SupportStatus *string `json:"SupportStatus,omitempty"`
	// The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days.
	Timeout *int64 `json:"Timeout,omitempty"`
	// The timeout policy for the task. * `Timeout` - The enum specifies the option as Timeout where task will be timed out after the specified time in Timeout property. * `Retry` - The enum specifies the option as Retry where task will be re-tried.
	TimeoutPolicy        *string `json:"TimeoutPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPropertiesAllOf WorkflowPropertiesAllOf

// NewWorkflowPropertiesAllOf instantiates a new WorkflowPropertiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPropertiesAllOf(classId string, objectType string) *WorkflowPropertiesAllOf {
	this := WorkflowPropertiesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var externalMeta bool = false
	this.ExternalMeta = &externalMeta
	var retryCount int64 = 3
	this.RetryCount = &retryCount
	var retryDelay int64 = 60
	this.RetryDelay = &retryDelay
	var retryPolicy string = "Fixed"
	this.RetryPolicy = &retryPolicy
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	var timeout int64 = 600
	this.Timeout = &timeout
	var timeoutPolicy string = "Timeout"
	this.TimeoutPolicy = &timeoutPolicy
	return &this
}

// NewWorkflowPropertiesAllOfWithDefaults instantiates a new WorkflowPropertiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPropertiesAllOfWithDefaults() *WorkflowPropertiesAllOf {
	this := WorkflowPropertiesAllOf{}
	var classId string = "workflow.Properties"
	this.ClassId = classId
	var objectType string = "workflow.Properties"
	this.ObjectType = objectType
	var externalMeta bool = false
	this.ExternalMeta = &externalMeta
	var retryCount int64 = 3
	this.RetryCount = &retryCount
	var retryDelay int64 = 60
	this.RetryDelay = &retryDelay
	var retryPolicy string = "Fixed"
	this.RetryPolicy = &retryPolicy
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	var timeout int64 = 600
	this.Timeout = &timeout
	var timeoutPolicy string = "Timeout"
	this.TimeoutPolicy = &timeoutPolicy
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPropertiesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPropertiesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPropertiesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPropertiesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloneable returns the Cloneable field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetCloneable() bool {
	if o == nil || o.Cloneable == nil {
		var ret bool
		return ret
	}
	return *o.Cloneable
}

// GetCloneableOk returns a tuple with the Cloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetCloneableOk() (*bool, bool) {
	if o == nil || o.Cloneable == nil {
		return nil, false
	}
	return o.Cloneable, true
}

// HasCloneable returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasCloneable() bool {
	if o != nil && o.Cloneable != nil {
		return true
	}

	return false
}

// SetCloneable gets a reference to the given bool and assigns it to the Cloneable field.
func (o *WorkflowPropertiesAllOf) SetCloneable(v bool) {
	o.Cloneable = &v
}

// GetExternalMeta returns the ExternalMeta field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetExternalMeta() bool {
	if o == nil || o.ExternalMeta == nil {
		var ret bool
		return ret
	}
	return *o.ExternalMeta
}

// GetExternalMetaOk returns a tuple with the ExternalMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetExternalMetaOk() (*bool, bool) {
	if o == nil || o.ExternalMeta == nil {
		return nil, false
	}
	return o.ExternalMeta, true
}

// HasExternalMeta returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasExternalMeta() bool {
	if o != nil && o.ExternalMeta != nil {
		return true
	}

	return false
}

// SetExternalMeta gets a reference to the given bool and assigns it to the ExternalMeta field.
func (o *WorkflowPropertiesAllOf) SetExternalMeta(v bool) {
	o.ExternalMeta = &v
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPropertiesAllOf) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPropertiesAllOf) GetInputDefinitionOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.InputDefinition == nil {
		return nil, false
	}
	return o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasInputDefinition() bool {
	if o != nil && o.InputDefinition != nil {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowPropertiesAllOf) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = v
}

// GetOutputDefinition returns the OutputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPropertiesAllOf) GetOutputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.OutputDefinition
}

// GetOutputDefinitionOk returns a tuple with the OutputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPropertiesAllOf) GetOutputDefinitionOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.OutputDefinition == nil {
		return nil, false
	}
	return o.OutputDefinition, true
}

// HasOutputDefinition returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasOutputDefinition() bool {
	if o != nil && o.OutputDefinition != nil {
		return true
	}

	return false
}

// SetOutputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the OutputDefinition field.
func (o *WorkflowPropertiesAllOf) SetOutputDefinition(v []WorkflowBaseDataType) {
	o.OutputDefinition = v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetRetryCount() int64 {
	if o == nil || o.RetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetRetryCountOk() (*int64, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *WorkflowPropertiesAllOf) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRetryDelay returns the RetryDelay field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetRetryDelay() int64 {
	if o == nil || o.RetryDelay == nil {
		var ret int64
		return ret
	}
	return *o.RetryDelay
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetRetryDelayOk() (*int64, bool) {
	if o == nil || o.RetryDelay == nil {
		return nil, false
	}
	return o.RetryDelay, true
}

// HasRetryDelay returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasRetryDelay() bool {
	if o != nil && o.RetryDelay != nil {
		return true
	}

	return false
}

// SetRetryDelay gets a reference to the given int64 and assigns it to the RetryDelay field.
func (o *WorkflowPropertiesAllOf) SetRetryDelay(v int64) {
	o.RetryDelay = &v
}

// GetRetryPolicy returns the RetryPolicy field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetRetryPolicy() string {
	if o == nil || o.RetryPolicy == nil {
		var ret string
		return ret
	}
	return *o.RetryPolicy
}

// GetRetryPolicyOk returns a tuple with the RetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetRetryPolicyOk() (*string, bool) {
	if o == nil || o.RetryPolicy == nil {
		return nil, false
	}
	return o.RetryPolicy, true
}

// HasRetryPolicy returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasRetryPolicy() bool {
	if o != nil && o.RetryPolicy != nil {
		return true
	}

	return false
}

// SetRetryPolicy gets a reference to the given string and assigns it to the RetryPolicy field.
func (o *WorkflowPropertiesAllOf) SetRetryPolicy(v string) {
	o.RetryPolicy = &v
}

// GetStartsWorkflow returns the StartsWorkflow field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetStartsWorkflow() bool {
	if o == nil || o.StartsWorkflow == nil {
		var ret bool
		return ret
	}
	return *o.StartsWorkflow
}

// GetStartsWorkflowOk returns a tuple with the StartsWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetStartsWorkflowOk() (*bool, bool) {
	if o == nil || o.StartsWorkflow == nil {
		return nil, false
	}
	return o.StartsWorkflow, true
}

// HasStartsWorkflow returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasStartsWorkflow() bool {
	if o != nil && o.StartsWorkflow != nil {
		return true
	}

	return false
}

// SetStartsWorkflow gets a reference to the given bool and assigns it to the StartsWorkflow field.
func (o *WorkflowPropertiesAllOf) SetStartsWorkflow(v bool) {
	o.StartsWorkflow = &v
}

// GetSupportStatus returns the SupportStatus field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetSupportStatus() string {
	if o == nil || o.SupportStatus == nil {
		var ret string
		return ret
	}
	return *o.SupportStatus
}

// GetSupportStatusOk returns a tuple with the SupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetSupportStatusOk() (*string, bool) {
	if o == nil || o.SupportStatus == nil {
		return nil, false
	}
	return o.SupportStatus, true
}

// HasSupportStatus returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasSupportStatus() bool {
	if o != nil && o.SupportStatus != nil {
		return true
	}

	return false
}

// SetSupportStatus gets a reference to the given string and assigns it to the SupportStatus field.
func (o *WorkflowPropertiesAllOf) SetSupportStatus(v string) {
	o.SupportStatus = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *WorkflowPropertiesAllOf) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetTimeoutPolicy returns the TimeoutPolicy field value if set, zero value otherwise.
func (o *WorkflowPropertiesAllOf) GetTimeoutPolicy() string {
	if o == nil || o.TimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.TimeoutPolicy
}

// GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPropertiesAllOf) GetTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.TimeoutPolicy == nil {
		return nil, false
	}
	return o.TimeoutPolicy, true
}

// HasTimeoutPolicy returns a boolean if a field has been set.
func (o *WorkflowPropertiesAllOf) HasTimeoutPolicy() bool {
	if o != nil && o.TimeoutPolicy != nil {
		return true
	}

	return false
}

// SetTimeoutPolicy gets a reference to the given string and assigns it to the TimeoutPolicy field.
func (o *WorkflowPropertiesAllOf) SetTimeoutPolicy(v string) {
	o.TimeoutPolicy = &v
}

func (o WorkflowPropertiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cloneable != nil {
		toSerialize["Cloneable"] = o.Cloneable
	}
	if o.ExternalMeta != nil {
		toSerialize["ExternalMeta"] = o.ExternalMeta
	}
	if o.InputDefinition != nil {
		toSerialize["InputDefinition"] = o.InputDefinition
	}
	if o.OutputDefinition != nil {
		toSerialize["OutputDefinition"] = o.OutputDefinition
	}
	if o.RetryCount != nil {
		toSerialize["RetryCount"] = o.RetryCount
	}
	if o.RetryDelay != nil {
		toSerialize["RetryDelay"] = o.RetryDelay
	}
	if o.RetryPolicy != nil {
		toSerialize["RetryPolicy"] = o.RetryPolicy
	}
	if o.StartsWorkflow != nil {
		toSerialize["StartsWorkflow"] = o.StartsWorkflow
	}
	if o.SupportStatus != nil {
		toSerialize["SupportStatus"] = o.SupportStatus
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.TimeoutPolicy != nil {
		toSerialize["TimeoutPolicy"] = o.TimeoutPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPropertiesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowPropertiesAllOf := _WorkflowPropertiesAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowPropertiesAllOf); err == nil {
		*o = WorkflowPropertiesAllOf(varWorkflowPropertiesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cloneable")
		delete(additionalProperties, "ExternalMeta")
		delete(additionalProperties, "InputDefinition")
		delete(additionalProperties, "OutputDefinition")
		delete(additionalProperties, "RetryCount")
		delete(additionalProperties, "RetryDelay")
		delete(additionalProperties, "RetryPolicy")
		delete(additionalProperties, "StartsWorkflow")
		delete(additionalProperties, "SupportStatus")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "TimeoutPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPropertiesAllOf struct {
	value *WorkflowPropertiesAllOf
	isSet bool
}

func (v NullableWorkflowPropertiesAllOf) Get() *WorkflowPropertiesAllOf {
	return v.value
}

func (v *NullableWorkflowPropertiesAllOf) Set(val *WorkflowPropertiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPropertiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPropertiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPropertiesAllOf(val *WorkflowPropertiesAllOf) *NullableWorkflowPropertiesAllOf {
	return &NullableWorkflowPropertiesAllOf{value: val, isSet: true}
}

func (v NullableWorkflowPropertiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPropertiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
