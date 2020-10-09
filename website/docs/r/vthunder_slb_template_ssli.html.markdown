---
layout: "thunder"
page_title: "thunder: thunder_slb_template_ssli"
sidebar_current: "docs-thunder-resource-slb-template-ssli"
description: |-
    Provides details about thunder slb template ssli resource for A10
---

# thunder\_slb\_template\_ssli

`thunder_slb_template_ssli` provides details about slb template ssli
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_ssli" "ssli" {
	name = "testssli"
	type = "init"
	user_tag = "test_user"
}
```

## Argument Reference

* `name` - SSLi Template Name
* `type` - 'http’: HTTP service; 'xmpp’: XMPP service; 'smtp’: SMTP service; 'pop’: POP service;
* `user_tag` - Customized tag


