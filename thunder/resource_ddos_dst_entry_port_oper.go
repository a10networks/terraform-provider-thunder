package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_port_oper`: Operational Status for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryPortOperRead,

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
						"resource_limit_config": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"reource_limit_alloc": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"resource_limit_remain": {
							Type: schema.TypeString, Optional: true, Description: "",
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
						"all_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_src_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_ip_protos": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"app_stat": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sflow_source_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resource_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_ext_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_blacklisted": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"suffix_request_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"domain_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"pattern_recognition": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peace_pkt_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"war_pkt_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"war_pkt_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filter_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filter_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filter_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"processing_unit": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"filter_enabled": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"hardware_filter": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"filter_expr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"filter_desc": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"sample_ratio": {
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
			"pattern_recognition_pu_details": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"all_filters": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"processing_unit": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"timestamp": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"peace_pkt_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"war_pkt_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"war_pkt_percentage": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"filter_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"filter_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"filter_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"filter_enabled": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"hardware_filter": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"filter_expr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"filter_desc": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"sample_ratio": {
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
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
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
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port;",
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

func resourceDdosDstEntryPortOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryPortOperIpFilteringPolicyOper := setObjectDdosDstEntryPortOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstEntryPortOperIpFilteringPolicyOper)
		DdosDstEntryPortOperOper := setObjectDdosDstEntryPortOperOper(res)
		d.Set("oper", DdosDstEntryPortOperOper)
		DdosDstEntryPortOperPatternRecognition := setObjectDdosDstEntryPortOperPatternRecognition(res)
		d.Set("pattern_recognition", DdosDstEntryPortOperPatternRecognition)
		DdosDstEntryPortOperPatternRecognitionPuDetails := setObjectDdosDstEntryPortOperPatternRecognitionPuDetails(res)
		d.Set("pattern_recognition_pu_details", DdosDstEntryPortOperPatternRecognitionPuDetails)
		DdosDstEntryPortOperPortInd := setObjectDdosDstEntryPortOperPortInd(res)
		d.Set("port_ind", DdosDstEntryPortOperPortInd)
		DdosDstEntryPortOperProgressionTracking := setObjectDdosDstEntryPortOperProgressionTracking(res)
		d.Set("progression_tracking", DdosDstEntryPortOperProgressionTracking)
		DdosDstEntryPortOperTopkSources := setObjectDdosDstEntryPortOperTopkSources(res)
		d.Set("topk_sources", DdosDstEntryPortOperTopkSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryPortOperIpFilteringPolicyOper(ret edpt.DataDdosDstEntryPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortOperIpFilteringPolicyOperOper(ret.DtDdosDstEntryPortOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstEntryPortOperIpFilteringPolicyOperOper(d edpt.DdosDstEntryPortOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryPortOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryPortOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryPortOperOper(ret edpt.DataDdosDstEntryPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstEntryPortOperOperDdos_entry_list(ret.DtDdosDstEntryPortOper.Oper.Ddos_entry_list),
			"resource_limit_config":   ret.DtDdosDstEntryPortOper.Oper.ResourceLimitConfig,
			"reource_limit_alloc":     ret.DtDdosDstEntryPortOper.Oper.ReourceLimitAlloc,
			"resource_limit_remain":   ret.DtDdosDstEntryPortOper.Oper.ResourceLimitRemain,
			"entry_displayed_count":   ret.DtDdosDstEntryPortOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstEntryPortOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstEntryPortOper.Oper.ReportingStatus,
			"all_ports":               ret.DtDdosDstEntryPortOper.Oper.AllPorts,
			"all_src_ports":           ret.DtDdosDstEntryPortOper.Oper.AllSrcPorts,
			"all_ip_protos":           ret.DtDdosDstEntryPortOper.Oper.AllIpProtos,
			"port_protocol":           ret.DtDdosDstEntryPortOper.Oper.PortProtocol,
			"app_stat":                ret.DtDdosDstEntryPortOper.Oper.AppStat,
			"sflow_source_id":         ret.DtDdosDstEntryPortOper.Oper.SflowSourceId,
			"resource_usage":          ret.DtDdosDstEntryPortOper.Oper.ResourceUsage,
			"l4_ext_rate":             ret.DtDdosDstEntryPortOper.Oper.L4ExtRate,
			"hw_blacklisted":          ret.DtDdosDstEntryPortOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstEntryPortOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstEntryPortOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstEntryPortOperOperDdos_entry_list(d []edpt.DdosDstEntryPortOperOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstEntryPortOperPatternRecognition(ret edpt.DataDdosDstEntryPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortOperPatternRecognitionOper(ret.DtDdosDstEntryPortOper.PatternRecognition.Oper),
		},
	}
}

func setObjectDdosDstEntryPortOperPatternRecognitionOper(d edpt.DdosDstEntryPortOperPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstEntryPortOperPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortOperPatternRecognitionOperFilterList(d []edpt.DdosDstEntryPortOperPatternRecognitionOperFilterList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["processing_unit"] = item.ProcessingUnit
		in["filter_enabled"] = item.FilterEnabled
		in["hardware_filter"] = item.HardwareFilter
		in["filter_expr"] = item.FilterExpr
		in["filter_desc"] = item.FilterDesc
		in["sample_ratio"] = item.SampleRatio
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryPortOperPatternRecognitionPuDetails(ret edpt.DataDdosDstEntryPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortOperPatternRecognitionPuDetailsOper(ret.DtDdosDstEntryPortOper.PatternRecognitionPuDetails.Oper),
		},
	}
}

func setObjectDdosDstEntryPortOperPatternRecognitionPuDetailsOper(d edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["processing_unit"] = item.ProcessingUnit
		in["state"] = item.State
		in["timestamp"] = item.Timestamp
		in["peace_pkt_count"] = item.PeacePktCount
		in["war_pkt_count"] = item.WarPktCount
		in["war_pkt_percentage"] = item.WarPktPercentage
		in["filter_threshold"] = item.FilterThreshold
		in["filter_count"] = item.FilterCount
		in["filter_list"] = setSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["filter_enabled"] = item.FilterEnabled
		in["hardware_filter"] = item.HardwareFilter
		in["filter_expr"] = item.FilterExpr
		in["filter_desc"] = item.FilterDesc
		in["sample_ratio"] = item.SampleRatio
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryPortOperPortInd(ret edpt.DataDdosDstEntryPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortOperPortIndOper(ret.DtDdosDstEntryPortOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstEntryPortOperPortIndOper(d edpt.DdosDstEntryPortOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryPortOperPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortOperPortIndOperIndicators(d []edpt.DdosDstEntryPortOperPortIndOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryPortOperProgressionTracking(ret edpt.DataDdosDstEntryPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortOperProgressionTrackingOper(ret.DtDdosDstEntryPortOper.ProgressionTracking.Oper),
		},
	}
}

func setObjectDdosDstEntryPortOperProgressionTrackingOper(d edpt.DdosDstEntryPortOperProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryPortOperProgressionTrackingOperIndicators(d.Indicators)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortOperProgressionTrackingOperIndicators(d []edpt.DdosDstEntryPortOperProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryPortOperTopkSources(ret edpt.DataDdosDstEntryPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortOperTopkSourcesOper(ret.DtDdosDstEntryPortOper.TopkSources.Oper),
		},
	}
}

func setObjectDdosDstEntryPortOperTopkSourcesOper(d edpt.DdosDstEntryPortOperTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryPortOperTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstEntryPortOperTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortOperTopkSourcesOperIndicators(d []edpt.DdosDstEntryPortOperTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstEntryPortOperTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortOperTopkSourcesOperIndicatorsSources(d []edpt.DdosDstEntryPortOperTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortOperTopkSourcesOperEntryList(d []edpt.DdosDstEntryPortOperTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryPortOperTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortOperTopkSourcesOperEntryListIndicators(d []edpt.DdosDstEntryPortOperTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstEntryPortOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryPortOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryPortOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryPortOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryPortOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortOperOper(d []interface{}) edpt.DdosDstEntryPortOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryPortOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.ResourceLimitConfig = in["resource_limit_config"].(string)
		ret.ReourceLimitAlloc = in["reource_limit_alloc"].(string)
		ret.ResourceLimitRemain = in["resource_limit_remain"].(string)
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.AllPorts = in["all_ports"].(int)
		ret.AllSrcPorts = in["all_src_ports"].(int)
		ret.AllIpProtos = in["all_ip_protos"].(int)
		ret.PortProtocol = in["port_protocol"].(string)
		ret.AppStat = in["app_stat"].(int)
		ret.SflowSourceId = in["sflow_source_id"].(int)
		ret.ResourceUsage = in["resource_usage"].(int)
		ret.L4ExtRate = in["l4_ext_rate"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstEntryPortOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryPortOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperOperDdos_entry_list
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

func getObjectDdosDstEntryPortOperPatternRecognition(d []interface{}) edpt.DdosDstEntryPortOperPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortOperPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortOperPatternRecognitionOper(d []interface{}) edpt.DdosDstEntryPortOperPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstEntryPortOperPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortOperPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstEntryPortOperPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperPatternRecognitionOperFilterList
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortOperPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstEntryPortOperPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortOperPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortOperPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortOperPortInd(d []interface{}) edpt.DdosDstEntryPortOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortOperPortIndOper(d []interface{}) edpt.DdosDstEntryPortOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortOperPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
	}
	return ret
}

func getSliceDdosDstEntryPortOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstEntryPortOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperPortIndOperIndicators
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

func getObjectDdosDstEntryPortOperProgressionTracking(d []interface{}) edpt.DdosDstEntryPortOperProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortOperProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortOperProgressionTrackingOper(d []interface{}) edpt.DdosDstEntryPortOperProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortOperProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortOperProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstEntryPortOperProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperProgressionTrackingOperIndicators
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

func getObjectDdosDstEntryPortOperTopkSources(d []interface{}) edpt.DdosDstEntryPortOperTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortOperTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortOperTopkSourcesOper(d []interface{}) edpt.DdosDstEntryPortOperTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortOperTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortOperTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryPortOperTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryPortOperTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstEntryPortOperTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstEntryPortOperTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortOperTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstEntryPortOperTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortOperTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstEntryPortOperTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryPortOperTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortOperTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryPortOperTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortOperTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortOperTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryPortOper(d *schema.ResourceData) edpt.DdosDstEntryPortOper {
	var ret edpt.DdosDstEntryPortOper

	ret.IpFilteringPolicyOper = getObjectDdosDstEntryPortOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstEntryPortOperOper(d.Get("oper").([]interface{}))

	ret.PatternRecognition = getObjectDdosDstEntryPortOperPatternRecognition(d.Get("pattern_recognition").([]interface{}))

	ret.PatternRecognitionPuDetails = getObjectDdosDstEntryPortOperPatternRecognitionPuDetails(d.Get("pattern_recognition_pu_details").([]interface{}))

	ret.PortInd = getObjectDdosDstEntryPortOperPortInd(d.Get("port_ind").([]interface{}))

	ret.PortNum = d.Get("port_num").(int)

	ret.ProgressionTracking = getObjectDdosDstEntryPortOperProgressionTracking(d.Get("progression_tracking").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.TopkSources = getObjectDdosDstEntryPortOperTopkSources(d.Get("topk_sources").([]interface{}))

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
