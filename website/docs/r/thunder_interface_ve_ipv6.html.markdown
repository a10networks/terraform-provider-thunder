---
layout: "thunder"
page_title: "thunder: thunder_interface_ve_ipv6"
sidebar_current: "docs-thunder-resource-interface-ve-ipv6"
description: |-
	Provides details about thunder interface ve ipv6 resource for A10
---

# thunder\_interface\_ve\_ipv6

`thunder_interface_ve_ipv6` Provides details about thunder interface ve ipv6
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_interface_ve_ipv6" "Interface_Ve_Ipv6_Test" {
uuid = "string"
inbound = 0
address-list {   
        address_type =  "string" 
        ipv6_addr =  "string" 
        }
inside = 0
ipv6_enable = 0
rip {  
 split_horizon_cfg {  
        state =  "string" 
        }
        uuid =  "string" 
        }
outside = 0
stateful_firewall {  
        uuid =  "string" 
        class_list =  "string" 
        acl_name =  "string" 
        inside =  0 
        outside =  0 
        access_list =  0 
        }
v6_acl_name = "string"
ttl_ignore = 0
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
```

## Argument Reference

* `ifnum` - Interface no.
* `inbound` - ACL applied on incoming packets to this interface
* `inside` - Inside (private) interface for stateful firewall
* `ipv6_enable` - Enable IPv6 processing
* `outside` - Outside (public) interface for stateful firewall
* `ttl_ignore` - Ignore TTL decrement for a received packet
* `uuid` - uuid of the object
* `v6_acl_name` - Apply ACL rules to incoming packets on this interface (Named Access List)
* `address_type` - ‘anycast’: Configure an IPv6 anycast address; ‘link-local’: Configure an IPv6 link local address;
* `ipv6_addr` - Set the IPv6 address of an interface
* `state` - ‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;
* `access_list` - Access-list for traffic from the outside
* `acl_name` - Access-list Name
* `class_list` - Class List (Class List Name)
* `rip` - RIP Routing for IPv6
* `area_id_addr` - OSPF area ID in IP address format
* `area_id_num` - OSPF area ID as a decimal value
* `instance_id` - Specify the interface instance ID
* `tag` - ISO routing area tag
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `disable` - Disable BFD
* `cost` - Interface cost
* `hello_interval` - Time between HELLO packets (Seconds)
* `priority` - Router priority
* `mtu_ignore` - Ignores the MTU in DBD packets
* `retransmit_interval` - Time between retransmitting lost link state advertisements (Seconds)
* `broadcast_type` - ‘broadcast’: Specify OSPF broadcast multi-access network; ‘non-broadcast’: Specify OSPF NBMA network; ‘point-to-point’: Specify OSPF point-to-point network; ‘point-to-multipoint’: Specify OSPF point-to-multipoint network;
* `network_instance_id` - Specify the interface instance ID
* `p2mp_nbma` - Specify non-broadcast point-to-multipoint network
* `transmit_delay` - Link state transmit delay (Seconds)
* `neig_inst` - Specify the interface instance ID
* `neighbor` - OSPFv3 neighbor (Neighbor IPv6 address)
* `neighbor_cost` - OSPF cost for point-to-multipoint neighbor (metric)
* `neighbor_poll_interval` - OSPF dead-router polling interval (Seconds)
* `neighbor_priority` - OSPF priority of non-broadcast neighbor
* `dead_interval` - Interval after which a neighbor is declared dead (Seconds)
* `action` - ‘enable’: Enable Router Advertisements on this interface; ‘disable’: Disable Router Advertisements on this interface;
* `adver_mtu` - Set Router Advertisement MTU Option
* `adver_mtu_disable` - Disable Router Advertisement MTU Option
* `adver_vrid` - Vrid
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

