---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_lsn_all_sessions_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_lsn_all_sessions_oper: Operational Status for the object all-sessions
  PLACEHOLDER
---

# thunder_cgnv6_lsn_all_sessions_oper (Data Source)

`thunder_cgnv6_lsn_all_sessions_oper`: Operational Status for the object all-sessions

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_all_sessions_oper" "thunder_cgnv6_lsn_all_sessions_oper" {
}
output "get_cgnv6_lsn_all_sessions_oper" {
  value = ["${data.thunder_cgnv6_lsn_all_sessions_oper.thunder_cgnv6_lsn_all_sessions_oper}"]
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

- `nat_pool_name` (String)
- `status` (String)

