---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_transparent_tcp_template"
sidebar_current: "docs-vthunder-slb-transparent-tcp-template"
description: |-
    Provides details about vthunder SLB transparent tcp template resource for A10
---

# vthunder\_slb\_transparent\_tcp\_template

`vthunder_slb_transparent_tcp_template` Provides details about vthunder SLB transparent tcp template
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_transparent_tcp_template" "tcp_template" {
	name = "testtransperenttcptemplate"
}
```

## Argument Reference

* `name` - Specify template name





