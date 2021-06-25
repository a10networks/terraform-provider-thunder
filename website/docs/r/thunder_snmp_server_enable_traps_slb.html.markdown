---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_slb"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-slb"
description: |-
	Provides details about thunder snmp server enable traps slb resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_slb

`thunder_snmp_server_enable_traps_slb` Provides details about thunder snmp server enable traps slb
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_snmp_server_enable_traps_slb" "Snmp_Server_Enable_Traps_Slb_Test" {

all = 0
server_down = 0
vip_port_connratelimit = 0
server_selection_failure = 0
service_group_down = 0
server_conn_limit = 0
service_group_member_up = 0
uuid = "string"
server_conn_resume = 0
service_up = 0
service_conn_limit = 0
gateway_up = 0
service_group_up = 0
application_buffer_limit = 0
vip_connratelimit = 0
vip_connlimit = 0
service_group_member_down = 0
service_down = 0
bw_rate_limit_exceed = 0
server_disabled = 0
server_up = 0
vip_port_connlimit = 0
vip_port_down = 0
bw_rate_limit_resume = 0
gateway_down = 0
vip_up = 0
vip_port_up = 0
vip_down = 0
service_conn_resume = 0
 
}
```

## Argument Reference

* `all` - Enable all SLB traps
* `application_buffer_limit` - Enable application buffer reach limit trap
* `bw_rate_limit_exceed` - Enable SLB server/port bandwidth rate limit exceed trap
* `bw_rate_limit_resume` - Enable SLB server/port bandwidth rate limit resume trap
* `gateway_down` - Enable SLB server gateway down trap
* `gateway_up` - Enable SLB server gateway up trap
* `server_conn_limit` - Enable SLB server connection limit trap
* `server_conn_resume` - Enable SLB server connection resume trap
* `server_disabled` - Enable SLB server-disabled trap
* `server_down` - Enable SLB server-down trap
* `server_selection_failure` - Enable SLB server selection failure trap
* `server_up` - Enable slb server up trap
* `service_conn_limit` - Enable SLB service connection limit trap
* `service_conn_resume` - Enable SLB service connection resume trap
* `service_down` - Enable SLB service-down trap
* `service_group_down` - Enable SLB service-group-down trap
* `service_group_member_down` - Enable SLB service-group-member-down trap
* `service_group_member_up` - Enable SLB service-group-member-up trap
* `service_group_up` - Enable SLB service-group-up trap
* `service_up` - Enable SLB service-up trap
* `uuid` - uuid of the object
* `vip_connlimit` - Enable the virtual server reach conn-limit trap
* `vip_connratelimit` - Enable the virtual server reach conn-rate-limit trap
* `vip_down` - Enable SLB virtual server down trap
* `vip_port_connlimit` - Enable the virtual port reach conn-limit trap
* `vip_port_connratelimit` - Enable the virtual port reach conn-rate-limit trap
* `vip_port_down` - Enable SLB virtual port down trap
* `vip_port_up` - Enable SLB virtual port up trap
* `vip_up` - Enable SLB virtual server up trap

