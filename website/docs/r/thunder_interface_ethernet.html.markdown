---
layout: "thunder"
page_title: "thunder: thunder_interface_ethernet"
sidebar_current: "docs-thunder-resource-interface-ethernet"
description: |-
    Provides details about thunder interface ethernet resource for A10
---

# thunder\_interface\_ethernet

`thunder_interface_ethernet` Provides details about thunder interface ethernet
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_interface_ethernet" "resourceInterfaceEthernetTest" {
	ifnum = 0
name = "string"
port_scan_detection = "string"
ping_sweep_detection = "string"
l3_vlan_fwd_disable = 0
load_interval = 0
media_type_copper = 0
auto_neg_enable = 0
fec_forced_on = 0
fec_forced_off = 0
port_breakout = "string"
speed_forced_1g = 0
speed_forced_10g = 0
speed_forced_40g = 0
ipg_bit_time = 0
remove_vlan_tag = 0
mtu = 0
trap_source = 0
duplexity = "string"
speed = "string"
flow_control = 0
action = "string"
icmp_rate_limit {  
 	normal =  0 
	lockup =  0 
	lockup_period =  0 
	}
icmpv6_rate_limit {  
 	normal_v6 =  0 
	lockup_v6 =  0 
	lockup_period_v6 =  0 
	}
monitor-list {   
	monitor =  "string" 
	mirror_index =  0 
	monitor_vlan =  0 
	}
cpu_process = 0
cpu_process_dir = "string"
traffic_distribution_mode = "string"
virtual_wire = 0
update_l2_info = 0
vlan_learning = "string"
mac_learning = "string"
access_list {  
 	acl_id =  0 
	acl_name =  "string" 
	}
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
packet_capture_template = "string"
lldp {  
 enable_cfg {  
 	rt_enable =  0 
	rx =  0 
	tx =  0 
	}
notification_cfg {  
 	notification =  0 
	notif_enable =  0 
	}
tx_dot1_cfg {  
 	tx_dot1_tlvs =  0 
	link_aggregation =  0 
	vlan =  0 
	}
tx_tlvs_cfg {  
 	tx_tlvs =  0 
	exclude =  0 
	management_address =  0 
	port_description =  0 
	system_capabilities =  0 
	system_description =  0 
	system_name =  0 
	}
	uuid =  "string" 
	}
ddos {  
 	outside =  0 
	inside =  0 
	uuid =  "string" 
	}
ip {  
 	dhcp =  0 
address-list {   
	ipv6_addr =  "string" 
	address_type =  "string" 
	}
	allow_promiscuous_vip =  0 
	cache_spoofing_port =  0 
helper-address-list {   
	helper_address =  "string" 
	}
	inside =  0 
	outside =  0 
	ttl_ignore =  0 
	syn_cookie =  0 
	slb_partition_redirect =  0 
	generate_membership_query =  0 
	query_interval =  0 
	max_resp_time =  0 
	client =  0 
	server =  0 
	uuid =  "string" 
stateful_firewall {  
 	inside =  0 
	class_list =  "string" 
	outside =  0 
	access_list =  0 
	acl_id =  0 
	uuid =  "string" 
	}
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
	address_type =  "string" 
	}
	inside =  0 
	outside =  0 
	ipv6_enable =  0 
	ttl_ignore =  0 
access_list_cfg {  
 	v6_acl_name =  "string" 
	inbound =  0 
	}
router_adver {  
 	action =  "string" 
	hop_limit =  0 
	max_interval =  0 
	min_interval =  0 
	default_lifetime =  0 
	rate_limit =  0 
	reachable_time =  0 
	retransmit_timer =  0 
	adver_mtu_disable =  0 
	adver_mtu =  0 
prefix-list {   
	prefix =  "string" 
	not_autonomous =  0 
	not_on_link =  0 
	preferred_lifetime =  0 
	valid_lifetime =  0 
	}
	managed_config_action =  "string" 
	other_config_action =  "string" 
	adver_vrid =  0 
	use_floating_ip =  0 
	floating_ip =  "string" 
	adver_vrid_default =  0 
	use_floating_ip_default_vrid =  0 
	floating_ip_default_vrid =  "string" 
	}
	uuid =  "string" 
