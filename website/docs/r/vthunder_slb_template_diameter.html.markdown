---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_diameter"
sidebar_current: "docs-vthunder-resource-slb-template-diameter"
description: |-
    Provides details about vthunder slb template diameter resource for A10
---

# vthunder\_slb\_template\_diameter

`vthunder_slb_template_diameter` provides details about slb template diameter
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_diameter" "testname" {
	name = "testdiameter"
	user_tag = "test_tag"
	terminate_on_cca_t = 0
	message_code_list {
		message_code = 100
	}
	avp_list {
		int32 = 0
		mandatory = 0
		avp = 0
	}
	idle_timeout = 10
	customize_cea = 0
	product_name = "yyeuey"
	dwr_up_retry = 4
	forward_unknown_session_id = 0
	vendor_id = 0
	session_age = 50
	load_balance_on_session_id = 0
	dwr_time = 0
	origin_realm = "tttest"
	origin_host {
		origin_host_name = "testtt"
	}
	multiple_origin_host = 0
	forward_to_latest_server = 0
}
```

## Argument Reference

* `name` - diameter template Name
* `user_tag` - Customized tag
* `terminate_on_cca_t` - remove diameter session when receiving CCA-T message
* `message_code_list` - 
    * `message_code` - minimum: 1, maximum: 2147483647 
* `avp_list` -
    * `int32` -  32 bits integer
    * `mandatory` - mandatory avp
    * `avp` -  customize avps for cer to the server (avp number)
* `idle_timeout` - user sesison idle timeout (in minutes, default is 5) 
* `customize_cea` - customizing cea response
* `product_name` - product name avp
* `dwr_up_retry` - number of successful dwr health-check before declaring target up
* `forward_unknown_session_id` - Forward server message even it has unknown session id
* `vendor_id` - vendor-id avp (Vendor Id) 
* `session_age` - user session age allowed (default 10), this is not idle-time (in minutes)
* `load_balance_on_session_id` - Load balance based on the session id
* `dwr_time` - dwr health-check timer interval (in 100 milli second unit, default is 100, 0 means unset this option) 
* `origin_realm` - origin-realm name avp
* `origin_host` - 
    * `origin_host_name` - origin-host name avp
* `multiple_origin_host` - allowing multiple origin-host to a single server
* `forward_to_latest_server` - Forward client message to the latest server that sends message with the same session id


