---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_client_ssh"
sidebar_current: "docs-vthunder-resource-slb-template-client-ssh"
description: |-
    Provides details about vthunder slb template client-ssh resource for A10
---

# vthunder\_slb\_template\_client\_ssh

`vthunder_slb_template_client_ssh` provides details about slb template client_ssh
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_client_ssh" "client_ssh1" {
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

