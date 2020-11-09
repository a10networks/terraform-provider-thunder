---
layout: "thunder"
page_title: "thunder: thunder_write_memory"
sidebar_current: "docs-thunder-resource-write-memory"
description: |-
	Provides details about thunder write memory resource for A10
---

# thunder\_write\_memory

`thunder_write_memory` Provides details about thunder write memory
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_write_memory" "test_write_memory"{
    
    profile = "my_profile_name"
    destination = "local"
    partition =  "all"
    specified_partition = "string"
}
```

## Argument Reference

* `destination` - ‘primary’: Write to default Primary Configuration; ‘secondary’: Write to default Secondary Configuration; ‘local’: Local Configuration Profile Name;
* `partition` - ‘all’: All partition configurations; ‘shared’: Shared partition; ‘specified’: Specified partition;
* `profile` - Local Configuration Profile Name
* `specified_partition` - Specified partition

