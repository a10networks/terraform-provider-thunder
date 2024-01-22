package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry`: Configure IP/IPv6 static entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryCreate,
		UpdateContext: resourceDdosDstEntryUpdate,
		ReadContext:   resourceDdosDstEntryRead,
		DeleteContext: resourceDdosDstEntryDelete,

		Schema: map[string]*schema.Schema{
			"advertised_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP advertised",
			},
			"blackhole_on_glid_exceed": {
				Type: schema.TypeInt, Optional: true, Description: "Blackhole destination entry for X minutes upon glid limit exceeded",
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
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description for this Destination Entry",
			},
			"dest_nat_ip": {
				Type: schema.TypeString, Optional: true, Description: "Destination NAT IP address",
			},
			"dest_nat_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Destination NAT IPv6 address",
			},
			"drop_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable certain drops during packet processing",
			},
			"drop_disable_fwd_immediate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Immediately forward L4 drops",
			},
			"drop_frag_pkt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
			},
			"drop_on_no_src_dst_default": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop if no match with src-based-policy class-list, and default is not configured",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"dynamic_entry_overflow_policy_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dummy_name": {
							Type: schema.TypeString, Required: true, Description: "'configuration': Configure src dst dynamic entry count overflow policy;",
						},
						"bypass": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
						},
						"exceed_log_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
									},
								},
							},
						},
						"log_periodic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logging": {
										Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
									},
								},
							},
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"l4_type_src_dst_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
												},
												"other": {
													Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
												},
												"template_icmp_v4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
												},
												"template_icmp_v6": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
						"app_type_src_dst_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
			"exceed_log_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
						},
						"log_with_sflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Turn on sflow sample with log",
						},
						"log_high_frequency": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable High frequency logging for non-event logs per entry",
						},
						"rate_limit": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate limit per second per entry(Default : 1 per second)",
						},
					},
				},
			},
			"exceed_log_dep_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exceed_log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "(Deprecated)Enable logging of limit exceed drop's",
						},
						"log_with_sflow_dep": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Turn on sflow sample with log",
						},
					},
				},
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"glid_exceed_action": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stateless_encap_action_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stateless_encap_action": {
										Type: schema.TypeString, Optional: true, Description: "'stateless-tunnel-encap': Encapsulate all packets; 'stateless-tunnel-encap-scrubbed': Encapsulate all packets and allow packets to go through other DDoS checks before sent (conn-limit exceeded packet can not be scrubbed, it will default to stateless-tunnel-encap);",
									},
									"encap_template": {
										Type: schema.TypeString, Optional: true, Description: "Apply legacy encap template for encap action",
									},
								},
							},
						},
					},
				},
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
			"ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"ip_proto_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Protocol Number",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
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
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_exceed_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stateless_encap_action_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stateless_encap_action": {
													Type: schema.TypeString, Optional: true, Description: "'stateless-tunnel-encap': Encapsulate all packets; 'stateless-tunnel-encap-scrubbed': Encapsulate all packets and allow packets to go through other DDoS checks before sent (conn-limit exceeded packet can not be scrubbed, it will default to stateless-tunnel-encap);",
												},
												"encap_template": {
													Type: schema.TypeString, Optional: true, Description: "Apply legacy encap template for encap action",
												},
											},
										},
									},
								},
							},
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"other": {
										Type: schema.TypeString, Optional: true, Description: "DDOS other template",
									},
								},
							},
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
					},
				},
			},
			"ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"l4_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': L4-Type TCP; 'udp': L4-Type UDP; 'icmp': L4-Type ICMP; 'other': L4-Type OTHER;",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_exceed_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stateless_encap_action_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stateless_encap_action": {
													Type: schema.TypeString, Optional: true, Description: "'stateless-tunnel-encap': Encapsulate all packets; 'stateless-tunnel-encap-scrubbed': Encapsulate all packets and allow packets to go through other DDoS checks before sent (conn-limit exceeded packet can not be scrubbed, it will default to stateless-tunnel-encap);",
												},
												"encap_template": {
													Type: schema.TypeString, Optional: true, Description: "Apply legacy encap template for encap action",
												},
											},
										},
									},
								},
							},
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"max_rexmit_syn_per_flow": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of re-transmit SYN per flow",
						},
						"max_rexmit_syn_per_flow_exceed_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the packet; 'black-list': Add the source IP into black list;",
						},
						"disable_syn_auth": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable TCP SYN Authentication",
						},
						"syn_auth": {
							Type: schema.TypeString, Optional: true, Default: "send-rst", Description: "'send-rst': Send RST to client upon client ACK; 'force-rst-by-ack': Force client RST via the use of ACK; 'force-rst-by-synack': Force client RST via the use of bad SYN|ACK; 'disable': Disable TCP SYN Authentication;",
						},
						"syn_cookie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SYN Cookie",
						},
						"tcp_reset_client": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset to client when rate exceeds or session ages out",
						},
						"tcp_reset_server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset to server when rate exceeds or session ages out",
						},
						"drop_on_no_port_match": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'disable': disable; 'enable': enable;",
						},
						"stateful": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateful tracking of sessions (Default is stateless)",
						},
						"tunnel_decap": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_decap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IP Tunnel decapsulation",
									},
									"gre_decap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GRE Tunnel decapsulation",
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
								},
							},
						},
						"tunnel_rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_rate_limit": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable inner IP rate limiting on IPinIP traffic",
									},
									"gre_rate_limit": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable inner IP rate limiting on GRE traffic",
									},
								},
							},
						},
						"drop_frag_pkt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
						},
						"undefined_port_hit_statistics": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"undefined_port_hit_statistics": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable port scanning statistics",
									},
									"reset_interval": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Configure port scanning counter reset interval (minutes), Default 60 mins",
									},
								},
							},
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_icmp_v4": {
										Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
									},
									"template_icmp_v6": {
										Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
									},
								},
							},
						},
						"detection_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos detection",
						},
						"enable_top_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k entries",
						},
						"topk_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
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
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max;",
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
					},
				},
			},
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
			},
			"operational_mode": {
				Type: schema.TypeString, Optional: true, Default: "protection", Description: "'protection': Protection mode; 'bypass': Bypass mode;",
			},
			"outbound_forward_dscp": {
				Type: schema.TypeInt, Optional: true, Description: "To set dscp value for outbound",
			},
			"pattern_recognition_hw_filter_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "to enable pattern recognition hardware filter",
			},
			"pattern_recognition_sensitivity": {
				Type: schema.TypeString, Optional: true, Description: "'high': High sensitive pattern recognition; 'medium': Medium sensitive pattern recognition; 'low': Low sensitive pattern recognition;",
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
						"detection_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos detection",
						},
						"enable_top_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k entries",
						},
						"topk_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_exceed_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stateless_encap_action_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stateless_encap_action": {
													Type: schema.TypeString, Optional: true, Description: "'stateless-tunnel-encap': Encapsulate all packets; 'stateless-tunnel-encap-scrubbed': Encapsulate all packets and allow packets to go through other DDoS checks before sent (conn-limit exceeded packet can not be scrubbed, it will default to stateless-tunnel-encap);",
												},
												"encap_template": {
													Type: schema.TypeString, Optional: true, Description: "Apply legacy encap template for encap action",
												},
											},
										},
									},
								},
							},
						},
						"dns_cache": {
							Type: schema.TypeString, Optional: true, Description: "DNS Cache Instance",
						},
						"template": {
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
										Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
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
								},
							},
						},
						"sflow": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"polling": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
											},
										},
									},
								},
							},
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
						"ip_filtering_policy": {
							Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
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
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max;",
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
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"detection_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos detection",
						},
						"enable_top_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k entries",
						},
						"topk_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_exceed_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stateless_encap_action_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stateless_encap_action": {
													Type: schema.TypeString, Optional: true, Description: "'stateless-tunnel-encap': Encapsulate all packets; 'stateless-tunnel-encap-scrubbed': Encapsulate all packets and allow packets to go through other DDoS checks before sent (conn-limit exceeded packet can not be scrubbed, it will default to stateless-tunnel-encap);",
												},
												"encap_template": {
													Type: schema.TypeString, Optional: true, Description: "Apply legacy encap template for encap action",
												},
											},
										},
									},
								},
							},
						},
						"template": {
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
										Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
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
								},
							},
						},
						"sflow": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"polling": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
											},
										},
									},
								},
							},
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
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max;",
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
					},
				},
			},
			"reporting_disabled": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Reporting",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'dst_tcp_any_exceed': TCP Dst L4-Type Rate: Total Exceeded; 'dst_tcp_pkt_rate_exceed': TCP Dst L4-Type Rate: Packet Exceeded; 'dst_tcp_conn_rate_exceed': TCP Dst L4-Type Rate: Conn Exceeded; 'dst_udp_any_exceed': UDP Dst L4-Type Rate: Total Exceeded; 'dst_udp_pkt_rate_exceed': UDP Dst L4-Type Rate: Packet Exceeded; 'dst_udp_conn_limit_exceed': UDP Dst L4-Type Limit: Conn Exceeded; 'dst_udp_conn_rate_exceed': UDP Dst L4-Type Rate: Conn Exceeded; 'dst_icmp_pkt_rate_exceed': ICMP Dst Rate: Packet Exceeded; 'dst_other_pkt_rate_exceed': OTHER Dst L4-Type Rate: Packet Exceeded; 'dst_other_frag_pkt_rate_exceed': OTHER Dst L4-Type Rate: Frag Exceeded; 'dst_port_pkt_rate_exceed': Port Rate: Packet Exceeded; 'dst_port_conn_limit_exceed': Port Limit: Conn Exceeded; 'dst_port_conn_rate_exceed': Port Rate: Conn Exceeded; 'dst_pkt_sent': Inbound: Packets Forwarded; 'dst_udp_pkt_sent': UDP Total Packets Forwarded; 'dst_tcp_pkt_sent': TCP Total Packets Forwarded; 'dst_icmp_pkt_sent': ICMP Total Packets Forwarded; 'dst_other_pkt_sent': OTHER Total Packets Forwarded; 'dst_tcp_conn_limit_exceed': TCP Dst L4-Type Limit: Conn Exceeded; 'dst_tcp_pkt_rcvd': TCP Total Packets Received; 'dst_udp_pkt_rcvd': UDP Total Packets Received; 'dst_icmp_pkt_rcvd': ICMP Total Packets Received; 'dst_other_pkt_rcvd': OTHER Total Packets Received; 'dst_udp_filter_match': UDP Filter Match; 'dst_udp_filter_not_match': UDP Filter Not Matched on Pkt; 'dst_udp_filter_action_blacklist': UDP Filter Action Blacklist; 'dst_udp_filter_action_drop': UDP Filter Action Drop; 'dst_tcp_syn': TCP Total SYN Received; 'dst_tcp_syn_drop': TCP SYN Packets Dropped; 'dst_tcp_src_rate_drop': TCP Src Rate: Total Exceeded; 'dst_udp_src_rate_drop': UDP Src Rate: Total Exceeded; 'dst_icmp_src_rate_drop': ICMP Src Rate: Total Exceeded; 'dst_other_frag_src_rate_drop': OTHER Src Rate: Frag Exceeded; 'dst_other_src_rate_drop': OTHER Src Rate: Total Exceeded; 'dst_tcp_drop': TCP Total Packets Dropped; 'dst_udp_drop': UDP Total Packets Dropped; 'dst_icmp_drop': ICMP Total Packets Dropped; 'dst_frag_drop': Fragmented Packets Dropped; 'dst_other_drop': OTHER Total Packets Dropped; 'dst_tcp_auth': TCP Auth: SYN Cookie Sent; 'dst_udp_filter_action_default_pass': UDP Filter Action Default Pass; 'dst_tcp_filter_match': TCP Filter Match; 'dst_tcp_filter_not_match': TCP Filter Not Matched on Pkt; 'dst_tcp_filter_action_blacklist': TCP Filter Action Blacklist; 'dst_tcp_filter_action_drop': TCP Filter Action Drop; 'dst_tcp_filter_action_default_pass': TCP Filter Action Default Pass; 'dst_udp_filter_action_whitelist': UDP Filter Action WL; 'dst_over_limit_on': DST overlimit Trigger ON; 'dst_over_limit_off': DST overlimit Trigger OFF; 'dst_port_over_limit_on': DST port overlimit Trigger ON; 'dst_port_over_limit_off': DST port overlimit Trigger OFF; 'dst_over_limit_action': DST overlimit action; 'dst_port_over_limit_action': DST port overlimit action; 'scanning_detected_drop': Scanning Detected drop (deprecated); 'scanning_detected_blacklist': Scanning Detected blacklist (deprecated); 'dst_udp_kibit_rate_drop': UDP Dst L4-Type Rate: KiBit Exceeded; 'dst_tcp_kibit_rate_drop': TCP Dst L4-Type Rate: KiBit Exceeded; 'dst_icmp_kibit_rate_drop': ICMP Dst Rate: KiBit Exceeded; 'dst_other_kibit_rate_drop': OTHER Dst L4-Type Rate: KiBit Exceeded; 'dst_port_undef_drop': Dst Port Undefined Dropped; 'dst_port_bl': Dst Port Blacklist Packets Dropped; 'dst_src_port_bl': Dst SrcPort Blacklist Packets Dropped; 'dst_port_kbit_rate_exceed': Port Rate: KiBit Exceeded; 'dst_tcp_src_drop': TCP Src Packets Dropped; 'dst_udp_src_drop': UDP Src Packets Dropped; 'dst_icmp_src_drop': ICMP Src Packets Dropped; 'dst_other_src_drop': OTHER Src Packets Dropped; 'tcp_syn_rcvd': TCP Inbound SYN Received; 'tcp_syn_ack_rcvd': TCP SYN ACK Received; 'tcp_ack_rcvd': TCP ACK Received; 'tcp_fin_rcvd': TCP FIN Received; 'tcp_rst_rcvd': TCP RST Received; 'ingress_bytes': Inbound: Bytes Received; 'egress_bytes': Outbound: Bytes Received; 'ingress_packets': Inbound: Packets Received; 'egress_packets': Outbound: Packets Received; 'tcp_fwd_recv': TCP Inbound Packets Received; 'udp_fwd_recv': UDP Inbound Packets Received; 'icmp_fwd_recv': ICMP Inbound Packets Received; 'tcp_syn_cookie_fail': TCP Auth: SYN Cookie Failed; 'dst_tcp_session_created': TCP Sessions Created; 'dst_udp_session_created': UDP Sessions Created; 'dst_tcp_filter_action_whitelist': TCP Filter Action WL; 'dst_other_filter_match': OTHER Filter Match; 'dst_other_filter_not_match': OTHER Filter Not Matched on Pkt; 'dst_other_filter_action_blacklist': OTHER Filter Action Blacklist; 'dst_other_filter_action_drop': OTHER Filter Action Drop; 'dst_other_filter_action_whitelist': OTHER Filter Action WL; 'dst_other_filter_action_default_pass': OTHER Filter Action Default Pass; 'dst_blackhole_inject': Dst Blackhole Inject; 'dst_blackhole_withdraw': Dst Blackhole Withdraw; 'dst_tcp_out_of_seq_excd': TCP Out-Of-Seq Exceeded; 'dst_tcp_retransmit_excd': TCP Retransmit Exceeded; 'dst_tcp_zero_window_excd': TCP Zero-Window Exceeded; 'dst_tcp_conn_prate_excd': TCP Rate: Conn Pkt Exceeded; 'dst_tcp_action_on_ack_init': TCP Auth: ACK Retry Init; 'dst_tcp_action_on_ack_gap_drop': TCP Auth: ACK Retry Retry-Gap Dropped; 'dst_tcp_action_on_ack_fail': TCP Auth: ACK Retry Dropped; 'dst_tcp_action_on_ack_pass': TCP Auth: ACK Retry Passed; 'dst_tcp_action_on_syn_init': TCP Auth: SYN Retry Init; 'dst_tcp_action_on_syn_gap_drop': TCP Auth: SYN Retry-Gap Dropped; 'dst_tcp_action_on_syn_fail': TCP Auth: SYN Retry Dropped; 'dst_tcp_action_on_syn_pass': TCP Auth: SYN Retry Passed; 'udp_payload_too_small': UDP Payload Too Small; 'udp_payload_too_big': UDP Payload Too Large; 'dst_udp_conn_prate_excd': UDP Rate: Conn Pkt Exceeded; 'dst_udp_ntp_monlist_req': UDP NTP Monlist Request; 'dst_udp_ntp_monlist_resp': UDP NTP Monlist Response; 'dst_udp_wellknown_sport_drop': UDP SrcPort Wellknown; 'dst_udp_retry_init': UDP Auth: Retry Init; 'dst_udp_retry_pass': UDP Auth: Retry Passed; 'dst_tcp_bytes_drop': TCP Total Bytes Dropped; 'dst_udp_bytes_drop': UDP Total Bytes Dropped; 'dst_icmp_bytes_drop': ICMP Total Bytes Dropped; 'dst_other_bytes_drop': OTHER Total Bytes Dropped; 'dst_out_no_route': Dst IPv4/v6 Out No Route; 'outbound_bytes_sent': Outbound: Bytes Forwarded; 'outbound_pkt_drop': Outbound: Packets Dropped; 'outbound_bytes_drop': Outbound: Bytes Dropped; 'outbound_pkt_sent': Outbound: Packets Forwarded; 'inbound_bytes_sent': Inbound: Bytes Forwarded; 'inbound_bytes_drop': Inbound: Bytes Dropped; 'dst_src_port_pkt_rate_exceed': SrcPort Rate: Packet Exceeded; 'dst_src_port_kbit_rate_exceed': SrcPort Rate: KiBit Exceeded; 'dst_src_port_conn_limit_exceed': SrcPort Limit: Conn Exceeded; 'dst_src_port_conn_rate_exceed': SrcPort Rate: Conn Exceeded; 'dst_ip_proto_pkt_rate_exceed': IP-Proto Rate: Packet Exceeded; 'dst_ip_proto_kbit_rate_exceed': IP-Proto Rate: KiBit Exceeded; 'dst_tcp_port_any_exceed': TCP Port Rate: Total Exceed; 'dst_udp_port_any_exceed': UDP Port Rate: Total Exceed; 'dst_tcp_auth_pass': TCP Auth: SYN Auth Passed; 'dst_tcp_rst_cookie_fail': TCP Auth: RST Cookie Failed; 'dst_tcp_unauth_drop': TCP Auth: Unauth Dropped; 'src_tcp_syn_auth_fail': Src TCP Auth: SYN Auth Failed; 'src_tcp_syn_cookie_sent': Src TCP Auth: SYN Cookie Sent; 'src_tcp_syn_cookie_fail': Src TCP Auth: SYN Cookie Failed; 'src_tcp_rst_cookie_fail': Src TCP Auth: RST Cookie Failed; 'src_tcp_unauth_drop': Src TCP Auth: Unauth Dropped; 'src_tcp_action_on_syn_init': Src TCP Auth: SYN Retry Init;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'src_tcp_action_on_syn_gap_drop': Src TCP Auth: SYN Retry-Gap Dropped; 'src_tcp_action_on_syn_fail': Src TCP Auth: SYN Retry Dropped; 'src_tcp_action_on_ack_init': Src TCP Auth: ACK Retry Init; 'src_tcp_action_on_ack_gap_drop': Src TCP Auth: ACK Retry Retry-Gap Dropped; 'src_tcp_action_on_ack_fail': Src TCP Auth: ACK Retry Dropped; 'src_tcp_out_of_seq_excd': Src TCP Out-Of-Seq Exceeded; 'src_tcp_retransmit_excd': Src TCP Retransmit Exceeded; 'src_tcp_zero_window_excd': Src TCP Zero-Window Exceeded; 'src_tcp_conn_prate_excd': Src TCP Rate: Conn Pkt Exceeded; 'src_udp_min_payload': Src UDP Payload Too Small; 'src_udp_max_payload': Src UDP Payload Too Large; 'src_udp_conn_prate_excd': Src UDP Rate: Conn Pkt Exceeded; 'src_udp_ntp_monlist_req': Src UDP NTP Monlist Request; 'src_udp_ntp_monlist_resp': Src UDP NTP Monlist Response; 'src_udp_wellknown_sport_drop': Src UDP SrcPort Wellknown; 'src_udp_retry_init': Src UDP Auth: Retry Init; 'dst_udp_retry_gap_drop': UDP Auth: Retry-Gap Dropped; 'dst_udp_retry_fail': UDP Auth: Retry Timeout; 'dst_tcp_session_aged': TCP Sessions Aged; 'dst_udp_session_aged': UDP Sessions Aged; 'dst_tcp_conn_close': TCP Connections Closed; 'dst_tcp_conn_close_half_open': TCP Half Open Connections Closed; 'dst_l4_tcp_auth': TCP Dst L4-Type Auth: SYN Cookie Sent; 'tcp_l4_syn_cookie_fail': TCP Dst L4-Type Auth: SYN Cookie Failed; 'tcp_l4_rst_cookie_fail': TCP Dst L4-Type Auth: RST Cookie Failed; 'tcp_l4_unauth_drop': TCP Dst L4-Type Auth: Unauth Dropped; 'dst_drop_frag_pkt': Dst Fragmented Packets Dropped; 'src_tcp_filter_action_blacklist': Src TCP Filter Action Blacklist; 'src_tcp_filter_action_whitelist': Src TCP Filter Action WL; 'src_tcp_filter_action_drop': Src TCP Filter Action Drop; 'src_tcp_filter_action_default_pass': Src TCP Filter Action Default Pass; 'src_udp_filter_action_blacklist': Src UDP Filter Action Blacklist; 'src_udp_filter_action_whitelist': Src UDP Filter Action WL; 'src_udp_filter_action_drop': Src UDP Filter Action Drop; 'src_udp_filter_action_default_pass': Src UDP Filter Action Default Pass; 'src_other_filter_action_blacklist': Src OTHER Filter Action Blacklist; 'src_other_filter_action_whitelist': Src OTHER Filter Action WL; 'src_other_filter_action_drop': Src OTHER Filter Action Drop; 'src_other_filter_action_default_pass': Src OTHER Filter Action Default Pass; 'tcp_invalid_syn': TCP Invalid SYN Received; 'dst_tcp_conn_close_w_rst': TCP RST Connections Closed; 'dst_tcp_conn_close_w_fin': TCP FIN Connections Closed; 'dst_tcp_conn_close_w_idle': TCP Idle Connections Closed; 'dst_tcp_conn_create_from_syn': TCP Connections Created From SYN; 'dst_tcp_conn_create_from_ack': TCP Connections Created From ACK; 'src_frag_drop': Src Fragmented Packets Dropped; 'dst_l4_tcp_blacklist_drop': Dst L4-type TCP Blacklist Dropped; 'dst_l4_udp_blacklist_drop': Dst L4-type UDP Blacklist Dropped; 'dst_l4_icmp_blacklist_drop': Dst L4-type ICMP Blacklist Dropped; 'dst_l4_other_blacklist_drop': Dst L4-type OTHER Blacklist Dropped; 'src_l4_tcp_blacklist_drop': Src L4-type TCP Blacklist Dropped; 'src_l4_udp_blacklist_drop': Src L4-type UDP Blacklist Dropped; 'src_l4_icmp_blacklist_drop': Src L4-type ICMP Blacklist Dropped; 'src_l4_other_blacklist_drop': Src L4-type OTHER Blacklist Dropped; 'drop_frag_timeout_drop': Fragment Reassemble Timeout Drop; 'dst_port_kbit_rate_exceed_pkt': Port Rate: KiBit Pkt Exceeded; 'dst_tcp_bytes_rcv': TCP Total Bytes Received; 'dst_udp_bytes_rcv': UDP Total Bytes Received; 'dst_icmp_bytes_rcv': ICMP Total Bytes Received; 'dst_other_bytes_rcv': OTHER Total Bytes Received; 'dst_tcp_bytes_sent': TCP Total Bytes Forwarded; 'dst_udp_bytes_sent': UDP Total Bytes Forwarded; 'dst_icmp_bytes_sent': ICMP Total Bytes Forwarded; 'dst_other_bytes_sent': OTHER Total Bytes Forwarded; 'dst_udp_auth_drop': UDP Auth: Dropped; 'dst_tcp_auth_drop': TCP Auth: Dropped; 'dst_tcp_auth_resp': TCP Auth: Responded; 'inbound_pkt_drop': Inbound: Packets Dropped; 'dst_entry_pkt_rate_exceed': Entry Rate: Packet Exceeded; 'dst_entry_kbit_rate_exceed': Entry Rate: KiBit Exceeded; 'dst_entry_conn_limit_exceed': Entry Limit: Conn Exceeded; 'dst_entry_conn_rate_exceed': Entry Rate: Conn Exceeded; 'dst_entry_frag_pkt_rate_exceed': Entry Rate: Frag Packet Exceeded; 'dst_icmp_any_exceed': ICMP Rate: Total Exceed; 'dst_other_any_exceed': OTHER Rate: Total Exceed; 'src_dst_pair_entry_total': Src-Dst Pair Entry Total Count; 'src_dst_pair_entry_udp': Src-Dst Pair Entry UDP Count; 'src_dst_pair_entry_tcp': Src-Dst Pair Entry TCP Count; 'src_dst_pair_entry_icmp': Src-Dst Pair Entry ICMP Count; 'src_dst_pair_entry_other': Src-Dst Pair Entry OTHER Count; 'dst_clist_overflow_policy_at_learning': Dst Src-Based Overflow Policy Hit; 'tcp_rexmit_syn_limit_drop': TCP SYN Retransmit Exceeded Drop; 'tcp_rexmit_syn_limit_bl': TCP SYN Retransmit Exceeded Blacklist; 'dst_tcp_wellknown_sport_drop': TCP SrcPort Wellknown; 'src_tcp_wellknown_sport_drop': Src TCP SrcPort Wellknown; 'dst_frag_rcvd': Fragmented Packets Received; 'no_policy_class_list_match': No Policy Class-list Match; 'src_udp_retry_gap_drop': Src UDP Auth: Retry-Gap Dropped; 'dst_entry_kbit_rate_exceed_count': Entry Rate: KiBit Exceeded Count; 'dst_port_undef_hit': Dst Port Undefined Hit; 'dst_tcp_action_on_ack_timeout': TCP Auth: ACK Retry Timeout; 'dst_tcp_action_on_ack_reset': TCP Auth: ACK Retry Timeout Reset; 'dst_tcp_action_on_ack_blacklist': TCP Auth: ACK Retry Timeout Blacklisted; 'src_tcp_action_on_ack_timeout': Src TCP Auth: ACK Retry Timeout; 'src_tcp_action_on_ack_reset': Src TCP Auth: ACK Retry Timeout Reset; 'src_tcp_action_on_ack_blacklist': Src TCP Auth: ACK Retry Timeout Blacklisted; 'dst_tcp_action_on_syn_timeout': TCP Auth: SYN Retry Timeout; 'dst_tcp_action_on_syn_reset': TCP Auth: SYN Retry Timeout Reset; 'dst_tcp_action_on_syn_blacklist': TCP Auth: SYN Retry Timeout Blacklisted; 'src_tcp_action_on_syn_timeout': Src TCP Auth: SYN Retry Timeout; 'src_tcp_action_on_syn_reset': Src TCP Auth: SYN Retry Timeout Reset; 'src_tcp_action_on_syn_blacklist': Src TCP Auth: SYN Retry Timeout Blacklisted; 'dst_udp_frag_pkt_rate_exceed': UDP Dst L4-Type Rate: Frag Exceeded; 'dst_udp_frag_src_rate_drop': UDP Src Rate: Frag Exceeded; 'dst_tcp_frag_pkt_rate_exceed': TCP Dst L4-Type Rate: Frag Exceeded; 'dst_tcp_frag_src_rate_drop': TCP Src Rate: Frag Exceeded; 'dst_icmp_frag_pkt_rate_exceed': ICMP Dst L4-Type Rate: Frag Exceeded; 'dst_icmp_frag_src_rate_drop': ICMP Src Rate: Frag Exceeded; 'sflow_internal_samples_packed': Sflow Internal Samples Packed; 'sflow_external_samples_packed': Sflow External Samples Packed; 'sflow_internal_packets_sent': Sflow Internal Packets Sent; 'sflow_external_packets_sent': Sflow External Packets Sent; 'dns_outbound_total_query': DNS Outbound Total Query; 'dns_outbound_query_malformed': DNS Outbound Query Malformed; 'dns_outbound_query_resp_chk_failed': DNS Outbound Query Resp Check Failed; 'dns_outbound_query_resp_chk_blacklisted': DNS Outbound Query Resp Check Blacklisted; 'dns_outbound_query_resp_chk_refused_sent': DNS Outbound Query Resp Check REFUSED Sent; 'dns_outbound_query_resp_chk_reset_sent': DNS Outbound Query Resp Check RESET Sent; 'dns_outbound_query_resp_chk_no_resp_sent': DNS Outbound Query Resp Check No Response Sent; 'dns_outbound_query_resp_size_exceed': DNS Outbound Query Response Size Exceed; 'dns_outbound_query_sess_timed_out': DNS Outbound Query Session Timed Out; 'dst_exceed_action_tunnel': Entry Exceed Action: Tunnel; 'src_udp_auth_timeout': Src UDP Auth: Retry Timeout; 'src_udp_retry_pass': Src UDP Retry Passed;",
						},
						"counters3": {
							Type: schema.TypeString, Optional: true, Description: "'dst_hw_drop_rule_insert': Dst Hardware Drop Rules Inserted; 'dst_hw_drop_rule_remove': Dst Hardware Drop Rules Removed; 'src_hw_drop_rule_insert': Src Hardware Drop Rules Inserted; 'src_hw_drop_rule_remove': Src Hardware Drop Rules Removed; 'prog_first_req_time_exceed': Req-Resp: First Request Time Exceed; 'prog_req_resp_time_exceed': Req-Resp: Request to Response Time Exceed; 'prog_request_len_exceed': Req-Resp: Request Length Exceed; 'prog_response_len_exceed': Req-Resp: Response Length Exceed; 'prog_resp_req_ratio_exceed': Req-Resp: Response to Request Ratio Exceed; 'prog_resp_req_time_exceed': Req-Resp: Response to Request Time Exceed; 'entry_sync_message_received': Entry Sync Message Received; 'entry_sync_message_sent': Entry Sync Message Sent; 'prog_conn_sent_exceed': Connection: Sent Exceed; 'prog_conn_rcvd_exceed': Connection: Received Exceed; 'prog_conn_time_exceed': Connection: Time Exceed; 'prog_conn_rcvd_sent_ratio_exceed': Connection: Received to Sent Ratio Exceed; 'prog_win_sent_exceed': Time Window: Sent Exceed; 'prog_win_rcvd_exceed': Time Window: Received Exceed; 'prog_win_rcvd_sent_ratio_exceed': Time Window: Received to Sent Exceed; 'prog_exceed_drop': Req-Resp: Violation Exceed Dropped; 'prog_exceed_bl': Req-Resp: Violation Exceed Blacklisted; 'prog_conn_exceed_drop': Connection: Violation Exceed Dropped; 'prog_conn_exceed_bl': Connection: Violation Exceed Blacklisted; 'prog_win_exceed_drop': Time Window: Violation Exceed Dropped; 'prog_win_exceed_bl': Time Window: Violation Exceed Blacklisted; 'dst_exceed_action_drop': Entry Exceed Action: Dropped; 'prog_conn_samples': Sample Collected: Connection; 'prog_req_samples': Sample Collected: Req-Resp; 'prog_win_samples': Sample Collected: Time Window; 'prog_conn_samples_processed': Sample Processed: Connnection; 'prog_req_samples_processed': Sample Processed: Req-Resp; 'prog_win_samples_processed': Sample Processed: Time Window; 'src_hw_drop': Src Hardware Packets Dropped; 'dst_tcp_auth_rst': TCP Auth: Reset; 'dst_src_learn_overflow': Src Dynamic Entry Count Overflow; 'tcp_fwd_sent': TCP Inbound Packets Forwarded; 'udp_fwd_sent': UDP Inbound Packets Forwarded;",
						},
					},
				},
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"sflow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"polling": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sflow_packets": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow packet-level counter polling. WARNING: Entry level Sflow polling might induce heavy CPU load depending on the tota",
									},
									"sflow_layer_4": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow Layer 4 counter polling. WARNING: Entry level Sflow polling might induce heavy CPU load depending on the total num",
									},
									"sflow_tcp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sflow_tcp_basic": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow basic TCP counter polling. WARNING: Entry level Sflow polling might induce heavy CPU load depending on the total n",
												},
												"sflow_tcp_stateful": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow stateful TCP counter polling. WARNING: Entry level Sflow polling might induce heavy CPU load depending on the tota",
												},
											},
										},
									},
									"sflow_http": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow HTTP counter polling. WARNING: Entry level Sflow polling might induce heavy CPU load depending on the total number",
									},
									"sflow_undef_port_hit_stats": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow undefined-port-hit-statistics polling",
									},
									"sflow_undef_port_hit_stats_brief": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow undefined-port-hit-statistics polling in brief mode",
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
					},
				},
			},
			"source_nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Configure source NAT",
			},
			"src_dst_pair": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure default",
						},
						"bypass": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
						},
						"exceed_log_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
									},
								},
							},
						},
						"log_periodic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logging": {
										Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
									},
								},
							},
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"l4_type_src_dst_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
												},
												"other": {
													Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
												},
												"template_icmp_v4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
												},
												"template_icmp_v6": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
						"app_type_src_dst_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
			"src_dst_pair_class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_name": {
							Type: schema.TypeString, Required: true, Description: "Class-list name",
						},
						"exceed_log_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
									},
								},
							},
						},
						"log_periodic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"l4_type_src_dst_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
												},
												"other": {
													Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
												},
												"template_icmp_v4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
												},
												"template_icmp_v6": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
						"app_type_src_dst_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
						"cid_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cid_num": {
										Type: schema.TypeInt, Required: true, Description: "Class-list id",
									},
									"exceed_log_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"log_enable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
												},
											},
										},
									},
									"log_periodic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"l4_type_src_dst_cid_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
												},
												"deny": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
															},
															"other": {
																Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
															},
															"template_icmp_v4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
															},
															"template_icmp_v6": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
									"app_type_src_dst_cid_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
												},
												"template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
															},
															"dns": {
																Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
															},
															"http": {
																Type: schema.TypeString, Optional: true, Description: "DDOS http template",
															},
															"sip": {
																Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
			"src_dst_pair_policy_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_based_policy_name": {
							Type: schema.TypeString, Required: true, Description: "Src-based-policy name",
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
									"bypass": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
									},
									"exceed_log_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"log_enable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
												},
											},
										},
									},
									"log_periodic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"logging": {
													Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
												},
											},
										},
									},
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"max_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry under class-list",
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
									"l4_type_src_dst_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
												},
												"deny": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tcp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
															},
															"udp": {
																Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
															},
															"other": {
																Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
															},
															"template_icmp_v4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
															},
															"template_icmp_v6": {
																Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
									"app_type_src_dst_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
												},
												"template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ssl_l4": {
																Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
															},
															"dns": {
																Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
															},
															"http": {
																Type: schema.TypeString, Optional: true, Description: "DDOS http template",
															},
															"sip": {
																Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
									"class_list_overflow_policy_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dummy_name": {
													Type: schema.TypeString, Required: true, Description: "'configuration': Configure src dst dynamic entry count overflow policy for class-list;",
												},
												"bypass": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
												},
												"exceed_log_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"log_enable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
															},
														},
													},
												},
												"log_periodic": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
												},
												"template": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logging": {
																Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
															},
														},
													},
												},
												"glid": {
													Type: schema.TypeString, Optional: true, Description: "Global limit ID",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"l4_type_src_dst_overflow_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"protocol": {
																Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
															},
															"deny": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
															},
															"glid": {
																Type: schema.TypeString, Optional: true, Description: "Global limit ID",
															},
															"template": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"tcp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
																		},
																		"udp": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
																		},
																		"other": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
																		},
																		"template_icmp_v4": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
																		},
																		"template_icmp_v6": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
												"app_type_src_dst_overflow_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"protocol": {
																Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
															},
															"template": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ssl_l4": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
																		},
																		"dns": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
																		},
																		"http": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS http template",
																		},
																		"sip": {
																			Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
					},
				},
			},
			"src_dst_pair_settings_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_types": {
							Type: schema.TypeString, Required: true, Description: "'all-types': Settings for all types (default or class-list);",
						},
						"age": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
						},
						"max_dynamic_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry",
						},
						"apply_policy_on_overflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
						},
						"unlimited_dynamic_entry_count": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
						},
						"enable_class_list_overflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for DST entry or each class-list",
						},
						"src_prefix_len": {
							Type: schema.TypeInt, Optional: true, Description: "Specify src prefix length for IPv6 (default: not set)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"l4_type_src_dst_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
									},
									"max_dynamic_entry_count": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry",
									},
									"apply_policy_on_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
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
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"outbound_src_tracking": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
						},
						"template": {
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
						"set_counter_base_val": {
							Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
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
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"template": {
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
						"set_counter_base_val": {
							Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
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
			"subnet_ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "IP Subnet",
			},
			"subnet_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 Subnet",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
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
			"traffic_distribution_mode": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Distribute traffic to one slot using default distribution mechanism; 'source-ip-based': Distribute traffic between slots, based on source ip;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDstEntryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntry(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntry(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntry(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntry(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntryCaptureConfigList(d []interface{}) []edpt.DdosDstEntryCaptureConfigList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryCaptureConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryCaptureConfigList
		oi.Name = in["name"].(string)
		oi.Mode = in["mode"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryDynamicEntryOverflowPolicyList(d []interface{}) []edpt.DdosDstEntryDynamicEntryOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryDynamicEntryOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryDynamicEntryOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Bypass = in["bypass"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntryDynamicEntryOverflowPolicyListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.Template = getObjectDdosDstEntryDynamicEntryOverflowPolicyListTemplate(in["template"].([]interface{}))
		oi.Glid = in["glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstList = getSliceDdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstList(in["l4_type_src_dst_list"].([]interface{}))
		oi.AppTypeSrcDstList = getSliceDdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstList(in["app_type_src_dst_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryDynamicEntryOverflowPolicyListExceedLogCfg(d []interface{}) edpt.DdosDstEntryDynamicEntryOverflowPolicyListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryDynamicEntryOverflowPolicyListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosDstEntryDynamicEntryOverflowPolicyListTemplate(d []interface{}) edpt.DdosDstEntryDynamicEntryOverflowPolicyListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryDynamicEntryOverflowPolicyListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstList(d []interface{}) []edpt.DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntryEnableTopK(d []interface{}) []edpt.DdosDstEntryEnableTopK {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryEnableTopK, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryEnableTopK
		oi.TopkType = in["topk_type"].(string)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryExceedLogCfg(d []interface{}) edpt.DdosDstEntryExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
		ret.LogWithSflow = in["log_with_sflow"].(int)
		ret.LogHighFrequency = in["log_high_frequency"].(int)
		ret.RateLimit = in["rate_limit"].(int)
	}
	return ret
}

func getObjectDdosDstEntryExceedLogDepCfg(d []interface{}) edpt.DdosDstEntryExceedLogDepCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryExceedLogDepCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExceedLogEnable = in["exceed_log_enable"].(int)
		ret.LogWithSflowDep = in["log_with_sflow_dep"].(int)
	}
	return ret
}

func getObjectDdosDstEntryGlidExceedAction(d []interface{}) edpt.DdosDstEntryGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryHwBlacklistBlocking178(d []interface{}) edpt.DdosDstEntryHwBlacklistBlocking178 {

	count1 := len(d)
	var ret edpt.DdosDstEntryHwBlacklistBlocking178
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstEnable = in["dst_enable"].(int)
		ret.SrcEnable = in["src_enable"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosDstEntryIpProtoList(d []interface{}) []edpt.DdosDstEntryIpProtoList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryIpProtoList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryIpProtoList
		oi.PortNum = in["port_num"].(int)
		oi.Deny = in["deny"].(int)
		oi.EspInspect = getObjectDdosDstEntryIpProtoListEspInspect(in["esp_inspect"].([]interface{}))
		oi.Glid = in["glid"].(string)
		oi.GlidExceedAction = getObjectDdosDstEntryIpProtoListGlidExceedAction(in["glid_exceed_action"].([]interface{}))
		oi.Template = getObjectDdosDstEntryIpProtoListTemplate(in["template"].([]interface{}))
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryIpProtoListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryIpProtoListEspInspect(d []interface{}) edpt.DdosDstEntryIpProtoListEspInspect {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoListEspInspect
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthAlgorithm = in["auth_algorithm"].(string)
		ret.EncryptAlgorithm = in["encrypt_algorithm"].(string)
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectDdosDstEntryIpProtoListGlidExceedAction(d []interface{}) edpt.DdosDstEntryIpProtoListGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoListGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryIpProtoListGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryIpProtoListGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryIpProtoListGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoListGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryIpProtoListTemplate(d []interface{}) edpt.DdosDstEntryIpProtoListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Other = in["other"].(string)
	}
	return ret
}

func getObjectDdosDstEntryIpProtoListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryIpProtoListIpFilteringPolicyOper {

	var ret edpt.DdosDstEntryIpProtoListIpFilteringPolicyOper
	return ret
}

func getSliceDdosDstEntryL4TypeList(d []interface{}) []edpt.DdosDstEntryL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidExceedAction = getObjectDdosDstEntryL4TypeListGlidExceedAction(in["glid_exceed_action"].([]interface{}))
		oi.Deny = in["deny"].(int)
		oi.MaxRexmitSynPerFlow = in["max_rexmit_syn_per_flow"].(int)
		oi.MaxRexmitSynPerFlowExceedAction = in["max_rexmit_syn_per_flow_exceed_action"].(string)
		oi.DisableSynAuth = in["disable_syn_auth"].(int)
		oi.SynAuth = in["syn_auth"].(string)
		oi.SynCookie = in["syn_cookie"].(int)
		oi.TcpResetClient = in["tcp_reset_client"].(int)
		oi.TcpResetServer = in["tcp_reset_server"].(int)
		oi.DropOnNoPortMatch = in["drop_on_no_port_match"].(string)
		oi.Stateful = in["stateful"].(int)
		oi.TunnelDecap = getObjectDdosDstEntryL4TypeListTunnelDecap(in["tunnel_decap"].([]interface{}))
		oi.TunnelRateLimit = getObjectDdosDstEntryL4TypeListTunnelRateLimit(in["tunnel_rate_limit"].([]interface{}))
		oi.DropFragPkt = in["drop_frag_pkt"].(int)
		oi.UndefinedPortHitStatistics = getObjectDdosDstEntryL4TypeListUndefinedPortHitStatistics(in["undefined_port_hit_statistics"].([]interface{}))
		oi.Template = getObjectDdosDstEntryL4TypeListTemplate(in["template"].([]interface{}))
		oi.DetectionEnable = in["detection_enable"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryL4TypeListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstEntryL4TypeListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstEntryL4TypeListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstEntryL4TypeListProgressionTracking(in["progression_tracking"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListGlidExceedAction(d []interface{}) edpt.DdosDstEntryL4TypeListGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeListGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryL4TypeListGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryL4TypeListGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeListGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListTunnelDecap(d []interface{}) edpt.DdosDstEntryL4TypeListTunnelDecap {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeListTunnelDecap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDecap = in["ip_decap"].(int)
		ret.GreDecap = in["gre_decap"].(int)
		ret.KeyCfg = getSliceDdosDstEntryL4TypeListTunnelDecapKeyCfg(in["key_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryL4TypeListTunnelDecapKeyCfg(d []interface{}) []edpt.DdosDstEntryL4TypeListTunnelDecapKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeListTunnelDecapKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeListTunnelDecapKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListTunnelRateLimit(d []interface{}) edpt.DdosDstEntryL4TypeListTunnelRateLimit {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeListTunnelRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpRateLimit = in["ip_rate_limit"].(int)
		ret.GreRateLimit = in["gre_rate_limit"].(int)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListUndefinedPortHitStatistics(d []interface{}) edpt.DdosDstEntryL4TypeListUndefinedPortHitStatistics {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeListUndefinedPortHitStatistics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UndefinedPortHitStatistics = in["undefined_port_hit_statistics"].(int)
		ret.ResetInterval = in["reset_interval"].(int)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListTemplate(d []interface{}) edpt.DdosDstEntryL4TypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryL4TypeListIpFilteringPolicyOper {

	var ret edpt.DdosDstEntryL4TypeListIpFilteringPolicyOper
	return ret
}

func getObjectDdosDstEntryL4TypeListPortInd(d []interface{}) edpt.DdosDstEntryL4TypeListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstEntryL4TypeListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryL4TypeListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstEntryL4TypeListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeListTopkSources(d []interface{}) edpt.DdosDstEntryL4TypeListTopkSources {

	var ret edpt.DdosDstEntryL4TypeListTopkSources
	return ret
}

func getObjectDdosDstEntryL4TypeListProgressionTracking(d []interface{}) edpt.DdosDstEntryL4TypeListProgressionTracking {

	var ret edpt.DdosDstEntryL4TypeListProgressionTracking
	return ret
}

func getSliceDdosDstEntryPortList(d []interface{}) []edpt.DdosDstEntryPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.DetectionEnable = in["detection_enable"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.GlidExceedAction = getObjectDdosDstEntryPortListGlidExceedAction(in["glid_exceed_action"].([]interface{}))
		oi.DnsCache = in["dns_cache"].(string)
		oi.Template = getObjectDdosDstEntryPortListTemplate(in["template"].([]interface{}))
		oi.Sflow = getObjectDdosDstEntryPortListSflow(in["sflow"].([]interface{}))
		oi.CaptureConfig = getObjectDdosDstEntryPortListCaptureConfig(in["capture_config"].([]interface{}))
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PortInd = getObjectDdosDstEntryPortListPortInd(in["port_ind"].([]interface{}))
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryPortListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.TopkSources = getObjectDdosDstEntryPortListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstEntryPortListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.SignatureExtraction = getObjectDdosDstEntryPortListSignatureExtraction(in["signature_extraction"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstEntryPortListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstEntryPortListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortListGlidExceedAction(d []interface{}) edpt.DdosDstEntryPortListGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryPortListGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortListGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryPortListGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortListTemplate(d []interface{}) edpt.DdosDstEntryPortListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortListSflow(d []interface{}) edpt.DdosDstEntryPortListSflow {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListSflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Polling = getObjectDdosDstEntryPortListSflowPolling(in["polling"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortListSflowPolling(d []interface{}) edpt.DdosDstEntryPortListSflowPolling {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListSflowPolling
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPackets = in["sflow_packets"].(int)
		ret.SflowTcp = getObjectDdosDstEntryPortListSflowPollingSflowTcp(in["sflow_tcp"].([]interface{}))
		ret.SflowHttp = in["sflow_http"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortListSflowPollingSflowTcp(d []interface{}) edpt.DdosDstEntryPortListSflowPollingSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListSflowPollingSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortListCaptureConfig(d []interface{}) edpt.DdosDstEntryPortListCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortListPortInd(d []interface{}) edpt.DdosDstEntryPortListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstEntryPortListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstEntryPortListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryPortListIpFilteringPolicyOper {

	var ret edpt.DdosDstEntryPortListIpFilteringPolicyOper
	return ret
}

func getObjectDdosDstEntryPortListTopkSources(d []interface{}) edpt.DdosDstEntryPortListTopkSources {

	var ret edpt.DdosDstEntryPortListTopkSources
	return ret
}

func getObjectDdosDstEntryPortListProgressionTracking(d []interface{}) edpt.DdosDstEntryPortListProgressionTracking {

	var ret edpt.DdosDstEntryPortListProgressionTracking
	return ret
}

func getObjectDdosDstEntryPortListSignatureExtraction(d []interface{}) edpt.DdosDstEntryPortListSignatureExtraction {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListSignatureExtraction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.ManualMode = in["manual_mode"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstEntryPortListPatternRecognition(d []interface{}) edpt.DdosDstEntryPortListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.Mode = in["mode"].(string)
		ret.Sensitivity = in["sensitivity"].(string)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterInactiveThreshold = in["filter_inactive_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstEntryPortListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstEntryPortListPatternRecognitionPuDetails {

	var ret edpt.DdosDstEntryPortListPatternRecognitionPuDetails
	return ret
}

func getSliceDdosDstEntryPortRangeList(d []interface{}) []edpt.DdosDstEntryPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeList
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.DetectionEnable = in["detection_enable"].(int)
		oi.EnableTopK = in["enable_top_k"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		oi.Glid = in["glid"].(string)
		oi.GlidExceedAction = getObjectDdosDstEntryPortRangeListGlidExceedAction(in["glid_exceed_action"].([]interface{}))
		oi.Template = getObjectDdosDstEntryPortRangeListTemplate(in["template"].([]interface{}))
		oi.Sflow = getObjectDdosDstEntryPortRangeListSflow(in["sflow"].([]interface{}))
		oi.CaptureConfig = getObjectDdosDstEntryPortRangeListCaptureConfig(in["capture_config"].([]interface{}))
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		oi.IpFilteringPolicy = in["ip_filtering_policy"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IpFilteringPolicyOper = getObjectDdosDstEntryPortRangeListIpFilteringPolicyOper(in["ip_filtering_policy_oper"].([]interface{}))
		oi.PortInd = getObjectDdosDstEntryPortRangeListPortInd(in["port_ind"].([]interface{}))
		oi.TopkSources = getObjectDdosDstEntryPortRangeListTopkSources(in["topk_sources"].([]interface{}))
		oi.ProgressionTracking = getObjectDdosDstEntryPortRangeListProgressionTracking(in["progression_tracking"].([]interface{}))
		oi.PatternRecognition = getObjectDdosDstEntryPortRangeListPatternRecognition(in["pattern_recognition"].([]interface{}))
		oi.PatternRecognitionPuDetails = getObjectDdosDstEntryPortRangeListPatternRecognitionPuDetails(in["pattern_recognition_pu_details"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListGlidExceedAction(d []interface{}) edpt.DdosDstEntryPortRangeListGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryPortRangeListGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryPortRangeListGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListTemplate(d []interface{}) edpt.DdosDstEntryPortRangeListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListSflow(d []interface{}) edpt.DdosDstEntryPortRangeListSflow {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListSflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Polling = getObjectDdosDstEntryPortRangeListSflowPolling(in["polling"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListSflowPolling(d []interface{}) edpt.DdosDstEntryPortRangeListSflowPolling {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListSflowPolling
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPackets = in["sflow_packets"].(int)
		ret.SflowTcp = getObjectDdosDstEntryPortRangeListSflowPollingSflowTcp(in["sflow_tcp"].([]interface{}))
		ret.SflowHttp = in["sflow_http"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListSflowPollingSflowTcp(d []interface{}) edpt.DdosDstEntryPortRangeListSflowPollingSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListSflowPollingSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListCaptureConfig(d []interface{}) edpt.DdosDstEntryPortRangeListCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryPortRangeListIpFilteringPolicyOper {

	var ret edpt.DdosDstEntryPortRangeListIpFilteringPolicyOper
	return ret
}

func getObjectDdosDstEntryPortRangeListPortInd(d []interface{}) edpt.DdosDstEntryPortRangeListPortInd {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstEntryPortRangeListPortIndSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangeListPortIndSamplingEnable(d []interface{}) []edpt.DdosDstEntryPortRangeListPortIndSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeListPortIndSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeListPortIndSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListTopkSources(d []interface{}) edpt.DdosDstEntryPortRangeListTopkSources {

	var ret edpt.DdosDstEntryPortRangeListTopkSources
	return ret
}

func getObjectDdosDstEntryPortRangeListProgressionTracking(d []interface{}) edpt.DdosDstEntryPortRangeListProgressionTracking {

	var ret edpt.DdosDstEntryPortRangeListProgressionTracking
	return ret
}

func getObjectDdosDstEntryPortRangeListPatternRecognition(d []interface{}) edpt.DdosDstEntryPortRangeListPatternRecognition {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeListPatternRecognition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.Mode = in["mode"].(string)
		ret.Sensitivity = in["sensitivity"].(string)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterInactiveThreshold = in["filter_inactive_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstEntryPortRangeListPatternRecognitionPuDetails(d []interface{}) edpt.DdosDstEntryPortRangeListPatternRecognitionPuDetails {

	var ret edpt.DdosDstEntryPortRangeListPatternRecognitionPuDetails
	return ret
}

func getSliceDdosDstEntrySamplingEnable(d []interface{}) []edpt.DdosDstEntrySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		oi.Counters3 = in["counters3"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySflow(d []interface{}) edpt.DdosDstEntrySflow {

	count1 := len(d)
	var ret edpt.DdosDstEntrySflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Polling = getObjectDdosDstEntrySflowPolling(in["polling"].([]interface{}))
		ret.Collector = getSliceDdosDstEntrySflowCollector(in["collector"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntrySflowPolling(d []interface{}) edpt.DdosDstEntrySflowPolling {

	count1 := len(d)
	var ret edpt.DdosDstEntrySflowPolling
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPackets = in["sflow_packets"].(int)
		ret.SflowLayer4 = in["sflow_layer_4"].(int)
		ret.SflowTcp = getObjectDdosDstEntrySflowPollingSflowTcp(in["sflow_tcp"].([]interface{}))
		ret.SflowHttp = in["sflow_http"].(int)
		ret.SflowUndefPortHitStats = in["sflow_undef_port_hit_stats"].(int)
		ret.SflowUndefPortHitStatsBrief = in["sflow_undef_port_hit_stats_brief"].(int)
	}
	return ret
}

func getObjectDdosDstEntrySflowPollingSflowTcp(d []interface{}) edpt.DdosDstEntrySflowPollingSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstEntrySflowPollingSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySflowCollector(d []interface{}) []edpt.DdosDstEntrySflowCollector {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySflowCollector, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySflowCollector
		oi.SflowName = in["sflow_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPair179(d []interface{}) edpt.DdosDstEntrySrcDstPair179 {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPair179
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = in["default"].(int)
		ret.Bypass = in["bypass"].(int)
		ret.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairExceedLogCfg180(in["exceed_log_cfg"].([]interface{}))
		ret.LogPeriodic = in["log_periodic"].(int)
		ret.Template = getObjectDdosDstEntrySrcDstPairTemplate181(in["template"].([]interface{}))
		ret.Glid = in["glid"].(string)
		//omit uuid
		ret.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairL4TypeSrcDstList182(in["l4_type_src_dst_list"].([]interface{}))
		ret.AppTypeSrcDstList = getSliceDdosDstEntrySrcDstPairAppTypeSrcDstList184(in["app_type_src_dst_list"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairExceedLogCfg180(d []interface{}) edpt.DdosDstEntrySrcDstPairExceedLogCfg180 {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairExceedLogCfg180
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairTemplate181(d []interface{}) edpt.DdosDstEntrySrcDstPairTemplate181 {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairTemplate181
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairL4TypeSrcDstList182(d []interface{}) []edpt.DdosDstEntrySrcDstPairL4TypeSrcDstList182 {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairL4TypeSrcDstList182, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairL4TypeSrcDstList182
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairL4TypeSrcDstListTemplate183(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairL4TypeSrcDstListTemplate183(d []interface{}) edpt.DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate183 {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate183
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairAppTypeSrcDstList184(d []interface{}) []edpt.DdosDstEntrySrcDstPairAppTypeSrcDstList184 {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairAppTypeSrcDstList184, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairAppTypeSrcDstList184
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairAppTypeSrcDstListTemplate185(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairAppTypeSrcDstListTemplate185(d []interface{}) edpt.DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate185 {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate185
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairClassListListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairClassListListL4TypeSrcDstList(in["l4_type_src_dst_list"].([]interface{}))
		oi.AppTypeSrcDstList = getSliceDdosDstEntrySrcDstPairClassListListAppTypeSrcDstList(in["app_type_src_dst_list"].([]interface{}))
		oi.CidList = getSliceDdosDstEntrySrcDstPairClassListListCidList(in["cid_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListListL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListListL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListListL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListListL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListListL4TypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListListL4TypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListListL4TypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListListL4TypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListListAppTypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListListAppTypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListListAppTypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListListAppTypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListListAppTypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListListAppTypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListListAppTypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListListAppTypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListListCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListListCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListListCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListListCidList
		oi.CidNum = in["cid_num"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairClassListListCidListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstCidList = getSliceDdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidList(in["l4_type_src_dst_cid_list"].([]interface{}))
		oi.AppTypeSrcDstCidList = getSliceDdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidList(in["app_type_src_dst_cid_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListListCidListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListListCidListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListListCidListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyList
		oi.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.PolicyClassListList = getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListList(in["policy_class_list_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Bypass = in["bypass"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListTemplate(in["template"].([]interface{}))
		oi.Glid = in["glid"].(string)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstList(in["l4_type_src_dst_list"].([]interface{}))
		oi.AppTypeSrcDstList = getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstList(in["app_type_src_dst_list"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Bypass = in["bypass"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListTemplate(in["template"].([]interface{}))
		oi.Glid = in["glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList(in["l4_type_src_dst_overflow_list"].([]interface{}))
		oi.AppTypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList(in["app_type_src_dst_overflow_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairSettingsList(d []interface{}) []edpt.DdosDstEntrySrcDstPairSettingsList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairSettingsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairSettingsList
		oi.AllTypes = in["all_types"].(string)
		oi.Age = in["age"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		oi.UnlimitedDynamicEntryCount = in["unlimited_dynamic_entry_count"].(int)
		oi.EnableClassListOverflow = in["enable_class_list_overflow"].(int)
		oi.SrcPrefixLen = in["src_prefix_len"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairSettingsListL4TypeSrcDstList(in["l4_type_src_dst_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairSettingsListL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairSettingsListL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairSettingsListL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairSettingsListL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntrySrcPortList(d []interface{}) []edpt.DdosDstEntrySrcPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.OutboundSrcTracking = in["outbound_src_tracking"].(string)
		oi.Template = getObjectDdosDstEntrySrcPortListTemplate(in["template"].([]interface{}))
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcPortListTemplate(d []interface{}) edpt.DdosDstEntrySrcPortListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcPortListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
		ret.SrcDns = in["src_dns"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcPortRangeList(d []interface{}) []edpt.DdosDstEntrySrcPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcPortRangeList
		oi.SrcPortRangeStart = in["src_port_range_start"].(int)
		oi.SrcPortRangeEnd = in["src_port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcPortRangeListTemplate(in["template"].([]interface{}))
		oi.SetCounterBaseVal = in["set_counter_base_val"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcPortRangeListTemplate(d []interface{}) edpt.DdosDstEntrySrcPortRangeListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcPortRangeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
	}
	return ret
}

func getObjectDdosDstEntryTemplate(d []interface{}) edpt.DdosDstEntryTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getObjectDdosDstEntryTopkDestinations186(d []interface{}) edpt.DdosDstEntryTopkDestinations186 {

	var ret edpt.DdosDstEntryTopkDestinations186
	return ret
}

func dataToEndpointDdosDstEntry(d *schema.ResourceData) edpt.DdosDstEntry {
	var ret edpt.DdosDstEntry
	ret.Inst.AdvertisedEnable = d.Get("advertised_enable").(int)
	ret.Inst.BlackholeOnGlidExceed = d.Get("blackhole_on_glid_exceed").(int)
	ret.Inst.CaptureConfigList = getSliceDdosDstEntryCaptureConfigList(d.Get("capture_config_list").([]interface{}))
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.DestNatIp = d.Get("dest_nat_ip").(string)
	ret.Inst.DestNatIpv6 = d.Get("dest_nat_ipv6").(string)
	ret.Inst.DropDisable = d.Get("drop_disable").(int)
	ret.Inst.DropDisableFwdImmediate = d.Get("drop_disable_fwd_immediate").(int)
	ret.Inst.DropFragPkt = d.Get("drop_frag_pkt").(int)
	ret.Inst.DropOnNoSrcDstDefault = d.Get("drop_on_no_src_dst_default").(int)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	ret.Inst.DynamicEntryOverflowPolicyList = getSliceDdosDstEntryDynamicEntryOverflowPolicyList(d.Get("dynamic_entry_overflow_policy_list").([]interface{}))
	ret.Inst.EnableTopK = getSliceDdosDstEntryEnableTopK(d.Get("enable_top_k").([]interface{}))
	ret.Inst.ExceedLogCfg = getObjectDdosDstEntryExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.ExceedLogDepCfg = getObjectDdosDstEntryExceedLogDepCfg(d.Get("exceed_log_dep_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.GlidExceedAction = getObjectDdosDstEntryGlidExceedAction(d.Get("glid_exceed_action").([]interface{}))
	ret.Inst.HwBlacklistBlocking = getObjectDdosDstEntryHwBlacklistBlocking178(d.Get("hw_blacklist_blocking").([]interface{}))
	ret.Inst.InboundForwardDscp = d.Get("inbound_forward_dscp").(int)
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.IpProtoList = getSliceDdosDstEntryIpProtoList(d.Get("ip_proto_list").([]interface{}))
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.L4TypeList = getSliceDdosDstEntryL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.OperationalMode = d.Get("operational_mode").(string)
	ret.Inst.OutboundForwardDscp = d.Get("outbound_forward_dscp").(int)
	ret.Inst.PatternRecognitionHwFilterEnable = d.Get("pattern_recognition_hw_filter_enable").(int)
	ret.Inst.PatternRecognitionSensitivity = d.Get("pattern_recognition_sensitivity").(string)
	ret.Inst.PortList = getSliceDdosDstEntryPortList(d.Get("port_list").([]interface{}))
	ret.Inst.PortRangeList = getSliceDdosDstEntryPortRangeList(d.Get("port_range_list").([]interface{}))
	ret.Inst.ReportingDisabled = d.Get("reporting_disabled").(int)
	ret.Inst.SamplingEnable = getSliceDdosDstEntrySamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.Sflow = getObjectDdosDstEntrySflow(d.Get("sflow").([]interface{}))
	ret.Inst.SourceNatPool = d.Get("source_nat_pool").(string)
	ret.Inst.SrcDstPair = getObjectDdosDstEntrySrcDstPair179(d.Get("src_dst_pair").([]interface{}))
	ret.Inst.SrcDstPairClassListList = getSliceDdosDstEntrySrcDstPairClassListList(d.Get("src_dst_pair_class_list_list").([]interface{}))
	ret.Inst.SrcDstPairPolicyList = getSliceDdosDstEntrySrcDstPairPolicyList(d.Get("src_dst_pair_policy_list").([]interface{}))
	ret.Inst.SrcDstPairSettingsList = getSliceDdosDstEntrySrcDstPairSettingsList(d.Get("src_dst_pair_settings_list").([]interface{}))
	ret.Inst.SrcPortList = getSliceDdosDstEntrySrcPortList(d.Get("src_port_list").([]interface{}))
	ret.Inst.SrcPortRangeList = getSliceDdosDstEntrySrcPortRangeList(d.Get("src_port_range_list").([]interface{}))
	ret.Inst.SubnetIpAddr = d.Get("subnet_ip_addr").(string)
	ret.Inst.SubnetIpv6Addr = d.Get("subnet_ipv6_addr").(string)
	ret.Inst.Template = getObjectDdosDstEntryTemplate(d.Get("template").([]interface{}))
	ret.Inst.TopkDestinations = getObjectDdosDstEntryTopkDestinations186(d.Get("topk_destinations").([]interface{}))
	ret.Inst.TrafficDistributionMode = d.Get("traffic_distribution_mode").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
