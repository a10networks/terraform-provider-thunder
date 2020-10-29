---
layout: "thunder"
page_title: "thunder: thunder_slb_template_logging"
sidebar_current: "docs-thunder-resource-slb-template-logging"
description: |-
	Provides details about thunder slb template logging resource for A10
---

# thunder\_slb\_template\_logging

`thunder_slb_template_logging` Provides details about thunder slb template logging
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_logging" "Slb_Template_Logging_Test" {

pool_shared = "string"
name = "string"
format = "string"
auto = "string"
keep_end = 0
local_logging = 0
mask = "string"
template_tcp_proxy_shared = "string"
shared_partition_tcp_proxy_template = 0
keep_start = 0
service_group = "string"
pcre_mask = "string"
user_tag = "string"
tcp_proxy = "string"
shared_partition_pool = 0
pool = "string"
uuid = "string"
 
}

```

## Argument Reference

* `auto` - ‘auto’: Configure auto NAT for logging, default is auto enabled;
* `format` - Specfiy a format string for web logging (format string(less than 250 characters) for web logging)
* `keep_end` - Number of unmasked characters at the end (default: 0)
* `keep_start` - Number of unmasked characters at the beginning (default: 0)
* `local_logging` - 1 to enable local logging (1 to enable local logging, default 0)
* `mask` - Character to mask the matched pattern (default: X)
* `name` - Logging Template Name
* `pcre_mask` - Mask matched PCRE pattern in the log
* `pool` - Specify NAT pool or pool group
* `pool_shared` - Specify NAT pool or pool group
* `service_group` - Bind a Service Group to the logging template (Service Group Name)
* `shared_partition_pool` - Reference a NAT pool or pool group from shared partition
* `shared_partition_tcp_proxy_template` - Reference a TCP Proxy template from shared partition
* `tcp_proxy` - TCP proxy template (TCP Proxy Config name)
* `template_tcp_proxy_shared` - TCP Proxy Template name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
