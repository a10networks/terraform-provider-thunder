---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_sip"
sidebar_current: "docs-vthunder-resource-slb_template_sip"
description: |-
    Provides details about vthunder SLB template sip resource for A10
---

# vthunder\_slb\_template\_sip

`vthunder_slb_template_sip` Provides details about vthunder SLB template sip
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_sip" "templatesip" {
	name = "testsip"
	user_tag = "test_tag"
	smp_call_id_rtp_session = 1
	client_keep_alive = 1
	alg_source_nat = 1
	alg_dest_nat = 1
	server_keep_alive = 1
	interval = 11
	dialog_aware = 1
	failed_server_selection = 1
	timeout = 1
	drop_when_server_fail = 1
}
```

## Argument Reference

* `name` - SIP Template Name
* `user_tag` - Customized tag
* `smp_call_id_rtp_session` - Create the across cpu call-id rtp session
* `client_keep_alive` - Respond client keep-alive packet directly instead of forwarding to server
* `alg_source_nat` - Translate source IP to NAT IP in SIP message when source NAT is used
* `alg_dest_nat` - Translate VIP to real server IP in SIP message when destination NAT is used
* `server_keep_alive` - Send server keep-alive packet for every persist connection when enable conn-reuse
* `interval` - The interval of keep-alive packet for each persist connection (second)
* `dialog_aware` - Permit system processes dialog session
* `failed_server_selection` - Define action when select server fail
* `timeout` - Time in minutes
* `drop_when_server_fail` - Drop current SIP message when select server fail



