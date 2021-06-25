---
layout: "thunder"
page_title: "thunder: thunder_interface_ve_bfd"
sidebar_current: "docs-thunder-resource-interface-ve-bfd"
description: |-
	Provides details about thunder interface ve bfd resource for A10
---

# thunder\_interface\_ve\_bfd

`thunder_interface_ve_bfd` Provides details about thunder interface ve bfd
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_interface_ve_bfd" "Interface_Ve_Bfd_Test" {

interval_cfg {  
        interval =  0 
        min_rx =  0 
        multiplier =  0 
        }
authentication {  
        encrypted =  "string" 
        password =  "string" 
        method =  "string" 
        key_id =  0 
        }
echo = 0
uuid = "string"
demand = 0
 
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
