---
layout: "thunder"
page_title: "thunder: thunder_rib_route"
sidebar_current: "docs-thunder-resource-rib-route"
description: |-
	Provides details about thunder rib route resource for A10
---

# thunder\_rib\_route

`thunder_rib_route` Provides details about thunder rib route 
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_rib_route" "Ip_Route_Rib_Test" {
ip-nexthop-lif {   
        description_nexthop_lif =  "string" 
        lif =  0 
        }
ip-nexthop-ipv4 {   
        description_nexthop_ip =  "string" 
        ip_next_hop =  "string" 
        distance_nexthop_ip =  0 
        }
uuid = "string"
ip_dest_addr = "string"
ip-nexthop-tunnel {   
        tunnel =  0 
        ip_next_hop_tunnel =  "string" 
        distance_nexthop_tunnel =  0 
        description_nexthop_tunnel =  "string" 
        }
ip-nexthop-partition {   
        partition_name =  "string" 
        vrid_num_in_partition =  0 
        description_nexthop_partition =  "string" 
        description_partition_vrid =  "string" 
        }
ip_mask = "string"
 
}

```

## Argument Reference

* `ip_dest_addr` - Destination prefix
* `ip_mask` - Destination prefix mask
* `uuid` - uuid of the object
* `description_nexthop_lif` - Description for static route
* `lif` - LIF Interface (Logical tunnel interface number)
* `description_nexthop_ip` - Description for static route
* `distance_nexthop_ip` - Distance value for this route
* `ip_next_hop` - Forwarding router’s address
* `description_nexthop_tunnel` - Description for static route
* `distance_nexthop_tunnel` - Distance value for this route
* `ip_next_hop_tunnel` - Forwarding router’s address
* `tunnel` - Tunnel interface (Tunnel interface number)
* `description_nexthop_partition` - Description for static route
* `description_partition_vrid` - Description for static route
* `partition_name` - Name of network partition
* `vrid_num_in_partition` - Specify ha VRRP-A vrid
