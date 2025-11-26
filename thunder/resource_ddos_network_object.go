package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObject() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object`: Configure DDoS a static Monitor Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectCreate,
		UpdateContext: resourceDdosNetworkObjectUpdate,
		ReadContext:   resourceDdosNetworkObjectRead,
		DeleteContext: resourceDdosNetworkObjectDelete,

		Schema: map[string]*schema.Schema{
			"anomaly_detection_trigger": {
				Type: schema.TypeString, Optional: true, Description: "'all': Use both learned and static thresholds (static thresholds take precedence); 'static-threshold-only': Use static thresholds only;",
			},
			"enable_top_k": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topk_type": {
							Type: schema.TypeString, Optional: true, Description: "'destination': Topk destination IP;",
						},
						"topk_dst_num_records": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
						},
						"topk_sort_key": {
							Type: schema.TypeString, Optional: true, Default: "average", Description: "'average': window average; 'max-peak': max peak;",
						},
					},
				},
			},
			"flooding_multiplier": {
				Type: schema.TypeInt, Optional: true, Description: "multiplier for flooding detection threshold in network objects (default 2x threshold)",
			},
			"histogram_mode": {
				Type: schema.TypeString, Optional: true, Description: "'off': histogram feature disabled; 'monitor': histogram feature enabled with anomaly escalation; 'observe': histogram feature enabled and observe only;",
			},
			"host_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Forward packet rate of per host",
						},
						"host_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Forward bit rate of per host",
						},
						"host_rev_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse packet rate of per host",
						},
						"host_rev_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse bit rate of per host",
						},
						"host_undiscovered_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Undiscovered forward packet rate of per host",
						},
						"host_flow_count": {
							Type: schema.TypeInt, Optional: true, Description: "Flow count of per host",
						},
						"host_syn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "SYN packet rate of per host",
						},
						"host_fin_rate": {
							Type: schema.TypeInt, Optional: true, Description: "FIN packet rate of per host",
						},
						"host_rst_rate": {
							Type: schema.TypeInt, Optional: true, Description: "RST packet rate of per host",
						},
						"host_tcp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Tcp packet rate of per host",
						},
						"host_udp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Udp packet rate of per host",
						},
						"host_icmp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP packet rate of per host",
						},
						"host_undiscovered_host_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "forward packet rate of per undiscovered host",
						},
						"host_undiscovered_host_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Forward bit rate of per undiscovered host",
						},
					},
				},
			},
			"host_sport_discovery": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable source port discovery.; 'disable': Disable source port discovery.;",
			},
			"indicators_to_monitor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"monitor_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward packet rate",
						},
						"monitor_bit_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward bit rate",
						},
						"monitor_rev_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse packet rate",
						},
						"monitor_rev_bit_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse bit rate",
						},
						"monitor_undiscovered_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Undiscovered forward packet rate",
						},
						"monitor_flow_count": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Flow count",
						},
						"monitor_syn_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "SYN packet rate",
						},
						"monitor_fin_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "FIN packet rate",
						},
						"monitor_rst_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "RST packet rate",
						},
						"monitor_tcp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP packet rate",
						},
						"monitor_udp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "UDP packet rate",
						},
						"monitor_icmp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ICMP packet rate",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ip_addr": {
							Type: schema.TypeString, Required: true, Description: "IP Subnet, supported prefix range is from 8 to 32",
						},
						"prefix_anomaly_threshold": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix_pkt_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of the prefix subnet",
									},
									"prefix_bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of the prefix subnet",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_rate': PPS; 'bit_rate': B(bits)PS;",
									},
								},
							},
						},
					},
				},
			},
			"ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Required: true, Description: "IPV6 Subnet, supported prefix range is from 40 to 64",
						},
						"prefix_anomaly_threshold": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix_pkt_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of the prefix subnet",
									},
									"prefix_bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of the prefix subnet",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_rate': PPS; 'bit_rate': B(bits)PS;",
									},
								},
							},
						},
					},
				},
			},
			"network_object_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_object_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of the network-object",
						},
						"network_object_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of the network-object",
						},
					},
				},
			},
			"network_object_template": {
				Type: schema.TypeString, Optional: true, Description: "The template applied for the network-object",
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
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"operational_mode": {
				Type: schema.TypeString, Optional: true, Description: "'monitor': Monitor mode; 'learning': Learning mode;",
			},
			"relative_auto_break_down_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "percentage of parent node",
						},
						"permil": {
							Type: schema.TypeInt, Optional: true, Description: "permil of root node",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'subnet_learned': Subnet Entry Learned; 'subnet_aged': Subnet Entry Aged; 'subnet_create_fail': Subnet Entry Create Failures; 'ip_learned': IP Entry Learned; 'ip_aged': IP Entry Aged; 'ip_create_fail': IP Entry Create Failures; 'service_learned': Service Entry Learned; 'service_aged': Service Entry Aged; 'service_create_fail': Service Entry Create Failures; 'packet_rate': PPS; 'bit_rate': B(bits)PS; 'topk_allocate_fail': Topk Allocate Failures; 'sport_learned': Source Port Entry Learned; 'sport_aged': Source Port Entry Aged; 'sport_create_fail': Source Port Entry Create Failures; 'agent_group_learned': Agent Group Entry Learned; 'agent_group_aged': Agent Group Entry Aged; 'agent_group_create_fail': Agent Group Entry Create Failures; 'duplicate_sample_pkt_rcv': Duplicate Sample Packet Received;",
						},
					},
				},
			},
			"service_break_down_threshold_local": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"svc_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "percentage of parent ip node",
						},
					},
				},
			},
			"service_discovery": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable service discovery for hosts (default: enabled);",
			},
			"sport_anomaly_detection": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable source port anomaly detection (default: enabled);",
			},
			"sport_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"packet_rate_percentage": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"bit_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"bit_rate_percentage": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_addr": {
										Type: schema.TypeString, Required: true, Description: "Override threshold",
									},
									"packet_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate': Packet rate of a source port entry;",
									},
									"packet_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"bit_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate': Bit rate of a source port entry;",
									},
									"bit_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"packet_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
									},
									"packet_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
									},
									"bit_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"sport_num": {
										Type: schema.TypeInt, Required: true, Description: "Source port number",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
									},
									"ip_sport_packet_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate': Packet rate of a source port entry;",
									},
									"ip_sport_packet_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"ip_sport_bit_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate': Bit rate of a source port entry;",
									},
									"ip_sport_bit_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"ip_sport_packet_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
									},
									"ip_sport_packet_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"ip_sport_bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
									},
									"ip_sport_bit_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ipv6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_addr": {
										Type: schema.TypeString, Required: true, Description: "Override threshold",
									},
									"packet_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate': Packet rate of a source port entry;",
									},
									"packet_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"bit_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate': Bit rate of a source port entry;",
									},
									"bit_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"packet_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
									},
									"packet_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
									},
									"bit_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"sport_num": {
										Type: schema.TypeInt, Required: true, Description: "Source port number",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
									},
									"ip_sport_packet_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate': Packet rate of a source port entry;",
									},
									"ip_sport_packet_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"ip_sport_bit_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate': Bit rate of a source port entry;",
									},
									"ip_sport_bit_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"ip_sport_packet_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
									},
									"ip_sport_packet_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"ip_sport_bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
									},
									"ip_sport_bit_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"sport_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sport_num": {
										Type: schema.TypeInt, Required: true, Description: "Port Number",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
									},
									"packet_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate': Packet rate of a source port entry;",
									},
									"packet_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'packet-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"bit_rate_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate': Bit rate of a source port entry;",
									},
									"bit_rate_percentage_str": {
										Type: schema.TypeString, Required: true, Description: "'bit-rate-percentage': Percentage of source port entry's parent entry;",
									},
									"packet_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
									},
									"packet_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
									},
									"bit_rate_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
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
			"sport_discovery_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sport_heavy_hitter_percentage": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Percentage of the bit rate of undiscovered source ports (default: 10)",
						},
						"sport_discovery_bit_rate_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "Percentage of the bit rate of source port's parent entry",
						},
					},
				},
			},
			"sport_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"static_auto_break_down_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "packet rate of current node",
						},
					},
				},
			},
			"sub_network": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_network_v4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet_ip_addr": {
										Type: schema.TypeString, Required: true, Description: "IPv4 Subnet/host, supported prefix range is from 24 to 32",
									},
									"host_anomaly_threshold": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"static_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
												},
												"static_rev_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
												},
												"static_bit_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
												},
												"static_rev_bit_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
												},
												"static_undiscovered_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Undiscovered packet rate of per host",
												},
												"static_flow_count_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Flow count of per host",
												},
												"static_syn_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "SYN packet rate of per host",
												},
												"static_fin_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "FIN packet rate of per host",
												},
												"static_rst_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "RST packet rate of per host",
												},
												"static_tcp_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "TCP packet rate of per host",
												},
												"static_udp_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "UDP packet rate of per host",
												},
												"static_icmp_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "ICMP packet rate of per host",
												},
												"static_undiscovered_host_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "packet rate of per undiscovered host",
												},
												"static_undiscovered_host_bit_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of per undiscovered host",
												},
											},
										},
									},
									"sub_network_anomaly_threshold": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"static_sub_network_pkt_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Packet rate of the sub-network",
												},
												"static_sub_network_bit_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of the sub-network",
												},
											},
										},
									},
									"subnet_breakdown": {
										Type: schema.TypeInt, Optional: true, Description: "additional layer of breakdown subnet",
									},
									"breakdown_subnet_threshold": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"breakdown_subnet_pkt_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
												},
												"breakdown_subnet_bit_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_rate': PPS; 'bit_rate': B(bits)PS;",
												},
											},
										},
									},
								},
							},
						},
						"sub_network_v6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet_ipv6_addr": {
										Type: schema.TypeString, Required: true, Description: "IPv6 Subnet/host, supported prefix range is from 56 to 64",
									},
									"host_anomaly_threshold": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"static_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
												},
												"static_rev_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
												},
												"static_bit_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
												},
												"static_rev_bit_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
												},
												"static_undiscovered_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Undiscovered packet rate of per host",
												},
												"static_flow_count_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Flow count of per host",
												},
												"static_syn_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "SYN packet rate of per host",
												},
												"static_fin_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "FIN packet rate of per host",
												},
												"static_rst_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "RST packet rate of per host",
												},
												"static_tcp_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "TCP packet rate of per host",
												},
												"static_udp_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "UDP packet rate of per host",
												},
												"static_icmp_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "ICMP packet rate of per host",
												},
												"static_undiscovered_host_pkt_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "packet rate of per undiscovered host",
												},
												"static_undiscovered_host_bit_rate_threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of per undiscovered host",
												},
											},
										},
									},
									"sub_network_anomaly_threshold": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"static_sub_network_pkt_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Packet rate of the sub-network",
												},
												"static_sub_network_bit_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Bit rate of the sub-network",
												},
											},
										},
									},
									"subnet_breakdown": {
										Type: schema.TypeInt, Optional: true, Description: "additional layer of breakdown subnet",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_rate': PPS; 'bit_rate': B(bits)PS;",
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
			"threshold_sensitivity": {
				Type: schema.TypeString, Optional: true, Description: "tune threshold ranges with levels LOW/MEDIUM/HIGH/OFF(default) or multiplier of threshold value (available options are LOW=5x/MEDIUM=3x/HIGH=1.5x/OFF=1x, or float value between 1.0-10.0)",
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
			"trustlist": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v4_class_list": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Class-list name",
						},
						"v6_class_list": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Class-list name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
func resourceDdosNetworkObjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosNetworkObjectEnableTopK(d []interface{}) []edpt.DdosNetworkObjectEnableTopK {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectEnableTopK, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectEnableTopK
		oi.TopkType = in["topk_type"].(string)
		oi.TopkDstNumRecords = in["topk_dst_num_records"].(int)
		oi.TopkSortKey = in["topk_sort_key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectHostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectHostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectHostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostPktRate = in["host_pkt_rate"].(int)
		ret.HostBitRate = in["host_bit_rate"].(int)
		ret.HostRevPktRate = in["host_rev_pkt_rate"].(int)
		ret.HostRevBitRate = in["host_rev_bit_rate"].(int)
		ret.HostUndiscoveredPktRate = in["host_undiscovered_pkt_rate"].(int)
		ret.HostFlowCount = in["host_flow_count"].(int)
		ret.HostSynRate = in["host_syn_rate"].(int)
		ret.HostFinRate = in["host_fin_rate"].(int)
		ret.HostRstRate = in["host_rst_rate"].(int)
		ret.HostTcpPktRate = in["host_tcp_pkt_rate"].(int)
		ret.HostUdpPktRate = in["host_udp_pkt_rate"].(int)
		ret.HostIcmpPktRate = in["host_icmp_pkt_rate"].(int)
		ret.HostUndiscoveredHostPktRate = in["host_undiscovered_host_pkt_rate"].(int)
		ret.HostUndiscoveredHostBitRate = in["host_undiscovered_host_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectIndicatorsToMonitor312(d []interface{}) edpt.DdosNetworkObjectIndicatorsToMonitor312 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIndicatorsToMonitor312
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.MonitorPktRate = in["monitor_pkt_rate"].(int)
		ret.MonitorBitRate = in["monitor_bit_rate"].(int)
		ret.MonitorRevPktRate = in["monitor_rev_pkt_rate"].(int)
		ret.MonitorRevBitRate = in["monitor_rev_bit_rate"].(int)
		ret.MonitorUndiscoveredPktRate = in["monitor_undiscovered_pkt_rate"].(int)
		ret.MonitorFlowCount = in["monitor_flow_count"].(int)
		ret.MonitorSynRate = in["monitor_syn_rate"].(int)
		ret.MonitorFinRate = in["monitor_fin_rate"].(int)
		ret.MonitorRstRate = in["monitor_rst_rate"].(int)
		ret.MonitorTcpPktRate = in["monitor_tcp_pkt_rate"].(int)
		ret.MonitorUdpPktRate = in["monitor_udp_pkt_rate"].(int)
		ret.MonitorIcmpPktRate = in["monitor_icmp_pkt_rate"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosNetworkObjectIpList(d []interface{}) []edpt.DdosNetworkObjectIpList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpList
		oi.SubnetIpAddr = in["subnet_ip_addr"].(string)
		oi.PrefixAnomalyThreshold = getObjectDdosNetworkObjectIpListPrefixAnomalyThreshold(in["prefix_anomaly_threshold"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosNetworkObjectIpListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectIpListPrefixAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpListPrefixAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpListPrefixAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixPktRate = in["prefix_pkt_rate"].(int)
		ret.PrefixBitRate = in["prefix_bit_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectIpListSamplingEnable(d []interface{}) []edpt.DdosNetworkObjectIpListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectIpv6List(d []interface{}) []edpt.DdosNetworkObjectIpv6List {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpv6List
		oi.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		oi.PrefixAnomalyThreshold = getObjectDdosNetworkObjectIpv6ListPrefixAnomalyThreshold(in["prefix_anomaly_threshold"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosNetworkObjectIpv6ListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectIpv6ListPrefixAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpv6ListPrefixAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpv6ListPrefixAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixPktRate = in["prefix_pkt_rate"].(int)
		ret.PrefixBitRate = in["prefix_bit_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectIpv6ListSamplingEnable(d []interface{}) []edpt.DdosNetworkObjectIpv6ListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpv6ListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpv6ListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectNetworkObjectAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectNetworkObjectAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectNetworkObjectAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkObjectPktRate = in["network_object_pkt_rate"].(int)
		ret.NetworkObjectBitRate = in["network_object_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectNotification313(d []interface{}) edpt.DdosNetworkObjectNotification313 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectNotification313
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Notification = getSliceDdosNetworkObjectNotificationNotification314(in["notification"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosNetworkObjectNotificationNotification314(d []interface{}) []edpt.DdosNetworkObjectNotificationNotification314 {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectNotificationNotification314, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectNotificationNotification314
		oi.NotificationTemplateName = in["notification_template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectRelativeAutoBreakDownThreshold(d []interface{}) edpt.DdosNetworkObjectRelativeAutoBreakDownThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectRelativeAutoBreakDownThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkPercentage = in["network_percentage"].(int)
		ret.Permil = in["permil"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectSamplingEnable(d []interface{}) []edpt.DdosNetworkObjectSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectServiceBreakDownThresholdLocal(d []interface{}) edpt.DdosNetworkObjectServiceBreakDownThresholdLocal {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectServiceBreakDownThresholdLocal
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SvcPercentage = in["svc_percentage"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSportAnomalyThreshold315(d []interface{}) edpt.DdosNetworkObjectSportAnomalyThreshold315 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSportAnomalyThreshold315
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketRate = getObjectDdosNetworkObjectSportAnomalyThresholdPacketRate316(in["packet_rate"].([]interface{}))
		ret.PacketRatePercentage = getObjectDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage317(in["packet_rate_percentage"].([]interface{}))
		ret.BitRate = getObjectDdosNetworkObjectSportAnomalyThresholdBitRate318(in["bit_rate"].([]interface{}))
		ret.BitRatePercentage = getObjectDdosNetworkObjectSportAnomalyThresholdBitRatePercentage319(in["bit_rate_percentage"].([]interface{}))
		ret.IpList = getSliceDdosNetworkObjectSportAnomalyThresholdIpList(in["ip_list"].([]interface{}))
		ret.Ipv6List = getSliceDdosNetworkObjectSportAnomalyThresholdIpv6List(in["ipv6_list"].([]interface{}))
		ret.SportList = getSliceDdosNetworkObjectSportAnomalyThresholdSportList(in["sport_list"].([]interface{}))
	}
	return ret
}

func getObjectDdosNetworkObjectSportAnomalyThresholdPacketRate316(d []interface{}) edpt.DdosNetworkObjectSportAnomalyThresholdPacketRate316 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdPacketRate316
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage317(d []interface{}) edpt.DdosNetworkObjectSportAnomalyThresholdPacketRatePercentage317 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdPacketRatePercentage317
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosNetworkObjectSportAnomalyThresholdBitRate318(d []interface{}) edpt.DdosNetworkObjectSportAnomalyThresholdBitRate318 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdBitRate318
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosNetworkObjectSportAnomalyThresholdBitRatePercentage319(d []interface{}) edpt.DdosNetworkObjectSportAnomalyThresholdBitRatePercentage319 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdBitRatePercentage319
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosNetworkObjectSportAnomalyThresholdIpList(d []interface{}) []edpt.DdosNetworkObjectSportAnomalyThresholdIpList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSportAnomalyThresholdIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSportAnomalyThresholdIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.PacketRateStr = in["packet_rate_str"].(string)
		oi.PacketRatePercentageStr = in["packet_rate_percentage_str"].(string)
		oi.BitRateStr = in["bit_rate_str"].(string)
		oi.BitRatePercentageStr = in["bit_rate_percentage_str"].(string)
		oi.PacketRate = in["packet_rate"].(int)
		oi.PacketRatePercentage = in["packet_rate_percentage"].(int)
		oi.BitRate = in["bit_rate"].(int)
		oi.BitRatePercentage = in["bit_rate_percentage"].(int)
		oi.SportNum = in["sport_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.IpSportPacketRateStr = in["ip_sport_packet_rate_str"].(string)
		oi.IpSportPacketRatePercentageStr = in["ip_sport_packet_rate_percentage_str"].(string)
		oi.IpSportBitRateStr = in["ip_sport_bit_rate_str"].(string)
		oi.IpSportBitRatePercentageStr = in["ip_sport_bit_rate_percentage_str"].(string)
		oi.IpSportPacketRate = in["ip_sport_packet_rate"].(int)
		oi.IpSportPacketRatePercentage = in["ip_sport_packet_rate_percentage"].(int)
		oi.IpSportBitRate = in["ip_sport_bit_rate"].(int)
		oi.IpSportBitRatePercentage = in["ip_sport_bit_rate_percentage"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectSportAnomalyThresholdIpv6List(d []interface{}) []edpt.DdosNetworkObjectSportAnomalyThresholdIpv6List {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSportAnomalyThresholdIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSportAnomalyThresholdIpv6List
		oi.IpAddr = in["ip_addr"].(string)
		oi.PacketRateStr = in["packet_rate_str"].(string)
		oi.PacketRatePercentageStr = in["packet_rate_percentage_str"].(string)
		oi.BitRateStr = in["bit_rate_str"].(string)
		oi.BitRatePercentageStr = in["bit_rate_percentage_str"].(string)
		oi.PacketRate = in["packet_rate"].(int)
		oi.PacketRatePercentage = in["packet_rate_percentage"].(int)
		oi.BitRate = in["bit_rate"].(int)
		oi.BitRatePercentage = in["bit_rate_percentage"].(int)
		oi.SportNum = in["sport_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.IpSportPacketRateStr = in["ip_sport_packet_rate_str"].(string)
		oi.IpSportPacketRatePercentageStr = in["ip_sport_packet_rate_percentage_str"].(string)
		oi.IpSportBitRateStr = in["ip_sport_bit_rate_str"].(string)
		oi.IpSportBitRatePercentageStr = in["ip_sport_bit_rate_percentage_str"].(string)
		oi.IpSportPacketRate = in["ip_sport_packet_rate"].(int)
		oi.IpSportPacketRatePercentage = in["ip_sport_packet_rate_percentage"].(int)
		oi.IpSportBitRate = in["ip_sport_bit_rate"].(int)
		oi.IpSportBitRatePercentage = in["ip_sport_bit_rate_percentage"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectSportAnomalyThresholdSportList(d []interface{}) []edpt.DdosNetworkObjectSportAnomalyThresholdSportList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSportAnomalyThresholdSportList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSportAnomalyThresholdSportList
		oi.SportNum = in["sport_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.PacketRateStr = in["packet_rate_str"].(string)
		oi.PacketRatePercentageStr = in["packet_rate_percentage_str"].(string)
		oi.BitRateStr = in["bit_rate_str"].(string)
		oi.BitRatePercentageStr = in["bit_rate_percentage_str"].(string)
		oi.PacketRate = in["packet_rate"].(int)
		oi.PacketRatePercentage = in["packet_rate_percentage"].(int)
		oi.BitRate = in["bit_rate"].(int)
		oi.BitRatePercentage = in["bit_rate_percentage"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectSportDiscoveryThreshold(d []interface{}) edpt.DdosNetworkObjectSportDiscoveryThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSportDiscoveryThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SportHeavyHitterPercentage = in["sport_heavy_hitter_percentage"].(int)
		ret.SportDiscoveryBitRatePercentage = in["sport_discovery_bit_rate_percentage"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectSportList(d []interface{}) []edpt.DdosNetworkObjectSportList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSportList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSportList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectStaticAutoBreakDownThreshold(d []interface{}) edpt.DdosNetworkObjectStaticAutoBreakDownThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectStaticAutoBreakDownThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkPktRate = in["network_pkt_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetwork320(d []interface{}) edpt.DdosNetworkObjectSubNetwork320 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetwork320
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SubNetworkV4List = getSliceDdosNetworkObjectSubNetworkSubNetworkV4List(in["sub_network_v4_list"].([]interface{}))
		ret.SubNetworkV6List = getSliceDdosNetworkObjectSubNetworkSubNetworkV6List(in["sub_network_v6_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosNetworkObjectSubNetworkSubNetworkV4List(d []interface{}) []edpt.DdosNetworkObjectSubNetworkSubNetworkV4List {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSubNetworkSubNetworkV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSubNetworkSubNetworkV4List
		oi.SubnetIpAddr = in["subnet_ip_addr"].(string)
		oi.HostAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV4ListHostAnomalyThreshold(in["host_anomaly_threshold"].([]interface{}))
		oi.SubNetworkAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV4ListSubNetworkAnomalyThreshold(in["sub_network_anomaly_threshold"].([]interface{}))
		oi.SubnetBreakdown = in["subnet_breakdown"].(int)
		oi.BreakdownSubnetThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV4ListBreakdownSubnetThreshold(in["breakdown_subnet_threshold"].([]interface{}))
		//omit uuid
		oi.SamplingEnable = getSliceDdosNetworkObjectSubNetworkSubNetworkV4ListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV4ListHostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListHostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListHostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticPktRateThreshold = in["static_pkt_rate_threshold"].(int)
		ret.StaticRevPktRateThreshold = in["static_rev_pkt_rate_threshold"].(int)
		ret.StaticBitRateThreshold = in["static_bit_rate_threshold"].(int)
		ret.StaticRevBitRateThreshold = in["static_rev_bit_rate_threshold"].(int)
		ret.StaticUndiscoveredPktRateThreshold = in["static_undiscovered_pkt_rate_threshold"].(int)
		ret.StaticFlowCountThreshold = in["static_flow_count_threshold"].(int)
		ret.StaticSynRateThreshold = in["static_syn_rate_threshold"].(int)
		ret.StaticFinRateThreshold = in["static_fin_rate_threshold"].(int)
		ret.StaticRstRateThreshold = in["static_rst_rate_threshold"].(int)
		ret.StaticTcpPktRateThreshold = in["static_tcp_pkt_rate_threshold"].(int)
		ret.StaticUdpPktRateThreshold = in["static_udp_pkt_rate_threshold"].(int)
		ret.StaticIcmpPktRateThreshold = in["static_icmp_pkt_rate_threshold"].(int)
		ret.StaticUndiscoveredHostPktRateThreshold = in["static_undiscovered_host_pkt_rate_threshold"].(int)
		ret.StaticUndiscoveredHostBitRateThreshold = in["static_undiscovered_host_bit_rate_threshold"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV4ListSubNetworkAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListSubNetworkAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListSubNetworkAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticSubNetworkPktRate = in["static_sub_network_pkt_rate"].(int)
		ret.StaticSubNetworkBitRate = in["static_sub_network_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV4ListBreakdownSubnetThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListBreakdownSubnetThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListBreakdownSubnetThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BreakdownSubnetPktRate = in["breakdown_subnet_pkt_rate"].(int)
		ret.BreakdownSubnetBitRate = in["breakdown_subnet_bit_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectSubNetworkSubNetworkV4ListSamplingEnable(d []interface{}) []edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSubNetworkSubNetworkV4ListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectSubNetworkSubNetworkV6List(d []interface{}) []edpt.DdosNetworkObjectSubNetworkSubNetworkV6List {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSubNetworkSubNetworkV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSubNetworkSubNetworkV6List
		oi.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		oi.HostAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV6ListHostAnomalyThreshold(in["host_anomaly_threshold"].([]interface{}))
		oi.SubNetworkAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV6ListSubNetworkAnomalyThreshold(in["sub_network_anomaly_threshold"].([]interface{}))
		oi.SubnetBreakdown = in["subnet_breakdown"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceDdosNetworkObjectSubNetworkSubNetworkV6ListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV6ListHostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV6ListHostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV6ListHostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticPktRateThreshold = in["static_pkt_rate_threshold"].(int)
		ret.StaticRevPktRateThreshold = in["static_rev_pkt_rate_threshold"].(int)
		ret.StaticBitRateThreshold = in["static_bit_rate_threshold"].(int)
		ret.StaticRevBitRateThreshold = in["static_rev_bit_rate_threshold"].(int)
		ret.StaticUndiscoveredPktRateThreshold = in["static_undiscovered_pkt_rate_threshold"].(int)
		ret.StaticFlowCountThreshold = in["static_flow_count_threshold"].(int)
		ret.StaticSynRateThreshold = in["static_syn_rate_threshold"].(int)
		ret.StaticFinRateThreshold = in["static_fin_rate_threshold"].(int)
		ret.StaticRstRateThreshold = in["static_rst_rate_threshold"].(int)
		ret.StaticTcpPktRateThreshold = in["static_tcp_pkt_rate_threshold"].(int)
		ret.StaticUdpPktRateThreshold = in["static_udp_pkt_rate_threshold"].(int)
		ret.StaticIcmpPktRateThreshold = in["static_icmp_pkt_rate_threshold"].(int)
		ret.StaticUndiscoveredHostPktRateThreshold = in["static_undiscovered_host_pkt_rate_threshold"].(int)
		ret.StaticUndiscoveredHostBitRateThreshold = in["static_undiscovered_host_bit_rate_threshold"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV6ListSubNetworkAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV6ListSubNetworkAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV6ListSubNetworkAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticSubNetworkPktRate = in["static_sub_network_pkt_rate"].(int)
		ret.StaticSubNetworkBitRate = in["static_sub_network_bit_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectSubNetworkSubNetworkV6ListSamplingEnable(d []interface{}) []edpt.DdosNetworkObjectSubNetworkSubNetworkV6ListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSubNetworkSubNetworkV6ListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSubNetworkSubNetworkV6ListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectTopkDestinations321(d []interface{}) edpt.DdosNetworkObjectTopkDestinations321 {

	var ret edpt.DdosNetworkObjectTopkDestinations321
	return ret
}

func getObjectDdosNetworkObjectTrustlist322(d []interface{}) edpt.DdosNetworkObjectTrustlist322 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTrustlist322
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V4ClassList = in["v4_class_list"].(string)
		ret.V6ClassList = in["v6_class_list"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosNetworkObject(d *schema.ResourceData) edpt.DdosNetworkObject {
	var ret edpt.DdosNetworkObject
	ret.Inst.AnomalyDetectionTrigger = d.Get("anomaly_detection_trigger").(string)
	ret.Inst.EnableTopK = getSliceDdosNetworkObjectEnableTopK(d.Get("enable_top_k").([]interface{}))
	ret.Inst.FloodingMultiplier = d.Get("flooding_multiplier").(int)
	ret.Inst.HistogramMode = d.Get("histogram_mode").(string)
	ret.Inst.HostAnomalyThreshold = getObjectDdosNetworkObjectHostAnomalyThreshold(d.Get("host_anomaly_threshold").([]interface{}))
	ret.Inst.HostSportDiscovery = d.Get("host_sport_discovery").(string)
	ret.Inst.IndicatorsToMonitor = getObjectDdosNetworkObjectIndicatorsToMonitor312(d.Get("indicators_to_monitor").([]interface{}))
	ret.Inst.IpList = getSliceDdosNetworkObjectIpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Ipv6List = getSliceDdosNetworkObjectIpv6List(d.Get("ipv6_list").([]interface{}))
	ret.Inst.NetworkObjectAnomalyThreshold = getObjectDdosNetworkObjectNetworkObjectAnomalyThreshold(d.Get("network_object_anomaly_threshold").([]interface{}))
	ret.Inst.NetworkObjectTemplate = d.Get("network_object_template").(string)
	ret.Inst.Notification = getObjectDdosNetworkObjectNotification313(d.Get("notification").([]interface{}))
	ret.Inst.ObjectName = d.Get("object_name").(string)
	ret.Inst.OperationalMode = d.Get("operational_mode").(string)
	ret.Inst.RelativeAutoBreakDownThreshold = getObjectDdosNetworkObjectRelativeAutoBreakDownThreshold(d.Get("relative_auto_break_down_threshold").([]interface{}))
	ret.Inst.SamplingEnable = getSliceDdosNetworkObjectSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceBreakDownThresholdLocal = getObjectDdosNetworkObjectServiceBreakDownThresholdLocal(d.Get("service_break_down_threshold_local").([]interface{}))
	ret.Inst.ServiceDiscovery = d.Get("service_discovery").(string)
	ret.Inst.SportAnomalyDetection = d.Get("sport_anomaly_detection").(string)
	ret.Inst.SportAnomalyThreshold = getObjectDdosNetworkObjectSportAnomalyThreshold315(d.Get("sport_anomaly_threshold").([]interface{}))
	ret.Inst.SportDiscoveryThreshold = getObjectDdosNetworkObjectSportDiscoveryThreshold(d.Get("sport_discovery_threshold").([]interface{}))
	ret.Inst.SportList = getSliceDdosNetworkObjectSportList(d.Get("sport_list").([]interface{}))
	ret.Inst.StaticAutoBreakDownThreshold = getObjectDdosNetworkObjectStaticAutoBreakDownThreshold(d.Get("static_auto_break_down_threshold").([]interface{}))
	ret.Inst.SubNetwork = getObjectDdosNetworkObjectSubNetwork320(d.Get("sub_network").([]interface{}))
	ret.Inst.ThresholdSensitivity = d.Get("threshold_sensitivity").(string)
	ret.Inst.TopkDestinations = getObjectDdosNetworkObjectTopkDestinations321(d.Get("topk_destinations").([]interface{}))
	ret.Inst.Trustlist = getObjectDdosNetworkObjectTrustlist322(d.Get("trustlist").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
