---
layout: "thunder"
page_title: "thunder: thunder_slb_template_external_service"
sidebar_current: "docs-thunder-resource-slb-template-external-service"
description: |-
	Provides details about thunder slb template external service resource for A10
---

# thunder\_slb\_template\_external\_service

`thunder_slb_template_external_service` Provides details about thunder slb template external service
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_external_service" "Slb_Template_External_Service_Test" {

name = "string"
shared_partition_persist_source_ip_template = 0
type = "string"
source_ip = "string"
request-header-forward-list {   
        request_header_forward =  "string" 
        }
template_tcp_proxy_shared = "string"
user_tag = "string"
shared_partition_tcp_proxy_template = 0
action = "string"
service_group = "string"
failure_action = "string"
timeout = 0
tcp_proxy = "string"
template_persist_source_ip_shared = "string"
bypass-ip-cfg {   
        mask =  "string" 
        bypass_ip =  "string" 
        }
uuid = "string"
 
}

```

## Argument Reference

* `action` - ‘continue’: Continue; ‘drop’: Drop; ‘reset’: Reset;
* `failure_action` - ‘continue’: Continue; ‘drop’: Drop; ‘reset’: Reset;
* `name` - External Service Template Name
* `service_group` - Bind a Service Group to the template (Service Group Name)
* `shared_partition_persist_source_ip_template` - Reference a persist source ip template from shared partition
* `shared_partition_tcp_proxy_template` - Reference a TCP Proxy template from shared partition
* `source_ip` - Source IP persistence template (Source IP persistence template name)
* `tcp_proxy` - TCP proxy template (TCP proxy template name)
* `template_persist_source_ip_shared` - Source IP Persistence Template Name
* `template_tcp_proxy_shared` - TCP Proxy Template name
* `timeout` - Timeout value 1 - 200 in units of 200ms, default is 5 (default is 1000ms) (1 - 200 in units of 200ms, default is 5 (1000ms))
* `type` - ‘skyfire-icap’: Skyfire ICAP service; ‘url-filter’: URL filtering service;
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `bypass_ip` - ip address to bypass external service
* `mask` - IP prefix mask
* `request_header_forward` - Request header to be forwarded to external service (Header Name)
