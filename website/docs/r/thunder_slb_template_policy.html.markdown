---
layout: "thunder"
page_title: "thunder: thunder_slb_template_policy"
sidebar_current: "docs-thunder-resource-slb-template-policy"
description: |-
	Provides details about thunder slb template policy resource for A10
---

# thunder\_slb\_template\_policy

`thunder_slb_template_policy` Provides details about thunder slb template policy
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_policy" "Slb_Template_Policy_Test" {

forward_policy {  
 filtering {   
        ssli_url_filtering =  "string" 
        }
        uuid =  "string" 
        local_logging =  0 
san-filtering {   
        ssli_url_filtering_san =  "string" 
        }
action-list {   
        log =  0 
        real_sg =  "string" 
        drop_response_code =  0 
        forward_snat =  "string" 
        uuid =  "string" 
        http_status_code =  "string" 
        action1 =  "string" 
        fake_sg =  "string" 
        user_tag =  "string" 
        proxy_chaining_bypass =  0 
        drop_message =  "string" 
sampling-enable {   
        counters1 =  "string" 
        }
        fall_back =  "string" 
        fall_back_snat =  "string" 
        drop_redirect_url =  "string" 
        name =  "string" 
        }
        no_client_conn_reuse =  0 
        require_web_category =  0 
        acos_event_log =  0 
source-list {   
        match_any =  0 
        name =  "string" 
        match_authorize_policy =  "string" 
destination {  
 class-list-list {   
        uuid =  "string" 
        dest_class_list =  "string" 
        priority =  0 
sampling-enable {   
        counters1 =  "string" 
        }
        action =  "string" 
        type =  "string" 
        }
web-category-list-list {   
        uuid =  "string" 
        web_category_list =  "string" 
        priority =  0 
sampling-enable {   
        counters1 =  "string" 
        }
        action =  "string" 
        type =  "string" 
        }
any {  
        action =  "string" 
sampling-enable {   
        counters1 =  "string" 
        }
        uuid =  "string" 
        }
        }
        user_tag =  "string" 
        priority =  0 
sampling-enable {   
        counters1 =  "string" 
        }
        match_class_list =  "string" 
        uuid =  "string" 
        }
        }
use_destination_ip = 0
name = "string"
over_limit = 0
class_list {  
        header_name =  "string" 
lid-list {   
        request_rate_limit =  0 
        action_value =  "string" 
        request_per =  0 
        bw_rate_limit =  0 
        conn_limit =  0 
        log =  0 
        direct_action_value =  "string" 
        conn_per =  0 
        direct_fail =  0 
        conn_rate_limit =  0 
        direct_pbslb_logging =  0 
dns64 {  
        prefix =  "string" 
        exclusive_answer =  0 
        disable =  0 
        }
        lidnum =  0 
        over_limit_action =  0 
response-code-rate-limit {   
        threshold =  0 
        code_range_end =  0 
        code_range_start =  0 
        period =  0 
        }
        direct_service_group =  "string" 
        uuid =  "string" 
        request_limit =  0 
        direct_action_interval =  0 
        bw_per =  0 
        interval =  0 
        user_tag =  "string" 
        direct_action =  0 
        lockout =  0 
        direct_logging_drp_rst =  0 
        direct_pbslb_interval =  0 
        }
        name =  "string" 
        client_ip_l3_dest =  0 
        client_ip_l7_header =  0 
        uuid =  "string" 
        }
interval = 0
share = 0
full_domain_tree = 0
over_limit_logging = 0
bw_list_name = "string"
timeout = 0
sampling-enable {   
        counters1 =  "string" 
        }
user_tag = "string"
bw-list-id {   
        pbslb_interval =  0 
        action_interval =  0 
        service_group =  "string" 
        logging_drp_rst =  0 
        fail =  0 
        pbslb_logging =  0 
        id =  0 
        bw_list_action =  "string" 
        }
