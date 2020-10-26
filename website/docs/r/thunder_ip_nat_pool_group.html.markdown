---
layout: "thunder"
page_title: "thunder: thunder_ip_nat_pool_group"
sidebar_current: "docs-thunder-resource-ip-nat-pool-group"
description: |-
	Provides details about thunder ip nat pool group resource for A10
---

# thunder\_ip\_nat\_pool\_group

`thunder_ip_nat_pool_group` Provides details about thunder ip nat pool group
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

// Put working JSON here
```

## Argument Reference

* `pool_group_name` - Specify pool group name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `vrid` - Specify VRRP-A vrid (Specify ha VRRP-A vrid)
* `pool_name` - Specify NAT pool name

