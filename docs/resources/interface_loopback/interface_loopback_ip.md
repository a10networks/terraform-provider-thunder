---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_loopback_ip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_interface_loopback_ip: Global IP configuration subcommands
  PLACEHOLDER
---

# thunder_interface_loopback_ip (Resource)

`thunder_interface_loopback_ip`: Global IP configuration subcommands

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_ip" "thunder_interface_loopback_ip" {

  ifnum = 4
  address_list {
    ipv4_address = "10.0.11.17"
    ipv4_netmask = "255.255.255.0"
  }
  ospf {
    ospf_global {
      authentication_cfg {
        authentication = 1
      }
      authentication_key = "8"
      bfd_cfg {
        bfd     = 1
        disable = 1
      }
      cost = 57908
      database_filter_cfg {
        database_filter = "all"
        out             = 1
      }
      dead_interval  = 40
      disable        = "all"
      hello_interval = 10
      message_digest_cfg {
        message_digest_key = 197
        md5 {
          md5_value = "16"
        }
      }
      mtu                 = 12994
      mtu_ignore          = 1
      priority            = 1
      retransmit_interval = 5
      transmit_delay      = 1
    }
  }
  rip {
    authentication {
      key_chain {
        key_chain = "g26"
      }
    }
    send_packet    = 1
    receive_packet = 1
    send_cfg {
      send    = 1
      version = "1"
    }
    receive_cfg {
      receive = 1
      version = "1"
    }
    split_horizon_cfg {
      state = "poisoned"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ifnum` (String) Ifnum

### Optional

- `address_list` (Block List) (see [below for nested schema](#nestedblock--address_list))
- `ospf` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf))
- `rip` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip))
- `router` (Block List, Max: 1) (see [below for nested schema](#nestedblock--router))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--address_list"></a>
### Nested Schema for `address_list`

Optional:

- `ipv4_address` (String) IP address
- `ipv4_netmask` (String) IP subnet mask


<a id="nestedblock--ospf"></a>
### Nested Schema for `ospf`

Optional:

- `ospf_global` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf--ospf_global))
- `ospf_ip_list` (Block List) (see [below for nested schema](#nestedblock--ospf--ospf_ip_list))

<a id="nestedblock--ospf--ospf_global"></a>
### Nested Schema for `ospf.ospf_global`

Optional:

- `authentication_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf--ospf_global--authentication_cfg))
- `authentication_key` (String) Authentication password (key) (The OSPF password (key))
- `bfd_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf--ospf_global--bfd_cfg))
- `cost` (Number) Interface cost
- `database_filter_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf--ospf_global--database_filter_cfg))
- `dead_interval` (Number) Interval after which a neighbor is declared dead (Seconds)
- `disable` (String) 'all': All functionality;
- `hello_interval` (Number) Time between HELLO packets (Seconds)
- `message_digest_cfg` (Block List) (see [below for nested schema](#nestedblock--ospf--ospf_global--message_digest_cfg))
- `mtu` (Number) OSPF interface MTU (MTU size)
- `mtu_ignore` (Number) Ignores the MTU in DBD packets
- `priority` (Number) Router priority
- `retransmit_interval` (Number) Time between retransmitting lost link state advertisements (Seconds)
- `transmit_delay` (Number) Link state transmit delay (Seconds)
- `uuid` (String) uuid of the object

<a id="nestedblock--ospf--ospf_global--authentication_cfg"></a>
### Nested Schema for `ospf.ospf_global.authentication_cfg`

Optional:

- `authentication` (Number) Enable authentication
- `value` (String) 'message-digest': Use message-digest authentication; 'null': Use no authentication;


<a id="nestedblock--ospf--ospf_global--bfd_cfg"></a>
### Nested Schema for `ospf.ospf_global.bfd_cfg`

Optional:

- `bfd` (Number) Bidirectional Forwarding Detection (BFD)
- `disable` (Number) Disable BFD


<a id="nestedblock--ospf--ospf_global--database_filter_cfg"></a>
### Nested Schema for `ospf.ospf_global.database_filter_cfg`

Optional:

- `database_filter` (String) 'all': Filter all LSA;
- `out` (Number) Outgoing LSA


<a id="nestedblock--ospf--ospf_global--message_digest_cfg"></a>
### Nested Schema for `ospf.ospf_global.message_digest_cfg`

Optional:

- `md5` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf--ospf_global--message_digest_cfg--md5))
- `message_digest_key` (Number) Message digest authentication password (key) (Key id)

