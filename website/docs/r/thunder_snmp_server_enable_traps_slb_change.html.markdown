---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_slb_change"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-slb-change"
description: |-
	Provides details about thunder snmp server enable traps slb change resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_slb\_change

`thunder_snmp_server_enable_traps_slb_change` Provides details about thunder snmp server enable traps slb change
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_snmp_server_enable_traps_slb_change" "SnmpServerEnableTrapsSlbChangeTest" {
	all = 0
resource_usage_warning = 0
uuid = "string"
ssl_cert_change = 0
ssl_cert_expire = 0
system_threshold = 0
server = 0
vip = 0
connection_resource_event = 0
server_port = 0
vip_port = 0
 
}

```

## Argument Reference

* `all` - Enable all system group traps
* `connection_resource_event` - Enable system connection resource event trap
* `resource_usage_warning` - Enable partition resource usage warning trap
* `server` - Enable slb server create/delete trap
* `server_port` - Enable slb server port create/delete trap
* `ssl_cert_change` - Enable SSL certificate change trap
* `ssl_cert_expire` - Enable SSL certificate expiring trap
* `system_threshold` - Enable slb system threshold trap
* `uuid` - uuid of the object
* `vip` - Enable slb vip create/delete trap
* `vip_port` - Enable slb vip-port create/delete trap

