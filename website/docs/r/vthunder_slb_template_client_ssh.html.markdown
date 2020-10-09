---
layout: "thunder"
page_title: "thunder: thunder_slb_template_client_ssh"
sidebar_current: "docs-thunder-resource-slb-template-client-ssh"
description: |-
    Provides details about thunder slb template client-ssh resource for A10
---

# thunder\_slb\_template\_client\_ssh

`thunder_slb_template_client_ssh` provides details about slb template client_ssh
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_client_ssh" "client_ssh1" {
	name = "testssh"
	user_tag = "test_tag"
	forward_proxy_enable = 1
	forward_proxy_hostkey = "test"
	passphrase = "testing123"
}
```

## Argument Reference

* `name` - Client SSH Template Name
* `user_tag` - Customized tag
* `forward_proxy_enable` - Enable SSH forward proxy
* `forward_proxy_hostkey` - Specify private-key
* `passphrase` - Password Phrase

