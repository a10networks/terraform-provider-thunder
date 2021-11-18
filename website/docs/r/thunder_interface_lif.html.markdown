---
layout: "thunder"
page_title: "thunder: thunder_interface_lif"
sidebar_current: "docs-thunder-resource-interface-lif"
description: |-
    Provides details about thunder interface lif resource for A10
---

# thunder\_interface\_lif

`thunder_interface_lif` Provides details about thunder interface lif
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_interface_lif" "resourceInterfaceLifTest" {
	ifname = "string"
mtu = 0
action = "string"
name = "string"
access_list {  
 	acl_id =  0 
	acl_name =  "string" 
	}
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
ip {  
 	dhcp =  0 
address-list {   
	ipv6_addr =  "string" 
	anycast =  0 
	link_local =  0 
	}
	allow_promiscuous_vip =  0 
	cache_spoofing_port =  0 
	inside =  0 
	outside =  0 
	generate_membership_query =  0 
	query_interval =  0 
	max_resp_time =  0 
	uuid =  "string" 
router {  
 isis {  
 	tag =  "string" 
	uuid =  "string" 
	}
	}
rip {  
 authentication {  
 str {  
 	string =  "string" 
	}
mode {  
 	mode =  "string" 
	}
key_chain {  
 	key_chain =  "string" 
	}
	}
	send_packet =  0 
	receive_packet =  0 
send_cfg {  
 	send =  0 
	version =  "string" 
	}
receive_cfg {  
 	receive =  0 
	version =  "string" 
	}
split_horizon_cfg {  
 	state =  "string" 
	}
	uuid =  "string" 
	}
ospf {  
 ospf_global {  
 authentication_cfg {  
 	authentication =  0 
	value =  "string" 
	}
	authentication_key =  "string" 
bfd_cfg {  
 	bfd =  0 
	disable =  0 
	}
	cost =  0 
database_filter_cfg {  
 	database_filter =  "string" 
	out =  0 
	}
	dead_interval =  0 
	disable =  "string" 
	hello_interval =  0 
message-digest-cfg {   
	message_digest_key =  0 
	md5_value =  "string" 
	}
	mtu =  0 
	mtu_ignore =  0 
network {  
 	broadcast =  0 
	non_broadcast =  0 
	point_to_point =  0 
	point_to_multipoint =  0 
	p2mp_nbma =  0 
	}
	priority =  0 
	retransmit_interval =  0 
	transmit_delay =  0 
	uuid =  "string" 
	}
ospf-ip-list {   
	ip_addr =  "string" 
	authentication =  0 
	value =  "string" 
	authentication_key =  "string" 
	cost =  0 
	database_filter =  "string" 
	out =  0 
	dead_interval =  0 
	hello_interval =  0 
message-digest-cfg {   
	message_digest_key =  0 
	md5_value =  "string" 
	}
	mtu_ignore =  0 
	priority =  0 
	retransmit_interval =  0 
	transmit_delay =  0 
	uuid =  "string" 
	}
	}
	}
ipv6 {  
 address-list {   
	ipv6_addr =  "string" 
	anycast =  0 
	link_local =  0 
	}
	ipv6_enable =  0 
	inside =  0 
	outside =  0 
	uuid =  "string" 
router {  
 ripng {  
 	rip =  0 
	uuid =  "string" 
	}
ospf {  
 area-list {   
	area_id_num =  0 
	area_id_addr =  "string" 
	tag =  "string" 
	instance_id =  0 
	}
	uuid =  "string" 
	}
isis {  
 	tag =  "string" 
	uuid =  "string" 
	}
	}
ospf {  
 network-list {   
	broadcast_type =  "string" 
	p2mp_nbma =  0 
	network_instance_id =  0 
	}
	bfd =  0 
	disable =  0 
cost-cfg {   
	cost =  0 
	instance_id =  0 
	}
dead-interval-cfg {   
	dead_interval =  0 
	instance_id =  0 
	}
hello-interval-cfg {   
	hello_interval =  0 
	instance_id =  0 
	}
mtu-ignore-cfg {   
	mtu_ignore =  0 
	instance_id =  0 
	}
neighbor-cfg {   
	neighbor =  "string" 
	neig_inst =  0 
	neighbor_cost =  0 
	neighbor_poll_interval =  0 
	neighbor_priority =  0 
	}
priority-cfg {   
	priority =  0 
	instance_id =  0 
	}
retransmit-interval-cfg {   
	retransmit_interval =  0 
	instance_id =  0 
	}
transmit-delay-cfg {   
	transmit_delay =  0 
	instance_id =  0 
	}
	uuid =  "string" 
	}
	}
