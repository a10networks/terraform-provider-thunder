package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystem() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system`: Configure System Parameters\n\n",
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
							ValidateFunc: validation.IntBetween(0, 128),
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
							ValidateFunc: validation.IntBetween(0, 32),
						},
					},
				},
			},
			"all_vlan_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bcast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "broadcast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"ipmcast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "IP multicast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"mcast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "multicast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"unknown_ucast": {
							Type: schema.TypeInt, Optional: true, Default: 5000, Description: "unknown unicast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"anomaly_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "log system anomalies",
				ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "total-throughput-bits-per-sec", "l4-conns-per-sec", "l7-conns-per-sec", "l7-trans-per-sec", "ssl-conns-per-sec", "ip-nat-conns-per-sec", "total-new-conns-per-sec", "total-curr-conns", "l4-bandwidth", "l7-bandwidth", "serv-ssl-conns-per-sec", "fw-conns-per-sec", "gifw-conns-per-sec"}, false),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send TCP session creation log on completion of 3-way handshake [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"msl_time": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure maximum session life, default is 2 seconds (1-39 seconds, default is 2 seconds) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 39),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable L2L3 ASIC traffic discard/drop events and Dump debug information [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"attack_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "log attack anomalies",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"bandwidth": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'input-bytes-per-sec': In Bytes per second; 'output-bytes-per-sec': Out Bytes per second;",
										ValidateFunc: validation.StringInSlice([]string{"all", "input-bytes-per-sec", "output-bytes-per-sec"}, false),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "ip_checksum_error", "udp_checksum_error", "session_not_found", "multihop_mismatch", "version_mismatch", "length_too_small", "data_is_short", "invalid_detect_mult", "invalid_multipoint", "invalid_my_disc", "invalid_ttl", "auth_length_invalid", "auth_mismatch", "auth_type_mismatch", "auth_key_id_mismatch", "auth_key_mismatch", "auth_seqnum_invalid", "auth_failed", "local_state_admin_down", "dest_unreachable", "no_ipv6_enable", "other_error"}, false),
									},
								},
							},
						},
					},
				},
			},
			"class_list_hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable class list hit count [shared partition only]",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"cm_update_file_name_ref": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_name": {
							Type: schema.TypeString, Optional: true, Description: "bind source name",
							ValidateFunc: validation.StringLenBetween(1, 63),
						},
						"dest_name": {
							Type: schema.TypeString, Optional: true, Description: "bind dest name",
							ValidateFunc: validation.StringLenBetween(1, 63),
						},
						"id": {
							Type: schema.TypeInt, ForceNew: true, Optional: true, Description: "Specify unique Partition id",
						},
					},
				},
			},
			//omit field 'control_cpu' which contains uuid only
			//omit field 'core' which contains uuid only
			//omit field 'cosq_show' which contains uuid only
			//omit field 'cosq_stats' which contains uuid only
			//omit field 'counter_lib_accounting' which contains uuid only
			"cpu_hyper_thread": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPU Hyperthreading",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU Hyperthreading",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			//omit field 'cpu_list' which contains uuid only
			"cpu_load_sharing": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU load sharing in overload situations [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"packets_per_second": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"min": {
										Type: schema.TypeInt, Optional: true, Default: 100000, Description: "Minimum packets-per-second threshold (per CPU) before redistribution will take effect (Minimum packets-per-second threshold (per CPU) before redistribution will take effect (default: 100000)) [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 30000000),
									},
								},
							},
						},
						"cpu_usage": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"low": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "CPU usage threshold (percentage) that will restore the normal packet distribution (default: 60) [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 100),
									},
									"high": {
										Type: schema.TypeInt, Optional: true, Default: 75, Description: "CPU usage threshold (percentage) that will trigger the redistribution (default: 75) [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 100),
									},
								},
							},
						},
						"tcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new TCP sessions [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"udp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new UDP sessions [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			//omit field 'cpu_map' which contains uuid only
			//omit field 'data_cpu' which contains uuid only
			"ddos_attack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "System DDoS Attack [shared partition only]",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"ddos_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "log DDoS attack anomalies",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"deep_hrxq": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: " [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			"del_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
							ValidateFunc: validation.IntBetween(0, 32),
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
							ValidateFunc: validation.IntBetween(0, 128),
						},
					},
				},
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'slb_req': No. of requests; 'slb_resp': No. of responses; 'slb_no_resp': No. of requests with no response; 'slb_req_rexmit': No. of requests retransmit; 'slb_resp_no_match': No. of requests and responses with no match; 'slb_no_resource': No. of resource failures; 'nat_req': (NAT) No. of requests; 'nat_resp': (NAT) No. of responses; 'nat_no_resp': (NAT) No. of resource failures; 'nat_req_rexmit': (NAT) No. of request retransmits; 'nat_resp_no_match': (NAT) No. of requests with no response; 'nat_no_resource': (NAT) No. of resource failures; 'nat_xid_reused': (NAT) No. of requests reusing a transaction id; 'filter_type_drop': Total Query Type Drop; 'filter_class_drop': Total Query Class Drop; 'filter_type_any_drop': Total Query ANY Type Drop; 'slb_dns_client_ssl_succ': No. of client ssl success; 'slb_dns_server_ssl_succ': No. of server ssl success; 'slb_dns_udp_conn': No. of backend udp connections; 'slb_dns_udp_conn_succ': No. of backend udp conn established; 'slb_dns_padding_to_server_removed': some help string; 'slb_dns_padding_to_client_added': some help string; 'slb_dns_edns_subnet_to_server_removed': some help string; 'slb_dns_udp_retransmit': some help string; 'slb_dns_udp_retransmit_fail': some help string; 'rpz_action_drop': RPZ Action Drop; 'rpz_action_pass_thru': RPZ Action Pass Through; 'rpz_action_tcp_only': RPZ Action TCP Only; 'rpz_action_nxdomain': RPZ Action NXDOMAIN; 'rpz_action_nodata': RPZ Action NODATA; 'rpz_action_local_data': RPZ Action Local Data; 'slb_drop': DNS requests drop; 'nat_slb_drop': (NAT)DNS requests drop; 'invalid_q_len_to_udp': invalid query length to conver to UDP;",
										ValidateFunc: validation.StringInSlice([]string{"all", "slb_req", "slb_resp", "slb_no_resp", "slb_req_rexmit", "slb_resp_no_match", "slb_no_resource", "nat_req", "nat_resp", "nat_no_resp", "nat_req_rexmit", "nat_resp_no_match", "nat_no_resource", "nat_xid_reused", "filter_type_drop", "filter_class_drop", "filter_type_any_drop", "slb_dns_client_ssl_succ", "slb_dns_server_ssl_succ", "slb_dns_udp_conn", "slb_dns_udp_conn_succ", "slb_dns_padding_to_server_removed", "slb_dns_padding_to_client_added", "slb_dns_edns_subnet_to_server_removed", "slb_dns_udp_retransmit", "slb_dns_udp_retransmit_fail", "rpz_action_drop", "rpz_action_pass_thru", "rpz_action_tcp_only", "rpz_action_nxdomain", "rpz_action_nodata", "rpz_action_local_data", "slb_drop", "nat_slb_drop", "invalid_q_len_to_udp"}, false),
									},
								},
							},
						},
						"recursive_nameserver": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"follow_shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the configured name servers of shared partition [private partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"server_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv4 server address",
													ValidateFunc: validation.IsIPv4Address,
												},
												"v4_desc": {
													Type: schema.TypeString, Optional: true, Description: "Description for this ipv4 address",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv6 server address",
													ValidateFunc: validation.IsIPv6Address,
												},
												"v6_desc": {
													Type: schema.TypeString, Optional: true, Description: "Description for this ipv6 address",
													ValidateFunc: validation.StringLenBetween(1, 63),
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_q': Total query; 'total_r': Total server response; 'hit': Total cache hit; 'bad_q': Query not passed; 'encode_q': Query encoded; 'multiple_q': Query with multiple questions; 'oversize_q': Query exceed cache size; 'bad_r': Response not passed; 'oversize_r': Response exceed cache size; 'encode_r': Response encoded; 'multiple_r': Response with multiple questions; 'answer_r': Response with multiple answers; 'ttl_r': Response with short TTL; 'ageout': Total aged out; 'bad_answer': Bad Answer; 'ageout_weight': Total aged for lower weight; 'total_log': Total stats log sent; 'total_alloc': Total allocated; 'total_freed': Total freed; 'current_allocate': Current allocate; 'current_data_allocate': Current data allocate; 'resolver_queue_full': Resolver task queue full; 'truncated_r': Response with Truncation bit set;",
										ValidateFunc: validation.StringInSlice([]string{"all", "total_q", "total_r", "hit", "bad_q", "encode_q", "multiple_q", "oversize_q", "bad_r", "oversize_r", "encode_r", "multiple_r", "answer_r", "ttl_r", "ageout", "bad_answer", "ageout_weight", "total_log", "total_alloc", "total_freed", "current_allocate", "current_data_allocate", "resolver_queue_full", "truncated_r"}, false),
									},
								},
							},
						},
					},
				},
			},
			"domain_list_hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable class list hit count",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			//omit field 'domain_list_info' which contains uuid only
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
										ValidateFunc: validation.StringInSlice([]string{"all", "pkt-drop", "pkt-lnk-down-drop", "err-pkt-drop", "rx-err", "tx-err", "tx-drop", "rx-len-err", "rx-over-err", "rx-crc-err", "rx-frame-err", "rx-no-buff-err", "rx-miss-err", "tx-abort-err", "tx-carrier-err", "tx-fifo-err", "tx-hbeat-err", "tx-windows-err", "rx-long-len-err", "rx-short-len-err", "rx-align-err", "rx-csum-offload-err", "io-rx-que-drop", "io-tx-que-drop", "io-ring-drop", "w-tx-que-drop", "w-link-down-drop", "w-ring-drop"}, false),
									},
								},
							},
						},
					},
				},
			},
			"dynamic_service_dns_socket_pool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable socket pool for dynamic-service DNS",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			//omit field 'environment' which contains uuid only
			"fpga_core_crc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable FPGA Core CRC error monitoring and act on it [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"reboot_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system reboot if system encounters FPGA Core CRC error [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "mrx-drop", "hrx-drop", "siz-drop", "fcs-drop", "land-drop", "empty-frag-drop", "mic-frag-drop", "ipv4-opt-drop", "ipv4-frag", "bad-ip-hdr-len", "bad-ip-flags-drop", "bad-ip-ttl-drop", "no-ip-payload-drop", "oversize-ip-payload", "bad-ip-payload-len", "bad-ip-frag-offset", "bad-ip-chksum-drop", "icmp-pod-drop", "tcp-bad-urg-offet", "tcp-short-hdr", "tcp-bad-ip-len", "tcp-null-flags", "tcp-null-scan", "tcp-fin-sin", "tcp-xmas-flags", "tcp-xmas-scan", "tcp-syn-frag", "tcp-frag-hdr", "tcp-bad-chksum", "udp-short-hdr", "udp-bad-ip-len", "udp-kb-frags", "udp-port-lb", "udp-bad-chksum", "runt-ip-hdr", "runt-tcpudp-hdr", "tun-mismatch", "qdr-drop"}, false),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable application memory pool [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"application_flow": {
							Type: schema.TypeInt, Optional: true, Description: "Number of flows [shared partition only]",
						},
						"basic_dpi_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable basic dpi [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"geo_db_hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Geolocation database hit count [shared partition only]",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"geo_location": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_iana": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Load built-in IANA Database [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"geo_location_geolite2_city": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-City database. Database available from http://www.maxmind.com [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"geolite2_city_include_ipv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"geo_location_geolite2_country": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-Country database. Database available from http://www.maxmind.com [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"geolite2_country_include_ipv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"geoloc_load_file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_location_load_filename": {
										Type: schema.TypeString, Optional: true, Description: "Specify file to be loaded [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"template_name": {
										Type: schema.TypeString, Optional: true, Description: "CSV template to load this file [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 63),
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
										Type: schema.TypeString, Required: true, Description: "Specify geo-location name, section range is (1-15) [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"geo_locn_multiple_addresses": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"first_ip_address": {
													Type: schema.TypeString, Optional: true, Description: "Specify IP information (Specify IP address) [shared partition only]",
													ValidateFunc: validation.IsIPv4Address,
												},
												"geol_ipv4_mask": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv4 mask [shared partition only]",
													ValidateFunc:  axapi.IsIpv4NetmaskBrief,
												},
												"ip_addr2": {
													Type: schema.TypeString, Optional: true, Description: "Specify IP address range [shared partition only]",
													ValidateFunc:  validation.IsIPv4Address,
												},
												"first_ipv6_address": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address [shared partition only]",
													ValidateFunc: validation.IsIPv6Address,
												},
												"geol_ipv6_mask": {
													Type: schema.TypeInt, Optional: true, Description: "Specify IPv6 mask [shared partition only]",
													ValidateFunc:  validation.IntBetween(0, 128),
												},
												"ipv6_addr2": {
													Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address range [shared partition only]",
													ValidateFunc:  validation.IsIPv6Address,
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 127),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "place-holder"}, false),
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
							ValidateFunc: validation.StringLenBetween(1, 63),
						},
						"shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sharing with other partitions [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"include_geoloc_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include_geoloc_name_val": {
										Type: schema.TypeString, Optional: true, Description: "Geolocation name to add",
										ValidateFunc: validation.StringLenBetween(1, 127),
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
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
							ValidateFunc: validation.StringLenBetween(1, 127),
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-count': some help string; 'total-geoloc': some help string; 'total-active': some help string;",
										ValidateFunc: validation.StringInSlice([]string{"all", "hit-count", "total-geoloc", "total-active"}, false),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "place-holder"}, false),
									},
								},
							},
						},
					},
				},
			},
			// omit "geolocation_file"
			"glid": {
				Type: schema.TypeInt, Optional: true, Description: "Apply limits to the whole system [shared partition only]",
				ValidateFunc: validation.IntBetween(1, 1023),
			},
			//omit field 'guest_file' which contains uuid only
			//omit field 'gui_image_list' which contains uuid only
			//omit field 'hardware' which contains uuid only
			"hardware_forward": {
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-counts': Total packts hit counts; 'hit-index': HW Fwd hit index; 'ipv4-forward-counts': Total IPv4 hardware forwarded packets; 'ipv6-forward-counts': Total IPv6 hardware forwarded packets; 'hw-fwd-module-status': hardware forwarder status flags; 'hw-fwd-prog-reqs': hardware forward programming requests; 'hw-fwd-prog-errors': hardware forward programming Errors; 'hw-fwd-flow-singlebit-errors': hardware forward singlebit Errors; 'hw-fwd-flow-tag-mismatch': hardware forward tag mismatch errors; 'hw-fwd-flow-seq-mismatch': hardware forward sequence mismatch errors; 'hw-fwd-ageout-drop-count': hardware forward ageout drop count; 'hw-fwd-invalidation-drop': hardware forward invalid drop count; 'hw-fwd-flow-hit-index': hardware forward flow hit index; 'hw-fwd-flow-reason-flags': hardware forward flow reason flags; 'hw-fwd-flow-drop-count': hardware forward flow drop count; 'hw-fwd-flow-error-count': hardware forward flow error count; 'hw-fwd-flow-unalign-count': hardware forward flow unalign count; 'hw-fwd-flow-underflow-count': hardware forward flow underflow count; 'hw-fwd-flow-tx-full-drop': hardware forward flow tx full drop count; 'hw-fwd-flow-qdr-full-drop': hardware forward flow qdr full drop count; 'hw-fwd-phyport-mismatch-drop': hardware forward phyport mismatch count; 'hw-fwd-vlanid-mismatch-drop': hardware forward vlanid mismatch count; 'hw-fwd-vmid-drop': hardware forward vmid mismatch count; 'hw-fwd-protocol-mismatch-drop': hardware forward protocol mismatch count; 'hw-fwd-avail-ipv4-entry': hardware forward available ipv4 entries count; 'hw-fwd-avail-ipv6-entry': hardware forward available ipv6 entries count;",
										ValidateFunc: validation.StringInSlice([]string{"all", "hit-counts", "hit-index", "ipv4-forward-counts", "ipv6-forward-counts", "hw-fwd-module-status", "hw-fwd-prog-reqs", "hw-fwd-prog-errors", "hw-fwd-flow-singlebit-errors", "hw-fwd-flow-tag-mismatch", "hw-fwd-flow-seq-mismatch", "hw-fwd-ageout-drop-count", "hw-fwd-invalidation-drop", "hw-fwd-flow-hit-index", "hw-fwd-flow-reason-flags", "hw-fwd-flow-drop-count", "hw-fwd-flow-error-count", "hw-fwd-flow-unalign-count", "hw-fwd-flow-underflow-count", "hw-fwd-flow-tx-full-drop", "hw-fwd-flow-qdr-full-drop", "hw-fwd-phyport-mismatch-drop", "hw-fwd-vlanid-mismatch-drop", "hw-fwd-vmid-drop", "hw-fwd-protocol-mismatch-drop", "hw-fwd-avail-ipv4-entry", "hw-fwd-avail-ipv6-entry"}, false),
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
													ValidateFunc: validation.StringInSlice([]string{"all", "entry-create", "entry-create-failure", "entry-create-fail-server-down", "entry-create-fail-max-entry", "entry-free", "entry-free-opp-entry", "entry-free-no-hw-prog", "entry-free-no-conn", "entry-free-no-sw-entry", "entry-counter", "entry-age-out", "entry-age-out-idle", "entry-age-out-tcp-fin", "entry-age-out-tcp-rst", "entry-age-out-invalid-dst", "entry-force-hw-invalidate", "entry-invalidate-server-down", "tcam-create", "tcam-free", "tcam-counter"}, false),
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
			"high_memory_l4_session": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable high memory l4 session support [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			//omit field 'hrxq_status' which contains uuid only
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
										ValidateFunc: validation.StringInSlice([]string{"all", "num", "inmsgs", "inerrors", "indestunreachs", "intimeexcds", "inparmprobs", "insrcquenchs", "inredirects", "inechos", "inechoreps", "intimestamps", "intimestampreps", "inaddrmasks", "inaddrmaskreps", "outmsgs", "outerrors", "outdestunreachs", "outtimeexcds", "outparmprobs", "outsrcquenchs", "outredirects", "outechos", "outechoreps", "outtimestamps", "outtimestampreps", "outaddrmasks", "outaddrmaskreps"}, false),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "over_limit_drop", "limit_intf_drop", "limit_vserver_drop", "limit_total_drop", "lockup_time_left", "curr_rate", "v6_over_limit_drop", "v6_limit_intf_drop", "v6_limit_vserver_drop", "v6_limit_total_drop", "v6_lockup_time_left", "v6_curr_rate"}, false),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "in_msgs", "in_errors", "in_dest_un_reach", "in_pkt_too_big", "in_time_exceeds", "in_param_prob", "in_echoes", "in_exho_reply", "in_grp_mem_query", "in_grp_mem_resp", "in_grp_mem_reduction", "in_router_sol", "in_ra", "in_ns", "in_na", "in_redirect", "out_msg", "out_dst_un_reach", "out_pkt_too_big", "out_time_exceeds", "out_param_prob", "out_echo_req", "out_echo_replies", "out_rs", "out_ra", "out_ns", "out_na", "out_redirects", "out_mem_resp", "out_mem_reductions", "err_rs", "err_ra", "err_ns", "err_na", "err_redirects", "err_echoes", "err_echo_replies"}, false),
									},
								},
							},
						},
					},
				},
			},
			//omit field 'inuse_cpu_list' which contains uuid only
			//omit field 'inuse_port_list' which contains uuid only
			"io_cpu": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_cores": {
							Type: schema.TypeInt, Optional: true, Description: "max number of IO cores (Specify number of cores)",
							ValidateFunc: validation.IntBetween(0, 128),
						},
					},
				},
			},
			//omit field 'ip_dns_cache' which contains uuid only
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
										ValidateFunc: validation.StringInSlice([]string{"all", "inreceives", "inhdrerrors", "intoobigerrors", "innoroutes", "inaddrerrors", "inunknownprotos", "intruncatedpkts", "indiscards", "indelivers", "outforwdatagrams", "outrequests", "outdiscards", "outnoroutes", "reasmtimeout", "reasmreqds", "reasmoks", "reasmfails", "fragoks", "fragfails", "fragcreates", "inmcastpkts", "outmcastpkts"}, false),
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_hit_count_in_sw': Packet Hit Count in SW; 'packet_hit_count_in_spe': Packet Hit Count in SPE; 'entries_added_in_sw': Entries Added in SW; 'entries_removed_from_sw': Entries Removed from SW; 'entries_added_in_spe': Entries Added in SPE; 'entries_removed_from_spe': Entries Removed from SPE; 'error_out_of_memory': Out of memory Error; 'error_out_of_spe_entries': Out of SPE Entries Error; [shared partition only]",
										ValidateFunc: validation.StringInSlice([]string{"all", "packet_hit_count_in_sw", "packet_hit_count_in_spe", "entries_added_in_sw", "entries_removed_from_sw", "entries_added_in_spe", "entries_removed_from_spe", "error_out_of_memory", "error_out_of_spe_entries"}, false),
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
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name) [shared partition only]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 8),
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
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name) [shared partition only]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 8),
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
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name) [shared partition only]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 8),
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
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name) [shared partition only]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 8),
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
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name) [shared partition only]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 8),
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
									"class_list_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name) [shared partition only]",
													ValidateFunc: validation.StringLenBetween(1, 63),
												},
												"ip_threat_action_tmpl": {
													Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 8),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "inreceives", "inhdrerrors", "intoobigerrors", "innoroutes", "inaddrerrors", "inunknownprotos", "intruncatedpkts", "indiscards", "indelivers", "outforwdatagrams", "outrequests", "outdiscards", "outnoroutes", "reasmtimeout", "reasmreqds", "reasmoks", "reasmfails", "fragoks", "fragfails", "fragcreates", "inmcastpkts", "outmcastpkts"}, false),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset IPMI Controller [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"ip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type: schema.TypeString, Optional: true, Description: "IP address [shared partition only]",
										ValidateFunc: validation.IsIPv4Address,
									},
									"ipv4_netmask": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask [shared partition only]",
										ValidateFunc: axapi.IsIpv4NetmaskBrief,
									},
									"default_gateway": {
										Type: schema.TypeString, Optional: true, Description: "Default gateway address [shared partition only]",
										ValidateFunc: validation.IsIPv4Address,
									},
								},
							},
						},
						"ipsrc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dhcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP addr obtained by BMC running DHCP [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"static": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually configured static IP address [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
								},
							},
						},
						"user": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"add": {
										Type: schema.TypeString, Optional: true, Description: "Add a new IPMI user (IPMI User Name) [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 16),
									},
									"password": {
										Type: schema.TypeString, Optional: true, Description: "Password [shared partition only]",
									},
									"administrator": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Full control [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"callback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lowest privilege level [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"operator": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most BMC commands are allowed [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"user": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only 'benign' commands are allowed [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"disable": {
										Type: schema.TypeString, Optional: true, Description: "Disable an existing IPMI user (IPMI User Name) [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 16),
									},
									"privilege": {
										Type: schema.TypeString, Optional: true, Description: "Change an existing IPMI user privilege (IPMI User Name) [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 16),
									},
									"setname": {
										Type: schema.TypeString, Optional: true, Description: "Change User Name (Current IPMI User Name) [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 16),
									},
									"newname": {
										Type: schema.TypeString, Optional: true, Description: "New IPMI User Name [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 16),
									},
									"setpass": {
										Type: schema.TypeString, Optional: true, Description: "Change Password (IPMI User Name) [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 16),
									},
									"newpass": {
										Type: schema.TypeString, Optional: true, Description: "New Password [shared partition only]",
									},
								},
							},
						},
						"tool": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cmd": {
										Type: schema.TypeString, Optional: true, Description: "Command to execute in double quotes [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable IPMI on platform [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet round robin for IPsec packets [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"crypto_core": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Crypto cores assigned for IPsec processing [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 64),
						},
						"crypto_mem": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Crypto memory percentage assigned for IPsec processing (rounded to increments of 10) [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 100),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"fpga_decrypt": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable FPGA decryption offload; 'disable': Disable FPGA decryption offload; [shared partition only]",
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
									},
								},
							},
						},
					},
				},
			},
			"ipv6_prefix_length": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Length of IPv6 prefix used to determine the user-group and the PU, by default 128 [shared partition only]",
				ValidateFunc: validation.IntBetween(16, 128),
			},
			"link_capability": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable link capabilities [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Link Monitoring",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Large Receive Offload",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"non_dedicated": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set management interface in non-dedicated mode",
							ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "usage-percentage"}, false),
									},
								},
							},
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
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"second_factor": {
							Type: schema.TypeString, Optional: true, Description: "Input second factor paramter",
							ValidateFunc: validation.StringLenBetween(1, 128),
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
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "'tftp': Use tftp for connection; 'ftp': Use ftp for connection; 'scp': Use scp for connection; 'http': Use http for connection; 'https': Use https for connection; 'sftp': Use sftp for connection;",
							ValidateFunc: validation.StringInSlice([]string{"tftp", "ftp", "scp", "http", "https", "sftp"}, false),
						},
						"cert_store_path": {
							Type: schema.TypeString, Optional: true, Description: "Configure certificate store path",
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"username": {
							Type: schema.TypeString, Optional: true, Description: "Certificate store host username",
							ValidateFunc: validation.StringLenBetween(1, 31),
						},
						"passwd_string": {
							Type: schema.TypeString, Optional: true, Description: "Certificate store host password",
						},
						//omit encrypted field: 'encrypted'
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
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.StringLenBetween(1, 128),
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
							ValidateFunc: validation.IntBetween(0, 32),
						},
						"mac_address": {
							Type: schema.TypeString, Optional: true, Description: "mac-address to be configured as mgmt port",
							ValidateFunc: validation.StringLenBetween(14, 14),
						},
						"pci_address": {
							Type: schema.TypeString, Optional: true, Description: "pci-address to be configured as mgmt port",
							ValidateFunc: validation.StringLenBetween(1, 63),
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
							ValidateFunc: validation.IntBetween(0, 32),
						},
						"port_number": {
							Type: schema.TypeInt, Optional: true, Description: "port number to be configured (Specify port number)",
							ValidateFunc: validation.IntBetween(1, 32),
						},
					},
				},
			},
			"module_ctrl_cpu": {
				Type: schema.TypeString, Optional: true, Description: "'high': high cpu usage; 'low': low cpu usage; 'medium': medium cpu usage; [shared partition only]",
				ValidateFunc: validation.StringInSlice([]string{"high", "low", "medium"}, false),
			},
			"mon_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt, Required: true, Description: "Monitor template ID Number [shared partition only]",
										ValidateFunc: validation.IntBetween(1, 16),
									},
									"clear_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sessions": {
													Type: schema.TypeString, Optional: true, Description: "'all': Clear all sessions; 'sequence': Sequence number; [shared partition only]",
													ValidateFunc: validation.StringInSlice([]string{"all", "sequence"}, false),
												},
												"clear_all_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
												"clear_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
											},
										},
									},
									"link_disable_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"diseth": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"dis_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
											},
										},
									},
									"link_enable_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enaeth": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"ena_sequence": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
											},
										},
									},
									"monitor_relation": {
										Type: schema.TypeString, Optional: true, Default: "monitor-and", Description: "'monitor-and': Configures the monitors in current template to work with AND logic; 'monitor-or': Configures the monitors in current template to work with OR logic; [shared partition only]",
										ValidateFunc: validation.StringInSlice([]string{"monitor-and", "monitor-or"}, false),
									},
									"link_up_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkup_ethernet1": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"link_up_sequence1": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
												"linkup_ethernet2": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"link_up_sequence2": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
												"linkup_ethernet3": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"link_up_sequence3": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequece number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
											},
										},
									},
									"link_down_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkdown_ethernet1": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"link_down_sequence1": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
												"linkdown_ethernet2": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"link_down_sequence2": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the seqeuence number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
												"linkdown_ethernet3": {
													Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number) [shared partition only]",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"link_down_sequence3": {
													Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 16),
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
								},
							},
						},
						"link_block_as_down": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: " [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
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
										Type: schema.TypeInt, Optional: true, Default: 0, Description: " [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Multi-Queue-Support [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "good_recv", "periodic_sent", "rate_limit", "bad_hop_limit", "truncated", "bad_icmpv6_csum", "bad_icmpv6_code", "bad_icmpv6_option", "l2_addr_and_unspec", "no_free_buffers"}, false),
									},
								},
							},
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
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Min length:8, Min Lower Case:2, Min Upper Case:2, Min Numbers:2, Min Special Character:1; 'Medium': Medium: Min length:6, Min Lower Case:2, Min Upper Case:2, Min Numbers:1, Min Special Character:1; 'Simple': Simple: Min length:4, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:0; [shared partition only]",
							ValidateFunc: validation.StringInSlice([]string{"Strict", "Medium", "Simple"}, false),
						},
						"aging": {
							Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Max Age-60 Days; 'Medium': Medium: Max Age- 90 Days; 'Simple': Simple: Max Age-120 Days; [shared partition only]",
							ValidateFunc: validation.StringInSlice([]string{"Strict", "Medium", "Simple"}, false),
						},
						"history": {
							Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Does not allow upto 5 old passwords; 'Medium': Medium: Does not allow upto 4 old passwords; 'Simple': Simple: Does not allow upto 3 old passwords; [shared partition only]",
							ValidateFunc: validation.StringInSlice([]string{"Strict", "Medium", "Simple"}, false),
						},
						"min_pswd_len": {
							Type: schema.TypeInt, Optional: true, Description: "Configure custom password length [shared partition only]",
							ValidateFunc: validation.IntBetween(8, 63),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"per_vlan_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bcast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "broadcast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"ipmcast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "IP multicast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"mcast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "multicast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"unknown_ucast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "unknown unicast packets (per second limit) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			//omit field 'platformtype' which contains uuid only
			//omit field 'port_info' which contains uuid only
			//omit field 'port_list' which contains uuid only
			"ports": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_detection_interval": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Link detection interval in msecs [shared partition only]",
							ValidateFunc: validation.IntBetween(50, 60000),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			//omit empty field 'probe_network_devices'
			"promiscuous_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run in promiscous mode settings [shared partition only]",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			//omit field 'psu_info' which contains uuid only
			"q_in_q": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inner_tpid": {
							Type: schema.TypeString, Optional: true, Description: "TPID for inner VLAN (Inner TPID, 16 bit hex value, default is 8100) [shared partition only]",
							ValidateFunc: validation.StringLenBetween(1, 4),
						},
						"outer_tpid": {
							Type: schema.TypeString, Optional: true, Description: "TPID for outer VLAN (Outer TPID, 16 bit hex value, default is 8100) [shared partition only]",
							ValidateFunc: validation.StringLenBetween(1, 4),
						},
						"enable_all_ports": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 802.1QinQ on all physical ports [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable micro-burst traffic support [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.IntBetween(1024, 65535),
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
																ValidateFunc: validation.StringLenBetween(1, 63),
															},
															"ip_list_secret": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure shared secret",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"ip_list_secret_string": {
																Type: schema.TypeString, Optional: true, Description: "The RADIUS secret",
															},
															//omit encrypted field: 'ip_list_encrypted'
														},
													},
												},
											},
										},
									},
									"secret": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure shared secret",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"secret_string": {
										Type: schema.TypeString, Optional: true, Description: "The RADIUS secret",
									},
									//omit encrypted field: 'encrypted'
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "Join a VRRP-A failover group",
										ValidateFunc: validation.IntBetween(1, 31),
									},
									"attribute": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_value": {
													Type: schema.TypeString, Optional: true, Description: "'inside-ipv6-prefix': Framed IPv6 Prefix; 'inside-ip': Inside IP address; 'inside-ipv6': Inside IPv6 address; 'imei': International Mobile Equipment Identity (IMEI); 'imsi': International Mobile Subscriber Identity (IMSI); 'msisdn': Mobile Subscriber Integrated Services Digital Network-Number (MSISDN); 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3; 'custom4': Customized attribute 4; 'custom5': Customized attribute 5; 'custom6': Customized attribute 6;",
													ValidateFunc: validation.StringInSlice([]string{"inside-ipv6-prefix", "inside-ip", "inside-ipv6", "imei", "imsi", "msisdn", "custom1", "custom2", "custom3", "custom4", "custom5", "custom6"}, false),
												},
												"prefix_length": {
													Type: schema.TypeString, Optional: true, Description: "'32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;",
													ValidateFunc: validation.StringInSlice([]string{"32", "48", "64", "80", "96", "112"}, false),
												},
												"prefix_vendor": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"prefix_number": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"name": {
													Type: schema.TypeString, Optional: true, Description: "Customized attribute name",
													ValidateFunc: validation.StringLenBetween(1, 15),
												},
												"value": {
													Type: schema.TypeString, Optional: true, Description: "'hexadecimal': Type of attribute value is hexadecimal;",
													ValidateFunc: validation.StringInSlice([]string{"hexadecimal"}, false),
												},
												"custom_vendor": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"custom_number": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"vendor": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"number": {
													Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
													ValidateFunc: validation.IntBetween(1, 255),
												},
											},
										},
									},
									"disable_reply": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle option for RADIUS reply packet(Default: Accounting response will be sent)",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"accounting_start": {
										Type: schema.TypeString, Optional: true, Default: "append-entry", Description: "'ignore': Ignore; 'append-entry': Append the AVPs to existing entry (default); 'replace-entry': Replace the AVPs of existing entry;",
										ValidateFunc: validation.StringInSlice([]string{"ignore", "append-entry", "replace-entry"}, false),
									},
									"accounting_stop": {
										Type: schema.TypeString, Optional: true, Default: "delete-entry", Description: "'ignore': Ignore; 'delete-entry': Delete the entry (default); 'delete-entry-and-sessions': Delete the entry and data sessions associated(CGN only);",
										ValidateFunc: validation.StringInSlice([]string{"ignore", "delete-entry", "delete-entry-and-sessions"}, false),
									},
									"accounting_interim_update": {
										Type: schema.TypeString, Optional: true, Default: "ignore", Description: "'ignore': Ignore (default); 'append-entry': Append the AVPs to existing entry; 'replace-entry': Replace the AVPs of existing entry;",
										ValidateFunc: validation.StringInSlice([]string{"ignore", "append-entry", "replace-entry"}, false),
									},
									"accounting_on": {
										Type: schema.TypeString, Optional: true, Default: "ignore", Description: "'ignore': Ignore (default); 'delete-entries-using-attribute': Delete entries matching attribute in RADIUS Table;",
										ValidateFunc: validation.StringInSlice([]string{"ignore", "delete-entries-using-attribute"}, false),
									},
									"attribute_name": {
										Type: schema.TypeString, Optional: true, Description: "'msisdn': Clear using MSISDN; 'imei': Clear using IMEI; 'imsi': Clear using IMSI; 'custom1': Clear using CUSTOM1 attribute configured; 'custom2': Clear using CUSTOM2 attribute configured; 'custom3': Clear using CUSTOM3 attribute configured; 'custom4': Clear using CUSTOM4 attribute configured; 'custom5': Clear using CUSTOM5 attribute configured; 'custom6': Clear using CUSTOM6 attribute configured;",
										ValidateFunc:  validation.StringInSlice([]string{"msisdn", "imei", "imsi", "custom1", "custom2", "custom3", "custom4", "custom5", "custom6"}, false),
										ConflictsWith: []string{"radius.0.server.0.custom_attribute_name"},
									},
									"custom_attribute_name": {
										Type: schema.TypeString, Optional: true, Description: "Clear using customized attribute",
										ValidateFunc:  validation.StringLenBetween(1, 15),
										ConflictsWith: []string{"radius.0.server.0.attribute_name"},
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
													ValidateFunc: validation.StringInSlice([]string{"all", "msisdn-received", "imei-received", "imsi-received", "custom-received", "radius-request-received", "radius-request-dropped", "request-bad-secret-dropped", "request-no-key-vap-dropped", "request-malformed-dropped", "request-ignored", "radius-table-full", "secret-not-configured-dropped", "ha-standby-dropped", "ipv6-prefix-length-mismatch", "invalid-key", "smp-created", "smp-deleted", "smp-mem-allocated", "smp-mem-alloc-failed", "smp-mem-freed", "smp-in-rml", "mem-allocated", "mem-alloc-failed", "mem-freed", "ha-sync-create-sent", "ha-sync-delete-sent", "ha-sync-create-recv", "ha-sync-delete-recv", "acct-on-filters-full", "acct-on-dup-request", "ip-mismatch-delete", "ip-add-race-drop", "ha-sync-no-key-vap-dropped", "inter-card-msg-fail-drop", "radius-packets-redirected", "radius-packets-redirect-fail-dropped", "radius-packets-process-local", "radius-packets-dropped-not-lo", "radius-inter-card-dup-redir"}, false),
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
			//omit field 'reboot' which contains uuid only
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
										Type: schema.TypeString, Required: true, Description: "Enter template name [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag [shared partition only]",
										ValidateFunc: validation.StringLenBetween(1, 127),
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
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-device allowed (gslb-device count (default is max-value)) [shared partition only]",
															},
															"gslb_device_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_geo_location_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_geo_location_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-geo-location allowed (gslb-geo-location count (default is max-value)) [shared partition only]",
															},
															"gslb_geo_location_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_ip_list_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_ip_list_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-ip-list allowed (gslb-ip-list count (default is max-value)) [shared partition only]",
															},
															"gslb_ip_list_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_policy_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_policy_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-policy allowed (gslb-policy count (default is max-value)) [shared partition only]",
															},
															"gslb_policy_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_service_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service allowed (gslb-service count (default is max-value) [shared partition only]",
															},
															"gslb_service_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_service_ip_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_ip_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-ip allowed (gslb-service-ip count (default is max-value)) [shared partition only]",
															},
															"gslb_service_ip_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_service_port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_port_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-port allowed ( gslb-service-port count (default is max-value)) [shared partition only]",
															},
															"gslb_service_port_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_site_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_site_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-site allowed (gslb-site count (default is max-value)) [shared partition only]",
															},
															"gslb_site_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_svc_group_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_svc_group_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-svc-group allowed (gslb-svc-group count (default is max-value)) [shared partition only]",
															},
															"gslb_svc_group_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-template allowed (gslb-template count (default is max-value)) [shared partition only]",
															},
															"gslb_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"gslb_zone_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_zone_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-zone allowed (gslb-zone count (default is max-value)) [shared partition only]",
															},
															"gslb_zone_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"health_monitor_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"health_monitor_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of health monitors allowed (health-monitor count (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 1023),
															},
															"health_monitor_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 1023),
															},
														},
													},
												},
												"real_port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"real_port_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-ports allowed (real-port count (default is max-value)) [shared partition only]",
															},
															"real_port_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"real_server_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"real_server_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-servers allowed (real-server count (default is max-value)) [shared partition only]",
															},
															"real_server_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"service_group_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"service_group_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of service groups allowed (service-group count (default is max-value)) [shared partition only]",
															},
															"service_group_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"virtual_server_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"virtual_server_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-servers allowed (virtual-server count (default is max-value)) [shared partition only]",
															},
															"virtual_server_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"virtual_port_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"virtual_port_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-port allowed (virtual-port count (default is max-value)) [shared partition only]",
															},
															"virtual_port_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"cache_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cache_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of cache-template allowed (cache-template count (default is max-value)) [shared partition only]",
															},
															"cache_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"client_ssl_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"client_ssl_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of client-ssl-template allowed (client-ssl-template count (default is max-value)) [shared partition only]",
															},
															"client_ssl_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"conn_reuse_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"conn_reuse_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of conn-reuse-template allowed (conn-reuse-template count (default is max-value)) [shared partition only]",
															},
															"conn_reuse_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"fast_tcp_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fast_tcp_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-tcp-template allowed (fast-tcp-template count (default is max-value)) [shared partition only]",
															},
															"fast_tcp_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"fast_udp_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fast_udp_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-udp-template allowed (fast-udp-template count (default is max-value)) [shared partition only]",
															},
															"fast_udp_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"fix_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fix_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of fix-template allowed (fix-template count (default is max-value)) [shared partition only]",
															},
															"fix_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"http_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"http_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of http-template allowed (http-template count (default is max-value)) [shared partition only]",
															},
															"http_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"link_cost_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"link_cost_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of link-cost-template allowed (link-cost-template count (default is max-value)) [shared partition only]",
															},
															"link_cost_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"persist_cookie_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"persist_cookie_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-cookie-template allowed (persist-cookie-template count (default is max-value)) [shared partition only]",
															},
															"persist_cookie_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"persist_srcip_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"persist_srcip_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-srcip-template allowed (persist-source-ip-template count (default is max-value)) [shared partition only]",
															},
															"persist_srcip_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"server_ssl_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server_ssl_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of server-ssl-template allowed (server-ssl-template count (default is max-value)) [shared partition only]",
															},
															"server_ssl_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"proxy_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"proxy_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of proxy-template allowed (server-ssl-template count (default is max-value)) [shared partition only]",
															},
															"proxy_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"stream_template_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"stream_template_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of stream-template allowed (server-ssl-template count (default is max-value)) [shared partition only]",
															},
															"stream_template_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%)) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 99),
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
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv4 routes allowed (Static ipv4 routes (default is max-value)) [shared partition only]",
															},
															"static_ipv4_route_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"static_ipv6_route_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_ipv6_route_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv6 routes allowed (Static ipv6 routes (default is max-value)) [shared partition only]",
															},
															"static_ipv6_route_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
															},
														},
													},
												},
												"ipv4_acl_line_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv4_acl_line_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV4 ACL lines (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 16000),
															},
															"ipv4_acl_line_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 16000),
															},
														},
													},
												},
												"ipv6_acl_line_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_acl_line_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV6 ACL lines (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 16000),
															},
															"ipv6_acl_line_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 16000),
															},
														},
													},
												},
												"static_arp_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_arp_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static arp entries allowed (Static arp (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 128),
															},
															"static_arp_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 128),
															},
														},
													},
												},
												"static_neighbor_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_neighbor_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static neighbor entries allowed (Static neighbors (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 128),
															},
															"static_neighbor_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 128),
															},
														},
													},
												},
												"static_mac_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_mac_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of static MAC entries allowed (Static MACs (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 500),
															},
															"static_mac_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 500),
															},
														},
													},
												},
												"object_group_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_group_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of object groups allowed (Object group (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 8000),
															},
															"object_group_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 8000),
															},
														},
													},
												},
												"object_group_clause_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_group_clause_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the number of object group clauses allowed (Object group clauses (default is max-value)) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 8192000),
															},
															"object_group_clause_min_guarantee": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 8192000),
															},
														},
													},
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%)) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 99),
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
																Type: schema.TypeInt, Optional: true, Description: "Enter the bandwidth limit in mbps (Bandwidth limit in Mbit/s (no limits applied by default)) [shared partition only]",
																ValidateFunc: validation.IntBetween(10, 10000000),
															},
															"bw_limit_watermark_disable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 1),
															},
														},
													},
												},
												"concurrent_session_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"concurrent_session_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the Concurrent Session limit (cps) (Concurrent-Session cps limit (no limits applied by default)) [shared partition only]",
															},
														},
													},
												},
												"l4_session_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l4_session_limit_max": {
																Type: schema.TypeString, Optional: true, Description: "Enter the l4 session limit in % (0.01% to 99.99%) (Enter a number from 0.01 to 99.99 (up to 2 digits precision)) [shared partition only]",
																ValidateFunc: validation.StringLenBetween(1, 5),
															},
															"l4_session_limit_min_guarantee": {
																Type: schema.TypeString, Optional: true, Default: "0", Description: "minimum guaranteed value in % (up to 2 digits precision) (Enter a number from 0 to 99.99) [shared partition only]",
																ValidateFunc: validation.StringLenBetween(1, 5),
															},
														},
													},
												},
												"l4cps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l4cps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the L4 cps limit (L4 cps limit (no limits applied by default)) [shared partition only]",
																ValidateFunc: validation.IntBetween(100, 1000000),
															},
														},
													},
												},
												"l7cps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l7cps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "L7cps-limit (L7 cps limit (no limits applied by default)) [shared partition only]",
																ValidateFunc: validation.IntBetween(100, 300000),
															},
														},
													},
												},
												"natcps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"natcps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the Nat cps limit (NAT cps limit (no limits applied by default)) [shared partition only]",
																ValidateFunc: validation.IntBetween(100, 1000000),
															},
														},
													},
												},
												"fwcps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fwcps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the Firewall cps limit (Firewall cps limit (no limits applied by default)) [shared partition only]",
																ValidateFunc: validation.IntBetween(100, 1000000),
															},
														},
													},
												},
												"ssl_throughput_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ssl_throughput_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the ssl throughput limit in mbps (SSL Througput limit in Mbit/s (no limits applied by default)) [shared partition only]",
																ValidateFunc: validation.IntBetween(10, 10000000),
															},
															"ssl_throughput_limit_watermark_disable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions) [shared partition only]",
																ValidateFunc: validation.IntBetween(0, 1),
															},
														},
													},
												},
												"sslcps_limit_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"sslcps_limit_max": {
																Type: schema.TypeInt, Optional: true, Description: "Enter the SSL cps limit (SSL cps limit (no limits applied by default)) [shared partition only]",
																ValidateFunc: validation.IntBetween(100, 300000),
															},
														},
													},
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%)) [shared partition only]",
													ValidateFunc: validation.IntBetween(1, 99),
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
						"ssl_context_memory": {
							Type: schema.TypeInt, Optional: true, Default: 2048, Description: "Total SSL context memory needed in units of MB. Will be rounded to closest multiple of 2MB [shared partition only]",
							ValidateFunc: validation.IntBetween(256, 16384),
						},
						"ssl_dma_memory": {
							Type: schema.TypeInt, Optional: true, Default: 256, Description: "Total SSL DMA memory needed in units of MB. Will be rounded to closest multiple of 2MB [shared partition only]",
							ValidateFunc: validation.IntBetween(32, 2048),
						},
						"nat_pool_addr_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total configurable NAT Pool addresses in the System [shared partition only]",
						},
						"l4_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total Sessions in the System [shared partition only]",
						},
						"auth_portal_html_file_size": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Specify maximum html file size for each html page in auth portal (in KB) [shared partition only]",
							ValidateFunc: validation.IntBetween(4, 120),
						},
						"auth_portal_image_file_size": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Specify maximum image file size for default portal (in KB) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 80),
						},
						"max_aflex_file_size": {
							Type: schema.TypeInt, Optional: true, Default: 32, Description: "Set maximum aFleX file size (Maximum file size in KBytes, default is 32K) [shared partition only]",
							ValidateFunc: validation.IntBetween(16, 256),
						},
						"aflex_table_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total aFleX table entry in the system (Total aFlex entry in the system) [shared partition only]",
						},
						"class_list_ipv6_addr_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total IPv6 addresses for class-list [shared partition only]",
						},
						"class_list_ac_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total entries for AC class-list [shared partition only]",
						},
						"class_list_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total entries for class-list [shared partition only]",
						},
						"max_aflex_authz_collection_number": {
							Type: schema.TypeInt, Optional: true, Default: 512, Description: "Specify the maximum number of collections supported by aFleX authorization [shared partition only]",
							ValidateFunc: validation.IntBetween(256, 4096),
						},
						"radius_table_size": {
							Type: schema.TypeInt, Optional: true, Description: "Total configurable CGNV6 RADIUS Table entries [shared partition only]",
						},
						"authz_policy_number": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the maximum number of authorization policies [shared partition only]",
						},
						"ipsec_sa_number": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the maximum number of IPsec SA [shared partition only]",
						},
						"ram_cache_memory_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the maximum memory used by ram cache [shared partition only]",
						},
						"waf_template_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total configurable WAF Templates in the System [shared partition only]",
						},
						"auth_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total auth sessions in the system [shared partition only]",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"visibility": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"monitored_entity_count": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of monitored entities for visibility [shared partition only]",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_l4_conn': Total L4 Count; 'conn_counter': Conn Count; 'conn_freed_counter': Conn Freed; 'total_l4_packet_count': Total L4 Packet Count; 'total_l7_packet_count': Total L7 Packet Count; 'total_l4_conn_proxy': Total L4 Conn Proxy Count; 'total_l7_conn': Total L7 Conn; 'total_tcp_conn': Total TCP Conn; 'curr_free_conn': Curr Free Conn; 'tcp_est_counter': TCP Established; 'tcp_half_open_counter': TCP Half Open; 'tcp_half_close_counter': TCP Half Closed; 'udp_counter': UDP Count; 'ip_counter': IP Count; 'other_counter': Non TCP/UDP IP sessions; 'reverse_nat_tcp_counter': Reverse NAT TCP; 'reverse_nat_udp_counter': Reverse NAT UDP; 'tcp_syn_half_open_counter': TCP SYN Half Open; 'conn_smp_alloc_counter': Conn SMP Alloc; 'conn_smp_free_counter': Conn SMP Free; 'conn_smp_aged_counter': Conn SMP Aged; 'ssl_count_curr': Curr SSL Count; 'ssl_count_total': Total SSL Count; 'server_ssl_count_curr': Current SSL Server Count; 'server_ssl_count_total': Total SSL Server Count; 'client_ssl_reuse_total': Total SSL Client Reuse; 'server_ssl_reuse_total': Total SSL Server Reuse; 'ssl_failed_total': Total SSL Failures; 'ssl_failed_ca_verification': SSL Cert Auth Verification Errors; 'ssl_server_cert_error': SSL Server Cert Errors; 'ssl_client_cert_auth_fail': SSL Client Cert Auth Failures; 'total_ip_nat_conn': Total IP Nat Conn; 'total_l2l3_conn': Totl L2/L3 Connections; 'client_ssl_ctx_malloc_failure': Client SSL Ctx malloc Failures; 'conn_type_0_available': Conn Type 0 Available; 'conn_type_1_available': Conn Type 1 Available; 'conn_type_2_available': Conn Type 2 Available; 'conn_type_3_available': Conn Type 3 Available; 'conn_type_4_available': Conn Type 4 Available; 'conn_smp_type_0_available': Conn SMP Type 0 Available; 'conn_smp_type_1_available': Conn SMP Type 1 Available; 'conn_smp_type_2_available': Conn SMP Type 2 Available; 'conn_smp_type_3_available': Conn SMP Type 3 Available; 'conn_smp_type_4_available': Conn SMP Type 4 Available; 'sctp-half-open-counter': SCTP Half Open; 'sctp-est-counter': SCTP Established; 'nonssl_bypass': NON SSL Bypass Count; 'ssl_failsafe_total': Total SSL Failsafe Count; 'ssl_forward_proxy_failed_handshake_total': Total SSL Forward Proxy Failed Handshake Count; 'ssl_forward_proxy_failed_tcp_total': Total SSL Forward Proxy Failed TCP Count; 'ssl_forward_proxy_failed_crypto_total': Total SSL Forward Proxy Failed Crypto Count; 'ssl_forward_proxy_failed_cert_verify_total': Total SSL Forward Proxy Failed Certificate Verification Count; 'ssl_forward_proxy_invalid_ocsp_stapling_total': Total SSL Forward Proxy Invalid OCSP Stapling Count; 'ssl_forward_proxy_revoked_ocsp_total': Total SSL Forward Proxy Revoked OCSP Response Count; 'ssl_forward_proxy_failed_cert_signing_total': Total SSL Forward Proxy Failed Certificate Signing Count; 'ssl_forward_proxy_failed_ssl_version_total': Total SSL Forward Proxy Unsupported version Count; 'ssl_forward_proxy_sni_bypass_total': Total SSL Forward Proxy SNI Bypass Count; 'ssl_forward_proxy_client_auth_bypass_total': Total SSL Forward Proxy Client Auth Bypass Count; 'conn_app_smp_alloc_counter': Conn APP SMP Alloc; 'diameter_conn_counter': Diameter Conn Count; 'diameter_conn_freed_counter': Diameter Conn Freed; 'debug_tcp_counter': Hidden TCP sessions for CGNv6 Stateless Technologies; 'debug_udp_counter': Hidden UDP sessions for CGNv6 Stateless Technologies; 'total_fw_conn': Total Firewall Conn; 'total_local_conn': Total Local Conn; 'total_curr_conn': Total Curr Conn; 'client_ssl_fatal_alert': client ssl fatal alert; 'client_ssl_fin_rst': client ssl fin rst; 'fp_session_fin_rst': FP Session FIN/RST; 'server_ssl_fatal_alert': server ssl fatal alert; 'server_ssl_fin_rst': server ssl fin rst; 'client_template_int_err': client template internal error; 'client_template_unknown_err': client template unknown error; 'server_template_int_err': server template int error; 'server_template_unknown_err': server template unknown error; 'total_debug_conn': Total Debug Conn; 'ssl_forward_proxy_failed_aflex_total': Total SSL Forward Proxy AFLEX Count; 'ssl_forward_proxy_cert_subject_bypass_total': Total SSL Forward Proxy Certificate Subject Bypass Count; 'ssl_forward_proxy_cert_issuer_bypass_total': Total SSL Forward Proxy Certificate Issuer Bypass Count; 'ssl_forward_proxy_cert_san_bypass_total': Total SSL Forward Proxy Certificate SAN Bypass Count; 'ssl_forward_proxy_no_sni_bypass_total': Total SSL Forward Proxy No SNI Bypass Count; 'ssl_forward_proxy_no_sni_reset_total': Total SSL Forward Proxy No SNI reset Count; 'ssl_forward_proxy_username_bypass_total': Total SSL Forward Proxy Username Bypass Count; 'ssl_forward_proxy_ad_grpup_bypass_total': Total SSL Forward Proxy AD-Group Bypass Count; 'diameter_concurrent_user_sessions_counter': Diameter Concurrent User-Sessions; 'client_ssl_session_ticket_reuse_total': Total SSL Client Session Ticket Reuse; 'server_ssl_session_ticket_reuse_total': Total SSL Server Session Ticket Reuse; 'total_clientside_early_data_connections': Total clientside early data connections; 'total_serverside_early_data_connections': Total serverside early data connections; 'total_clientside_failed_early_data-connections': Total clientside failed early data connections; 'total_serverside_failed_early_data-connections': Total serverside failed early data connections; 'ssl_forward_proxy_esni_bypass_total': Total SSL Forward Proxy ESNI Bypass Count; 'ssl_forward_proxy_esni_reset_total': Total SSL Forward Proxy ESNI Drop Count; 'total_logging_conn': Total Logging Conn; 'gtp_c_est_counter': GTP-C Established; 'gtp_c_half_open_counter': GTP-C Half Open; 'gtp_u_counter': GTP-U Count; 'gtp_c_echo_counter': GTP-C Echo Count; 'gtp_u_echo_counter': GTP-U Echo Count; 'gtp_curr_free_conn': GTP Current Available Conn; 'gtp_cum_conn_counter': GTP cumulative Conn Count; 'gtp_cum_conn_freed_counter': GTP cumulative Conn Freed;",
										ValidateFunc: validation.StringInSlice([]string{"all", "total_l4_conn", "conn_counter", "conn_freed_counter", "total_l4_packet_count", "total_l7_packet_count", "total_l4_conn_proxy", "total_l7_conn", "total_tcp_conn", "curr_free_conn", "tcp_est_counter", "tcp_half_open_counter", "tcp_half_close_counter", "udp_counter", "ip_counter", "other_counter", "reverse_nat_tcp_counter", "reverse_nat_udp_counter", "tcp_syn_half_open_counter", "conn_smp_alloc_counter", "conn_smp_free_counter", "conn_smp_aged_counter", "ssl_count_curr", "ssl_count_total", "server_ssl_count_curr", "server_ssl_count_total", "client_ssl_reuse_total", "server_ssl_reuse_total", "ssl_failed_total", "ssl_failed_ca_verification", "ssl_server_cert_error", "ssl_client_cert_auth_fail", "total_ip_nat_conn", "total_l2l3_conn", "client_ssl_ctx_malloc_failure", "conn_type_0_available", "conn_type_1_available", "conn_type_2_available", "conn_type_3_available", "conn_type_4_available", "conn_smp_type_0_available", "conn_smp_type_1_available", "conn_smp_type_2_available", "conn_smp_type_3_available", "conn_smp_type_4_available", "sctp-half-open-counter", "sctp-est-counter", "nonssl_bypass", "ssl_failsafe_total", "ssl_forward_proxy_failed_handshake_total", "ssl_forward_proxy_failed_tcp_total", "ssl_forward_proxy_failed_crypto_total", "ssl_forward_proxy_failed_cert_verify_total", "ssl_forward_proxy_invalid_ocsp_stapling_total", "ssl_forward_proxy_revoked_ocsp_total", "ssl_forward_proxy_failed_cert_signing_total", "ssl_forward_proxy_failed_ssl_version_total", "ssl_forward_proxy_sni_bypass_total", "ssl_forward_proxy_client_auth_bypass_total", "conn_app_smp_alloc_counter", "diameter_conn_counter", "diameter_conn_freed_counter", "debug_tcp_counter", "debug_udp_counter", "total_fw_conn", "total_local_conn", "total_curr_conn", "client_ssl_fatal_alert", "client_ssl_fin_rst", "fp_session_fin_rst", "server_ssl_fatal_alert", "server_ssl_fin_rst", "client_template_int_err", "client_template_unknown_err", "server_template_int_err", "server_template_unknown_err", "total_debug_conn", "ssl_forward_proxy_failed_aflex_total", "ssl_forward_proxy_cert_subject_bypass_total", "ssl_forward_proxy_cert_issuer_bypass_total", "ssl_forward_proxy_cert_san_bypass_total", "ssl_forward_proxy_no_sni_bypass_total", "ssl_forward_proxy_no_sni_reset_total", "ssl_forward_proxy_username_bypass_total", "ssl_forward_proxy_ad_grpup_bypass_total", "diameter_concurrent_user_sessions_counter", "client_ssl_session_ticket_reuse_total", "server_ssl_session_ticket_reuse_total", "total_clientside_early_data_connections", "total_serverside_early_data_connections", "total_clientside_failed_early_data-connections", "total_serverside_failed_early_data-connections", "ssl_forward_proxy_esni_bypass_total", "ssl_forward_proxy_esni_reset_total", "total_logging_conn", "gtp_c_est_counter", "gtp_c_half_open_counter", "gtp_u_counter", "gtp_c_echo_counter", "gtp_u_echo_counter", "gtp_curr_free_conn", "gtp_cum_conn_counter", "gtp_cum_conn_freed_counter"}, false),
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
							Type: schema.TypeInt, Optional: true, Default: 4096, Description: "smp session scan limit (number of smp sessions per scan) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 20480),
						},
						"scan_freq": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "smp session scan frequency (scan per second) [shared partition only]",
							ValidateFunc: validation.IntBetween(5, 50),
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
							ValidateFunc: validation.IntBetween(0, 32),
						},
						"rxd_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set new rx-descriptor size",
							ValidateFunc: validation.IntBetween(0, 65536),
						},
						"txd_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set new tx-descriptor size",
							ValidateFunc: validation.IntBetween(0, 65536),
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
							ValidateFunc: validation.IntBetween(0, 32),
						},
						"rxq_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set number of new rx queues",
							ValidateFunc: validation.IntBetween(0, 1024),
						},
						"txq_size": {
							Type: schema.TypeInt, Optional: true, Description: "Set number of new tx queues",
							ValidateFunc: validation.IntBetween(0, 1024),
						},
					},
				},
			},
			"set_tcp_syn_per_sec": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_syn_value": {
							Type: schema.TypeInt, Optional: true, Default: 70, Description: "Configure Tcp SYN's per sec, default 70 [shared partition only]",
							ValidateFunc: validation.IntBetween(25, 6000),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable shared poll mode",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			"shell_privileges": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_shell_privileges": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the shell privileges for a given customer [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			//omit field 'shutdown' which contains uuid only
			"sockstress_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable sockstress protection",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"spe_profile": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "ipv4-ipv6", Description: "'ipv4-only': Enable IPv4 HW forward entries only; 'ipv6-only': Enable IPv6 HW forward entries only; 'ipv4-ipv6': Enable Both IPv4/IPv6 HW forward entries (shared); [shared partition only]",
							ValidateFunc: validation.StringInSlice([]string{"ipv4-only", "ipv6-only", "ipv4-ipv6"}, false),
						},
					},
				},
			},
			//omit field 'spe_status' which contains uuid only
			"src_ip_hash_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable source ip hash",
				ValidateFunc: validation.IntBetween(0, 1),
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'num-ssl-queues': some help string; 'ssl-req-q-depth-tot': some help string; 'ssl-req-q-inuse-tot': some help string; 'ssl-hw-q-depth-tot': some help string; 'ssl-hw-q-inuse-tot': some help string;",
										ValidateFunc: validation.StringInSlice([]string{"all", "num-ssl-queues", "ssl-req-q-depth-tot", "ssl-req-q-inuse-tot", "ssl-hw-q-depth-tot", "ssl-hw-q-inuse-tot"}, false),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable server certificate validation for all SSL connections [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable verify CRL signature during SCV [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable verify host during SCV [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable setting common cipher suite in management plane [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			//omit field 'ssl_status' which contains uuid only
			"syslog_time_msec": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_flag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: " [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			"table_integrity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"table": {
							Type: schema.TypeString, Optional: true, Default: "all", Description: "'all': All tables; [shared partition only]",
							ValidateFunc: validation.StringInSlice([]string{"all"}, false),
						},
						"audit_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable table integrity audit; 'disable': Disable table integrity audit; [shared partition only]",
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
						},
						"auto_sync_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable auto-sync; 'disable': Disable auto-sync; [shared partition only]",
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "arp-tbl-sync-start-ts-m-1st", "nd6-tbl-sync-start-ts-m-1st", "ipv4-fib-tbl-sync-start-ts-m-1st", "ipv6-fib-tbl-sync-start-ts-m-1st", "mac-tbl-sync-start-ts-m-1st", "arp-tbl-sync-start-ts-b-1st", "nd6-tbl-sync-start-ts-b-1st", "ipv4-fib-tbl-sync-start-ts-b-1st", "ipv6-fib-tbl-sync-start-ts-b-1st", "mac-tbl-sync-start-ts-b-1st", "arp-tbl-sync-entries-sent-m-1st", "nd6-tbl-sync-entries-sent-m-1st", "ipv4-fib-tbl-sync-entries-sent-m-1st", "ipv6-fib-tbl-sync-entries-sent-m-1st", "mac-tbl-sync-entries-sent-m-1st", "arp-tbl-sync-entries-rcvd-b-1st", "nd6-tbl-sync-entries-rcvd-b-1st", "ipv4-fib-tbl-sync-entries-rcvd-b-1st", "ipv6-fib-tbl-sync-entries-rcvd-b-1st", "mac-tbl-sync-entries-rcvd-b-1st", "arp-tbl-sync-entries-added-b-1st", "nd6-tbl-sync-entries-added-b-1st", "ipv4-fib-tbl-sync-entries-added-b-1st", "ipv6-fib-tbl-sync-entries-added-b-1st", "mac-tbl-sync-entries-added-b-1st", "arp-tbl-sync-entries-removed-b-1st", "nd6-tbl-sync-entries-removed-b-1st", "ipv4-fib-tbl-sync-entries-removed-b-1st", "ipv6-fib-tbl-sync-entries-removed-b-1st", "mac-tbl-sync-entries-removed-b-1st", "arp-tbl-sync-end-ts-m-1st", "nd6-tbl-sync-end-ts-m-1st", "ipv4-fib-tbl-sync-end-ts-m-1st", "ipv6-fib-tbl-sync-end-ts-m-1st", "mac-tbl-sync-end-ts-m-1st", "arp-tbl-sync-end-ts-b-1st", "nd6-tbl-sync-end-ts-b-1st", "ipv4-fib-tbl-sync-end-ts-b-1st", "ipv6-fib-tbl-sync-end-ts-b-1st", "mac-tbl-sync-end-ts-b-1st", "arp-tbl-sync-start-ts-m-2nd", "nd6-tbl-sync-start-ts-m-2nd", "ipv4-fib-tbl-sync-start-ts-m-2nd", "ipv6-fib-tbl-sync-start-ts-m-2nd", "mac-tbl-sync-start-ts-m-2nd", "arp-tbl-sync-start-ts-b-2nd", "nd6-tbl-sync-start-ts-b-2nd", "ipv4-fib-tbl-sync-start-ts-b-2nd", "ipv6-fib-tbl-sync-start-ts-b-2nd", "mac-tbl-sync-start-ts-b-2nd", "arp-tbl-sync-entries-sent-m-2nd", "nd6-tbl-sync-entries-sent-m-2nd", "ipv4-fib-tbl-sync-entries-sent-m-2nd", "ipv6-fib-tbl-sync-entries-sent-m-2nd", "mac-tbl-sync-entries-sent-m-2nd", "arp-tbl-sync-entries-rcvd-b-2nd", "nd6-tbl-sync-entries-rcvd-b-2nd", "ipv4-fib-tbl-sync-entries-rcvd-b-2nd", "ipv6-fib-tbl-sync-entries-rcvd-b-2nd", "mac-tbl-sync-entries-rcvd-b-2nd", "arp-tbl-sync-entries-added-b-2nd", "nd6-tbl-sync-entries-added-b-2nd", "ipv4-fib-tbl-sync-entries-added-b-2nd", "ipv6-fib-tbl-sync-entries-added-b-2nd", "mac-tbl-sync-entries-added-b-2nd", "arp-tbl-sync-entries-removed-b-2nd", "nd6-tbl-sync-entries-removed-b-2nd", "ipv4-fib-tbl-sync-entries-removed-b-2nd", "ipv6-fib-tbl-sync-entries-removed-b-2nd", "mac-tbl-sync-entries-removed-b-2nd", "arp-tbl-sync-end-ts-m-2nd", "nd6-tbl-sync-end-ts-m-2nd", "ipv4-fib-tbl-sync-end-ts-m-2nd", "ipv6-fib-tbl-sync-end-ts-m-2nd", "mac-tbl-sync-end-ts-m-2nd", "arp-tbl-sync-end-ts-b-2nd", "nd6-tbl-sync-end-ts-b-2nd", "ipv4-fib-tbl-sync-end-ts-b-2nd", "ipv6-fib-tbl-sync-end-ts-b-2nd", "mac-tbl-sync-end-ts-b-2nd", "arp-tbl-sync-start-ts-m-3rd", "nd6-tbl-sync-start-ts-m-3rd"}, false),
									},
									"counters2": {
										Type: schema.TypeString, Optional: true, Description: "'ipv4-fib-tbl-sync-start-ts-m-3rd': ipv4-fib table sync start time stamp master for T-2 synchronization; 'ipv6-fib-tbl-sync-start-ts-m-3rd': ipv6-fib table sync start time stamp master for T-2 synchronization; 'mac-tbl-sync-start-ts-m-3rd': mac table sync start time stamp master for T-2 synchronization; 'arp-tbl-sync-start-ts-b-3rd': arp table sync start time stamp blade for T-2 synchronization; 'nd6-tbl-sync-start-ts-b-3rd': nd6 table sync start time stamp blade for T-2 synchronization; 'ipv4-fib-tbl-sync-start-ts-b-3rd': ipv4-fib table sync start time stamp blade for T-2 synchronization; 'ipv6-fib-tbl-sync-start-ts-b-3rd': ipv6-fib table sync start time stamp blade for T-2 synchronization; 'mac-tbl-sync-start-ts-b-3rd': mac table sync start time stamp blade for T-2 synchronization; 'arp-tbl-sync-entries-sent-m-3rd': arp table entries sent from master for T-2 synchronization; 'nd6-tbl-sync-entries-sent-m-3rd': nd6 table entries sent from master for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-3rd': ipv4-fib table entries sent from master for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-3rd': ipv6-fib table entries sent from master for T-2 synchronization; 'mac-tbl-sync-entries-sent-m-3rd': mac table entries sent from master for T-2 synchronization; 'arp-tbl-sync-entries-rcvd-b-3rd': arp table entries received in blade for T-2 synchronization; 'nd6-tbl-sync-entries-rcvd-b-3rd': nd6 table entries received in blade for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-3rd': ipv4-fib table entries received in blade for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-3rd': ipv6-fib table entries received in blade for T-2 synchronization; 'mac-tbl-sync-entries-rcvd-b-3rd': mac table entries received in blade for T-2 synchronization; 'arp-tbl-sync-entries-added-b-3rd': arp table entries added in blade for T-2 synchronization; 'nd6-tbl-sync-entries-added-b-3rd': nd6 table entries added in blade for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-3rd': ipv4-fib table entries added in blade for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-3rd': ipv6-fib table entries added in blade for T-2 synchronization; 'mac-tbl-sync-entries-added-b-3rd': mac table entries added in blade for T-2 synchronization; 'arp-tbl-sync-entries-removed-b-3rd': arp table entries removed in blade for T-2 synchronization; 'nd6-tbl-sync-entries-removed-b-3rd': nd6 table entries removed in blade for T-2 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-3rd': ipv4-fib table entries removed in blade for T-2 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-3rd': ipv6-fib table entries removed in blade for T-2 synchronization; 'mac-tbl-sync-entries-removed-b-3rd': mac table entries removed in blade for T-2 synchronization; 'arp-tbl-sync-end-ts-m-3rd': arp table sync end time stamp master for T-2 synchronization; 'nd6-tbl-sync-end-ts-m-3rd': nd6 table sync end time stamp master for T-2 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-3rd': ipv4-fib table sync end time stamp master for T-2 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-3rd': ipv6-fib table sync end time stamp master for T-2 synchronization; 'mac-tbl-sync-end-ts-m-3rd': mac table sync end time stamp master for T-2 synchronization; 'arp-tbl-sync-end-ts-b-3rd': arp table sync end time stamp blade for T-2 synchronization; 'nd6-tbl-sync-end-ts-b-3rd': nd6 table sync end time stamp blade for T-2 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-3rd': ipv4-fib table sync end time stamp blade for T-2 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-3rd': ipv6-fib table sync end time stamp blade for T-2 synchronization; 'mac-tbl-sync-end-ts-b-3rd': mac table sync end time stamp blade for T-2 synchronization; 'arp-tbl-sync-start-ts-m-4th': arp table sync start time stamp master for T-3 synchronization; 'nd6-tbl-sync-start-ts-m-4th': nd6 table sync start time stamp master for T-3 synchronization; 'ipv4-fib-tbl-sync-start-ts-m-4th': ipv4-fib table sync start time stamp master for T-3 synchronization; 'ipv6-fib-tbl-sync-start-ts-m-4th': ipv6-fib table sync start time stamp master for T-3 synchronization; 'mac-tbl-sync-start-ts-m-4th': mac table sync start time stamp master for T-3 synchronization; 'arp-tbl-sync-start-ts-b-4th': arp table sync start time stamp blade for T-3 synchronization; 'nd6-tbl-sync-start-ts-b-4th': nd6 table sync start time stamp blade for T-3 synchronization; 'ipv4-fib-tbl-sync-start-ts-b-4th': ipv4-fib table sync start time stamp blade for T-3 synchronization; 'ipv6-fib-tbl-sync-start-ts-b-4th': ipv6-fib table sync start time stamp blade for T-3 synchronization; 'mac-tbl-sync-start-ts-b-4th': mac table sync start time stamp blade for T-3 synchronization; 'arp-tbl-sync-entries-sent-m-4th': arp table entries sent from master for T-3 synchronization; 'nd6-tbl-sync-entries-sent-m-4th': nd6 table entries sent from master for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-4th': ipv4-fib table entries sent from master for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-4th': ipv6-fib table entries sent from master for T-3 synchronization; 'mac-tbl-sync-entries-sent-m-4th': mac table entries sent from master for T-3 synchronization; 'arp-tbl-sync-entries-rcvd-b-4th': arp table entries received in blade for T-3 synchronization; 'nd6-tbl-sync-entries-rcvd-b-4th': nd6 table entries received in blade for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-4th': ipv4-fib table entries received in blade for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-4th': ipv6-fib table entries received in blade for T-3 synchronization; 'mac-tbl-sync-entries-rcvd-b-4th': mac table entries received in blade for T-3 synchronization; 'arp-tbl-sync-entries-added-b-4th': arp table entries added in blade for T-3 synchronization; 'nd6-tbl-sync-entries-added-b-4th': nd6 table entries added in blade for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-4th': ipv4-fib table entries added in blade for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-4th': ipv6-fib table entries added in blade for T-3 synchronization; 'mac-tbl-sync-entries-added-b-4th': mac table entries added in blade for T-3 synchronization; 'arp-tbl-sync-entries-removed-b-4th': arp table entries removed in blade for T-3 synchronization; 'nd6-tbl-sync-entries-removed-b-4th': nd6 table entries removed in blade for T-3 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-4th': ipv4-fib table entries removed in blade for T-3 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-4th': ipv6-fib table entries removed in blade for T-3 synchronization; 'mac-tbl-sync-entries-removed-b-4th': mac table entries removed in blade for T-3 synchronization; 'arp-tbl-sync-end-ts-m-4th': arp table sync end time stamp master for T-3 synchronization; 'nd6-tbl-sync-end-ts-m-4th': nd6 table sync end time stamp master for T-3 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-4th': ipv4-fib table sync end time stamp master for T-3 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-4th': ipv6-fib table sync end time stamp master for T-3 synchronization; 'mac-tbl-sync-end-ts-m-4th': mac table sync end time stamp master for T-3 synchronization; 'arp-tbl-sync-end-ts-b-4th': arp table sync end time stamp blade for T-3 synchronization; 'nd6-tbl-sync-end-ts-b-4th': nd6 table sync end time stamp blade for T-3 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-4th': ipv4-fib table sync end time stamp blade for T-3 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-4th': ipv6-fib table sync end time stamp blade for T-3 synchronization; 'mac-tbl-sync-end-ts-b-4th': mac table sync end time stamp blade for T-3 synchronization; 'arp-tbl-sync-start-ts-m-5th': arp table sync start time stamp master for T-4 synchronization;",
										ValidateFunc: validation.StringInSlice([]string{"ipv4-fib-tbl-sync-start-ts-m-3rd", "ipv6-fib-tbl-sync-start-ts-m-3rd", "mac-tbl-sync-start-ts-m-3rd", "arp-tbl-sync-start-ts-b-3rd", "nd6-tbl-sync-start-ts-b-3rd", "ipv4-fib-tbl-sync-start-ts-b-3rd", "ipv6-fib-tbl-sync-start-ts-b-3rd", "mac-tbl-sync-start-ts-b-3rd", "arp-tbl-sync-entries-sent-m-3rd", "nd6-tbl-sync-entries-sent-m-3rd", "ipv4-fib-tbl-sync-entries-sent-m-3rd", "ipv6-fib-tbl-sync-entries-sent-m-3rd", "mac-tbl-sync-entries-sent-m-3rd", "arp-tbl-sync-entries-rcvd-b-3rd", "nd6-tbl-sync-entries-rcvd-b-3rd", "ipv4-fib-tbl-sync-entries-rcvd-b-3rd", "ipv6-fib-tbl-sync-entries-rcvd-b-3rd", "mac-tbl-sync-entries-rcvd-b-3rd", "arp-tbl-sync-entries-added-b-3rd", "nd6-tbl-sync-entries-added-b-3rd", "ipv4-fib-tbl-sync-entries-added-b-3rd", "ipv6-fib-tbl-sync-entries-added-b-3rd", "mac-tbl-sync-entries-added-b-3rd", "arp-tbl-sync-entries-removed-b-3rd", "nd6-tbl-sync-entries-removed-b-3rd", "ipv4-fib-tbl-sync-entries-removed-b-3rd", "ipv6-fib-tbl-sync-entries-removed-b-3rd", "mac-tbl-sync-entries-removed-b-3rd", "arp-tbl-sync-end-ts-m-3rd", "nd6-tbl-sync-end-ts-m-3rd", "ipv4-fib-tbl-sync-end-ts-m-3rd", "ipv6-fib-tbl-sync-end-ts-m-3rd", "mac-tbl-sync-end-ts-m-3rd", "arp-tbl-sync-end-ts-b-3rd", "nd6-tbl-sync-end-ts-b-3rd", "ipv4-fib-tbl-sync-end-ts-b-3rd", "ipv6-fib-tbl-sync-end-ts-b-3rd", "mac-tbl-sync-end-ts-b-3rd", "arp-tbl-sync-start-ts-m-4th", "nd6-tbl-sync-start-ts-m-4th", "ipv4-fib-tbl-sync-start-ts-m-4th", "ipv6-fib-tbl-sync-start-ts-m-4th", "mac-tbl-sync-start-ts-m-4th", "arp-tbl-sync-start-ts-b-4th", "nd6-tbl-sync-start-ts-b-4th", "ipv4-fib-tbl-sync-start-ts-b-4th", "ipv6-fib-tbl-sync-start-ts-b-4th", "mac-tbl-sync-start-ts-b-4th", "arp-tbl-sync-entries-sent-m-4th", "nd6-tbl-sync-entries-sent-m-4th", "ipv4-fib-tbl-sync-entries-sent-m-4th", "ipv6-fib-tbl-sync-entries-sent-m-4th", "mac-tbl-sync-entries-sent-m-4th", "arp-tbl-sync-entries-rcvd-b-4th", "nd6-tbl-sync-entries-rcvd-b-4th", "ipv4-fib-tbl-sync-entries-rcvd-b-4th", "ipv6-fib-tbl-sync-entries-rcvd-b-4th", "mac-tbl-sync-entries-rcvd-b-4th", "arp-tbl-sync-entries-added-b-4th", "nd6-tbl-sync-entries-added-b-4th", "ipv4-fib-tbl-sync-entries-added-b-4th", "ipv6-fib-tbl-sync-entries-added-b-4th", "mac-tbl-sync-entries-added-b-4th", "arp-tbl-sync-entries-removed-b-4th", "nd6-tbl-sync-entries-removed-b-4th", "ipv4-fib-tbl-sync-entries-removed-b-4th", "ipv6-fib-tbl-sync-entries-removed-b-4th", "mac-tbl-sync-entries-removed-b-4th", "arp-tbl-sync-end-ts-m-4th", "nd6-tbl-sync-end-ts-m-4th", "ipv4-fib-tbl-sync-end-ts-m-4th", "ipv6-fib-tbl-sync-end-ts-m-4th", "mac-tbl-sync-end-ts-m-4th", "arp-tbl-sync-end-ts-b-4th", "nd6-tbl-sync-end-ts-b-4th", "ipv4-fib-tbl-sync-end-ts-b-4th", "ipv6-fib-tbl-sync-end-ts-b-4th", "mac-tbl-sync-end-ts-b-4th", "arp-tbl-sync-start-ts-m-5th"}, false),
									},
									"counters3": {
										Type: schema.TypeString, Optional: true, Description: "'nd6-tbl-sync-start-ts-m-5th': nd6 table sync start time stamp master for T-4 synchronization; 'ipv4-fib-tbl-sync-start-ts-m-5th': ipv4-fib table sync start time stamp master for T-4 synchronization; 'ipv6-fib-tbl-sync-start-ts-m-5th': ipv6-fib table sync start time stamp master for T-4 synchronization; 'mac-tbl-sync-start-ts-m-5th': mac table sync start time stamp master for T-4 synchronization; 'arp-tbl-sync-start-ts-b-5th': arp table sync start time stamp blade for T-4 synchronization; 'nd6-tbl-sync-start-ts-b-5th': nd6 table sync start time stamp blade for T-4 synchronization; 'ipv4-fib-tbl-sync-start-ts-b-5th': ipv4-fib table sync start time stamp blade for T-4 synchronization; 'ipv6-fib-tbl-sync-start-ts-b-5th': ipv6-fib table sync start time stamp blade for T-4 synchronization; 'mac-tbl-sync-start-ts-b-5th': mac table sync start time stamp blade for T-4 synchronization; 'arp-tbl-sync-entries-sent-m-5th': arp table sync start time stamp blade for T-4 synchronization; 'nd6-tbl-sync-entries-sent-m-5th': nd6 table sync start time stamp blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-sent-m-5th': ipv4-fib table sync start time stamp blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-sent-m-5th': ipv6-fib table sync start time stamp blade for T-4 synchronization; 'mac-tbl-sync-entries-sent-m-5th': mac table sync start time stamp blade for T-4 synchronization; 'arp-tbl-sync-entries-rcvd-b-5th': arp table entries received in blade for T-4 synchronization; 'nd6-tbl-sync-entries-rcvd-b-5th': nd6 table entries received in blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-rcvd-b-5th': ipv4-fib table entries received in blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-rcvd-b-5th': ipv6-fib table entries received in blade for T-4 synchronization; 'mac-tbl-sync-entries-rcvd-b-5th': mac table entries received in blade for T-4 synchronization; 'arp-tbl-sync-entries-added-b-5th': arp table entries added in blade for T-4 synchronization; 'nd6-tbl-sync-entries-added-b-5th': nd6 table entries added in blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-added-b-5th': ipv4-fib table entries added in blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-added-b-5th': ipv6-fib table entries added in blade for T-4 synchronization; 'mac-tbl-sync-entries-added-b-5th': mac table entries added in blade for T-4 synchronization; 'arp-tbl-sync-entries-removed-b-5th': arp table entries removed in blade for T-4 synchronization; 'nd6-tbl-sync-entries-removed-b-5th': nd6 table entries removed in blade for T-4 synchronization; 'ipv4-fib-tbl-sync-entries-removed-b-5th': ipv4-fib table entries removed in blade for T-4 synchronization; 'ipv6-fib-tbl-sync-entries-removed-b-5th': ipv6-fib table entries removed in blade for T-4 synchronization; 'mac-tbl-sync-entries-removed-b-5th': mac table entries removed in blade for T-4 synchronization; 'arp-tbl-sync-end-ts-m-5th': arp table sync end time stamp master for T-4 synchronization; 'nd6-tbl-sync-end-ts-m-5th': nd6 table sync end time stamp master for T-4 synchronization; 'ipv4-fib-tbl-sync-end-ts-m-5th': ipv4-fib table sync end time stamp master for T-4 synchronization; 'ipv6-fib-tbl-sync-end-ts-m-5th': ipv6-fib table sync end time stamp master for T-4 synchronization; 'mac-tbl-sync-end-ts-m-5th': mac table sync end time stamp master for T-4 synchronization; 'arp-tbl-sync-end-ts-b-5th': arp table sync end time stamp blade for T-4 synchronization; 'nd6-tbl-sync-end-ts-b-5th': nd6 table sync end time stamp blade for T-4 synchronization; 'ipv4-fib-tbl-sync-end-ts-b-5th': ipv4-fib table sync end time stamp blade for T-4 synchronization; 'ipv6-fib-tbl-sync-end-ts-b-5th': ipv6-fib table sync end time stamp blade for T-4 synchronization; 'mac-tbl-sync-end-ts-b-5th': mac table sync end time stamp blade for T-4 synchronization; 'arp-tbl-sync-m': arp table sync count in master; 'nd6-tbl-sync-m': nd6 table sync count in master; 'ipv4-fib-tbl-sync-m': ipv4-fib table sync count in master; 'ipv6-fib-tbl-sync-m': ipv6-fib table sync count in master; 'mac-tbl-sync-m': mac table sync count in master; 'arp-tbl-sync-b': arp table sync count in blade; 'nd6-tbl-sync-b': nd6 table sync count in blade; 'ipv4-fib-tbl-sync-b': ipv4-fib table sync count in blade; 'ipv6-fib-tbl-sync-b': ipv6-fib table sync count in blade; 'mac-tbl-sync-b': mac table sync count in blade; 'arp-tbl-cksum-m': arp table checksum count in master; 'nd6-tbl-cksum-m': nd6 table checksum count in master; 'ipv4-fib-tbl-cksum-m': ipv4-fib table checksum count in master; 'ipv6-fib-tbl-cksum-m': ipv6-fib table checksum count in master; 'mac-tbl-cksum-m': mac table checksum count in master; 'arp-tbl-cksum-b': arp table checksum count in blade; 'nd6-tbl-cksum-b': nd6 table checksum count in blade; 'ipv4-fib-tbl-cksum-b': ipv4-fib table checksum count in blade; 'ipv6-fib-tbl-cksum-b': ipv6-fib table checksum count in blade; 'mac-tbl-cksum-b': mac table checksum count in blade; 'arp-tbl-cksum-mismatch-b': arp table checksum mismatch count in blade; 'nd6-tbl-cksum-mismatch-b': nd6 table checksum mismatch count in blade; 'ipv4-fib-tbl-cksum-mismatch-b': ipv4-fib table checksum mismatch count in blade; 'ipv6-fib-tbl-cksum-mismatch-b': ipv6-fib table checksum mismatch count in blade; 'mac-tbl-cksum-mismatch-b': mac table checksum mismatch count in blade; 'arp-tbl-cksum-cancel-m': arp table checksum cancelled count in master; 'nd6-tbl-cksum-cancel-m': nd6 table checksum cancelled count in master; 'ipv4-fib-tbl-cksum-cancel-m': ipv4-fib table checksum cancelled count in master; 'ipv6-fib-tbl-cksum-cancel-m': ipv6-fib table checksum cancelled count in master; 'mac-tbl-cksum-cancel-m': mac table checksum cancelled count in master;",
										ValidateFunc: validation.StringInSlice([]string{"nd6-tbl-sync-start-ts-m-5th", "ipv4-fib-tbl-sync-start-ts-m-5th", "ipv6-fib-tbl-sync-start-ts-m-5th", "mac-tbl-sync-start-ts-m-5th", "arp-tbl-sync-start-ts-b-5th", "nd6-tbl-sync-start-ts-b-5th", "ipv4-fib-tbl-sync-start-ts-b-5th", "ipv6-fib-tbl-sync-start-ts-b-5th", "mac-tbl-sync-start-ts-b-5th", "arp-tbl-sync-entries-sent-m-5th", "nd6-tbl-sync-entries-sent-m-5th", "ipv4-fib-tbl-sync-entries-sent-m-5th", "ipv6-fib-tbl-sync-entries-sent-m-5th", "mac-tbl-sync-entries-sent-m-5th", "arp-tbl-sync-entries-rcvd-b-5th", "nd6-tbl-sync-entries-rcvd-b-5th", "ipv4-fib-tbl-sync-entries-rcvd-b-5th", "ipv6-fib-tbl-sync-entries-rcvd-b-5th", "mac-tbl-sync-entries-rcvd-b-5th", "arp-tbl-sync-entries-added-b-5th", "nd6-tbl-sync-entries-added-b-5th", "ipv4-fib-tbl-sync-entries-added-b-5th", "ipv6-fib-tbl-sync-entries-added-b-5th", "mac-tbl-sync-entries-added-b-5th", "arp-tbl-sync-entries-removed-b-5th", "nd6-tbl-sync-entries-removed-b-5th", "ipv4-fib-tbl-sync-entries-removed-b-5th", "ipv6-fib-tbl-sync-entries-removed-b-5th", "mac-tbl-sync-entries-removed-b-5th", "arp-tbl-sync-end-ts-m-5th", "nd6-tbl-sync-end-ts-m-5th", "ipv4-fib-tbl-sync-end-ts-m-5th", "ipv6-fib-tbl-sync-end-ts-m-5th", "mac-tbl-sync-end-ts-m-5th", "arp-tbl-sync-end-ts-b-5th", "nd6-tbl-sync-end-ts-b-5th", "ipv4-fib-tbl-sync-end-ts-b-5th", "ipv6-fib-tbl-sync-end-ts-b-5th", "mac-tbl-sync-end-ts-b-5th", "arp-tbl-sync-m", "nd6-tbl-sync-m", "ipv4-fib-tbl-sync-m", "ipv6-fib-tbl-sync-m", "mac-tbl-sync-m", "arp-tbl-sync-b", "nd6-tbl-sync-b", "ipv4-fib-tbl-sync-b", "ipv6-fib-tbl-sync-b", "mac-tbl-sync-b", "arp-tbl-cksum-m", "nd6-tbl-cksum-m", "ipv4-fib-tbl-cksum-m", "ipv6-fib-tbl-cksum-m", "mac-tbl-cksum-m", "arp-tbl-cksum-b", "nd6-tbl-cksum-b", "ipv4-fib-tbl-cksum-b", "ipv6-fib-tbl-cksum-b", "mac-tbl-cksum-b", "arp-tbl-cksum-mismatch-b", "nd6-tbl-cksum-mismatch-b", "ipv4-fib-tbl-cksum-mismatch-b", "ipv6-fib-tbl-cksum-mismatch-b", "mac-tbl-cksum-mismatch-b", "arp-tbl-cksum-cancel-m", "nd6-tbl-cksum-cancel-m", "ipv4-fib-tbl-cksum-cancel-m", "ipv6-fib-tbl-cksum-cancel-m", "mac-tbl-cksum-cancel-m"}, false),
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'activeopens': Active open conns; 'passiveopens': Passive open conns; 'attemptfails': Connect attemp failures; 'estabresets': Resets rcvd on EST conn; 'insegs': Total in TCP packets; 'outsegs': Total out TCP packets; 'retranssegs': Retransmited packets; 'inerrs': Input errors; 'outrsts': Reset Sent; 'sock_alloc': Sockets allocated; 'orphan_count': Orphan sockets; 'mem_alloc': Memory alloc; 'recv_mem': Total rx buffer; 'send_mem': Total tx buffer; 'currestab': Currently EST conns; 'currsyssnt': TCP in SYN-SNT state; 'currsynrcv': TCP in SYN-RCV state; 'currfinw1': TCP in FIN-W1 state; 'currfinw2': TCP FIN-W2 state; 'currtimew': TCP TimeW state; 'currclose': TCP in Close state; 'currclsw': TCP in CloseW state; 'currlack': TCP in LastACK state; 'currlstn': TCP in Listen state; 'currclsg': TCP in Closing state; 'pawsactiverejected': TCP paw active rej; 'syn_rcv_rstack': Rcv RST|ACK on SYN; 'syn_rcv_rst': Rcv RST on SYN; 'syn_rcv_ack': Rcv ACK on SYN; 'ax_rexmit_syn': TCP rexmit SYN; 'tcpabortontimeout': TCP abort on timeout; 'noroute': TCPIP out noroute; 'exceedmss': MSS exceeded pkt dropped; 'tfo_conns': TFO Total Connections; 'tfo_actives': TFO Current Actives; 'tfo_denied': TFO Denied;",
										ValidateFunc: validation.StringInSlice([]string{"all", "activeopens", "passiveopens", "attemptfails", "estabresets", "insegs", "outsegs", "retranssegs", "inerrs", "outrsts", "sock_alloc", "orphan_count", "mem_alloc", "recv_mem", "send_mem", "currestab", "currsyssnt", "currsynrcv", "currfinw1", "currfinw2", "currtimew", "currclose", "currclsw", "currlack", "currlstn", "currclsg", "pawsactiverejected", "syn_rcv_rstack", "syn_rcv_rst", "syn_rcv_ack", "ax_rexmit_syn", "tcpabortontimeout", "noroute", "exceedmss", "tfo_conns", "tfo_actives", "tfo_denied"}, false),
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
										ValidateFunc: validation.IntBetween(1, 1048575),
									},
									"log_for_reset_unknown_conn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when rate exceed",
										ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "connattempt", "connects", "drops", "conndrops", "closed", "segstimed", "rttupdated", "delack", "timeoutdrop", "rexmttimeo", "persisttimeo", "keeptimeo", "keepprobe", "keepdrops", "sndtotal", "sndpack", "sndbyte", "sndrexmitpack", "sndrexmitbyte", "sndrexmitbad", "sndacks", "sndprobe", "sndurg", "sndwinup", "sndctrl", "sndrst", "sndfin", "sndsyn", "rcvtotal", "rcvpack", "rcvbyte", "rcvbadoff", "rcvmemdrop", "rcvduppack", "rcvdupbyte", "rcvpartduppack", "rcvpartdupbyte", "rcvoopack", "rcvoobyte", "rcvpackafterwin", "rcvbyteafterwin", "rcvwinprobe", "rcvdupack", "rcvacktoomuch", "rcvackpack", "rcvackbyte", "rcvwinupd", "pawsdrop", "predack", "preddat", "persistdrop", "badrst", "finwait2_drops", "sack_recovery_episode", "sack_rexmits", "sack_rexmit_bytes", "sack_rcv_blocks", "sack_send_blocks", "sndcack", "cacklim", "reassmemdrop", "reasstimeout", "cc_idle", "cc_reduce", "rcvdsack", "a2brcvwnd", "a2bsackpresent", "a2bdupack", "a2brxdata", "a2btcpoptions", "a2boodata", "a2bpartialack", "a2bfsmtransition", "a2btransitionnum", "b2atransitionnum", "bad_iochan", "atcpforward", "atcpsent", "atcprexmitsadrop", "atcpsendbackack", "atcprexmit", "atcpbuffallocfail", "a2bappbuffering", "atcpsendfail", "earlyrexmit", "mburstlim", "a2bsndwnd", "proxyheaderv1", "proxyheaderv2"}, false),
									},
								},
							},
						},
					},
				},
			},
			//omit field 'tcp_syn_per_sec' which contains uuid only
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_policy": {
							Type: schema.TypeString, Optional: true, Description: "Apply policy template to the whole system (Policy template name) [shared partition only]",
							ValidateFunc: validation.StringLenBetween(1, 127),
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
										Type: schema.TypeInt, Required: true, Description: "Monitor template ID Number [shared partition only]",
										ValidateFunc: validation.IntBetween(1, 16),
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
										ValidateFunc: validation.StringInSlice([]string{"all", "global-system-throughput-bits-per-sec", "per-part-throughput-bits-per-sec"}, false),
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
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "set timeout to stop ftp transfer in seconds, 0 is no limit [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 9999),
						},
						"scp": {
							Type: schema.TypeInt, Optional: true, Default: 300, Description: "set timeout to stop scp transfer in seconds, 0 is no limit [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 9999),
						},
						"sftp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop sftp transfer in seconds, 0 is no limit [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 9999),
						},
						"tftp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop tftp transfer in seconds, 0 is no limit [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 9999),
						},
						"http": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop http transfer in seconds, 0 is no limit [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 9999),
						},
						"https": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop https transfer in seconds, 0 is no limit [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 9999),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tls_1_3_mgmt": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable TLS 1.3 support on ACOS management plane [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer-3 Header based load balancing [shared partition only]",
										ValidateFunc:  validation.IntBetween(0, 1),
										ConflictsWith: []string{"trunk.0.load_balance.0.use_l4"},
									},
									"use_l4": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer-3/4 Header based load balancing [shared partition only]",
										ValidateFunc:  validation.IntBetween(0, 1),
										ConflictsWith: []string{"trunk.0.load_balance.0.use_l3"},
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
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Set HW hash mode, default is 6 (1:dst-mac 2:src-mac 3:src-dst-mac 4:src-ip 5:dst-ip 6:rtag6 7:rtag7) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 7),
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
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Set HW hash mode, default is 6 (1:dst-mac 2:src-mac 3:src-dst-mac 4:src-ip 5:dst-ip 6:rtag6 7:rtag7) [shared partition only]",
							ValidateFunc: validation.IntBetween(1, 7),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable TCP Segmentation Offload",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			//omit field 'upgrade_status' which contains uuid only
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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

func getObjectSystemAddCpuCore(d []interface{}) edpt.SystemAddCpuCore {
	var ret edpt.SystemAddCpuCore
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.CoreIndex = in["core_index"].(int)
	}
	return ret
}

func getObjectSystemAddPort(d []interface{}) edpt.SystemAddPort {
	var ret edpt.SystemAddPort
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
	}
	return ret
}

func getObjectSystemAllVlanLimit(d []interface{}) edpt.SystemAllVlanLimit {
	var ret edpt.SystemAllVlanLimit
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bcast = in["bcast"].(int)
		ret.Ipmcast = in["ipmcast"].(int)
		ret.Mcast = in["mcast"].(int)
		ret.UnknownUcast = in["unknown_ucast"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemAppPerformance(d []interface{}) edpt.SystemAppPerformance {
	var ret edpt.SystemAppPerformance
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemAppPerformanceSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemAppPerformanceSamplingEnable(d []interface{}) []edpt.SystemAppPerformanceSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemAppPerformanceSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemAppPerformanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemAppsGlobal(d []interface{}) edpt.SystemAppsGlobal {
	var ret edpt.SystemAppsGlobal
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.LogSessionOnEstablished = in["log_session_on_established"].(int)
		ret.MslTime = in["msl_time"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemAsicDebugDump(d []interface{}) edpt.SystemAsicDebugDump {
	var ret edpt.SystemAsicDebugDump
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemBandwidth(d []interface{}) edpt.SystemBandwidth {
	var ret edpt.SystemBandwidth
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemBandwidthSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemBandwidthSamplingEnable(d []interface{}) []edpt.SystemBandwidthSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemBandwidthSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemBandwidthSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemBfd(d []interface{}) edpt.SystemBfd {
	var ret edpt.SystemBfd
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemBfdSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemBfdSamplingEnable(d []interface{}) []edpt.SystemBfdSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemBfdSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemBfdSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemCmUpdateFileNameRef(d []interface{}) edpt.SystemCmUpdateFileNameRef {
	var ret edpt.SystemCmUpdateFileNameRef
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Source_name = in["source_name"].(string)
		ret.Dest_name = in["dest_name"].(string)
		ret.Id = in["id"].(int)
	}
	return ret
}

//omit object-type field 'control_cpu' which contains uuid only

//omit object-type field 'core' which contains uuid only

//omit object-type field 'cosq_show' which contains uuid only

//omit object-type field 'cosq_stats' which contains uuid only

//omit object-type field 'counter_lib_accounting' which contains uuid only

func getObjectSystemCpuHyperThread(d []interface{}) edpt.SystemCpuHyperThread {
	var ret edpt.SystemCpuHyperThread
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

//omit object-type field 'cpu_list' which contains uuid only

func getObjectSystemCpuLoadSharing(d []interface{}) edpt.SystemCpuLoadSharing {
	var ret edpt.SystemCpuLoadSharing
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.PacketsPerSecond = getObjectSystemCpuLoadSharingPacketsPerSecond(in["packets_per_second"].([]interface{}))
		ret.CpuUsage = getObjectSystemCpuLoadSharingCpuUsage(in["cpu_usage"].([]interface{}))
		ret.Tcp = in["tcp"].(int)
		ret.Udp = in["udp"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemCpuLoadSharingPacketsPerSecond(d []interface{}) edpt.SystemCpuLoadSharingPacketsPerSecond {
	var ret edpt.SystemCpuLoadSharingPacketsPerSecond
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Min = in["min"].(int)
	}
	return ret
}

func getObjectSystemCpuLoadSharingCpuUsage(d []interface{}) edpt.SystemCpuLoadSharingCpuUsage {
	var ret edpt.SystemCpuLoadSharingCpuUsage
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Low = in["low"].(int)
		ret.High = in["high"].(int)
	}
	return ret
}

//omit object-type field 'cpu_map' which contains uuid only

//omit object-type field 'data_cpu' which contains uuid only

func getObjectSystemDeepHrxq(d []interface{}) edpt.SystemDeepHrxq {
	var ret edpt.SystemDeepHrxq
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
	}
	return ret
}

func getObjectSystemDelPort(d []interface{}) edpt.SystemDelPort {
	var ret edpt.SystemDelPort
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
	}
	return ret
}

func getObjectSystemDeleteCpuCore(d []interface{}) edpt.SystemDeleteCpuCore {
	var ret edpt.SystemDeleteCpuCore
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.CoreIndex = in["core_index"].(int)
	}
	return ret
}

func getObjectSystemDns(d []interface{}) edpt.SystemDns {
	var ret edpt.SystemDns
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemDnsSamplingEnable(in["sampling_enable"].([]interface{}))
		ret.RecursiveNameserver = getObjectSystemDnsRecursiveNameserver(in["recursive_nameserver"].([]interface{}))
	}
	return ret
}

func getSliceSystemDnsSamplingEnable(d []interface{}) []edpt.SystemDnsSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemDnsSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemDnsRecursiveNameserver(d []interface{}) edpt.SystemDnsRecursiveNameserver {
	var ret edpt.SystemDnsRecursiveNameserver
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		ret.ServerList = getSliceSystemDnsRecursiveNameserverServerList(in["server_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemDnsRecursiveNameserverServerList(d []interface{}) []edpt.SystemDnsRecursiveNameserverServerList {
	count := len(d)
	ret := make([]edpt.SystemDnsRecursiveNameserverServerList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsRecursiveNameserverServerList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.V4Desc = in["v4_desc"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.V6Desc = in["v6_desc"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemDnsCache(d []interface{}) edpt.SystemDnsCache {
	var ret edpt.SystemDnsCache
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemDnsCacheSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemDnsCacheSamplingEnable(d []interface{}) []edpt.SystemDnsCacheSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemDnsCacheSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsCacheSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

//omit object-type field 'domain_list_info' which contains uuid only

func getObjectSystemDpdkStats(d []interface{}) edpt.SystemDpdkStats {
	var ret edpt.SystemDpdkStats
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemDpdkStatsSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemDpdkStatsSamplingEnable(d []interface{}) []edpt.SystemDpdkStatsSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemDpdkStatsSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemDpdkStatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

//omit object-type field 'environment' which contains uuid only

func getObjectSystemFpgaCoreCrc(d []interface{}) edpt.SystemFpgaCoreCrc {
	var ret edpt.SystemFpgaCoreCrc
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.MonitorDisable = in["monitor_disable"].(int)
		ret.RebootEnable = in["reboot_enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemFpgaDrop(d []interface{}) edpt.SystemFpgaDrop {
	var ret edpt.SystemFpgaDrop
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemFpgaDropSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemFpgaDropSamplingEnable(d []interface{}) []edpt.SystemFpgaDropSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemFpgaDropSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemFpgaDropSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemFw(d []interface{}) edpt.SystemFw {
	var ret edpt.SystemFw
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ApplicationMempool = in["application_mempool"].(int)
		ret.ApplicationFlow = in["application_flow"].(int)
		ret.BasicDpiEnable = in["basic_dpi_enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemGeoLocation(d []interface{}) edpt.SystemGeoLocation {
	var ret edpt.SystemGeoLocation
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GeoLocationIana = in["geo_location_iana"].(int)
		ret.GeoLocationGeolite2City = in["geo_location_geolite2_city"].(int)
		ret.Geolite2CityIncludeIpv6 = in["geolite2_city_include_ipv6"].(int)
		ret.GeoLocationGeolite2Country = in["geo_location_geolite2_country"].(int)
		ret.Geolite2CountryIncludeIpv6 = in["geolite2_country_include_ipv6"].(int)
		ret.GeolocLoadFileList = getSliceSystemGeoLocationGeolocLoadFileList(in["geoloc_load_file_list"].([]interface{}))
		//omit uuid
		ret.EntryList = getSliceSystemGeoLocationEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeoLocationGeolocLoadFileList(d []interface{}) []edpt.SystemGeoLocationGeolocLoadFileList {
	count := len(d)
	ret := make([]edpt.SystemGeoLocationGeolocLoadFileList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationGeolocLoadFileList
		oi.GeoLocationLoadFilename = in["geo_location_load_filename"].(string)
		oi.TemplateName = in["template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeoLocationEntryList(d []interface{}) []edpt.SystemGeoLocationEntryList {
	count := len(d)
	ret := make([]edpt.SystemGeoLocationEntryList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationEntryList
		oi.GeoLocnObjName = in["geo_locn_obj_name"].(string)
		oi.GeoLocnMultipleAddresses = getSliceSystemGeoLocationEntryListGeoLocnMultipleAddresses(in["geo_locn_multiple_addresses"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeoLocationEntryListGeoLocnMultipleAddresses(d []interface{}) []edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses {
	count := len(d)
	ret := make([]edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses
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

func getObjectSystemGeoloc(d []interface{}) edpt.SystemGeoloc {
	var ret edpt.SystemGeoloc
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemGeolocSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeolocSamplingEnable(d []interface{}) []edpt.SystemGeolocSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemGeolocSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListList(d []interface{}) []edpt.SystemGeolocListList {
	count := len(d)
	ret := make([]edpt.SystemGeolocListList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.SystemGeolocListListIncludeGeolocNameList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListListIncludeGeolocNameList
		oi.IncludeGeolocNameVal = in["include_geoloc_name_val"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListListExcludeGeolocNameList(d []interface{}) []edpt.SystemGeolocListListExcludeGeolocNameList {
	count := len(d)
	ret := make([]edpt.SystemGeolocListListExcludeGeolocNameList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListListExcludeGeolocNameList
		oi.ExcludeGeolocNameVal = in["exclude_geoloc_name_val"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeolocListListSamplingEnable(d []interface{}) []edpt.SystemGeolocListListSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemGeolocListListSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemGeolocNameHelper(d []interface{}) edpt.SystemGeolocNameHelper {
	var ret edpt.SystemGeolocNameHelper
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemGeolocNameHelperSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeolocNameHelperSamplingEnable(d []interface{}) []edpt.SystemGeolocNameHelperSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemGeolocNameHelperSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocNameHelperSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemHardwareForward(d []interface{}) edpt.SystemHardwareForward {
	var ret edpt.SystemHardwareForward
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemHardwareForwardSamplingEnable(in["sampling_enable"].([]interface{}))
		ret.Slb = getObjectSystemHardwareForwardSlb(in["slb"].([]interface{}))
	}
	return ret
}

func getSliceSystemHardwareForwardSamplingEnable(d []interface{}) []edpt.SystemHardwareForwardSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemHardwareForwardSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareForwardSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemHardwareForwardSlb(d []interface{}) edpt.SystemHardwareForwardSlb {
	var ret edpt.SystemHardwareForwardSlb
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemHardwareForwardSlbSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemHardwareForwardSlbSamplingEnable(d []interface{}) []edpt.SystemHardwareForwardSlbSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemHardwareForwardSlbSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareForwardSlbSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemHighMemoryL4Session(d []interface{}) edpt.SystemHighMemoryL4Session {
	var ret edpt.SystemHighMemoryL4Session
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

//omit object-type field 'hrxq_status' which contains uuid only

func getObjectSystemIcmp(d []interface{}) edpt.SystemIcmp {
	var ret edpt.SystemIcmp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIcmpSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIcmpSamplingEnable(d []interface{}) []edpt.SystemIcmpSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemIcmpSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIcmpRate(d []interface{}) edpt.SystemIcmpRate {
	var ret edpt.SystemIcmpRate
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIcmpRateSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIcmpRateSamplingEnable(d []interface{}) []edpt.SystemIcmpRateSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemIcmpRateSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmpRateSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIcmp6(d []interface{}) edpt.SystemIcmp6 {
	var ret edpt.SystemIcmp6
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIcmp6SamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIcmp6SamplingEnable(d []interface{}) []edpt.SystemIcmp6SamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemIcmp6SamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmp6SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

//omit object-type field 'inuse_cpu_list' which contains uuid only

//omit object-type field 'inuse_port_list' which contains uuid only

func getObjectSystemIoCpu(d []interface{}) edpt.SystemIoCpu {
	var ret edpt.SystemIoCpu
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.MaxCores = in["max_cores"].(int)
	}
	return ret
}

//omit object-type field 'ip_dns_cache' which contains uuid only

func getObjectSystemIpStats(d []interface{}) edpt.SystemIpStats {
	var ret edpt.SystemIpStats
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIpStatsSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIpStatsSamplingEnable(d []interface{}) []edpt.SystemIpStatsSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemIpStatsSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpStatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatList(d []interface{}) edpt.SystemIpThreatList {
	var ret edpt.SystemIpThreatList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIpThreatListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret.Ipv4SourceList = getObjectSystemIpThreatListIpv4SourceList(in["ipv4_source_list"].([]interface{}))
		ret.Ipv4DestList = getObjectSystemIpThreatListIpv4DestList(in["ipv4_dest_list"].([]interface{}))
		ret.Ipv6SourceList = getObjectSystemIpThreatListIpv6SourceList(in["ipv6_source_list"].([]interface{}))
		ret.Ipv6DestList = getObjectSystemIpThreatListIpv6DestList(in["ipv6_dest_list"].([]interface{}))
		ret.Ipv4InternetHostList = getObjectSystemIpThreatListIpv4InternetHostList(in["ipv4_internet_host_list"].([]interface{}))
		ret.Ipv6InternetHostList = getObjectSystemIpThreatListIpv6InternetHostList(in["ipv6_internet_host_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemIpThreatListSamplingEnable(d []interface{}) []edpt.SystemIpThreatListSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemIpThreatListSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4SourceList(d []interface{}) edpt.SystemIpThreatListIpv4SourceList {
	var ret edpt.SystemIpThreatListIpv4SourceList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4SourceListClassListCfg(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4SourceListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv4SourceListClassListCfg {
	count := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4SourceListClassListCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4SourceListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4DestList(d []interface{}) edpt.SystemIpThreatListIpv4DestList {
	var ret edpt.SystemIpThreatListIpv4DestList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4DestListClassListCfg(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4DestListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv4DestListClassListCfg {
	count := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4DestListClassListCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4DestListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6SourceList(d []interface{}) edpt.SystemIpThreatListIpv6SourceList {
	var ret edpt.SystemIpThreatListIpv6SourceList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6SourceListClassListCfg(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6SourceListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv6SourceListClassListCfg {
	count := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6SourceListClassListCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6SourceListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6DestList(d []interface{}) edpt.SystemIpThreatListIpv6DestList {
	var ret edpt.SystemIpThreatListIpv6DestList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6DestListClassListCfg(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6DestListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv6DestListClassListCfg {
	count := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6DestListClassListCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6DestListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4InternetHostList(d []interface{}) edpt.SystemIpThreatListIpv4InternetHostList {
	var ret edpt.SystemIpThreatListIpv4InternetHostList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4InternetHostListClassListCfg(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4InternetHostListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv4InternetHostListClassListCfg {
	count := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4InternetHostListClassListCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4InternetHostListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6InternetHostList(d []interface{}) edpt.SystemIpThreatListIpv6InternetHostList {
	var ret edpt.SystemIpThreatListIpv6InternetHostList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6InternetHostListClassListCfg(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6InternetHostListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv6InternetHostListClassListCfg {
	count := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6InternetHostListClassListCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6InternetHostListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIp6Stats(d []interface{}) edpt.SystemIp6Stats {
	var ret edpt.SystemIp6Stats
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemIp6StatsSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemIp6StatsSamplingEnable(d []interface{}) []edpt.SystemIp6StatsSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemIp6StatsSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemIp6StatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpmi(d []interface{}) edpt.SystemIpmi {
	var ret edpt.SystemIpmi
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Reset = in["reset"].(int)
		ret.Ip = getObjectSystemIpmiIp(in["ip"].([]interface{}))
		ret.Ipsrc = getObjectSystemIpmiIpsrc(in["ipsrc"].([]interface{}))
		ret.User = getObjectSystemIpmiUser(in["user"].([]interface{}))
		ret.Tool = getObjectSystemIpmiTool(in["tool"].([]interface{}))
	}
	return ret
}

func getObjectSystemIpmiIp(d []interface{}) edpt.SystemIpmiIp {
	var ret edpt.SystemIpmiIp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.DefaultGateway = in["default_gateway"].(string)
	}
	return ret
}

func getObjectSystemIpmiIpsrc(d []interface{}) edpt.SystemIpmiIpsrc {
	var ret edpt.SystemIpmiIpsrc
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.Static = in["static"].(int)
	}
	return ret
}

func getObjectSystemIpmiUser(d []interface{}) edpt.SystemIpmiUser {
	var ret edpt.SystemIpmiUser
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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

func getObjectSystemIpmiTool(d []interface{}) edpt.SystemIpmiTool {
	var ret edpt.SystemIpmiTool
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Cmd = in["cmd"].(string)
	}
	return ret
}

func getObjectSystemIpmiService(d []interface{}) edpt.SystemIpmiService {
	var ret edpt.SystemIpmiService
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Disable = in["disable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemIpsec(d []interface{}) edpt.SystemIpsec {
	var ret edpt.SystemIpsec
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PacketRoundRobin = in["packet_round_robin"].(int)
		ret.CryptoCore = in["crypto_core"].(int)
		ret.CryptoMem = in["crypto_mem"].(int)
		//omit uuid
		ret.FpgaDecrypt = getObjectSystemIpsecFpgaDecrypt(in["fpga_decrypt"].([]interface{}))
	}
	return ret
}

func getObjectSystemIpsecFpgaDecrypt(d []interface{}) edpt.SystemIpsecFpgaDecrypt {
	var ret edpt.SystemIpsecFpgaDecrypt
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Action = in["action"].(string)
	}
	return ret
}

func getObjectSystemLinkCapability(d []interface{}) edpt.SystemLinkCapability {
	var ret edpt.SystemLinkCapability
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemLinkMonitor(d []interface{}) edpt.SystemLinkMonitor {
	var ret edpt.SystemLinkMonitor
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemLro(d []interface{}) edpt.SystemLro {
	var ret edpt.SystemLro
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemManagementInterfaceMode(d []interface{}) edpt.SystemManagementInterfaceMode {
	var ret edpt.SystemManagementInterfaceMode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Dedicated = in["dedicated"].(int)
		ret.NonDedicated = in["non_dedicated"].(int)
	}
	return ret
}

func getObjectSystemMemory(d []interface{}) edpt.SystemMemory {
	var ret edpt.SystemMemory
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemMemorySamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemMemorySamplingEnable(d []interface{}) []edpt.SystemMemorySamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemMemorySamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemMemorySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemMfaAuth(d []interface{}) edpt.SystemMfaAuth {
	var ret edpt.SystemMfaAuth
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Username = in["username"].(string)
		ret.SecondFactor = in["second_factor"].(string)
	}
	return ret
}

func getObjectSystemMfaCertStore(d []interface{}) edpt.SystemMfaCertStore {
	var ret edpt.SystemMfaCertStore
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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

func getObjectSystemMfaManagement(d []interface{}) edpt.SystemMfaManagement {
	var ret edpt.SystemMfaManagement
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemMfaValidationType(d []interface{}) edpt.SystemMfaValidationType {
	var ret edpt.SystemMfaValidationType
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.CaCert = in["ca_cert"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemMgmtPort(d []interface{}) edpt.SystemMgmtPort {
	var ret edpt.SystemMgmtPort
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.MacAddress = in["mac_address"].(string)
		ret.PciAddress = in["pci_address"].(string)
	}
	return ret
}

func getObjectSystemModifyPort(d []interface{}) edpt.SystemModifyPort {
	var ret edpt.SystemModifyPort
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.PortNumber = in["port_number"].(int)
	}
	return ret
}

func getObjectSystemMonTemplate(d []interface{}) edpt.SystemMonTemplate {
	var ret edpt.SystemMonTemplate
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.MonitorList = getSliceSystemMonTemplateMonitorList(in["monitor_list"].([]interface{}))
		ret.LinkBlockAsDown = getObjectSystemMonTemplateLinkBlockAsDown(in["link_block_as_down"].([]interface{}))
		ret.LinkDownOnRestart = getObjectSystemMonTemplateLinkDownOnRestart(in["link_down_on_restart"].([]interface{}))
	}
	return ret
}

func getSliceSystemMonTemplateMonitorList(d []interface{}) []edpt.SystemMonTemplateMonitorList {
	count := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorList
		oi.Id = in["id"].(int)
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
	count := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListClearCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListClearCfg
		oi.Sessions = in["sessions"].(string)
		oi.ClearAllSequence = in["clear_all_sequence"].(int)
		oi.ClearSequence = in["clear_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListLinkDisableCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListLinkDisableCfg {
	count := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkDisableCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListLinkDisableCfg
		oi.Diseth = in["diseth"].(int)
		oi.DisSequence = in["dis_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListLinkEnableCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListLinkEnableCfg {
	count := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkEnableCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorListLinkEnableCfg
		oi.Enaeth = in["enaeth"].(int)
		oi.EnaSequence = in["ena_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorListLinkUpCfg(d []interface{}) []edpt.SystemMonTemplateMonitorListLinkUpCfg {
	count := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkUpCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorListLinkDownCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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

func getObjectSystemMonTemplateLinkBlockAsDown(d []interface{}) edpt.SystemMonTemplateLinkBlockAsDown {
	var ret edpt.SystemMonTemplateLinkBlockAsDown
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemMonTemplateLinkDownOnRestart(d []interface{}) edpt.SystemMonTemplateLinkDownOnRestart {
	var ret edpt.SystemMonTemplateLinkDownOnRestart
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemMultiQueueSupport(d []interface{}) edpt.SystemMultiQueueSupport {
	var ret edpt.SystemMultiQueueSupport
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
	}
	return ret
}

func getObjectSystemNdiscRa(d []interface{}) edpt.SystemNdiscRa {
	var ret edpt.SystemNdiscRa
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemNdiscRaSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemNdiscRaSamplingEnable(d []interface{}) []edpt.SystemNdiscRaSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemNdiscRaSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemNdiscRaSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemNsmA10lb(d []interface{}) edpt.SystemNsmA10lb {
	var ret edpt.SystemNsmA10lb
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Kill = in["kill"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemPasswordPolicy(d []interface{}) edpt.SystemPasswordPolicy {
	var ret edpt.SystemPasswordPolicy
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Complexity = in["complexity"].(string)
		ret.Aging = in["aging"].(string)
		ret.History = in["history"].(string)
		ret.MinPswdLen = in["min_pswd_len"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemPerVlanLimit(d []interface{}) edpt.SystemPerVlanLimit {
	var ret edpt.SystemPerVlanLimit
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bcast = in["bcast"].(int)
		ret.Ipmcast = in["ipmcast"].(int)
		ret.Mcast = in["mcast"].(int)
		ret.UnknownUcast = in["unknown_ucast"].(int)
		//omit uuid
	}
	return ret
}

//omit object-type field 'platformtype' which contains uuid only

//omit object-type field 'port_info' which contains uuid only

//omit object-type field 'port_list' which contains uuid only

func getObjectSystemPorts(d []interface{}) edpt.SystemPorts {
	var ret edpt.SystemPorts
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.LinkDetectionInterval = in["link_detection_interval"].(int)
		//omit uuid
	}
	return ret
}

//omit empty object-type field 'probe_network_devices'

//omit object-type field 'psu_info' which contains uuid only

func getObjectSystemQInQ(d []interface{}) edpt.SystemQInQ {
	var ret edpt.SystemQInQ
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.InnerTpid = in["inner_tpid"].(string)
		ret.OuterTpid = in["outer_tpid"].(string)
		ret.EnableAllPorts = in["enable_all_ports"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemQueuingBuffer(d []interface{}) edpt.SystemQueuingBuffer {
	var ret edpt.SystemQueuingBuffer
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemRadius(d []interface{}) edpt.SystemRadius {
	var ret edpt.SystemRadius
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Server = getObjectSystemRadiusServer(in["server"].([]interface{}))
	}
	return ret
}

func getObjectSystemRadiusServer(d []interface{}) edpt.SystemRadiusServer {
	var ret edpt.SystemRadiusServer
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ListenPort = in["listen_port"].(int)
		ret.Remote = getObjectSystemRadiusServerRemote(in["remote"].([]interface{}))
		ret.Secret = in["secret"].(int)
		ret.SecretString = in["secret_string"].(string)
		//omit encrypted
		ret.Vrid = in["vrid"].(int)
		ret.Attribute = getSliceSystemRadiusServerAttribute(in["attribute"].([]interface{}))
		ret.DisableReply = in["disable_reply"].(int)
		ret.AccountingStart = in["accounting_start"].(string)
		ret.AccountingStop = in["accounting_stop"].(string)
		ret.AccountingInterimUpdate = in["accounting_interim_update"].(string)
		ret.AccountingOn = in["accounting_on"].(string)
		ret.AttributeName = in["attribute_name"].(string)
		ret.CustomAttributeName = in["custom_attribute_name"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSystemRadiusServerSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getObjectSystemRadiusServerRemote(d []interface{}) edpt.SystemRadiusServerRemote {
	var ret edpt.SystemRadiusServerRemote
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.IpList = getSliceSystemRadiusServerRemoteIpList(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemRadiusServerRemoteIpList(d []interface{}) []edpt.SystemRadiusServerRemoteIpList {
	count := len(d)
	ret := make([]edpt.SystemRadiusServerRemoteIpList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerRemoteIpList
		oi.IpListName = in["ip_list_name"].(string)
		oi.IpListSecret = in["ip_list_secret"].(int)
		oi.IpListSecretString = in["ip_list_secret_string"].(string)
		//omit ip_list_encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemRadiusServerAttribute(d []interface{}) []edpt.SystemRadiusServerAttribute {
	count := len(d)
	ret := make([]edpt.SystemRadiusServerAttribute, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerAttribute
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

func getSliceSystemRadiusServerSamplingEnable(d []interface{}) []edpt.SystemRadiusServerSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemRadiusServerSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

//omit object-type field 'reboot' which contains uuid only

func getObjectSystemResourceAccounting(d []interface{}) edpt.SystemResourceAccounting {
	var ret edpt.SystemResourceAccounting
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.TemplateList = getSliceSystemResourceAccountingTemplateList(in["template_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemResourceAccountingTemplateList(d []interface{}) []edpt.SystemResourceAccountingTemplateList {
	count := len(d)
	ret := make([]edpt.SystemResourceAccountingTemplateList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemResourceAccountingTemplateList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.AppResources = getObjectSystemResourceAccountingTemplateListAppResources(in["app_resources"].([]interface{}))
		oi.NetworkResources = getObjectSystemResourceAccountingTemplateListNetworkResources(in["network_resources"].([]interface{}))
		oi.SystemResources = getObjectSystemResourceAccountingTemplateListSystemResources(in["system_resources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResources(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResources {
	var ret edpt.SystemResourceAccountingTemplateListAppResources
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbDeviceCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg(in["gslb_device_cfg"].([]interface{}))
		ret.GslbGeoLocationCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg(in["gslb_geo_location_cfg"].([]interface{}))
		ret.GslbIpListCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbIpListCfg(in["gslb_ip_list_cfg"].([]interface{}))
		ret.GslbPolicyCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg(in["gslb_policy_cfg"].([]interface{}))
		ret.GslbServiceCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceCfg(in["gslb_service_cfg"].([]interface{}))
		ret.GslbServiceIpCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg(in["gslb_service_ip_cfg"].([]interface{}))
		ret.GslbServicePortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg(in["gslb_service_port_cfg"].([]interface{}))
		ret.GslbSiteCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbSiteCfg(in["gslb_site_cfg"].([]interface{}))
		ret.GslbSvcGroupCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg(in["gslb_svc_group_cfg"].([]interface{}))
		ret.GslbTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg(in["gslb_template_cfg"].([]interface{}))
		ret.GslbZoneCfg = getObjectSystemResourceAccountingTemplateListAppResourcesGslbZoneCfg(in["gslb_zone_cfg"].([]interface{}))
		ret.HealthMonitorCfg = getObjectSystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg(in["health_monitor_cfg"].([]interface{}))
		ret.RealPortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesRealPortCfg(in["real_port_cfg"].([]interface{}))
		ret.RealServerCfg = getObjectSystemResourceAccountingTemplateListAppResourcesRealServerCfg(in["real_server_cfg"].([]interface{}))
		ret.ServiceGroupCfg = getObjectSystemResourceAccountingTemplateListAppResourcesServiceGroupCfg(in["service_group_cfg"].([]interface{}))
		ret.VirtualServerCfg = getObjectSystemResourceAccountingTemplateListAppResourcesVirtualServerCfg(in["virtual_server_cfg"].([]interface{}))
		ret.VirtualPortCfg = getObjectSystemResourceAccountingTemplateListAppResourcesVirtualPortCfg(in["virtual_port_cfg"].([]interface{}))
		ret.CacheTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg(in["cache_template_cfg"].([]interface{}))
		ret.ClientSslTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg(in["client_ssl_template_cfg"].([]interface{}))
		ret.ConnReuseTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg(in["conn_reuse_template_cfg"].([]interface{}))
		ret.FastTcpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg(in["fast_tcp_template_cfg"].([]interface{}))
		ret.FastUdpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg(in["fast_udp_template_cfg"].([]interface{}))
		ret.FixTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesFixTemplateCfg(in["fix_template_cfg"].([]interface{}))
		ret.HttpTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg(in["http_template_cfg"].([]interface{}))
		ret.LinkCostTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg(in["link_cost_template_cfg"].([]interface{}))
		ret.PersistCookieTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg(in["persist_cookie_template_cfg"].([]interface{}))
		ret.PersistSrcipTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg(in["persist_srcip_template_cfg"].([]interface{}))
		ret.ServerSslTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg(in["server_ssl_template_cfg"].([]interface{}))
		ret.ProxyTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg(in["proxy_template_cfg"].([]interface{}))
		ret.StreamTemplateCfg = getObjectSystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg(in["stream_template_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbDeviceMax = in["gslb_device_max"].(int)
		ret.GslbDeviceMinGuarantee = in["gslb_device_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbGeoLocationMax = in["gslb_geo_location_max"].(int)
		ret.GslbGeoLocationMinGuarantee = in["gslb_geo_location_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbIpListCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbIpListMax = in["gslb_ip_list_max"].(int)
		ret.GslbIpListMinGuarantee = in["gslb_ip_list_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbPolicyMax = in["gslb_policy_max"].(int)
		ret.GslbPolicyMinGuarantee = in["gslb_policy_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbServiceMax = in["gslb_service_max"].(int)
		ret.GslbServiceMinGuarantee = in["gslb_service_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbServiceIpMax = in["gslb_service_ip_max"].(int)
		ret.GslbServiceIpMinGuarantee = in["gslb_service_ip_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbServicePortMax = in["gslb_service_port_max"].(int)
		ret.GslbServicePortMinGuarantee = in["gslb_service_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbSiteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbSiteMax = in["gslb_site_max"].(int)
		ret.GslbSiteMinGuarantee = in["gslb_site_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbSvcGroupMax = in["gslb_svc_group_max"].(int)
		ret.GslbSvcGroupMinGuarantee = in["gslb_svc_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbTemplateMax = in["gslb_template_max"].(int)
		ret.GslbTemplateMinGuarantee = in["gslb_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesGslbZoneCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.GslbZoneMax = in["gslb_zone_max"].(int)
		ret.GslbZoneMinGuarantee = in["gslb_zone_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.HealthMonitorMax = in["health_monitor_max"].(int)
		ret.HealthMonitorMinGuarantee = in["health_monitor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesRealPortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesRealPortCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesRealPortCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.RealPortMax = in["real_port_max"].(int)
		ret.RealPortMinGuarantee = in["real_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesRealServerCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesRealServerCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesRealServerCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.RealServerMax = in["real_server_max"].(int)
		ret.RealServerMinGuarantee = in["real_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesServiceGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ServiceGroupMax = in["service_group_max"].(int)
		ret.ServiceGroupMinGuarantee = in["service_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesVirtualServerCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.VirtualServerMax = in["virtual_server_max"].(int)
		ret.VirtualServerMinGuarantee = in["virtual_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesVirtualPortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.VirtualPortMax = in["virtual_port_max"].(int)
		ret.VirtualPortMinGuarantee = in["virtual_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.CacheTemplateMax = in["cache_template_max"].(int)
		ret.CacheTemplateMinGuarantee = in["cache_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ClientSslTemplateMax = in["client_ssl_template_max"].(int)
		ret.ClientSslTemplateMinGuarantee = in["client_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ConnReuseTemplateMax = in["conn_reuse_template_max"].(int)
		ret.ConnReuseTemplateMinGuarantee = in["conn_reuse_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.FastTcpTemplateMax = in["fast_tcp_template_max"].(int)
		ret.FastTcpTemplateMinGuarantee = in["fast_tcp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.FastUdpTemplateMax = in["fast_udp_template_max"].(int)
		ret.FastUdpTemplateMinGuarantee = in["fast_udp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesFixTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.FixTemplateMax = in["fix_template_max"].(int)
		ret.FixTemplateMinGuarantee = in["fix_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.HttpTemplateMax = in["http_template_max"].(int)
		ret.HttpTemplateMinGuarantee = in["http_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.LinkCostTemplateMax = in["link_cost_template_max"].(int)
		ret.LinkCostTemplateMinGuarantee = in["link_cost_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PersistCookieTemplateMax = in["persist_cookie_template_max"].(int)
		ret.PersistCookieTemplateMinGuarantee = in["persist_cookie_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PersistSrcipTemplateMax = in["persist_srcip_template_max"].(int)
		ret.PersistSrcipTemplateMinGuarantee = in["persist_srcip_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ServerSslTemplateMax = in["server_ssl_template_max"].(int)
		ret.ServerSslTemplateMinGuarantee = in["server_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ProxyTemplateMax = in["proxy_template_max"].(int)
		ret.ProxyTemplateMinGuarantee = in["proxy_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg {
	var ret edpt.SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StreamTemplateMax = in["stream_template_max"].(int)
		ret.StreamTemplateMinGuarantee = in["stream_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResources(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResources {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResources
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticIpv4RouteCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg(in["static_ipv4_route_cfg"].([]interface{}))
		ret.StaticIpv6RouteCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg(in["static_ipv6_route_cfg"].([]interface{}))
		ret.Ipv4AclLineCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg(in["ipv4_acl_line_cfg"].([]interface{}))
		ret.Ipv6AclLineCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg(in["ipv6_acl_line_cfg"].([]interface{}))
		ret.StaticArpCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg(in["static_arp_cfg"].([]interface{}))
		ret.StaticNeighborCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg(in["static_neighbor_cfg"].([]interface{}))
		ret.StaticMacCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg(in["static_mac_cfg"].([]interface{}))
		ret.ObjectGroupCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg(in["object_group_cfg"].([]interface{}))
		ret.ObjectGroupClauseCfg = getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg(in["object_group_clause_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticIpv4RouteMax = in["static_ipv4_route_max"].(int)
		ret.StaticIpv4RouteMinGuarantee = in["static_ipv4_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticIpv6RouteMax = in["static_ipv6_route_max"].(int)
		ret.StaticIpv6RouteMinGuarantee = in["static_ipv6_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ipv4AclLineMax = in["ipv4_acl_line_max"].(int)
		ret.Ipv4AclLineMinGuarantee = in["ipv4_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ipv6AclLineMax = in["ipv6_acl_line_max"].(int)
		ret.Ipv6AclLineMinGuarantee = in["ipv6_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticArpMax = in["static_arp_max"].(int)
		ret.StaticArpMinGuarantee = in["static_arp_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticNeighborMax = in["static_neighbor_max"].(int)
		ret.StaticNeighborMinGuarantee = in["static_neighbor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticMacMax = in["static_mac_max"].(int)
		ret.StaticMacMinGuarantee = in["static_mac_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ObjectGroupMax = in["object_group_max"].(int)
		ret.ObjectGroupMinGuarantee = in["object_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg {
	var ret edpt.SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ObjectGroupClauseMax = in["object_group_clause_max"].(int)
		ret.ObjectGroupClauseMinGuarantee = in["object_group_clause_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResources(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResources {
	var ret edpt.SystemResourceAccountingTemplateListSystemResources
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.BwLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesBwLimitCfg(in["bw_limit_cfg"].([]interface{}))
		ret.ConcurrentSessionLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg(in["concurrent_session_limit_cfg"].([]interface{}))
		ret.L4SessionLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg(in["l4_session_limit_cfg"].([]interface{}))
		ret.L4cpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg(in["l4cps_limit_cfg"].([]interface{}))
		ret.L7cpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg(in["l7cps_limit_cfg"].([]interface{}))
		ret.NatcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg(in["natcps_limit_cfg"].([]interface{}))
		ret.FwcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg(in["fwcps_limit_cfg"].([]interface{}))
		ret.SslThroughputLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg(in["ssl_throughput_limit_cfg"].([]interface{}))
		ret.SslcpsLimitCfg = getObjectSystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg(in["sslcps_limit_cfg"].([]interface{}))
		ret.Threshold = in["threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesBwLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.BwLimitMax = in["bw_limit_max"].(int)
		ret.BwLimitWatermarkDisable = in["bw_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ConcurrentSessionLimitMax = in["concurrent_session_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.L4SessionLimitMax = in["l4_session_limit_max"].(string)
		ret.L4SessionLimitMinGuarantee = in["l4_session_limit_min_guarantee"].(string)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.L4cpsLimitMax = in["l4cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.L7cpsLimitMax = in["l7cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NatcpsLimitMax = in["natcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.FwcpsLimitMax = in["fwcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SslThroughputLimitMax = in["ssl_throughput_limit_max"].(int)
		ret.SslThroughputLimitWatermarkDisable = in["ssl_throughput_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg {
	var ret edpt.SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SslcpsLimitMax = in["sslcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceUsage(d []interface{}) edpt.SystemResourceUsage {
	var ret edpt.SystemResourceUsage
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SslContextMemory = in["ssl_context_memory"].(int)
		ret.SslDmaMemory = in["ssl_dma_memory"].(int)
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
		ret.WafTemplateCount = in["waf_template_count"].(int)
		ret.AuthSessionCount = in["auth_session_count"].(int)
		//omit uuid
		ret.Visibility = getObjectSystemResourceUsageVisibility(in["visibility"].([]interface{}))
	}
	return ret
}

func getObjectSystemResourceUsageVisibility(d []interface{}) edpt.SystemResourceUsageVisibility {
	var ret edpt.SystemResourceUsageVisibility
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.MonitoredEntityCount = in["monitored_entity_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSession(d []interface{}) edpt.SystemSession {
	var ret edpt.SystemSession
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemSessionSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemSessionSamplingEnable(d []interface{}) []edpt.SystemSessionSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemSessionSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemSessionSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemSessionReclaimLimit(d []interface{}) edpt.SystemSessionReclaimLimit {
	var ret edpt.SystemSessionReclaimLimit
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NscanLimit = in["nscan_limit"].(int)
		ret.ScanFreq = in["scan_freq"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSetRxtxDescSize(d []interface{}) edpt.SystemSetRxtxDescSize {
	var ret edpt.SystemSetRxtxDescSize
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.RxdSize = in["rxd_size"].(int)
		ret.TxdSize = in["txd_size"].(int)
	}
	return ret
}

func getObjectSystemSetRxtxQueue(d []interface{}) edpt.SystemSetRxtxQueue {
	var ret edpt.SystemSetRxtxQueue
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PortIndex = in["port_index"].(int)
		ret.RxqSize = in["rxq_size"].(int)
		ret.TxqSize = in["txq_size"].(int)
	}
	return ret
}

func getObjectSystemSetTcpSynPerSec(d []interface{}) edpt.SystemSetTcpSynPerSec {
	var ret edpt.SystemSetTcpSynPerSec
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.TcpSynValue = in["tcp_syn_value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSharedPollMode(d []interface{}) edpt.SystemSharedPollMode {
	var ret edpt.SystemSharedPollMode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectSystemShellPrivileges(d []interface{}) edpt.SystemShellPrivileges {
	var ret edpt.SystemShellPrivileges
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.EnableShellPrivileges = in["enable_shell_privileges"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemShmLogging(d []interface{}) edpt.SystemShmLogging {
	var ret edpt.SystemShmLogging
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

//omit object-type field 'shutdown' which contains uuid only

func getObjectSystemSpeProfile(d []interface{}) edpt.SystemSpeProfile {
	var ret edpt.SystemSpeProfile
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Action = in["action"].(string)
	}
	return ret
}

//omit object-type field 'spe_status' which contains uuid only

func getObjectSystemSslReqQ(d []interface{}) edpt.SystemSslReqQ {
	var ret edpt.SystemSslReqQ
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemSslReqQSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemSslReqQSamplingEnable(d []interface{}) []edpt.SystemSslReqQSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemSslReqQSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemSslReqQSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemSslScv(d []interface{}) edpt.SystemSslScv {
	var ret edpt.SystemSslScv
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSslScvVerifyCrlSign(d []interface{}) edpt.SystemSslScvVerifyCrlSign {
	var ret edpt.SystemSslScvVerifyCrlSign
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSslScvVerifyHost(d []interface{}) edpt.SystemSslScvVerifyHost {
	var ret edpt.SystemSslScvVerifyHost
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Disable = in["disable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemSslSetCompatibleCipher(d []interface{}) edpt.SystemSslSetCompatibleCipher {
	var ret edpt.SystemSslSetCompatibleCipher
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Disable = in["disable"].(int)
		//omit uuid
	}
	return ret
}

//omit object-type field 'ssl_status' which contains uuid only

func getObjectSystemSyslogTimeMsec(d []interface{}) edpt.SystemSyslogTimeMsec {
	var ret edpt.SystemSyslogTimeMsec
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.EnableFlag = in["enable_flag"].(int)
	}
	return ret
}

func getObjectSystemTableIntegrity(d []interface{}) edpt.SystemTableIntegrity {
	var ret edpt.SystemTableIntegrity
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Table = in["table"].(string)
		ret.AuditAction = in["audit_action"].(string)
		ret.AutoSyncAction = in["auto_sync_action"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSystemTableIntegritySamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemTableIntegritySamplingEnable(d []interface{}) []edpt.SystemTableIntegritySamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemTableIntegritySamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemTableIntegritySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		oi.Counters3 = in["counters3"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTcp(d []interface{}) edpt.SystemTcp {
	var ret edpt.SystemTcp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemTcpSamplingEnable(in["sampling_enable"].([]interface{}))
		ret.RateLimitResetUnknownConn = getObjectSystemTcpRateLimitResetUnknownConn(in["rate_limit_reset_unknown_conn"].([]interface{}))
	}
	return ret
}

func getSliceSystemTcpSamplingEnable(d []interface{}) []edpt.SystemTcpSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemTcpSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemTcpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTcpRateLimitResetUnknownConn(d []interface{}) edpt.SystemTcpRateLimitResetUnknownConn {
	var ret edpt.SystemTcpRateLimitResetUnknownConn
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PktRateForResetUnknownConn = in["pkt_rate_for_reset_unknown_conn"].(int)
		ret.LogForResetUnknownConn = in["log_for_reset_unknown_conn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTcpStats(d []interface{}) edpt.SystemTcpStats {
	var ret edpt.SystemTcpStats
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemTcpStatsSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemTcpStatsSamplingEnable(d []interface{}) []edpt.SystemTcpStatsSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemTcpStatsSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemTcpStatsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTemplate(d []interface{}) edpt.SystemTemplate {
	var ret edpt.SystemTemplate
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.TemplatePolicy = in["template_policy"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSystemTemplateBind(d []interface{}) edpt.SystemTemplateBind {
	var ret edpt.SystemTemplateBind
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.MonitorList = getSliceSystemTemplateBindMonitorList(in["monitor_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemTemplateBindMonitorList(d []interface{}) []edpt.SystemTemplateBindMonitorList {
	count := len(d)
	ret := make([]edpt.SystemTemplateBindMonitorList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemTemplateBindMonitorList
		oi.TemplateMonitor = in["template_monitor"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemThroughput(d []interface{}) edpt.SystemThroughput {
	var ret edpt.SystemThroughput
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemThroughputSamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemThroughputSamplingEnable(d []interface{}) []edpt.SystemThroughputSamplingEnable {
	count := len(d)
	ret := make([]edpt.SystemThroughputSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SystemThroughputSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTimeoutValue(d []interface{}) edpt.SystemTimeoutValue {
	var ret edpt.SystemTimeoutValue
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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

func getObjectSystemTls13Mgmt(d []interface{}) edpt.SystemTls13Mgmt {
	var ret edpt.SystemTls13Mgmt
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTrunk(d []interface{}) edpt.SystemTrunk {
	var ret edpt.SystemTrunk
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.LoadBalance = getObjectSystemTrunkLoadBalance(in["load_balance"].([]interface{}))
	}
	return ret
}

func getObjectSystemTrunkLoadBalance(d []interface{}) edpt.SystemTrunkLoadBalance {
	var ret edpt.SystemTrunkLoadBalance
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.UseL3 = in["use_l3"].(int)
		ret.UseL4 = in["use_l4"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTrunkHwHash(d []interface{}) edpt.SystemTrunkHwHash {
	var ret edpt.SystemTrunkHwHash
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Mode = in["mode"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTrunkXauiHwHash(d []interface{}) edpt.SystemTrunkXauiHwHash {
	var ret edpt.SystemTrunkXauiHwHash
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Mode = in["mode"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemTso(d []interface{}) edpt.SystemTso {
	var ret edpt.SystemTso
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func dataToEndpointSystem(d *schema.ResourceData) edpt.System {
	var ret edpt.System
	ret.Inst.AddCpuCore = getObjectSystemAddCpuCore(d.Get("add_cpu_core").([]interface{}))
	ret.Inst.AddPort = getObjectSystemAddPort(d.Get("add_port").([]interface{}))
	ret.Inst.AllVlanLimit = getObjectSystemAllVlanLimit(d.Get("all_vlan_limit").([]interface{}))
	ret.Inst.AnomalyLog = d.Get("anomaly_log").(int)
	ret.Inst.AppPerformance = getObjectSystemAppPerformance(d.Get("app_performance").([]interface{}))
	ret.Inst.AppsGlobal = getObjectSystemAppsGlobal(d.Get("apps_global").([]interface{}))
	ret.Inst.AsicDebugDump = getObjectSystemAsicDebugDump(d.Get("asic_debug_dump").([]interface{}))
	ret.Inst.AttackLog = d.Get("attack_log").(int)
	ret.Inst.Bandwidth = getObjectSystemBandwidth(d.Get("bandwidth").([]interface{}))
	ret.Inst.Bfd = getObjectSystemBfd(d.Get("bfd").([]interface{}))
	ret.Inst.ClassListHitcountEnable = d.Get("class_list_hitcount_enable").(int)
	ret.Inst.CmUpdateFileNameRef = getObjectSystemCmUpdateFileNameRef(d.Get("cm_update_file_name_ref").([]interface{}))
	//omit field 'control_cpu' which contains uuid only
	//omit field 'core' which contains uuid only
	//omit field 'cosq_show' which contains uuid only
	//omit field 'cosq_stats' which contains uuid only
	//omit field 'counter_lib_accounting' which contains uuid only
	ret.Inst.CpuHyperThread = getObjectSystemCpuHyperThread(d.Get("cpu_hyper_thread").([]interface{}))
	//omit field 'cpu_list' which contains uuid only
	ret.Inst.CpuLoadSharing = getObjectSystemCpuLoadSharing(d.Get("cpu_load_sharing").([]interface{}))
	//omit field 'cpu_map' which contains uuid only
	//omit field 'data_cpu' which contains uuid only
	ret.Inst.DdosAttack = d.Get("ddos_attack").(int)
	ret.Inst.DdosLog = d.Get("ddos_log").(int)
	ret.Inst.DeepHrxq = getObjectSystemDeepHrxq(d.Get("deep_hrxq").([]interface{}))
	ret.Inst.DelPort = getObjectSystemDelPort(d.Get("del_port").([]interface{}))
	ret.Inst.DeleteCpuCore = getObjectSystemDeleteCpuCore(d.Get("delete_cpu_core").([]interface{}))
	ret.Inst.Dns = getObjectSystemDns(d.Get("dns").([]interface{}))
	ret.Inst.DnsCache = getObjectSystemDnsCache(d.Get("dns_cache").([]interface{}))
	ret.Inst.DomainListHitcountEnable = d.Get("domain_list_hitcount_enable").(int)
	//omit field 'domain_list_info' which contains uuid only
	ret.Inst.DpdkStats = getObjectSystemDpdkStats(d.Get("dpdk_stats").([]interface{}))
	ret.Inst.DynamicServiceDnsSocketPool = d.Get("dynamic_service_dns_socket_pool").(int)
	//omit field 'environment' which contains uuid only
	ret.Inst.FpgaCoreCrc = getObjectSystemFpgaCoreCrc(d.Get("fpga_core_crc").([]interface{}))
	ret.Inst.FpgaDrop = getObjectSystemFpgaDrop(d.Get("fpga_drop").([]interface{}))
	ret.Inst.Fw = getObjectSystemFw(d.Get("fw").([]interface{}))
	ret.Inst.GeoDbHitcountEnable = d.Get("geo_db_hitcount_enable").(int)
	ret.Inst.GeoLocation = getObjectSystemGeoLocation(d.Get("geo_location").([]interface{}))
	ret.Inst.Geoloc = getObjectSystemGeoloc(d.Get("geoloc").([]interface{}))
	ret.Inst.GeolocListList = getSliceSystemGeolocListList(d.Get("geoloc_list_list").([]interface{}))
	ret.Inst.GeolocNameHelper = getObjectSystemGeolocNameHelper(d.Get("geoloc_name_helper").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(int)
	//omit field 'guest_file' which contains uuid only
	//omit field 'gui_image_list' which contains uuid only
	//omit field 'hardware' which contains uuid only
	ret.Inst.HardwareForward = getObjectSystemHardwareForward(d.Get("hardware_forward").([]interface{}))
	ret.Inst.HighMemoryL4Session = getObjectSystemHighMemoryL4Session(d.Get("high_memory_l4_session").([]interface{}))
	//omit field 'hrxq_status' which contains uuid only
	ret.Inst.Icmp = getObjectSystemIcmp(d.Get("icmp").([]interface{}))
	ret.Inst.IcmpRate = getObjectSystemIcmpRate(d.Get("icmp_rate").([]interface{}))
	ret.Inst.Icmp6 = getObjectSystemIcmp6(d.Get("icmp6").([]interface{}))
	//omit field 'inuse_cpu_list' which contains uuid only
	//omit field 'inuse_port_list' which contains uuid only
	ret.Inst.IoCpu = getObjectSystemIoCpu(d.Get("io_cpu").([]interface{}))
	//omit field 'ip_dns_cache' which contains uuid only
	ret.Inst.IpStats = getObjectSystemIpStats(d.Get("ip_stats").([]interface{}))
	ret.Inst.IpThreatList = getObjectSystemIpThreatList(d.Get("ip_threat_list").([]interface{}))
	ret.Inst.Ip6Stats = getObjectSystemIp6Stats(d.Get("ip6_stats").([]interface{}))
	ret.Inst.Ipmi = getObjectSystemIpmi(d.Get("ipmi").([]interface{}))
	ret.Inst.IpmiService = getObjectSystemIpmiService(d.Get("ipmi_service").([]interface{}))
	ret.Inst.Ipsec = getObjectSystemIpsec(d.Get("ipsec").([]interface{}))
	ret.Inst.Ipv6PrefixLength = d.Get("ipv6_prefix_length").(int)
	ret.Inst.LinkCapability = getObjectSystemLinkCapability(d.Get("link_capability").([]interface{}))
	ret.Inst.LinkMonitor = getObjectSystemLinkMonitor(d.Get("link_monitor").([]interface{}))
	ret.Inst.Lro = getObjectSystemLro(d.Get("lro").([]interface{}))
	ret.Inst.ManagementInterfaceMode = getObjectSystemManagementInterfaceMode(d.Get("management_interface_mode").([]interface{}))
	ret.Inst.Memory = getObjectSystemMemory(d.Get("memory").([]interface{}))
	ret.Inst.MfaAuth = getObjectSystemMfaAuth(d.Get("mfa_auth").([]interface{}))
	ret.Inst.MfaCertStore = getObjectSystemMfaCertStore(d.Get("mfa_cert_store").([]interface{}))
	ret.Inst.MfaManagement = getObjectSystemMfaManagement(d.Get("mfa_management").([]interface{}))
	ret.Inst.MfaValidationType = getObjectSystemMfaValidationType(d.Get("mfa_validation_type").([]interface{}))
	ret.Inst.MgmtPort = getObjectSystemMgmtPort(d.Get("mgmt_port").([]interface{}))
	ret.Inst.ModifyPort = getObjectSystemModifyPort(d.Get("modify_port").([]interface{}))
	ret.Inst.ModuleCtrlCpu = d.Get("module_ctrl_cpu").(string)
	ret.Inst.MonTemplate = getObjectSystemMonTemplate(d.Get("mon_template").([]interface{}))
	ret.Inst.MultiQueueSupport = getObjectSystemMultiQueueSupport(d.Get("multi_queue_support").([]interface{}))
	ret.Inst.NdiscRa = getObjectSystemNdiscRa(d.Get("ndisc_ra").([]interface{}))
	ret.Inst.NsmA10lb = getObjectSystemNsmA10lb(d.Get("nsm_a10lb").([]interface{}))
	ret.Inst.PasswordPolicy = getObjectSystemPasswordPolicy(d.Get("password_policy").([]interface{}))
	ret.Inst.PerVlanLimit = getObjectSystemPerVlanLimit(d.Get("per_vlan_limit").([]interface{}))
	//omit field 'platformtype' which contains uuid only
	//omit field 'port_info' which contains uuid only
	//omit field 'port_list' which contains uuid only
	ret.Inst.Ports = getObjectSystemPorts(d.Get("ports").([]interface{}))
	//omit empty field 'probe_network_devices'
	ret.Inst.PromiscuousMode = d.Get("promiscuous_mode").(int)
	//omit field 'psu_info' which contains uuid only
	ret.Inst.QInQ = getObjectSystemQInQ(d.Get("q_in_q").([]interface{}))
	ret.Inst.QueuingBuffer = getObjectSystemQueuingBuffer(d.Get("queuing_buffer").([]interface{}))
	ret.Inst.Radius = getObjectSystemRadius(d.Get("radius").([]interface{}))
	//omit field 'reboot' which contains uuid only
	ret.Inst.ResourceAccounting = getObjectSystemResourceAccounting(d.Get("resource_accounting").([]interface{}))
	ret.Inst.ResourceUsage = getObjectSystemResourceUsage(d.Get("resource_usage").([]interface{}))
	ret.Inst.Session = getObjectSystemSession(d.Get("session").([]interface{}))
	ret.Inst.SessionReclaimLimit = getObjectSystemSessionReclaimLimit(d.Get("session_reclaim_limit").([]interface{}))
	ret.Inst.SetRxtxDescSize = getObjectSystemSetRxtxDescSize(d.Get("set_rxtx_desc_size").([]interface{}))
	ret.Inst.SetRxtxQueue = getObjectSystemSetRxtxQueue(d.Get("set_rxtx_queue").([]interface{}))
	ret.Inst.SetTcpSynPerSec = getObjectSystemSetTcpSynPerSec(d.Get("set_tcp_syn_per_sec").([]interface{}))
	ret.Inst.SharedPollMode = getObjectSystemSharedPollMode(d.Get("shared_poll_mode").([]interface{}))
	ret.Inst.ShellPrivileges = getObjectSystemShellPrivileges(d.Get("shell_privileges").([]interface{}))
	ret.Inst.ShmLogging = getObjectSystemShmLogging(d.Get("shm_logging").([]interface{}))
	//omit field 'shutdown' which contains uuid only
	ret.Inst.SockstressDisable = d.Get("sockstress_disable").(int)
	ret.Inst.SpeProfile = getObjectSystemSpeProfile(d.Get("spe_profile").([]interface{}))
	//omit field 'spe_status' which contains uuid only
	ret.Inst.SrcIpHashEnable = d.Get("src_ip_hash_enable").(int)
	ret.Inst.SslReqQ = getObjectSystemSslReqQ(d.Get("ssl_req_q").([]interface{}))
	ret.Inst.SslScv = getObjectSystemSslScv(d.Get("ssl_scv").([]interface{}))
	ret.Inst.SslScvVerifyCrlSign = getObjectSystemSslScvVerifyCrlSign(d.Get("ssl_scv_verify_crl_sign").([]interface{}))
	ret.Inst.SslScvVerifyHost = getObjectSystemSslScvVerifyHost(d.Get("ssl_scv_verify_host").([]interface{}))
	ret.Inst.SslSetCompatibleCipher = getObjectSystemSslSetCompatibleCipher(d.Get("ssl_set_compatible_cipher").([]interface{}))
	//omit field 'ssl_status' which contains uuid only
	ret.Inst.SyslogTimeMsec = getObjectSystemSyslogTimeMsec(d.Get("syslog_time_msec").([]interface{}))
	ret.Inst.TableIntegrity = getObjectSystemTableIntegrity(d.Get("table_integrity").([]interface{}))
	ret.Inst.Tcp = getObjectSystemTcp(d.Get("tcp").([]interface{}))
	ret.Inst.TcpStats = getObjectSystemTcpStats(d.Get("tcp_stats").([]interface{}))
	//omit field 'tcp_syn_per_sec' which contains uuid only
	ret.Inst.Template = getObjectSystemTemplate(d.Get("template").([]interface{}))
	ret.Inst.TemplateBind = getObjectSystemTemplateBind(d.Get("template_bind").([]interface{}))
	ret.Inst.Throughput = getObjectSystemThroughput(d.Get("throughput").([]interface{}))
	ret.Inst.TimeoutValue = getObjectSystemTimeoutValue(d.Get("timeout_value").([]interface{}))
	ret.Inst.Tls13Mgmt = getObjectSystemTls13Mgmt(d.Get("tls_1_3_mgmt").([]interface{}))
	ret.Inst.Trunk = getObjectSystemTrunk(d.Get("trunk").([]interface{}))
	ret.Inst.TrunkHwHash = getObjectSystemTrunkHwHash(d.Get("trunk_hw_hash").([]interface{}))
	ret.Inst.TrunkXauiHwHash = getObjectSystemTrunkXauiHwHash(d.Get("trunk_xaui_hw_hash").([]interface{}))
	ret.Inst.Tso = getObjectSystemTso(d.Get("tso").([]interface{}))
	//omit field 'upgrade_status' which contains uuid only
	//omit uuid
	return ret
}
