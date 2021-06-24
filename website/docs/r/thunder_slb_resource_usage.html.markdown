---
layout: "thunder"
page_title: "thunder: thunder_slb_resource_usage"
sidebar_current: "docs-thunder-resource-slb-resource-usage"
description: |-
	Provides details about thunder slb resource usage resource for A10
---

# thunder\_slb\_resource\_usage

`thunder_slb_resource_usage` Provides details about thunder slb resource usage
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_resource_usage" "Slb_Resource_Usage_Test" {

slb_threshold_res_usage_percent = 0
proxy_template_count = 0
nat_pool_addr_count = 0
fast_tcp_template_count = 0
cache_template_count = 0
health_monitor_count = 0
fast_udp_template_count = 0
virtual_port_count = 0
client_ssl_template_count = 0
persist_srcip_template_count = 0
server_ssl_template_count = 0
http_template_count = 0
pbslb_subnet_count = 0
persist_cookie_template_count = 0
virtual_server_count = 0
stream_template_count = 0
conn_reuse_template_count = 0
real_server_count = 0
real_port_count = 0
service_group_count = 0
uuid = "string"
 
}
```

## Argument Reference

* `cache_template_count` - Total configurable HTTP Cache Templates in the System
* `client_ssl_template_count` - Total configurable Client SSL Templates in the System
* `conn_reuse_template_count` - Total configurable Connection reuse Templates in the System
* `fast_tcp_template_count` - Total configurable Fast TCP Templates in the System
* `fast_udp_template_count` - Total configurable Fast UDP Templates in the System
* `health_monitor_count` - Total Health Monitors in the System
* `http_template_count` - Total configurable HTTP Templates in the System
* `nat_pool_addr_count` - Total configurable NAT Pool addresses in the System
* `pbslb_subnet_count` - Total PBSLB Subnets in the System
* `persist_cookie_template_count` - Total configurable Persistent cookie Templates in the System
* `persist_srcip_template_count` - Total configurable Source IP Persistent Templates in the System
* `proxy_template_count` - Total configurable Proxy Templates in the System
* `real_port_count` - Total Real Server Ports in the System
* `real_server_count` - Total Real Servers in the System
* `server_ssl_template_count` - Total configurable Server SSL Templates in the System
* `service_group_count` - Total Service Groups in the System
* `slb_threshold_res_usage_percent` - Enter the threshold as a percentage (Threshold in percentage(default is 0%))
* `stream_template_count` - Total configurable Streaming media in the System
* `uuid` - uuid of the object
* `virtual_port_count` - Total Virtual Server Ports in the System
* `virtual_server_count` - Total Virtual Servers in the System
