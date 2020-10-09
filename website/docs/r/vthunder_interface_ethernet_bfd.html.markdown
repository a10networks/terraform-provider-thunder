---
layout: "thunder"
page_title: "thunder: thunder_interface_ethernet_bfd"
sidebar_current: "docs-thunder-resource-interface-ethernet-bfd"
description: |-
  Provides details about thunder interface ethernet bfd resource for A10
---

# thunder\_interface\_ethernet\_bfd

`thunder_interface_ethernet_bfd` Provides details about thunder interface ethernet bfd
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_interface_ethernet_bfd" "ethernetbfd" {
  ifnum=1
  authentication {
    method = "md5"
    key_id = 0
    password = "joker"
  }
}
```

## Argument Reference

* `ifnum` - Interface no.
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

