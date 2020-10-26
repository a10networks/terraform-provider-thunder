---
layout: "thunder"
page_title: "thunder: thunder_slb_template_server_ssh"
sidebar_current: "docs-thunder-resource-slb-template-server-ssh"
description: |-
    Provides details about thunder slb template server-ssh resource for A10
---

# thunder\_slb\_template\_server\_ssh

`thunder_slb_template_server_ssh` provides details about slb template server_ssh
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_server_ssh" "server_ssh" {
	name = "testserverssh"
	user_tag = "test_tag"
	forward_proxy_enable = 1
}
```

## Argument Reference

* `name` - Server SSH Template Name
* `user_tag` - Customized tag
* `forward_proxy_enable` - Enable SSH forward proxy

