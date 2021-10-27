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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_common" "resourceSlbCommonTest" {
	port_scan_detection = "string"
ping_sweep_detection = "string"
extended_stats = 0
stats_data_disable = 0
graceful_shutdown_enable = 0
graceful_shutdown = 0
entity = "string"
after_disable = 0
rate_limit_logging = 0
max_local_rate = 0
max_remote_rate = 0
exclude_destination = "string"
auto_translate_port = 0
range = 0
range_start = 0
range_end = 0
use_default_sess_count = 0
per_thr_percent = 0
dsr_health_check_enable = 0
one_server_conn_hm_rate = 0
aflex_table_entry_aging_interval = 0
override_port = 0
health_check_to_all_vip = 0
reset_stale_session = 0
dns_cache_enable = 0
response_type = "string"
ttl_threshold = 0
dns_cache_age = 0
compress_block_size = 0
dns_cache_entry_size = 0
dns_vip_stateless = 0
honor_server_response_ttl = 0
buff_thresh = 0
buff_thresh_hw_buff = 0
buff_thresh_relieve_thresh = 0
buff_thresh_sys_buff_low = 0
buff_thresh_sys_buff_high = 0
max_buff_queued_per_conn = 0
pkt_rate_for_reset_unknown_conn = 0
log_for_reset_unknown_conn = 0
gateway_health_check = 0
interval = 0
timeout = 0
msl_time = 0
fast_path_disable = 0
http_fast_enable = 0
l2l3_trunk_lb_disable = 0
snat_gwy_for_l3 = 0
allow_in_gateway_mode = 0
disable_server_auto_reselect = 0
enable_l7_req_acct = 0
disable_adaptive_resource_check = 0
ddos_pkt_size_thresh = 0
ddos_pkt_count_thresh = 0
snat_on_vip = 0
low_latency = 0
mss_table = 0
resolve_port_conflict = 0
no_auto_up_on_aflex = 0
hw_compression = 0
hw_syn_rr = 0
max_http_header_count = 0
scale_out = 0
scale_out_traffic_map = 0
show_slb_server_legacy_cmd = 0
show_slb_service_group_legacy_cmd = 0
show_slb_virtual_server_legacy_cmd = 0
traffic_map_type = "string"
sort_res = 0
use_mss_tab = 0
auto_nat_no_ip_refresh = "string"
ddos_protection {  
 	ipd_enable_toggle =  "string" 
logging {  
 	ipd_logging_toggle =  "string" 
	}
packets_per_second {  
 	ipd_tcp =  0 
	ipd_udp =  0 
	}
	}
ssli_sni_hash_enable = 0
software = 0
software_tls13 = 0
qat = 0
n5_new = 0
n5_old = 0
software_tls13_offload = 0
substitute_source_mac = 0
drop_icmp_to_vip_when_vip_down = 0
player_id_check_enable = 0
stateless_sg_multi_binding = 0
ecmp_hash = "string"
service_group_on_no_dest_nat_vports = "string"
disable_port_masking = 0
snat_preserve {  
 range {   
	port1 =  0 
	port2 =  0 
	}
	}
disable_persist_scoring = 0
ipv4_offset = 0
pre_process_enable = 0
attack_resp_code = 0
uuid = "string"
cert_pinning {  
 	ttl =  0 
	uuid =  "string" 
candidate_list_feedback_opt_in {  
 	enable =  0 
	schedule =  0 
	weekly =  0 
	week_day =  "string" 
	week_time =  "string" 
	daily =  0 
	day_time =  "string" 
	use_mgmt_port =  0 
	uuid =  "string" 
	}
	}
aflex_table_entry_sync {  
 	aflex_table_entry_sync_enable =  0 
	aflex_table_entry_sync_max_key_len =  0 
	aflex_table_entry_sync_max_value_len =  0 
	aflex_table_entry_sync_min_lifetime =  0 
	uuid =  "string" 
	}
conn_rate_limit {  
 src-ip-list {   
	protocol =  "string" 
	limit =  0 
	limit_period =  "string" 
	shared =  0 
	exceed_action =  0 
	log =  0 
	lock_out =  0 
	uuid =  "string" 
	}
	}
