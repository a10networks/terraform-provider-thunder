package vthunder

// vThunder resource Slb Common

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbCommon() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbCommonCreate,
		Update: resourceSlbCommonUpdate,
		Read:   resourceSlbCommonRead,
		Delete: resourceSlbCommonDelete,

		Schema: map[string]*schema.Schema{
			"auto_nat_no_ip_refresh": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_http_header_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"range_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"software": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_cache_age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
									"lock_out": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"exceed_action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"protocol": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"log": {
										Type:        schema.TypeInt,
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
			"dns_cache_entry_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_local_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_cache_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sort_res": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_hw_buff": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"override_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"max_remote_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_sg_multi_binding": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_sys_buff_high": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss_table": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fast_path_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_buff_queued_per_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_adaptive_resource_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"drop_icmp_to_vip_when_vip_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"scale_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_auto_up_on_aflex": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_sys_buff_low": {
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
			"rate_limit_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssli_sni_hash_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hw_compression": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"graceful_shutdown": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"after_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pkt_rate_for_reset_unknown_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compress_block_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"graceful_shutdown_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_mss_tab": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snat_on_vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"range": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"buff_thresh_relieve_thresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_vip_stateless": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_server_auto_reselect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ddos_protection": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"ipd_enable_toggle": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
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
			"dsr_health_check_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"resolve_port_conflict": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stats_data_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hw_syn_rr": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"range_end": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exclude_destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ttl_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gateway_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_stale_session": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"honor_server_response_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"low_latency": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"player_id_check_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_l7_req_acct": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"response_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"msl_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log_for_reset_unknown_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"entity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceSlbCommonCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating common (Inside resourceSlbCommonCreate)")

	if client.Host != "" {
		vc := dataToSlbCommon(d)
		d.SetId("1")
		go_vthunder.PostSlbCommon(client.Token, vc, client.Host)

		return resourceSlbCommonRead(d, meta)
	}
	return nil
}

func resourceSlbCommonRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading common (Inside resourceSlbCommonRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbCommon(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No common found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbCommonUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbCommonRead(d, meta)
}

func resourceSlbCommonDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbCommonRead(d, meta)
}

