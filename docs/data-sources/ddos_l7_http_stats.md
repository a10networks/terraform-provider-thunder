---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_l7_http_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_l7_http_stats: Statistics for the object l7-http
  PLACEHOLDER
---

# thunder_ddos_l7_http_stats (Data Source)

`thunder_ddos_l7_http_stats`: Statistics for the object l7-http

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_l7_http_stats" "thunder_ddos_l7_http_stats" {
}
output "get_ddos_l7_http_stats" {
  value = ["${data.thunder_ddos_l7_http_stats.thunder_ddos_l7_http_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `agent_filter_blacklist` (Number) Agent Filter Blacklisted
- `agent_filter_match` (Number) Agent Filter Match
- `alloc_fail` (Number) Alloc Failed
- `challenge_fail` (Number) Challenge Failed
- `challenge_js_fail` (Number) Challenge Javascript Failed
- `challenge_js_sent` (Number) Challenge Javascript Sent
- `challenge_ud_fail` (Number) Challenge URL Redirect Failed
- `challenge_ud_sent` (Number) Challenge URL Redirect Sent
- `chunk_bad` (Number) Bad HTTP Chunk
- `chunk_sz_1k` (Number) Payload Chunk Size Less Than or Equal to 1K
- `chunk_sz_2k` (Number) Payload Chunk Size Less Than or Equal to 2K
- `chunk_sz_4k` (Number) Payload Chunk Size Less Than or Equal to 4K
- `chunk_sz_512` (Number) Payload Chunk Size Less Than or Equal to 512
- `chunk_sz_gt_4k` (Number) Payload Chunk Size Larger Than 4K
- `client_rst` (Number) Client TCP RST Received
- `ddos_policy_violation` (Number) Policy Violation
- `dst_filter_action_blacklist` (Number) Dst Filter Action Blacklist
- `dst_filter_action_default_pass` (Number) Dst Filter Action Default Pass
- `dst_filter_action_drop` (Number) Dst Filter Action Drop
- `dst_filter_action_ignore` (Number) Dst Filter Action Ignore
- `dst_filter_action_reset` (Number) Dst Filter Action Reset
- `dst_filter_action_whitelist` (Number) Dst Filter Action WL
- `dst_filter_match` (Number) Dst Filter Match
- `dst_filter_not_match` (Number) Dst Filter No Match
- `dst_filter_rate_exceed` (Number) Dst Filter Rate Exceed
- `dst_post_rate_exceed` (Number) Dst Post Rate Exceeded
- `dst_req_rate_exceed` (Number) Dst Request Rate Exceeded
- `dst_resp_rate_exceed` (Number) Dst Response Rate Exceeded
- `error_condition` (Number) Error Condition
- `header_name_too_long` (Number) HTTP Header Name Too Long
- `header_processing_incomplete` (Number) Header Process Incomplete
- `header_processing_time_0` (Number) Header Process Time Less Than 1s
- `header_processing_time_1` (Number) Header Process Time Less Than 10s
- `header_processing_time_2` (Number) Header Process Time Less Than 30s
- `header_processing_time_3` (Number) Header Process Time Larger or Equal to 30s
- `hps_req_sz_16k` (Number) Request Payload Size Less Than or Equal to 16K
- `hps_req_sz_1k` (Number) Request Payload Size Less Than or Equal to 1K
- `hps_req_sz_256k` (Number) Request Payload Size Less Than or Equal to 256K
- `hps_req_sz_256k_plus` (Number) Request Payload Size Larger Than 256K
- `hps_req_sz_2k` (Number) Request Payload Size Less Than or Equal to 2K
- `hps_req_sz_32k` (Number) Request Payload Size Less Than or Equal to 32K
- `hps_req_sz_4k` (Number) Request Payload Size Less Than or Equal to 4K
- `hps_req_sz_64k` (Number) Request Payload Size Less Than or Equal to 64K
- `hps_req_sz_8k` (Number) Request Payload Size Less Than or Equal to 8K
- `hps_rsp_10` (Number) Response HTTP 1.0
- `hps_rsp_11` (Number) Response HTTP 1.1
- `hps_rsp_status_1xx` (Number) Status Code 1XX
- `hps_rsp_status_2xx` (Number) Status Code 2XX
- `hps_rsp_status_3xx` (Number) Status Code 3XX
- `hps_rsp_status_4xx` (Number) Status Code 4XX
- `hps_rsp_status_504_ax` (Number) Status Code 504 AX-Gen
- `hps_rsp_status_5xx` (Number) Status Code 5XX
- `hps_rsp_status_6xx` (Number) Status Code 6XX
- `hps_rsp_status_unknown` (Number) Status Code Unknown
- `hps_rsp_sz_16k` (Number) Response Payload Size Less Than or Equal to 16K
- `hps_rsp_sz_1k` (Number) Response Payload Size Less Than or Equal to 1K
- `hps_rsp_sz_256k` (Number) Response Payload Size Less Than or Equal to 256K
- `hps_rsp_sz_256k_plus` (Number) Response Payload Size Larger Than 256K
- `hps_rsp_sz_2k` (Number) Response Payload Size Less Than or Equal to 2K
- `hps_rsp_sz_32k` (Number) Response Payload Size Less Than or Equal to 32K
- `hps_rsp_sz_4k` (Number) Response Payload Size Less Than or Equal to 4K
- `hps_rsp_sz_64k` (Number) Response Payload Size Less Than or Equal to 64K
- `hps_rsp_sz_8k` (Number) Response Payload Size Less Than or Equal to 8K
- `hps_server_rst` (Number) Server TCP RST Received
- `http10` (Number) Request HTTP 1.0
- `http11` (Number) Request HTTP 1.1
- `http_auth_drop` (Number) HTTP Auth Dropped
- `http_auth_resp` (Number) HTTP Auth Responded
- `http_connect` (Number) Request Method CONNECT
- `http_del` (Number) Request Method DELETE
- `http_get` (Number) Request Method GET
- `http_head` (Number) Request Method HEAD
- `http_idle_timeout` (Number) Http Idle Timeout
- `http_options` (Number) Request Method OPTIONS
- `http_post` (Number) Request Method POST
- `http_put` (Number) Request Method PUT
- `http_trace` (Number) Request Method TRACE
- `http_unknown` (Number) Request Method UNKNOWN
- `invalid_hdr_name` (Number) HTTP Header Name Invalid
- `invalid_hdr_val` (Number) HTTP Header Value Invalid
- `invalid_header` (Number) HTTP Header Invalid
- `line_too_long` (Number) Line Too Long
- `lower_than_mss_exceed` (Number) Min Payload Size Fail Exceeded
- `malform_bad_chunk` (Number) Malform Bad Chunk
- `malform_content_len_too_long` (Number) Malform Content Length Too Long
- `malform_header_name_too_long` (Number) Malform Header Name Too Long
- `malform_line_too_long` (Number) Malform Line Too Long
- `malform_req_line_invalid_method` (Number) Malform Request Line Invalid Method
- `malform_req_line_too_long` (Number) Malform Request Line Too Long
- `malform_req_line_too_small` (Number) Malform Request Line Too Small
- `malform_too_many_headers` (Number) Malform Too Many Headers
- `neg_req_remain` (Number) Negative Request Remain
- `neg_rsp_remain` (Number) Negative Response Remain
- `new_syn` (Number) TCP SYN
- `ofo` (Number) Out-Of-Order Packets
- `ofo_queue_exceed` (Number) Out-Of-Order Queue Exceeded
- `ofo_timer_expired` (Number) Out-Of-Order Timeout
- `parsereq_fail` (Number) Parse Request Failed
- `partial_hdr` (Number) Partial Header
- `policy_drop` (Number) Policy Dropped
- `referer_filter_blacklist` (Number) Referer Filter Blacklisted
- `referer_filter_match` (Number) Referer Filter Match
- `req_content_len` (Number) Request Content-Length Received
- `req_ofo` (Number) Out-Of-Order Request
- `req_processed` (Number) Packets Processed
- `req_retrans` (Number) Retransmit Request
- `request` (Number) Request Total
- `retrans` (Number) TCP Retransmit
- `retrans_fin` (Number) TCP Retransmit FIN
- `retrans_push` (Number) TCP Retransmit PSH
- `retrans_rst` (Number) TCP Retransmit RST
- `rsp_chunk` (Number) Response Chunk
- `src_dst_filter_action_blacklist` (Number) SrcDst Filter Action Blacklist
- `src_dst_filter_action_default_pass` (Number) SrcDst Filter Action Default Pass
- `src_dst_filter_action_drop` (Number) SrcDst Filter Action Drop
- `src_dst_filter_action_whitelist` (Number) SrcDst Filter Action WL
- `src_dst_filter_match` (Number) SrcDst Filter Match
- `src_dst_filter_not_match` (Number) SrcDst Filter No Match
- `src_filter_action_blacklist` (Number) Src Filter Action Blacklist
- `src_filter_action_default_pass` (Number) Src Filter Action Default Pass
- `src_filter_action_drop` (Number) Src Filter Action Drop
- `src_filter_action_whitelist` (Number) Src Filter Action WL
- `src_filter_match` (Number) Src Filter Match
- `src_filter_not_match` (Number) Src Filter No Match
- `src_post_rate_exceed` (Number) Src Post Rate Exceeded
- `src_req_rate_exceed` (Number) Src Request Rate Exceeded
- `too_many_headers` (Number) HTTP Header Too Many
- `uri_filter_match` (Number) URI Filter Match
- `use_hdr_ip_as_source` (Number) Use IP In Header As Src
- `window_small` (Number) Window Size Small
- `window_small_drop` (Number) Window Size Small Dropped

