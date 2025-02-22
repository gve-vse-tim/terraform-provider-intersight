---
subcategory: "security"
layout: "intersight"
page_title: "Intersight: intersight_security_unit"
description: |-
        The crypto card present on a server.

---

# Data Source: intersight_security_unit
The crypto card present on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_security_unit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of the security unit. 
* `operability`:(string) Operability state of the security unit. 
* `part_number`:(string) The part number of the security unit. 
* `pci_slot`:(string) PCIe slot of the security unit in the server. 
* `power`:(string) Power state of the security unit. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `thermal`:(string) Thermal state of the security unit. 
* `unit_id`:(int) The unique identifier assigned to the security unit within the server. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) The vendor identifier of the security unit. 
* `voltage`:(string) The voltage state of the security unit. 
 
