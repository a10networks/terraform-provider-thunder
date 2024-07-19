---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_src_based_policy_policy_class_list Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_src_based_policy_policy_class_list: Configure class-list
  PLACEHOLDER
---

# thunder_ddos_src_based_policy_policy_class_list (Resource)

`thunder_ddos_src_based_policy_policy_class_list`: Configure class-list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_based_policy_policy_class_list" "thunder_ddos_src_based_policy_policy_class_list" {
  name            = "test"
  class_list_name = "10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `class_list_name` (String) Class-list name
- `name` (String) Name

### Optional

- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

