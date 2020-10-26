---
layout: "thunder"
page_title: "thunder: thunder_partition"
sidebar_current: "docs-thunder-resource-partition"
description: |-
    Provides details about thunder partition resource for A10
---

# thunder\_partition

`thunder_partition` provides details about partition
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_partition" "partition"{
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

