---
layout: "thunder"
page_title: "thunder: thunder_interface_ethernet_ipv6"
sidebar_current: "docs-thunder-resource-interface-ethernet-ipv6"
description: |-
    Provides details about thunder interface ethernet ipv6 resource for A10
---

# thunder\_interface\_ethernet\_ipv6

`thunder_interface_ethernet_ipv6` Provides details about thunder interface ethernet ipv6
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_interface_ethernet_ipv6" "resourceInterfaceEthernetIpv6Test" {
	address-list {   
	ipv6_addr =  "string" 
	address_type =  "string" 
	}
inside = 0
outside = 0
ipv6_enable = 0
ttl_ignore = 0
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
uuid = "string"
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

```

## Argument Reference

* `ipv6` - Global IPv6 configuration subcommands
* `ipv6-addr` - Set the IPv6 address of an interface
* `address-type` - 'anycast': Configure an IPv6 anycast address; 'link-local': Configure an IPv6 link local address;
* `inside` - Inside (private) interface for stateful firewall
* `outside` - Outside (public) interface for stateful firewall
* `ipv6-enable` - Enable IPv6 processing
* `ttl-ignore` - Ignore TTL decrement for a received packet before sending out
* `v6-acl-name` - Apply ACL rules to incoming packets on this interface (Named Access List)
* `inbound` - ACL applied on incoming packets to this interface
* `action` - 'enable': Enable Router Advertisements on this interface; 'disable': Disable Router Advertisements on this interface;
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
* `uuid` - uuid of the object
* `class-list` - Class List (Class List Name)
* `access-list` - Access-list for traffic from the outside
* `acl-name` - Access-list Name
* `rip` - RIP Routing for IPv6
* `area-id-num` - OSPF area ID as a decimal value
* `area-id-addr` - OSPF area ID in IP address format
* `tag` - ISO routing area tag
* `instance-id` - Specify the interface instance ID
* `state` - 'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;
* `broadcast-type` - 'broadcast': Specify OSPF broadcast multi-access network; 'non-broadcast': Specify OSPF NBMA network; 'point-to-point': Specify OSPF point-to-point network; 'point-to-multipoint': Specify OSPF point-to-multipoint network;
* `p2mp-nbma` - Specify non-broadcast point-to-multipoint network
* `network-instance-id` - Specify the interface instance ID
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `disable` - Disable BFD
* `cost` - Interface cost
* `dead-interval` - Interval after which a neighbor is declared dead (Seconds)
* `hello-interval` - Time between HELLO packets (Seconds)
* `mtu-ignore` - Ignores the MTU in DBD packets
* `neighbor` - OSPFv3 neighbor (Neighbor IPv6 address)
* `neig-inst` - Specify the interface instance ID
* `neighbor-cost` - OSPF cost for point-to-multipoint neighbor (metric)
* `neighbor-poll-interval` - OSPF dead-router polling interval (Seconds)
* `neighbor-priority` - OSPF priority of non-broadcast neighbor
* `priority` - Router priority
* `retransmit-interval` - Time between retransmitting lost link state advertisements (Seconds)
* `transmit-delay` - Link state transmit delay (Seconds)

