---
layout: "thunder"
page_title: "thunder: thunder_interface_management"
sidebar_current: "docs-thunder-resource-interface-management"
description: |-
	Provides details about thunder interface management resource for A10
---

# thunder\_interface\_management

`thunder_interface_management` Provides details about thunder interface management
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_interface_management" "Interface_Management_Test" {
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
flow_control = 0
broadcast_rate_limit {  
        rate =  0 
        bcast_rate_limit_enable =  0 
        }
uuid = "string"
duplexity = "string"
ip {  
        dhcp =  0 
        ipv4_address =  "string" 
        control_apps_use_mgmt_port =  0 
        default_gateway =  "string" 
        ipv4_netmask =  "string" 
        }
secondary_ip {  
        ipv4_netmask =  "string" 
        control_apps_use_mgmt_port =  0 
        secondary_ip =  0 
        default_gateway =  "string" 
        dhcp =  0 
        ipv4_address =  "string" 
        }
access_list {  
        acl_name =  "string" 
        acl_id =  0 
        }
sampling-enable {   
        counters1 =  "string" 
        }
ipv6 {   
        default_ipv6_gateway =  "string" 
        inbound =  0 
        address_type =  "string" 
        ipv6_addr =  "string" 
        v6_acl_name =  "string" 
        }
action = "string"
speed = "string"
 
}

```

## Argument Reference

* `action` - ‘enable’: Enable Management Port; ‘disable’: Disable Management Port;
* `duplexity` - ‘Full’: Full; ‘Half’: Half; ‘auto’: Auto;
* `flow_control` - Enable 802.3x flow control on full duplex port
* `speed` - ‘10’: 10 Mbs/sec; ‘100’: 100 Mbs/sec; ‘1000’: 1 Gb/sec; ‘auto’: Auto Negotiate Speed;  (Interface Speed)
* `uuid` - uuid of the object
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
* `bcast_rate_limit_enable` - Rate limit the l2 broadcast packet on mgmt port
* `rate` - packets per second. Default is 500. (packets per second. Please specify an even number. Default is 500)
* `address_type` - ‘link-local’: Configure an IPv6 link local address;
* `default_ipv6_gateway` - Set default gateway (Default gateway address)
* `inbound` - ACL applied on incoming packets to this interface
* `ipv6_addr` - Set the IPv6 address of an interface
* `v6_acl_name` - Apply ACL rules to incoming packets on this interface (Named Access List)
* `control_apps_use_mgmt_port` - Control applications use management port
* `default_gateway` - Set default gateway (Default gateway address)
* `dhcp` - Use DHCP to configure IP address
* `ipv4_address` - IP address
* `ipv4_netmask` - IP subnet mask
* `secondary_ip` - Global IP configuration subcommands
* `counters1` - ‘all’: all; ‘packets_input’: Input packets; ‘bytes_input’: Input bytes; ‘received_broadcasts’: Received broadcasts; ‘received_multicasts’: Received multicasts; ‘received_unicasts’: Received unicasts; ‘input_errors’: Input errors; ‘crc’: CRC; ‘frame’: Frames; ‘input_err_short’: Runts; ‘input_err_long’: Giants; ‘packets_output’: Output packets; ‘bytes_output’: Output bytes; ‘transmitted_broadcasts’: Transmitted broadcasts; ‘transmitted_multicasts’: Transmitted multicasts; ‘transmitted_unicasts’: Transmitted unicasts; ‘output_errors’: Output errors; ‘collisions’: Collisions;
* `acl_id` - ACL id
* `acl_name` - Apply an access list (Named Access List)

