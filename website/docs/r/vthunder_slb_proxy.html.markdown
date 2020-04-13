---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_proxy"
sidebar_current: "docs-vthunder-resource-slb-proxy"
description: |-
    Provides details about vthunder SLB proxy resource for A10
---

# vthunder\_slb\_proxy

`vthunder_slb_proxy` Provides details about vthunder SLB proxy
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_proxy" "testname" {

	sampling_enable  {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'num’: Num; 'tcp_event’: TCP stack event; 'est_event’: Connection established; 'data_event’: Data received; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'wbuf_event’: Ready to send data; 'err_event’: Error occured; 'no_mem’: No memory; 'client_rst’: Client RST; 'server_rst’: Server RST; 'queue_depth_over_limit’: Queue depth over limit; 'event_failed’: Event failed; 'conn_not_exist’: Conn not exist; 'service_alloc_cb’: Service alloc callback; 'service_alloc_cb_failed’: Service alloc callback failed; 'service_free_cb’: Service free callback; 'service_free_cb_failed’: Service free callback failed; 'est_cb_failed’: App EST callback failed; 'data_cb_failed’: App DATA callback failed; 'wbuf_cb_failed’: App WBUF callback failed; 'err_cb_failed’: App ERR callback failed; 'start_server_conn’: Start server conn; 'start_server_conn_succ’: Success; 'start_server_conn_no_route’: No route to server; 'start_server_conn_fail_mem’: No memory; 'start_server_conn_fail_snat’: Failed Source NAT; 'start_server_conn_fail_persist’: Fail Persistence; 'start_server_conn_fail_server’: Fail Server issue; 'start_server_conn_fail_tuple’: Fail Tuple Issue; 'line_too_long’: Line too long;