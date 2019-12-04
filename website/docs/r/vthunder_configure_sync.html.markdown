---
layout: "vthunder"
page_title: "vthunder: vthunder_configure_sync"
sidebar_current: "docs-vthunder-resource-configure-sync"
description: |-
    Provides details about vthunder configure sync resource for A10
---

# vthunder\_configure\_sync

`vthunder_configure_sync` provides details about Configure Sync
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_configure_sync" "configure_sync" {
		all_partitions = 1
	    private_key = "smita_key.pem"
}
```

## Argument Reference

* `all_partitions` - All partition configurations
* `private_key` - Use private key for authentication
