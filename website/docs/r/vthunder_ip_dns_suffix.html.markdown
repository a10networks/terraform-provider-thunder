---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_dns_suffix"
sidebar_current: "docs-vthunder-resource-ip-dns-suffix"
description: |-
	Provides details about vthunder ip dns suffix resource for A10
---

# vthunder\_ip\_dns\_suffix

`vthunder_ip_dns_suffix` Provides details about vthunder ip dns suffix
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_dns_suffix" "testname" {
	domain_name = "domainName"
}
```

## Argument Reference

* `domain_name` - DNS suffix
* `uuid` - uuid of the object

