---
layout: "thunder"
page_title: "thunder: thunder_ethernet"
sidebar_current: "docs-thunder-resource-ethernet"
description: |-
    Provides details about thunder ethernet resource for A10
---

# thunder\_ethernet

`thunder_ethernet` provides details about how to enable ethernet interfaces on a device
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ethernet" "Interface_Ethernet_Test" {
fec_forced_on = 0
trap_source = 0
ip {  
 	uuid =  "string" 
address-list {   
	address_type =  "string" 
	ipv6_addr =  "string" 
	}
	generate_membership_query =  0 
	cache_spoofing_port =  0 
	inside =  0 
	allow_promiscuous_vip =  0 
	client =  0 
	max_resp_time =  0 
	query_interval =  0 
	outside =  0 
helper-address-list {   
	helper_address =  "string" 
	}
stateful_firewall {  
 	uuid =  "string" 
	class_list =  "string" 
	inside =  0 
	outside =  0 
	acl_id =  0 
	access_list =  0 
	}
rip {  
 receive_cfg {  
 	receive =  0 
	version =  "string" 
	}
	uuid =  "string" 
	receive_packet =  0 
split_horizon_cfg {  
 	state =  "string" 
	}
authentication {  
 key_chain {  
 	key_chain =  "string" 
	}
mode {  
 	mode =  "string" 
	}
str {  
 	string =  "string" 
	}
	}
send_cfg {  
 	version =  "string" 
	send =  0 
	}
	send_packet =  0 
	}
	ttl_ignore =  0 
router {  
 isis {  
 	tag =  "string" 
	uuid =  "string" 
	}
	}
	dhcp =  0 
	server =  0 
ospf {  
 ospf-ip-list {   
	dead_interval =  0 
	authentication_key =  "string" 
	uuid =  "string" 
	mtu_ignore =  0 
	transmit_delay =  0 
	value =  "string" 
	priority =  0 
	authentication =  0 
	cost =  0 
	database_filter =  "string" 
	hello_interval =  0 
	ip_addr =  "string" 
	retransmit_interval =  0 
message-digest-cfg {   
	message_digest_key =  0 
md5 {  
 	md5_value =  "string" 
	encrypted =  "Unknown Type: encrypted" 
	}
	}
	out =  0 
	}
ospf_global {  
 	cost =  0 
	dead_interval =  0 
	authentication_key =  "string" 
network {  
 	broadcast =  0 
	point_to_multipoint =  0 
	non_broadcast =  0 
	point_to_point =  0 
	p2mp_nbma =  0 
	}
	mtu_ignore =  0 
	transmit_delay =  0 
authentication_cfg {  
 	authentication =  0 
	value =  "string" 
	}
	retransmit_interval =  0 
bfd_cfg {  
 	disable =  0 
	bfd =  0 
	}
	disable =  "string" 
	hello_interval =  0 
database_filter_cfg {  
 	database_filter =  "string" 
	out =  0 
	}
	priority =  0 
	mtu =  0 
message-digest-cfg {   
	message_digest_key =  0 
md5 {  
 	md5_value =  "string" 
	encrypted =  "Unknown Type: encrypted" 
	}
	}
	uuid =  "string" 
	}
	}
	slb_partition_redirect =  0 
	}
ddos {  
 	outside =  0 
	inside =  0 
	uuid =  "string" 
	}
l3_vlan_fwd_disable = 0
access_list {  
 	acl_name =  "string" 
	acl_id =  0 
	}
speed = "string"
speed_forced_40g = 0
lldp {  
 tx_dot1_cfg {  
 	link_aggregation =  0 
	vlan =  0 
	tx_dot1_tlvs =  0 
	}
notification_cfg {  
 	notification =  0 
	notif_enable =  0 
	}
enable_cfg {  
 	rx =  0 
	tx =  0 
	rt_enable =  0 
	}
tx_tlvs_cfg {  
 	system_capabilities =  0 
	system_description =  0 
	management_address =  0 
	tx_tlvs =  0 
	exclude =  0 
	port_description =  0 
	system_name =  0 
	}
	uuid =  "string" 
	}
speed_forced_1g = 0
bfd {  
 interval_cfg {  
 	interval =  0 
	min_rx =  0 
	multiplier =  0 
	}
authentication {  
 	encrypted =  "Unknown Type: encrypted" 
	password =  "string" 
	method =  "string" 
	key_id =  0 
	}
	echo =  0 
	uuid =  "string" 
	demand =  0 
	}
