---
layout: "thunder"
page_title: "thunder: thunder_fw_alg_tftp"
sidebar_current: "docs-thunder-resource-fw-alg-tftp"
description: |-
	Provides details about thunder fw alg tftp resource for A10
---

# thunder\_fw\_alg\_tftp

`thunder_fw_alg_tftp` Provides details about thunder fw alg tftp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_alg_tftp" "Fw_Alg_Tftp_Test" {
default_port_disable = "string"
sampling-enable {   
        counters1 =  "string" 
        }
uuid = "string"
 
}

```

## Argument Reference

* `default_port_disable` - ‘default-port-disable’: Disable TFTP ALG default port 69;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘session-created’: TFTP Client Sessions Created; ‘helper-created’: TFTP Helper Sessions created; ‘helper-freed’: TFTP Helper Sessions freed; ‘helper-freed-used’: TFTP Helper Sessions freed used; ‘helper-freed-unused’: TFTP Helper Sessions freed unused; ‘helper-already-used’: TFTP Helper Session already used; ‘helper-in-rml’: TFTP Helper Session in Remove List;

