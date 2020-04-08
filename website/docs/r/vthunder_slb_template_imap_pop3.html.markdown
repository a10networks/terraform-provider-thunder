---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_imap_pop3"
sidebar_current: "docs-vthunder-resource-slb-template-imap-pop3"
description: |-
    Provides details about vthunder slb template imap-pop3 resource for A10
---

# vthunder\_slb\_template\_imap\_pop3

`vthunder_slb_template_imap_pop3` provides details about slb template imap-pop3
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_imap_pop3" "testname" {
	name = "Testimap"
	logindisabled = 0
	starttls = "disabled"
	user_tag = "test_user"
}
```

## Argument Reference

* `name` - IMAP-POP3 Template Name
* `logindisabled` - Disable Login before STARTTLS.Works only for imap
* `starttls` - 'disabled’: Disable STARTTLS; 'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before imap transaction;
* `user_tag` - Customized tag


