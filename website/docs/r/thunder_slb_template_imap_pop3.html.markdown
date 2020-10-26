---
layout: "thunder"
page_title: "thunder: thunder_slb_template_imap_pop3"
sidebar_current: "docs-thunder-resource-slb-template-imap-pop3"
description: |-
    Provides details about thunder slb template imap-pop3 resource for A10
---

# thunder\_slb\_template\_imap\_pop3

`thunder_slb_template_imap_pop3` provides details about slb template imap-pop3
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_imap_pop3" "imap_pop3" {
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


