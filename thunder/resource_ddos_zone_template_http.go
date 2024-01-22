package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_http`: HTTP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateHttpCreate,
		UpdateContext: resourceDdosZoneTemplateHttpUpdate,
		ReadContext:   resourceDdosZoneTemplateHttpRead,
		DeleteContext: resourceDdosZoneTemplateHttpDelete,

		Schema: map[string]*schema.Schema{
			"challenge": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"challenge_method": {
							Type: schema.TypeString, Optional: true, Description: "'http-redirect': http-redirect; 'javascript': javascript;",
						},
						"challenge_redirect_code": {
							Type: schema.TypeString, Optional: true, Default: "302", Description: "'302': 302 Found; '307': 307 Temporary Redirect;",
						},
						"challenge_uri_encode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Encode the challenge phrase in uri instead of in http cookie. Default encoded in http cookie",
						},
						"challenge_cookie_name": {
							Type: schema.TypeString, Optional: true, Default: "sto-idd", Description: "Set the cookie name used to send back to client. Default is sto-idd",
						},
						"challenge_keep_cookie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep the challenge cookie from client and forward to backend. Default is do not keep",
						},
						"challenge_interval": {
							Type: schema.TypeInt, Optional: true, Default: 8, Description: "Specify the challenge interval. Default is 8 seconds",
						},
						"challenge_pass_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for passing the authentication",
						},
						"challenge_pass_action": {
							Type: schema.TypeString, Optional: true, Description: "'authenticate-src': Authenticate-src (Default);",
						},
						"challenge_fail_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for failing the authentication",
						},
						"challenge_fail_action": {
							Type: schema.TypeString, Optional: true, Default: "reset", Description: "'blacklist-src': Blacklist-src; 'reset': Reset client connection(Default);",
						},
					},
				},
			},
			"client_source_ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_source_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mitigate on src ip specified by http header for example X-Forwarded-For header. Default is disabled",
						},
						"http_header_name": {
							Type: schema.TypeString, Optional: true, Default: "X-Forwarded-For", Description: "Set the http header name to parse for client ip. Default is X-Forwarded-For",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this template",
			},
			"disallow_connect_method": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not allow HTTP Connect method (asymmetric mode only)",
			},
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_post": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dst_post_rate_limit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dst_post_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"dst_post_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets(Default); 'ignore': Take no action; 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
												},
											},
										},
									},
									"http_request": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dst_request_rate": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dst_request_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"dst_request_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
												},
											},
										},
									},
									"response_size": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"less_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"obj_less": {
																Type: schema.TypeInt, Optional: true, Description: "Response size configuration",
															},
															"obj_less_rate": {
																Type: schema.TypeInt, Optional: true, Description: "Response rate limit",
															},
														},
													},
												},
												"greater_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"obj_greater": {
																Type: schema.TypeInt, Optional: true, Description: "Response size configuration",
															},
															"obj_greater_rate": {
																Type: schema.TypeInt, Optional: true, Description: "Response rate limit",
															},
														},
													},
												},
												"between_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"obj_between1": {
																Type: schema.TypeInt, Optional: true, Description: "Response size configuration",
															},
															"obj_between2": {
																Type: schema.TypeInt, Optional: true, Description: "Response size configuration",
															},
															"obj_between_rate": {
																Type: schema.TypeInt, Optional: true, Description: "Response rate limit",
															},
														},
													},
												},
												"response_size_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"response_size_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
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
						"http_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"http_filter_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"http_header_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_header_regex": {
										Type: schema.TypeString, Optional: true, Description: "Regex Expression",
									},
									"http_filter_header_inverse_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
								},
							},
						},
						"http_referer_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"referer_equals_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_referer_equals": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"referer_contains_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_referer_contains": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"referer_starts_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_referer_starts_with": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"referer_ends_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_referer_ends_with": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"http_agent_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agent_equals_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_agent_equals": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"agent_contains_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_agent_contains": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"agent_starts_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_agent_starts_with": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"agent_ends_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_agent_ends_with": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"http_uri_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uri_equal_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_uri_equals": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"uri_contains_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_uri_contains": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"uri_starts_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_uri_starts_with": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"uri_ends_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_uri_ends_with": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"dst": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Set rate limit",
									},
								},
							},
						},
						"http_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"http_filter_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src; 'reset': Reset client connection;",
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
			"http_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "DDOS HTTP Template Name",
			},
			"idle_timeout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idle_timeout_value": {
							Type: schema.TypeInt, Optional: true, Description: "Set the the idle timeout value in seconds for HTTP connections",
						},
						"ignore_zero_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't reset idle timer on packets with zero payload length from clients",
						},
						"idle_timeout_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"idle_timeout_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"malformed_http": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"malformed_http": {
							Type: schema.TypeString, Optional: true, Default: "check", Description: "'check': Configure malformed HTTP parameters;",
						},
						"malformed_http_max_line_size": {
							Type: schema.TypeInt, Optional: true, Default: 32512, Description: "Set the maximum line size. Default value is 32512",
						},
						"malformed_http_max_num_headers": {
							Type: schema.TypeInt, Optional: true, Default: 90, Description: "Set the maximum number of headers. Default value is 90",
						},
						"malformed_http_max_req_line_size": {
							Type: schema.TypeInt, Optional: true, Default: 32512, Description: "Set the maximum request line size. Default value is 32512",
						},
						"malformed_http_max_header_name_size": {
							Type: schema.TypeInt, Optional: true, Default: 64, Description: "Set the maxinum header name length. Default value is 64.",
						},
						"malformed_http_max_content_length": {
							Type: schema.TypeInt, Optional: true, Default: 4294967295, Description: "Set the maxinum content-length header. Default value is 4294967295 bytes",
						},
						"malformed_http_bad_chunk_mon_enabled": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enabling bad chunk monitoring. Default is disabled",
						},
						"malformed_http_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"malformed_http_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mss_timeout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mss_percent": {
							Type: schema.TypeInt, Optional: true, Description: "Configure percentage of mss such that if a packet size is below the mss times mss-percent, packet is considered bad.",
						},
						"number_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Specify percentage of mss. Default is 0, mss-timeout is not enabled.",
						},
						"mss_timeout_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"mss_timeout_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"multi_pu_threshold_distribution": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_pu_threshold_distribution_value": {
							Type: schema.TypeInt, Optional: true, Description: "Destination side rate limit only. Default: 0",
						},
						"multi_pu_threshold_distribution_disable": {
							Type: schema.TypeString, Optional: true, Description: "'disable': Destination side rate limit only. Default: Enable;",
						},
					},
				},
			},
			"non_http_bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass non-http traffic instead of dropping",
			},
			"out_of_order_queue_size": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set the number of packets for the out-of-order HTTP queue (asym mode only)",
			},
			"out_of_order_queue_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set the timeout value in seconds for out-of-order queue in HTTP (asym mode only)",
			},
			"request_header": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"header_timeout_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"header_timeout_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"slow_read": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_window_size": {
							Type: schema.TypeInt, Optional: true, Description: "minimum window size",
						},
						"min_window_count": {
							Type: schema.TypeInt, Optional: true, Description: "Number of packets",
						},
						"slow_read_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"slow_read_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'ignore': Take no action; 'reset': Reset client connection;",
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
									"http_post": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_post_rate_limit": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"src_post_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"src_post_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets(Default); 'ignore': Take no action; 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
												},
											},
										},
									},
									"http_request": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_request_rate": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"src_request_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"src_request_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
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
func resourceDdosZoneTemplateHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateHttpChallenge(d []interface{}) edpt.DdosZoneTemplateHttpChallenge {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpChallenge
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ChallengeMethod = in["challenge_method"].(string)
		ret.ChallengeRedirectCode = in["challenge_redirect_code"].(string)
		ret.ChallengeUriEncode = in["challenge_uri_encode"].(int)
		ret.ChallengeCookieName = in["challenge_cookie_name"].(string)
		ret.ChallengeKeepCookie = in["challenge_keep_cookie"].(int)
		ret.ChallengeInterval = in["challenge_interval"].(int)
		ret.ChallengePassActionListName = in["challenge_pass_action_list_name"].(string)
		ret.ChallengePassAction = in["challenge_pass_action"].(string)
		ret.ChallengeFailActionListName = in["challenge_fail_action_list_name"].(string)
		ret.ChallengeFailAction = in["challenge_fail_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpClientSourceIp(d []interface{}) edpt.DdosZoneTemplateHttpClientSourceIp {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpClientSourceIp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientSourceIp = in["client_source_ip"].(int)
		ret.HttpHeaderName = in["http_header_name"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpDst(d []interface{}) edpt.DdosZoneTemplateHttpDst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateHttpDstRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateHttpDstRateLimit(d []interface{}) edpt.DdosZoneTemplateHttpDstRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpDstRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpPost = getObjectDdosZoneTemplateHttpDstRateLimitHttpPost(in["http_post"].([]interface{}))
		ret.HttpRequest = getObjectDdosZoneTemplateHttpDstRateLimitHttpRequest(in["http_request"].([]interface{}))
		ret.ResponseSize = getObjectDdosZoneTemplateHttpDstRateLimitResponseSize(in["response_size"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateHttpDstRateLimitHttpPost(d []interface{}) edpt.DdosZoneTemplateHttpDstRateLimitHttpPost {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpDstRateLimitHttpPost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstPostRateLimit = in["dst_post_rate_limit"].(int)
		ret.DstPostRateLimitActionListName = in["dst_post_rate_limit_action_list_name"].(string)
		ret.DstPostRateLimitAction = in["dst_post_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpDstRateLimitHttpRequest(d []interface{}) edpt.DdosZoneTemplateHttpDstRateLimitHttpRequest {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpDstRateLimitHttpRequest
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstRequestRate = in["dst_request_rate"].(int)
		ret.DstRequestRateLimitActionListName = in["dst_request_rate_limit_action_list_name"].(string)
		ret.DstRequestRateLimitAction = in["dst_request_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpDstRateLimitResponseSize(d []interface{}) edpt.DdosZoneTemplateHttpDstRateLimitResponseSize {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpDstRateLimitResponseSize
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LessCfg = getSliceDdosZoneTemplateHttpDstRateLimitResponseSizeLessCfg(in["less_cfg"].([]interface{}))
		ret.GreaterCfg = getSliceDdosZoneTemplateHttpDstRateLimitResponseSizeGreaterCfg(in["greater_cfg"].([]interface{}))
		ret.BetweenCfg = getSliceDdosZoneTemplateHttpDstRateLimitResponseSizeBetweenCfg(in["between_cfg"].([]interface{}))
		ret.ResponseSizeActionListName = in["response_size_action_list_name"].(string)
		ret.ResponseSizeAction = in["response_size_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpDstRateLimitResponseSizeLessCfg(d []interface{}) []edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeLessCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeLessCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeLessCfg
		oi.ObjLess = in["obj_less"].(int)
		oi.ObjLessRate = in["obj_less_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpDstRateLimitResponseSizeGreaterCfg(d []interface{}) []edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeGreaterCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeGreaterCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeGreaterCfg
		oi.ObjGreater = in["obj_greater"].(int)
		oi.ObjGreaterRate = in["obj_greater_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpDstRateLimitResponseSizeBetweenCfg(d []interface{}) []edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeBetweenCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeBetweenCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpDstRateLimitResponseSizeBetweenCfg
		oi.ObjBetween1 = in["obj_between1"].(int)
		oi.ObjBetween2 = in["obj_between2"].(int)
		oi.ObjBetweenRate = in["obj_between_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterList(d []interface{}) []edpt.DdosZoneTemplateHttpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterList
		oi.HttpFilterName = in["http_filter_name"].(string)
		oi.HttpFilterSeq = in["http_filter_seq"].(int)
		oi.HttpHeaderCfg = getObjectDdosZoneTemplateHttpFilterListHttpHeaderCfg(in["http_header_cfg"].([]interface{}))
		oi.HttpRefererCfg = getObjectDdosZoneTemplateHttpFilterListHttpRefererCfg(in["http_referer_cfg"].([]interface{}))
		oi.HttpAgentCfg = getObjectDdosZoneTemplateHttpFilterListHttpAgentCfg(in["http_agent_cfg"].([]interface{}))
		oi.HttpUriCfg = getObjectDdosZoneTemplateHttpFilterListHttpUriCfg(in["http_uri_cfg"].([]interface{}))
		oi.Dst = getObjectDdosZoneTemplateHttpFilterListDst(in["dst"].([]interface{}))
		oi.HttpFilterActionListName = in["http_filter_action_list_name"].(string)
		oi.HttpFilterAction = in["http_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterListHttpHeaderCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterListHttpHeaderCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterListHttpHeaderCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpFilterHeaderRegex = in["http_filter_header_regex"].(string)
		ret.HttpFilterHeaderInverseMatch = in["http_filter_header_inverse_match"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterListHttpRefererCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterListHttpRefererCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterListHttpRefererCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RefererEqualsCfg = getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererEqualsCfg(in["referer_equals_cfg"].([]interface{}))
		ret.RefererContainsCfg = getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererContainsCfg(in["referer_contains_cfg"].([]interface{}))
		ret.RefererStartsCfg = getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererStartsCfg(in["referer_starts_cfg"].([]interface{}))
		ret.RefererEndsCfg = getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererEndsCfg(in["referer_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererEqualsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEqualsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEqualsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEqualsCfg
		oi.HttpFilterRefererEquals = in["http_filter_referer_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererContainsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererContainsCfg
		oi.HttpFilterRefererContains = in["http_filter_referer_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererStartsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererStartsCfg
		oi.HttpFilterRefererStartsWith = in["http_filter_referer_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpRefererCfgRefererEndsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEndsCfg
		oi.HttpFilterRefererEndsWith = in["http_filter_referer_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterListHttpAgentCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterListHttpAgentCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterListHttpAgentCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AgentEqualsCfg = getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentEqualsCfg(in["agent_equals_cfg"].([]interface{}))
		ret.AgentContainsCfg = getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentContainsCfg(in["agent_contains_cfg"].([]interface{}))
		ret.AgentStartsCfg = getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentStartsCfg(in["agent_starts_cfg"].([]interface{}))
		ret.AgentEndsCfg = getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentEndsCfg(in["agent_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentEqualsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEqualsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEqualsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEqualsCfg
		oi.HttpFilterAgentEquals = in["http_filter_agent_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentContainsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentContainsCfg
		oi.HttpFilterAgentContains = in["http_filter_agent_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentStartsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentStartsCfg
		oi.HttpFilterAgentStartsWith = in["http_filter_agent_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpAgentCfgAgentEndsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEndsCfg
		oi.HttpFilterAgentEndsWith = in["http_filter_agent_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterListHttpUriCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterListHttpUriCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterListHttpUriCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UriEqualCfg = getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriEqualCfg(in["uri_equal_cfg"].([]interface{}))
		ret.UriContainsCfg = getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriContainsCfg(in["uri_contains_cfg"].([]interface{}))
		ret.UriStartsCfg = getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriStartsCfg(in["uri_starts_cfg"].([]interface{}))
		ret.UriEndsCfg = getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriEndsCfg(in["uri_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriEqualCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriEqualCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriEqualCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriEqualCfg
		oi.HttpFilterUriEquals = in["http_filter_uri_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriContainsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriContainsCfg
		oi.HttpFilterUriContains = in["http_filter_uri_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriStartsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriStartsCfg
		oi.HttpFilterUriStartsWith = in["http_filter_uri_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterListHttpUriCfgUriEndsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterListHttpUriCfgUriEndsCfg
		oi.HttpFilterUriEndsWith = in["http_filter_uri_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterListDst(d []interface{}) edpt.DdosZoneTemplateHttpFilterListDst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterListDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpFilterRateLimit = in["http_filter_rate_limit"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpIdleTimeout(d []interface{}) edpt.DdosZoneTemplateHttpIdleTimeout {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpIdleTimeout
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IdleTimeoutValue = in["idle_timeout_value"].(int)
		ret.IgnoreZeroPayload = in["ignore_zero_payload"].(int)
		ret.IdleTimeoutActionListName = in["idle_timeout_action_list_name"].(string)
		ret.IdleTimeoutAction = in["idle_timeout_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpMalformedHttp308(d []interface{}) edpt.DdosZoneTemplateHttpMalformedHttp308 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpMalformedHttp308
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MalformedHttp = in["malformed_http"].(string)
		ret.MalformedHttpMaxLineSize = in["malformed_http_max_line_size"].(int)
		ret.MalformedHttpMaxNumHeaders = in["malformed_http_max_num_headers"].(int)
		ret.MalformedHttpMaxReqLineSize = in["malformed_http_max_req_line_size"].(int)
		ret.MalformedHttpMaxHeaderNameSize = in["malformed_http_max_header_name_size"].(int)
		ret.MalformedHttpMaxContentLength = in["malformed_http_max_content_length"].(int)
		ret.MalformedHttpBadChunkMonEnabled = in["malformed_http_bad_chunk_mon_enabled"].(int)
		ret.MalformedHttpActionListName = in["malformed_http_action_list_name"].(string)
		ret.MalformedHttpAction = in["malformed_http_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosZoneTemplateHttpMssTimeout(d []interface{}) edpt.DdosZoneTemplateHttpMssTimeout {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpMssTimeout
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MssPercent = in["mss_percent"].(int)
		ret.NumberPackets = in["number_packets"].(int)
		ret.MssTimeoutActionListName = in["mss_timeout_action_list_name"].(string)
		ret.MssTimeoutAction = in["mss_timeout_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpMultiPuThresholdDistribution(d []interface{}) edpt.DdosZoneTemplateHttpMultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpMultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpRequestHeader(d []interface{}) edpt.DdosZoneTemplateHttpRequestHeader {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpRequestHeader
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Timeout = in["timeout"].(int)
		ret.HeaderTimeoutActionListName = in["header_timeout_action_list_name"].(string)
		ret.HeaderTimeoutAction = in["header_timeout_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpSlowRead(d []interface{}) edpt.DdosZoneTemplateHttpSlowRead {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpSlowRead
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinWindowSize = in["min_window_size"].(int)
		ret.MinWindowCount = in["min_window_count"].(int)
		ret.SlowReadActionListName = in["slow_read_action_list_name"].(string)
		ret.SlowReadAction = in["slow_read_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpSrc(d []interface{}) edpt.DdosZoneTemplateHttpSrc {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpSrc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateHttpSrcRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateHttpSrcRateLimit(d []interface{}) edpt.DdosZoneTemplateHttpSrcRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpSrcRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpPost = getObjectDdosZoneTemplateHttpSrcRateLimitHttpPost(in["http_post"].([]interface{}))
		ret.HttpRequest = getObjectDdosZoneTemplateHttpSrcRateLimitHttpRequest(in["http_request"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateHttpSrcRateLimitHttpPost(d []interface{}) edpt.DdosZoneTemplateHttpSrcRateLimitHttpPost {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpSrcRateLimitHttpPost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPostRateLimit = in["src_post_rate_limit"].(int)
		ret.SrcPostRateLimitActionListName = in["src_post_rate_limit_action_list_name"].(string)
		ret.SrcPostRateLimitAction = in["src_post_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpSrcRateLimitHttpRequest(d []interface{}) edpt.DdosZoneTemplateHttpSrcRateLimitHttpRequest {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpSrcRateLimitHttpRequest
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcRequestRate = in["src_request_rate"].(int)
		ret.SrcRequestRateLimitActionListName = in["src_request_rate_limit_action_list_name"].(string)
		ret.SrcRequestRateLimitAction = in["src_request_rate_limit_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateHttp(d *schema.ResourceData) edpt.DdosZoneTemplateHttp {
	var ret edpt.DdosZoneTemplateHttp
	ret.Inst.Challenge = getObjectDdosZoneTemplateHttpChallenge(d.Get("challenge").([]interface{}))
	ret.Inst.ClientSourceIp = getObjectDdosZoneTemplateHttpClientSourceIp(d.Get("client_source_ip").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DisallowConnectMethod = d.Get("disallow_connect_method").(int)
	ret.Inst.Dst = getObjectDdosZoneTemplateHttpDst(d.Get("dst").([]interface{}))
	ret.Inst.FilterList = getSliceDdosZoneTemplateHttpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.HttpTmplName = d.Get("http_tmpl_name").(string)
	ret.Inst.IdleTimeout = getObjectDdosZoneTemplateHttpIdleTimeout(d.Get("idle_timeout").([]interface{}))
	ret.Inst.MalformedHttp = getObjectDdosZoneTemplateHttpMalformedHttp308(d.Get("malformed_http").([]interface{}))
	ret.Inst.MssTimeout = getObjectDdosZoneTemplateHttpMssTimeout(d.Get("mss_timeout").([]interface{}))
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosZoneTemplateHttpMultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.NonHttpBypass = d.Get("non_http_bypass").(int)
	ret.Inst.OutOfOrderQueueSize = d.Get("out_of_order_queue_size").(int)
	ret.Inst.OutOfOrderQueueTimeout = d.Get("out_of_order_queue_timeout").(int)
	ret.Inst.RequestHeader = getObjectDdosZoneTemplateHttpRequestHeader(d.Get("request_header").([]interface{}))
	ret.Inst.SlowRead = getObjectDdosZoneTemplateHttpSlowRead(d.Get("slow_read").([]interface{}))
	ret.Inst.Src = getObjectDdosZoneTemplateHttpSrc(d.Get("src").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
