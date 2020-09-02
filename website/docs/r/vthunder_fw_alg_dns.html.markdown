---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_alg_dns"
sidebar_current: "docs-vthunder-resource-fw-alg-dns"
description: |-
	Provides details about vthunder fw alg dns resource for A10
---

# vthunder\_fw\_alg\_dns

`vthunder_fw_alg_dns` Provides details about vthunder fw alg dns
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_alg_dns" "FwAlgTest" {
	default_port_disable = "default-port-disable" 
}
```

## Argument Reference

* `default_port_disable` - ‘default-port-disable’: Disable DNS ALG default port 53;
* `uuid` - uuid of the object

