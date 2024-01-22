package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_name_oper`: Operational Status for the object proto-name\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNameOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6; 'other': ip-proto other; 'gre': ip-proto gre; 'ipv4-encap': ip-proto IPv4 Encapsulation; 'ipv6-encap': ip-proto IPv6 Encapsulation;",
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

func resourceDdosDstZoneIpProtoProtoNameOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper := setObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper)
		DdosDstZoneIpProtoProtoNameOperOper := setObjectDdosDstZoneIpProtoProtoNameOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNameOperOper)
		DdosDstZoneIpProtoProtoNameOperPortInd := setObjectDdosDstZoneIpProtoProtoNameOperPortInd(res)
		d.Set("port_ind", DdosDstZoneIpProtoProtoNameOperPortInd)
		DdosDstZoneIpProtoProtoNameOperProgressionTracking := setObjectDdosDstZoneIpProtoProtoNameOperProgressionTracking(res)
		d.Set("progression_tracking", DdosDstZoneIpProtoProtoNameOperProgressionTracking)
		DdosDstZoneIpProtoProtoNameOperTopkDestinations := setObjectDdosDstZoneIpProtoProtoNameOperTopkDestinations(res)
		d.Set("topk_destinations", DdosDstZoneIpProtoProtoNameOperTopkDestinations)
		DdosDstZoneIpProtoProtoNameOperTopkSources := setObjectDdosDstZoneIpProtoProtoNameOperTopkSources(res)
		d.Set("topk_sources", DdosDstZoneIpProtoProtoNameOperTopkSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper(ret edpt.DataDdosDstZoneIpProtoProtoNameOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper(ret.DtDdosDstZoneIpProtoProtoNameOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper(d edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneIpProtoProtoNameOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNameOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstZoneIpProtoProtoNameOperOperDdos_entry_list(ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.ReportingStatus,
			"sources":                 ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.Sources,
			"overflow_policy":         ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.OverflowPolicy,
			"sources_all_entries":     ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.SourcesAllEntries,
			"class_list":              ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.ClassList,
			"subnet_ip_addr":          ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":        ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.SubnetIpv6Addr,
			"ipv6":                    ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.Ipv6,
			"exceeded":                ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.Exceeded,
			"black_listed":            ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.BlackListed,
			"white_listed":            ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.WhiteListed,
			"authenticated":           ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.Authenticated,
			"level":                   ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.Level,
			"app_stat":                ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.AppStat,
			"indicators":              ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.Indicators,
			"indicator_detail":        ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.IndicatorDetail,
			"hw_blacklisted":          ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstZoneIpProtoProtoNameOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNameOperOperDdos_entry_list(d []edpt.DdosDstZoneIpProtoProtoNameOperOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstZoneIpProtoProtoNameOperPortInd(ret edpt.DataDdosDstZoneIpProtoProtoNameOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNameOperPortIndOper(ret.DtDdosDstZoneIpProtoProtoNameOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNameOperPortIndOper(d edpt.DdosDstZoneIpProtoProtoNameOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList(d []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZoneIpProtoProtoNameOperProgressionTracking(ret edpt.DataDdosDstZoneIpProtoProtoNameOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNameOperProgressionTrackingOper(ret.DtDdosDstZoneIpProtoProtoNameOper.ProgressionTracking.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNameOperProgressionTrackingOper(d edpt.DdosDstZoneIpProtoProtoNameOperProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneIpProtoProtoNameOperTopkDestinations(ret edpt.DataDdosDstZoneIpProtoProtoNameOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNameOperTopkDestinationsOper(ret.DtDdosDstZoneIpProtoProtoNameOper.TopkDestinations.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNameOperTopkDestinationsOper(d edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneIpProtoProtoNameOperTopkSources(ret edpt.DataDdosDstZoneIpProtoProtoNameOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoNameOperTopkSourcesOper(ret.DtDdosDstZoneIpProtoProtoNameOper.TopkSources.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoNameOperTopkSourcesOper(d edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneIpProtoProtoNameOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZoneIpProtoProtoNameOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperOperDdos_entry_list
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

func getObjectDdosDstZoneIpProtoProtoNameOperPortInd(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNameOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperPortIndOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperProgressionTracking(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNameOperProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperProgressionTrackingOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators
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

func getObjectDdosDstZoneIpProtoProtoNameOperTopkDestinations(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNameOperTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperTopkSources(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoNameOperTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameOperTopkSourcesOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNameOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameOper {
	var ret edpt.DdosDstZoneIpProtoProtoNameOper

	ret.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNameOperOper(d.Get("oper").([]interface{}))

	ret.PortInd = getObjectDdosDstZoneIpProtoProtoNameOperPortInd(d.Get("port_ind").([]interface{}))

	ret.ProgressionTracking = getObjectDdosDstZoneIpProtoProtoNameOperProgressionTracking(d.Get("progression_tracking").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.TopkDestinations = getObjectDdosDstZoneIpProtoProtoNameOperTopkDestinations(d.Get("topk_destinations").([]interface{}))

	ret.TopkSources = getObjectDdosDstZoneIpProtoProtoNameOperTopkSources(d.Get("topk_sources").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
