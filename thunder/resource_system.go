package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystem() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system`: Configure System Parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemCreate,
		UpdateContext: resourceSystemUpdate,
		ReadContext:   resourceSystemRead,
		DeleteContext: resourceSystemDelete,

		Schema: map[string]*schema.Schema{
			"add_cpu_core": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"core_index": {
							Type: schema.TypeInt, Optional: true, Description: "core index to be added (Specify core index)",
						},
					},
				},
			},
			"add_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
						},
					},
				},
			},
			"all_vlan_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bcast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "broadcast packets (per second limit)",
						},
						"ipmcast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "IP multicast packets (per second limit)",
						},
						"mcast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "multicast packets (per second limit)",
						},
						"unknown_ucast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "unknown unicast packets (per second limit)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"anomaly_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "log system anomalies",
			},
			"anomaly_log_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Anomaly log rate-limit per second, default 32",
			},
			"app_performance": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total-throughput-bits-per-sec': Total Throughput in bits/sec; 'l4-conns-per-sec': L4 Connections/sec; 'l7-conns-per-sec': L7 Connections/sec; 'l7-trans-per-sec': L7 Transactions/sec; 'ssl-conns-per-sec': SSL Connections/sec; 'ip-nat-conns-per-sec': IP NAT Connections/sec; 'total-new-conns-per-sec': Total New Connections Established/sec; 'total-curr-conns': Total Current Established Connections; 'l4-bandwidth': L4 Bandwidth in bits/sec; 'l7-bandwidth': L7 Bandwidth in bits/sec; 'serv-ssl-conns-per-sec': Server SSL Connections/sec; 'fw-conns-per-sec': FW Connections/sec; 'gifw-conns-per-sec': GiFW Connections/sec;",
									},
								},
							},
						},
					},
				},
			},
			"apps_global": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_session_on_established": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send TCP session creation log on completion of 3-way handshake",
						},
						"msl_time": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure maximum session life, default is 2 seconds (1-39 seconds, default is 2 seconds)",
						},
						"timer_wheel_walk_limit": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Set timer wheel walk limit (0-1024, 0 is unlimited, default is 100)",
						},
						"sessions_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set threshold for total sessions across the system (Enter threshold number)",
						},
						"cps_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set threshold for the total Connections Per Second across the system (Enter threshold number)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"asic_debug_dump": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable L2L3 ASIC traffic discard/drop events and Dump debug information",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"asic_mmu_fail_safe": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recovery_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "ASIC Fail-safe recovery threshold in Errors (Units of 1 Errors (default 2))",
						},
						"monitor_interval": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "ASIC Fail-safe monitoring intervals in Seconds (Units of 1 Seconds (default 60))",
						},
						"monitor_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Fail-safe software error monitoring and act on it",
						},
						"reboot_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable system reboot if system encounters mmu error",
						},
						"inject_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inject MMU SER/Parity errors",
						},
						"test_pattern_type": {
							Type: schema.TypeString, Optional: true, Default: "lcb", Description: "'all-zeros': Inject all bits 0s in a byte; 'all-ones': Inject all bits 1s in a byte; 'lcb': Logical checker board; 'inverse-lcb': Inverse Logical checker board;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"attack_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "log attack anomalies",
			},
			"bandwidth": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"warning_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 75, Description: "Adjust bandwidth warning threshold",
						},
						"critical_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 95, Description: "Adjust bandwidth critical threshold",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'input-bytes-per-sec': In Bytes per second; 'output-bytes-per-sec': Out Bytes per second; 'ppsl_drop_egr': Packet-Per-Sec Limit Drop at egress; 'ppsl_drop_ing': Packet-Per-Sec Limit Drop at ingress; 'ppsl_ignore_limit': Packet-Per-Sec Limit ignored packets count; 'licexpire_drop': License Expire Drop; 'bwl_drop': BW Limit Drop;",
									},
								},
							},
						},
					},
				},
			},
			"bfd": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip_checksum_error': IP packet checksum errors; 'udp_checksum_error': UDP packet checksum errors; 'session_not_found': Session not found; 'multihop_mismatch': Multihop session or packet mismatch; 'version_mismatch': BFD version mismatch; 'length_too_small': Packets too small; 'data_is_short': Packet data length too short; 'invalid_detect_mult': Invalid detect multiplier; 'invalid_multipoint': Invalid multipoint setting; 'invalid_my_disc': Invalid my descriptor; 'invalid_ttl': Invalid TTL; 'auth_length_invalid': Invalid authentication length; 'auth_mismatch': Authentication mismatch; 'auth_type_mismatch': Authentication type mismatch; 'auth_key_id_mismatch': Authentication key-id mismatch; 'auth_key_mismatch': Authentication key mismatch; 'auth_seqnum_invalid': Invalid authentication sequence number; 'auth_failed': Authentication failures; 'local_state_admin_down': Local admin down session state; 'dest_unreachable': Destination unreachable; 'no_ipv6_enable': No IPv6 enable; 'other_error': Other errors;",
									},
								},
							},
						},
					},
				},
			},
			"cl_threat_category": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"class_list_hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable class list hit count",
			},
			"cli_monitor_interval": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt, Optional: true, Description: "one interval is 300ms (0 = disable)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cm_update_file_name_ref": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_name": {
							Type: schema.TypeString, Optional: true, Description: "bind source name",
						},
						"dest_name": {
							Type: schema.TypeString, Optional: true, Description: "bind dest name",
						},
						"id1": {
							Type: schema.TypeInt, Optional: true, Description: "Specify unique Partition id",
						},
					},
				},
			},
			"config_mgmt": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_referenced_tagged_objects": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Allow deletion of referenced tagged objects. Default option.; 'disable': Block deletion of referenced tagged objects;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"pu_sync_detection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "Time interval (seconds) for detection. Default is 30 seconds.",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable pu-sync-detection feature; 'disable': Disable pu-sync-detection feature;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"mpm": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_workers": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set max workers count. Default is 1",
									},
									"min_idle_workers": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set minimum idle workers count. Default is 1",
									},
									"start_workers": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set starting workers count. Default is 1",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"notification": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"period": {
										Type: schema.TypeInt, Optional: true, Default: 15, Description: "Time interval (seconds) for kafka notification. Default is 15 seconds.",
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
			"control_cpu": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"core": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cosq_show": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cosq_stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"counter_lib_accounting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cpu_hyper_thread": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPU Hyperthreading",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU Hyperthreading",
						},
					},
				},
			},
			"cpu_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cpu_load_sharing": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU load sharing in overload situations",
						},
						"packets_per_second": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"min": {
										Type: schema.TypeInt, Optional: true, Default: 100000, Description: "Minimum packets-per-second threshold (per CPU) before redistribution will take effect (Minimum packets-per-second threshold (per CPU) before redistribution will take effect (default: 100000))",
									},
								},
							},
						},
						"cpu_usage": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"low": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "CPU usage threshold (percentage) that will restore the normal packet distribution (default: 60)",
									},
									"high": {
										Type: schema.TypeInt, Optional: true, Default: 75, Description: "CPU usage threshold (percentage) that will trigger the redistribution (default: 75)",
									},
								},
							},
						},
						"allow_l7_sessions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow L7 sessions forward to home cpu",
						},
						"tcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new TCP sessions",
						},
						"udp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new UDP sessions",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new non TCP/UDP IP sessions",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cpu_map": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cpu_packet_prio_support": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPU packet prioritization Support",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU packet prioritization Support",
						},
					},
				},
			},
			"data_cpu": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ddos_attack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "System DDoS Attack",
			},
			"ddos_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "log DDoS attack anomalies",
			},
			"default_mtu": {
				Type: schema.TypeInt, Optional: true, Description: "Set all interfaces default mtu (Interface MTU, default 1 (System jumbo needs to be enabled))",
			},
			"del_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
						},
					},
				},
			},
			"delete_cpu_core": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"core_index": {
							Type: schema.TypeInt, Optional: true, Description: "core index to be deleted (Specify core index)",
						},
					},
				},
			},
			"disable_ssh_agent_forwarding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable sshd agent forwarding",
			},
			"dns": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'slb_req': No. of requests; 'slb_resp': No. of responses; 'slb_no_resp': No. of requests with no response; 'slb_req_rexmit': No. of requests retransmit; 'slb_resp_no_match': No. of requests and responses with no match; 'slb_no_resource': No. of resource failures; 'nat_req': (NAT) No. of requests; 'nat_resp': (NAT) No. of responses; 'nat_no_resp': (NAT) No. of resource failures; 'nat_req_rexmit': (NAT) No. of request retransmits; 'nat_resp_no_match': (NAT) No. of requests with no response; 'nat_no_resource': (NAT) No. of resource failures; 'nat_xid_reused': (NAT) No. of requests reusing a transaction id; 'filter_type_drop': Total Query Type Drop; 'filter_class_drop': Total Query Class Drop; 'filter_type_any_drop': Total Query ANY Type Drop; 'slb_dns_client_ssl_succ': No. of client ssl success; 'slb_dns_server_ssl_succ': No. of server ssl success; 'slb_dns_udp_conn': No. of backend udp connections; 'slb_dns_udp_conn_succ': No. of backend udp conn established; 'slb_dns_padding_to_server_removed': slb_dns_padding_to_server_removed; 'slb_dns_padding_to_client_added': slb_dns_padding_to_client_added; 'slb_dns_edns_subnet_to_server_removed': slb_dns_edns_subnet_to_server_removed; 'slb_dns_udp_retransmit': slb_dns_udp_retransmit; 'slb_dns_udp_retransmit_fail': slb_dns_udp_retransmit_fail; 'rpz_action_drop': RPZ Action Drop; 'rpz_action_pass_thru': RPZ Action Pass Through; 'rpz_action_tcp_only': RPZ Action TCP Only; 'rpz_action_nxdomain': RPZ Action NXDOMAIN; 'rpz_action_nodata': RPZ Action NODATA; 'rpz_action_local_data': RPZ Action Local Data; 'slb_drop': DNS requests drop; 'nat_slb_drop': (NAT)DNS requests drop; 'invalid_q_len_to_udp': invalid query length to conver to UDP; 'slb_dns_edns_ecs_received': Number of ecs from client received; 'slb_dns_edns_ecs_inserted': Number of ecs inserted;",
									},
								},
							},
						},
						"recursive_nameserver": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"follow_shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the configured name servers of shared partition",
									},
									"server_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv4 server address",
												},
												"v4_desc": {
													Type: schema.TypeString, Optional: true, Description: "Description for this ipv4 address",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv6 server address",
												},
												"v6_desc": {
													Type: schema.TypeString, Optional: true, Description: "Description for this ipv6 address",
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
					},
				},
			},
			"dns_cache": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_q': Total query; 'total_r': Total server response; 'hit': Total cache hit; 'bad_q': Query not passed; 'encode_q': Query encoded; 'multiple_q': Query with multiple questions; 'oversize_q': Query exceed cache size; 'bad_r': Response not passed; 'oversize_r': Response exceed cache size; 'encode_r': Response encoded; 'multiple_r': Response with multiple questions; 'answer_r': Response with multiple answers; 'ttl_r': Response with short TTL; 'ageout': Total aged out; 'bad_answer': Bad Answer; 'ageout_weight': Total aged for lower weight; 'total_log': Total stats log sent; 'total_alloc': Total allocated; 'total_freed': Total freed; 'current_allocate': Current allocate; 'current_data_allocate': Current data allocate; 'resolver_queue_full': Resolver task queue full; 'truncated_r': Response with Truncation bit set; 'qps': Cache Queries-per-second; 'hit_rate_per_sec': Cache hit rate per second;",
									},
								},
							},
						},
					},
				},
			},
			"domain_list_hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable class list hit count",
			},
			"domain_list_info": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"domain_list_settings": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"polling_interval": {
							Type: schema.TypeString, Optional: true, Default: "10-second", Description: "'1-second': Set interval to 1 second; '5-second': Set interval to 5 seconds; '10-second': Set interval to 10 seconds (Default);",
						},
						"concurrent_task": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Configure max concurrent AXFR task (Default 6)",
						},
						"domain_list_per_group": {
							Type: schema.TypeString, Optional: true, Default: "16", Description: "'16': Allow 16 domain-list per group (Default); '32': Allow 32 domain-list per group;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dpdk_stats": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'pkt-drop': Total packet drop; 'pkt-lnk-down-drop': Total packets link down drop; 'err-pkt-drop': Total error packet drop; 'rx-err': Total RX packet error; 'tx-err': Total TX packet error; 'tx-drop': Total TX packet drop; 'rx-len-err': Total RX packet length error; 'rx-over-err': Total RX packet over error; 'rx-crc-err': Total RX packet CRC error; 'rx-frame-err': Total RX packet frame error; 'rx-no-buff-err': Total RX packet no buffer error; 'rx-miss-err': Total RX packet miss error; 'tx-abort-err': Total TX packet abort error; 'tx-carrier-err': Total TX packert carrier error; 'tx-fifo-err': Total TX packet fifo error; 'tx-hbeat-err': Total TX packet HBEAT error; 'tx-windows-err': Total TX windows error; 'rx-long-len-err': Total RX packet long length error; 'rx-short-len-err': Total RX packet short length error; 'rx-align-err': Total RX packet align error; 'rx-csum-offload-err': Total Rx packet check-sum offload error; 'io-rx-que-drop': Total IO core Rx queue drop; 'io-tx-que-drop': Total IO core TX queue drop; 'io-ring-drop': Total IO core ring drop; 'w-tx-que-drop': Total worker core queue drop; 'w-link-down-drop': Total worker core link down drop; 'w-ring-drop': Total worker core ring drop;",
									},
								},
							},
						},
					},
				},
			},
			"drop_linux_closed_port_syn": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable (default); 'disable': disable;",
			},
			"dynamic_service_dns_socket_pool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable socket pool for dynamic-service DNS",
			},
			"enable_disk_encryption": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher": {
							Type: schema.TypeString, Optional: true, Default: "aes", Description: "'aes': cipher aes; 'serpent': cipher serpent; 'twofish': cipher twofish;",
						},
						"passphrase": {
							Type: schema.TypeString, Optional: true, Description: "Enter phassphrase in plain text format",
						},
						"passphrase_base64": {
							Type: schema.TypeString, Optional: true, Description: "Enter phassphrase in base64 format",
						},
					},
				},
			},
			"enable_password": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"follow_password_policy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable-password will follow password policy complexity",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"environment": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"even_port_hash_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable even src/dest port number hash",
			},
			"ext_only_logging": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable external only logging for packet driven DDOS logs",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"forced_group_speed_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eth01_to_04": {
							Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet  1 ~  4",
						},
						"eth05_to_08": {
							Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet  5 ~  8",
						},
						"eth09_to_12": {
							Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet  9 ~ 12",
						},
						"eth13_to_16": {
							Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet 13 ~ 16",
						},
						"eth17_to_20": {
							Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet 17 ~ 20",
						},
						"eth21_to_24": {
							Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet 21 ~ 24",
						},
						"speed": {
							Type: schema.TypeString, Optional: true, Default: "10G", Description: "'1G': Speed 1G; '10G': Speed 10G (default); '25G': Speed 25G;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"fpga_core_crc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable FPGA Core CRC error monitoring and act on it",
						},
						"reboot_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system reboot if system encounters FPGA Core CRC error",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"fpga_drop": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'mrx-drop': Total MRX Drop; 'hrx-drop': Total HRX Drop; 'siz-drop': Total Size Drop; 'fcs-drop': Total FCS Drop; 'land-drop': Total LAND Attack Drop; 'empty-frag-drop': Total Empty frag Drop; 'mic-frag-drop': Total Micro frag Drop; 'ipv4-opt-drop': Total IPv4 opt Drop; 'ipv4-frag': Total IP frag Drop; 'bad-ip-hdr-len': Total Bad IP hdr len Drop; 'bad-ip-flags-drop': Total Bad IP Flags Drop; 'bad-ip-ttl-drop': Total Bad IP TTL Drop; 'no-ip-payload-drop': Total No IP Payload Drop; 'oversize-ip-payload': Total Oversize IP PL Drop; 'bad-ip-payload-len': Total Bad IP PL len Drop; 'bad-ip-frag-offset': Total Bad IP frag off Drop; 'bad-ip-chksum-drop': Total Bad IP csum Drop; 'icmp-pod-drop': Total ICMP POD Drop; 'tcp-bad-urg-offet': Total TCP bad urg off Drop; 'tcp-short-hdr': Total TCP short hdr Drop; 'tcp-bad-ip-len': Total TCP Bad IP Len Drop; 'tcp-null-flags': Total TCP null flags Drop; 'tcp-null-scan': Total TCP null scan Drop; 'tcp-fin-sin': Total TCP SYN&FIN Drop; 'tcp-xmas-flags': Total TCP XMAS FLAGS Drop; 'tcp-xmas-scan': Total TCP XMAS scan Drop; 'tcp-syn-frag': Total TCP SYN frag Drop; 'tcp-frag-hdr': Total TCP frag header Drop; 'tcp-bad-chksum': Total TCP bad csum Drop; 'udp-short-hdr': Total UDP short hdr Drop; 'udp-bad-ip-len': Total UDP bad leng Drop; 'udp-kb-frags': Total UDP KB frag Drop; 'udp-port-lb': Total UDP port LB Drop; 'udp-bad-chksum': Total UDP bad csum Drop; 'runt-ip-hdr': Total Runt IP hdr Drop; 'runt-tcpudp-hdr': Total Runt TCPUDP hdr Drop; 'tun-mismatch': Total Tun mismatch Drop; 'qdr-drop': Total QDR Drop;",
									},
								},
							},
						},
					},
				},
			},
			"fw": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_mempool": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable application memory pool",
						},
						"application_flow": {
							Type: schema.TypeInt, Optional: true, Description: "Number of flows",
						},
						"basic_dpi_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable basic dpi",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"geo_db_hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Geolocation database hit count",
			},
			"geo_location": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_iana": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Load built-in IANA Database",
						},
						"geo_location_iana_system": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in IANA Database",
						},
						"geo_location_geolite2_asn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-ASN database. Database available from http://www.maxmind.com",
						},
						"geolite2_asn_include_ipv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
						},
						"geo_location_geolite2_city": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-City database. Database available from http://www.maxmind.com",
						},
						"geolite2_city_include_ipv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
						},
						"geo_location_geolite2_country": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-Country database. Database available from http://www.maxmind.com",
						},
						"geolite2_country_include_ipv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
						},
						"geoloc_load_file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_location_load_filename": {
										Type: schema.TypeString, Optional: true, Description: "Specify file to be loaded",
									},
									"geo_location_load_file_include_ipv6": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
									},
									"template_name": {
										Type: schema.TypeString, Optional: true, Description: "Use CSV template to load this file",
									},
									"geo_location_load_temp_include_ipv6": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_locn_obj_name": {
										Type: schema.TypeString, Required: true, Description: "Specify geo-location name, section range is (1-15)",
									},
									"geo_locn_multiple_addresses": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"first_ip_address": {
													Type: schema.TypeString, Optional: true, Description: "Specify IP information (Specify IP address)",
												},
												"geol_ipv4_mask": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv4 mask",
												},
												"ip_addr2": {
													Type: schema.TypeString, Optional: true, Description: "Specify IP address range",
												},
												"first_ipv6_address": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
												},
												"geol_ipv6_mask": {
													Type: schema.TypeInt, Optional: true, Description: "Specify IPv6 mask",
												},
												"ipv6_addr2": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address range",
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
			"geoloc": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'place-holder': place-holder;",
									},
								},
							},
						},
					},
				},
			},
			"geoloc_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify name of Geolocation list",
						},
						"shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sharing with other partitions",
						},
						"include_geoloc_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include_geoloc_name_val": {
										Type: schema.TypeString, Optional: true, Description: "Geolocation name to add",
									},
								},
							},
						},
						"exclude_geoloc_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exclude_geoloc_name_val": {
										Type: schema.TypeString, Optional: true, Description: "Geolocation name to exclude",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-count': hit-count; 'total-geoloc': total-geoloc; 'total-active': total-active;",
									},
								},
							},
						},
					},
				},
			},
			"geoloc_name_helper": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'place-holder': place-holder;",
									},
								},
							},
						},
					},
				},
			},
			"geolocation_file": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"error_info": {
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
			"glid": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"glid_id": {
							Type: schema.TypeString, Optional: true, Description: "Apply limits to the whole system",
						},
						"non_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply global limit ID to the whole system at per data cpu level (default disabled)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"guest_file": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"gui_image_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"hardware": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"hardware_accelerate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_forwarding": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure session-forwarding in Hardware (default:off)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-counts': Total packets hit counts; 'hit-index': HW Fwd hit index; 'ipv4-forward-counts': Total IPv4 hardware forwarded packets; 'ipv6-forward-counts': Total IPv6 hardware forwarded packets; 'hw-fwd-module-status': Hardware forwarder status flags; 'hw-fwd-prog-reqs': Hardware forward programming requests; 'hw-fwd-prog-errors': Hardware forward programming Errors; 'hw-fwd-flow-singlebit-errors': Hardware forward singlebit Errors; 'hw-fwd-flow-tag-mismatch': Hardware forward tag mismatch errors; 'hw-fwd-flow-seq-mismatch': Hardware forward sequence mismatch errors; 'hw-fwd-ageout-drop-count': Hardware forward ageout drop count; 'hw-fwd-invalidation-drop': Hardware forward invalid drop count; 'hw-fwd-flow-hit-index': Hardware forward flow hit index; 'hw-fwd-flow-reason-flags': Hardware forward flow reason flags; 'hw-fwd-flow-drop-count': Hardware forward flow drop count; 'hw-fwd-flow-error-count': Hardware forward flow error count; 'hw-fwd-flow-unalign-count': Hardware forward flow unalign count; 'hw-fwd-flow-underflow-count': Hardware forward flow underflow count; 'hw-fwd-flow-tx-full-drop': Hardware forward flow tx full drop count; 'hw-fwd-flow-qdr-full-drop': Hardware forward flow qdr full drop count; 'hw-fwd-phyport-mismatch-drop': Hardware forward phyport mismatch count; 'hw-fwd-vlanid-mismatch-drop': Hardware forward vlanid mismatch count; 'hw-fwd-vmid-drop': Hardware forward vmid mismatch count; 'hw-fwd-protocol-mismatch-drop': Hardware forward protocol mismatch count; 'hw-fwd-avail-ipv4-entry': Hardware forward available ipv4 entries count; 'hw-fwd-avail-ipv6-entry': Hardware forward available ipv6 entries count; 'hw-fwd-rate-drop-count': Hardware forward rate drop count; 'hw-fwd-normal-ageout-rcvd': Hardware forward normal ageout received count; 'hw-fwd-tcp-fin-ageout-rcvd': Hardware forward tcp FIN ageout received count; 'hw-fwd-tcp-rst-ageout-rcvd': Hardware forward tcp RST ageout received count; 'hw-fwd-lookup-fail-rcvd': Hardware forward entry lookup fail count; 'hw-fwd-stats-update-rcvd': Hardware forward entry stats update count; 'hw-fwd-flow-sflow-count': hardware forward rate drop count;",
									},
								},
							},
						},
						"slb": {
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
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'entry-create': Hardware Entries Created; 'entry-create-failure': Hardware Entries Created; 'entry-create-fail-server-down': Hardware Entries Created; 'entry-create-fail-max-entry': Hardware Entries Created; 'entry-free': Hardware Entries Freed; 'entry-free-opp-entry': Hardware Entries Free due to opposite tuple entry ageout event; 'entry-free-no-hw-prog': Hardware Entry Free no hw prog; 'entry-free-no-conn': Hardware Entry Free no matched conn; 'entry-free-no-sw-entry': Hardware Entry Free no software entry; 'entry-counter': Hardware Entry Count; 'entry-age-out': Hardware Entries Aged Out; 'entry-age-out-idle': Hardware Entries Aged Out Idle; 'entry-age-out-tcp-fin': Hardware Entries Aged Out TCP FIN; 'entry-age-out-tcp-rst': Hardware Entries Aged Out TCP RST; 'entry-age-out-invalid-dst': Hardware Entries Aged Out invalid dst; 'entry-force-hw-invalidate': Hardware Entries Force HW Invalidate; 'entry-invalidate-server-down': Hardware Entries Invalidate due to server down; 'tcam-create': TCAM Entries Created; 'tcam-free': TCAM Entries Freed; 'tcam-counter': TCAM Entry Count;",
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
			"health_check_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l2hm_hc_name": {
							Type: schema.TypeString, Required: true, Description: "Monitor Name",
						},
						"method_l2bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Method is l2bfd",
						},
						"l2bfd_tx_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Transmit interval between BFD packets",
						},
						"l2bfd_rx_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum receive interval capability (Milliseconds (default: 800))",
						},
						"l2bfd_multiplier": {
							Type: schema.TypeInt, Optional: true, Description: "Multiplier value used to compute holddown (value used to multiply the interval (default: 4))",
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
			"high_memory_l4_session": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable high memory l4 session support",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"hrxq_status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"hw_blocking_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system hardware blocking (default disabled)",
			},
			"icmp": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'num': Total number; 'inmsgs': In Messages; 'inerrors': In Errors; 'indestunreachs': In Destination Unreachable; 'intimeexcds': In TTL Exceeds; 'inparmprobs': In Parameter Problem; 'insrcquenchs': In Source Quench Error; 'inredirects': In Redirects; 'inechos': In Echo requests; 'inechoreps': In Echo replies; 'intimestamps': In Timestamp; 'intimestampreps': In Timestamp Rep; 'inaddrmasks': In Address Masks; 'inaddrmaskreps': In Address Mask Rep; 'outmsgs': Out Message; 'outerrors': Out Errors; 'outdestunreachs': Out Destination Unreachable; 'outtimeexcds': Out TTL Exceeds; 'outparmprobs': Out Parameter Problem; 'outsrcquenchs': Out Source Quench Error; 'outredirects': Out Redirects; 'outechos': Out Echo Requests; 'outechoreps': Out Echo Replies; 'outtimestamps': Out Time Stamp; 'outtimestampreps': Out Time Stamp Rep; 'outaddrmasks': Out Address Mask; 'outaddrmaskreps': Out Address Mask Rep;",
									},
								},
							},
						},
					},
				},
			},
			"icmp_rate": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'over_limit_drop': Over limit drops; 'limit_intf_drop': Interfaces rate limit drops; 'limit_vserver_drop': Virtual Server rate limit drops; 'limit_total_drop': Total rate limit drops; 'lockup_time_left': Lockup time left; 'curr_rate': Current rate; 'v6_over_limit_drop': Over limit drops (v6); 'v6_limit_intf_drop': Interfaces rate limit drops (v6); 'v6_limit_vserver_drop': Virtual Server rate limit drops (v6); 'v6_limit_total_drop': Total rate limit drops (v6); 'v6_lockup_time_left': Lockup time left (v6); 'v6_curr_rate': Current rate (v6);",
									},
								},
							},
						},
					},
				},
			},
			"icmp6": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'in_msgs': In messages; 'in_errors': In Errors; 'in_dest_un_reach': In Destunation Unreachable; 'in_pkt_too_big': In Packet too big; 'in_time_exceeds': In TTL Exceeds; 'in_param_prob': In Parameter Problem; 'in_echoes': In Echo requests; 'in_exho_reply': In Echo replies; 'in_grp_mem_query': In Group member query; 'in_grp_mem_resp': In Group member reply; 'in_grp_mem_reduction': In Group member reduction; 'in_router_sol': In Router solicitation; 'in_ra': In Router advertisement; 'in_ns': In neighbor solicitation; 'in_na': In neighbor advertisement; 'in_redirect': In Redirects; 'out_msg': Out Messages; 'out_dst_un_reach': Out Destination Unreachable; 'out_pkt_too_big': Out Packet too big; 'out_time_exceeds': Out TTL Exceeds; 'out_param_prob': Out Parameter Problem; 'out_echo_req': Out Echo requests; 'out_echo_replies': Out Echo replies; 'out_rs': Out Router solicitation; 'out_ra': Out Router advertisement; 'out_ns': Out neighbor solicitation; 'out_na': Out neighbor advertisement; 'out_redirects': Out Redirects; 'out_mem_resp': Out Group member reply; 'out_mem_reductions': Out Group member reduction; 'err_rs': Error Router solicitation; 'err_ra': Error Router advertisement; 'err_ns': Error Neighbor solicitation; 'err_na': Error Neighbor advertisement; 'err_redirects': Error Redirects; 'err_echoes': Error Echo requests; 'err_echo_replies': Error Echo replies;",
									},
								},
							},
						},
					},
				},
			},
			"inuse_cpu_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"inuse_port_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"io_cpu": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_cores": {
							Type: schema.TypeInt, Optional: true, Description: "max number of IO cores (Specify number of cores)",
						},
					},
				},
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_unreachable_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmp unreachable messages",
						},
						"icmp_redirect_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmp redirect messages",
						},
						"rpf_check_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable reverse path filter (strict mode)",
						},
						"source_route_pkt_drop_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv4 source routed packet drop",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ip_dns_cache": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ip_stats": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'inreceives': Incoming packets received; 'inhdrerrors': Incoming packet header errors; 'intoobigerrors': Incoming packet too big errors; 'innoroutes': Incoming no route packet drops; 'inaddrerrors': Incoming packet address errors; 'inunknownprotos': Incoming unkown protocol packet drops; 'intruncatedpkts': Incoming truncated packets; 'indiscards': Incoming packets discarded; 'indelivers': Incoming packets delivered; 'outforwdatagrams': Outgoing forwarded datagrams; 'outrequests': Outgoing packets; 'outdiscards': Outgoing packets discarded; 'outnoroutes': Outgoing no route packet drops; 'reasmtimeout': Reassembly timed out packet drops; 'reasmreqds': Incoming reassembly requests; 'reasmoks': Incoming reassembled packets; 'reasmfails': Incoming reassembly requests failed; 'fragoks': Outgoing packets fragmented; 'fragfails': Outgoing packets fragmentation failed; 'fragcreates': Outgoing fragmented packets; 'inmcastpkts': Incoming multicast packets; 'outmcastpkts': Outgoing multicast packets;",
									},
								},
							},
						},
					},
				},
			},
			"ip_threat_list": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_hit_count_in_sw': Packet Hit Count in SW; 'packet_hit_count_in_spe': Packet Hit Count in SPE; 'entries_added_in_sw': Entries Added in SW; 'entries_removed_from_sw': Entries Removed from SW; 'entries_added_in_spe': Entries Added in SPE; 'entries_removed_from_spe': Entries Removed from SPE; 'error_out_of_memory': Out of memory Error; 'error_out_of_spe_entries': Out of SPE Entries Error;",
									},
								},
							},
						},
						"ipv4_source_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
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
						"ipv4_dest_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
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
						"ipv6_source_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
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
						"ipv6_dest_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
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
						"ipv4_internet_host_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"white_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind exception-list (class-list name)",
									},
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
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
						"ipv6_internet_host_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"white_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind exception-list (class-list name)",
									},
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
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
					},
				},
			},
			"ip6_stats": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'inreceives': Incoming packets received; 'inhdrerrors': Incoming packet header errors; 'intoobigerrors': Incoming packet too big errors; 'innoroutes': Incoming no route packet drops; 'inaddrerrors': Incoming packet address errors; 'inunknownprotos': Incoming unkown protocol packet drops; 'intruncatedpkts': Incoming truncated packets; 'indiscards': Incoming packets discarded; 'indelivers': Incoming packets delivered; 'outforwdatagrams': Outgoing forwarded datagrams; 'outrequests': Outgoing packets; 'outdiscards': Outgoing packets discarded; 'outnoroutes': Outgoing no route packet drops; 'reasmtimeout': Reassembly timed out packet drops; 'reasmreqds': Incoming reassembly requests; 'reasmoks': Incoming reassembled packets; 'reasmfails': Incoming reassembly requests failed; 'fragoks': Outgoing packets fragmented; 'fragfails': Outgoing packets fragmentation failed; 'fragcreates': Outgoing fragmented packets; 'inmcastpkts': Incoming multicast packets; 'outmcastpkts': Outgoing multicast packets;",
									},
								},
							},
						},
					},
				},
			},
			"ipmi": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset IPMI Controller",
						},
						"ip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"ipv4_netmask": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
									},
									"default_gateway": {
										Type: schema.TypeString, Optional: true, Description: "Default gateway address",
									},
								},
							},
						},
						"ipsrc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dhcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP addr obtained by BMC running DHCP",
									},
									"static": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually configured static IP address",
									},
								},
							},
						},
						"user": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"add": {
										Type: schema.TypeString, Optional: true, Description: "Add a new IPMI user (IPMI User Name)",
									},
									"password": {
										Type: schema.TypeString, Optional: true, Description: "Password",
									},
									"administrator": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Full control",
									},
									"callback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lowest privilege level",
									},
									"operator": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most BMC commands are allowed",
									},
									"user": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only 'benign' commands are allowed",
									},
									"disable": {
										Type: schema.TypeString, Optional: true, Description: "Disable an existing IPMI user (IPMI User Name)",
									},
									"privilege": {
										Type: schema.TypeString, Optional: true, Description: "Change an existing IPMI user privilege (IPMI User Name)",
									},
									"setname": {
										Type: schema.TypeString, Optional: true, Description: "Change User Name (Current IPMI User Name)",
									},
									"newname": {
										Type: schema.TypeString, Optional: true, Description: "New IPMI User Name",
									},
									"setpass": {
										Type: schema.TypeString, Optional: true, Description: "Change Password (IPMI User Name)",
									},
									"newpass": {
										Type: schema.TypeString, Optional: true, Description: "New Password",
									},
								},
							},
						},
						"tool": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cmd": {
										Type: schema.TypeString, Optional: true, Description: "Command to execute in double quotes",
									},
								},
							},
						},
					},
				},
			},
			"ipmi_service": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable IPMI on platform",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipsec": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_round_robin": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet round robin for IPsec packets",
						},
						"crypto_core": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Crypto cores assigned for IPsec processing",
						},
						"crypto_mem": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Crypto memory percentage assigned for IPsec processing (rounded to increments of 10)",
						},
						"qat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "HW assisted QAT SSL module",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"fpga_decrypt": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable FPGA decryption offload; 'disable': Disable FPGA decryption offload;",
									},
								},
							},
						},
					},
				},
			},
			"ipv6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmpv6_unreachable_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmpv6 unreachable messages",
						},
						"icmpv6_redirect_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable icmpv6 redirect messages",
						},
						"rpf_check_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable reverse path filter (strict mode)",
						},
						"source_route_pkt_drop_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 source routed packet drop",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv6_prefix_length": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Length of IPv6 prefix used to determine the user-group and the PU, by default 128",
			},
			"job_offload": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'jobs': Current Jobs; 'submit': Jobs Submitted; 'receive': Jobs Received; 'execute': Jobs Executed; 'snt_home': Jobs Sent Back Home; 'rcv_home': Jobs Received Home; 'complete': Jobs Completed; 'fail_submit': Jobs Failed to Submit; 'q_no_space': No More Space in Q; 'fail_execute': Failed to Execute Job; 'fail_complete': Failed to Complete Job;",
									},
								},
							},
						},
					},
				},
			},
			"link_capability": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable link capabilities",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"link_monitor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Link Monitoring",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Link Monitoring",
						},
					},
				},
			},
			"lro": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Large Receive Offload",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Large Receive Offload",
						},
					},
				},
			},
			"management_interface_mode": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dedicated": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set management interface in dedicated mode",
						},
						"non_dedicated": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set management interface in non-dedicated mode",
						},
					},
				},
			},
			"memory": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'usage-percentage': Memory Usage percentage;",
									},
								},
							},
						},
					},
				},
			},
			"memory_block_debug": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assert_block": {
							Type: schema.TypeInt, Optional: true, Default: 65536, Description: "Over size block allocation (Assert memory block over size (default: 65536))",
						},
						"pktdump_block": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable pktdump Oversize block request packet",
						},
						"first_blk": {
							Type: schema.TypeInt, Optional: true, Default: 8192, Description: "First memory block ascending order (default: 8192) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
						},
						"second_blk": {
							Type: schema.TypeInt, Optional: true, Default: 16384, Description: "Second memory block (default: 16384) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
						},
						"third_blk": {
							Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Third memory block (default: 32768) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
						},
						"fourth_blk": {
							Type: schema.TypeInt, Optional: true, Default: 65536, Description: "Fourth memory block (default: 65536) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mfa_auth": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type: schema.TypeString, Optional: true, Description: "Username for MFA validation",
						},
						"second_factor": {
							Type: schema.TypeString, Optional: true, Description: "Input second factor paramter",
						},
					},
				},
			},
			"mfa_cert_store": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_host": {
							Type: schema.TypeString, Optional: true, Description: "Configure certificate store host",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "'tftp': Use tftp for connection; 'ftp': Use ftp for connection; 'scp': Use scp for connection; 'http': Use http for connection; 'https': Use https for connection; 'sftp': Use sftp for connection;",
						},
						"cert_store_path": {
							Type: schema.TypeString, Optional: true, Description: "Configure certificate store path",
						},
						"username": {
							Type: schema.TypeString, Optional: true, Description: "Certificate store host username",
						},
						"passwd_string": {
							Type: schema.TypeString, Optional: true, Description: "Certificate store host password",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mfa_management": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 2FA for management plane",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mfa_validation_type": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_cert": {
							Type: schema.TypeString, Optional: true, Description: "Configure CA Certificate",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mgmt_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
						},
						"mac_address": {
							Type: schema.TypeString, Optional: true, Description: "mac-address to be configured as mgmt port",
						},
						"pci_address": {
							Type: schema.TypeString, Optional: true, Description: "pci-address to be configured as mgmt port",
						},
					},
				},
			},
			"modify_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
						},
						"port_number": {
							Type: schema.TypeInt, Optional: true, Description: "port number to be configured (Specify port number)",
						},
					},
				},
			},
			"module_ctrl_cpu": {
				Type: schema.TypeString, Optional: true, Description: "'high': high cpu usage; 'low': low cpu usage; 'medium': medium cpu usage;",
			},
			"mon_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id1": {
										Type: schema.TypeInt, Required: true, Description: "Monitor template ID Number",
									},
									"clear_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sessions": {
													Type: schema.TypeString, Optional: true, Description: "'all': Clear all sessions; 'sequence': Sequence number;",
												},
												"clear_all_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
												},
												"clear_all_partition": {
													Type: schema.TypeString, Optional: true, Description: "Partition name to clear",
												},
												"clear_all_partition_all": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Clear all partitions",
												},
												"clear_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number",
												},
												"clear_partition": {
													Type: schema.TypeString, Optional: true, Description: "Partition name to clear",
												},
												"clear_partition_all": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Clear all partitions",
												},
											},
										},
									},
									"link_disable_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"diseth": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the physical port number (Ethernet interface number)",
												},
												"dis_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
												},
											},
										},
									},
									"link_enable_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enaeth": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the physical port number (Ethernet interface number)",
												},
												"ena_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
												},
											},
										},
									},
									"monitor_relation": {
										Type: schema.TypeString, Optional: true, Default: "monitor-and", Description: "'monitor-and': Configures the monitors in current template to work with AND logic; 'monitor-or': Configures the monitors in current template to work with OR logic;",
									},
									"link_up_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkup_ethernet1": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
												},
												"link_up_sequence1": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
												},
												"linkup_ethernet2": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
												},
												"link_up_sequence2": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
												},
												"linkup_ethernet3": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
												},
												"link_up_sequence3": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequece number)",
												},
											},
										},
									},
									"link_down_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkdown_ethernet1": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
												},
												"link_down_sequence1": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
												},
												"linkdown_ethernet2": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
												},
												"link_down_sequence2": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the seqeuence number)",
												},
												"linkdown_ethernet3": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
												},
												"link_down_sequence3": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
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
						"link_block_as_down": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"link_down_on_restart": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"monitoring_mode": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mmode": {
										Type: schema.TypeString, Optional: true, Default: "and", Description: "'interdependent': INTERDEPENDENT monitoring behaviour; 'and': AND monitoring behaviour, Default;",
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
			"multi_queue_support": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Multi-Queue-Support",
						},
					},
				},
			},
			"ndisc_ra": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'good_recv': Good Router Solicitations (R.S.) Received; 'periodic_sent': Periodic Router Advertisements (R.A.) Sent; 'rate_limit': R.S. Rate Limited; 'bad_hop_limit': R.S. Bad Hop Limit; 'truncated': R.S. Truncated; 'bad_icmpv6_csum': R.S. Bad ICMPv6 Checksum; 'bad_icmpv6_code': R.S. Unknown ICMPv6 Code; 'bad_icmpv6_option': R.S. Bad ICMPv6 Option; 'l2_addr_and_unspec': R.S. Src Link-Layer Option and Unspecified Address; 'no_free_buffers': No Free Buffers to send R.A.;",
									},
								},
							},
						},
					},
				},
			},
			"netvsc_monitor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable auto-recovery from Rx/Tx hang",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"nsm_a10lb": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kill": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NSM will terminate a10lb if no response received",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"password_policy": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complexity": {
							Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Min length:8, Min Lower Case:2, Min Upper Case:2, Min Numbers:2, Min Special Character:1, CHANGE Min 8 Characters; 'Medium': Medium: Min length:6, Min Lower Case:2, Min Upper Case:2, Min Numbers:1, Min Special Character:1, CHANGE Min 6 Characters; 'Default': Default: Min length:9, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:1, CHANGE Min 1 Characters; 'Simple': Simple: Min length:4, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:0, CHANGE Min 4 Characters;",
						},
						"aging": {
							Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Max Age-60 Days; 'Medium': Medium: Max Age- 90 Days; 'Simple': Simple: Max Age-120 Days;",
						},
						"history": {
							Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Does not allow upto 5 old passwords; 'Medium': Medium: Does not allow upto 4 old passwords; 'Simple': Simple: Does not allow upto 3 old passwords;",
						},
						"min_pswd_len": {
							Type: schema.TypeInt, Optional: true, Description: "Configure custom password length",
						},
						"username_check": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Prohibition to set password contains user account, case sensitive; 'disable': Will not check if the password contains user account;",
						},
						"repeat_character_check": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Prohibition of consecutive repeated input of the same letter/number, case sensitive; 'disable': Will not check if the password contains repeat characters;",
						},
						"forbid_consecutive_character": {
							Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': Will disable the check; '3': Three consecutive characters on keyboard will not be allowed.; '4': Four consecutive characters on keyboard will not be allowed.; '5': Five consecutive characters on keyboard will not be allowed.;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"path_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l2hm_path_name": {
							Type: schema.TypeString, Required: true, Description: "Monitor Name",
						},
						"l2hm_vlan": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN id",
						},
						"l2hm_setup_test_api": {
							Type: schema.TypeInt, Optional: true, Description: "Test-API Interface (Ethernet Interface)",
						},
						"ifpair_eth_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
						},
						"ifpair_eth_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
						},
						"ifpair_trunk_start": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
						},
						"ifpair_trunk_end": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
						},
						"l2hm_attach": {
							Type: schema.TypeString, Optional: true, Description: "Monitor Name",
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
			"pbslb": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sockstress_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable sockstress protection",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_entries': Current PBSLB Entry Count; 'total_v4_entries_created': Total V4 Entry Count Created; 'total_v4_entries_freed': Total V4 Entry Count Freed; 'total_v6_entries_created': Total V6 Entry Count Created; 'total_v6_entries_freed': Total V6 Entry Count Freed; 'total_domain_entries_created': Total Domain Entry Count Created; 'total_domain_entries_freed': Total Domain Entry Count Freed; 'total_direct_action_entries_created': Total Direct Action Entry Count Created; 'total_direct_action_entries_freed': Total Direct Action Entry Count Freed; 'curr_entries_target_global': Current Entry Target Global; 'curr_entries_target_vserver': Current Entry Target Vserver; 'curr_entries_target_vport': Current Entry Target Vport; 'curr_entries_target_LOC': Current Entry Target LOC; 'curr_entries_target_rserver': Current Entry Target Rserver; 'curr_entries_target_rport': Current Entry Target Rport; 'curr_entries_target_service': Current Entry Target Service; 'curr_entries_stats': Current Entry Stats Count; 'curr_entries_target_global_dns_cache': Current Entry Target Global DNS Cache;",
									},
								},
							},
						},
					},
				},
			},
			"per_vlan_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bcast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "broadcast packets (per second limit)",
						},
						"ipmcast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "IP multicast packets (per second limit)",
						},
						"mcast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "multicast packets (per second limit)",
						},
						"unknown_ucast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "unknown unicast packets (per second limit)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"platformtype": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port_count": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_count_kernel": {
							Type: schema.TypeInt, Optional: true, Default: 18000, Description: "Total Ports to be allocated for kernel.",
						},
						"port_count_hm": {
							Type: schema.TypeInt, Optional: true, Default: 1024, Description: "Total Ports to be allocated for hm.",
						},
						"port_count_logging": {
							Type: schema.TypeInt, Optional: true, Default: 4096, Description: "Total Ports to be allocated for logging.",
						},
						"port_count_alg": {
							Type: schema.TypeInt, Optional: true, Default: 6000, Description: "Total Ports to be allocated for alg types.",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port_info": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ports": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_detection_interval": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Link detection interval in msecs",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"power_on_self_test": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"probe_network_devices": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"promiscuous_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run in promiscous mode settings",
			},
			"psu_info": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"q_in_q": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_all_ports": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 802.1QinQ on all physical ports",
						},
						"inner_tpid": {
							Type: schema.TypeString, Optional: true, Description: "TPID for inner VLAN (Inner TPID, 16 bit hex value, default is 8100)",
						},
						"outer_tpid": {
							Type: schema.TypeString, Optional: true, Description: "TPID for outer VLAN (Outer TPID, 16 bit hex value, default is 8100)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"queuing_buffer": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable micro-burst traffic support",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"radius": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listen_port": {
										Type: schema.TypeInt, Optional: true, Default: 1813, Description: "Configure the listen port of RADIUS server (default 1813) (Port number)",
									},
									"remote": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_list_name": {
																Type: schema.TypeString, Optional: true, Description: "IP-list name",
															},
															"ip_list_secret": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure shared secret",
															},
															"ip_list_secret_string": {
																Type: schema.TypeString, Optional: true, Description: "The RADIUS secret",
															},
														},
													},
												},
											},
										},
									},
									"secret": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure shared secret",
									},
									"secret_string": {
										Type: schema.TypeString, Optional: true, Description: "The RADIUS secret",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "Join a VRRP-A failover group",
									},
									"attribute": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_value": {
													Type: schema.TypeString, Optional: true, Description: "'inside-ipv6-prefix': Framed IPv6 Prefix; 'inside-ip': Inside IP address; 'inside-ipv6': Inside IPv6 address; 'imei': International Mobile Equipment Identity (IMEI); 'imsi': International Mobile Subscriber Identity (IMSI); 'msisdn': Mobile Subscriber Integrated Services Digital Network-Number (MSISDN); 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3; 'custom4': Customized attribute 4; 'custom5': Customized attribute 5; 'custom6': Customized attribute 6;",
												},
												"prefix_length": {
													Type: schema.TypeString, Optional: true, Description: "'32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;",
												},
												"prefix_vendor": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
												},
												"prefix_number": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
												},
												"name": {
													Type: schema.TypeString, Optional: true, Description: "Customized attribute name",
												},
												"value": {
													Type: schema.TypeString, Optional: true, Description: "'hexadecimal': Type of attribute value is hexadecimal;",
												},
												"custom_vendor": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
												},
												"custom_number": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
												},
												"vendor": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
												},
												"number": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
												},
											},
										},
									},
									"disable_reply": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle option for RADIUS reply packet(Default: Accounting response will be sent)",
									},
									"accounting_start": {
										Type: schema.TypeString, Optional: true, Default: "append-entry", Description: "'ignore': Ignore; 'append-entry': Append the AVPs to existing entry (default); 'replace-entry': Replace the AVPs of existing entry;",
									},
									"accounting_stop": {
										Type: schema.TypeString, Optional: true, Default: "delete-entry", Description: "'ignore': Ignore; 'delete-entry': Delete the entry (default); 'delete-entry-and-sessions': Delete the entry and data sessions associated(CGN only);",
									},
									"accounting_interim_update": {
										Type: schema.TypeString, Optional: true, Default: "ignore", Description: "'ignore': Ignore (default); 'append-entry': Append the AVPs to existing entry; 'replace-entry': Replace the AVPs of existing entry;",
									},
									"accounting_on": {
										Type: schema.TypeString, Optional: true, Default: "ignore", Description: "'ignore': Ignore (default); 'delete-entries-using-attribute': Delete entries matching attribute in RADIUS Table;",
									},
									"attribute_name": {
										Type: schema.TypeString, Optional: true, Description: "'msisdn': Clear using MSISDN; 'imei': Clear using IMEI; 'imsi': Clear using IMSI; 'custom1': Clear using CUSTOM1 attribute configured; 'custom2': Clear using CUSTOM2 attribute configured; 'custom3': Clear using CUSTOM3 attribute configured; 'custom4': Clear using CUSTOM4 attribute configured; 'custom5': Clear using CUSTOM5 attribute configured; 'custom6': Clear using CUSTOM6 attribute configured;",
									},
									"custom_attribute_name": {
										Type: schema.TypeString, Optional: true, Description: "Clear using customized attribute",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'msisdn-received': MSISDN Received; 'imei-received': IMEI Received; 'imsi-received': IMSI Received; 'custom-received': Custom attribute Received; 'radius-request-received': RADIUS Request Received; 'radius-request-dropped': RADIUS Request Dropped (Malformed Packet); 'request-bad-secret-dropped': RADIUS Request Bad Secret Dropped; 'request-no-key-vap-dropped': RADIUS Request No Key Attribute Dropped; 'request-malformed-dropped': RADIUS Request Malformed Dropped; 'request-ignored': RADIUS Request Ignored; 'radius-table-full': RADIUS Request Dropped (Table Full); 'secret-not-configured-dropped': RADIUS Secret Not Configured Dropped; 'ha-standby-dropped': HA Standby Dropped; 'ipv6-prefix-length-mismatch': Framed IPV6 Prefix Length Mismatch; 'invalid-key': Radius Request has Invalid Key Field; 'smp-created': RADIUS SMP Created; 'smp-deleted': RADIUS SMP Deleted; 'smp-mem-allocated': RADIUS SMP Memory Allocated; 'smp-mem-alloc-failed': RADIUS SMP Memory Allocation Failed; 'smp-mem-freed': RADIUS SMP Memory Freed; 'smp-in-rml': RADIUS SMP in RML; 'mem-allocated': RADIUS Memory Allocated; 'mem-alloc-failed': RADIUS Memory Allocation Failed; 'mem-freed': RADIUS Memory Freed; 'ha-sync-create-sent': HA Record Sync Create Sent; 'ha-sync-delete-sent': HA Record Sync Delete Sent; 'ha-sync-create-recv': HA Record Sync Create Received; 'ha-sync-delete-recv': HA Record Sync Delete Received; 'acct-on-filters-full': RADIUS Acct On Request Ignored(Filters Full); 'acct-on-dup-request': Duplicate RADIUS Acct On Request; 'ip-mismatch-delete': Radius Entry IP Mismatch Delete; 'ip-add-race-drop': Radius Entry IP Add Race Drop; 'ha-sync-no-key-vap-dropped': HA Record Sync No key dropped; 'inter-card-msg-fail-drop': Inter-Card Message Fail Drop; 'radius-packets-redirected': RADIUS packets redirected (SO); 'radius-packets-redirect-fail-dropped': RADIUS packets dropped due to redirect failure (SO); 'radius-packets-process-local': RADIUS packets processed locally without redirection (SO); 'radius-packets-dropped-not-lo': RADIUS packets dropped dest not loopback (SO); 'radius-inter-card-dup-redir': RADIUS packet dropped as redirected by other blade (SO);",
												},
											},
										},
									},
									"derived_attribute": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"usergroup": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"attribute": {
																Type: schema.TypeString, Optional: true, Description: "'imei': Specify the IMEI attribute.; 'imsi': Specify the IMSI attribute.; 'msisdn': Specify the MSISDN attribute.; 'custom1': Specify the custom1 attribute.; 'custom2': Specify the custom2 attribute.; 'custom3': Specify the custom3 attribute.; 'custom4': Specify the custom4 attribute.; 'custom5': Specify the custom5 attribute.; 'custom6': Specify the custom6 attribute.;",
															},
															"regex": {
																Type: schema.TypeString, Optional: true, Description: "Specify the regular expression to parse the value from a RADIUS attribute.",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"userid": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"attribute": {
																Type: schema.TypeString, Optional: true, Description: "'imei': Specify the IMEI attribute.; 'imsi': Specify the IMSI attribute.; 'msisdn': Specify the MSISDN attribute.; 'custom1': Specify the custom1 attribute.; 'custom2': Specify the custom2 attribute.; 'custom3': Specify the custom3 attribute.; 'custom4': Specify the custom4 attribute.; 'custom5': Specify the custom5 attribute.; 'custom6': Specify the custom6 attribute.;",
															},
															"regex": {
																Type: schema.TypeString, Optional: true, Description: "Specify the regular expression to parse the value from a RADIUS attribute.",
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
								},
							},
						},
					},
				},
			},
			"reboot": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"resource_accounting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"template_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Enter template name",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"app_resources": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gslb_device_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_device_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-device allowed (gslb-device count (default is max-value))",
															},
															"gslb_device_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_geo_location_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_geo_location_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-geo-location allowed (gslb-geo-location count (default is max-value))",
															},
															"gslb_geo_location_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_ip_list_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_ip_list_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-ip-list allowed (gslb-ip-list count (default is max-value))",
															},
															"gslb_ip_list_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_policy_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_policy_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-policy allowed (gslb-policy count (default is max-value))",
															},
															"gslb_policy_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_service_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service allowed (gslb-service count (default is max-value)",
															},
															"gslb_service_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_service_ip_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_ip_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-ip allowed (gslb-service-ip count (default is max-value))",
															},
															"gslb_service_ip_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_service_port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_port_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-port allowed ( gslb-service-port count (default is max-value))",
															},
															"gslb_service_port_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_site_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_site_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-site allowed (gslb-site count (default is max-value))",
															},
															"gslb_site_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_svc_group_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_svc_group_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-svc-group allowed (gslb-svc-group count (default is max-value))",
															},
															"gslb_svc_group_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-template allowed (gslb-template count (default is max-value))",
															},
															"gslb_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"gslb_zone_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_zone_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-zone allowed (gslb-zone count (default is max-value))",
															},
															"gslb_zone_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"health_monitor_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"health_monitor_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of health monitors allowed (health-monitor count (default is max-value))",
															},
															"health_monitor_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"real_port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"real_port_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-ports allowed (real-port count (default is max-value))",
															},
															"real_port_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"real_server_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"real_server_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-servers allowed (real-server count (default is max-value))",
															},
															"real_server_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"service_group_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"service_group_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of service groups allowed (service-group count (default is max-value))",
															},
															"service_group_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"virtual_server_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"virtual_server_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-servers allowed (virtual-server count (default is max-value))",
															},
															"virtual_server_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"virtual_port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"virtual_port_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-port allowed (virtual-port count (default is max-value))",
															},
															"virtual_port_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"cache_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cache_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of cache-template allowed (cache-template count (default is max-value))",
															},
															"cache_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"client_ssl_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"client_ssl_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of client-ssl-template allowed (client-ssl-template count (default is max-value))",
															},
															"client_ssl_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"conn_reuse_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"conn_reuse_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of conn-reuse-template allowed (conn-reuse-template count (default is max-value))",
															},
															"conn_reuse_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"fast_tcp_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fast_tcp_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-tcp-template allowed (fast-tcp-template count (default is max-value))",
															},
															"fast_tcp_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"fast_udp_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fast_udp_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-udp-template allowed (fast-udp-template count (default is max-value))",
															},
															"fast_udp_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"fix_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fix_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of fix-template allowed (fix-template count (default is max-value))",
															},
															"fix_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"http_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"http_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of http-template allowed (http-template count (default is max-value))",
															},
															"http_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"link_cost_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"link_cost_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of link-cost-template allowed (link-cost-template count (default is max-value))",
															},
															"link_cost_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"pbslb_entry_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"pbslb_entry_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of pbslb-entry allowed (pbslb-entry count)",
															},
															"pbslb_entry_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"persist_cookie_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"persist_cookie_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-cookie-template allowed (persist-cookie-template count (default is max-value))",
															},
															"persist_cookie_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"persist_srcip_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"persist_srcip_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-srcip-template allowed (persist-source-ip-template count (default is max-value))",
															},
															"persist_srcip_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"server_ssl_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server_ssl_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of server-ssl-template allowed (server-ssl-template count (default is max-value))",
															},
															"server_ssl_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"proxy_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"proxy_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of proxy-template allowed (server-ssl-template count (default is max-value))",
															},
															"proxy_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"stream_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"stream_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of stream-template allowed (server-ssl-template count (default is max-value))",
															},
															"stream_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"network_resources": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"static_ipv4_route_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_ipv4_route_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv4 routes allowed (Static ipv4 routes (default is max-value))",
															},
															"static_ipv4_route_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"static_ipv6_route_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_ipv6_route_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv6 routes allowed (Static ipv6 routes (default is max-value))",
															},
															"static_ipv6_route_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"ipv4_acl_line_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv4_acl_line_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV4 ACL lines (default is max-value))",
															},
															"ipv4_acl_line_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"ipv6_acl_line_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_acl_line_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV6 ACL lines (default is max-value))",
															},
															"ipv6_acl_line_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"static_arp_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_arp_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static arp entries allowed (Static arp (default is max-value))",
															},
															"static_arp_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"static_neighbor_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_neighbor_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static neighbor entries allowed (Static neighbors (default is max-value))",
															},
															"static_neighbor_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"static_mac_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_mac_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static MAC entries allowed (Static MACs (default is max-value))",
															},
															"static_mac_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"object_group_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_group_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of object groups allowed (Object group (default is max-value))",
															},
															"object_group_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"object_group_clause_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_group_clause_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of object group clauses allowed (Object group clauses (default is max-value))",
															},
															"object_group_clause_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
															},
														},
													},
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"system_resources": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bw_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bw_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the bandwidth limit in mbps (Bandwidth limit in Mbit/s (no limits applied by default))",
															},
															"bw_limit_watermark_disable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions)",
															},
														},
													},
												},
												"concurrent_session_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"concurrent_session_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the Concurrent Session limit (cps) (Concurrent-Session cps limit (no limits applied by default))",
															},
														},
													},
												},
												"l4_session_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l4_session_limit_max": {
																Type: schema.TypeString, Optional: true, Description: "Enter the l4 session limit in % (0.01% to 99.99%) (Enter a number from 0.01 to 99.99 (up to 2 digits precision))",
															},
															"l4_session_limit_min_guarantee": {
																Type: schema.TypeString, Optional: true, Default: "0", Description: "minimum guaranteed value in % (up to 2 digits precision) (Enter a number from 0 to 99.99)",
															},
														},
													},
												},
												"l4cps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l4cps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the L4 cps limit (L4 cps limit (no limits applied by default))",
															},
														},
													},
												},
												"l7cps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l7cps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "L7cps-limit (L7 cps limit (no limits applied by default))",
															},
														},
													},
												},
												"natcps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"natcps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the Nat cps limit (NAT cps limit (no limits applied by default))",
															},
														},
													},
												},
												"fwcps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fwcps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the Firewall cps limit (Firewall cps limit (no limits applied by default))",
															},
														},
													},
												},
												"ssl_throughput_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ssl_throughput_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the ssl throughput limit in mbps (SSL Througput limit in Mbit/s (no limits applied by default))",
															},
															"ssl_throughput_limit_watermark_disable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions)",
															},
														},
													},
												},
												"sslcps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"sslcps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the SSL cps limit (SSL cps limit (no limits applied by default))",
															},
														},
													},
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
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
					},
				},
			},
			"resource_usage": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat_pool_addr_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total configurable NAT Pool addresses in the System",
						},
						"l4_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total Sessions in the System",
						},
						"auth_portal_html_file_size": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Specify maximum html file size for each html page in auth portal (in KB)",
						},
						"auth_portal_image_file_size": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Specify maximum image file size for default portal (in KB)",
						},
						"max_aflex_file_size": {
							Type: schema.TypeInt, Optional: true, Default: 32, Description: "Set maximum aFleX file size (Maximum file size in KBytes, default is 32K)",
						},
						"aflex_table_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total aFleX table entry in the system (Total aFlex entry in the system)",
						},
						"class_list_ipv6_addr_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total IPv6 addresses for class-list",
						},
						"class_list_ac_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total entries for AC class-list",
						},
						"class_list_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total entries for class-list",
						},
						"max_aflex_authz_collection_number": {
							Type: schema.TypeInt, Optional: true, Default: 512, Description: "Specify the maximum number of collections supported by aFleX authorization",
						},
						"radius_table_size": {
							Type: schema.TypeInt, Optional: true, Description: "Total configurable CGNV6 RADIUS Table entries",
						},
						"authz_policy_number": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the maximum number of authorization policies",
						},
						"ipsec_sa_number": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the maximum number of IPsec SA",
						},
						"ram_cache_memory_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the maximum memory used by ram cache",
						},
						"auth_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total auth sessions in the system",
						},
						"ngwaf_cache_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the maximum cache entries for NGWAF",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"visibility": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"monitored_entity_count": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of monitored entities for visibility",
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
			"rfc_ipfix_ie_spec": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Use RFC-defined IPFIX information element lengths; 'disable': Use non-standard IPFIX information element lengths;",
			},
			"session": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_l4_conn': Total L4 Count; 'conn_counter': Conn Count; 'conn_freed_counter': Conn Freed; 'total_l4_packet_count': Total L4 Packet Count; 'total_l7_packet_count': Total L7 Packet Count; 'total_l4_conn_proxy': Total L4 Conn Proxy Count; 'total_l7_conn': Total L7 Conn; 'total_tcp_conn': Total TCP Conn; 'curr_free_conn': Curr Free Conn; 'tcp_est_counter': TCP Established; 'tcp_half_open_counter': TCP Half Open; 'tcp_half_close_counter': TCP Half Closed; 'udp_counter': UDP Count; 'ip_counter': IP Count; 'other_counter': Non TCP/UDP IP sessions; 'reverse_nat_tcp_counter': Reverse NAT TCP; 'reverse_nat_udp_counter': Reverse NAT UDP; 'tcp_syn_half_open_counter': TCP SYN Half Open; 'conn_smp_alloc_counter': Conn SMP Alloc; 'conn_smp_free_counter': Conn SMP Free; 'conn_smp_aged_counter': Conn SMP Aged; 'ssl_count_curr': Curr SSL Count; 'ssl_count_total': Total SSL Count; 'server_ssl_count_curr': Current SSL Server Count; 'server_ssl_count_total': Total SSL Server Count; 'client_ssl_reuse_total': Total SSL Client Reuse; 'server_ssl_reuse_total': Total SSL Server Reuse; 'ssl_failed_total': Total SSL Failures; 'ssl_failed_ca_verification': SSL Cert Auth Verification Errors; 'ssl_server_cert_error': SSL Server Cert Errors; 'ssl_client_cert_auth_fail': SSL Client Cert Auth Failures; 'total_ip_nat_conn': Total IP Nat Conn; 'total_l2l3_conn': Totl L2/L3 Connections; 'client_ssl_ctx_malloc_failure': Client SSL Ctx malloc Failures; 'conn_type_0_available': Conn Type 0 Available; 'conn_type_1_available': Conn Type 1 Available; 'conn_type_2_available': Conn Type 2 Available; 'conn_type_3_available': Conn Type 3 Available; 'conn_type_4_available': Conn Type 4 Available; 'conn_smp_type_0_available': Conn SMP Type 0 Available; 'conn_smp_type_1_available': Conn SMP Type 1 Available; 'conn_smp_type_2_available': Conn SMP Type 2 Available; 'conn_smp_type_3_available': Conn SMP Type 3 Available; 'conn_smp_type_4_available': Conn SMP Type 4 Available; 'sctp-half-open-counter': SCTP Half Open; 'sctp-est-counter': SCTP Established; 'nonssl_bypass': NON SSL Bypass Count; 'ssl_failsafe_total': Total SSL Failsafe Count; 'ssl_forward_proxy_failed_handshake_total': Total SSL Forward Proxy Failed Handshake Count; 'ssl_forward_proxy_failed_tcp_total': Total SSL Forward Proxy Failed TCP Count; 'ssl_forward_proxy_failed_crypto_total': Total SSL Forward Proxy Failed Crypto Count; 'ssl_forward_proxy_failed_cert_verify_total': Total SSL Forward Proxy Failed Certificate Verification Count; 'ssl_forward_proxy_invalid_ocsp_stapling_total': Total SSL Forward Proxy Invalid OCSP Stapling Count; 'ssl_forward_proxy_revoked_ocsp_total': Total SSL Forward Proxy Revoked OCSP Response Count; 'ssl_forward_proxy_failed_cert_signing_total': Total SSL Forward Proxy Failed Certificate Signing Count; 'ssl_forward_proxy_failed_ssl_version_total': Total SSL Forward Proxy Unsupported version Count; 'ssl_forward_proxy_sni_bypass_total': Total SSL Forward Proxy SNI Bypass Count; 'ssl_forward_proxy_client_auth_bypass_total': Total SSL Forward Proxy Client Auth Bypass Count; 'conn_app_smp_alloc_counter': Conn APP SMP Alloc; 'diameter_conn_counter': Diameter Conn Count; 'diameter_conn_freed_counter': Diameter Conn Freed; 'debug_tcp_counter': Hidden TCP sessions for CGNv6 Stateless Technologies; 'debug_udp_counter': Hidden UDP sessions for CGNv6 Stateless Technologies; 'total_fw_conn': Total Firewall Conn; 'total_local_conn': Total Local Conn; 'total_curr_conn': Total Curr Conn; 'client_ssl_fatal_alert': client ssl fatal alert; 'client_ssl_fin_rst': client ssl fin rst; 'fp_session_fin_rst': FP Session FIN/RST; 'server_ssl_fatal_alert': server ssl fatal alert; 'server_ssl_fin_rst': server ssl fin rst; 'client_template_int_err': client template internal error; 'client_template_unknown_err': client template unknown error; 'server_template_int_err': server template int error; 'server_template_unknown_err': server template unknown error; 'total_debug_conn': Total Debug Conn; 'ssl_forward_proxy_failed_aflex_total': Total SSL Forward Proxy AFLEX Count; 'ssl_forward_proxy_cert_subject_bypass_total': Total SSL Forward Proxy Certificate Subject Bypass Count; 'ssl_forward_proxy_cert_issuer_bypass_total': Total SSL Forward Proxy Certificate Issuer Bypass Count; 'ssl_forward_proxy_cert_san_bypass_total': Total SSL Forward Proxy Certificate SAN Bypass Count; 'ssl_forward_proxy_no_sni_bypass_total': Total SSL Forward Proxy No SNI Bypass Count; 'ssl_forward_proxy_no_sni_reset_total': Total SSL Forward Proxy No SNI reset Count; 'ssl_forward_proxy_username_bypass_total': Total SSL Forward Proxy Username Bypass Count; 'ssl_forward_proxy_ad_grpup_bypass_total': Total SSL Forward Proxy AD-Group Bypass Count; 'diameter_concurrent_user_sessions_counter': Diameter Concurrent User-Sessions; 'client_ssl_session_ticket_reuse_total': Total SSL Client Session Ticket Reuse; 'server_ssl_session_ticket_reuse_total': Total SSL Server Session Ticket Reuse; 'total_clientside_early_data_connections': Total clientside early data connections; 'total_serverside_early_data_connections': Total serverside early data connections; 'total_clientside_failed_early_data-connections': Total clientside failed early data connections; 'total_serverside_failed_early_data-connections': Total serverside failed early data connections; 'ssl_forward_proxy_esni_bypass_total': Total SSL Forward Proxy ESNI Bypass Count; 'ssl_forward_proxy_esni_reset_total': Total SSL Forward Proxy ESNI Drop Count; 'total_logging_conn': Total Logging Conn; 'gtp_c_est_counter': GTP-C Established; 'gtp_c_half_open_counter': GTP-C Half Open; 'gtp_u_counter': GTP-U Count; 'gtp_c_echo_counter': GTP-C Echo Count; 'gtp_u_echo_counter': GTP-U Echo Count; 'gtp_curr_free_conn': GTP Current Available Conn; 'gtp_cum_conn_counter': GTP cumulative Conn Count; 'gtp_cum_conn_freed_counter': GTP cumulative Conn Freed; 'fw_blacklist_sess': Blacklist Sessions; 'fw_blacklist_sess_created': Blacklist Session Created; 'fw_blacklist_sess_freed': Blacklist Session Freed; 'server_tcp_est_counter': Server TCP Established; 'server_tcp_half_open_counter': Server TCP Half Open; 'sched_conn_with_wrong_next_idx_to_rml': Attempt to Put a Conn to RML Whose next_idx is NOT Invalid; 'free_conn_not_in_sp': Attempt to Free a Conn Whoes Address Not in Session Pool; 'total_system_sessions': Total system-wide sessions;",
									},
								},
							},
						},
					},
				},
			},
			"session_reclaim_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nscan_limit": {
							Type: schema.TypeInt, Optional: true, Default: 4096, Description: "smp session scan limit (number of smp sessions per scan)",
						},
						"scan_freq": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "smp session scan frequency (scan per second)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"set_rxtx_desc_size": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
						},
						"rxd_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set new rx-descriptor size",
						},
						"txd_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set new tx-descriptor size",
						},
					},
				},
			},
			"set_rxtx_queue": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
						},
						"rxq_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set number of new rx queues",
						},
						"txq_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set number of new tx queues",
						},
					},
				},
			},
			"set_tcp_syn_per_sec": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_syn_value": {
							Type: schema.TypeInt, Optional: true, Default: 70, Description: "Configure Tcp SYN's per sec, default 70",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"shared_poll_mode": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable shared poll mode",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable shared poll mode",
						},
					},
				},
			},
			"shell_privileges": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_shell_privileges": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the shell privileges for a given customer",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"shm_logging": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable shared memory based logging",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"shutdown": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"software_tcam": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable software TCAM",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"spe_profile": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "ipv4-ipv6", Description: "'ipv4-only': Enable IPv4 HW forward entries only; 'ipv6-only': Enable IPv6 HW forward entries only; 'ipv4-ipv6': Enable Both IPv4/IPv6 HW forward entries (shared);",
						},
					},
				},
			},
			"spe_status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"src_ip_hash_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable source ip hash",
			},
			"ssl_hw_memory": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mem_block_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mem_block": {
										Type: schema.TypeString, Optional: true, Description: "'ssl_mem': Shared SSL memory; 'ssl_context': Cipher session context memory for SSL engine 1; 'ssl_context_2': Cipher session context memory for SSL engine 2; 'ssl_context_3': Cipher session context memory for SSL engine 3; 'ssl_context_4': Cipher session context memory for SSL engine 4; 'ssl_context_5': Cipher session context memory for SSL engine 5; 'ssl_context_6': Cipher session context memory for SSL engine 6; 'ssl_context_7': Cipher session context memory for SSL engine 7; 'ssl_context_8': Cipher session context memory for SSL engine 8; 'ssl_context_9': Cipher session context memory for SSL engine 9; 'ssl_context_10': Cipher session context memory for SSL engine 10; 'ssl_context_11': Cipher session context memory for SSL engine 11; 'ssl_context_12': Cipher session context memory for SSL engine 12; 'ssl_context_13': Cipher session context memory for SSL engine 13; 'ssl_context_14': Cipher session context memory for SSL engine 14; 'ssl_context_15': Cipher session context memory for SSL engine 15; 'ssl_context_16': Cipher session context memory for SSL engine 16;",
									},
									"size": {
										Type: schema.TypeInt, Optional: true, Description: "Size of the block",
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
			"ssl_req_q": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'num-ssl-queues': num-ssl-queues; 'ssl-req-q-depth-tot': ssl-req-q-depth-tot; 'ssl-req-q-inuse-tot': ssl-req-q-inuse-tot; 'ssl-hw-q-depth-tot': ssl-hw-q-depth-tot; 'ssl-hw-q-inuse-tot': ssl-hw-q-inuse-tot;",
									},
								},
							},
						},
					},
				},
			},
			"ssl_scv": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable server certificate validation for all SSL connections",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ssl_scv_verify_crl_sign": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable verify CRL signature during SCV",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ssl_scv_verify_host": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable verify host during SCV",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ssl_set_compatible_cipher": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable setting common cipher suite in management plane",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ssl_status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"syslog_time_msec": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_flag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
					},
				},
			},
			"system_chassis_port_split_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable port split for the chassis",
			},
			"table_integrity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"table": {
							Type: schema.TypeString, Optional: true, Default: "all", Description: "'all': All tables;",
						},
						"audit_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable table integrity audit; 'disable': Disable table integrity audit;",
						},
						"auto_sync_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable auto-sync; 'disable': Disable auto-sync;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'arp-tbl-sync-start-ts-m-1st': arp table sync start time stamp master; 'nd6-tbl-sync-start-ts-m-1st': nd6 table sync start time stamp master; 'ipv4-fib-tbl-sync-start-ts-m-1st': ipv4-fib table sync start time stamp master; 'ipv6-fib-tbl-sync-start-ts-m-1st': ipv6-fib table sync start time stamp master; 'mac-tbl-sync-start-ts-m-1st': mac table sync start time stamp master; 'arp-tbl-sync-start-ts-b-1st': arp table sync start time stamp blade; 'nd6-tbl-sync-start-ts-b-1st': nd6 table sync start time stamp blade; 'ipv4-fib-tbl-sync-start-ts-b-1st': ipv4-fib table sync start time stamp blade; 'ipv6-fib-tbl-sync-start-ts-b-1st': ipv6-fib table sync start time stamp blade; 'mac-tbl-sync-start-ts-b-1st': mac table sync start time stamp blade; 'arp-tbl-sync-entries-sent-m-1st': arp table entries sent from master for T0 synchronization; 'nd6-tbl-sync-entries-sent-m-1st': nd6 table entries sent from master for T0 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-1st': ipv4-fib table entries sent from master for T0 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-1st': ipv6-fib table entries sent from master for T0 synchronization; 'mac-tbl-sync-entries-sent-m-1st': mac table entries sent from master for T0 synchronization; 'arp-tbl-sync-entries-rcvd-b-1st': arp table entries received on blade for T0 synchronization; 'nd6-tbl-sync-entries-rcvd-b-1st': nd6 table entries received on blade for T0 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-1st': ipv4-fib table entries received on blade for T0 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-1st': ipv6-fib table entries received on blade for T0 synchronization; 'mac-tbl-sync-entries-rcvd-b-1st': mac table entries received on blade for T0 synchronization; 'arp-tbl-sync-entries-added-b-1st': arp table entries added on blade for T0 synchronization; 'nd6-tbl-sync-entries-added-b-1st': nd6 table entries added on blade for T0 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-1st': ipv4-fib table entries added on blade for T0 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-1st': ipv6-fib table entries added on blade for T0 synchronization; 'mac-tbl-sync-entries-added-b-1st': mac table entries added on blade for T0 synchronization; 'arp-tbl-sync-entries-removed-b-1st': arp table entries removed on blade for T0 synchronization; 'nd6-tbl-sync-entries-removed-b-1st': arp table entries removed on blade for T0 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-1st': arp table entries removed on blade for T0 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-1st': arp table entries removed on blade for T0 synchronization; 'mac-tbl-sync-entries-removed-b-1st': arp table entries removed on blade for T0 synchronization; 'arp-tbl-sync-end-ts-m-1st': arp table sync end time stamp master for T0 synchronization; 'nd6-tbl-sync-end-ts-m-1st': nd6 table sync end time stamp master for T0 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-1st': ipv4-fib table sync end time stamp master for T0 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-1st': ipv6-fib table sync end time stamp master for T0 synchronization; 'mac-tbl-sync-end-ts-m-1st': mac table sync end time stamp master for T0 synchronization; 'arp-tbl-sync-end-ts-b-1st': arp table sync end time stamp blade for T0 synchronization; 'nd6-tbl-sync-end-ts-b-1st': nd6 table sync end time stamp blade for T0 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-1st': ipv4-fib table sync end time stamp blade for T0 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-1st': ipv6-fib table sync end time stamp blade for T0 synchronization; 'mac-tbl-sync-end-ts-b-1st': mac table sync end time stamp blade for T0 synchronization; 'arp-tbl-sync-start-ts-m-2nd': arp table sync start time stamp master for T-1 synchronization; 'nd6-tbl-sync-start-ts-m-2nd': nd6 table sync start time stamp master for T-1 synchronization; 'ipv4-fib-tbl-sync-start-ts-m-2nd': ipv4-fib table sync start time stamp master for T-1 synchronization; 'ipv6-fib-tbl-sync-start-ts-m-2nd': ipv6-fib table sync start time stamp master for T-1 synchronization; 'mac-tbl-sync-start-ts-m-2nd': mac table sync start time stamp master for T-1 synchronization; 'arp-tbl-sync-start-ts-b-2nd': arp table sync start time stamp blade for T-1 synchronization; 'nd6-tbl-sync-start-ts-b-2nd': nd6 table sync start time stamp blade for T-1 synchronization; 'ipv4-fib-tbl-sync-start-ts-b-2nd': ipv4-fib table sync start time stamp blade for T-1 synchronization; 'ipv6-fib-tbl-sync-start-ts-b-2nd': ipv6-fib table sync start time stamp blade for T-1 synchronization; 'mac-tbl-sync-start-ts-b-2nd': mac table sync start time stamp blade for T-1 synchronization; 'arp-tbl-sync-entries-sent-m-2nd': arp table entries sent from master for T-1 synchronization; 'nd6-tbl-sync-entries-sent-m-2nd': nd6 table entries sent from master for T-1 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-2nd': ipv4-fib table entries sent from master for T-1 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-2nd': ipv6-fib table entries sent from master for T-1 synchronization; 'mac-tbl-sync-entries-sent-m-2nd': mac table entries sent from master for T-1 synchronization; 'arp-tbl-sync-entries-rcvd-b-2nd': arp table entries received in blade for T-1 synchronization; 'nd6-tbl-sync-entries-rcvd-b-2nd': nd6 table entries received in blade for T-1 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-2nd': ipv4-fib table entries received in blade for T-1 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-2nd': ipv6-fib table entries received in blade for T-1 synchronization; 'mac-tbl-sync-entries-rcvd-b-2nd': mac table entries received in blade for T-1 synchronization; 'arp-tbl-sync-entries-added-b-2nd': arp table entries added in blade for T-1 synchronization; 'nd6-tbl-sync-entries-added-b-2nd': nd6 table entries added in blade for T-1 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-2nd': ipv4-fib table entries added in blade for T-1 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-2nd': ipv6-fib table entries added in blade for T-1 synchronization; 'mac-tbl-sync-entries-added-b-2nd': mac table entries added in blade for T-1 synchronization; 'arp-tbl-sync-entries-removed-b-2nd': arp table entries removed in blade for T-1 synchronization; 'nd6-tbl-sync-entries-removed-b-2nd': nd6 table entries removed in blade for T-1 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-2nd': ipv4-fib table entries removed in blade for T-1 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-2nd': ipv6-fib table entries removed in blade for T-1 synchronization; 'mac-tbl-sync-entries-removed-b-2nd': mac table entries removed in blade for T-1 synchronization; 'arp-tbl-sync-end-ts-m-2nd': arp table sync end time stamp master for T-1 synchronization; 'nd6-tbl-sync-end-ts-m-2nd': nd6 table sync end time stamp master for T-1 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-2nd': ipv4-fib table sync end time stamp master for T-1 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-2nd': ipv6-fib table sync end time stamp master for T-1 synchronization; 'mac-tbl-sync-end-ts-m-2nd': mac table sync end time stamp master for T-1 synchronization; 'arp-tbl-sync-end-ts-b-2nd': arp table sync end time stamp blade for T-1 synchronization; 'nd6-tbl-sync-end-ts-b-2nd': nd6 table sync end time stamp blade for T-1 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-2nd': ipv4-fib table sync end time stamp blade for T-1 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-2nd': ipv6-fib table sync end time stamp blade for T-1 synchronization; 'mac-tbl-sync-end-ts-b-2nd': mac table sync end time stamp blade for T-1 synchronization; 'arp-tbl-sync-start-ts-m-3rd': arp table sync start time stamp master for T-2 synchronization; 'nd6-tbl-sync-start-ts-m-3rd': nd6 table sync start time stamp master for T-2 synchronization;",
									},
									"counters2": {
										Type: schema.TypeString, Optional: true, Description: "'ipv4-fib-tbl-sync-start-ts-m-3rd': ipv4-fib table sync start time stamp master for T-2 synchronization; 'ipv6-fib-tbl-sync-start-ts-m-3rd': ipv6-fib table sync start time stamp master for T-2 synchronization; 'mac-tbl-sync-start-ts-m-3rd': mac table sync start time stamp master for T-2 synchronization; 'arp-tbl-sync-start-ts-b-3rd': arp table sync start time stamp blade for T-2 synchronization; 'nd6-tbl-sync-start-ts-b-3rd': nd6 table sync start time stamp blade for T-2 synchronization; 'ipv4-fib-tbl-sync-start-ts-b-3rd': ipv4-fib table sync start time stamp blade for T-2 synchronization; 'ipv6-fib-tbl-sync-start-ts-b-3rd': ipv6-fib table sync start time stamp blade for T-2 synchronization; 'mac-tbl-sync-start-ts-b-3rd': mac table sync start time stamp blade for T-2 synchronization; 'arp-tbl-sync-entries-sent-m-3rd': arp table entries sent from master for T-2 synchronization; 'nd6-tbl-sync-entries-sent-m-3rd': nd6 table entries sent from master for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-3rd': ipv4-fib table entries sent from master for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-3rd': ipv6-fib table entries sent from master for T-2 synchronization; 'mac-tbl-sync-entries-sent-m-3rd': mac table entries sent from master for T-2 synchronization; 'arp-tbl-sync-entries-rcvd-b-3rd': arp table entries received in blade for T-2 synchronization; 'nd6-tbl-sync-entries-rcvd-b-3rd': nd6 table entries received in blade for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-3rd': ipv4-fib table entries received in blade for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-3rd': ipv6-fib table entries received in blade for T-2 synchronization; 'mac-tbl-sync-entries-rcvd-b-3rd': mac table entries received in blade for T-2 synchronization; 'arp-tbl-sync-entries-added-b-3rd': arp table entries added in blade for T-2 synchronization; 'nd6-tbl-sync-entries-added-b-3rd': nd6 table entries added in blade for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-3rd': ipv4-fib table entries added in blade for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-3rd': ipv6-fib table entries added in blade for T-2 synchronization; 'mac-tbl-sync-entries-added-b-3rd': mac table entries added in blade for T-2 synchronization; 'arp-tbl-sync-entries-removed-b-3rd': arp table entries removed in blade for T-2 synchronization; 'nd6-tbl-sync-entries-removed-b-3rd': nd6 table entries removed in blade for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-3rd': ipv4-fib table entries removed in blade for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-3rd': ipv6-fib table entries removed in blade for T-2 synchronization; 'mac-tbl-sync-entries-removed-b-3rd': mac table entries removed in blade for T-2 synchronization; 'arp-tbl-sync-end-ts-m-3rd': arp table sync end time stamp master for T-2 synchronization; 'nd6-tbl-sync-end-ts-m-3rd': nd6 table sync end time stamp master for T-2 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-3rd': ipv4-fib table sync end time stamp master for T-2 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-3rd': ipv6-fib table sync end time stamp master for T-2 synchronization; 'mac-tbl-sync-end-ts-m-3rd': mac table sync end time stamp master for T-2 synchronization; 'arp-tbl-sync-end-ts-b-3rd': arp table sync end time stamp blade for T-2 synchronization; 'nd6-tbl-sync-end-ts-b-3rd': nd6 table sync end time stamp blade for T-2 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-3rd': ipv4-fib table sync end time stamp blade for T-2 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-3rd': ipv6-fib table sync end time stamp blade for T-2 synchronization; 'mac-tbl-sync-end-ts-b-3rd': mac table sync end time stamp blade for T-2 synchronization; 'arp-tbl-sync-start-ts-m-4th': arp table sync start time stamp master for T-3 synchronization; 'nd6-tbl-sync-start-ts-m-4th': nd6 table sync start time stamp master for T-3 synchronization; 'ipv4-fib-tbl-sync-start-ts-m-4th': ipv4-fib table sync start time stamp master for T-3 synchronization; 'ipv6-fib-tbl-sync-start-ts-m-4th': ipv6-fib table sync start time stamp master for T-3 synchronization; 'mac-tbl-sync-start-ts-m-4th': mac table sync start time stamp master for T-3 synchronization; 'arp-tbl-sync-start-ts-b-4th': arp table sync start time stamp blade for T-3 synchronization; 'nd6-tbl-sync-start-ts-b-4th': nd6 table sync start time stamp blade for T-3 synchronization; 'ipv4-fib-tbl-sync-start-ts-b-4th': ipv4-fib table sync start time stamp blade for T-3 synchronization; 'ipv6-fib-tbl-sync-start-ts-b-4th': ipv6-fib table sync start time stamp blade for T-3 synchronization; 'mac-tbl-sync-start-ts-b-4th': mac table sync start time stamp blade for T-3 synchronization; 'arp-tbl-sync-entries-sent-m-4th': arp table entries sent from master for T-3 synchronization; 'nd6-tbl-sync-entries-sent-m-4th': nd6 table entries sent from master for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-4th': ipv4-fib table entries sent from master for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-4th': ipv6-fib table entries sent from master for T-3 synchronization; 'mac-tbl-sync-entries-sent-m-4th': mac table entries sent from master for T-3 synchronization; 'arp-tbl-sync-entries-rcvd-b-4th': arp table entries received in blade for T-3 synchronization; 'nd6-tbl-sync-entries-rcvd-b-4th': nd6 table entries received in blade for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-4th': ipv4-fib table entries received in blade for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-4th': ipv6-fib table entries received in blade for T-3 synchronization; 'mac-tbl-sync-entries-rcvd-b-4th': mac table entries received in blade for T-3 synchronization; 'arp-tbl-sync-entries-added-b-4th': arp table entries added in blade for T-3 synchronization; 'nd6-tbl-sync-entries-added-b-4th': nd6 table entries added in blade for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-4th': ipv4-fib table entries added in blade for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-4th': ipv6-fib table entries added in blade for T-3 synchronization; 'mac-tbl-sync-entries-added-b-4th': mac table entries added in blade for T-3 synchronization; 'arp-tbl-sync-entries-removed-b-4th': arp table entries removed in blade for T-3 synchronization; 'nd6-tbl-sync-entries-removed-b-4th': nd6 table entries removed in blade for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-4th': ipv4-fib table entries removed in blade for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-4th': ipv6-fib table entries removed in blade for T-3 synchronization; 'mac-tbl-sync-entries-removed-b-4th': mac table entries removed in blade for T-3 synchronization; 'arp-tbl-sync-end-ts-m-4th': arp table sync end time stamp master for T-3 synchronization; 'nd6-tbl-sync-end-ts-m-4th': nd6 table sync end time stamp master for T-3 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-4th': ipv4-fib table sync end time stamp master for T-3 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-4th': ipv6-fib table sync end time stamp master for T-3 synchronization; 'mac-tbl-sync-end-ts-m-4th': mac table sync end time stamp master for T-3 synchronization; 'arp-tbl-sync-end-ts-b-4th': arp table sync end time stamp blade for T-3 synchronization; 'nd6-tbl-sync-end-ts-b-4th': nd6 table sync end time stamp blade for T-3 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-4th': ipv4-fib table sync end time stamp blade for T-3 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-4th': ipv6-fib table sync end time stamp blade for T-3 synchronization; 'mac-tbl-sync-end-ts-b-4th': mac table sync end time stamp blade for T-3 synchronization; 'arp-tbl-sync-start-ts-m-5th': arp table sync start time stamp master for T-4 synchronization;",
									},
									"counters3": {
										Type: schema.TypeString, Optional: true, Description: "'nd6-tbl-sync-start-ts-m-5th': nd6 table sync start time stamp master for T-4 synchronization; 'ipv4-fib-tbl-sync-start-ts-m-5th': ipv4-fib table sync start time stamp master for T-4 synchronization; 'ipv6-fib-tbl-sync-start-ts-m-5th': ipv6-fib table sync start time stamp master for T-4 synchronization; 'mac-tbl-sync-start-ts-m-5th': mac table sync start time stamp master for T-4 synchronization; 'arp-tbl-sync-start-ts-b-5th': arp table sync start time stamp blade for T-4 synchronization; 'nd6-tbl-sync-start-ts-b-5th': nd6 table sync start time stamp blade for T-4 synchronization; 'ipv4-fib-tbl-sync-start-ts-b-5th': ipv4-fib table sync start time stamp blade for T-4 synchronization; 'ipv6-fib-tbl-sync-start-ts-b-5th': ipv6-fib table sync start time stamp blade for T-4 synchronization; 'mac-tbl-sync-start-ts-b-5th': mac table sync start time stamp blade for T-4 synchronization; 'arp-tbl-sync-entries-sent-m-5th': arp table sync start time stamp blade for T-4 synchronization; 'nd6-tbl-sync-entries-sent-m-5th': nd6 table sync start time stamp blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-5th': ipv4-fib table sync start time stamp blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-5th': ipv6-fib table sync start time stamp blade for T-4 synchronization; 'mac-tbl-sync-entries-sent-m-5th': mac table sync start time stamp blade for T-4 synchronization; 'arp-tbl-sync-entries-rcvd-b-5th': arp table entries received in blade for T-4 synchronization; 'nd6-tbl-sync-entries-rcvd-b-5th': nd6 table entries received in blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-5th': ipv4-fib table entries received in blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-5th': ipv6-fib table entries received in blade for T-4 synchronization; 'mac-tbl-sync-entries-rcvd-b-5th': mac table entries received in blade for T-4 synchronization; 'arp-tbl-sync-entries-added-b-5th': arp table entries added in blade for T-4 synchronization; 'nd6-tbl-sync-entries-added-b-5th': nd6 table entries added in blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-5th': ipv4-fib table entries added in blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-5th': ipv6-fib table entries added in blade for T-4 synchronization; 'mac-tbl-sync-entries-added-b-5th': mac table entries added in blade for T-4 synchronization; 'arp-tbl-sync-entries-removed-b-5th': arp table entries removed in blade for T-4 synchronization; 'nd6-tbl-sync-entries-removed-b-5th': nd6 table entries removed in blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-5th': ipv4-fib table entries removed in blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-5th': ipv6-fib table entries removed in blade for T-4 synchronization; 'mac-tbl-sync-entries-removed-b-5th': mac table entries removed in blade for T-4 synchronization; 'arp-tbl-sync-end-ts-m-5th': arp table sync end time stamp master for T-4 synchronization; 'nd6-tbl-sync-end-ts-m-5th': nd6 table sync end time stamp master for T-4 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-5th': ipv4-fib table sync end time stamp master for T-4 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-5th': ipv6-fib table sync end time stamp master for T-4 synchronization; 'mac-tbl-sync-end-ts-m-5th': mac table sync end time stamp master for T-4 synchronization; 'arp-tbl-sync-end-ts-b-5th': arp table sync end time stamp blade for T-4 synchronization; 'nd6-tbl-sync-end-ts-b-5th': nd6 table sync end time stamp blade for T-4 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-5th': ipv4-fib table sync end time stamp blade for T-4 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-5th': ipv6-fib table sync end time stamp blade for T-4 synchronization; 'mac-tbl-sync-end-ts-b-5th': mac table sync end time stamp blade for T-4 synchronization; 'arp-tbl-sync-m': arp table sync count in master; 'nd6-tbl-sync-m': nd6 table sync count in master; 'ipv4-fib-tbl-sync-m': ipv4-fib table sync count in master; 'ipv6-fib-tbl-sync-m': ipv6-fib table sync count in master; 'mac-tbl-sync-m': mac table sync count in master; 'arp-tbl-sync-b': arp table sync count in blade; 'nd6-tbl-sync-b': nd6 table sync count in blade; 'ipv4-fib-tbl-sync-b': ipv4-fib table sync count in blade; 'ipv6-fib-tbl-sync-b': ipv6-fib table sync count in blade; 'mac-tbl-sync-b': mac table sync count in blade; 'arp-tbl-cksum-m': arp table checksum count in master; 'nd6-tbl-cksum-m': nd6 table checksum count in master; 'ipv4-fib-tbl-cksum-m': ipv4-fib table checksum count in master; 'ipv6-fib-tbl-cksum-m': ipv6-fib table checksum count in master; 'mac-tbl-cksum-m': mac table checksum count in master; 'arp-tbl-cksum-b': arp table checksum count in blade; 'nd6-tbl-cksum-b': nd6 table checksum count in blade; 'ipv4-fib-tbl-cksum-b': ipv4-fib table checksum count in blade; 'ipv6-fib-tbl-cksum-b': ipv6-fib table checksum count in blade; 'mac-tbl-cksum-b': mac table checksum count in blade; 'arp-tbl-cksum-mismatch-b': arp table checksum mismatch count in blade; 'nd6-tbl-cksum-mismatch-b': nd6 table checksum mismatch count in blade; 'ipv4-fib-tbl-cksum-mismatch-b': ipv4-fib table checksum mismatch count in blade; 'ipv6-fib-tbl-cksum-mismatch-b': ipv6-fib table checksum mismatch count in blade; 'mac-tbl-cksum-mismatch-b': mac table checksum mismatch count in blade; 'arp-tbl-cksum-cancel-m': arp table checksum cancelled count in master; 'nd6-tbl-cksum-cancel-m': nd6 table checksum cancelled count in master; 'ipv4-fib-tbl-cksum-cancel-m': ipv4-fib table checksum cancelled count in master; 'ipv6-fib-tbl-cksum-cancel-m': ipv6-fib table checksum cancelled count in master; 'mac-tbl-cksum-cancel-m': mac table checksum cancelled count in master;",
									},
								},
							},
						},
					},
				},
			},
			"tcp": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'activeopens': Active open conns; 'passiveopens': Passive open conns; 'attemptfails': Connect attemp failures; 'estabresets': Resets rcvd on EST conn; 'insegs': Total in TCP packets; 'outsegs': Total out TCP packets; 'retranssegs': Retransmited packets; 'inerrs': Input errors; 'outrsts': Reset Sent; 'sock_alloc': Sockets allocated; 'orphan_count': Orphan sockets; 'mem_alloc': Memory alloc; 'recv_mem': Total rx buffer; 'send_mem': Total tx buffer; 'currestab': Currently EST conns; 'currsyssnt': TCP in SYN-SNT state; 'currsynrcv': TCP in SYN-RCV state; 'currfinw1': TCP in FIN-W1 state; 'currfinw2': TCP FIN-W2 state; 'currtimew': TCP TimeW state; 'currclose': TCP in Close state; 'currclsw': TCP in CloseW state; 'currlack': TCP in LastACK state; 'currlstn': TCP in Listen state; 'currclsg': TCP in Closing state; 'pawsactiverejected': TCP paw active rej; 'syn_rcv_rstack': Rcv RST|ACK on SYN; 'syn_rcv_rst': Rcv RST on SYN; 'syn_rcv_ack': Rcv ACK on SYN; 'ax_rexmit_syn': TCP rexmit SYN; 'tcpabortontimeout': TCP abort on timeout; 'noroute': TCPIP out noroute; 'exceedmss': MSS exceeded pkt dropped; 'tfo_conns': TFO Total Connections; 'tfo_actives': TFO Current Actives; 'tfo_denied': TFO Denied; 'syn_rcv_rexmit': Rcv SYN rexmit; 'sock_init': Socket init; 'invalid_drop': Invalid packet drop; 'delayed_free': Delayed free return;",
									},
								},
							},
						},
						"rate_limit_reset_unknown_conn": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pkt_rate_for_reset_unknown_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"log_for_reset_unknown_conn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when rate exceed",
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
			"tcp_stats": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'connattempt': Connect initiated; 'connects': Connect established; 'drops': Connect dropped; 'conndrops': Embryonic connect dropped; 'closed': Connect closed; 'segstimed': Segs to get RTT; 'rttupdated': Update RTT; 'delack': Delayed acks sent; 'timeoutdrop': Conn dropped in rxmt timeout; 'rexmttimeo': Retransmit timeout; 'persisttimeo': Persist timeout; 'keeptimeo': Keepalive timeout; 'keepprobe': Keepalive probe sent; 'keepdrops': Connect dropped in keepalive; 'sndtotal': Total packet sent; 'sndpack': Data packet sent; 'sndbyte': Data bytes sent; 'sndrexmitpack': Data packet retransmit; 'sndrexmitbyte': Data byte retransmit; 'sndrexmitbad': Unnecessary packet retransmit; 'sndacks': Ack packet sent; 'sndprobe': Window probe sent; 'sndurg': URG packet sent; 'sndwinup': Window update packet sent; 'sndctrl': SYN|FIN|RST packet sent; 'sndrst': RST packet sent; 'sndfin': FIN packet sent; 'sndsyn': SYN packet sent; 'rcvtotal': Total packet received; 'rcvpack': Packet received; 'rcvbyte': Bytes received; 'rcvbadoff': Packet received with bad offset; 'rcvmemdrop': Packet dropped for lack of memory; 'rcvduppack': Duplicate packet received; 'rcvdupbyte': Duplicate bytes received; 'rcvpartduppack': Packet with some duplicate data; 'rcvpartdupbyte': Dup. bytes in part-dup. packets; 'rcvoopack': Out-of-order packet received; 'rcvoobyte': Out-of-order bytes received; 'rcvpackafterwin': Packets with data after window; 'rcvbyteafterwin': Bytes rcvd after window; 'rcvwinprobe': Rcvd window probe packet; 'rcvdupack': Rcvd duplicate acks; 'rcvacktoomuch': Rcvd acks for unsent data; 'rcvackpack': Rcvd ack packets; 'rcvackbyte': Bytes acked by rcvd acks; 'rcvwinupd': Rcvd window update packets; 'pawsdrop': Segments dropped due to PAWS; 'predack': Hdr predict for acks; 'preddat': Hdr predict for data pkts; 'persistdrop': Timeout in persist state; 'badrst': Ignored RST; 'finwait2_drops': Drop FIN_WAIT_2 connection after time limit; 'sack_recovery_episode': SACK recovery episodes; 'sack_rexmits': SACK rexmit segments; 'sack_rexmit_bytes': SACK rexmit bytes; 'sack_rcv_blocks': SACK received; 'sack_send_blocks': SACK sent; 'sndcack': Challenge ACK sent; 'cacklim': Challenge ACK limited; 'reassmemdrop': Packet dropped during reassembly; 'reasstimeout': Reassembly Time Out; 'cc_idle': Congestion control window set do to idle; 'cc_reduce': Congestion control window reduced by event; 'rcvdsack': Rcvd DSACK packets; 'a2brcvwnd': ATCP to BTCP receive window; 'a2bsackpresent': ATCP to BTCP SACK options present; 'a2bdupack': ATCP to BTCP Dup/OO ACK; 'a2brxdata': ATCP to BTCP Rxmitted data; 'a2btcpoptions': ATCP to BTCP unsupported TCP options; 'a2boodata': ATCP to BTCP oo data received; 'a2bpartialack': ATCP to BTCP partial ack received; 'a2bfsmtransition': ATCP to BTCP state machine transition; 'a2btransitionnum': ATCP to BTCP total transitions; 'b2atransitionnum': ATCP to BTCP total transitions; 'bad_iochan': IO Channel Modified; 'atcpforward': Adaptive TCP forward; 'atcpsent': Adaptive TCP sent; 'atcprexmitsadrop': Adaptive TCP transmit SA drops; 'atcpsendbackack': Adaptive TCP sendback ACK; 'atcprexmit': Adaptive TCP retransmits; 'atcpbuffallocfail': Adaptive TCP buffer allocation fails; 'a2bappbuffering': Transition to full stack on when application buffers too much data; 'atcpsendfail': Adaptive TCP sent fails; 'earlyrexmit': Early Retransmission sent; 'mburstlim': Maxburst limited tx; 'a2bsndwnd': ATCP to BTCP send window; 'proxyheaderv1': Proxy header v1; 'proxyheaderv2': Proxy header v2;",
									},
								},
							},
						},
					},
				},
			},
			"tcp_syn_per_sec": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"telemetry_log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top_k_source_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"top_k_app_svc_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"device_status": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"environment": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"partition_metrics": {
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
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_policy": {
							Type: schema.TypeString, Optional: true, Description: "Apply policy template to the whole system (Policy template name)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"template_bind": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_monitor": {
										Type: schema.TypeInt, Required: true, Description: "Monitor template ID Number",
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
			"throughput": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'global-system-throughput-bits-per-sec': Global System throughput in bits/sec; 'per-part-throughput-bits-per-sec': Partition throughput in bits/sec;",
									},
								},
							},
						},
					},
				},
			},
			"timeout_value": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ftp": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "set timeout to stop ftp transfer in seconds, 0 is no limit",
						},
						"scp": {
							Type: schema.TypeInt, Optional: true, Default: 300, Description: "set timeout to stop scp transfer in seconds, 0 is no limit",
						},
						"sftp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop sftp transfer in seconds, 0 is no limit",
						},
						"tftp": {
							Type: schema.TypeInt, Optional: true, Default: 300, Description: "set timeout to stop tftp transfer in seconds, 0 is no limit",
						},
						"http": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "set timeout to stop http transfer in seconds",
						},
						"https": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "set timeout to stop https transfer in seconds",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trunk": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balance": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"use_l3": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer-3 Header based load balancing",
									},
									"use_l4": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer-3/4 Header based load balancing",
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
			"trunk_hw_hash": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Set HW hash mode, default is 6 (1:dst-mac 2:src-mac 3:src-dst-mac 4:src-ip 5:dst-ip 6:rtag6 7:rtag7)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trunk_xaui_hw_hash": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Set HW hash mode, default is 6 (1:dst-mac 2:src-mac 3:src-dst-mac 4:src-ip 5:dst-ip 6:rtag6 7:rtag7)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tso": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable TCP Segmentation Offload",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable TCP Segmentation Offload",
						},
					},
				},
			},
			"udp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"skip_checksum_when_zero": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "When enabled, the system does not modify the L4 checksum if the original passthrough packet was received with UDP checksum valu",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"upgrade_status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve_mac_scheme": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ve_mac_scheme_val": {
							Type: schema.TypeString, Optional: true, Default: "hash-based", Description: "'hash-based': Hash-based using the VE number; 'round-robin': Round Robin scheme; 'system-mac': Use system MAC address;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"xaui_dlb_mode": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable Dynamic Load Balancing traffic distribution support",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemAddCpuCore1755(d []interface{}) edpt.SystemAddCpuCore1755 {

	count1 := len(d)
	var ret edpt.SystemAddCpuCore1755
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CoreIndex = in["core_index"].(int)
	}
	return ret
}

func getObjectSystemAddPort1756(d []interface{}) edpt.SystemAddPort1756 {

	count1 := len(d)
	var ret edpt.SystemAddPort1756
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
	}
	return ret
}

func getObjectSystemAllVlanLimit1757(d []interface{}) edpt.SystemAllVlanLimit1757 {

	count1 := len(d)
	var ret edpt.SystemAllVlanLimit1757
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bcast = in["bcast"].(int)
		ret.Ipmcast = in["ipmcast"].(int)
		ret.Mcast = in["mcast"].(int)
		ret.UnknownUcast = in["unknown_ucast"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemAppPerformance1758(d []interface{}) edpt.SystemAppPerformance1758 {

	count1 := len(d)
	var ret edpt.SystemAppPerformance1758
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemAppPerformanceSamplingEnable1759(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemAppPerformanceSamplingEnable1759(d []interface{}) []edpt.SystemAppPerformanceSamplingEnable1759 {

	count1 := len(d)
	ret := make([]edpt.SystemAppPerformanceSamplingEnable1759, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemAppPerformanceSamplingEnable1759
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemAppsGlobal1760(d []interface{}) edpt.SystemAppsGlobal1760 {

	count1 := len(d)
	var ret edpt.SystemAppsGlobal1760
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogSessionOnEstablished = in["log_session_on_established"].(int)
		ret.MslTime = in["msl_time"].(int)
		ret.TimerWheelWalkLimit = in["timer_wheel_walk_limit"].(int)
		ret.SessionsThreshold = in["sessions_threshold"].(int)
		ret.CpsThreshold = in["cps_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemAsicDebugDump1761(d []interface{}) edpt.SystemAsicDebugDump1761 {

	count1 := len(d)
	var ret edpt.SystemAsicDebugDump1761
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemAsicMmuFailSafe1762(d []interface{}) edpt.SystemAsicMmuFailSafe1762 {

	count1 := len(d)
	var ret edpt.SystemAsicMmuFailSafe1762
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RecoveryThreshold = in["recovery_threshold"].(int)
		ret.MonitorInterval = in["monitor_interval"].(int)
		ret.MonitorDisable = in["monitor_disable"].(int)
		ret.RebootDisable = in["reboot_disable"].(int)
		ret.InjectError = in["inject_error"].(int)
		ret.TestPatternType = in["test_pattern_type"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemBandwidth1763(d []interface{}) edpt.SystemBandwidth1763 {

	count1 := len(d)
	var ret edpt.SystemBandwidth1763
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WarningThreshold = in["warning_threshold"].(int)
		ret.CriticalThreshold = in["critical_threshold"].(int)
		//omit uuid
		ret.SamplingEnable = getSliceSystemBandwidthSamplingEnable1764(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemBandwidthSamplingEnable1764(d []interface{}) []edpt.SystemBandwidthSamplingEnable1764 {

	count1 := len(d)
	ret := make([]edpt.SystemBandwidthSamplingEnable1764, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemBandwidthSamplingEnable1764
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemBfd1765(d []interface{}) edpt.SystemBfd1765 {

	count1 := len(d)
	var ret edpt.SystemBfd1765
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemBfdSamplingEnable1766(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemBfdSamplingEnable1766(d []interface{}) []edpt.SystemBfdSamplingEnable1766 {

	count1 := len(d)
	ret := make([]edpt.SystemBfdSamplingEnable1766, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemBfdSamplingEnable1766
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemClThreatCategory1767(d []interface{}) edpt.SystemClThreatCategory1767 {

	var ret edpt.SystemClThreatCategory1767
	return ret
}

func getObjectSystemCliMonitorInterval1768(d []interface{}) edpt.SystemCliMonitorInterval1768 {

	count1 := len(d)
	var ret edpt.SystemCliMonitorInterval1768
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemCmUpdateFileNameRef1769(d []interface{}) edpt.SystemCmUpdateFileNameRef1769 {

	count1 := len(d)
	var ret edpt.SystemCmUpdateFileNameRef1769
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Source_name = in["source_name"].(string)
		ret.Dest_name = in["dest_name"].(string)
		ret.Id1 = in["id1"].(int)
	}
	return ret
}

func getObjectSystemConfigMgmt1770(d []interface{}) edpt.SystemConfigMgmt1770 {

	count1 := len(d)
	var ret edpt.SystemConfigMgmt1770
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DeleteReferencedTaggedObjects = in["delete_referenced_tagged_objects"].(string)
		//omit uuid
		ret.PuSyncDetection = getObjectSystemConfigMgmtPuSyncDetection1771(in["pu_sync_detection"].([]interface{}))
		ret.Mpm = getObjectSystemConfigMgmtMpm1772(in["mpm"].([]interface{}))
		ret.Notification = getObjectSystemConfigMgmtNotification1773(in["notification"].([]interface{}))
	}
	return ret
}

func getObjectSystemConfigMgmtPuSyncDetection1771(d []interface{}) edpt.SystemConfigMgmtPuSyncDetection1771 {

	count1 := len(d)
	var ret edpt.SystemConfigMgmtPuSyncDetection1771
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemConfigMgmtMpm1772(d []interface{}) edpt.SystemConfigMgmtMpm1772 {

	count1 := len(d)
	var ret edpt.SystemConfigMgmtMpm1772
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxWorkers = in["max_workers"].(int)
		ret.MinIdleWorkers = in["min_idle_workers"].(int)
		ret.StartWorkers = in["start_workers"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemConfigMgmtNotification1773(d []interface{}) edpt.SystemConfigMgmtNotification1773 {

	count1 := len(d)
	var ret edpt.SystemConfigMgmtNotification1773
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Period = in["period"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemControlCpu1774(d []interface{}) edpt.SystemControlCpu1774 {

	var ret edpt.SystemControlCpu1774
	return ret
}

func getObjectSystemCore1775(d []interface{}) edpt.SystemCore1775 {

	var ret edpt.SystemCore1775
	return ret
}

func getObjectSystemCosqShow1776(d []interface{}) edpt.SystemCosqShow1776 {

	var ret edpt.SystemCosqShow1776
	return ret
}

func getObjectSystemCosqStats1777(d []interface{}) edpt.SystemCosqStats1777 {

	var ret edpt.SystemCosqStats1777
	return ret
}

func getObjectSystemCounterLibAccounting1778(d []interface{}) edpt.SystemCounterLibAccounting1778 {

	var ret edpt.SystemCounterLibAccounting1778
	return ret
}

func getObjectSystemCpuHyperThread1779(d []interface{}) edpt.SystemCpuHyperThread1779 {

	count1 := len(d)
	var ret edpt.SystemCpuHyperThread1779
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemCpuList1780(d []interface{}) edpt.SystemCpuList1780 {

	var ret edpt.SystemCpuList1780
	return ret
}

func getObjectSystemCpuLoadSharing1781(d []interface{}) edpt.SystemCpuLoadSharing1781 {

	count1 := len(d)
	var ret edpt.SystemCpuLoadSharing1781
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.PacketsPerSecond = getObjectSystemCpuLoadSharingPacketsPerSecond1782(in["packets_per_second"].([]interface{}))
		ret.CpuUsage = getObjectSystemCpuLoadSharingCpuUsage1783(in["cpu_usage"].([]interface{}))
		ret.AllowL7Sessions = in["allow_l7_sessions"].(int)
		ret.Tcp = in["tcp"].(int)
		ret.Udp = in["udp"].(int)
		ret.Others = in["others"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemCpuLoadSharingPacketsPerSecond1782(d []interface{}) edpt.SystemCpuLoadSharingPacketsPerSecond1782 {

	count1 := len(d)
	var ret edpt.SystemCpuLoadSharingPacketsPerSecond1782
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Min = in["min"].(int)
	}
	return ret
}

func getObjectSystemCpuLoadSharingCpuUsage1783(d []interface{}) edpt.SystemCpuLoadSharingCpuUsage1783 {

	count1 := len(d)
	var ret edpt.SystemCpuLoadSharingCpuUsage1783
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Low = in["low"].(int)
		ret.High = in["high"].(int)
	}
	return ret
}

func getObjectSystemCpuMap1784(d []interface{}) edpt.SystemCpuMap1784 {

	var ret edpt.SystemCpuMap1784
	return ret
}

func getObjectSystemCpuPacketPrioSupport1785(d []interface{}) edpt.SystemCpuPacketPrioSupport1785 {

	count1 := len(d)
	var ret edpt.SystemCpuPacketPrioSupport1785
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemDataCpu1786(d []interface{}) edpt.SystemDataCpu1786 {

	var ret edpt.SystemDataCpu1786
	return ret
}

func getObjectSystemDelPort1787(d []interface{}) edpt.SystemDelPort1787 {

	count1 := len(d)
	var ret edpt.SystemDelPort1787
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
	}
	return ret
}

func getObjectSystemDeleteCpuCore1788(d []interface{}) edpt.SystemDeleteCpuCore1788 {

	count1 := len(d)
	var ret edpt.SystemDeleteCpuCore1788
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CoreIndex = in["core_index"].(int)
	}
	return ret
}

func getObjectSystemDns1789(d []interface{}) edpt.SystemDns1789 {

	count1 := len(d)
	var ret edpt.SystemDns1789
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemDnsSamplingEnable1790(in["sampling_enable"].([]interface{}))
		ret.RecursiveNameserver = getObjectSystemDnsRecursiveNameserver1791(in["recursive_nameserver"].([]interface{}))
	}
	return ret
}

func getSliceSystemDnsSamplingEnable1790(d []interface{}) []edpt.SystemDnsSamplingEnable1790 {

	count1 := len(d)
	ret := make([]edpt.SystemDnsSamplingEnable1790, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsSamplingEnable1790
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemDnsRecursiveNameserver1791(d []interface{}) edpt.SystemDnsRecursiveNameserver1791 {

	count1 := len(d)
	var ret edpt.SystemDnsRecursiveNameserver1791
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		ret.ServerList = getSliceSystemDnsRecursiveNameserverServerList1792(in["server_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemDnsRecursiveNameserverServerList1792(d []interface{}) []edpt.SystemDnsRecursiveNameserverServerList1792 {

	count1 := len(d)
	ret := make([]edpt.SystemDnsRecursiveNameserverServerList1792, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsRecursiveNameserverServerList1792
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.V4Desc = in["v4_desc"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.V6Desc = in["v6_desc"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemDnsCache1793(d []interface{}) edpt.SystemDnsCache1793 {

	count1 := len(d)
	var ret edpt.SystemDnsCache1793
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemDnsCacheSamplingEnable1794(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemDnsCacheSamplingEnable1794(d []interface{}) []edpt.SystemDnsCacheSamplingEnable1794 {

	count1 := len(d)
	ret := make([]edpt.SystemDnsCacheSamplingEnable1794, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsCacheSamplingEnable1794
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemDomainListInfo1795(d []interface{}) edpt.SystemDomainListInfo1795 {

	var ret edpt.SystemDomainListInfo1795
	return ret
}

func getObjectSystemDomainListSettings1796(d []interface{}) edpt.SystemDomainListSettings1796 {

	count1 := len(d)
	var ret edpt.SystemDomainListSettings1796
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PollingInterval = in["polling_interval"].(string)
		ret.ConcurrentTask = in["concurrent_task"].(int)
		ret.DomainListPerGroup = in["domain_list_per_group"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemDpdkStats1797(d []interface{}) edpt.SystemDpdkStats1797 {

	count1 := len(d)
	var ret edpt.SystemDpdkStats1797
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemDpdkStatsSamplingEnable1798(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemDpdkStatsSamplingEnable1798(d []interface{}) []edpt.SystemDpdkStatsSamplingEnable1798 {

	count1 := len(d)
	ret := make([]edpt.SystemDpdkStatsSamplingEnable1798, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDpdkStatsSamplingEnable1798
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemEnableDiskEncryption1799(d []interface{}) edpt.SystemEnableDiskEncryption1799 {

	count1 := len(d)
	var ret edpt.SystemEnableDiskEncryption1799
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cipher = in["cipher"].(string)
		ret.Passphrase = in["passphrase"].(string)
		ret.PassphraseBase64 = in["passphrase_base64"].(string)
	}
	return ret
}

func getObjectSystemEnablePassword1800(d []interface{}) edpt.SystemEnablePassword1800 {

	count1 := len(d)
	var ret edpt.SystemEnablePassword1800
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowPasswordPolicy = in["follow_password_policy"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemEnvironment1801(d []interface{}) edpt.SystemEnvironment1801 {

	var ret edpt.SystemEnvironment1801
	return ret
}

func getObjectSystemExtOnlyLogging1802(d []interface{}) edpt.SystemExtOnlyLogging1802 {

	count1 := len(d)
	var ret edpt.SystemExtOnlyLogging1802
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSystemForcedGroupSpeedList(d []interface{}) []edpt.SystemForcedGroupSpeedList {

	count1 := len(d)
	ret := make([]edpt.SystemForcedGroupSpeedList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemForcedGroupSpeedList
		oi.Eth01_to_04 = in["eth01_to_04"].(int)
		oi.Eth05_to_08 = in["eth05_to_08"].(int)
		oi.Eth09_to_12 = in["eth09_to_12"].(int)
		oi.Eth13_to_16 = in["eth13_to_16"].(int)
		oi.Eth17_to_20 = in["eth17_to_20"].(int)
		oi.Eth21_to_24 = in["eth21_to_24"].(int)
		oi.Speed = in["speed"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemFpgaCoreCrc1803(d []interface{}) edpt.SystemFpgaCoreCrc1803 {

	count1 := len(d)
	var ret edpt.SystemFpgaCoreCrc1803
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonitorDisable = in["monitor_disable"].(int)
		ret.RebootEnable = in["reboot_enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemFpgaDrop1804(d []interface{}) edpt.SystemFpgaDrop1804 {

	count1 := len(d)
	var ret edpt.SystemFpgaDrop1804
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemFpgaDropSamplingEnable1805(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemFpgaDropSamplingEnable1805(d []interface{}) []edpt.SystemFpgaDropSamplingEnable1805 {

	count1 := len(d)
	ret := make([]edpt.SystemFpgaDropSamplingEnable1805, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemFpgaDropSamplingEnable1805
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemFw1806(d []interface{}) edpt.SystemFw1806 {

	count1 := len(d)
	var ret edpt.SystemFw1806
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ApplicationMempool = in["application_mempool"].(int)
		ret.ApplicationFlow = in["application_flow"].(int)
		ret.BasicDpiEnable = in["basic_dpi_enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemGeoLocation1807(d []interface{}) edpt.SystemGeoLocation1807 {

	count1 := len(d)
	var ret edpt.SystemGeoLocation1807
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeoLocationIana = in["geo_location_iana"].(int)
		ret.GeoLocationIanaSystem = in["geo_location_iana_system"].(int)
		ret.GeoLocationGeolite2Asn = in["geo_location_geolite2_asn"].(int)
		ret.Geolite2AsnIncludeIpv6 = in["geolite2_asn_include_ipv6"].(int)
		ret.GeoLocationGeolite2City = in["geo_location_geolite2_city"].(int)
		ret.Geolite2CityIncludeIpv6 = in["geolite2_city_include_ipv6"].(int)
		ret.GeoLocationGeolite2Country = in["geo_location_geolite2_country"].(int)
		ret.Geolite2CountryIncludeIpv6 = in["geolite2_country_include_ipv6"].(int)
		ret.GeolocLoadFileList = getSliceSystemGeoLocationGeolocLoadFileList1808(in["geoloc_load_file_list"].([]interface{}))
		//omit uuid
		ret.EntryList = getSliceSystemGeoLocationEntryList1809(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeoLocationGeolocLoadFileList1808(d []interface{}) []edpt.SystemGeoLocationGeolocLoadFileList1808 {

	count1 := len(d)
	ret := make([]edpt.SystemGeoLocationGeolocLoadFileList1808, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationGeolocLoadFileList1808
		oi.GeoLocationLoadFilename = in["geo_location_load_filename"].(string)
		oi.GeoLocationLoadFileIncludeIpv6 = in["geo_location_load_file_include_ipv6"].(int)
		oi.TemplateName = in["template_name"].(string)
		oi.GeoLocationLoadTempIncludeIpv6 = in["geo_location_load_temp_include_ipv6"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeoLocationEntryList1809(d []interface{}) []edpt.SystemGeoLocationEntryList1809 {

	count1 := len(d)
	ret := make([]edpt.SystemGeoLocationEntryList1809, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationEntryList1809
		oi.GeoLocnObjName = in["geo_locn_obj_name"].(string)
		oi.GeoLocnMultipleAddresses = getSliceSystemGeoLocationEntryListGeoLocnMultipleAddresses1810(in["geo_locn_multiple_addresses"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeoLocationEntryListGeoLocnMultipleAddresses1810(d []interface{}) []edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses1810 {

	count1 := len(d)
	ret := make([]edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses1810, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses1810
		oi.FirstIpAddress = in["first_ip_address"].(string)
		oi.GeolIpv4Mask = in["geol_ipv4_mask"].(string)
		oi.IpAddr2 = in["ip_addr2"].(string)
		oi.FirstIpv6Address = in["first_ipv6_address"].(string)
		oi.GeolIpv6Mask = in["geol_ipv6_mask"].(int)
		oi.Ipv6Addr2 = in["ipv6_addr2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemGeoloc1811(d []interface{}) edpt.SystemGeoloc1811 {

	count1 := len(d)
	var ret edpt.SystemGeoloc1811
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemGeolocSamplingEnable1812(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeolocSamplingEnable1812(d []interface{}) []edpt.SystemGeolocSamplingEnable1812 {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocSamplingEnable1812, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocSamplingEnable1812
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListList(d []interface{}) []edpt.SystemGeolocListList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListList
		oi.Name = in["name"].(string)
		oi.Shared = in["shared"].(int)
		oi.IncludeGeolocNameList = getSliceSystemGeolocListListIncludeGeolocNameList(in["include_geoloc_name_list"].([]interface{}))
		oi.ExcludeGeolocNameList = getSliceSystemGeolocListListExcludeGeolocNameList(in["exclude_geoloc_name_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSystemGeolocListListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListListIncludeGeolocNameList(d []interface{}) []edpt.SystemGeolocListListIncludeGeolocNameList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListListIncludeGeolocNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListListIncludeGeolocNameList
		oi.IncludeGeolocNameVal = in["include_geoloc_name_val"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListListExcludeGeolocNameList(d []interface{}) []edpt.SystemGeolocListListExcludeGeolocNameList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListListExcludeGeolocNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListListExcludeGeolocNameList
		oi.ExcludeGeolocNameVal = in["exclude_geoloc_name_val"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListListSamplingEnable(d []interface{}) []edpt.SystemGeolocListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemGeolocNameHelper1813(d []interface{}) edpt.SystemGeolocNameHelper1813 {

	count1 := len(d)
	var ret edpt.SystemGeolocNameHelper1813
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemGeolocNameHelperSamplingEnable1814(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeolocNameHelperSamplingEnable1814(d []interface{}) []edpt.SystemGeolocNameHelperSamplingEnable1814 {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocNameHelperSamplingEnable1814, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocNameHelperSamplingEnable1814
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemGeolocationFile1815(d []interface{}) edpt.SystemGeolocationFile1815 {

	count1 := len(d)
	var ret edpt.SystemGeolocationFile1815
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.ErrorInfo = getObjectSystemGeolocationFileErrorInfo1816(in["error_info"].([]interface{}))
	}
	return ret
}

func getObjectSystemGeolocationFileErrorInfo1816(d []interface{}) edpt.SystemGeolocationFileErrorInfo1816 {

	var ret edpt.SystemGeolocationFileErrorInfo1816
	return ret
}

func getObjectSystemGlid1817(d []interface{}) edpt.SystemGlid1817 {

	count1 := len(d)
	var ret edpt.SystemGlid1817
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GlidId = in["glid_id"].(string)
		ret.NonShared = in["non_shared"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemGuestFile1818(d []interface{}) edpt.SystemGuestFile1818 {

	var ret edpt.SystemGuestFile1818
	return ret
}

func getObjectSystemGuiImageList1819(d []interface{}) edpt.SystemGuiImageList1819 {

	var ret edpt.SystemGuiImageList1819
	return ret
}

func getObjectSystemHardware1820(d []interface{}) edpt.SystemHardware1820 {

	var ret edpt.SystemHardware1820
	return ret
}

func getObjectSystemHardwareAccelerate1821(d []interface{}) edpt.SystemHardwareAccelerate1821 {

	count1 := len(d)
	var ret edpt.SystemHardwareAccelerate1821
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionForwarding = in["session_forwarding"].(int)
		//omit uuid
		ret.SamplingEnable = getSliceSystemHardwareAccelerateSamplingEnable1822(in["sampling_enable"].([]interface{}))
		ret.Slb = getObjectSystemHardwareAccelerateSlb1823(in["slb"].([]interface{}))
	}
	return ret
}

func getSliceSystemHardwareAccelerateSamplingEnable1822(d []interface{}) []edpt.SystemHardwareAccelerateSamplingEnable1822 {

	count1 := len(d)
	ret := make([]edpt.SystemHardwareAccelerateSamplingEnable1822, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareAccelerateSamplingEnable1822
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemHardwareAccelerateSlb1823(d []interface{}) edpt.SystemHardwareAccelerateSlb1823 {

	count1 := len(d)
	var ret edpt.SystemHardwareAccelerateSlb1823
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemHardwareAccelerateSlbSamplingEnable1824(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemHardwareAccelerateSlbSamplingEnable1824(d []interface{}) []edpt.SystemHardwareAccelerateSlbSamplingEnable1824 {

	count1 := len(d)
	ret := make([]edpt.SystemHardwareAccelerateSlbSamplingEnable1824, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareAccelerateSlbSamplingEnable1824
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemHealthCheckList(d []interface{}) []edpt.SystemHealthCheckList {

	count1 := len(d)
	ret := make([]edpt.SystemHealthCheckList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHealthCheckList
		oi.L2hmHcName = in["l2hm_hc_name"].(string)
		oi.MethodL2bfd = in["method_l2bfd"].(int)
		oi.L2bfdTxInterval = in["l2bfd_tx_interval"].(int)
		oi.L2bfdRxInterval = in["l2bfd_rx_interval"].(int)
		oi.L2bfdMultiplier = in["l2bfd_multiplier"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemHighMemoryL4Session1825(d []interface{}) edpt.SystemHighMemoryL4Session1825 {

	count1 := len(d)
	var ret edpt.SystemHighMemoryL4Session1825
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemHrxqStatus1826(d []interface{}) edpt.SystemHrxqStatus1826 {

	var ret edpt.SystemHrxqStatus1826
	return ret
}

func getObjectSystemIcmp1827(d []interface{}) edpt.SystemIcmp1827 {

	count1 := len(d)
	var ret edpt.SystemIcmp1827
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIcmpSamplingEnable1828(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIcmpSamplingEnable1828(d []interface{}) []edpt.SystemIcmpSamplingEnable1828 {

	count1 := len(d)
	ret := make([]edpt.SystemIcmpSamplingEnable1828, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmpSamplingEnable1828
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIcmpRate1829(d []interface{}) edpt.SystemIcmpRate1829 {

	count1 := len(d)
	var ret edpt.SystemIcmpRate1829
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIcmpRateSamplingEnable1830(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIcmpRateSamplingEnable1830(d []interface{}) []edpt.SystemIcmpRateSamplingEnable1830 {

	count1 := len(d)
	ret := make([]edpt.SystemIcmpRateSamplingEnable1830, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmpRateSamplingEnable1830
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIcmp61831(d []interface{}) edpt.SystemIcmp61831 {

	count1 := len(d)
	var ret edpt.SystemIcmp61831
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIcmp6SamplingEnable1832(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIcmp6SamplingEnable1832(d []interface{}) []edpt.SystemIcmp6SamplingEnable1832 {

	count1 := len(d)
	ret := make([]edpt.SystemIcmp6SamplingEnable1832, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmp6SamplingEnable1832
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemInuseCpuList1833(d []interface{}) edpt.SystemInuseCpuList1833 {

	var ret edpt.SystemInuseCpuList1833
	return ret
}

func getObjectSystemInusePortList1834(d []interface{}) edpt.SystemInusePortList1834 {

	var ret edpt.SystemInusePortList1834
	return ret
}

func getObjectSystemIoCpu1835(d []interface{}) edpt.SystemIoCpu1835 {

	count1 := len(d)
	var ret edpt.SystemIoCpu1835
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxCores = in["max_cores"].(int)
	}
	return ret
}

func getObjectSystemIp1836(d []interface{}) edpt.SystemIp1836 {

	count1 := len(d)
	var ret edpt.SystemIp1836
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpUnreachableDisable = in["icmp_unreachable_disable"].(int)
		ret.IcmpRedirectDisable = in["icmp_redirect_disable"].(int)
		ret.RpfCheckEnable = in["rpf_check_enable"].(int)
		ret.SourceRoutePktDropEnable = in["source_route_pkt_drop_enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemIpDnsCache1837(d []interface{}) edpt.SystemIpDnsCache1837 {

	var ret edpt.SystemIpDnsCache1837
	return ret
}

func getObjectSystemIpStats1838(d []interface{}) edpt.SystemIpStats1838 {

	count1 := len(d)
	var ret edpt.SystemIpStats1838
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIpStatsSamplingEnable1839(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIpStatsSamplingEnable1839(d []interface{}) []edpt.SystemIpStatsSamplingEnable1839 {

	count1 := len(d)
	ret := make([]edpt.SystemIpStatsSamplingEnable1839, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpStatsSamplingEnable1839
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatList1840(d []interface{}) edpt.SystemIpThreatList1840 {

	count1 := len(d)
	var ret edpt.SystemIpThreatList1840
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIpThreatListSamplingEnable1841(in["sampling_enable"].([]interface{}))
		ret.Ipv4SourceList = getObjectSystemIpThreatListIpv4SourceList1842(in["ipv4_source_list"].([]interface{}))
		ret.Ipv4DestList = getObjectSystemIpThreatListIpv4DestList1844(in["ipv4_dest_list"].([]interface{}))
		ret.Ipv6SourceList = getObjectSystemIpThreatListIpv6SourceList1846(in["ipv6_source_list"].([]interface{}))
		ret.Ipv6DestList = getObjectSystemIpThreatListIpv6DestList1848(in["ipv6_dest_list"].([]interface{}))
		ret.Ipv4InternetHostList = getObjectSystemIpThreatListIpv4InternetHostList1850(in["ipv4_internet_host_list"].([]interface{}))
		ret.Ipv6InternetHostList = getObjectSystemIpThreatListIpv6InternetHostList1852(in["ipv6_internet_host_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemIpThreatListSamplingEnable1841(d []interface{}) []edpt.SystemIpThreatListSamplingEnable1841 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListSamplingEnable1841, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListSamplingEnable1841
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4SourceList1842(d []interface{}) edpt.SystemIpThreatListIpv4SourceList1842 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4SourceList1842
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4SourceListClassListCfg1843(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4SourceListClassListCfg1843(d []interface{}) []edpt.SystemIpThreatListIpv4SourceListClassListCfg1843 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4SourceListClassListCfg1843, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4SourceListClassListCfg1843
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4DestList1844(d []interface{}) edpt.SystemIpThreatListIpv4DestList1844 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4DestList1844
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4DestListClassListCfg1845(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4DestListClassListCfg1845(d []interface{}) []edpt.SystemIpThreatListIpv4DestListClassListCfg1845 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4DestListClassListCfg1845, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4DestListClassListCfg1845
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6SourceList1846(d []interface{}) edpt.SystemIpThreatListIpv6SourceList1846 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6SourceList1846
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6SourceListClassListCfg1847(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6SourceListClassListCfg1847(d []interface{}) []edpt.SystemIpThreatListIpv6SourceListClassListCfg1847 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6SourceListClassListCfg1847, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6SourceListClassListCfg1847
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6DestList1848(d []interface{}) edpt.SystemIpThreatListIpv6DestList1848 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6DestList1848
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6DestListClassListCfg1849(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6DestListClassListCfg1849(d []interface{}) []edpt.SystemIpThreatListIpv6DestListClassListCfg1849 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6DestListClassListCfg1849, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6DestListClassListCfg1849
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4InternetHostList1850(d []interface{}) edpt.SystemIpThreatListIpv4InternetHostList1850 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4InternetHostList1850
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WhiteList = in["white_list"].(string)
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4InternetHostListClassListCfg1851(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4InternetHostListClassListCfg1851(d []interface{}) []edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1851 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1851, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1851
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6InternetHostList1852(d []interface{}) edpt.SystemIpThreatListIpv6InternetHostList1852 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6InternetHostList1852
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WhiteList = in["white_list"].(string)
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6InternetHostListClassListCfg1853(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6InternetHostListClassListCfg1853(d []interface{}) []edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1853 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1853, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1853
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIp6Stats1854(d []interface{}) edpt.SystemIp6Stats1854 {

	count1 := len(d)
	var ret edpt.SystemIp6Stats1854
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIp6StatsSamplingEnable1855(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIp6StatsSamplingEnable1855(d []interface{}) []edpt.SystemIp6StatsSamplingEnable1855 {

	count1 := len(d)
	ret := make([]edpt.SystemIp6StatsSamplingEnable1855, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIp6StatsSamplingEnable1855
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpmi1856(d []interface{}) edpt.SystemIpmi1856 {

	count1 := len(d)
	var ret edpt.SystemIpmi1856
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Reset = in["reset"].(int)
		ret.Ip = getObjectSystemIpmiIp1857(in["ip"].([]interface{}))
		ret.Ipsrc = getObjectSystemIpmiIpsrc1858(in["ipsrc"].([]interface{}))
		ret.User = getObjectSystemIpmiUser1859(in["user"].([]interface{}))
		ret.Tool = getObjectSystemIpmiTool1860(in["tool"].([]interface{}))
	}
	return ret
}

func getObjectSystemIpmiIp1857(d []interface{}) edpt.SystemIpmiIp1857 {

	count1 := len(d)
	var ret edpt.SystemIpmiIp1857
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.DefaultGateway = in["default_gateway"].(string)
	}
	return ret
}

func getObjectSystemIpmiIpsrc1858(d []interface{}) edpt.SystemIpmiIpsrc1858 {

	count1 := len(d)
	var ret edpt.SystemIpmiIpsrc1858
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.Static = in["static"].(int)
	}
	return ret
}

func getObjectSystemIpmiUser1859(d []interface{}) edpt.SystemIpmiUser1859 {

	count1 := len(d)
	var ret edpt.SystemIpmiUser1859
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Add = in["add"].(string)
		ret.Password = in["password"].(string)
		ret.Administrator = in["administrator"].(int)
		ret.Callback = in["callback"].(int)
		ret.Operator = in["operator"].(int)
		ret.User = in["user"].(int)
		ret.Disable = in["disable"].(string)
		ret.Privilege = in["privilege"].(string)
		ret.Setname = in["setname"].(string)
		ret.Newname = in["newname"].(string)
		ret.Setpass = in["setpass"].(string)
		ret.Newpass = in["newpass"].(string)
	}
	return ret
}

func getObjectSystemIpmiTool1860(d []interface{}) edpt.SystemIpmiTool1860 {

	count1 := len(d)
	var ret edpt.SystemIpmiTool1860
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cmd = in["cmd"].(string)
	}
	return ret
}

func getObjectSystemIpmiService1861(d []interface{}) edpt.SystemIpmiService1861 {

	count1 := len(d)
	var ret edpt.SystemIpmiService1861
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemIpsec1862(d []interface{}) edpt.SystemIpsec1862 {

	count1 := len(d)
	var ret edpt.SystemIpsec1862
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketRoundRobin = in["packet_round_robin"].(int)
		ret.CryptoCore = in["crypto_core"].(int)
		ret.CryptoMem = in["crypto_mem"].(int)
		ret.Qat = in["qat"].(int)
		//omit uuid
		ret.FpgaDecrypt = getObjectSystemIpsecFpgaDecrypt1863(in["fpga_decrypt"].([]interface{}))
	}
	return ret
}

func getObjectSystemIpsecFpgaDecrypt1863(d []interface{}) edpt.SystemIpsecFpgaDecrypt1863 {

	count1 := len(d)
	var ret edpt.SystemIpsecFpgaDecrypt1863
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
	}
	return ret
}

func getObjectSystemIpv61864(d []interface{}) edpt.SystemIpv61864 {

	count1 := len(d)
	var ret edpt.SystemIpv61864
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Icmpv6UnreachableDisable = in["icmpv6_unreachable_disable"].(int)
		ret.Icmpv6RedirectDisable = in["icmpv6_redirect_disable"].(int)
		ret.RpfCheckEnable = in["rpf_check_enable"].(int)
		ret.SourceRoutePktDropEnable = in["source_route_pkt_drop_enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemJobOffload1865(d []interface{}) edpt.SystemJobOffload1865 {

	count1 := len(d)
	var ret edpt.SystemJobOffload1865
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemJobOffloadSamplingEnable1866(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemJobOffloadSamplingEnable1866(d []interface{}) []edpt.SystemJobOffloadSamplingEnable1866 {

	count1 := len(d)
	ret := make([]edpt.SystemJobOffloadSamplingEnable1866, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemJobOffloadSamplingEnable1866
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemLinkCapability1867(d []interface{}) edpt.SystemLinkCapability1867 {

	count1 := len(d)
	var ret edpt.SystemLinkCapability1867
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemLinkMonitor1868(d []interface{}) edpt.SystemLinkMonitor1868 {

	count1 := len(d)
	var ret edpt.SystemLinkMonitor1868
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemLro1869(d []interface{}) edpt.SystemLro1869 {

	count1 := len(d)
	var ret edpt.SystemLro1869
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemManagementInterfaceMode1870(d []interface{}) edpt.SystemManagementInterfaceMode1870 {

	count1 := len(d)
	var ret edpt.SystemManagementInterfaceMode1870
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dedicated = in["dedicated"].(int)
		ret.NonDedicated = in["non_dedicated"].(int)
	}
	return ret
}

func getObjectSystemMemory1871(d []interface{}) edpt.SystemMemory1871 {

	count1 := len(d)
	var ret edpt.SystemMemory1871
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemMemorySamplingEnable1872(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemMemorySamplingEnable1872(d []interface{}) []edpt.SystemMemorySamplingEnable1872 {

	count1 := len(d)
	ret := make([]edpt.SystemMemorySamplingEnable1872, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemorySamplingEnable1872
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemMemoryBlockDebug1873(d []interface{}) edpt.SystemMemoryBlockDebug1873 {

	count1 := len(d)
	var ret edpt.SystemMemoryBlockDebug1873
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AssertBlock = in["assert_block"].(int)
		ret.PktdumpBlock = in["pktdump_block"].(int)
		ret.FirstBlk = in["first_blk"].(int)
		ret.SecondBlk = in["second_blk"].(int)
		ret.ThirdBlk = in["third_blk"].(int)
		ret.FourthBlk = in["fourth_blk"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemMfaAuth1874(d []interface{}) edpt.SystemMfaAuth1874 {

	count1 := len(d)
	var ret edpt.SystemMfaAuth1874
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Username = in["username"].(string)
		ret.SecondFactor = in["second_factor"].(string)
	}
	return ret
}

func getObjectSystemMfaCertStore1875(d []interface{}) edpt.SystemMfaCertStore1875 {

	count1 := len(d)
	var ret edpt.SystemMfaCertStore1875
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CertHost = in["cert_host"].(string)
		ret.Protocol = in["protocol"].(string)
		ret.CertStorePath = in["cert_store_path"].(string)
		ret.Username = in["username"].(string)
		ret.PasswdString = in["passwd_string"].(string)
		//omit encrypted
		//omit uuid
	}
	return ret
}

func getObjectSystemMfaManagement1876(d []interface{}) edpt.SystemMfaManagement1876 {

	count1 := len(d)
	var ret edpt.SystemMfaManagement1876
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemMfaValidationType1877(d []interface{}) edpt.SystemMfaValidationType1877 {

	count1 := len(d)
	var ret edpt.SystemMfaValidationType1877
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaCert = in["ca_cert"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemMgmtPort1878(d []interface{}) edpt.SystemMgmtPort1878 {

	count1 := len(d)
	var ret edpt.SystemMgmtPort1878
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.MacAddress = in["mac_address"].(string)
		ret.PciAddress = in["pci_address"].(string)
	}
	return ret
}

func getObjectSystemModifyPort1879(d []interface{}) edpt.SystemModifyPort1879 {

	count1 := len(d)
	var ret edpt.SystemModifyPort1879
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.PortNumber = in["port_number"].(int)
	}
	return ret
}

func getObjectSystemMonTemplate1880(d []interface{}) edpt.SystemMonTemplate1880 {

	count1 := len(d)
	var ret edpt.SystemMonTemplate1880
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonitorList = getSliceSystemMonTemplateMonitorList(in["monitor_list"].([]interface{}))
		ret.LinkBlockAsDown = getObjectSystemMonTemplateLinkBlockAsDown1881(in["link_block_as_down"].([]interface{}))
		ret.LinkDownOnRestart = getObjectSystemMonTemplateLinkDownOnRestart1882(in["link_down_on_restart"].([]interface{}))
		ret.MonitoringMode = getObjectSystemMonTemplateMonitoringMode1883(in["monitoring_mode"].([]interface{}))
	}
	return ret
}

func getSliceSystemMonTemplateMonitorList(d []interface{}) []edpt.SystemMonTemplateMonitorList {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorList
		oi.Id1 = in["id1"].(int)
		oi.ClearCfg = getSliceSystemMonTemplateMonitorListClearCfg(in["clear_cfg"].([]interface{}))
		oi.LinkDisableCfg = getSliceSystemMonTemplateMonitorListLinkDisableCfg(in["link_disable_cfg"].([]interface{}))
		oi.LinkEnableCfg = getSliceSystemMonTemplateMonitorListLinkEnableCfg(in["link_enable_cfg"].([]interface{}))
		oi.MonitorRelation = in["monitor_relation"].(string)
		oi.LinkUpCfg = getSliceSystemMonTemplateMonitorListLinkUpCfg(in["link_up_cfg"].([]interface{}))
		oi.LinkDownCfg = getSliceSystemMonTemplateMonitorListLinkDownCfg(in["link_down_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListClearCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListClearCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListClearCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListClearCfg
		oi.Sessions = in["sessions"].(string)
		oi.ClearAllSequence = in["clear_all_sequence"].(int)
		oi.ClearAllPartition = in["clear_all_partition"].(string)
		oi.ClearAllPartitionAll = in["clear_all_partition_all"].(int)
		oi.ClearSequence = in["clear_sequence"].(int)
		oi.ClearPartition = in["clear_partition"].(string)
		oi.ClearPartitionAll = in["clear_partition_all"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListLinkDisableCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListLinkDisableCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkDisableCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListLinkDisableCfg
		oi.Diseth = in["diseth"].(int)
		oi.DisSequence = in["dis_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListLinkEnableCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListLinkEnableCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkEnableCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListLinkEnableCfg
		oi.Enaeth = in["enaeth"].(int)
		oi.EnaSequence = in["ena_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListLinkUpCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListLinkUpCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkUpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListLinkUpCfg
		oi.LinkupEthernet1 = in["linkup_ethernet1"].(int)
		oi.LinkUpSequence1 = in["link_up_sequence1"].(int)
		oi.LinkupEthernet2 = in["linkup_ethernet2"].(int)
		oi.LinkUpSequence2 = in["link_up_sequence2"].(int)
		oi.LinkupEthernet3 = in["linkup_ethernet3"].(int)
		oi.LinkUpSequence3 = in["link_up_sequence3"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListLinkDownCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListLinkDownCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkDownCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListLinkDownCfg
		oi.LinkdownEthernet1 = in["linkdown_ethernet1"].(int)
		oi.LinkDownSequence1 = in["link_down_sequence1"].(int)
		oi.LinkdownEthernet2 = in["linkdown_ethernet2"].(int)
		oi.LinkDownSequence2 = in["link_down_sequence2"].(int)
		oi.LinkdownEthernet3 = in["linkdown_ethernet3"].(int)
		oi.LinkDownSequence3 = in["link_down_sequence3"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemMonTemplateLinkBlockAsDown1881(d []interface{}) edpt.SystemMonTemplateLinkBlockAsDown1881 {

	count1 := len(d)
	var ret edpt.SystemMonTemplateLinkBlockAsDown1881
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemMonTemplateLinkDownOnRestart1882(d []interface{}) edpt.SystemMonTemplateLinkDownOnRestart1882 {

	count1 := len(d)
	var ret edpt.SystemMonTemplateLinkDownOnRestart1882
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemMonTemplateMonitoringMode1883(d []interface{}) edpt.SystemMonTemplateMonitoringMode1883 {

	count1 := len(d)
	var ret edpt.SystemMonTemplateMonitoringMode1883
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mmode = in["mmode"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemMultiQueueSupport1884(d []interface{}) edpt.SystemMultiQueueSupport1884 {

	count1 := len(d)
	var ret edpt.SystemMultiQueueSupport1884
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
	}
	return ret
}

func getObjectSystemNdiscRa1885(d []interface{}) edpt.SystemNdiscRa1885 {

	count1 := len(d)
	var ret edpt.SystemNdiscRa1885
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemNdiscRaSamplingEnable1886(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemNdiscRaSamplingEnable1886(d []interface{}) []edpt.SystemNdiscRaSamplingEnable1886 {

	count1 := len(d)
	ret := make([]edpt.SystemNdiscRaSamplingEnable1886, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemNdiscRaSamplingEnable1886
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemNetvscMonitor1887(d []interface{}) edpt.SystemNetvscMonitor1887 {

	count1 := len(d)
	var ret edpt.SystemNetvscMonitor1887
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemNsmA10lb1888(d []interface{}) edpt.SystemNsmA10lb1888 {

	count1 := len(d)
	var ret edpt.SystemNsmA10lb1888
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Kill = in["kill"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemPasswordPolicy1889(d []interface{}) edpt.SystemPasswordPolicy1889 {

	count1 := len(d)
	var ret edpt.SystemPasswordPolicy1889
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Complexity = in["complexity"].(string)
		ret.Aging = in["aging"].(string)
		ret.History = in["history"].(string)
		ret.MinPswdLen = in["min_pswd_len"].(int)
		ret.UsernameCheck = in["username_check"].(string)
		ret.RepeatCharacterCheck = in["repeat_character_check"].(string)
		ret.ForbidConsecutiveCharacter = in["forbid_consecutive_character"].(string)
		//omit uuid
	}
	return ret
}

func getSliceSystemPathList(d []interface{}) []edpt.SystemPathList {

	count1 := len(d)
	ret := make([]edpt.SystemPathList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemPathList
		oi.L2hmPathName = in["l2hm_path_name"].(string)
		oi.L2hmVlan = in["l2hm_vlan"].(int)
		oi.L2hmSetupTestApi = in["l2hm_setup_test_api"].(int)
		oi.IfpairEthStart = in["ifpair_eth_start"].(int)
		oi.IfpairEthEnd = in["ifpair_eth_end"].(int)
		oi.IfpairTrunkStart = in["ifpair_trunk_start"].(int)
		oi.IfpairTrunkEnd = in["ifpair_trunk_end"].(int)
		oi.L2hmAttach = in["l2hm_attach"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemPbslb1890(d []interface{}) edpt.SystemPbslb1890 {

	count1 := len(d)
	var ret edpt.SystemPbslb1890
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SockstressDisable = in["sockstress_disable"].(int)
		//omit uuid
		ret.SamplingEnable = getSliceSystemPbslbSamplingEnable1891(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemPbslbSamplingEnable1891(d []interface{}) []edpt.SystemPbslbSamplingEnable1891 {

	count1 := len(d)
	ret := make([]edpt.SystemPbslbSamplingEnable1891, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemPbslbSamplingEnable1891
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemPerVlanLimit1892(d []interface{}) edpt.SystemPerVlanLimit1892 {

	count1 := len(d)
	var ret edpt.SystemPerVlanLimit1892
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bcast = in["bcast"].(int)
		ret.Ipmcast = in["ipmcast"].(int)
		ret.Mcast = in["mcast"].(int)
		ret.UnknownUcast = in["unknown_ucast"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemPlatformtype1893(d []interface{}) edpt.SystemPlatformtype1893 {

	var ret edpt.SystemPlatformtype1893
	return ret
}

func getObjectSystemPortCount1894(d []interface{}) edpt.SystemPortCount1894 {

	count1 := len(d)
	var ret edpt.SystemPortCount1894
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortCountKernel = in["port_count_kernel"].(int)
		ret.PortCountHm = in["port_count_hm"].(int)
		ret.PortCountLogging = in["port_count_logging"].(int)
		ret.PortCountAlg = in["port_count_alg"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemPortInfo1895(d []interface{}) edpt.SystemPortInfo1895 {

	var ret edpt.SystemPortInfo1895
	return ret
}

func getObjectSystemPortList1896(d []interface{}) edpt.SystemPortList1896 {

	var ret edpt.SystemPortList1896
	return ret
}

func getObjectSystemPorts1897(d []interface{}) edpt.SystemPorts1897 {

	count1 := len(d)
	var ret edpt.SystemPorts1897
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LinkDetectionInterval = in["link_detection_interval"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemPowerOnSelfTest1898(d []interface{}) edpt.SystemPowerOnSelfTest1898 {

	var ret edpt.SystemPowerOnSelfTest1898
	return ret
}

func getObjectSystemProbeNetworkDevices1899(d []interface{}) edpt.SystemProbeNetworkDevices1899 {

	var ret edpt.SystemProbeNetworkDevices1899
	return ret
}

func getObjectSystemPsuInfo1900(d []interface{}) edpt.SystemPsuInfo1900 {

	var ret edpt.SystemPsuInfo1900
	return ret
}

func getObjectSystemQInQ1901(d []interface{}) edpt.SystemQInQ1901 {

	count1 := len(d)
	var ret edpt.SystemQInQ1901
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableAllPorts = in["enable_all_ports"].(int)
		ret.InnerTpid = in["inner_tpid"].(string)
		ret.OuterTpid = in["outer_tpid"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemQueuingBuffer1902(d []interface{}) edpt.SystemQueuingBuffer1902 {

	count1 := len(d)
	var ret edpt.SystemQueuingBuffer1902
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemRadius1903(d []interface{}) edpt.SystemRadius1903 {

	count1 := len(d)
	var ret edpt.SystemRadius1903
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server = getObjectSystemRadiusServer1904(in["server"].([]interface{}))
	}
	return ret
}

func getObjectSystemRadiusServer1904(d []interface{}) edpt.SystemRadiusServer1904 {

	count1 := len(d)
	var ret edpt.SystemRadiusServer1904
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListenPort = in["listen_port"].(int)
		ret.Remote = getObjectSystemRadiusServerRemote1905(in["remote"].([]interface{}))
		ret.Secret = in["secret"].(int)
		ret.SecretString = in["secret_string"].(string)
		//omit encrypted
		ret.Vrid = in["vrid"].(int)
		ret.Attribute = getSliceSystemRadiusServerAttribute1907(in["attribute"].([]interface{}))
		ret.DisableReply = in["disable_reply"].(int)
		ret.AccountingStart = in["accounting_start"].(string)
		ret.AccountingStop = in["accounting_stop"].(string)
		ret.AccountingInterimUpdate = in["accounting_interim_update"].(string)
		ret.AccountingOn = in["accounting_on"].(string)
		ret.AttributeName = in["attribute_name"].(string)
		ret.CustomAttributeName = in["custom_attribute_name"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSystemRadiusServerSamplingEnable1908(in["sampling_enable"].([]interface{}))
		ret.DerivedAttribute = getObjectSystemRadiusServerDerivedAttribute1909(in["derived_attribute"].([]interface{}))
	}
	return ret
}

func getObjectSystemRadiusServerRemote1905(d []interface{}) edpt.SystemRadiusServerRemote1905 {

	count1 := len(d)
	var ret edpt.SystemRadiusServerRemote1905
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpList = getSliceSystemRadiusServerRemoteIpList1906(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemRadiusServerRemoteIpList1906(d []interface{}) []edpt.SystemRadiusServerRemoteIpList1906 {

	count1 := len(d)
	ret := make([]edpt.SystemRadiusServerRemoteIpList1906, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerRemoteIpList1906
		oi.IpListName = in["ip_list_name"].(string)
		oi.IpListSecret = in["ip_list_secret"].(int)
		oi.IpListSecretString = in["ip_list_secret_string"].(string)
		//omit ip_list_encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemRadiusServerAttribute1907(d []interface{}) []edpt.SystemRadiusServerAttribute1907 {

	count1 := len(d)
	ret := make([]edpt.SystemRadiusServerAttribute1907, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerAttribute1907
		oi.AttributeValue = in["attribute_value"].(string)
		oi.PrefixLength = in["prefix_length"].(string)
		oi.PrefixVendor = in["prefix_vendor"].(int)
		oi.PrefixNumber = in["prefix_number"].(int)
		oi.Name = in["name"].(string)
		oi.Value = in["value"].(string)
		oi.CustomVendor = in["custom_vendor"].(int)
		oi.CustomNumber = in["custom_number"].(int)
		oi.Vendor = in["vendor"].(int)
		oi.Number = in["number"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemRadiusServerSamplingEnable1908(d []interface{}) []edpt.SystemRadiusServerSamplingEnable1908 {

	count1 := len(d)
	ret := make([]edpt.SystemRadiusServerSamplingEnable1908, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerSamplingEnable1908
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemRadiusServerDerivedAttribute1909(d []interface{}) edpt.SystemRadiusServerDerivedAttribute1909 {

	count1 := len(d)
	var ret edpt.SystemRadiusServerDerivedAttribute1909
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Usergroup = getObjectSystemRadiusServerDerivedAttributeUsergroup1910(in["usergroup"].([]interface{}))
		ret.Userid = getObjectSystemRadiusServerDerivedAttributeUserid1911(in["userid"].([]interface{}))
	}
	return ret
}

func getObjectSystemRadiusServerDerivedAttributeUsergroup1910(d []interface{}) edpt.SystemRadiusServerDerivedAttributeUsergroup1910 {

	count1 := len(d)
	var ret edpt.SystemRadiusServerDerivedAttributeUsergroup1910
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Attribute = in["attribute"].(string)
		ret.Regex = in["regex"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemRadiusServerDerivedAttributeUserid1911(d []interface{}) edpt.SystemRadiusServerDerivedAttributeUserid1911 {

	count1 := len(d)
	var ret edpt.SystemRadiusServerDerivedAttributeUserid1911
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Attribute = in["attribute"].(string)
		ret.Regex = in["regex"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemReboot1912(d []interface{}) edpt.SystemReboot1912 {

	var ret edpt.SystemReboot1912
	return ret
}

func getObjectSystemResourceAccounting1913(d []interface{}) edpt.SystemResourceAccounting1913 {

	count1 := len(d)
	var ret edpt.SystemResourceAccounting1913
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TemplateList = getSliceSystemResourceAccountingTemplateList1914(in["template_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemResourceAccountingTemplateList1914(d []interface{}) []edpt.SystemResourceAccountingTemplateList1914 {

	count1 := len(d)
	ret := make([]edpt.SystemResourceAccountingTemplateList1914, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemResourceAccountingTemplateList1914
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.AppResources = getObjectSystemResourceAccountingTemplateListAppResources1915(in["app_resources"].([]interface{}))
		oi.NetworkResources = getObjectSystemResourceAccountingTemplateListNetworkResources1947(in["network_resources"].([]interface{}))
		oi.SystemResources = getObjectSystemResourceAccountingTemplateListSystemResources1957(in["system_resources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResources1915(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResources1915 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResources1915
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbDeviceCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1916(in["gslb_device_cfg"].([]interface{}))
		ret.GslbGeoLocationCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1917(in["gslb_geo_location_cfg"].([]interface{}))
		ret.GslbIpListCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1918(in["gslb_ip_list_cfg"].([]interface{}))
		ret.GslbPolicyCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1919(in["gslb_policy_cfg"].([]interface{}))
		ret.GslbServiceCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1920(in["gslb_service_cfg"].([]interface{}))
		ret.GslbServiceIpCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1921(in["gslb_service_ip_cfg"].([]interface{}))
		ret.GslbServicePortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1922(in["gslb_service_port_cfg"].([]interface{}))
		ret.GslbSiteCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1923(in["gslb_site_cfg"].([]interface{}))
		ret.GslbSvcGroupCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1924(in["gslb_svc_group_cfg"].([]interface{}))
		ret.GslbTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1925(in["gslb_template_cfg"].([]interface{}))
		ret.GslbZoneCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1926(in["gslb_zone_cfg"].([]interface{}))
		ret.HealthMonitorCfg = getObjectSystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1927(in["health_monitor_cfg"].([]interface{}))
		ret.RealPortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesRealPortCfg1928(in["real_port_cfg"].([]interface{}))
		ret.RealServerCfg = getObjectSystemResourceAccountingTemplateListAppResourcesRealServerCfg1929(in["real_server_cfg"].([]interface{}))
		ret.ServiceGroupCfg = getObjectSystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1930(in["service_group_cfg"].([]interface{}))
		ret.VirtualServerCfg = getObjectSystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1931(in["virtual_server_cfg"].([]interface{}))
		ret.VirtualPortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1932(in["virtual_port_cfg"].([]interface{}))
		ret.CacheTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1933(in["cache_template_cfg"].([]interface{}))
		ret.ClientSslTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1934(in["client_ssl_template_cfg"].([]interface{}))
		ret.ConnReuseTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1935(in["conn_reuse_template_cfg"].([]interface{}))
		ret.FastTcpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1936(in["fast_tcp_template_cfg"].([]interface{}))
		ret.FastUdpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1937(in["fast_udp_template_cfg"].([]interface{}))
		ret.FixTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1938(in["fix_template_cfg"].([]interface{}))
		ret.HttpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1939(in["http_template_cfg"].([]interface{}))
		ret.LinkCostTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1940(in["link_cost_template_cfg"].([]interface{}))
		ret.PbslbEntryCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1941(in["pbslb_entry_cfg"].([]interface{}))
		ret.PersistCookieTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1942(in["persist_cookie_template_cfg"].([]interface{}))
		ret.PersistSrcipTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1943(in["persist_srcip_template_cfg"].([]interface{}))
		ret.ServerSslTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1944(in["server_ssl_template_cfg"].([]interface{}))
		ret.ProxyTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1945(in["proxy_template_cfg"].([]interface{}))
		ret.StreamTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1946(in["stream_template_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1916(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1916 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1916
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbDeviceMax = in["gslb_device_max"].(int)
		ret.GslbDeviceMinGuarantee = in["gslb_device_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1917(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1917 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1917
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbGeoLocationMax = in["gslb_geo_location_max"].(int)
		ret.GslbGeoLocationMinGuarantee = in["gslb_geo_location_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1918(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1918 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1918
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbIpListMax = in["gslb_ip_list_max"].(int)
		ret.GslbIpListMinGuarantee = in["gslb_ip_list_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1919(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1919 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1919
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbPolicyMax = in["gslb_policy_max"].(int)
		ret.GslbPolicyMinGuarantee = in["gslb_policy_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1920(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1920 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1920
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceMax = in["gslb_service_max"].(int)
		ret.GslbServiceMinGuarantee = in["gslb_service_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1921(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1921 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1921
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceIpMax = in["gslb_service_ip_max"].(int)
		ret.GslbServiceIpMinGuarantee = in["gslb_service_ip_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1922(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1922 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1922
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServicePortMax = in["gslb_service_port_max"].(int)
		ret.GslbServicePortMinGuarantee = in["gslb_service_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1923(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1923 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1923
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSiteMax = in["gslb_site_max"].(int)
		ret.GslbSiteMinGuarantee = in["gslb_site_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1924(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1924 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1924
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSvcGroupMax = in["gslb_svc_group_max"].(int)
		ret.GslbSvcGroupMinGuarantee = in["gslb_svc_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1925(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1925 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1925
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbTemplateMax = in["gslb_template_max"].(int)
		ret.GslbTemplateMinGuarantee = in["gslb_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1926(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1926 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1926
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbZoneMax = in["gslb_zone_max"].(int)
		ret.GslbZoneMinGuarantee = in["gslb_zone_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1927(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1927 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1927
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HealthMonitorMax = in["health_monitor_max"].(int)
		ret.HealthMonitorMinGuarantee = in["health_monitor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesRealPortCfg1928(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesRealPortCfg1928 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesRealPortCfg1928
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealPortMax = in["real_port_max"].(int)
		ret.RealPortMinGuarantee = in["real_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesRealServerCfg1929(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesRealServerCfg1929 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesRealServerCfg1929
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealServerMax = in["real_server_max"].(int)
		ret.RealServerMinGuarantee = in["real_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1930(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1930 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1930
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceGroupMax = in["service_group_max"].(int)
		ret.ServiceGroupMinGuarantee = in["service_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1931(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1931 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1931
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualServerMax = in["virtual_server_max"].(int)
		ret.VirtualServerMinGuarantee = in["virtual_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1932(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1932 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1932
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualPortMax = in["virtual_port_max"].(int)
		ret.VirtualPortMinGuarantee = in["virtual_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1933(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1933 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1933
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheTemplateMax = in["cache_template_max"].(int)
		ret.CacheTemplateMinGuarantee = in["cache_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1934(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1934 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1934
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientSslTemplateMax = in["client_ssl_template_max"].(int)
		ret.ClientSslTemplateMinGuarantee = in["client_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1935(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1935 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1935
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnReuseTemplateMax = in["conn_reuse_template_max"].(int)
		ret.ConnReuseTemplateMinGuarantee = in["conn_reuse_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1936(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1936 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1936
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastTcpTemplateMax = in["fast_tcp_template_max"].(int)
		ret.FastTcpTemplateMinGuarantee = in["fast_tcp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1937(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1937 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1937
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastUdpTemplateMax = in["fast_udp_template_max"].(int)
		ret.FastUdpTemplateMinGuarantee = in["fast_udp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1938(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1938 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1938
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixTemplateMax = in["fix_template_max"].(int)
		ret.FixTemplateMinGuarantee = in["fix_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1939(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1939 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1939
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpTemplateMax = in["http_template_max"].(int)
		ret.HttpTemplateMinGuarantee = in["http_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1940(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1940 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1940
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LinkCostTemplateMax = in["link_cost_template_max"].(int)
		ret.LinkCostTemplateMinGuarantee = in["link_cost_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1941(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1941 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1941
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PbslbEntryMax = in["pbslb_entry_max"].(int)
		ret.PbslbEntryMinGuarantee = in["pbslb_entry_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1942(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1942 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1942
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistCookieTemplateMax = in["persist_cookie_template_max"].(int)
		ret.PersistCookieTemplateMinGuarantee = in["persist_cookie_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1943(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1943 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1943
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistSrcipTemplateMax = in["persist_srcip_template_max"].(int)
		ret.PersistSrcipTemplateMinGuarantee = in["persist_srcip_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1944(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1944 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1944
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerSslTemplateMax = in["server_ssl_template_max"].(int)
		ret.ServerSslTemplateMinGuarantee = in["server_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1945(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1945 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1945
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyTemplateMax = in["proxy_template_max"].(int)
		ret.ProxyTemplateMinGuarantee = in["proxy_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1946(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1946 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1946
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StreamTemplateMax = in["stream_template_max"].(int)
		ret.StreamTemplateMinGuarantee = in["stream_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResources1947(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResources1947 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResources1947
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv4RouteCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1948(in["static_ipv4_route_cfg"].([]interface{}))
		ret.StaticIpv6RouteCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1949(in["static_ipv6_route_cfg"].([]interface{}))
		ret.Ipv4AclLineCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1950(in["ipv4_acl_line_cfg"].([]interface{}))
		ret.Ipv6AclLineCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1951(in["ipv6_acl_line_cfg"].([]interface{}))
		ret.StaticArpCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1952(in["static_arp_cfg"].([]interface{}))
		ret.StaticNeighborCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1953(in["static_neighbor_cfg"].([]interface{}))
		ret.StaticMacCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1954(in["static_mac_cfg"].([]interface{}))
		ret.ObjectGroupCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1955(in["object_group_cfg"].([]interface{}))
		ret.ObjectGroupClauseCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1956(in["object_group_clause_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1948(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1948 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1948
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv4RouteMax = in["static_ipv4_route_max"].(int)
		ret.StaticIpv4RouteMinGuarantee = in["static_ipv4_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1949(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1949 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1949
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv6RouteMax = in["static_ipv6_route_max"].(int)
		ret.StaticIpv6RouteMinGuarantee = in["static_ipv6_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1950(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1950 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1950
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4AclLineMax = in["ipv4_acl_line_max"].(int)
		ret.Ipv4AclLineMinGuarantee = in["ipv4_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1951(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1951 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1951
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6AclLineMax = in["ipv6_acl_line_max"].(int)
		ret.Ipv6AclLineMinGuarantee = in["ipv6_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1952(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1952 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1952
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticArpMax = in["static_arp_max"].(int)
		ret.StaticArpMinGuarantee = in["static_arp_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1953(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1953 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1953
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNeighborMax = in["static_neighbor_max"].(int)
		ret.StaticNeighborMinGuarantee = in["static_neighbor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1954(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1954 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1954
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticMacMax = in["static_mac_max"].(int)
		ret.StaticMacMinGuarantee = in["static_mac_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1955(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1955 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1955
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupMax = in["object_group_max"].(int)
		ret.ObjectGroupMinGuarantee = in["object_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1956(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1956 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1956
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupClauseMax = in["object_group_clause_max"].(int)
		ret.ObjectGroupClauseMinGuarantee = in["object_group_clause_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResources1957(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResources1957 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResources1957
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BwLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1958(in["bw_limit_cfg"].([]interface{}))
		ret.ConcurrentSessionLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1959(in["concurrent_session_limit_cfg"].([]interface{}))
		ret.L4SessionLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1960(in["l4_session_limit_cfg"].([]interface{}))
		ret.L4cpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1961(in["l4cps_limit_cfg"].([]interface{}))
		ret.L7cpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1962(in["l7cps_limit_cfg"].([]interface{}))
		ret.NatcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1963(in["natcps_limit_cfg"].([]interface{}))
		ret.FwcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1964(in["fwcps_limit_cfg"].([]interface{}))
		ret.SslThroughputLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1965(in["ssl_throughput_limit_cfg"].([]interface{}))
		ret.SslcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1966(in["sslcps_limit_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1958(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1958 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1958
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BwLimitMax = in["bw_limit_max"].(int)
		ret.BwLimitWatermarkDisable = in["bw_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1959(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1959 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1959
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConcurrentSessionLimitMax = in["concurrent_session_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1960(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1960 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1960
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4SessionLimitMax = in["l4_session_limit_max"].(string)
		ret.L4SessionLimitMinGuarantee = in["l4_session_limit_min_guarantee"].(string)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1961(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1961 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1961
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4cpsLimitMax = in["l4cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1962(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1962 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1962
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L7cpsLimitMax = in["l7cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1963(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1963 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1963
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatcpsLimitMax = in["natcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1964(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1964 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1964
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FwcpsLimitMax = in["fwcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1965(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1965 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1965
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslThroughputLimitMax = in["ssl_throughput_limit_max"].(int)
		ret.SslThroughputLimitWatermarkDisable = in["ssl_throughput_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1966(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1966 {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1966
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslcpsLimitMax = in["sslcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceUsage1967(d []interface{}) edpt.SystemResourceUsage1967 {

	count1 := len(d)
	var ret edpt.SystemResourceUsage1967
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatPoolAddrCount = in["nat_pool_addr_count"].(int)
		ret.L4SessionCount = in["l4_session_count"].(int)
		ret.AuthPortalHtmlFileSize = in["auth_portal_html_file_size"].(int)
		ret.AuthPortalImageFileSize = in["auth_portal_image_file_size"].(int)
		ret.MaxAflexFileSize = in["max_aflex_file_size"].(int)
		ret.AflexTableEntryCount = in["aflex_table_entry_count"].(int)
		ret.ClassListIpv6AddrCount = in["class_list_ipv6_addr_count"].(int)
		ret.ClassListAcEntryCount = in["class_list_ac_entry_count"].(int)
		ret.ClassListEntryCount = in["class_list_entry_count"].(int)
		ret.MaxAflexAuthzCollectionNumber = in["max_aflex_authz_collection_number"].(int)
		ret.RadiusTableSize = in["radius_table_size"].(int)
		ret.AuthzPolicyNumber = in["authz_policy_number"].(int)
		ret.IpsecSaNumber = in["ipsec_sa_number"].(int)
		ret.RamCacheMemoryLimit = in["ram_cache_memory_limit"].(int)
		ret.AuthSessionCount = in["auth_session_count"].(int)
		ret.NgwafCacheEntry = in["ngwaf_cache_entry"].(int)
		//omit uuid
		ret.Visibility = getObjectSystemResourceUsageVisibility1968(in["visibility"].([]interface{}))
	}
	return ret
}

func getObjectSystemResourceUsageVisibility1968(d []interface{}) edpt.SystemResourceUsageVisibility1968 {

	count1 := len(d)
	var ret edpt.SystemResourceUsageVisibility1968
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonitoredEntityCount = in["monitored_entity_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSession1969(d []interface{}) edpt.SystemSession1969 {

	count1 := len(d)
	var ret edpt.SystemSession1969
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemSessionSamplingEnable1970(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemSessionSamplingEnable1970(d []interface{}) []edpt.SystemSessionSamplingEnable1970 {

	count1 := len(d)
	ret := make([]edpt.SystemSessionSamplingEnable1970, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemSessionSamplingEnable1970
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemSessionReclaimLimit1971(d []interface{}) edpt.SystemSessionReclaimLimit1971 {

	count1 := len(d)
	var ret edpt.SystemSessionReclaimLimit1971
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NscanLimit = in["nscan_limit"].(int)
		ret.ScanFreq = in["scan_freq"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSetRxtxDescSize1972(d []interface{}) edpt.SystemSetRxtxDescSize1972 {

	count1 := len(d)
	var ret edpt.SystemSetRxtxDescSize1972
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.RxdSize = in["rxd_size"].(int)
		ret.TxdSize = in["txd_size"].(int)
	}
	return ret
}

func getObjectSystemSetRxtxQueue1973(d []interface{}) edpt.SystemSetRxtxQueue1973 {

	count1 := len(d)
	var ret edpt.SystemSetRxtxQueue1973
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.RxqSize = in["rxq_size"].(int)
		ret.TxqSize = in["txq_size"].(int)
	}
	return ret
}

func getObjectSystemSetTcpSynPerSec1974(d []interface{}) edpt.SystemSetTcpSynPerSec1974 {

	count1 := len(d)
	var ret edpt.SystemSetTcpSynPerSec1974
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpSynValue = in["tcp_syn_value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSharedPollMode1975(d []interface{}) edpt.SystemSharedPollMode1975 {

	count1 := len(d)
	var ret edpt.SystemSharedPollMode1975
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemShellPrivileges1976(d []interface{}) edpt.SystemShellPrivileges1976 {

	count1 := len(d)
	var ret edpt.SystemShellPrivileges1976
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableShellPrivileges = in["enable_shell_privileges"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemShmLogging1977(d []interface{}) edpt.SystemShmLogging1977 {

	count1 := len(d)
	var ret edpt.SystemShmLogging1977
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemShutdown1978(d []interface{}) edpt.SystemShutdown1978 {

	var ret edpt.SystemShutdown1978
	return ret
}

func getObjectSystemSoftwareTcam1979(d []interface{}) edpt.SystemSoftwareTcam1979 {

	count1 := len(d)
	var ret edpt.SystemSoftwareTcam1979
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSpeProfile1980(d []interface{}) edpt.SystemSpeProfile1980 {

	count1 := len(d)
	var ret edpt.SystemSpeProfile1980
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
	}
	return ret
}

func getObjectSystemSpeStatus1981(d []interface{}) edpt.SystemSpeStatus1981 {

	var ret edpt.SystemSpeStatus1981
	return ret
}

func getObjectSystemSslHwMemory1982(d []interface{}) edpt.SystemSslHwMemory1982 {

	count1 := len(d)
	var ret edpt.SystemSslHwMemory1982
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MemBlockCfg = getSliceSystemSslHwMemoryMemBlockCfg1983(in["mem_block_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemSslHwMemoryMemBlockCfg1983(d []interface{}) []edpt.SystemSslHwMemoryMemBlockCfg1983 {

	count1 := len(d)
	ret := make([]edpt.SystemSslHwMemoryMemBlockCfg1983, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemSslHwMemoryMemBlockCfg1983
		oi.MemBlock = in["mem_block"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemSslReqQ1984(d []interface{}) edpt.SystemSslReqQ1984 {

	count1 := len(d)
	var ret edpt.SystemSslReqQ1984
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemSslReqQSamplingEnable1985(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemSslReqQSamplingEnable1985(d []interface{}) []edpt.SystemSslReqQSamplingEnable1985 {

	count1 := len(d)
	ret := make([]edpt.SystemSslReqQSamplingEnable1985, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemSslReqQSamplingEnable1985
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemSslScv1986(d []interface{}) edpt.SystemSslScv1986 {

	count1 := len(d)
	var ret edpt.SystemSslScv1986
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSslScvVerifyCrlSign1987(d []interface{}) edpt.SystemSslScvVerifyCrlSign1987 {

	count1 := len(d)
	var ret edpt.SystemSslScvVerifyCrlSign1987
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSslScvVerifyHost1988(d []interface{}) edpt.SystemSslScvVerifyHost1988 {

	count1 := len(d)
	var ret edpt.SystemSslScvVerifyHost1988
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSslSetCompatibleCipher1989(d []interface{}) edpt.SystemSslSetCompatibleCipher1989 {

	count1 := len(d)
	var ret edpt.SystemSslSetCompatibleCipher1989
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSslStatus1990(d []interface{}) edpt.SystemSslStatus1990 {

	var ret edpt.SystemSslStatus1990
	return ret
}

func getObjectSystemSyslogTimeMsec1991(d []interface{}) edpt.SystemSyslogTimeMsec1991 {

	count1 := len(d)
	var ret edpt.SystemSyslogTimeMsec1991
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableFlag = in["enable_flag"].(int)
	}
	return ret
}

func getObjectSystemTableIntegrity1992(d []interface{}) edpt.SystemTableIntegrity1992 {

	count1 := len(d)
	var ret edpt.SystemTableIntegrity1992
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Table = in["table"].(string)
		ret.AuditAction = in["audit_action"].(string)
		ret.AutoSyncAction = in["auto_sync_action"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSystemTableIntegritySamplingEnable1993(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemTableIntegritySamplingEnable1993(d []interface{}) []edpt.SystemTableIntegritySamplingEnable1993 {

	count1 := len(d)
	ret := make([]edpt.SystemTableIntegritySamplingEnable1993, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTableIntegritySamplingEnable1993
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		oi.Counters3 = in["counters3"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTcp1994(d []interface{}) edpt.SystemTcp1994 {

	count1 := len(d)
	var ret edpt.SystemTcp1994
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemTcpSamplingEnable1995(in["sampling_enable"].([]interface{}))
		ret.RateLimitResetUnknownConn = getObjectSystemTcpRateLimitResetUnknownConn1996(in["rate_limit_reset_unknown_conn"].([]interface{}))
	}
	return ret
}

func getSliceSystemTcpSamplingEnable1995(d []interface{}) []edpt.SystemTcpSamplingEnable1995 {

	count1 := len(d)
	ret := make([]edpt.SystemTcpSamplingEnable1995, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTcpSamplingEnable1995
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTcpRateLimitResetUnknownConn1996(d []interface{}) edpt.SystemTcpRateLimitResetUnknownConn1996 {

	count1 := len(d)
	var ret edpt.SystemTcpRateLimitResetUnknownConn1996
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktRateForResetUnknownConn = in["pkt_rate_for_reset_unknown_conn"].(int)
		ret.LogForResetUnknownConn = in["log_for_reset_unknown_conn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTcpStats1997(d []interface{}) edpt.SystemTcpStats1997 {

	count1 := len(d)
	var ret edpt.SystemTcpStats1997
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemTcpStatsSamplingEnable1998(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemTcpStatsSamplingEnable1998(d []interface{}) []edpt.SystemTcpStatsSamplingEnable1998 {

	count1 := len(d)
	ret := make([]edpt.SystemTcpStatsSamplingEnable1998, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTcpStatsSamplingEnable1998
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTcpSynPerSec1999(d []interface{}) edpt.SystemTcpSynPerSec1999 {

	var ret edpt.SystemTcpSynPerSec1999
	return ret
}

func getObjectSystemTelemetryLog2000(d []interface{}) edpt.SystemTelemetryLog2000 {

	count1 := len(d)
	var ret edpt.SystemTelemetryLog2000
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TopKSourceList = getObjectSystemTelemetryLogTopKSourceList2001(in["top_k_source_list"].([]interface{}))
		ret.TopKAppSvcList = getObjectSystemTelemetryLogTopKAppSvcList2002(in["top_k_app_svc_list"].([]interface{}))
		ret.DeviceStatus = getObjectSystemTelemetryLogDeviceStatus2003(in["device_status"].([]interface{}))
		ret.Environment = getObjectSystemTelemetryLogEnvironment2004(in["environment"].([]interface{}))
		ret.PartitionMetrics = getObjectSystemTelemetryLogPartitionMetrics2005(in["partition_metrics"].([]interface{}))
	}
	return ret
}

func getObjectSystemTelemetryLogTopKSourceList2001(d []interface{}) edpt.SystemTelemetryLogTopKSourceList2001 {

	var ret edpt.SystemTelemetryLogTopKSourceList2001
	return ret
}

func getObjectSystemTelemetryLogTopKAppSvcList2002(d []interface{}) edpt.SystemTelemetryLogTopKAppSvcList2002 {

	var ret edpt.SystemTelemetryLogTopKAppSvcList2002
	return ret
}

func getObjectSystemTelemetryLogDeviceStatus2003(d []interface{}) edpt.SystemTelemetryLogDeviceStatus2003 {

	var ret edpt.SystemTelemetryLogDeviceStatus2003
	return ret
}

func getObjectSystemTelemetryLogEnvironment2004(d []interface{}) edpt.SystemTelemetryLogEnvironment2004 {

	var ret edpt.SystemTelemetryLogEnvironment2004
	return ret
}

func getObjectSystemTelemetryLogPartitionMetrics2005(d []interface{}) edpt.SystemTelemetryLogPartitionMetrics2005 {

	var ret edpt.SystemTelemetryLogPartitionMetrics2005
	return ret
}

func getObjectSystemTemplate2006(d []interface{}) edpt.SystemTemplate2006 {

	count1 := len(d)
	var ret edpt.SystemTemplate2006
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplatePolicy = in["template_policy"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemTemplateBind2007(d []interface{}) edpt.SystemTemplateBind2007 {

	count1 := len(d)
	var ret edpt.SystemTemplateBind2007
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonitorList = getSliceSystemTemplateBindMonitorList(in["monitor_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemTemplateBindMonitorList(d []interface{}) []edpt.SystemTemplateBindMonitorList {

	count1 := len(d)
	ret := make([]edpt.SystemTemplateBindMonitorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTemplateBindMonitorList
		oi.TemplateMonitor = in["template_monitor"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemThroughput2008(d []interface{}) edpt.SystemThroughput2008 {

	count1 := len(d)
	var ret edpt.SystemThroughput2008
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemThroughputSamplingEnable2009(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemThroughputSamplingEnable2009(d []interface{}) []edpt.SystemThroughputSamplingEnable2009 {

	count1 := len(d)
	ret := make([]edpt.SystemThroughputSamplingEnable2009, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemThroughputSamplingEnable2009
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTimeoutValue2010(d []interface{}) edpt.SystemTimeoutValue2010 {

	count1 := len(d)
	var ret edpt.SystemTimeoutValue2010
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ftp = in["ftp"].(int)
		ret.Scp = in["scp"].(int)
		ret.Sftp = in["sftp"].(int)
		ret.Tftp = in["tftp"].(int)
		ret.Http = in["http"].(int)
		ret.Https = in["https"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTrunk2011(d []interface{}) edpt.SystemTrunk2011 {

	count1 := len(d)
	var ret edpt.SystemTrunk2011
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LoadBalance = getObjectSystemTrunkLoadBalance2012(in["load_balance"].([]interface{}))
	}
	return ret
}

func getObjectSystemTrunkLoadBalance2012(d []interface{}) edpt.SystemTrunkLoadBalance2012 {

	count1 := len(d)
	var ret edpt.SystemTrunkLoadBalance2012
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UseL3 = in["use_l3"].(int)
		ret.UseL4 = in["use_l4"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTrunkHwHash2013(d []interface{}) edpt.SystemTrunkHwHash2013 {

	count1 := len(d)
	var ret edpt.SystemTrunkHwHash2013
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTrunkXauiHwHash2014(d []interface{}) edpt.SystemTrunkXauiHwHash2014 {

	count1 := len(d)
	var ret edpt.SystemTrunkXauiHwHash2014
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTso2015(d []interface{}) edpt.SystemTso2015 {

	count1 := len(d)
	var ret edpt.SystemTso2015
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemUdp2016(d []interface{}) edpt.SystemUdp2016 {

	count1 := len(d)
	var ret edpt.SystemUdp2016
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipChecksumWhenZero = in["skip_checksum_when_zero"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemUpgradeStatus2017(d []interface{}) edpt.SystemUpgradeStatus2017 {

	var ret edpt.SystemUpgradeStatus2017
	return ret
}

func getObjectSystemVeMacScheme2018(d []interface{}) edpt.SystemVeMacScheme2018 {

	count1 := len(d)
	var ret edpt.SystemVeMacScheme2018
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VeMacSchemeVal = in["ve_mac_scheme_val"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemXauiDlbMode2019(d []interface{}) edpt.SystemXauiDlbMode2019 {

	count1 := len(d)
	var ret edpt.SystemXauiDlbMode2019
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSystem(d *schema.ResourceData) edpt.System {
	var ret edpt.System
	ret.Inst.AddCpuCore = getObjectSystemAddCpuCore1755(d.Get("add_cpu_core").([]interface{}))
	ret.Inst.AddPort = getObjectSystemAddPort1756(d.Get("add_port").([]interface{}))
	ret.Inst.AllVlanLimit = getObjectSystemAllVlanLimit1757(d.Get("all_vlan_limit").([]interface{}))
	ret.Inst.AnomalyLog = d.Get("anomaly_log").(int)
	ret.Inst.AnomalyLogRateLimit = d.Get("anomaly_log_rate_limit").(int)
	ret.Inst.AppPerformance = getObjectSystemAppPerformance1758(d.Get("app_performance").([]interface{}))
	ret.Inst.AppsGlobal = getObjectSystemAppsGlobal1760(d.Get("apps_global").([]interface{}))
	ret.Inst.AsicDebugDump = getObjectSystemAsicDebugDump1761(d.Get("asic_debug_dump").([]interface{}))
	ret.Inst.AsicMmuFailSafe = getObjectSystemAsicMmuFailSafe1762(d.Get("asic_mmu_fail_safe").([]interface{}))
	ret.Inst.AttackLog = d.Get("attack_log").(int)
	ret.Inst.Bandwidth = getObjectSystemBandwidth1763(d.Get("bandwidth").([]interface{}))
	ret.Inst.Bfd = getObjectSystemBfd1765(d.Get("bfd").([]interface{}))
	ret.Inst.ClThreatCategory = getObjectSystemClThreatCategory1767(d.Get("cl_threat_category").([]interface{}))
	ret.Inst.ClassListHitcountEnable = d.Get("class_list_hitcount_enable").(int)
	ret.Inst.CliMonitorInterval = getObjectSystemCliMonitorInterval1768(d.Get("cli_monitor_interval").([]interface{}))
	ret.Inst.CmUpdateFileNameRef = getObjectSystemCmUpdateFileNameRef1769(d.Get("cm_update_file_name_ref").([]interface{}))
	ret.Inst.ConfigMgmt = getObjectSystemConfigMgmt1770(d.Get("config_mgmt").([]interface{}))
	ret.Inst.ControlCpu = getObjectSystemControlCpu1774(d.Get("control_cpu").([]interface{}))
	ret.Inst.Core = getObjectSystemCore1775(d.Get("core").([]interface{}))
	ret.Inst.CosqShow = getObjectSystemCosqShow1776(d.Get("cosq_show").([]interface{}))
	ret.Inst.CosqStats = getObjectSystemCosqStats1777(d.Get("cosq_stats").([]interface{}))
	ret.Inst.CounterLibAccounting = getObjectSystemCounterLibAccounting1778(d.Get("counter_lib_accounting").([]interface{}))
	ret.Inst.CpuHyperThread = getObjectSystemCpuHyperThread1779(d.Get("cpu_hyper_thread").([]interface{}))
	ret.Inst.CpuList = getObjectSystemCpuList1780(d.Get("cpu_list").([]interface{}))
	ret.Inst.CpuLoadSharing = getObjectSystemCpuLoadSharing1781(d.Get("cpu_load_sharing").([]interface{}))
	ret.Inst.CpuMap = getObjectSystemCpuMap1784(d.Get("cpu_map").([]interface{}))
	ret.Inst.CpuPacketPrioSupport = getObjectSystemCpuPacketPrioSupport1785(d.Get("cpu_packet_prio_support").([]interface{}))
	ret.Inst.DataCpu = getObjectSystemDataCpu1786(d.Get("data_cpu").([]interface{}))
	ret.Inst.DdosAttack = d.Get("ddos_attack").(int)
	ret.Inst.DdosLog = d.Get("ddos_log").(int)
	ret.Inst.DefaultMtu = d.Get("default_mtu").(int)
	ret.Inst.DelPort = getObjectSystemDelPort1787(d.Get("del_port").([]interface{}))
	ret.Inst.DeleteCpuCore = getObjectSystemDeleteCpuCore1788(d.Get("delete_cpu_core").([]interface{}))
	ret.Inst.DisableSshAgentForwarding = d.Get("disable_ssh_agent_forwarding").(int)
	ret.Inst.Dns = getObjectSystemDns1789(d.Get("dns").([]interface{}))
	ret.Inst.DnsCache = getObjectSystemDnsCache1793(d.Get("dns_cache").([]interface{}))
	ret.Inst.DomainListHitcountEnable = d.Get("domain_list_hitcount_enable").(int)
	ret.Inst.DomainListInfo = getObjectSystemDomainListInfo1795(d.Get("domain_list_info").([]interface{}))
	ret.Inst.DomainListSettings = getObjectSystemDomainListSettings1796(d.Get("domain_list_settings").([]interface{}))
	ret.Inst.DpdkStats = getObjectSystemDpdkStats1797(d.Get("dpdk_stats").([]interface{}))
	ret.Inst.DropLinuxClosedPortSyn = d.Get("drop_linux_closed_port_syn").(string)
	ret.Inst.DynamicServiceDnsSocketPool = d.Get("dynamic_service_dns_socket_pool").(int)
	ret.Inst.EnableDiskEncryption = getObjectSystemEnableDiskEncryption1799(d.Get("enable_disk_encryption").([]interface{}))
	ret.Inst.EnablePassword = getObjectSystemEnablePassword1800(d.Get("enable_password").([]interface{}))
	ret.Inst.Environment = getObjectSystemEnvironment1801(d.Get("environment").([]interface{}))
	ret.Inst.EvenPortHashEnable = d.Get("even_port_hash_enable").(int)
	ret.Inst.ExtOnlyLogging = getObjectSystemExtOnlyLogging1802(d.Get("ext_only_logging").([]interface{}))
	ret.Inst.ForcedGroupSpeedList = getSliceSystemForcedGroupSpeedList(d.Get("forced_group_speed_list").([]interface{}))
	ret.Inst.FpgaCoreCrc = getObjectSystemFpgaCoreCrc1803(d.Get("fpga_core_crc").([]interface{}))
	ret.Inst.FpgaDrop = getObjectSystemFpgaDrop1804(d.Get("fpga_drop").([]interface{}))
	ret.Inst.Fw = getObjectSystemFw1806(d.Get("fw").([]interface{}))
	ret.Inst.GeoDbHitcountEnable = d.Get("geo_db_hitcount_enable").(int)
	ret.Inst.GeoLocation = getObjectSystemGeoLocation1807(d.Get("geo_location").([]interface{}))
	ret.Inst.Geoloc = getObjectSystemGeoloc1811(d.Get("geoloc").([]interface{}))
	ret.Inst.GeolocListList = getSliceSystemGeolocListList(d.Get("geoloc_list_list").([]interface{}))
	ret.Inst.GeolocNameHelper = getObjectSystemGeolocNameHelper1813(d.Get("geoloc_name_helper").([]interface{}))
	ret.Inst.GeolocationFile = getObjectSystemGeolocationFile1815(d.Get("geolocation_file").([]interface{}))
	ret.Inst.Glid = getObjectSystemGlid1817(d.Get("glid").([]interface{}))
	ret.Inst.GuestFile = getObjectSystemGuestFile1818(d.Get("guest_file").([]interface{}))
	ret.Inst.GuiImageList = getObjectSystemGuiImageList1819(d.Get("gui_image_list").([]interface{}))
	ret.Inst.Hardware = getObjectSystemHardware1820(d.Get("hardware").([]interface{}))
	ret.Inst.HardwareAccelerate = getObjectSystemHardwareAccelerate1821(d.Get("hardware_accelerate").([]interface{}))
	ret.Inst.HealthCheckList = getSliceSystemHealthCheckList(d.Get("health_check_list").([]interface{}))
	ret.Inst.HighMemoryL4Session = getObjectSystemHighMemoryL4Session1825(d.Get("high_memory_l4_session").([]interface{}))
	ret.Inst.HrxqStatus = getObjectSystemHrxqStatus1826(d.Get("hrxq_status").([]interface{}))
	ret.Inst.HwBlockingEnable = d.Get("hw_blocking_enable").(int)
	ret.Inst.Icmp = getObjectSystemIcmp1827(d.Get("icmp").([]interface{}))
	ret.Inst.IcmpRate = getObjectSystemIcmpRate1829(d.Get("icmp_rate").([]interface{}))
	ret.Inst.Icmp6 = getObjectSystemIcmp61831(d.Get("icmp6").([]interface{}))
	ret.Inst.InuseCpuList = getObjectSystemInuseCpuList1833(d.Get("inuse_cpu_list").([]interface{}))
	ret.Inst.InusePortList = getObjectSystemInusePortList1834(d.Get("inuse_port_list").([]interface{}))
	ret.Inst.IoCpu = getObjectSystemIoCpu1835(d.Get("io_cpu").([]interface{}))
	ret.Inst.Ip = getObjectSystemIp1836(d.Get("ip").([]interface{}))
	ret.Inst.IpDnsCache = getObjectSystemIpDnsCache1837(d.Get("ip_dns_cache").([]interface{}))
	ret.Inst.IpStats = getObjectSystemIpStats1838(d.Get("ip_stats").([]interface{}))
	ret.Inst.IpThreatList = getObjectSystemIpThreatList1840(d.Get("ip_threat_list").([]interface{}))
	ret.Inst.Ip6Stats = getObjectSystemIp6Stats1854(d.Get("ip6_stats").([]interface{}))
	ret.Inst.Ipmi = getObjectSystemIpmi1856(d.Get("ipmi").([]interface{}))
	ret.Inst.IpmiService = getObjectSystemIpmiService1861(d.Get("ipmi_service").([]interface{}))
	ret.Inst.Ipsec = getObjectSystemIpsec1862(d.Get("ipsec").([]interface{}))
	ret.Inst.Ipv6 = getObjectSystemIpv61864(d.Get("ipv6").([]interface{}))
	ret.Inst.Ipv6PrefixLength = d.Get("ipv6_prefix_length").(int)
	ret.Inst.JobOffload = getObjectSystemJobOffload1865(d.Get("job_offload").([]interface{}))
	ret.Inst.LinkCapability = getObjectSystemLinkCapability1867(d.Get("link_capability").([]interface{}))
	ret.Inst.LinkMonitor = getObjectSystemLinkMonitor1868(d.Get("link_monitor").([]interface{}))
	ret.Inst.Lro = getObjectSystemLro1869(d.Get("lro").([]interface{}))
	ret.Inst.ManagementInterfaceMode = getObjectSystemManagementInterfaceMode1870(d.Get("management_interface_mode").([]interface{}))
	ret.Inst.Memory = getObjectSystemMemory1871(d.Get("memory").([]interface{}))
	ret.Inst.MemoryBlockDebug = getObjectSystemMemoryBlockDebug1873(d.Get("memory_block_debug").([]interface{}))
	ret.Inst.MfaAuth = getObjectSystemMfaAuth1874(d.Get("mfa_auth").([]interface{}))
	ret.Inst.MfaCertStore = getObjectSystemMfaCertStore1875(d.Get("mfa_cert_store").([]interface{}))
	ret.Inst.MfaManagement = getObjectSystemMfaManagement1876(d.Get("mfa_management").([]interface{}))
	ret.Inst.MfaValidationType = getObjectSystemMfaValidationType1877(d.Get("mfa_validation_type").([]interface{}))
	ret.Inst.MgmtPort = getObjectSystemMgmtPort1878(d.Get("mgmt_port").([]interface{}))
	ret.Inst.ModifyPort = getObjectSystemModifyPort1879(d.Get("modify_port").([]interface{}))
	ret.Inst.ModuleCtrlCpu = d.Get("module_ctrl_cpu").(string)
	ret.Inst.MonTemplate = getObjectSystemMonTemplate1880(d.Get("mon_template").([]interface{}))
	ret.Inst.MultiQueueSupport = getObjectSystemMultiQueueSupport1884(d.Get("multi_queue_support").([]interface{}))
	ret.Inst.NdiscRa = getObjectSystemNdiscRa1885(d.Get("ndisc_ra").([]interface{}))
	ret.Inst.NetvscMonitor = getObjectSystemNetvscMonitor1887(d.Get("netvsc_monitor").([]interface{}))
	ret.Inst.NsmA10lb = getObjectSystemNsmA10lb1888(d.Get("nsm_a10lb").([]interface{}))
	ret.Inst.PasswordPolicy = getObjectSystemPasswordPolicy1889(d.Get("password_policy").([]interface{}))
	ret.Inst.PathList = getSliceSystemPathList(d.Get("path_list").([]interface{}))
	ret.Inst.Pbslb = getObjectSystemPbslb1890(d.Get("pbslb").([]interface{}))
	ret.Inst.PerVlanLimit = getObjectSystemPerVlanLimit1892(d.Get("per_vlan_limit").([]interface{}))
	ret.Inst.Platformtype = getObjectSystemPlatformtype1893(d.Get("platformtype").([]interface{}))
	ret.Inst.PortCount = getObjectSystemPortCount1894(d.Get("port_count").([]interface{}))
	ret.Inst.PortInfo = getObjectSystemPortInfo1895(d.Get("port_info").([]interface{}))
	ret.Inst.PortList = getObjectSystemPortList1896(d.Get("port_list").([]interface{}))
	ret.Inst.Ports = getObjectSystemPorts1897(d.Get("ports").([]interface{}))
	ret.Inst.PowerOnSelfTest = getObjectSystemPowerOnSelfTest1898(d.Get("power_on_self_test").([]interface{}))
	ret.Inst.ProbeNetworkDevices = getObjectSystemProbeNetworkDevices1899(d.Get("probe_network_devices").([]interface{}))
	ret.Inst.PromiscuousMode = d.Get("promiscuous_mode").(int)
	ret.Inst.PsuInfo = getObjectSystemPsuInfo1900(d.Get("psu_info").([]interface{}))
	ret.Inst.QInQ = getObjectSystemQInQ1901(d.Get("q_in_q").([]interface{}))
	ret.Inst.QueuingBuffer = getObjectSystemQueuingBuffer1902(d.Get("queuing_buffer").([]interface{}))
	ret.Inst.Radius = getObjectSystemRadius1903(d.Get("radius").([]interface{}))
	ret.Inst.Reboot = getObjectSystemReboot1912(d.Get("reboot").([]interface{}))
	ret.Inst.ResourceAccounting = getObjectSystemResourceAccounting1913(d.Get("resource_accounting").([]interface{}))
	ret.Inst.ResourceUsage = getObjectSystemResourceUsage1967(d.Get("resource_usage").([]interface{}))
	ret.Inst.RfcIpfixIeSpec = d.Get("rfc_ipfix_ie_spec").(string)
	ret.Inst.Session = getObjectSystemSession1969(d.Get("session").([]interface{}))
	ret.Inst.SessionReclaimLimit = getObjectSystemSessionReclaimLimit1971(d.Get("session_reclaim_limit").([]interface{}))
	ret.Inst.SetRxtxDescSize = getObjectSystemSetRxtxDescSize1972(d.Get("set_rxtx_desc_size").([]interface{}))
	ret.Inst.SetRxtxQueue = getObjectSystemSetRxtxQueue1973(d.Get("set_rxtx_queue").([]interface{}))
	ret.Inst.SetTcpSynPerSec = getObjectSystemSetTcpSynPerSec1974(d.Get("set_tcp_syn_per_sec").([]interface{}))
	ret.Inst.SharedPollMode = getObjectSystemSharedPollMode1975(d.Get("shared_poll_mode").([]interface{}))
	ret.Inst.ShellPrivileges = getObjectSystemShellPrivileges1976(d.Get("shell_privileges").([]interface{}))
	ret.Inst.ShmLogging = getObjectSystemShmLogging1977(d.Get("shm_logging").([]interface{}))
	ret.Inst.Shutdown = getObjectSystemShutdown1978(d.Get("shutdown").([]interface{}))
	ret.Inst.SoftwareTcam = getObjectSystemSoftwareTcam1979(d.Get("software_tcam").([]interface{}))
	ret.Inst.SpeProfile = getObjectSystemSpeProfile1980(d.Get("spe_profile").([]interface{}))
	ret.Inst.SpeStatus = getObjectSystemSpeStatus1981(d.Get("spe_status").([]interface{}))
	ret.Inst.SrcIpHashEnable = d.Get("src_ip_hash_enable").(int)
	ret.Inst.SslHwMemory = getObjectSystemSslHwMemory1982(d.Get("ssl_hw_memory").([]interface{}))
	ret.Inst.SslReqQ = getObjectSystemSslReqQ1984(d.Get("ssl_req_q").([]interface{}))
	ret.Inst.SslScv = getObjectSystemSslScv1986(d.Get("ssl_scv").([]interface{}))
	ret.Inst.SslScvVerifyCrlSign = getObjectSystemSslScvVerifyCrlSign1987(d.Get("ssl_scv_verify_crl_sign").([]interface{}))
	ret.Inst.SslScvVerifyHost = getObjectSystemSslScvVerifyHost1988(d.Get("ssl_scv_verify_host").([]interface{}))
	ret.Inst.SslSetCompatibleCipher = getObjectSystemSslSetCompatibleCipher1989(d.Get("ssl_set_compatible_cipher").([]interface{}))
	ret.Inst.SslStatus = getObjectSystemSslStatus1990(d.Get("ssl_status").([]interface{}))
	ret.Inst.SyslogTimeMsec = getObjectSystemSyslogTimeMsec1991(d.Get("syslog_time_msec").([]interface{}))
	ret.Inst.SystemChassisPortSplitEnable = d.Get("system_chassis_port_split_enable").(int)
	ret.Inst.TableIntegrity = getObjectSystemTableIntegrity1992(d.Get("table_integrity").([]interface{}))
	ret.Inst.Tcp = getObjectSystemTcp1994(d.Get("tcp").([]interface{}))
	ret.Inst.TcpStats = getObjectSystemTcpStats1997(d.Get("tcp_stats").([]interface{}))
	ret.Inst.TcpSynPerSec = getObjectSystemTcpSynPerSec1999(d.Get("tcp_syn_per_sec").([]interface{}))
	ret.Inst.TelemetryLog = getObjectSystemTelemetryLog2000(d.Get("telemetry_log").([]interface{}))
	ret.Inst.Template = getObjectSystemTemplate2006(d.Get("template").([]interface{}))
	ret.Inst.TemplateBind = getObjectSystemTemplateBind2007(d.Get("template_bind").([]interface{}))
	ret.Inst.Throughput = getObjectSystemThroughput2008(d.Get("throughput").([]interface{}))
	ret.Inst.TimeoutValue = getObjectSystemTimeoutValue2010(d.Get("timeout_value").([]interface{}))
	ret.Inst.Trunk = getObjectSystemTrunk2011(d.Get("trunk").([]interface{}))
	ret.Inst.TrunkHwHash = getObjectSystemTrunkHwHash2013(d.Get("trunk_hw_hash").([]interface{}))
	ret.Inst.TrunkXauiHwHash = getObjectSystemTrunkXauiHwHash2014(d.Get("trunk_xaui_hw_hash").([]interface{}))
	ret.Inst.Tso = getObjectSystemTso2015(d.Get("tso").([]interface{}))
	ret.Inst.Udp = getObjectSystemUdp2016(d.Get("udp").([]interface{}))
	ret.Inst.UpgradeStatus = getObjectSystemUpgradeStatus2017(d.Get("upgrade_status").([]interface{}))
	//omit uuid
	ret.Inst.VeMacScheme = getObjectSystemVeMacScheme2018(d.Get("ve_mac_scheme").([]interface{}))
	ret.Inst.XauiDlbMode = getObjectSystemXauiDlbMode2019(d.Get("xaui_dlb_mode").([]interface{}))
	return ret
}
