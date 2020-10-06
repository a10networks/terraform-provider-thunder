---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_respmod_icap"
sidebar_current: "docs-vthunder-resource-slb-template-respmod-icap"
description: |-
    Provides details about vthunder slb template respmod-icap resource for A10
---

# vthunder\_slb\_template\_respmod\_icap

`vthunder_slb_template_respmod_icap` provides details about slb template respmod-icap
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_respmod_icap" "respmod_icap" {
	name = "testrespmodicap"
	user_tag = "test_tag"
	min_payload_size = 30
	preview = 200
	disable_http_server_reset = 0
	fail_close = 0
	include_protocol_in_uri = 0
	log_only_allowed_method = 0
	cylance = 0

}
```

## Argument Reference

* `name` - Respmod ICAP Template Name
* `user_tag` - Customized tag
* `min_payload_size` - min-payload-size value 0 - 65535, default is 0
* `preview` - Preview value 1 - 32768, default is 32768
* `disable_http_server_reset` - Don’t reset http server
* `fail_close` - When template sg is down mark vport down
* `include_protocol_in_uri` - Include protocol and port in HTTP URI
* `log_only_allowed_method` - Only log allowed HTTP method
* `cylance` - cylance external server


