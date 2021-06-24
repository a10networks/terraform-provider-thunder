---
layout: "thunder"
page_title: "thunder: thunder_active_partition"
sidebar_current: "docs-thunder-resource-active-partition"
description: |-
    Provides details about thunder active partition resource for A10
---

# thunder\_active\_partition

`thunder_active_partition` Provides details about thunder active partition
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_active_partition" "resourceActivePartitionTest" {
	shared = 0
curr_part_name = "string"
 
}

```

## Argument Reference

* `active-partition` - Change active partition
* `shared` - Shared partition
* `curr_part_name` - current active patition

