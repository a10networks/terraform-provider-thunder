---
layout: "thunder"
page_title: "thunder: thunder_slb_ssl_expire_check"
sidebar_current: "docs-thunder-resource-slb-ssl-expire-check"
description: |-
	Provides details about thunder slb ssl expire check resource for A10
---

# thunder\_slb\_ssl\_expire\_check

`thunder_slb_ssl_expire_check` Provides details about thunder slb ssl expire check
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_ssl_expire_check" "Slb_Ssl_Expire_Check_Test" {
exception {  
        action =  "string" 
        certificate_name =  "string" 
        }
uuid = "string"
ssl_expire_email_address = "string"
interval_days = 0
expire_address1 = "string"
before = 0
 
}

```

## Argument Reference

* `before` - The number of days in advance notice before expiration, default is 5
* `expire_address1` - Email address for certificate expiration check
* `interval_days` - The interval of days notice after expiration, default is 2
* `ssl_expire_email_address` - Config Email address for certificate expiration check
* `uuid` - uuid of the object
* `action` - ‘add’: Add an exception; ‘delete’: Delete an exception; ‘clean’: Delete all exception;
* `certificate_name` - The certificate name
