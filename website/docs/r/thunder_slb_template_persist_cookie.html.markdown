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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_persist_cookie" "resourceSlbTemplatePersistCookieTest" {
	name = "string"
dont_honor_conn_rules = 0
expire = 0
insert_always = 0
encrypt_level = 0
pass_phrase = "string"
cookie_name = "string"
prefix = "string"
domain = "string"
samesite = "string"
path = "string"
pass_thru = 0
secure = 0
httponly = 0
match_type = 0
server = 0
server_service_group = 0
scan_all_members = 0
service_group = 0
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `cookie` - Cookie persistence
* `name` - Cookie persistence (Cookie persistence template name)
* `dont-honor-conn-rules` - Do not observe connection rate rules
* `expire` - Set cookie expiration time (Expiration in seconds)
* `insert-always` - Insert persist cookie to every reponse
* `encrypt-level` - Encryption level for cookie name / value
* `pass-phrase` - Set passphrase for encryption
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `cookie-name` - Set cookie name
* `prefix` - 'host': the cookie will have been set with a Secure attribute, a Path attribute with a value of /, and no Domain attribute; 'secure': the cookie will have been set with a Secure attribute;
* `domain` - Set cookie domain
* `samesite` - 'none': none; 'lax': lax; 'strict': strict;
* `path` - Set cookie path (Cookie path, default is "/")
* `pass-thru` - Pass thru mode - Server sends the persist cookie
* `secure` - Enable secure attribute
* `httponly` - Enable HttpOnly attribute
* `match-type` - Persist for server, default is port
* `server` - Persist to the same server, default is port
* `server-service-group` - Persist to the same server and within the same service group
* `scan-all-members` - Persist within the same server SCAN
* `service-group` - Persist within the same service group
* `uuid` - uuid of the object
* `user-tag` - Customized tag

