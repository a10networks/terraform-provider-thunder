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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_dns" "resourceSlbTemplateDnsTest" {
	name = "string"
default_policy = "string"
cache_record_serving_policy = "string"
remove_aa_flag = 0
disable_dns_template = 0
period = 0
drop = 0
forward = "string"
max_query_length = 0
max_cache_entry_size = 0
max_cache_size = 0
enable_cache_sharing = 0
remove_padding_to_server = 0
add_padding_to_client = "string"
remove_edns_csubnet_to_server = 0
insert_ipv4 = 0
insert_ipv6 = 0
redirect_to_tcp_port = 0
query_id_switch = 0
dnssec_service_group = "string"
disable_rpz_attach_soa = 0
dns_logging = "string"
uuid = "string"
user_tag = "string"
dns64 {  
 	enable =  0 
	cache =  0 
	change_query =  0 
	parallel_query =  0 
	retry =  0 
	single_response_disable =  0 
	timeout =  0 
	uuid =  "string" 
	}
udp_retransmit {  
 	retry_interval =  0 
	max_trials =  0 
	uuid =  "string" 
	}
query_type_filter {  
 	query_type_action =  "string" 
query-type {   
	str_query_type =  "string" 
	num_query_type =  0 
	order =  "string" 
	}
	uuid =  "string" 
	}
query_class_filter {  
 	query_class_action =  "string" 
query-class {   
	str_query_class =  "string" 
	num_query_class =  0 
	}
	uuid =  "string" 
	}
rpz-list {   
	seq_id =  0 
	name =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
logging {  
 	enable =  0 
rpz-action {   
	str_rpz_action =  "string" 
	}
	uuid =  "string" 
	}
	}
class_list {  
 	name =  "string" 
	uuid =  "string" 
lid-list {   
	lidnum =  0 
	lid_response_rate =  0 
	lid_slip_rate =  0 
	lid_window =  0 
	lid_enable_log =  0 
	lid_action =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
	}
	}
response_rate_limiting {  
 	response_rate =  0 
	filter_response_rate =  0 
	slip_rate =  0 
	window =  0 
	enable_log =  0 
	action =  "string" 
	uuid =  "string" 
rrl-class-list-list {   
	name =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
lid-list {   
	lidnum =  0 
	lid_response_rate =  0 
	lid_slip_rate =  0 
	lid_window =  0 
	lid_enable_log =  0 
	lid_action =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
	}
	}
	}
local_dns_resolution {  
 host-list-cfg {   
	hostnames =  "string" 
	}
local-resolver-cfg {   
	local_resolver =  "string" 
	}
	uuid =  "string" 
	}
recursive_dns_resolution {  
 host-list-cfg {   
	hostnames =  "string" 
	}
	csubnet_retry =  0 
	ipv4_nat_pool =  "string" 
	ipv6_nat_pool =  "string" 
	retries_per_level =  0 
	minimal_response =  0 
	max_trials =  0 
	default_recursive =  0 
	uuid =  "string" 
lookup_order {  
 query-type {   
	str_query_type =  "string" 
	num_query_type =  0 
	order =  "string" 
	}
	uuid =  "string" 
	}
	}
 
}

