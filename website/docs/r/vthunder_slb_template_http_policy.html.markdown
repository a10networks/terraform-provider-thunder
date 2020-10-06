---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_http_policy"
sidebar_current: "docs-vthunder-resource-slb-template-http-policy"
description: |-
    Provides details about vthunder slb template http-policy resource for A10
---

# vthunder\_slb\_template\_http_policy

`vthunder_slb_template_http_policy` provides details about slb template http-policy
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_http_policy" "http_policy"{
name= "httppol1"
cookie_name="cookie-two"
user_tag="u1"
}
```

## Argument Reference

* `name` - http-policy template name
* `user_tag` - Customized tag
* `cookie_name` - name of cookie to match (Cookie Name)


