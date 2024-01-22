package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_oper`: Operational Status for the object entry\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryOperRead,

		Schema: map[string]*schema.Schema{
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"ip_proto_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Protocol Number",
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
					},
				},
			},
			"l4_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': L4-Type TCP; 'udp': L4-Type UDP; 'icmp': L4-Type ICMP; 'other': L4-Type OTHER;",
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
						"entry_address_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_dynamic_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"udp_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"udp_dynamic_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tcp_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tcp_dynamic_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"icmp_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"icmp_dynamic_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"other_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"other_dynamic_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"operational_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"traffic_distribution_status": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"master_pu": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"active_pu": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pu_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"dst_entry_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_entry_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_entry_alloc": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_entry_remain": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_service_limit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_service_alloc": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_service_remain": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"entry_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"no_t2_idx_port_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_all_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sources": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sources_all_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"overflow_policy": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sflow_source_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"l4_type_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"app_type": {
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
						"class_list": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_proto_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_range_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_range_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_port_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_port_range_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_port_range_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"opt_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sport_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"opt_sport_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"app_stat": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_app_stat": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_ip_protos": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_l4_types": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_src_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"black_holed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resource_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"display_traffic_distribution_status": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entry_status": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_ext_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_blacklisted": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port;",
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
					},
				},
			},
			"port_range_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_range_start": {
							Type: schema.TypeInt, Required: true, Description: "Port-Range Start Port Number",
						},
						"port_range_end": {
							Type: schema.TypeInt, Required: true, Description: "Port-Range End Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port;",
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
					},
				},
			},
			"src_port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns-udp': DNS-UDP Port; 'dns-tcp': DNS-TCP Port; 'udp': UDP Port; 'tcp': TCP Port;",
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
					},
				},
			},
			"src_port_range_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_port_range_start": {
							Type: schema.TypeInt, Required: true, Description: "Src Port-Range Start Port Number",
						},
						"src_port_range_end": {
							Type: schema.TypeInt, Required: true, Description: "Src Port-Range End Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'udp': UDP Port; 'tcp': TCP Port;",
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
					},
				},
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
		},
	}
}

func resourceDdosDstEntryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryOperIpProtoList := setSliceDdosDstEntryOperIpProtoList(res)
		d.Set("ip_proto_list", DdosDstEntryOperIpProtoList)
		DdosDstEntryOperL4TypeList := setSliceDdosDstEntryOperL4TypeList(res)
		d.Set("l4_type_list", DdosDstEntryOperL4TypeList)
		DdosDstEntryOperOper := setObjectDdosDstEntryOperOper(res)
		d.Set("oper", DdosDstEntryOperOper)
		DdosDstEntryOperPortList := setSliceDdosDstEntryOperPortList(res)
		d.Set("port_list", DdosDstEntryOperPortList)
		DdosDstEntryOperPortRangeList := setSliceDdosDstEntryOperPortRangeList(res)
		d.Set("port_range_list", DdosDstEntryOperPortRangeList)
		DdosDstEntryOperSrcPortList := setSliceDdosDstEntryOperSrcPortList(res)
		d.Set("src_port_list", DdosDstEntryOperSrcPortList)
		DdosDstEntryOperSrcPortRangeList := setSliceDdosDstEntryOperSrcPortRangeList(res)
		d.Set("src_port_range_list", DdosDstEntryOperSrcPortRangeList)
		DdosDstEntryOperTopkDestinations := setObjectDdosDstEntryOperTopkDestinations(res)
		d.Set("topk_destinations", DdosDstEntryOperTopkDestinations)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceDdosDstEntryOperIpProtoList(d edpt.DataDdosDstEntryOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstEntryOper.IpProtoList {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["oper"] = setObjectDdosDstEntryOperIpProtoListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperIpProtoListOper(d edpt.DdosDstEntryOperIpProtoListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstEntryOperIpProtoListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["all_ports"] = d.AllPorts

	in["all_src_ports"] = d.AllSrcPorts

	in["all_ip_protos"] = d.AllIpProtos

	in["port_protocol"] = d.PortProtocol

	in["app_stat"] = d.AppStat

	in["sflow_source_id"] = d.SflowSourceId

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperIpProtoListOperDdos_entry_list(d []edpt.DdosDstEntryOperIpProtoListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOper(d edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOperOper(d edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperL4TypeList(d edpt.DataDdosDstEntryOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstEntryOper.L4TypeList {
		in := make(map[string]interface{})
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstEntryOperL4TypeListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["port_ind"] = setObjectDdosDstEntryOperL4TypeListPortInd(item.PortInd)
		in["topk_sources"] = setObjectDdosDstEntryOperL4TypeListTopkSources(item.TopkSources)
		in["progression_tracking"] = setObjectDdosDstEntryOperL4TypeListProgressionTracking(item.ProgressionTracking)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperL4TypeListOper(d edpt.DdosDstEntryOperL4TypeListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstEntryOperL4TypeListOperDdos_entry_list(d.Ddos_entry_list)
	in["undefined_port_hit_stats_wellknown"] = setSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown(d.UndefinedPortHitStatsWellknown)
	in["undefined_port_hit_stats_non_wellknown"] = setSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown(d.UndefinedPortHitStatsNonWellknown)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["undefined_port_hit_statistics"] = d.UndefinedPortHitStatistics

	in["undefined_stats_port_num"] = d.UndefinedStatsPortNum

	in["all_l4_types"] = d.AllL4Types

	in["hw_blacklisted"] = d.HwBlacklisted
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperL4TypeListOperDdos_entry_list(d []edpt.DdosDstEntryOperL4TypeListOperDdos_entry_list) []map[string]interface{} {
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

func setSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown(d []edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port"] = item.Port
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown(d []edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown) []map[string]interface{} {
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

func setObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOper(d edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOperOper(d edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperL4TypeListPortInd(d edpt.DdosDstEntryOperL4TypeListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperL4TypeListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperL4TypeListPortIndOper(d edpt.DdosDstEntryOperL4TypeListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperL4TypeListPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperL4TypeListPortIndOperIndicators(d []edpt.DdosDstEntryOperL4TypeListPortIndOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperL4TypeListTopkSources(d edpt.DdosDstEntryOperL4TypeListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperL4TypeListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperL4TypeListTopkSourcesOper(d edpt.DdosDstEntryOperL4TypeListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicators(d []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryList(d []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperL4TypeListProgressionTracking(d edpt.DdosDstEntryOperL4TypeListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperL4TypeListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperL4TypeListProgressionTrackingOper(d edpt.DdosDstEntryOperL4TypeListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperL4TypeListProgressionTrackingOperIndicators(d.Indicators)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperL4TypeListProgressionTrackingOperIndicators(d []edpt.DdosDstEntryOperL4TypeListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperOper(ret edpt.DataDdosDstEntryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":                     setSliceDdosDstEntryOperOperDdos_entry_list(ret.DtDdosDstEntryOper.Oper.Ddos_entry_list),
			"entry_address_str":                   ret.DtDdosDstEntryOper.Oper.EntryAddressStr,
			"total_dynamic_entry_count":           ret.DtDdosDstEntryOper.Oper.TotalDynamicEntryCount,
			"total_dynamic_entry_limit":           ret.DtDdosDstEntryOper.Oper.TotalDynamicEntryLimit,
			"udp_dynamic_entry_count":             ret.DtDdosDstEntryOper.Oper.UdpDynamicEntryCount,
			"udp_dynamic_entry_limit":             ret.DtDdosDstEntryOper.Oper.UdpDynamicEntryLimit,
			"tcp_dynamic_entry_count":             ret.DtDdosDstEntryOper.Oper.TcpDynamicEntryCount,
			"tcp_dynamic_entry_limit":             ret.DtDdosDstEntryOper.Oper.TcpDynamicEntryLimit,
			"icmp_dynamic_entry_count":            ret.DtDdosDstEntryOper.Oper.IcmpDynamicEntryCount,
			"icmp_dynamic_entry_limit":            ret.DtDdosDstEntryOper.Oper.IcmpDynamicEntryLimit,
			"other_dynamic_entry_count":           ret.DtDdosDstEntryOper.Oper.OtherDynamicEntryCount,
			"other_dynamic_entry_limit":           ret.DtDdosDstEntryOper.Oper.OtherDynamicEntryLimit,
			"operational_mode":                    ret.DtDdosDstEntryOper.Oper.OperationalMode,
			"traffic_distribution_status":         setSliceDdosDstEntryOperOperTrafficDistributionStatus(ret.DtDdosDstEntryOper.Oper.TrafficDistributionStatus),
			"dst_entry_name":                      ret.DtDdosDstEntryOper.Oper.DstEntryName,
			"source_entry_limit":                  ret.DtDdosDstEntryOper.Oper.SourceEntryLimit,
			"source_entry_alloc":                  ret.DtDdosDstEntryOper.Oper.SourceEntryAlloc,
			"source_entry_remain":                 ret.DtDdosDstEntryOper.Oper.SourceEntryRemain,
			"dst_service_limit":                   ret.DtDdosDstEntryOper.Oper.DstServiceLimit,
			"dst_service_alloc":                   ret.DtDdosDstEntryOper.Oper.DstServiceAlloc,
			"dst_service_remain":                  ret.DtDdosDstEntryOper.Oper.DstServiceRemain,
			"entry_displayed_count":               ret.DtDdosDstEntryOper.Oper.EntryDisplayedCount,
			"service_displayed_count":             ret.DtDdosDstEntryOper.Oper.ServiceDisplayedCount,
			"no_t2_idx_port_count":                ret.DtDdosDstEntryOper.Oper.NoT2IdxPortCount,
			"dst_all_entries":                     ret.DtDdosDstEntryOper.Oper.DstAllEntries,
			"sources":                             ret.DtDdosDstEntryOper.Oper.Sources,
			"sources_all_entries":                 ret.DtDdosDstEntryOper.Oper.SourcesAllEntries,
			"overflow_policy":                     ret.DtDdosDstEntryOper.Oper.OverflowPolicy,
			"entry_count":                         ret.DtDdosDstEntryOper.Oper.EntryCount,
			"sflow_source_id":                     ret.DtDdosDstEntryOper.Oper.SflowSourceId,
			"ipv6":                                ret.DtDdosDstEntryOper.Oper.Ipv6,
			"subnet_ip_addr":                      ret.DtDdosDstEntryOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":                    ret.DtDdosDstEntryOper.Oper.SubnetIpv6Addr,
			"l4_type_str":                         ret.DtDdosDstEntryOper.Oper.L4TypeStr,
			"app_type":                            ret.DtDdosDstEntryOper.Oper.AppType,
			"exceeded":                            ret.DtDdosDstEntryOper.Oper.Exceeded,
			"black_listed":                        ret.DtDdosDstEntryOper.Oper.BlackListed,
			"white_listed":                        ret.DtDdosDstEntryOper.Oper.WhiteListed,
			"authenticated":                       ret.DtDdosDstEntryOper.Oper.Authenticated,
			"class_list":                          ret.DtDdosDstEntryOper.Oper.ClassList,
			"ip_proto_num":                        ret.DtDdosDstEntryOper.Oper.IpProtoNum,
			"port_num":                            ret.DtDdosDstEntryOper.Oper.PortNum,
			"port_range_start":                    ret.DtDdosDstEntryOper.Oper.PortRangeStart,
			"port_range_end":                      ret.DtDdosDstEntryOper.Oper.PortRangeEnd,
			"src_port_num":                        ret.DtDdosDstEntryOper.Oper.SrcPortNum,
			"src_port_range_start":                ret.DtDdosDstEntryOper.Oper.SrcPortRangeStart,
			"src_port_range_end":                  ret.DtDdosDstEntryOper.Oper.SrcPortRangeEnd,
			"protocol":                            ret.DtDdosDstEntryOper.Oper.Protocol,
			"opt_protocol":                        ret.DtDdosDstEntryOper.Oper.OptProtocol,
			"sport_protocol":                      ret.DtDdosDstEntryOper.Oper.SportProtocol,
			"opt_sport_protocol":                  ret.DtDdosDstEntryOper.Oper.OptSportProtocol,
			"app_stat":                            ret.DtDdosDstEntryOper.Oper.AppStat,
			"port_app_stat":                       ret.DtDdosDstEntryOper.Oper.PortAppStat,
			"all_ip_protos":                       ret.DtDdosDstEntryOper.Oper.AllIpProtos,
			"all_l4_types":                        ret.DtDdosDstEntryOper.Oper.AllL4Types,
			"all_ports":                           ret.DtDdosDstEntryOper.Oper.AllPorts,
			"all_src_ports":                       ret.DtDdosDstEntryOper.Oper.AllSrcPorts,
			"black_holed":                         ret.DtDdosDstEntryOper.Oper.BlackHoled,
			"resource_usage":                      ret.DtDdosDstEntryOper.Oper.ResourceUsage,
			"display_traffic_distribution_status": ret.DtDdosDstEntryOper.Oper.DisplayTrafficDistributionStatus,
			"entry_status":                        ret.DtDdosDstEntryOper.Oper.EntryStatus,
			"l4_ext_rate":                         ret.DtDdosDstEntryOper.Oper.L4ExtRate,
			"hw_blacklisted":                      ret.DtDdosDstEntryOper.Oper.HwBlacklisted,
		},
	}
}

func setSliceDdosDstEntryOperOperDdos_entry_list(d []edpt.DdosDstEntryOperOperDdos_entry_list) []map[string]interface{} {
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

func setSliceDdosDstEntryOperOperTrafficDistributionStatus(d []edpt.DdosDstEntryOperOperTrafficDistributionStatus) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["master_pu"] = item.MasterPu
		in["active_pu"] = setSliceDdosDstEntryOperOperTrafficDistributionStatusActivePu(item.ActivePu)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperOperTrafficDistributionStatusActivePu(d []edpt.DdosDstEntryOperOperTrafficDistributionStatusActivePu) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pu_id"] = item.PuId
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortList(d edpt.DataDdosDstEntryOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstEntryOper.PortList {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstEntryOperPortListOper(item.Oper)
		in["port_ind"] = setObjectDdosDstEntryOperPortListPortInd(item.PortInd)
		in["ip_filtering_policy_oper"] = setObjectDdosDstEntryOperPortListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["topk_sources"] = setObjectDdosDstEntryOperPortListTopkSources(item.TopkSources)
		in["progression_tracking"] = setObjectDdosDstEntryOperPortListProgressionTracking(item.ProgressionTracking)
		in["pattern_recognition"] = setObjectDdosDstEntryOperPortListPatternRecognition(item.PatternRecognition)
		in["pattern_recognition_pu_details"] = setObjectDdosDstEntryOperPortListPatternRecognitionPuDetails(item.PatternRecognitionPuDetails)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperPortListOper(d edpt.DdosDstEntryOperPortListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstEntryOperPortListOperDdos_entry_list(d.Ddos_entry_list)

	in["resource_limit_config"] = d.ResourceLimitConfig

	in["reource_limit_alloc"] = d.ReourceLimitAlloc

	in["resource_limit_remain"] = d.ResourceLimitRemain

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["all_ports"] = d.AllPorts

	in["all_src_ports"] = d.AllSrcPorts

	in["all_ip_protos"] = d.AllIpProtos

	in["port_protocol"] = d.PortProtocol

	in["app_stat"] = d.AppStat

	in["sflow_source_id"] = d.SflowSourceId

	in["resource_usage"] = d.ResourceUsage

	in["l4_ext_rate"] = d.L4ExtRate

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortListOperDdos_entry_list(d []edpt.DdosDstEntryOperPortListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortListPortInd(d edpt.DdosDstEntryOperPortListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortListPortIndOper(d edpt.DdosDstEntryOperPortListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperPortListPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortListPortIndOperIndicators(d []edpt.DdosDstEntryOperPortListPortIndOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortListIpFilteringPolicyOper(d edpt.DdosDstEntryOperPortListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortListIpFilteringPolicyOperOper(d edpt.DdosDstEntryOperPortListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperPortListTopkSources(d edpt.DdosDstEntryOperPortListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortListTopkSourcesOper(d edpt.DdosDstEntryOperPortListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperPortListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstEntryOperPortListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortListTopkSourcesOperIndicators(d []edpt.DdosDstEntryOperPortListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstEntryOperPortListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstEntryOperPortListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortListTopkSourcesOperEntryList(d []edpt.DdosDstEntryOperPortListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryOperPortListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstEntryOperPortListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortListProgressionTracking(d edpt.DdosDstEntryOperPortListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortListProgressionTrackingOper(d edpt.DdosDstEntryOperPortListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperPortListProgressionTrackingOperIndicators(d.Indicators)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortListProgressionTrackingOperIndicators(d []edpt.DdosDstEntryOperPortListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortListPatternRecognition(d edpt.DdosDstEntryOperPortListPatternRecognition) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortListPatternRecognitionOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortListPatternRecognitionOper(d edpt.DdosDstEntryOperPortListPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstEntryOperPortListPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortListPatternRecognitionOperFilterList(d []edpt.DdosDstEntryOperPortListPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortListPatternRecognitionPuDetails(d edpt.DdosDstEntryOperPortListPatternRecognitionPuDetails) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortListPatternRecognitionPuDetailsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortListPatternRecognitionPuDetailsOper(d edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setSliceDdosDstEntryOperPortRangeList(d edpt.DataDdosDstEntryOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstEntryOper.PortRangeList {
		in := make(map[string]interface{})
		in["port_range_start"] = item.PortRangeStart
		in["port_range_end"] = item.PortRangeEnd
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstEntryOperPortRangeListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["port_ind"] = setObjectDdosDstEntryOperPortRangeListPortInd(item.PortInd)
		in["topk_sources"] = setObjectDdosDstEntryOperPortRangeListTopkSources(item.TopkSources)
		in["progression_tracking"] = setObjectDdosDstEntryOperPortRangeListProgressionTracking(item.ProgressionTracking)
		in["pattern_recognition"] = setObjectDdosDstEntryOperPortRangeListPatternRecognition(item.PatternRecognition)
		in["pattern_recognition_pu_details"] = setObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetails(item.PatternRecognitionPuDetails)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperPortRangeListOper(d edpt.DdosDstEntryOperPortRangeListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstEntryOperPortRangeListOperDdos_entry_list(d.Ddos_entry_list)

	in["resource_limit_config"] = d.ResourceLimitConfig

	in["reource_limit_alloc"] = d.ReourceLimitAlloc

	in["resource_limit_remain"] = d.ResourceLimitRemain

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["all_ports"] = d.AllPorts

	in["all_src_ports"] = d.AllSrcPorts

	in["all_ip_protos"] = d.AllIpProtos

	in["port_protocol"] = d.PortProtocol

	in["app_stat"] = d.AppStat

	in["sflow_source_id"] = d.SflowSourceId

	in["resource_usage"] = d.ResourceUsage

	in["l4_ext_rate"] = d.L4ExtRate

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortRangeListOperDdos_entry_list(d []edpt.DdosDstEntryOperPortRangeListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOper(d edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOperOper(d edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperPortRangeListPortInd(d edpt.DdosDstEntryOperPortRangeListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortRangeListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortRangeListPortIndOper(d edpt.DdosDstEntryOperPortRangeListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperPortRangeListPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortRangeListPortIndOperIndicators(d []edpt.DdosDstEntryOperPortRangeListPortIndOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortRangeListTopkSources(d edpt.DdosDstEntryOperPortRangeListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortRangeListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortRangeListTopkSourcesOper(d edpt.DdosDstEntryOperPortRangeListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicators(d []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryList(d []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortRangeListProgressionTracking(d edpt.DdosDstEntryOperPortRangeListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortRangeListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortRangeListProgressionTrackingOper(d edpt.DdosDstEntryOperPortRangeListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperPortRangeListProgressionTrackingOperIndicators(d.Indicators)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortRangeListProgressionTrackingOperIndicators(d []edpt.DdosDstEntryOperPortRangeListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortRangeListPatternRecognition(d edpt.DdosDstEntryOperPortRangeListPatternRecognition) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortRangeListPatternRecognitionOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortRangeListPatternRecognitionOper(d edpt.DdosDstEntryOperPortRangeListPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstEntryOperPortRangeListPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortRangeListPatternRecognitionOperFilterList(d []edpt.DdosDstEntryOperPortRangeListPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetails(d edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetails) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper(d edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setSliceDdosDstEntryOperSrcPortList(d edpt.DataDdosDstEntryOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstEntryOper.SrcPortList {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstEntryOperSrcPortListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperSrcPortListOper(d edpt.DdosDstEntryOperSrcPortListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstEntryOperSrcPortListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["all_ports"] = d.AllPorts

	in["all_src_ports"] = d.AllSrcPorts

	in["all_ip_protos"] = d.AllIpProtos

	in["port_protocol"] = d.PortProtocol

	in["app_stat"] = d.AppStat

	in["sflow_source_id"] = d.SflowSourceId

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperSrcPortListOperDdos_entry_list(d []edpt.DdosDstEntryOperSrcPortListOperDdos_entry_list) []map[string]interface{} {
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

func setSliceDdosDstEntryOperSrcPortRangeList(d edpt.DataDdosDstEntryOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstEntryOper.SrcPortRangeList {
		in := make(map[string]interface{})
		in["src_port_range_start"] = item.SrcPortRangeStart
		in["src_port_range_end"] = item.SrcPortRangeEnd
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstEntryOperSrcPortRangeListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryOperSrcPortRangeListOper(d edpt.DdosDstEntryOperSrcPortRangeListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstEntryOperSrcPortRangeListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["all_ports"] = d.AllPorts

	in["all_src_ports"] = d.AllSrcPorts

	in["all_ip_protos"] = d.AllIpProtos

	in["port_protocol"] = d.PortProtocol

	in["app_stat"] = d.AppStat

	in["sflow_source_id"] = d.SflowSourceId

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperSrcPortRangeListOperDdos_entry_list(d []edpt.DdosDstEntryOperSrcPortRangeListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstEntryOperTopkDestinations(ret edpt.DataDdosDstEntryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryOperTopkDestinationsOper(ret.DtDdosDstEntryOper.TopkDestinations.Oper),
		},
	}
}

func setObjectDdosDstEntryOperTopkDestinationsOper(d edpt.DdosDstEntryOperTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstEntryOperTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstEntryOperTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryOperTopkDestinationsOperIndicators(d []edpt.DdosDstEntryOperTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstEntryOperTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstEntryOperTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperTopkDestinationsOperEntryList(d []edpt.DdosDstEntryOperTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryOperTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryOperTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstEntryOperTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func getSliceDdosDstEntryOperIpProtoList(d []interface{}) []edpt.DdosDstEntryOperIpProtoList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperIpProtoList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperIpProtoList
		oi.PortNum = in["port_num"].(int)
		oi.Oper = getObjectDdosDstEntryOperIpProtoListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperIpProtoListOper(d []interface{}) edpt.DdosDstEntryOperIpProtoListOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperIpProtoListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryOperIpProtoListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.AllPorts = in["all_ports"].(int)
		ret.AllSrcPorts = in["all_src_ports"].(int)
		ret.AllIpProtos = in["all_ip_protos"].(int)
		ret.PortProtocol = in["port_protocol"].(string)
		ret.AppStat = in["app_stat"].(int)
		ret.SflowSourceId = in["sflow_source_id"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperIpProtoListOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryOperIpProtoListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperIpProtoListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperIpProtoListOperDdos_entry_list
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

func getObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperIpProtoListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeList(d []interface{}) []edpt.DdosDstEntryOperL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstEntryOperL4TypeListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstEntryOperL4TypeListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstEntryOperL4TypeListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstEntryOperL4TypeListProgressionTracking(in["progression_tracking"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListOper(d []interface{}) edpt.DdosDstEntryOperL4TypeListOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryOperL4TypeListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.UndefinedPortHitStatsWellknown = getSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown(in["undefined_port_hit_stats_wellknown"].([]interface{}))
		ret.UndefinedPortHitStatsNonWellknown = getSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown(in["undefined_port_hit_stats_non_wellknown"].([]interface{}))
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

func getSliceDdosDstEntryOperL4TypeListOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryOperL4TypeListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListOperDdos_entry_list
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

func getSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown(d []interface{}) []edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown
		oi.Port = in["port"].(string)
		oi.Counter = in["counter"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown(d []interface{}) []edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown
		oi.PortStart = in["port_start"].(string)
		oi.PortEnd = in["port_end"].(string)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListPortInd(d []interface{}) edpt.DdosDstEntryOperL4TypeListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperL4TypeListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListPortIndOper(d []interface{}) edpt.DdosDstEntryOperL4TypeListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperL4TypeListPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListPortIndOperIndicators(d []interface{}) []edpt.DdosDstEntryOperL4TypeListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListPortIndOperIndicators
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

func getObjectDdosDstEntryOperL4TypeListTopkSources(d []interface{}) edpt.DdosDstEntryOperL4TypeListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperL4TypeListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListTopkSourcesOper(d []interface{}) edpt.DdosDstEntryOperL4TypeListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListProgressionTracking(d []interface{}) edpt.DdosDstEntryOperL4TypeListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperL4TypeListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperL4TypeListProgressionTrackingOper(d []interface{}) edpt.DdosDstEntryOperL4TypeListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperL4TypeListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperL4TypeListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperL4TypeListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstEntryOperL4TypeListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperL4TypeListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperL4TypeListProgressionTrackingOperIndicators
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

func getObjectDdosDstEntryOperOper(d []interface{}) edpt.DdosDstEntryOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryAddressStr = in["entry_address_str"].(string)
		ret.TotalDynamicEntryCount = in["total_dynamic_entry_count"].(string)
		ret.TotalDynamicEntryLimit = in["total_dynamic_entry_limit"].(string)
		ret.UdpDynamicEntryCount = in["udp_dynamic_entry_count"].(string)
		ret.UdpDynamicEntryLimit = in["udp_dynamic_entry_limit"].(string)
		ret.TcpDynamicEntryCount = in["tcp_dynamic_entry_count"].(string)
		ret.TcpDynamicEntryLimit = in["tcp_dynamic_entry_limit"].(string)
		ret.IcmpDynamicEntryCount = in["icmp_dynamic_entry_count"].(string)
		ret.IcmpDynamicEntryLimit = in["icmp_dynamic_entry_limit"].(string)
		ret.OtherDynamicEntryCount = in["other_dynamic_entry_count"].(string)
		ret.OtherDynamicEntryLimit = in["other_dynamic_entry_limit"].(string)
		ret.OperationalMode = in["operational_mode"].(string)
		ret.TrafficDistributionStatus = getSliceDdosDstEntryOperOperTrafficDistributionStatus(in["traffic_distribution_status"].([]interface{}))
		ret.DstEntryName = in["dst_entry_name"].(string)
		ret.SourceEntryLimit = in["source_entry_limit"].(string)
		ret.SourceEntryAlloc = in["source_entry_alloc"].(string)
		ret.SourceEntryRemain = in["source_entry_remain"].(string)
		ret.DstServiceLimit = in["dst_service_limit"].(string)
		ret.DstServiceAlloc = in["dst_service_alloc"].(string)
		ret.DstServiceRemain = in["dst_service_remain"].(string)
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.NoT2IdxPortCount = in["no_t2_idx_port_count"].(int)
		ret.DstAllEntries = in["dst_all_entries"].(int)
		ret.Sources = in["sources"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.OverflowPolicy = in["overflow_policy"].(int)
		ret.EntryCount = in["entry_count"].(int)
		ret.SflowSourceId = in["sflow_source_id"].(int)
		ret.Ipv6 = in["ipv6"].(string)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.L4TypeStr = in["l4_type_str"].(string)
		ret.AppType = in["app_type"].(string)
		ret.Exceeded = in["exceeded"].(int)
		ret.BlackListed = in["black_listed"].(int)
		ret.WhiteListed = in["white_listed"].(int)
		ret.Authenticated = in["authenticated"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.IpProtoNum = in["ip_proto_num"].(int)
		ret.PortNum = in["port_num"].(int)
		ret.PortRangeStart = in["port_range_start"].(int)
		ret.PortRangeEnd = in["port_range_end"].(int)
		ret.SrcPortNum = in["src_port_num"].(int)
		ret.SrcPortRangeStart = in["src_port_range_start"].(int)
		ret.SrcPortRangeEnd = in["src_port_range_end"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.OptProtocol = in["opt_protocol"].(string)
		ret.SportProtocol = in["sport_protocol"].(string)
		ret.OptSportProtocol = in["opt_sport_protocol"].(string)
		ret.AppStat = in["app_stat"].(int)
		ret.PortAppStat = in["port_app_stat"].(int)
		ret.AllIpProtos = in["all_ip_protos"].(int)
		ret.AllL4Types = in["all_l4_types"].(int)
		ret.AllPorts = in["all_ports"].(int)
		ret.AllSrcPorts = in["all_src_ports"].(int)
		ret.BlackHoled = in["black_holed"].(int)
		ret.ResourceUsage = in["resource_usage"].(int)
		ret.DisplayTrafficDistributionStatus = in["display_traffic_distribution_status"].(int)
		ret.EntryStatus = in["entry_status"].(int)
		ret.L4ExtRate = in["l4_ext_rate"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperOperDdos_entry_list
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

func getSliceDdosDstEntryOperOperTrafficDistributionStatus(d []interface{}) []edpt.DdosDstEntryOperOperTrafficDistributionStatus {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperOperTrafficDistributionStatus, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperOperTrafficDistributionStatus
		oi.MasterPu = in["master_pu"].(string)
		oi.ActivePu = getSliceDdosDstEntryOperOperTrafficDistributionStatusActivePu(in["active_pu"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperOperTrafficDistributionStatusActivePu(d []interface{}) []edpt.DdosDstEntryOperOperTrafficDistributionStatusActivePu {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperOperTrafficDistributionStatusActivePu, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperOperTrafficDistributionStatusActivePu
		oi.PuId = in["pu_id"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortList(d []interface{}) []edpt.DdosDstEntryOperPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstEntryOperPortListOper(in["oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstEntryOperPortListPortInd(in["port_ind"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryOperPortListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.TopkSources = getObjectDdosDstEntryOperPortListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstEntryOperPortListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstEntryOperPortListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstEntryOperPortListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperPortListOper(d []interface{}) edpt.DdosDstEntryOperPortListOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryOperPortListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstEntryOperPortListOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryOperPortListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListOperDdos_entry_list
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

func getObjectDdosDstEntryOperPortListPortInd(d []interface{}) edpt.DdosDstEntryOperPortListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortListPortIndOper(d []interface{}) edpt.DdosDstEntryOperPortListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperPortListPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperPortListPortIndOperIndicators(d []interface{}) []edpt.DdosDstEntryOperPortListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListPortIndOperIndicators
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

func getObjectDdosDstEntryOperPortListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryOperPortListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryOperPortListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperPortListTopkSources(d []interface{}) edpt.DdosDstEntryOperPortListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortListTopkSourcesOper(d []interface{}) edpt.DdosDstEntryOperPortListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperPortListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryOperPortListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperPortListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstEntryOperPortListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstEntryOperPortListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstEntryOperPortListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstEntryOperPortListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryOperPortListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryOperPortListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperPortListProgressionTracking(d []interface{}) edpt.DdosDstEntryOperPortListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortListProgressionTrackingOper(d []interface{}) edpt.DdosDstEntryOperPortListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperPortListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstEntryOperPortListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListProgressionTrackingOperIndicators
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

func getObjectDdosDstEntryOperPortListPatternRecognition(d []interface{}) edpt.DdosDstEntryOperPortListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortListPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortListPatternRecognitionOper(d []interface{}) edpt.DdosDstEntryOperPortListPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstEntryOperPortListPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortListPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstEntryOperPortListPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListPatternRecognitionOperFilterList
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

func getObjectDdosDstEntryOperPortListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstEntryOperPortListPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortListPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortListPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeList(d []interface{}) []edpt.DdosDstEntryOperPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeList
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstEntryOperPortRangeListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstEntryOperPortRangeListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstEntryOperPortRangeListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstEntryOperPortRangeListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstEntryOperPortRangeListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryOperPortRangeListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstEntryOperPortRangeListOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryOperPortRangeListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListOperDdos_entry_list
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

func getObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListPortInd(d []interface{}) edpt.DdosDstEntryOperPortRangeListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortRangeListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListPortIndOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperPortRangeListPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListPortIndOperIndicators(d []interface{}) []edpt.DdosDstEntryOperPortRangeListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListPortIndOperIndicators
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

func getObjectDdosDstEntryOperPortRangeListTopkSources(d []interface{}) edpt.DdosDstEntryOperPortRangeListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortRangeListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListTopkSourcesOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListProgressionTracking(d []interface{}) edpt.DdosDstEntryOperPortRangeListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortRangeListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListProgressionTrackingOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperPortRangeListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstEntryOperPortRangeListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListProgressionTrackingOperIndicators
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

func getObjectDdosDstEntryOperPortRangeListPatternRecognition(d []interface{}) edpt.DdosDstEntryOperPortRangeListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortRangeListPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListPatternRecognitionOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstEntryOperPortRangeListPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstEntryOperPortRangeListPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListPatternRecognitionOperFilterList
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

func getObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperSrcPortList(d []interface{}) []edpt.DdosDstEntryOperSrcPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperSrcPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperSrcPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstEntryOperSrcPortListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperSrcPortListOper(d []interface{}) edpt.DdosDstEntryOperSrcPortListOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperSrcPortListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryOperSrcPortListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.AllPorts = in["all_ports"].(int)
		ret.AllSrcPorts = in["all_src_ports"].(int)
		ret.AllIpProtos = in["all_ip_protos"].(int)
		ret.PortProtocol = in["port_protocol"].(string)
		ret.AppStat = in["app_stat"].(int)
		ret.SflowSourceId = in["sflow_source_id"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperSrcPortListOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryOperSrcPortListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperSrcPortListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperSrcPortListOperDdos_entry_list
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

func getSliceDdosDstEntryOperSrcPortRangeList(d []interface{}) []edpt.DdosDstEntryOperSrcPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperSrcPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperSrcPortRangeList
		oi.SrcPortRangeStart = in["src_port_range_start"].(int)
		oi.SrcPortRangeEnd = in["src_port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstEntryOperSrcPortRangeListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryOperSrcPortRangeListOper(d []interface{}) edpt.DdosDstEntryOperSrcPortRangeListOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperSrcPortRangeListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryOperSrcPortRangeListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.AllPorts = in["all_ports"].(int)
		ret.AllSrcPorts = in["all_src_ports"].(int)
		ret.AllIpProtos = in["all_ip_protos"].(int)
		ret.PortProtocol = in["port_protocol"].(string)
		ret.AppStat = in["app_stat"].(int)
		ret.SflowSourceId = in["sflow_source_id"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperSrcPortRangeListOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryOperSrcPortRangeListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperSrcPortRangeListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperSrcPortRangeListOperDdos_entry_list
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

func getObjectDdosDstEntryOperTopkDestinations(d []interface{}) edpt.DdosDstEntryOperTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryOperTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryOperTopkDestinationsOper(d []interface{}) edpt.DdosDstEntryOperTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryOperTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryOperTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryOperTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryOperTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstEntryOperTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstEntryOperTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstEntryOperTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstEntryOperTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryOperTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryOperTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryOperTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryOperTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryOperTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryOper(d *schema.ResourceData) edpt.DdosDstEntryOper {
	var ret edpt.DdosDstEntryOper

	ret.DstEntryName = d.Get("dst_entry_name").(string)

	ret.IpProtoList = getSliceDdosDstEntryOperIpProtoList(d.Get("ip_proto_list").([]interface{}))

	ret.L4TypeList = getSliceDdosDstEntryOperL4TypeList(d.Get("l4_type_list").([]interface{}))

	ret.Oper = getObjectDdosDstEntryOperOper(d.Get("oper").([]interface{}))

	ret.PortList = getSliceDdosDstEntryOperPortList(d.Get("port_list").([]interface{}))

	ret.PortRangeList = getSliceDdosDstEntryOperPortRangeList(d.Get("port_range_list").([]interface{}))

	ret.SrcPortList = getSliceDdosDstEntryOperSrcPortList(d.Get("src_port_list").([]interface{}))

	ret.SrcPortRangeList = getSliceDdosDstEntryOperSrcPortRangeList(d.Get("src_port_range_list").([]interface{}))

	ret.TopkDestinations = getObjectDdosDstEntryOperTopkDestinations(d.Get("topk_destinations").([]interface{}))
	return ret
}
