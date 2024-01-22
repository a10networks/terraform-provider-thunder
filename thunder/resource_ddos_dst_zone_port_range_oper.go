package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_oper`: Operational Status for the object port-range\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeOperRead,

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
			"ips": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"signature_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"match_count": {
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
									"http_filter_rates": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_rate_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"is_http_filter_rate_limit_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"current_http_filter_rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"http_filter_rate_limit": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"response_size_rates": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"response_size_rate_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"is_response_size_rate_limit_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"current_response_size_rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"response_size_rate_limit": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
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
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port; 'quic': QUIC Port;",
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

func resourceDdosDstZonePortRangeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeOperIpFilteringPolicyOper := setObjectDdosDstZonePortRangeOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstZonePortRangeOperIpFilteringPolicyOper)
		DdosDstZonePortRangeOperIps := setObjectDdosDstZonePortRangeOperIps(res)
		d.Set("ips", DdosDstZonePortRangeOperIps)
		DdosDstZonePortRangeOperOper := setObjectDdosDstZonePortRangeOperOper(res)
		d.Set("oper", DdosDstZonePortRangeOperOper)
		DdosDstZonePortRangeOperPatternRecognition := setObjectDdosDstZonePortRangeOperPatternRecognition(res)
		d.Set("pattern_recognition", DdosDstZonePortRangeOperPatternRecognition)
		DdosDstZonePortRangeOperPatternRecognitionPuDetails := setObjectDdosDstZonePortRangeOperPatternRecognitionPuDetails(res)
		d.Set("pattern_recognition_pu_details", DdosDstZonePortRangeOperPatternRecognitionPuDetails)
		DdosDstZonePortRangeOperPortInd := setObjectDdosDstZonePortRangeOperPortInd(res)
		d.Set("port_ind", DdosDstZonePortRangeOperPortInd)
		DdosDstZonePortRangeOperProgressionTracking := setObjectDdosDstZonePortRangeOperProgressionTracking(res)
		d.Set("progression_tracking", DdosDstZonePortRangeOperProgressionTracking)
		DdosDstZonePortRangeOperTopkDestinations := setObjectDdosDstZonePortRangeOperTopkDestinations(res)
		d.Set("topk_destinations", DdosDstZonePortRangeOperTopkDestinations)
		DdosDstZonePortRangeOperTopkSources := setObjectDdosDstZonePortRangeOperTopkSources(res)
		d.Set("topk_sources", DdosDstZonePortRangeOperTopkSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeOperIpFilteringPolicyOper(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperIpFilteringPolicyOperOper(ret.DtDdosDstZonePortRangeOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperIpFilteringPolicyOperOper(d edpt.DdosDstZonePortRangeOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZonePortRangeOperIps(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperIpsOper(ret.DtDdosDstZonePortRangeOper.Ips.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperIpsOper(d edpt.DdosDstZonePortRangeOperIpsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["signature_list"] = setSliceDdosDstZonePortRangeOperIpsOperSignatureList(d.SignatureList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortRangeOperIpsOperSignatureList(d []edpt.DdosDstZonePortRangeOperIpsOperSignatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["match_count"] = item.MatchCount
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZonePortRangeOperOper(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstZonePortRangeOperOperDdos_entry_list(ret.DtDdosDstZonePortRangeOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstZonePortRangeOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstZonePortRangeOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstZonePortRangeOper.Oper.ReportingStatus,
			"sources":                 ret.DtDdosDstZonePortRangeOper.Oper.Sources,
			"overflow_policy":         ret.DtDdosDstZonePortRangeOper.Oper.OverflowPolicy,
			"sources_all_entries":     ret.DtDdosDstZonePortRangeOper.Oper.SourcesAllEntries,
			"class_list":              ret.DtDdosDstZonePortRangeOper.Oper.ClassList,
			"subnet_ip_addr":          ret.DtDdosDstZonePortRangeOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":        ret.DtDdosDstZonePortRangeOper.Oper.SubnetIpv6Addr,
			"ipv6":                    ret.DtDdosDstZonePortRangeOper.Oper.Ipv6,
			"exceeded":                ret.DtDdosDstZonePortRangeOper.Oper.Exceeded,
			"black_listed":            ret.DtDdosDstZonePortRangeOper.Oper.BlackListed,
			"white_listed":            ret.DtDdosDstZonePortRangeOper.Oper.WhiteListed,
			"authenticated":           ret.DtDdosDstZonePortRangeOper.Oper.Authenticated,
			"level":                   ret.DtDdosDstZonePortRangeOper.Oper.Level,
			"app_stat":                ret.DtDdosDstZonePortRangeOper.Oper.AppStat,
			"indicators":              ret.DtDdosDstZonePortRangeOper.Oper.Indicators,
			"indicator_detail":        ret.DtDdosDstZonePortRangeOper.Oper.IndicatorDetail,
			"l4_ext_rate":             ret.DtDdosDstZonePortRangeOper.Oper.L4ExtRate,
			"hw_blacklisted":          ret.DtDdosDstZonePortRangeOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstZonePortRangeOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstZonePortRangeOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstZonePortRangeOperOperDdos_entry_list(d []edpt.DdosDstZonePortRangeOperOperDdos_entry_list) []map[string]interface{} {
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
		in["http_filter_rates"] = setSliceDdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates(item.HttpFilterRates)
		in["response_size_rates"] = setSliceDdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates(item.ResponseSizeRates)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates(d []edpt.DdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["http_filter_rate_name"] = item.HttpFilterRateName
		in["is_http_filter_rate_limit_exceed"] = item.IsHttpFilterRateLimitExceed
		in["current_http_filter_rate"] = item.CurrentHttpFilterRate
		in["http_filter_rate_limit"] = item.HttpFilterRateLimit
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates(d []edpt.DdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["response_size_rate_name"] = item.ResponseSizeRateName
		in["is_response_size_rate_limit_exceed"] = item.IsResponseSizeRateLimitExceed
		in["current_response_size_rate"] = item.CurrentResponseSizeRate
		in["response_size_rate_limit"] = item.ResponseSizeRateLimit
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZonePortRangeOperPatternRecognition(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperPatternRecognitionOper(ret.DtDdosDstZonePortRangeOper.PatternRecognition.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperPatternRecognitionOper(d edpt.DdosDstZonePortRangeOperPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstZonePortRangeOperPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortRangeOperPatternRecognitionOperFilterList(d []edpt.DdosDstZonePortRangeOperPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstZonePortRangeOperPatternRecognitionPuDetails(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperPatternRecognitionPuDetailsOper(ret.DtDdosDstZonePortRangeOper.PatternRecognitionPuDetails.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperPatternRecognitionPuDetailsOper(d edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setObjectDdosDstZonePortRangeOperPortInd(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperPortIndOper(ret.DtDdosDstZonePortRangeOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperPortIndOper(d edpt.DdosDstZonePortRangeOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZonePortRangeOperPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZonePortRangeOperPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZonePortRangeOperPortIndOperSrcEntryList(d []edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZonePortRangeOperPortIndOperIndicators(d []edpt.DdosDstZonePortRangeOperPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZonePortRangeOperProgressionTracking(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperProgressionTrackingOper(ret.DtDdosDstZonePortRangeOper.ProgressionTracking.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperProgressionTrackingOper(d edpt.DdosDstZonePortRangeOperProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZonePortRangeOperProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortRangeOperProgressionTrackingOperIndicators(d []edpt.DdosDstZonePortRangeOperProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZonePortRangeOperTopkDestinations(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperTopkDestinationsOper(ret.DtDdosDstZonePortRangeOper.TopkDestinations.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperTopkDestinationsOper(d edpt.DdosDstZonePortRangeOperTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicators(d []edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryList(d []edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZonePortRangeOperTopkSources(ret edpt.DataDdosDstZonePortRangeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZonePortRangeOperTopkSourcesOper(ret.DtDdosDstZonePortRangeOper.TopkSources.Oper),
		},
	}
}

func setObjectDdosDstZonePortRangeOperTopkSourcesOper(d edpt.DdosDstZonePortRangeOperTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZonePortRangeOperTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZonePortRangeOperTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZonePortRangeOperTopkSourcesOperIndicators(d []edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperTopkSourcesOperEntryList(d []edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortRangeOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZonePortRangeOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZonePortRangeOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperIps(d []interface{}) edpt.DdosDstZonePortRangeOperIps {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperIps
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperIpsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperIpsOper(d []interface{}) edpt.DdosDstZonePortRangeOperIpsOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperIpsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignatureList = getSliceDdosDstZonePortRangeOperIpsOperSignatureList(in["signature_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperIpsOperSignatureList(d []interface{}) []edpt.DdosDstZonePortRangeOperIpsOperSignatureList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperIpsOperSignatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperIpsOperSignatureList
		oi.Sid = in["sid"].(int)
		oi.MatchCount = in["match_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperOper(d []interface{}) edpt.DdosDstZonePortRangeOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZonePortRangeOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZonePortRangeOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZonePortRangeOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperOperDdos_entry_list
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
		oi.HttpFilterRates = getSliceDdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates(in["http_filter_rates"].([]interface{}))
		oi.ResponseSizeRates = getSliceDdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates(in["response_size_rates"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates(d []interface{}) []edpt.DdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates
		oi.HttpFilterRateName = in["http_filter_rate_name"].(string)
		oi.IsHttpFilterRateLimitExceed = in["is_http_filter_rate_limit_exceed"].(int)
		oi.CurrentHttpFilterRate = in["current_http_filter_rate"].(string)
		oi.HttpFilterRateLimit = in["http_filter_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates(d []interface{}) []edpt.DdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates
		oi.ResponseSizeRateName = in["response_size_rate_name"].(string)
		oi.IsResponseSizeRateLimitExceed = in["is_response_size_rate_limit_exceed"].(int)
		oi.CurrentResponseSizeRate = in["current_response_size_rate"].(string)
		oi.ResponseSizeRateLimit = in["response_size_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperPatternRecognition(d []interface{}) edpt.DdosDstZonePortRangeOperPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperPatternRecognitionOper(d []interface{}) edpt.DdosDstZonePortRangeOperPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstZonePortRangeOperPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstZonePortRangeOperPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperPatternRecognitionOperFilterList
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

func getObjectDdosDstZonePortRangeOperPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperPortInd(d []interface{}) edpt.DdosDstZonePortRangeOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperPortIndOper(d []interface{}) edpt.DdosDstZonePortRangeOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZonePortRangeOperPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZonePortRangeOperPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZonePortRangeOperPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZonePortRangeOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstZonePortRangeOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperProgressionTracking(d []interface{}) edpt.DdosDstZonePortRangeOperProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperProgressionTrackingOper(d []interface{}) edpt.DdosDstZonePortRangeOperProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortRangeOperProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZonePortRangeOperProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperProgressionTrackingOperIndicators
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

func getObjectDdosDstZonePortRangeOperTopkDestinations(d []interface{}) edpt.DdosDstZonePortRangeOperTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperTopkDestinationsOper(d []interface{}) edpt.DdosDstZonePortRangeOperTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperTopkSources(d []interface{}) edpt.DdosDstZonePortRangeOperTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZonePortRangeOperTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortRangeOperTopkSourcesOper(d []interface{}) edpt.DdosDstZonePortRangeOperTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeOperTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortRangeOperTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortRangeOperTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeOper(d *schema.ResourceData) edpt.DdosDstZonePortRangeOper {
	var ret edpt.DdosDstZonePortRangeOper

	ret.IpFilteringPolicyOper = getObjectDdosDstZonePortRangeOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Ips = getObjectDdosDstZonePortRangeOperIps(d.Get("ips").([]interface{}))

	ret.Oper = getObjectDdosDstZonePortRangeOperOper(d.Get("oper").([]interface{}))

	ret.PatternRecognition = getObjectDdosDstZonePortRangeOperPatternRecognition(d.Get("pattern_recognition").([]interface{}))

	ret.PatternRecognitionPuDetails = getObjectDdosDstZonePortRangeOperPatternRecognitionPuDetails(d.Get("pattern_recognition_pu_details").([]interface{}))

	ret.PortInd = getObjectDdosDstZonePortRangeOperPortInd(d.Get("port_ind").([]interface{}))

	ret.PortRangeEnd = d.Get("port_range_end").(int)

	ret.PortRangeStart = d.Get("port_range_start").(int)

	ret.ProgressionTracking = getObjectDdosDstZonePortRangeOperProgressionTracking(d.Get("progression_tracking").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.TopkDestinations = getObjectDdosDstZonePortRangeOperTopkDestinations(d.Get("topk_destinations").([]interface{}))

	ret.TopkSources = getObjectDdosDstZonePortRangeOperTopkSources(d.Get("topk_sources").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
