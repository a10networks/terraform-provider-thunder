---
layout: "thunder"
page_title: "thunder: thunder_fw_server"
sidebar_current: "docs-thunder-resource-fw-server"
description: |-
    Provides details about thunder fw server resource for A10
---

# thunder\_fw\_server

`thunder_fw_server` Provides details about thunder fw server
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_fw_server" "resourceFwServerTest" {
	name = "string"
server_ipv6_addr = "string"
host = "string"
fqdn_name = "string"
resolve_as = "string"
action = "string"
health_check = "string"
health_check_disable = 0
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
port-list {   
	port_number =  0 
	protocol =  "string" 
	action =  "string" 
	health_check =  "string" 
	health_check_disable =  0 
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

* `server` - Firewall logging Server
* `name` - Server Name
* `server-ipv6-addr` - IPV6 address
* `host` - IP Address
* `fqdn-name` - Server hostname
* `resolve-as` - 'resolve-to-ipv4': Use A Query only to resolve FQDN; 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;
* `action` - 'enable': enable; 'disable': disable;
* `health-check` - Health Check (Monitor Name)
* `health-check-disable` - Disable health check
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'curr_conn': Current connections; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total request success; 'total_fwd_bytes': Forward bytes; 'total_fwd_pkts': Forward packets; 'total_rev_bytes': Reverse bytes; 'total_rev_pkts': Reverse packets; 'total_conn': Total connections; 'last_total_conn': Last total connections; 'peak_conn': Peak connections; 'es_resp_200': Response status 200; 'es_resp_300': Response status 300; 'es_resp_400': Response status 400; 'es_resp_500': Response status 500; 'es_resp_other': Response status other; 'es_req_count': Total proxy request; 'es_resp_count': Total proxy Response; 'es_resp_invalid_http': Total non-http response; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_good_status_code': Total reverse packets with good status code inspected; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time;
* `port-number` - Port Number
* `protocol` - 'tcp': TCP Port; 'udp': UDP Port;
* `packet-capture-template` - Name of the packet capture template to be bind with this object

