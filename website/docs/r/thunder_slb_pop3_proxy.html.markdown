---
layout: "thunder"
page_title: "thunder: thunder_slb_pop3_proxy"
sidebar_current: "docs-thunder-resource-slb-pop3-proxy"
description: |-
    Provides details about thunder SLB pop3 proxy resource for A10
---

# thunder\_slb\_pop3\_proxy

`thunder_slb_pop3_proxy` Provides details about thunder SLB pop3 proxy
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_pop3_proxy" "pop3_proxy" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'num’: Num; 'curr’: Current proxy conns; 'total’: Total proxy conns; 'svrsel_fail’: Server selection failure; 'no_route’: no route failure; 'snat_fail’: source nat failure; 'line_too_long’: line too long; 'line_mem_freed’: request line freed; 'invalid_start_line’: invalid start line; 'stls’: stls cmd; 'request_dont_care’: other cmd; 'unsupported_command’: Unsupported cmd; 'bad_sequence’: Bad Sequence; 'rsv_persist_conn_fail’: Serv Sel Persist fail; 'smp_v6_fail’: Serv Sel SMPv6 fail; 'smp_v4_fail’: Serv Sel SMPv4 fail; 'insert_tuple_fail’: Serv Sel insert tuple fail; 'cl_est_err’: Client EST state erro; 'ser_connecting_err’: Serv CTNG state error; 'server_response_err’: Serv RESP state error; 'cl_request_err’: Client RQ state error; 'request’: Total POP3 Request; 'control_to_ssl’: Control chn ssl;