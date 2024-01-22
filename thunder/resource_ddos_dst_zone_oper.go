package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_oper`: Operational Status for the object zone\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneOperRead,

		Schema: map[string]*schema.Schema{
			"detection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"outbound_detection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"discovery_timestamp": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"entry_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"location_type": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"location_name": {
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
																		"maximum": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"minimum": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"non_zero_minimum": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"average": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"adaptive_threshold": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																	},
																},
															},
															"data_source": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"anomaly": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"anomaly_timestamp": {
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
											},
										},
									},
									"topk_source_subnet": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"oper": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"entry_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"location_type": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"location_name": {
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
																					"source_subnets": {
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
						"service_discovery": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"discovered_service_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"protocol": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"rate": {
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
						"victim_ip_detection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_entry_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_address_str": {
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
																		"value": {
																			Type: schema.TypeList, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"current": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"threshold": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																				},
																			},
																		},
																		"is_anomaly": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"is_learning_done": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"is_histogram_learning_done": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"is_ip_anomaly": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"is_static_threshold": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"escalation_timestamp": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"de_escalation_timestamp": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"ip_entry_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"total_ip_entry_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"active_list": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"victim_list": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ipv4_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv6_ip": {
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
			"ip_proto": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"proto_number_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol_num": {
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
								},
							},
						},
						"proto_tcp_udp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': ip-proto tcp; 'udp': ip-proto udp;",
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
						"proto_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6; 'other': ip-proto other; 'gre': ip-proto gre; 'ipv4-encap': ip-proto IPv4 Encapsulation; 'ipv6-encap': ip-proto IPv6 Encapsulation;",
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
									"port_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"operational_mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bw_state": {
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
									"dynamic_entry_count": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lockup_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
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
						"total_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"udp_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tcp_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"icmp_dynamic_entry_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"other_dynamic_entry_count": {
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
						"entry_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"no_t2_idx_port_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"addresses": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"all_addresses": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip_proto_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_ip_protos": {
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
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"all_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dynamic_expand_subnet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"blackhole": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"outbound_policy": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"policy_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"no_class_list_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"policy_class_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list_name": {
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
												"current_frag_packet_rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"is_frag_packet_rate_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"frag_packet_rate_limit": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"age_str": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lockup_time": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packet_received": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packet_dropped": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packet_rate_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"kbit_rate_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"kbit_rate_exceed_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"connections_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"connection_rate_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"frag_packet_rate": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"geo_tracking_statistics": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"packet_received": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packet_dropped": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packet_rate_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"kbit_rate_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"kbit_rate_exceed_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"connections_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"connection_rate_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"frag_packet_rate": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"tracking_entry_learn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"tracking_entry_aged": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"tracking_entry_learning_thre_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"tracking_entry_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"geo_location_name": {
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
											},
										},
									},
									"policy_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"policy_statistics": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tracking_entry_filter": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"packet_anomaly_detection": {
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
												"maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"threshold": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"is_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"data_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"zone_service_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_num": {
										Type: schema.TypeInt, Required: true, Description: "Port Number",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port; 'quic': QUIC Port;",
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
							},
						},
						"zone_service_other_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_other": {
										Type: schema.TypeString, Required: true, Description: "'other': other;",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
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
							Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port; 'quic': QUIC Port;",
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
					},
				},
			},
			"src_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"zone_src_port_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_num": {
										Type: schema.TypeInt, Required: true, Description: "Source Port Number",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'dns-udp': DNS-UDP Port; 'dns-tcp': DNS-TCP Port; 'udp': UDP port; 'tcp': TCP Port;",
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
												"hw_blacklisted": {
													Type: schema.TypeInt, Optional: true, Description: "",
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
																	},
																},
															},
															"detection_data_source": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"current_level": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"details": {
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
						"zone_src_port_other_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_other": {
										Type: schema.TypeString, Required: true, Description: "'other': other;",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
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
												"hw_blacklisted": {
													Type: schema.TypeInt, Optional: true, Description: "",
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
																	},
																},
															},
															"detection_data_source": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"current_level": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"details": {
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
							Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
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
									"hw_blacklisted": {
										Type: schema.TypeInt, Optional: true, Description: "",
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
														},
													},
												},
												"detection_data_source": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"current_level": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"details": {
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
		},
	}
}

func resourceDdosDstZoneOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneOperDetection := setObjectDdosDstZoneOperDetection(res)
		d.Set("detection", DdosDstZoneOperDetection)
		DdosDstZoneOperIpProto := setObjectDdosDstZoneOperIpProto(res)
		d.Set("ip_proto", DdosDstZoneOperIpProto)
		DdosDstZoneOperOper := setObjectDdosDstZoneOperOper(res)
		d.Set("oper", DdosDstZoneOperOper)
		DdosDstZoneOperOutboundPolicy := setObjectDdosDstZoneOperOutboundPolicy(res)
		d.Set("outbound_policy", DdosDstZoneOperOutboundPolicy)
		DdosDstZoneOperPacketAnomalyDetection := setObjectDdosDstZoneOperPacketAnomalyDetection(res)
		d.Set("packet_anomaly_detection", DdosDstZoneOperPacketAnomalyDetection)
		DdosDstZoneOperPort := setObjectDdosDstZoneOperPort(res)
		d.Set("port", DdosDstZoneOperPort)
		DdosDstZoneOperPortRangeList := setSliceDdosDstZoneOperPortRangeList(res)
		d.Set("port_range_list", DdosDstZoneOperPortRangeList)
		DdosDstZoneOperSrcPort := setObjectDdosDstZoneOperSrcPort(res)
		d.Set("src_port", DdosDstZoneOperSrcPort)
		DdosDstZoneOperSrcPortRangeList := setSliceDdosDstZoneOperSrcPortRangeList(res)
		d.Set("src_port_range_list", DdosDstZoneOperSrcPortRangeList)
		DdosDstZoneOperTopkDestinations := setObjectDdosDstZoneOperTopkDestinations(res)
		d.Set("topk_destinations", DdosDstZoneOperTopkDestinations)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneOperDetection(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":                setObjectDdosDstZoneOperDetectionOper(ret.DtDdosDstZoneOper.Detection.Oper),
			"outbound_detection":  setObjectDdosDstZoneOperDetectionOutboundDetection(ret.DtDdosDstZoneOper.Detection.OutboundDetection),
			"service_discovery":   setObjectDdosDstZoneOperDetectionServiceDiscovery(ret.DtDdosDstZoneOper.Detection.ServiceDiscovery),
			"victim_ip_detection": setObjectDdosDstZoneOperDetectionVictimIpDetection(ret.DtDdosDstZoneOper.Detection.VictimIpDetection),
		},
	}
}

func setObjectDdosDstZoneOperDetectionOper(d edpt.DdosDstZoneOperDetectionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperDetectionOutboundDetection(d edpt.DdosDstZoneOperDetectionOutboundDetection) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperDetectionOutboundDetectionOper(d.Oper)
	in["topk_source_subnet"] = setObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet(d.TopkSourceSubnet)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperDetectionOutboundDetectionOper(d edpt.DdosDstZoneOperDetectionOutboundDetectionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["discovery_timestamp"] = d.DiscoveryTimestamp
	in["entry_list"] = setSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryList(d.EntryList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryList(d []edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["location_type"] = item.LocationType
		in["location_name"] = item.LocationName
		in["indicators"] = setSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators(item.Indicators)
		in["data_source"] = item.DataSource
		in["anomaly"] = item.Anomaly
		in["anomaly_timestamp"] = item.AnomalyTimestamp
		in["initial_learning"] = item.InitialLearning
		in["active_time"] = item.ActiveTime
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators(d []edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["maximum"] = item.Maximum
		in["minimum"] = item.Minimum
		in["non_zero_minimum"] = item.NonZeroMinimum
		in["average"] = item.Average
		in["adaptive_threshold"] = item.AdaptiveThreshold
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet(d edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper(d edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["entry_list"] = setSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList(d.EntryList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList(d []edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["location_type"] = item.LocationType
		in["location_name"] = item.LocationName
		in["indicators"] = setSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators(d []edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["source_subnets"] = setSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(item.SourceSubnets)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(d []edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperDetectionServiceDiscovery(d edpt.DdosDstZoneOperDetectionServiceDiscovery) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperDetectionServiceDiscoveryOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperDetectionServiceDiscoveryOper(d edpt.DdosDstZoneOperDetectionServiceDiscoveryOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["discovered_service_list"] = setSliceDdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList(d.DiscoveredServiceList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList(d []edpt.DdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port"] = item.Port
		in["protocol"] = item.Protocol
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperDetectionVictimIpDetection(d edpt.DdosDstZoneOperDetectionVictimIpDetection) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperDetectionVictimIpDetectionOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperDetectionVictimIpDetectionOper(d edpt.DdosDstZoneOperDetectionVictimIpDetectionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ip_entry_list"] = setSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList(d.IpEntryList)

	in["ip_entry_count"] = d.IpEntryCount

	in["total_ip_entry_count"] = d.TotalIpEntryCount

	in["active_list"] = d.ActiveList

	in["victim_list"] = d.VictimList

	in["ipv4_ip"] = d.Ipv4Ip

	in["ipv6_ip"] = d.Ipv6Ip
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList(d []edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_address_str"] = item.IpAddressStr
		in["indicators"] = setSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators(item.Indicators)
		in["is_learning_done"] = item.IsLearningDone
		in["is_histogram_learning_done"] = item.IsHistogramLearningDone
		in["is_ip_anomaly"] = item.IsIpAnomaly
		in["is_static_threshold"] = item.Is_static_threshold
		in["escalation_timestamp"] = item.EscalationTimestamp
		in["de_escalation_timestamp"] = item.DeEscalationTimestamp
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators(d []edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["value"] = setSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue(item.Value)
		in["is_anomaly"] = item.IsAnomaly
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue(d []edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["current"] = item.Current
		in["threshold"] = item.Threshold
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperIpProto(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":               setObjectDdosDstZoneOperIpProtoOper(ret.DtDdosDstZoneOper.IpProto.Oper),
			"proto_number_list":  setSliceDdosDstZoneOperIpProtoProtoNumberList(ret.DtDdosDstZoneOper.IpProto.ProtoNumberList),
			"proto_tcp_udp_list": setSliceDdosDstZoneOperIpProtoProtoTcpUdpList(ret.DtDdosDstZoneOper.IpProto.ProtoTcpUdpList),
			"proto_name_list":    setSliceDdosDstZoneOperIpProtoProtoNameList(ret.DtDdosDstZoneOper.IpProto.ProtoNameList),
		},
	}
}

func setObjectDdosDstZoneOperIpProtoOper(d edpt.DdosDstZoneOperIpProtoOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberList(d []edpt.DdosDstZoneOperIpProtoProtoNumberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol_num"] = item.ProtocolNum
		in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNumberListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["port_ind"] = setObjectDdosDstZoneOperIpProtoProtoNumberListPortInd(item.PortInd)
		in["topk_sources"] = setObjectDdosDstZoneOperIpProtoProtoNumberListTopkSources(item.TopkSources)
		in["topk_destinations"] = setObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinations(item.TopkDestinations)
		in["progression_tracking"] = setObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTracking(item.ProgressionTracking)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNumberListOper(d edpt.DdosDstZoneOperIpProtoProtoNumberListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["overflow_policy"] = d.OverflowPolicy

	in["sources_all_entries"] = d.SourcesAllEntries

	in["class_list"] = d.ClassList

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["exceeded"] = d.Exceeded

	in["black_listed"] = d.BlackListed

	in["white_listed"] = d.WhiteListed

	in["authenticated"] = d.Authenticated

	in["level"] = d.Level

	in["app_stat"] = d.AppStat

	in["indicators"] = d.Indicators

	in["indicator_detail"] = d.IndicatorDetail

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list(d []edpt.DdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOper(d edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOper(d edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNumberListPortInd(d edpt.DdosDstZoneOperIpProtoProtoNumberListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNumberListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNumberListPortIndOper(d edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList(d []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNumberListTopkSources(d edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper(d edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinations(d edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper(d edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTracking(d edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper(d edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperIpProtoProtoTcpUdpList(d []edpt.DdosDstZoneOperIpProtoProtoTcpUdpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperIpProtoProtoTcpUdpListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoTcpUdpListOper(d edpt.DdosDstZoneOperIpProtoProtoTcpUdpListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["overflow_policy"] = d.OverflowPolicy

	in["sources_all_entries"] = d.SourcesAllEntries

	in["class_list"] = d.ClassList

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["exceeded"] = d.Exceeded

	in["black_listed"] = d.BlackListed

	in["white_listed"] = d.WhiteListed

	in["authenticated"] = d.Authenticated

	in["level"] = d.Level

	in["app_stat"] = d.AppStat

	in["indicators"] = d.Indicators

	in["indicator_detail"] = d.IndicatorDetail

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list(d []edpt.DdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOper(d edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOper(d edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameList(d []edpt.DdosDstZoneOperIpProtoProtoNameList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNameListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["port_ind"] = setObjectDdosDstZoneOperIpProtoProtoNameListPortInd(item.PortInd)
		in["topk_sources"] = setObjectDdosDstZoneOperIpProtoProtoNameListTopkSources(item.TopkSources)
		in["progression_tracking"] = setObjectDdosDstZoneOperIpProtoProtoNameListProgressionTracking(item.ProgressionTracking)
		in["topk_destinations"] = setObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinations(item.TopkDestinations)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNameListOper(d edpt.DdosDstZoneOperIpProtoProtoNameListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["overflow_policy"] = d.OverflowPolicy

	in["sources_all_entries"] = d.SourcesAllEntries

	in["class_list"] = d.ClassList

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["exceeded"] = d.Exceeded

	in["black_listed"] = d.BlackListed

	in["white_listed"] = d.WhiteListed

	in["authenticated"] = d.Authenticated

	in["level"] = d.Level

	in["app_stat"] = d.AppStat

	in["indicators"] = d.Indicators

	in["indicator_detail"] = d.IndicatorDetail

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list(d []edpt.DdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOper(d edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOper(d edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNameListPortInd(d edpt.DdosDstZoneOperIpProtoProtoNameListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNameListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNameListPortIndOper(d edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList(d []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNameListTopkSources(d edpt.DdosDstZoneOperIpProtoProtoNameListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNameListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNameListTopkSourcesOper(d edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNameListProgressionTracking(d edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper(d edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinations(d edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper(d edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperOper(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":             setSliceDdosDstZoneOperOperDdos_entry_list(ret.DtDdosDstZoneOper.Oper.Ddos_entry_list),
			"total_dynamic_entry_count":   ret.DtDdosDstZoneOper.Oper.TotalDynamicEntryCount,
			"udp_dynamic_entry_count":     ret.DtDdosDstZoneOper.Oper.UdpDynamicEntryCount,
			"tcp_dynamic_entry_count":     ret.DtDdosDstZoneOper.Oper.TcpDynamicEntryCount,
			"icmp_dynamic_entry_count":    ret.DtDdosDstZoneOper.Oper.IcmpDynamicEntryCount,
			"other_dynamic_entry_count":   ret.DtDdosDstZoneOper.Oper.OtherDynamicEntryCount,
			"traffic_distribution_status": setSliceDdosDstZoneOperOperTrafficDistributionStatus(ret.DtDdosDstZoneOper.Oper.TrafficDistributionStatus),
			"entry_displayed_count":       ret.DtDdosDstZoneOper.Oper.EntryDisplayedCount,
			"service_displayed_count":     ret.DtDdosDstZoneOper.Oper.ServiceDisplayedCount,
			"no_t2_idx_port_count":        ret.DtDdosDstZoneOper.Oper.NoT2IdxPortCount,
			"addresses":                   ret.DtDdosDstZoneOper.Oper.Addresses,
			"subnet_ip_addr":              ret.DtDdosDstZoneOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":            ret.DtDdosDstZoneOper.Oper.SubnetIpv6Addr,
			"all_addresses":               ret.DtDdosDstZoneOper.Oper.AllAddresses,
			"ip_proto_num":                ret.DtDdosDstZoneOper.Oper.IpProtoNum,
			"all_ip_protos":               ret.DtDdosDstZoneOper.Oper.AllIpProtos,
			"port_num":                    ret.DtDdosDstZoneOper.Oper.PortNum,
			"port_range_start":            ret.DtDdosDstZoneOper.Oper.PortRangeStart,
			"port_range_end":              ret.DtDdosDstZoneOper.Oper.PortRangeEnd,
			"protocol":                    ret.DtDdosDstZoneOper.Oper.Protocol,
			"all_ports":                   ret.DtDdosDstZoneOper.Oper.AllPorts,
			"dynamic_expand_subnet":       ret.DtDdosDstZoneOper.Oper.DynamicExpandSubnet,
			"blackhole":                   ret.DtDdosDstZoneOper.Oper.Blackhole,
		},
	}
}

func setSliceDdosDstZoneOperOperDdos_entry_list(d []edpt.DdosDstZoneOperOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["port_str"] = item.PortStr
		in["operational_mode"] = item.OperationalMode
		in["bw_state"] = item.BwState
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
		in["dynamic_entry_count"] = item.DynamicEntryCount
		in["dynamic_entry_limit"] = item.DynamicEntryLimit
		in["age_str"] = item.AgeStr
		in["lockup_time"] = item.LockupTime
		in["sflow_source_id"] = item.SflowSourceId
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperOperTrafficDistributionStatus(d []edpt.DdosDstZoneOperOperTrafficDistributionStatus) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["master_pu"] = item.MasterPu
		in["active_pu"] = setSliceDdosDstZoneOperOperTrafficDistributionStatusActivePu(item.ActivePu)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperOperTrafficDistributionStatusActivePu(d []edpt.DdosDstZoneOperOperTrafficDistributionStatusActivePu) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pu_id"] = item.PuId
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperOutboundPolicy(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneOperOutboundPolicyOper(ret.DtDdosDstZoneOper.OutboundPolicy.Oper),
		},
	}
}

func setObjectDdosDstZoneOperOutboundPolicyOper(d edpt.DdosDstZoneOperOutboundPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["policy_name"] = d.PolicyName

	in["no_class_list_match"] = d.NoClassListMatch
	in["policy_class_list"] = setSliceDdosDstZoneOperOutboundPolicyOperPolicyClassList(d.PolicyClassList)
	in["geo_tracking_statistics"] = setObjectDdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics(d.GeoTrackingStatistics)
	in["tracking_entry_list"] = setSliceDdosDstZoneOperOutboundPolicyOperTrackingEntryList(d.TrackingEntryList)

	in["policy_rate"] = d.PolicyRate

	in["policy_statistics"] = d.PolicyStatistics

	in["tracking_entry_filter"] = d.TrackingEntryFilter
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperOutboundPolicyOperPolicyClassList(d []edpt.DdosDstZoneOperOutboundPolicyOperPolicyClassList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["class_list_name"] = item.ClassListName
		in["current_packet_rate"] = item.CurrentPacketRate
		in["is_packet_rate_exceed"] = item.IsPacketRateExceed
		in["packet_rate_limit"] = item.PacketRateLimit
		in["current_kbit_rate"] = item.CurrentKbitRate
		in["is_kbit_rate_exceed"] = item.IsKbitRateExceed
		in["kbit_rate_limit"] = item.KbitRateLimit
		in["current_connections"] = item.CurrentConnections
		in["is_connections_exceed"] = item.IsConnectionsExceed
		in["connection_limit"] = item.ConnectionLimit
		in["current_connection_rate"] = item.CurrentConnectionRate
		in["is_connection_rate_exceed"] = item.IsConnectionRateExceed
		in["connection_rate_limit"] = item.ConnectionRateLimit
		in["current_frag_packet_rate"] = item.CurrentFragPacketRate
		in["is_frag_packet_rate_exceed"] = item.IsFragPacketRateExceed
		in["frag_packet_rate_limit"] = item.FragPacketRateLimit
		in["age_str"] = item.AgeStr
		in["lockup_time"] = item.LockupTime
		in["packet_received"] = item.PacketReceived
		in["packet_dropped"] = item.PacketDropped
		in["packet_rate_exceed"] = item.PacketRateExceed
		in["kbit_rate_exceed"] = item.KbitRateExceed
		in["kbit_rate_exceed_count"] = item.KbitRateExceedCount
		in["connections_exceed"] = item.ConnectionsExceed
		in["connection_rate_exceed"] = item.ConnectionRateExceed
		in["frag_packet_rate"] = item.FragPacketRate
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics(d edpt.DdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["packet_received"] = d.PacketReceived

	in["packet_dropped"] = d.PacketDropped

	in["packet_rate_exceed"] = d.PacketRateExceed

	in["kbit_rate_exceed"] = d.KbitRateExceed

	in["kbit_rate_exceed_count"] = d.KbitRateExceedCount

	in["connections_exceed"] = d.ConnectionsExceed

	in["connection_rate_exceed"] = d.ConnectionRateExceed

	in["frag_packet_rate"] = d.FragPacketRate

	in["tracking_entry_learn"] = d.TrackingEntryLearn

	in["tracking_entry_aged"] = d.TrackingEntryAged

	in["tracking_entry_learning_thre_exceed"] = d.TrackingEntryLearningThreExceed
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperOutboundPolicyOperTrackingEntryList(d []edpt.DdosDstZoneOperOutboundPolicyOperTrackingEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["geo_location_name"] = item.GeoLocationName
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
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPacketAnomalyDetection(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneOperPacketAnomalyDetectionOper(ret.DtDdosDstZoneOper.PacketAnomalyDetection.Oper),
		},
	}
}

func setObjectDdosDstZoneOperPacketAnomalyDetectionOper(d edpt.DdosDstZoneOperPacketAnomalyDetectionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPacketAnomalyDetectionOperIndicators(d.Indicators)

	in["data_source"] = d.DataSource
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPacketAnomalyDetectionOperIndicators(d []edpt.DdosDstZoneOperPacketAnomalyDetectionOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["maximum"] = item.Maximum
		in["minimum"] = item.Minimum
		in["threshold"] = item.Threshold
		in["is_anomaly"] = item.IsAnomaly
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPort(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":                    setObjectDdosDstZoneOperPortOper(ret.DtDdosDstZoneOper.Port.Oper),
			"zone_service_list":       setSliceDdosDstZoneOperPortZoneServiceList(ret.DtDdosDstZoneOper.Port.ZoneServiceList),
			"zone_service_other_list": setSliceDdosDstZoneOperPortZoneServiceOtherList(ret.DtDdosDstZoneOper.Port.ZoneServiceOtherList),
		},
	}
}

func setObjectDdosDstZoneOperPortOper(d edpt.DdosDstZoneOperPortOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceList(d []edpt.DdosDstZoneOperPortZoneServiceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListOper(item.Oper)
		in["pattern_recognition"] = setObjectDdosDstZoneOperPortZoneServiceListPatternRecognition(item.PatternRecognition)
		in["pattern_recognition_pu_details"] = setObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails(item.PatternRecognitionPuDetails)
		in["ip_filtering_policy_oper"] = setObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["ips"] = setObjectDdosDstZoneOperPortZoneServiceListIps(item.Ips)
		in["port_ind"] = setObjectDdosDstZoneOperPortZoneServiceListPortInd(item.PortInd)
		in["topk_sources"] = setObjectDdosDstZoneOperPortZoneServiceListTopkSources(item.TopkSources)
		in["progression_tracking"] = setObjectDdosDstZoneOperPortZoneServiceListProgressionTracking(item.ProgressionTracking)
		in["topk_destinations"] = setObjectDdosDstZoneOperPortZoneServiceListTopkDestinations(item.TopkDestinations)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListOper(d edpt.DdosDstZoneOperPortZoneServiceListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["overflow_policy"] = d.OverflowPolicy

	in["sources_all_entries"] = d.SourcesAllEntries

	in["class_list"] = d.ClassList

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["exceeded"] = d.Exceeded

	in["black_listed"] = d.BlackListed

	in["white_listed"] = d.WhiteListed

	in["authenticated"] = d.Authenticated

	in["level"] = d.Level

	in["app_stat"] = d.AppStat

	in["indicators"] = d.Indicators

	in["indicator_detail"] = d.IndicatorDetail

	in["l4_ext_rate"] = d.L4ExtRate

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_list(d []edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_list) []map[string]interface{} {
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
		in["http_filter_rates"] = setSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates(item.HttpFilterRates)
		in["response_size_rates"] = setSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates(item.ResponseSizeRates)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates(d []edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates) []map[string]interface{} {
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

func setSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates(d []edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceListPatternRecognition(d edpt.DdosDstZoneOperPortZoneServiceListPatternRecognition) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionOper(d edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList(d []edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails(d edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper(d edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOper(d edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOper(d edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListIps(d edpt.DdosDstZoneOperPortZoneServiceListIps) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListIpsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListIpsOper(d edpt.DdosDstZoneOperPortZoneServiceListIpsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["signature_list"] = setSliceDdosDstZoneOperPortZoneServiceListIpsOperSignatureList(d.SignatureList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListIpsOperSignatureList(d []edpt.DdosDstZoneOperPortZoneServiceListIpsOperSignatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["match_count"] = item.MatchCount
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListPortInd(d edpt.DdosDstZoneOperPortZoneServiceListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListPortIndOper(d edpt.DdosDstZoneOperPortZoneServiceListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList(d []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceListTopkSources(d edpt.DdosDstZoneOperPortZoneServiceListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListTopkSourcesOper(d edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList(d []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceListProgressionTracking(d edpt.DdosDstZoneOperPortZoneServiceListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListProgressionTrackingOper(d edpt.DdosDstZoneOperPortZoneServiceListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceListTopkDestinations(d edpt.DdosDstZoneOperPortZoneServiceListTopkDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceListTopkDestinationsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceListTopkDestinationsOper(d edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList(d []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperPortZoneServiceOtherList(d []edpt.DdosDstZoneOperPortZoneServiceOtherList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_other"] = item.PortOther
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["pattern_recognition"] = setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognition(item.PatternRecognition)
		in["pattern_recognition_pu_details"] = setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails(item.PatternRecognitionPuDetails)
		in["port_ind"] = setObjectDdosDstZoneOperPortZoneServiceOtherListPortInd(item.PortInd)
		in["topk_sources"] = setObjectDdosDstZoneOperPortZoneServiceOtherListTopkSources(item.TopkSources)
		in["progression_tracking"] = setObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTracking(item.ProgressionTracking)
		in["topk_destinations"] = setObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinations(item.TopkDestinations)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["overflow_policy"] = d.OverflowPolicy

	in["sources_all_entries"] = d.SourcesAllEntries

	in["class_list"] = d.ClassList

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["exceeded"] = d.Exceeded

	in["black_listed"] = d.BlackListed

	in["white_listed"] = d.WhiteListed

	in["authenticated"] = d.Authenticated

	in["level"] = d.Level

	in["app_stat"] = d.AppStat

	in["indicators"] = d.Indicators

	in["indicator_detail"] = d.IndicatorDetail

	in["l4_ext_rate"] = d.L4ExtRate

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list(d []edpt.DdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognition(d edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognition) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList(d []edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails(d edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceOtherListPortInd(d edpt.DdosDstZoneOperPortZoneServiceOtherListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListPortIndOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList(d []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceOtherListTopkSources(d edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTracking(d edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinations(d edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper(d edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperPortRangeList(d edpt.DataDdosDstZoneOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstZoneOper.PortRangeList {
		in := make(map[string]interface{})
		in["port_range_start"] = item.PortRangeStart
		in["port_range_end"] = item.PortRangeEnd
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperPortRangeListOper(item.Oper)
		in["ip_filtering_policy_oper"] = setObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOper(item.IpFilteringPolicyOper)
		in["pattern_recognition"] = setObjectDdosDstZoneOperPortRangeListPatternRecognition(item.PatternRecognition)
		in["pattern_recognition_pu_details"] = setObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetails(item.PatternRecognitionPuDetails)
		in["ips"] = setObjectDdosDstZoneOperPortRangeListIps(item.Ips)
		in["port_ind"] = setObjectDdosDstZoneOperPortRangeListPortInd(item.PortInd)
		in["topk_sources"] = setObjectDdosDstZoneOperPortRangeListTopkSources(item.TopkSources)
		in["topk_destinations"] = setObjectDdosDstZoneOperPortRangeListTopkDestinations(item.TopkDestinations)
		in["progression_tracking"] = setObjectDdosDstZoneOperPortRangeListProgressionTracking(item.ProgressionTracking)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortRangeListOper(d edpt.DdosDstZoneOperPortRangeListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperPortRangeListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["overflow_policy"] = d.OverflowPolicy

	in["sources_all_entries"] = d.SourcesAllEntries

	in["class_list"] = d.ClassList

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["exceeded"] = d.Exceeded

	in["black_listed"] = d.BlackListed

	in["white_listed"] = d.WhiteListed

	in["authenticated"] = d.Authenticated

	in["level"] = d.Level

	in["app_stat"] = d.AppStat

	in["indicators"] = d.Indicators

	in["indicator_detail"] = d.IndicatorDetail

	in["l4_ext_rate"] = d.L4ExtRate

	in["hw_blacklisted"] = d.HwBlacklisted

	in["suffix_request_rate"] = d.SuffixRequestRate

	in["domain_name"] = d.DomainName
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListOperDdos_entry_list(d []edpt.DdosDstZoneOperPortRangeListOperDdos_entry_list) []map[string]interface{} {
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
		in["http_filter_rates"] = setSliceDdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates(item.HttpFilterRates)
		in["response_size_rates"] = setSliceDdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates(item.ResponseSizeRates)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates(d []edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates) []map[string]interface{} {
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

func setSliceDdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates(d []edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOper(d edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOperOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOperOper(d edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortRangeListPatternRecognition(d edpt.DdosDstZoneOperPortRangeListPatternRecognition) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListPatternRecognitionOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListPatternRecognitionOper(d edpt.DdosDstZoneOperPortRangeListPatternRecognitionOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["timestamp"] = d.Timestamp

	in["peace_pkt_count"] = d.PeacePktCount

	in["war_pkt_count"] = d.WarPktCount

	in["war_pkt_percentage"] = d.WarPktPercentage

	in["filter_threshold"] = d.FilterThreshold

	in["filter_count"] = d.FilterCount
	in["filter_list"] = setSliceDdosDstZoneOperPortRangeListPatternRecognitionOperFilterList(d.FilterList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListPatternRecognitionOperFilterList(d []edpt.DdosDstZoneOperPortRangeListPatternRecognitionOperFilterList) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetails(d edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetails) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper(d edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["all_filters"] = setSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(d.AllFilters)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(d []edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortRangeListIps(d edpt.DdosDstZoneOperPortRangeListIps) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListIpsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListIpsOper(d edpt.DdosDstZoneOperPortRangeListIpsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["signature_list"] = setSliceDdosDstZoneOperPortRangeListIpsOperSignatureList(d.SignatureList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListIpsOperSignatureList(d []edpt.DdosDstZoneOperPortRangeListIpsOperSignatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["match_count"] = item.MatchCount
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperPortRangeListPortInd(d edpt.DdosDstZoneOperPortRangeListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListPortIndOper(d edpt.DdosDstZoneOperPortRangeListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["src_entry_list"] = setSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryList(d.SrcEntryList)
	in["indicators"] = setSliceDdosDstZoneOperPortRangeListPortIndOperIndicators(d.Indicators)

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

func setSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryList(d []edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators(item.Indicators)
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

func setSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators(d []edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators) []map[string]interface{} {
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

func setSliceDdosDstZoneOperPortRangeListPortIndOperIndicators(d []edpt.DdosDstZoneOperPortRangeListPortIndOperIndicators) []map[string]interface{} {
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
		in["indicator_cfg"] = setSliceDdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortRangeListTopkSources(d edpt.DdosDstZoneOperPortRangeListTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListTopkSourcesOper(d edpt.DdosDstZoneOperPortRangeListTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicators(d []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources(d []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryList(d []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators(d []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortRangeListTopkDestinations(d edpt.DdosDstZoneOperPortRangeListTopkDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListTopkDestinationsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListTopkDestinationsOper(d edpt.DdosDstZoneOperPortRangeListTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicators(d []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryList(d []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperPortRangeListProgressionTracking(d edpt.DdosDstZoneOperPortRangeListProgressionTracking) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperPortRangeListProgressionTrackingOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperPortRangeListProgressionTrackingOper(d edpt.DdosDstZoneOperPortRangeListProgressionTrackingOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperPortRangeListProgressionTrackingOperIndicators(d.Indicators)

	in["learning_details"] = d.LearningDetails

	in["learning_brief"] = d.LearningBrief

	in["recommended_template"] = d.RecommendedTemplate

	in["template_debug_table"] = d.TemplateDebugTable
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperPortRangeListProgressionTrackingOperIndicators(d []edpt.DdosDstZoneOperPortRangeListProgressionTrackingOperIndicators) []map[string]interface{} {
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

func setObjectDdosDstZoneOperSrcPort(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":                     setObjectDdosDstZoneOperSrcPortOper(ret.DtDdosDstZoneOper.SrcPort.Oper),
			"zone_src_port_list":       setSliceDdosDstZoneOperSrcPortZoneSrcPortList(ret.DtDdosDstZoneOper.SrcPort.ZoneSrcPortList),
			"zone_src_port_other_list": setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherList(ret.DtDdosDstZoneOper.SrcPort.ZoneSrcPortOtherList),
		},
	}
}

func setObjectDdosDstZoneOperSrcPortOper(d edpt.DdosDstZoneOperSrcPortOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperSrcPortZoneSrcPortList(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperSrcPortZoneSrcPortListOper(item.Oper)
		in["port_ind"] = setObjectDdosDstZoneOperSrcPortZoneSrcPortListPortInd(item.PortInd)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperSrcPortZoneSrcPortListOper(d edpt.DdosDstZoneOperSrcPortZoneSrcPortListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["sources_all_entries"] = d.SourcesAllEntries

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["hw_blacklisted"] = d.HwBlacklisted
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["bw_state"] = item.BwState
		in["is_auth_passed"] = item.Is_auth_passed
		in["level"] = item.Level
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

func setObjectDdosDstZoneOperSrcPortZoneSrcPortListPortInd(d edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperSrcPortZoneSrcPortListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperSrcPortZoneSrcPortListPortIndOper(d edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource

	in["current_level"] = d.CurrentLevel

	in["details"] = d.Details
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["indicator_cfg"] = setSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherList(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_other"] = item.PortOther
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListOper(item.Oper)
		in["port_ind"] = setObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd(item.PortInd)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListOper(d edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["sources_all_entries"] = d.SourcesAllEntries

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["hw_blacklisted"] = d.HwBlacklisted
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["bw_state"] = item.BwState
		in["is_auth_passed"] = item.Is_auth_passed
		in["level"] = item.Level
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

func setObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd(d edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper(d edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource

	in["current_level"] = d.CurrentLevel

	in["details"] = d.Details
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["indicator_cfg"] = setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setSliceDdosDstZoneOperSrcPortRangeList(d edpt.DataDdosDstZoneOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosDstZoneOper.SrcPortRangeList {
		in := make(map[string]interface{})
		in["src_port_range_start"] = item.SrcPortRangeStart
		in["src_port_range_end"] = item.SrcPortRangeEnd
		in["protocol"] = item.Protocol
		in["oper"] = setObjectDdosDstZoneOperSrcPortRangeListOper(item.Oper)
		in["port_ind"] = setObjectDdosDstZoneOperSrcPortRangeListPortInd(item.PortInd)
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOperSrcPortRangeListOper(d edpt.DdosDstZoneOperSrcPortRangeListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ddos_entry_list"] = setSliceDdosDstZoneOperSrcPortRangeListOperDdos_entry_list(d.Ddos_entry_list)

	in["entry_displayed_count"] = d.EntryDisplayedCount

	in["service_displayed_count"] = d.ServiceDisplayedCount

	in["reporting_status"] = d.ReportingStatus

	in["sources"] = d.Sources

	in["sources_all_entries"] = d.SourcesAllEntries

	in["subnet_ip_addr"] = d.SubnetIpAddr

	in["subnet_ipv6_addr"] = d.SubnetIpv6Addr

	in["ipv6"] = d.Ipv6

	in["hw_blacklisted"] = d.HwBlacklisted
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperSrcPortRangeListOperDdos_entry_list(d []edpt.DdosDstZoneOperSrcPortRangeListOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["bw_state"] = item.BwState
		in["is_auth_passed"] = item.Is_auth_passed
		in["level"] = item.Level
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

func setObjectDdosDstZoneOperSrcPortRangeListPortInd(d edpt.DdosDstZoneOperSrcPortRangeListPortInd) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectDdosDstZoneOperSrcPortRangeListPortIndOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectDdosDstZoneOperSrcPortRangeListPortIndOper(d edpt.DdosDstZoneOperSrcPortRangeListPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource

	in["current_level"] = d.CurrentLevel

	in["details"] = d.Details
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicators(d []edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["indicator_cfg"] = setSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func setObjectDdosDstZoneOperTopkDestinations(ret edpt.DataDdosDstZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneOperTopkDestinationsOper(ret.DtDdosDstZoneOper.TopkDestinations.Oper),
		},
	}
}

func setObjectDdosDstZoneOperTopkDestinationsOper(d edpt.DdosDstZoneOperTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneOperTopkDestinationsOperIndicators(d.Indicators)
	in["entry_list"] = setSliceDdosDstZoneOperTopkDestinationsOperEntryList(d.EntryList)

	in["next_indicator"] = d.NextIndicator

	in["finished"] = d.Finished

	in["details"] = d.Details

	in["top_k_key"] = d.TopKKey
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOperTopkDestinationsOperIndicators(d []edpt.DdosDstZoneOperTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneOperTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosDstZoneOperTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperTopkDestinationsOperEntryList(d []edpt.DdosDstZoneOperTopkDestinationsOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneOperTopkDestinationsOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneOperTopkDestinationsOperEntryListIndicators(d []edpt.DdosDstZoneOperTopkDestinationsOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneOperDetection(d []interface{}) edpt.DdosDstZoneOperDetection {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetection
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperDetectionOper(in["oper"].([]interface{}))
		ret.OutboundDetection = getObjectDdosDstZoneOperDetectionOutboundDetection(in["outbound_detection"].([]interface{}))
		ret.ServiceDiscovery = getObjectDdosDstZoneOperDetectionServiceDiscovery(in["service_discovery"].([]interface{}))
		ret.VictimIpDetection = getObjectDdosDstZoneOperDetectionVictimIpDetection(in["victim_ip_detection"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionOper(d []interface{}) edpt.DdosDstZoneOperDetectionOper {

	var ret edpt.DdosDstZoneOperDetectionOper
	return ret
}

func getObjectDdosDstZoneOperDetectionOutboundDetection(d []interface{}) edpt.DdosDstZoneOperDetectionOutboundDetection {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionOutboundDetection
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperDetectionOutboundDetectionOper(in["oper"].([]interface{}))
		ret.TopkSourceSubnet = getObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet(in["topk_source_subnet"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionOutboundDetectionOper(d []interface{}) edpt.DdosDstZoneOperDetectionOutboundDetectionOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionOutboundDetectionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DiscoveryTimestamp = in["discovery_timestamp"].(string)
		ret.EntryList = getSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryList(d []interface{}) []edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryList
		oi.LocationType = in["location_type"].(string)
		oi.LocationName = in["location_name"].(string)
		oi.Indicators = getSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators(in["indicators"].([]interface{}))
		oi.DataSource = in["data_source"].(string)
		oi.Anomaly = in["anomaly"].(string)
		oi.AnomalyTimestamp = in["anomaly_timestamp"].(string)
		oi.InitialLearning = in["initial_learning"].(string)
		oi.ActiveTime = in["active_time"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.Maximum = in["maximum"].(string)
		oi.Minimum = in["minimum"].(string)
		oi.NonZeroMinimum = in["non_zero_minimum"].(string)
		oi.Average = in["average"].(string)
		oi.AdaptiveThreshold = in["adaptive_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet(d []interface{}) edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper(d []interface{}) edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList(d []interface{}) []edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList
		oi.LocationType = in["location_type"].(string)
		oi.LocationName = in["location_name"].(string)
		oi.Indicators = getSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.SourceSubnets = getSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(in["source_subnets"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(d []interface{}) []edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionServiceDiscovery(d []interface{}) edpt.DdosDstZoneOperDetectionServiceDiscovery {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionServiceDiscovery
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperDetectionServiceDiscoveryOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionServiceDiscoveryOper(d []interface{}) edpt.DdosDstZoneOperDetectionServiceDiscoveryOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionServiceDiscoveryOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DiscoveredServiceList = getSliceDdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList(in["discovered_service_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList(d []interface{}) []edpt.DdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Rate = in["rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionVictimIpDetection(d []interface{}) edpt.DdosDstZoneOperDetectionVictimIpDetection {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionVictimIpDetection
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperDetectionVictimIpDetectionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperDetectionVictimIpDetectionOper(d []interface{}) edpt.DdosDstZoneOperDetectionVictimIpDetectionOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperDetectionVictimIpDetectionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpEntryList = getSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList(in["ip_entry_list"].([]interface{}))
		ret.IpEntryCount = in["ip_entry_count"].(int)
		ret.TotalIpEntryCount = in["total_ip_entry_count"].(int)
		ret.ActiveList = in["active_list"].(int)
		ret.VictimList = in["victim_list"].(int)
		ret.Ipv4Ip = in["ipv4_ip"].(string)
		ret.Ipv6Ip = in["ipv6_ip"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList(d []interface{}) []edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList
		oi.IpAddressStr = in["ip_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators(in["indicators"].([]interface{}))
		oi.IsLearningDone = in["is_learning_done"].(int)
		oi.IsHistogramLearningDone = in["is_histogram_learning_done"].(int)
		oi.IsIpAnomaly = in["is_ip_anomaly"].(int)
		oi.Is_static_threshold = in["is_static_threshold"].(int)
		oi.EscalationTimestamp = in["escalation_timestamp"].(string)
		oi.DeEscalationTimestamp = in["de_escalation_timestamp"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Value = getSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue(in["value"].([]interface{}))
		oi.IsAnomaly = in["is_anomaly"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue(d []interface{}) []edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue
		oi.Current = in["current"].(string)
		oi.Threshold = in["threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProto(d []interface{}) edpt.DdosDstZoneOperIpProto {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProto
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoOper(in["oper"].([]interface{}))
		ret.ProtoNumberList = getSliceDdosDstZoneOperIpProtoProtoNumberList(in["proto_number_list"].([]interface{}))
		ret.ProtoTcpUdpList = getSliceDdosDstZoneOperIpProtoProtoTcpUdpList(in["proto_tcp_udp_list"].([]interface{}))
		ret.ProtoNameList = getSliceDdosDstZoneOperIpProtoProtoNameList(in["proto_name_list"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoOper(d []interface{}) edpt.DdosDstZoneOperIpProtoOper {

	var ret edpt.DdosDstZoneOperIpProtoOper
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberList
		oi.ProtocolNum = in["protocol_num"].(int)
		oi.Oper = getObjectDdosDstZoneOperIpProtoProtoNumberListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperIpProtoProtoNumberListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZoneOperIpProtoProtoNumberListTopkSources(in["topk_sources"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinations(in["topk_destinations"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTracking(in["progression_tracking"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list
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

func getObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListPortInd(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNumberListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListPortIndOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListTopkSources(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinations(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTracking(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators
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

func getSliceDdosDstZoneOperIpProtoProtoTcpUdpList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoTcpUdpList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoTcpUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoTcpUdpList
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperIpProtoProtoTcpUdpListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoTcpUdpListOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoTcpUdpListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoTcpUdpListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list
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

func getObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameList
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperIpProtoProtoNameListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperIpProtoProtoNameListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZoneOperIpProtoProtoNameListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZoneOperIpProtoProtoNameListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinations(in["topk_destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list
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

func getObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListPortInd(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNameListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListPortIndOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListTopkSources(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNameListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListTopkSourcesOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListProgressionTracking(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators
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

func getObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinations(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperOper(d []interface{}) edpt.DdosDstZoneOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.TotalDynamicEntryCount = in["total_dynamic_entry_count"].(string)
		ret.UdpDynamicEntryCount = in["udp_dynamic_entry_count"].(string)
		ret.TcpDynamicEntryCount = in["tcp_dynamic_entry_count"].(string)
		ret.IcmpDynamicEntryCount = in["icmp_dynamic_entry_count"].(string)
		ret.OtherDynamicEntryCount = in["other_dynamic_entry_count"].(string)
		ret.TrafficDistributionStatus = getSliceDdosDstZoneOperOperTrafficDistributionStatus(in["traffic_distribution_status"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.NoT2IdxPortCount = in["no_t2_idx_port_count"].(int)
		ret.Addresses = in["addresses"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.AllAddresses = in["all_addresses"].(int)
		ret.IpProtoNum = in["ip_proto_num"].(int)
		ret.AllIpProtos = in["all_ip_protos"].(int)
		ret.PortNum = in["port_num"].(int)
		ret.PortRangeStart = in["port_range_start"].(int)
		ret.PortRangeEnd = in["port_range_end"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.AllPorts = in["all_ports"].(int)
		ret.DynamicExpandSubnet = in["dynamic_expand_subnet"].(int)
		ret.Blackhole = in["blackhole"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.PortStr = in["port_str"].(string)
		oi.OperationalMode = in["operational_mode"].(string)
		oi.BwState = in["bw_state"].(string)
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
		oi.DynamicEntryCount = in["dynamic_entry_count"].(string)
		oi.DynamicEntryLimit = in["dynamic_entry_limit"].(string)
		oi.AgeStr = in["age_str"].(string)
		oi.LockupTime = in["lockup_time"].(int)
		oi.SflowSourceId = in["sflow_source_id"].(int)
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperOperTrafficDistributionStatus(d []interface{}) []edpt.DdosDstZoneOperOperTrafficDistributionStatus {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperOperTrafficDistributionStatus, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperOperTrafficDistributionStatus
		oi.MasterPu = in["master_pu"].(string)
		oi.ActivePu = getSliceDdosDstZoneOperOperTrafficDistributionStatusActivePu(in["active_pu"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperOperTrafficDistributionStatusActivePu(d []interface{}) []edpt.DdosDstZoneOperOperTrafficDistributionStatusActivePu {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperOperTrafficDistributionStatusActivePu, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperOperTrafficDistributionStatusActivePu
		oi.PuId = in["pu_id"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperOutboundPolicy(d []interface{}) edpt.DdosDstZoneOperOutboundPolicy {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperOutboundPolicy
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperOutboundPolicyOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperOutboundPolicyOper(d []interface{}) edpt.DdosDstZoneOperOutboundPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperOutboundPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PolicyName = in["policy_name"].(string)
		ret.NoClassListMatch = in["no_class_list_match"].(int)
		ret.PolicyClassList = getSliceDdosDstZoneOperOutboundPolicyOperPolicyClassList(in["policy_class_list"].([]interface{}))
		ret.GeoTrackingStatistics = getObjectDdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics(in["geo_tracking_statistics"].([]interface{}))
		ret.TrackingEntryList = getSliceDdosDstZoneOperOutboundPolicyOperTrackingEntryList(in["tracking_entry_list"].([]interface{}))
		ret.PolicyRate = in["policy_rate"].(int)
		ret.PolicyStatistics = in["policy_statistics"].(int)
		ret.TrackingEntryFilter = in["tracking_entry_filter"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperOutboundPolicyOperPolicyClassList(d []interface{}) []edpt.DdosDstZoneOperOutboundPolicyOperPolicyClassList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperOutboundPolicyOperPolicyClassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperOutboundPolicyOperPolicyClassList
		oi.ClassListName = in["class_list_name"].(string)
		oi.CurrentPacketRate = in["current_packet_rate"].(string)
		oi.IsPacketRateExceed = in["is_packet_rate_exceed"].(int)
		oi.PacketRateLimit = in["packet_rate_limit"].(string)
		oi.CurrentKbitRate = in["current_kbit_rate"].(string)
		oi.IsKbitRateExceed = in["is_kbit_rate_exceed"].(int)
		oi.KbitRateLimit = in["kbit_rate_limit"].(string)
		oi.CurrentConnections = in["current_connections"].(string)
		oi.IsConnectionsExceed = in["is_connections_exceed"].(int)
		oi.ConnectionLimit = in["connection_limit"].(string)
		oi.CurrentConnectionRate = in["current_connection_rate"].(string)
		oi.IsConnectionRateExceed = in["is_connection_rate_exceed"].(int)
		oi.ConnectionRateLimit = in["connection_rate_limit"].(string)
		oi.CurrentFragPacketRate = in["current_frag_packet_rate"].(string)
		oi.IsFragPacketRateExceed = in["is_frag_packet_rate_exceed"].(int)
		oi.FragPacketRateLimit = in["frag_packet_rate_limit"].(string)
		oi.AgeStr = in["age_str"].(string)
		oi.LockupTime = in["lockup_time"].(int)
		oi.PacketReceived = in["packet_received"].(int)
		oi.PacketDropped = in["packet_dropped"].(int)
		oi.PacketRateExceed = in["packet_rate_exceed"].(int)
		oi.KbitRateExceed = in["kbit_rate_exceed"].(int)
		oi.KbitRateExceedCount = in["kbit_rate_exceed_count"].(int)
		oi.ConnectionsExceed = in["connections_exceed"].(int)
		oi.ConnectionRateExceed = in["connection_rate_exceed"].(int)
		oi.FragPacketRate = in["frag_packet_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics(d []interface{}) edpt.DdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketReceived = in["packet_received"].(int)
		ret.PacketDropped = in["packet_dropped"].(int)
		ret.PacketRateExceed = in["packet_rate_exceed"].(int)
		ret.KbitRateExceed = in["kbit_rate_exceed"].(int)
		ret.KbitRateExceedCount = in["kbit_rate_exceed_count"].(int)
		ret.ConnectionsExceed = in["connections_exceed"].(int)
		ret.ConnectionRateExceed = in["connection_rate_exceed"].(int)
		ret.FragPacketRate = in["frag_packet_rate"].(int)
		ret.TrackingEntryLearn = in["tracking_entry_learn"].(int)
		ret.TrackingEntryAged = in["tracking_entry_aged"].(int)
		ret.TrackingEntryLearningThreExceed = in["tracking_entry_learning_thre_exceed"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperOutboundPolicyOperTrackingEntryList(d []interface{}) []edpt.DdosDstZoneOperOutboundPolicyOperTrackingEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperOutboundPolicyOperTrackingEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperOutboundPolicyOperTrackingEntryList
		oi.GeoLocationName = in["geo_location_name"].(string)
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
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPacketAnomalyDetection(d []interface{}) edpt.DdosDstZoneOperPacketAnomalyDetection {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPacketAnomalyDetection
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPacketAnomalyDetectionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPacketAnomalyDetectionOper(d []interface{}) edpt.DdosDstZoneOperPacketAnomalyDetectionOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPacketAnomalyDetectionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPacketAnomalyDetectionOperIndicators(in["indicators"].([]interface{}))
		ret.DataSource = in["data_source"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperPacketAnomalyDetectionOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPacketAnomalyDetectionOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPacketAnomalyDetectionOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPacketAnomalyDetectionOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.Maximum = in["maximum"].(string)
		oi.Minimum = in["minimum"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.IsAnomaly = in["is_anomaly"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPort(d []interface{}) edpt.DdosDstZoneOperPort {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPort
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortOper(in["oper"].([]interface{}))
		ret.ZoneServiceList = getSliceDdosDstZoneOperPortZoneServiceList(in["zone_service_list"].([]interface{}))
		ret.ZoneServiceOtherList = getSliceDdosDstZoneOperPortZoneServiceOtherList(in["zone_service_other_list"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortOper(d []interface{}) edpt.DdosDstZoneOperPortOper {

	var ret edpt.DdosDstZoneOperPortOper
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperPortZoneServiceListOper(in["oper"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstZoneOperPortZoneServiceListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.Ips = getObjectDdosDstZoneOperPortZoneServiceListIps(in["ips"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperPortZoneServiceListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZoneOperPortZoneServiceListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZoneOperPortZoneServiceListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZoneOperPortZoneServiceListTopkDestinations(in["topk_destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_list
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
		oi.HttpFilterRates = getSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates(in["http_filter_rates"].([]interface{}))
		oi.ResponseSizeRates = getSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates(in["response_size_rates"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates
		oi.HttpFilterRateName = in["http_filter_rate_name"].(string)
		oi.IsHttpFilterRateLimitExceed = in["is_http_filter_rate_limit_exceed"].(int)
		oi.CurrentHttpFilterRate = in["current_http_filter_rate"].(string)
		oi.HttpFilterRateLimit = in["http_filter_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates
		oi.ResponseSizeRateName = in["response_size_rate_name"].(string)
		oi.IsResponseSizeRateLimitExceed = in["is_response_size_rate_limit_exceed"].(int)
		oi.CurrentResponseSizeRate = in["current_response_size_rate"].(string)
		oi.ResponseSizeRateLimit = in["response_size_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListPatternRecognition(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList
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

func getObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListIps(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListIps {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListIps
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListIpsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListIpsOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListIpsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListIpsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignatureList = getSliceDdosDstZoneOperPortZoneServiceListIpsOperSignatureList(in["signature_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListIpsOperSignatureList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListIpsOperSignatureList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListIpsOperSignatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListIpsOperSignatureList
		oi.Sid = in["sid"].(int)
		oi.MatchCount = in["match_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListPortInd(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListPortIndOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListTopkSources(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListTopkSourcesOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListProgressionTracking(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListProgressionTrackingOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators
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

func getObjectDdosDstZoneOperPortZoneServiceListTopkDestinations(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceListTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceListTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherList
		oi.PortOther = in["port_other"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperPortZoneServiceOtherListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZoneOperPortZoneServiceOtherListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinations(in["topk_destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list
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

func getObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognition(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList
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

func getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListPortInd(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListPortIndOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListTopkSources(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTracking(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators
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

func getObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinations(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeList(d []interface{}) []edpt.DdosDstZoneOperPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeList
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperPortRangeListOper(in["oper"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstZoneOperPortRangeListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		oi.Ips = getObjectDdosDstZoneOperPortRangeListIps(in["ips"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperPortRangeListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZoneOperPortRangeListTopkSources(in["topk_sources"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZoneOperPortRangeListTopkDestinations(in["topk_destinations"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZoneOperPortRangeListProgressionTracking(in["progression_tracking"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperPortRangeListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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

func getSliceDdosDstZoneOperPortRangeListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperPortRangeListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListOperDdos_entry_list
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
		oi.HttpFilterRates = getSliceDdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates(in["http_filter_rates"].([]interface{}))
		oi.ResponseSizeRates = getSliceDdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates(in["response_size_rates"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates(d []interface{}) []edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates
		oi.HttpFilterRateName = in["http_filter_rate_name"].(string)
		oi.IsHttpFilterRateLimitExceed = in["is_http_filter_rate_limit_exceed"].(int)
		oi.CurrentHttpFilterRate = in["current_http_filter_rate"].(string)
		oi.HttpFilterRateLimit = in["http_filter_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates(d []interface{}) []edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates
		oi.ResponseSizeRateName = in["response_size_rate_name"].(string)
		oi.IsResponseSizeRateLimitExceed = in["is_response_size_rate_limit_exceed"].(int)
		oi.CurrentResponseSizeRate = in["current_response_size_rate"].(string)
		oi.ResponseSizeRateLimit = in["response_size_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListPatternRecognition(d []interface{}) edpt.DdosDstZoneOperPortRangeListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListPatternRecognitionOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListPatternRecognitionOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListPatternRecognitionOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListPatternRecognitionOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstZoneOperPortRangeListPatternRecognitionOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListPatternRecognitionOperFilterList(d []interface{}) []edpt.DdosDstZoneOperPortRangeListPatternRecognitionOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListPatternRecognitionOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListPatternRecognitionOperFilterList
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

func getObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetails {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetails
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters(d []interface{}) []edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListIps(d []interface{}) edpt.DdosDstZoneOperPortRangeListIps {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListIps
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListIpsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListIpsOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListIpsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListIpsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignatureList = getSliceDdosDstZoneOperPortRangeListIpsOperSignatureList(in["signature_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListIpsOperSignatureList(d []interface{}) []edpt.DdosDstZoneOperPortRangeListIpsOperSignatureList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListIpsOperSignatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListIpsOperSignatureList
		oi.Sid = in["sid"].(int)
		oi.MatchCount = in["match_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListPortInd(d []interface{}) edpt.DdosDstZoneOperPortRangeListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListPortIndOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneOperPortRangeListPortIndOperIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators(in["indicators"].([]interface{}))
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

func getSliceDdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators
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

func getSliceDdosDstZoneOperPortRangeListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortRangeListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListTopkSources(d []interface{}) edpt.DdosDstZoneOperPortRangeListTopkSources {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListTopkSourcesOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListTopkSourcesOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryList(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListTopkDestinations(d []interface{}) edpt.DdosDstZoneOperPortRangeListTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListProgressionTracking(d []interface{}) edpt.DdosDstZoneOperPortRangeListProgressionTracking {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListProgressionTracking
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperPortRangeListProgressionTrackingOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperPortRangeListProgressionTrackingOper(d []interface{}) edpt.DdosDstZoneOperPortRangeListProgressionTrackingOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperPortRangeListProgressionTrackingOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperPortRangeListProgressionTrackingOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperPortRangeListProgressionTrackingOperIndicators(d []interface{}) []edpt.DdosDstZoneOperPortRangeListProgressionTrackingOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperPortRangeListProgressionTrackingOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperPortRangeListProgressionTrackingOperIndicators
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

func getObjectDdosDstZoneOperSrcPort(d []interface{}) edpt.DdosDstZoneOperSrcPort {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPort
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperSrcPortOper(in["oper"].([]interface{}))
		ret.ZoneSrcPortList = getSliceDdosDstZoneOperSrcPortZoneSrcPortList(in["zone_src_port_list"].([]interface{}))
		ret.ZoneSrcPortOtherList = getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherList(in["zone_src_port_other_list"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperSrcPortOper(d []interface{}) edpt.DdosDstZoneOperSrcPortOper {

	var ret edpt.DdosDstZoneOperSrcPortOper
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortList(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperSrcPortZoneSrcPortListOper(in["oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperSrcPortZoneSrcPortListPortInd(in["port_ind"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperSrcPortZoneSrcPortListOper(d []interface{}) edpt.DdosDstZoneOperSrcPortZoneSrcPortListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortZoneSrcPortListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.Sources = in["sources"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.BwState = in["bw_state"].(string)
		oi.Is_auth_passed = in["is_auth_passed"].(string)
		oi.Level = in["level"].(int)
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

func getObjectDdosDstZoneOperSrcPortZoneSrcPortListPortInd(d []interface{}) edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperSrcPortZoneSrcPortListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperSrcPortZoneSrcPortListPortIndOper(d []interface{}) edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.Details = in["details"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherList(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherList
		oi.PortOther = in["port_other"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListOper(in["oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd(in["port_ind"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListOper(d []interface{}) edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.Sources = in["sources"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.BwState = in["bw_state"].(string)
		oi.Is_auth_passed = in["is_auth_passed"].(string)
		oi.Level = in["level"].(int)
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

func getObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd(d []interface{}) edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper(d []interface{}) edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.Details = in["details"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortRangeList(d []interface{}) []edpt.DdosDstZoneOperSrcPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortRangeList
		oi.SrcPortRangeStart = in["src_port_range_start"].(int)
		oi.SrcPortRangeEnd = in["src_port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectDdosDstZoneOperSrcPortRangeListOper(in["oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneOperSrcPortRangeListPortInd(in["port_ind"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperSrcPortRangeListOper(d []interface{}) edpt.DdosDstZoneOperSrcPortRangeListOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortRangeListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneOperSrcPortRangeListOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.Sources = in["sources"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortRangeListOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneOperSrcPortRangeListOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortRangeListOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortRangeListOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.BwState = in["bw_state"].(string)
		oi.Is_auth_passed = in["is_auth_passed"].(string)
		oi.Level = in["level"].(int)
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

func getObjectDdosDstZoneOperSrcPortRangeListPortInd(d []interface{}) edpt.DdosDstZoneOperSrcPortRangeListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortRangeListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperSrcPortRangeListPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperSrcPortRangeListPortIndOper(d []interface{}) edpt.DdosDstZoneOperSrcPortRangeListPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperSrcPortRangeListPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.Details = in["details"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOperTopkDestinations(d []interface{}) edpt.DdosDstZoneOperTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneOperTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneOperTopkDestinationsOper(d []interface{}) edpt.DdosDstZoneOperTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOperTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneOperTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneOperTopkDestinationsOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneOperTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosDstZoneOperTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneOperTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneOperTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperTopkDestinationsOperEntryList(d []interface{}) []edpt.DdosDstZoneOperTopkDestinationsOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperTopkDestinationsOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperTopkDestinationsOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneOperTopkDestinationsOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneOperTopkDestinationsOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneOperTopkDestinationsOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOperTopkDestinationsOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOperTopkDestinationsOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneOper(d *schema.ResourceData) edpt.DdosDstZoneOper {
	var ret edpt.DdosDstZoneOper

	ret.Detection = getObjectDdosDstZoneOperDetection(d.Get("detection").([]interface{}))

	ret.IpProto = getObjectDdosDstZoneOperIpProto(d.Get("ip_proto").([]interface{}))

	ret.Oper = getObjectDdosDstZoneOperOper(d.Get("oper").([]interface{}))

	ret.OutboundPolicy = getObjectDdosDstZoneOperOutboundPolicy(d.Get("outbound_policy").([]interface{}))

	ret.PacketAnomalyDetection = getObjectDdosDstZoneOperPacketAnomalyDetection(d.Get("packet_anomaly_detection").([]interface{}))

	ret.Port = getObjectDdosDstZoneOperPort(d.Get("port").([]interface{}))

	ret.PortRangeList = getSliceDdosDstZoneOperPortRangeList(d.Get("port_range_list").([]interface{}))

	ret.SrcPort = getObjectDdosDstZoneOperSrcPort(d.Get("src_port").([]interface{}))

	ret.SrcPortRangeList = getSliceDdosDstZoneOperSrcPortRangeList(d.Get("src_port_range_list").([]interface{}))

	ret.TopkDestinations = getObjectDdosDstZoneOperTopkDestinations(d.Get("topk_destinations").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
