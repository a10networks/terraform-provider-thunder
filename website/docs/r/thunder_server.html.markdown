---
layout: "thunder"
page_title: "thunder: thunder_server"
sidebar_current: "docs-thunder-resource-server"
description: |-
    Provides details about thunder server resource for A10
---

# thunder\_server

`thunder_server` Provides details about thunder server
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_server" "resourceServerTest" {
	name = "string"
server_ipv6_addr = "string"
host = "string"
fqdn_name = "string"
resolve_as = "string"
use_aam_server = 0
ethernet = 0
trunk = 0
action = "string"
external_ip = "string"
ipv6 = "string"
template_server = "string"
shared_partition_server_template = 0
template_server_shared = "string"
template_link_cost = "string"
health_check = "string"
l2_health_check_path = "string"
shared_partition_health_check = 0
health_check_shared = "string"
health_check_disable = 0
conn_limit = 0
no_logging = 0
conn_resume = 0
weight = 0
slow_start = 0
spoofing_cache = 0
stats_data_action = "string"
extended_stats = 0
alternate-server {   
	alternate =  0 
	alternate_name =  "string" 
	}
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
port-list {   
	port_number =  0 
	protocol =  "string" 
	range =  0 
	template_port =  "string" 
	shared_partition_port_template =  0 
	template_port_shared =  "string" 
	template_server_ssl =  "string" 
	action =  "string" 
	no_ssl =  0 
	health_check =  "string" 
	shared_rport_health_check =  0 
	rport_health_check_shared =  "string" 
	health_check_follow_port =  0 
	follow_port_protocol =  "string" 
	health_check_disable =  0 
	support_http2 =  0 
	weight =  0 
	conn_limit =  0 
	no_logging =  0 
	conn_resume =  0 
	stats_data_action =  "string" 
	extended_stats =  0 
alternate-port {   
	alternate =  0 
	alternate_name =  "string" 
	alternate_server_port =  0 
	}
auth_cfg {  
 	service_principal_name =  "string" 
	}
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

* `server` - Server
* `name` - Server Name
* `server-ipv6-addr` - IPV6 address
* `host` - IP Address
* `fqdn-name` - Server hostname
* `resolve-as` - 'resolve-to-ipv4': Use A Query only to resolve FQDN; 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;
* `use-aam-server` - Using aam server. For health check, please configure it in aam server
* `ethernet` - ethernet interface
* `trunk` - trunk interface
* `action` - 'enable': enable; 'disable': disable; 'disable-with-health-check': disable port, but health check work;
* `external-ip` - External IP address for NAT of GSLB
* `ipv6` - IPv6 address Mapping of GSLB
* `template-server` - Server template (Server template name)
* `shared-partition-server-template` - Reference a server template from shared partition
* `template-server-shared` - Server Template Name
* `template-link-cost` - Link-Cost template (Link-Cost template name)
* `health-check` - Health Check (Monitor Name)
* `l2-health-check-path` - L2 health check path
* `shared-partition-health-check` - Reference a health-check from shared partition
* `health-check-shared` - Health Check Monitor (Health monitor name)
* `health-check-disable` - Disable health check
* `conn-limit` - Connection Limit
* `no-logging` - Do not log connection over limit event
* `conn-resume` - Connection Resume
* `weight` - Port Weight (Connection Weight)
* `slow-start` - Slowly ramp up the connection number after server is up (start from 128, then double every 10 sec till 4096)
* `spoofing-cache` - This server is a spoofing cache
* `stats-data-action` - 'stats-data-enable': Enable statistical data collection for real server port; 'stats-data-disable': Disable statistical data collection for real server port;
* `extended-stats` - Enable extended statistics on real server port
* `alternate` - Alternate Server Number
* `alternate-name` - Alternate Name
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'curr_req': Current requests; 'total_req': Total Requests; 'total_req_succ': Total requests succ; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total connections; 'last_total_conn': Last total connections; 'peak_conn': Peak connections; 'es_resp_200': Response status 200; 'es_resp_300': Response status 300; 'es_resp_400': Response status 400; 'es_resp_500': Response status 500; 'es_resp_other': Response status other; 'es_req_count': Total proxy requests; 'es_resp_count': Total proxy response; 'es_resp_invalid_http': Total non-http response; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_good_status_code': Total reverse packets with good status code inspected; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'resp-count': Total Response Count; 'resp-1xx': Response status 1xx; 'resp-2xx': Response status 2xx; 'resp-3xx': Response status 3xx; 'resp-4xx': Response status 4xx; 'resp-5xx': Response status 5xx; 'resp-other': Response status Other; 'resp-latency': Time to First Response Byte; 'curr_pconn': Current persistent connections;
* `port-number` - Port Number
* `protocol` - 'tcp': TCP Port; 'udp': UDP Port;
* `range` - Port range (Port range value - used for vip-to-rport-mapping and vport-rport range mapping)
* `template-port` - Port template (Port template name)
* `shared-partition-port-template` - Reference a port template from shared partition
* `template-port-shared` - Port Template Name
* `template-server-ssl` - Server side SSL template (Server side SSL Name)
* `no-ssl` - No SSL
* `shared-rport-health-check` - Reference a health-check from shared partition
* `rport-health-check-shared` - Health Check (Monitor Name)
* `health-check-follow-port` - Specify which port to follow for health status (Port Number)
* `follow-port-protocol` - 'tcp': TCP Port; 'udp': UDP Port;
* `support-http2` - Starting HTTP/2 with Prior Knowledge
* `alternate-server-port` - Port (Alternate Server Port Value)
* `service-principal-name` - Service Principal Name (Kerberos principal name)
* `packet-capture-template` - Name of the packet capture template to be bind with this object

