---
layout: "thunder"
page_title: "thunder: thunder_slb_template_http_policy"
sidebar_current: "docs-thunder-resource-slb-template-http-policy"
description: |-
    Provides details about thunder slb template http-policy resource for A10
---

# thunder\_slb\_template\_http_policy

`thunder_slb_template_http_policy` provides details about slb template http-policy
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_http_policy" "http_policy"{
name= "httppol1"
cookie_name="cookie-two"
user_tag="u1"
}
```

## Argument Reference

* `name` - http-policy template name
* `user_tag` - Customized tag
* `cookie_name` - name of cookie to match (Cookie Name)


