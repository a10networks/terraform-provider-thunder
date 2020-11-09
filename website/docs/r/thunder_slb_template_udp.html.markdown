---
layout: "thunder"
page_title: "thunder: thunder_slb_template_udp"
sidebar_current: "docs-thunder-resource-slb-template-udp"
description: |-
	Provides details about thunder slb template udp resource for A10
---

# thunder\_slb\_template\_udp

`thunder_slb_template_udp` Provides details about thunder slb template udp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_udp" "Slb_Template_Udp_Test" {

short = 0
qos = 0
name = "string"
age = 0
stateless_conn_timeout = 0
idle_timeout = 0
re_select_if_server_down = 0
immediate = 0
user_tag = "string"
disable_clear_session = 0
uuid = "string"
 
}
```

## Argument Reference

* `age` - short age (in sec), default is 31
* `disable_clear_session` - Disable immediate clearing of session
* `idle_timeout` - Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)
* `immediate` - Immediate Removal after Transaction
* `name` - Fast UDP Template Name
* `qos` - QOS level (number)
* `re_select_if_server_down` - re-select another server if service port is down
* `short` - Short lived session
* `stateless_conn_timeout` - Stateless Current Connection Timeout value (5 - 120 seconds) (idle timeout in second, default 120)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
