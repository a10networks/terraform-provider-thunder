---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_template_limit_policy_limit_cps Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_template_limit_policy_limit_cps: Enable Connections Per Second Rate Limit
  PLACEHOLDER
---

# thunder_template_limit_policy_limit_cps (Resource)

`thunder_template_limit_policy_limit_cps`: Enable Connections Per Second Rate Limit

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_limit_policy_limit_cps" "thunder_template_limit_policy_limit_cps" {

  policy_number = 2
  value         = 22
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy_number` (String) PolicyNumber

### Optional

- `burstsize` (Number) CPS Token Bucket Size (Must Exceed Configured Rate) (In Connections per second)
- `relaxed` (Number) Relax the limitation when the policy has more tokens from the parent of policy
- `uuid` (String) uuid of the object
- `value` (Number) Connections Per Second Rate Limit (Number of Connections per second)

### Read-Only

- `id` (String) The ID of this resource.

