---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_ssli"
sidebar_current: "docs-vthunder-resource-slb-template-ssli"
description: |-
    Provides details about vthunder slb template ssli resource for A10
---

# vthunder\_slb\_template\_ssli

`vthunder_slb_template_ssli` provides details about slb template ssli
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_ssli" "ssli" {
	name = "testssli"
	type = "init"
	user_tag = "test_user"
}
```

## Argument Reference

* `name` - SSLi Template Name
* `type` - 'http’: HTTP service; 'xmpp’: XMPP service; 'smtp’: SMTP service; 'pop’: POP service;
* `user_tag` - Customized tag


