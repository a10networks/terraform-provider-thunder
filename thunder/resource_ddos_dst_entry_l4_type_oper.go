package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryL4TypeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_l4_type_oper`: Operational Status for the object l4-type\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryL4TypeOperRead,

		Schema: map[string]*schema.Schema{
			"ip_filtering_policy_oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"seq": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"hits": {
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
						"ddos_entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_address_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src_address_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"level_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connections": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"connection_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connection_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"connection_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_kbit_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"kbit_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_frag_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"frag_packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat1_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat2": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat2_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat3": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat3_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat4_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat5": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat5_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat6_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat7": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat7_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat8": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat8_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lockup_time_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_count": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sflow_source_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"debug_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"undefined_port_hit_stats_wellknown": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"counter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"undefined_port_hit_stats_non_wellknown": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_start": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_end": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"entry_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reporting_status": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"undefined_port_hit_statistics": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"undefined_stats_port_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_l4_types": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_blacklisted": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"port_ind": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicators": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"indicator_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicator_index": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"entry_maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"entry_minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"entry_non_zero_minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"entry_average": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"detection_data_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"progression_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicators": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"indicator_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicator_index": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"num_sample": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"average": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"standard_deviation": {
													Type: schema.TypeString, Optional: true, Description: "",
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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': L4-Type TCP; 'udp': L4-Type UDP; 'icmp': L4-Type ICMP; 'other': L4-Type OTHER;",
			},
			"topk_sources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicators": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"indicator_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicator_index": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sources": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"rate": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
									"entry_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address_str": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicators": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"indicator_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"indicator_index": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"rate": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"max_peak": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"psd_wdw_cnt": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
									"next_indicator": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"finished": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"details": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"top_k_key": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryL4TypeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4TypeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryL4TypeOperIpFilteringPolicyOper := setObjectDdosDstEntryL4TypeOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstEntryL4TypeOperIpFilteringPolicyOper)
		DdosDstEntryL4TypeOperOper := setObjectDdosDstEntryL4TypeOperOper(res)
		d.Set("oper", DdosDstEntryL4TypeOperOper)
		DdosDstEntryL4TypeOperPortInd := setObjectDdosDstEntryL4TypeOperPortInd(res)
		d.Set("port_ind", DdosDstEntryL4TypeOperPortInd)
		DdosDstEntryL4TypeOperProgressionTracking := setObjectDdosDstEntryL4TypeOperProgressionTracking(res)
		d.Set("progression_tracking", DdosDstEntryL4TypeOperProgressionTracking)
		DdosDstEntryL4TypeOperTopkSources := setObjectDdosDstEntryL4TypeOperTopkSources(res)
		d.Set("topk_sources", DdosDstEntryL4TypeOperTopkSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryL4TypeOperIpFilteringPolicyOper(ret edpt.DataDdosDstEntryL4TypeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryL4TypeOperIpFilteringPolicyOperOper(ret.DtDdosDstEntryL4TypeOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstEntryL4TypeOperIpFilteringPolicyOperOper(d edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryL4TypeOperOper(ret edpt.DataDdosDstEntryL4TypeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":                        setSliceDdosDstEntryL4TypeOperOperDdos_entry_list(ret.DtDdosDstEntryL4TypeOper.Oper.Ddos_entry_list),
			"undefined_port_hit_stats_wellknown":     setSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown(ret.DtDdosDstEntryL4TypeOper.Oper.UndefinedPortHitStatsWellknown),
			"undefined_port_hit_stats_non_wellknown": setSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown(ret.DtDdosDstEntryL4TypeOper.Oper.UndefinedPortHitStatsNonWellknown),
			"entry_displayed_count":                  ret.DtDdosDstEntryL4TypeOper.Oper.EntryDisplayedCount,
			"service_displayed_count":                ret.DtDdosDstEntryL4TypeOper.Oper.ServiceDisplayedCount,
			"reporting_status":                       ret.DtDdosDstEntryL4TypeOper.Oper.ReportingStatus,
			"undefined_port_hit_statistics":          ret.DtDdosDstEntryL4TypeOper.Oper.UndefinedPortHitStatistics,
			"undefined_stats_port_num":               ret.DtDdosDstEntryL4TypeOper.Oper.UndefinedStatsPortNum,
			"all_l4_types":                           ret.DtDdosDstEntryL4TypeOper.Oper.AllL4Types,
			"hw_blacklisted":                         ret.DtDdosDstEntryL4TypeOper.Oper.HwBlacklisted,
		},
	}
}

func setSliceDdosDstEntryL4TypeOperOperDdos_entry_list(d []edpt.DdosDstEntryL4TypeOperOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["src_address_str"] = item.SrcAddressStr
		in["port_str"] = item.PortStr
		in["state_str"] = item.StateStr
		in["level_str"] = item.LevelStr
		in["current_connections"] = item.CurrentConnections
		in["connection_limit"] = item.ConnectionLimit
		in["current_connection_rate"] = item.CurrentConnectionRate
		in["connection_rate_limit"] = item.ConnectionRateLimit
		in["current_packet_rate"] = item.CurrentPacketRate
		in["packet_rate_limit"] = item.PacketRateLimit
		in["current_kbit_rate"] = item.CurrentKbitRate
		in["kbit_rate_limit"] = item.KbitRateLimit
		in["current_frag_packet_rate"] = item.CurrentFragPacketRate
		in["frag_packet_rate_limit"] = item.FragPacketRateLimit
		in["current_app_stat1"] = item.CurrentAppStat1
		in["app_stat1_limit"] = item.AppStat1Limit
		in["current_app_stat2"] = item.CurrentAppStat2
		in["app_stat2_limit"] = item.AppStat2Limit
		in["current_app_stat3"] = item.CurrentAppStat3
		in["app_stat3_limit"] = item.AppStat3Limit
		in["current_app_stat4"] = item.CurrentAppStat4
		in["app_stat4_limit"] = item.AppStat4Limit
		in["current_app_stat5"] = item.CurrentAppStat5
		in["app_stat5_limit"] = item.AppStat5Limit
		in["current_app_stat6"] = item.CurrentAppStat6
		in["app_stat6_limit"] = item.AppStat6Limit
		in["current_app_stat7"] = item.CurrentAppStat7
		in["app_stat7_limit"] = item.AppStat7Limit
		in["current_app_stat8"] = item.CurrentAppStat8
		in["app_stat8_limit"] = item.AppStat8Limit
		in["age_str"] = item.AgeStr
		in["lockup_time_str"] = item.LockupTimeStr
		in["dynamic_entry_count"] = item.DynamicEntryCount
		in["dynamic_entry_limit"] = item.DynamicEntryLimit
		in["sflow_source_id"] = item.SflowSourceId
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown(d []edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port"] = item.Port
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown(d []edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_start"] = item.PortStart
		in["port_end"] = item.PortEnd
		in["status"] = item.Status
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryL4TypeOperPortInd(ret edpt.DataDdosDstEntryL4TypeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryL4TypeOperPortIndOper(ret.DtDdosDstEntryL4TypeOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstEntryL4TypeOperPortIndOper(d edpt.DdosDstEntryL4TypeOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryL4TypeOperPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryL4TypeOperPortIndOperIndicators(d []edpt.DdosDstEntryL4TypeOperPortIndOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["entry_maximum"] = item.EntryMaximum
		in["entry_minimum"] = item.EntryMinimum
		in["entry_non_zero_minimum"] = item.EntryNonZeroMinimum
		in["entry_average"] = item.EntryAverage
		in["src_maximum"] = item.SrcMaximum
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryL4TypeOperProgressionTracking(ret edpt.DataDdosDstEntryL4TypeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryL4TypeOperProgressionTrackingOper(ret.DtDdosDstEntryL4TypeOper.ProgressionTracking.Oper),
		},
	}
}

func setObjectDdosDstEntryL4TypeOperProgressionTrackingOper(d edpt.DdosDstEntryL4TypeOperProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryL4TypeOperProgressionTrackingOperIndicators(d.Indicators)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryL4TypeOperProgressionTrackingOperIndicators(d []edpt.DdosDstEntryL4TypeOperProgressionTrackingOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["num_sample"] = item.NumSample
		in["average"] = item.Average
		in["maximum"] = item.Maximum
		in["minimum"] = item.Minimum
		in["standard_deviation"] = item.StandardDeviation
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryL4TypeOperTopkSources(ret edpt.DataDdosDstEntryL4TypeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryL4TypeOperTopkSourcesOper(ret.DtDdosDstEntryL4TypeOper.TopkSources.Oper),
		},
	}
}

func setObjectDdosDstEntryL4TypeOperTopkSourcesOper(d edpt.DdosDstEntryL4TypeOperTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicators(d []edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources(d []edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryList(d []edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators(d []edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["max_peak"] = item.MaxPeak
		in["psd_wdw_cnt"] = item.PsdWdwCnt
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstEntryL4TypeOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryL4TypeOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperOper(d []interface{}) edpt.DdosDstEntryL4TypeOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryL4TypeOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.UndefinedPortHitStatsWellknown = getSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown(in["undefined_port_hit_stats_wellknown"].([]interface{}))
		ret.UndefinedPortHitStatsNonWellknown = getSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown(in["undefined_port_hit_stats_non_wellknown"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.UndefinedPortHitStatistics = in["undefined_port_hit_statistics"].(int)
		ret.UndefinedStatsPortNum = in["undefined_stats_port_num"].(int)
		ret.AllL4Types = in["all_l4_types"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryL4TypeOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.PortStr = in["port_str"].(string)
		oi.StateStr = in["state_str"].(string)
		oi.LevelStr = in["level_str"].(string)
		oi.CurrentConnections = in["current_connections"].(string)
		oi.ConnectionLimit = in["connection_limit"].(string)
		oi.CurrentConnectionRate = in["current_connection_rate"].(string)
		oi.ConnectionRateLimit = in["connection_rate_limit"].(string)
		oi.CurrentPacketRate = in["current_packet_rate"].(string)
		oi.PacketRateLimit = in["packet_rate_limit"].(string)
		oi.CurrentKbitRate = in["current_kbit_rate"].(string)
		oi.KbitRateLimit = in["kbit_rate_limit"].(string)
		oi.CurrentFragPacketRate = in["current_frag_packet_rate"].(string)
		oi.FragPacketRateLimit = in["frag_packet_rate_limit"].(string)
		oi.CurrentAppStat1 = in["current_app_stat1"].(string)
		oi.AppStat1Limit = in["app_stat1_limit"].(string)
		oi.CurrentAppStat2 = in["current_app_stat2"].(string)
		oi.AppStat2Limit = in["app_stat2_limit"].(string)
		oi.CurrentAppStat3 = in["current_app_stat3"].(string)
		oi.AppStat3Limit = in["app_stat3_limit"].(string)
		oi.CurrentAppStat4 = in["current_app_stat4"].(string)
		oi.AppStat4Limit = in["app_stat4_limit"].(string)
		oi.CurrentAppStat5 = in["current_app_stat5"].(string)
		oi.AppStat5Limit = in["app_stat5_limit"].(string)
		oi.CurrentAppStat6 = in["current_app_stat6"].(string)
		oi.AppStat6Limit = in["app_stat6_limit"].(string)
		oi.CurrentAppStat7 = in["current_app_stat7"].(string)
		oi.AppStat7Limit = in["app_stat7_limit"].(string)
		oi.CurrentAppStat8 = in["current_app_stat8"].(string)
		oi.AppStat8Limit = in["app_stat8_limit"].(string)
		oi.AgeStr = in["age_str"].(string)
		oi.LockupTimeStr = in["lockup_time_str"].(string)
		oi.DynamicEntryCount = in["dynamic_entry_count"].(string)
		oi.DynamicEntryLimit = in["dynamic_entry_limit"].(string)
		oi.SflowSourceId = in["sflow_source_id"].(string)
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown(d []interface{}) []edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown
		oi.Port = in["port"].(string)
		oi.Counter = in["counter"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown(d []interface{}) []edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown
		oi.PortStart = in["port_start"].(string)
		oi.PortEnd = in["port_end"].(string)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperPortInd(d []interface{}) edpt.DdosDstEntryL4TypeOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryL4TypeOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperPortIndOper(d []interface{}) edpt.DdosDstEntryL4TypeOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryL4TypeOperPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstEntryL4TypeOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.EntryMaximum = in["entry_maximum"].(string)
		oi.EntryMinimum = in["entry_minimum"].(string)
		oi.EntryNonZeroMinimum = in["entry_non_zero_minimum"].(string)
		oi.EntryAverage = in["entry_average"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperProgressionTracking(d []interface{}) edpt.DdosDstEntryL4TypeOperProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryL4TypeOperProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperProgressionTrackingOper(d []interface{}) edpt.DdosDstEntryL4TypeOperProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryL4TypeOperProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstEntryL4TypeOperProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperProgressionTrackingOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.NumSample = in["num_sample"].(int)
		oi.Average = in["average"].(string)
		oi.Maximum = in["maximum"].(string)
		oi.Minimum = in["minimum"].(string)
		oi.StandardDeviation = in["standard_deviation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperTopkSources(d []interface{}) edpt.DdosDstEntryL4TypeOperTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryL4TypeOperTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryL4TypeOperTopkSourcesOper(d []interface{}) edpt.DdosDstEntryL4TypeOperTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeOperTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryL4TypeOper(d *schema.ResourceData) edpt.DdosDstEntryL4TypeOper {
	var ret edpt.DdosDstEntryL4TypeOper

	ret.IpFilteringPolicyOper = getObjectDdosDstEntryL4TypeOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstEntryL4TypeOperOper(d.Get("oper").([]interface{}))

	ret.PortInd = getObjectDdosDstEntryL4TypeOperPortInd(d.Get("port_ind").([]interface{}))

	ret.ProgressionTracking = getObjectDdosDstEntryL4TypeOperProgressionTracking(d.Get("progression_tracking").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.TopkSources = getObjectDdosDstEntryL4TypeOperTopkSources(d.Get("topk_sources").([]interface{}))

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
