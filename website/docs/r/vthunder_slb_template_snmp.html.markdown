---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_snmp"
sidebar_current: "docs-vthunder-resource-slb-template-snmp"
description: |-
    Provides details about vthunder slb template snmp resource for A10
---

# vthunder\_slb\_template\_snmp

`vthunder_slb_template_snmp` provides details about slb template snmp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_snmp" "testname" {
	user_tag = "test_tag"
	priv_proto = "aes"
	context_name = "testcont"
	interval = 10
	security_level = "no-auth"
	community = "tttest"
	auth_proto = "sha"
	version = "v1"
	interface = 0
	port = 1770
	snmp_name = "testsnmp"
}
```

## Argument Reference

* `user_tag` - Customized tag
* `priv_proto` - 'aes’: AES; 'des’: DES;
* `context_name` - Specify context name
* `interval` - Specify interval, default is 3 (Interval, unit: second, default is 3)
* `security_level` - 'no-auth’: No authentication; 'auth-no-priv’: Authentication, but no privacy; 'auth-priv’: Authentication and privacy;
* `community` - Specify community for version 2c (Community name)
* `auth_proto` - 'sha’: SHA; 'md5’: MD5;
* `version` - 'v1’: Version 1; 'v2c’: Version 2c; 'v3’: Version 3;
* `interface` - Specify Interface ID
* `port` - Specify port, default is 161 (Port Number, default is 161)
* `snmp_name` - Specify name of snmp template