media_type_copper = 0
ifnum = 0
remove_vlan_tag = 0
monitor-list {   
	monitor_vlan =  0 
	monitor =  "string" 
	mirror_index =  0 
	}
cpu_process = 0
auto_neg_enable = 0
port_breakout = "string"
spanning_tree {  
 	auto_edge =  0 
	path_cost =  0 
	uuid =  "string" 
	admin_edge =  0 
	}
speed_forced_10g = 0
map {  
 	inside =  0 
	map_t_inside =  0 
	uuid =  "string" 
	map_t_outside =  0 
	outside =  0 
	}
traffic_distribution_mode = "string"
trunk-group-list {   
	uuid =  "string" 
	trunk_number =  0 
	user_tag =  "string" 
udld_timeout_cfg {  
 	slow =  0 
	fast =  0 
	}
	mode =  "string" 
	timeout =  "string" 
	type =  "string" 
	admin_key =  0 
	port_priority =  0 
	}
nptv6 {  
 domain-list {   
	domain_name =  "string" 
	bind_type =  "string" 
	uuid =  "string" 
	}
	}
uuid = "string"
cpu_process_dir = "string"
isis {  
 priority-list {   
	priority =  0 
	level =  "string" 
	}
	padding =  0 
hello-interval-minimal-list {   
	hello_interval_minimal =  0 
	level =  "string" 
	}
mesh_group {  
 	value =  0 
	blocked =  0 
	}
	network =  "string" 
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
csnp-interval-list {   
	csnp_interval =  0 
	level =  "string" 
	}
	retransmit_interval =  0 
password-list {   
	password =  "string" 
	level =  "string" 
	}
bfd_cfg {  
 	disable =  0 
	bfd =  0 
	}
wide-metric-list {   
	wide_metric =  0 
	level =  "string" 
	}
hello-interval-list {   
	hello_interval =  0 
	level =  "string" 
	}
	circuit_type =  "string" 
hello-multiplier-list {   
	hello_multiplier =  0 
	level =  "string" 
	}
metric-list {   
	metric =  0 
	level =  "string" 
	}
	lsp_interval =  0 
	uuid =  "string" 
	}
name = "string"
duplexity = "string"
icmpv6_rate_limit {  
 	lockup_period_v6 =  0 
	normal_v6 =  0 
	lockup_v6 =  0 
	}
user_tag = "string"
mtu = 0
ipv6 {  
 	uuid =  "string" 
address-list {   
	address_type =  "string" 
	ipv6_addr =  "string" 
	}
	inside =  0 
	ipv6_enable =  0 
rip {  
 split_horizon_cfg {  
 	state =  "string" 
	}
	uuid =  "string" 
	}
	outside =  0 
stateful_firewall {  
 	uuid =  "string" 
	class_list =  "string" 
	acl_name =  "string" 
	inside =  0 
	outside =  0 
	access_list =  0 
	}
	ttl_ignore =  0 
router {  
 ripng {  
 	uuid =  "string" 
	rip =  0 
	}
ospf {  
 area-list {   
	area_id_addr =  "string" 
	tag =  "string" 
	instance_id =  0 
	area_id_num =  0 
	}
	uuid =  "string" 
	}
isis {  
 	tag =  "string" 
	uuid =  "string" 
	}
	}
access_list_cfg {  
 	inbound =  0 
	v6_acl_name =  "string" 
	}
ospf {  
 	uuid =  "string" 
	bfd =  0 
cost-cfg {   
	cost =  0 
	instance_id =  0 
	}
priority-cfg {   
	priority =  0 
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
retransmit-interval-cfg {   
	retransmit_interval =  0 
	instance_id =  0 
	}
	disable =  0 
transmit-delay-cfg {   
	transmit_delay =  0 
	instance_id =  0 
	}
neighbor-cfg {   
	neighbor_priority =  0 
	neig_inst =  0 
	neighbor_poll_interval =  0 
	neighbor_cost =  0 
	neighbor =  "string" 
	}
network-list {   
	broadcast_type =  "string" 
	p2mp_nbma =  0 
	network_instance_id =  0 
	}
dead-interval-cfg {   
	dead_interval =  0 
	instance_id =  0 
	}
	}
router_adver {  
 	max_interval =  0 
	default_lifetime =  0 
	reachable_time =  0 
	other_config_action =  "string" 
	floating_ip_default_vrid =  "string" 
	managed_config_action =  "string" 
	min_interval =  0 
	rate_limit =  0 
	adver_mtu_disable =  0 
prefix-list {   
	not_autonomous =  0 
	valid_lifetime =  0 
	not_on_link =  0 
	prefix =  "string" 
	preferred_lifetime =  0 
	}
	floating_ip =  "string" 
	adver_vrid =  0 
	use_floating_ip_default_vrid =  0 
	action =  "string" 
	adver_vrid_default =  0 
	adver_mtu =  0 
	retransmit_timer =  0 
	hop_limit =  0 
	use_floating_ip =  0 
	}
	}
sampling-enable {   
	counters1 =  "string" 
	}
load_interval = 0
lw_4o6 {  
 	outside =  0 
	inside =  0 
	uuid =  "string" 
	}
action = "string"
fec_forced_off = 0
icmp_rate_limit {  
 	lockup =  0 
	lockup_period =  0 
	normal =  0 
	}
flow_control = 0
 
}

