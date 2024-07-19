---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_password_retry_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_password_retry_oper: Operational Status for the object password-retry
  PLACEHOLDER
---

# thunder_aam_authentication_password_retry_oper (Data Source)

`thunder_aam_authentication_password_retry_oper`: Operational Status for the object password-retry

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_password_retry_oper" "thunder_aam_authentication_password_retry_oper" {
}
output "get_aam_authentication_password_retry_oper" {
  value = ["${data.thunder_aam_authentication_password_retry_oper.thunder_aam_authentication_password_retry_oper}"]
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

- `entry_list` (Block List) (see [below for nested schema](#nestedblock--oper--entry_list))
- `logon_name` (String)
- `name` (String)

<a id="nestedblock--oper--entry_list"></a>
### Nested Schema for `oper.entry_list`

Optional:

- `account` (String)
- `locked_out` (String)
- `logon` (String)
- `pw_failure` (Number)
- `ttl` (Number)

