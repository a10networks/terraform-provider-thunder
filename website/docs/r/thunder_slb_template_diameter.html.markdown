---
layout: "thunder"
page_title: "thunder: thunder_slb_template_diameter"
sidebar_current: "docs-thunder-resource-slb-template-diameter"
description: |-
	Provides details about thunder slb template diameter resource for A10
---

# thunder\_slb\_template\_diameter

`thunder_slb_template_diameter` Provides details about thunder slb template diameter
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_dblb" "Slb_Template_Dblb_Test" {

avp_string = "string"
terminate_on_cca_t = 0
message-code-list {   
        message_code =  0 
        }
avp-list {   
        int32 =  0 
        mandatory =  0 
        string =  "string" 
        avp =  0 
        int64 =  0 
        }
service_group_name = "string"
uuid = "string"
idle_timeout = 0
customize_cea = 0
product_name = "string"
dwr_up_retry = 0
forward_unknown_session_id = 0
avp_code = 0
vendor_id = 0
session_age = 0
load_balance_on_session_id = 0
name = "string"
dwr_time = 0
user_tag = "string"
origin_realm = "string"
origin_host {  
        uuid =  "string" 
        origin_host_name =  "string" 
        }
multiple_origin_host = 0
forward_to_latest_server = 0
 
}

```

## Argument Reference

* `avp_code` - avp code
* `avp_string` - pattern to be matched in the avp string name, max length 127 bytes
* `customize_cea` - customizing cea response
* `dwr_time` - dwr health-check timer interval (in 100 milli second unit, default is 100, 0 means unset this option)
* `dwr_up_retry` - number of successful dwr health-check before declaring target up
* `forward_to_latest_server` - Forward client message to the latest server that sends message with the same session id
* `forward_unknown_session_id` - Forward server message even it has unknown session id
* `idle_timeout` -  user sesison idle timeout (in minutes, default is 5)
* `load_balance_on_session_id` - Load balance based on the session id
* `multiple_origin_host` - allowing multiple origin-host to a single server
* `name` - diameter template Name
* `origin_realm` - origin-realm name avp
* `product_name` - product name avp
* `service_group_name` - service group name, this is the service group that the message needs to be copied to
* `session_age` - user session age allowed (default 10), this is not idle-time (in minutes)
* `terminate_on_cca_t` - remove diameter session when receiving CCA-T message
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `vendor_id` - vendor-id avp (Vendon Id)
* `avp` - customize avps for cer to the server (avp number)
* `int32` - 32 bits integer
* `int64` - 64 bits integer
* `mandatory` - mandatory avp
* `string` - String (string name, max length 127 bytes)
* `origin_host_name` - origin-host name avp
