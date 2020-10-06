---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_port"
sidebar_current: "docs-vthunder-resource-slb_template_port"
description: |-
    Provides details about vthunder SLB template port resource for A10
---

# vthunder\_slb\_template\_port

`vthunder_slb_template_port` Provides details about vthunder SLB template port
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_port" "port" {
	name = "testport"
	user_tag = "test_tag"
	slow_start = 1
	conn_limit = 2
	weight = 5
	extended_stats = 1
	del_session_on_server_down = 1
}
```

## Argument Reference

* `name` - PORT Template Name
* `user_tag` - Customized tag
* `slow_start` - Slowly ramp up the connection number after port is up
* `conn_limit` - Connection limit
* `weight` - Weight (port weight)
* `extended_stats` - Enable extended statistics on real server port
* `del_session_on_server_down` - Delete session if the server/port goes down (either disabled/hm down)




