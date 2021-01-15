---
layout: "thunder"
page_title: "thunder: thunder_router_ospf_default_information"
sidebar_current: "docs-thunder-resource-router-ospf-default-information"
description: |-
    Provides details about thunder router ospf default information resource for A10
---

# thunder\_router\_ospf\_default\_information

`thunder_router_ospf_default_information` Provides details about thunder router ospf default information
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_ospf_default_information" "resourceRouterOspfDefaultInformationTest" {
	originate = 0
always = 0
metric = 0
metric_type = 0
route_map = "string"
uuid = "string"
 
}

```

## Argument Reference

* `default-information` - Control distribution of default information
* `originate` - Distribute a default route
* `always` - Always advertise default route
* `metric` - OSPF default metric (OSPF metric)
* `metric-type` - OSPF metric type for default routes
* `route-map` - Route map reference (Pointer to route-map entries)
* `uuid` - uuid of the object

