---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_community_read"
sidebar_current: "docs-thunder-resource-snmp-server-community-read"
description: |-
	Provides details about thunder snmp server community read resource for A10
---

# thunder\_snmp\_server\_community\_read

`thunder_snmp_server_community_read` Provides details about thunder snmp server community read
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_snmp_server_community_read" "Snmp_Server_Community_Read_Test" {
uuid = "string"
remote {  
 host-list {   
        dns_host =  "string" 
        ipv4_mask =  "string" 
        }
ipv4-list {   
        ipv4_host =  "string" 
        ipv4_mask =  "string" 
        }
ipv6-list {   
        ipv6_host =  "string" 
        ipv6_mask =  0 
        }
        }
oid-list {   
remote {  
 host-list {   
        dns_host =  "string" 
        ipv4_mask =  "string" 
        }
ipv4-list {   
        ipv4_host =  "string" 
        ipv4_mask =  "string" 
        }
ipv6-list {   
        ipv6_host =  "string" 
        ipv6_mask =  0 
        }
        }
        oid_val =  "string" 
        user_tag =  "string" 
        uuid =  "string" 
        }
user_tag = "string"
user = "string"
 
}

```

## Argument Reference

* `user` - SNMPv1/v2c community string
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `oid_val` - specific the oid (The oid value, object-key)
* `dns_host` - DNS remote host
* `ipv4_mask` - IPV4 mask
* `ipv4_host` - IPV4 remote host
* `ipv6_host` - IPV6 remote host
* `ipv6_mask` - IPV6 mask

