package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_oper`: Operational Status for the object zone-service-other\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherOperRead,

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
									"current_app_stat1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat1_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat1_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat2": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat2_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat2_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat3": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat3_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat3_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat4_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat4_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat5": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat5_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat5_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat6_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat6_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat7": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat7_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat7_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat8": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat8_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat8_limit": {
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
						"l4_ext_rate": {
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
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "'other': other;",
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
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
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

func resourceDdosDstZonePortZoneServiceOtherOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper := setObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper)
		DdosDstZonePortZoneServiceOtherOperOper := setObjectDdosDstZonePortZoneServiceOtherOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceOtherOperOper)
		DdosDstZonePortZoneServiceOtherOperPatternRecognition := setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognition(res)
		d.Set("pattern_recognition", DdosDstZonePortZoneServiceOtherOperPatternRecognition)
		DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails := setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails(res)
		d.Set("pattern_recognition_pu_details", DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails)
		DdosDstZonePortZoneServiceOtherOperPortInd := setObjectDdosDstZonePortZoneServiceOtherOperPortInd(res)
		d.Set("port_ind", DdosDstZonePortZoneServiceOtherOperPortInd)
		DdosDstZonePortZoneServiceOtherOperProgressionTracking := setObjectDdosDstZonePortZoneServiceOtherOperProgressionTracking(res)
		d.Set("progression_tracking", DdosDstZonePortZoneServiceOtherOperProgressionTracking)
		DdosDstZonePortZoneServiceOtherOperTopkDestinations := setObjectDdosDstZonePortZoneServiceOtherOperTopkDestinations(res)
		d.Set("topk_destinations", DdosDstZonePortZoneServiceOtherOperTopkDestinations)
		DdosDstZonePortZoneServiceOtherOperTopkSources := setObjectDdosDstZonePortZoneServiceOtherOperTopkSources(res)
		d.Set("topk_sources", DdosDstZonePortZoneServiceOtherOperTopkSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper(ret.DtDdosDstZonePortZoneServiceOtherOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper(d edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZonePortZoneServiceOtherOperOper(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstZonePortZoneServiceOtherOperOperDdos_entry_list(ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.ReportingStatus,
			"sources":                 ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.Sources,
			"overflow_policy":         ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.OverflowPolicy,
			"sources_all_entries":     ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.SourcesAllEntries,
			"class_list":              ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.ClassList,
			"subnet_ip_addr":          ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":        ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.SubnetIpv6Addr,
			"ipv6":                    ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.Ipv6,
			"exceeded":                ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.Exceeded,
			"black_listed":            ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.BlackListed,
			"white_listed":            ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.WhiteListed,
			"authenticated":           ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.Authenticated,
			"level":                   ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.Level,
			"app_stat":                ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.AppStat,
			"indicators":              ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.Indicators,
			"indicator_detail":        ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.IndicatorDetail,
			"l4_ext_rate":             ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.L4ExtRate,
			"hw_blacklisted":          ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstZonePortZoneServiceOtherOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstZonePortZoneServiceOtherOperOperDdos_entry_list(d []edpt.DdosDstZonePortZoneServiceOtherOperOperDdos_entry_list) []map[string]interface{} {
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
		in["current_app_stat1"] = item.CurrentAppStat1
		in["is_app_stat1_exceed"] = item.IsAppStat1Exceed
		in["app_stat1_limit"] = item.AppStat1Limit
		in["current_app_stat2"] = item.CurrentAppStat2
		in["is_app_stat2_exceed"] = item.IsAppStat2Exceed
		in["app_stat2_limit"] = item.AppStat2Limit
		in["current_app_stat3"] = item.CurrentAppStat3
		in["is_app_stat3_exceed"] = item.IsAppStat3Exceed
		in["app_stat3_limit"] = item.AppStat3Limit
		in["current_app_stat4"] = item.CurrentAppStat4
		in["is_app_stat4_exceed"] = item.IsAppStat4Exceed
		in["app_stat4_limit"] = item.AppStat4Limit
		in["current_app_stat5"] = item.CurrentAppStat5
		in["is_app_stat5_exceed"] = item.IsAppStat5Exceed
		in["app_stat5_limit"] = item.AppStat5Limit
		in["current_app_stat6"] = item.CurrentAppStat6
		in["is_app_stat6_exceed"] = item.IsAppStat6Exceed
		in["app_stat6_limit"] = item.AppStat6Limit
		in["current_app_stat7"] = item.CurrentAppStat7
		in["is_app_stat7_exceed"] = item.IsAppStat7Exceed
		in["app_stat7_limit"] = item.AppStat7Limit
		in["current_app_stat8"] = item.CurrentAppStat8
		in["is_app_stat8_exceed"] = item.IsAppStat8Exceed
		in["app_stat8_limit"] = item.AppStat8Limit
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

func setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognition(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionOper(ret.DtDdosDstZonePortZoneServiceOtherOper.PatternRecognition.Oper),
		},
	}
}

func setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionOper(d edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList(d []edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper(ret.DtDdosDstZonePortZoneServiceOtherOper.PatternRecognitionPuDetails.Oper),
		},
	}
}

func setObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper(d edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setObjectDdosDstZonePortZoneServiceOtherOperPortInd(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortZoneServiceOtherOperPortIndOper(ret.DtDdosDstZonePortZoneServiceOtherOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstZonePortZoneServiceOtherOperPortIndOper(d edpt.DdosDstZonePortZoneServiceOtherOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList(d []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicators(d []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZonePortZoneServiceOtherOperProgressionTracking(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortZoneServiceOtherOperProgressionTrackingOper(ret.DtDdosDstZonePortZoneServiceOtherOper.ProgressionTracking.Oper),
		},
	}
}

func setObjectDdosDstZonePortZoneServiceOtherOperProgressionTrackingOper(d edpt.DdosDstZonePortZoneServiceOtherOperProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators(d []edpt.DdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZonePortZoneServiceOtherOperTopkDestinations(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortZoneServiceOtherOperTopkDestinationsOper(ret.DtDdosDstZonePortZoneServiceOtherOper.TopkDestinations.Oper),
		},
	}
}

func setObjectDdosDstZonePortZoneServiceOtherOperTopkDestinationsOper(d edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZonePortZoneServiceOtherOperTopkSources(ret edpt.DataDdosDstZonePortZoneServiceOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortZoneServiceOtherOperTopkSourcesOper(ret.DtDdosDstZonePortZoneServiceOtherOper.TopkSources.Oper),
		},
	}
}

func setObjectDdosDstZonePortZoneServiceOtherOperTopkSourcesOper(d edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZonePortZoneServiceOtherOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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
		ret.L4ExtRate = in["l4_ext_rate"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperOperDdos_entry_list
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
		oi.CurrentAppStat1 = in["current_app_stat1"].(string)
		oi.IsAppStat1Exceed = in["is_app_stat1_exceed"].(int)
		oi.AppStat1Limit = in["app_stat1_limit"].(string)
		oi.CurrentAppStat2 = in["current_app_stat2"].(string)
		oi.IsAppStat2Exceed = in["is_app_stat2_exceed"].(int)
		oi.AppStat2Limit = in["app_stat2_limit"].(string)
		oi.CurrentAppStat3 = in["current_app_stat3"].(string)
		oi.IsAppStat3Exceed = in["is_app_stat3_exceed"].(int)
		oi.AppStat3Limit = in["app_stat3_limit"].(string)
		oi.CurrentAppStat4 = in["current_app_stat4"].(string)
		oi.IsAppStat4Exceed = in["is_app_stat4_exceed"].(int)
		oi.AppStat4Limit = in["app_stat4_limit"].(string)
		oi.CurrentAppStat5 = in["current_app_stat5"].(string)
		oi.IsAppStat5Exceed = in["is_app_stat5_exceed"].(int)
		oi.AppStat5Limit = in["app_stat5_limit"].(string)
		oi.CurrentAppStat6 = in["current_app_stat6"].(string)
		oi.IsAppStat6Exceed = in["is_app_stat6_exceed"].(int)
		oi.AppStat6Limit = in["app_stat6_limit"].(string)
		oi.CurrentAppStat7 = in["current_app_stat7"].(string)
		oi.IsAppStat7Exceed = in["is_app_stat7_exceed"].(int)
		oi.AppStat7Limit = in["app_stat7_limit"].(string)
		oi.CurrentAppStat8 = in["current_app_stat8"].(string)
		oi.IsAppStat8Exceed = in["is_app_stat8_exceed"].(int)
		oi.AppStat8Limit = in["app_stat8_limit"].(string)
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

func getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognition(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList
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

func getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperPortInd(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperPortIndOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperProgressionTracking(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperProgressionTrackingOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators
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

func getObjectDdosDstZonePortZoneServiceOtherOperTopkDestinations(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperTopkDestinationsOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperTopkSources(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherOperTopkSourcesOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherOper {
	var ret edpt.DdosDstZonePortZoneServiceOtherOper

	ret.IpFilteringPolicyOper = getObjectDdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstZonePortZoneServiceOtherOperOper(d.Get("oper").([]interface{}))

	ret.PatternRecognition = getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognition(d.Get("pattern_recognition").([]interface{}))

	ret.PatternRecognitionPuDetails = getObjectDdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails(d.Get("pattern_recognition_pu_details").([]interface{}))

	ret.PortInd = getObjectDdosDstZonePortZoneServiceOtherOperPortInd(d.Get("port_ind").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ProgressionTracking = getObjectDdosDstZonePortZoneServiceOtherOperProgressionTracking(d.Get("progression_tracking").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.TopkDestinations = getObjectDdosDstZonePortZoneServiceOtherOperTopkDestinations(d.Get("topk_destinations").([]interface{}))

	ret.TopkSources = getObjectDdosDstZonePortZoneServiceOtherOperTopkSources(d.Get("topk_sources").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
