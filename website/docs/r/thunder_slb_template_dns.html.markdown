---
layout: "thunder"
page_title: "thunder: thunder_slb_template_dns"
sidebar_current: "docs-thunder-resource-slb-template-dns"
description: |-
	Provides details about thunder slb template dns resource for A10
---

# thunder\_slb\_template\_dns

`thunder_slb_template_dns` Provides details about thunder slb template dns
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_dns" "Slb_Template_Dns_Test" {

dnssec_service_group = "string"
remove_aa_flag = 0
name = "string"
class_list {  
 lid-list {   
        lid_slip_rate =  0 
        lid_action =  "string" 
        uuid =  "string" 
        lid_enable_log =  0 
        lidnum =  0 
        user_tag =  "string" 
        lid_response_rate =  0 
        lid_window =  0 
        }
        name =  "string" 
        uuid =  "string" 
        }
response_rate_limiting {  
 rrl-class-list-list {   
lid-list {   
        lid_slip_rate =  0 
        lid_action =  "string" 
        uuid =  "string" 
        lid_enable_log =  0 
        lidnum =  0 
        user_tag =  "string" 
        lid_response_rate =  0 
        lid_window =  0 
        }
        name =  "string" 
        user_tag =  "string" 
        uuid =  "string" 
        }
        filter_response_rate =  0 
        slip_rate =  0 
        response_rate =  0 
        window =  0 
        action =  "string" 
        enable_log =  0 
        uuid =  "string" 
        }
drop = 0
period = 0
dns_logging = "string"
query_id_switch = 0
enable_cache_sharing = 0
redirect_to_tcp_port = 0
user_tag = "string"
max_query_length = 0
disable_dns_template = 0
forward = "string"
max_cache_size = 0
local_dns_resolution {  
 host-list-cfg {   
        hostnames =  "string" 
        }
        uuid =  "string" 
local-resolver-cfg {   
        local_resolver =  "string" 
        }
        }
default_policy = "string"
max_cache_entry_size = 0
uuid = "string"
 
}
```

## Argument Reference

* `default_policy` - ‘nocache’: Cache disable; ‘cache’: Cache enable;
* `disable_dns_template` - Disable DNS template
* `dnssec_service_group` - Use different service group if DNSSEC DO bit set (Service Group Name)
* `drop` - Drop the malformed query
* `enable_cache_sharing` - Enable DNS cache sharing
* `forward` - Forward to service group (Service group name)
* `max_cache_entry_size` - Define maximum cache entry size (Maximum cache entry size per VIP)
* `max_cache_size` - Define maximum cache size (Maximum cache entry per VIP)
* `max_query_length` - Define Maximum DNS Query Length, default is unlimited (Specify Maximum Length)
* `name` - Specify a class list name
* `period` - Period in minutes
* `query_id_switch` - Use DNS query ID to create sesion
* `redirect_to_tcp_port` - Direct the client to retry with TCP for DNS UDP request
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `action_value` - ‘dns-cache-disable’: Disable DNS cache when it exceeds limit; ‘dns-cache-enable’: Enable DNS cache when it exceeds limit; ‘forward’: Forward the traffic even it exceeds limit;
* `conn_rate_limit` - Connection rate limit
* `lidnum` - Specify a limit ID
* `lockout` - Don’t accept any new connection for certain time (Lockout duration in minutes)
* `log` - Log a message
* `log_interval` - Log interval (minute, by default system will log every over limit instance)
* `over_limit_action` - Action when exceeds limit
* `per` - Per (Number of 100ms)
* `cache_action` - ‘cache-disable’: Disable dns cache; ‘cache-enable’: Enable dns cache;
* `honor_server_response_ttl` - Honor the server reponse TTL
* `ttl` - TTL for cache entry (TTL in seconds)
* `weight` - Weight for cache entry
* `action` - ‘log-only’: Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; ‘rate-limit’: Rate-Limit based on configuration (Default); ‘whitelist’: Whitelist, disable rate-limiting;
* `enable_log` - Enable logging
* `filter_response_rate` - Maximum allowed request rate for the filter. This should match average traffic. (default 20 per two seconds)
* `response_rate` - Responses exceeding this rate within the window will be dropped (default 5 per second)
* `slip_rate` - Every n’th response that would be rate-limited will be let through instead
* `window` - Rate-Limiting Interval in Seconds (default is one)
