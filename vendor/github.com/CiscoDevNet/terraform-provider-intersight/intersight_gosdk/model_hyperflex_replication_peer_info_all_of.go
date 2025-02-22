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

// HyperflexReplicationPeerInfoAllOf Definition of the list of properties defined in 'hyperflex.ReplicationPeerInfo', excluding properties defined in parent classes.
type HyperflexReplicationPeerInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                  `json:"ObjectType"`
	Datastores []HyperflexReplicationPlatDatastorePair `json:"Datastores,omitempty"`
	// Data Cluster IP for the replication peer.
	Dcip *string                          `json:"Dcip,omitempty"`
	Er   NullableHyperflexEntityReference `json:"Er,omitempty"`
	// Management Cluster IP for the replication peer.
	Mcip  *string                            `json:"Mcip,omitempty"`
	Ports []HyperflexPortTypeToPortNumberMap `json:"Ports,omitempty"`
	// Replication Cluster IP for the replication peer.
	ReplCip *string `json:"ReplCip,omitempty"`
	// Peer Cluster Status for the replication peer.
	Status *string `json:"Status,omitempty"`
	// Peer Cluster Status Details for the replication peer.
	StatusDetails        *string `json:"StatusDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReplicationPeerInfoAllOf HyperflexReplicationPeerInfoAllOf

// NewHyperflexReplicationPeerInfoAllOf instantiates a new HyperflexReplicationPeerInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationPeerInfoAllOf(classId string, objectType string) *HyperflexReplicationPeerInfoAllOf {
	this := HyperflexReplicationPeerInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexReplicationPeerInfoAllOfWithDefaults instantiates a new HyperflexReplicationPeerInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationPeerInfoAllOfWithDefaults() *HyperflexReplicationPeerInfoAllOf {
	this := HyperflexReplicationPeerInfoAllOf{}
	var classId string = "hyperflex.ReplicationPeerInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationPeerInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationPeerInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationPeerInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationPeerInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationPeerInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatastores returns the Datastores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPeerInfoAllOf) GetDatastores() []HyperflexReplicationPlatDatastorePair {
	if o == nil {
		var ret []HyperflexReplicationPlatDatastorePair
		return ret
	}
	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPeerInfoAllOf) GetDatastoresOk() ([]HyperflexReplicationPlatDatastorePair, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []HyperflexReplicationPlatDatastorePair and assigns it to the Datastores field.
func (o *HyperflexReplicationPeerInfoAllOf) SetDatastores(v []HyperflexReplicationPlatDatastorePair) {
	o.Datastores = v
}

// GetDcip returns the Dcip field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfoAllOf) GetDcip() string {
	if o == nil || o.Dcip == nil {
		var ret string
		return ret
	}
	return *o.Dcip
}

// GetDcipOk returns a tuple with the Dcip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfoAllOf) GetDcipOk() (*string, bool) {
	if o == nil || o.Dcip == nil {
		return nil, false
	}
	return o.Dcip, true
}

// HasDcip returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasDcip() bool {
	if o != nil && o.Dcip != nil {
		return true
	}

	return false
}

// SetDcip gets a reference to the given string and assigns it to the Dcip field.
func (o *HyperflexReplicationPeerInfoAllOf) SetDcip(v string) {
	o.Dcip = &v
}

// GetEr returns the Er field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPeerInfoAllOf) GetEr() HyperflexEntityReference {
	if o == nil || o.Er.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.Er.Get()
}

// GetErOk returns a tuple with the Er field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPeerInfoAllOf) GetErOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Er.Get(), o.Er.IsSet()
}

// HasEr returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasEr() bool {
	if o != nil && o.Er.IsSet() {
		return true
	}

	return false
}

// SetEr gets a reference to the given NullableHyperflexEntityReference and assigns it to the Er field.
func (o *HyperflexReplicationPeerInfoAllOf) SetEr(v HyperflexEntityReference) {
	o.Er.Set(&v)
}

// SetErNil sets the value for Er to be an explicit nil
func (o *HyperflexReplicationPeerInfoAllOf) SetErNil() {
	o.Er.Set(nil)
}

// UnsetEr ensures that no value is present for Er, not even an explicit nil
func (o *HyperflexReplicationPeerInfoAllOf) UnsetEr() {
	o.Er.Unset()
}

// GetMcip returns the Mcip field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfoAllOf) GetMcip() string {
	if o == nil || o.Mcip == nil {
		var ret string
		return ret
	}
	return *o.Mcip
}

// GetMcipOk returns a tuple with the Mcip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfoAllOf) GetMcipOk() (*string, bool) {
	if o == nil || o.Mcip == nil {
		return nil, false
	}
	return o.Mcip, true
}

// HasMcip returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasMcip() bool {
	if o != nil && o.Mcip != nil {
		return true
	}

	return false
}

// SetMcip gets a reference to the given string and assigns it to the Mcip field.
func (o *HyperflexReplicationPeerInfoAllOf) SetMcip(v string) {
	o.Mcip = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPeerInfoAllOf) GetPorts() []HyperflexPortTypeToPortNumberMap {
	if o == nil {
		var ret []HyperflexPortTypeToPortNumberMap
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPeerInfoAllOf) GetPortsOk() ([]HyperflexPortTypeToPortNumberMap, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []HyperflexPortTypeToPortNumberMap and assigns it to the Ports field.
func (o *HyperflexReplicationPeerInfoAllOf) SetPorts(v []HyperflexPortTypeToPortNumberMap) {
	o.Ports = v
}

// GetReplCip returns the ReplCip field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfoAllOf) GetReplCip() string {
	if o == nil || o.ReplCip == nil {
		var ret string
		return ret
	}
	return *o.ReplCip
}

// GetReplCipOk returns a tuple with the ReplCip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfoAllOf) GetReplCipOk() (*string, bool) {
	if o == nil || o.ReplCip == nil {
		return nil, false
	}
	return o.ReplCip, true
}

// HasReplCip returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasReplCip() bool {
	if o != nil && o.ReplCip != nil {
		return true
	}

	return false
}

// SetReplCip gets a reference to the given string and assigns it to the ReplCip field.
func (o *HyperflexReplicationPeerInfoAllOf) SetReplCip(v string) {
	o.ReplCip = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexReplicationPeerInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfoAllOf) GetStatusDetails() string {
	if o == nil || o.StatusDetails == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfoAllOf) GetStatusDetailsOk() (*string, bool) {
	if o == nil || o.StatusDetails == nil {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfoAllOf) HasStatusDetails() bool {
	if o != nil && o.StatusDetails != nil {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *HyperflexReplicationPeerInfoAllOf) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

func (o HyperflexReplicationPeerInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Datastores != nil {
		toSerialize["Datastores"] = o.Datastores
	}
	if o.Dcip != nil {
		toSerialize["Dcip"] = o.Dcip
	}
	if o.Er.IsSet() {
		toSerialize["Er"] = o.Er.Get()
	}
	if o.Mcip != nil {
		toSerialize["Mcip"] = o.Mcip
	}
	if o.Ports != nil {
		toSerialize["Ports"] = o.Ports
	}
	if o.ReplCip != nil {
		toSerialize["ReplCip"] = o.ReplCip
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusDetails != nil {
		toSerialize["StatusDetails"] = o.StatusDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexReplicationPeerInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexReplicationPeerInfoAllOf := _HyperflexReplicationPeerInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexReplicationPeerInfoAllOf); err == nil {
		*o = HyperflexReplicationPeerInfoAllOf(varHyperflexReplicationPeerInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Datastores")
		delete(additionalProperties, "Dcip")
		delete(additionalProperties, "Er")
		delete(additionalProperties, "Mcip")
		delete(additionalProperties, "Ports")
		delete(additionalProperties, "ReplCip")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexReplicationPeerInfoAllOf struct {
	value *HyperflexReplicationPeerInfoAllOf
	isSet bool
}

func (v NullableHyperflexReplicationPeerInfoAllOf) Get() *HyperflexReplicationPeerInfoAllOf {
	return v.value
}

func (v *NullableHyperflexReplicationPeerInfoAllOf) Set(val *HyperflexReplicationPeerInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationPeerInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationPeerInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationPeerInfoAllOf(val *HyperflexReplicationPeerInfoAllOf) *NullableHyperflexReplicationPeerInfoAllOf {
	return &NullableHyperflexReplicationPeerInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexReplicationPeerInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationPeerInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
