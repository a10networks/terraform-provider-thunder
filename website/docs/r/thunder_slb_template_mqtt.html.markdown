---
layout: "thunder"
page_title: "thunder: thunder_slb_template_mqtt"
sidebar_current: "docs-thunder-resource-slb_template_mqtt"
description: |-
    Provides details about thunder SLB template mqtt resource for A10
---

# thunder\_slb\_template\_mqtt

`thunder_slb_template_mqtt` Provides details about thunder SLB template mqtt
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_mqtt" "mqtt"{
    user_tag="tag2"
    clientid_hash_persist=0
    name="mqtt2"
}
```

## Argument Reference

* `name` - MQTT Template Name
* `clientid_hash_persist` - Immediate Removal after Transaction
* `user_tag` - Customized tag



