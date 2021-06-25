---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_contact"
sidebar_current: "docs-thunder-resource-snmp-server-contact"
description: |-
	Provides details about thunder snmp server contact resource for A10
---

# thunder\_snmp\_server\_contact

`thunder_snmp_server_contact` Provides details about thunder snmp server contact
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_snmp_server_contact" "Snmp_Server_Contact_Test" {
        contact_name = "string"
uuid = "string"
 
}


```

## Argument Reference

* `contact_name` - Identification of the contact person for this managed node
* `uuid` - uuid of the object

