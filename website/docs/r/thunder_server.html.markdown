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
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_server" "Server_Test" {
health_check_disable = 0
port-list {   
        health_check_disable =  0 
        protocol =  "string" 
        weight =  0 
        shared_rport_health_check =  0 
        stats_data_action =  "string" 
        health_check_follow_port =  0 
        template_port =  "string" 
        conn_limit =  0 
        template_port_shared =  "string" 
        uuid =  "string" 
        support_http2 =  0 
sampling-enable {   
        counters1 =  "string" 
        }
        no_ssl =  0 
        follow_port_protocol =  "string" 
        template_server_ssl =  "string" 
alternate-port {   
        alternate_name =  "string" 
        alternate =  0 
        alternate_server_port =  0 
        }
        port_number =  0 
        extended_stats =  0 
        rport_health_check_shared =  "string" 
        conn_resume =  0 
        user_tag =  "string" 
        range =  0 
auth_cfg {  
        service_principal_name =  "string" 
        }
        action =  "string" 
        no_logging =  0 
        health_check =  "string" 
        shared_partition_port_template =  0 
        }
stats_data_action = "string"
spoofing_cache = 0
weight = 0
template_server_shared = "string"
slow_start = 0
shared_partition_server_template = 0
resolve_as = "string"
conn_limit = 0
uuid = "string"
fqdn_name = "string"
external_ip = "string"
health_check_shared = "string"
ipv6 = "string"
template_server = "string"
server_ipv6_addr = "string"
alternate-server {   
        alternate_name =  "string" 
        alternate =  0 
        }
shared_partition_health_check = 0
host = "string"
extended_stats = 0
conn_resume = 0
name = "string"
user_tag = "string"
sampling-enable {   
        counters1 =  "string" 
        }
action = "string"
health_check = "string"
no_logging = 0
 
}
```

## Argument Reference

* `action` - ‘enable’: enable; ‘disable’: disable; ‘disable-with-health-check’: disable port, but health check work;
* `conn_limit` - Connection Limit
* `conn_resume` - Connection Resume
* `extended_stats` - Enable extended statistics on real server port
* `external_ip` - External IP address for NAT of GSLB
* `fqdn_name` - Server hostname
* `health_check` - Health Check (Monitor Name)
* `health_check_disable` - Disable health check
* `health_check_shared` - Health Check Monitor (Health monitor name)
* `host` - IP Address
* `ipv6` - IPv6 address Mapping of GSLB
* `name` - Server Name
* `no_logging` - Do not log connection over limit event
* `resolve_as` - ‘resolve-to-ipv4’: Use A Query only to resolve FQDN; ‘resolve-to-ipv6’: Use AAAA Query only to resolve FQDN; ‘resolve-to-ipv4-and-ipv6’: Use A as well as AAAA Query to resolve FQDN;
* `server_ipv6_addr` - IPV6 address
* `shared_partition_health_check` - Reference a health-check from shared partition
* `slow_start` - Slowly ramp up the connection number after server is up (start from 128, then double every 10 sec till 4096)
* `spoofing_cache` - This server is a spoofing cache
* `stats_data_action` - ‘stats-data-enable’: Enable statistical data collection for real server port; ‘stats-data-disable’: Disable statistical data collection for real server port;
* `template_server` - Server template (Server template name)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `weight` - Port Weight (Connection Weight)
* `follow_port_protocol` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `health_check_follow_port` - Specify which port to follow for health status (Port Number)
* `no_ssl` - No SSL
* `port_number` - Port Number
* `protocol` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `range` - Port range (Port range value - used for vip-to-rport-mapping)
* `rport_health_check_shared` - Health Check (Monitor Name)
* `shared_rport_health_check` - Reference a health-check from shared partition
* `support_http2` - Starting HTTP/2 with Prior Knowledge
* `template_port` - Port template (Port template name)
* `template_server_ssl` - Server side SSL template (Server side SSL Name)
* `counters1` - ‘all’: all; ‘total-conn’: Total connections; ‘fwd-pkt’: Forward packets; ‘rev-pkt’: Reverse packets; ‘peak-conn’: Peak connections; ‘total_req’: Total Requests; ‘total_req_succ’: Total requests succ; ‘curr_ssl_conn’: Current SSL connections; ‘total_ssl_conn’: Total SSL connections;
* `alternate` - Alternate Server (Alternate Server Number)
* `alternate_name` - Alternate Name
* `alternate_server_port` - Port (Alternate Server Port Value)
* `service_principal_name` - Service Principal Name (Kerberos principal name)
