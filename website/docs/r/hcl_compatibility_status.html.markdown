---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_compatibility_status"
description: |-
        Check the compatibility status for the given hardware and software configurations.

---

# Resource: intersight_hcl_compatibility_status
Check the compatibility status for the given hardware and software configurations.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profile_list`:(Array)
This complex property has following sub-properties:
  + `driver_iso_url`:(string) Url for the ISO with the drivers supported for the server. 
  + `error_code`:(string)(ReadOnly) Error code indicating the compatibility status.* `Success` - The input combination is valid.* `Unknown` - Unknown API request to the service.* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.* `InvalidUcsVersion` - UCS Version is not in expected format.* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.* `OSUnknown` - OS vendor or version is not known as per the HCL database.* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.* `ProductUnknown` - Product is not known as per the HCL database.* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.* `InternalError` - API requests to the service have either failed or timed out.* `MarshallingError` - Error in JSON marshalling.* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. 
  + `id`:(string) Identifier of the hardware compatibility profile. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `os_vendor`:(string) Vendor of the Operating System running on the server. 
  + `os_version`:(string) Version of the Operating System running on the server. 
  + `personality`:(string) Personality indicating the personality with the which the sever is used in a Hyperflex environment. 
  + `processor_model`:(string) Model of the processor present in the server. 
  + `products`:(Array)
This complex property has following sub-properties:
    + `driver_names`:
                (Array of schema.TypeString) -
    + `error_code`:(string)(ReadOnly) Error code indicating the support status.* `Success` - The input combination is valid.* `Unknown` - Unknown API request to the service.* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.* `InvalidUcsVersion` - UCS Version is not in expected format.* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.* `OSUnknown` - OS vendor or version is not known as per the HCL database.* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.* `ProductUnknown` - Product is not known as per the HCL database.* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.* `InternalError` - API requests to the service have either failed or timed out.* `MarshallingError` - Error in JSON marshalling.* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. 
    + `firmwares`:(Array)
This complex property has following sub-properties:
    + `driver_name`:(string) Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3. 
    + `driver_version`:(string) Version of the Driver supported. 
    + `error_code`:(string)(ReadOnly) Error code for the support status.* `Success` - The input combination is valid.* `Unknown` - Unknown API request to the service.* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.* `InvalidUcsVersion` - UCS Version is not in expected format.* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.* `OSUnknown` - OS vendor or version is not known as per the HCL database.* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.* `ProductUnknown` - Product is not known as per the HCL database.* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.* `InternalError` - API requests to the service have either failed or timed out.* `MarshallingError` - Error in JSON marshalling.* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. 
    + `firmware_version`:(string) Firmware version of the product/adapter supported. 
    + `id`:(string) Identifier of the firmware. 
    + `latest_driver`:(bool)(ReadOnly) True if the driver is latest recommended driver. 
    + `latest_firmware`:(bool)(ReadOnly) True if the firmware is latest recommended firmware. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `id`:(string) Identifier of the product. 
  + `model`:(string) Model/PID of the product/adapter. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `revision`:(string) Revision of the adapter model. 
  + `type`:(string) Type of the product/adapter say OCP, PT, GPU.* `` - Default type of the Product.* `Adapter` - Represents network adapter cards.* `StorageController` - Represents storage controllers.* `GPU` - Represents graphics cards. 
  + `vendor`:(string) Vendor of the product or adapter. 
  + `server_model`:(string) Model of the server as returned by UCSM/CIMC XML API. 
  + `server_revision`:(string) Revision of the server model. 
  + `ucs_version`:(string) Version of the UCS software. 
  + `version_type`:(string) Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.* `UCSM` - The server is managed by UCS Manager.* `IMC` - The server is standalone managed by CIMC. 
* `request_type`:(string) Type of the request to be served.* `FillSupportedVersions` - Responds with the supported firmware and driver versions. The API doesn't expect firmware and driver versions to be passed in the request and ignores if passed.* `CheckCompatibility` - Checks the compatibility for the given firmware and driver versions. This request type expects the firmware and driver versions to filled and the service validates the values and responds back with the error codes.* `GetRecommendedDrivers` - Responds with the supported drivers. The API expects firmware version to be filled. The API populates driver ISO url for the given server model. Today the link is same for all servers managed by UCSM whereas it depends on the model for Standalone servers. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_hcl_compatibility_status` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hcl_compatibility_status.example 1234567890987654321abcde
``` 
