---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_http"
sidebar_current: "docs-vthunder-resource-slb_template_http"
description: |-
    Provides details about vthunder SLB template http resource for A10
---

# vthunder\_slb\_template\_http

`vthunder_slb_template_http` Provides details about vthunder SLB template http
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_http" "testname" {
	name = "testhttp"
	user_tag = "test_tag"
	keep_client_alive = 1
	req_hdr_wait_time = 1
}
```

## Argument Reference

* `name` - HTTP Template Name
* `keep_client_alive` - Keep client alive
* `req_hdr_wait_time` - HTTP request header wait time before abort connection
* `user_tag` - Customized tag