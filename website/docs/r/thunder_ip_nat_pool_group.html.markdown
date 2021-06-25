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
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_ip_nat_pool_group" "Ip_Nat_Pool_Group_Test" {
member-list {   
        uuid =  "string" 
        pool_name =  "string" 
        }
pool_group_name = "string"
vrid = 0
user_tag = "string"
uuid = "string"
 
}
```

## Argument Reference

* `pool_group_name` - Specify pool group name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `vrid` - Specify VRRP-A vrid (Specify ha VRRP-A vrid)
* `pool_name` - Specify NAT pool name

