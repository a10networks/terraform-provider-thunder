---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_prefix_list"
sidebar_current: "docs-vthunder-resource-ip-prefix-list"
description: |-
  Provides details about vthunder ip prefix list resource for A10
---

# vthunder\_ip\_prefix\_list

`vthunder_ip_prefix_list` Provides details about vthunder ip prefix list
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_prefix_list" "IpPrefixList" {
  name = "testprefixlist"
  rules {
        any= 1
        action = "deny"
        seq = 1
        le = 32
        ipaddr = "23.0.0.0/8"
        description = "string"
        ge = 24
  }
  uuid = "string"
}
```

## Argument Reference

* `name` - Name of a prefix list
* `uuid` - uuid of the object
* `action` - ‘deny’: Specify packets to reject; ‘permit’: Specify packets to forward;
* `any` - Any prefix match. Same as “0.0.0.0/0 le 32”
* `description` - Prefix-list specific description (Up to 80 characters describing this prefix-list)
* `ge` - Minimum prefix length to be matched
* `ipaddr` - IP prefix, e.g., 35.0.0.0/8
* `le` - Maximum prefix length to be matched
* `seq` - Sequence number of an entry

