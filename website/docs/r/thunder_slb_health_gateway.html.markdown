---
layout: "thunder"
page_title: "thunder: thunder_slb_health_gateway"
sidebar_current: "docs-thunder-resource-slb-health-gateway"
description: |-
    Provides details about thunder SLB health-gateway resource for A10
---

# thunder\_slb\_health_gateway

`thunder_slb_health_gateway` Provides details about thunder SLB health-gateway
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_health_gateway" "health_gateway" {
	sampling_enable {
		counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable` - 
	* `counters1` - 'all’: all; 'total_sent’: Number of Total health-check sent; 'total_retry_sent’: Number of Total health-check retry sent; 'total_timeout’: Number of Total health-check timeout;


