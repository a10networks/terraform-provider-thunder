---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_common_conn_rate_limit_src_ip"
sidebar_current: "docs-vthunder-resource-slb-common-conn-rate-limit-src-ip"
description: |-
    Provides details about vthunder SLB common conn-rate-limit src-ip resource for A10
---

# vthunder\_slb\_common\_conn_rate_limit\_src_ip

`vthunder_slb_conn_rate_limit_src_ip` Provides details about vthunder SLB common conn-rate-limit src-ip
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_common_conn_rate_limit_src_ip" "src_ip" {
	protocol = "tcp"
	limit_period = 1000
	limit = 50
	exceed_action = 0
	shared = 0
}
```

## Argument Reference

* `protocol` - 'tcp’: Set TCP connection rate limit; 'udp’: Set UDP packet rate limit;
* `limit_period` - '100’: 100 ms; '1000’: 1000 ms;
* `limit` - Set max connections per period
* `exceed_action` - Set action if threshold exceeded
* `shared` -  Set threshold shared amongst all virtual ports
