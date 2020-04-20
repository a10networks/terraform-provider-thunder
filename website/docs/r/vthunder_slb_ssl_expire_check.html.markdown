---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_ssl_expire_check"
sidebar_current: "docs-vthunder-resource-slb-ssl-expire-check"
description: |-
    Provides details about vthunder SLB ssl expire check resource for A10
---

# vthunder\_slb\_ssl\_expire\_check

`vthunder_slb_ssl_expire_check` Provides details about vthunder SLB ssl expire check
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_ssl_expire_check" "SSLExpireCheck" {
	before = 5
	expire_address1 = "abc@gmail.com"
	interval_days = 2
	ssl_expire_email_address = "abc@gmail.com"
}
```

## Argument Reference

* `ssl_expire_email_address` - Config Email address for certificate expiration check
* `before` - The number of days in advance notice before expiration, default is 5
* `interval_days` - The interval of days notice after expiration, default is 2
* `expire_address1` - Email address for certificate expiration check
