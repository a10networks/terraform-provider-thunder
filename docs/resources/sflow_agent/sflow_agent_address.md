---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_sflow_agent_address Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_sflow_agent_address: Configure agent address used in sFlow datagram, default use management IP address
  PLACEHOLDER
---

# thunder_sflow_agent_address (Resource)

`thunder_sflow_agent_address`: Configure agent address used in sFlow datagram, default use management IP address

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_agent_address" "thunder_sflow_agent_address" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ip` (String) Configure sFlow agent IP address
- `ipv6` (String) Configure sFlow agent IPv6 address
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

