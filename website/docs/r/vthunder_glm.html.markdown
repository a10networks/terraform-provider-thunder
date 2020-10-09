---
layout: "thunder"
page_title: "thunder: thunder_glm"
sidebar_current: "docs-thunder-resource-glm"
description: |-
    Provides details about thunder GLM resource for A10
---

# thunder\_glm

`thunder_glm provides` details about GLM
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_glm" "glm1" {
		token = "vTh7aa90d73d"
		use_mgmt_port = 1
	    enable_requests = 1
	    allocate_bandwidth = 1000
}
```

## Argument Reference

* `token` - License entitlement token
* `use_mgmt_port` - Use management port to connect to GLM
* `enable_requests` - Turn on periodic GLM license requests (default license retrieval interval is every 24 hours)
* `allocate_bandwidth` - Enter the requested bandwidth in Mbps for Capacity Pool
