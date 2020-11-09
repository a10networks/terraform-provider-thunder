---
layout: "thunder"
page_title: "thunder: thunder_vrrp_peer_group"
sidebar_current: "docs-thunder-resource-vrrp-peer-group"
description: |-
	Provides details about thunder vrrp peer group resource for A10
---

# thunder\_vrrp\_peer\_group

`thunder_vrrp_peer_group` Provides details about thunder vrrp peer group
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_vrrp_peer_group" "Vrrp_Peer_Group_Test" {

peer {  
 ip-peer-address-cfg {   
        ip_peer_address =  "string" 
        }
ipv6-peer-address-cfg {   
        ipv6_peer_address =  "string" 
        }
        }
uuid = "string"
 
}
```

## Argument Reference

* `uuid` - uuid of the object
* `ip_peer_address` - IP Address
* `ipv6_peer_address` - IPV6 address
