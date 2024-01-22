package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortRangeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_port_range_oper`: Operational Status for the object port-range\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryPortRangeOperRead,

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
			"port_range_end": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range End Port Number",
			},
			"port_range_start": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range Start Port Number",
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

func resourceDdosDstEntryPortRangeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryPortRangeOperIpFilteringPolicyOper := setObjectDdosDstEntryPortRangeOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstEntryPortRangeOperIpFilteringPolicyOper)
		DdosDstEntryPortRangeOperOper := setObjectDdosDstEntryPortRangeOperOper(res)
		d.Set("oper", DdosDstEntryPortRangeOperOper)
		DdosDstEntryPortRangeOperPatternRecognition := setObjectDdosDstEntryPortRangeOperPatternRecognition(res)
		d.Set("pattern_recognition", DdosDstEntryPortRangeOperPatternRecognition)
		DdosDstEntryPortRangeOperPatternRecognitionPuDetails := setObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetails(res)
		d.Set("pattern_recognition_pu_details", DdosDstEntryPortRangeOperPatternRecognitionPuDetails)
		DdosDstEntryPortRangeOperPortInd := setObjectDdosDstEntryPortRangeOperPortInd(res)
		d.Set("port_ind", DdosDstEntryPortRangeOperPortInd)
		DdosDstEntryPortRangeOperProgressionTracking := setObjectDdosDstEntryPortRangeOperProgressionTracking(res)
		d.Set("progression_tracking", DdosDstEntryPortRangeOperProgressionTracking)
		DdosDstEntryPortRangeOperTopkSources := setObjectDdosDstEntryPortRangeOperTopkSources(res)
		d.Set("topk_sources", DdosDstEntryPortRangeOperTopkSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryPortRangeOperIpFilteringPolicyOper(ret edpt.DataDdosDstEntryPortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortRangeOperIpFilteringPolicyOperOper(ret.DtDdosDstEntryPortRangeOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstEntryPortRangeOperIpFilteringPolicyOperOper(d edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryPortRangeOperOper(ret edpt.DataDdosDstEntryPortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstEntryPortRangeOperOperDdos_entry_list(ret.DtDdosDstEntryPortRangeOper.Oper.Ddos_entry_list),
			"resource_limit_config":   ret.DtDdosDstEntryPortRangeOper.Oper.ResourceLimitConfig,
			"reource_limit_alloc":     ret.DtDdosDstEntryPortRangeOper.Oper.ReourceLimitAlloc,
			"resource_limit_remain":   ret.DtDdosDstEntryPortRangeOper.Oper.ResourceLimitRemain,
			"entry_displayed_count":   ret.DtDdosDstEntryPortRangeOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstEntryPortRangeOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstEntryPortRangeOper.Oper.ReportingStatus,
			"all_ports":               ret.DtDdosDstEntryPortRangeOper.Oper.AllPorts,
			"all_src_ports":           ret.DtDdosDstEntryPortRangeOper.Oper.AllSrcPorts,
			"all_ip_protos":           ret.DtDdosDstEntryPortRangeOper.Oper.AllIpProtos,
			"port_protocol":           ret.DtDdosDstEntryPortRangeOper.Oper.PortProtocol,
			"app_stat":                ret.DtDdosDstEntryPortRangeOper.Oper.AppStat,
			"sflow_source_id":         ret.DtDdosDstEntryPortRangeOper.Oper.SflowSourceId,
			"resource_usage":          ret.DtDdosDstEntryPortRangeOper.Oper.ResourceUsage,
			"l4_ext_rate":             ret.DtDdosDstEntryPortRangeOper.Oper.L4ExtRate,
			"hw_blacklisted":          ret.DtDdosDstEntryPortRangeOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstEntryPortRangeOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstEntryPortRangeOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstEntryPortRangeOperOperDdos_entry_list(d []edpt.DdosDstEntryPortRangeOperOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstEntryPortRangeOperPatternRecognition(ret edpt.DataDdosDstEntryPortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortRangeOperPatternRecognitionOper(ret.DtDdosDstEntryPortRangeOper.PatternRecognition.Oper),
		},
	}
}

func setObjectDdosDstEntryPortRangeOperPatternRecognitionOper(d edpt.DdosDstEntryPortRangeOperPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstEntryPortRangeOperPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortRangeOperPatternRecognitionOperFilterList(d []edpt.DdosDstEntryPortRangeOperPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetails(ret edpt.DataDdosDstEntryPortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper(ret.DtDdosDstEntryPortRangeOper.PatternRecognitionPuDetails.Oper),
		},
	}
}

func setObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper(d edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setObjectDdosDstEntryPortRangeOperPortInd(ret edpt.DataDdosDstEntryPortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortRangeOperPortIndOper(ret.DtDdosDstEntryPortRangeOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstEntryPortRangeOperPortIndOper(d edpt.DdosDstEntryPortRangeOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryPortRangeOperPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortRangeOperPortIndOperIndicators(d []edpt.DdosDstEntryPortRangeOperPortIndOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryPortRangeOperProgressionTracking(ret edpt.DataDdosDstEntryPortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortRangeOperProgressionTrackingOper(ret.DtDdosDstEntryPortRangeOper.ProgressionTracking.Oper),
		},
	}
}

func setObjectDdosDstEntryPortRangeOperProgressionTrackingOper(d edpt.DdosDstEntryPortRangeOperProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryPortRangeOperProgressionTrackingOperIndicators(d.Indicators)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortRangeOperProgressionTrackingOperIndicators(d []edpt.DdosDstEntryPortRangeOperProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryPortRangeOperTopkSources(ret edpt.DataDdosDstEntryPortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryPortRangeOperTopkSourcesOper(ret.DtDdosDstEntryPortRangeOper.TopkSources.Oper),
		},
	}
}

func setObjectDdosDstEntryPortRangeOperTopkSourcesOper(d edpt.DdosDstEntryPortRangeOperTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicators(d []edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources(d []edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryList(d []edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators(d []edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstEntryPortRangeOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortRangeOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperOper(d []interface{}) edpt.DdosDstEntryPortRangeOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryPortRangeOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstEntryPortRangeOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryPortRangeOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperOperDdos_entry_list
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

func getObjectDdosDstEntryPortRangeOperPatternRecognition(d []interface{}) edpt.DdosDstEntryPortRangeOperPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortRangeOperPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperPatternRecognitionOper(d []interface{}) edpt.DdosDstEntryPortRangeOperPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstEntryPortRangeOperPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstEntryPortRangeOperPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperPatternRecognitionOperFilterList
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

func getObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperPortInd(d []interface{}) edpt.DdosDstEntryPortRangeOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortRangeOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperPortIndOper(d []interface{}) edpt.DdosDstEntryPortRangeOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortRangeOperPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstEntryPortRangeOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperPortIndOperIndicators
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

func getObjectDdosDstEntryPortRangeOperProgressionTracking(d []interface{}) edpt.DdosDstEntryPortRangeOperProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortRangeOperProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperProgressionTrackingOper(d []interface{}) edpt.DdosDstEntryPortRangeOperProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortRangeOperProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstEntryPortRangeOperProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperProgressionTrackingOperIndicators
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

func getObjectDdosDstEntryPortRangeOperTopkSources(d []interface{}) edpt.DdosDstEntryPortRangeOperTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryPortRangeOperTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeOperTopkSourcesOper(d []interface{}) edpt.DdosDstEntryPortRangeOperTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeOperTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryPortRangeOper(d *schema.ResourceData) edpt.DdosDstEntryPortRangeOper {
	var ret edpt.DdosDstEntryPortRangeOper

	ret.IpFilteringPolicyOper = getObjectDdosDstEntryPortRangeOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstEntryPortRangeOperOper(d.Get("oper").([]interface{}))

	ret.PatternRecognition = getObjectDdosDstEntryPortRangeOperPatternRecognition(d.Get("pattern_recognition").([]interface{}))

	ret.PatternRecognitionPuDetails = getObjectDdosDstEntryPortRangeOperPatternRecognitionPuDetails(d.Get("pattern_recognition_pu_details").([]interface{}))

	ret.PortInd = getObjectDdosDstEntryPortRangeOperPortInd(d.Get("port_ind").([]interface{}))

	ret.PortRangeEnd = d.Get("port_range_end").(int)

	ret.PortRangeStart = d.Get("port_range_start").(int)

	ret.ProgressionTracking = getObjectDdosDstEntryPortRangeOperProgressionTracking(d.Get("progression_tracking").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.TopkSources = getObjectDdosDstEntryPortRangeOperTopkSources(d.Get("topk_sources").([]interface{}))

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
