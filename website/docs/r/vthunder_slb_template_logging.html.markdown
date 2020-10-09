---
layout: "thunder"
page_title: "thunder: thunder_slb_template_logging"
sidebar_current: "docs-thunder-resource-slb-template-logging"
description: |-
    Provides details about thunder slb template logging resource for A10
---

# thunder\_slb\_template\_logging

`thunder_slb_template_logging` provides details about slb template logging
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_logging" "logging"{
name = "log1"
format= "abc"
local_logging= 1
}
```

## Argument Reference

* `name` - Logging Template Name
* `format` - Specify a format string for web logging (format string(less than 250 characters) for web logging)
* `local_logging` - 1 to enable local logging (1 to enable local logging, default 0)


