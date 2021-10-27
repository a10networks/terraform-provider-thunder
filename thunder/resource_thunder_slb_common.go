package thunder

//Thunder resource SlbCommon

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbCommon() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbCommonCreate,
		UpdateContext: resourceSlbCommonUpdate,
		ReadContext:   resourceSlbCommonRead,
		DeleteContext: resourceSlbCommonDelete,
		Schema: map[string]*schema.Schema{
			"port_scan_detection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ping_sweep_detection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stats_data_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"graceful_shutdown_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"graceful_shutdown": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"entity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"after_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rate_limit_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_local_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_remote_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exclude_destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auto_translate_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"range": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"range_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"range_end": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_default_sess_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"per_thr_percent": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dsr_health_check_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"one_server_conn_hm_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"aflex_table_entry_aging_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"override_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check_to_all_vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_stale_session": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_cache_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"response_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ttl_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_cache_age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compress_block_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_cache_entry_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_vip_stateless": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"honor_server_response_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_hw_buff": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_relieve_thresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_sys_buff_low": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_sys_buff_high": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_buff_queued_per_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pkt_rate_for_reset_unknown_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log_for_reset_unknown_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gateway_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"msl_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fast_path_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"http_fast_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l2l3_trunk_lb_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snat_gwy_for_l3": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"allow_in_gateway_mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_server_auto_reselect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_l7_req_acct": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_adaptive_resource_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ddos_pkt_size_thresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ddos_pkt_count_thresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snat_on_vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"low_latency": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss_table": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"resolve_port_conflict": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_auto_up_on_aflex": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hw_compression": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hw_syn_rr": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_http_header_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"scale_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"scale_out_traffic_map": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"show_slb_server_legacy_cmd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"show_slb_service_group_legacy_cmd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"show_slb_virtual_server_legacy_cmd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_map_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sort_res": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_mss_tab": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auto_nat_no_ip_refresh": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ddos_protection": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipd_enable_toggle": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"logging": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipd_logging_toggle": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"packets_per_second": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipd_tcp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ipd_udp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"ssli_sni_hash_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"software": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"software_tls13": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"qat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"n5_new": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"n5_old": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"software_tls13_offload": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"substitute_source_mac": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"drop_icmp_to_vip_when_vip_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"player_id_check_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_sg_multi_binding": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ecmp_hash": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"service_group_on_no_dest_nat_vports": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"disable_port_masking": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snat_preserve": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port1": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port2": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"disable_persist_scoring": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ipv4_offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pre_process_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"attack_resp_code": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cert_pinning": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"candidate_list_feedback_opt_in": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"schedule": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"weekly": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"week_day": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"week_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"daily": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"day_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"use_mgmt_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"aflex_table_entry_sync": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aflex_table_entry_sync_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"aflex_table_entry_sync_max_key_len": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"aflex_table_entry_sync_max_value_len": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"aflex_table_entry_sync_min_lifetime": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"conn_rate_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ip_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"limit_period": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"exceed_action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lock_out": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"dns_response_rate_limiting": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_table_entries": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbCommonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbCommon (Inside resourceSlbCommonCreate) ")

		data := dataToSlbCommon(d)
		logger.Println("[INFO] received formatted data from method data to SlbCommon --")
		d.SetId("1")
		err := go_thunder.PostSlbCommon(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbCommonRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbCommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbCommon (Inside resourceSlbCommonRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSlbCommon(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return diags
}

func resourceSlbCommonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbCommonRead(ctx, d, meta)
}

func resourceSlbCommonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbCommonRead(ctx, d, meta)
}
func dataToSlbCommon(d *schema.ResourceData) go_thunder.SlbCommon {
	var vc go_thunder.SlbCommon
	var c go_thunder.SlbCommonInstance
	c.SlbCommonInstancePortScanDetection = d.Get("port_scan_detection").(string)
	c.SlbCommonInstancePingSweepDetection = d.Get("ping_sweep_detection").(string)
	c.SlbCommonInstanceExtendedStats = d.Get("extended_stats").(int)
	c.SlbCommonInstanceStatsDataDisable = d.Get("stats_data_disable").(int)
	c.SlbCommonInstanceGracefulShutdownEnable = d.Get("graceful_shutdown_enable").(int)
	c.SlbCommonInstanceGracefulShutdown = d.Get("graceful_shutdown").(int)
	c.SlbCommonInstanceEntity = d.Get("entity").(string)
	c.SlbCommonInstanceAfterDisable = d.Get("after_disable").(int)
	c.SlbCommonInstanceRateLimitLogging = d.Get("rate_limit_logging").(int)
	c.SlbCommonInstanceMaxLocalRate = d.Get("max_local_rate").(int)
	c.SlbCommonInstanceMaxRemoteRate = d.Get("max_remote_rate").(int)
	c.SlbCommonInstanceExcludeDestination = d.Get("exclude_destination").(string)
	c.SlbCommonInstanceAutoTranslatePort = d.Get("auto_translate_port").(int)
	c.SlbCommonInstanceRange = d.Get("range").(int)
	c.SlbCommonInstanceRangeStart = d.Get("range_start").(int)
	c.SlbCommonInstanceRangeEnd = d.Get("range_end").(int)
	c.SlbCommonInstanceUseDefaultSessCount = d.Get("use_default_sess_count").(int)
	c.SlbCommonInstancePerThrPercent = d.Get("per_thr_percent").(int)
	c.SlbCommonInstanceDsrHealthCheckEnable = d.Get("dsr_health_check_enable").(int)
	c.SlbCommonInstanceOneServerConnHmRate = d.Get("one_server_conn_hm_rate").(int)
	c.SlbCommonInstanceAflexTableEntryAgingInterval = d.Get("aflex_table_entry_aging_interval").(int)
	c.SlbCommonInstanceOverridePort = d.Get("override_port").(int)
	c.SlbCommonInstanceHealthCheckToAllVip = d.Get("health_check_to_all_vip").(int)
	c.SlbCommonInstanceResetStaleSession = d.Get("reset_stale_session").(int)
	c.SlbCommonInstanceDNSCacheEnable = d.Get("dns_cache_enable").(int)
	c.SlbCommonInstanceResponseType = d.Get("response_type").(string)
	c.SlbCommonInstanceTTLThreshold = d.Get("ttl_threshold").(int)
	c.SlbCommonInstanceDNSCacheAge = d.Get("dns_cache_age").(int)
	c.SlbCommonInstanceCompressBlockSize = d.Get("compress_block_size").(int)
	c.SlbCommonInstanceDNSCacheEntrySize = d.Get("dns_cache_entry_size").(int)
	c.SlbCommonInstanceDNSVipStateless = d.Get("dns_vip_stateless").(int)
	c.SlbCommonInstanceHonorServerResponseTTL = d.Get("honor_server_response_ttl").(int)
	c.SlbCommonInstanceBuffThresh = d.Get("buff_thresh").(int)
	c.SlbCommonInstanceBuffThreshHwBuff = d.Get("buff_thresh_hw_buff").(int)
	c.SlbCommonInstanceBuffThreshRelieveThresh = d.Get("buff_thresh_relieve_thresh").(int)
	c.SlbCommonInstanceBuffThreshSysBuffLow = d.Get("buff_thresh_sys_buff_low").(int)
	c.SlbCommonInstanceBuffThreshSysBuffHigh = d.Get("buff_thresh_sys_buff_high").(int)
	c.SlbCommonInstanceMaxBuffQueuedPerConn = d.Get("max_buff_queued_per_conn").(int)
	c.SlbCommonInstancePktRateForResetUnknownConn = d.Get("pkt_rate_for_reset_unknown_conn").(int)
	c.SlbCommonInstanceLogForResetUnknownConn = d.Get("log_for_reset_unknown_conn").(int)
	c.SlbCommonInstanceGatewayHealthCheck = d.Get("gateway_health_check").(int)
	c.SlbCommonInstanceInterval = d.Get("interval").(int)
	c.SlbCommonInstanceTimeout = d.Get("timeout").(int)
	c.SlbCommonInstanceMslTime = d.Get("msl_time").(int)
	c.SlbCommonInstanceFastPathDisable = d.Get("fast_path_disable").(int)
	c.SlbCommonInstanceHTTPFastEnable = d.Get("http_fast_enable").(int)
	c.SlbCommonInstanceL2L3TrunkLbDisable = d.Get("l2l3_trunk_lb_disable").(int)
	c.SlbCommonInstanceSnatGwyForL3 = d.Get("snat_gwy_for_l3").(int)
	c.SlbCommonInstanceAllowInGatewayMode = d.Get("allow_in_gateway_mode").(int)
	c.SlbCommonInstanceDisableServerAutoReselect = d.Get("disable_server_auto_reselect").(int)
	c.SlbCommonInstanceEnableL7ReqAcct = d.Get("enable_l7_req_acct").(int)
	c.SlbCommonInstanceDisableAdaptiveResourceCheck = d.Get("disable_adaptive_resource_check").(int)
	c.SlbCommonInstanceDdosPktSizeThresh = d.Get("ddos_pkt_size_thresh").(int)
	c.SlbCommonInstanceDdosPktCountThresh = d.Get("ddos_pkt_count_thresh").(int)
	c.SlbCommonInstanceSnatOnVip = d.Get("snat_on_vip").(int)
	c.SlbCommonInstanceLowLatency = d.Get("low_latency").(int)
	c.SlbCommonInstanceMssTable = d.Get("mss_table").(int)
	c.SlbCommonInstanceResolvePortConflict = d.Get("resolve_port_conflict").(int)
	c.SlbCommonInstanceNoAutoUpOnAflex = d.Get("no_auto_up_on_aflex").(int)
	c.SlbCommonInstanceHwCompression = d.Get("hw_compression").(int)
	c.SlbCommonInstanceHwSynRr = d.Get("hw_syn_rr").(int)
	c.SlbCommonInstanceMaxHTTPHeaderCount = d.Get("max_http_header_count").(int)
	c.SlbCommonInstanceScaleOut = d.Get("scale_out").(int)
	c.SlbCommonInstanceScaleOutTrafficMap = d.Get("scale_out_traffic_map").(int)
	c.SlbCommonInstanceShowSlbServerLegacyCmd = d.Get("show_slb_server_legacy_cmd").(int)
	c.SlbCommonInstanceShowSlbServiceGroupLegacyCmd = d.Get("show_slb_service_group_legacy_cmd").(int)
	c.SlbCommonInstanceShowSlbVirtualServerLegacyCmd = d.Get("show_slb_virtual_server_legacy_cmd").(int)
	c.SlbCommonInstanceTrafficMapType = d.Get("traffic_map_type").(string)
	c.SlbCommonInstanceSortRes = d.Get("sort_res").(int)
	c.SlbCommonInstanceUseMssTab = d.Get("use_mss_tab").(int)
	c.SlbCommonInstanceAutoNatNoIPRefresh = d.Get("auto_nat_no_ip_refresh").(string)

	var obj1 go_thunder.SlbCommonInstanceDdosProtection
	prefix1 := "ddos_protection.0."
	obj1.SlbCommonInstanceDdosProtectionIpdEnableToggle = d.Get(prefix1 + "ipd_enable_toggle").(string)

	var obj1_1 go_thunder.SlbCommonInstanceDdosProtectionLogging
	prefix1_1 := prefix1 + "logging.0."
	obj1_1.SlbCommonInstanceDdosProtectionLoggingIpdLoggingToggle = d.Get(prefix1_1 + "ipd_logging_toggle").(string)

	obj1.SlbCommonInstanceDdosProtectionLoggingIpdLoggingToggle = obj1_1

	var obj1_2 go_thunder.SlbCommonInstanceDdosProtectionPacketsPerSecond
	prefix1_2 := prefix1 + "packets_per_second.0."
	obj1_2.SlbCommonInstanceDdosProtectionPacketsPerSecondIpdTCP = d.Get(prefix1_2 + "ipd_tcp").(int)
	obj1_2.SlbCommonInstanceDdosProtectionPacketsPerSecondIpdUDP = d.Get(prefix1_2 + "ipd_udp").(int)

	obj1.SlbCommonInstanceDdosProtectionPacketsPerSecondIpdTCP = obj1_2

	c.SlbCommonInstanceDdosProtectionIpdEnableToggle = obj1

	c.SlbCommonInstanceSsliSniHashEnable = d.Get("ssli_sni_hash_enable").(int)
	c.SlbCommonInstanceSoftware = d.Get("software").(int)
	c.SlbCommonInstanceSoftwareTLS13 = d.Get("software_tls13").(int)
	c.SlbCommonInstanceQAT = d.Get("qat").(int)
	c.SlbCommonInstanceN5New = d.Get("n5_new").(int)
	c.SlbCommonInstanceN5Old = d.Get("n5_old").(int)
	c.SlbCommonInstanceSoftwareTLS13Offload = d.Get("software_tls13_offload").(int)
	c.SlbCommonInstanceSubstituteSourceMac = d.Get("substitute_source_mac").(int)
	c.SlbCommonInstanceDropIcmpToVipWhenVipDown = d.Get("drop_icmp_to_vip_when_vip_down").(int)
	c.SlbCommonInstancePlayerIDCheckEnable = d.Get("player_id_check_enable").(int)
	c.SlbCommonInstanceStatelessSgMultiBinding = d.Get("stateless_sg_multi_binding").(int)
	c.SlbCommonInstanceEcmpHash = d.Get("ecmp_hash").(string)
	c.SlbCommonInstanceServiceGroupOnNoDestNatVports = d.Get("service_group_on_no_dest_nat_vports").(string)
	c.SlbCommonInstanceDisablePortMasking = d.Get("disable_port_masking").(int)

	var obj2 go_thunder.SlbCommonInstanceSnatPreserve
	prefix2 := "snat_preserve.0."

	SlbCommonInstanceSnatPreserveRangeCount := d.Get(prefix2 + "range.#").(int)
	obj2.SlbCommonInstanceSnatPreserveRangePort1 = make([]go_thunder.SlbCommonInstanceSnatPreserveRange, 0, SlbCommonInstanceSnatPreserveRangeCount)

	for i := 0; i < SlbCommonInstanceSnatPreserveRangeCount; i++ {
		var obj2_1 go_thunder.SlbCommonInstanceSnatPreserveRange
		prefix2_1 := prefix2 + fmt.Sprintf("range.%d.", i)
		obj2_1.SlbCommonInstanceSnatPreserveRangePort1 = d.Get(prefix2_1 + "port1").(int)
		obj2_1.SlbCommonInstanceSnatPreserveRangePort2 = d.Get(prefix2_1 + "port2").(int)
		obj2.SlbCommonInstanceSnatPreserveRangePort1 = append(obj2.SlbCommonInstanceSnatPreserveRangePort1, obj2_1)
	}

	c.SlbCommonInstanceSnatPreserveRange = obj2

	c.SlbCommonInstanceDisablePersistScoring = d.Get("disable_persist_scoring").(int)
	c.SlbCommonInstanceIpv4Offset = d.Get("ipv4_offset").(int)
	c.SlbCommonInstancePreProcessEnable = d.Get("pre_process_enable").(int)
	c.SlbCommonInstanceAttackRespCode = d.Get("attack_resp_code").(int)

	var obj3 go_thunder.SlbCommonInstanceCertPinning
	prefix3 := "cert_pinning.0."
	obj3.SlbCommonInstanceCertPinningTTL = d.Get(prefix3 + "ttl").(int)

	var obj3_1 go_thunder.SlbCommonInstanceCertPinningCandidateListFeedbackOptIn
	prefix3_1 := prefix3 + "candidate_list_feedback_opt_in.0."
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInEnable = d.Get(prefix3_1 + "enable").(int)
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInSchedule = d.Get(prefix3_1 + "schedule").(int)
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInWeekly = d.Get(prefix3_1 + "weekly").(int)
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInWeekDay = d.Get(prefix3_1 + "week_day").(string)
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInWeekTime = d.Get(prefix3_1 + "week_time").(string)
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInDaily = d.Get(prefix3_1 + "daily").(int)
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInDayTime = d.Get(prefix3_1 + "day_time").(string)
	obj3_1.SlbCommonInstanceCertPinningCandidateListFeedbackOptInUseMgmtPort = d.Get(prefix3_1 + "use_mgmt_port").(int)

	obj3.SlbCommonInstanceCertPinningCandidateListFeedbackOptInEnable = obj3_1

	c.SlbCommonInstanceCertPinningTTL = obj3

	var obj4 go_thunder.SlbCommonInstanceAflexTableEntrySync
	prefix4 := "aflex_table_entry_sync.0."
	obj4.SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncEnable = d.Get(prefix4 + "aflex_table_entry_sync_enable").(int)
	obj4.SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncMaxKeyLen = d.Get(prefix4 + "aflex_table_entry_sync_max_key_len").(int)
	obj4.SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncMaxValueLen = d.Get(prefix4 + "aflex_table_entry_sync_max_value_len").(int)
	obj4.SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncMinLifetime = d.Get(prefix4 + "aflex_table_entry_sync_min_lifetime").(int)

	c.SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncEnable = obj4

	var obj5 go_thunder.SlbCommonInstanceConnRateLimit
	prefix5 := "conn_rate_limit.0."

	SlbCommonInstanceConnRateLimitSrcIPListCount := d.Get(prefix5 + "src_ip_list.#").(int)
	obj5.SlbCommonInstanceConnRateLimitSrcIPListProtocol = make([]go_thunder.SlbCommonInstanceConnRateLimitSrcIPList, 0, SlbCommonInstanceConnRateLimitSrcIPListCount)

	for i := 0; i < SlbCommonInstanceConnRateLimitSrcIPListCount; i++ {
		var obj5_1 go_thunder.SlbCommonInstanceConnRateLimitSrcIPList
		prefix5_1 := prefix5 + fmt.Sprintf("src_ip_list.%d.", i)
		obj5_1.SlbCommonInstanceConnRateLimitSrcIPListProtocol = d.Get(prefix5_1 + "protocol").(string)
		obj5_1.SlbCommonInstanceConnRateLimitSrcIPListLimit = d.Get(prefix5_1 + "limit").(int)
		obj5_1.SlbCommonInstanceConnRateLimitSrcIPListLimitPeriod = d.Get(prefix5_1 + "limit_period").(string)
		obj5_1.SlbCommonInstanceConnRateLimitSrcIPListShared = d.Get(prefix5_1 + "shared").(int)
		obj5_1.SlbCommonInstanceConnRateLimitSrcIPListExceedAction = d.Get(prefix5_1 + "exceed_action").(int)
		obj5_1.SlbCommonInstanceConnRateLimitSrcIPListLog = d.Get(prefix5_1 + "log").(int)
		obj5_1.SlbCommonInstanceConnRateLimitSrcIPListLockOut = d.Get(prefix5_1 + "lock_out").(int)
		obj5.SlbCommonInstanceConnRateLimitSrcIPListProtocol = append(obj5.SlbCommonInstanceConnRateLimitSrcIPListProtocol, obj5_1)
	}

	c.SlbCommonInstanceConnRateLimitSrcIPList = obj5

	var obj6 go_thunder.SlbCommonInstanceDNSResponseRateLimiting
	prefix6 := "dns_response_rate_limiting.0."
	obj6.SlbCommonInstanceDNSResponseRateLimitingMaxTableEntries = d.Get(prefix6 + "max_table_entries").(int)

	c.SlbCommonInstanceDNSResponseRateLimitingMaxTableEntries = obj6

	vc.SlbCommonInstancePortScanDetection = c
	return vc
}
