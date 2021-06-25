---
layout: "thunder"
page_title: "thunder: thunder_slb_template_persist_source_ip"
sidebar_current: "docs-thunder-resource-slb-template-persist-source-ip"
description: |-
	Provides details about thunder slb template persist source ip resource for A10
---

# thunder\_slb\_template\_persist\_source\_ip

`thunder_slb_template_persist_source_ip` Provides details about thunder slb template persist source ip
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_persist_source_ip" "Slb_Template_Persist_Source_Ip_Test" {

netmask6 = 0
incl_dst_ip = 0
hash_persist = 0
name = "string"
enforce_higher_priority = 0
dont_honor_conn_rules = 0
primary_port = 0
user_tag = "string"
server = 0
service_group = 0
timeout = 0
scan_all_members = 0
netmask = "string"
incl_sport = 0
match_type = 0
uuid = "string"
 
}
```

## Argument Reference

* `dont_honor_conn_rules` - Do not observe connection rate rules
* `enforce_higher_priority` - Enforce to use high priority node if available
* `hash_persist` - Use hash value of source IP address
* `incl_dst_ip` - Include destination IP on the persist
* `incl_sport` - Include source port on the persist
* `match_type` - Persistence type
* `name` - Source IP persistence template name
* `netmask` - IP subnet mask
* `netmask6` - IPV6 subnet mask
* `primary_port` - Primary port to create the persist session
* `scan_all_members` - Persist with SCAN of all members
* `server` - Persist to the same server, default is port
* `service_group` - Persist within the same service group
* `timeout` - Persistence timeout (in minutes)
* `user_tag` - Customized tag
* `uuid` - uuid of the object

