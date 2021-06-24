---
layout: "thunder"
page_title: "thunder: thunder_slb_common"
sidebar_current: "docs-thunder-resource-slb-common"
description: |-
	Provides details about thunder slb common resource for A10
---

# thunder\_slb\_common

`thunder_slb_common` Provides details about thunder slb common
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_common" "Slb_Common_Test" {
low_latency = 0
QAT = 0
use_mss_tab = 0
stats_data_disable = 0
compress_block_size = 0
player_id_check_enable = 0
after_disable = 0
msl_time = 0
graceful_shutdown_enable = 0
buff_thresh_hw_buff = 0
hw_syn_rr = 0
entity = "string"
reset_stale_session = 0
N5_old = 0
gateway_health_check = 0
scale_out = 0
graceful_shutdown = 0
rate_limit_logging = 0
fast_path_disable = 0
drop_icmp_to_vip_when_vip_down = 0
ssli_sni_hash_enable = 0
hw_compression = 0
dns_vip_stateless = 0
buff_thresh_sys_buff_low = 0
range_end = 0
dns_response_rate_limiting {  
        uuid =  "string" 
        max_table_entries =  0 
        }
dns_cache_enable = 0
max_local_rate = 0
exclude_destination = "string"
dns_cache_age = 0
max_http_header_count = 0
l2l3_trunk_lb_disable = 0
resolve_port_conflict = 0
sort_res = 0
snat_gwy_for_l3 = 0
buff_thresh_relieve_thresh = 0
N5_new = 0
dsr_health_check_enable = 0
buff_thresh = 0
dns_cache_entry_size = 0
log_for_reset_unknown_conn = 0
auto_nat_no_ip_refresh = "string"
software_tls13 = 0
pkt_rate_for_reset_unknown_conn = 0
buff_thresh_sys_buff_high = 0
max_buff_queued_per_conn = 0
max_remote_rate = 0
ttl_threshold = 0
extended_stats = 0
enable_l7_req_acct = 0
uuid = "string"
snat_on_vip = 0
substitute_source_mac = 0
range_start = 0
honor_server_response_ttl = 0
interval = 0
stateless_sg_multi_binding = 0
disable_adaptive_resource_check = 0
range = 0
conn_rate_limit {  
 src-ip-list {   
        protocol =  "string" 
        log =  0 
        lock_out =  0 
        limit_period =  "string" 
        limit =  0 
        exceed_action =  0 
        shared =  0 
        uuid =  "string" 
        }
        }
mss_table = 0
timeout = 0
response_type = "string"
ddos_protection {  
 packets_per_second {  
        ipd_tcp =  0 
        ipd_udp =  0 
        }
logging {  
        ipd_logging_toggle =  "string" 
        }
        ipd_enable_toggle =  "string" 
        }
