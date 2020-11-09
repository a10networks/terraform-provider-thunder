---
layout: "thunder"
page_title: "thunder: thunder_dns_primary"
sidebar_current: "docs-thunder-resource-dns-primary"
description: |-
    Provides details about thunder dns primary resource for A10
---

# thunder\_dns\_primary

`thunder_dns_primary provides` details about Dns Primary
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_dns_primary" "Ip_Dns_Primary_Test" {

ip_v4_addr = "string"
ip_v6_addr = "string"
uuid = "string"
 
}
```

## Argument Reference

* `ip_v4_addr` - DNS server address
* `ip_v6_addr` - DNS server address
* `uuid` - uuid of the object
