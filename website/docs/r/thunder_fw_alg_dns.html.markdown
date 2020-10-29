---
layout: "thunder"
page_title: "thunder: thunder_fw_alg_dns"
sidebar_current: "docs-thunder-resource-fw-alg-dns"
description: |-
	Provides details about thunder fw alg dns resource for A10
---

# thunder\_fw\_alg\_dns

`thunder_fw_alg_dns` Provides details about thunder fw alg dns
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_alg_dns" "Fw_Alg_Dns_Test" {
default_port_disable = "string"
uuid = "string"
}

```

## Argument Reference

* `default_port_disable` - ‘default-port-disable’: Disable DNS ALG default port 53;
* `uuid` - uuid of the object

