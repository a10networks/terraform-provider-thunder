---
layout: "thunder"
page_title: "thunder: thunder_router_ospf"
sidebar_current: "docs-thunder-resource-router-ospf"
description: |-
    Provides details about thunder router ospf resource for A10
---

# thunder\_router\_ospf

`thunder_router_ospf` Provides details about thunder router ospf
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_ospf" "resourceRouterOspfTest" {
	process_id = 0
auto_cost_reference_bandwidth = 0
bfd_all_interfaces = 0
rfc1583_compatible = 0
default_metric = 0
distance {  
 	distance_value =  0 
distance_ospf {  
 	distance_external =  0 
	distance_inter_area =  0 
	distance_intra_area =  0 
	}
	}
distribute-internal-list {   
	di_type =  "string" 
	di_area_ipv4 =  "string" 
	di_area_num =  0 
	di_cost =  0 
	}
distribute-lists {   
	value =  "string" 
	direction =  "string" 
	protocol =  "string" 
	ospf_id =  0 
	option =  "string" 
	}
ha-standby-extra-cost {   
	extra_cost =  0 
	group =  0 
	}
host-list {   
	host_address =  "string" 
area_cfg {  
 	area_ipv4 =  "string" 
	area_num =  0 
	cost =  0 
	}
	}
log_adjacency_changes_cfg {  
 	state =  "string" 
	}
max_concurrent_dd = 0
maximum_area = 0
neighbor-list {   
	address =  "string" 
	cost =  0 
	poll_interval =  0 
	priority =  0 
	}
network-list {   
	network_ipv4 =  "string" 
	network_ipv4_mask =  "string" 
	network_ipv4_cidr =  "string" 
network_area {  
 	network_area_ipv4 =  "string" 
	network_area_num =  0 
	instance_value =  0 
	}
	}
ospf_1 {  
 abr_type {  
 	option =  "string" 
	}
	}
router_id {  
 	value =  "string" 
	}
overflow {  
 database {  
 	count =  0 
	limit =  "string" 
	db_external =  0 
	recovery_time =  0 
	}
	}
passive_interface {  
 loopback-cfg {   
	loopback =  0 
	loopback_address =  "string" 
	}
trunk-cfg {   
	trunk =  0 
	trunk_address =  "string" 
	}
ve-cfg {   
	ve =  0 
	ve_address =  "string" 
	}
tunnel-cfg {   
	tunnel =  0 
	tunnel_address =  "string" 
	}
lif-cfg {   
	lif =  0 
	lif_address =  "string" 
	}
eth-cfg {   
	ethernet =  0 
	eth_address =  "string" 
	}
	}
summary-address-list {   
	summary_address =  "string" 
	not_advertise =  0 
	tag =  0 
	}
timers {  
 spf {  
 exp {  
 	min_delay =  0 
	max_delay =  0 
	}
	}
	}
uuid = "string"
user_tag = "string"
default_information {  
 	originate =  0 
	always =  0 
	metric =  0 
	metric_type =  0 
	route_map =  "string" 
	uuid =  "string" 
	}
area-list {   
	area_ipv4 =  "string" 
	area_num =  0 
auth_cfg {  
 	authentication =  0 
	message_digest =  0 
	}
filter-lists {   
	filter_list =  0 
	acl_name =  "string" 
	acl_direction =  "string" 
	plist_name =  "string" 
	plist_direction =  "string" 
	}
nssa_cfg {  
 	nssa =  0 
	no_redistribution =  0 
	no_summary =  0 
	translator_role =  "string" 
	default_information_originate =  0 
	metric =  0 
	metric_type =  0 
	}
	default_cost =  0 
range-list {   
	area_range_prefix =  "string" 
	option =  "string" 
	}
	shortcut =  "string" 
stub_cfg {  
 	stub =  0 
	no_summary =  0 
	}
virtual-link-list {   
	virtual_link_ip_addr =  "string" 
	bfd =  0 
	hello_interval =  0 
	dead_interval =  0 
	retransmit_interval =  0 
	transmit_delay =  0 
	virtual_link_authentication =  0 
	virtual_link_auth_type =  "string" 
	authentication_key =  "string" 
	message_digest_key =  0 
	md5 =  "string" 
	}
	uuid =  "string" 
	}
