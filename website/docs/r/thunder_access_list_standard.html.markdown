---
layout: "thunder"
page_title: "thunder: thunder_access_list_standard"
sidebar_current: "docs-thunder-resource-access-list-standard"
description: |-
    Provides details about thunder access list standard resource for A10
---

# thunder\_access\_list\_standard

`thunder_access_list_standard` Provides details about thunder access list standard
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_access_list_standard" "resourceAccessListStandardTest" {
	std = 0
stdrules {   
	seq_num =  0 
	std_remark =  "string" 
	action =  "string" 
	any =  0 
	host =  "string" 
	subnet =  "string" 
	rev_subnet_mask =  "string" 
	log =  0 
	transparent_session_only =  0 
	}
uuid = "string"
 
}

```

## Argument Reference

* `standard` - Configure Standard Access List
* `std` - IP standard access list
* `seq-num` - Sequence number
* `std-remark` - Access list entry comment (Notes for this ACL)
* `action` - 'deny': Deny; 'permit': Permit; 'l3-vlan-fwd-disable': Disable L3 forwarding between VLANs;
* `any` - Any source host
* `host` - A single source host (Host address)
* `subnet` - Source Address
* `rev-subnet-mask` - Network Mask 0=apply 255=ignore
* `log` - Log matches against this entry
* `transparent-session-only` - Only log transparent sessions
* `uuid` - uuid of the object

