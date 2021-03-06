---
layout: "thunder"
page_title: "thunder: thunder_slb_template_persist_cookie"
sidebar_current: "docs-thunder-resource-slb-template-persist-cookie"
description: |-
	Provides details about thunder slb template persist cookie resource for A10
---

# thunder\_slb\_template\_persist\_cookie

`thunder_slb_template_persist_cookie` Provides details about thunder slb template persist cookie
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_persist_cookie" "Slb_Template_Persist_Cookie_Test" {

pass_phrase = "string"
domain = "string"
cookie_name = "string"
secure = 0
encrypted = "string"
dont_honor_conn_rules = 0
encrypt_level = 0
uuid = "string"
user_tag = "string"
server = 0
server_service_group = 0
service_group = 0
expire = 0
httponly = 0
path = "string"
pass_thru = 0
scan_all_members = 0
insert_always = 0
match_type = 0
name = "string"
 
}
```

## Argument Reference

* `cookie_name` - Set cookie name (Cookie name, default “sto-id”)
* `domain` - Set cookie domain
* `dont_honor_conn_rules` - Do not observe connection rate rules
* `encrypt_level` - Encryption level for cookie name / value
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `expire` - Set cookie expiration time (Expiration in seconds)
* `httponly` - Enable HttpOnly attribute
* `insert_always` - Insert persist cookie to every reponse
* `match_type` - Persist for server, default is port
* `name` - Cookie persistence (Cookie persistence template name)
* `pass_phrase` - Set passphrase for encryption
* `pass_thru` - Pass thru mode - Server sends the persist cookie
* `path` - Set cookie path (Cookie path, default is “/”)
* `scan_all_members` - Persist within the same server SCAN
* `secure` - Enable secure attribute
* `server` - Persist to the same server, default is port
* `server_service_group` - Persist to the same server and within the same service group
* `service_group` - Persist within the same service group
* `user_tag` - Customized tag
* `uuid` - uuid of the object