stateful_firewall {  
 	inside =  0 
	class_list =  "string" 
	outside =  0 
	access_list =  0 
	acl_name =  "string" 
	uuid =  "string" 
	}
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
rip {  
 split_horizon_cfg {  
 	state =  "string" 
	}
	uuid =  "string" 
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
nptv6 {  
 domain-list {   
	domain_name =  "string" 
	bind_type =  "string" 
	uuid =  "string" 
	}
	}
map {  
 	inside =  0 
	outside =  0 
	map_t_inside =  0 
	map_t_outside =  0 
	uuid =  "string" 
	}
lw_4o6 {  
 	outside =  0 
	inside =  0 
	uuid =  "string" 
	}
trunk-group-list {   
	trunk_number =  0 
	type =  "string" 
	admin_key =  0 
	port_priority =  0 
udld_timeout_cfg {  
 	fast =  0 
	slow =  0 
	}
	mode =  "string" 
	timeout =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
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
spanning_tree {  
 	auto_edge =  0 
	admin_edge =  0 
instance-list {   
	instance_start =  0 
	mstp_path_cost =  0 
	}
	path_cost =  0 
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `ethernet` - Ethernet interface
* `ifnum` - Ethernet interface number
* `name` - Name for the interface
* `port-scan-detection` - 'enable': Enable port scan detection; 'disable': Disable port scan detection(default);
* `ping-sweep-detection` - 'enable': Enabl ping sweep detection; 'disable': Disable ping sweep detection(default);
* `load-interval` - Configure Load Interval (Seconds (5-300, Multiple of 5), default 300)
* `media-type-copper` - Set the media type to copper
* `auto-neg-enable` - enable auto-negotiation
* `fec-forced-on` - turn on the FEC
* `fec-forced-off` - turn off the FEC
* `port-breakout` - '4x10G': Breakout 100G/40G ports into 4x10G ports; '4x25G': Breakout 100G ports into 4x25G ports; '2x50G': Breakout 100G ports into 2x50G ports;
* `speed-forced-1g` - force the speed to be 1G on 25G link
* `speed-forced-10g` - force the speed to be 10G on 25G link
* `speed-forced-40g` - force the speed to be 40G on 100G link
* `ipg-bit-time` - Set Inter-packet-gap interval in bit timing, default is 96
* `remove-vlan-tag` - Remove the vlan tag for egressing packets
* `mtu` - OSPF interface MTU (MTU size)
* `trap-source` - The trap source
* `duplexity` - 'Full': Full; 'Half': Half; 'auto': auto;
* `speed` - '10': 10; '100': 100; '1000': 1000; 'auto': auto;
* `flow-control` - Enable 802.3x flow control on full duplex port
* `action` - 'enable': Enable Router Advertisements on this interface; 'disable': Disable Router Advertisements on this interface;
* `normal` - Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit
* `lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `lockup-period` - Lockup period (second)
* `normal-v6` - Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit
* `lockup-v6` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `lockup-period-v6` - Lockup period (second)
* `monitor` - 'input': Incoming packets; 'output': Outgoing packets; 'both': Both incoming and outgoing packets;
* `mirror-index` - Mirror index
* `monitor-vlan` - VLAN number
* `cpu-process` - All Packets to this port are processed by CPU
* `cpu-process-dir` - 'primary': Primary board; 'blade': blade board; 'hash-dip': Hash based on the Destination IP; 'hash-sip': Hash based on the Source IP; 'hash-dmac': Hash based on the Destination MAC; 'hash-smac': Hash based on the Source MAC;
* `traffic-distribution-mode` - 'sip': sip; 'dip': dip; 'primary': primary; 'blade': blade; 'l4-src-port': l4-src-port; 'l4-dst-port': l4-dst-port;
* `virtual-wire` - Mark ethernet as a virtual wire interface
* `update-l2-info` - Update and use received L2 information
* `vlan-learning` - 'enable': Enable VLAN learning; 'disable': Disable VLAN learning;
* `mac-learning` - 'enable': Enable MAC learning; 'disable': Disable MAC learning; 'dmac-only': Enable destination MAC learning only;
* `acl-id` - ACL id
* `acl-name` - Access-list Name
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'packets_input': Input packets; 'bytes_input': Input bytes; 'received_broadcasts': Received broadcasts; 'received_multicasts': Received multicasts; 'received_unicasts': Received unicasts; 'input_errors': Input errors; 'crc': CRC; 'frame': Frames; 'runts': Runts; 'giants': Giants; 'packets_output': Output packets; 'bytes_output': Output bytes; 'transmitted_broadcasts': Transmitted broadcasts; 'transmitted_multicasts': Transmitted multicasts; 'transmitted_unicasts': Transmitted unicasts; 'output_errors': Output errors; 'collisions': Collisions; 'giants_output': Output Giants; 'rate_pkt_sent': Packet sent rate packets/sec; 'rate_byte_sent': Byte sent rate bits/sec; 'rate_pkt_rcvd': Packet received rate packets/sec; 'rate_byte_rcvd': Byte received rate bits/sec; 'load_interval': Load Interval;
* `packet-capture-template` - Name of the packet capture template to be bind with this object
* `rt-enable` - Interface lldp enable/disable
* `rx` - Enable lldp rx
* `tx` - Enable lldp tx
* `notification` - Interface lldp notification configuration
* `notif-enable` - Interface lldp notification enable
* `tx-dot1-tlvs` - Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration
* `link-aggregation` - Interface link aggregation information
* `vlan` - Interface vlan information
* `tx-tlvs` - Interface lldp tx TLVs configuration
* `exclude` - Configure which TLVs excluded. All basic TLVs will be included by default
* `management-address` - Interface lldp management address
* `port-description` - Interface lldp port description
* `system-capabilities` - Interface lldp system capabilities
* `system-description` - Interface lldp system description
* `system-name` - Interface lldp system name
* `outside` - Configure LW-4over6 inside interface
* `inside` - Configure LW-4over6 outside interface
* `dhcp` - Use DHCP to configure IP address
* `ipv6-addr` - Set the IPv6 address of an interface
* `address-type` - 'anycast': Configure an IPv6 anycast address; 'link-local': Configure an IPv6 link local address;
* `allow-promiscuous-vip` - Allow traffic to be associated with promiscuous VIP
* `cache-spoofing-port` - This interface connects to spoofing cache
* `helper-address` - Helper address for DHCP packets (IP address)
* `ttl-ignore` - Ignore TTL decrement for a received packet before sending out
* `syn-cookie` - Configure Enable SYN-cookie on the interface
* `slb-partition-redirect` - Redirect SLB traffic across partition
* `generate-membership-query` - Enable Membership Query
* `query-interval` - 1 - 255 (Default is 125)
* `max-resp-time` - Maximum Response Time (Max Response Time (Default is 100))
* `client` - Client facing interface for IPv4/v6 traffic
* `server` - Server facing interface for IPv4/v6 traffic
* `class-list` - Class List (Class List Name)
* `access-list` - Access-list for traffic from the outside
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
* `v6-acl-name` - Apply ACL rules to incoming packets on this interface (Named Access List)
* `inbound` - ACL applied on incoming packets to this interface
* `hop-limit` - Set Router Advertisement Hop Limit (default: 255)
* `max-interval` - Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))
* `min-interval` - Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))
* `default-lifetime` - Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))
* `rate-limit` - Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)
* `reachable-time` - Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))
* `retransmit-timer` - Set Router Advertisement Retransmit Timer (default: 0)
* `adver-mtu-disable` - Disable Router Advertisement MTU Option
* `adver-mtu` - Set Router Advertisement MTU Option
* `prefix` - Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)
* `not-autonomous` - Specify that the Prefix is not usable for autoconfiguration (default:autonomous)
* `not-on-link` - Specify that the Prefix is not On-Link (default: on-link)
* `preferred-lifetime` - Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))
* `valid-lifetime` - Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))
* `managed-config-action` - 'enable': Enable the Managed Address Configuration flag; 'disable': Disable the Managed Address Configuration flag (default);
* `other-config-action` - 'enable': Enable the Other Stateful Configuration flag; 'disable': Disable the Other Stateful Configuration flag (default);
* `adver-vrid` - Specify ha VRRP-A vrid
* `use-floating-ip` - Use a floating IP as the source address for Router advertisements
* `floating-ip` - Use a floating IP as the source address for Router advertisements
* `adver-vrid-default` - Default VRRP-A vrid
* `use-floating-ip-default-vrid` - Use a floating IP as the source address for Router advertisements
* `floating-ip-default-vrid` - Use a floating IP as the source address for Router advertisements
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
* `domain-name` - NPTv6 domain name
* `bind-type` - 'inside': This interface is connected to NPTv6 inside networks; 'outside': This interface is connected to NPTv6 outside networks;
* `map-t-inside` - Configure MAP inside interface (connected to MAP domains)
* `map-t-outside` - Configure MAP outside interface
* `trunk-number` - Trunk Number
* `type` - 'static': Static (default); 'lacp': lacp; 'lacp-udld': lacp-udld;
* `admin-key` - LACP admin key (Admin key value)
* `port-priority` - Set LACP priority for a port (LACP port priority)
* `fast` - fast timeout in unit of milli-seconds(default 1000)
* `slow` - slow timeout in unit of seconds
* `timeout` - 'long': Set LACP long timeout (default); 'short': Set LACP short timeout;
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
* `auto-edge` - Enable auto-edge
* `admin-edge` - Enable admin-edge
* `instance-start` - Instance ID
* `mstp-path-cost` - Path cost (Limit)
* `path-cost` - Path cost (Limit)

