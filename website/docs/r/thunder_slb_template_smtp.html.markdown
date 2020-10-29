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
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_smtp" "Slb_Template_Smtp_Test" {

name = "string"
user_tag = "string"
server_domain = "string"
client-domain-switching {   
        service_group =  "string" 
        match_string =  "string" 
        switching_type =  "string" 
        }
client_starttls_type = "string"
LF_to_CRLF = 0
template {  
        logging =  "string" 
        }
command-disable {   
        disable_type =  "string" 
        }
server_starttls_type = "string"
error_code_to_client = 0
service_ready_msg = "string"
uuid = "string"
 
}
```

## Argument Reference

* `client_starttls_type` - ‘optional’: STARTTLS is optional requirement; ‘enforced’: Must issue STARTTLS command before mail transaction;
* `name` - SMTP Template Name
* `server_domain` - Config the domain of the email servers (Server’s domain, default is “mail-server-domain”)
* `server_starttls_type` - ‘optional’: STARTTLS is optional requirement; ‘enforced’: Must issue STARTTLS command before mail transaction;
* `service_ready_msg` - Set SMTP service ready message (SMTP service ready message, default is “ESMTP mail service ready”)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `match_string` - Domain name string
* `service_group` - Select service group (Service group name)
* `switching_type` - ‘contains’: Specify domain name string if domain contains another string; ‘ends-with’: Specify domain name string if domain ends with another string; ‘starts-with’: Specify domain string if domain starts with another string;
* `logging` - Logging template (Logging Config name)
* `disable_type` - ‘expn’: Disable SMTP EXPN commands; ‘turn’: Disable SMTP TURN commands; ‘vrfy’: Disable SMTP VRFY commands;
