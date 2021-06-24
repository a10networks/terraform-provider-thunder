---
layout: "thunder"
page_title: "thunder: thunder_hostname"
sidebar_current: "docs-thunder-resource-hostname"
description: |-
	Provides details about thunder hostname resource for A10
---

# thunder\_hostname

`thunder_hostname` Provides details about thunder hostname
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_hostname" "test_hostname"{

    value = "my_hostname"
}
```

## Argument Reference

* `uuid` - uuid of the object
* `value` - This System’s Network Name

