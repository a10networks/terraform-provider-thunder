---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_health_gateway"
sidebar_current: "docs-vthunder-resource-slb-health-gateway"
description: |-
    Provides details about vthunder SLB health-gateway resource for A10
---

# vthunder\_slb\_health_gateway

`vthunder_slb_health_gateway` Provides details about vthunder SLB health-gateway
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_health_gateway" "health_gateway" {
	sampling_enable {
		counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable` - 
	* `counters1` - 'all’: all; 'total_sent’: Number of Total health-check sent; 'total_retry_sent’: Number of Total health-check retry sent; 'total_timeout’: Number of Total health-check timeout;