<a id="nestedblock--ospf--ospf_global--message_digest_cfg--md5"></a>
### Nested Schema for `ospf.ospf_global.message_digest_cfg.md5`

Optional:

- `md5_value` (String) The OSPF password (1-16)




<a id="nestedblock--ospf--ospf_ip_list"></a>
### Nested Schema for `ospf.ospf_ip_list`

Required:

- `ip_addr` (String) Address of interface

Optional:

- `authentication` (Number) Enable authentication
- `authentication_key` (String) Authentication password (key) (The OSPF password (key))
- `cost` (Number) Interface cost
- `database_filter` (String) 'all': Filter all LSA;
- `dead_interval` (Number) Interval after which a neighbor is declared dead (Seconds)
- `hello_interval` (Number) Time between HELLO packets (Seconds)
- `message_digest_cfg` (Block List) (see [below for nested schema](#nestedblock--ospf--ospf_ip_list--message_digest_cfg))
- `mtu_ignore` (Number) Ignores the MTU in DBD packets
- `out` (Number) Outgoing LSA
- `priority` (Number) Router priority
- `retransmit_interval` (Number) Time between retransmitting lost link state advertisements (Seconds)
- `transmit_delay` (Number) Link state transmit delay (Seconds)
- `uuid` (String) uuid of the object
- `value` (String) 'message-digest': Use message-digest authentication; 'null': Use no authentication;

<a id="nestedblock--ospf--ospf_ip_list--message_digest_cfg"></a>
### Nested Schema for `ospf.ospf_ip_list.message_digest_cfg`

Optional:

- `md5_value` (String) The OSPF password (1-16)
- `message_digest_key` (Number) Message digest authentication password (key) (Key id)




<a id="nestedblock--rip"></a>
### Nested Schema for `rip`

Optional:

- `authentication` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip--authentication))
- `receive_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip--receive_cfg))
- `receive_packet` (Number) Enable receiving packet through the specified interface
- `send_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip--send_cfg))
- `send_packet` (Number) Enable sending packets through the specified interface
- `split_horizon_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip--split_horizon_cfg))
- `uuid` (String) uuid of the object

<a id="nestedblock--rip--authentication"></a>
### Nested Schema for `rip.authentication`

Optional:

- `key_chain` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip--authentication--key_chain))
- `mode` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip--authentication--mode))
- `str` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip--authentication--str))

<a id="nestedblock--rip--authentication--key_chain"></a>
### Nested Schema for `rip.authentication.key_chain`

Optional:

- `key_chain` (String) Authentication key-chain (Name of key-chain)


<a id="nestedblock--rip--authentication--mode"></a>
### Nested Schema for `rip.authentication.mode`

Optional:

- `mode` (String) 'md5': Keyed message digest; 'text': Clear text authentication;


<a id="nestedblock--rip--authentication--str"></a>
### Nested Schema for `rip.authentication.str`

Optional:

- `string` (String) The RIP authentication string



<a id="nestedblock--rip--receive_cfg"></a>
### Nested Schema for `rip.receive_cfg`

Optional:

- `receive` (Number) Advertisement reception
- `version` (String) '1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;


<a id="nestedblock--rip--send_cfg"></a>
### Nested Schema for `rip.send_cfg`

Optional:

- `send` (Number) Advertisement transmission
- `version` (String) '1': RIP version 1; '2': RIP version 2; '1-compatible': RIPv1-compatible; '1-2': RIP version 1 & 2;


<a id="nestedblock--rip--split_horizon_cfg"></a>
### Nested Schema for `rip.split_horizon_cfg`

Optional:

- `state` (String) 'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;



<a id="nestedblock--router"></a>
### Nested Schema for `router`

Optional:

- `isis` (Block List, Max: 1) (see [below for nested schema](#nestedblock--router--isis))

<a id="nestedblock--router--isis"></a>
### Nested Schema for `router.isis`

Optional:

- `tag` (String) ISO routing area tag
- `uuid` (String) uuid of the object

