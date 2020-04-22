---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_mssql"
sidebar_current: "docs-vthunder-resource-slb-mssql"
description: |-
    Provides details about vthunder SLB mssql resource for A10
---

# vthunder\_slb\_mssql

`vthunder_slb_mssql` Provides details about vthunder SLB mssql
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_mssql" "mssql" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'curr_be_enc’: Curr BE Encryption Conns; 'total_be_enc’: Total BE Encryption Conns; 'curr_fe_enc’: Curr FE Encryption Conns; 'total_fe_enc’: Total FE Encryption Conns; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'session_err’: Session err; 'queries’: DB Queries; 'commands’: DB commands reply; 'auth_success’: Authentication Success; 'auth_failure’: Authentication Failure;
