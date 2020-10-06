---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_templateFix"
sidebar_current: "docs-vthunder-resource-slb_templateFix"
description: |-
    Provides details about vthunder SLB template fix resource for A10
---

# vthunder\_slb\_template\_fix

`vthunder_TemplateFix` Provides details about vthunder SLB template fix
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_fix" "templatefix" {
	name = "testfix"
	logging = "init"
	tag_switching {
		equals = "test"
		service_group = "testsg"
		switching_type = "sender-comp-id"
	} 
	insert_client_ip = 1111
	user_tag = "test_user"
}
```

## Argument Reference

* `name` - FIX Template Name
* `logging` - 'init’: init only log; 'term’: termination only log; 'both’: both initial and termination log
* `service_group` - Create a Service Group comprising Servers (Service Group Name)
* `switching_type` - 'sender-comp-id’: Select service group based on SenderCompID; 'target-comp-id’: Select service group based on TargetComp
* `insert_client_ip` - Insert client ip to tag 11447
* `user_tag` - Customized tag




