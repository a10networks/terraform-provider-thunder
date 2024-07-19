---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_server_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_server_stats: Statistics for the object server
  PLACEHOLDER
---

# thunder_fw_server_stats (Data Source)

`thunder_fw_server_stats`: Statistics for the object server

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_server_stats" "thunder_fw_server_stats" {

  name = "test"
}
output "get_fw_server_stats" {
  value = ["${data.thunder_fw_server_stats.thunder_fw_server_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Server Name

### Optional

- `port_list` (Block List) (see [below for nested schema](#nestedblock--port_list))
- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--port_list"></a>
### Nested Schema for `port_list`

Required:

- `port_number` (Number) Port Number
- `protocol` (String) 'tcp': TCP Port; 'udp': UDP Port;

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--port_list--stats))

<a id="nestedblock--port_list--stats"></a>
### Nested Schema for `port_list.stats`

Optional:

- `curr_conn` (Number) Current connections
- `curr_req` (Number) Current requests
- `es_req_count` (Number) Total proxy request
- `es_resp_200` (Number) Response status 200
- `es_resp_300` (Number) Response status 300
- `es_resp_400` (Number) Response status 400
- `es_resp_500` (Number) Response status 500
- `es_resp_count` (Number) Total proxy Response
- `es_resp_invalid_http` (Number) Total non-http response
- `es_resp_other` (Number) Response status other
- `fastest_rsp_time` (Number) Fastest response time
- `last_total_conn` (Number) Last total connections
- `peak_conn` (Number) Peak connections
- `response_time` (Number) Response time
- `slowest_rsp_time` (Number) Slowest response time
- `total_conn` (Number) Total connections
- `total_fwd_bytes` (Number) Forward bytes
- `total_fwd_pkts` (Number) Forward packets
- `total_req` (Number) Total requests
- `total_req_succ` (Number) Total request success
- `total_rev_bytes` (Number) Reverse bytes
- `total_rev_pkts` (Number) Reverse packets
- `total_rev_pkts_inspected` (Number) Total reverse packets inspected
- `total_rev_pkts_inspected_good_status_code` (Number) Total reverse packets with good status code inspected



<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `curr_conn` (Number) Current connections
- `fwd_pkt` (Number) Forward packets
- `peak_conn` (Number) Peak connections
- `rev_pkt` (Number) Reverse Packets
- `total_conn` (Number) Total connections

