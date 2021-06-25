---
layout: "thunder"
page_title: "thunder: thunder_slb_server_port"
sidebar_current: "docs-thunder-resource-slb-server-port"
description: |-
	Provides details about thunder slb server port resource for A10
---

# thunder\_slb\_server\_port

`thunder_slb_server_port` Provides details about thunder slb server port
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_server_port" "Slb_Server_Port_Test" {

health_check_disable = 0
protocol = "string"
weight = 0
shared_rport_health_check = 0
stats_data_action = "string"
health_check_follow_port = 0
template_port = "string"
conn_limit = 0
template_port_shared = "string"
uuid = "string"
support_http2 = 0
sampling-enable {   
        counters1 =  "string" 
        }
no_ssl = 0
follow_port_protocol = "string"
template_server_ssl = "string"
alternate-port {   
        alternate_name =  "string" 
        alternate =  0 
        alternate_server_port =  0 
        }
port_number = 0
extended_stats = 0
rport_health_check_shared = "string"
conn_resume = 0
user_tag = "string"
range = 0
auth_cfg {  
        service_principal_name =  "string" 
        }
action = "string"
no_logging = 0
health_check = "string"
shared_partition_port_template = 0
 
}
```

## Argument Reference

* `action` - ‘enable’: enable; ‘disable’: disable; ‘disable-with-health-check’: disable port, but health check work;
* `conn_limit` - Connection Limit
* `conn_resume` - Connection Resume
* `extended_stats` - Enable extended statistics on real server port
* `follow_port_protocol` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `health_check` - Health Check (Monitor Name)
* `health_check_disable` - Disable health check
* `health_check_follow_port` - Specify which port to follow for health status (Port Number)
* `no_logging` - Do not log connection over limit event
* `no_ssl` - No SSL
* `port_number` - Port Number
* `protocol` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `range` - Port range (Port range value - used for vip-to-rport-mapping)
* `rport_health_check_shared` - Health Check (Monitor Name)
* `shared_rport_health_check` - Reference a health-check from shared partition
* `stats_data_action` - ‘stats-data-enable’: Enable statistical data collection for real server port; ‘stats-data-disable’: Disable statistical data collection for real server port;
* `support_http2` - Starting HTTP/2 with Prior Knowledge
* `template_port` - Port template (Port template name)
* `template_server_ssl` - Server side SSL template (Server side SSL Name)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `weight` - Port Weight (Connection Weight)
* `counters1` - ‘all’: all; ‘curr_req’: Current requests; ‘total_req’: Total Requests; ‘total_req_succ’: Total requests succ; ‘total_fwd_bytes’: Forward bytes; ‘total_fwd_pkts’: Forward packets; ‘total_rev_bytes’: Reverse bytes; ‘total_rev_pkts’: Reverse packets; ‘total_conn’: Total connections; ‘last_total_conn’: Last total connections; ‘peak_conn’: Peak connections; ‘es_resp_200’: Response status 200; ‘es_resp_300’: Response status 300; ‘es_resp_400’: Response status 400; ‘es_resp_500’: Response status 500; ‘es_resp_other’: Response status other; ‘es_req_count’: Total proxy request; ‘es_resp_count’: Total proxy response; ‘es_resp_invalid_http’: Total non-http response; ‘total_rev_pkts_inspected’: Total reverse packets inspected; ‘total_rev_pkts_inspected_good_status_code’: Total reverse packets with good status code inspected; ‘response_time’: Response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time; ‘curr_ssl_conn’: Current SSL connections; ‘total_ssl_conn’: Total SSL connections;
* `alternate` - Alternate Server Number
* `alternate_name` - Alternate Name
* `alternate_server_port` - Port (Alternate Server Port Value)
* `service_principal_name` - Service Principal Name (Kerberos principal name)

