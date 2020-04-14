---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_server_ssh"
sidebar_current: "docs-vthunder-resource-slb-template-server-ssh"
description: |-
    Provides details about vthunder slb template server-ssh resource for A10
---

# vthunder\_slb\_template\_server\_ssh

`vthunder_slb_template_server_ssh` provides details about slb template server_ssh
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_server_ssh" "server_ssh" {
	name = "testserverssh"
	user_tag = "test_tag"
	forward_proxy_enable = 1
}
```

## Argument Reference

* `name` - Server SSH Template Name
* `user_tag` - Customized tag
* `forward_proxy_enable` - Enable SSH forward proxy

