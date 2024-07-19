---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_radius_server_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_radius_server_oper: Operational Status for the object server
  PLACEHOLDER
---

# thunder_fw_radius_server_oper (Data Source)

`thunder_fw_radius_server_oper`: Operational Status for the object server

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_radius_server_oper" "thunder_fw_radius_server_oper" {

}
output "get_fw_radius_server_oper" {
  value = ["${data.thunder_fw_radius_server_oper.thunder_fw_radius_server_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `case_insensitive` (Number)
- `custom_attr_name` (String)
- `custom_attr_value` (String)
- `radius_table_entries_list` (Block List) (see [below for nested schema](#nestedblock--oper--radius_table_entries_list))
- `starts_with` (Number)
- `total_entries` (Number)

<a id="nestedblock--oper--radius_table_entries_list"></a>
### Nested Schema for `oper.radius_table_entries_list`

Optional:

- `custom1_attr_value` (String)
- `custom2_attr_value` (String)
- `custom3_attr_value` (String)
- `custom4_attr_value` (String)
- `custom5_attr_value` (String)
- `custom6_attr_value` (String)
- `imei` (String)
- `imsi` (String)
- `inside_ip` (String)
- `inside_ipv6` (String)
- `is_obsolete` (Number)
- `msisdn` (String)
- `prefix_len` (Number)

