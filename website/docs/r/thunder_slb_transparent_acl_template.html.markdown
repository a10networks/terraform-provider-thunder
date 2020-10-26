---
layout: "thunder"
page_title: "thunder: thunder_slb_transparent_acl_template"
sidebar_current: "docs-thunder-slb-transparent-acl-template"
description: |-
    Provides details about thunder SLB transparent acl template resource for A10
---

# thunder\_slb\_transparent\_acl\_template

`thunder_slb_transparent_acl_template` Provides details about thunder SLB transparent acl template
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_transparent_acl_template" "acl_template" {
	name = "testtransparentacltemplate"
}
```

## Argument Reference

* `name` - Specify template name





