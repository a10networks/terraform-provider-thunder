---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_templateFtp"
sidebar_current: "docs-vthunder-resource-slb_templateFtp"
description: |-
    Provides details about vthunder SLB template ftp resource for A10
---

# vthunder\_slb\_template\_ftp

`vthunder_TemplateFtp` Provides details about vthunder SLB template ftp
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_ftp" "templateftp" {
	name = "testftp"
	user_tag = "test_tag"
	active_mode_port = 1
	active_mode_port_val = 1
}
```

## Argument Reference

* `name` - FTP Template Name
* `active_mode_port` - Non-Standard FTP Active mode port
* `active_mode_port_val` - Non-Standard FTP Active mode port
* `to` - End range of FTP Active mode port
* `user_tag` - Customized tag




