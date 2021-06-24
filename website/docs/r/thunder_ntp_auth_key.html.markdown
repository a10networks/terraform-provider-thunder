---
layout: "thunder"
page_title: "thunder: thunder_ntp_auth_key"
sidebar_current: "docs-thunder-resource-ntp-auth-key"
description: |-
	Provides details about thunder ntp auth key resource for A10
---

# thunder\_ntp\_auth\_key

`thunder_ntp_auth_key` Provides details about thunder ntp auth key
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_ntp_auth_key" "resourceNtpAuthKeyTest" {
	key = 0
alg_type = "string"
key_type = "string"
asc_key = "string"
encrypted = "string"
hex_key = "string"
hex_encrypted = "string"
uuid = "string"
 
}

```

## Argument Reference

* `alg_type` - ‘M’: encryption using MD5; ‘SHA’: encryption using SHA; ‘SHA1’: encryption using SHA1;
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `hex_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `key` - authentication key
* `key_type` - ‘ascii’: key string in ASCII form; ‘hex’: key string in hex form;
* `uuid` - uuid of the object

