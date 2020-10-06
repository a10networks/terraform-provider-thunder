---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_dynamic_service"
sidebar_current: "docs-vthunder-resource-slb-template-dynamic-service"
description: |-
    Provides details about vthunder slb template dynamic-service resource for A10
---

# vthunder\_slb\_template\_dynamic\_service

`vthunder_slb_template_dynamic_service` provides details about slb template dynamic-service
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_dynamic_service" "dynamic_service" {
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

