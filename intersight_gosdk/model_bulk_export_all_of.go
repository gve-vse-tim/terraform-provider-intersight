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

// BulkExportAllOf Definition of the list of properties defined in 'bulk.Export', excluding properties defined in parent classes.
type BulkExportAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be performed on the export operation. * `Start` - Starts the export operation. * `Cancel` - Cancels the export operation that is in progress.
	Action *string `json:"Action,omitempty"`
	// Specifies whether tags must be exported and will be considered for all the items MOs.
	ExportTags      *bool            `json:"ExportTags,omitempty"`
	ExportedObjects []BulkSubRequest `json:"ExportedObjects,omitempty"`
	// Contains the list of import order.
	ImportOrder interface{} `json:"ImportOrder,omitempty"`
	Items       []MoMoRef   `json:"Items,omitempty"`
	// An identifier for the export instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_).
	Name *string `json:"Name,omitempty"`
	// Status of the export operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `OrderInProgress` - The archive operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `OperationTimedOut` - The operation has timed out. * `OperationCancelled` - The operation has been cancelled. * `CancelInProgress` - The operation is being cancelled.
	Status *string `json:"Status,omitempty"`
	// Status message associated with failures or progress indication.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// An array of relationships to bulkExportedItem resources.
	ExportedItems        []BulkExportedItemRelationship        `json:"ExportedItems,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkExportAllOf BulkExportAllOf

// NewBulkExportAllOf instantiates a new BulkExportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkExportAllOf(classId string, objectType string) *BulkExportAllOf {
	this := BulkExportAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "Start"
	this.Action = &action
	var exportTags bool = true
	this.ExportTags = &exportTags
	return &this
}

// NewBulkExportAllOfWithDefaults instantiates a new BulkExportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkExportAllOfWithDefaults() *BulkExportAllOf {
	this := BulkExportAllOf{}
	var classId string = "bulk.Export"
	this.ClassId = classId
	var objectType string = "bulk.Export"
	this.ObjectType = objectType
	var action string = "Start"
	this.Action = &action
	var exportTags bool = true
	this.ExportTags = &exportTags
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkExportAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkExportAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkExportAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkExportAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BulkExportAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BulkExportAllOf) SetAction(v string) {
	o.Action = &v
}

// GetExportTags returns the ExportTags field value if set, zero value otherwise.
func (o *BulkExportAllOf) GetExportTags() bool {
	if o == nil || o.ExportTags == nil {
		var ret bool
		return ret
	}
	return *o.ExportTags
}

// GetExportTagsOk returns a tuple with the ExportTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetExportTagsOk() (*bool, bool) {
	if o == nil || o.ExportTags == nil {
		return nil, false
	}
	return o.ExportTags, true
}

// HasExportTags returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasExportTags() bool {
	if o != nil && o.ExportTags != nil {
		return true
	}

	return false
}

// SetExportTags gets a reference to the given bool and assigns it to the ExportTags field.
func (o *BulkExportAllOf) SetExportTags(v bool) {
	o.ExportTags = &v
}

// GetExportedObjects returns the ExportedObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExportAllOf) GetExportedObjects() []BulkSubRequest {
	if o == nil {
		var ret []BulkSubRequest
		return ret
	}
	return o.ExportedObjects
}

// GetExportedObjectsOk returns a tuple with the ExportedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExportAllOf) GetExportedObjectsOk() ([]BulkSubRequest, bool) {
	if o == nil || o.ExportedObjects == nil {
		return nil, false
	}
	return o.ExportedObjects, true
}

// HasExportedObjects returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasExportedObjects() bool {
	if o != nil && o.ExportedObjects != nil {
		return true
	}

	return false
}

// SetExportedObjects gets a reference to the given []BulkSubRequest and assigns it to the ExportedObjects field.
func (o *BulkExportAllOf) SetExportedObjects(v []BulkSubRequest) {
	o.ExportedObjects = v
}

// GetImportOrder returns the ImportOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExportAllOf) GetImportOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImportOrder
}

// GetImportOrderOk returns a tuple with the ImportOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExportAllOf) GetImportOrderOk() (*interface{}, bool) {
	if o == nil || o.ImportOrder == nil {
		return nil, false
	}
	return &o.ImportOrder, true
}

// HasImportOrder returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasImportOrder() bool {
	if o != nil && o.ImportOrder != nil {
		return true
	}

	return false
}

// SetImportOrder gets a reference to the given interface{} and assigns it to the ImportOrder field.
func (o *BulkExportAllOf) SetImportOrder(v interface{}) {
	o.ImportOrder = v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExportAllOf) GetItems() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExportAllOf) GetItemsOk() ([]MoMoRef, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MoMoRef and assigns it to the Items field.
func (o *BulkExportAllOf) SetItems(v []MoMoRef) {
	o.Items = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BulkExportAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BulkExportAllOf) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkExportAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkExportAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkExportAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkExportAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetExportedItems returns the ExportedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExportAllOf) GetExportedItems() []BulkExportedItemRelationship {
	if o == nil {
		var ret []BulkExportedItemRelationship
		return ret
	}
	return o.ExportedItems
}

// GetExportedItemsOk returns a tuple with the ExportedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExportAllOf) GetExportedItemsOk() ([]BulkExportedItemRelationship, bool) {
	if o == nil || o.ExportedItems == nil {
		return nil, false
	}
	return o.ExportedItems, true
}

// HasExportedItems returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasExportedItems() bool {
	if o != nil && o.ExportedItems != nil {
		return true
	}

	return false
}

// SetExportedItems gets a reference to the given []BulkExportedItemRelationship and assigns it to the ExportedItems field.
func (o *BulkExportAllOf) SetExportedItems(v []BulkExportedItemRelationship) {
	o.ExportedItems = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkExportAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkExportAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkExportAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o BulkExportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ExportTags != nil {
		toSerialize["ExportTags"] = o.ExportTags
	}
	if o.ExportedObjects != nil {
		toSerialize["ExportedObjects"] = o.ExportedObjects
	}
	if o.ImportOrder != nil {
		toSerialize["ImportOrder"] = o.ImportOrder
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.ExportedItems != nil {
		toSerialize["ExportedItems"] = o.ExportedItems
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkExportAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkExportAllOf := _BulkExportAllOf{}

	if err = json.Unmarshal(bytes, &varBulkExportAllOf); err == nil {
		*o = BulkExportAllOf(varBulkExportAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "ExportTags")
		delete(additionalProperties, "ExportedObjects")
		delete(additionalProperties, "ImportOrder")
		delete(additionalProperties, "Items")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "ExportedItems")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkExportAllOf struct {
	value *BulkExportAllOf
	isSet bool
}

func (v NullableBulkExportAllOf) Get() *BulkExportAllOf {
	return v.value
}

func (v *NullableBulkExportAllOf) Set(val *BulkExportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkExportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkExportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkExportAllOf(val *BulkExportAllOf) *NullableBulkExportAllOf {
	return &NullableBulkExportAllOf{value: val, isSet: true}
}

func (v NullableBulkExportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkExportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
