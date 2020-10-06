---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_mlb"
sidebar_current: "docs-vthunder-resource-slb-mlb"
description: |-
    Provides details about vthunder SLB mlb resource for A10
---

# vthunder\_slb\_mlb

`vthunder_slb_mlb` Provides details about vthunder SLB mlb
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_mlb" "mlb" {
	sampling_enable {
	    counters1 = "all"
	 }
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'client_msg_sent’: Client message sent; 'server_msg_received’: Server message received; 'server_conn_created’: Server connection created; 'server_conn_rst’: Server connection reset; 'server_conn_failed’: Server connection failed; 'server_conn_closed’: Server connection closed; 'client_conn_created’: Client connection created; 'client_conn_closed’: Client connection closed; 'client_conn_not_found’: Client connection not found;
