---
layout: "thunder"
page_title: "thunder: thunder_slb_mysql"
sidebar_current: "docs-thunder-resource-slb-mysql"
description: |-
    Provides details about thunder SLB mysql resource for A10
---

# thunder\_slb\_mysql

`thunder_slb_mysql` Provides details about thunder SLB mysql
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_mysql" "mysql" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'curr_be_enc’: Curr BE Encryption Conns; 'total_be_enc’: Total BE Encryption Conns; 'curr_fe_enc’: Curr FE Encryption Conns; 'total_fe_enc’: Total FE Encryption Conns; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'session_err’: Session err; 'queries’: DB Queries; 'commands’: DB commands reply;