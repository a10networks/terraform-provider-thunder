---
layout: "thunder"
page_title: "thunder: thunder_slb_template_tcp"
sidebar_current: "docs-thunder-resource-slb-template-tcp"
description: |-
	Provides details about thunder slb template tcp resource for A10
---

# thunder\_slb\_template\_tcp

`thunder_slb_template_tcp` Provides details about thunder slb template tcp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_tcp" "Slb_Template_Tcp_Test" {

del_session_on_server_down = 0
initial_window_size = 0
half_open_idle_timeout = 0
logging = "string"
name = "string"
reset_fwd = 0
reset_follow_fin = 0
alive_if_active = 0
idle_timeout = 0
force_delete_timeout = 0
user_tag = "string"
down = 0
disable = 0
reset_rev = 0
insert_client_ip = 0
lan_fast_ack = 0
half_close_idle_timeout = 0
force_delete_timeout_100ms = 0
qos = 0
uuid = "string"
 
}
```

## Argument Reference

* `alive_if_active` - keep connection alive if active traffic
* `del_session_on_server_down` - Delete session if the server/port goes down (either disabled/hm down)
* `disable` - send reset to client when server is disabled
* `down` - send reset to client when server is down
* `force_delete_timeout` - The maximum time that a session can stay in the system before being delete (number (second))
* `force_delete_timeout_100ms` - The maximum time that a session can stay in the system before being delete (number in 100ms)
* `half_close_idle_timeout` - TCP Half Close Idle Timeout (sec), default off (half close idle timeout in second, default off)
* `half_open_idle_timeout` - TCP Half Open Idle Timeout (sec), default off (half open idle timeout in second, default off)
* `idle_timeout` - Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)
* `initial_window_size` - Set the initial window size (number)
* `insert_client_ip` - Insert client ip into TCP option
* `lan_fast_ack` - Enable fast TCP ack on LAN
* `logging` - ‘init’: init only log; ‘term’: termination only log; ‘both’: both initial and termination log;
* `name` - Fast TCP Template Name
* `qos` - QOS level (number)
* `reset_follow_fin` - send reset to client or server upon receiving first fin
* `reset_fwd` - send reset to server if error happens
* `reset_rev` - send reset to client if error happens
* `user_tag` - Customized tag
* `uuid` - uuid of the object
