---
layout: "thunder"
page_title: "thunder: thunder_slb_template_smpp"
sidebar_current: "docs-thunder-resource-slb-template-smpp"
description: |-
	Provides details about thunder slb template smpp resource for A10
---

# thunder\_slb\_template\_smpp

`thunder_slb_template_smpp` Provides details about thunder slb template smpp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_smpp" "Slb_Template_Smpp_Test" {

name = "string"
server_enquire_link = 0
server_selection_per_request = 0
client_enquire_link = 0
user_tag = "string"
server_enquire_link_val = 0
user = "string"
password = "string"
uuid = "string"
 
}

```

## Argument Reference

* `client_enquire_link` - Respond client ENQUIRE_LINK packet directly instead of forwarding to server
* `name` - SMPP Template Name
* `password` - Configure the password used to bind
* `server_enquire_link` - Send server ENQUIRE_LINK packet for every persist connection when enable conn-reuse
* `server_enquire_link_val` - Set interval of keep-alive packet for each persistent connection (second, default is 30)
* `server_selection_per_request` - Force server selection on every SMPP request when enable conn-reuse
* `user` - Configure the user to bind (The name used to bind)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
