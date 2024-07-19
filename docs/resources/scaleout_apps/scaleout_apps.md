---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_scaleout_apps Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_scaleout_apps: Enable Scaleout for apps
  PLACEHOLDER
---

# thunder_scaleout_apps (Resource)

`thunder_scaleout_apps`: Enable Scaleout for apps

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_apps" "thunder_scaleout_apps" {

  enable = 1
  skip_mac_overwrite {
    enable = 1
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `enable` (Number) Enable Scaleout for apps
- `skip_mac_overwrite` (Block List, Max: 1) (see [below for nested schema](#nestedblock--skip_mac_overwrite))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--skip_mac_overwrite"></a>
### Nested Schema for `skip_mac_overwrite`

Optional:

- `enable` (Number) Skips overwriting dest MAC of flooded packets on Active node
- `uuid` (String) uuid of the object

