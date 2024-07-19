---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_network_twamp_responder_ipv6_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_network_twamp_responder_ipv6_stats: Statistics for the object ipv6
  PLACEHOLDER
---

# thunder_network_twamp_responder_ipv6_stats (Data Source)

`thunder_network_twamp_responder_ipv6_stats`: Statistics for the object ipv6

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_twamp_responder_ipv6_stats" "thunder_network_twamp_responder_ipv6_stats" {
}
output "get_network_twamp_responder_ipv6_stats" {
  value = ["${data.thunder_network_twamp_responder_ipv6_stats.thunder_network_twamp_responder_ipv6_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `no_route_err_v6` (Number) Tx IPv6 no route error drop
- `other_err_v6` (Number) IPv6 Other error drop
- `rx_acl_drop_v6` (Number) Rx IPv6 client-list drop
- `rx_drop_not_enabled_v6` (Number) Rx IPv6 disabled drop
- `rx_pkts_v6` (Number) Rx IPv6 TWAMP test packets
- `twamp_hdr_len_err_v6` (Number) Rx IPv6 TWAMP hdr length error drop
- `tx_pkts_v6` (Number) Tx IPv6 TWAMP test packets

