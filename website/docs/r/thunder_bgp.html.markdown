---
layout: "thunder"
page_title: "thunder: thunder_bgp"
sidebar_current: "docs-thunder-resource-bgp"
description: |-
    Provides details about thunder bgp resource for A10
---

# thunder\_bgp

`thunder_bgp` Provides details about thunder bgp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_bgp" "resourceBgpTest" {
	extended_asn_cap = 0
nexthop_trigger {  
 	enable =  0 
	delay =  0 
	}
uuid = "string"
 
}

```

## Argument Reference

* `bgp` - Border Gateway Protocol (BGP)
* `extended-asn-cap` - Enable the router to send 4-octet ASN capabilities
* `enable` - Enable the nexthop tracking functionality
* `delay` - Configure nexthop trigger delay time interval (Nexthop Trigger Delay time interval between 1 and 100)
* `uuid` - uuid of the object

