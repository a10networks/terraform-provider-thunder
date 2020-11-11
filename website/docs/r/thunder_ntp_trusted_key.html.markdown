---
layout: "thunder"
page_title: "thunder: thunder_ntp_trusted_key"
sidebar_current: "docs-thunder-resource-ntp-trusted-key"
description: |-
	Provides details about thunder ntp trusted key resource for A10
---

# thunder\_ntp\_trusted\_key

`thunder_ntp_trusted_key` Provides details about thunder ntp trusted key
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_ntp_trusted_key" "resourceNtpTrustedKeyTest" {
	key = 0
uuid = "string"
 
}

```

## Argument Reference

* `key` - trusted key
* `uuid` - uuid of the object

