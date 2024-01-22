package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZone() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone`: Configure a static zone entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneCreate,
		UpdateContext: resourceDdosDstZoneUpdate,
		ReadContext:   resourceDdosDstZoneRead,
		DeleteContext: resourceDdosDstZoneDelete,

		Schema: map[string]*schema.Schema{
			"action_list": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"advertised_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP advertised",
			},
			"capture_config_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Capture-config name",
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Apply capture-config to dropped packets; 'forward': Apply capture-config to forwarded packets; 'all': Apply capture-config to both dropped and forwarded packets;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"collector": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sflow_name": {
							Type: schema.TypeString, Optional: true, Description: "Name of configured custom sFlow collector",
						},
					},
				},
			},
			"continuous_learning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Continuous learning of detection",
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description for this Destination Zone",
			},
			"dest_nat_ip": {
				Type: schema.TypeString, Optional: true, Description: "Destination NAT IP address",
			},
			"dest_nat_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Destination NAT IPv6 address",
			},
			"detection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"settings": {
							Type: schema.TypeString, Optional: true, Description: "'settings': settings;",
						},
						"toggle": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable detection; 'disable': Disable detection;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"notification": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configuration": {
										Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
									},
									"notification": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"notification_template_name": {
													Type: schema.TypeString, Optional: true, Description: "Specify the notification template name",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"outbound_detection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configuration": {
										Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
									},
									"toggle": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable outbound detection; 'disable': Disable outbound detection;",
									},
									"discovery_method": {
										Type: schema.TypeString, Optional: true, Description: "'asn': Autonomous Systems number; 'country': Country;",
									},
									"discovery_record": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Maximum number of top locations",
									},
									"enable_top_k": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"topk_type": {
													Type: schema.TypeString, Optional: true, Description: "'source-subnet': Topk source subnet;",
												},
												"topk_netmask": {
													Type: schema.TypeInt, Optional: true, Default: 128, Description: "Subnet mask. The value should be less than or equal to the minimum zone subnet mask + 8 (IPv6 Subnet mask)",
												},
												"topk_num_records": {
													Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets;",
												},
												"tcp_window_size": {
													Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
												},
												"data_packet_size": {
													Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
												},
												"threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
												},
												"threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
												},
												"threshold_str": {
													Type: schema.TypeString, Optional: true, Description: "Threshold for each geo-location (Non-zero floating point)",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
									"topk_source_subnet": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
									"configuration": {
										Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
									},
									"toggle": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable service discovery; 'disable': Disable service discovery;",
									},
									"pkt_rate_threshold": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "packet rate threshold for discovery (default 10 packets per second)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"packet_anomaly_detection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configuration": {
										Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
									},
									"toggle": {
										Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable packet anomaly; 'disable': Disable packet anomaly;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Required: true, Description: "'port-zero-pkt-rate': Port Zero Packet Rate (default 100 packet per second);",
												},
												"threshold_num": {
													Type: schema.TypeInt, Optional: true, Default: 100, Description: "Threshold for each indicator",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
									"configuration": {
										Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
									},
									"toggle": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable victim IP detection; 'disable': Disable victim IP detection;",
									},
									"histogram_toggle": {
										Type: schema.TypeString, Optional: true, Default: "histogram-disable", Description: "'histogram-enable': Enable histogram statistics of victim IP detection; 'histogram-disable': Disable histogram statistics of victim IP detection;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'reverse-pkt-rate': rate of reverse coming packets; 'fwd-byte-rate': rate of incoming bytes; 'rev-byte-rate': rate of reverse coming bytes;",
												},
												"ip_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for IP",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
			"drop_frag_pkt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
			},
			"enable_top_k": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topk_type": {
							Type: schema.TypeString, Optional: true, Description: "'destination': Topk destination IP;",
						},
						"topk_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
						},
					},
				},
			},
			"force_operational_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force configure operational mode",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
			},
			"hw_blacklist_blocking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Dst side hardware blocking",
						},
						"src_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Src side hardware blocking",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"inbound_forward_dscp": {
				Type: schema.TypeInt, Optional: true, Description: "To set dscp value for inbound packets (DSCP Value for the clear traffic marking)",
			},
			"ip": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address",
						},
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "IP Subnet",
						},
						"expand_ip_subnet": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Expand this subnet to individual IP address",
						},
						"expand_ip_subnet_mode": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Default learning mechanism (Default: Dynamic); 'dynamic': Dynamic learning; 'static': Static learning;",
						},
					},
				},
			},
			"ip_proto": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto_number_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol_num": {
										Type: schema.TypeInt, Required: true, Description: "Protocol Number",
									},
									"manual_mode_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle manual mode to use fix templates",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for this ip-proto",
									},
									"esp_inspect": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auth_algorithm": {
													Type: schema.TypeString, Optional: true, Description: "'AUTH_NULL': No Integrity Check Value; 'HMAC-SHA-1-96': 96 bit Auth Algo; 'HMAC-SHA-256-96': 96 bit Auth Algo; 'HMAC-SHA-256-128': 128 bit Auth Algo; 'HMAC-SHA-384-192': 192 bit Auth Algo; 'HMAC-SHA-512-256': 256 bit Auth Algo; 'HMAC-MD5-96': 96 bit Auth Algo; 'MAC-RIPEMD-160-96': 96 bit Auth Algo;",
												},
												"encrypt_algorithm": {
													Type: schema.TypeString, Optional: true, Description: "'NULL': Null Encryption Algorithm;",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "'transport': Transport mode;",
												},
											},
										},
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
												},
												"action_list": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"per_addr_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
												},
											},
										},
									},
									"drop_frag_pkt": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
									},
									"unlimited_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
									},
									"max_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry",
									},
									"apply_policy_on_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
									},
									"enable_top_k": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k source IP detection",
									},
									"topk_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"enable_top_k_destination": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k destination IP detection",
									},
									"topk_dst_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"set_counter_base_val": {
										Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
									},
									"enable_class_list_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for zone-port or class-list",
									},
									"faster_de_escalation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "De-escalate faster in standalone mode",
									},
									"ip_filtering_policy": {
										Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"ip_filtering_policy_oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"src_based_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_based_policy_name": {
													Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"policy_class_list_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"class_list_name": {
																Type: schema.TypeString, Required: true, Description: "Class-list name",
															},
															"glid": {
																Type: schema.TypeString, Optional: true, Description: "Global limit ID",
															},
															"glid_action": {
																Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
															},
															"log_enable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
															},
															"log_periodic": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
															},
															"max_dynamic_entry_count": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry allowed for this class-list",
															},
															"zone_template": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"logging": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
																		},
																		"ip_proto": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
																		},
																	},
																},
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
															"sampling_enable": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;",
																		},
																	},
																},
															},
															"class_list_overflow_policy_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dummy_name": {
																			Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy for class-list;",
																		},
																		"glid": {
																			Type: schema.TypeString, Optional: true, Description: "Global limit ID",
																		},
																		"action": {
																			Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
																		},
																		"log_enable": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
																		},
																		"log_periodic": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
																		},
																		"zone_template": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"ip_proto": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
																					},
																				},
																			},
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																		"user_tag": {
																			Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
									"dynamic_entry_overflow_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dummy_name": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy;",
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_proto": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
									"level_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_num": {
													Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1; '2': Policy level 2; '3': Policy level 3; '4': Policy level 4;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Zone activation score of this level",
												},
												"zone_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to zone escalate from this level",
												},
												"src_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Source activation score of this level",
												},
												"src_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to source escalate from this level",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_proto": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"indicator_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'frag-rate': rate of incoming fragmented packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
															},
															"data_packet_size": {
																Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
															},
															"score": {
																Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
															},
															"src_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
															},
															"src_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
															},
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
															},
															"zone_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this zone indicator threshold reaches",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
														},
													},
												},
											},
										},
									},
									"manual_mode_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"config": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Manual-mode configuration;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_proto": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
									"port_ind": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"sampling_enable": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"counters1": {
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_rate_adaptive_threshold': Pkt Rate Adaptive Threshold; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_pkt_drop_rate_adaptive_threshold': Pkt Drop Rate Adaptive Threshold; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_syn_rate_adaptive_threshold': TCP SYN Rate Adaptive Threshold; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_fin_rate_adaptive_threshold': TCP FIN Rate Adaptive Threshold; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_rst_rate_adaptive_threshold': TCP RST Rate Adaptive Threshold; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_small_window_ack_rate_adaptive_threshold': TCP Small Window ACK Rate Adaptive Threshold; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_empty_ack_rate_adaptive_threshold': TCP Empty ACK Rate Adaptive Threshold; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_small_payload_rate_adaptive_threshold': TCP Small Payload Rate Adaptive Threshold; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_pkt_drop_ratio_adaptive_threshold': Pkt Drop / Pkt Rcvd Adaptive Threshold; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_inb_per_outb_adaptive_threshold': Bytes-to / Bytes-from Adaptive Threshold; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_syn_per_fin_rate_adaptive_threshold': TCP SYN Rate / FIN Rate Adaptive Threshold; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_conn_miss_rate_adaptive_threshold': TCP Session Miss Rate Adaptive Threshold; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_concurrent_conns_adaptive_threshold': TCP/UDP Concurrent Sessions Adaptive Threshold; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_data_cpu_util_adaptive_threshold': Data CPU Utilization Adaptive Threshold; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_outside_intf_util_adaptive_threshold': Outside Interface Utilization Adaptive Threshold; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_frag_rate_adaptive_threshold': Frag Pkt Rate Adaptive Threshold; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max; 'ddet_ind_bit_rate_adaptive_threshold': Bit Rate Adaptive Threshold;",
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
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"topk_destinations": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"progression_tracking": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for this ip-proto",
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
												},
												"per_addr_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
												},
												"action_list": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
											},
										},
									},
									"drop_frag_pkt": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
									},
									"set_counter_base_val": {
										Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
									},
									"ip_filtering_policy": {
										Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"ip_filtering_policy_oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
									"manual_mode_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle manual mode to use fix templates",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for ip-proto icmp-v4",
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
												},
												"action_list": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"per_addr_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
												},
											},
										},
									},
									"tunnel_decap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable tunnel decapsulation",
									},
									"key_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type: schema.TypeString, Optional: true, Description: "Only decapsulate GRE packet with this key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
												},
											},
										},
									},
									"tunnel_rate_limit": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DDOS-protection on tunnel traffic",
									},
									"drop_frag_pkt": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
									},
									"unlimited_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
									},
									"max_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry",
									},
									"apply_policy_on_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
									},
									"enable_top_k": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k source IP detection",
									},
									"topk_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"enable_top_k_destination": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k destination IP detection",
									},
									"topk_dst_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"set_counter_base_val": {
										Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
									},
									"enable_class_list_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for zone-port or class-list",
									},
									"faster_de_escalation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "De-escalate faster in standalone mode",
									},
									"ip_filtering_policy": {
										Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"ip_filtering_policy_oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"level_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_num": {
													Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1; '2': Policy level 2; '3': Policy level 3; '4': Policy level 4;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Zone activation score of this level",
												},
												"zone_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to zone escalate from this level",
												},
												"src_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Source activation score of this level",
												},
												"src_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to source escalate from this level",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"icmp_v4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
															},
															"icmp_v6": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
															},
															"ip_proto": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"indicator_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'frag-rate': rate of incoming fragmented packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
															},
															"data_packet_size": {
																Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
															},
															"score": {
																Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
															},
															"src_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
															},
															"src_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
															},
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
															},
															"zone_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this zone indicator threshold reaches",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
														},
													},
												},
											},
										},
									},
									"manual_mode_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"config": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Manual-mode configuration;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"icmp_v4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
															},
															"icmp_v6": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
															},
															"ip_proto": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
									"src_based_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_based_policy_name": {
													Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"policy_class_list_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"class_list_name": {
																Type: schema.TypeString, Required: true, Description: "Class-list name",
															},
															"glid": {
																Type: schema.TypeString, Optional: true, Description: "Global limit ID",
															},
															"glid_action": {
																Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
															},
															"log_enable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
															},
															"log_periodic": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
															},
															"max_dynamic_entry_count": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry allowed for this class-list",
															},
															"zone_template": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"logging": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
																		},
																		"icmp_v4": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
																		},
																		"icmp_v6": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
																		},
																		"ip_proto": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
																		},
																		"encap": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
																		},
																	},
																},
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
															"sampling_enable": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;",
																		},
																	},
																},
															},
															"class_list_overflow_policy_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dummy_name": {
																			Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy for class-list;",
																		},
																		"glid": {
																			Type: schema.TypeString, Optional: true, Description: "Global limit ID",
																		},
																		"action": {
																			Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
																		},
																		"log_enable": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
																		},
																		"log_periodic": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
																		},
																		"zone_template": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"icmp_v4": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
																					},
																					"icmp_v6": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
																					},
																					"ip_proto": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
																					},
																					"encap": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
																					},
																				},
																			},
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																		"user_tag": {
																			Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
									"dynamic_entry_overflow_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dummy_name": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy;",
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"icmp_v4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
															},
															"icmp_v6": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
															},
															"ip_proto": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
									"port_ind": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"sampling_enable": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"counters1": {
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_rate_adaptive_threshold': Pkt Rate Adaptive Threshold; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_pkt_drop_rate_adaptive_threshold': Pkt Drop Rate Adaptive Threshold; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_syn_rate_adaptive_threshold': TCP SYN Rate Adaptive Threshold; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_fin_rate_adaptive_threshold': TCP FIN Rate Adaptive Threshold; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_rst_rate_adaptive_threshold': TCP RST Rate Adaptive Threshold; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_small_window_ack_rate_adaptive_threshold': TCP Small Window ACK Rate Adaptive Threshold; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_empty_ack_rate_adaptive_threshold': TCP Empty ACK Rate Adaptive Threshold; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_small_payload_rate_adaptive_threshold': TCP Small Payload Rate Adaptive Threshold; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_pkt_drop_ratio_adaptive_threshold': Pkt Drop / Pkt Rcvd Adaptive Threshold; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_inb_per_outb_adaptive_threshold': Bytes-to / Bytes-from Adaptive Threshold; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_syn_per_fin_rate_adaptive_threshold': TCP SYN Rate / FIN Rate Adaptive Threshold; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_conn_miss_rate_adaptive_threshold': TCP Session Miss Rate Adaptive Threshold; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_concurrent_conns_adaptive_threshold': TCP/UDP Concurrent Sessions Adaptive Threshold; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_data_cpu_util_adaptive_threshold': Data CPU Utilization Adaptive Threshold; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_outside_intf_util_adaptive_threshold': Outside Interface Utilization Adaptive Threshold; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_frag_rate_adaptive_threshold': Frag Pkt Rate Adaptive Threshold; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max; 'ddet_ind_bit_rate_adaptive_threshold': Bit Rate Adaptive Threshold;",
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
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"progression_tracking": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"topk_destinations": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"ipv6": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
						},
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 Subnet",
						},
						"expand_ipv6_subnet": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Expand this subnet to individual IPv6 address",
						},
						"expand_ipv6_subnet_mode": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Default learning mechanism (Default: Dynamic); 'dynamic': Dynamic learning; 'static': Static learning;",
						},
					},
				},
			},
			"is_from_wizard": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Is It Created from Onbox GUI Wizard",
			},
			"log_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"log_high_frequency": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable High frequency logging for non-event logs per zone",
			},
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
			},
			"non_restrictive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Non-restrictive mode ignores Zero Thresholds Indicators",
			},
			"operational_mode": {
				Type: schema.TypeString, Optional: true, Default: "idle", Description: "'idle': Idle mode; 'monitor': Monitor mode; 'learning': Learning mode;",
			},
			"outbound_forward_dscp": {
				Type: schema.TypeInt, Optional: true, Description: "To set dscp value for outbound",
			},
			"outbound_policy": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the outbound policy",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"packet_anomaly_detection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pattern_recognition_hw_filter_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "to enable pattern recognition hardware filter",
			},
			"pattern_recognition_sensitivity": {
				Type: schema.TypeString, Optional: true, Description: "'high': High sensitive pattern recognition; 'medium': Medium sensitive pattern recognition; 'low': Low sensitive pattern recognition;",
			},
			"per_addr_glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
			},
			"port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
									"manual_mode_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle manual mode to use fix templates",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default if default-action-list is not configured); 'ignore': Do nothing for glid exceed;",
												},
												"action_list": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"per_addr_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
												},
											},
										},
									},
									"stateful": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateful tracking of sessions (Default is stateless)",
									},
									"default_action_list": {
										Type: schema.TypeString, Optional: true, Description: "Configure default-action-list",
									},
									"sflow_common": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all sFlow polling options under this zone port",
									},
									"sflow_packets": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow packet-level counter polling",
									},
									"sflow_tcp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sflow_tcp_basic": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow basic TCP counter polling",
												},
												"sflow_tcp_stateful": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow stateful TCP counter polling",
												},
											},
										},
									},
									"sflow_http": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow HTTP counter polling",
									},
									"unlimited_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
									},
									"max_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry",
									},
									"apply_policy_on_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
									},
									"enable_class_list_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for zone-port or class-list",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
									},
									"enable_top_k": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k source IP detection",
									},
									"topk_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"enable_top_k_destination": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k destination IP detection",
									},
									"topk_dst_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"capture_config": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"capture_config_name": {
													Type: schema.TypeString, Optional: true, Description: "Capture-config name",
												},
												"capture_config_mode": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Apply capture-config to dropped packets; 'forward': Apply capture-config to forwarded packets; 'all': Apply capture-config to both dropped and forwarded packets;",
												},
											},
										},
									},
									"set_counter_base_val": {
										Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
									},
									"outbound_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only allow outbound traffic",
									},
									"faster_de_escalation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "De-escalate faster in standalone mode",
									},
									"ip_filtering_policy": {
										Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"signature_extraction": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"algorithm": {
													Type: schema.TypeString, Optional: true, Description: "'heuristic': heuristic algorithm;",
												},
												"manual_mode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable manual mode",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"pattern_recognition": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"algorithm": {
													Type: schema.TypeString, Optional: true, Description: "'heuristic': heuristic algorithm;",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "'capture-never-expire': War-time capture without rate exceeding and never expires; 'manual': Manual mode;",
												},
												"sensitivity": {
													Type: schema.TypeString, Optional: true, Description: "'high': High Sensitivity; 'medium': Medium Sensitivity; 'low': Low Sensitivity;",
												},
												"filter_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Extracted filter threshold",
												},
												"filter_inactive_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Extracted filter inactive threshold",
												},
												"triggered_by": {
													Type: schema.TypeString, Optional: true, Description: "'zone-escalation': Zone escalation trigger pattern recognition; 'packet-rate-exceeds': Packet rate limit exceeds trigger pattern recognition (default);",
												},
												"capture_traffic": {
													Type: schema.TypeString, Optional: true, Description: "'all': Capture all packets; 'dropped': Capture dropped packets (default);",
												},
												"app_payload_offset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set offset of the payload",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"pattern_recognition_pu_details": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"ip_filtering_policy_oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"level_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_num": {
													Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1; '2': Policy level 2; '3': Policy level 3; '4': Policy level 4;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Zone activation score of this level",
												},
												"zone_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to zone escalate from this level",
												},
												"src_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Source activation score of this level",
												},
												"src_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to source escalate from this level",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"quic": {
																Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
															},
															"dns": {
																Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
															},
															"http": {
																Type: schema.TypeString, Optional: true, Description: "DDOS http template",
															},
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
															},
															"sip": {
																Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
															},
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"close_sessions_for_unauth_sources": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for unauthenticated sources",
												},
												"start_signature_extraction": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Start signature extraction from this level",
												},
												"start_pattern_recognition": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Start pattern recognition from this level",
												},
												"apply_extracted_filters": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply extracted filters from this level",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"indicator_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'concurrent-conns': number of concurrent connections; 'conn-miss-rate': rate of incoming packets for which no previously established connection exists; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
															},
															"tcp_window_size": {
																Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
															},
															"data_packet_size": {
																Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
															},
															"score": {
																Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
															},
															"src_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
															},
															"src_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
															},
															"zone_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this zone indicator threshold reaches",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
														},
													},
												},
											},
										},
									},
									"manual_mode_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"config": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Manual-mode configuration;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"quic": {
																Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
															},
															"dns": {
																Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
															},
															"http": {
																Type: schema.TypeString, Optional: true, Description: "DDOS http template",
															},
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
															},
															"sip": {
																Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
															},
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"close_sessions_for_unauth_sources": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for unauthenticated sources",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
									"ips": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"sampling_enable": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"counters1": {
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'ips_matched_total': IPS Matched Total; 'ips_matched_action_pass': IPS Matched Action Pass; 'ips_matched_action_drop': IPS Matched Action Drop; 'ips_matched_action_blacklist': IPS Matched Action Blacklist; 'ips_matched_severity_high': IPS Matched Severity High; 'ips_matched_severity_medium': IPS Matched Severity Medium; 'ips_matched_severity_low': IPS Matched Severity Low; 'src_ips_matched_action_pass': Src IPS Matched Action Pass; 'src_ips_matched_action_drop': Src IPS Matched Action Drop; 'src_ips_matched_action_blacklist': Src IPS Matched Action Blacklist; 'src_ips_matched_severity_high': Src IPS Matched Severity High; 'src_ips_matched_severity_medium': Src IPS Matched Severity Medium; 'src_ips_matched_severity_low': Src IPS Matched Severity Low;",
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
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"sampling_enable": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"counters1": {
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_rate_adaptive_threshold': Pkt Rate Adaptive Threshold; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_pkt_drop_rate_adaptive_threshold': Pkt Drop Rate Adaptive Threshold; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_syn_rate_adaptive_threshold': TCP SYN Rate Adaptive Threshold; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_fin_rate_adaptive_threshold': TCP FIN Rate Adaptive Threshold; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_rst_rate_adaptive_threshold': TCP RST Rate Adaptive Threshold; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_small_window_ack_rate_adaptive_threshold': TCP Small Window ACK Rate Adaptive Threshold; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_empty_ack_rate_adaptive_threshold': TCP Empty ACK Rate Adaptive Threshold; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_small_payload_rate_adaptive_threshold': TCP Small Payload Rate Adaptive Threshold; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_pkt_drop_ratio_adaptive_threshold': Pkt Drop / Pkt Rcvd Adaptive Threshold; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_inb_per_outb_adaptive_threshold': Bytes-to / Bytes-from Adaptive Threshold; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_syn_per_fin_rate_adaptive_threshold': TCP SYN Rate / FIN Rate Adaptive Threshold; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_conn_miss_rate_adaptive_threshold': TCP Session Miss Rate Adaptive Threshold; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_concurrent_conns_adaptive_threshold': TCP/UDP Concurrent Sessions Adaptive Threshold; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_data_cpu_util_adaptive_threshold': Data CPU Utilization Adaptive Threshold; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_outside_intf_util_adaptive_threshold': Outside Interface Utilization Adaptive Threshold; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_frag_rate_adaptive_threshold': Frag Pkt Rate Adaptive Threshold; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max; 'ddet_ind_bit_rate_adaptive_threshold': Bit Rate Adaptive Threshold;",
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
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"progression_tracking": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"topk_destinations": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"src_based_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_based_policy_name": {
													Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"policy_class_list_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"class_list_name": {
																Type: schema.TypeString, Required: true, Description: "Class-list name",
															},
															"glid": {
																Type: schema.TypeString, Optional: true, Description: "Global limit ID",
															},
															"glid_action": {
																Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
															},
															"log_enable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
															},
															"log_periodic": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
															},
															"max_dynamic_entry_count": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry allowed for this class-list",
															},
															"zone_template": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"quic": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
																		},
																		"dns": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
																		},
																		"http": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS http template",
																		},
																		"ssl_l4": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
																		},
																		"sip": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
																		},
																		"tcp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
																		},
																		"udp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
																		},
																		"encap": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
																		},
																		"logging": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
																		},
																	},
																},
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
															"sampling_enable": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;",
																		},
																	},
																},
															},
															"class_list_overflow_policy_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dummy_name": {
																			Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy for class-list;",
																		},
																		"glid": {
																			Type: schema.TypeString, Optional: true, Description: "Global limit ID",
																		},
																		"action": {
																			Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
																		},
																		"log_enable": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
																		},
																		"log_periodic": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
																		},
																		"zone_template": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"quic": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
																					},
																					"dns": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
																					},
																					"http": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS http template",
																					},
																					"ssl_l4": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
																					},
																					"sip": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
																					},
																					"tcp": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
																					},
																					"udp": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
																					},
																					"encap": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
																					},
																					"logging": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
																					},
																				},
																			},
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																		"user_tag": {
																			Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
									"dynamic_entry_overflow_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dummy_name": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy;",
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
												},
												"log_enable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
												},
												"log_periodic": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"quic": {
																Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
															},
															"dns": {
																Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
															},
															"http": {
																Type: schema.TypeString, Optional: true, Description: "DDOS http template",
															},
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
															},
															"sip": {
																Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
															},
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
															"logging": {
																Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
									"manual_mode_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle manual mode to use fix templates",
									},
									"enable_top_k": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k source IP detection",
									},
									"topk_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"enable_top_k_destination": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k destination IP detection",
									},
									"topk_dst_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default if default-action-list is not configured); 'ignore': Do nothing for glid exceed;",
												},
												"action_list": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"per_addr_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
												},
											},
										},
									},
									"stateful": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateful tracking of sessions (Default is stateless)",
									},
									"default_action_list": {
										Type: schema.TypeString, Optional: true, Description: "Configure default-action-list",
									},
									"sflow_common": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all sFlow polling options under this zone port",
									},
									"sflow_packets": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow packet-level counter polling",
									},
									"sflow_tcp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sflow_tcp_basic": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow basic TCP counter polling",
												},
												"sflow_tcp_stateful": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow stateful TCP counter polling",
												},
											},
										},
									},
									"unlimited_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
									},
									"max_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry",
									},
									"apply_policy_on_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
									},
									"set_counter_base_val": {
										Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
									},
									"enable_class_list_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for this zone port or each class-list",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
									},
									"outbound_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only allow outbound traffic",
									},
									"faster_de_escalation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "De-escalate faster in standalone mode",
									},
									"ip_filtering_policy": {
										Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"ip_filtering_policy_oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"pattern_recognition": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"algorithm": {
													Type: schema.TypeString, Optional: true, Description: "'heuristic': heuristic algorithm;",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "'capture-never-expire': War-time capture without rate exceeding and never expires; 'manual': Manual mode;",
												},
												"sensitivity": {
													Type: schema.TypeString, Optional: true, Description: "'high': High Sensitivity; 'medium': Medium Sensitivity; 'low': Low Sensitivity;",
												},
												"filter_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Extracted filter threshold",
												},
												"filter_inactive_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Extracted filter inactive threshold",
												},
												"triggered_by": {
													Type: schema.TypeString, Optional: true, Description: "'zone-escalation': Zone escalation trigger pattern recognition; 'packet-rate-exceeds': Packet rate limit exceeds trigger pattern recognition (default);",
												},
												"capture_traffic": {
													Type: schema.TypeString, Optional: true, Description: "'all': Capture all packets; 'dropped': Capture dropped packets (default);",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"pattern_recognition_pu_details": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"level_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_num": {
													Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1; '2': Policy level 2; '3': Policy level 3; '4': Policy level 4;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Zone activation score of this level",
												},
												"zone_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to zone escalate from this level",
												},
												"src_escalation_score": {
													Type: schema.TypeInt, Optional: true, Description: "Source activation score of this level",
												},
												"src_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to source escalate from this level",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"close_sessions_for_unauth_sources": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for unauthenticated sources",
												},
												"start_pattern_recognition": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Start pattern recognition from this level",
												},
												"apply_extracted_filters": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply extracted filters from this level",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"indicator_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'concurrent-conns': number of concurrent connections; 'conn-miss-rate': rate of incoming packets for which no previously established connection exists; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
															},
															"tcp_window_size": {
																Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
															},
															"data_packet_size": {
																Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
															},
															"score": {
																Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
															},
															"src_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
															},
															"src_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
															},
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
															},
															"zone_violation_actions": {
																Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this zone indicator threshold reaches",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
														},
													},
												},
											},
										},
									},
									"manual_mode_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"config": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Manual-mode configuration;",
												},
												"src_default_glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
														},
													},
												},
												"close_sessions_for_unauth_sources": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for unauthenticated sources",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
									"port_ind": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"sampling_enable": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"counters1": {
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_rate_adaptive_threshold': Pkt Rate Adaptive Threshold; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_pkt_drop_rate_adaptive_threshold': Pkt Drop Rate Adaptive Threshold; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_syn_rate_adaptive_threshold': TCP SYN Rate Adaptive Threshold; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_fin_rate_adaptive_threshold': TCP FIN Rate Adaptive Threshold; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_rst_rate_adaptive_threshold': TCP RST Rate Adaptive Threshold; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_small_window_ack_rate_adaptive_threshold': TCP Small Window ACK Rate Adaptive Threshold; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_empty_ack_rate_adaptive_threshold': TCP Empty ACK Rate Adaptive Threshold; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_small_payload_rate_adaptive_threshold': TCP Small Payload Rate Adaptive Threshold; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_pkt_drop_ratio_adaptive_threshold': Pkt Drop / Pkt Rcvd Adaptive Threshold; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_inb_per_outb_adaptive_threshold': Bytes-to / Bytes-from Adaptive Threshold; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_syn_per_fin_rate_adaptive_threshold': TCP SYN Rate / FIN Rate Adaptive Threshold; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_conn_miss_rate_adaptive_threshold': TCP Session Miss Rate Adaptive Threshold; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_concurrent_conns_adaptive_threshold': TCP/UDP Concurrent Sessions Adaptive Threshold; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_data_cpu_util_adaptive_threshold': Data CPU Utilization Adaptive Threshold; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_outside_intf_util_adaptive_threshold': Outside Interface Utilization Adaptive Threshold; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_frag_rate_adaptive_threshold': Frag Pkt Rate Adaptive Threshold; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max; 'ddet_ind_bit_rate_adaptive_threshold': Bit Rate Adaptive Threshold;",
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
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"progression_tracking": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"topk_destinations": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"src_based_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_based_policy_name": {
													Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"policy_class_list_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"class_list_name": {
																Type: schema.TypeString, Required: true, Description: "Class-list name",
															},
															"glid": {
																Type: schema.TypeString, Optional: true, Description: "Global limit ID",
															},
															"glid_action": {
																Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
															},
															"max_dynamic_entry_count": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry allowed for this class-list",
															},
															"zone_template": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"tcp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
																		},
																		"udp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
																		},
																		"encap": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
																		},
																		"logging": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
																		},
																	},
																},
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
															"sampling_enable": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;",
																		},
																	},
																},
															},
															"class_list_overflow_policy_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dummy_name": {
																			Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy for class-list;",
																		},
																		"glid": {
																			Type: schema.TypeString, Optional: true, Description: "Global limit ID",
																		},
																		"action": {
																			Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
																		},
																		"log_enable": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
																		},
																		"log_periodic": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
																		},
																		"zone_template": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"dns": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
																					},
																					"http": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS http template",
																					},
																					"ssl_l4": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
																					},
																					"sip": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
																					},
																					"tcp": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
																					},
																					"udp": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
																					},
																					"encap": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
																					},
																					"logging": {
																						Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
																					},
																				},
																			},
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																		"user_tag": {
																			Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
									"dynamic_entry_overflow_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dummy_name": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy;",
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
												},
												"log_enable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
												},
												"log_periodic": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dns": {
																Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
															},
															"http": {
																Type: schema.TypeString, Optional: true, Description: "DDOS http template",
															},
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
															},
															"sip": {
																Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
															},
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
															"logging": {
																Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
						"manual_mode_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle manual mode to use fix templates",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID for the whole zone",
									},
									"glid_action": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default if default-action-list is not configured); 'ignore': Do nothing for glid exceed;",
									},
									"action_list": {
										Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
									},
									"per_addr_glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID per address",
									},
								},
							},
						},
						"stateful": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateful tracking of sessions (Default is stateless)",
						},
						"default_action_list": {
							Type: schema.TypeString, Optional: true, Description: "Configure default-action-list",
						},
						"sflow_common": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all sFlow polling options under this zone port",
						},
						"sflow_packets": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow packet-level counter polling",
						},
						"sflow_tcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sflow_tcp_basic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow basic TCP counter polling",
									},
									"sflow_tcp_stateful": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow stateful TCP counter polling",
									},
								},
							},
						},
						"sflow_http": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow HTTP counter polling",
						},
						"unlimited_dynamic_entry_count": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
						},
						"max_dynamic_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry",
						},
						"apply_policy_on_overflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
						},
						"enable_class_list_overflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified under zone port or each class-list",
						},
						"enable_top_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k source IP detection",
						},
						"topk_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
						},
						"enable_top_k_destination": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k destination IP detection",
						},
						"topk_dst_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
						},
						"set_counter_base_val": {
							Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
						},
						"age": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
						},
						"outbound_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only allow outbound traffic",
						},
						"faster_de_escalation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "De-escalate faster in standalone mode",
						},
						"ip_filtering_policy": {
							Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"ip_filtering_policy_oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"pattern_recognition": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"algorithm": {
										Type: schema.TypeString, Optional: true, Description: "'heuristic': heuristic algorithm;",
									},
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "'capture-never-expire': War-time capture without rate exceeding and never expires; 'manual': Manual mode;",
									},
									"sensitivity": {
										Type: schema.TypeString, Optional: true, Description: "'high': High Sensitivity; 'medium': Medium Sensitivity; 'low': Low Sensitivity;",
									},
									"filter_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted filter threshold",
									},
									"filter_inactive_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted filter inactive threshold",
									},
									"triggered_by": {
										Type: schema.TypeString, Optional: true, Description: "'zone-escalation': Zone escalation trigger pattern recognition; 'packet-rate-exceeds': Packet rate limit exceeds trigger pattern recognition (default);",
									},
									"capture_traffic": {
										Type: schema.TypeString, Optional: true, Description: "'all': Capture all packets; 'dropped': Capture dropped packets (default);",
									},
									"app_payload_offset": {
										Type: schema.TypeInt, Optional: true, Description: "Set offset of the payload, default 0",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"pattern_recognition_pu_details": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"level_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level_num": {
										Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1; '2': Policy level 2; '3': Policy level 3; '4': Policy level 4;",
									},
									"src_default_glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"glid_action": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
									},
									"zone_escalation_score": {
										Type: schema.TypeInt, Optional: true, Description: "Zone activation score of this level",
									},
									"zone_violation_actions": {
										Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to zone escalate from this level",
									},
									"src_escalation_score": {
										Type: schema.TypeInt, Optional: true, Description: "Source activation score of this level",
									},
									"src_violation_actions": {
										Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to source escalate from this level",
									},
									"zone_template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"quic": {
													Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
												},
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
												},
												"encap": {
													Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
											},
										},
									},
									"close_sessions_for_unauth_sources": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for unauthenticated sources",
									},
									"start_pattern_recognition": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Start pattern recognition from this level",
									},
									"apply_extracted_filters": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply extracted filters from this level",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'concurrent-conns': number of concurrent connections; 'conn-miss-rate': rate of incoming packets for which no previously established connection exists; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
												},
												"tcp_window_size": {
													Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
												},
												"data_packet_size": {
													Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
												},
												"score": {
													Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
												},
												"src_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
												},
												"src_threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
												},
												"src_threshold_str": {
													Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
												},
												"src_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
												},
												"zone_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
												},
												"zone_threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
												},
												"zone_threshold_str": {
													Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
												},
												"zone_violation_actions": {
													Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this zone indicator threshold reaches",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
								},
							},
						},
						"manual_mode_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config": {
										Type: schema.TypeString, Required: true, Description: "'configuration': Manual-mode configuration;",
									},
									"src_default_glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"glid_action": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
									},
									"zone_template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"quic": {
													Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
												},
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
												},
												"encap": {
													Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
											},
										},
									},
									"close_sessions_for_unauth_sources": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for unauthenticated sources",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
						"ips": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'ips_matched_total': IPS Matched Total; 'ips_matched_action_pass': IPS Matched Action Pass; 'ips_matched_action_drop': IPS Matched Action Drop; 'ips_matched_action_blacklist': IPS Matched Action Blacklist; 'ips_matched_severity_high': IPS Matched Severity High; 'ips_matched_severity_medium': IPS Matched Severity Medium; 'ips_matched_severity_low': IPS Matched Severity Low; 'src_ips_matched_action_pass': Src IPS Matched Action Pass; 'src_ips_matched_action_drop': Src IPS Matched Action Drop; 'src_ips_matched_action_blacklist': Src IPS Matched Action Blacklist; 'src_ips_matched_severity_high': Src IPS Matched Severity High; 'src_ips_matched_severity_medium': Src IPS Matched Severity Medium; 'src_ips_matched_severity_low': Src IPS Matched Severity Low;",
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
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_rate_adaptive_threshold': Pkt Rate Adaptive Threshold; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_pkt_drop_rate_adaptive_threshold': Pkt Drop Rate Adaptive Threshold; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_syn_rate_adaptive_threshold': TCP SYN Rate Adaptive Threshold; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_fin_rate_adaptive_threshold': TCP FIN Rate Adaptive Threshold; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_rst_rate_adaptive_threshold': TCP RST Rate Adaptive Threshold; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_small_window_ack_rate_adaptive_threshold': TCP Small Window ACK Rate Adaptive Threshold; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_empty_ack_rate_adaptive_threshold': TCP Empty ACK Rate Adaptive Threshold; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_small_payload_rate_adaptive_threshold': TCP Small Payload Rate Adaptive Threshold; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_pkt_drop_ratio_adaptive_threshold': Pkt Drop / Pkt Rcvd Adaptive Threshold; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_inb_per_outb_adaptive_threshold': Bytes-to / Bytes-from Adaptive Threshold; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_syn_per_fin_rate_adaptive_threshold': TCP SYN Rate / FIN Rate Adaptive Threshold; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_conn_miss_rate_adaptive_threshold': TCP Session Miss Rate Adaptive Threshold; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_concurrent_conns_adaptive_threshold': TCP/UDP Concurrent Sessions Adaptive Threshold; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_data_cpu_util_adaptive_threshold': Data CPU Utilization Adaptive Threshold; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_outside_intf_util_adaptive_threshold': Outside Interface Utilization Adaptive Threshold; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_frag_rate_adaptive_threshold': Frag Pkt Rate Adaptive Threshold; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max; 'ddet_ind_bit_rate_adaptive_threshold': Bit Rate Adaptive Threshold;",
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
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"topk_destinations": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"progression_tracking": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"src_based_policy_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_based_policy_name": {
										Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"policy_class_list_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list_name": {
													Type: schema.TypeString, Required: true, Description: "Class-list name",
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
												},
												"max_dynamic_entry_count": {
													Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry allowed for this class-list",
												},
												"zone_template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"quic": {
																Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
															},
															"dns": {
																Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
															},
															"http": {
																Type: schema.TypeString, Optional: true, Description: "DDOS http template",
															},
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
															},
															"sip": {
																Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
															},
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
															},
															"encap": {
																Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
															},
															"logging": {
																Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"sampling_enable": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"counters1": {
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;",
															},
														},
													},
												},
												"class_list_overflow_policy_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dummy_name": {
																Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy for class-list;",
															},
															"glid": {
																Type: schema.TypeString, Optional: true, Description: "Global limit ID",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
															},
															"log_enable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
															},
															"log_periodic": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
															},
															"zone_template": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"quic": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
																		},
																		"dns": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
																		},
																		"http": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS http template",
																		},
																		"ssl_l4": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
																		},
																		"sip": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
																		},
																		"tcp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
																		},
																		"udp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
																		},
																		"encap": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
																		},
																		"logging": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
																		},
																	},
																},
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
						"dynamic_entry_overflow_policy_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dummy_name": {
										Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy;",
									},
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
									},
									"log_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
									},
									"log_periodic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
									},
									"zone_template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"quic": {
													Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
												},
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
												},
												"encap": {
													Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"logging": {
													Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"rate_limit": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate limit per second per zone(Default : 1 per second)",
			},
			"reporting_disabled": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Reporting",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'zone_tcp_any_exceed': TCP Dst IP-Proto Rate: Total Exceeded; 'zone_tcp_pkt_rate_exceed': TCP Dst IP-Proto Rate: Packet Exceeded; 'zone_tcp_conn_rate_exceed': TCP Dst IP-Proto Rate: Conn Exceeded; 'zone_udp_any_exceed': UDP Dst IP-Proto Rate: Total Exceeded; 'zone_udp_pkt_rate_exceed': UDP Dst IP-Proto Rate: Packet Exceeded; 'zone_udp_conn_limit_exceed': UDP Dst IP-Proto Limit: Conn Exceeded; 'zone_udp_conn_rate_exceed': UDP Dst IP-Proto Rate: Conn Exceeded; 'zone_icmp_pkt_rate_exceed': ICMP Dst Rate: Packet Exceeded; 'zone_other_pkt_rate_exceed': OTHER Dst IP-Proto Rate: Packet Exceeded; 'zone_other_frag_pkt_rate_exceed': OTHER Dst IP-Proto Rate: Frag Exceeded; 'zone_port_pkt_rate_exceed': Port Rate: Packet Exceeded; 'zone_port_conn_limit_exceed': Port Limit: Conn Exceeded; 'zone_port_conn_rate_exceed': Port Rate: Conn Exceeded; 'zone_pkt_sent': Inbound: Packets Forwarded; 'zone_udp_pkt_sent': UDP Total Packets Forwarded; 'zone_tcp_pkt_sent': TCP Total Packets Forwarded; 'zone_icmp_pkt_sent': ICMP Total Packets Forwarded; 'zone_other_pkt_sent': OTHER Total Packets Forwarded; 'zone_tcp_conn_limit_exceed': TCP Dst IP-Proto Limit: Conn Exceeded; 'zone_tcp_pkt_rcvd': TCP Total Packets Received; 'zone_udp_pkt_rcvd': UDP Total Packets Received; 'zone_icmp_pkt_rcvd': ICMP Total Packets Received; 'zone_other_pkt_rcvd': OTHER Total Packets Received; 'zone_udp_filter_match': UDP Filter Match; 'zone_udp_filter_not_match': UDP Filter Not Matched on Pkt; 'zone_udp_filter_action_blacklist': UDP Filter Action Blacklist; 'zone_udp_filter_action_drop': UDP Filter Action Drop; 'zone_tcp_syn': TCP Total SYN Received; 'zone_tcp_syn_drop': TCP SYN Packets Dropped; 'zone_tcp_src_rate_drop': TCP Src Rate: Total Exceeded; 'zone_udp_src_rate_drop': UDP Src Rate: Total Exceeded; 'zone_icmp_src_rate_drop': ICMP Src Rate: Total Exceeded; 'zone_other_frag_src_rate_drop': OTHER Src Rate: Frag Exceeded; 'zone_other_src_rate_drop': OTHER Src Rate: Total Exceeded; 'zone_tcp_drop': TCP Total Packets Dropped; 'zone_udp_drop': UDP Total Packets Dropped; 'zone_icmp_drop': ICMP Total Packets Dropped; 'zone_frag_drop': Fragmented Packets Dropped; 'zone_other_drop': OTHER Total Packets Dropped; 'zone_tcp_auth': TCP Auth: SYN Cookie Sent; 'zone_udp_filter_action_default_pass': UDP Filter Action Default Pass; 'zone_tcp_filter_match': TCP Filter Match; 'zone_tcp_filter_not_match': TCP Filter Not Matched on Pkt; 'zone_tcp_filter_action_blacklist': TCP Filter Action Blacklist; 'zone_tcp_filter_action_drop': TCP Filter Action Drop; 'zone_tcp_filter_action_default_pass': TCP Filter Action Default Pass; 'zone_udp_filter_action_whitelist': UDP Filter Action WL; 'zone_over_limit_on': Zone overlimit Trigger ON; 'zone_over_limit_off': Zone overlimit Trigger OFF; 'zone_port_over_limit_on': Zone port overlimit Trigger ON; 'zone_port_over_limit_off': Zone port overlimit Trigger OFF; 'zone_over_limit_action': Zone overlimit action; 'zone_port_over_limit_action': Zone port overlimit action; 'scanning_detected_drop': Scanning Detected drop (deprecated); 'scanning_detected_blacklist': Scanning Detected blacklist (deprecated); 'zone_udp_kibit_rate_drop': UDP Dst IP-Proto Rate: KiBit Exceeded; 'zone_tcp_kibit_rate_drop': TCP Dst IP-Proto Rate: KiBit Exceeded; 'zone_icmp_kibit_rate_drop': ICMP Dst Rate: KiBit Exceeded; 'zone_other_kibit_rate_drop': OTHER Dst IP-Proto Rate: KiBit Exceeded; 'zone_port_undef_drop': Dst Port Undefined Dropped; 'zone_port_bl': Dst Port Blacklist Packets Dropped; 'zone_src_port_bl': Dst SrcPort Blacklist Packets Dropped; 'zone_port_kbit_rate_exceed': Port Rate: KiBit Exceeded; 'zone_tcp_src_drop': TCP Src Packets Dropped; 'zone_udp_src_drop': UDP Src Packets Dropped; 'zone_icmp_src_drop': ICMP Src Packets Dropped; 'zone_other_src_drop': OTHER Src Packets Dropped; 'tcp_syn_rcvd': TCP Inbound SYN Received; 'tcp_syn_ack_rcvd': TCP SYN ACK Received; 'tcp_ack_rcvd': TCP ACK Received; 'tcp_fin_rcvd': TCP FIN Received; 'tcp_rst_rcvd': TCP RST Received; 'ingress_bytes': Inbound: Bytes Received; 'egress_bytes': Outbound: Bytes Received; 'ingress_packets': Inbound: Packets Received; 'egress_packets': Outbound: Packets Received; 'tcp_fwd_recv': TCP Inbound Packets Received; 'udp_fwd_recv': UDP Inbound Packets Received; 'icmp_fwd_recv': ICMP Inbound Packets Received; 'tcp_syn_cookie_fail': TCP Auth: SYN Cookie Failed; 'zone_tcp_session_created': TCP Sessions Created; 'zone_udp_session_created': UDP Sessions Created; 'zone_tcp_filter_action_whitelist': TCP Filter Action WL; 'zone_other_filter_match': OTHER Filter Match; 'zone_other_filter_not_match': OTHER Filter Not Matched on Pkt; 'zone_other_filter_action_blacklist': OTHER Filter Action Blacklist; 'zone_other_filter_action_drop': OTHER Filter Action Drop; 'zone_other_filter_action_whitelist': OTHER Filter Action WL; 'zone_other_filter_action_default_pass': OTHER Filter Action Default Pass; 'zone_blackhole_inject': Dst Blackhole Inject; 'zone_blackhole_withdraw': Dst Blackhole Withdraw; 'zone_tcp_out_of_seq_excd': TCP Out-Of-Seq Exceeded; 'zone_tcp_retransmit_excd': TCP Retransmit Exceeded; 'zone_tcp_zero_window_excd': TCP Zero-Window Exceeded; 'zone_tcp_conn_prate_excd': TCP Rate: Conn Pkt Exceeded; 'zone_tcp_action_on_ack_init': TCP Auth: ACK Retry Init; 'zone_tcp_action_on_ack_gap_drop': TCP Auth: ACK Retry Retry-Gap Dropped; 'zone_tcp_action_on_ack_fail': TCP Auth: ACK Retry Dropped; 'zone_tcp_action_on_ack_pass': TCP Auth: ACK Retry Passed; 'zone_tcp_action_on_syn_init': TCP Auth: SYN Retry Init; 'zone_tcp_action_on_syn_gap_drop': TCP Auth: SYN Retry-Gap Dropped; 'zone_tcp_action_on_syn_fail': TCP Auth: SYN Retry Dropped; 'zone_tcp_action_on_syn_pass': TCP Auth: SYN Retry Passed; 'zone_payload_too_small': UDP Payload Too Small; 'zone_payload_too_big': UDP Payload Too Large; 'zone_udp_conn_prate_excd': UDP Rate: Conn Pkt Exceeded; 'zone_udp_ntp_monlist_req': UDP NTP Monlist Request; 'zone_udp_ntp_monlist_resp': UDP NTP Monlist Response; 'zone_udp_wellknown_sport_drop': UDP SrcPort Wellknown; 'zone_udp_retry_init': UDP Auth: Retry Init; 'zone_udp_retry_pass': UDP Auth: Retry Passed; 'zone_tcp_bytes_drop': TCP Total Bytes Dropped; 'zone_udp_bytes_drop': UDP Total Bytes Dropped; 'zone_icmp_bytes_drop': ICMP Total Bytes Dropped; 'zone_other_bytes_drop': OTHER Total Bytes Dropped; 'zone_out_no_route': Dst IPv4/v6 Out No Route; 'outbound_bytes_sent': Outbound: Bytes Forwarded; 'outbound_drop': Outbound: Packets Dropped; 'outbound_bytes_drop': Outbound: Bytes Dropped; 'outbound_pkt_sent': Outbound: Packets Forwarded; 'inbound_bytes_sent': Inbound: Bytes Forwarded; 'inbound_bytes_drop': Inbound: Bytes Dropped; 'zone_src_port_pkt_rate_exceed': SrcPort Rate: Packet Exceeded; 'zone_src_port_kbit_rate_exceed': SrcPort Rate: KiBit Exceeded; 'zone_src_port_conn_limit_exceed': SrcPort Limit: Conn Exceeded; 'zone_src_port_conn_rate_exceed': SrcPort Rate: Conn Exceeded; 'zone_ip_proto_pkt_rate_exceed': IP-Proto Rate: Packet Exceeded; 'zone_ip_proto_kbit_rate_exceed': IP-Proto Rate: KiBit Exceeded; 'zone_tcp_port_any_exceed': TCP Port Rate: Total Exceed; 'zone_udp_port_any_exceed': UDP Port Rate: Total Exceed; 'zone_tcp_auth_pass': TCP Auth: SYN Auth Passed; 'zone_tcp_rst_cookie_fail': TCP Auth: RST Cookie Failed; 'zone_tcp_unauth_drop': TCP Auth: Unauth Dropped; 'src_tcp_syn_auth_fail': Src TCP Auth: SYN Auth Failed; 'src_tcp_syn_cookie_sent': Src TCP Auth: SYN Cookie Sent; 'src_tcp_syn_cookie_fail': Src TCP Auth: SYN Cookie Failed; 'src_tcp_rst_cookie_fail': Src TCP Auth: RST Cookie Failed;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'src_tcp_unauth_drop': Src TCP Auth: Unauth Dropped; 'src_tcp_action_on_syn_init': Src TCP Auth: SYN Retry Init; 'src_tcp_action_on_syn_gap_drop': Src TCP Auth: SYN Retry-Gap Dropped; 'src_tcp_action_on_syn_fail': Src TCP Auth: SYN Retry Dropped; 'src_tcp_action_on_ack_init': Src TCP Auth: ACK Retry Init; 'src_tcp_action_on_ack_gap_drop': Src TCP Auth: ACK Retry Retry-Gap Dropped; 'src_tcp_action_on_ack_fail': Src TCP Auth: ACK Retry Dropped; 'src_tcp_out_of_seq_excd': Src TCP Out-Of-Seq Exceeded; 'src_tcp_retransmit_excd': Src TCP Retransmit Exceeded; 'src_tcp_zero_window_excd': Src TCP Zero-Window Exceeded; 'src_tcp_conn_prate_excd': Src TCP Rate: Conn Pkt Exceeded; 'src_udp_min_payload': Src UDP Payload Too Small; 'src_udp_max_payload': Src UDP Payload Too Large; 'src_udp_conn_prate_excd': Src UDP Rate: Conn Pkt Exceeded; 'src_udp_ntp_monlist_req': Src UDP NTP Monlist Request; 'src_udp_ntp_monlist_resp': Src UDP NTP Monlist Response; 'src_udp_wellknown_sport_drop': Src UDP SrcPort Wellknown; 'src_udp_retry_init': Src UDP Auth: Retry Init; 'dst_udp_retry_gap_drop': UDP Auth: Retry-Gap Dropped; 'dst_udp_retry_fail': UDP Auth: Retry Timeout; 'dst_tcp_session_aged': TCP Sessions Aged; 'dst_udp_session_aged': UDP Sessions Aged; 'dst_tcp_conn_close': TCP Connections Closed; 'dst_tcp_conn_close_half_open': TCP Half Open Connections Closed; 'dst_drop_frag_pkt': Fragmented Packets Dropped; 'src_tcp_filter_action_blacklist': Src TCP Filter Action Blacklist; 'src_tcp_filter_action_whitelist': Src TCP Filter Action WL; 'src_tcp_filter_action_drop': Src TCP Filter Action Drop; 'src_tcp_filter_action_default_pass': Src TCP Filter Action Default Pass; 'src_udp_filter_action_blacklist': Src UDP Filter Action Blacklist; 'src_udp_filter_action_whitelist': Src UDP Filter Action WL; 'src_udp_filter_action_drop': Src UDP Filter Action Drop; 'src_udp_filter_action_default_pass': Src UDP Filter Action Default Pass; 'src_other_filter_action_blacklist': Src OTHER Filter Action Blacklist; 'src_other_filter_action_whitelist': Src OTHER Filter Action WL; 'src_other_filter_action_drop': Src OTHER Filter Action Drop; 'src_other_filter_action_default_pass': Src OTHER Filter Action Default Pass; 'tcp_invalid_syn': TCP Invalid SYN Received; 'dst_tcp_conn_close_w_rst': TCP RST Connections Closed; 'dst_tcp_conn_close_w_fin': TCP FIN Connections Closed; 'dst_tcp_conn_close_w_idle': TCP Idle Connections Closed; 'dst_tcp_conn_create_from_syn': TCP Connections Created From SYN; 'dst_tcp_conn_create_from_ack': TCP Connections Created From ACK; 'src_frag_drop': Src Fragmented Packets Dropped; 'zone_port_kbit_rate_exceed_pkt': Port Rate: KiBit Pkt Exceeded; 'dst_tcp_bytes_rcv': TCP Total Bytes Received; 'dst_udp_bytes_rcv': UDP Total Bytes Received; 'dst_icmp_bytes_rcv': ICMP Total Bytes Received; 'dst_other_bytes_rcv': OTHER Total Bytes Received; 'dst_tcp_bytes_sent': TCP Total Bytes Forwarded; 'dst_udp_bytes_sent': UDP Total Bytes Forwarded; 'dst_icmp_bytes_sent': ICMP Total Bytes Forwarded; 'dst_other_bytes_sent': OTHER Total Bytes Forwarded; 'dst_udp_auth_drop': UDP Auth: Dropped; 'dst_tcp_auth_drop': TCP Auth: Dropped; 'dst_tcp_auth_resp': TCP Auth: Responded; 'dst_drop': Inbound: Packets Dropped; 'dst_entry_pkt_rate_exceed': Entry Rate: Packet Exceeded; 'dst_entry_kbit_rate_exceed': Entry Rate: KiBit Exceeded; 'dst_entry_conn_limit_exceed': Entry Limit: Conn Exceeded; 'dst_entry_conn_rate_exceed': Entry Rate: Conn Exceeded; 'dst_entry_frag_pkt_rate_exceed': Entry Rate: Frag Packet Exceeded; 'dst_l4_tcp_blacklist_drop': Dst TCP IP-Proto Blacklist Dropped; 'dst_l4_udp_blacklist_drop': Dst UDP IP-Proto Blacklist Dropped; 'dst_l4_icmp_blacklist_drop': Dst ICMP IP-Proto Blacklist Dropped; 'dst_l4_other_blacklist_drop': Dst OTHER IP-Proto Blacklist Dropped; 'dst_frag_timeout_drop': Fragment Reassemble Timeout Drop; 'dst_icmp_any_exceed': ICMP Rate: Total Exceed; 'dst_other_any_exceed': OTHER Rate: Total Exceed; 'tcp_rexmit_syn_limit_drop': TCP SYN Retransmit Exceeded Drop; 'tcp_rexmit_syn_limit_bl': TCP SYN Retransmit Exceeded Blacklist; 'dst_clist_overflow_policy_at_learning': Dst Src-Based Overflow Policy Hit; 'zone_frag_rcvd': Fragmented Packets Received; 'zone_tcp_wellknown_sport_drop': TCP SrcPort Wellknown; 'src_tcp_wellknown_sport_drop': Src TCP SrcPort Wellknown; 'secondary_dst_entry_pkt_rate_exceed': Per Addr Rate: Packet Exceeded; 'secondary_dst_entry_kbit_rate_exceed': Per Addr Rate: KiBit Exceeded; 'secondary_dst_entry_conn_limit_exceed': Per Addr Limit: Conn Exceeded; 'secondary_dst_entry_conn_rate_exceed': Per Addr Rate: Conn Exceeded; 'secondary_dst_entry_frag_pkt_rate_exceed': Per Addr Rate: Frag Packet Exceeded; 'src_udp_retry_gap_drop': Src UDP Auth: Retry-Gap Dropped; 'dst_entry_kbit_rate_exceed_count': Entry Rate: KiBit Exceeded Count; 'secondary_entry_learn': Per Addr Entry Learned; 'secondary_entry_hit': Per Addr Entry Hit; 'secondary_entry_miss': Per Addr Entry Missed; 'secondary_entry_aged': Per Addr Entry Aged; 'secondary_entry_learning_thre_exceed': Per Addr Entry Count Overflow; 'zone_port_undef_hit': Dst Port undefined Hit; 'zone_tcp_action_on_ack_timeout': TCP Auth: ACK Retry Timeout; 'zone_tcp_action_on_ack_reset': TCP Auth: ACK Retry Timeout Reset; 'zone_tcp_action_on_ack_blacklist': TCP Auth: ACK Retry Timeout Blacklisted; 'src_tcp_action_on_ack_timeout': Src TCP Auth: ACK Retry Timeout; 'src_tcp_action_on_ack_reset': Src TCP Auth: ACK Retry Timeout Reset; 'src_tcp_action_on_ack_blacklist': Src TCP Auth: ACK Retry Timeout Blacklisted; 'zone_tcp_action_on_syn_timeout': TCP Auth: SYN Retry Timeout; 'zone_tcp_action_on_syn_reset': TCP Auth: SYN Retry Timeout Reset; 'zone_tcp_action_on_syn_blacklist': TCP Auth: SYN Retry Timeout Blacklisted; 'src_tcp_action_on_syn_timeout': Src TCP Auth: SYN Retry Timeout; 'src_tcp_action_on_syn_reset': Src TCP Auth: SYN Retry Timeout Reset; 'src_tcp_action_on_syn_blacklist': Src TCP Auth: SYN Retry Timeout Blacklisted; 'zone_udp_frag_pkt_rate_exceed': UDP Dst IP-Proto Rate: Frag Exceeded; 'zone_udp_frag_src_rate_drop': UDP Src Rate: Frag Exceeded; 'zone_tcp_frag_pkt_rate_exceed': TCP Dst IP-Proto Rate: Frag Exceeded; 'zone_tcp_frag_src_rate_drop': TCP Src Rate: Frag Exceeded; 'zone_icmp_frag_pkt_rate_exceed': ICMP Dst IP-Proto Rate: Frag Exceeded; 'zone_icmp_frag_src_rate_drop': ICMP Src Rate: Frag Exceeded; 'sflow_internal_samples_packed': Sflow Internal Samples Packed; 'sflow_external_samples_packed': Sflow External Samples Packed; 'sflow_internal_packets_sent': Sflow Internal Packets Sent; 'sflow_external_packets_sent': Sflow External Packets Sent; 'dns_outbound_total_query': DNS Outbound Total Query; 'dns_outbound_query_malformed': DNS Outbound Query Malformed; 'dns_outbound_query_resp_chk_failed': DNS Outbound Query Resp Check Failed; 'dns_outbound_query_resp_chk_blacklisted': DNS Outbound Query Resp Check Blacklisted; 'dns_outbound_query_resp_chk_refused_sent': DNS Outbound Query Resp Check REFUSED Sent; 'dns_outbound_query_resp_chk_reset_sent': DNS Outbound Query Resp Check RESET Sent; 'dns_outbound_query_resp_chk_no_resp_sent': DNS Outbound Query Resp Check No Response Sent; 'dns_outbound_query_resp_size_exceed': DNS Outbound Query Response Size Exceed; 'dns_outbound_query_sess_timed_out': DNS Outbound Query Session Timed Out; 'source_entry_total': Source Entry Total Count; 'source_entry_udp': Source Entry UDP Count; 'source_entry_tcp': Source Entry TCP Count; 'source_entry_icmp': Source Entry ICMP Count; 'source_entry_other': Source Entry OTHER Count; 'dst_exceed_action_tunnel': Entry Exceed Action: Tunnel;",
						},
						"counters3": {
							Type: schema.TypeString, Optional: true, Description: "'dst_udp_retry_timeout_blacklist': UDP Auth: Retry Timeout Blacklisted; 'src_udp_auth_timeout': Src UDP Auth: Retry Timeout; 'zone_src_udp_retry_timeout_blacklist': Src UDP Auth: Retry Timeout Blacklisted; 'src_udp_retry_pass': Src UDP Retry Passed; 'secondary_port_learn': Per Addr Port Learned; 'secondary_port_aged': Per Addr Port Aged; 'dst_entry_outbound_udp_session_created': Outbound: UDP Sessions Created; 'dst_entry_outbound_udp_session_aged': Outbound: UDP Sessions Aged; 'dst_entry_outbound_tcp_session_created': Outbound: TCP Sessions Created; 'dst_entry_outbound_tcp_session_aged': Outbound: TCP Sessions Aged; 'dst_entry_outbound_pkt_rate_exceed': Outbound Rate: Packet Exceeded; 'dst_entry_outbound_kbit_rate_exceed': Outbound Rate: KiBit Exceeded; 'dst_entry_outbound_kbit_rate_exceed_count': Outbound Rate: KiBit Exceeded Count; 'dst_entry_outbound_conn_limit_exceed': Outbound Limit: Conn Exceeded; 'dst_entry_outbound_conn_rate_exceed': Outbound Rate: Conn Exceeded; 'dst_entry_outbound_frag_pkt_rate_exceed': Outbound Rate: Frag Packet Exceeded; 'prog_first_req_time_exceed': Req-Resp: First Request Time Exceed; 'prog_req_resp_time_exceed': Req-Resp: Request to Response Time Exceed; 'prog_request_len_exceed': Req-Resp: Request Length Exceed; 'prog_response_len_exceed': Req-Resp: Response Length Exceed; 'prog_resp_req_ratio_exceed': Req-Resp: Response to Request Ratio Exceed; 'prog_resp_req_time_exceed': Req-Resp: Response to Request Time Exceed; 'entry_sync_message_received': Entry Sync Message Received; 'entry_sync_message_sent': Entry Sync Message Sent; 'prog_conn_sent_exceed': Connection: Sent Exceed; 'prog_conn_rcvd_exceed': Connection: Received Exceed; 'prog_conn_time_exceed': Connection: Time Exceed; 'prog_conn_rcvd_sent_ratio_exceed': Connection: Received to Sent Ratio Exceed; 'prog_win_sent_exceed': Time Window: Sent Exceed; 'prog_win_rcvd_exceed': Time Window: Received Exceed; 'prog_win_rcvd_sent_ratio_exceed': Time Window: Received to Sent Exceed; 'prog_exceed_drop': Req-Resp: Violation Exceed Dropped; 'prog_exceed_bl': Req-Resp: Violation Exceed Blacklisted; 'prog_conn_exceed_drop': Connection: Violation Exceed Dropped; 'prog_conn_exceed_bl': Connection: Violation Exceed Blacklisted; 'prog_win_exceed_drop': Time Window: Violation Exceed Dropped; 'prog_win_exceed_bl': Time Window: Violation Exceed Blacklisted; 'east_west_inbound_rcv_pkt': East West: Inbound Packets Received; 'east_west_inbound_drop_pkt': East West: Inbound Packets Dropped; 'east_west_inbound_fwd_pkt': East West: Inbound Packets Forwarded; 'east_west_inbound_rcv_byte': East West: Inbound Bytes Received; 'east_west_inbound_drop_byte': East West: Inbound Bytes Dropped; 'east_west_inbound_fwd_byte': East West: Inbound Bytes Forwarded; 'east_west_outbound_rcv_pkt': East West: Outbound Packets Received; 'east_west_outbound_drop_pkt': East West: Outbound Packets Dropped; 'east_west_outbound_fwd_pkt': East West: Outbound Packets Forwarded; 'east_west_outbound_rcv_byte': East West: Outbound Bytes Received; 'east_west_outbound_drop_byte': East West: Outbound Bytes Dropped; 'east_west_outbound_fwd_byte': East West: Outbound Bytes Forwarded; 'dst_exceed_action_drop': Entry Exceed Action: Dropped; 'prog_conn_samples': Sample Collected: Connection; 'prog_req_samples': Sample Collected: Req-Resp; 'prog_win_samples': Sample Collected: Time Window; 'victim_ip_learned': Victim Identification: IP Entry Learned; 'victim_ip_aged': Victim Identification: IP Entry Aged; 'prog_conn_samples_processed': Sample Processed: Connnection; 'prog_req_samples_processed': Sample Processed: Req-Resp; 'prog_win_samples_processed': Sample Processed: Time Window; 'dst_src_learn_overflow': Src Dynamic Entry Count Overflow; 'dst_tcp_auth_rst': TCP Auth: Reset;",
						},
					},
				},
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"sflow_common": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow counter polling packets, tcp-basic, tcp-stateful and http. WARNING: Zone level Sflow polling might induce heavy CP",
			},
			"sflow_http": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow HTTP counter polling. WARNING: Zone level Sflow polling might induce heavy CPU load depending on the total number",
			},
			"sflow_layer_4": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow Layer 4 counter polling. WARNING: Zone level Sflow polling might induce heavy CPU load depending on the number of",
			},
			"sflow_packets": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow packet-level counter polling. WARNING: Zone level Sflow polling might induce heavy CPU load depending on the total",
			},
			"sflow_tcp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sflow_tcp_basic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow basic TCP counter polling. WARNING: Zone level Sflow polling might induce heavy CPU load depending on the total nu",
						},
						"sflow_tcp_stateful": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow stateful TCP counter polling. WARNING: Zone level Sflow polling might induce heavy CPU load depending on the total",
						},
					},
				},
			},
			"source_nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Configure source NAT",
			},
			"src_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
												},
											},
										},
									},
									"outbound_src_tracking": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
									},
									"zone_template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS udp src template",
												},
												"src_tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS tcp src template",
												},
												"src_dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns src template",
												},
											},
										},
									},
									"default_action_list": {
										Type: schema.TypeString, Optional: true, Description: "Configure default-action-list",
									},
									"set_counter_base_val": {
										Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"port_ind": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"level_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_num": {
													Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1;",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"indicator_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'bit-rate': rate of incoming bits;",
															},
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
												},
											},
										},
									},
									"zone_template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS udp src template",
												},
												"src_tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS tcp src template",
												},
											},
										},
									},
									"default_action_list": {
										Type: schema.TypeString, Optional: true, Description: "Configure default-action-list",
									},
									"set_counter_base_val": {
										Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"port_ind": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"level_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_num": {
													Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1;",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"indicator_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'bit-rate': rate of incoming bits;",
															},
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"glid_action": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
									},
								},
							},
						},
						"zone_template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_udp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS udp src template",
									},
									"src_tcp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS tcp src template",
									},
								},
							},
						},
						"default_action_list": {
							Type: schema.TypeString, Optional: true, Description: "Configure default-action-list",
						},
						"capture_config": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"capture_config_name": {
										Type: schema.TypeString, Optional: true, Description: "Capture-config name",
									},
									"capture_config_mode": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Apply capture-config to dropped packets; 'forward': Apply capture-config to forwarded packets; 'all': Apply capture-config to both dropped and forwarded packets;",
									},
								},
							},
						},
						"set_counter_base_val": {
							Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"port_ind": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"level_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level_num": {
										Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'bit-rate': rate of incoming bits;",
												},
												"zone_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the port-range",
												},
												"zone_threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the port-range",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
			"telemetry_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable from-l3-peer flag for the zone, thus all the ip entries in the zone will be dynamically created/deleted based on the BGP",
			},
			"topk_destinations": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"traffic_distribution_mode": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Distribute traffic to one slot using default distribution mechanism; 'source-ip-based': Distribute traffic between slots, based on source ip;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"web_gui": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type: schema.TypeString, Optional: true, Default: "newly", Description: "'newly': newly; 'learning': learning; 'learned': learned; 'activated': activated;",
						},
						"activated_after_learning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Activate it after learning",
						},
						"create_time": {
							Type: schema.TypeString, Optional: true, Description: "Configure create time",
						},
						"modify_time": {
							Type: schema.TypeString, Optional: true, Description: "Configure modify time",
						},
						"sensitivity": {
							Type: schema.TypeString, Optional: true, Default: "3", Description: "'5': Low; '3': Medium; '1.5': High;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"learning": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"duration": {
										Type: schema.TypeString, Optional: true, Default: "6hour", Description: "'1minute': 1 minute; '6hour': 6 hours; '12hour': 12 hours; '24hour': 24 hours; '7day': 7 days;",
									},
									"starting_time": {
										Type: schema.TypeString, Optional: true, Description: "Configure learning starting time",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"protection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_service_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"port_num": {
																Type: schema.TypeInt, Required: true, Description: "Port Number",
															},
															"protocol": {
																Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port;",
															},
															"pbe": {
																Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
															"pbe": {
																Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
												"proto_name_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"protocol": {
																Type: schema.TypeString, Required: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6;",
															},
															"pbe": {
																Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
													Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port;",
												},
												"pbe": {
													Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"zone_profile": {
				Type: schema.TypeString, Optional: true, Description: "Apply threshold profile",
			},
			"zone_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
						},
					},
				},
			},
		},
	}
}
func resourceDdosDstZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZone(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZone(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZone(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZone(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneCaptureConfigList(d []interface{}) []edpt.DdosDstZoneCaptureConfigList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneCaptureConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneCaptureConfigList
		oi.Name = in["name"].(string)
		oi.Mode = in["mode"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneCollector(d []interface{}) []edpt.DdosDstZoneCollector {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneCollector, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneCollector
		oi.SflowName = in["sflow_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetection254(d []interface{}) edpt.DdosDstZoneDetection254 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetection254
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Settings = in["settings"].(string)
		ret.Toggle = in["toggle"].(string)
		//omit uuid
		ret.Notification = getObjectDdosDstZoneDetectionNotification255(in["notification"].([]interface{}))
		ret.OutboundDetection = getObjectDdosDstZoneDetectionOutboundDetection257(in["outbound_detection"].([]interface{}))
		ret.ServiceDiscovery = getObjectDdosDstZoneDetectionServiceDiscovery261(in["service_discovery"].([]interface{}))
		ret.PacketAnomalyDetection = getObjectDdosDstZoneDetectionPacketAnomalyDetection262(in["packet_anomaly_detection"].([]interface{}))
		ret.VictimIpDetection = getObjectDdosDstZoneDetectionVictimIpDetection264(in["victim_ip_detection"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneDetectionNotification255(d []interface{}) edpt.DdosDstZoneDetectionNotification255 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionNotification255
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Notification = getSliceDdosDstZoneDetectionNotificationNotification256(in["notification"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDstZoneDetectionNotificationNotification256(d []interface{}) []edpt.DdosDstZoneDetectionNotificationNotification256 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionNotificationNotification256, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionNotificationNotification256
		oi.NotificationTemplateName = in["notification_template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionOutboundDetection257(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetection257 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionOutboundDetection257
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		ret.DiscoveryMethod = in["discovery_method"].(string)
		ret.DiscoveryRecord = in["discovery_record"].(int)
		ret.EnableTopK = getSliceDdosDstZoneDetectionOutboundDetectionEnableTopK258(in["enable_top_k"].([]interface{}))
		//omit uuid
		ret.IndicatorList = getSliceDdosDstZoneDetectionOutboundDetectionIndicatorList259(in["indicator_list"].([]interface{}))
		ret.TopkSourceSubnet = getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnet260(in["topk_source_subnet"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionEnableTopK258(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK258 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK258, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK258
		oi.TopkType = in["topk_type"].(string)
		oi.TopkNetmask = in["topk_netmask"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionIndicatorList259(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList259 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList259, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList259
		oi.Type = in["type"].(string)
		oi.TcpWindowSize = in["tcp_window_size"].(int)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.ThresholdNum = in["threshold_num"].(int)
		oi.ThresholdLargeNum = in["threshold_large_num"].(int)
		oi.ThresholdStr = in["threshold_str"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnet260(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet260 {

	var ret edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet260
	return ret
}

func getObjectDdosDstZoneDetectionServiceDiscovery261(d []interface{}) edpt.DdosDstZoneDetectionServiceDiscovery261 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionServiceDiscovery261
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		ret.PktRateThreshold = in["pkt_rate_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZoneDetectionPacketAnomalyDetection262(d []interface{}) edpt.DdosDstZoneDetectionPacketAnomalyDetection262 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionPacketAnomalyDetection262
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		//omit uuid
		ret.IndicatorList = getSliceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorList263(in["indicator_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorList263(d []interface{}) []edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList263 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList263, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList263
		oi.Type = in["type"].(string)
		oi.ThresholdNum = in["threshold_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionVictimIpDetection264(d []interface{}) edpt.DdosDstZoneDetectionVictimIpDetection264 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionVictimIpDetection264
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		ret.HistogramToggle = in["histogram_toggle"].(string)
		//omit uuid
		ret.IndicatorList = getSliceDdosDstZoneDetectionVictimIpDetectionIndicatorList265(in["indicator_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionVictimIpDetectionIndicatorList265(d []interface{}) []edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList265 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList265, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList265
		oi.Type = in["type"].(string)
		oi.IpThresholdNum = in["ip_threshold_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneEnableTopK(d []interface{}) []edpt.DdosDstZoneEnableTopK {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneEnableTopK, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneEnableTopK
		oi.TopkType = in["topk_type"].(string)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneHwBlacklistBlocking266(d []interface{}) edpt.DdosDstZoneHwBlacklistBlocking266 {

	count1 := len(d)
	var ret edpt.DdosDstZoneHwBlacklistBlocking266
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstEnable = in["dst_enable"].(int)
		ret.SrcEnable = in["src_enable"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosDstZoneIp(d []interface{}) []edpt.DdosDstZoneIp {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIp
		oi.IpAddr = in["ip_addr"].(string)
		oi.SubnetIpAddr = in["subnet_ip_addr"].(string)
		oi.ExpandIpSubnet = in["expand_ip_subnet"].(int)
		oi.ExpandIpSubnetMode = in["expand_ip_subnet_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProto267(d []interface{}) edpt.DdosDstZoneIpProto267 {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProto267
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProtoNumberList = getSliceDdosDstZoneIpProtoProtoNumberList(in["proto_number_list"].([]interface{}))
		ret.ProtoTcpUdpList = getSliceDdosDstZoneIpProtoProtoTcpUdpList(in["proto_tcp_udp_list"].([]interface{}))
		ret.ProtoNameList = getSliceDdosDstZoneIpProtoProtoNameList(in["proto_name_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberList
		oi.ProtocolNum = in["protocol_num"].(int)
		oi.ManualModeEnable = in["manual_mode_enable"].(int)
		oi.Deny = in["deny"].(int)
		oi.EspInspect = getObjectDdosDstZoneIpProtoProtoNumberListEspInspect(in["esp_inspect"].([]interface{}))
		oi.GlidCfg = getObjectDdosDstZoneIpProtoProtoNumberListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.DropFragPkt = in["drop_frag_pkt"].(int)
		oi.UnlimitedDynamicEntryCount = in["unlimited_dynamic_entry_count"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.EnableTopKDestination = in["enable_top_k_destination"].(int)
		oi.TopkDstNumRecords = in["topk_dst_num_records"].(int)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.Age = in["age"].(int)
		oi.EnableClassListOverflow = in["enable_class_list_overflow"].(int)
		oi.FasterDeEscalation = in["faster_de_escalation"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoNumberListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.SrcBasedPolicyList = getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyList(in["src_based_policy_list"].([]interface{}))
		oi.DynamicEntryOverflowPolicyList = getSliceDdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyList(in["dynamic_entry_overflow_policy_list"].([]interface{}))
		oi.LevelList = getSliceDdosDstZoneIpProtoProtoNumberListLevelList(in["level_list"].([]interface{}))
		oi.ManualModeList = getSliceDdosDstZoneIpProtoProtoNumberListManualModeList(in["manual_mode_list"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneIpProtoProtoNumberListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZoneIpProtoProtoNumberListTopkSources(in["topk_sources"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZoneIpProtoProtoNumberListTopkDestinations(in["topk_destinations"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZoneIpProtoProtoNumberListProgressionTracking(in["progression_tracking"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListEspInspect(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListEspInspect {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListEspInspect
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthAlgorithm = in["auth_algorithm"].(string)
		ret.EncryptAlgorithm = in["encrypt_algorithm"].(string)
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListGlidCfg(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.ActionList = in["action_list"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListIpFilteringPolicyOper {

	var ret edpt.DdosDstZoneIpProtoProtoNumberListIpFilteringPolicyOper
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListLevelList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneEscalationScore = in["zone_escalation_score"].(int)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		oi.SrcEscalationScore = in["src_escalation_score"].(int)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberListLevelListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneIpProtoProtoNumberListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListLevelListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.Score = in["score"].(int)
		oi.SrcThresholdNum = in["src_threshold_num"].(int)
		oi.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		oi.SrcThresholdStr = in["src_threshold_str"].(string)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		oi.ZoneThresholdStr = in["zone_threshold_str"].(string)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListManualModeList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListManualModeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListManualModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListManualModeList
		oi.Config = in["config"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberListManualModeListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListManualModeListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListManualModeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListManualModeListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListPortInd(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZoneIpProtoProtoNumberListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListTopkSources(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListTopkSources {

	var ret edpt.DdosDstZoneIpProtoProtoNumberListTopkSources
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListTopkDestinations(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListTopkDestinations {

	var ret edpt.DdosDstZoneIpProtoProtoNumberListTopkDestinations
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberListProgressionTracking(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberListProgressionTracking {

	var ret edpt.DdosDstZoneIpProtoProtoNumberListProgressionTracking
	return ret
}

func getSliceDdosDstZoneIpProtoProtoTcpUdpList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoTcpUdpList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoTcpUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoTcpUdpList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZoneIpProtoProtoTcpUdpListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.DropFragPkt = in["drop_frag_pkt"].(int)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoTcpUdpListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpListGlidCfg(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
		ret.ActionList = in["action_list"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpListIpFilteringPolicyOper {

	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpListIpFilteringPolicyOper
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameList
		oi.Protocol = in["protocol"].(string)
		oi.ManualModeEnable = in["manual_mode_enable"].(int)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZoneIpProtoProtoNameListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.TunnelDecap = in["tunnel_decap"].(int)
		oi.KeyCfg = getSliceDdosDstZoneIpProtoProtoNameListKeyCfg(in["key_cfg"].([]interface{}))
		oi.TunnelRateLimit = in["tunnel_rate_limit"].(int)
		oi.DropFragPkt = in["drop_frag_pkt"].(int)
		oi.UnlimitedDynamicEntryCount = in["unlimited_dynamic_entry_count"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.EnableTopKDestination = in["enable_top_k_destination"].(int)
		oi.TopkDstNumRecords = in["topk_dst_num_records"].(int)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.Age = in["age"].(int)
		oi.EnableClassListOverflow = in["enable_class_list_overflow"].(int)
		oi.FasterDeEscalation = in["faster_de_escalation"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoNameListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.LevelList = getSliceDdosDstZoneIpProtoProtoNameListLevelList(in["level_list"].([]interface{}))
		oi.ManualModeList = getSliceDdosDstZoneIpProtoProtoNameListManualModeList(in["manual_mode_list"].([]interface{}))
		oi.SrcBasedPolicyList = getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyList(in["src_based_policy_list"].([]interface{}))
		oi.DynamicEntryOverflowPolicyList = getSliceDdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyList(in["dynamic_entry_overflow_policy_list"].([]interface{}))
		oi.PortInd = getObjectDdosDstZoneIpProtoProtoNameListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZoneIpProtoProtoNameListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZoneIpProtoProtoNameListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZoneIpProtoProtoNameListTopkDestinations(in["topk_destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListGlidCfg(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.ActionList = in["action_list"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListKeyCfg(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListIpFilteringPolicyOper {

	var ret edpt.DdosDstZoneIpProtoProtoNameListIpFilteringPolicyOper
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListLevelList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneEscalationScore = in["zone_escalation_score"].(int)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		oi.SrcEscalationScore = in["src_escalation_score"].(int)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameListLevelListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneIpProtoProtoNameListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListLevelListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameListLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.Score = in["score"].(int)
		oi.SrcThresholdNum = in["src_threshold_num"].(int)
		oi.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		oi.SrcThresholdStr = in["src_threshold_str"].(string)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		oi.ZoneThresholdStr = in["zone_threshold_str"].(string)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListManualModeList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListManualModeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListManualModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListManualModeList
		oi.Config = in["config"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameListManualModeListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListManualModeListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListManualModeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameListManualModeListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListPortInd(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZoneIpProtoProtoNameListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListTopkSources(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListTopkSources {

	var ret edpt.DdosDstZoneIpProtoProtoNameListTopkSources
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListProgressionTracking(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListProgressionTracking {

	var ret edpt.DdosDstZoneIpProtoProtoNameListProgressionTracking
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNameListTopkDestinations(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameListTopkDestinations {

	var ret edpt.DdosDstZoneIpProtoProtoNameListTopkDestinations
	return ret
}

func getSliceDdosDstZoneIpv6(d []interface{}) []edpt.DdosDstZoneIpv6 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpv6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpv6
		oi.Ip6Addr = in["ip6_addr"].(string)
		oi.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		oi.ExpandIpv6Subnet = in["expand_ipv6_subnet"].(int)
		oi.ExpandIpv6SubnetMode = in["expand_ipv6_subnet_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOutboundPolicy268(d []interface{}) edpt.DdosDstZoneOutboundPolicy268 {

	count1 := len(d)
	var ret edpt.DdosDstZoneOutboundPolicy268
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZonePacketAnomalyDetection269(d []interface{}) edpt.DdosDstZonePacketAnomalyDetection269 {

	var ret edpt.DdosDstZonePacketAnomalyDetection269
	return ret
}

func getObjectDdosDstZonePort270(d []interface{}) edpt.DdosDstZonePort270 {

	count1 := len(d)
	var ret edpt.DdosDstZonePort270
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneServiceList = getSliceDdosDstZonePortZoneServiceList(in["zone_service_list"].([]interface{}))
		ret.ZoneServiceOtherList = getSliceDdosDstZonePortZoneServiceOtherList(in["zone_service_other_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceList(d []interface{}) []edpt.DdosDstZonePortZoneServiceList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.ManualModeEnable = in["manual_mode_enable"].(int)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZonePortZoneServiceListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.Stateful = in["stateful"].(int)
		oi.DefaultActionList = in["default_action_list"].(string)
		oi.SflowCommon = in["sflow_common"].(int)
		oi.SflowPackets = in["sflow_packets"].(int)
		oi.SflowTcp = getObjectDdosDstZonePortZoneServiceListSflowTcp(in["sflow_tcp"].([]interface{}))
		oi.SflowHttp = in["sflow_http"].(int)
		oi.UnlimitedDynamicEntryCount = in["unlimited_dynamic_entry_count"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		oi.EnableClassListOverflow = in["enable_class_list_overflow"].(int)
		oi.Age = in["age"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.EnableTopKDestination = in["enable_top_k_destination"].(int)
		oi.TopkDstNumRecords = in["topk_dst_num_records"].(int)
		oi.CaptureConfig = getObjectDdosDstZonePortZoneServiceListCaptureConfig(in["capture_config"].([]interface{}))
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.OutboundOnly = in["outbound_only"].(int)
		oi.FasterDeEscalation = in["faster_de_escalation"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.SignatureExtraction = getObjectDdosDstZonePortZoneServiceListSignatureExtraction(in["signature_extraction"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstZonePortZoneServiceListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstZonePortZoneServiceListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstZonePortZoneServiceListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.LevelList = getSliceDdosDstZonePortZoneServiceListLevelList(in["level_list"].([]interface{}))
		oi.ManualModeList = getSliceDdosDstZonePortZoneServiceListManualModeList(in["manual_mode_list"].([]interface{}))
		oi.Ips = getObjectDdosDstZonePortZoneServiceListIps(in["ips"].([]interface{}))
		oi.PortInd = getObjectDdosDstZonePortZoneServiceListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZonePortZoneServiceListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZonePortZoneServiceListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZonePortZoneServiceListTopkDestinations(in["topk_destinations"].([]interface{}))
		oi.SrcBasedPolicyList = getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyList(in["src_based_policy_list"].([]interface{}))
		oi.DynamicEntryOverflowPolicyList = getSliceDdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyList(in["dynamic_entry_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListGlidCfg(d []interface{}) edpt.DdosDstZonePortZoneServiceListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.ActionList = in["action_list"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListSflowTcp(d []interface{}) edpt.DdosDstZonePortZoneServiceListSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListCaptureConfig(d []interface{}) edpt.DdosDstZonePortZoneServiceListCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListSignatureExtraction(d []interface{}) edpt.DdosDstZonePortZoneServiceListSignatureExtraction {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListSignatureExtraction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.ManualMode = in["manual_mode"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListPatternRecognition(d []interface{}) edpt.DdosDstZonePortZoneServiceListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.Mode = in["mode"].(string)
		ret.Sensitivity = in["sensitivity"].(string)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterInactiveThreshold = in["filter_inactive_threshold"].(int)
		ret.TriggeredBy = in["triggered_by"].(string)
		ret.CaptureTraffic = in["capture_traffic"].(string)
		ret.AppPayloadOffset = in["app_payload_offset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZonePortZoneServiceListPatternRecognitionPuDetails {

	var ret edpt.DdosDstZonePortZoneServiceListPatternRecognitionPuDetails
	return ret
}

func getObjectDdosDstZonePortZoneServiceListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZonePortZoneServiceListIpFilteringPolicyOper {

	var ret edpt.DdosDstZonePortZoneServiceListIpFilteringPolicyOper
	return ret
}

func getSliceDdosDstZonePortZoneServiceListLevelList(d []interface{}) []edpt.DdosDstZonePortZoneServiceListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneEscalationScore = in["zone_escalation_score"].(int)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		oi.SrcEscalationScore = in["src_escalation_score"].(int)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceListLevelListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		oi.StartSignatureExtraction = in["start_signature_extraction"].(int)
		oi.StartPatternRecognition = in["start_pattern_recognition"].(int)
		oi.ApplyExtractedFilters = in["apply_extracted_filters"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZonePortZoneServiceListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListLevelListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceListLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZonePortZoneServiceListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.TcpWindowSize = in["tcp_window_size"].(int)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.Score = in["score"].(int)
		oi.SrcThresholdNum = in["src_threshold_num"].(int)
		oi.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		oi.SrcThresholdStr = in["src_threshold_str"].(string)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdStr = in["zone_threshold_str"].(string)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListManualModeList(d []interface{}) []edpt.DdosDstZonePortZoneServiceListManualModeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListManualModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListManualModeList
		oi.Config = in["config"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceListManualModeListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListManualModeListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceListManualModeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListManualModeListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListIps(d []interface{}) edpt.DdosDstZonePortZoneServiceListIps {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListIps
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZonePortZoneServiceListIpsSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListIpsSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceListIpsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListIpsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListIpsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListPortInd(d []interface{}) edpt.DdosDstZonePortZoneServiceListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZonePortZoneServiceListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListTopkSources(d []interface{}) edpt.DdosDstZonePortZoneServiceListTopkSources {

	var ret edpt.DdosDstZonePortZoneServiceListTopkSources
	return ret
}

func getObjectDdosDstZonePortZoneServiceListProgressionTracking(d []interface{}) edpt.DdosDstZonePortZoneServiceListProgressionTracking {

	var ret edpt.DdosDstZonePortZoneServiceListProgressionTracking
	return ret
}

func getObjectDdosDstZonePortZoneServiceListTopkDestinations(d []interface{}) edpt.DdosDstZonePortZoneServiceListTopkDestinations {

	var ret edpt.DdosDstZonePortZoneServiceListTopkDestinations
	return ret
}

func getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherList
		oi.PortOther = in["port_other"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.ManualModeEnable = in["manual_mode_enable"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.EnableTopKDestination = in["enable_top_k_destination"].(int)
		oi.TopkDstNumRecords = in["topk_dst_num_records"].(int)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZonePortZoneServiceOtherListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.Stateful = in["stateful"].(int)
		oi.DefaultActionList = in["default_action_list"].(string)
		oi.SflowCommon = in["sflow_common"].(int)
		oi.SflowPackets = in["sflow_packets"].(int)
		oi.SflowTcp = getObjectDdosDstZonePortZoneServiceOtherListSflowTcp(in["sflow_tcp"].([]interface{}))
		oi.UnlimitedDynamicEntryCount = in["unlimited_dynamic_entry_count"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.EnableClassListOverflow = in["enable_class_list_overflow"].(int)
		oi.Age = in["age"].(int)
		oi.OutboundOnly = in["outbound_only"].(int)
		oi.FasterDeEscalation = in["faster_de_escalation"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.IpFilteringPolicyOper = getObjectDdosDstZonePortZoneServiceOtherListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstZonePortZoneServiceOtherListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstZonePortZoneServiceOtherListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		oi.LevelList = getSliceDdosDstZonePortZoneServiceOtherListLevelList(in["level_list"].([]interface{}))
		oi.ManualModeList = getSliceDdosDstZonePortZoneServiceOtherListManualModeList(in["manual_mode_list"].([]interface{}))
		oi.PortInd = getObjectDdosDstZonePortZoneServiceOtherListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZonePortZoneServiceOtherListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZonePortZoneServiceOtherListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZonePortZoneServiceOtherListTopkDestinations(in["topk_destinations"].([]interface{}))
		oi.SrcBasedPolicyList = getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyList(in["src_based_policy_list"].([]interface{}))
		oi.DynamicEntryOverflowPolicyList = getSliceDdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyList(in["dynamic_entry_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListGlidCfg(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.ActionList = in["action_list"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListSflowTcp(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListIpFilteringPolicyOper {

	var ret edpt.DdosDstZonePortZoneServiceOtherListIpFilteringPolicyOper
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListPatternRecognition(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.Mode = in["mode"].(string)
		ret.Sensitivity = in["sensitivity"].(string)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterInactiveThreshold = in["filter_inactive_threshold"].(int)
		ret.TriggeredBy = in["triggered_by"].(string)
		ret.CaptureTraffic = in["capture_traffic"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListPatternRecognitionPuDetails {

	var ret edpt.DdosDstZonePortZoneServiceOtherListPatternRecognitionPuDetails
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListLevelList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneEscalationScore = in["zone_escalation_score"].(int)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		oi.SrcEscalationScore = in["src_escalation_score"].(int)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceOtherListLevelListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		oi.StartPatternRecognition = in["start_pattern_recognition"].(int)
		oi.ApplyExtractedFilters = in["apply_extracted_filters"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZonePortZoneServiceOtherListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListLevelListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.TcpWindowSize = in["tcp_window_size"].(int)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.Score = in["score"].(int)
		oi.SrcThresholdNum = in["src_threshold_num"].(int)
		oi.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		oi.SrcThresholdStr = in["src_threshold_str"].(string)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		oi.ZoneThresholdStr = in["zone_threshold_str"].(string)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListManualModeList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListManualModeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListManualModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListManualModeList
		oi.Config = in["config"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceOtherListManualModeListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListManualModeListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListManualModeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListManualModeListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListPortInd(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZonePortZoneServiceOtherListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListTopkSources(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListTopkSources {

	var ret edpt.DdosDstZonePortZoneServiceOtherListTopkSources
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListProgressionTracking(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListProgressionTracking {

	var ret edpt.DdosDstZonePortZoneServiceOtherListProgressionTracking
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListTopkDestinations(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListTopkDestinations {

	var ret edpt.DdosDstZonePortZoneServiceOtherListTopkDestinations
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeList(d []interface{}) []edpt.DdosDstZonePortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeList
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.ManualModeEnable = in["manual_mode_enable"].(int)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZonePortRangeListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.Stateful = in["stateful"].(int)
		oi.DefaultActionList = in["default_action_list"].(string)
		oi.SflowCommon = in["sflow_common"].(int)
		oi.SflowPackets = in["sflow_packets"].(int)
		oi.SflowTcp = getObjectDdosDstZonePortRangeListSflowTcp(in["sflow_tcp"].([]interface{}))
		oi.SflowHttp = in["sflow_http"].(int)
		oi.UnlimitedDynamicEntryCount = in["unlimited_dynamic_entry_count"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		oi.EnableClassListOverflow = in["enable_class_list_overflow"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.EnableTopKDestination = in["enable_top_k_destination"].(int)
		oi.TopkDstNumRecords = in["topk_dst_num_records"].(int)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.Age = in["age"].(int)
		oi.OutboundOnly = in["outbound_only"].(int)
		oi.FasterDeEscalation = in["faster_de_escalation"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IpFilteringPolicyOper = getObjectDdosDstZonePortRangeListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstZonePortRangeListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstZonePortRangeListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		oi.LevelList = getSliceDdosDstZonePortRangeListLevelList(in["level_list"].([]interface{}))
		oi.ManualModeList = getSliceDdosDstZonePortRangeListManualModeList(in["manual_mode_list"].([]interface{}))
		oi.Ips = getObjectDdosDstZonePortRangeListIps(in["ips"].([]interface{}))
		oi.PortInd = getObjectDdosDstZonePortRangeListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstZonePortRangeListTopkSources(in["topk_sources"].([]interface{}))
		oi.TopkDestinations = getObjectDdosDstZonePortRangeListTopkDestinations(in["topk_destinations"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstZonePortRangeListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.SrcBasedPolicyList = getSliceDdosDstZonePortRangeListSrcBasedPolicyList(in["src_based_policy_list"].([]interface{}))
		oi.DynamicEntryOverflowPolicyList = getSliceDdosDstZonePortRangeListDynamicEntryOverflowPolicyList(in["dynamic_entry_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListGlidCfg(d []interface{}) edpt.DdosDstZonePortRangeListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.ActionList = in["action_list"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListSflowTcp(d []interface{}) edpt.DdosDstZonePortRangeListSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZonePortRangeListIpFilteringPolicyOper {

	var ret edpt.DdosDstZonePortRangeListIpFilteringPolicyOper
	return ret
}

func getObjectDdosDstZonePortRangeListPatternRecognition(d []interface{}) edpt.DdosDstZonePortRangeListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.Mode = in["mode"].(string)
		ret.Sensitivity = in["sensitivity"].(string)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterInactiveThreshold = in["filter_inactive_threshold"].(int)
		ret.TriggeredBy = in["triggered_by"].(string)
		ret.CaptureTraffic = in["capture_traffic"].(string)
		ret.AppPayloadOffset = in["app_payload_offset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZonePortRangeListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstZonePortRangeListPatternRecognitionPuDetails {

	var ret edpt.DdosDstZonePortRangeListPatternRecognitionPuDetails
	return ret
}

func getSliceDdosDstZonePortRangeListLevelList(d []interface{}) []edpt.DdosDstZonePortRangeListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneEscalationScore = in["zone_escalation_score"].(int)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		oi.SrcEscalationScore = in["src_escalation_score"].(int)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortRangeListLevelListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		oi.StartPatternRecognition = in["start_pattern_recognition"].(int)
		oi.ApplyExtractedFilters = in["apply_extracted_filters"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZonePortRangeListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListLevelListZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeListLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZonePortRangeListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.TcpWindowSize = in["tcp_window_size"].(int)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.Score = in["score"].(int)
		oi.SrcThresholdNum = in["src_threshold_num"].(int)
		oi.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		oi.SrcThresholdStr = in["src_threshold_str"].(string)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		oi.ZoneThresholdStr = in["zone_threshold_str"].(string)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeListManualModeList(d []interface{}) []edpt.DdosDstZonePortRangeListManualModeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListManualModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListManualModeList
		oi.Config = in["config"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortRangeListManualModeListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListManualModeListZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeListManualModeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListManualModeListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListIps(d []interface{}) edpt.DdosDstZonePortRangeListIps {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListIps
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZonePortRangeListIpsSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeListIpsSamplingEnable(d []interface{}) []edpt.DdosDstZonePortRangeListIpsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListIpsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListIpsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListPortInd(d []interface{}) edpt.DdosDstZonePortRangeListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZonePortRangeListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstZonePortRangeListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListTopkSources(d []interface{}) edpt.DdosDstZonePortRangeListTopkSources {

	var ret edpt.DdosDstZonePortRangeListTopkSources
	return ret
}

func getObjectDdosDstZonePortRangeListTopkDestinations(d []interface{}) edpt.DdosDstZonePortRangeListTopkDestinations {

	var ret edpt.DdosDstZonePortRangeListTopkDestinations
	return ret
}

func getObjectDdosDstZonePortRangeListProgressionTracking(d []interface{}) edpt.DdosDstZonePortRangeListProgressionTracking {

	var ret edpt.DdosDstZonePortRangeListProgressionTracking
	return ret
}

func getSliceDdosDstZonePortRangeListSrcBasedPolicyList(d []interface{}) []edpt.DdosDstZonePortRangeListSrcBasedPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListSrcBasedPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListSrcBasedPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeListDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortRangeListDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeListDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeListDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortRangeListDynamicEntryOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeListDynamicEntryOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeListDynamicEntryOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeListDynamicEntryOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZoneSamplingEnable(d []interface{}) []edpt.DdosDstZoneSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		oi.Counters3 = in["counters3"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSflowTcp(d []interface{}) edpt.DdosDstZoneSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstZoneSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstZoneSrcPort271(d []interface{}) edpt.DdosDstZoneSrcPort271 {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPort271
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneSrcPortList = getSliceDdosDstZoneSrcPortZoneSrcPortList(in["zone_src_port_list"].([]interface{}))
		ret.ZoneSrcPortOtherList = getSliceDdosDstZoneSrcPortZoneSrcPortOtherList(in["zone_src_port_other_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZoneSrcPortZoneSrcPortListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.OutboundSrcTracking = in["outbound_src_tracking"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneSrcPortZoneSrcPortListZoneTemplate(in["zone_template"].([]interface{}))
		oi.DefaultActionList = in["default_action_list"].(string)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		//omit uuid
		oi.PortInd = getObjectDdosDstZoneSrcPortZoneSrcPortListPortInd(in["port_ind"].([]interface{}))
		oi.LevelList = getSliceDdosDstZoneSrcPortZoneSrcPortListLevelList(in["level_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortListGlidCfg(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortListZoneTemplate(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
		ret.SrcDns = in["src_dns"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortListPortInd(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortListPortInd {

	var ret edpt.DdosDstZoneSrcPortZoneSrcPortListPortInd
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortListLevelList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortListLevelList
		oi.LevelNum = in["level_num"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneSrcPortZoneSrcPortListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherList
		oi.PortOther = in["port_other"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZoneSrcPortZoneSrcPortOtherListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.ZoneTemplate = getObjectDdosDstZoneSrcPortZoneSrcPortOtherListZoneTemplate(in["zone_template"].([]interface{}))
		oi.DefaultActionList = in["default_action_list"].(string)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		//omit uuid
		oi.PortInd = getObjectDdosDstZoneSrcPortZoneSrcPortOtherListPortInd(in["port_ind"].([]interface{}))
		oi.LevelList = getSliceDdosDstZoneSrcPortZoneSrcPortOtherListLevelList(in["level_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherListGlidCfg(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherListZoneTemplate(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherListPortInd(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherListPortInd {

	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherListPortInd
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherListLevelList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherListLevelList
		oi.LevelNum = in["level_num"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneSrcPortZoneSrcPortOtherListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortRangeList(d []interface{}) []edpt.DdosDstZoneSrcPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortRangeList
		oi.SrcPortRangeStart = in["src_port_range_start"].(int)
		oi.SrcPortRangeEnd = in["src_port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.GlidCfg = getObjectDdosDstZoneSrcPortRangeListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.ZoneTemplate = getObjectDdosDstZoneSrcPortRangeListZoneTemplate(in["zone_template"].([]interface{}))
		oi.DefaultActionList = in["default_action_list"].(string)
		oi.CaptureConfig = getObjectDdosDstZoneSrcPortRangeListCaptureConfig(in["capture_config"].([]interface{}))
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PortInd = getObjectDdosDstZoneSrcPortRangeListPortInd(in["port_ind"].([]interface{}))
		oi.LevelList = getSliceDdosDstZoneSrcPortRangeListLevelList(in["level_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortRangeListGlidCfg(d []interface{}) edpt.DdosDstZoneSrcPortRangeListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortRangeListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortRangeListZoneTemplate(d []interface{}) edpt.DdosDstZoneSrcPortRangeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortRangeListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortRangeListCaptureConfig(d []interface{}) edpt.DdosDstZoneSrcPortRangeListCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortRangeListCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortRangeListPortInd(d []interface{}) edpt.DdosDstZoneSrcPortRangeListPortInd {

	var ret edpt.DdosDstZoneSrcPortRangeListPortInd
	return ret
}

func getSliceDdosDstZoneSrcPortRangeListLevelList(d []interface{}) []edpt.DdosDstZoneSrcPortRangeListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortRangeListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortRangeListLevelList
		oi.LevelNum = in["level_num"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneSrcPortRangeListLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortRangeListLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneSrcPortRangeListLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortRangeListLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortRangeListLevelListIndicatorList
		oi.Type = in["type"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneTopkDestinations272(d []interface{}) edpt.DdosDstZoneTopkDestinations272 {

	var ret edpt.DdosDstZoneTopkDestinations272
	return ret
}

func getObjectDdosDstZoneWebGui273(d []interface{}) edpt.DdosDstZoneWebGui273 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGui273
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(string)
		ret.ActivatedAfterLearning = in["activated_after_learning"].(int)
		ret.CreateTime = in["create_time"].(string)
		ret.ModifyTime = in["modify_time"].(string)
		ret.Sensitivity = in["sensitivity"].(string)
		//omit uuid
		ret.Learning = getObjectDdosDstZoneWebGuiLearning274(in["learning"].([]interface{}))
		ret.Protection = getObjectDdosDstZoneWebGuiProtection275(in["protection"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneWebGuiLearning274(d []interface{}) edpt.DdosDstZoneWebGuiLearning274 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiLearning274
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Duration = in["duration"].(string)
		ret.StartingTime = in["starting_time"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZoneWebGuiProtection275(d []interface{}) edpt.DdosDstZoneWebGuiProtection275 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtection275
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = getObjectDdosDstZoneWebGuiProtectionPort276(in["port"].([]interface{}))
		ret.IpProto = getObjectDdosDstZoneWebGuiProtectionIpProto279(in["ip_proto"].([]interface{}))
		ret.PortRangeList = getSliceDdosDstZoneWebGuiProtectionPortRangeList281(in["port_range_list"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneWebGuiProtectionPort276(d []interface{}) edpt.DdosDstZoneWebGuiProtectionPort276 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtectionPort276
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneServiceList = getSliceDdosDstZoneWebGuiProtectionPortZoneServiceList277(in["zone_service_list"].([]interface{}))
		ret.ZoneServiceOtherList = getSliceDdosDstZoneWebGuiProtectionPortZoneServiceOtherList278(in["zone_service_other_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortZoneServiceList277(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList277 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList277, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList277
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortZoneServiceOtherList278(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList278 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList278, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList278
		oi.PortOther = in["port_other"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneWebGuiProtectionIpProto279(d []interface{}) edpt.DdosDstZoneWebGuiProtectionIpProto279 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtectionIpProto279
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProtoNameList = getSliceDdosDstZoneWebGuiProtectionIpProtoProtoNameList280(in["proto_name_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionIpProtoProtoNameList280(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList280 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList280, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList280
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortRangeList281(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortRangeList281 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortRangeList281, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortRangeList281
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneZoneTemplate(d []interface{}) edpt.DdosDstZoneZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZone(d *schema.ResourceData) edpt.DdosDstZone {
	var ret edpt.DdosDstZone
	ret.Inst.ActionList = d.Get("action_list").(string)
	ret.Inst.AdvertisedEnable = d.Get("advertised_enable").(int)
	ret.Inst.CaptureConfigList = getSliceDdosDstZoneCaptureConfigList(d.Get("capture_config_list").([]interface{}))
	ret.Inst.Collector = getSliceDdosDstZoneCollector(d.Get("collector").([]interface{}))
	ret.Inst.ContinuousLearning = d.Get("continuous_learning").(int)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.DestNatIp = d.Get("dest_nat_ip").(string)
	ret.Inst.DestNatIpv6 = d.Get("dest_nat_ipv6").(string)
	ret.Inst.Detection = getObjectDdosDstZoneDetection254(d.Get("detection").([]interface{}))
	ret.Inst.DropFragPkt = d.Get("drop_frag_pkt").(int)
	ret.Inst.EnableTopK = getSliceDdosDstZoneEnableTopK(d.Get("enable_top_k").([]interface{}))
	ret.Inst.ForceOperationalMode = d.Get("force_operational_mode").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.HwBlacklistBlocking = getObjectDdosDstZoneHwBlacklistBlocking266(d.Get("hw_blacklist_blocking").([]interface{}))
	ret.Inst.InboundForwardDscp = d.Get("inbound_forward_dscp").(int)
	ret.Inst.Ip = getSliceDdosDstZoneIp(d.Get("ip").([]interface{}))
	ret.Inst.IpProto = getObjectDdosDstZoneIpProto267(d.Get("ip_proto").([]interface{}))
	ret.Inst.Ipv6 = getSliceDdosDstZoneIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.IsFromWizard = d.Get("is_from_wizard").(int)
	ret.Inst.LogEnable = d.Get("log_enable").(int)
	ret.Inst.LogHighFrequency = d.Get("log_high_frequency").(int)
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.NonRestrictive = d.Get("non_restrictive").(int)
	ret.Inst.OperationalMode = d.Get("operational_mode").(string)
	ret.Inst.OutboundForwardDscp = d.Get("outbound_forward_dscp").(int)
	ret.Inst.OutboundPolicy = getObjectDdosDstZoneOutboundPolicy268(d.Get("outbound_policy").([]interface{}))
	ret.Inst.PacketAnomalyDetection = getObjectDdosDstZonePacketAnomalyDetection269(d.Get("packet_anomaly_detection").([]interface{}))
	ret.Inst.PatternRecognitionHwFilterEnable = d.Get("pattern_recognition_hw_filter_enable").(int)
	ret.Inst.PatternRecognitionSensitivity = d.Get("pattern_recognition_sensitivity").(string)
	ret.Inst.PerAddrGlid = d.Get("per_addr_glid").(string)
	ret.Inst.Port = getObjectDdosDstZonePort270(d.Get("port").([]interface{}))
	ret.Inst.PortRangeList = getSliceDdosDstZonePortRangeList(d.Get("port_range_list").([]interface{}))
	ret.Inst.RateLimit = d.Get("rate_limit").(int)
	ret.Inst.ReportingDisabled = d.Get("reporting_disabled").(int)
	ret.Inst.SamplingEnable = getSliceDdosDstZoneSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.SflowCommon = d.Get("sflow_common").(int)
	ret.Inst.SflowHttp = d.Get("sflow_http").(int)
	ret.Inst.SflowLayer4 = d.Get("sflow_layer_4").(int)
	ret.Inst.SflowPackets = d.Get("sflow_packets").(int)
	ret.Inst.SflowTcp = getObjectDdosDstZoneSflowTcp(d.Get("sflow_tcp").([]interface{}))
	ret.Inst.SourceNatPool = d.Get("source_nat_pool").(string)
	ret.Inst.SrcPort = getObjectDdosDstZoneSrcPort271(d.Get("src_port").([]interface{}))
	ret.Inst.SrcPortRangeList = getSliceDdosDstZoneSrcPortRangeList(d.Get("src_port_range_list").([]interface{}))
	ret.Inst.TelemetryEnable = d.Get("telemetry_enable").(int)
	ret.Inst.TopkDestinations = getObjectDdosDstZoneTopkDestinations272(d.Get("topk_destinations").([]interface{}))
	ret.Inst.TrafficDistributionMode = d.Get("traffic_distribution_mode").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.WebGui = getObjectDdosDstZoneWebGui273(d.Get("web_gui").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ZoneProfile = d.Get("zone_profile").(string)
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneZoneTemplate(d.Get("zone_template").([]interface{}))
	return ret
}
