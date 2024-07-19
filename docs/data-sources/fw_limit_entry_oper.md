---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_limit_entry_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_limit_entry_oper: Operational Status for the object limit-entry
  PLACEHOLDER
---

# thunder_fw_limit_entry_oper (Data Source)

`thunder_fw_limit_entry_oper`: Operational Status for the object limit-entry

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_limit_entry_oper" "thunder_fw_limit_entry_oper" {

}
output "get_fw_limit_entry_oper" {
  value = ["${data.thunder_fw_limit_entry_oper.thunder_fw_limit_entry_oper}"]
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

- `limit_entry_count` (Number)
- `limit_entry_list` (Block List) (see [below for nested schema](#nestedblock--oper--limit_entry_list))
- `prefix4` (String)
- `prefix6` (String)
- `prefix_len4` (Number)
- `prefix_len6` (Number)

<a id="nestedblock--oper--limit_entry_list"></a>
### Nested Schema for `oper.limit_entry_list`

Optional:

- `address` (String)
- `curr_count` (Number)
- `max_count` (Number)
- `prefix_len` (Number)
- `rule_name` (String)
- `type` (String)

