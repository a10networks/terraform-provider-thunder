---
layout: "thunder"
page_title: "thunder: thunder_router_bgp"
sidebar_current: "docs-thunder-resource-router-bgp"
description: |-
    Provides details about thunder router bgp resource for A10
---

# thunder\_router\_bgp

`thunder_router_bgp` Provides details about thunder router bgp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp" "resourceRouterBgpTest" {
	as_number = 0
aggregate-address-list {   
	aggregate_address =  "string" 
	as_set =  0 
	summary_only =  0 
	}
bgp {  
 	always_compare_med =  0 
bestpath_cfg {  
 	ignore =  0 
	compare_routerid =  0 
	remove_recv_med =  0 
	remove_send_med =  0 
	missing_as_worst =  0 
	}
dampening_cfg {  
 	dampening =  0 
	dampening_half_time =  0 
	dampening_reuse =  0 
	dampening_supress =  0 
	dampening_max_supress =  0 
	dampening_penalty =  0 
	route_map =  "string" 
	}
	local_preference_value =  0 
	deterministic_med =  0 
	enforce_first_as =  0 
	fast_external_failover =  0 
	log_neighbor_changes =  0 
	nexthop_trigger_count =  0 
	router_id =  "string" 
	override_validation =  0 
	scan_time =  0 
	graceful_restart =  0 
	bgp_restart_time =  0 
	bgp_stalepath_time =  0 
	}
distance-list {   
	admin_distance =  0 
	src_prefix =  "string" 
	acl_str =  "string" 
	ext_routes_dist =  0 
	int_routes_dist =  0 
	local_routes_dist =  0 
	}
maximum_paths_value = 0
originate = 0
timers {  
 	bgp_keepalive =  0 
	bgp_holdtime =  0 
	}
synchronization = 0
auto_summary = 0
uuid = "string"
user_tag = "string"
network {  
 synchronization {  
 	network_synchronization =  0 
	uuid =  "string" 
	}
ip-cidr-list {   
	network_ipv4_cidr =  "string" 
	route_map =  "string" 
	backdoor =  0 
	description =  "string" 
	comm_value =  "string" 
	uuid =  "string" 
	}
	}
neighbor {  
 peer-group-neighbor-list {   
	peer_group =  "string" 
	activate =  0 
	allowas_in =  0 
	allowas_in_count =  0 
	prefix_list_direction =  "string" 
	default_originate =  0 
	route_map =  "string" 
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
	maximum_prefix =  0 
	maximum_prefix_thres =  0 
	next_hop_self =  0 
neighbor-prefix-lists {   
	nbr_prefix_list =  "string" 
	nbr_prefix_list_direction =  "string" 
	}
	remove_private_as =  0 
neighbor-route-map-lists {   
	nbr_route_map =  "string" 
	nbr_rmap_direction =  "string" 
	}
	send_community_val =  "string" 
	inbound =  0 
	unsuppress_map =  "string" 
	weight =  0 
	uuid =  "string" 
	}
ipv4-neighbor-list {   
	neighbor_ipv4 =  "string" 
	peer_group_name =  "string" 
	activate =  0 
	allowas_in =  0 
	allowas_in_count =  0 
	prefix_list_direction =  "string" 
	graceful_restart =  0 
	default_originate =  0 
	route_map =  "string" 
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
	maximum_prefix =  0 
	maximum_prefix_thres =  0 
	next_hop_self =  0 
neighbor-prefix-lists {   
	nbr_prefix_list =  "string" 
	nbr_prefix_list_direction =  "string" 
	}
	remove_private_as =  0 
neighbor-route-map-lists {   
	nbr_route_map =  "string" 
	nbr_rmap_direction =  "string" 
	}
	send_community_val =  "string" 
	inbound =  0 
	unsuppress_map =  "string" 
	weight =  0 
	uuid =  "string" 
	}
ipv6-neighbor-list {   
	neighbor_ipv6 =  "string" 
	peer_group_name =  "string" 
	activate =  0 
	allowas_in =  0 
	allowas_in_count =  0 
	prefix_list_direction =  "string" 
	graceful_restart =  0 
	default_originate =  0 
	route_map =  "string" 
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
	maximum_prefix =  0 
	maximum_prefix_thres =  0 
	next_hop_self =  0 
neighbor-prefix-lists {   
	nbr_prefix_list =  "string" 
	nbr_prefix_list_direction =  "string" 
	}
	remove_private_as =  0 
neighbor-route-map-lists {   
	nbr_route_map =  "string" 
	nbr_rmap_direction =  "string" 
	}
	send_community_val =  "string" 
	inbound =  0 
	unsuppress_map =  "string" 
	weight =  0 
	uuid =  "string" 
	}
ethernet-neighbor-list {   
	ethernet =  0 
	unnumbered =  0 
	peer_group_name =  "string" 
	uuid =  "string" 
	}
ve-neighbor-list {   
	ve =  0 
	unnumbered =  0 
	peer_group_name =  "string" 
	uuid =  "string" 
	}
trunk-neighbor-list {   
	trunk =  0 
	unnumbered =  0 
	peer_group_name =  "string" 
	uuid =  "string" 
	}
	}
