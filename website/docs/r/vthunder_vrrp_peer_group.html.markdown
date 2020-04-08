---
layout: "vthunder"
page_title: "vthunder: vthunder_vrrp_peer_group"
sidebar_current: "docs-vthunder-resource-vrrp-peer-group"
description: |-
    Provides details about vthunder vrrp peer group resource for A10
---

# vthunder\_vrrp\_peer\_group

`vthunder_vrrp_peer_group` provides details about vrrp peer group
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_vrrp_peer_group" "peer_group" {
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

