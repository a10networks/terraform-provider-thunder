---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_SNMPv1_v2c_user"
sidebar_current: "docs-thunder-resource-snmp-server-SNMPv1-v2c-user"
description: |-
	Provides details about thunder snmp server SNMPv1 v2c user resource for A10
---

# thunder\_snmp\_server\_SNMPv1\_v2c\_user

`thunder_snmp_server_SNMPv1_v2c_user` Provides details about thunder snmp server SNMPv1 v2c user
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_snmp_server_snmpv1_v2c_user" "Snmp_Server_SNMPv1_V2c_User_Test" {
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
uuid = "string"
passwd = "string"
encrypted = "Unknown Type: encrypted"
user_tag = "string"
user = "string"
oid_list {   
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
        oid_val =  "string" 
        user_tag =  "string" 
        uuid =  "string" 
        }
 
}


```

## Argument Reference

* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED community string)
* `passwd` - define value of community string
* `user` - SNMPv1/v2c community configuration key
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `dns_host` - DNS remote host
* `ipv4_mask` - IPV4 mask
* `ipv4_host` - IPV4 remote host
* `ipv6_host` - IPV6 remote host
* `ipv6_mask` - IPV6 mask
* `oid_val` - specific the oid (The oid value, object-key)

