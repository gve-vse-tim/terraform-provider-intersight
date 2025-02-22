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

// ServerProfileAllOf Definition of the list of properties defined in 'server.Profile', excluding properties defined in parent classes.
type ServerProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType          string                            `json:"ObjectType"`
	ConfigChangeContext NullablePolicyConfigChangeContext `json:"ConfigChangeContext,omitempty"`
	ConfigChanges       NullablePolicyConfigChange        `json:"ConfigChanges,omitempty"`
	// Indicates whether the value of the 'pmcDeployedSecurePassphrase' property has been set.
	IsPmcDeployedSecurePassphraseSet *bool `json:"IsPmcDeployedSecurePassphraseSet,omitempty"`
	// Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy.
	PmcDeployedSecurePassphrase *string `json:"PmcDeployedSecurePassphrase,omitempty"`
	// Source of the server assigned to the server profile. Values can be Static, Pool or None. Static is used if a server is attached directly to server profile. Pool is used if a resource pool is attached to server profile. None is used if no server or resource pool is attached to server profile. * `None` - No server is assigned to the server profile. * `Static` - Server is directly assigned to server profile using assign server. * `Pool` - Server is assigned from a resource pool.
	ServerAssignmentMode *string `json:"ServerAssignmentMode,omitempty"`
	// The UUID address for the server must include UUID prefix xxxxxxxx-xxxx-xxxx along with the UUID suffix of format xxxx-xxxxxxxxxxxx.
	StaticUuidAddress *string `json:"StaticUuidAddress,omitempty"`
	// The UUID address that is assigned to the server based on the UUID pool.
	Uuid                 *string                       `json:"Uuid,omitempty"`
	AssignedServer       *ComputePhysicalRelationship  `json:"AssignedServer,omitempty"`
	AssociatedServer     *ComputePhysicalRelationship  `json:"AssociatedServer,omitempty"`
	AssociatedServerPool *ResourcepoolPoolRelationship `json:"AssociatedServerPool,omitempty"`
	// An array of relationships to serverConfigChangeDetail resources.
	ConfigChangeDetails []ServerConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
	LeasedServer        *ComputePhysicalRelationship           `json:"LeasedServer,omitempty"`
	Organization        *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	ResourceLease       *ResourcepoolLeaseRelationship         `json:"ResourceLease,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	ServerPool           *ResourcepoolPoolRelationship      `json:"ServerPool,omitempty"`
	UuidLease            *UuidpoolUuidLeaseRelationship     `json:"UuidLease,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerProfileAllOf ServerProfileAllOf

// NewServerProfileAllOf instantiates a new ServerProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProfileAllOf(classId string, objectType string) *ServerProfileAllOf {
	this := ServerProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var serverAssignmentMode string = "None"
	this.ServerAssignmentMode = &serverAssignmentMode
	return &this
}

// NewServerProfileAllOfWithDefaults instantiates a new ServerProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerProfileAllOfWithDefaults() *ServerProfileAllOf {
	this := ServerProfileAllOf{}
	var classId string = "server.Profile"
	this.ClassId = classId
	var objectType string = "server.Profile"
	this.ObjectType = objectType
	var serverAssignmentMode string = "None"
	this.ServerAssignmentMode = &serverAssignmentMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChangeContext returns the ConfigChangeContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetConfigChangeContext() PolicyConfigChangeContext {
	if o == nil || o.ConfigChangeContext.Get() == nil {
		var ret PolicyConfigChangeContext
		return ret
	}
	return *o.ConfigChangeContext.Get()
}

// GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetConfigChangeContextOk() (*PolicyConfigChangeContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChangeContext.Get(), o.ConfigChangeContext.IsSet()
}

// HasConfigChangeContext returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasConfigChangeContext() bool {
	if o != nil && o.ConfigChangeContext.IsSet() {
		return true
	}

	return false
}

// SetConfigChangeContext gets a reference to the given NullablePolicyConfigChangeContext and assigns it to the ConfigChangeContext field.
func (o *ServerProfileAllOf) SetConfigChangeContext(v PolicyConfigChangeContext) {
	o.ConfigChangeContext.Set(&v)
}

// SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil
func (o *ServerProfileAllOf) SetConfigChangeContextNil() {
	o.ConfigChangeContext.Set(nil)
}

// UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
func (o *ServerProfileAllOf) UnsetConfigChangeContext() {
	o.ConfigChangeContext.Unset()
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ServerProfileAllOf) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ServerProfileAllOf) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ServerProfileAllOf) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetIsPmcDeployedSecurePassphraseSet returns the IsPmcDeployedSecurePassphraseSet field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetIsPmcDeployedSecurePassphraseSet() bool {
	if o == nil || o.IsPmcDeployedSecurePassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPmcDeployedSecurePassphraseSet
}

// GetIsPmcDeployedSecurePassphraseSetOk returns a tuple with the IsPmcDeployedSecurePassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetIsPmcDeployedSecurePassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsPmcDeployedSecurePassphraseSet == nil {
		return nil, false
	}
	return o.IsPmcDeployedSecurePassphraseSet, true
}

// HasIsPmcDeployedSecurePassphraseSet returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasIsPmcDeployedSecurePassphraseSet() bool {
	if o != nil && o.IsPmcDeployedSecurePassphraseSet != nil {
		return true
	}

	return false
}

// SetIsPmcDeployedSecurePassphraseSet gets a reference to the given bool and assigns it to the IsPmcDeployedSecurePassphraseSet field.
func (o *ServerProfileAllOf) SetIsPmcDeployedSecurePassphraseSet(v bool) {
	o.IsPmcDeployedSecurePassphraseSet = &v
}

// GetPmcDeployedSecurePassphrase returns the PmcDeployedSecurePassphrase field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetPmcDeployedSecurePassphrase() string {
	if o == nil || o.PmcDeployedSecurePassphrase == nil {
		var ret string
		return ret
	}
	return *o.PmcDeployedSecurePassphrase
}

// GetPmcDeployedSecurePassphraseOk returns a tuple with the PmcDeployedSecurePassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetPmcDeployedSecurePassphraseOk() (*string, bool) {
	if o == nil || o.PmcDeployedSecurePassphrase == nil {
		return nil, false
	}
	return o.PmcDeployedSecurePassphrase, true
}

// HasPmcDeployedSecurePassphrase returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasPmcDeployedSecurePassphrase() bool {
	if o != nil && o.PmcDeployedSecurePassphrase != nil {
		return true
	}

	return false
}

// SetPmcDeployedSecurePassphrase gets a reference to the given string and assigns it to the PmcDeployedSecurePassphrase field.
func (o *ServerProfileAllOf) SetPmcDeployedSecurePassphrase(v string) {
	o.PmcDeployedSecurePassphrase = &v
}

// GetServerAssignmentMode returns the ServerAssignmentMode field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetServerAssignmentMode() string {
	if o == nil || o.ServerAssignmentMode == nil {
		var ret string
		return ret
	}
	return *o.ServerAssignmentMode
}

// GetServerAssignmentModeOk returns a tuple with the ServerAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetServerAssignmentModeOk() (*string, bool) {
	if o == nil || o.ServerAssignmentMode == nil {
		return nil, false
	}
	return o.ServerAssignmentMode, true
}

// HasServerAssignmentMode returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasServerAssignmentMode() bool {
	if o != nil && o.ServerAssignmentMode != nil {
		return true
	}

	return false
}

// SetServerAssignmentMode gets a reference to the given string and assigns it to the ServerAssignmentMode field.
func (o *ServerProfileAllOf) SetServerAssignmentMode(v string) {
	o.ServerAssignmentMode = &v
}

// GetStaticUuidAddress returns the StaticUuidAddress field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetStaticUuidAddress() string {
	if o == nil || o.StaticUuidAddress == nil {
		var ret string
		return ret
	}
	return *o.StaticUuidAddress
}

// GetStaticUuidAddressOk returns a tuple with the StaticUuidAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetStaticUuidAddressOk() (*string, bool) {
	if o == nil || o.StaticUuidAddress == nil {
		return nil, false
	}
	return o.StaticUuidAddress, true
}

// HasStaticUuidAddress returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasStaticUuidAddress() bool {
	if o != nil && o.StaticUuidAddress != nil {
		return true
	}

	return false
}

// SetStaticUuidAddress gets a reference to the given string and assigns it to the StaticUuidAddress field.
func (o *ServerProfileAllOf) SetStaticUuidAddress(v string) {
	o.StaticUuidAddress = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ServerProfileAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetAssignedServer returns the AssignedServer field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetAssignedServer() ComputePhysicalRelationship {
	if o == nil || o.AssignedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.AssignedServer
}

// GetAssignedServerOk returns a tuple with the AssignedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetAssignedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.AssignedServer == nil {
		return nil, false
	}
	return o.AssignedServer, true
}

