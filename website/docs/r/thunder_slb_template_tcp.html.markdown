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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_tcp" "resourceSlbTemplateTcpTest" {
	name = "string"
logging = "string"
idle_timeout = 0
half_open_idle_timeout = 0
half_close_idle_timeout = 0
initial_window_size = 0
force_delete_timeout = 0
force_delete_timeout_100ms = 0
alive_if_active = 0
qos = 0
insert_client_ip = 0
lan_fast_ack = 0
reset_fwd = 0
reset_rev = 0
reset_follow_fin = 0
disable = 0
down = 0
re_select_if_server_down = 0
del_session_on_server_down = 0
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `tcp` - L4 TCP switch config
* `name` - Fast TCP Template Name
* `logging` - 'init': init only log; 'term': termination only log; 'both': both initial and termination log;
* `idle-timeout` - Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)
* `half-open-idle-timeout` - TCP Half Open Idle Timeout (sec), default off (half open idle timeout in second, default off)
* `half-close-idle-timeout` - TCP Half Close Idle Timeout (sec), default off (half close idle timeout in second, default off)
* `initial-window-size` - Set the initial window size (number)
* `force-delete-timeout` - The maximum time that a session can stay in the system before being delete (number (second))
* `force-delete-timeout-100ms` - The maximum time that a session can stay in the system before being delete (number in 100ms)
* `alive-if-active` - keep connection alive if active traffic
* `qos` - QOS level (number)
* `insert-client-ip` - Insert client ip into TCP option
* `lan-fast-ack` - Enable fast TCP ack on LAN
* `reset-fwd` - send reset to server if error happens
* `reset-rev` - send reset to client if error happens
* `reset-follow-fin` - send reset to client or server upon receiving first fin
* `disable` - send reset to client when server is disabled
* `down` - send reset to client when server is down
* `re-select-if-server-down` - re-select another server if service port is down
* `del-session-on-server-down` - Delete session if the server/port goes down (either disabled/hm down)
* `uuid` - uuid of the object
* `user-tag` - Customized tag

