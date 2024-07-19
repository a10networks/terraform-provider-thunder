---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vcs_showdebug_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vcs_showdebug_oper: Operational Status for the object showdebug
  PLACEHOLDER
---

# thunder_vcs_showdebug_oper (Data Source)

`thunder_vcs_showdebug_oper`: Operational Status for the object showdebug

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vcs_showdebug_oper" "thunder_vcs_showdebug_oper" {
}
output "get_vcs_showdebug_oper" {
  value = ["${data.thunder_vcs_showdebug_oper.thunder_vcs_showdebug_oper}"]
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

- `debugging_buff_size` (Number)
- `debugging_switches` (String)