over_limit_lockup = 0
uuid = "string"
over_limit_reset = 0
overlap = 0
 
}
```

## Argument Reference

* `bw_list_name` - Specify a blacklist/whitelist name
* `full_domain_tree` - Share counters between geo-location and sub regions
* `interval` - Specify log interval in minutes, by default system will log every over limit instance
* `name` - Class list name or geo-location-class-list name
* `over_limit` - Specify operation in case over limit
* `over_limit_lockup` - Don’t accept any new connection for certain time (Lockup duration (minute))
* `over_limit_logging` - Log a message
* `over_limit_reset` - Reset the connection when it exceeds limit
* `overlap` - Use overlap mode for geo-location to do longest match
* `share` - Share counters between virtual ports and virtual servers
* `timeout` - Define timeout value of PBSLB dynamic entry (Timeout value (minute, default is 5))
* `use_destination_ip` - Use destination IP to match the policy
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `acos_event_log` - Enable acos event logging
* `local_logging` - Enable local logging
* `no_client_conn_reuse` - Inspects only first request of a connection
* `require_web_category` - Wait for web category to be resolved before taking proxy decision
* `ssli_url_filtering` - ‘bypassed-sni-disable’: Disable SNI filtering for bypassed URL’s(enabled by default); ‘intercepted-sni-enable’: Enable SNI filtering for intercepted URL’s(disabled by default); ‘intercepted-http-disable’: Disable HTTP(host/URL) filtering for intercepted URL’s(enabled by default); ‘no-sni-allow’: Allow connection if SNI filtering is enabled and SNI header is not present(Drop by default);
* `ssli_url_filtering_san` - ‘enable-san’: Enable SAN filtering(disabled by default); ‘bypassed-san-disable’: Disable SAN filtering for bypassed URL’s(enabled by default); ‘intercepted-san-enable’: Enable SAN filtering for intercepted URL’s(disabled by default); ‘no-san-allow’: Allow connection if SAN filtering is enabled and SAN field is not present(Drop by default);
* `action1` - ‘forward-to-internet’: Forward request to Internet; ‘forward-to-service-group’: Forward request to service group; ‘drop’: Drop request;
* `drop_message` - drop-message sent to the client as webpage(html tags are included and quotation marks are required for white spaces)
* `drop_redirect_url` - Specify URL to which client request is redirected upon being dropped
* `drop_response_code` - Specify response code for drop action
* `fake_sg` - service group to forward the packets to Internet
* `fall_back` - Fallback service group for Internet
* `fall_back_snat` - Source NAT pool or pool group for fallback server
* `forward_snat` - Source NAT pool or pool group
* `http_status_code` - ‘301’: Moved permanently; ‘302’: Found;
* `log` - Log a message
* `real_sg` - service group to forward the packets
* `counters1` - ‘all’: all; ‘fwd-policy-dns-unresolved’: Forward-policy unresolved DNS queries; ‘fwd-policy-dns-outstanding’: Forward-policy current DNS outstanding requests; ‘fwd-policy-snat-fail’: Forward-policy source-nat translation failure; ‘fwd-policy-hits’: Number of forward-policy requests for this policy template; ‘fwd-policy-forward-to-internet’: Number of forward-policy requests forwarded to internet; ‘fwd-policy-forward-to-service-group’: Number of forward-policy requests forwarded to service group; ‘fwd-policy-policy-drop’: Number of forward-policy requests dropped; ‘fwd-policy-source-match-not-found’: Forward-policy requests without matching source rule; ‘exp_client_hello_not_found’: Expected Client HELLO requests not found;
* `match_any` - Match any source
* `match_authorize_policy` - Authorize-policy for user and group based policy
* `match_class_list` - Class List Name
* `priority` - Priority value of the action(higher the number higher the priority)
* `action` - Action to be performed
* `dest_class_list` - Destination Class List Name
* `type` - ‘host’: Match hostname; ‘url’: match URL;
* `web_category_list` - Destination Class List Name
* `client_ip_l3_dest` - Use destination IP as client IP address
* `client_ip_l7_header` - Use extract client IP address from L7 header
* `header_name` - Specify L7 header name
* `action_value` - ‘forward’: Forward the traffic even it exceeds limit; ‘reset’: Reset the connection when it exceeds limit;
* `bw_per` - Per (Specify interval in number of 100ms)
* `bw_rate_limit` - Specify bandwidth rate limit (Bandwidth rate limit in bytes)
* `conn_limit` - Connection limit
* `conn_per` - Per (Specify interval in number of 100ms)
* `conn_rate_limit` - Specify connection rate limit
* `direct_action` - Set action when match the lid
* `direct_action_interval` - Specify logging interval in minute (default is 3)
* `direct_action_value` - ‘drop’: drop the packet; ‘reset’: Send reset back;
* `direct_fail` - Only log unsuccessful connections
* `direct_logging_drp_rst` - Configure PBSLB logging
* `direct_pbslb_interval` - Specify logging interval in minutes(default is 3)
* `direct_pbslb_logging` - Configure PBSLB logging
* `direct_service_group` - Specify a service group (Specify the service group name)
* `lidnum` - Specify a limit ID
* `lockout` - Don’t accept any new connection for certain time (Lockout duration in minutes)
* `over_limit_action` - Set action when exceeds limit
* `request_limit` - Request limit (Specify request limit)
* `request_per` - Per (Specify interval in number of 100ms)
* `request_rate_limit` - Request rate limit (Specify request rate limit)
* `disable` - Disable
* `exclusive_answer` - Exclusive Answer in DNS Response
* `prefix` - IPv6 prefix
* `code_range_end` - server response code range end
* `code_range_start` - server response code range start
* `period` - seconds
* `threshold` - the times of getting the response code
* `action_interval` - Specify logging interval in minute (default is 3)
* `bw_list_action` - ‘drop’: drop the packet; ‘reset’: Send reset back;
* `fail` - Only log unsuccessful connections
* `id` - Specify id that maps to service group (The id number)
* `logging_drp_rst` - Configure PBSLB logging
* `pbslb_interval` - Specify logging interval in minutes
* `pbslb_logging` - Configure PBSLB logging
* `service_group` - Specify a service group (Specify the service group name)
