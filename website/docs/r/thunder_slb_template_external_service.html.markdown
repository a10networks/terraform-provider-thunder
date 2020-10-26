---
layout: "thunder"
page_title: "thunder: thunder_slb_template_external_service"
sidebar_current: "docs-thunder-resource-slb-template-external-service"
description: |-
    Provides details about thunder slb template external-service resource for A10
---

# thunder\_slb\_template\_external\_service

`thunder_slb_template_external_service` provides details about slb template external-service
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_external_service" "external_service" {
	name = "testexternalservice"
	user_tag = "test_tag"
	type = "skyfire-icap"
	action = "continue"
	failure_action = "continue"
	timeout = 10
	request_header_forward_list {
        request_header_forward = "testf"
      }
}
```

## Argument Reference

* `name` - External Service Template Name
* `user_tag` - Customized tag
* `type` - 'skyfire-icap’: Skyfire ICAP service; 'url-filter’: URL filtering service;
* `action` - 'continue’: Continue; 'drop’: Drop; 'reset’: Reset; 
* `failure_action` - 'continue’: Continue; 'drop’: Drop; 'reset’: Reset;
* `timeout` - Timeout value 1 - 200 in units of 200ms, default is 5 (default is 1000ms) (1 - 200 in units of 200ms, default is 5 (1000ms)) 
* `request_header_forward_list` -
    * `request_header_forward` - Request header to be forwarded to external service (Header Name)


