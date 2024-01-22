package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_tcp`: TCP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateTcpCreate,
		UpdateContext: resourceDdosZoneTemplateTcpUpdate,
		ReadContext:   resourceDdosZoneTemplateTcpRead,
		DeleteContext: resourceDdosZoneTemplateTcpDelete,

		Schema: map[string]*schema.Schema{
			"ack_authentication": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ack_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "ack retransmit timeout in seconds(default timeout: 5 seconds)",
						},
						"ack_auth_min_delay": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum delay (in 100ms intervals) between ACK retransmits for retransmit-check to pass",
						},
						"ack_auth_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply retransmit-check only once per source address for authentication purpose",
						},
						"ack_auth_rto": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Estimate the RTO and apply the exponential back-off for authentication",
						},
						"ack_auth_pass_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for passing the authentication",
						},
						"ack_auth_pass_action": {
							Type: schema.TypeString, Optional: true, Description: "'authenticate-src': authenticate-src (Default);",
						},
						"ack_auth_fail_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for failing the authentication.",
						},
						"ack_auth_fail_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Send reset to client;",
						},
					},
				},
			},
			"ack_authentication_synack_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset client TCP SYN+ACK for authentication (DST support only)",
			},
			"action_on_ack_rto_retry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Take action if ack-auth RTO-authentication fail over retry time(default:5)",
			},
			"action_on_syn_rto_retry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Take action if syn-auth RTO-authentication fail over retry time(default:5)",
			},
			"age": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Session age in minutes",
			},
			"allow_syn_otherflags": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Treat TCP SYN+PSH as a TCP SYN (DST tcp ports support only)",
			},
			"allow_synack_skip_authentications": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow create sessions on SYNACK without syn-auth and ack-auth (ASYM Mode only)",
			},
			"allow_tcp_tfo": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow TCP Fast Open",
			},
			"concurrent": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable concurrent port access for non-matching ports (DST support only)",
			},
			"conn_rate_limit_on_syn_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only count SYN-initiated connections towards connection-rate tracking",
			},
			"create_conn_on_syn_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable connection establishment on SYN only",
			},
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"syn_rate_limit": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dst_syn_rate_limit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dst_syn_rate_action": {
													Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for syn-rate exceed (Default); 'ignore': Ignore syn-rate-exceed;",
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
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"tcp_filter_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"tcp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"tcp_filter_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
						},
						"tcp_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"tcp_filter_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src;",
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
			"filter_match_type": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Stop matching on drop/blacklist action; 'stop-on-first-match': Stop matching on first match;",
			},
			"known_resp_src_port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"known_resp_src_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take action if src-port is less than 1024",
						},
						"known_resp_src_port_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for well-known src-port",
						},
						"known_resp_src_port_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets from well-known src-port(Default); 'blacklist-src': Blacklist-src from well-known src-port; 'ignore': Ignore well-known src-port;",
						},
						"exclude_src_resp_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Exclude src port equal to dst port",
						},
					},
				},
			},
			"max_rexmit_syn_per_flow_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_rexmit_syn_per_flow": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of re-transmit SYN per flow",
						},
						"max_rexmit_syn_per_flow_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for max-rexmit-syn-per-flow exceed",
						},
						"max_rexmit_syn_per_flow_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop SYN packets for max-rexmit-syn-per-flow exceed (Default); 'blacklist-src': help Blacklist-src for max-rexmit-syn-per-flow exceed;",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"out_of_seq_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"out_of_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Take action if out-of-seq pkts exceed configured threshold",
						},
						"out_of_seq_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for out-of-seq exceed",
						},
						"out_of_seq_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for out-of-seq exceed (Default); 'blacklist-src': help Blacklist-src for out-of-seq exceed; 'ignore': help Ignore out-of-seq exceed;",
						},
					},
				},
			},
			"per_conn_out_of_seq_rate_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per_conn_out_of_seq_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Take action if out-of-seq pkt rate exceed configured threshold",
						},
						"per_conn_out_of_seq_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for out-of-seq rate exceed",
						},
						"per_conn_out_of_seq_rate_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for out-of-seq rate exceed (Default); 'blacklist-src': help Blacklist-src for out-of-seq rate exceed; 'ignore': help Ignore out-of-seq rate exceed;",
						},
					},
				},
			},
			"per_conn_pkt_rate_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per_conn_pkt_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate limit per connection per rate-interval",
						},
						"per_conn_pkt_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for per-conn-pkt-rate exceed",
						},
						"per_conn_pkt_rate_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for per-conn-pkt-rate exceed (Default); 'blacklist-src': help Blacklist-src for per-conn-pkt-rate exceed; 'ignore': Ignore per-conn-pkt-rate-exceed;",
						},
					},
				},
			},
			"per_conn_rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "1sec", Description: "'100ms': 100ms; '1sec': 1sec; '10sec': 10sec;",
			},
			"per_conn_retransmit_rate_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per_conn_retransmit_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Take action if retransmit pkt rate exceed configured threshold",
						},
						"per_conn_retransmit_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for retransmit rate exceed",
						},
						"per_conn_retransmit_rate_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for retrans rate exceed (Default); 'blacklist-src': help Blacklist-src for retrans rate exceed; 'ignore': help Ignore retrans rate exceed;",
						},
					},
				},
			},
			"per_conn_zero_win_rate_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per_conn_zero_win_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Take action if zero window pkt rate exceed configured threshold",
						},
						"per_conn_zero_win_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for zero window rate exceed",
						},
						"per_conn_zero_win_rate_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for zero-win rate exceed (Default); 'blacklist-src': help Blacklist-src for zero-win rate exceed; 'ignore': Ignore zero-win rate exceed;",
						},
					},
				},
			},
			"progression_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"progression_tracking_enabled": {
							Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable Progression Tracking Check;",
						},
						"request_response_model": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Request Response Model; 'disable': Disable Request Response Model;",
						},
						"violation": {
							Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
						},
						"ignore_tls_handshake": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TLS handshake",
						},
						"response_length_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum response length",
						},
						"response_length_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum response length",
						},
						"request_length_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum request length",
						},
						"request_length_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum request length",
						},
						"response_request_min_ratio": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum response to request ratio (in unit of 0.1% [1:1000])",
						},
						"response_request_max_ratio": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum response to request ratio (in unit of 0.1% [1:1000])",
						},
						"first_request_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum wait time from connection creation until the first data is transmitted over the connection (100 ms)",
						},
						"request_to_response_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum request to response time (100 ms)",
						},
						"response_to_request_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum response to request time (100 ms)",
						},
						"profiling_request_response_model": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for Request Response model",
						},
						"profiling_connection_life_model": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for connection model",
						},
						"profiling_time_window_model": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for time window model",
						},
						"progression_tracking_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
						},
						"progression_tracking_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"connection_tracking": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"progression_tracking_conn_enabled": {
										Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable General Progression Tracking per Connection;",
									},
									"conn_sent_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum total sent byte",
									},
									"conn_sent_min": {
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum total sent byte",
									},
									"conn_rcvd_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum total received byte",
									},
									"conn_rcvd_min": {
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum total received byte",
									},
									"conn_rcvd_sent_ratio_min": {
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum received to sent ratio (in unit of milli-, 0.001)",
									},
									"conn_rcvd_sent_ratio_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum received to sent ratio (in unit of milli-, 0.001)",
									},
									"conn_duration_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum duration time (in unit of 100ms, up to 24 hours)",
									},
									"conn_duration_min": {
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum duration time (in unit of 100ms, up to 24 hours)",
									},
									"conn_violation": {
										Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
									},
									"progression_tracking_conn_action_list_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
									},
									"progression_tracking_conn_action": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"time_window_tracking": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"progression_tracking_win_enabled": {
										Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable Progression Tracking per Time Window;",
									},
									"window_sent_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum total sent byte",
									},
									"window_sent_min": {
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum total sent byte",
									},
									"window_rcvd_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum total received byte",
									},
									"window_rcvd_min": {
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum total received byte",
									},
									"window_rcvd_sent_ratio_min": {
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum received to sent ratio (in unit of 0.1% [1:1000])",
									},
									"window_rcvd_sent_ratio_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum received to sent ratio (in unit of 0.1% [1:1000])",
									},
									"window_violation": {
										Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
									},
									"progression_tracking_windows_action_list_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
									},
									"progression_tracking_windows_action": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
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
			"retransmit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retransmit": {
							Type: schema.TypeInt, Optional: true, Description: "Take action if retransmit pkts exceed configured threshold",
						},
						"retransmit_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for retransmit exceed",
						},
						"retransmit_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for retrans exceed (Default); 'blacklist-src': help Blacklist-src for retrans exceed; 'ignore': help Ignore retrans exceed;",
						},
					},
				},
			},
			"src": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"syn_rate_limit": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_syn_rate_limit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"src_syn_rate_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for syn-rate exceed",
												},
												"src_syn_rate_action": {
													Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for syn-rate exceed (Default); 'blacklist-src': Blacklist-src for syn-rate exceed; 'ignore': Ignore syn-rate-exceed;",
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
			"syn_authentication": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"syn_auth_type": {
							Type: schema.TypeString, Optional: true, Description: "'send-rst': Send reset to all concurrent client auth attempts after syn cookie check pass; 'force-rst-by-ack': Send client a bad ack after syn cookie check pass; 'force-rst-by-synack': Send client a bad synack after syn cookie check pass; 'send-rst-once': Send RST to one client concurrent auth attempts;",
						},
						"syn_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "syn retransmit timeout in seconds(default timeout: 5 seconds)",
						},
						"syn_auth_min_delay": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum delay (in 100ms intervals) between SYN retransmits for retransmit-check to pass",
						},
						"syn_auth_rto": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Estimate the RTO and apply the exponential back-off for authentication",
						},
						"syn_auth_pass_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for passing the authentication",
						},
						"syn_auth_pass_action": {
							Type: schema.TypeString, Optional: true, Description: "'authenticate-src': authenticate-src (Default);",
						},
						"syn_auth_fail_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for failing the authentication.",
						},
						"syn_auth_fail_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Send reset to client (Applicable to retransmit-check only);",
						},
						"allow_ra": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow RA packets to be used for auth",
						},
					},
				},
			},
			"syn_cookie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SYN Cookie",
			},
			"synack_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Config SYNACK rate limit",
			},
			"track_together_with_syn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SYNACK will be counted in Dst Syn-rate limit",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zero_win_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zero_win": {
							Type: schema.TypeInt, Optional: true, Description: "Take action if zero window pkts exceed configured threshold",
						},
						"zero_win_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for zero window exceed",
						},
						"zero_win_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for zero-win exceed (Default); 'blacklist-src': help Blacklist-src for zero-win exceed; 'ignore': Ignore zero-win exceed;",
						},
					},
				},
			},
		},
	}
}
func resourceDdosZoneTemplateTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateTcpAckAuthentication(d []interface{}) edpt.DdosZoneTemplateTcpAckAuthentication {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpAckAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AckAuthTimeout = in["ack_auth_timeout"].(int)
		ret.AckAuthMinDelay = in["ack_auth_min_delay"].(int)
		ret.AckAuthOnly = in["ack_auth_only"].(int)
		ret.AckAuthRto = in["ack_auth_rto"].(int)
		ret.AckAuthPassActionListName = in["ack_auth_pass_action_list_name"].(string)
		ret.AckAuthPassAction = in["ack_auth_pass_action"].(string)
		ret.AckAuthFailActionListName = in["ack_auth_fail_action_list_name"].(string)
		ret.AckAuthFailAction = in["ack_auth_fail_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpDst(d []interface{}) edpt.DdosZoneTemplateTcpDst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateTcpDstRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateTcpDstRateLimit(d []interface{}) edpt.DdosZoneTemplateTcpDstRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpDstRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SynRateLimit = getObjectDdosZoneTemplateTcpDstRateLimitSynRateLimit(in["syn_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateTcpDstRateLimitSynRateLimit(d []interface{}) edpt.DdosZoneTemplateTcpDstRateLimitSynRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpDstRateLimitSynRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSynRateLimit = in["dst_syn_rate_limit"].(int)
		ret.DstSynRateAction = in["dst_syn_rate_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateTcpFilterList(d []interface{}) []edpt.DdosZoneTemplateTcpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateTcpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateTcpFilterList
		oi.TcpFilterName = in["tcp_filter_name"].(string)
		oi.TcpFilterSeq = in["tcp_filter_seq"].(int)
		oi.TcpFilterRegex = in["tcp_filter_regex"].(string)
		oi.TcpFilterInverseMatch = in["tcp_filter_inverse_match"].(int)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.TcpFilterActionListName = in["tcp_filter_action_list_name"].(string)
		oi.TcpFilterAction = in["tcp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpKnownRespSrcPortCfg(d []interface{}) edpt.DdosZoneTemplateTcpKnownRespSrcPortCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpKnownRespSrcPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KnownRespSrcPort = in["known_resp_src_port"].(int)
		ret.KnownRespSrcPortActionListName = in["known_resp_src_port_action_list_name"].(string)
		ret.KnownRespSrcPortAction = in["known_resp_src_port_action"].(string)
		ret.ExcludeSrcRespPort = in["exclude_src_resp_port"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpMaxRexmitSynPerFlowCfg(d []interface{}) edpt.DdosZoneTemplateTcpMaxRexmitSynPerFlowCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpMaxRexmitSynPerFlowCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxRexmitSynPerFlow = in["max_rexmit_syn_per_flow"].(int)
		ret.MaxRexmitSynPerFlowActionListName = in["max_rexmit_syn_per_flow_action_list_name"].(string)
		ret.MaxRexmitSynPerFlowAction = in["max_rexmit_syn_per_flow_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpOutOfSeqCfg(d []interface{}) edpt.DdosZoneTemplateTcpOutOfSeqCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpOutOfSeqCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutOfSeq = in["out_of_seq"].(int)
		ret.OutOfSeqActionListName = in["out_of_seq_action_list_name"].(string)
		ret.OutOfSeqAction = in["out_of_seq_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpPerConnOutOfSeqRateCfg(d []interface{}) edpt.DdosZoneTemplateTcpPerConnOutOfSeqRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpPerConnOutOfSeqRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PerConnOutOfSeqRateLimit = in["per_conn_out_of_seq_rate_limit"].(int)
		ret.PerConnOutOfSeqRateActionListName = in["per_conn_out_of_seq_rate_action_list_name"].(string)
		ret.PerConnOutOfSeqRateAction = in["per_conn_out_of_seq_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpPerConnPktRateCfg(d []interface{}) edpt.DdosZoneTemplateTcpPerConnPktRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpPerConnPktRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PerConnPktRateLimit = in["per_conn_pkt_rate_limit"].(int)
		ret.PerConnPktRateActionListName = in["per_conn_pkt_rate_action_list_name"].(string)
		ret.PerConnPktRateAction = in["per_conn_pkt_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpPerConnRetransmitRateCfg(d []interface{}) edpt.DdosZoneTemplateTcpPerConnRetransmitRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpPerConnRetransmitRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PerConnRetransmitRateLimit = in["per_conn_retransmit_rate_limit"].(int)
		ret.PerConnRetransmitRateActionListName = in["per_conn_retransmit_rate_action_list_name"].(string)
		ret.PerConnRetransmitRateAction = in["per_conn_retransmit_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpPerConnZeroWinRateCfg(d []interface{}) edpt.DdosZoneTemplateTcpPerConnZeroWinRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpPerConnZeroWinRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PerConnZeroWinRateLimit = in["per_conn_zero_win_rate_limit"].(int)
		ret.PerConnZeroWinRateActionListName = in["per_conn_zero_win_rate_action_list_name"].(string)
		ret.PerConnZeroWinRateAction = in["per_conn_zero_win_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpProgressionTracking320(d []interface{}) edpt.DdosZoneTemplateTcpProgressionTracking320 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpProgressionTracking320
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProgressionTrackingEnabled = in["progression_tracking_enabled"].(string)
		ret.RequestResponseModel = in["request_response_model"].(string)
		ret.Violation = in["violation"].(int)
		ret.IgnoreTlsHandshake = in["ignore_tls_handshake"].(int)
		ret.ResponseLengthMax = in["response_length_max"].(int)
		ret.ResponseLengthMin = in["response_length_min"].(int)
		ret.RequestLengthMin = in["request_length_min"].(int)
		ret.RequestLengthMax = in["request_length_max"].(int)
		ret.ResponseRequestMinRatio = in["response_request_min_ratio"].(int)
		ret.ResponseRequestMaxRatio = in["response_request_max_ratio"].(int)
		ret.FirstRequestMaxTime = in["first_request_max_time"].(int)
		ret.RequestToResponseMaxTime = in["request_to_response_max_time"].(int)
		ret.ResponseToRequestMaxTime = in["response_to_request_max_time"].(int)
		ret.ProfilingRequestResponseModel = in["profiling_request_response_model"].(int)
		ret.ProfilingConnectionLifeModel = in["profiling_connection_life_model"].(int)
		ret.ProfilingTimeWindowModel = in["profiling_time_window_model"].(int)
		ret.ProgressionTrackingActionListName = in["progression_tracking_action_list_name"].(string)
		ret.ProgressionTrackingAction = in["progression_tracking_action"].(string)
		//omit uuid
		ret.ConnectionTracking = getObjectDdosZoneTemplateTcpProgressionTrackingConnectionTracking321(in["connection_tracking"].([]interface{}))
		ret.TimeWindowTracking = getObjectDdosZoneTemplateTcpProgressionTrackingTimeWindowTracking322(in["time_window_tracking"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateTcpProgressionTrackingConnectionTracking321(d []interface{}) edpt.DdosZoneTemplateTcpProgressionTrackingConnectionTracking321 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpProgressionTrackingConnectionTracking321
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProgressionTrackingConnEnabled = in["progression_tracking_conn_enabled"].(string)
		ret.ConnSentMax = in["conn_sent_max"].(int)
		ret.ConnSentMin = in["conn_sent_min"].(int)
		ret.ConnRcvdMax = in["conn_rcvd_max"].(int)
		ret.ConnRcvdMin = in["conn_rcvd_min"].(int)
		ret.ConnRcvdSentRatioMin = in["conn_rcvd_sent_ratio_min"].(int)
		ret.ConnRcvdSentRatioMax = in["conn_rcvd_sent_ratio_max"].(int)
		ret.ConnDurationMax = in["conn_duration_max"].(int)
		ret.ConnDurationMin = in["conn_duration_min"].(int)
		ret.ConnViolation = in["conn_violation"].(int)
		ret.ProgressionTrackingConnActionListName = in["progression_tracking_conn_action_list_name"].(string)
		ret.ProgressionTrackingConnAction = in["progression_tracking_conn_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosZoneTemplateTcpProgressionTrackingTimeWindowTracking322(d []interface{}) edpt.DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking322 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking322
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProgressionTrackingWinEnabled = in["progression_tracking_win_enabled"].(string)
		ret.WindowSentMax = in["window_sent_max"].(int)
		ret.WindowSentMin = in["window_sent_min"].(int)
		ret.WindowRcvdMax = in["window_rcvd_max"].(int)
		ret.WindowRcvdMin = in["window_rcvd_min"].(int)
		ret.WindowRcvdSentRatioMin = in["window_rcvd_sent_ratio_min"].(int)
		ret.WindowRcvdSentRatioMax = in["window_rcvd_sent_ratio_max"].(int)
		ret.WindowViolation = in["window_violation"].(int)
		ret.ProgressionTrackingWindowsActionListName = in["progression_tracking_windows_action_list_name"].(string)
		ret.ProgressionTrackingWindowsAction = in["progression_tracking_windows_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosZoneTemplateTcpRetransmitCfg(d []interface{}) edpt.DdosZoneTemplateTcpRetransmitCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpRetransmitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Retransmit = in["retransmit"].(int)
		ret.RetransmitActionListName = in["retransmit_action_list_name"].(string)
		ret.RetransmitAction = in["retransmit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpSrc(d []interface{}) edpt.DdosZoneTemplateTcpSrc {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpSrc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateTcpSrcRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateTcpSrcRateLimit(d []interface{}) edpt.DdosZoneTemplateTcpSrcRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpSrcRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SynRateLimit = getObjectDdosZoneTemplateTcpSrcRateLimitSynRateLimit(in["syn_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateTcpSrcRateLimitSynRateLimit(d []interface{}) edpt.DdosZoneTemplateTcpSrcRateLimitSynRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpSrcRateLimitSynRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSynRateLimit = in["src_syn_rate_limit"].(int)
		ret.SrcSynRateActionListName = in["src_syn_rate_action_list_name"].(string)
		ret.SrcSynRateAction = in["src_syn_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpSynAuthentication(d []interface{}) edpt.DdosZoneTemplateTcpSynAuthentication {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpSynAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SynAuthType = in["syn_auth_type"].(string)
		ret.SynAuthTimeout = in["syn_auth_timeout"].(int)
		ret.SynAuthMinDelay = in["syn_auth_min_delay"].(int)
		ret.SynAuthRto = in["syn_auth_rto"].(int)
		ret.SynAuthPassActionListName = in["syn_auth_pass_action_list_name"].(string)
		ret.SynAuthPassAction = in["syn_auth_pass_action"].(string)
		ret.SynAuthFailActionListName = in["syn_auth_fail_action_list_name"].(string)
		ret.SynAuthFailAction = in["syn_auth_fail_action"].(string)
		ret.AllowRa = in["allow_ra"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateTcpZeroWinCfg(d []interface{}) edpt.DdosZoneTemplateTcpZeroWinCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpZeroWinCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZeroWin = in["zero_win"].(int)
		ret.ZeroWinActionListName = in["zero_win_action_list_name"].(string)
		ret.ZeroWinAction = in["zero_win_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateTcp(d *schema.ResourceData) edpt.DdosZoneTemplateTcp {
	var ret edpt.DdosZoneTemplateTcp
	ret.Inst.AckAuthentication = getObjectDdosZoneTemplateTcpAckAuthentication(d.Get("ack_authentication").([]interface{}))
	ret.Inst.AckAuthenticationSynackReset = d.Get("ack_authentication_synack_reset").(int)
	ret.Inst.ActionOnAckRtoRetryCount = d.Get("action_on_ack_rto_retry_count").(int)
	ret.Inst.ActionOnSynRtoRetryCount = d.Get("action_on_syn_rto_retry_count").(int)
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.AllowSynOtherflags = d.Get("allow_syn_otherflags").(int)
	ret.Inst.AllowSynackSkipAuthentications = d.Get("allow_synack_skip_authentications").(int)
	ret.Inst.AllowTcpTfo = d.Get("allow_tcp_tfo").(int)
	ret.Inst.Concurrent = d.Get("concurrent").(int)
	ret.Inst.ConnRateLimitOnSynOnly = d.Get("conn_rate_limit_on_syn_only").(int)
	ret.Inst.CreateConnOnSynOnly = d.Get("create_conn_on_syn_only").(int)
	ret.Inst.Dst = getObjectDdosZoneTemplateTcpDst(d.Get("dst").([]interface{}))
	ret.Inst.FilterList = getSliceDdosZoneTemplateTcpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.FilterMatchType = d.Get("filter_match_type").(string)
	ret.Inst.KnownRespSrcPortCfg = getObjectDdosZoneTemplateTcpKnownRespSrcPortCfg(d.Get("known_resp_src_port_cfg").([]interface{}))
	ret.Inst.MaxRexmitSynPerFlowCfg = getObjectDdosZoneTemplateTcpMaxRexmitSynPerFlowCfg(d.Get("max_rexmit_syn_per_flow_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OutOfSeqCfg = getObjectDdosZoneTemplateTcpOutOfSeqCfg(d.Get("out_of_seq_cfg").([]interface{}))
	ret.Inst.PerConnOutOfSeqRateCfg = getObjectDdosZoneTemplateTcpPerConnOutOfSeqRateCfg(d.Get("per_conn_out_of_seq_rate_cfg").([]interface{}))
	ret.Inst.PerConnPktRateCfg = getObjectDdosZoneTemplateTcpPerConnPktRateCfg(d.Get("per_conn_pkt_rate_cfg").([]interface{}))
	ret.Inst.PerConnRateInterval = d.Get("per_conn_rate_interval").(string)
	ret.Inst.PerConnRetransmitRateCfg = getObjectDdosZoneTemplateTcpPerConnRetransmitRateCfg(d.Get("per_conn_retransmit_rate_cfg").([]interface{}))
	ret.Inst.PerConnZeroWinRateCfg = getObjectDdosZoneTemplateTcpPerConnZeroWinRateCfg(d.Get("per_conn_zero_win_rate_cfg").([]interface{}))
	ret.Inst.ProgressionTracking = getObjectDdosZoneTemplateTcpProgressionTracking320(d.Get("progression_tracking").([]interface{}))
	ret.Inst.RetransmitCfg = getObjectDdosZoneTemplateTcpRetransmitCfg(d.Get("retransmit_cfg").([]interface{}))
	ret.Inst.Src = getObjectDdosZoneTemplateTcpSrc(d.Get("src").([]interface{}))
	ret.Inst.SynAuthentication = getObjectDdosZoneTemplateTcpSynAuthentication(d.Get("syn_authentication").([]interface{}))
	ret.Inst.SynCookie = d.Get("syn_cookie").(int)
	ret.Inst.SynackRateLimit = d.Get("synack_rate_limit").(int)
	ret.Inst.TrackTogetherWithSyn = d.Get("track_together_with_syn").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZeroWinCfg = getObjectDdosZoneTemplateTcpZeroWinCfg(d.Get("zero_win_cfg").([]interface{}))
	return ret
}
