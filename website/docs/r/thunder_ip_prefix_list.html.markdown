---
layout: "thunder"
page_title: "thunder: thunder_ip_prefix_list"
sidebar_current: "docs-thunder-resource-ip-prefix-list"
description: |-
    Provides details about thunder ip prefix list resource for A10
---

# thunder\_ip\_prefix\_list

`thunder_ip_prefix_list` Provides details about thunder ip prefix list
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_ip_prefix_list" "resourceIpPrefixListTest" {
	name = "string"
rules {   
	seq =  0 
	description =  "string" 
	action =  "string" 
	any =  0 
	ipaddr =  "string" 
	ge =  0 
	le =  0 
	}
uuid = "string"
 
}

```

## Argument Reference

* `prefix-list` - Configure Prefix-list
* `name` - Name of a prefix list
* `seq` - Sequence number of an entry
* `description` - Prefix-list specific description (Up to 80 characters describing this prefix-list)
* `action` - 'deny': Specify packets to reject; 'permit': Specify packets to forward;
* `any` - Any prefix match. Same as "0.0.0.0/0 le 32"
* `ipaddr` - IP prefix, e.g., 35.0.0.0/8
* `ge` - Minimum prefix length to be matched
* `le` - Maximum prefix length to be matched
* `uuid` - uuid of the object

