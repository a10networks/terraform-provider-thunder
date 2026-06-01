package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneService() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service`: DDOS Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceDelete,

		Schema: map[string]*schema.Schema{
			"age": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
			},
			"apply_policy_on_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
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
			"default_action_list": {
				Type: schema.TypeString, Optional: true, Description: "Configure default-action-list in zone-service",
			},
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"dynamic_entry_count_warn_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Set threshold percentage of \"max-src-dst-entry\" for generating warning logs. Including start and end.",
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
			"enable_class_list_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for zone-port or class-list",
			},
			"enable_top_k": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k source IP detection",
			},
			"enable_top_k_destination": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k destination IP detection",
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
			"ip_filtering_policy": {
				Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
			},
			"ip_filtering_policy_statistics": {
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
						"close_sessions_for_all_sources": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for all sources",
						},
						"clear_sources_upon_deescalation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Clear sources upon de-escalation from level 1 to 0 or manual to 0",
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
										Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'concurrent-conns': number of concurrent connections; 'conn-miss-rate': rate of incoming packets for which no previously established connection exists; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'syn-ack-rate': rate on incoming SYN-ACK packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization; 'learnt-sources': learnt sources;",
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
			"log_src_default_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable src default logging in zone-service",
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
						"close_sessions_for_all_sources": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for all sources",
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
			"outbound_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only allow outbound traffic",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_rate_adaptive_threshold': Pkt Rate Adaptive Threshold; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_pkt_drop_rate_adaptive_threshold': Pkt Drop Rate Adaptive Threshold; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_syn_rate_adaptive_threshold': TCP SYN Rate Adaptive Threshold; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_fin_rate_adaptive_threshold': TCP FIN Rate Adaptive Threshold; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_rst_rate_adaptive_threshold': TCP RST Rate Adaptive Threshold; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_small_window_ack_rate_adaptive_threshold': TCP Small Window ACK Rate Adaptive Threshold; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_empty_ack_rate_adaptive_threshold': TCP Empty ACK Rate Adaptive Threshold; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_small_payload_rate_adaptive_threshold': TCP Small Payload Rate Adaptive Threshold; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_pkt_drop_ratio_adaptive_threshold': Pkt Drop / Pkt Rcvd Adaptive Threshold; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_inb_per_outb_adaptive_threshold': Bytes-to / Bytes-from Adaptive Threshold; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_syn_per_fin_rate_adaptive_threshold': TCP SYN Rate / FIN Rate Adaptive Threshold; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_conn_miss_rate_adaptive_threshold': TCP Session Miss Rate Adaptive Threshold; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_concurrent_conns_adaptive_threshold': TCP/UDP Concurrent Sessions Adaptive Threshold; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_data_cpu_util_adaptive_threshold': Data CPU Utilization Adaptive Threshold; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_outside_intf_util_adaptive_threshold': Outside Interface Utilization Adaptive Threshold; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_frag_rate_adaptive_threshold': Frag Pkt Rate Adaptive Threshold; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max; 'ddet_ind_bit_rate_adaptive_threshold': Bit Rate Adaptive Threshold; 'ddet_ind_total_szp_current': Total Learnt Sources Current; 'ddet_ind_total_szp_min': Total Learnt Sources Min; 'ddet_ind_total_szp_max': Total Learnt Sources Max; 'ddet_ind_total_szp_adaptive_threshold': Total Learnt Sources Adaptive Threshold; 'ddet_ind_syn_ack_rate_current': TCP SYN ACK Rate Current; 'ddet_ind_syn_ack_rate_min': TCP SYN ACK Rate Min; 'ddet_ind_syn_ack_rate_max': TCP SYN ACK Rate Max; 'ddet_ind_syn_ack_rate_adaptive_threshold': TCP SYN ACK Adaptive Threshold; 'ddet_ind_inside_out_concurrent_conns_current': TCP/UDP Inside Out Concurrent Sessions Current; 'ddet_ind_inside_out_concurrent_conns_min': TCP/UDP Inside Out Concurrent Sessions Min; 'ddet_ind_inside_out_concurrent_conns_max': TCP/UDP Inside Out Concurrent Sessions Max; 'ddet_ind_inside_out_concurrent_conns_adaptive_threshold': TCP/UDP Inside Out Concurrent Sessions Adaptive Threshold; 'ddet_ind_pkt_rate_adaptive_baseline': Pkt Rate Adaptive Baseline; 'ddet_ind_bit_rate_adaptive_baseline': Bit Rate Adaptive Baseline;",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port; 'quic': QUIC Port;",
			},
			"same_source_dest_port_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop packet with same Source Port and Dest Port",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"sflow_common": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all sFlow polling options under this zone port",
			},
			"sflow_http": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow HTTP counter polling",
			},
			"sflow_ip_filtering_policy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow IP filtering policy per port per rule counter polling",
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
						"sflow_samp_sni": {
							Type: schema.TypeInt, Optional: true, Description: "Enable sFlow Auto-Discoverd SNI Sampling",
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
									"class_list_glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID (class-list based)",
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
									"dynamic_entry_count_warn_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "Set threshold percentage of \"max-src-dst-entry\" for generating warning logs. Including start and end.",
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
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow; 'exceed_drop_pkt_rate_clist': Packet Rate Exceeded; 'exceed_drop_conn_rate_clist': Conn Rate Exceeded; 'exceed_drop_conn_limit_clist': Conn Limit Exceeded; 'exceed_drop_kbit_rate_clist': KiBit Rate Exceeded; 'exceed_drop_kbit_rate_clist_pkt': KiBit Rate Exceeded Count; 'exceed_drop_frag_rate_clist': Frag Rate Exceeded;",
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
			"stateful": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateful tracking of sessions (Default is stateless)",
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
			"topk_dst_sort_key": {
				Type: schema.TypeString, Optional: true, Default: "avg", Description: "'avg': window average; 'max-peak': max peak;",
			},
			"topk_num_records": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
			},
			"topk_sort_key": {
				Type: schema.TypeString, Optional: true, Default: "avg", Description: "'avg': window average; 'max-peak': max peak;",
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
			"virtualhosts": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vhosts_config": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': configure virtualhost based mitigation for ssl services;",
						},
						"source_tracking_all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable creation of source entries for all virtualhosts",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"virtualhost_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vhost": {
										Type: schema.TypeString, Required: true, Description: "name for virtualhost",
									},
									"servername": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"match_type": {
													Type: schema.TypeString, Optional: true, Description: "'contains': match servername extension when contains this string; 'ends-with': match servername extension when ends with this string; 'equals': match servername extension when equals this string; 'starts-with': match servername extension when starts with this string;",
												},
												"host_match_string": {
													Type: schema.TypeString, Optional: true, Description: "SNI String",
												},
											},
										},
									},
									"servername_list": {
										Type: schema.TypeString, Optional: true, Description: "Class List to match servername (AC type Class List Name)",
									},
									"servername_match_any": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match any SNI extension",
									},
									"servername_no_sni": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match when there is no SNI extension found",
									},
									"source_tracking": {
										Type: schema.TypeString, Optional: true, Default: "follow", Description: "'follow': enable creation of source entries when source-tracking-all is enabled (default); 'enable': enable creation of source entries on this virtualhost; 'disable': disable creation of source entries on this virtualhost;",
									},
									"glid_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"glid_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed; 'ignore': Do nothing for glid exceed;",
												},
											},
										},
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"level_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_num": {
													Type: schema.TypeString, Required: true, Description: "'0': Default policy level;",
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
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
															},
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneService(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneService(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneService(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneService(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZonePortZoneServiceCaptureConfig(d []interface{}) edpt.DdosDstZonePortZoneServiceCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyListZoneTemplate
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

func getObjectDdosDstZonePortZoneServiceGlidCfg(d []interface{}) edpt.DdosDstZonePortZoneServiceGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
		ret.ActionList = in["action_list"].(string)
		ret.PerAddrGlid = in["per_addr_glid"].(string)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceIpFilteringPolicyStatistics254(d []interface{}) edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatistics254 {

	var ret edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatistics254
	return ret
}

func getSliceDdosDstZonePortZoneServiceLevelList(d []interface{}) []edpt.DdosDstZonePortZoneServiceLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneEscalationScore = in["zone_escalation_score"].(int)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		oi.SrcEscalationScore = in["src_escalation_score"].(int)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceLevelListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		oi.CloseSessionsForAllSources = in["close_sessions_for_all_sources"].(int)
		oi.ClearSourcesUponDeescalation = in["clear_sources_upon_deescalation"].(int)
		oi.StartSignatureExtraction = in["start_signature_extraction"].(int)
		oi.StartPatternRecognition = in["start_pattern_recognition"].(int)
		oi.ApplyExtractedFilters = in["apply_extracted_filters"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosDstZonePortZoneServiceLevelListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceLevelListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceLevelListZoneTemplate
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

func getSliceDdosDstZonePortZoneServiceLevelListIndicatorList(d []interface{}) []edpt.DdosDstZonePortZoneServiceLevelListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceLevelListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceLevelListIndicatorList
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

func getSliceDdosDstZonePortZoneServiceManualModeList(d []interface{}) []edpt.DdosDstZonePortZoneServiceManualModeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceManualModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceManualModeList
		oi.Config = in["config"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceManualModeListZoneTemplate(in["zone_template"].([]interface{}))
		oi.CloseSessionsForUnauthSources = in["close_sessions_for_unauth_sources"].(int)
		oi.CloseSessionsForAllSources = in["close_sessions_for_all_sources"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceManualModeListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceManualModeListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceManualModeListZoneTemplate
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

func getObjectDdosDstZonePortZoneServicePatternRecognition255(d []interface{}) edpt.DdosDstZonePortZoneServicePatternRecognition255 {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServicePatternRecognition255
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

func getObjectDdosDstZonePortZoneServicePatternRecognitionPuDetails256(d []interface{}) edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetails256 {

	var ret edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetails256
	return ret
}

func getObjectDdosDstZonePortZoneServicePortInd257(d []interface{}) edpt.DdosDstZonePortZoneServicePortInd257 {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServicePortInd257
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstZonePortZoneServicePortIndSamplingEnable258(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServicePortIndSamplingEnable258(d []interface{}) []edpt.DdosDstZonePortZoneServicePortIndSamplingEnable258 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServicePortIndSamplingEnable258, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServicePortIndSamplingEnable258
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceProgressionTracking259(d []interface{}) edpt.DdosDstZonePortZoneServiceProgressionTracking259 {

	var ret edpt.DdosDstZonePortZoneServiceProgressionTracking259
	return ret
}

func getObjectDdosDstZonePortZoneServiceSflowTcp(d []interface{}) edpt.DdosDstZonePortZoneServiceSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
		ret.SflowSampSni = in["sflow_samp_sni"].(int)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.ClassListGlid = in["class_list_glid"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.DynamicEntryCountWarnThreshold = in["dynamic_entry_count_warn_threshold"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListZoneTemplate
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

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate
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

func getObjectDdosDstZonePortZoneServiceTopkDestinations260(d []interface{}) edpt.DdosDstZonePortZoneServiceTopkDestinations260 {

	var ret edpt.DdosDstZonePortZoneServiceTopkDestinations260
	return ret
}

func getObjectDdosDstZonePortZoneServiceTopkSources261(d []interface{}) edpt.DdosDstZonePortZoneServiceTopkSources261 {

	var ret edpt.DdosDstZonePortZoneServiceTopkSources261
	return ret
}

func getObjectDdosDstZonePortZoneServiceVirtualhosts262(d []interface{}) edpt.DdosDstZonePortZoneServiceVirtualhosts262 {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceVirtualhosts262
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VhostsConfig = in["vhosts_config"].(string)
		ret.SourceTrackingAll = in["source_tracking_all"].(int)
		//omit uuid
		ret.VirtualhostList = getSliceDdosDstZonePortZoneServiceVirtualhostsVirtualhostList263(in["virtualhost_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceVirtualhostsVirtualhostList263(d []interface{}) []edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostList263 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostList263, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostList263
		oi.Vhost = in["vhost"].(string)
		oi.Servername = getSliceDdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername264(in["servername"].([]interface{}))
		oi.ServernameList = in["servername_list"].(string)
		oi.ServernameMatchAny = in["servername_match_any"].(int)
		oi.ServernameNoSni = in["servername_no_sni"].(int)
		oi.SourceTracking = in["source_tracking"].(string)
		oi.GlidCfg = getObjectDdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg265(in["glid_cfg"].([]interface{}))
		oi.Deny = in["deny"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.LevelList = getSliceDdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList266(in["level_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername264(d []interface{}) []edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername264 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername264, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername264
		oi.MatchType = in["match_type"].(string)
		oi.HostMatchString = in["host_match_string"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg265(d []interface{}) edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg265 {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg265
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList266(d []interface{}) []edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList266 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList266, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList266
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate267(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate267(d []interface{}) edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate267 {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate267
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Tcp = in["tcp"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneService(d *schema.ResourceData) edpt.DdosDstZonePortZoneService {
	var ret edpt.DdosDstZonePortZoneService
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.ApplyPolicyOnOverflow = d.Get("apply_policy_on_overflow").(int)
	ret.Inst.CaptureConfig = getObjectDdosDstZonePortZoneServiceCaptureConfig(d.Get("capture_config").([]interface{}))
	ret.Inst.DefaultActionList = d.Get("default_action_list").(string)
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DynamicEntryCountWarnThreshold = d.Get("dynamic_entry_count_warn_threshold").(int)
	ret.Inst.DynamicEntryOverflowPolicyList = getSliceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyList(d.Get("dynamic_entry_overflow_policy_list").([]interface{}))
	ret.Inst.EnableClassListOverflow = d.Get("enable_class_list_overflow").(int)
	ret.Inst.EnableTopK = d.Get("enable_top_k").(int)
	ret.Inst.EnableTopKDestination = d.Get("enable_top_k_destination").(int)
	ret.Inst.FasterDeEscalation = d.Get("faster_de_escalation").(int)
	ret.Inst.GlidCfg = getObjectDdosDstZonePortZoneServiceGlidCfg(d.Get("glid_cfg").([]interface{}))
	ret.Inst.IpFilteringPolicy = d.Get("ip_filtering_policy").(string)
	ret.Inst.IpFilteringPolicyStatistics = getObjectDdosDstZonePortZoneServiceIpFilteringPolicyStatistics254(d.Get("ip_filtering_policy_statistics").([]interface{}))
	ret.Inst.LevelList = getSliceDdosDstZonePortZoneServiceLevelList(d.Get("level_list").([]interface{}))
	ret.Inst.LogSrcDefaultEnable = d.Get("log_src_default_enable").(int)
	ret.Inst.ManualModeEnable = d.Get("manual_mode_enable").(int)
	ret.Inst.ManualModeList = getSliceDdosDstZonePortZoneServiceManualModeList(d.Get("manual_mode_list").([]interface{}))
	ret.Inst.MaxDynamicEntryCount = d.Get("max_dynamic_entry_count").(int)
	ret.Inst.OutboundOnly = d.Get("outbound_only").(int)
	ret.Inst.PatternRecognition = getObjectDdosDstZonePortZoneServicePatternRecognition255(d.Get("pattern_recognition").([]interface{}))
	ret.Inst.PatternRecognitionPuDetails = getObjectDdosDstZonePortZoneServicePatternRecognitionPuDetails256(d.Get("pattern_recognition_pu_details").([]interface{}))
	ret.Inst.PortInd = getObjectDdosDstZonePortZoneServicePortInd257(d.Get("port_ind").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.ProgressionTracking = getObjectDdosDstZonePortZoneServiceProgressionTracking259(d.Get("progression_tracking").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SameSourceDestPortDrop = d.Get("same_source_dest_port_drop").(int)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.SflowCommon = d.Get("sflow_common").(int)
	ret.Inst.SflowHttp = d.Get("sflow_http").(int)
	ret.Inst.SflowIpFilteringPolicy = d.Get("sflow_ip_filtering_policy").(int)
	ret.Inst.SflowPackets = d.Get("sflow_packets").(int)
	ret.Inst.SflowTcp = getObjectDdosDstZonePortZoneServiceSflowTcp(d.Get("sflow_tcp").([]interface{}))
	ret.Inst.SrcBasedPolicyList = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyList(d.Get("src_based_policy_list").([]interface{}))
	ret.Inst.Stateful = d.Get("stateful").(int)
	ret.Inst.TopkDestinations = getObjectDdosDstZonePortZoneServiceTopkDestinations260(d.Get("topk_destinations").([]interface{}))
	ret.Inst.TopkDstNumRecords = d.Get("topk_dst_num_records").(int)
	ret.Inst.TopkDstSortKey = d.Get("topk_dst_sort_key").(string)
	ret.Inst.TopkNumRecords = d.Get("topk_num_records").(int)
	ret.Inst.TopkSortKey = d.Get("topk_sort_key").(string)
	ret.Inst.TopkSources = getObjectDdosDstZonePortZoneServiceTopkSources261(d.Get("topk_sources").([]interface{}))
	ret.Inst.UnlimitedDynamicEntryCount = d.Get("unlimited_dynamic_entry_count").(int)
	//omit uuid
	ret.Inst.Virtualhosts = getObjectDdosDstZonePortZoneServiceVirtualhosts262(d.Get("virtualhosts").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
