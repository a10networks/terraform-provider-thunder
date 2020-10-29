---
layout: "thunder"
page_title: "thunder: thunder_slb_template_respmod_icap"
sidebar_current: "docs-thunder-resource-slb-template-respmod-icap"
description: |-
	Provides details about thunder slb template respmod icap resource for A10
---

# thunder\_slb\_template\_respmod\_icap

`thunder_slb_template_respmod_icap` Provides details about thunder slb template respmod icap
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_reqmod_icap" "Slb_Template_Reqmod_Icap_Test" {
        min_payload_size = 0
shared_partition_persist_source_ip_template = 0
template_tcp_proxy_shared = "string"
allowed_http_methods = "string"
uuid = "string"
source_ip = "string"
shared_partition_tcp_proxy_template = 0
service_group = "string"
tcp_proxy = "string"
preview = 0
disable_http_server_reset = 0
server_ssl = "string"
fail_close = 0
bypass-ip-cfg {   
        mask =  "string" 
        bypass_ip =  "string" 
        }
template_persist_source_ip_shared = "string"
include_protocol_in_uri = 0
logging = "string"
name = "string"
user_tag = "string"
x_auth_url = 0
log_only_allowed_method = 0
action = "string"
cylance = 0
service_url = "string"
 
}
```

## Argument Reference

* `action` - ‘continue’: Continue; ‘drop’: Drop; ‘reset’: Reset;
* `cylance` - cylance external server
* `disable_http_server_reset` - Don’t reset http server
* `fail_close` - When template sg is down mark vport down
* `include_protocol_in_uri` - Include protocol and port in HTTP URI
* `log_only_allowed_method` - Only log allowed HTTP method
* `logging` - logging template (Logging template name)
* `min_payload_size` - min-payload-size value 1 - 65536, default is 4096
* `name` - Reqmod ICAP Template Name
* `preview` - Preview value 1 - 32768, default is 32768
* `server_ssl` - Server SSL template (Server SSL template name)
* `service_group` - Bind a Service Group to the template (Service Group Name)
* `service_url` - URL to send to ICAP server (Service URL Name)
* `shared_partition_persist_source_ip_template` - Reference a persist source ip template from shared partition
* `shared_partition_tcp_proxy_template` - Reference a TCP Proxy template from shared partition
* `source_ip` - Source IP persistence template (Source IP persistence template name)
* `tcp_proxy` - TCP proxy template (TCP proxy template name)
* `template_persist_source_ip_shared` - Source IP Persistence Template Name
* `template_tcp_proxy_shared` - TCP Proxy Template name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `bypass_ip` - ip address to bypass respmod-icap service
* `mask` - IP prefix mask
