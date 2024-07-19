---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_domain_group_domain_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_domain_group_domain_rate: Configure Domain group per domain rate
  PLACEHOLDER
---

# thunder_domain_group_domain_rate (Resource)

`thunder_domain_group_domain_rate`: Configure Domain group per domain rate

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_domain_group_domain_rate" "thunder_domain_group_domain_rate" {
  name       = "test_domain_group"
  dummy_name = "configuration"
  user_tag   = "test"
  domain_list_list {
    name            = "test"
    per_suffix_rate = 3
    user_tag        = "test"
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dummy_name` (String) 'configuration': Configure domain group rate limit;
- `name` (String) Name

### Optional

- `domain_list_list` (Block List) (see [below for nested schema](#nestedblock--domain_list_list))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--domain_list_list"></a>
### Nested Schema for `domain_list_list`

Required:

- `name` (String) Specify name of the domain list

Optional:

- `per_suffix_rate` (Number) Per suffix domain rate limit
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