```

## Argument Reference

* `action` - ‘enable’: Enable Router Advertisements on this interface; ‘disable’: Disable Router Advertisements on this interface;
* `auto_neg_enable` - enable auto-negotiation
* `cpu_process` - All Packets to this port are processed by CPU
* `cpu_process_dir` - ‘primary’: Primary board; ‘blade’: blade board; ‘hash-dip’: Hash based on the Destination IP; ‘hash-sip’: Hash based on the Source IP; ‘hash-dmac’: Hash based on the Destination MAC; ‘hash-smac’: Hash based on the Source MAC;
* `duplexity` - ‘Full’: Full; ‘Half’: Half; ‘auto’: auto;
* `fec_forced_off` - turn off the FEC
* `fec_forced_on` - turn on the FEC
* `flow_control` - Enable 802.3x flow control on full duplex port
* `ifnum` - Ethernet interface number
* `load_interval` - Configure Load Interval (Seconds (5-300, Multiple of 5), default 300)
* `media_type_copper` - Set the media type to copper
* `mtu` - OSPF interface MTU (MTU size)
* `name` - Name for the interface
* `remove_vlan_tag` - Remove the vlan tag for egressing packets
* `speed` - ‘10’: 10; ‘100’: 100; ‘1000’: 1000; ‘auto’: auto;
* `speed_forced_40g` - force the speed to be 40G on 100G link
* `traffic_distribution_mode` - ‘sip’: sip; ‘dip’: dip; ‘primary’: primary; ‘blade’: blade; ‘l4-src-port’: l4-src-port; ‘l4-dst-port’: l4-dst-port;
* `trap_source` - The trap source
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `allow_promiscuous_vip` - Allow traffic to be associated with promiscuous VIP
* `cache_spoofing_port` - This interface connects to spoofing cache
* `client` - Client facing interface for IPv4/v6 traffic
* `dhcp` - Use DHCP to configure IP address
* `generate_membership_query` - Enable Membership Query
* `inside` - Configure LW-4over6 outside interface
* `max_resp_time` - Maximum Response Time (Max Response Time (Default is 100))
* `outside` - Configure LW-4over6 inside interface
* `query_interval` - 1 - 255 (Default is 125)
* `server` - Server facing interface for IPv4/v6 traffic
* `slb_partition_redirect` - Redirect SLB traffic across partition
* `ttl_ignore` - Ignore TTL decrement for a received packet before sending out
* `tag` - ISO routing area tag
* `helper_address` - Helper address for DHCP packets (IP address)
* `access_list` - Access-list for traffic from the outside
* `acl_id` - ACL id
* `class_list` - Class List (Class List Name)
* `receive_packet` - Enable receiving packet through the specified interface
* `send_packet` - Enable sending packets through the specified interface
* `receive` - Advertisement reception
* `version` - ‘1’: RIP version 1; ‘2’: RIP version 2; ‘1-compatible’: RIPv1-compatible; ‘1-2’: RIP version 1 & 2;
* `state` - ‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;
* `key_chain` - Authentication key-chain (Name of key-chain)
* `mode` - ‘md5’: Keyed message digest;
* `string` - The RIP authentication string
* `send` - Advertisement transmission
* `ipv4_address` - IP address
* `ipv4_netmask` - IP subnet mask
* `authentication` - Enable authentication
* `authentication_key` - Authentication password (key) (The OSPF password (key))
* `cost` - Interface cost
* `database_filter` - ‘all’: Filter all LSA;
* `dead_interval` - Interval after which a neighbor is declared dead (Seconds)
* `hello_interval` - Time between HELLO packets (Seconds)
* `ip_addr` - Address of interface
* `mtu_ignore` - Ignores the MTU in DBD packets
* `out` - Outgoing LSA
* `priority` - Router priority
* `retransmit_interval` - Time between retransmitting lost link state advertisements (Seconds)
* `transmit_delay` - Link state transmit delay (Seconds)
* `value` - Mesh group number
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `md5_value` - The OSPF password (1-16)
* `message_digest_key` - Message digest authentication password (key) (Key id)
* `disable` - Disable BFD
* `broadcast` - Specify OSPF broadcast multi-access network
* `non_broadcast` - Specify OSPF NBMA network
* `p2mp_nbma` - Specify non-broadcast point-to-multipoint network
* `point_to_multipoint` - Specify OSPF point-to-multipoint network
* `point_to_point` - Specify OSPF point-to-point network
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `acl_name` - Access-list Name
* `lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `lockup_period` - Lockup period (second)
* `normal` - Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit
* `link_aggregation` - Interface link aggregation information
* `tx_dot1_tlvs` - Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration
* `vlan` - Interface vlan information
* `notif_enable` - Interface lldp notification enable
* `notification` - Interface lldp notification configuration
* `exclude` - Configure which TLVs excluded. All basic TLVs will be included by default
* `management_address` - Interface lldp management address
* `port_description` - Interface lldp port description
* `system_capabilities` - Interface lldp system capabilities
* `system_description` - Interface lldp system description
* `system_name` - Interface lldp system name
* `tx_tlvs` - Interface lldp tx TLVs configuration
* `rt_enable` - Interface lldp enable/disable
* `rx` - Enable lldp rx
* `tx` - Enable lldp tx
* `demand` - Demand mode
* `echo` - Enable BFD Echo
* `key_id` - Key ID
* `method` - ‘md5’: Keyed MD5; ‘meticulous-md5’: Meticulous Keyed MD5; ‘meticulous-sha1’: Meticulous Keyed SHA1; ‘sha1’: Keyed SHA1; ‘simple’: Simple Password;
* `password` - Configure the authentication password for interface
* `interval` - Transmit interval between BFD packets (Milliseconds)
* `min_rx` - Minimum receive interval capability (Milliseconds)
* `multiplier` - Multiplier value used to compute holddown (value used to multiply the interval)
* `counters1` - ‘all’: all; ‘packets_input’: Input packets; ‘bytes_input’: Input bytes; ‘received_broadcasts’: Received broadcasts; ‘received_multicasts’: Received multicasts; ‘received_unicasts’: Received unicasts; ‘input_errors’: Input errors; ‘crc’: CRC; ‘frame’: Frames; ‘runts’: Runts; ‘giants’: Giants; ‘packets_output’: Output packets; ‘bytes_output’: Output bytes; ‘transmitted_broadcasts’: Transmitted braodcasts; ‘transmitted_multicasts’: Transmitted multicasts; ‘transmitted_unicasts’: Transmitted unicasts; ‘output_errors’: Output errors; ‘collisions’: Collisions;
* `map_t_inside` - Configure MAP inside interface (connected to MAP domains)
* `map_t_outside` - Configure MAP outside interface
* `admin_key` - LACP admin key (Admin key value)
* `port_priority` - Set LACP priority for a port (LACP port priority)
* `timeout` - ‘long’: Set LACP long timeout (default); ‘short’: Set LACP short timeout;
* `trunk_number` - Trunk Number
* `type` - ‘static’: Static (default); ‘lacp’: lacp; ‘lacp-udld’: lacp-udld;
* `fast` - fast timeout in unit of milli-seconds(default 1000)
* `slow` - slow timeout in unit of seconds
* `bind_type` - ‘inside’: This interface is connected to NPTv6 inside networks; ‘outside’: This interface is connected to NPTv6 outside networks;
* `domain_name` - NPTv6 domain name
* `circuit_type` - ‘level-1’: Level-1 only adjacencies are formed; ‘level-1-2’: Level-1-2 adjacencies are formed; ‘level-2-only’: Level-2 only adjacencies are formed;
* `lsp_interval` - Set LSP transmission interval (LSP transmission interval (milliseconds))
* `network` - ‘broadcast’: Specify IS-IS broadcast multi-access network; ‘point-to-point’: Specify IS-IS point-to-point network;
* `padding` - Add padding to IS-IS hello packets
* `level` - ‘level-1’: Speficy interval for level-1 CSNPs; ‘level-2’: Specify interval for level-2 CSNPs;
* `hello_interval_minimal` - Set Hello holdtime 1 second, interval depends on multiplier
* `blocked` - Block LSPs on this interface
* `send_only` - Authentication send-only
* `wide_metric` - Configure the wide metric for interface
* `hello_multiplier` - Set multiplier for Hello holding time (Hello multiplier value)
* `metric` - Configure the metric for interface (Default metric)
* `csnp_interval` - Set CSNP interval in seconds (CSNP interval value)
* `lockup_period_v6` - Lockup period (second)
* `lockup_v6` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `normal_v6` - Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit
* `ipv6_enable` - Enable IPv6 processing
* `address_type` - ‘anycast’: Configure an IPv6 anycast address; ‘link-local’: Configure an IPv6 link local address;
* `ipv6_addr` - Set the IPv6 address of an interface
* `rip` - RIP Routing for IPv6
* `area_id_addr` - OSPF area ID in IP address format
* `area_id_num` - OSPF area ID as a decimal value
* `instance_id` - Specify the interface instance ID
* `inbound` - ACL applied on incoming packets to this interface
* `v6_acl_name` - Apply ACL rules to incoming packets on this interface (Named Access List)
* `broadcast_type` - ‘broadcast’: Specify OSPF broadcast multi-access network; ‘non-broadcast’: Specify OSPF NBMA network; ‘point-to-point’: Specify OSPF point-to-point network; ‘point-to-multipoint’: Specify OSPF point-to-multipoint network;
* `network_instance_id` - Specify the interface instance ID
* `neig_inst` - Specify the interface instance ID
* `neighbor` - OSPFv3 neighbor (Neighbor IPv6 address)
* `neighbor_cost` - OSPF cost for point-to-multipoint neighbor (metric)
* `neighbor_poll_interval` - OSPF dead-router polling interval (Seconds)
* `neighbor_priority` - OSPF priority of non-broadcast neighbor
* `adver_mtu` - Set Router Advertisement MTU Option
* `adver_mtu_disable` - Disable Router Advertisement MTU Option
* `adver_vrid` - Specify ha VRRP-A vrid
* `adver_vrid_default` - Default VRRP-A vrid
* `default_lifetime` - Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))
* `floating_ip` - Use a floating IP as the source address for Router advertisements
* `floating_ip_default_vrid` - Use a floating IP as the source address for Router advertisements
* `hop_limit` - Set Router Advertisement Hop Limit (default: 255)
* `managed_config_action` - ‘enable’: Enable the Managed Address Configuration flag; ‘disable’: Disable the Managed Address Configuration flag (default);
* `max_interval` - Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))
* `min_interval` - Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))
* `other_config_action` - ‘enable’: Enable the Other Stateful Configuration flag; ‘disable’: Disable the Other Stateful Configuration flag (default);
* `rate_limit` - Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)
* `reachable_time` - Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))
* `retransmit_timer` - Set Router Advertisement Retransmit Timer (default: 0)
* `use_floating_ip` - Use a floating IP as the source address for Router advertisements
* `use_floating_ip_default_vrid` - Use a floating IP as the source address for Router advertisements
* `not_autonomous` - Specify that the Prefix is not usable for autoconfiguration (default:autonomous)
* `not_on_link` - Specify that the Prefix is not On-Link (default: on-link)
* `preferred_lifetime` - Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))
* `prefix` - Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)
* `valid_lifetime` - Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))
* `mirror_index` - Mirror index
* `monitor` - ‘input’: Incoming packets; ‘output’: Outgoing packets; ‘both’: Both incoming and outgoing packets;
* `monitor_vlan` - VLAN number
