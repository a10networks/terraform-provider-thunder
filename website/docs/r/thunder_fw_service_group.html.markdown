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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_fw_service_group" "resourceFwServiceGroupTest" {
	name = "string"
protocol = "string"
health_check = "string"
traffic_replication_mirror_ip_repl = 0
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
packet_capture_template = "string"
member-list {   
	name =  "string" 
	port =  0 
	uuid =  "string" 
	user_tag =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	packet_capture_template =  "string" 
	}
 
}

```

## Argument Reference

* `service-group` - Service Group
* `name` - Member name
* `protocol` - 'tcp': TCP LB service; 'udp': UDP LB service;
* `health-check` - Health Check (Monitor Name)
* `traffic-replication-mirror-ip-repl` - Replaces IP with server-IP
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'curr_conn': Current connections; 'total_fwd_bytes': Total forward bytes; 'total_fwd_pkts': Total forward packets; 'total_rev_bytes': Total reverse bytes; 'total_rev_pkts': Total reverse packets; 'total_conn': Total connections; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_status_code_2xx': Total reverse packets inspected status code 2xx; 'total_rev_pkts_inspected_status_code_non_5xx': Total reverse packets inspected status code non 5xx; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total requests success; 'peak_conn': Peak connections; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'curr_conn_overflow': Current connection counter overflow count; 'state_flaps': State flaps count;
* `packet-capture-template` - Name of the packet capture template to be bind with this object
* `port` - Port number

