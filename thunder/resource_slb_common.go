package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommon() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common`: SLB related commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonCreate,
		UpdateContext: resourceSlbCommonUpdate,
		ReadContext:   resourceSlbCommonRead,
		DeleteContext: resourceSlbCommonDelete,

		Schema: map[string]*schema.Schema{
			"aflex_table_entry_aging_interval": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "aFleX table entry aging interval in second",
			},
			"aflex_table_entry_sync": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aflex_table_entry_sync_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable aflex table sync",
						},
						"aflex_table_entry_sync_max_key_len": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "aflex table entry max key length to sync",
						},
						"aflex_table_entry_sync_max_value_len": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "aflex table entry max value length to sync",
						},
						"aflex_table_entry_sync_min_lifetime": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "aflex table entry minimum lifetime to sync",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"after_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Graceful shutdown after disable server/port and/or virtual server/port",
			},
			"allow_in_gateway_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use source NAT gateway for L3 traffic for gateway mode",
			},
			"attack_resp_code": {
				Type: schema.TypeInt, Optional: true, Default: 410, Description: "Custom response code",
			},
			"auto_nat_no_ip_refresh": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"auto_translate_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto Translate Port range",
			},
			"buff_thresh": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set buffer threshold",
			},
			"buff_thresh_hw_buff": {
				Type: schema.TypeInt, Optional: true, Description: "Set hardware buffer threshold",
			},
			"buff_thresh_relieve_thresh": {
				Type: schema.TypeInt, Optional: true, Description: "Relieve threshold",
			},
			"buff_thresh_sys_buff_high": {
				Type: schema.TypeInt, Optional: true, Description: "Set high water mark of system buffer",
			},
			"buff_thresh_sys_buff_low": {
				Type: schema.TypeInt, Optional: true, Description: "Set low water mark of system buffer",
			},
			"cache_expire_time": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Cache expiration time, default is 1 minute",
			},
			"cancel_stream_loop_limit": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set global cancel stream loop limit (cancel stream loop limit, default is 5)",
			},
			"cert_pinning": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ttl": {
							Type: schema.TypeInt, Optional: true, Default: 144, Description: "The ttl of local cert pinning candidate list, multiple of 10 minutes, default is 144 (1440 minutes)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"candidate_list_feedback_opt_in": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the feedback function",
									},
									"schedule": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "schedule the uploading time, default is daily 00:00",
									},
									"weekly": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every week",
									},
									"week_day": {
										Type: schema.TypeString, Optional: true, Description: "'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;",
									},
									"week_time": {
										Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
									},
									"daily": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every day",
									},
									"day_time": {
										Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
									},
									"use_mgmt_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to connect",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"clientside_ip": {
				Type: schema.TypeString, Optional: true, Description: "Clientside IP address",
			},
			"clientside_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Clientside IPv6 address",
			},
			"compress_block_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set compression block size (Compression block size in bytes)",
			},
			"conn_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disable_ipv6_support": {
										Type: schema.TypeInt, Required: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': Set TCP connection rate limit; 'udp': Set UDP packet rate limit;",
									},
									"limit": {
										Type: schema.TypeInt, Optional: true, Description: "Set max connections per period",
									},
									"limit_period": {
										Type: schema.TypeString, Optional: true, Description: "'100': 100 ms; '1000': 1000 ms;",
									},
									"shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set threshold shared amongst all virtual ports",
									},
									"exceed_action": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action if threshold exceeded",
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send log if threshold exceeded",
									},
									"lock_out": {
										Type: schema.TypeInt, Optional: true, Description: "Set lockout period in seconds if threshold exceeded",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"custom_message": {
				Type: schema.TypeString, Optional: true, Description: "Block message",
			},
			"custom_page": {
				Type: schema.TypeString, Optional: true, Description: "Specify the custom webpage name",
			},
			"custom_signal_clist": {
				Type: schema.TypeString, Optional: true, Description: "Provide custom signal names",
			},
			"ddos_pkt_count_thresh": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Set packet count threshold for DDOS, default is 100",
			},
			"ddos_pkt_size_thresh": {
				Type: schema.TypeInt, Optional: true, Default: 64, Description: "Set data packet size threshold for DDOS, default is 64 bytes",
			},
			"ddos_protection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipd_enable_toggle": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable SLB DDoS protection; 'disable': Disable SLB DDoS protection (default);",
						},
						"logging": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipd_logging_toggle": {
										Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable SLB DDoS protection logging (default); 'disable': Disable SLB DDoS protection logging;",
									},
								},
							},
						},
						"packets_per_second": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipd_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 200, Description: "Configure packets-per-second threshold per TCP port (default: 200)",
									},
									"ipd_udp": {
										Type: schema.TypeInt, Optional: true, Default: 200, Description: "Configure packets-per-second threshold per UDP port (default: 200)",
									},
								},
							},
						},
					},
				},
			},
			"disable_adaptive_resource_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable adaptive resource check based on buffer usage",
			},
			"disable_persist_scoring": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Persist Scoring",
			},
			"disable_port_masking": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable masking of ports for CPU hashing",
			},
			"disable_server_auto_reselect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable auto reselection of server",
			},
			"dns_cache_age": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Set DNS cache entry age, default is 300 seconds (1-1000000 seconds, default is 300 seconds)",
			},
			"dns_cache_age_min_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set DNS cache entry age minimum threshold, default is 0 seconds (1-1000000 seconds, default is 0 seconds)",
			},
			"dns_cache_aging_weight": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set DNS cache entry weight, default is 1",
			},
			"dns_cache_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS cache",
			},
			"dns_cache_entry_size": {
				Type: schema.TypeInt, Optional: true, Default: 256, Description: "Set DNS cache entry size, default is 256 bytes (1-4096 bytes, default is 256 bytes)",
			},
			"dns_cache_sync": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS cache HA sync",
			},
			"dns_cache_sync_entry_size": {
				Type: schema.TypeInt, Optional: true, Default: 256, Description: "Only sync DNS cache with smaller size (1-4096 bytes, default is 256 bytes)",
			},
			"dns_cache_sync_ttl_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only sync DNS cache with longer TTL (0-10000000 seconds, default is 0 second)",
			},
			"dns_cache_ttl_adjustment_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS cache response ttl adjustment",
			},
			"dns_negative_cache_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS negative cache",
			},
			"dns_persistent_cache_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable persistent DNS cache",
			},
			"dns_persistent_cache_hit_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only save DNS cache with larger hit count (0-10000000, default is 0)",
			},
			"dns_persistent_cache_ttl_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only save DNS cache with longer TTL (0-10000000 seconds, default is 0 second)",
			},
			"dns_response_rate_limiting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_table_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of entries allowed",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dns_vip_stateless": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS VIP stateless mode",
			},
			"drop_icmp_to_vip_when_vip_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop ICMP to VIP when VIP down",
			},
			"dsr_health_check_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable dsr-health-check (direct server return health check)",
			},
			"ecmp_hash": {
				Type: schema.TypeString, Optional: true, Default: "system-default", Description: "'system-default': Use system default ecmp hashing algorithm; 'connection-based': Use connection information for hashing;",
			},
			"enable_ddos": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DDoS protection",
			},
			"enable_fast_path_rerouting": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Fast-Path Rerouting",
			},
			"enable_l7_req_acct": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable L7 request accounting",
			},
			"entity": {
				Type: schema.TypeString, Optional: true, Description: "'server': Graceful shutdown server/port only; 'virtual-server': Graceful shutdown virtual server/port only;",
			},
			"exclude_destination": {
				Type: schema.TypeString, Optional: true, Description: "'local': Maximum local rate; 'remote': Maximum remote rate;  (Maximum rates)",
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable global slb extended statistics",
			},
			"fast_path_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable fast path in SLB processing",
			},
			"gateway_health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable gateway health check",
			},
			"graceful_shutdown": {
				Type: schema.TypeInt, Optional: true, Description: "1-65535, in unit of seconds",
			},
			"graceful_shutdown_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable graceful shutdown",
			},
			"health_check_to_all_vip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"honor_server_response_ttl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Honor the server reponse TTL",
			},
			"http_fast_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Http Fast in SLB processing",
			},
			"hw_compression": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use hardware compression",
			},
			"hw_syn_rr": {
				Type: schema.TypeInt, Optional: true, Description: "Configure hardware SYN round robin (range 1-500000)",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the healthcheck interval, default is 5 seconds (Interval Value, in seconds (default 5))",
			},
			"ipv4_offset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv4 Octet Offset for Hash",
			},
			"ipv6_subnet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv6 Octet Valid Subnet Length for Hash",
			},
			"l2l3_trunk_lb_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable L2/L3 trunk LB",
			},
			"log_for_reset_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when rate exceed",
			},
			"low_latency": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable low latency mode",
			},
			"max_buff_queued_per_conn": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Set per connection buffer threshold (Buffer value range 128-4096)",
			},
			"max_http_header_count": {
				Type: schema.TypeInt, Optional: true, Default: 90, Description: "Set maximum number of HTTP headers allowed",
			},
			"max_local_rate": {
				Type: schema.TypeInt, Optional: true, Default: 32, Description: "Set maximum local rate",
			},
			"max_persistent_cache": {
				Type: schema.TypeInt, Optional: true, Description: "Define maximum persistent cache (Maximum persistent cache entry)",
			},
			"max_remote_rate": {
				Type: schema.TypeInt, Optional: true, Default: 15000, Description: "Set maximum remote rate",
			},
			"monitor_mode_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable NG-WAF monitor mode",
			},
			"msl_time": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure maximum session life, default is 2 seconds (1-39 seconds, default is 2 seconds)",
			},
			"mss_table": {
				Type: schema.TypeInt, Optional: true, Default: 536, Description: "Set MSS table (128-750, default is 536)",
			},
			"multi_cpu": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specific NGWAF CPU",
			},
			"n5_new": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HW assisted N5 SSL module with TLS 1.3 and TLS 1.2 support using OpenSSL 1.1.1",
			},
			"n5_old": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HW assisted N5 SSL module with TLS 1.2 support using OpenSSL 0.9.7",
			},
			"ngwaf_proxy_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 address",
			},
			"ngwaf_proxy_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 address",
			},
			"ngwaf_proxy_port": {
				Type: schema.TypeInt, Optional: true, Description: "Port",
			},
			"no_auto_up_on_aflex": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't automatically mark vport up when aFleX is bound",
			},
			"odd_even_nat_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable odd even nat pool allocation in dual blade systems",
			},
			"one_server_conn_hm_rate": {
				Type: schema.TypeInt, Optional: true, Description: "One Server Conn Health Check Rate",
			},
			"override_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable override port in DSR health check mode",
			},
			"pbslb_entry_age": {
				Type: schema.TypeInt, Optional: true, Default: 6, Description: "Set global pbslb entry age (minute)",
			},
			"pbslb_overflow_glid": {
				Type: schema.TypeString, Optional: true, Description: "Apply global limit id to overflow pbslb entry",
			},
			"per_thr_percent": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of default session count to use for per thread session table size",
			},
			"ping_sweep_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable ping sweep detection; 'disable': Disable ping sweep detection(default);",
			},
			"pkt_rate_for_reset_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"player_id_check_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Player id check",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Serverside port number for SNI transmission",
			},
			"port_scan_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable port scan detection; 'disable': Disable port scan detection(default);",
			},
			"pre_process_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable NG-WAF pre-processing",
			},
			"qat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HW assisted QAT SSL module",
			},
			"quic": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cid_len": {
							Type: schema.TypeInt, Optional: true, Default: 4, Description: "Length of CID",
						},
						"signature": {
							Type: schema.TypeString, Optional: true, Description: "Set CID Signature",
						},
						"signature_len": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Offset for CID Signature",
						},
						"signature_offset": {
							Type: schema.TypeInt, Optional: true, Default: 4, Description: "Offset for CID Signature",
						},
						"cpu_offset": {
							Type: schema.TypeInt, Optional: true, Default: 12, Description: "Offset for Encoded CPU",
						},
						"quic_lb_offset": {
							Type: schema.TypeInt, Optional: true, Default: 8, Description: "Offset for QUIC-LB",
						},
						"enable_hash": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CID Hashing",
						},
						"enable_signature": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CID Signature Validation",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"range": {
				Type: schema.TypeInt, Optional: true, Description: "auto translate port range",
			},
			"range_end": {
				Type: schema.TypeInt, Optional: true, Description: "port range end",
			},
			"range_start": {
				Type: schema.TypeInt, Optional: true, Description: "port range start",
			},
			"rate_limit_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure rate limit logging",
			},
			"recursive_ns_cache": {
				Type: schema.TypeString, Optional: true, Default: "honor-packet-ttl", Description: "'honor-packet-ttl': Honor the lowest TTL among NS records in the server response; 'honor-age-config': Honor the ttl/age settings based on acos dns cache configuration;",
			},
			"reset_stale_session": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset if session in delete queue receives a SYN packet",
			},
			"resolve_port_conflict": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable client port service port conflicts",
			},
			"response_type": {
				Type: schema.TypeString, Optional: true, Description: "'single-answer': Only cache DNS response with single answer; 'round-robin': Round robin;",
			},
			"scale_out": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB scale out",
			},
			"scale_out_traffic_map": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set SLB scaleout traffic-map",
			},
			"serverside_ip": {
				Type: schema.TypeString, Optional: true, Description: "Serverside IP address",
			},
			"serverside_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Serverside IPv6 address",
			},
			"service_group_on_no_dest_nat_vports": {
				Type: schema.TypeString, Optional: true, Default: "enforce-different", Description: "'allow-same': Allow the binding service-group on no-dest-nat virtual ports; 'enforce-different': Enforce that the same service-group can not be bound on different no-dest-nat virtual ports;",
			},
			"show_slb_server_legacy_cmd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable show slb server legacy command",
			},
			"show_slb_service_group_legacy_cmd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable show slb service-group legacy command",
			},
			"show_slb_virtual_server_legacy_cmd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable show slb virtual-server legacy command",
			},
			"snat_gwy_for_l3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use source NAT gateway for L3 traffic for transparent mode",
			},
			"snat_on_vip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable source NAT traffic against VIP",
			},
			"snat_preserve": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port1": {
										Type: schema.TypeInt, Optional: true, Default: 1025, Description: "start port",
									},
									"port2": {
										Type: schema.TypeInt, Optional: true, Default: 1025, Description: "end port which is greater than start",
									},
								},
							},
						},
					},
				},
			},
			"software": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Software",
			},
			"software_tls13": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Software TLS1.3",
			},
			"software_tls13_offload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Software TLS1.3 with CPU Offload Support",
			},
			"sort_res": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB sorting of resource names",
			},
			"ssl_module_usage_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL module usage calculations for QAT",
			},
			"ssl_n5_delay_tx_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable delay transmission for N5-new",
			},
			"ssl_ratelimit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable HW SSL Rate limit for N5-new",
						},
						"tls12_rate": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "Enabling Rateliming for TLS1.2 HW requests per chip in 1K - default 120",
						},
						"tls13_rate": {
							Type: schema.TypeInt, Optional: true, Default: 72, Description: "Enabling Rateliming for TLS1.3 HW requests per chip in 1K - default 72",
						},
					},
				},
			},
			"ssli_cert_not_ready_inspect_limit": {
				Type: schema.TypeInt, Optional: true, Default: 2000, Description: "SSLI asynchronized connection max number, default is 2000 (set to 0 for unlimited size)",
			},
			"ssli_cert_not_ready_inspect_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "SSLI asynchronized connection timeout, default is 10 seconds (seconds, set to 0 for never timeout)",
			},
			"ssli_silent_termination_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Terminate the SSLi sessions silently without sending RST/FIN packet",
			},
			"ssli_sni_hash_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSLi SNI hash table",
			},
			"stateless_sg_multi_binding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateless service groups to be assigned to multiple L2/L3 DSR VIPs",
			},
			"stats_data_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable global slb data statistics",
			},
			"substitute_source_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Substitute Source MAC Address to that of the outgoing interface",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "Specify the healthcheck timeout value, default is 15 seconds (Timeout Value, in seconds (default 15))",
			},
			"traffic_map_type": {
				Type: schema.TypeString, Optional: true, Default: "vport", Description: "'vport': traffic-map per vport; 'global': global traffic-map;",
			},
			"ttl_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Only cache DNS response with longer TTL",
			},
			"use_default_sess_count": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use default session count",
			},
			"use_https_proxy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NG-WAF connects to Cloud through proxy server",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to connect",
			},
			"use_mss_tab": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use MSS based on internal table for SLB processing",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vport_global": {
				Type: schema.TypeInt, Optional: true, Description: "Configure periodic showtech vport paging global limit",
			},
			"vport_l3v": {
				Type: schema.TypeInt, Optional: true, Description: "Configure periodic showtech vport paging l3v limit",
			},
		},
	}
}
func resourceSlbCommonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommon(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommon(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommon(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommon(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbCommonAflexTableEntrySync1407(d []interface{}) edpt.SlbCommonAflexTableEntrySync1407 {

	count1 := len(d)
	var ret edpt.SlbCommonAflexTableEntrySync1407
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AflexTableEntrySyncEnable = in["aflex_table_entry_sync_enable"].(int)
		ret.AflexTableEntrySyncMaxKeyLen = in["aflex_table_entry_sync_max_key_len"].(int)
		ret.AflexTableEntrySyncMaxValueLen = in["aflex_table_entry_sync_max_value_len"].(int)
		ret.AflexTableEntrySyncMinLifetime = in["aflex_table_entry_sync_min_lifetime"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbCommonCertPinning1408(d []interface{}) edpt.SlbCommonCertPinning1408 {

	count1 := len(d)
	var ret edpt.SlbCommonCertPinning1408
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ttl = in["ttl"].(int)
		//omit uuid
		ret.CandidateListFeedbackOptIn = getObjectSlbCommonCertPinningCandidateListFeedbackOptIn1409(in["candidate_list_feedback_opt_in"].([]interface{}))
	}
	return ret
}

func getObjectSlbCommonCertPinningCandidateListFeedbackOptIn1409(d []interface{}) edpt.SlbCommonCertPinningCandidateListFeedbackOptIn1409 {

	count1 := len(d)
	var ret edpt.SlbCommonCertPinningCandidateListFeedbackOptIn1409
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Schedule = in["schedule"].(int)
		ret.Weekly = in["weekly"].(int)
		ret.WeekDay = in["week_day"].(string)
		ret.WeekTime = in["week_time"].(string)
		ret.Daily = in["daily"].(int)
		ret.DayTime = in["day_time"].(string)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbCommonConnRateLimit1410(d []interface{}) edpt.SlbCommonConnRateLimit1410 {

	count1 := len(d)
	var ret edpt.SlbCommonConnRateLimit1410
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcIpList = getSliceSlbCommonConnRateLimitSrcIpList(in["src_ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbCommonConnRateLimitSrcIpList(d []interface{}) []edpt.SlbCommonConnRateLimitSrcIpList {

	count1 := len(d)
	ret := make([]edpt.SlbCommonConnRateLimitSrcIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbCommonConnRateLimitSrcIpList
		oi.DisableIpv6Support = in["disable_ipv6_support"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Limit = in["limit"].(int)
		oi.LimitPeriod = in["limit_period"].(string)
		oi.Shared = in["shared"].(int)
		oi.ExceedAction = in["exceed_action"].(int)
		oi.Log = in["log"].(int)
		oi.LockOut = in["lock_out"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbCommonDdosProtection(d []interface{}) edpt.SlbCommonDdosProtection {

	count1 := len(d)
	var ret edpt.SlbCommonDdosProtection
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpdEnableToggle = in["ipd_enable_toggle"].(string)
		ret.Logging = getObjectSlbCommonDdosProtectionLogging(in["logging"].([]interface{}))
		ret.PacketsPerSecond = getObjectSlbCommonDdosProtectionPacketsPerSecond(in["packets_per_second"].([]interface{}))
	}
	return ret
}

func getObjectSlbCommonDdosProtectionLogging(d []interface{}) edpt.SlbCommonDdosProtectionLogging {

	count1 := len(d)
	var ret edpt.SlbCommonDdosProtectionLogging
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpdLoggingToggle = in["ipd_logging_toggle"].(string)
	}
	return ret
}

func getObjectSlbCommonDdosProtectionPacketsPerSecond(d []interface{}) edpt.SlbCommonDdosProtectionPacketsPerSecond {

	count1 := len(d)
	var ret edpt.SlbCommonDdosProtectionPacketsPerSecond
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpdTcp = in["ipd_tcp"].(int)
		ret.IpdUdp = in["ipd_udp"].(int)
	}
	return ret
}

func getObjectSlbCommonDnsResponseRateLimiting1411(d []interface{}) edpt.SlbCommonDnsResponseRateLimiting1411 {

	count1 := len(d)
	var ret edpt.SlbCommonDnsResponseRateLimiting1411
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxTableEntries = in["max_table_entries"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbCommonQuic1412(d []interface{}) edpt.SlbCommonQuic1412 {

	count1 := len(d)
	var ret edpt.SlbCommonQuic1412
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CidLen = in["cid_len"].(int)
		ret.Signature = in["signature"].(string)
		ret.SignatureLen = in["signature_len"].(int)
		ret.SignatureOffset = in["signature_offset"].(int)
		ret.CpuOffset = in["cpu_offset"].(int)
		ret.QuicLbOffset = in["quic_lb_offset"].(int)
		ret.EnableHash = in["enable_hash"].(int)
		ret.EnableSignature = in["enable_signature"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbCommonSnatPreserve(d []interface{}) edpt.SlbCommonSnatPreserve {

	count1 := len(d)
	var ret edpt.SlbCommonSnatPreserve
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Range = getSliceSlbCommonSnatPreserveRange(in["range"].([]interface{}))
	}
	return ret
}

func getSliceSlbCommonSnatPreserveRange(d []interface{}) []edpt.SlbCommonSnatPreserveRange {

	count1 := len(d)
	ret := make([]edpt.SlbCommonSnatPreserveRange, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbCommonSnatPreserveRange
		oi.Port1 = in["port1"].(int)
		oi.Port2 = in["port2"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbCommonSslRatelimitCfg(d []interface{}) edpt.SlbCommonSslRatelimitCfg {

	count1 := len(d)
	var ret edpt.SlbCommonSslRatelimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DisableRate = in["disable_rate"].(int)
		ret.Tls12Rate = in["tls12_rate"].(int)
		ret.Tls13Rate = in["tls13_rate"].(int)
	}
	return ret
}

func dataToEndpointSlbCommon(d *schema.ResourceData) edpt.SlbCommon {
	var ret edpt.SlbCommon
	ret.Inst.AflexTableEntryAgingInterval = d.Get("aflex_table_entry_aging_interval").(int)
	ret.Inst.AflexTableEntrySync = getObjectSlbCommonAflexTableEntrySync1407(d.Get("aflex_table_entry_sync").([]interface{}))
	ret.Inst.AfterDisable = d.Get("after_disable").(int)
	ret.Inst.AllowInGatewayMode = d.Get("allow_in_gateway_mode").(int)
	ret.Inst.AttackRespCode = d.Get("attack_resp_code").(int)
	ret.Inst.AutoNatNoIpRefresh = d.Get("auto_nat_no_ip_refresh").(string)
	ret.Inst.AutoTranslatePort = d.Get("auto_translate_port").(int)
	ret.Inst.BuffThresh = d.Get("buff_thresh").(int)
	ret.Inst.BuffThreshHwBuff = d.Get("buff_thresh_hw_buff").(int)
	ret.Inst.BuffThreshRelieveThresh = d.Get("buff_thresh_relieve_thresh").(int)
	ret.Inst.BuffThreshSysBuffHigh = d.Get("buff_thresh_sys_buff_high").(int)
	ret.Inst.BuffThreshSysBuffLow = d.Get("buff_thresh_sys_buff_low").(int)
	ret.Inst.CacheExpireTime = d.Get("cache_expire_time").(int)
	ret.Inst.CancelStreamLoopLimit = d.Get("cancel_stream_loop_limit").(int)
	ret.Inst.CertPinning = getObjectSlbCommonCertPinning1408(d.Get("cert_pinning").([]interface{}))
	ret.Inst.ClientsideIp = d.Get("clientside_ip").(string)
	ret.Inst.ClientsideIpv6 = d.Get("clientside_ipv6").(string)
	ret.Inst.CompressBlockSize = d.Get("compress_block_size").(int)
	ret.Inst.ConnRateLimit = getObjectSlbCommonConnRateLimit1410(d.Get("conn_rate_limit").([]interface{}))
	ret.Inst.CustomMessage = d.Get("custom_message").(string)
	ret.Inst.CustomPage = d.Get("custom_page").(string)
	ret.Inst.CustomSignalClist = d.Get("custom_signal_clist").(string)
	ret.Inst.DdosPktCountThresh = d.Get("ddos_pkt_count_thresh").(int)
	ret.Inst.DdosPktSizeThresh = d.Get("ddos_pkt_size_thresh").(int)
	ret.Inst.DdosProtection = getObjectSlbCommonDdosProtection(d.Get("ddos_protection").([]interface{}))
	ret.Inst.DisableAdaptiveResourceCheck = d.Get("disable_adaptive_resource_check").(int)
	ret.Inst.DisablePersistScoring = d.Get("disable_persist_scoring").(int)
	ret.Inst.DisablePortMasking = d.Get("disable_port_masking").(int)
	ret.Inst.DisableServerAutoReselect = d.Get("disable_server_auto_reselect").(int)
	ret.Inst.DnsCacheAge = d.Get("dns_cache_age").(int)
	ret.Inst.DnsCacheAgeMinThreshold = d.Get("dns_cache_age_min_threshold").(int)
	ret.Inst.DnsCacheAgingWeight = d.Get("dns_cache_aging_weight").(int)
	ret.Inst.DnsCacheEnable = d.Get("dns_cache_enable").(int)
	ret.Inst.DnsCacheEntrySize = d.Get("dns_cache_entry_size").(int)
	ret.Inst.DnsCacheSync = d.Get("dns_cache_sync").(int)
	ret.Inst.DnsCacheSyncEntrySize = d.Get("dns_cache_sync_entry_size").(int)
	ret.Inst.DnsCacheSyncTtlThreshold = d.Get("dns_cache_sync_ttl_threshold").(int)
	ret.Inst.DnsCacheTtlAdjustmentEnable = d.Get("dns_cache_ttl_adjustment_enable").(int)
	ret.Inst.DnsNegativeCacheEnable = d.Get("dns_negative_cache_enable").(int)
	ret.Inst.DnsPersistentCacheEnable = d.Get("dns_persistent_cache_enable").(int)
	ret.Inst.DnsPersistentCacheHitThreshold = d.Get("dns_persistent_cache_hit_threshold").(int)
	ret.Inst.DnsPersistentCacheTtlThreshold = d.Get("dns_persistent_cache_ttl_threshold").(int)
	ret.Inst.DnsResponseRateLimiting = getObjectSlbCommonDnsResponseRateLimiting1411(d.Get("dns_response_rate_limiting").([]interface{}))
	ret.Inst.DnsVipStateless = d.Get("dns_vip_stateless").(int)
	ret.Inst.DropIcmpToVipWhenVipDown = d.Get("drop_icmp_to_vip_when_vip_down").(int)
	ret.Inst.DsrHealthCheckEnable = d.Get("dsr_health_check_enable").(int)
	ret.Inst.EcmpHash = d.Get("ecmp_hash").(string)
	ret.Inst.EnableDdos = d.Get("enable_ddos").(int)
	ret.Inst.EnableFastPathRerouting = d.Get("enable_fast_path_rerouting").(int)
	ret.Inst.EnableL7ReqAcct = d.Get("enable_l7_req_acct").(int)
	ret.Inst.Entity = d.Get("entity").(string)
	ret.Inst.ExcludeDestination = d.Get("exclude_destination").(string)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.FastPathDisable = d.Get("fast_path_disable").(int)
	ret.Inst.GatewayHealthCheck = d.Get("gateway_health_check").(int)
	ret.Inst.GracefulShutdown = d.Get("graceful_shutdown").(int)
	ret.Inst.GracefulShutdownEnable = d.Get("graceful_shutdown_enable").(int)
	ret.Inst.HealthCheckToAllVip = d.Get("health_check_to_all_vip").(int)
	ret.Inst.HonorServerResponseTtl = d.Get("honor_server_response_ttl").(int)
	ret.Inst.HttpFastEnable = d.Get("http_fast_enable").(int)
	ret.Inst.HwCompression = d.Get("hw_compression").(int)
	ret.Inst.HwSynRr = d.Get("hw_syn_rr").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Ipv4Offset = d.Get("ipv4_offset").(int)
	ret.Inst.Ipv6Subnet = d.Get("ipv6_subnet").(int)
	ret.Inst.L2l3TrunkLbDisable = d.Get("l2l3_trunk_lb_disable").(int)
	ret.Inst.LogForResetUnknownConn = d.Get("log_for_reset_unknown_conn").(int)
	ret.Inst.LowLatency = d.Get("low_latency").(int)
	ret.Inst.MaxBuffQueuedPerConn = d.Get("max_buff_queued_per_conn").(int)
	ret.Inst.MaxHttpHeaderCount = d.Get("max_http_header_count").(int)
	ret.Inst.MaxLocalRate = d.Get("max_local_rate").(int)
	ret.Inst.MaxPersistentCache = d.Get("max_persistent_cache").(int)
	ret.Inst.MaxRemoteRate = d.Get("max_remote_rate").(int)
	ret.Inst.MonitorModeEnable = d.Get("monitor_mode_enable").(int)
	ret.Inst.MslTime = d.Get("msl_time").(int)
	ret.Inst.MssTable = d.Get("mss_table").(int)
	ret.Inst.MultiCpu = d.Get("multi_cpu").(int)
	ret.Inst.N5New = d.Get("n5_new").(int)
	ret.Inst.N5Old = d.Get("n5_old").(int)
	ret.Inst.NgwafProxyIpv4 = d.Get("ngwaf_proxy_ipv4").(string)
	ret.Inst.NgwafProxyIpv6 = d.Get("ngwaf_proxy_ipv6").(string)
	ret.Inst.NgwafProxyPort = d.Get("ngwaf_proxy_port").(int)
	ret.Inst.NoAutoUpOnAflex = d.Get("no_auto_up_on_aflex").(int)
	ret.Inst.OddEvenNatEnable = d.Get("odd_even_nat_enable").(int)
	ret.Inst.OneServerConnHmRate = d.Get("one_server_conn_hm_rate").(int)
	ret.Inst.OverridePort = d.Get("override_port").(int)
	ret.Inst.PbslbEntryAge = d.Get("pbslb_entry_age").(int)
	ret.Inst.PbslbOverflowGlid = d.Get("pbslb_overflow_glid").(string)
	ret.Inst.PerThrPercent = d.Get("per_thr_percent").(int)
	ret.Inst.PingSweepDetection = d.Get("ping_sweep_detection").(string)
	ret.Inst.PktRateForResetUnknownConn = d.Get("pkt_rate_for_reset_unknown_conn").(int)
	ret.Inst.PlayerIdCheckEnable = d.Get("player_id_check_enable").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortScanDetection = d.Get("port_scan_detection").(string)
	ret.Inst.PreProcessEnable = d.Get("pre_process_enable").(int)
	ret.Inst.Qat = d.Get("qat").(int)
	ret.Inst.Quic = getObjectSlbCommonQuic1412(d.Get("quic").([]interface{}))
	ret.Inst.Range = d.Get("range").(int)
	ret.Inst.RangeEnd = d.Get("range_end").(int)
	ret.Inst.RangeStart = d.Get("range_start").(int)
	ret.Inst.RateLimitLogging = d.Get("rate_limit_logging").(int)
	ret.Inst.RecursiveNsCache = d.Get("recursive_ns_cache").(string)
	ret.Inst.ResetStaleSession = d.Get("reset_stale_session").(int)
	ret.Inst.ResolvePortConflict = d.Get("resolve_port_conflict").(int)
	ret.Inst.ResponseType = d.Get("response_type").(string)
	ret.Inst.ScaleOut = d.Get("scale_out").(int)
	ret.Inst.ScaleOutTrafficMap = d.Get("scale_out_traffic_map").(int)
	ret.Inst.ServersideIp = d.Get("serverside_ip").(string)
	ret.Inst.ServersideIpv6 = d.Get("serverside_ipv6").(string)
	ret.Inst.ServiceGroupOnNoDestNatVports = d.Get("service_group_on_no_dest_nat_vports").(string)
	ret.Inst.ShowSlbServerLegacyCmd = d.Get("show_slb_server_legacy_cmd").(int)
	ret.Inst.ShowSlbServiceGroupLegacyCmd = d.Get("show_slb_service_group_legacy_cmd").(int)
	ret.Inst.ShowSlbVirtualServerLegacyCmd = d.Get("show_slb_virtual_server_legacy_cmd").(int)
	ret.Inst.SnatGwyForL3 = d.Get("snat_gwy_for_l3").(int)
	ret.Inst.SnatOnVip = d.Get("snat_on_vip").(int)
	ret.Inst.SnatPreserve = getObjectSlbCommonSnatPreserve(d.Get("snat_preserve").([]interface{}))
	ret.Inst.Software = d.Get("software").(int)
	ret.Inst.SoftwareTls13 = d.Get("software_tls13").(int)
	ret.Inst.SoftwareTls13Offload = d.Get("software_tls13_offload").(int)
	ret.Inst.SortRes = d.Get("sort_res").(int)
	ret.Inst.SslModuleUsageEnable = d.Get("ssl_module_usage_enable").(int)
	ret.Inst.SslN5DelayTxEnable = d.Get("ssl_n5_delay_tx_enable").(int)
	ret.Inst.SslRatelimitCfg = getObjectSlbCommonSslRatelimitCfg(d.Get("ssl_ratelimit_cfg").([]interface{}))
	ret.Inst.SsliCertNotReadyInspectLimit = d.Get("ssli_cert_not_ready_inspect_limit").(int)
	ret.Inst.SsliCertNotReadyInspectTimeout = d.Get("ssli_cert_not_ready_inspect_timeout").(int)
	ret.Inst.SsliSilentTerminationEnable = d.Get("ssli_silent_termination_enable").(int)
	ret.Inst.SsliSniHashEnable = d.Get("ssli_sni_hash_enable").(int)
	ret.Inst.StatelessSgMultiBinding = d.Get("stateless_sg_multi_binding").(int)
	ret.Inst.StatsDataDisable = d.Get("stats_data_disable").(int)
	ret.Inst.SubstituteSourceMac = d.Get("substitute_source_mac").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.TrafficMapType = d.Get("traffic_map_type").(string)
	ret.Inst.TtlThreshold = d.Get("ttl_threshold").(int)
	ret.Inst.UseDefaultSessCount = d.Get("use_default_sess_count").(int)
	ret.Inst.UseHttpsProxy = d.Get("use_https_proxy").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UseMssTab = d.Get("use_mss_tab").(int)
	//omit uuid
	ret.Inst.VportGlobal = d.Get("vport_global").(int)
	ret.Inst.VportL3v = d.Get("vport_l3v").(int)
	return ret
}
