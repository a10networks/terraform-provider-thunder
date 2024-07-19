---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_jwks_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_jwks_oper: Operational Status for the object jwks
  PLACEHOLDER
---

# thunder_aam_authentication_jwks_oper (Data Source)

`thunder_aam_authentication_jwks_oper`: Operational Status for the object jwks

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_jwks_oper" "thunder_aam_authentication_jwks_oper" {
}
output "get_aam_authentication_jwks_oper" {
  value = ["${data.thunder_aam_authentication_jwks_oper.thunder_aam_authentication_jwks_oper}"]
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

- `jwk_list` (Block List) (see [below for nested schema](#nestedblock--oper--jwk_list))

<a id="nestedblock--oper--jwk_list"></a>
### Nested Schema for `oper.jwk_list`

Optional:

- `jwk_name` (String)
- `jwk_size` (Number)

