---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_geoloc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_geoloc:
  PLACEHOLDER
---

# thunder_system_geoloc (Resource)

`thunder_system_geoloc`: 

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_geoloc" "thunder_system_geoloc" {
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'place-holder': place-holder;

