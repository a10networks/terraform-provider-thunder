---
layout: "thunder"
page_title: "thunder: thunder_route_map"
sidebar_current: "docs-thunder-resource-route-map"
description: |-
    Provides details about thunder route map resource for A10
---

# thunder\_route\_map

`thunder_route_map` Provides details about thunder route map
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_route_map" "resourceRouteMapTest" {
	tag = "string"
action = "string"
sequence = 0
uuid = "string"
user_tag = "string"
match {  
 as_path {  
 	name =  "string" 
	}
community {  
 name_cfg {  
 	name =  "string" 
	exact_match =  0 
	}
	}
extcommunity {  
 extcommunity_l_name {  
 	name =  "string" 
	exact_match =  0 
	}
	}
group {  
 	group_id =  0 
	ha_state =  "string" 
	}
scaleout {  
 	cluster_id =  0 
	operational_state =  "string" 
	}
interface {  
 	ethernet =  0 
	loopback =  0 
	trunk =  0 
	ve =  0 
	tunnel =  0 
	}
local_preference {  
 	val =  0 
	}
origin {  
 	egp =  0 
	igp =  0 
	incomplete =  0 
	}
ip {  
 address {  
 	acl1 =  0 
	acl2 =  0 
	name =  "string" 
prefix_list {  
 	name =  "string" 
	}
	}
next_hop {  
 	acl1 =  0 
	acl2 =  0 
	name =  "string" 
prefix_list_1 {  
 	name =  "string" 
	}
	}
peer {  
 	acl1 =  0 
	acl2 =  0 
	name =  "string" 
	}
rib {  
 	exact =  "string" 
	reachable =  "string" 
	unreachable =  "string" 
	}
	}
ipv6 {  
 address_1 {  
 	name =  "string" 
prefix_list_2 {  
 	name =  "string" 
	}
	}
next_hop_1 {  
 	next_hop_acl_name =  "string" 
	v6_addr =  "string" 
	prefix_list_name =  "string" 
	}
peer_1 {  
 	acl1 =  0 
	acl2 =  0 
	name =  "string" 
	}
rib {  
 	exact =  "string" 
	reachable =  "string" 
	unreachable =  "string" 
	}
	}
metric {  
 	value =  0 
	}
route_type {  
 external {  
 	value =  "string" 
	}
	}
tag {  
 	value =  0 
	}
	uuid =  "string" 
	}
set {  
 ip {  
 next_hop {  
 	address =  "string" 
	}
	}
ddos {  
 	class_list_name =  "string" 
	class_list_cid =  0 
	zone =  "string" 
	}
ipv6 {  
 next_hop_1 {  
 	address =  "string" 
local {  
 	address =  "string" 
	}
	}
	}
level {  
 	value =  "string" 
	}
metric {  
 	value =  "string" 
	}
metric_type {  
 	value =  "string" 
	}
tag {  
 	value =  0 
	}
aggregator {  
 aggregator_as {  
 	asn =  0 
	ip =  "string" 
	}
	}
as_path {  
 	prepend =  "string" 
	num =  "string" 
	num2 =  "string" 
	}
	atomic_aggregate =  0 
comm_list {  
 	v_std =  0 
	delete =  0 
	v_exp =  0 
	v_exp_delete =  0 
	name =  "string" 
	name_delete =  0 
	}
	community =  "string" 
dampening_cfg {  
 	dampening =  0 
	dampening_half_time =  0 
	dampening_reuse =  0 
	dampening_supress =  0 
	dampening_max_supress =  0 
	dampening_penalty =  0 
	}
extcommunity {  
 rt {  
 	value =  "string" 
	}
soo {  
 	value =  "string" 
	}
	}
local_preference {  
 	val =  0 
	}
originator_id {  
 	originator_ip =  "string" 
	}
weight {  
 	weight_val =  0 
	}
origin {  
 	egp =  0 
	igp =  0 
	incomplete =  0 
	}
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `route-map` - Configure route-map
* `tag` - Route map tag
* `action` - 'permit': Route map permits set operations; 'deny': Route map denies set operations;
* `sequence` - Sequence to insert to/delete from existing route-map entry
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `name` - Community-list name
* `exact-match` - Do exact matching of ecommunities
* `group-id` - HA or VRRP-A group id
* `ha-state` - 'active': HA or VRRP-A in Active State; 'standby': HA or VRRP-A in Standby State;
* `cluster-id` - Scaleout Cluster-id
* `operational-state` - 'up': Scaleout is up and running; 'down': Scaleout is down or disabled;
* `ethernet` - Ethernet interface (Port number)
* `loopback` - Loopback interface (Port number)
* `trunk` - Trunk Interface (Trunk interface number)
* `ve` - Virtual ethernet interface (Virtual ethernet interface number)
* `tunnel` - Tunnel interface (Tunnel interface number)
* `val` - Preference value
* `egp` - remote EGP
* `igp` - local IGP
* `incomplete` - unknown heritage
* `acl1` - IPv6 access-list number
* `acl2` - IP access-list number (expanded range)
* `exact` - Exact match a prefix (Prefix IPv6 address)
* `reachable` - IPv6 address is reachable
* `unreachable` - IPv6 address is unreachable
* `next-hop-acl-name` - IPv6 access-list name
* `v6-addr` - IPv6 address of next hop
* `prefix-list-name` - IPv6 prefix-list name
* `value` - VPN extended community
* `address` - IPv6 address of next hop
* `class-list-name` - Class-List Name
* `class-list-cid` - Class-List Cid
* `zone` - Zone Name
* `asn` - AS number
* `ip` - IP address of aggregator
* `prepend` - Prepend to the as-path (AS number)
* `num` - AS number
* `num2` - AS number
* `atomic-aggregate` - BGP atomic aggregate attribute
* `v-std` - Community-list number (standard)
* `delete` - Delete matching communities
* `v-exp` - Community-list number (expanded)
* `v-exp-delete` - Delete matching communities
* `name-delete` - Delete matching communities
* `community` - BGP community attribute
* `dampening` - Enable route-flap dampening
* `dampening-half-time` - Reachability Half-life time for the penalty(minutes)
* `dampening-reuse` - Value to start reusing a route
* `dampening-supress` - Value to start suppressing a route
* `dampening-max-supress` - Maximum duration to suppress a stable route(minutes)
* `dampening-penalty` - Un-reachability Half-life time for the penalty(minutes)
* `originator-ip` - IP address of originator
* `weight-val` - Weight value

