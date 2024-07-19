---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_nat64_alg_pptp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_nat64_alg_pptp: NAT64 PPTP ALG (default: disabled)
  PLACEHOLDER
---

# thunder_cgnv6_nat64_alg_pptp (Resource)

`thunder_cgnv6_nat64_alg_pptp`: NAT64 PPTP ALG (default: disabled)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_alg_pptp" "thunder_cgnv6_nat64_alg_pptp" {
  pptp_enable = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `pptp_enable` (String) 'enable': Enable NAT64 PPTP ALG;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

