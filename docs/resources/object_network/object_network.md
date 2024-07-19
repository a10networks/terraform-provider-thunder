---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_object_network Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_object_network: Configure Network Object
  PLACEHOLDER
---

# thunder_object_network (Resource)

`thunder_object_network`: Configure Network Object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_object_network" "thunder_object_network" {
  net_name       = "test"
  description    = "109"
  ip_range_end   = "10.10.10.101"
  ip_range_start = "10.10.10.10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `net_name` (String) Network Object Name

### Optional

- `description` (String) Description of the object instance
- `ip_range_end` (String) IPV4 Host address end
- `ip_range_start` (String) IPv4 Host Address start
- `ipv6_range_end` (String) IPV6 Host address end
- `ipv6_range_start` (String) IPv6 Host Address start
- `ipv6_subnet` (String) IPv6 Network Address
- `subnet` (String) IPv4 Network Address
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

