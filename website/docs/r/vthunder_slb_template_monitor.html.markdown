---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_monitor"
sidebar_current: "docs-vthunder-resource-slb_template_monitor"
description: |-
    Provides details about vthunder SLB template monitor resource for A10
---

# vthunder\_slb\_template\_monitor

`vthunder_slb_template_monitor` Provides details about vthunder SLB template monitor
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_monitor" "monitor" {
	user_tag = "test_tag"
	link_disable_cfg {
		dis_sequence = 1 
		diseth = 1
		}
	monitor_relation = "monitor-and"
	id2 = 1
}
```

## Argument Reference

* `monitor_relation` - 'monitor-and’: Configures the monitors in current template to work with AND logic; 'monitor-or’: Configures the monitors in current template to work with OR logic;
* `user_tag` - Customized tag
* `dis_sequence` - Sequence number (Specify the physical port number)
* `diseth` - Specify the physical port number (Ethernet interface number)
* `id2` - Monitor template ID Number




