---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_community_read_oid"
sidebar_current: "docs-thunder-resource-snmp-server-community-read-oid"
description: |-
	Provides details about thunder snmp server community read oid resource for A10
---

# thunder\_snmp\_server\_community\_read\_oid

`thunder_snmp_server_community_read_oid` Provides details about thunder snmp server community read oid
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_snmp_server_community_read_oid" "Snmp_Server_Community_Read_Oid_Test" {

 remote {  
 host_list {   
        dns_host =  "string" 
        ipv4_mask =  "string" 
        }
ipv4_list {   
        ipv4_host =  "string" 
        ipv4_mask =  "string" 
        }
ipv6_list {   
        ipv6_host =  "string" 
        ipv6_mask =  0 
        }
        }
oid_val = "string"
user_tag = "string"
uuid = "string"
 
}
```

## Argument Reference

* `oid_val` - specific the oid (The oid value, object-key)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `dns_host` - DNS remote host
* `ipv4_mask` - IPV4 mask
* `ipv4_host` - IPV4 remote host
* `ipv6_host` - IPV6 remote host
* `ipv6_mask` - IPV6 mask

