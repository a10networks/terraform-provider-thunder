---
layout: "thunder"
page_title: "thunder: thunder_web_category_web_reputation"
sidebar_current: "docs-thunder-resource-web-category-web-reputation"
description: |-
    Provides details about thunder web category web reputation resource for A10
---

# thunder\_web\_category\_web\_reputation

`thunder_web_category_web_reputation` Provides details about thunder web category web reputation
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_web_category_web_reputation" "resourceWebCategoryWebReputationTest" {
	uuid = "string"
intercepted_urls {  
 	uuid =  "string" 
	}
bypassed_urls {  
 	uuid =  "string" 
	}
url {  
 	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `web-reputation` - Used for Web-Reputation Show Commands
* `uuid` - uuid of the object