dns_response_rate_limiting {  
 	max_table_entries =  0 
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `common` - SLB related commands
* `port-scan-detection` - 'enable': Enable port scan detection; 'disable': Disable port scan detection(default);
* `ping-sweep-detection` - 'enable': Enable ping sweep detection; 'disable': Disable ping sweep detection(default);
* `extended-stats` - Enable global slb extended statistics
* `stats-data-disable` - Disable global slb data statistics
* `graceful-shutdown-enable` - Enable graceful shutdown
* `graceful-shutdown` - 1-65535, in unit of seconds
* `entity` - 'server': Graceful shutdown server/port only; 'virtual-server': Graceful shutdown virtual server/port only;
* `after-disable` - Graceful shutdown after disable server/port and/or virtual server/port
* `rate-limit-logging` - Configure rate limit logging
* `max-local-rate` - Set maximum local rate
* `max-remote-rate` - Set maximum remote rate
* `exclude-destination` - 'local': Maximum local rate; 'remote': Maximum remote rate;  (Maximum rates)
* `auto-translate-port` - Auto Translate Port range
* `range` - auto translate port range
* `range-start` - port range start
* `range-end` - port range end
* `use-default-sess-count` - Use default session count
* `per-thr-percent` - Percentage of default session count to use for per thread session table size
* `dsr-health-check-enable` - Enable dsr-health-check (direct server return health check)
* `one-server-conn-hm-rate` - One Server Conn Health Check Rate
* `aflex-table-entry-aging-interval` - aFleX table entry aging interval in second
* `override-port` - Enable override port in DSR health check mode
* `reset-stale-session` - Send reset if session in delete queue receives a SYN packet
* `dns-cache-enable` - Enable DNS cache
* `response-type` - 'single-answer': Only cache DNS response with single answer; 'round-robin': Round robin;
* `ttl-threshold` - Only cache DNS response with longer TTL
* `dns-cache-age` - Set DNS cache entry age, default is 300 seconds (1-1000000 seconds, default is 300 seconds)
* `compress-block-size` - Set compression block size (Compression block size in bytes)
* `dns-cache-entry-size` - Set DNS cache entry size, default is 256 bytes (1-4096 bytes, default is 256 bytes)
* `dns-vip-stateless` - Enable DNS VIP stateless mode
* `honor-server-response-ttl` - Honor the server reponse TTL
* `buff-thresh` - Set buffer threshold
* `buff-thresh-hw-buff` - Set hardware buffer threshold
* `buff-thresh-relieve-thresh` - Relieve threshold
* `buff-thresh-sys-buff-low` - Set low water mark of system buffer
* `buff-thresh-sys-buff-high` - Set high water mark of system buffer
* `max-buff-queued-per-conn` - Set per connection buffer threshold (Buffer value range 128-4096)
* `log-for-reset-unknown-conn` - Log when rate exceed
* `gateway-health-check` - Enable gateway health check
* `interval` - Specify the healthcheck interval, default is 5 seconds (Interval Value, in seconds (default 5))
* `timeout` - Specify the healthcheck timeout value, default is 15 seconds (Timeout Value, in seconds (default 15))
* `msl-time` - Configure maximum session life, default is 2 seconds (1-40 seconds, default is 2 seconds)
* `fast-path-disable` - Disable fast path in SLB processing
* `http-fast-enable` - Enable Http Fast in SLB processing
* `l2l3-trunk-lb-disable` - Disable L2/L3 trunk LB
* `snat-gwy-for-l3` - Use source NAT gateway for L3 traffic for transparent mode
* `allow-in-gateway-mode` - Use source NAT gateway for L3 traffic for gateway mode
* `disable-server-auto-reselect` - Disable auto reselection of server
* `enable-l7-req-acct` - Enable L7 request accounting
* `disable-adaptive-resource-check` - Disable adaptive resource check based on buffer usage
* `ddos-pkt-size-thresh` - Set data packet size threshold for DDOS, default is 64 bytes
* `ddos-pkt-count-thresh` - Set packet count threshold for DDOS, default is 100
* `snat-on-vip` - Enable source NAT traffic against VIP
* `low-latency` - Enable low latency mode
* `mss-table` - Set MSS table (128-750, default is 536)
* `resolve-port-conflict` - Enable client port service port conflicts
* `no-auto-up-on-aflex` - Don't automatically mark vport up when aFleX is bound
* `hw-compression` - Use hardware compression
* `hw-syn-rr` - Configure hardware SYN round robin (range 1-500000)
* `max-http-header-count` - Set maximum number of HTTP headers allowed
* `scale-out` - Enable SLB scale out
* `scale-out-traffic-map` - Set SLB scaleout traffic-map
* `show-slb-server-legacy-cmd` - Enable show slb server legacy command
* `show-slb-service-group-legacy-cmd` - Enable show slb service-group legacy command
* `show-slb-virtual-server-legacy-cmd` - Enable show slb virtual-server legacy command
* `traffic-map-type` - 'vport': traffic-map per vport; 'global': global traffic-map;
* `sort-res` - Enable SLB sorting of resource names
* `use-mss-tab` - Use MSS based on internal table for SLB processing
* `auto-nat-no-ip-refresh` - 'enable': enable; 'disable': disable;
* `ipd-enable-toggle` - 'enable': Enable SLB DDoS protection; 'disable': Disable SLB DDoS protection (default);
* `ipd-logging-toggle` - 'enable': Enable SLB DDoS protection logging (default); 'disable': Disable SLB DDoS protection logging;
* `ipd-tcp` - Configure packets-per-second threshold per TCP port (default: 200)
* `ipd-udp` - Configure packets-per-second threshold per UDP port (default: 200)
* `ssli-sni-hash-enable` - Enable SSLi SNI hash table
* `software` - Software
* `software-tls13` - Software TLS1.3
* `qat` - HW assisted QAT SSL module
* `n5-new` - HW assisted N5 SSL module with TLS 1.3 and TLS 1.2 support using OpenSSL 1.1.1
* `n5-old` - HW assisted N5 SSL module with TLS 1.2 support using OpenSSL 0.9.7
* `software-tls13-offload` - Software TLS1.3 with CPU Offload Support
* `substitute-source-mac` - Substitute Source MAC Address to that of the outgoing interface
* `drop-icmp-to-vip-when-vip-down` - Drop ICMP to VIP when VIP down
* `player-id-check-enable` - Enable the Player id check
* `stateless-sg-multi-binding` - Enable stateless service groups to be assigned to multiple L2/L3 DSR VIPs
* `ecmp-hash` - 'system-default': Use system default ecmp hashing algorithm; 'connection-based': Use connection information for hashing;
* `service-group-on-no-dest-nat-vports` - 'allow-same': Allow the binding service-group on no-dest-nat virtual ports; 'enforce-different': Enforce that the same service-group can not be bound on different no-dest-nat virtual ports;
* `disable-port-masking` - Disable masking of ports for CPU hashing
* `port1` - start port
* `port2` - end port which is greater than start
* `disable-persist-scoring` - Disable Persist Scoring
* `ipv4-offset` - IPv4 Octet Offset for Hash
* `pre-process-enable` - Enable NG-WAF pre-processing
* `attack-resp-code` - Custom response
* `uuid` - uuid of the object
* `ttl` - The ttl of local cert pinning candidate list, multiple of 10 minutes, default is 144 (1440 minutes)
* `enable` - Enable the feedback function
* `schedule` - schedule the uploading time, default is daily 00:00
* `weekly` - Every week
* `week-day` - 'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;
* `week-time` - Time of day to update (hh:mm) in 24 hour local time
* `daily` - Every day
* `day-time` - Time of day to update (hh:mm) in 24 hour local time
* `use-mgmt-port` - Use management port to connect
* `aflex-table-entry-sync-enable` - Enable aflex table sync
* `aflex-table-entry-sync-max-key-len` - aflex table entry max key length to sync
* `aflex-table-entry-sync-max-value-len` - aflex table entry max value length to sync
* `aflex-table-entry-sync-min-lifetime` - aflex table entry minimum lifetime to sync
* `protocol` - 'tcp': Set TCP connection rate limit; 'udp': Set UDP packet rate limit;
* `limit` - Set max connections per period
* `limit-period` - '100': 100 ms; '1000': 1000 ms;
* `shared` - Set threshold shared amongst all virtual ports
* `exceed-action` - Set action if threshold exceeded
* `log` - Send log if threshold exceeded
* `lock-out` - Set lockout period in seconds if threshold exceeded
* `max-table-entries` - Maximum number of entries allowed