// HasAssignedServer returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasAssignedServer() bool {
	if o != nil && o.AssignedServer != nil {
		return true
	}

	return false
}

// SetAssignedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the AssignedServer field.
func (o *ServerProfileAllOf) SetAssignedServer(v ComputePhysicalRelationship) {
	o.AssignedServer = &v
}

// GetAssociatedServer returns the AssociatedServer field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetAssociatedServer() ComputePhysicalRelationship {
	if o == nil || o.AssociatedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.AssociatedServer
}

// GetAssociatedServerOk returns a tuple with the AssociatedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetAssociatedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.AssociatedServer == nil {
		return nil, false
	}
	return o.AssociatedServer, true
}

// HasAssociatedServer returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasAssociatedServer() bool {
	if o != nil && o.AssociatedServer != nil {
		return true
	}

	return false
}

// SetAssociatedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the AssociatedServer field.
func (o *ServerProfileAllOf) SetAssociatedServer(v ComputePhysicalRelationship) {
	o.AssociatedServer = &v
}

// GetAssociatedServerPool returns the AssociatedServerPool field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetAssociatedServerPool() ResourcepoolPoolRelationship {
	if o == nil || o.AssociatedServerPool == nil {
		var ret ResourcepoolPoolRelationship
		return ret
	}
	return *o.AssociatedServerPool
}

// GetAssociatedServerPoolOk returns a tuple with the AssociatedServerPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetAssociatedServerPoolOk() (*ResourcepoolPoolRelationship, bool) {
	if o == nil || o.AssociatedServerPool == nil {
		return nil, false
	}
	return o.AssociatedServerPool, true
}

// HasAssociatedServerPool returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasAssociatedServerPool() bool {
	if o != nil && o.AssociatedServerPool != nil {
		return true
	}

	return false
}

// SetAssociatedServerPool gets a reference to the given ResourcepoolPoolRelationship and assigns it to the AssociatedServerPool field.
func (o *ServerProfileAllOf) SetAssociatedServerPool(v ResourcepoolPoolRelationship) {
	o.AssociatedServerPool = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetConfigChangeDetails() []ServerConfigChangeDetailRelationship {
	if o == nil {
		var ret []ServerConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetConfigChangeDetailsOk() ([]ServerConfigChangeDetailRelationship, bool) {
	if o == nil || o.ConfigChangeDetails == nil {
		return nil, false
	}
	return o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasConfigChangeDetails() bool {
	if o != nil && o.ConfigChangeDetails != nil {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []ServerConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *ServerProfileAllOf) SetConfigChangeDetails(v []ServerConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetLeasedServer returns the LeasedServer field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetLeasedServer() ComputePhysicalRelationship {
	if o == nil || o.LeasedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.LeasedServer
}

// GetLeasedServerOk returns a tuple with the LeasedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetLeasedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.LeasedServer == nil {
		return nil, false
	}
	return o.LeasedServer, true
}

// HasLeasedServer returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasLeasedServer() bool {
	if o != nil && o.LeasedServer != nil {
		return true
	}

	return false
}

// SetLeasedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the LeasedServer field.
func (o *ServerProfileAllOf) SetLeasedServer(v ComputePhysicalRelationship) {
	o.LeasedServer = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ServerProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetResourceLease returns the ResourceLease field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetResourceLease() ResourcepoolLeaseRelationship {
	if o == nil || o.ResourceLease == nil {
		var ret ResourcepoolLeaseRelationship
		return ret
	}
	return *o.ResourceLease
}

// GetResourceLeaseOk returns a tuple with the ResourceLease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetResourceLeaseOk() (*ResourcepoolLeaseRelationship, bool) {
	if o == nil || o.ResourceLease == nil {
		return nil, false
	}
	return o.ResourceLease, true
}

// HasResourceLease returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasResourceLease() bool {
	if o != nil && o.ResourceLease != nil {
		return true
	}

	return false
}

// SetResourceLease gets a reference to the given ResourcepoolLeaseRelationship and assigns it to the ResourceLease field.
func (o *ServerProfileAllOf) SetResourceLease(v ResourcepoolLeaseRelationship) {
	o.ResourceLease = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerProfileAllOf) GetRunningWorkflowsOk() ([]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *ServerProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

// GetServerPool returns the ServerPool field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetServerPool() ResourcepoolPoolRelationship {
	if o == nil || o.ServerPool == nil {
		var ret ResourcepoolPoolRelationship
		return ret
	}
	return *o.ServerPool
}

// GetServerPoolOk returns a tuple with the ServerPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetServerPoolOk() (*ResourcepoolPoolRelationship, bool) {
	if o == nil || o.ServerPool == nil {
		return nil, false
	}
	return o.ServerPool, true
}

// HasServerPool returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasServerPool() bool {
	if o != nil && o.ServerPool != nil {
		return true
	}

	return false
}

