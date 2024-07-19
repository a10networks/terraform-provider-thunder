---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_preferred_session_sync_port_ethernet Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_preferred_session_sync_port_ethernet: preferred-session-sync-port ethernet
  PLACEHOLDER
---

# thunder_vrrp_a_preferred_session_sync_port_ethernet (Resource)

`thunder_vrrp_a_preferred_session_sync_port_ethernet`: preferred-session-sync-port ethernet

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_preferred_session_sync_port_ethernet" "thunder_vrrp_a_preferred_session_sync_port_ethernet" {
  pre_eth  = 1
  pre_vlan = 514
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `pre_eth` (Number) Ethernet interface number

### Optional

- `pre_vlan` (Number) Interface VLAN (VLAN ID)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

