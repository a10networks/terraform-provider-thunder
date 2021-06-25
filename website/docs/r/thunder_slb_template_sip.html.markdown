---
layout: "thunder"
page_title: "thunder: thunder_slb_template_sip"
sidebar_current: "docs-thunder-resource-slb-template-sip"
description: |-
	Provides details about thunder slb template sip resource for A10
---

# thunder\_slb\_template\_sip

`thunder_slb_template_sip` Provides details about thunder slb template sip
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_sip" "Slb_Template_Sip_Test" {

server-request-header {   
        server_request_header_insert =  "string" 
        server_request_erase_all =  0 
        insert_condition_server_request =  "string" 
        server_request_header_erase =  "string" 
        }
smp_call_id_rtp_session = 0
keep_server_ip_if_match_acl = 0
client_keep_alive = 0
alg_source_nat = 0
uuid = "string"
server-response-header {   
        server_response_header_insert =  "string" 
        insert_condition_server_response =  "string" 
        server_response_header_erase =  "string" 
        server_response_erase_all =  0 
        }
server_selection_per_request = 0
client-request-header {   
        client_request_header_erase =  "string" 
        client_request_header_insert =  "string" 
        client_request_erase_all =  0 
        insert_condition_client_request =  "string" 
        }
pstn_gw = "string"
service_group = "string"
insert_client_ip = 0
failed_client_selection = 0
failed_client_selection_message = "string"
call_id_persist_disable = 0
acl_id = 0
alg_dest_nat = 0
server_keep_alive = 0
client-response-header {   
        client_response_erase_all =  0 
        insert_condition_client_response =  "string" 
        client_response_header_erase =  "string" 
        client_response_header_insert =  "string" 
        }
failed_server_selection_message = "string"
name = "string"
exclude-translation {   
        translation_value =  "string" 
        header_string =  "string" 
        }
interval = 0
user_tag = "string"
dialog_aware = 0
failed_server_selection = 0
drop_when_client_fail = 0
timeout = 0
drop_when_server_fail = 0
acl_name_value = "string"
 
}

```

## Argument Reference

* `acl_id` - ACL id
* `acl_name_value` - IPv4 Access List Name
* `alg_dest_nat` - Translate VIP to real server IP in SIP message when destination NAT is used
* `alg_source_nat` - Translate source IP to NAT IP in SIP message when source NAT is used
* `call_id_persist_disable` - Disable call-ID persistence
* `client_keep_alive` - Respond client keep-alive packet directly instead of forwarding to server
* `dialog_aware` - Permit system processes dialog session
* `drop_when_client_fail` - Drop current SIP message when select client fail
* `drop_when_server_fail` - Drop current SIP message when select server fail
* `failed_client_selection` - Define action when select client fail
* `failed_client_selection_message` - Send SIP message (includs status code) to server when select client fail(Format: 3 digits(1XX~6XX) space reason)
* `failed_server_selection` - Define action when select server fail
* `failed_server_selection_message` - Send SIP message (includs status code) to client when select server fail(Format: 3 digits(1XX~6XX) space reason)
* `insert_client_ip` - Insert Client IP address into SIP header
* `interval` - The interval of keep-alive packet for each persist connection (second)
* `keep_server_ip_if_match_acl` - Use Real Server IP for addresses matching the ACL for a Call-Id
* `name` - SIP Template Name
* `pstn_gw` - configure pstn gw host name for tel: uri translate to sip: uri (Hostname String, default is “pstn”)
* `server_keep_alive` - Send server keep-alive packet for every persist connection when enable conn-reuse
* `server_selection_per_request` - Force server selection on every SIP request
* `service_group` - service group name
* `smp_call_id_rtp_session` - Create the across cpu call-id rtp session
* `timeout` - Time in minutes
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `insert_condition_server_request` - ‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;
* `server_request_erase_all` - Erase all headers
* `server_request_header_erase` - Erase a SIP header (Header Name)
* `server_request_header_insert` - Insert a SIP header (Header Content (Format: “name: value”))
* `insert_condition_server_response` - ‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;
* `server_response_erase_all` - Erase all headers
* `server_response_header_erase` - Erase a SIP header (Header Name)
* `server_response_header_insert` - Insert a SIP header (Header Content (Format: “name: value”))
* `client_request_erase_all` - Erase all headers
* `client_request_header_erase` - Erase a SIP header (Header Name)
* `client_request_header_insert` - Insert a SIP header (Header Content (Format: “name: value”))
* `insert_condition_client_request` - ‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;
* `client_response_erase_all` - Erase all headers
* `client_response_header_erase` - Erase a SIP header (Header Name)
* `client_response_header_insert` - Insert a SIP header (Header Content (Format: “name: value”))
* `insert_condition_client_response` - ‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;
* `header_string` - SIP header name
* `translation_value` - ‘start-line’: SIP request line or status line; ‘header’: SIP message headers; ‘body’: SIP message body;
