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

resource "thunder_configure_sync" "Configure_Sync_Test" {

all_partitions = 0
private_key = "string"
partition_name = "string"
pwd = "string"
auto_authentication = 0
address = "string"
shared = 0
type = "string"
pwd_enc = "string"
usr = "string" 

}

```

## Argument Reference

* `address` - Specify the destination ip address to sync
* `all_partitions` - All partition configurations
* `auto_authentication` - Authenticate with local username and password
* `partition_name` - Partition name
* `private_key` - Use private key for authentication
* `shared` - Shared partition
* `type` - ‘running’: Sync local running to peer’s running configuration; ‘all’: Sync local running to peer’s running configuration, and local startup to peer’s startup configuration;
