---
layout: "thunder"
page_title: "thunder: thunder_slb_template_monitor"
sidebar_current: "docs-thunder-resource-slb-template-monitor"
description: |-
	Provides details about thunder slb template monitor resource for A10
---

# thunder\_slb\_template\_monitor

`thunder_slb_template_monitor` Provides details about thunder slb template monitor
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_monitor" "Slb_Template_Monitor_Test" {

clear-cfg {   
        clear_sequence =  0 
        clear_all_sequence =  0 
        sessions =  "string" 
        }
uuid = "string"
link-enable-cfg {   
        ena_sequence =  0 
        enaeth =  0 
        }
link-up-cfg {   
        linkup_ethernet3 =  0 
        linkup_ethernet2 =  0 
        linkup_ethernet1 =  0 
        link_up_sequence1 =  0 
        link_up_sequence3 =  0 
        link_up_sequence2 =  0 
        }
link-down-cfg {   
        link_down_sequence1 =  0 
        link_down_sequence2 =  0 
        link_down_sequence3 =  0 
        linkdown_ethernet2 =  0 
        linkdown_ethernet3 =  0 
        linkdown_ethernet1 =  0 
        }
user_tag = "string"
link-disable-cfg {   
        dis_sequence =  0 
        diseth =  0 
        }
monitor_relation = "string"
id = 0
 
}
```

## Argument Reference

* `id` - Monitor template ID Number
* `monitor_relation` - ‘monitor-and’: Configures the monitors in current template to work with AND logic; ‘monitor-or’: Configures the monitors in current template to work with OR logic;
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `clear_all_sequence` - Sequence number (Specify the port physical port number)
* `clear_sequence` - Specify the port physical port number
* `sessions` - ‘all’: Clear all sessions; ‘sequence’: Sequence number;
* `ena_sequence` - Sequence number (Specify the physical port number)
* `enaeth` - Specify the physical port number (Ethernet interface number)
* `link_up_sequence1` - Sequence number (Specify the port physical port number)
* `link_up_sequence2` - Sequence number (Specify the port physical port number)
* `link_up_sequence3` - Sequence number (Specify the port physical port number)
* `linkup_ethernet1` - Specify the port physical port number (Ethernet interface number)
* `linkup_ethernet2` - Specify the port physical port number (Ethernet interface number)
* `linkup_ethernet3` - Specify the port physical port number (Ethernet interface number)
* `link_down_sequence1` - Sequence number (Specify the port physical port number)
* `link_down_sequence2` - Sequence number (Specify the port physical port number)
* `link_down_sequence3` - Sequence number (Specify the port physical port number)
* `linkdown_ethernet1` - Specify the port physical port number (Ethernet interface number)
* `linkdown_ethernet2` - Specify the port physical port number (Ethernet interface number)
* `linkdown_ethernet3` - Specify the port physical port number (Ethernet interface number)
* `dis_sequence` - Sequence number (Specify the physical port number)
* `diseth` - Specify the physical port number (Ethernet interface number)
