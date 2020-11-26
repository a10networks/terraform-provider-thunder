---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_slb_data_cache_timeout"
sidebar_current: "docs-thunder-resource-snmp-server-slb-data-cache-timeout"
description: |-
	Provides details about thunder snmp server slb data cache timeout resource for A10
---

# thunder\_snmp\_server\_slb\_data\_cache\_timeout

`thunder_snmp_server_slb_data_cache_timeout` Provides details about thunder snmp server slb data cache timeout
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_slb_data_cache_timeout" "resourceSnmpServerSlbDataCacheTimeoutTest" {
	uuid = "string"
slblimit = 0
 
}

```

## Argument Reference

* `slblimit` - Cache time-out in seconds, default as 60 seconds
* `uuid` - uuid of the object

