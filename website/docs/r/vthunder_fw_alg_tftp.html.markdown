---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_alg_tftp"
sidebar_current: "docs-vthunder-resource-fw-alg-tftp"
description: |-
	Provides details about vthunder fw alg tftp resource for A10
---

# vthunder\_fw\_alg\_tftp

`vthunder_fw_alg_tftp` Provides details about vthunder fw alg tftp
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_alg_tftp" "FwAlgTest" {
	default_port_disable = "default-port-disable" 
}
```

## Argument Reference

* `default_port_disable` - ‘default-port-disable’: Disable TFTP ALG default port 69;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘session-created’: TFTP Client Sessions Created; ‘helper-created’: TFTP Helper Sessions created; ‘helper-freed’: TFTP Helper Sessions freed; ‘helper-freed-used’: TFTP Helper Sessions freed used; ‘helper-freed-unused’: TFTP Helper Sessions freed unused; ‘helper-already-used’: TFTP Helper Session already used; ‘helper-in-rml’: TFTP Helper Session in Remove List;

