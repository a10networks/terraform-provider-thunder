---
layout: "thunder"
page_title: "thunder: thunder_interface_ve_ip"
sidebar_current: "docs-thunder-resource-interface-ve-ip"
description: |-
	Provides details about thunder interface ve ip resource for A10
---

# thunder\_interface\_ve\_ip

`thunder_interface_ve_ip` Provides details about thunder interface ve ip
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_interface_ve_ip" "Interface_Ve_Ip_Test" {
uuid = "string"
generate_membership_query = 0
address-list {   
        ipv4_address =  "string" 
        ipv4_netmask =  "string" 
        }
inside = 0
allow_promiscuous_vip = 0
helper-address-list {   
        helper_address =  "string" 
        }
max_resp_time = 0
query_interval = 0
outside = 0
client = 0
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
ttl_ignore = 0
router {  
 isis {  
        tag =  "string" 
        uuid =  "string" 
        }
        }
dhcp = 0
server = 0
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
        encrypted =  "string" 
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
        encrypted =  "string" 
        }
        }
        uuid =  "string" 
        }
        }
slb_partition_redirect = 0
 
}

```

## Argument Reference

* `allow_promiscuous_vip` - Allow traffic to be associated with promiscuous VIP
* `client` - Client facing interface for IPv4/v6 traffic
* `dhcp` - Use DHCP to configure IP address
* `generate_membership_query` - Enable Membership Query
* `inside` - Inside (private) interface for stateful firewall
* `max_resp_time` - Maximum Response Time (Max Response Time (Default is 100))
* `outside` - Outside (public) interface for stateful firewall
* `query_interval` - 1 - 255 (Default is 125)
* `server` - Server facing interface for IPv4/v6 traffic
* `slb_partition_redirect` - Redirect SLB traffic across partition
* `ttl_ignore` - Ignore TTL decrement for a received packet
* `uuid` - uuid of the object
* `ipv4_address` - IP address
* `ipv4_netmask` - IP subnet mask
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
* `mode` - ‘md5’: Keyed message digest; ‘text’: Clear text authentication;
* `string` - The RIP authentication string
* `send` - Advertisement transmission
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
* `value` - ‘message-digest’: Use message-digest authentication; ‘null’: Use no authentication;
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `md5_value` - The OSPF password (1-16)
* `message_digest_key` - Message digest authentication password (key) (Key id)
* `disable` - Disable BFD
* `mtu` - OSPF interface MTU (MTU size)
* `broadcast` - Specify OSPF broadcast multi-access network
* `non_broadcast` - Specify OSPF NBMA network
* `p2mp_nbma` - Specify non-broadcast point-to-multipoint network
* `point_to_multipoint` - Specify OSPF point-to-multipoint network
* `point_to_point` - Specify OSPF point-to-point network
* `bfd` - Bidirectional Forwarding Detection (BFD)

