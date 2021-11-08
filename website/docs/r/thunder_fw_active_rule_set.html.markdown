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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_fw_active_rule_set" "resourceFwActiveRuleSetTest" {
	name = "string"
session_aging = "string"
override_nat_aging = 0
uuid = "string"
 
}

```

## Argument Reference

* `active-rule-set` - Active firewall policy
* `name` - Rule set name
* `session-aging` - Session Aging Template
* `override-nat-aging` - Override NAT idle-timeout
* `uuid` - uuid of the object

