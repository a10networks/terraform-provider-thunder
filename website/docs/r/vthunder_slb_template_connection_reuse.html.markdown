---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_connection_reuse"
sidebar_current: "docs-vthunder-resource-slb-template-connection-reuse"
description: |-
    Provides details about vthunder slb template connection-reuse resource for A10
---

# vthunder\_slb\_template\_connection\_reuse

`vthunder_slb_template_connection_reuse` provides details about slb template connection-reuse
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_connection_reuse" "connection_reuse" {
	name = "testConn"
	keep_alive_conn = 0
	limit_per_server = 10
	timeout = 120
	user_tag = "testtag"
}
```

## Argument Reference

* `name` - Connection Reuse Template Name
* `keep_alive_conn` - Keep a number of server connections open
* `limit_per_server` - Max Server Connections allowed (Connections per Server Port (default 1000))
* `timeout` - Timeout in seconds. Multiple of 60 (default 2400)
* `user_tag` - Customized tag

