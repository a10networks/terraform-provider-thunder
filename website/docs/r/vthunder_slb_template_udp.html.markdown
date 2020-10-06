---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_udp"
sidebar_current: "docs-vthunder-resource-slb_template_udp"
description: |-
    Provides details about vthunder SLB template udp resource for A10
---

# vthunder\_slb\_template\_udp

`vthunder_slb_template_udp` Provides details about vthunder SLB template udp
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_udp" "udp"{
    qos = 12
    name = "udp2"
    stateless_conn_timeout = 120
    idle_timeout = 120
    re_select_if_server_down = 1
    immediate = 1
    user_tag = "tag1"
}
```

## Argument Reference

* `qos` - QOS level (number)
* `name` - Fast UDP Template Name
* `stateless_conn_timeout` - Stateless Current Connection Timeout value (5 - 120 seconds) (idle timeout in second, default 120)
* `idle_timeout` - Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)
* `re_select_if_server_down` - re-select another server if service port is down
* `immediate` - Immediate Removal after Transaction
* `user_tag` - Customized tag



