---
layout: "thunder"
page_title: "thunder: thunder_slb_resource_usage"
sidebar_current: "docs-thunder-resource-slb-resource-usage"
description: |-
    Provides details about thunder SLB resource usage resource for A10
---

# thunder\_slb\__resource\_usage

`thunder_slb_resource_usage` Provides details about thunder SLB resource usage
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_resource_usage" "resource_usage" {
	real_server_count = 128
	stream_template_count = 128
	proxy_template_count = 128
	http_template_count = 128
	persist_srcip_template_count = 128
	server_ssl_template_count = 128
	service_group_count = 128
	persist_cookie_template_count = 128
	virtual_server_count = 64
	fast_udp_template_count = 128
	fast_tcp_template_count = 128
	virtual_port_count = 128
	conn_reuse_template_count = 128
	real_port_count = 256
	client_ssl_template_count = 128
	nat_pool_addr_count = 10
	health_monitor_count = 1023
	pbslb_subnet_count = 65536 
}
```

## Argument Reference

* `real_server_count` - Total Real Servers in the System
* `stream_template_count` - Total configurable Streaming media in the System
* `proxy_template_count` - Total configurable Proxy Templates in the System
* `http_template_count` - Total configurable HTTP Templates in the System
* `persist_srcip_template_count` - Total configurable Source IP Persistent Templates in the System
* `server_ssl_template_count` - Total configurable Server SSL Templates in the System
* `service_group_count` - Total Service Groups in the System
* `persist_cookie_template_count` - Total configurable Persistent cookie Templates in the System
* `virtual_server_count` - Total Virtual Servers in the System
* `fast_udp_template_count` - Total configurable Fast UDP Templates in the System
* `fast_tcp_template_count` - Total configurable Fast TCP Templates in the System
* `virtual_port_count` - Total Virtual Server Ports in the System
* `conn_reuse_template_count` - Total configurable Connection reuse Templates in the System
* `real_port_count` - Total Real Server Ports in the System
* `client_ssl_template_count` - Total configurable Client SSL Templates in the System
* `nat_pool_addr_count` - Total configurable NAT Pool addresses in the System (deprecated)
* `pbslb_subnet_count` - Total PBSLB Subnets in the System