```

## Argument Reference

* `dns` - DNS template
* `name` - Class-list name
* `default-policy` - 'nocache': Cache disable; 'cache': Cache enable;
* `cache-record-serving-policy` - 'global': Follow global cofiguration (Default); 'no-change': No change in record order; 'round-robin': Round-robin;
* `remove-aa-flag` - Make answers created from cache non-authoritative
* `disable-dns-template` - Disable DNS template
* `period` - Period in minutes
* `drop` - Drop the malformed query
* `forward` - Forward to service group (Service group name)
* `max-query-length` - Define Maximum DNS Query Length, default is unlimited (Specify Maximum Length)
* `max-cache-entry-size` - Define maximum cache entry size (Maximum cache entry size per VIP (default 1024))
* `max-cache-size` - Define maximum cache size (Maximum cache entry per VIP)
* `enable-cache-sharing` - Enable DNS cache sharing
* `remove-padding-to-server` - Remove EDNS(0) padding to server
* `add-padding-to-client` - 'block-length': Block-Length Padding; 'random-block-length': Random-Block-Length Padding;
* `remove-edns-csubnet-to-server` - Remove EDNS(0) client subnet from client queries
* `insert-ipv4` - prefix-length to insert for IPv4
* `insert-ipv6` - prefix-length to insert for IPv6
* `redirect-to-tcp-port` - Direct the client to retry with TCP for DNS UDP request
* `query-id-switch` - Use DNS query ID to create sesion
* `dnssec-service-group` - Use different service group if DNSSEC DO bit set (Service Group Name)
* `disable-rpz-attach-soa` - Disable attaching SOA due to RPZ
* `dns-logging` - dns logging template (DNS Logging template name)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `enable` - Log RPZ triggered action
* `cache` - Use a cached A-query response to provide AAAA query responses for the same hostname
* `change-query` - Always change incoming AAAA DNS Query to A
* `parallel-query` - Forward AAAA Query & generate A Query in parallel
* `retry` - Retry count, default is 3 (Retry Number)
* `single-response-disable` - Disable Single Response which is used to avoid ambiguity
* `timeout` - Timeout to send additional Queries, unit: second, default is 1
* `retry-interval` - DNS Retry Interval value 1 - 400 in units of 100ms, default is 10 (default is 1000ms) (1 - 400 in units of 100ms, default is 10 (1000ms/1sec))
* `max-trials` - Total number of times to try DNS query to server before closing client connection, default 30
* `query-type-action` - 'allow': Allow only certain DNS query types; 'deny': Deny only certain DNS query types;
* `str-query-type` - 'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record; 'ANY': All cached record;
* `num-query-type` - Other query type value
* `order` - 'ipv4-precede-ipv6': Recursive lookup via IPv4 then IPv6; 'ipv6-precede-ipv4': Recursive lookup via IPv6 then IPv4;
* `query-class-action` - 'allow': Allow only certain DNS query classes; 'deny': Deny only certain DNS query classes;
* `str-query-class` - 'INTERNET': INTERNET query class; 'CHAOS': CHAOS query class; 'HESIOD': HESIOD query class; 'NONE': NONE query class; 'ANY': ANY query class;
* `num-query-class` - Other query class value
* `seq-id` - sequential id of RPZ
* `str-rpz-action` - 'drop': Log RPZ due to drop action; 'pass-thru': Log RPZ due to pass-thru action; 'nxdomain': Log RPZ due to nxdomain action; 'nodata': Log RPZ due to nodata action; 'tcp-only': Log RPZ due to tcp-only action; 'local-data': Log RPZ due to local-data action;
* `lidnum` - Specify a limit ID
* `lid-response-rate` - Responses exceeding this rate within the window will be dropped (default 5 per second)
* `lid-slip-rate` - Every n'th response that would be rate-limited will be let through instead
* `lid-window` - Rate-Limiting Interval in Seconds (default is one)
* `lid-enable-log` - Enable logging
* `lid-action` - 'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;
* `response-rate` - Responses exceeding this rate within the window will be dropped (default 5 per second)
* `filter-response-rate` - Maximum allowed request rate for the filter. This should match average traffic. (default 10 per seconds)
* `slip-rate` - Every n'th response that would be rate-limited will be let through instead
* `window` - Rate-Limiting Interval in Seconds (default is one)
* `enable-log` - Enable logging
* `action` - 'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;
* `hostnames` - Hostnames class-list name (ac type), perform resolution while query name matched
* `local-resolver` - Local dns servers (address)
* `csubnet-retry` - retry when server REFUSED AX inserted EDNS(0) subnet, works only when insert-client-subnet is configured
* `ipv4-nat-pool` - IPv4 Source NAT pool or pool group
* `ipv6-nat-pool` - IPv6 Source NAT pool or pool group
* `retries-per-level` - Number of DNS query retries at each server level before closing client connection, default 6
* `minimal-response` - Remove authority and additional records when applicable
* `default-recursive` - Default recurisve mode, forward query to bound service-group if hostnames matched

