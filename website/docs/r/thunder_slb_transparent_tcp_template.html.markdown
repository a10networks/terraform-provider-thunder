---
layout: "thunder"
page_title: "thunder: thunder_slb_transparent_tcp_template"
sidebar_current: "docs-thunder-slb-transparent-tcp-template"
description: |-
    Provides details about thunder SLB transparent tcp template resource for A10
---

# thunder\_slb\_transparent\_tcp\_template

`thunder_slb_transparent_tcp_template` Provides details about thunder SLB transparent tcp template
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_transparent_tcp_template" "tcp_template" {
	name = "testtransperenttcptemplate"
}
```

## Argument Reference

* `name` - Specify template name





