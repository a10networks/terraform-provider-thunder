---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_mqtt"
sidebar_current: "docs-vthunder-resource-slb_template_mqtt"
description: |-
    Provides details about vthunder SLB template mqtt resource for A10
---

# vthunder\_slb\_template\_mqtt

`vthunder_slb_template_mqtt` Provides details about vthunder SLB template mqtt
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_mqtt" "mqtt"{
    user_tag="tag2"
    clientid_hash_persist=0
    name="mqtt2"
}
```

## Argument Reference

* `name` - MQTT Template Name
* `clientid_hash_persist` - Immediate Removal after Transaction
* `user_tag` - Customized tag



