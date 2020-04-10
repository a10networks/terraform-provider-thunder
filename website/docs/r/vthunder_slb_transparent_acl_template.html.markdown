---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_transparent_acl_template"
sidebar_current: "docs-vthunder-slb-transparent-acl-template"
description: |-
    Provides details about vthunder SLB transparent acl template resource for A10
---

# vthunder\_slb\_transparent\_acl\_template

`vthunder_slb_transparent_acl_template` Provides details about vthunder SLB transparent acl template
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_transparent_acl_template" "testname" {
	name = "testtransparentacltemplate"
}
```

## Argument Reference

* `name` - Specify template name