redistribute {  
 connected_cfg {  
 	connected =  0 
	route_map =  "string" 
	}
floating_ip_cfg {  
 	floating_ip =  0 
	route_map =  "string" 
	}
lw4o6_cfg {  
 	lw4o6 =  0 
	route_map =  "string" 
	}
static_nat_cfg {  
 	static_nat =  0 
	route_map =  "string" 
	}
ip_nat_cfg {  
 	ip_nat =  0 
	route_map =  "string" 
	}
ip_nat_list_cfg {  
 	ip_nat_list =  0 
	route_map =  "string" 
	}
isis_cfg {  
 	isis =  0 
	route_map =  "string" 
	}
ospf_cfg {  
 	ospf =  0 
	route_map =  "string" 
	}
rip_cfg {  
 	rip =  0 
	route_map =  "string" 
	}
static_cfg {  
 	static =  0 
	route_map =  "string" 
	}
nat_map_cfg {  
 	nat_map =  0 
	route_map =  "string" 
	}
vip {  
 only_flagged_cfg {  
 	only_flagged =  0 
	route_map =  "string" 
	}
only_not_flagged_cfg {  
 	only_not_flagged =  0 
	route_map =  "string" 
	}
	}
	uuid =  "string" 
	}
address_family {  
 ipv6 {  
 bgp {  
 	dampening =  0 
	dampening_half =  0 
	dampening_start_reuse =  0 
	dampening_start_supress =  0 
	dampening_max_supress =  0 
	dampening_unreachability =  0 
	route_map =  "string" 
	}
distance {  
 	distance_ext =  0 
	distance_int =  0 
	distance_local =  0 
	}
	maximum_paths_value =  0 
	originate =  0 
aggregate-address-list {   
	aggregate_address =  "string" 
	as_set =  0 
	summary_only =  0 
	}
	auto_summary =  0 
	synchronization =  0 
	uuid =  "string" 
network {  
 synchronization {  
 	network_synchronization =  0 
	uuid =  "string" 
	}
ipv6-network-list {   
	network_ipv6 =  "string" 
	route_map =  "string" 
	backdoor =  0 
	description =  "string" 
	comm_value =  "string" 
	uuid =  "string" 
	}
	}
neighbor {  
 peer-group-neighbor-list {   
	peer_group =  "string" 
	activate =  0 
	allowas_in =  0 
	allowas_in_count =  0 
	prefix_list_direction =  "string" 
	default_originate =  0 
	route_map =  "string" 
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
	maximum_prefix =  0 
	maximum_prefix_thres =  0 
	next_hop_self =  0 
neighbor-prefix-lists {   
	nbr_prefix_list =  "string" 
	nbr_prefix_list_direction =  "string" 
	}
	remove_private_as =  0 
neighbor-route-map-lists {   
	nbr_route_map =  "string" 
	nbr_rmap_direction =  "string" 
	}
	send_community_val =  "string" 
	inbound =  0 
	unsuppress_map =  "string" 
	weight =  0 
	uuid =  "string" 
	}
ipv4-neighbor-list {   
	neighbor_ipv4 =  "string" 
	peer_group_name =  "string" 
	activate =  0 
	allowas_in =  0 
	allowas_in_count =  0 
	prefix_list_direction =  "string" 
	graceful_restart =  0 
	default_originate =  0 
	route_map =  "string" 
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
	maximum_prefix =  0 
	maximum_prefix_thres =  0 
	next_hop_self =  0 
neighbor-prefix-lists {   
	nbr_prefix_list =  "string" 
	nbr_prefix_list_direction =  "string" 
	}
	remove_private_as =  0 
neighbor-route-map-lists {   
	nbr_route_map =  "string" 
	nbr_rmap_direction =  "string" 
	}
	send_community_val =  "string" 
	inbound =  0 
	unsuppress_map =  "string" 
	weight =  0 
	uuid =  "string" 
	}
ipv6-neighbor-list {   
	neighbor_ipv6 =  "string" 
	peer_group_name =  "string" 
	activate =  0 
	allowas_in =  0 
	allowas_in_count =  0 
	prefix_list_direction =  "string" 
	graceful_restart =  0 
	default_originate =  0 
	route_map =  "string" 
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
	maximum_prefix =  0 
	maximum_prefix_thres =  0 
	next_hop_self =  0 
neighbor-prefix-lists {   
	nbr_prefix_list =  "string" 
	nbr_prefix_list_direction =  "string" 
	}
	remove_private_as =  0 
neighbor-route-map-lists {   
	nbr_route_map =  "string" 
	nbr_rmap_direction =  "string" 
	}
	send_community_val =  "string" 
	inbound =  0 
	unsuppress_map =  "string" 
	weight =  0 
	uuid =  "string" 
	}
ethernet-neighbor-ipv6-list {   
	ethernet =  0 
	peer_group_name =  "string" 
	uuid =  "string" 
	}
ve-neighbor-ipv6-list {   
	ve =  0 
	peer_group_name =  "string" 
	uuid =  "string" 
	}
trunk-neighbor-ipv6-list {   
	trunk =  0 
	peer_group_name =  "string" 
	uuid =  "string" 
	}
	}
redistribute {  
 connected_cfg {  
 	connected =  0 
	route_map =  "string" 
	}
floating_ip_cfg {  
 	floating_ip =  0 
	route_map =  "string" 
	}
nat64_cfg {  
 	nat64 =  0 
	route_map =  "string" 
	}
nat_map_cfg {  
 	nat_map =  0 
	route_map =  "string" 
	}
lw4o6_cfg {  
 	lw4o6 =  0 
	route_map =  "string" 
	}
static_nat_cfg {  
 	static_nat =  0 
	route_map =  "string" 
	}
ip_nat_cfg {  
 	ip_nat =  0 
	route_map =  "string" 
	}
ip_nat_list_cfg {  
 	ip_nat_list =  0 
	route_map =  "string" 
	}
isis_cfg {  
 	isis =  0 
	route_map =  "string" 
	}
ospf_cfg {  
 	ospf =  0 
	route_map =  "string" 
	}
rip_cfg {  
 	rip =  0 
	route_map =  "string" 
	}
static_cfg {  
 	static =  0 
	route_map =  "string" 
	}
vip {  
 only_flagged_cfg {  
 	only_flagged =  0 
	route_map =  "string" 
	}
only_not_flagged_cfg {  
 	only_not_flagged =  0 
	route_map =  "string" 
	}
	}
	uuid =  "string" 
	}
	}
	}
 
}

