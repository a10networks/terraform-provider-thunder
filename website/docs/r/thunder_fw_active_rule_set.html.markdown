---
layout: "thunder"
page_title: "thunder: thunder_fw_active_rule_set"
sidebar_current: "docs-thunder-resource-fw-active-rule-set"
description: |-
	Provides details about thunder fw active rule set resource for A10
---

# thunder\_fw\_active\_rule\_set

`thunder_fw_active_rule_set` Provides details about thunder fw active rule set
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_active_rule_set" "Fw_Active_Rule_Set_Test" {
session_aging = "string"
override_nat_aging = 0
name = "string"
uuid = "string"
 
}

```

## Argument Reference

* `name` - Rule set name
* `override_nat_aging` - Override NAT idle-timeout
* `session_aging` - Session Aging Template
* `uuid` - uuid of the object
