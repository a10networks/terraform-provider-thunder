---
layout: "thunder"
page_title: "thunder: thunder_slb_template_dynamic_service"
sidebar_current: "docs-thunder-resource-slb-template-dynamic-service"
description: |-
    Provides details about thunder slb template dynamic-service resource for A10
---

# thunder\_slb\_template\_dynamic\_service

`thunder_slb_template_dynamic_service` provides details about slb template dynamic-service
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_dynamic_service" "dynamic_service" {
	name = "testdynamicservice"
	user_tag = "test_tag"
	dns_server {
		ipv4_dns_server = "10.0.0.10"
	}
}
```

## Argument Reference

* `name` - Dynamic Service Template Name
* `user_tag` - Customized tag
* `dns_server` - blade parameter priority
    * `ipv4_dns_server` - DNS Server IPv4 Address

