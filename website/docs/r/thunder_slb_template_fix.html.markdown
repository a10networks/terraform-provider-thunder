---
layout: "thunder"
page_title: "thunder: thunder_slb_template_fix"
sidebar_current: "docs-thunder-resource-slb-template-fix"
description: |-
	Provides details about thunder slb template fix resource for A10
---

# thunder\_slb\_template\_fix

`thunder_slb_template_fix` Provides details about thunder slb template fix
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_fix" "Slb_Template_Fix_Test" {

logging = "string"
name = "string"
tag-switching {   
        service_group =  "string" 
        equals =  "string" 
        switching_type =  "string" 
        }
user_tag = "string"
insert_client_ip = 0
uuid = "string"
 
}

```

## Argument Reference

* `insert_client_ip` - Insert client ip to tag 11447
* `logging` - ‘init’: init only log; ‘term’: termination only log; ‘both’: both initial and termination log;
* `name` - FIX Template Name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `equals` - Equals (Tag String)
* `service_group` - Create a Service Group comprising Servers (Service Group Name)
* `switching_type` - ‘sender-comp-id’: Select service group based on SenderCompID; ‘target-comp-id’: Select service group based on TargetCompID;