// SetServerPool gets a reference to the given ResourcepoolPoolRelationship and assigns it to the ServerPool field.
func (o *ServerProfileAllOf) SetServerPool(v ResourcepoolPoolRelationship) {
	o.ServerPool = &v
}

// GetUuidLease returns the UuidLease field value if set, zero value otherwise.
func (o *ServerProfileAllOf) GetUuidLease() UuidpoolUuidLeaseRelationship {
	if o == nil || o.UuidLease == nil {
		var ret UuidpoolUuidLeaseRelationship
		return ret
	}
	return *o.UuidLease
}

// GetUuidLeaseOk returns a tuple with the UuidLease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileAllOf) GetUuidLeaseOk() (*UuidpoolUuidLeaseRelationship, bool) {
	if o == nil || o.UuidLease == nil {
		return nil, false
	}
	return o.UuidLease, true
}

// HasUuidLease returns a boolean if a field has been set.
func (o *ServerProfileAllOf) HasUuidLease() bool {
	if o != nil && o.UuidLease != nil {
		return true
	}

	return false
}

// SetUuidLease gets a reference to the given UuidpoolUuidLeaseRelationship and assigns it to the UuidLease field.
func (o *ServerProfileAllOf) SetUuidLease(v UuidpoolUuidLeaseRelationship) {
	o.UuidLease = &v
}

func (o ServerProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigChangeContext.IsSet() {
		toSerialize["ConfigChangeContext"] = o.ConfigChangeContext.Get()
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if o.IsPmcDeployedSecurePassphraseSet != nil {
		toSerialize["IsPmcDeployedSecurePassphraseSet"] = o.IsPmcDeployedSecurePassphraseSet
	}
	if o.PmcDeployedSecurePassphrase != nil {
		toSerialize["PmcDeployedSecurePassphrase"] = o.PmcDeployedSecurePassphrase
	}
	if o.ServerAssignmentMode != nil {
		toSerialize["ServerAssignmentMode"] = o.ServerAssignmentMode
	}
	if o.StaticUuidAddress != nil {
		toSerialize["StaticUuidAddress"] = o.StaticUuidAddress
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.AssignedServer != nil {
		toSerialize["AssignedServer"] = o.AssignedServer
	}
	if o.AssociatedServer != nil {
		toSerialize["AssociatedServer"] = o.AssociatedServer
	}
	if o.AssociatedServerPool != nil {
		toSerialize["AssociatedServerPool"] = o.AssociatedServerPool
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.LeasedServer != nil {
		toSerialize["LeasedServer"] = o.LeasedServer
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ResourceLease != nil {
		toSerialize["ResourceLease"] = o.ResourceLease
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}
	if o.ServerPool != nil {
		toSerialize["ServerPool"] = o.ServerPool
	}
	if o.UuidLease != nil {
		toSerialize["UuidLease"] = o.UuidLease
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerProfileAllOf := _ServerProfileAllOf{}

	if err = json.Unmarshal(bytes, &varServerProfileAllOf); err == nil {
		*o = ServerProfileAllOf(varServerProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChangeContext")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "IsPmcDeployedSecurePassphraseSet")
		delete(additionalProperties, "PmcDeployedSecurePassphrase")
		delete(additionalProperties, "ServerAssignmentMode")
		delete(additionalProperties, "StaticUuidAddress")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "AssignedServer")
		delete(additionalProperties, "AssociatedServer")
		delete(additionalProperties, "AssociatedServerPool")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "LeasedServer")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "ResourceLease")
		delete(additionalProperties, "RunningWorkflows")
		delete(additionalProperties, "ServerPool")
		delete(additionalProperties, "UuidLease")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerProfileAllOf struct {
	value *ServerProfileAllOf
	isSet bool
}

func (v NullableServerProfileAllOf) Get() *ServerProfileAllOf {
	return v.value
}

func (v *NullableServerProfileAllOf) Set(val *ServerProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProfileAllOf(val *ServerProfileAllOf) *NullableServerProfileAllOf {
	return &NullableServerProfileAllOf{value: val, isSet: true}
}

func (v NullableServerProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
