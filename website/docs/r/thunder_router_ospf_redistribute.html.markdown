---
layout: "thunder"
page_title: "thunder: thunder_router_ospf_redistribute"
sidebar_current: "docs-thunder-resource-router-ospf-redistribute"
description: |-
    Provides details about thunder router ospf redistribute resource for A10
---

# thunder\_router\_ospf\_redistribute

`thunder_router_ospf_redistribute` Provides details about thunder router ospf redistribute
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_ospf_redistribute" "resourceRouterOspfRedistributeTest" {
	redist-list {   
	type =  "string" 
	metric =  0 
	metric_type =  "string" 
	route_map =  "string" 
	tag =  0 
	}
ospf-list {   
	ospf =  0 
	process_id =  0 
	metric_ospf =  0 
	metric_type_ospf =  "string" 
	route_map_ospf =  "string" 
	tag_ospf =  0 
	}
ip_nat = 0
metric_ip_nat = 0
metric_type_ip_nat = "string"
route_map_ip_nat = "string"
tag_ip_nat = 0
ip-nat-floating-list {   
	ip_nat_prefix =  "string" 
	ip_nat_floating_IP_forward =  "string" 
	}
vip-list {   
	type_vip =  "string" 
	metric_vip =  0 
	metric_type_vip =  "string" 
	route_map_vip =  "string" 
	tag_vip =  0 
	}
vip-floating-list {   
	vip_address =  "string" 
	vip_floating_IP_forward =  "string" 
	}
uuid = "string"
 
}

```

## Argument Reference

* `redistribute` - Redistribute information from another routing protocol
* `type` - 'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;
* `metric` - OSPF default metric (OSPF metric)
* `metric-type` - '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
* `route-map` - Route map reference (Pointer to route-map entries)
* `tag` - Set tag for routes redistributed into OSPF (32-bit tag value)
* `ospf` - Open Shortest Path First (OSPF)
* `process-id` - OSPF process ID
* `metric-ospf` - OSPF default metric (OSPF metric)
* `metric-type-ospf` - '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
* `route-map-ospf` - Route map reference (Pointer to route-map entries)
* `tag-ospf` - Set tag for routes redistributed into OSPF (32-bit tag value)
* `ip-nat` - IP-NAT
* `metric-ip-nat` - OSPF default metric (OSPF metric)
* `metric-type-ip-nat` - '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
* `route-map-ip-nat` - Route map reference (Pointer to route-map entries)
* `tag-ip-nat` - Set tag for routes redistributed into OSPF (32-bit tag value)
* `ip-nat-prefix` - Address
* `ip-nat-floating-IP-forward` - Floating-IP as forward address
* `type-vip` - 'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;
* `metric-vip` - OSPF default metric (OSPF metric)
* `metric-type-vip` - '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
* `route-map-vip` - Route map reference (Pointer to route-map entries)
* `tag-vip` - Set tag for routes redistributed into OSPF (32-bit tag value)
* `vip-address` - Address
* `vip-floating-IP-forward` - Floating-IP as forward address
* `uuid` - uuid of the object