```

## Argument Reference

* `bgp` - Border Gateway Protocol (BGP)
* `as-number` - AS number
* `aggregate-address` - Configure BGP aggregate entries (Aggregate IPv6 prefix)
* `as-set` - Generate AS set path information
* `summary-only` - Filter more specific routes from updates
* `always-compare-med` - Allow comparing MED from different neighbors
* `ignore` - Ignore as-path length in selecting a route
* `compare-routerid` - Compare router-id for identical EBGP paths
* `remove-recv-med` - To remove rcvd MED attribute
* `remove-send-med` - To remove send MED attribute
* `missing-as-worst` - Treat missing MED as the least preferred one
* `dampening` - Enable route-flap dampening
* `dampening-half-time` - Reachability Half-life time for the penalty(minutes)
* `dampening-reuse` - Value to start reusing a route
* `dampening-supress` - Value to start suppressing a route
* `dampening-max-supress` - Maximum duration to suppress a stable route(minutes)
* `dampening-penalty` - Un-reachability Half-life time for the penalty(minutes)
* `route-map` - Route map reference (Pointer to route-map entries)
* `local-preference-value` - Configure default local preference value
* `deterministic-med` - Pick the best-MED path among paths advertised from the neighboring AS
* `enforce-first-as` - Enforce the first AS for EBGP routes
* `fast-external-failover` - Immediately reset session if a link to a directly connected external peer goes down
* `log-neighbor-changes` - Log neighbor up/down and reset reason
* `nexthop-trigger-count` - BGP nexthop-tracking status (count)
* `router-id` - Override current router identifier (peers will reset) (Manually configured router identifier)
* `override-validation` - override router-id validation
* `scan-time` - Configure background scan interval (Scan interval (sec) [Default:60 Disable:0])
* `graceful-restart` - enable graceful-restart helper for this neighbor
* `bgp-restart-time` - BGP Peer Graceful Restart time in seconds (default 90)
* `bgp-stalepath-time` - BGP Graceful Restart Stalepath retention time in seconds (default 360)
* `admin-distance` - Define an administrative distance
* `src-prefix` - IP source prefix
* `acl-str` - Access list name
* `ext-routes-dist` - Distance for routes external to the AS
* `int-routes-dist` - Distance for routes internal to the AS
* `local-routes-dist` - Distance for local routes
* `maximum-paths-value` - Supported BGP multipath numbers
* `originate` - Distribute an IPv6 default route
* `bgp-keepalive` - Keepalive interval
* `bgp-holdtime` - Holdtime
* `synchronization` - Perform IGP synchronization
* `auto-summary` - Enable automatic network number summarization
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `network-synchronization` - Perform IGP synchronization
* `network-ipv4-cidr` - Specify network mask
* `backdoor` - Specify a BGP backdoor route
* `description` - Network specific description (Up to 80 characters describing this network)
* `comm-value` - community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export
* `peer-group` - Neighbor tag
* `activate` - Enable the Address Family for this Neighbor
* `allowas-in` - Accept as-path with my AS present in it
* `allowas-in-count` - Number of occurrences of AS number
* `prefix-list-direction` - 'both': both; 'receive': receive; 'send': send;
* `default-originate` - Originate default route to this neighbor
* `distribute-list` - Filter updates to/from this neighbor (IP standard/extended/named access list)
* `distribute-list-direction` - 'in': in; 'out': out;
* `filter-list` - Establish BGP filters (AS path access-list name)
* `filter-list-direction` - 'in': in; 'out': out;
* `maximum-prefix` - Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))
* `maximum-prefix-thres` - threshold-value, 1 to 100 percent
* `next-hop-self` - Disable the next hop calculation for this neighbor
* `nbr-prefix-list` - Filter updates to/from this neighbor (Name of a prefix list)
* `nbr-prefix-list-direction` - 'in': in; 'out': out;
* `remove-private-as` - Remove private AS number from outbound updates
* `nbr-route-map` - Apply route map to neighbor (Name of route map)
* `nbr-rmap-direction` - 'in': in; 'out': out;
* `send-community-val` - 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;
* `inbound` - Allow inbound soft reconfiguration for this neighbor
* `unsuppress-map` - Route-map to selectively unsuppress suppressed routes (Name of route map)
* `weight` - Set default weight for routes from this neighbor
* `neighbor-ipv4` - Neighbor address
* `peer-group-name` - Configure peer-group (peer-group name)
* `neighbor-ipv6` - Neighbor IPv6 address
* `ethernet` - Ethernet interface number
* `ve` - Virtual ethernet interface number
* `trunk` - Trunk interface number
* `connected` - Connected
* `floating-ip` - Floating IP
* `lw4o6` - LW4O6 Prefix
* `static-nat` - Static NAT Prefix
* `ip-nat` - IP NAT
* `ip-nat-list` - IP NAT list
* `isis` - ISO IS-IS
* `ospf` - Open Shortest Path First (OSPF)
* `rip` - Routing Information Protocol (RIP)
* `static` - Static routes
* `nat-map` - NAT MAP Prefix
* `only-flagged` - Selected Virtual IP (VIP)
* `only-not-flagged` - Only not flagged
* `dampening-half` - Reachability Half-life time for the penalty(minutes)
* `dampening-start-reuse` - Value to start reusing a route
* `dampening-start-supress` - Value to start suppressing a route
* `dampening-unreachability` - Un-reachability Half-life time for the penalty(minutes)
* `distance-ext` - Distance for routes external to the AS
* `distance-int` - Distance for routes internal to the AS
* `distance-local` - Distance for local routes
* `network-ipv6` - Specify a network to announce via BGP
* `nat64` - NAT64 Prefix

