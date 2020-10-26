---
layout: "thunder"
page_title: "thunder: thunder_slb_icap"
sidebar_current: "docs-thunder-resource-slb-icap"
description: |-
    Provides details about thunder SLB icap resource for A10
---

# thunder\_slb\_icap

`thunder_slb_icap` Provides details about thunder SLB icap
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_icap" "icap" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'reqmod_request’: Reqmod Request Stats; 'respmod_request’: Respmod Request Stats; 'reqmod_request_after_100’: Reqmod Request Sent After 100 Cont Stats; 'respmod_request_after_100’: Respmod Request Sent After 100 Cont Stats; 'reqmod_response’: Reqmod Response Stats; 'respmod_response’: Respmod Response Stats; 'reqmod_response_after_100’: Reqmod Response After 100 Cont Stats; 'respmod_response_after_100’: Respmod Response After 100 Cont Stats; 'chunk_no_allow_204’: Chunk so no Allow 204 Stats; 'len_exceed_no_allow_204’: Length Exceeded so no Allow 204 Stats; 'result_continue’: Result Continue Stats; 'result_icap_response’: Result ICAP Response Stats; 'result_100_continue’: Result 100 Continue Stats; 'result_other’: Result Other Stats; 'status_2xx’: Status 2xx Stats; 'status_200’: Status 200 Stats; 'status_201’: Status 201 Stats; 'status_202’: Status 202 Stats; 'status_203’: Status 203 Stats; 'status_204’: Status 204 Stats; 'status_205’: Status 205 Stats; 'status_206’: Status 206 Stats; 'status_207’: Status 207 Stats; 'status_1xx’: Status 1xx Stats; 'status_100’: Status 100 Stats; 'status_101’: Status 101 Stats; 'status_102’: Status 102 Stats; 'status_3xx’: Status 3xx Stats; 'status_300’: Status 300 Stats; 'status_301’: Status 301 Stats; 'status_302’: Status 302 Stats; 'status_303’: Status 303 Stats; 'status_304’: Status 304 Stats; 'status_305’: Status 305 Stats; 'status_306’: Status 306 Stats; 'status_307’: Status 307 Stats; 'status_4xx’: Status 4xx Stats; 'status_400’: Status 400 Stats; 'status_401’: Status 401 Stats; 'status_402’: Status 402 Stats; 'status_403’: Status 403 Stats; 'status_404’: Status 404 Stats; 'status_405’: Status 405 Stats; 'status_406’: Status 406 Stats; 'status_407’: Status 407 Stats; 'status_408’: Status 408 Stats; 'status_409’: Status 409 Stats; 'status_410’: Status 410 Stats; 'status_411’: Status 411 Stats; 'status_412’: Status 412 Stats; 'status_413’: Status 413 Stats; 'status_414’: Status 414 Stats; 'status_415’: Status 415 Stats; 'status_416’: Status 416 Stats; 'status_417’: Status 417 Stats; 'status_418’: Status 418 Stats; 'status_419’: Status 419 Stats; 'status_420’: Status 420 Stats; 'status_422’: Status 422 Stats; 'status_423’: Status 423 Stats; 'status_424’: Status 424 Stats; 'status_425’: Status 425 Stats; 'status_426’: Status 426 Stats; 'status_449’: Status 449 Stats; 'status_450’: Status 450 Stats; 'status_5xx’: Status 5xx Stats; 'status_500’: Status 500 Stats; 'status_501’: Status 501 Stats; 'status_502’: Status 502 Stats; 'status_503’: Status 503 Stats; 'status_504’: Status 504 Stats; 'status_505’: Status 505 Stats; 'status_506’: Status 506 Stats; 'status_507’: Status 507 Stats; 'status_508’: Status 508 Stats; 'status_509’: Status 509 Stats; 'status_510’: Status 510 Stats; 'status_6xx’: Status 6xx Stats; 'status_unknown’: Status Unknown Stats; 'send_option_req’: Send Option Req Stats; 'app_serv_conn_no_pcb_err’: App Server Conn no ES PCB Err Stats; 'app_serv_conn_err’: App Server Conn Err Stats; 'chunk1_hdr_err’: Chunk Hdr Err1 Stats; 'chunk2_hdr_err’: Chunk Hdr Err2 Stats; 'chunk_bad_trail_err’: Chunk Bad Trail Err Stats; 'no_payload_next_buff_err’: No Payload In Next Buff Err Stats; 'no_payload_buff_err’: No Payload Buff Err Stats; 'resp_hdr_incomplete_err’: Resp Hdr Incomplete Err Stats; 'serv_sel_fail_err’: Server Select Fail Err Stats; 'start_icap_conn_fail_err’: Start ICAP conn fail Stats; 'prep_req_fail_err’: Prepare ICAP req fail Err Stats; 'icap_ver_err’: ICAP Ver Err Stats; 'icap_line_err’: ICAP Line Err Stats; 'encap_hdr_incomplete_err’: Encap HDR Incomplete Err Stats; 'no_icap_resp_err’: No ICAP Resp Err Stats; 'resp_line_read_err’: Resp Line Read Err Stats; 'resp_line_parse_err’: Resp Line Parse Err Stats; 'resp_hdr_err’: Resp Hdr Err Stats; 'req_hdr_incomplete_err’: Req Hdr Incomplete Err Stats; 'no_status_code_err’: No Status Code Err Stats; 'http_resp_line_read_err’: HTTP Response Line Read Err Stats; 'http_resp_line_parse_err’: HTTP Response Line Parse Err Stats; 'http_resp_hdr_err’: HTTP Resp Hdr Err Stats; 'recv_option_resp’: Send Option Req Stats;