//Utility method to instantiate Common Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbCommon(d *schema.ResourceData) go_vthunder.SlbCommon {
	var vc go_vthunder.SlbCommon

	var c go_vthunder.SlbCommonInstance

	c.LowLatency = d.Get("low_latency").(int)
	c.UseMssTab = d.Get("use_mss_tab").(int)
	c.StatsDataDisable = d.Get("stats_data_disable").(int)
	c.CompressBlockSize = d.Get("compress_block_size").(int)
	c.PlayerIDCheckEnable = d.Get("player_id_check_enable").(int)
	c.AfterDisable = d.Get("after_disable").(int)
	c.MslTime = d.Get("msl_time").(int)
	c.GracefulShutdownEnable = d.Get("graceful_shutdown_enable").(int)
	c.BuffThreshHwBuff = d.Get("buff_thresh_hw_buff").(int)
	c.HwSynRr = d.Get("hw_syn_rr").(int)
	c.Entity = d.Get("entity").(string)
	c.ResetStaleSession = d.Get("reset_stale_session").(int)
	c.GatewayHealthCheck = d.Get("gateway_health_check").(int)
	c.ScaleOut = d.Get("scale_out").(int)
	c.GracefulShutdown = d.Get("graceful_shutdown").(int)
	c.RateLimitLogging = d.Get("rate_limit_logging").(int)
	c.FastPathDisable = d.Get("fast_path_disable").(int)
	c.DropIcmpToVipWhenVipDown = d.Get("drop_icmp_to_vip_when_vip_down").(int)
	c.SsliSniHashEnable = d.Get("ssli_sni_hash_enable").(int)
	c.HwCompression = d.Get("hw_compression").(int)
	c.DNSVipStateless = d.Get("dns_vip_stateless").(int)
	c.BuffThreshSysBuffLow = d.Get("buff_thresh_sys_buff_low").(int)
	c.RangeEnd = d.Get("range_end").(int)

	var obj1 go_vthunder.DNSResponseRateLimiting
	prefix1 := "dns_response_rate_limiting.0."
	obj1.MaxTableEntries = d.Get(prefix1 + "max_table_entries").(int)
	c.MaxTableEntries = obj1

	c.DNSCacheEnable = d.Get("dns_cache_enable").(int)
	c.MaxLocalRate = d.Get("max_local_rate").(int)
	c.ExcludeDestination = d.Get("exclude_destination").(string)
	c.DNSCacheAge = d.Get("dns_cache_age").(int)
	c.MaxHTTPHeaderCount = d.Get("max_http_header_count").(int)
	c.L2L3TrunkLbDisable = d.Get("l2l3_trunk_lb_disable").(int)
	c.ResolvePortConflict = d.Get("resolve_port_conflict").(int)
	c.SortRes = d.Get("sort_res").(int)
	c.SnatGwyForL3 = d.Get("snat_gwy_for_l3").(int)
	c.BuffThreshRelieveThresh = d.Get("buff_thresh_relieve_thresh").(int)
	c.DsrHealthCheckEnable = d.Get("dsr_health_check_enable").(int)
	c.BuffThresh = d.Get("buff_thresh").(int)
	c.DNSCacheEntrySize = d.Get("dns_cache_entry_size").(int)
	c.LogForResetUnknownConn = d.Get("log_for_reset_unknown_conn").(int)
	c.AutoNatNoIPRefresh = d.Get("auto_nat_no_ip_refresh").(string)
	c.PktRateForResetUnknownConn = d.Get("pkt_rate_for_reset_unknown_conn").(int)
	c.BuffThreshSysBuffHigh = d.Get("buff_thresh_sys_buff_high").(int)
	c.MaxBuffQueuedPerConn = d.Get("max_buff_queued_per_conn").(int)
	c.MaxRemoteRate = d.Get("max_remote_rate").(int)
	c.TTLThreshold = d.Get("ttl_threshold").(int)
	c.ExtendedStats = d.Get("extended_stats").(int)
	c.EnableL7ReqAcct = d.Get("enable_l7_req_acct").(int)
	c.SnatOnVip = d.Get("snat_on_vip").(int)
	c.RangeStart = d.Get("range_start").(int)
	c.HonorServerResponseTTL = d.Get("honor_server_response_ttl").(int)
	c.Interval = d.Get("interval").(int)
	c.StatelessSgMultiBinding = d.Get("stateless_sg_multi_binding").(int)
	c.DisableAdaptiveResourceCheck = d.Get("disable_adaptive_resource_check").(int)
	c.Range = d.Get("range").(int)

	var obj2 go_vthunder.ConnRateLimit
	prefix2 := "conn_rate_limit.0."

	SrcIpListCount := d.Get(prefix2 + "src_ip_list.#").(int)
	obj2.Protocol = make([]go_vthunder.SrcIPList, 0, SrcIpListCount)

	for i := 0; i < SrcIpListCount; i++ {
		var obj2_1 go_vthunder.SrcIPList
		prefix3 := fmt.Sprintf(prefix2+"src_ip_list.%d.", i)
		obj2_1.Protocol = d.Get(prefix3 + "protocol").(string)
		obj2_1.Log = d.Get(prefix3 + "log").(int)
		obj2_1.LockOut = d.Get(prefix3 + "lock_out").(int)
		obj2_1.LimitPeriod = d.Get(prefix3 + "limit_period").(string)
		obj2_1.Limit = d.Get(prefix3 + "limit").(int)
		obj2_1.ExceedAction = d.Get(prefix3 + "exceed_action").(int)
		obj2_1.Shared = d.Get(prefix3 + "shared").(int)
		obj2.Protocol = append(obj2.Protocol, obj2_1)
	}

	c.Protocol = obj2

	c.MssTable = d.Get("mss_table").(int)
	c.Timeout = d.Get("timeout").(int)
	c.ResponseType = d.Get("response_type").(string)

	var obj3 go_vthunder.DdosProtection
	prefix1 = "ddos_protection.0."

	var obj3_1 go_vthunder.PacketsPerSecond
	prefix2 = prefix1 + "packets_per_second.0."
	obj3_1.IpdTCP = d.Get(prefix2 + "ipd_tcp").(int)
	obj3_1.IpdUDP = d.Get(prefix2 + "ipd_udp").(int)
	obj3.IpdTCP = obj3_1

	var obj3_2 go_vthunder.Logging2
	prefix2 = prefix1 + "logging.0."
	obj3_2.IpdLoggingToggle = d.Get(prefix2 + "ipd_logging_toggle").(string)
	obj3.IpdLoggingToggle = obj3_2

	obj3.IpdEnableToggle = d.Get(prefix1 + "ipd_enable_toggle").(string)
	c.IpdTCP = obj3

	c.OverridePort = d.Get("override_port").(int)
	c.NoAutoUpOnAflex = d.Get("no_auto_up_on_aflex").(int)
	c.DisableServerAutoReselect = d.Get("disable_server_auto_reselect").(int)
	c.Software = d.Get("software").(int)

	vc.UUID = c

	return vc
}
