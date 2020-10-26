---
layout: "thunder"
page_title: "thunder: thunder_overlay_tunnel_options"
sidebar_current: "docs-thunder-resource-overlay_tunnel_options"
description: |-
    Provides details about thunder overlay tunnel options resource for A10
---

# thunder\_overlay\_tunnel\_options

`thunder_overlay_tunnel_options` Provides details about thunder overlay tunnel options
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_overlay_tunnel_options" "options" {
		tcp_mss_adjust_disable=1
		nvgre_disable_flow_id=1
		ip_dscp_preserve=1
		vxlan_dest_port=100
		uuid="uuid1"
		gateway_mac="00:0a:95:9d:68:16"
		nvgre_key_mode_lower24=1		
}
```

## Argument Reference

* `tcp_mss_adjust_disable` - Disable TCP MSS adjustment in SYN packet for tunnels
* `nvgre_disable_flow_id` - Disable Flow-ID computation for NVGRE
* `ip_dscp_preserve` - Copy DSCP bits from inner IP to outer IP header
* `vxlan_dest_port` - VXLAN UDP Destination Port (UDP Port Number (default 4789))
* `uuid` - uuid of the object
* `gateway_mac` - MAC to be used with Gateway segment Id (MAC Address for the Gateway segment)
* `nvgre_key_mode_lower24` - Use the lower 24-bits of the GRE key as the VSID
