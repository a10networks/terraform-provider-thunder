---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_system"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-system"
description: |-
	Provides details about thunder snmp server enable traps system resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_system

`thunder_snmp_server_enable_traps_system` Provides details about thunder snmp server enable traps system
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_enable_traps_system" "resourceSnmpServerEnableTrapsSystemTest" {
	all = 0
data_cpu_high = 0
uuid = "string"
power = 0
high_disk_use = 0
high_memory_use = 0
control_cpu_high = 0
file_sys_read_only = 0
low_temp = 0
high_temp = 0
sec_disk = 0
license_management = 0
start = 0
fan = 0
shutdown = 0
pri_disk = 0
syslog_severity_one = 0
tacacs_server_up_down = 0
smp_resource_event = 0
restart = 0
packet_drop = 0
 
}

```

## Argument Reference

* `all` - Enable all system group traps
* `control_cpu_high` - Enable control CPU usage high trap
* `data_cpu_high` - Enable data CPU usage high trap
* `fan` - Enable system fan trap
* `file_sys_read_only` - Enable file system read-only trap
* `high_disk_use` - Enable system high disk usage trap
* `high_memory_use` - Enable system high memory usage trap
* `high_temp` - Enable system high temperature trap
* `license_management` - Enable system license management traps
* `low_temp` - Enable system low temperature trap
* `packet_drop` - Enable system packet dropped trap
* `power` - Enable system power supply trap
* `pri_disk` - Enable system primary hard disk trap
* `restart` - Enable system restart trap
* `sec_disk` - Enable system secondary hard disk trap
* `shutdown` - Enable system shutdown trap
* `smp_resource_event` - Enable system smp resource event trap
* `start` - Enable system start trap
* `tacacs_server_up_down` - Enable system TACACS monitor server up/down trap
* `uuid` - uuid of the object

