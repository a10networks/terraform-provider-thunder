---
layout: "thunder"
page_title: "thunder: thunder_vrrp_peer_group"
sidebar_current: "docs-thunder-resource-vrrp-peer-group"
description: |-
    Provides details about thunder vrrp peer group resource for A10
---

# thunder\_vrrp\_peer\_group

`thunder_vrrp_peer_group` provides details about vrrp peer group
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_vrrp_peer_group" "peer_group" {
		peer{
			ip_peer_address_cfg{
			ip_peer_address = "10.0.2.2"
		}
			ip_peer_address_cfg{
			ip_peer_address = "10.0.2.3"
		}
	}
}

```

## Argument Reference

* `ip_peer_address` - IP address of peer

