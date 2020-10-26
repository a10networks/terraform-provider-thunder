---
layout: "thunder"
page_title: "thunder: thunder_slb_template_http"
sidebar_current: "docs-thunder-resource-slb_template_http"
description: |-
    Provides details about thunder SLB template http resource for A10
---

# thunder\_slb\_template\_http

`thunder_slb_template_http` Provides details about thunder SLB template http
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_http" "http" {
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