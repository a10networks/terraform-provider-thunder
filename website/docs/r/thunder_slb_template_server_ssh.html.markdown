---
layout: "thunder"
page_title: "thunder: thunder_slb_template_server_ssh"
sidebar_current: "docs-thunder-resource-slb-template-server-ssh"
description: |-
	Provides details about thunder slb template server ssh resource for A10
---

# thunder\_slb\_template\_server\_ssh

`thunder_slb_template_server_ssh` Provides details about thunder slb template server ssh
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_server_ssh" "Slb_Template_Server_Ssh_Test" {
sampling-enable {   
counters1 =  "string" 
        }
forward_proxy_enable = 0
name = "string"
user_tag = "string"
uuid = "string"
 
}
```

## Argument Reference

* `forward_proxy_enable` - Enable SSH forward proxy
* `name` - Server SSH Template Name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘successful_handshakes’: successful_handshakes; ‘failed_handshakes’: failed_handshakes; ‘forwarding_errors’: forwarding_errors;
