---
layout: "thunder"
page_title: "thunder: thunder_slb_template_snmp"
sidebar_current: "docs-thunder-resource-slb-template-snmp"
description: |-
    Provides details about thunder slb template snmp resource for A10
---

# thunder\_slb\_template\_snmp

`thunder_slb_template_snmp` provides details about slb template snmp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_snmp" "snmp" {
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

