---
layout: "thunder"
page_title: "thunder: thunder_configure_sync"
sidebar_current: "docs-thunder-resource-configure-sync"
description: |-
    Provides details about thunder configure sync resource for A10
---

# thunder\_configure\_sync

`thunder_configure_sync` provides details about Configure Sync
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_configure_sync" "configure_sync" {
		all_partitions = 1
	    private_key = "smita_key.pem"
}
```

## Argument Reference

* `all_partitions` - All partition configurations
* `private_key` - Use private key for authentication