bfd {  
 authentication {  
 	key_id =  0 
	method =  "string" 
	password =  "string" 
	}
	echo =  0 
	demand =  0 
interval_cfg {  
 	interval =  0 
	min_rx =  0 
	multiplier =  0 
	}
	uuid =  "string" 
	}
isis {  
 authentication {  
 send-only-list {   
	send_only =  0 
	level =  "string" 
	}
mode-list {   
	mode =  "string" 
	level =  "string" 
	}
key-chain-list {   
	key_chain =  "string" 
	level =  "string" 
	}
	}
bfd_cfg {  
 	bfd =  0 
	disable =  0 
	}
	circuit_type =  "string" 
csnp-interval-list {   
	csnp_interval =  0 
	level =  "string" 
	}
	padding =  0 
hello-interval-list {   
	hello_interval =  0 
	level =  "string" 
	}
hello-interval-minimal-list {   
	hello_interval_minimal =  0 
	level =  "string" 
	}
hello-multiplier-list {   
	hello_multiplier =  0 
	level =  "string" 
	}
	lsp_interval =  0 
mesh_group {  
 	value =  0 
	blocked =  0 
	}
metric-list {   
	metric =  0 
	level =  "string" 
	}
	network =  "string" 
password-list {   
	password =  "string" 
	level =  "string" 
	}
priority-list {   
	priority =  0 
	level =  "string" 
	}
	retransmit_interval =  0 
wide-metric-list {   
	wide_metric =  0 
	level =  "string" 
	}
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `lif` - Logical interface
* `ifname` - Lif interface name
* `mtu` - OSPF interface MTU (MTU size)
* `action` - 'enable': Enable; 'disable': Disable;
* `name` - Name for the interface
* `acl-id` - ACL id
* `acl-name` - Apply an access list (Named Access List)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'num_pkts': num_pkts; 'num_total_bytes': num_total_bytes; 'num_unicast_pkts': num_unicast_pkts; 'num_broadcast_pkts': num_broadcast_pkts; 'num_multicast_pkts': num_multicast_pkts; 'num_tx_pkts': num_tx_pkts; 'num_total_tx_bytes': num_total_tx_bytes; 'num_unicast_tx_pkts': num_unicast_tx_pkts; 'num_broadcast_tx_pkts': num_broadcast_tx_pkts; 'num_multicast_tx_pkts': num_multicast_tx_pkts; 'dropped_dis_rx_pkts': dropped_dis_rx_pkts; 'dropped_rx_pkts': dropped_rx_pkts; 'dropped_dis_tx_pkts': dropped_dis_tx_pkts; 'dropped_tx_pkts': dropped_tx_pkts; 'dropped_rx_pkts_gre_key': dropped_rx_pkts_gre_key;
* `dhcp` - Use DHCP to configure IP address
* `ipv6-addr` - Set the IPv6 address of an interface
* `anycast` - Configure an IPv6 anycast address
* `link-local` - Configure an IPv6 link local address
* `allow-promiscuous-vip` - Allow traffic to be associated with promiscuous VIP
* `cache-spoofing-port` - This interface connects to spoofing cache
* `inside` - Configure interface as inside
* `outside` - Configure interface as outside
* `generate-membership-query` - Enable Membership Query
* `query-interval` - 1 - 255 (Default is 125)
* `max-resp-time` - Maximum Response Time (Max Response Time (Default is 100))
* `tag` - ISO routing area tag
* `string` - The RIP authentication string
* `mode` - 'md5': Keyed message digest;
* `key-chain` - Authentication key-chain (Name of key-chain)
* `send-packet` - Enable sending packets through the specified interface
* `receive-packet` - Enable receiving packet through the specified interface
* `send` - Advertisement transmission
* `version` - '1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;
* `receive` - Advertisement reception
* `state` - 'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;
* `authentication` - Enable authentication
* `value` - Mesh group number
* `authentication-key` - Authentication password (key) (The OSPF password (key))
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `disable` - Disable BFD
* `cost` - Interface cost
* `database-filter` - 'all': Filter all LSA;
* `out` - Outgoing LSA
* `dead-interval` - Interval after which a neighbor is declared dead (Seconds)
* `hello-interval` - Set Hello interval in seconds (Hello interval value)
* `message-digest-key` - Message digest authentication password (key) (Key id)
* `md5-value` - The OSPF password (1-16)
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `mtu-ignore` - Ignores the MTU in DBD packets
* `broadcast` - Specify OSPF broadcast multi-access network
* `non-broadcast` - Specify OSPF NBMA network
* `point-to-point` - Specify OSPF point-to-point network
* `point-to-multipoint` - Specify OSPF point-to-multipoint network
* `p2mp-nbma` - Specify non-broadcast point-to-multipoint network
* `priority` - Set priority for Designated Router election (Priority value)
* `retransmit-interval` - Set per-LSP retransmission interval (Interval between retransmissions of the same LSP (seconds))
* `transmit-delay` - Link state transmit delay (Seconds)
* `ip-addr` - Address of interface
* `ipv6-enable` - Enable IPv6 processing
* `rip` - RIP Routing for IPv6
* `area-id-num` - OSPF area ID as a decimal value
* `area-id-addr` - OSPF area ID in IP address format
* `instance-id` - Specify the interface instance ID
* `broadcast-type` - 'broadcast': Specify OSPF broadcast multi-access network; 'non-broadcast': Specify OSPF NBMA network; 'point-to-point': Specify OSPF point-to-point network; 'point-to-multipoint': Specify OSPF point-to-multipoint network;
* `network-instance-id` - Specify the interface instance ID
* `neighbor` - OSPFv3 neighbor (Neighbor IPv6 address)
* `neig-inst` - Specify the interface instance ID
* `neighbor-cost` - OSPF cost for point-to-multipoint neighbor (metric)
* `neighbor-poll-interval` - OSPF dead-router polling interval (Seconds)
* `neighbor-priority` - OSPF priority of non-broadcast neighbor
* `key-id` - Key ID
* `method` - 'md5': Keyed MD5; 'meticulous-md5': Meticulous Keyed MD5; 'meticulous-sha1': Meticulous Keyed SHA1; 'sha1': Keyed SHA1; 'simple': Simple Password;
* `password` - Configure the authentication password for interface
* `echo` - Enable BFD Echo
* `demand` - Demand mode
* `interval` - Transmit interval between BFD packets (Milliseconds)
* `min-rx` - Minimum receive interval capability (Milliseconds)
* `multiplier` - Multiplier value used to compute holddown (value used to multiply the interval)
* `send-only` - Authentication send-only
* `level` - 'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;
* `circuit-type` - 'level-1': Level-1 only adjacencies are formed; 'level-1-2': Level-1-2 adjacencies are formed; 'level-2-only': Level-2 only adjacencies are formed;
* `csnp-interval` - Set CSNP interval in seconds (CSNP interval value)
* `padding` - Add padding to IS-IS hello packets
* `hello-interval-minimal` - Set Hello holdtime 1 second, interval depends on multiplier
* `hello-multiplier` - Set multiplier for Hello holding time (Hello multiplier value)
* `lsp-interval` - Set LSP transmission interval (LSP transmission interval (milliseconds))
* `blocked` - Block LSPs on this interface
* `metric` - Configure the metric for interface (Default metric)
* `network` - 'broadcast': Specify IS-IS broadcast multi-access network; 'point-to-point': Specify IS-IS point-to-point network;
* `wide-metric` - Configure the wide metric for interface

