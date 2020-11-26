---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_user"
sidebar_current: "docs-thunder-resource-snmp-server-user"
description: |-
	Provides details about thunder snmp server user resource for A10
---

# thunder\_snmp\_server\_user

`thunder_snmp_server_user` Provides details about thunder snmp server user
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_user" "resourceSnmpServerUserTest" {
	username = "string"
auth_val = "string"
group = "string"
uuid = "string"
encpasswd = "string"
passwd = "string"
priv_pw_encrypted = "string"
v3 = "string"
pw_encrypted = "string"
priv = "string"
 
}

```

## Argument Reference

* `auth_val` - ‘md5’: Use HMAC MD5 algorithm for authentication; ‘sha’: Use HMAC SHA algorithm for authentication;
* `encpasswd` - Passphrase for encryption
* `group` - Group to which the user belongs
* `passwd` - Password of this user
* `priv` - ‘des’: DES encryption alogrithm; ‘aes’: AES encryption alogrithm;  (Encryption type)
* `priv_pw_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED passphrase string)
* `pw_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED passphrase string)
* `username` - Name of the user
* `uuid` - uuid of the object
* `v3` - ‘auth’: Using the authNoPriv Security Level; ‘noauth’: Using the noAuthNoPriv Security Level;

