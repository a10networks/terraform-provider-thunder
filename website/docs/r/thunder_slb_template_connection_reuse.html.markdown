---
layout: "thunder"
page_title: "thunder: thunder_slb_template_connection_reuse"
sidebar_current: "docs-thunder-resource-slb-template-connection-reuse"
description: |-
	Provides details about thunder slb template connection reuse resource for A10
---

# thunder\_slb\_template\_connection\_reuse

`thunder_slb_template_connection_reuse` Provides details about thunder slb template connection reuse
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_connection_reuse" "Slb_Template_Connection_Reuse_Test" {

preopen = 0
uuid = "string"
keep_alive_conn = 0
user_tag = "string"
limit_per_server = 0
timeout = 0
num_conn_per_port = 0
name = "string"
 
}

```

## Argument Reference

* `keep_alive_conn` - Keep a number of server connections open
* `limit_per_server` - Max Server Connections allowed (Connections per Server Port (default 1000))
* `name` - Connection Reuse Template Name
* `num_conn_per_port` - Connections per Server Port (default 100)
* `preopen` - Preopen server connection
* `timeout` - Timeout in seconds. Multiple of 60 (def 2400)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
