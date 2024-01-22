package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosProtectionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_protection_oper`: Operational Status for the object protection\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosProtectionOperRead,

		Schema: map[string]*schema.Schema{
			"ipv6_src_hash_mask_bits": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"offsets": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mask_bit_offset_1": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mask_bit_offset_2": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mask_bit_offset_3": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mask_bit_offset_4": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mask_bit_offset_5": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_protection": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"rate_interval": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"use_route": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tap_interfaces": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_auto_learning_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_auto_learning_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_auto_learning_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_auto_learning_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_delay_learning": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"one_arm_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hw_syn_cookie": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sync": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sync_auto_wl": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"bgp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"bgp_auto_wl": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vrrp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vrrp_auto_wl": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mpls_pkt_inspect": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"detection": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ddet_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ddet_cpus": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_dynamic_overflow_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_dynamic_overflow_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_dynamic_overflow_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_dynamic_overflow_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_ano_sec_l3": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_ano_sec_l4_tcp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_ano_sec_l4_udp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_ano_def_l3": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_ano_def_l4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dns_cache_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"warm_up": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dns_zone_transfer_dedicated_cpus": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_dst_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_zone_port_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"interblade_sync_accuracy": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pattern_recognition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pattern_recognition_cpus": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pattern_recognition_hardware_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"detection_window_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disallow_rst_ack_in_syn_auth": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"non_zero_win_size_syncookie": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hw_blocking": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hw_blocking_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"interface_http_health_check": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosProtectionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosProtectionOperIpv6SrcHashMaskBits := setObjectDdosProtectionOperIpv6SrcHashMaskBits(res)
		d.Set("ipv6_src_hash_mask_bits", DdosProtectionOperIpv6SrcHashMaskBits)
		DdosProtectionOperOper := setObjectDdosProtectionOperOper(res)
		d.Set("oper", DdosProtectionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosProtectionOperIpv6SrcHashMaskBits(ret edpt.DataDdosProtectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosProtectionOperIpv6SrcHashMaskBitsOper(ret.DtDdosProtectionOper.Ipv6SrcHashMaskBits.Oper),
		},
	}
}

func setObjectDdosProtectionOperIpv6SrcHashMaskBitsOper(d edpt.DdosProtectionOperIpv6SrcHashMaskBitsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["offsets"] = setSliceDdosProtectionOperIpv6SrcHashMaskBitsOperOffsets(d.Offsets)
	result = append(result, in)
	return result
}

func setSliceDdosProtectionOperIpv6SrcHashMaskBitsOperOffsets(d []edpt.DdosProtectionOperIpv6SrcHashMaskBitsOperOffsets) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["mask_bit_offset_1"] = item.MaskBitOffset1
		in["mask_bit_offset_2"] = item.MaskBitOffset2
		in["mask_bit_offset_3"] = item.MaskBitOffset3
		in["mask_bit_offset_4"] = item.MaskBitOffset4
		in["mask_bit_offset_5"] = item.MaskBitOffset5
		result = append(result, in)
	}
	return result
}

func setObjectDdosProtectionOperOper(ret edpt.DataDdosProtectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_protection":                     ret.DtDdosProtectionOper.Oper.DdosProtection,
			"rate_interval":                       ret.DtDdosProtectionOper.Oper.RateInterval,
			"mode":                                ret.DtDdosProtectionOper.Oper.Mode,
			"use_route":                           ret.DtDdosProtectionOper.Oper.UseRoute,
			"tap_interfaces":                      ret.DtDdosProtectionOper.Oper.TapInterfaces,
			"dst_auto_learning_ipv4":              ret.DtDdosProtectionOper.Oper.DstAutoLearningIpv4,
			"dst_auto_learning_ipv6":              ret.DtDdosProtectionOper.Oper.DstAutoLearningIpv6,
			"src_auto_learning_ipv4":              ret.DtDdosProtectionOper.Oper.SrcAutoLearningIpv4,
			"src_auto_learning_ipv6":              ret.DtDdosProtectionOper.Oper.SrcAutoLearningIpv6,
			"src_delay_learning":                  ret.DtDdosProtectionOper.Oper.SrcDelayLearning,
			"one_arm_mode":                        ret.DtDdosProtectionOper.Oper.OneArmMode,
			"hw_syn_cookie":                       ret.DtDdosProtectionOper.Oper.HwSynCookie,
			"sync":                                ret.DtDdosProtectionOper.Oper.Sync,
			"sync_auto_wl":                        ret.DtDdosProtectionOper.Oper.SyncAutoWl,
			"bgp":                                 ret.DtDdosProtectionOper.Oper.Bgp,
			"bgp_auto_wl":                         ret.DtDdosProtectionOper.Oper.BgpAutoWl,
			"vrrp":                                ret.DtDdosProtectionOper.Oper.Vrrp,
			"vrrp_auto_wl":                        ret.DtDdosProtectionOper.Oper.VrrpAutoWl,
			"mpls_pkt_inspect":                    ret.DtDdosProtectionOper.Oper.MplsPktInspect,
			"detection":                           ret.DtDdosProtectionOper.Oper.Detection,
			"ddet_mode":                           ret.DtDdosProtectionOper.Oper.DdetMode,
			"ddet_cpus":                           ret.DtDdosProtectionOper.Oper.DdetCpus,
			"dst_dynamic_overflow_ipv4":           ret.DtDdosProtectionOper.Oper.DstDynamicOverflowIpv4,
			"dst_dynamic_overflow_ipv6":           ret.DtDdosProtectionOper.Oper.DstDynamicOverflowIpv6,
			"src_dynamic_overflow_ipv4":           ret.DtDdosProtectionOper.Oper.SrcDynamicOverflowIpv4,
			"src_dynamic_overflow_ipv6":           ret.DtDdosProtectionOper.Oper.SrcDynamicOverflowIpv6,
			"ip_ano_sec_l3":                       ret.DtDdosProtectionOper.Oper.IpAnoSecL3,
			"ip_ano_sec_l4_tcp":                   ret.DtDdosProtectionOper.Oper.IpAnoSecL4Tcp,
			"ip_ano_sec_l4_udp":                   ret.DtDdosProtectionOper.Oper.IpAnoSecL4Udp,
			"ip_ano_def_l3":                       ret.DtDdosProtectionOper.Oper.IpAnoDefL3,
			"ip_ano_def_l4":                       ret.DtDdosProtectionOper.Oper.IpAnoDefL4,
			"dns_cache_mode":                      ret.DtDdosProtectionOper.Oper.DnsCacheMode,
			"warm_up":                             ret.DtDdosProtectionOper.Oper.WarmUp,
			"dns_zone_transfer_dedicated_cpus":    ret.DtDdosProtectionOper.Oper.DnsZoneTransferDedicatedCpus,
			"src_dst_entry_limit":                 ret.DtDdosProtectionOper.Oper.SrcDstEntryLimit,
			"src_zone_port_entry_limit":           ret.DtDdosProtectionOper.Oper.SrcZonePortEntryLimit,
			"interblade_sync_accuracy":            ret.DtDdosProtectionOper.Oper.InterbladeSyncAccuracy,
			"pattern_recognition":                 ret.DtDdosProtectionOper.Oper.PatternRecognition,
			"pattern_recognition_cpus":            ret.DtDdosProtectionOper.Oper.PatternRecognitionCpus,
			"pattern_recognition_hardware_filter": ret.DtDdosProtectionOper.Oper.PatternRecognitionHardwareFilter,
			"detection_window_size":               ret.DtDdosProtectionOper.Oper.DetectionWindowSize,
			"disallow_rst_ack_in_syn_auth":        ret.DtDdosProtectionOper.Oper.DisallowRstAckInSynAuth,
			"non_zero_win_size_syncookie":         ret.DtDdosProtectionOper.Oper.NonZeroWinSizeSyncookie,
			"hw_blocking":                         ret.DtDdosProtectionOper.Oper.HwBlocking,
			"hw_blocking_threshold":               ret.DtDdosProtectionOper.Oper.HwBlockingThreshold,
			"interface_http_health_check":         ret.DtDdosProtectionOper.Oper.InterfaceHttpHealthCheck,
		},
	}
}

func getObjectDdosProtectionOperIpv6SrcHashMaskBits(d []interface{}) edpt.DdosProtectionOperIpv6SrcHashMaskBits {

	count1 := len(d)
	var ret edpt.DdosProtectionOperIpv6SrcHashMaskBits
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosProtectionOperIpv6SrcHashMaskBitsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosProtectionOperIpv6SrcHashMaskBitsOper(d []interface{}) edpt.DdosProtectionOperIpv6SrcHashMaskBitsOper {

	count1 := len(d)
	var ret edpt.DdosProtectionOperIpv6SrcHashMaskBitsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Offsets = getSliceDdosProtectionOperIpv6SrcHashMaskBitsOperOffsets(in["offsets"].([]interface{}))
	}
	return ret
}

func getSliceDdosProtectionOperIpv6SrcHashMaskBitsOperOffsets(d []interface{}) []edpt.DdosProtectionOperIpv6SrcHashMaskBitsOperOffsets {

	count1 := len(d)
	ret := make([]edpt.DdosProtectionOperIpv6SrcHashMaskBitsOperOffsets, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosProtectionOperIpv6SrcHashMaskBitsOperOffsets
		oi.MaskBitOffset1 = in["mask_bit_offset_1"].(int)
		oi.MaskBitOffset2 = in["mask_bit_offset_2"].(int)
		oi.MaskBitOffset3 = in["mask_bit_offset_3"].(int)
		oi.MaskBitOffset4 = in["mask_bit_offset_4"].(int)
		oi.MaskBitOffset5 = in["mask_bit_offset_5"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosProtectionOperOper(d []interface{}) edpt.DdosProtectionOperOper {

	count1 := len(d)
	var ret edpt.DdosProtectionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DdosProtection = in["ddos_protection"].(string)
		ret.RateInterval = in["rate_interval"].(string)
		ret.Mode = in["mode"].(string)
		ret.UseRoute = in["use_route"].(string)
		ret.TapInterfaces = in["tap_interfaces"].(string)
		ret.DstAutoLearningIpv4 = in["dst_auto_learning_ipv4"].(string)
		ret.DstAutoLearningIpv6 = in["dst_auto_learning_ipv6"].(string)
		ret.SrcAutoLearningIpv4 = in["src_auto_learning_ipv4"].(string)
		ret.SrcAutoLearningIpv6 = in["src_auto_learning_ipv6"].(string)
		ret.SrcDelayLearning = in["src_delay_learning"].(string)
		ret.OneArmMode = in["one_arm_mode"].(string)
		ret.HwSynCookie = in["hw_syn_cookie"].(string)
		ret.Sync = in["sync"].(string)
		ret.SyncAutoWl = in["sync_auto_wl"].(string)
		ret.Bgp = in["bgp"].(string)
		ret.BgpAutoWl = in["bgp_auto_wl"].(string)
		ret.Vrrp = in["vrrp"].(string)
		ret.VrrpAutoWl = in["vrrp_auto_wl"].(string)
		ret.MplsPktInspect = in["mpls_pkt_inspect"].(string)
		ret.Detection = in["detection"].(string)
		ret.DdetMode = in["ddet_mode"].(string)
		ret.DdetCpus = in["ddet_cpus"].(int)
		ret.DstDynamicOverflowIpv4 = in["dst_dynamic_overflow_ipv4"].(string)
		ret.DstDynamicOverflowIpv6 = in["dst_dynamic_overflow_ipv6"].(string)
		ret.SrcDynamicOverflowIpv4 = in["src_dynamic_overflow_ipv4"].(string)
		ret.SrcDynamicOverflowIpv6 = in["src_dynamic_overflow_ipv6"].(string)
		ret.IpAnoSecL3 = in["ip_ano_sec_l3"].(string)
		ret.IpAnoSecL4Tcp = in["ip_ano_sec_l4_tcp"].(string)
		ret.IpAnoSecL4Udp = in["ip_ano_sec_l4_udp"].(string)
		ret.IpAnoDefL3 = in["ip_ano_def_l3"].(string)
		ret.IpAnoDefL4 = in["ip_ano_def_l4"].(string)
		ret.DnsCacheMode = in["dns_cache_mode"].(string)
		ret.WarmUp = in["warm_up"].(string)
		ret.DnsZoneTransferDedicatedCpus = in["dns_zone_transfer_dedicated_cpus"].(int)
		ret.SrcDstEntryLimit = in["src_dst_entry_limit"].(string)
		ret.SrcZonePortEntryLimit = in["src_zone_port_entry_limit"].(string)
		ret.InterbladeSyncAccuracy = in["interblade_sync_accuracy"].(string)
		ret.PatternRecognition = in["pattern_recognition"].(string)
		ret.PatternRecognitionCpus = in["pattern_recognition_cpus"].(int)
		ret.PatternRecognitionHardwareFilter = in["pattern_recognition_hardware_filter"].(string)
		ret.DetectionWindowSize = in["detection_window_size"].(int)
		ret.DisallowRstAckInSynAuth = in["disallow_rst_ack_in_syn_auth"].(string)
		ret.NonZeroWinSizeSyncookie = in["non_zero_win_size_syncookie"].(string)
		ret.HwBlocking = in["hw_blocking"].(string)
		ret.HwBlockingThreshold = in["hw_blocking_threshold"].(int)
		ret.InterfaceHttpHealthCheck = in["interface_http_health_check"].(string)
	}
	return ret
}

func dataToEndpointDdosProtectionOper(d *schema.ResourceData) edpt.DdosProtectionOper {
	var ret edpt.DdosProtectionOper

	ret.Ipv6SrcHashMaskBits = getObjectDdosProtectionOperIpv6SrcHashMaskBits(d.Get("ipv6_src_hash_mask_bits").([]interface{}))

	ret.Oper = getObjectDdosProtectionOperOper(d.Get("oper").([]interface{}))
	return ret
}
