---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_policy"
sidebar_current: "docs-vthunder-resource-slb_template_policy"
description: |-
    Provides details about vthunder SLB template policy resource for A10
---

# vthunder\_slb\_template\_policy

`vthunder_slb_template_policy` Provides details about vthunder SLB template policy
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_policy" "policy" {
	name = "testdynamicpolicy"
	user_tag = "test_tag"
	use_destination_ip = 1
	over_limit = 1
	timeout = 1
	over_limit_lockup = 2
	over_limit_reset = 1
	overlap = 1
}
```

## Argument Reference

* `name` - Policy template name
* `user_tag` - Customized tag
* `use_destination_ip` - Use destination IP to match the policy
* `over_limit` - Specify operation in case over limit
* `timeout` - Define timeout value of PBSLB dynamic entry (Timeout value (minute, default is 5))
* `over_limit_lockup` - Don’t accept any new connection for certain time (Lockup duration (minute))
* `over_limit_reset` - Reset the connection when it exceeds limit
* `overlap` - Use overlap mode for geo-location to do longest match