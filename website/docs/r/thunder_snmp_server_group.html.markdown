---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_group"
sidebar_current: "docs-thunder-resource-snmp-server-group"
description: |-
	Provides details about thunder snmp server group resource for A10
---

# thunder\_snmp\_server\_group

`thunder_snmp_server_group` Provides details about thunder snmp server group
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_snmp_server_group" "resourceSnmpServerGroupTest" {
	read = "string"
groupname = "string"
v3 = "string"
uuid = "string"
 
}

```

## Argument Reference

* `groupname` - Name of the group
* `read` - specify a read view for the group (read view name)
* `uuid` - uuid of the object
* `v3` - ‘auth’: group using the authNoPriv Security Level; ‘noauth’: group using the noAuthNoPriv Security Level; ‘priv’: group using SNMPv3 authPriv security level;

