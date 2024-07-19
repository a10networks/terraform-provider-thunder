---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vpn_ike_gateway_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vpn_ike_gateway_stats: Statistics for the object ike-gateway
  PLACEHOLDER
---

# thunder_vpn_ike_gateway_stats (Data Source)

`thunder_vpn_ike_gateway_stats`: Statistics for the object ike-gateway

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ike_gateway_stats" "thunder_vpn_ike_gateway_stats" {

  name = "test"

}
output "get_vpn_ike_gateway_stats" {
  value = ["${data.thunder_vpn_ike_gateway_stats.thunder_vpn_ike_gateway_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) IKE-gateway name

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `ike_current_version` (Number) IKE version
- `v1_child_sa_invalid_spi` (Number) Invalid SPI for Child SAs
- `v1_in_aggressive_req` (Number) Incoming Aggressive Request
- `v1_in_aggressive_rsp` (Number) Incoming Aggressive Response
- `v1_in_auth_only_req` (Number) Incoming Auth Only Request
- `v1_in_auth_only_rsp` (Number) Incoming Auth Only Response
- `v1_in_id_prot_req` (Number) Incoming ID Protection Request
- `v1_in_id_prot_rsp` (Number) Incoming ID Protection Response
- `v1_in_info_v1_req` (Number) Incoming Info Request
- `v1_in_info_v1_rsp` (Number) Incoming Info Response
- `v1_in_new_group_mode_req` (Number) Incoming New Group Mode Request
- `v1_in_new_group_mode_rsp` (Number) Incoming New Group Mode Response
- `v1_in_quick_mode_req` (Number) Incoming Quick Mode Request
- `v1_in_quick_mode_rsp` (Number) Incoming Quick Mode Response
- `v1_in_transaction_req` (Number) Incoming Transaction Request
- `v1_in_transaction_rsp` (Number) Incoming Transaction Response
- `v1_out_aggressive_req` (Number) Outgoing Aggressive Request
- `v1_out_aggressive_rsp` (Number) Outgoing Aggressive Response
- `v1_out_auth_only_req` (Number) Outgoing Auth Only Request
- `v1_out_auth_only_rsp` (Number) Outgoing Auth Only Response
- `v1_out_id_prot_req` (Number) Outgoing ID Protection Request
- `v1_out_id_prot_rsp` (Number) Outgoing ID Protection Response
- `v1_out_info_v1_req` (Number) Outgoing Info Request
- `v1_out_info_v1_rsp` (Number) Outgoing Info Response
- `v1_out_new_group_mode_req` (Number) Outgoing New Group Mode Request
- `v1_out_new_group_mode_rsp` (Number) Outgoing New Group Mode Response
- `v1_out_quick_mode_req` (Number) Outgoing Quick Mode Request
- `v1_out_quick_mode_rsp` (Number) Outgoing Quick Mode Response
- `v1_out_transaction_req` (Number) Outgoing Transaction Request
- `v1_out_transaction_rsp` (Number) Outgoing Transaction Response
- `v2_child_sa_invalid_spi` (Number) Invalid SPI for Child SAs
- `v2_child_sa_rekey` (Number) Child SA Rekey
- `v2_in_auth_req` (Number) Incoming Auth Request
- `v2_in_auth_rsp` (Number) Incoming Auth Response
- `v2_in_create_child_req` (Number) Incoming Create Child Request
- `v2_in_create_child_rsp` (Number) Incoming Create Child Response
- `v2_in_info_req` (Number) Incoming Info Request
- `v2_in_info_rsp` (Number) Incoming Info Response
- `v2_in_init_req` (Number) Incoming Init Request
- `v2_in_init_rsp` (Number) Incoming Init Response
- `v2_in_invalid` (Number) Incoming Invalid
- `v2_in_invalid_spi` (Number) Incoming Invalid SPI
- `v2_init_rekey` (Number) Initiate Rekey
- `v2_out_auth_req` (Number) Outgoing Auth Request
- `v2_out_auth_rsp` (Number) Outgoing Auth Response
- `v2_out_create_child_req` (Number) Outgoing Create Child Request
- `v2_out_create_child_rsp` (Number) Outgoing Create Child Response
- `v2_out_info_req` (Number) Outgoing Info Request
- `v2_out_info_rsp` (Number) Outgoing Info Response
- `v2_out_init_req` (Number) Outgoing Init Request
- `v2_out_init_rsp` (Number) Outgoing Init Response
- `v2_rsp_rekey` (Number) Respond Rekey

