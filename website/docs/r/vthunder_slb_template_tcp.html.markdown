---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_tcp"
sidebar_current: "docs-vthunder-resource-slb_template_tcp"
description: |-
    Provides details about vthunder SLB template tcp resource for A10
---

# vthunder\_slb\_template\_tcp

`vthunder_slb_template_tcp` Provides details about vthunder SLB template tcp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_tcp" "tcp"{
    name= "tcp2"
    idle_timeout= 120
    insert_client_ip= 0
    lan_fast_ack= 0
    reset_fwd= 0
    reset_rev= 0
    reset_follow_fin= 0
    del_session_on_server_down= 0
}
```

## Argument Reference

* `name` - TCP Template Name
* `idle_timeout` - Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)
* `insert_client_ip` - Insert client ip into TCP option
* `reset_fwd` - send reset to server if error happens
* `reset_rev` - send reset to client if error happens
* `del_session_on_server_down` - Delete session if the server/port goes down (either disabled/hm down)


