---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_mysql"
sidebar_current: "docs-vthunder-resource-slb-mysql"
description: |-
    Provides details about vthunder SLB mysql resource for A10
---

# vthunder\_slb\_mysql

`vthunder_slb_mysql` Provides details about vthunder SLB mysql
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_mysql" "mysql" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'curr_be_enc’: Curr BE Encryption Conns; 'total_be_enc’: Total BE Encryption Conns; 'curr_fe_enc’: Curr FE Encryption Conns; 'total_fe_enc’: Total FE Encryption Conns; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'session_err’: Session err; 'queries’: DB Queries; 'commands’: DB commands reply;