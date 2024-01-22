package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumber() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_number`: DDOS IP protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNumberCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNumberUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNumberRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNumberDelete,

		Schema: map[string]*schema.Schema{
			"age": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
			},
			"apply_policy_on_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
			},
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for this ip-proto",
			},
			"drop_frag_pkt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
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
			"enable_class_list_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for zone-port or class-list",
			},
			"enable_top_k": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k source IP detection",
			},
			"enable_top_k_destination": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k destination IP detection",
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
			"faster_de_escalation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "De-escalate faster in standalone mode",
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
			"ip_filtering_policy": {
				Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
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
			"manual_mode_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle manual mode to use fix templates",
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
			"max_dynamic_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry",
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
			"protocol_num": {
				Type: schema.TypeInt, Required: true, Description: "Protocol Number",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
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
			"topk_dst_num_records": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
			},
			"topk_num_records": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
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
			"unlimited_dynamic_entry_count": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneIpProtoProtoNumberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumber(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumber(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNumberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumber(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumber(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberEspInspect(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberEspInspect {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberEspInspect
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthAlgorithm = in["auth_algorithm"].(string)
		ret.EncryptAlgorithm = in["encrypt_algorithm"].(string)
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberGlidCfg(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.ActionList = in["action_list"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberIpFilteringPolicyOper205(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOper205 {

	var ret edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOper205
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberLevelList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneEscalationScore = in["zone_escalation_score"].(int)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		oi.SrcEscalationScore = in["src_escalation_score"].(int)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberLevelListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZoneIpProtoProtoNumberLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberLevelListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberLevelListIndicatorList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberLevelListIndicatorList
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

func getSliceDdosDstZoneIpProtoProtoNumberManualModeList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberManualModeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberManualModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberManualModeList
		oi.Config = in["config"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberManualModeListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberManualModeListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberManualModeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberManualModeListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberPortInd206(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberPortInd206 {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberPortInd206
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZoneIpProtoProtoNumberPortIndSamplingEnable207(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberPortIndSamplingEnable207(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberPortIndSamplingEnable207 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberPortIndSamplingEnable207, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberPortIndSamplingEnable207
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberProgressionTracking208(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberProgressionTracking208 {

	var ret edpt.DdosDstZoneIpProtoProtoNumberProgressionTracking208
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberTopkDestinations209(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberTopkDestinations209 {

	var ret edpt.DdosDstZoneIpProtoProtoNumberTopkDestinations209
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberTopkSources210(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberTopkSources210 {

	var ret edpt.DdosDstZoneIpProtoProtoNumberTopkSources210
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumber(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumber {
	var ret edpt.DdosDstZoneIpProtoProtoNumber
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.ApplyPolicyOnOverflow = d.Get("apply_policy_on_overflow").(int)
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DropFragPkt = d.Get("drop_frag_pkt").(int)
	ret.Inst.DynamicEntryOverflowPolicyList = getSliceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyList(d.Get("dynamic_entry_overflow_policy_list").([]interface{}))
	ret.Inst.EnableClassListOverflow = d.Get("enable_class_list_overflow").(int)
	ret.Inst.EnableTopK = d.Get("enable_top_k").(int)
	ret.Inst.EnableTopKDestination = d.Get("enable_top_k_destination").(int)
	ret.Inst.EspInspect = getObjectDdosDstZoneIpProtoProtoNumberEspInspect(d.Get("esp_inspect").([]interface{}))
	ret.Inst.FasterDeEscalation = d.Get("faster_de_escalation").(int)
	ret.Inst.GlidCfg = getObjectDdosDstZoneIpProtoProtoNumberGlidCfg(d.Get("glid_cfg").([]interface{}))
	ret.Inst.IpFilteringPolicy = d.Get("ip_filtering_policy").(string)
	ret.Inst.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoNumberIpFilteringPolicyOper205(d.Get("ip_filtering_policy_oper").([]interface{}))
	ret.Inst.LevelList = getSliceDdosDstZoneIpProtoProtoNumberLevelList(d.Get("level_list").([]interface{}))
	ret.Inst.ManualModeEnable = d.Get("manual_mode_enable").(int)
	ret.Inst.ManualModeList = getSliceDdosDstZoneIpProtoProtoNumberManualModeList(d.Get("manual_mode_list").([]interface{}))
	ret.Inst.MaxDynamicEntryCount = d.Get("max_dynamic_entry_count").(int)
	ret.Inst.PortInd = getObjectDdosDstZoneIpProtoProtoNumberPortInd206(d.Get("port_ind").([]interface{}))
	ret.Inst.ProgressionTracking = getObjectDdosDstZoneIpProtoProtoNumberProgressionTracking208(d.Get("progression_tracking").([]interface{}))
	ret.Inst.ProtocolNum = d.Get("protocol_num").(int)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.SrcBasedPolicyList = getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyList(d.Get("src_based_policy_list").([]interface{}))
	ret.Inst.TopkDestinations = getObjectDdosDstZoneIpProtoProtoNumberTopkDestinations209(d.Get("topk_destinations").([]interface{}))
	ret.Inst.TopkDstNumRecords = d.Get("topk_dst_num_records").(int)
	ret.Inst.TopkNumRecords = d.Get("topk_num_records").(int)
	ret.Inst.TopkSources = getObjectDdosDstZoneIpProtoProtoNumberTopkSources210(d.Get("topk_sources").([]interface{}))
	ret.Inst.UnlimitedDynamicEntryCount = d.Get("unlimited_dynamic_entry_count").(int)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
