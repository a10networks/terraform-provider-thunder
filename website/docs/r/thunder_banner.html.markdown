---
layout: "thunder"
page_title: "thunder: thunder_banner"
sidebar_current: "docs-thunder-resource-banner"
description: |-
    Provides details about thunder banner resource for A10
---

# thunder\_banner

`thunder_banner` Provides details about thunder banner
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_banner" "resourceBannerTest" {
	exec_banner_cfg {  
 	exec =  0 
	exec_banner =  "string" 
	}
login_banner_cfg {  
 	login =  0 
	login_banner =  "string" 
	}
uuid = "string"
 
}

```

## Argument Reference

* `banner` - Define a login banner
* `exec` - Set EXEC process creation banner
* `exec-banner` - Banner text, string \n is taken as line break of multi-line banner text, use \\n for \n, \077 for ? and \011 for tab
* `login` - Set login banner
* `login-banner` - Banner text, string \n is taken as line break of multi-line banner text, use \\n to indicate \n
* `uuid` - uuid of the object

