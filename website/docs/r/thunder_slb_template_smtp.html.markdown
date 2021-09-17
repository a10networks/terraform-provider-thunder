---
layout: "thunder"
page_title: "thunder: thunder_slb_template_smtp"
sidebar_current: "docs-thunder-resource-slb-template-smtp"
description: |-
    Provides details about thunder slb template smtp resource for A10
---

# thunder\_slb\_template\_smtp

`thunder_slb_template_smtp` Provides details about thunder slb template smtp
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_smtp" "resourceSlbTemplateSmtpTest" {
	name = "string"
server_domain = "string"
service_ready_msg = "string"
client_starttls_type = "string"
server_starttls_type = "string"
command-disable {   
	disable_type =  "string" 
	}
LF_to_CRLF = 0
error_code_to_client = 0
client-domain-switching {   
	switching_type =  "string" 
	match_string =  "string" 
	service_group =  "string" 
	}
template {  
 	logging =  "string" 
	}
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `smtp` - SMTP
* `name` - SMTP Template Name
* `server-domain` - Config the domain of the email servers (Server's domain, default is "mail-server-domain")
* `service-ready-msg` - Set SMTP service ready message (SMTP service ready message, default is "ESMTP mail service ready")
* `client-starttls-type` - 'optional': STARTTLS is optional requirement; 'enforced': Must issue STARTTLS command before mail transaction;
* `server-starttls-type` - 'optional': STARTTLS is optional requirement; 'enforced': Must issue STARTTLS command before mail transaction;
* `disable-type` - 'expn': Disable SMTP EXPN commands; 'turn': Disable SMTP TURN commands; 'vrfy': Disable SMTP VRFY commands;
* `LF-to-CRLF` - Change the LF to CRLF for smtp end of line
* `error-code-to-client` - Would transfer error code(554) to client, when getting it from connection establishing with real-server
* `switching-type` - 'contains': Specify domain name string if domain contains another string; 'ends-with': Specify domain name string if domain ends with another string; 'starts-with': Specify domain string if domain starts with another string;
* `match-string` - Domain name string
* `service-group` - Select service group (Service group name)
* `logging` - Logging template (Logging Config name)
* `uuid` - uuid of the object
* `user-tag` - Customized tag

