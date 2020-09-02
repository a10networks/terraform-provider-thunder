---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_tap_monitor"
sidebar_current: "docs-vthunder-resource-fw-tap-monitor"
description: |-
	Provides details about vthunder fw tap monitor resource for A10
---

# vthunder\_fw\_tap\_monitor

`vthunder_fw_tap_monitor` Provides details about vthunder fw tap monitor
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

// Put working JSON here
```

## Argument Reference

* `status` - ‘enable’: Enable tap monitor mode; ‘disable’: Disable tap monitor mode (Default:Disable);
* `uuid` - uuid of the object
* `tap_eth` - Ethernet interface number
* `tap_vlan` - Vlan number

