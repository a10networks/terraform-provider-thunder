---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_monitored_entity_detail Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_monitored_entity_detail: Display Monitoring entity detail
  PLACEHOLDER
---

# thunder_visibility_monitored_entity_detail (Resource)

`thunder_visibility_monitored_entity_detail`: Display Monitoring entity detail

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_monitored_entity_detail" "thunder_visibility_monitored_entity_detail" {

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `debug` (Block List, Max: 1) (see [below for nested schema](#nestedblock--debug))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--debug"></a>
### Nested Schema for `debug`

Optional:

- `uuid` (String) uuid of the object