redistribute {  
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
	ip_nat =  0 
	metric_ip_nat =  0 
	metric_type_ip_nat =  "string" 
	route_map_ip_nat =  "string" 
	tag_ip_nat =  0 
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
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `ospf` - Open Shortest Path First (OSPF)
* `process-id` - OSPF process ID
* `auto-cost-reference-bandwidth` - Use reference bandwidth method to assign OSPF cost (The reference bandwidth in terms of Mbits per second)
* `bfd-all-interfaces` - Enable BFD on all interfaces
* `rfc1583-compatible` - Compatible with RFC 1583
* `default-metric` - Set metric of redistributed routes (Default metric)
* `distance-value` - OSPF Administrative distance
* `distance-external` - External routes (Distance for external routes)
* `distance-inter-area` - Inter-area routes (Distance for inter-area routes)
* `distance-intra-area` - Intra-area routes (Distance for intra-area routes)
* `di-type` - 'lw4o6': LW4O6 Prefix; 'floating-ip': Floating IP; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'vip': Only not flagged Virtual IP (VIP); 'vip-only-flagged': Selected Virtual IP (VIP);
* `di-area-ipv4` - OSPF area ID as a IP address format
* `di-area-num` - OSPF area ID as a decimal value
* `di-cost` - Cost of route
* `value` - OSPF router-id in IPv4 address format
* `direction` - 'in': Filter incoming routing updates; 'out': Filter outgoing routing updates;
* `protocol` - 'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'lw4o6': LW4O6 Prefix; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'ospf': Open Shortest Path First (OSPF); 'rip': Routing Information Protocol (RIP); 'static': Static routes;
* `ospf-id` - OSPF process ID
* `option` - 'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;
* `extra-cost` - The extra cost value
* `group` - Group (Group ID)
* `host-address` - Host address
* `area-ipv4` - OSPF area ID in IP address format
* `area-num` - OSPF area ID as a decimal value
* `cost` - OSPF cost for point-to-multipoint neighbor (Metric)
* `state` - 'detail': Log changes in adjacency state; 'disable': Disable logging;
* `max-concurrent-dd` - Maximum number allowed to process DD concurrently (Number of DD process)
* `maximum-area` - Maximum number of non-backbone areas (OSPF area limit)
* `address` - Neighbor address
* `poll-interval` - OSPF dead-router polling interval (Seconds)
* `priority` - OSPF priority of non-broadcast neighbor
* `network-ipv4` - Network number
* `network-ipv4-mask` - OSPF wild card bits
* `network-ipv4-cidr` - OSPF network prefix
* `network-area-ipv4` - OSPF area ID in IP address format
* `network-area-num` - OSPF area ID as a decimal value
* `instance-value` - Instance ID
* `count` - Maximum number of LSAs
* `limit` - 'hard': Hard limit: Instance will be shutdown if exceeded; 'soft': Soft limit: Warning will be given if exceeded;
* `db-external` - Maximum number of LSAs
* `recovery-time` - Time to recover (0 not recover)
* `loopback` - Loopback interface (Port number)
* `loopback-address` - Address of Interface
* `trunk` - Trunk interface (Trunk interface number)
* `trunk-address` - Address of Interface
* `ve` - Virtual ethernet interface (Virtual ethernet interface number)
* `ve-address` - Address of Interface
* `tunnel` - Tunnel interface (Tunnel interface number)
* `tunnel-address` - Address of Interface
* `lif` - Logical interface (Lif interface number)
* `lif-address` - Address of Interface
* `ethernet` - Ethernet interface (Port number)
* `eth-address` - Address of Interface
* `summary-address` - Configure IP address summaries (Summary prefix)
* `not-advertise` - Suppress routes that match the prefix
* `tag` - Set tag for routes redistributed into OSPF (32-bit tag value)
* `min-delay` - Minimum delay between receiving a change to SPF calculation in milliseconds
* `max-delay` - Maximum delay between receiving a change to SPF calculation in milliseconds
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `originate` - Distribute a default route
* `always` - Always advertise default route
* `metric` - OSPF default metric (OSPF metric)
* `metric-type` - '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
* `route-map` - Route map reference (Pointer to route-map entries)
* `authentication` - Enable authentication
* `message-digest` - Use message-digest authentication
* `filter-list` - Filter networks between OSPF areas
* `acl-name` - Filter networks by access-list (Name of an access-list)
* `acl-direction` - 'in': Filter networks sent to this area; 'out': Filter networks sent from this area;
* `plist-name` - Filter networks by prefix-list (Name of an IP prefix-list)
* `plist-direction` - 'in': Filter networks sent to this area; 'out': Filter networks sent from this area;
* `nssa` - Specify a NSSA area
* `no-redistribution` - No redistribution into this NSSA area
* `no-summary` - Do not inject inter-area routes into area
* `translator-role` - 'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;
* `default-information-originate` - Originate Type 7 default into NSSA area
* `default-cost` - Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)
* `area-range-prefix` - Area range for IPv4 prefix
* `shortcut` - 'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;
* `stub` - Configure OSPF area as stub
* `virtual-link-ip-addr` - ID (IP addr) associated with virtual link neighbor
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `hello-interval` - Hello packet interval (Seconds)
* `dead-interval` - Dead router detection time (Seconds)
* `retransmit-interval` - LSA retransmit interval (Seconds)
* `transmit-delay` - LSA transmission delay (Seconds)
* `virtual-link-authentication` - Enable authentication
* `virtual-link-auth-type` - 'message-digest': Use message-digest authentication; 'null': Use null authentication;
* `authentication-key` - Set authentication key (Authentication key (8 chars))
* `message-digest-key` - Set message digest key (Key ID)
* `md5` - Use MD5 algorithm (Authentication key (16 chars))
* `type` - 'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;
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

