package thunder

import (
	"context"
	"fmt"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSlbCommon() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbCommonCreate,
		UpdateContext: resourceSlbCommonUpdate,
		ReadContext:   resourceSlbCommonRead,
		DeleteContext: resourceSlbCommonDelete,
		Schema: map[string]*schema.Schema{
			"aflex_table_entry_aging_interval": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "aFleX table entry aging interval in second",
				ValidateFunc: validation.IntBetween(1, 3600),
			},
			"aflex_table_entry_sync": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aflex_table_entry_sync_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable aflex table sync",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"aflex_table_entry_sync_max_key_len": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "aflex table entry max key length to sync",
							ValidateFunc: validation.IntBetween(0, 1000),
						},
						"aflex_table_entry_sync_max_value_len": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "aflex table entry max value length to sync",
							ValidateFunc: validation.IntBetween(0, 1000),
						},
						"aflex_table_entry_sync_min_lifetime": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "aflex table entry minimum lifetime to sync",
							ValidateFunc: validation.IntBetween(0, 65535),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"after_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Graceful shutdown after disable server/port and/or virtual server/port",
				RequiredWith: []string{"graceful_shutdown"}, ValidateFunc: validation.IntBetween(0, 1),
			},
			"allow_in_gateway_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use source NAT gateway for L3 traffic for gateway mode",
				RequiredWith: []string{"snat_gwy_for_l3"}, ValidateFunc: validation.IntBetween(0, 1),
			},
			"auto_nat_no_ip_refresh": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "Disable auto-nat IP refresh with route change or server state UP",
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
			},
			"auto_translate_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto Translate Port range",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"compress_block_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set compression block size (Compression block size in bytes)",
				ValidateFunc: validation.IntBetween(6000, 131008),
			},
			"conn_rate_limit": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': Set TCP connection rate limit; 'udp': Set UDP packet rate limit;",
										ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
									},
									"limit": {
										Type: schema.TypeInt, Optional: true, Description: "Set max connections per period",
										ValidateFunc: validation.IntBetween(1, 1000000),
									},
									"limit_period": {
										Type: schema.TypeString, Optional: true, Description: "'100': 100 ms; '1000': 1000 ms;",
										ValidateFunc: validation.StringInSlice([]string{"100", "1000"}, false),
									},
									"shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set threshold shared amongst all virtual ports",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"exceed_action": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action if threshold exceeded",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send log if threshold exceeded",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"lock_out": {
										Type: schema.TypeInt, Optional: true, Description: "Set lockout period in seconds if threshold exceeded",
										ValidateFunc: validation.IntBetween(1, 3600),
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
			"ddos_pkt_count_thresh": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Set packet count threshold for DDOS, default is 100",
				ValidateFunc: validation.IntBetween(1, 256),
			},
			"ddos_pkt_size_thresh": {
				Type: schema.TypeInt, Optional: true, Default: 64, Description: "Set data packet size threshold for DDOS, default is 64 bytes",
				ValidateFunc: validation.IntBetween(1, 256),
			},
			"ddos_protection": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "SLB DDos Protection",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipd_enable_toggle": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "Toggle SLB DDoS protection; 'disable': Disable SLB DDoS protection (default);",
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
						},
						"logging": {
							Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "Enable or disable logging (default enabled)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipd_logging_toggle": {
										Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable SLB DDoS protection logging (default); 'disable': Disable SLB DDoS protection logging;",
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
									},
								},
							},
						},
						"packets_per_second": {
							Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "Configue packets-per-second thresholds",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipd_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 200, Description: "Configure packets-per-second threshold per TCP port (default: 200)",
										ValidateFunc: validation.IntBetween(0, 65535),
									},
									"ipd_udp": {
										Type: schema.TypeInt, Optional: true, Default: 200, Description: "Configure packets-per-second threshold per UDP port (default: 200)",
										ValidateFunc: validation.IntBetween(0, 65535), 
									},
								},
							},
						},
					},
				},
			},
			"disable_adaptive_resource_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable adaptive resource check based on buffer usage",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"disable_persist_scoring": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Persist Scoring",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"disable_port_masking": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable masking of ports for CPU hashing",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"disable_server_auto_reselect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable auto reselection of server",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"dns_cache_age": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Set DNS cache entry age, default is 300 seconds (1-1000000 seconds, default is 300 seconds)",
				ValidateFunc: validation.IntBetween(1, 1000000),
			},
			"dns_cache_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS cache",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"dns_cache_entry_size": {
				Type: schema.TypeInt, Optional: true, Default: 256, Description: "Set DNS cache entry size, default is 256 bytes (1-4096 bytes, default is 256 bytes)",
				ValidateFunc: validation.IntBetween(1, 4096),
			},
			"dns_response_rate_limiting": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_table_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of entries allowed",
							ValidateFunc: validation.IntBetween(1000, 4194304),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dns_vip_stateless": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS VIP stateless mode",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"drop_icmp_to_vip_when_vip_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop ICMP to VIP when VIP down",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"dsr_health_check_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable dsr-health-check (direct server return health check)",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"ecmp_hash": {
				Type: schema.TypeString, Optional: true, Default: "system-default", Description: "'system-default': Use system default ecmp hashing algorithm; 'connection-based': Use connection information for hashing;",
				ValidateFunc: validation.StringInSlice([]string{"system-default", "connection-based"}, false),
			},
			"enable_l7_req_acct": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable L7 request accounting",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"entity": {
				Type: schema.TypeString, Optional: true, Description: "'server': Graceful shutdown server/port only; 'virtual-server': Graceful shutdown virtual server/port only;",
				RequiredWith: []string{"graceful_shutdown"}, ValidateFunc: validation.StringInSlice([]string{"server", "virtual-server"}, false),
			},
			"exclude_destination": {
				Type: schema.TypeString, Optional: true, Description: "'local': Maximum local rate; 'remote': Maximum remote rate;  (Maximum rates)",
				ValidateFunc: validation.StringInSlice([]string{"local", "remote"}, false),
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable global slb extended statistics",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"fast_path_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable fast path in SLB processing",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"gateway_health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable gateway health check",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"graceful_shutdown": {
				Type: schema.TypeInt, Optional: true, Description: "1-65535, in unit of seconds",
				RequiredWith: []string{"graceful_shutdown_enable"}, ValidateFunc: validation.IntBetween(1, 65535),
			},
			"graceful_shutdown_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable graceful shutdown",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"health_check_to_all_vip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
				RequiredWith: []string{"dsr_health_check_enable"}, ValidateFunc: validation.IntBetween(0, 1),
			},
			"honor_server_response_ttl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Honor the server reponse TTL",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"http_fast_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Http Fast in SLB processing",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"hw_compression": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use hardware compression",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"hw_syn_rr": {
				Type: schema.TypeInt, Optional: true, Description: "Configure hardware SYN round robin (range 1-500000)",
				ValidateFunc: validation.IntBetween(1, 500000),
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the healthcheck interval, default is 5 seconds (Interval Value, in seconds (default 5))",
				RequiredWith: []string{"gateway_health_check"}, ValidateFunc: validation.IntBetween(1, 180),
			},
			"ipv4_offset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv4 Octet Offset for Hash",
				ValidateFunc: validation.IntBetween(0, 3),
			},
			"l2l3_trunk_lb_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable L2/L3 trunk LB",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"log_for_reset_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when rate exceed",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"low_latency": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable low latency mode",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"max_buff_queued_per_conn": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Set per connection buffer threshold (Buffer value range 128-4096)",
				ValidateFunc: validation.IntBetween(128, 4096),
			},
			"max_http_header_count": {
				Type: schema.TypeInt, Optional: true, Default: 90, Description: "Set maximum number of HTTP headers allowed",
				ValidateFunc: validation.IntBetween(90, 255),
			},
			"max_local_rate": {
				Type: schema.TypeInt, Optional: true, Default: 32, Description: "Set maximum local rate",
				ValidateFunc: validation.IntBetween(1, 100),
			},
			"max_remote_rate": {
				Type: schema.TypeInt, Optional: true, Default: 15000, Description: "Set maximum remote rate",
				ValidateFunc: validation.IntBetween(1, 1000000),
			},
			"msl_time": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure maximum session life, default is 2 seconds (1-39 seconds, default is 2 seconds)",
				ValidateFunc: validation.IntBetween(1, 39),
			},
			"mss_table": {
				Type: schema.TypeInt, Optional: true, Default: 536, Description: "Set MSS table (128-750, default is 536)",
				ValidateFunc: validation.IntBetween(128, 750),
			},
			"no_auto_up_on_aflex": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't automatically mark vport up when aFleX is bound",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"one_server_conn_hm_rate": {
				Type: schema.TypeInt, Optional: true, Description: "One Server Conn Health Check Rate",
				ValidateFunc: validation.IntBetween(1, 60),
			},
			"override_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable override port in DSR health check mode",
				RequiredWith: []string{"dsr_health_check_enable"}, ValidateFunc: validation.IntBetween(0, 1),
			},
			"per_thr_percent": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of default session count to use for per thread session table size",
				RequiredWith: []string{"use_default_sess_count"}, ValidateFunc: validation.IntBetween(1, 100),
			},
			"ping_sweep_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable ping sweep detection; 'disable': Disable ping sweep detection(default);",
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
			},
			"pkt_rate_for_reset_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Description: "Configure rate limit for reset-unknown-conn feature",
				ValidateFunc: validation.IntBetween(1, 1048575),
			},
			"player_id_check_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Player id check",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"port_scan_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable port scan detection; 'disable': Disable port scan detection(default);",
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
			},
			"range": {
				Type: schema.TypeInt, Optional: true, Description: "auto translate port range",
				RequiredWith: []string{"auto_translate_port"}, ValidateFunc: validation.IntBetween(1, 3),
			},
			"range_end": {
				Type: schema.TypeInt, Optional: true, Description: "port range end",
				RequiredWith: []string{"range_start"}, ValidateFunc: validation.IntBetween(0, 65535),
			},
			"range_start": {
				Type: schema.TypeInt, Optional: true, Description: "port range start",
				RequiredWith: []string{"range"}, ValidateFunc: validation.IntBetween(0, 65535),
			},
			"rate_limit_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure rate limit logging",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"reset_stale_session": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset if session in delete queue receives a SYN packet",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"resolve_port_conflict": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable client port service port conflicts",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"response_type": {
				Type: schema.TypeString, Optional: true, Description: "'single-answer': Only cache DNS response with single answer; 'round-robin': Round robin;",
				RequiredWith: []string{"dns_cache_enable"}, ValidateFunc: validation.StringInSlice([]string{"single-answer", "round-robin"}, false),
			},
			"scale_out": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB scale out",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"scale_out_traffic_map": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set SLB scaleout traffic-map",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"service_group_on_no_dest_nat_vports": {
				Type: schema.TypeString, Optional: true, Default: "enforce-different", Description: "'allow-same': Allow the binding service-group on no-dest-nat virtual ports; 'enforce-different': Enforce that the same service-group can not be bound on different no-dest-nat virtual ports;",
				ValidateFunc: validation.StringInSlice([]string{"allow-same", "enforce-different"}, false),
			},
			"show_slb_server_legacy_cmd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable show slb server legacy command",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"show_slb_service_group_legacy_cmd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable show slb service-group legacy command",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"show_slb_virtual_server_legacy_cmd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable show slb virtual-server legacy command",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"snat_gwy_for_l3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use source NAT gateway for L3 traffic for transparent mode",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"snat_on_vip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable source NAT traffic against VIP",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"snat_preserve": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "Port preservation pairs", ConflictsWith: []string{"disable_port_masking"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type: schema.TypeList, Optional: true, Description: "Port range",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port1": {
										Type: schema.TypeInt, Optional: true, Default: 1025, Description: "start port",
										ValidateFunc: validation.IntBetween(1025, 65535),
									},
									"port2": {
										Type: schema.TypeInt, Optional: true, Default: 1025, Description: "end port which is greater than start",
										ValidateFunc: validation.IntBetween(1025, 65535),
									},
								},
							},
						},
					},
				},
			},
			"sort_res": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB sorting of resource names",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"ssl_module": {
				Type: schema.TypeString, Optional: true, Description: "SSL module",
				ValidateFunc: validation.StringInSlice([]string{"software", "software-tls13", "QAT", "N5-new", "N5-old"}, false),
			},
			"ssli_cert_not_ready_inspect_limit": {
				Type: schema.TypeInt, Optional: true, Default: 2000, Description: "SSLI asynchronized connection max number, default is 2000 (set to 0 for unlimited size)",
				ValidateFunc: validation.IntBetween(0, 2147483647),
			},
			"ssli_cert_not_ready_inspect_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "SSLI asynchronized connection timeout, default is 10 seconds (seconds, set to 0 for never timeout)",
				ValidateFunc: validation.IntBetween(0, 2147483647),
			},
			"ssli_sni_hash_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSLi SNI hash table",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"stateless_sg_multi_binding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateless service groups to be assigned to multiple L2/L3 DSR VIPs",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"stats_data_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable global slb data statistics",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"substitute_source_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Substitute Source MAC Address to that of the outgoing interface",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "Specify the healthcheck timeout value, default is 15 seconds (Timeout Value, in seconds (default 15))",
				RequiredWith: []string{"interval"}, ValidateFunc: validation.IntBetween(1, 360),
			},
			"traffic_map_type": {
				Type: schema.TypeString, Optional: true, Default: "vport", Description: "'vport': traffic-map per vport; 'global': global traffic-map;",
				RequiredWith: []string{"scale_out_traffic_map"}, ValidateFunc: validation.StringInSlice([]string{"vport", "global"}, false),
			},
			"ttl_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Only cache DNS response with longer TTL",
				RequiredWith: []string{"dns_cache_enable"}, ValidateFunc: validation.IntBetween(1, 10000000),
			},
			"use_default_sess_count": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use default session count",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"use_mss_tab": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use MSS based on internal table for SLB processing",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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

func resourceSlbCommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := edpt.SlbCommon{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
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
		obj := edpt.SlbCommon{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbCommon(d *schema.ResourceData) edpt.SlbCommon {
	var ret edpt.SlbCommon
	count := 0
	ret.Inst.AflexTableEntryAgingInterval = d.Get("aflex_table_entry_aging_interval").(int)
	ret.Inst.AflexTableEntrySync.AflexTableEntrySyncEnable = d.Get("aflex_table_entry_sync.0.aflex_table_entry_sync_enable").(int)
	ret.Inst.AflexTableEntrySync.AflexTableEntrySyncMaxKeyLen = d.Get("aflex_table_entry_sync.0.aflex_table_entry_sync_max_key_len").(int)
	ret.Inst.AflexTableEntrySync.AflexTableEntrySyncMaxValueLen = d.Get("aflex_table_entry_sync.0.aflex_table_entry_sync_max_value_len").(int)
	ret.Inst.AflexTableEntrySync.AflexTableEntrySyncMinLifetime = d.Get("aflex_table_entry_sync.0.aflex_table_entry_sync_min_lifetime").(int)
	//omit uuid
	ret.Inst.AfterDisable = d.Get("after_disable").(int)
	ret.Inst.AllowInGatewayMode = d.Get("allow_in_gateway_mode").(int)
	ret.Inst.AutoNatNoIpRefresh = d.Get("auto_nat_no_ip_refresh").(string)
	ret.Inst.AutoTranslatePort = d.Get("auto_translate_port").(int)
	ret.Inst.CompressBlockSize = d.Get("compress_block_size").(int)
	count = d.Get("conn_rate_limit.0.src_ip_list.#").(int)
	ret.Inst.ConnRateLimit.SrcIpList = make([]edpt.SlbCommonConnRateLimitSrcIpList, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.SlbCommonConnRateLimitSrcIpList
		prefix := fmt.Sprintf("conn_rate_limit.0.src_ip_list.%d.", i)
		obj.Protocol = d.Get(prefix + "protocol").(string)
		obj.Limit = d.Get(prefix + "limit").(int)
		obj.LimitPeriod = d.Get(prefix + "limit_period").(string)
		obj.Shared = d.Get(prefix + "shared").(int)
		obj.ExceedAction = d.Get(prefix + "exceed_action").(int)
		obj.Log = d.Get(prefix + "log").(int)
		obj.LockOut = d.Get(prefix + "lock_out").(int)
		//omit uuid
		ret.Inst.ConnRateLimit.SrcIpList = append(ret.Inst.ConnRateLimit.SrcIpList, obj)
	}
	ret.Inst.DdosPktCountThresh = d.Get("ddos_pkt_count_thresh").(int)
	ret.Inst.DdosPktSizeThresh = d.Get("ddos_pkt_size_thresh").(int)
	ret.Inst.DdosProtection.IpdEnableToggle = d.Get("ddos_protection.0.ipd_enable_toggle").(string)
	ret.Inst.DdosProtection.Logging.IpdLoggingToggle = d.Get("ddos_protection.0.logging.0.ipd_logging_toggle").(string)
	ret.Inst.DdosProtection.PacketsPerSecond.IpdTcp = d.Get("ddos_protection.0.packets_per_second.0.ipd_tcp").(int)
	ret.Inst.DdosProtection.PacketsPerSecond.IpdUdp = d.Get("ddos_protection.0.packets_per_second.0.ipd_udp").(int)
	ret.Inst.DisableAdaptiveResourceCheck = d.Get("disable_adaptive_resource_check").(int)
	ret.Inst.DisablePersistScoring = d.Get("disable_persist_scoring").(int)
	ret.Inst.DisablePortMasking = d.Get("disable_port_masking").(int)
	ret.Inst.DisableServerAutoReselect = d.Get("disable_server_auto_reselect").(int)
	ret.Inst.DnsCacheAge = d.Get("dns_cache_age").(int)
	ret.Inst.DnsCacheEnable = d.Get("dns_cache_enable").(int)
	ret.Inst.DnsCacheEntrySize = d.Get("dns_cache_entry_size").(int)
	ret.Inst.DnsResponseRateLimiting.MaxTableEntries = d.Get("dns_response_rate_limiting.0.max_table_entries").(int)
	//omit uuid
	ret.Inst.DnsVipStateless = d.Get("dns_vip_stateless").(int)
	ret.Inst.DropIcmpToVipWhenVipDown = d.Get("drop_icmp_to_vip_when_vip_down").(int)
	ret.Inst.DsrHealthCheckEnable = d.Get("dsr_health_check_enable").(int)
	ret.Inst.EcmpHash = d.Get("ecmp_hash").(string)
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
	ret.Inst.L2l3TrunkLbDisable = d.Get("l2l3_trunk_lb_disable").(int)
	ret.Inst.LogForResetUnknownConn = d.Get("log_for_reset_unknown_conn").(int)
	ret.Inst.LowLatency = d.Get("low_latency").(int)
	ret.Inst.MaxBuffQueuedPerConn = d.Get("max_buff_queued_per_conn").(int)
	ret.Inst.MaxHttpHeaderCount = d.Get("max_http_header_count").(int)
	ret.Inst.MaxLocalRate = d.Get("max_local_rate").(int)
	ret.Inst.MaxRemoteRate = d.Get("max_remote_rate").(int)
	ret.Inst.MslTime = d.Get("msl_time").(int)
	ret.Inst.MssTable = d.Get("mss_table").(int)
	ret.Inst.NoAutoUpOnAflex = d.Get("no_auto_up_on_aflex").(int)
	ret.Inst.OneServerConnHmRate = d.Get("one_server_conn_hm_rate").(int)
	ret.Inst.OverridePort = d.Get("override_port").(int)
	ret.Inst.PerThrPercent = d.Get("per_thr_percent").(int)
	ret.Inst.PingSweepDetection = d.Get("ping_sweep_detection").(string)
	ret.Inst.PktRateForResetUnknownConn = d.Get("pkt_rate_for_reset_unknown_conn").(int)
	ret.Inst.PlayerIdCheckEnable = d.Get("player_id_check_enable").(int)
	ret.Inst.PortScanDetection = d.Get("port_scan_detection").(string)
	ret.Inst.Range = d.Get("range").(int)
	ret.Inst.RangeEnd = d.Get("range_end").(int)
	ret.Inst.RangeStart = d.Get("range_start").(int)
	ret.Inst.RateLimitLogging = d.Get("rate_limit_logging").(int)
	ret.Inst.ResetStaleSession = d.Get("reset_stale_session").(int)
	ret.Inst.ResolvePortConflict = d.Get("resolve_port_conflict").(int)
	ret.Inst.ResponseType = d.Get("response_type").(string)
	ret.Inst.ScaleOut = d.Get("scale_out").(int)
	ret.Inst.ScaleOutTrafficMap = d.Get("scale_out_traffic_map").(int)
	ret.Inst.ServiceGroupOnNoDestNatVports = d.Get("service_group_on_no_dest_nat_vports").(string)
	ret.Inst.ShowSlbServerLegacyCmd = d.Get("show_slb_server_legacy_cmd").(int)
	ret.Inst.ShowSlbServiceGroupLegacyCmd = d.Get("show_slb_service_group_legacy_cmd").(int)
	ret.Inst.ShowSlbVirtualServerLegacyCmd = d.Get("show_slb_virtual_server_legacy_cmd").(int)
	ret.Inst.SnatGwyForL3 = d.Get("snat_gwy_for_l3").(int)
	ret.Inst.SnatOnVip = d.Get("snat_on_vip").(int)
	count = d.Get("snat_preserve.0.range.#").(int)
	ret.Inst.SnatPreserve.Range = make([]edpt.SlbCommonSnatPreserveRange, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.SlbCommonSnatPreserveRange
		prefix := fmt.Sprintf("snat_preserve.0.range.%d.", i)
		obj.Port1 = d.Get(prefix + "port1").(int)
		obj.Port2 = d.Get(prefix + "port2").(int)
		ret.Inst.SnatPreserve.Range = append(ret.Inst.SnatPreserve.Range, obj)
	}
	ret.Inst.SortRes = d.Get("sort_res").(int)
	ssl_module_str := d.Get("ssl_module").(string)
	if ssl_module_str == "N5-new" {
		ret.Inst.N5New = 1
	} else if ssl_module_str == "N5-old" {
		ret.Inst.N5Old = 1
	} else if ssl_module_str == "software" {
		ret.Inst.Software = 1
	} else if ssl_module_str == "software-tls13" {
		ret.Inst.SoftwareTls13 = 1
	} else if ssl_module_str == "QAT" {
		ret.Inst.Qat = 1
	}
	ret.Inst.SsliCertNotReadyInspectLimit = d.Get("ssli_cert_not_ready_inspect_limit").(int)
	ret.Inst.SsliCertNotReadyInspectTimeout = d.Get("ssli_cert_not_ready_inspect_timeout").(int)
	ret.Inst.SsliSniHashEnable = d.Get("ssli_sni_hash_enable").(int)
	ret.Inst.StatelessSgMultiBinding = d.Get("stateless_sg_multi_binding").(int)
	ret.Inst.StatsDataDisable = d.Get("stats_data_disable").(int)
	ret.Inst.SubstituteSourceMac = d.Get("substitute_source_mac").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.TrafficMapType = d.Get("traffic_map_type").(string)
	ret.Inst.TtlThreshold = d.Get("ttl_threshold").(int)
	ret.Inst.UseDefaultSessCount = d.Get("use_default_sess_count").(int)
	ret.Inst.UseMssTab = d.Get("use_mss_tab").(int)
	//omit uuid
	return ret
}
