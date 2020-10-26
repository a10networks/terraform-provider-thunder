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

resource "thunder_interface_management" "testname" {
	lldp {
    enable_cfg {
      rt_enable = 1
      rx = 1
      tx = 1
    }
  }
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

