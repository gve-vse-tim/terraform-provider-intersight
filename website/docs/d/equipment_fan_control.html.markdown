---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_fan_control"
description: |-
        This represents the states of Fan equipment.

---

# Data Source: intersight_equipment_fan_control
This represents the states of Fan equipment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_fan_control.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) This field identifies the Fan Control Mode on the endpoint.* `Balanced` - Value of Fan Speed is Balanced.* `LowPower` - Value of Fan Speed is LowPower.* `HighPower` - Value of Fan Speed is HighPower.* `MaximumPower` - Value of Fan Speed is MaximumPower.* `Acoustic` - Value of Fan Speed is Acoustic. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 
