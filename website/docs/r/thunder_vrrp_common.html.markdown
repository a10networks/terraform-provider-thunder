---
layout: "thunder"
page_title: "thunder: thunder_vrrp_common"
sidebar_current: "docs-thunder-resource-vrrp-common"
description: |-
	Provides details about thunder vrrp common resource for A10
---

# thunder\_vrrp\_common

`thunder_vrrp_common` Provides details about thunder vrrp common
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_vrrp_common" "Vrrp_Common_Test" {

forward_l4_packet_on_standby = 0
get_ready_time = 0
hello_interval = 0
uuid = "string"
preemption_delay = 0
set_id = 0
device_id = 0
arp_retry = 0
dead_timer = 0
disable_default_vrid = 0
enable_sync_session_seq_number = 0
track_event_delay = 0
action = "string"
hostid_append_to_vrid {  
        hostid_append_to_vrid_value =  0 
        hostid_append_to_vrid_default =  0 
        }
restart_time = 0
inline_mode_cfg {  
        inline_mode =  0 
        preferred_trunk =  0 
        preferred_port =  0 
        }
 
}
```

## Argument Reference

* `action` - ‘enable’: enable vrrp-a; ‘disable’: disable vrrp-a;
* `arp_retry` - Number of additional gratuitous ARPs sent out after HA failover (1-255, default is 4)
* `dead_timer` - VRRP-A dead timer in terms of how many hello messages missed, default is 5 (2-255, default is 5)
* `device_id` - Unique ID for each VRRP-A box (Device-id number)
* `disable_default_vrid` - Disable default vrid
* `forward_l4_packet_on_standby` - Enables Layer 2/3 forwarding of Layer 4 traffic on the Standby ACOS device
* `get_ready_time` - set get ready time after ax starting up (60-1200, in unit of 100millisec, default is 60)
* `hello_interval` - VRRP-A Hello Interval (1-255, in unit of 100millisec, default is 2)
* `preemption_delay` - Delay before changing state from Active to Standby (1-255, in unit of 100millisec, default is 60)
* `restart_time` - Time between restarting ports on standby system after transition
* `set_id` - Set-ID for HA configuration (Set id from 1 to 15)
* `track_event_delay` - Delay before changing state after up/down event (Units of 100 milliseconds (default 30))
* `uuid` - uuid of the object
* `hostid_append_to_vrid_default` - hostid append to vrid default
* `hostid_append_to_vrid_value` - hostid append to vrid num
* `inline_mode` - Enable Layer 2 Inline Hot Standby Mode
* `preferred_port` - Preferred ethernet Port
* `preferred_trunk` - Preferred trunk Port
