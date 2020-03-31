---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_logging"
sidebar_current: "docs-vthunder-resource-slb-template-logging"
description: |-
    Provides details about vthunder slb template logging resource for A10
---

# vthunder\_slb\_template\_logging

`vthunder_slb_template_logging provides details about slb template logging
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_logging" "l"{
name = "log1"
format= "abc"
local_logging= 1
}
```

## Argument Reference

* `name` - Logging Template Name
* `format` - Specify a format string for web logging (format string(less than 250 characters) for web logging)
* `local_logging` - 1 to enable local logging (1 to enable local logging, default 0)


