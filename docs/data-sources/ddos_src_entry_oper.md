---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_src_entry_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_src_entry_oper: Operational Status for the object entry
  PLACEHOLDER
---

# thunder_ddos_src_entry_oper (Data Source)

`thunder_ddos_src_entry_oper`: Operational Status for the object entry

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_src_entry_oper" "thunder_ddos_src_entry_oper" {
  src_entry_name = "2"

}
output "get_ddos_src_entry_oper" {
  value = ["${data.thunder_ddos_src_entry_oper.thunder_ddos_src_entry_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `src_entry_name` (String)

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `all_entries` (Number)
- `all_l4_types` (Number)
- `app_type` (String)
- `authenticated` (Number)
- `black_listed` (Number)
- `ddos_entry_list` (Block List) (see [below for nested schema](#nestedblock--oper--ddos_entry_list))
- `entry_displayed_count` (Number)
- `exceeded` (Number)
- `hw_blacklisted` (String)
- `l4_ext_rate` (String)
- `l4_type_str` (String)
- `service_displayed_count` (Number)
- `white_listed` (Number)

<a id="nestedblock--oper--ddos_entry_list"></a>
### Nested Schema for `oper.ddos_entry_list`

Optional:

- `age_str` (String)
- `app_stat1_limit` (String)
- `app_stat2_limit` (String)
- `app_stat3_limit` (String)
- `app_stat4_limit` (String)
- `app_stat5_limit` (String)
- `app_stat6_limit` (String)
- `app_stat7_limit` (String)
- `app_stat8_limit` (String)
- `connection_limit` (String)
- `connection_rate_limit` (String)
- `current_app_stat1` (String)
- `current_app_stat2` (String)
- `current_app_stat3` (String)
- `current_app_stat4` (String)
- `current_app_stat5` (String)
- `current_app_stat6` (String)
- `current_app_stat7` (String)
- `current_app_stat8` (String)
- `current_connection_rate` (String)
- `current_connections` (String)
- `current_frag_packet_rate` (String)
- `current_kbit_rate` (String)
- `current_packet_rate` (String)
- `debug_str` (String)
- `dst_address_str` (String)
- `dynamic_entry_count` (String)
- `dynamic_entry_limit` (String)
- `frag_packet_rate_limit` (String)
- `kbit_rate_limit` (String)
- `level_str` (String)
- `lockup_time_str` (String)
- `packet_rate_limit` (String)
- `port_str` (String)
- `sflow_source_id` (String)
- `src_address_str` (String)
- `state_str` (String)

