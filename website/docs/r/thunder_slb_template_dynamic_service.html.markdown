---
layout: "thunder"
page_title: "thunder: thunder_slb_template_dynamic_service"
sidebar_current: "docs-thunder-resource-slb-template-dynamic-service"
description: |-
	Provides details about thunder slb template dynamic service resource for A10
---

# thunder\_slb\_template\_dynamic\_service

`thunder_slb_template_dynamic_service` Provides details about thunder slb template dynamic service
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_dynamic_service" "Slb_Template_Dynamic_Service_Test" {

dns-server {   
        ipv6_dns_server =  "string" 
        ipv4_dns_server =  "string" 
        }
name = "string"
user_tag = "string"
uuid = "string"
 
}
```

## Argument Reference

* `name` - Dynamic Service Template Name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `ipv4_dns_server` - DNS Server IPv4 Address
* `ipv6_dns_server` - DNS Server IPv6 Address
