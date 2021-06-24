---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_SNMPv3_user"
sidebar_current: "docs-thunder-resource-snmp-server-SNMPv3-user"
description: |-
	Provides details about thunder snmp server SNMPv3 user resource for A10
---

# thunder\_snmp\_server\_SNMPv3\_user

`thunder_snmp_server_SNMPv3_user` Provides details about thunder snmp server SNMPv3 user
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_snmp_server_snmpv3_user" "Snmp_Server_SNMPv3_User_Test" {

username = "string"
group = "string"
v3 = "string"
auth_val = "string"
passwd = "string"
pw_encrypted = "string"
priv = "string"
encpasswd = "string"
priv_pw_encrypted = "string"
uuid = "string"
 
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

