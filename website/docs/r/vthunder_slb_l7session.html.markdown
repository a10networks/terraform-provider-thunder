---
layout: "thunder"
page_title: "thunder: thunder_slb_l7session"
sidebar_current: "docs-thunder-resource-slb-l7session"
description: |-
    Provides details about thunder SLB l7session resource for A10
---

# thunder\_slb\_l7session

`thunder_slb_l7session` Provides details about thunder SLB l7session
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_l7session" "l7session" {
	sampling_enable {
	    counters1=  "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'start_server_conn_succ’: Start Server Conn Success; 'conn_not_exist’: Conn does not exist; 'data_event’: Data event from TCP; 'client_fin’: FIN from client; 'server_fin’: FIN from server; 'wbuf_event’: Wbuf event from TCP; 'wbuf_cb_failed’: Wbuf event callback failed; 'err_event’: Err event from TCP; 'err_cb_failed’: Err event callback failed; 'server_conn_failed’: Server connection failed; 'client_rst’: RST from client; 'server_rst’: RST from server; 'client_rst_req’: RST from client - request; 'client_rst_connecting’: RST from client - connecting; 'client_rst_connected’: RST from client - connected; 'client_rst_rsp’: RST from client - response; 'server_rst_req’: RST from server - request; 'server_rst_connecting’: RST from server - connecting; 'server_rst_connected’: RST from server - connected; 'server_rst_rsp’: RST from server - response; 'proxy_v1_connection’: counter for Proxy v1 connection; 'proxy_v2_connection’: counter for Proxy v2 connection; 'curr_proxy’: Curr proxy conn; 'curr_proxy_client’: Curr proxy conn - client; 'curr_proxy_server’: Curr proxy conn - server; 'curr_proxy_es’: Curr proxy conn - ES; 'total_proxy’: Total proxy conn; 'total_proxy_client’: Total proxy conn - client; 'total_proxy_server’: Total proxy conn - server; 'total_proxy_es’: Total proxy conn - ES; 'server_select_fail’: Server selection fail; 'est_event’: Est event from TCP; 'est_cb_failed’: Est event callback fail; 'data_cb_failed’: Data event callback fail; 'hps_fwdreq_fail’: Fwd req fail; 'hps_fwdreq_fail_buff’: Fwd req fail - buff; 'hps_fwdreq_fail_rport’: Fwd req fail - rport; 'hps_fwdreq_fail_route’: Fwd req fail - route; 'hps_fwdreq_fail_persist’: Fwd req fail - persist; 'hps_fwdreq_fail_server’: Fwd req fail - server; 'hps_fwdreq_fail_tuple’: Fwd req fail - tuple;
