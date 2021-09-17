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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_udp" "resourceSlbTemplateUdpTest" {
	name = "string"
idle_timeout = 0
qos = 0
stateless_conn_timeout = 0
immediate = 0
short = 0
age = 0
re_select_if_server_down = 0
disable_clear_session = 0
radius_lb_method_hash_type = "string"
avp = "string"
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `udp` - L4 UDP switch config
* `name` - Fast UDP Template Name
* `idle-timeout` - Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)
* `qos` - QOS level (number)
* `stateless-conn-timeout` - Stateless Current Connection Timeout value (5 - 120 seconds) (idle timeout in second, default 120)
* `immediate` - Immediate Removal after Transaction
* `short` - Short lived session
* `age` - short age (in sec), default is 31
* `re-select-if-server-down` - re-select another server if service port is down
* `disable-clear-session` - Disable immediate clearing of session
* `radius-lb-method-hash-type` - 'ip': IP-Hash;
* `avp` - '4': NAS-IP-address; '8': Framed-IP-Address;
* `uuid` - uuid of the object
* `user-tag` - Customized tag

