---
layout: "thunder"
page_title: "thunder: thunder_ip_dns_suffix"
sidebar_current: "docs-thunder-resource-ip-dns-suffix"
description: |-
	Provides details about thunder ip dns suffix resource for A10
---

# thunder\_ip\_dns\_suffix

`thunder_ip_dns_suffix` Provides details about thunder ip dns suffix
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_dns_suffix" "testname" {
	domain_name = "domainName"
}
```

## Argument Reference

* `domain_name` - DNS suffix
* `uuid` - uuid of the object

