---
layout: "thunder"
page_title: "thunder: thunder_fw_service_group"
sidebar_current: "docs-thunder-resource-fw-service-group"
description: |-
	Provides details about thunder fw service group resource for A10
---

# thunder\_fw\_service\_group

`thunder_fw_service_group` Provides details about thunder fw service group
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_service_group" "Fw_Service_Group_Test" {
protocol = "string"
uuid = "string"
user_tag = "string"
traffic_replication_mirror_ip_repl = 0
sampling-enable {   
        counters1 =  "string" 
        }
member-list {   
        port =  0 
sampling-enable {   
        counters1 =  "string" 
        }
        uuid =  "string" 
        user_tag =  "string" 
        name =  "string" 
        }
health_check = "string"
name = "string"
 
}
```

## Argument Reference

* `health_check` - Health Check (Monitor Name)
* `name` - Member name
* `protocol` - ‘tcp’: TCP LB service; ‘udp’: UDP LB service;
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘curr_conn’: Current connections; ‘total_fwd_bytes’: Total forward bytes; ‘total_fwd_pkts’: Total forward packets; ‘total_rev_bytes’: Total reverse bytes; ‘total_rev_pkts’: Total reverse packets; ‘total_conn’: Total connections; ‘total_rev_pkts_inspected’: Total reverse packets inspected; ‘total_rev_pkts_inspected_status_code_2xx’: Total reverse packets inspected status code 2xx; ‘total_rev_pkts_inspected_status_code_non_5xx’: Total reverse packets inspected status code non 5xx; ‘curr_req’: Current requests; ‘total_req’: Total requests; ‘total_req_succ’: Total requests success; ‘peak_conn’: Peak connections; ‘response_time’: Response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time;
* `port` - Port number

