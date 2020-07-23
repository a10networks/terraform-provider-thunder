---
layout: "vthunder"
page_title: "vthunder: vthunder_interface_ve_bfd"
sidebar_current: "docs-vthunder-resource-interface-ve-bfd"
description: |-
	Provides details about vthunder interface ve bfd resource for A10
---

# vthunder\_interface\_ve\_bfd

`vthunder_interface_ve_bfd` Provides details about vthunder interface ve bfd
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_interface_ve_bfd" "vebfd" {
    ifnum=11
    authentication {
      method = "md5"
      key_id = 0
      password = "joker"
    }
}
```

## Argument Reference

* `demand` - Demand mode
* `echo` - Enable BFD Echo
* `uuid` - uuid of the object
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `key_id` - Key ID
* `method` - ‘md5’: Keyed MD5; ‘meticulous-md5’: Meticulous Keyed MD5; ‘meticulous-sha1’: Meticulous Keyed SHA1; ‘sha1’: Keyed SHA1; ‘simple’: Simple Password;
* `password` - Key String
* `interval` - Transmit interval between BFD packets (Milliseconds)
* `min_rx` - Minimum receive interval capability (Milliseconds)
* `multiplier` - Multiplier value used to compute holddown (value used to multiply the interval)

