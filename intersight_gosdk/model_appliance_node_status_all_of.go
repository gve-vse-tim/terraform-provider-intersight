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

// ApplianceNodeStatusAllOf Definition of the list of properties defined in 'appliance.NodeStatus', excluding properties defined in parent classes.
type ApplianceNodeStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Percentage of CPU currently in use.
	CpuUsage *float32 `json:"CpuUsage,omitempty"`
	// Percentage of memory currently in use.
	MemUsage *float32 `json:"MemUsage,omitempty"`
	// System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1.
	NodeId *int64 `json:"NodeId,omitempty"`
	// State of the node in terms of its readiness to host Kubernetes pods. * `Down` - The node is yet to come up and join as a member of theKubernetes cluster. * `Preparing` - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods. * `Ready` - The node is ready to host Kubernetes pods.
	NodeState *string `json:"NodeState,omitempty"`
	// Operational status of the Intersight Appliance node. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
	OperationalStatus *string                `json:"OperationalStatus,omitempty"`
	StatusChecks      []ApplianceStatusCheck `json:"StatusChecks,omitempty"`
	// An array of relationships to applianceFileSystemStatus resources.
	FileSystemStatuses   []ApplianceFileSystemStatusRelationship `json:"FileSystemStatuses,omitempty"`
	NodeInfo             *ApplianceNodeInfoRelationship          `json:"NodeInfo,omitempty"`
	SystemStatus         *ApplianceSystemStatusRelationship      `json:"SystemStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNodeStatusAllOf ApplianceNodeStatusAllOf

// NewApplianceNodeStatusAllOf instantiates a new ApplianceNodeStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNodeStatusAllOf(classId string, objectType string) *ApplianceNodeStatusAllOf {
	this := ApplianceNodeStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNodeStatusAllOfWithDefaults instantiates a new ApplianceNodeStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNodeStatusAllOfWithDefaults() *ApplianceNodeStatusAllOf {
	this := ApplianceNodeStatusAllOf{}
	var classId string = "appliance.NodeStatus"
	this.ClassId = classId
	var objectType string = "appliance.NodeStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNodeStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNodeStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNodeStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNodeStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise.
func (o *ApplianceNodeStatusAllOf) GetCpuUsage() float32 {
	if o == nil || o.CpuUsage == nil {
		var ret float32
		return ret
	}
	return *o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetCpuUsageOk() (*float32, bool) {
	if o == nil || o.CpuUsage == nil {
		return nil, false
	}
	return o.CpuUsage, true
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasCpuUsage() bool {
	if o != nil && o.CpuUsage != nil {
		return true
	}

	return false
}

// SetCpuUsage gets a reference to the given float32 and assigns it to the CpuUsage field.
func (o *ApplianceNodeStatusAllOf) SetCpuUsage(v float32) {
	o.CpuUsage = &v
}

// GetMemUsage returns the MemUsage field value if set, zero value otherwise.
func (o *ApplianceNodeStatusAllOf) GetMemUsage() float32 {
	if o == nil || o.MemUsage == nil {
		var ret float32
		return ret
	}
	return *o.MemUsage
}

// GetMemUsageOk returns a tuple with the MemUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetMemUsageOk() (*float32, bool) {
	if o == nil || o.MemUsage == nil {
		return nil, false
	}
	return o.MemUsage, true
}

// HasMemUsage returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasMemUsage() bool {
	if o != nil && o.MemUsage != nil {
		return true
	}

	return false
}

// SetMemUsage gets a reference to the given float32 and assigns it to the MemUsage field.
func (o *ApplianceNodeStatusAllOf) SetMemUsage(v float32) {
	o.MemUsage = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ApplianceNodeStatusAllOf) GetNodeId() int64 {
	if o == nil || o.NodeId == nil {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetNodeIdOk() (*int64, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *ApplianceNodeStatusAllOf) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetNodeState returns the NodeState field value if set, zero value otherwise.
func (o *ApplianceNodeStatusAllOf) GetNodeState() string {
	if o == nil || o.NodeState == nil {
		var ret string
		return ret
	}
	return *o.NodeState
}

// GetNodeStateOk returns a tuple with the NodeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetNodeStateOk() (*string, bool) {
	if o == nil || o.NodeState == nil {
		return nil, false
	}
	return o.NodeState, true
}

// HasNodeState returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasNodeState() bool {
	if o != nil && o.NodeState != nil {
		return true
	}

	return false
}

// SetNodeState gets a reference to the given string and assigns it to the NodeState field.
func (o *ApplianceNodeStatusAllOf) SetNodeState(v string) {
	o.NodeState = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceNodeStatusAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceNodeStatusAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetStatusChecks returns the StatusChecks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeStatusAllOf) GetStatusChecks() []ApplianceStatusCheck {
	if o == nil {
		var ret []ApplianceStatusCheck
		return ret
	}
	return o.StatusChecks
}

// GetStatusChecksOk returns a tuple with the StatusChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeStatusAllOf) GetStatusChecksOk() ([]ApplianceStatusCheck, bool) {
	if o == nil || o.StatusChecks == nil {
		return nil, false
	}
	return o.StatusChecks, true
}

// HasStatusChecks returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasStatusChecks() bool {
	if o != nil && o.StatusChecks != nil {
		return true
	}

	return false
}

// SetStatusChecks gets a reference to the given []ApplianceStatusCheck and assigns it to the StatusChecks field.
func (o *ApplianceNodeStatusAllOf) SetStatusChecks(v []ApplianceStatusCheck) {
	o.StatusChecks = v
}

// GetFileSystemStatuses returns the FileSystemStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeStatusAllOf) GetFileSystemStatuses() []ApplianceFileSystemStatusRelationship {
	if o == nil {
		var ret []ApplianceFileSystemStatusRelationship
		return ret
	}
	return o.FileSystemStatuses
}

// GetFileSystemStatusesOk returns a tuple with the FileSystemStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeStatusAllOf) GetFileSystemStatusesOk() ([]ApplianceFileSystemStatusRelationship, bool) {
	if o == nil || o.FileSystemStatuses == nil {
		return nil, false
	}
	return o.FileSystemStatuses, true
}

// HasFileSystemStatuses returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasFileSystemStatuses() bool {
	if o != nil && o.FileSystemStatuses != nil {
		return true
	}

	return false
}

// SetFileSystemStatuses gets a reference to the given []ApplianceFileSystemStatusRelationship and assigns it to the FileSystemStatuses field.
func (o *ApplianceNodeStatusAllOf) SetFileSystemStatuses(v []ApplianceFileSystemStatusRelationship) {
	o.FileSystemStatuses = v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *ApplianceNodeStatusAllOf) GetNodeInfo() ApplianceNodeInfoRelationship {
	if o == nil || o.NodeInfo == nil {
		var ret ApplianceNodeInfoRelationship
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetNodeInfoOk() (*ApplianceNodeInfoRelationship, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given ApplianceNodeInfoRelationship and assigns it to the NodeInfo field.
func (o *ApplianceNodeStatusAllOf) SetNodeInfo(v ApplianceNodeInfoRelationship) {
	o.NodeInfo = &v
}

// GetSystemStatus returns the SystemStatus field value if set, zero value otherwise.
func (o *ApplianceNodeStatusAllOf) GetSystemStatus() ApplianceSystemStatusRelationship {
	if o == nil || o.SystemStatus == nil {
		var ret ApplianceSystemStatusRelationship
		return ret
	}
	return *o.SystemStatus
}

// GetSystemStatusOk returns a tuple with the SystemStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatusAllOf) GetSystemStatusOk() (*ApplianceSystemStatusRelationship, bool) {
	if o == nil || o.SystemStatus == nil {
		return nil, false
	}
	return o.SystemStatus, true
}

// HasSystemStatus returns a boolean if a field has been set.
func (o *ApplianceNodeStatusAllOf) HasSystemStatus() bool {
	if o != nil && o.SystemStatus != nil {
		return true
	}

	return false
}

// SetSystemStatus gets a reference to the given ApplianceSystemStatusRelationship and assigns it to the SystemStatus field.
func (o *ApplianceNodeStatusAllOf) SetSystemStatus(v ApplianceSystemStatusRelationship) {
	o.SystemStatus = &v
}

func (o ApplianceNodeStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpuUsage != nil {
		toSerialize["CpuUsage"] = o.CpuUsage
	}
	if o.MemUsage != nil {
		toSerialize["MemUsage"] = o.MemUsage
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.NodeState != nil {
		toSerialize["NodeState"] = o.NodeState
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.StatusChecks != nil {
		toSerialize["StatusChecks"] = o.StatusChecks
	}
	if o.FileSystemStatuses != nil {
		toSerialize["FileSystemStatuses"] = o.FileSystemStatuses
	}
	if o.NodeInfo != nil {
		toSerialize["NodeInfo"] = o.NodeInfo
	}
	if o.SystemStatus != nil {
		toSerialize["SystemStatus"] = o.SystemStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceNodeStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceNodeStatusAllOf := _ApplianceNodeStatusAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceNodeStatusAllOf); err == nil {
		*o = ApplianceNodeStatusAllOf(varApplianceNodeStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuUsage")
		delete(additionalProperties, "MemUsage")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "NodeState")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "StatusChecks")
		delete(additionalProperties, "FileSystemStatuses")
		delete(additionalProperties, "NodeInfo")
		delete(additionalProperties, "SystemStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceNodeStatusAllOf struct {
	value *ApplianceNodeStatusAllOf
	isSet bool
}

func (v NullableApplianceNodeStatusAllOf) Get() *ApplianceNodeStatusAllOf {
	return v.value
}

func (v *NullableApplianceNodeStatusAllOf) Set(val *ApplianceNodeStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNodeStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNodeStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNodeStatusAllOf(val *ApplianceNodeStatusAllOf) *NullableApplianceNodeStatusAllOf {
	return &NullableApplianceNodeStatusAllOf{value: val, isSet: true}
}

func (v NullableApplianceNodeStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNodeStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
