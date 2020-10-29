---
layout: "thunder"
page_title: "thunder: thunder_fw_tap_monitor"
sidebar_current: "docs-thunder-resource-fw-tap-monitor"
description: |-
	Provides details about thunder fw tap monitor resource for A10
---

# thunder\_fw\_tap\_monitor

`thunder_fw_tap_monitor` Provides details about thunder fw tap monitor
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_tap_monitor" "Fw_Tap_Monitor_Test" {
        status = "string"
tap-port-cfg {   
        tap_eth =  0 
        tap_vlan =  0 
        }
uuid = "string"
 
}
```

## Argument Reference

* `status` - ‘enable’: Enable tap monitor mode; ‘disable’: Disable tap monitor mode (Default:Disable);
* `uuid` - uuid of the object
* `tap_eth` - Ethernet interface number
* `tap_vlan` - Vlan number

