---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_cpu_packet_prio_support Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_cpu_packet_prio_support: Enable/Disable CPU packet prioritization support
  PLACEHOLDER
---

# thunder_system_cpu_packet_prio_support (Resource)

`thunder_system_cpu_packet_prio_support`: Enable/Disable CPU packet prioritization support

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_cpu_packet_prio_support" "thunder_system_cpu_packet_prio_support" {
  disable = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `disable` (Number) Disable CPU packet prioritization Support
- `enable` (Number) Enable CPU packet prioritization Support

### Read-Only

- `id` (String) The ID of this resource.

