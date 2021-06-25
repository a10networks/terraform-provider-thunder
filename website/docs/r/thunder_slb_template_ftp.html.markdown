---
layout: "thunder"
page_title: "thunder: thunder_slb_template_ftp"
sidebar_current: "docs-thunder-resource-slb-template-ftp"
description: |-
	Provides details about thunder slb template ftp resource for A10
---

# thunder\_slb\_template\_ftp

`thunder_slb_template_ftp` Provides details about thunder slb template ftp
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_ftp" "Slb_Template_Ftp_Test" {

uuid = "string"
user_tag = "string"
to = 0
active_mode_port = 0
active_mode_port_val = 0
any = 0
name = "string"
 
}
```

## Argument Reference

* `active_mode_port` - Non-Standard FTP Active mode port
* `active_mode_port_val` - Non-Standard FTP Active mode port
* `any` - Allow any FTP Active mode port
* `name` - FTP template name
* `to` - End range of FTP Active mode port
* `user_tag` - Customized tag
* `uuid` - uuid of the object
