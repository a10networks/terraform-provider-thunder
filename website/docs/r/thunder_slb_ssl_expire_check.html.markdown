---
layout: "thunder"
page_title: "thunder: thunder_slb_ssl_expire_check"
sidebar_current: "docs-thunder-resource-slb-ssl-expire-check"
description: |-
    Provides details about thunder SLB ssl expire check resource for A10
---

# thunder\_slb\_ssl\_expire\_check

`thunder_slb_ssl_expire_check` Provides details about thunder SLB ssl expire check
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_ssl_expire_check" "SSLExpireCheck" {
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
