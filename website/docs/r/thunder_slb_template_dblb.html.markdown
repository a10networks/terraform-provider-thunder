---
layout: "thunder"
page_title: "thunder: thunder_slb_template_dblb"
sidebar_current: "docs-thunder-resource-slb-template-dblb"
description: |-
	Provides details about thunder slb template dblb resource for A10
---

# thunder\_slb\_template\_dblb

`thunder_slb_template_dblb` Provides details about thunder slb template dblb
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_dblb" "Slb_Template_Dblb_Test" {

server_version = "string"
name = "string"
class_list = "string"
user_tag = "string"
calc_sha1 {  
        sha1_value =  "string" 
        }
uuid = "string"
 
}

```

## Argument Reference

* `class_list` - Specify user/password string class list (Class list name)
* `name` - DBLB template name
* `server_version` - ‘MSSQL2008’: MSSQL server 2008 or 2008 R2; ‘MSSQL2012’: MSSQL server 2012; ‘MySQL’: MySQL server (any version);
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `sha1_value` - Cleartext password
