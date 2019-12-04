---
layout: "vthunder"
page_title: "vthunder: vthunder_glm"
sidebar_current: "docs-vthunder-resource-glm"
description: |-
    Provides details about vthunder GLM resource for A10
---

# vthunder\_glm

`vthunder_glm provides details about GLM
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_glm" "glm1" {
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
