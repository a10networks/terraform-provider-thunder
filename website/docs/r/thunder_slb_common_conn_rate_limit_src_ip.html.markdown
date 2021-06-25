---
layout: "thunder"
page_title: "thunder: thunder_slb_common_conn_rate_limit_src_ip"
sidebar_current: "docs-thunder-resource-slb-common-conn-rate-limit-src-ip"
description: |-
	Provides details about thunder slb common conn rate limit src ip resource for A10
---

# thunder\_slb\_common\_conn\_rate\_limit\_src\_ip

`thunder_slb_common_conn_rate_limit_src_ip` Provides details about thunder slb common conn rate limit src ip
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_common_conn_rate_limit_src_ip" "Slb_Common_Conn_Rate_Limit_Src_Ip_Test" {

protocol = "string"
log = 0
lock_out = 0
limit_period = "string"
limit = 0
exceed_action = 0
shared = 0
uuid = "string"
 
}

```

## Argument Reference

* `exceed_action` - Set action if threshold exceeded
* `limit` - Set max connections per period
* `limit_period` - ‘100’: 100 ms; ‘1000’: 1000 ms;
* `lock_out` - Set lockout period in seconds if threshold exceeded
* `log` - Send log if threshold exceeded
* `protocol` - ‘tcp’: Set TCP connection rate limit; ‘udp’: Set UDP packet rate limit;
* `shared` - Set threshold shared amongst all virtual ports
* `uuid` - uuid of the object
