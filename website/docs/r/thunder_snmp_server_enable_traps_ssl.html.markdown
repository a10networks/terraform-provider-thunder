---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_ssl"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-ssl"
description: |-
	Provides details about thunder snmp server enable traps ssl resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_ssl

`thunder_snmp_server_enable_traps_ssl` Provides details about thunder snmp server enable traps ssl
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_enable_traps_ssl" "resourceSnmpServerEnableTrapsSslTest" {
	server_certificate_error = 0
uuid = "string"
 
}

```

## Argument Reference

* `server_certificate_error` - Enable SSL server certificate error trap
* `uuid` - uuid of the object

