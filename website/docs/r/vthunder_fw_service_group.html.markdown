---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_service_group"
sidebar_current: "docs-vthunder-resource-fw-service-group"
description: |-
	Provides details about vthunder fw service group resource for A10
---

# vthunder\_fw\_service\_group

`vthunder_fw_service_group` Provides details about vthunder fw service group
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

// Put working JSON here
```

## Argument Reference

* `health_check` - Health Check (Monitor Name)
* `name` - Member name
* `protocol` - ‘tcp’: TCP LB service; ‘udp’: UDP LB service;
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘curr_conn’: Current connections; ‘total_fwd_bytes’: Total forward bytes; ‘total_fwd_pkts’: Total forward packets; ‘total_rev_bytes’: Total reverse bytes; ‘total_rev_pkts’: Total reverse packets; ‘total_conn’: Total connections; ‘total_rev_pkts_inspected’: Total reverse packets inspected; ‘total_rev_pkts_inspected_status_code_2xx’: Total reverse packets inspected status code 2xx; ‘total_rev_pkts_inspected_status_code_non_5xx’: Total reverse packets inspected status code non 5xx; ‘curr_req’: Current requests; ‘total_req’: Total requests; ‘total_req_succ’: Total requests success; ‘peak_conn’: Peak connections; ‘response_time’: Response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time;
* `port` - Port number

