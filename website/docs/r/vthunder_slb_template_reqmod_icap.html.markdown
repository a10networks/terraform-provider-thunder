---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_reqmod_icap"
sidebar_current: "docs-vthunder-resource-slb-template-reqmod-icap"
description: |-
    Provides details about vthunder slb template reqmod-icap resource for A10
---

# vthunder\_slb\_template\_reqmod\_icap

`vthunder_slb_template_reqmod_icap` provides details about slb template reqmod-icap
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_reqmod_icap" "reqmod_icap" {
	name = "testreqmodicap"
	user_tag = "test_tag"
	min_payload_size = 0
	allowed_http_methods = "POST"
	service_url = "10.0.0.11"
	preview = 200
	disable_http_server_reset = 0
	fail_close = 0
	include_protocol_in_uri = 0
	log_only_allowed_method = 0
	cylance = 0
	 
}
```

## Argument Reference

* `name` - Reqmod ICAP Template Name
* `user_tag` - Customized tag
* `min_payload_size` - min-payload-size value 0 - 65535, default is 0
* `allowed_http_methods` - List of allowed HTTP methods. Default is "Allow All". (List of HTTP methods allowed (default “Allow All”))
* `service_url` - URL to send to ICAP server (Service URL Name)
* `preview` - Preview value 1 - 32768, default is 32768
* `disable_http_server_reset` - Don’t reset http server
* `fail_close` - When template sg is down mark vport down
* `include_protocol_in_uri` - Include protocol and port in HTTP URI
* `log_only_allowed_method` - Only log allowed HTTP method
* `cylance` - cylance external server


