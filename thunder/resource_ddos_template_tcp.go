package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp`: TCP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpCreate,
		UpdateContext: resourceDdosTemplateTcpUpdate,
		ReadContext:   resourceDdosTemplateTcpRead,
		DeleteContext: resourceDdosTemplateTcpDelete,

		Schema: map[string]*schema.Schema{
			"ack_authentication_synack_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Reset client TCP SYN+ACK for authentication (DST support only)",
			},
			"action_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action_on_ack": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Monitor tcp ack for age-out session",
						},
						"reset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send RST to client",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "ACK retry timeout in sec",
						},
						"min_retry_gap": {
							Type: schema.TypeInt, Optional: true, Description: "Min gap between 2 ACKs for action-on-ack pass in 100ms interval",
						},
						"authenticate_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply action-on-ack once per source address for authentication purpose",
						},
						"rto_authentication": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Estimate the RTO and apply the exponential back-off for authentication",
						},
					},
				},
			},
			"action_on_ack_rto_retry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Take action if action-on-ack RTO-authentication fail over retry time(default:5)",
			},
			"action_on_syn_rto_retry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Take action if action-on-syn RTO-authentication fail over retry time(default:5)",
			},
			"action_syn_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action_on_syn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Monitor tcp syn for age-out session",
						},
						"action_on_syn_reset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send RST to client",
						},
						"action_on_syn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "SYN retry timeout in sec",
						},
						"action_on_syn_gap": {
							Type: schema.TypeInt, Optional: true, Description: "Min gap between 2 SYNs for action-on-syn pass in 100ms interval",
						},
						"action_on_syn_rto": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Estimate the RTO and apply the exponential back-off for authentication",
						},
					},
				},
			},
			"age": {
				Type: schema.TypeInt, Optional: true, Description: "Session age in minutes",
			},
			"allow_ra": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow RA packets to be used for auth",
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
			"black_list_out_of_seq": {
				Type: schema.TypeInt, Optional: true, Description: "Black list Src IP if out of seq pkts exceed configured threshold",
			},
			"black_list_retransmit": {
				Type: schema.TypeInt, Optional: true, Description: "Black list Src IP if retransmit pkts exceed configured threshold",
			},
			"black_list_zero_win": {
				Type: schema.TypeInt, Optional: true, Description: "Black list Src IP if zero window pkts exceed configured threshold",
			},
			"conn_rate_limit_on_syn_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only count SYN-initiated connections towards connection-rate tracking",
			},
			"create_conn_on_syn_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable connection establishment on SYN only",
			},
			"drop_known_resp_src_port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drop_known_resp_src_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop well-known if src-port is less than 1024",
						},
						"exclude_src_resp_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "excluding src port equal destination port",
						},
					},
				},
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
						"tcp_filter_seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"tcp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter Expression using Berkeley Packet Filter syntax",
						},
						"tcp_filter_unmatched": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
						},
						"tcp_filter_action": {
							Type: schema.TypeString, Optional: true, Description: "'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"per_conn_out_of_seq_rate_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for out-of-seq rate exceed (Default); 'blacklist-src': help Blacklist-src for out-of-seq rate exceed; 'ignore': help Ignore out-of-seq rate exceed;",
			},
			"per_conn_out_of_seq_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Take action if out-of-seq pkt rate exceed configured threshold",
			},
			"per_conn_pkt_rate_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for per-conn-pkt-rate exceed (Default); 'blacklist-src': help Blacklist-src for per-conn-pkt-rate exceed; 'ignore': Ignore per-conn-pkt-rate-exceed;",
			},
			"per_conn_pkt_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Packet rate limit per connection per rate-interval",
			},
			"per_conn_rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "1sec", Description: "'100ms': 100ms; '1sec': 1sec; '10sec': 10sec;",
			},
			"per_conn_retransmit_rate_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for retransmit rate exceed (Default); 'blacklist-src': help Blacklist-src for retransmit rate exceed; 'ignore': help Ignore retransmit rate exceed;",
			},
			"per_conn_retransmit_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Take action if retransmit pkt rate exceed configured threshold",
			},
			"per_conn_zero_win_rate_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for zero-win rate exceed (Default); 'blacklist-src': help Blacklist-src for zero-win rate exceed; 'ignore': help Ignore zero-win rate exceed;",
			},
			"per_conn_zero_win_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Take action if zero window pkt rate exceed configured threshold",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for request response model",
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
										Type: schema.TypeInt, Optional: true, Description: "Set the minimum received to sent ratio (in unit of 0.1% [1:1000])",
									},
									"conn_rcvd_sent_ratio_max": {
										Type: schema.TypeInt, Optional: true, Description: "Set the maximum received to sent ratio (in unit of 0.1% [1:1000])",
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
			"syn_auth": {
				Type: schema.TypeString, Optional: true, Default: "send-rst", Description: "'send-rst': Send RST to all client's concurrent auth attempts; 'force-rst-by-ack': Force client RST via the use of ACK; 'force-rst-by-synack': Force client RST via the use of bad SYN|ACK; 'disable': Disable TCP SYN Authentication; 'send-rst-once': Send RST to one client's concurrent auth attempts;",
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
			"tunnel_encap": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_encap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encapsulation using IP in IP",
									},
									"always": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"preserve_src_ipv4": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"preserve_src_ipv6": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
												},
											},
										},
									},
								},
							},
						},
						"gre_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gre_encap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encapsulation using GRE",
									},
									"gre_always": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gre_ipv4": {
													Type: schema.TypeString, Optional: true, Description: "IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"key_ipv4": {
													Type: schema.TypeString, Optional: true, Description: "Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
												},
												"preserve_src_ipv4_gre": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
												},
												"gre_ipv6": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"key_ipv6": {
													Type: schema.TypeString, Optional: true, Description: "Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
												},
												"preserve_src_ipv6_gre": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosTemplateTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateTcpActionCfg(d []interface{}) edpt.DdosTemplateTcpActionCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActionOnAck = in["action_on_ack"].(int)
		ret.Reset = in["reset"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.MinRetryGap = in["min_retry_gap"].(int)
		ret.AuthenticateOnly = in["authenticate_only"].(int)
		ret.RtoAuthentication = in["rto_authentication"].(int)
	}
	return ret
}

func getObjectDdosTemplateTcpActionSynCfg(d []interface{}) edpt.DdosTemplateTcpActionSynCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpActionSynCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActionOnSyn = in["action_on_syn"].(int)
		ret.ActionOnSynReset = in["action_on_syn_reset"].(int)
		ret.ActionOnSynTimeout = in["action_on_syn_timeout"].(int)
		ret.ActionOnSynGap = in["action_on_syn_gap"].(int)
		ret.ActionOnSynRto = in["action_on_syn_rto"].(int)
	}
	return ret
}

func getObjectDdosTemplateTcpDropKnownRespSrcPortCfg(d []interface{}) edpt.DdosTemplateTcpDropKnownRespSrcPortCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpDropKnownRespSrcPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DropKnownRespSrcPort = in["drop_known_resp_src_port"].(int)
		ret.ExcludeSrcRespPort = in["exclude_src_resp_port"].(int)
	}
	return ret
}

func getObjectDdosTemplateTcpDst(d []interface{}) edpt.DdosTemplateTcpDst {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosTemplateTcpDstRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpDstRateLimit(d []interface{}) edpt.DdosTemplateTcpDstRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpDstRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SynRateLimit = getObjectDdosTemplateTcpDstRateLimitSynRateLimit(in["syn_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpDstRateLimitSynRateLimit(d []interface{}) edpt.DdosTemplateTcpDstRateLimitSynRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpDstRateLimitSynRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSynRateLimit = in["dst_syn_rate_limit"].(int)
		ret.DstSynRateAction = in["dst_syn_rate_action"].(string)
	}
	return ret
}

func getSliceDdosTemplateTcpFilterList(d []interface{}) []edpt.DdosTemplateTcpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateTcpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateTcpFilterList
		oi.TcpFilterSeq = in["tcp_filter_seq"].(int)
		oi.TcpFilterRegex = in["tcp_filter_regex"].(string)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.TcpFilterUnmatched = in["tcp_filter_unmatched"].(int)
		oi.TcpFilterAction = in["tcp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateTcpProgressionTracking302(d []interface{}) edpt.DdosTemplateTcpProgressionTracking302 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTracking302
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProgressionTrackingEnabled = in["progression_tracking_enabled"].(string)
		ret.RequestResponseModel = in["request_response_model"].(string)
		ret.Violation = in["violation"].(int)
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
		ret.ConnectionTracking = getObjectDdosTemplateTcpProgressionTrackingConnectionTracking303(in["connection_tracking"].([]interface{}))
		ret.TimeWindowTracking = getObjectDdosTemplateTcpProgressionTrackingTimeWindowTracking304(in["time_window_tracking"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpProgressionTrackingConnectionTracking303(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingConnectionTracking303 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingConnectionTracking303
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

func getObjectDdosTemplateTcpProgressionTrackingTimeWindowTracking304(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingTimeWindowTracking304 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingTimeWindowTracking304
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

func getObjectDdosTemplateTcpSrc(d []interface{}) edpt.DdosTemplateTcpSrc {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpSrc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosTemplateTcpSrcRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpSrcRateLimit(d []interface{}) edpt.DdosTemplateTcpSrcRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpSrcRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SynRateLimit = getObjectDdosTemplateTcpSrcRateLimitSynRateLimit(in["syn_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpSrcRateLimitSynRateLimit(d []interface{}) edpt.DdosTemplateTcpSrcRateLimitSynRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpSrcRateLimitSynRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSynRateLimit = in["src_syn_rate_limit"].(int)
		ret.SrcSynRateAction = in["src_syn_rate_action"].(string)
	}
	return ret
}

func getObjectDdosTemplateTcpTunnelEncap(d []interface{}) edpt.DdosTemplateTcpTunnelEncap {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpTunnelEncap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpCfg = getObjectDdosTemplateTcpTunnelEncapIpCfg(in["ip_cfg"].([]interface{}))
		ret.GreCfg = getObjectDdosTemplateTcpTunnelEncapGreCfg(in["gre_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpTunnelEncapIpCfg(d []interface{}) edpt.DdosTemplateTcpTunnelEncapIpCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpTunnelEncapIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpEncap = in["ip_encap"].(int)
		ret.Always = getObjectDdosTemplateTcpTunnelEncapIpCfgAlways(in["always"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpTunnelEncapIpCfgAlways(d []interface{}) edpt.DdosTemplateTcpTunnelEncapIpCfgAlways {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpTunnelEncapIpCfgAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.PreserveSrcIpv4 = in["preserve_src_ipv4"].(int)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.PreserveSrcIpv6 = in["preserve_src_ipv6"].(int)
	}
	return ret
}

func getObjectDdosTemplateTcpTunnelEncapGreCfg(d []interface{}) edpt.DdosTemplateTcpTunnelEncapGreCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpTunnelEncapGreCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreEncap = in["gre_encap"].(int)
		ret.GreAlways = getObjectDdosTemplateTcpTunnelEncapGreCfgGreAlways(in["gre_always"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpTunnelEncapGreCfgGreAlways(d []interface{}) edpt.DdosTemplateTcpTunnelEncapGreCfgGreAlways {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpTunnelEncapGreCfgGreAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreIpv4 = in["gre_ipv4"].(string)
		ret.KeyIpv4 = in["key_ipv4"].(string)
		ret.PreserveSrcIpv4Gre = in["preserve_src_ipv4_gre"].(int)
		ret.GreIpv6 = in["gre_ipv6"].(string)
		ret.KeyIpv6 = in["key_ipv6"].(string)
		ret.PreserveSrcIpv6Gre = in["preserve_src_ipv6_gre"].(int)
	}
	return ret
}

func dataToEndpointDdosTemplateTcp(d *schema.ResourceData) edpt.DdosTemplateTcp {
	var ret edpt.DdosTemplateTcp
	ret.Inst.AckAuthenticationSynackReset = d.Get("ack_authentication_synack_reset").(int)
	ret.Inst.ActionCfg = getObjectDdosTemplateTcpActionCfg(d.Get("action_cfg").([]interface{}))
	ret.Inst.ActionOnAckRtoRetryCount = d.Get("action_on_ack_rto_retry_count").(int)
	ret.Inst.ActionOnSynRtoRetryCount = d.Get("action_on_syn_rto_retry_count").(int)
	ret.Inst.ActionSynCfg = getObjectDdosTemplateTcpActionSynCfg(d.Get("action_syn_cfg").([]interface{}))
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.AllowRa = d.Get("allow_ra").(int)
	ret.Inst.AllowSynOtherflags = d.Get("allow_syn_otherflags").(int)
	ret.Inst.AllowSynackSkipAuthentications = d.Get("allow_synack_skip_authentications").(int)
	ret.Inst.AllowTcpTfo = d.Get("allow_tcp_tfo").(int)
	ret.Inst.BlackListOutOfSeq = d.Get("black_list_out_of_seq").(int)
	ret.Inst.BlackListRetransmit = d.Get("black_list_retransmit").(int)
	ret.Inst.BlackListZeroWin = d.Get("black_list_zero_win").(int)
	ret.Inst.ConnRateLimitOnSynOnly = d.Get("conn_rate_limit_on_syn_only").(int)
	ret.Inst.CreateConnOnSynOnly = d.Get("create_conn_on_syn_only").(int)
	ret.Inst.DropKnownRespSrcPortCfg = getObjectDdosTemplateTcpDropKnownRespSrcPortCfg(d.Get("drop_known_resp_src_port_cfg").([]interface{}))
	ret.Inst.Dst = getObjectDdosTemplateTcpDst(d.Get("dst").([]interface{}))
	ret.Inst.FilterList = getSliceDdosTemplateTcpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PerConnOutOfSeqRateAction = d.Get("per_conn_out_of_seq_rate_action").(string)
	ret.Inst.PerConnOutOfSeqRateLimit = d.Get("per_conn_out_of_seq_rate_limit").(int)
	ret.Inst.PerConnPktRateAction = d.Get("per_conn_pkt_rate_action").(string)
	ret.Inst.PerConnPktRateLimit = d.Get("per_conn_pkt_rate_limit").(int)
	ret.Inst.PerConnRateInterval = d.Get("per_conn_rate_interval").(string)
	ret.Inst.PerConnRetransmitRateAction = d.Get("per_conn_retransmit_rate_action").(string)
	ret.Inst.PerConnRetransmitRateLimit = d.Get("per_conn_retransmit_rate_limit").(int)
	ret.Inst.PerConnZeroWinRateAction = d.Get("per_conn_zero_win_rate_action").(string)
	ret.Inst.PerConnZeroWinRateLimit = d.Get("per_conn_zero_win_rate_limit").(int)
	ret.Inst.ProgressionTracking = getObjectDdosTemplateTcpProgressionTracking302(d.Get("progression_tracking").([]interface{}))
	ret.Inst.Src = getObjectDdosTemplateTcpSrc(d.Get("src").([]interface{}))
	ret.Inst.SynAuth = d.Get("syn_auth").(string)
	ret.Inst.SynCookie = d.Get("syn_cookie").(int)
	ret.Inst.SynackRateLimit = d.Get("synack_rate_limit").(int)
	ret.Inst.TrackTogetherWithSyn = d.Get("track_together_with_syn").(int)
	ret.Inst.TunnelEncap = getObjectDdosTemplateTcpTunnelEncap(d.Get("tunnel_encap").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
