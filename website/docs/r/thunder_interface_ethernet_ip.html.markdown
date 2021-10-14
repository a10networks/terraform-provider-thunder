---
layout: "thunder"
page_title: "thunder: thunder_interface_ethernet_ip"
sidebar_current: "docs-thunder-resource-interface-ethernet-ip"
description: |-
    Provides details about thunder interface ethernet ip resource for A10
---

# thunder\_interface\_ethernet\_ip

`thunder_interface_ethernet_ip` Provides details about thunder interface ethernet ip
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_interface_ethernet_ip" "resourceInterfaceEthernetIpTest" {
	dhcp = 0
address-list {   
	ipv4_address =  "string" 
	ipv4_netmask =  "string" 
	}
allow_promiscuous_vip = 0
cache_spoofing_port = 0
helper-address-list {   
	helper_address =  "string" 
	}
inside = 0
outside = 0
ttl_ignore = 0
syn_cookie = 0
slb_partition_redirect = 0
generate_membership_query = 0
query_interval = 0
max_resp_time = 0
client = 0
server = 0
uuid = "string"
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

```

## Argument Reference

* `ip` - Global IP configuration subcommands
* `dhcp` - Use DHCP to configure IP address
* `ipv4-address` - IP address
* `ipv4-netmask` - IP subnet mask
* `allow-promiscuous-vip` - Allow traffic to be associated with promiscuous VIP
* `cache-spoofing-port` - This interface connects to spoofing cache
* `helper-address` - Helper address for DHCP packets (IP address)
* `inside` - Inside (private) interface for stateful firewall
* `outside` - Outside (public) interface for stateful firewall
* `ttl-ignore` - Ignore TTL decrement for a received packet before sending out
* `syn-cookie` - Configure Enable SYN-cookie on the interface
* `slb-partition-redirect` - Redirect SLB traffic across partition
* `generate-membership-query` - Enable Membership Query
* `query-interval` - 1 - 255 (Default is 125)
* `max-resp-time` - Maximum Response Time (Max Response Time (Default is 100))
* `client` - Client facing interface for IPv4/v6 traffic
* `server` - Server facing interface for IPv4/v6 traffic
* `uuid` - uuid of the object
* `class-list` - Class List (Class List Name)
* `access-list` - Access-list for traffic from the outside
* `acl-id` - ACL id
* `tag` - ISO routing area tag
* `string` - The RIP authentication string
* `mode` - 'md5': Keyed message digest; 'text': Clear text authentication;
* `key-chain` - Authentication key-chain (Name of key-chain)
* `send-packet` - Enable sending packets through the specified interface
* `receive-packet` - Enable receiving packet through the specified interface
* `send` - Advertisement transmission
* `version` - '1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;
* `receive` - Advertisement reception
* `state` - 'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;
* `authentication` - Enable authentication
* `value` - 'message-digest': Use message-digest authentication; 'null': Use no authentication;
* `authentication-key` - Authentication password (key) (The OSPF password (key))
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `disable` - 'all': All functionality;
* `cost` - Interface cost
* `database-filter` - 'all': Filter all LSA;
* `out` - Outgoing LSA
* `dead-interval` - Interval after which a neighbor is declared dead (Seconds)
* `hello-interval` - Time between HELLO packets (Seconds)
* `message-digest-key` - Message digest authentication password (key) (Key id)
* `md5-value` - The OSPF password (1-16)
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `mtu` - OSPF interface MTU (MTU size)
* `mtu-ignore` - Ignores the MTU in DBD packets
* `broadcast` - Specify OSPF broadcast multi-access network
* `non-broadcast` - Specify OSPF NBMA network
* `point-to-point` - Specify OSPF point-to-point network
* `point-to-multipoint` - Specify OSPF point-to-multipoint network
* `p2mp-nbma` - Specify non-broadcast point-to-multipoint network
* `priority` - Router priority
* `retransmit-interval` - Time between retransmitting lost link state advertisements (Seconds)
* `transmit-delay` - Link state transmit delay (Seconds)
* `ip-addr` - Address of interface

