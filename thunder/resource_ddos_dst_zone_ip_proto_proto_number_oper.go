package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_number_oper`: Operational Status for the object proto-number\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNumberOperRead,

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
									"bw_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_auth_passed": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"level": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bl_reasoning_rcode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bl_reasoning_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connections": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_connections_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connection_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_connection_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_packet_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_kbit_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_kbit_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"kbit_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_frag_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_frag_packet_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frag_packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lockup_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dynamic_entry_count": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sflow_source_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"debug_str": {
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
						"sources": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"overflow_policy": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sources_all_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"black_listed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"white_listed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"authenticated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"level": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"app_stat": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"indicators": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"indicator_detail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_blacklisted": {
							Type: schema.TypeInt, Optional: true, Description: "",
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
			"port_ind": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_entry_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_address_str": {
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
															"src_maximum": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"src_minimum": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"src_non_zero_minimum": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"src_average": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"score": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"detection_data_source": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"total_score": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"current_level": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_level": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"escalation_timestamp": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"initial_learning": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"active_time": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
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
												"zone_maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"zone_minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"zone_non_zero_minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"zone_average": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"zone_adaptive_threshold": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicator_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"level": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"zone_threshold": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"source_threshold": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"score": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"detection_data_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_score": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_level": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"escalation_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"initial_learning": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"active_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sources_all_entries": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"subnet_ip_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"subnet_ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"details": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sources": {
										Type: schema.TypeInt, Optional: true, Description: "",
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
									"learning_details": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"learning_brief": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recommended_template": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"template_debug_table": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"protocol_num": {
				Type: schema.TypeInt, Required: true, Description: "Protocol Number",
			},
			"topk_destinations": {
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
												"destinations": {
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNumberOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper := setObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper)
		DdosDstZoneIpProtoProtoNumberOperOper := setObjectDdosDstZoneIpProtoProtoNumberOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNumberOperOper)
		DdosDstZoneIpProtoProtoNumberOperPortInd := setObjectDdosDstZoneIpProtoProtoNumberOperPortInd(res)
		d.Set("port_ind", DdosDstZoneIpProtoProtoNumberOperPortInd)
		DdosDstZoneIpProtoProtoNumberOperProgressionTracking := setObjectDdosDstZoneIpProtoProtoNumberOperProgressionTracking(res)
		d.Set("progression_tracking", DdosDstZoneIpProtoProtoNumberOperProgressionTracking)
		DdosDstZoneIpProtoProtoNumberOperTopkDestinations := setObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinations(res)
		d.Set("topk_destinations", DdosDstZoneIpProtoProtoNumberOperTopkDestinations)
		DdosDstZoneIpProtoProtoNumberOperTopkSources := setObjectDdosDstZoneIpProtoProtoNumberOperTopkSources(res)
		d.Set("topk_sources", DdosDstZoneIpProtoProtoNumberOperTopkSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper(ret edpt.DataDdosDstZoneIpProtoProtoNumberOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper(ret.DtDdosDstZoneIpProtoProtoNumberOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper(d edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneIpProtoProtoNumberOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNumberOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list(ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.ReportingStatus,
			"sources":                 ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.Sources,
			"overflow_policy":         ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.OverflowPolicy,
			"sources_all_entries":     ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.SourcesAllEntries,
			"class_list":              ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.ClassList,
			"subnet_ip_addr":          ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":        ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.SubnetIpv6Addr,
			"ipv6":                    ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.Ipv6,
			"exceeded":                ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.Exceeded,
			"black_listed":            ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.BlackListed,
			"white_listed":            ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.WhiteListed,
			"authenticated":           ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.Authenticated,
			"level":                   ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.Level,
			"app_stat":                ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.AppStat,
			"indicators":              ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.Indicators,
			"indicator_detail":        ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.IndicatorDetail,
			"hw_blacklisted":          ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstZoneIpProtoProtoNumberOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list(d []edpt.DdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["bw_state"] = item.BwState
		in["is_auth_passed"] = item.Is_auth_passed
		in["level"] = item.Level
		in["bl_reasoning_rcode"] = item.BlReasoningRcode
		in["bl_reasoning_timestamp"] = item.BlReasoningTimestamp
		in["current_connections"] = item.CurrentConnections
		in["is_connections_exceed"] = item.IsConnectionsExceed
		in["connection_limit"] = item.ConnectionLimit
		in["current_connection_rate"] = item.CurrentConnectionRate
		in["is_connection_rate_exceed"] = item.IsConnectionRateExceed
		in["connection_rate_limit"] = item.ConnectionRateLimit
		in["current_packet_rate"] = item.CurrentPacketRate
		in["is_packet_rate_exceed"] = item.IsPacketRateExceed
		in["packet_rate_limit"] = item.PacketRateLimit
		in["current_kbit_rate"] = item.CurrentKbitRate
		in["is_kbit_rate_exceed"] = item.IsKbitRateExceed
		in["kbit_rate_limit"] = item.KbitRateLimit
		in["current_frag_packet_rate"] = item.CurrentFragPacketRate
		in["is_frag_packet_rate_exceed"] = item.IsFragPacketRateExceed
		in["frag_packet_rate_limit"] = item.FragPacketRateLimit
		in["age"] = item.Age
		in["lockup_time"] = item.LockupTime
		in["dynamic_entry_count"] = item.DynamicEntryCount
		in["dynamic_entry_limit"] = item.DynamicEntryLimit
		in["sflow_source_id"] = item.SflowSourceId
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneIpProtoProtoNumberOperPortInd(ret edpt.DataDdosDstZoneIpProtoProtoNumberOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNumberOperPortIndOper(ret.DtDdosDstZoneIpProtoProtoNumberOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNumberOperPortIndOper(d edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource

	in["total_score"] = d.TotalScore

	in["current_level"] = d.CurrentLevel

	in["escalation_timestamp"] = d.EscalationTimestamp

	in["initial_learning"] = d.InitialLearning

	in["active_time"] = d.ActiveTime

	in["sources_all_entries"] = d.SourcesAllEntries

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["details"] = d.Details

	in["sources"] = d.Sources
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList(d []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators(item.Indicators)
		in["detection_data_source"] = item.DetectionDataSource
		in["total_score"] = item.TotalScore
		in["current_level"] = item.CurrentLevel
		in["src_level"] = item.SrcLevel
		in["escalation_timestamp"] = item.EscalationTimestamp
		in["initial_learning"] = item.InitialLearning
		in["active_time"] = item.ActiveTime
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["src_maximum"] = item.SrcMaximum
		in["src_minimum"] = item.SrcMinimum
		in["src_non_zero_minimum"] = item.SrcNonZeroMinimum
		in["src_average"] = item.SrcAverage
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["zone_maximum"] = item.ZoneMaximum
		in["zone_minimum"] = item.ZoneMinimum
		in["zone_non_zero_minimum"] = item.ZoneNonZeroMinimum
		in["zone_average"] = item.ZoneAverage
		in["zone_adaptive_threshold"] = item.ZoneAdaptiveThreshold
		in["src_maximum"] = item.SrcMaximum
		in["indicator_cfg"] = setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["level"] = item.Level
		in["zone_threshold"] = item.ZoneThreshold
		in["source_threshold"] = item.SourceThreshold
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneIpProtoProtoNumberOperProgressionTracking(ret edpt.DataDdosDstZoneIpProtoProtoNumberOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper(ret.DtDdosDstZoneIpProtoProtoNumberOper.ProgressionTracking.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper(d edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinations(ret edpt.DataDdosDstZoneIpProtoProtoNumberOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper(ret.DtDdosDstZoneIpProtoProtoNumberOper.TopkDestinations.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper(d edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneIpProtoProtoNumberOperTopkSources(ret edpt.DataDdosDstZoneIpProtoProtoNumberOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNumberOperTopkSourcesOper(ret.DtDdosDstZoneIpProtoProtoNumberOper.TopkSources.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNumberOperTopkSourcesOper(d edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.Sources = in["sources"].(int)
		ret.OverflowPolicy = in["overflow_policy"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Exceeded = in["exceeded"].(int)
		ret.BlackListed = in["black_listed"].(int)
		ret.WhiteListed = in["white_listed"].(int)
		ret.Authenticated = in["authenticated"].(int)
		ret.Level = in["level"].(int)
		ret.AppStat = in["app_stat"].(int)
		ret.Indicators = in["indicators"].(int)
		ret.IndicatorDetail = in["indicator_detail"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.BwState = in["bw_state"].(string)
		oi.Is_auth_passed = in["is_auth_passed"].(string)
		oi.Level = in["level"].(int)
		oi.BlReasoningRcode = in["bl_reasoning_rcode"].(string)
		oi.BlReasoningTimestamp = in["bl_reasoning_timestamp"].(string)
		oi.CurrentConnections = in["current_connections"].(string)
		oi.IsConnectionsExceed = in["is_connections_exceed"].(int)
		oi.ConnectionLimit = in["connection_limit"].(string)
		oi.CurrentConnectionRate = in["current_connection_rate"].(string)
		oi.IsConnectionRateExceed = in["is_connection_rate_exceed"].(int)
		oi.ConnectionRateLimit = in["connection_rate_limit"].(string)
		oi.CurrentPacketRate = in["current_packet_rate"].(string)
		oi.IsPacketRateExceed = in["is_packet_rate_exceed"].(int)
		oi.PacketRateLimit = in["packet_rate_limit"].(string)
		oi.CurrentKbitRate = in["current_kbit_rate"].(string)
		oi.IsKbitRateExceed = in["is_kbit_rate_exceed"].(int)
		oi.KbitRateLimit = in["kbit_rate_limit"].(string)
		oi.CurrentFragPacketRate = in["current_frag_packet_rate"].(string)
		oi.IsFragPacketRateExceed = in["is_frag_packet_rate_exceed"].(int)
		oi.FragPacketRateLimit = in["frag_packet_rate_limit"].(string)
		oi.Age = in["age"].(int)
		oi.LockupTime = in["lockup_time"].(int)
		oi.DynamicEntryCount = in["dynamic_entry_count"].(string)
		oi.DynamicEntryLimit = in["dynamic_entry_limit"].(string)
		oi.SflowSourceId = in["sflow_source_id"].(int)
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperPortInd(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperPortIndOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.TotalScore = in["total_score"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.EscalationTimestamp = in["escalation_timestamp"].(string)
		ret.InitialLearning = in["initial_learning"].(string)
		ret.ActiveTime = in["active_time"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Details = in["details"].(int)
		ret.Sources = in["sources"].(int)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
		oi.DetectionDataSource = in["detection_data_source"].(string)
		oi.TotalScore = in["total_score"].(string)
		oi.CurrentLevel = in["current_level"].(string)
		oi.SrcLevel = in["src_level"].(string)
		oi.EscalationTimestamp = in["escalation_timestamp"].(string)
		oi.InitialLearning = in["initial_learning"].(string)
		oi.ActiveTime = in["active_time"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.SrcMinimum = in["src_minimum"].(string)
		oi.SrcNonZeroMinimum = in["src_non_zero_minimum"].(string)
		oi.SrcAverage = in["src_average"].(string)
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperProgressionTracking(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators
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

func getObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinations(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperTopkSources(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberOperTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberOperTopkSourcesOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberOper {
	var ret edpt.DdosDstZoneIpProtoProtoNumberOper

	ret.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberOperOper(d.Get("oper").([]interface{}))

	ret.PortInd = getObjectDdosDstZoneIpProtoProtoNumberOperPortInd(d.Get("port_ind").([]interface{}))

	ret.ProgressionTracking = getObjectDdosDstZoneIpProtoProtoNumberOperProgressionTracking(d.Get("progression_tracking").([]interface{}))

	ret.ProtocolNum = d.Get("protocol_num").(int)

	ret.TopkDestinations = getObjectDdosDstZoneIpProtoProtoNumberOperTopkDestinations(d.Get("topk_destinations").([]interface{}))

	ret.TopkSources = getObjectDdosDstZoneIpProtoProtoNumberOperTopkSources(d.Get("topk_sources").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
