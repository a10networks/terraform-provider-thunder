---
layout: "thunder"
page_title: "thunder: thunder_slb_template_client_ssh"
sidebar_current: "docs-thunder-resource-slb-template-client-ssh"
description: |-
	Provides details about thunder slb template client ssh resource for A10
---

# thunder\_slb\_template\_client\_ssh

`thunder_slb_template_client_ssh` Provides details about thunder slb template client ssh
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_client_ssh" "Slb_Template_Client_Ssh_Test" {

uuid = "string"
encrypted = "string"
user_tag = "string"
forward_proxy_enable = 0
passphrase = "string"
forward_proxy_hostkey = "string"
name = "string"
 
}


```

## Argument Reference

* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `forward_proxy_enable` - Enable SSH forward proxy
* `forward_proxy_hostkey` - Specify private-key (Key Name)
* `name` - Client SSH Template Name
* `passphrase` - Password Phrase
* `user_tag` - Customized tag
* `uuid` - uuid of the object
