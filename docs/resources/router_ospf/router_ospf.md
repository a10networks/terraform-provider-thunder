---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_ospf Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_router_ospf: Open Shortest Path First (OSPF)
  PLACEHOLDER
---

# thunder_router_ospf (Resource)

`thunder_router_ospf`: Open Shortest Path First (OSPF)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_ospf" "thunder_router_ospf" {
  process_id                    = 400
  auto_cost_reference_bandwidth = 30000
  bfd_all_interfaces            = 1
  rfc1583_compatible            = 1
  default_metric                = 30

  distance {
    distance_value = 10
    distance_ospf {
      distance_external   = 20
      distance_inter_area = 30
      distance_intra_area = 40
    }
  }
  distribute_internal_list {
    di_area_ipv4 = "1.1.1.1"
    di_area_num  = 200
    di_cost      = 3000
    di_type      = "lw4o6"
  }
  ha_standby_extra_cost {
    extra_cost = 500
    group      = 2
  }
  router_id {
    value = "5.5.5.5"
  }
  ospf_1 {
    abr_type {
      option = "standard"
    }
  }
  log_adjacency_changes_cfg {
    state = "detail"
  }
  max_concurrent_dd = 10
  maximum_area      = 2000
  user_tag          = "ospf"
  neighbor_list {
    address = "1.1.1.1"
    cost    = 400
  }
  distribute_lists {
    direction = "in"
    value     = "ACL1"
  }
  host_list {
    area_cfg {
      area_ipv4 = "0.0.0.0"
      area_num  = 0
      cost      = 65535
    }
    host_address = "2.2.2.2"
  }
  overflow {
    database {
      limit = "hard"
    }
  }
  passive_interface {
    eth_cfg {
      ethernet    = 2
      eth_address = "2.2.2.2"
    }
    lif_cfg {
      lif         = 3
      lif_address = "4.4.4.4"
    }
    ve_cfg {
      ve         = 300
      ve_address = "5.5.5.5"
    }
    trunk_cfg {
      trunk         = 5
      trunk_address = "6.6.6.6"
    }
    tunnel_cfg {
      tunnel         = 20
      tunnel_address = "12.12.12.13"
    }
    loopback_cfg {
      loopback         = 1
      loopback_address = "11.11.11.11"
    }

  }
  summary_address_list {
    summary_address = "20.20.20.0/24"
    tag             = 3000
    not_advertise   = 0
  }
  timers {
    spf {
      exp {
        min_delay = 3000
        max_delay = 40000
      }
    }
  }
  network_list {
    network_ipv4      = "1.1.1.1"
    network_ipv4_mask = "255.255.255.0"
    network_area {
      network_area_ipv4 = "0.0.0.0"
      network_area_num  = 0
      instance_value    = 0
    }
  }
  default_information {
    originate   = 1
    always      = 1
    metric      = 40
    metric_type = 1
    route_map   = "MAP"
  }
  redistribute {
    ip_nat             = 1
    metric_ip_nat      = 40
    metric_type_ip_nat = 1
    route_map_ip_nat   = "NAT"
    tag_ip_nat         = 800
    ip_nat_floating_list {
      ip_nat_prefix              = "9.9.9.9/24"
      ip_nat_floating_ip_forward = "2.2.2.2"
    }
    vip_list {
      type_vip        = "only-flagged"
      metric_type_vip = 1
      metric_vip      = 800
      route_map_vip   = "VIP"
      tag_vip         = 500
    }
    vip_floating_list {
      vip_address             = "5.5.5.5"
      vip_floating_ip_forward = "6.6.6.6"
    }
    ospf_list {
      process_id       = 200
      ospf             = 1
      metric_ospf      = 50
      metric_type_ospf = 1
      route_map_ospf   = "OSPF"
      tag_ospf         = 500
    }
  }
  area_list {
    area_ipv4 = "0.0.0.2"
    area_num  = 0
    auth_cfg {
      authentication = 1
      message_digest = 1
    }
    filter_lists {
      filter_list     = 1
      acl_name        = "ACL1"
      acl_direction   = "in"
      plist_name      = "PREFIX1"
      plist_direction = "in"
    }
    default_cost = 1

    range_list {
      area_range_prefix = "1.1.1.1/24"
      option            = "advertise"
    }
    virtual_link_list {
      virtual_link_ip_addr        = "3.3.3.3"
      bfd                         = 1
      hello_interval              = 10
      dead_interval               = 40
      retransmit_interval         = 300
      transmit_delay              = 50
      virtual_link_authentication = 1
      virtual_link_auth_type      = "message-digest"
      authentication_key          = "string"
      message_digest_key          = 20
      md5                         = "MD5"
    }
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `process_id` (Number) OSPF process ID

### Optional

- `area_list` (Block List) (see [below for nested schema](#nestedblock--area_list))
- `auto_cost_reference_bandwidth` (Number) Use reference bandwidth method to assign OSPF cost (The reference bandwidth in terms of Mbits per second)
- `bfd_all_interfaces` (Number) Enable BFD on all interfaces
- `default_information` (Block List, Max: 1) (see [below for nested schema](#nestedblock--default_information))
- `default_metric` (Number) Set metric of redistributed routes (Default metric)
- `distance` (Block List, Max: 1) (see [below for nested schema](#nestedblock--distance))
- `distribute_internal_list` (Block List) (see [below for nested schema](#nestedblock--distribute_internal_list))
- `distribute_lists` (Block List) (see [below for nested schema](#nestedblock--distribute_lists))
- `extern_lsa_equivalence_check` (Number) external LSA equivalance check
- `ha_standby_extra_cost` (Block List) (see [below for nested schema](#nestedblock--ha_standby_extra_cost))
- `host_list` (Block List) (see [below for nested schema](#nestedblock--host_list))
- `log_adjacency_changes_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--log_adjacency_changes_cfg))
- `max_concurrent_dd` (Number) Maximum number allowed to process DD concurrently (Number of DD process)
- `maximum_area` (Number) Maximum number of non-backbone areas (OSPF area limit)
- `neighbor_list` (Block List) (see [below for nested schema](#nestedblock--neighbor_list))
- `network_list` (Block List) (see [below for nested schema](#nestedblock--network_list))
- `ospf_1` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf_1))
- `overflow` (Block List, Max: 1) (see [below for nested schema](#nestedblock--overflow))
- `passive_interface` (Block List, Max: 1) (see [below for nested schema](#nestedblock--passive_interface))
- `redistribute` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute))
- `rfc1583_compatible` (Number) Compatible with RFC 1583
- `router_id` (Block List, Max: 1) (see [below for nested schema](#nestedblock--router_id))
- `summary_address_list` (Block List) (see [below for nested schema](#nestedblock--summary_address_list))
- `timers` (Block List, Max: 1) (see [below for nested schema](#nestedblock--timers))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--area_list"></a>
### Nested Schema for `area_list`

Required:

- `area_ipv4` (String) OSPF area ID in IP address format
- `area_num` (Number) OSPF area ID as a decimal value

Optional:

- `auth_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--area_list--auth_cfg))
- `default_cost` (Number) Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)
- `filter_lists` (Block List) (see [below for nested schema](#nestedblock--area_list--filter_lists))
- `nssa_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--area_list--nssa_cfg))
- `range_list` (Block List) (see [below for nested schema](#nestedblock--area_list--range_list))
- `shortcut` (String) 'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;
- `stub_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--area_list--stub_cfg))
- `uuid` (String) uuid of the object
- `virtual_link_list` (Block List) (see [below for nested schema](#nestedblock--area_list--virtual_link_list))

<a id="nestedblock--area_list--auth_cfg"></a>
### Nested Schema for `area_list.auth_cfg`

Optional:

- `authentication` (Number) Enable authentication
- `message_digest` (Number) Use message-digest authentication


<a id="nestedblock--area_list--filter_lists"></a>
### Nested Schema for `area_list.filter_lists`

Optional:

- `acl_direction` (String) 'in': Filter networks sent to this area; 'out': Filter networks sent from this area;
- `acl_name` (String) Filter networks by access-list (Name of an access-list)
- `filter_list` (Number) Filter networks between OSPF areas
- `plist_direction` (String) 'in': Filter networks sent to this area; 'out': Filter networks sent from this area;
- `plist_name` (String) Filter networks by prefix-list (Name of an IP prefix-list)


<a id="nestedblock--area_list--nssa_cfg"></a>
### Nested Schema for `area_list.nssa_cfg`

Optional:

- `default_information_originate` (Number) Originate Type 7 default into NSSA area
- `metric` (Number) OSPF default metric (OSPF metric)
- `metric_type` (Number) OSPF metric type (OSPF metric type for default routes)
- `no_redistribution` (Number) No redistribution into this NSSA area
- `no_summary` (Number) Do not send summary LSA into NSSA
- `nssa` (Number) Specify a NSSA area
- `translator_role` (String) 'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;


<a id="nestedblock--area_list--range_list"></a>
### Nested Schema for `area_list.range_list`

Optional:

- `area_range_prefix` (String) Area range for IPv4 prefix
- `option` (String) 'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;


<a id="nestedblock--area_list--stub_cfg"></a>
### Nested Schema for `area_list.stub_cfg`

Optional:

- `no_summary` (Number) Do not inject inter-area routes into area
- `stub` (Number) Configure OSPF area as stub


<a id="nestedblock--area_list--virtual_link_list"></a>
### Nested Schema for `area_list.virtual_link_list`

Optional:

- `authentication_key` (String) Set authentication key (Authentication key (8 chars))
- `bfd` (Number) Bidirectional Forwarding Detection (BFD)
- `dead_interval` (Number) Dead router detection time (Seconds)
- `hello_interval` (Number) Hello packet interval (Seconds)
- `md5` (String) Use MD5 algorithm (Authentication key (16 chars))
- `message_digest_key` (Number) Set message digest key (Key ID)
- `retransmit_interval` (Number) LSA retransmit interval (Seconds)
- `transmit_delay` (Number) LSA transmission delay (Seconds)
- `virtual_link_auth_type` (String) 'message-digest': Use message-digest authentication; 'null': Use null authentication;
- `virtual_link_authentication` (Number) Enable authentication
- `virtual_link_ip_addr` (String) ID (IP addr) associated with virtual link neighbor



<a id="nestedblock--default_information"></a>
### Nested Schema for `default_information`

Optional:

- `always` (Number) Always advertise default route
- `metric` (Number) OSPF default metric (OSPF metric)
- `metric_type` (Number) OSPF metric type for default routes
- `originate` (Number) Distribute a default route
- `route_map` (String) Route map reference (Pointer to route-map entries)
- `uuid` (String) uuid of the object


<a id="nestedblock--distance"></a>
### Nested Schema for `distance`

Optional:

- `distance_ospf` (Block List, Max: 1) (see [below for nested schema](#nestedblock--distance--distance_ospf))
- `distance_value` (Number) OSPF Administrative distance

<a id="nestedblock--distance--distance_ospf"></a>
### Nested Schema for `distance.distance_ospf`

Optional:

- `distance_external` (Number) External routes (Distance for external routes)
- `distance_inter_area` (Number) Inter-area routes (Distance for inter-area routes)
- `distance_intra_area` (Number) Intra-area routes (Distance for intra-area routes)



<a id="nestedblock--distribute_internal_list"></a>
### Nested Schema for `distribute_internal_list`

Optional:

- `di_area_ipv4` (String) OSPF area ID as a IP address format
- `di_area_num` (Number) OSPF area ID as a decimal value
- `di_cost` (Number) Cost of route
- `di_type` (String) 'lw4o6': LW4O6 Prefix; 'floating-ip': Floating IP; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'vip': Only not flagged Virtual IP (VIP); 'vip-only-flagged': Selected Virtual IP (VIP);


<a id="nestedblock--distribute_lists"></a>
### Nested Schema for `distribute_lists`

Optional:

- `direction` (String) 'in': Filter incoming routing updates; 'out': Filter outgoing routing updates;
- `option` (String) 'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;
- `ospf_id` (Number) OSPF process ID
- `protocol` (String) 'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'lw4o6': LW4O6 Prefix; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'ospf': Open Shortest Path First (OSPF); 'rip': Routing Information Protocol (RIP); 'static': Static routes;
- `value` (String) Access-list name


<a id="nestedblock--ha_standby_extra_cost"></a>
### Nested Schema for `ha_standby_extra_cost`

Optional:

- `extra_cost` (Number) The extra cost value
- `group` (Number) Group (Group ID)


<a id="nestedblock--host_list"></a>
### Nested Schema for `host_list`

Optional:

- `area_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host_list--area_cfg))
- `host_address` (String) Host address

<a id="nestedblock--host_list--area_cfg"></a>
### Nested Schema for `host_list.area_cfg`

Optional:

- `area_ipv4` (String) OSPF area ID in IP address format
- `area_num` (Number) OSPF area ID as a decimal value
- `cost` (Number) Cost of host



<a id="nestedblock--log_adjacency_changes_cfg"></a>
### Nested Schema for `log_adjacency_changes_cfg`

Optional:

- `state` (String) 'detail': Log changes in adjacency state; 'disable': Disable logging;


<a id="nestedblock--neighbor_list"></a>
### Nested Schema for `neighbor_list`

Optional:

- `address` (String) Neighbor address
- `cost` (Number) OSPF cost for point-to-multipoint neighbor (Metric)
- `poll_interval` (Number) OSPF dead-router polling interval (Seconds)
- `priority` (Number) OSPF priority of non-broadcast neighbor


<a id="nestedblock--network_list"></a>
### Nested Schema for `network_list`

Optional:

- `network_area` (Block List, Max: 1) (see [below for nested schema](#nestedblock--network_list--network_area))
- `network_ipv4` (String) Network number
- `network_ipv4_cidr` (String) OSPF network prefix
- `network_ipv4_mask` (String) OSPF wild card bits

<a id="nestedblock--network_list--network_area"></a>
### Nested Schema for `network_list.network_area`

Optional:

- `instance_value` (Number) Instance ID
- `network_area_ipv4` (String) OSPF area ID in IP address format
- `network_area_num` (Number) OSPF area ID as a decimal value



<a id="nestedblock--ospf_1"></a>
### Nested Schema for `ospf_1`

Optional:

- `abr_type` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf_1--abr_type))

<a id="nestedblock--ospf_1--abr_type"></a>
### Nested Schema for `ospf_1.abr_type`

Optional:

- `option` (String) 'cisco': Alternative ABR, Cisco implementation (RFC3509); 'ibm': Alternative ABR, IBM implementation (RFC3509); 'shortcut': Shortcut ABR; 'standard': Standard behavior (RFC2328);



<a id="nestedblock--overflow"></a>
### Nested Schema for `overflow`

Optional:

- `database` (Block List, Max: 1) (see [below for nested schema](#nestedblock--overflow--database))

<a id="nestedblock--overflow--database"></a>
### Nested Schema for `overflow.database`

Optional:

- `count1` (Number) Maximum number of LSAs
- `db_external` (Number) Maximum number of LSAs
- `limit` (String) 'hard': Hard limit: Instance will be shutdown if exceeded; 'soft': Soft limit: Warning will be given if exceeded;
- `recovery_time` (Number) Time to recover (0 not recover)



<a id="nestedblock--passive_interface"></a>
### Nested Schema for `passive_interface`

Optional:

- `eth_cfg` (Block List) (see [below for nested schema](#nestedblock--passive_interface--eth_cfg))
- `lif_cfg` (Block List) (see [below for nested schema](#nestedblock--passive_interface--lif_cfg))
- `loopback_cfg` (Block List) (see [below for nested schema](#nestedblock--passive_interface--loopback_cfg))
- `trunk_cfg` (Block List) (see [below for nested schema](#nestedblock--passive_interface--trunk_cfg))
- `tunnel_cfg` (Block List) (see [below for nested schema](#nestedblock--passive_interface--tunnel_cfg))
- `ve_cfg` (Block List) (see [below for nested schema](#nestedblock--passive_interface--ve_cfg))

<a id="nestedblock--passive_interface--eth_cfg"></a>
### Nested Schema for `passive_interface.eth_cfg`

Optional:

- `eth_address` (String) Address of Interface
- `ethernet` (Number) Ethernet interface (Port number)


<a id="nestedblock--passive_interface--lif_cfg"></a>
### Nested Schema for `passive_interface.lif_cfg`

Optional:

- `lif` (String) Logical interface (Lif interface name)
- `lif_address` (String) Address of Interface


<a id="nestedblock--passive_interface--loopback_cfg"></a>
### Nested Schema for `passive_interface.loopback_cfg`

Optional:

- `loopback` (Number) Loopback interface (Port number)
- `loopback_address` (String) Address of Interface


<a id="nestedblock--passive_interface--trunk_cfg"></a>
### Nested Schema for `passive_interface.trunk_cfg`

Optional:

- `trunk` (Number) Trunk interface (Trunk interface number)
- `trunk_address` (String) Address of Interface


<a id="nestedblock--passive_interface--tunnel_cfg"></a>
### Nested Schema for `passive_interface.tunnel_cfg`

Optional:

- `tunnel` (Number) Tunnel interface (Tunnel interface number)
- `tunnel_address` (String) Address of Interface


<a id="nestedblock--passive_interface--ve_cfg"></a>
### Nested Schema for `passive_interface.ve_cfg`

Optional:

- `ve` (Number) Virtual ethernet interface (Virtual ethernet interface number)
- `ve_address` (String) Address of Interface



<a id="nestedblock--redistribute"></a>
### Nested Schema for `redistribute`

Optional:

- `ip_nat` (Number) IP-NAT
- `ip_nat_floating_list` (Block List) (see [below for nested schema](#nestedblock--redistribute--ip_nat_floating_list))
- `metric_ip_nat` (Number) OSPF default metric (OSPF metric)
- `metric_type_ip_nat` (String) '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
- `ospf_list` (Block List) (see [below for nested schema](#nestedblock--redistribute--ospf_list))
- `redist_list` (Block List) (see [below for nested schema](#nestedblock--redistribute--redist_list))
- `route_map_ip_nat` (String) Route map reference (Pointer to route-map entries)
- `tag_ip_nat` (Number) Set tag for routes redistributed into OSPF (32-bit tag value)
- `uuid` (String) uuid of the object
- `vip_floating_list` (Block List) (see [below for nested schema](#nestedblock--redistribute--vip_floating_list))
- `vip_list` (Block List) (see [below for nested schema](#nestedblock--redistribute--vip_list))

<a id="nestedblock--redistribute--ip_nat_floating_list"></a>
### Nested Schema for `redistribute.ip_nat_floating_list`

Optional:

- `ip_nat_floating_ip_forward` (String) Floating-IP as forward address
- `ip_nat_prefix` (String) Address


<a id="nestedblock--redistribute--ospf_list"></a>
### Nested Schema for `redistribute.ospf_list`

Optional:

- `metric_ospf` (Number) OSPF default metric (OSPF metric)
- `metric_type_ospf` (String) '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
- `ospf` (Number) Open Shortest Path First (OSPF)
- `process_id` (Number) OSPF process ID
- `route_map_ospf` (String) Route map reference (Pointer to route-map entries)
- `tag_ospf` (Number) Set tag for routes redistributed into OSPF (32-bit tag value)


<a id="nestedblock--redistribute--redist_list"></a>
### Nested Schema for `redistribute.redist_list`

Optional:

- `metric` (Number) OSPF default metric (OSPF metric)
- `metric_type` (String) '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
- `route_map` (String) Route map reference (Pointer to route-map entries)
- `tag` (Number) Set tag for routes redistributed into OSPF (32-bit tag value)
- `type` (String) 'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;


<a id="nestedblock--redistribute--vip_floating_list"></a>
### Nested Schema for `redistribute.vip_floating_list`

Optional:

- `vip_address` (String) Address
- `vip_floating_ip_forward` (String) Floating-IP as forward address


<a id="nestedblock--redistribute--vip_list"></a>
### Nested Schema for `redistribute.vip_list`

Optional:

- `metric_type_vip` (String) '1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;
- `metric_vip` (Number) OSPF default metric (OSPF metric)
- `route_map_vip` (String) Route map reference (Pointer to route-map entries)
- `tag_vip` (Number) Set tag for routes redistributed into OSPF (32-bit tag value)
- `type_vip` (String) 'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;



<a id="nestedblock--router_id"></a>
### Nested Schema for `router_id`

Optional:

- `value` (String) OSPF router-id in IPv4 address format


<a id="nestedblock--summary_address_list"></a>
### Nested Schema for `summary_address_list`

Optional:

- `not_advertise` (Number) Suppress routes that match the prefix
- `summary_address` (String) Configure IP address summaries (Summary prefix)
- `tag` (Number) Set tag (32-bit tag value)


<a id="nestedblock--timers"></a>
### Nested Schema for `timers`

Optional:

- `spf` (Block List, Max: 1) (see [below for nested schema](#nestedblock--timers--spf))

<a id="nestedblock--timers--spf"></a>
### Nested Schema for `timers.spf`

Optional:

- `exp` (Block List, Max: 1) (see [below for nested schema](#nestedblock--timers--spf--exp))

<a id="nestedblock--timers--spf--exp"></a>
### Nested Schema for `timers.spf.exp`

Optional:

- `max_delay` (Number) Maximum delay between receiving a change to SPF calculation in milliseconds
- `min_delay` (Number) Minimum delay between receiving a change to SPF calculation in milliseconds

