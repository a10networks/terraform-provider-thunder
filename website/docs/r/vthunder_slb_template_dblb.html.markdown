---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_dblb"
sidebar_current: "docs-vthunder-resource-slb_template_dblb"
description: |-
    Provides details about vthunder SLB template dblb resource for A10
---

# vthunder\_slb\_template\_dblb

`vthunder_slb_template_dblb` Provides details about vthunder SLB template dblb
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_dblb" "dblb" {
	name = "testdblb"
	user_tag = "test_tag"
	server_version = "MSSQL2008"
}
```

## Argument Reference

* `name` - BDLB Template Name
* `user_tag` - Customized tag
* `server_version` - 'MSSQL2008’: MSSQL server 2008 or 2008 R2; 'MSSQL2012’: MSSQL server 2012; 'MySQL’: MySQL server (any version)