### Resource Creation

```hcl
resource "intersight_ippool_pool" "ippool_pool1" {
  name             = "ippool_pool1"
  description      = "ippool pool"
  assignment_order = "sequential"

  ip_v4_config {
    object_type = "ippool.IpV4Config"
    gateway     = "10.1.1.1"
    netmask     = "255.0.0.0"
    primary_dns = "8.8.8.8"
  }

  # Specify either "size" or "to", but not both!
  ip_v4_blocks {
    from = "10.1.1.100"
    size = 100
    # to   = "10.1.1.199"
  }

  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "ippool_ip_v4_config" {
  type = string
  description = " value for moid"
}

variable "organization" {
  type = string
  description = "value for organization"
}
```