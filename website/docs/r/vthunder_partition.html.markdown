---
layout: "vthunder"
page_title: "vthunder: vthunder_partition"
sidebar_current: "docs-vthunder-resource-partition"
description: |-
    Provides details about vthunder partition resource for A10
---

# vthunder\_partition

`vthunder_partition` provides details about partition
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_partition" "partition"{
user_tag="tag1"
partition_name="part8"
application_type="adc"
id2=8
}
```

## Argument Reference

* `user_tag` - Customized tag
* `partition_name` - Object partition name
* `application_type` - 'adc’: Application type ADC; 'cgnv6’: Application type CGNv6;
* `id2` - Specify unique Partition id

