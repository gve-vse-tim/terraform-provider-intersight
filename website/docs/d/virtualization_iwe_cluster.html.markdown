---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_iwe_cluster"
description: |-
        A Intersight Workload Engine compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.

---

# Data Source: intersight_virtualization_iwe_cluster
A Intersight Workload Engine compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_iwe_cluster.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity_runway`:(int) The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%.Default value is math.MaxInt32 to indicate that the capacity runway is \ Unknown\  for a cluster that is not connected or with not sufficient data. 
* `cluster_name`:(string) The name of this HyperFlex cluster. 
* `compute_node_count`:(int) The number of compute nodes that belong to this cluster. 
* `configured_cpu_over_sub_factor`:(float) CPU oversubscription factor configured on the cluster. 
* `converged_node_count`:(int) The number of converged nodes that belong to this cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `current_cpu_over_sub_factor`:(float) Current oversubscription factor of the cluster. 
* `datacenter_name`:(string) Datacenter to which the cluster belongs. 
* `deployment_type`:(string) The deployment type of the HyperFlex cluster.The cluster can have one of the following configurations:1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,the deployment type is set as 'NA' (not available).* `NA` - The deployment type of the cluster is not available.* `Datacenter` - The deployment type of a cluster consisting of UCS Fabric Interconnect-attached nodes on the same site.* `Stretched Cluster` - The deployment type of a cluster consisting of UCS Fabric Interconnect-attached nodes across different sites.* `Edge` - The deployment type of a cluster consisting of 2 or more standalone nodes. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `drive_type`:(string) The type of the drives used for storage in this cluster.* `NA` - The drive type of the cluster is not available.* `All-Flash` - Indicates that this cluster contains flash drives only.* `Hybrid` - Indicates that this cluster contains both flash and hard disk drives. 
* `failure_reason`:(string) Reason for the failure when cluster is in failed state. 
* `hx_version`:(string) The HyperFlex Data or Application Platform version of this cluster. 
* `hypervisor_build`:(string) Hypervisor version of HyperFlex compute cluster along with build number. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform.* `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `hypervisor_version`:(string) Identifies the version of the hypervisor running on this cluster. Keeping the hypervisor version in the cluster makes it convenient for applications to validate their deployment needs. Defaults to an empty string. 
* `identity`:(string) The internally generated identity of this cluster. This entity is not manipulated by users. 
* `management_ip_address`:(string) Management IP Address of the cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user-provided name for this cluster to facilitate identification. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Cluster health status as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `storage_capacity`:(int) The storage capacity in this cluster. 
* `storage_node_count`:(int) The number of storage nodes that belong to this cluster. 
* `storage_utilization`:(float) The storage utilization is computed based on total capacity and current capacity utilization. 
* `total_cores`:(int) Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. 
* `utilization_percentage`:(float) The storage utilization percentage is computed based on total capacity and current capacity utilization. 
* `utilization_trend_percentage`:(float) The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data. 
 
