---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_system_status_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_system_status_oper: Operational Status for the object system-status
  PLACEHOLDER
---

# thunder_fw_system_status_oper (Data Source)

`thunder_fw_system_status_oper`: Operational Status for the object system-status

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_system_status_oper" "thunder_fw_system_status_oper" {

}
output "get_fw_system_status_oper" {
  value = ["${data.thunder_fw_system_status_oper.thunder_fw_system_status_oper}"]
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

- `data_sessions_free` (Number)
- `data_sessions_used` (Number)
- `radius_entries_free` (Number)
- `radius_entries_used` (Number)
- `smp_sessions_free` (Number)
- `smp_sessions_used` (Number)

