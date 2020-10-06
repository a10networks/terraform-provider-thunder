---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_player_id_global"
sidebar_current: "docs-vthunder-resource-slb-player-id-global"
description: |-
    Provides details about vthunder SLB player id global resource for A10
---

# vthunder\_slb\_player\_id\_global

`vthunder_slb_player_id_global` Provides details about vthunder SLB player id global
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_player_id_global" "player_id" {
	enforcement_timer = 10
	abs_max_expiration = 10
	sampling_enable  {
	    counters1 = "all"
	}
	force_passive = 1
	pkt_activity_expiration = 10
	min_expiration = 10
	enable_64bit_player_id = 0 
}
```

## Argument Reference

* `enforcement_timer` - Time to playerid enforcement after bootup (default 480 seconds) (Time to playerid enforcement in seconds, default 480)
* `abs_max_expiration` - Absolute max record expiration value (default 10 minutes) (Absolute max record expiration time in minutes, default 10)
* `sampling_enable`
    * `counters1` - 'all’: all; 'total_playerids_created’: Playerid records created; 'total_playerids_deleted’: Playerid records deleted; 'total_abs_max_age_outs’: Playerid records max time aged out; 'total_pkt_activity_age_outs’: Playerid records idle timeout; 'total_invalid_playerid_pkts’: Invalid playerid packets; 'total_invalid_playerid_drops’: Invalid playerid packet drops; 'total_valid_playerid_pkts’: Valid playerid packets;
* `force_passive` - Forces the device to be in passive mode (Only stats and no packet drops)
* `pkt_activity_expiration` - Packet activity record expiration value (default 5 minutes) (Packet activity record expiration time in minutes, default 5)
* `min_expiration` - Minimum record expiration value (default 1 min) (Min record expiration time in minutes, default 1)
* `enable_64bit_player_id` - Enable 64 bit player id check. Default is 32 bit

