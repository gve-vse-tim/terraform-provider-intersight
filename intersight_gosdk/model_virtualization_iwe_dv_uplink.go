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

// VirtualizationIweDvUplink A Intersight Workload Engine cluster wise distributed uplink entity.
type VirtualizationIweDvUplink struct {
	VirtualizationBaseNetwork
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                          `json:"ObjectType"`
	BondState     NullableVirtualizationBondState `json:"BondState,omitempty"`
	NetInterfaces []string                        `json:"NetInterfaces,omitempty"`
	// The vlans associated with this this cluster wide uplink.
	Vlans   *string                               `json:"Vlans,omitempty"`
	Cluster *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to virtualizationIweHost resources.
	MemberHosts []VirtualizationIweHostRelationship `json:"MemberHosts,omitempty"`
	// An array of relationships to virtualizationIweHostInterface resources.
	MemberUplinks        []VirtualizationIweHostInterfaceRelationship `json:"MemberUplinks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweDvUplink VirtualizationIweDvUplink

// NewVirtualizationIweDvUplink instantiates a new VirtualizationIweDvUplink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweDvUplink(classId string, objectType string) *VirtualizationIweDvUplink {
	this := VirtualizationIweDvUplink{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationIweDvUplinkWithDefaults instantiates a new VirtualizationIweDvUplink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweDvUplinkWithDefaults() *VirtualizationIweDvUplink {
	this := VirtualizationIweDvUplink{}
	var classId string = "virtualization.IweDvUplink"
	this.ClassId = classId
	var objectType string = "virtualization.IweDvUplink"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweDvUplink) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvUplink) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweDvUplink) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweDvUplink) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvUplink) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweDvUplink) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBondState returns the BondState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvUplink) GetBondState() VirtualizationBondState {
	if o == nil || o.BondState.Get() == nil {
		var ret VirtualizationBondState
		return ret
	}
	return *o.BondState.Get()
}

// GetBondStateOk returns a tuple with the BondState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvUplink) GetBondStateOk() (*VirtualizationBondState, bool) {
	if o == nil {
		return nil, false
	}
	return o.BondState.Get(), o.BondState.IsSet()
}

// HasBondState returns a boolean if a field has been set.
func (o *VirtualizationIweDvUplink) HasBondState() bool {
	if o != nil && o.BondState.IsSet() {
		return true
	}

	return false
}

// SetBondState gets a reference to the given NullableVirtualizationBondState and assigns it to the BondState field.
func (o *VirtualizationIweDvUplink) SetBondState(v VirtualizationBondState) {
	o.BondState.Set(&v)
}

// SetBondStateNil sets the value for BondState to be an explicit nil
func (o *VirtualizationIweDvUplink) SetBondStateNil() {
	o.BondState.Set(nil)
}

// UnsetBondState ensures that no value is present for BondState, not even an explicit nil
func (o *VirtualizationIweDvUplink) UnsetBondState() {
	o.BondState.Unset()
}

// GetNetInterfaces returns the NetInterfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvUplink) GetNetInterfaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NetInterfaces
}

// GetNetInterfacesOk returns a tuple with the NetInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvUplink) GetNetInterfacesOk() ([]string, bool) {
	if o == nil || o.NetInterfaces == nil {
		return nil, false
	}
	return o.NetInterfaces, true
}

// HasNetInterfaces returns a boolean if a field has been set.
func (o *VirtualizationIweDvUplink) HasNetInterfaces() bool {
	if o != nil && o.NetInterfaces != nil {
		return true
	}

	return false
}

// SetNetInterfaces gets a reference to the given []string and assigns it to the NetInterfaces field.
func (o *VirtualizationIweDvUplink) SetNetInterfaces(v []string) {
	o.NetInterfaces = v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *VirtualizationIweDvUplink) GetVlans() string {
	if o == nil || o.Vlans == nil {
		var ret string
		return ret
	}
	return *o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvUplink) GetVlansOk() (*string, bool) {
	if o == nil || o.Vlans == nil {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *VirtualizationIweDvUplink) HasVlans() bool {
	if o != nil && o.Vlans != nil {
		return true
	}

	return false
}

// SetVlans gets a reference to the given string and assigns it to the Vlans field.
func (o *VirtualizationIweDvUplink) SetVlans(v string) {
	o.Vlans = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweDvUplink) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvUplink) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweDvUplink) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweDvUplink) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetMemberHosts returns the MemberHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvUplink) GetMemberHosts() []VirtualizationIweHostRelationship {
	if o == nil {
		var ret []VirtualizationIweHostRelationship
		return ret
	}
	return o.MemberHosts
}

// GetMemberHostsOk returns a tuple with the MemberHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvUplink) GetMemberHostsOk() ([]VirtualizationIweHostRelationship, bool) {
	if o == nil || o.MemberHosts == nil {
		return nil, false
	}
	return o.MemberHosts, true
}

// HasMemberHosts returns a boolean if a field has been set.
func (o *VirtualizationIweDvUplink) HasMemberHosts() bool {
	if o != nil && o.MemberHosts != nil {
		return true
	}

	return false
}

// SetMemberHosts gets a reference to the given []VirtualizationIweHostRelationship and assigns it to the MemberHosts field.
func (o *VirtualizationIweDvUplink) SetMemberHosts(v []VirtualizationIweHostRelationship) {
	o.MemberHosts = v
}

// GetMemberUplinks returns the MemberUplinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvUplink) GetMemberUplinks() []VirtualizationIweHostInterfaceRelationship {
	if o == nil {
		var ret []VirtualizationIweHostInterfaceRelationship
		return ret
	}
	return o.MemberUplinks
}

// GetMemberUplinksOk returns a tuple with the MemberUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvUplink) GetMemberUplinksOk() ([]VirtualizationIweHostInterfaceRelationship, bool) {
	if o == nil || o.MemberUplinks == nil {
		return nil, false
	}
	return o.MemberUplinks, true
}

// HasMemberUplinks returns a boolean if a field has been set.
func (o *VirtualizationIweDvUplink) HasMemberUplinks() bool {
	if o != nil && o.MemberUplinks != nil {
		return true
	}

	return false
}

// SetMemberUplinks gets a reference to the given []VirtualizationIweHostInterfaceRelationship and assigns it to the MemberUplinks field.
func (o *VirtualizationIweDvUplink) SetMemberUplinks(v []VirtualizationIweHostInterfaceRelationship) {
	o.MemberUplinks = v
}

func (o VirtualizationIweDvUplink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseNetwork, errVirtualizationBaseNetwork := json.Marshal(o.VirtualizationBaseNetwork)
	if errVirtualizationBaseNetwork != nil {
		return []byte{}, errVirtualizationBaseNetwork
	}
	errVirtualizationBaseNetwork = json.Unmarshal([]byte(serializedVirtualizationBaseNetwork), &toSerialize)
	if errVirtualizationBaseNetwork != nil {
		return []byte{}, errVirtualizationBaseNetwork
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BondState.IsSet() {
		toSerialize["BondState"] = o.BondState.Get()
	}
	if o.NetInterfaces != nil {
		toSerialize["NetInterfaces"] = o.NetInterfaces
	}
	if o.Vlans != nil {
		toSerialize["Vlans"] = o.Vlans
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.MemberHosts != nil {
		toSerialize["MemberHosts"] = o.MemberHosts
	}
	if o.MemberUplinks != nil {
		toSerialize["MemberUplinks"] = o.MemberUplinks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweDvUplink) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationIweDvUplinkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                          `json:"ObjectType"`
		BondState     NullableVirtualizationBondState `json:"BondState,omitempty"`
		NetInterfaces []string                        `json:"NetInterfaces,omitempty"`
		// The vlans associated with this this cluster wide uplink.
		Vlans   *string                               `json:"Vlans,omitempty"`
		Cluster *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
		// An array of relationships to virtualizationIweHost resources.
		MemberHosts []VirtualizationIweHostRelationship `json:"MemberHosts,omitempty"`
		// An array of relationships to virtualizationIweHostInterface resources.
		MemberUplinks []VirtualizationIweHostInterfaceRelationship `json:"MemberUplinks,omitempty"`
	}

	varVirtualizationIweDvUplinkWithoutEmbeddedStruct := VirtualizationIweDvUplinkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationIweDvUplinkWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationIweDvUplink := _VirtualizationIweDvUplink{}
		varVirtualizationIweDvUplink.ClassId = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.ClassId
		varVirtualizationIweDvUplink.ObjectType = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.ObjectType
		varVirtualizationIweDvUplink.BondState = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.BondState
		varVirtualizationIweDvUplink.NetInterfaces = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.NetInterfaces
		varVirtualizationIweDvUplink.Vlans = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.Vlans
		varVirtualizationIweDvUplink.Cluster = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.Cluster
		varVirtualizationIweDvUplink.MemberHosts = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.MemberHosts
		varVirtualizationIweDvUplink.MemberUplinks = varVirtualizationIweDvUplinkWithoutEmbeddedStruct.MemberUplinks
		*o = VirtualizationIweDvUplink(varVirtualizationIweDvUplink)
	} else {
		return err
	}

	varVirtualizationIweDvUplink := _VirtualizationIweDvUplink{}

	err = json.Unmarshal(bytes, &varVirtualizationIweDvUplink)
	if err == nil {
		o.VirtualizationBaseNetwork = varVirtualizationIweDvUplink.VirtualizationBaseNetwork
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BondState")
		delete(additionalProperties, "NetInterfaces")
		delete(additionalProperties, "Vlans")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "MemberHosts")
		delete(additionalProperties, "MemberUplinks")

		// remove fields from embedded structs
		reflectVirtualizationBaseNetwork := reflect.ValueOf(o.VirtualizationBaseNetwork)
		for i := 0; i < reflectVirtualizationBaseNetwork.Type().NumField(); i++ {
			t := reflectVirtualizationBaseNetwork.Type().Field(i)

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

type NullableVirtualizationIweDvUplink struct {
	value *VirtualizationIweDvUplink
	isSet bool
}

func (v NullableVirtualizationIweDvUplink) Get() *VirtualizationIweDvUplink {
	return v.value
}

func (v *NullableVirtualizationIweDvUplink) Set(val *VirtualizationIweDvUplink) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweDvUplink) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweDvUplink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweDvUplink(val *VirtualizationIweDvUplink) *NullableVirtualizationIweDvUplink {
	return &NullableVirtualizationIweDvUplink{value: val, isSet: true}
}

func (v NullableVirtualizationIweDvUplink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweDvUplink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
