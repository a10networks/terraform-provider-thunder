---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_smtp"
sidebar_current: "docs-vthunder-resource-slb_template_smtp"
description: |-
    Provides details about vthunder SLB template smtp resource for A10
---

# vthunder\_slb\_template\_smtp

`vthunder_slb_template_smtp` Provides details about vthunder SLB template smtp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_smtp" "testname" {
	name = "testsmtp"
	user_tag = "test_tag"
	server_domain= "gmail.com"
	client_starttls_type = "optional"
	command_disable {
		disable_type = "expn"
	}
	server_starttls_type = "optional"
	service_ready_msg = "ttt test"
}
```

## Argument Reference

* `name` - SMTP Template Name
* `user_tag` - Customized tag
* `server_domain` - Config the domain of the email servers (Server’s domain, default is “mail-server-domain”)
* `client_starttls_type` - 'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before mail transaction
* `disable_type` - 'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before mail transaction
* `server_starttls_type` - 'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before mail transaction;
* `service_ready_msg` - Set SMTP service ready message (SMTP service ready message, default is “ESMTP mail service ready”)