override_port = 0
no_auto_up_on_aflex = 0
disable_server_auto_reselect = 0
software = 0
 
}
```

## Argument Reference

* `after_disable` - Graceful shutdown after disable server/port and/or virtual server/port
* `auto_nat_no_ip_refresh` - ‘enable’: enable; ‘disable’: disable;
* `buff_thresh` - Set buffer threshold
* `buff_thresh_hw_buff` - Set hardware buffer threshold
* `buff_thresh_relieve_thresh` - Relieve threshold
* `buff_thresh_sys_buff_high` - Set high water mark of system buffer
* `buff_thresh_sys_buff_low` - Set low water mark of system buffer
* `compress_block_size` - Set compression block size (Compression block size in bytes)
* `disable_adaptive_resource_check` - Disable adaptive resource check based on buffer usage
* `disable_server_auto_reselect` - Disable auto reselection of server
* `dns_cache_age` - Set DNS cache entry age, default is 300 seconds (1-1000000 seconds, default is 300 seconds)
* `dns_cache_enable` - Enable DNS cache
* `dns_cache_entry_size` - Set DNS cache entry size, default is 256 bytes (1-4096 bytes, default is 256 bytes)
* `dns_vip_stateless` - Enable DNS VIP stateless mode
* `drop_icmp_to_vip_when_vip_down` - Drop ICMP to VIP when VIP down
* `dsr_health_check_enable` - Enable dsr-health-check (direct server return health check)
* `enable_l7_req_acct` - Enable L7 request accounting
* `entity` - ‘server’: Graceful shutdown server/port only; ‘virtual-server’: Graceful shutdown virtual server/port only;
* `exclude_destination` - ‘local’: Maximum local rate; ‘remote’: Maximum remote rate;  (Maximum rates)
* `extended_stats` - Enable global slb extended statistics
* `fast_path_disable` - Disable fast path in SLB processing
* `gateway_health_check` - Enable gateway health check
* `graceful_shutdown` - 1-65535, in unit of seconds
* `graceful_shutdown_enable` - Enable graceful shutdown
* `honor_server_response_ttl` - Honor the server reponse TTL
* `hw_compression` - Use hardware compression
* `hw_syn_rr` - Configure hardware SYN round robin (range 1-500000)
* `interval` - Specify the healthcheck interval, default is 5 seconds (Interval Value, in seconds (default 5))
* `l2l3_trunk_lb_disable` - Disable L2/L3 trunk LB
* `log_for_reset_unknown_conn` - Log when rate exceed
* `low_latency` - Enable low latency mode
* `max_buff_queued_per_conn` - Set per connection buffer threshold (Buffer value range 128-4096)
* `max_http_header_count` - Set maximum number of HTTP headers allowed
* `max_local_rate` - Set maximum local rate
* `max_remote_rate` - Set maximum remote rate
* `msl_time` - Configure maximum session life, default is 2 seconds (1-40 seconds, default is 2 seconds)
* `mss_table` - Set MSS table (128-750, default is 536)
* `no_auto_up_on_aflex` - Don’t automatically mark vport up when aFleX is bound
* `override_port` - Enable override port in DSR health check mode
* `player_id_check_enable` - Enable the Player id check
* `range` - auto translate port range
* `range_end` - port range end
* `range_start` - port range start
* `rate_limit_logging` - Configure rate limit logging
* `reset_stale_session` - Send reset if session in delete queue receives a SYN packet
* `resolve_port_conflict` - Enable client port service port conflicts
* `response_type` - ‘single-answer’: Only cache DNS response with single answer; ‘round-robin’: Round robin;
* `scale_out` - Enable SLB scale out
* `snat_gwy_for_l3` - Use source NAT gateway for L3 traffic
* `snat_on_vip` - Enable source NAT traffic against VIP
* `software` - Software
* `sort_res` - Enable SLB sorting of resource names
* `ssli_sni_hash_enable` - Enable SSLi SNI hash table
* `stateless_sg_multi_binding` - Enable stateless service groups to be assigned to multiple L2/L3 DSR VIPs
* `stats_data_disable` - Disable global slb data statistics
* `timeout` - Specify the healthcheck timeout value, default is 15 seconds (Timeout Value, in seconds (default 15))
* `ttl_threshold` - Only cache DNS response with longer TTL
* `use_mss_tab` - Use MSS based on internal table for SLB processing
* `uuid` - uuid of the object
* `max_table_entries` - Maximum number of entries allowed
* `ipd_enable_toggle` - ‘enable’: Enable SLB DDoS protection; ‘disable’: Disable SLB DDoS protection (default);
* `ipd_tcp` - Configure packets-per-second threshold per TCP port (default: 200)
* `ipd_udp` - Configure packets-per-second threshold per UDP port (default: 200)
* `ipd_logging_toggle` - ‘enable’: Enable SLB DDoS protection logging (default); ‘disable’: Disable SLB DDoS protection logging;
* `exceed_action` - Set action if threshold exceeded
* `limit` - Set max connections per period
* `limit_period` - ‘100’: 100 ms; ‘1000’: 1000 ms;
* `lock_out` - Set lockout period in seconds if threshold exceeded
* `log` - Send log if threshold exceeded
* `protocol` - ‘tcp’: Set TCP connection rate limit; ‘udp’: Set UDP packet rate limit;
* `shared` - Set threshold shared amongst all virtual ports
