---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_app_protocol_port_monitor Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_app_protocol_port_monitor: Monitors application ports
  PLACEHOLDER
---

# thunder_ip_app_protocol_port_monitor (Resource)

`thunder_ip_app_protocol_port_monitor`: Monitors application ports

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_app_protocol_port_monitor" "thunder_ip_app_protocol_port_monitor" {
  disable = 0
  enable  = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `disable` (Number) Enables global port processing
- `enable` (Number) Enables global port processing
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

