---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_vcs"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-vcs"
description: |-
	Provides details about thunder snmp server enable traps vcs resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_vcs

`thunder_snmp_server_enable_traps_vcs` Provides details about thunder snmp server enable traps vcs
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_enable_traps_vcs" "resourceSnmpServerEnableTrapsVcsTest" {
	state_change = 0
uuid = "string"
 
}

```

## Argument Reference

* `state_change` - Enable VCS state change trap
* `uuid` - uuid of the object

