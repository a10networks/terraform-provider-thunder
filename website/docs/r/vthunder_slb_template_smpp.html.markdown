---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_smpp"
sidebar_current: "docs-vthunder-resource-slb_template_smpp"
description: |-
    Provides details about vthunder SLB template smpp resource for A10
---

# vthunder\_slb\_template\_smpp

`vthunder_slb_template_smpp` Provides details about vthunder SLB template smpp
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_smpp" "smpp"{
    name = "smpp2"
    server_enquire_link = 1
    server_selection_per_request = 1
    client_enquire_link = 1
    user_tag = "utag1"
    server_enquire_link_val = 6
    user = "u1"
    password = "pwd1"
}
```

## Argument Reference

* `name` - SMPP Template Name
* `server_enquire_link` - Send server ENQUIRE_LINK packet for every persist connection when enable conn-reuse
* `server_selection_per_request` - Force server selection on every SMPP request when enable conn-reuse
* `client_enquire_link` - Respond client ENQUIRE_LINK packet directly instead of forwarding to server
* `user_tag` - Customized tag
* `server_enquire_link_val` - Set interval of keep-alive packet for each persistent connection (second, default is 30)
* `user` - Configure the user to bind (The name used to bind)
* `password` - Configure the password used to bind




