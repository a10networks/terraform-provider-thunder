---
layout: "vthunder"
page_title: "vthunder: vthunder_interface_ethernet_trunk_group"
sidebar_current: "docs-vthunder-resource-interface-ethernet-trunk-group"
description: |-
	Provides details about vthunder interface ethernet trunk group resource for A10
---

# vthunder\_interface\_ethernet\_trunk\_group

`vthunder_interface_ethernet_trunk_group` Provides details about vthunder interface ethernet trunk group
## Example Usage

## Argument Reference

* `admin_key` - LACP admin key (Admin key value)
* `mode` - ‘active’: enable initiation of LACP negotiation on a port(default); ‘passive’: disable initiation of LACP negotiation on a port;
* `port_priority` - Set LACP priority for a port (LACP port priority)
* `timeout` - ‘long’: Set LACP long timeout (default); ‘short’: Set LACP short timeout;
* `trunk_number` - Trunk Number
* `type` - ‘static’: Static (default); ‘lacp’: lacp; ‘lacp-udld’: lacp-udld;
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `fast` - fast timeout in unit of milli-seconds(default 1000)
* `slow` - slow timeout in unit of seconds
