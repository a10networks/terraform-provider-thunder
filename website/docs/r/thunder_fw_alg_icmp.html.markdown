---
layout: "thunder"
page_title: "thunder: thunder_fw_alg_icmp"
sidebar_current: "docs-thunder-resource-fw-alg-icmp"
description: |-
	Provides details about thunder fw alg icmp resource for A10
---

# thunder\_fw\_alg\_icmp

`thunder_fw_alg_icmp` Provides details about thunder fw alg icmp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_alg_icmp" "FwAlgTest" {
	disable = "disable"
  uuid = "string" 
}
```

## Argument Reference

* `disable` - ‘disable’: Disable ICMP ALG which allows ICMP errors to pass the firewall;
* `uuid` - uuid of the object

