---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_claim"
description: |-
        DeviceClaim managed object represents a user initiated claim request for claiming
        an endpoint device. There can be many DeviceClaim managed object for a given endpoint
        device when users claim and unclaim devices repeatedly.
        Claiming an endpoint device is a multi-step operation. The Intersight Appliance
        starts a workflow with multiple tasks to process the device claim request. The status
        of the device claim operation can be obtained from the claim workflow.

---

# Resource: intersight_appliance_device_claim
DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.
Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.
## Usage Example
### Resource Creation

```hcl
resource "intersight_appliance_device_claim" "appliance_device_claim1" {
  device_id      = "WZP23350KMU"
  hostname       = "10.106.233.221"
  message        = "Endpoint claimed successfully"
  platform_type  = "IMC"
  request_id     = "718"
  security_token = "8707352528DC"
  status         = "completed"
  username       = "admin"
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
}

variable "account" {
  type        = string
  description = "Moid of iam.Account Mo"
}
```
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) -(ReadOnly) A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `device_id`:(string)(ReadOnly) Device identifier of the endpoint device. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `hostname`:(string) Hostname or IP address of the endpoint device the user wants to claim. 
* `is_password_set`:(bool)(ReadOnly) Indicates whether the value of the 'password' property has been set. 
* `is_renew`:(bool) Tracks if this device is to be claimed or certificate renewal. 
* `message`:(string)(ReadOnly) Message set by the device claim process. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `password`:(string) Password to be used to login to the endpoint device. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `platform_type`:(string) Platform type of the endpoint device.* `` - The device reported an empty or unrecognized platform type.* `APIC` - An Application Policy Infrastructure Controller cluster.* `CAPIC` - An Application Policy Infrastructure Controller cloud instance.* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.* `IMC` - A standalone UCS Server Integrated Management Controller.* `IMCM4` - A standalone UCS M4 Server.* `IMCM5` - A standalone UCS M5 server.* `IMCRack` - A standalone UCS M6 and above server.* `UCSIOM` - An UCS Chassis IO module.* `HX` - A HyperFlex storage controller.* `HyperFlexAP` - A HyperFlex Application Platform.* `IWE` - An Intersight Workload Engine.* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.* `IntersightAssist` - A Cisco Intersight Assist.* `PureStorageFlashArray` - A Pure Storage FlashArray device.* `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on.* `ACISwitch` - A platform type to support ACI Switches.* `NexusSwitch` - A platform type to support Cisco Nexus Switches.* `MDSDevice` - A platform type to support MDS devices.* `UCSC890` - A standalone Cisco UCSC890 server.* `NetAppOntap` - A NetApp ONTAP storage system.* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.* `EmcScaleIo` - An EMC ScaleIO storage system.* `EmcVmax` - An EMC VMAX storage system.* `EmcVplex` - An EMC VPLEX storage system.* `EmcXtremIo` - An EMC XtremIO storage system.* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.* `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines.* `AppDynamics` - An AppDynamics controller that monitors applications.* `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation.* `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation.* `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks.* `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.* `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.* `MicrosoftSqlServer` - A Microsoft SQL database server.* `MySqlServer` - An instance of either Oracle MySQL Database or the open source MariaDB.* `Kubernetes` - A Kubernetes cluster that runs containerized applications.* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.* `DellCompellent` - A Dell Compellent storage system.* `HPE3Par` - A HPE 3PAR storage system.* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.* `IMCBlade` - An Intersight managed UCS Blade Server.* `TerraformCloud` - A Terraform Cloud account.* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.* `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow.* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token.* `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow.* `CiscoCatalyst` - A Cisco Catalyst networking switch device.* `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely. 
* `request_id`:(string) User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set. 
* `reservation`:(HashMap) - A reference to a resourceReservation resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `security_token`:(string)(ReadOnly) Device security token of the endpoint device. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string)(ReadOnly) Status of the device claim process.* `started` - Device claim operation has started.* `failed` - Device claim operation has failed.* `completed` - Device claim operation has completed. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `username`:(string) Username to log in to the endpoint device. 
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
`intersight_appliance_device_claim` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_appliance_device_claim.example 1234567890987654321abcde
``` 
