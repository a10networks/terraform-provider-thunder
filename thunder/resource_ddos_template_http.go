package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_http`: HTTP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateHttpCreate,
		UpdateContext: resourceDdosTemplateHttpUpdate,
		ReadContext:   resourceDdosTemplateHttpRead,
		DeleteContext: resourceDdosTemplateHttpDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for the connection; 'reset': Send RST for the connection;",
			},
			"agent_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_filter_blacklist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist the source if the user-agent matches",
						},
						"agent_equals_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agent_equals": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"agent_contains_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agent_contains": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"agent_starts_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agent_starts_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"agent_ends_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agent_ends_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"challenge_cookie_name": {
				Type: schema.TypeString, Optional: true, Default: "sto-idd", Description: "Set the cookie name used to send back to client. Default is sto-idd",
			},
			"challenge_interval": {
				Type: schema.TypeInt, Optional: true, Default: 8, Description: "Specify the challenge interval. Default is 8 seconds",
			},
			"challenge_keep_cookie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep the challenge cookie from client and forward to backend. Default is do not keep",
			},
			"challenge_method": {
				Type: schema.TypeString, Optional: true, Description: "'http-redirect': http-redirect; 'javascript': javascript;",
			},
			"challenge_redirect_code": {
				Type: schema.TypeString, Optional: true, Default: "302", Description: "'302': 302 Found; '307': 307 Temporary Redirect;",
			},
			"challenge_uri_encode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Encode the challenge phrase in uri instead of in http cookie. Default encoded in http cookie",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this template",
			},
			"disallow_connect_method": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not allow HTTP Connect method (asymmetric mode only)",
			},
			"filter_header_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_filter_header_seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"http_filter_header_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"http_filter_header_unmatched": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
						},
						"http_filter_header_blacklist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Also blacklist the source when action is taken",
						},
						"http_filter_header_whitelist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Whitelist the source after filter passes, packets are dropped until then",
						},
						"http_filter_header_count_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take no action and continue processing the next filter",
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
				Type: schema.TypeInt, Optional: true, Description: "Set the the idle timeout value in seconds for HTTP connections",
			},
			"ignore_zero_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't reset idle timer on packets with zero payload length from clients",
			},
			"malformed_http": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"malformed_http_enabled": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enabling ddos malformed http protection. Default value is disabled.",
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
							Type: schema.TypeInt, Optional: true, Default: 4294967295, Description: "Set the maximum content-length header. Default value is 4294967295 bytes",
						},
						"malformed_http_bad_chunk_mon_enabled": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enabling bad chunk monitoring. Default is disabled",
						},
					},
				},
			},
			"mss_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mss_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure DDOS detection based on mss and packet size",
						},
						"mss_percent": {
							Type: schema.TypeInt, Optional: true, Description: "Configure percentage of mss such that if a packet size is below the mss times mss-percent, packet is considered bad.",
						},
						"number_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Specify percentage of mss. Default is 0, mss-timeout is not enabled.",
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
			"post_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure rate limiting for HTTP POST request",
			},
			"referer_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ref_filter_blacklist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist the source if the referer matches",
						},
						"referer_equals_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"referer_equals": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"referer_contains_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"referer_contains": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"referer_starts_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"referer_starts_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"referer_ends_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"referer_ends_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"request_header": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"request_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_rate": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP request rate limit",
						},
						"uri": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"equal_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_equals": {
													Type: schema.TypeString, Optional: true, Description: "Request rate-limit HTTP URI matching a specified pattern",
												},
												"url_equals_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Request rate limit",
												},
											},
										},
									},
									"contains_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_contains": {
													Type: schema.TypeString, Optional: true, Description: "Request rate-limit HTTP URI containing a specified pattern",
												},
												"url_contains_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Request rate limit",
												},
											},
										},
									},
									"starts_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_starts_with": {
													Type: schema.TypeString, Optional: true, Description: "Request rate-limit HTTP URI strting with a specified pattern",
												},
												"url_starts_with_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Request rate limit",
												},
											},
										},
									},
									"ends_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_ends_with": {
													Type: schema.TypeString, Optional: true, Description: "Request rate-limit HTTP URI ending with a specified pattern",
												},
												"url_ends_with_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Request rate limit",
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
			"response_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"obj_size": {
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
								},
							},
						},
					},
				},
			},
			"slow_read_drop": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_window_size": {
							Type: schema.TypeInt, Optional: true, Description: "minimum window size",
						},
						"min_window_count": {
							Type: schema.TypeInt, Optional: true, Description: "Number of packets",
						},
					},
				},
			},
			"use_hdr_ip_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_hdr_ip_as_source": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mitigate on src ip specified by http header for example X-Forwarded-For header. Default is disabled",
						},
						"l7_hdr_name": {
							Type: schema.TypeString, Optional: true, Default: "X-Forwarded-For", Description: "Set the http header name to parse for client ip. Default is X-Forwarded-For",
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
func resourceDdosTemplateHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateHttpAgentFilter(d []interface{}) edpt.DdosTemplateHttpAgentFilter {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpAgentFilter
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AgentFilterBlacklist = in["agent_filter_blacklist"].(int)
		ret.AgentEqualsCfg = getSliceDdosTemplateHttpAgentFilterAgentEqualsCfg(in["agent_equals_cfg"].([]interface{}))
		ret.AgentContainsCfg = getSliceDdosTemplateHttpAgentFilterAgentContainsCfg(in["agent_contains_cfg"].([]interface{}))
		ret.AgentStartsCfg = getSliceDdosTemplateHttpAgentFilterAgentStartsCfg(in["agent_starts_cfg"].([]interface{}))
		ret.AgentEndsCfg = getSliceDdosTemplateHttpAgentFilterAgentEndsCfg(in["agent_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosTemplateHttpAgentFilterAgentEqualsCfg(d []interface{}) []edpt.DdosTemplateHttpAgentFilterAgentEqualsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpAgentFilterAgentEqualsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpAgentFilterAgentEqualsCfg
		oi.AgentEquals = in["agent_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpAgentFilterAgentContainsCfg(d []interface{}) []edpt.DdosTemplateHttpAgentFilterAgentContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpAgentFilterAgentContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpAgentFilterAgentContainsCfg
		oi.AgentContains = in["agent_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpAgentFilterAgentStartsCfg(d []interface{}) []edpt.DdosTemplateHttpAgentFilterAgentStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpAgentFilterAgentStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpAgentFilterAgentStartsCfg
		oi.AgentStartsWith = in["agent_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpAgentFilterAgentEndsCfg(d []interface{}) []edpt.DdosTemplateHttpAgentFilterAgentEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpAgentFilterAgentEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpAgentFilterAgentEndsCfg
		oi.AgentEndsWith = in["agent_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpFilterHeaderList(d []interface{}) []edpt.DdosTemplateHttpFilterHeaderList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpFilterHeaderList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpFilterHeaderList
		oi.HttpFilterHeaderSeq = in["http_filter_header_seq"].(int)
		oi.HttpFilterHeaderRegex = in["http_filter_header_regex"].(string)
		oi.HttpFilterHeaderUnmatched = in["http_filter_header_unmatched"].(int)
		oi.HttpFilterHeaderBlacklist = in["http_filter_header_blacklist"].(int)
		oi.HttpFilterHeaderWhitelist = in["http_filter_header_whitelist"].(int)
		oi.HttpFilterHeaderCountOnly = in["http_filter_header_count_only"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateHttpMalformedHttp(d []interface{}) edpt.DdosTemplateHttpMalformedHttp {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpMalformedHttp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MalformedHttpEnabled = in["malformed_http_enabled"].(int)
		ret.MalformedHttpMaxLineSize = in["malformed_http_max_line_size"].(int)
		ret.MalformedHttpMaxNumHeaders = in["malformed_http_max_num_headers"].(int)
		ret.MalformedHttpMaxReqLineSize = in["malformed_http_max_req_line_size"].(int)
		ret.MalformedHttpMaxHeaderNameSize = in["malformed_http_max_header_name_size"].(int)
		ret.MalformedHttpMaxContentLength = in["malformed_http_max_content_length"].(int)
		ret.MalformedHttpBadChunkMonEnabled = in["malformed_http_bad_chunk_mon_enabled"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpMssCfg(d []interface{}) edpt.DdosTemplateHttpMssCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpMssCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MssTimeout = in["mss_timeout"].(int)
		ret.MssPercent = in["mss_percent"].(int)
		ret.NumberPackets = in["number_packets"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpMultiPuThresholdDistribution(d []interface{}) edpt.DdosTemplateHttpMultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpMultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getObjectDdosTemplateHttpRefererFilter(d []interface{}) edpt.DdosTemplateHttpRefererFilter {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpRefererFilter
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RefFilterBlacklist = in["ref_filter_blacklist"].(int)
		ret.RefererEqualsCfg = getSliceDdosTemplateHttpRefererFilterRefererEqualsCfg(in["referer_equals_cfg"].([]interface{}))
		ret.RefererContainsCfg = getSliceDdosTemplateHttpRefererFilterRefererContainsCfg(in["referer_contains_cfg"].([]interface{}))
		ret.RefererStartsCfg = getSliceDdosTemplateHttpRefererFilterRefererStartsCfg(in["referer_starts_cfg"].([]interface{}))
		ret.RefererEndsCfg = getSliceDdosTemplateHttpRefererFilterRefererEndsCfg(in["referer_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosTemplateHttpRefererFilterRefererEqualsCfg(d []interface{}) []edpt.DdosTemplateHttpRefererFilterRefererEqualsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpRefererFilterRefererEqualsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpRefererFilterRefererEqualsCfg
		oi.RefererEquals = in["referer_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpRefererFilterRefererContainsCfg(d []interface{}) []edpt.DdosTemplateHttpRefererFilterRefererContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpRefererFilterRefererContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpRefererFilterRefererContainsCfg
		oi.RefererContains = in["referer_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpRefererFilterRefererStartsCfg(d []interface{}) []edpt.DdosTemplateHttpRefererFilterRefererStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpRefererFilterRefererStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpRefererFilterRefererStartsCfg
		oi.RefererStartsWith = in["referer_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpRefererFilterRefererEndsCfg(d []interface{}) []edpt.DdosTemplateHttpRefererFilterRefererEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpRefererFilterRefererEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpRefererFilterRefererEndsCfg
		oi.RefererEndsWith = in["referer_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateHttpRequestHeader(d []interface{}) edpt.DdosTemplateHttpRequestHeader {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpRequestHeader
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpRequestRateLimit(d []interface{}) edpt.DdosTemplateHttpRequestRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpRequestRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestRate = in["request_rate"].(int)
		ret.Uri = getSliceDdosTemplateHttpRequestRateLimitUri(in["uri"].([]interface{}))
	}
	return ret
}

func getSliceDdosTemplateHttpRequestRateLimitUri(d []interface{}) []edpt.DdosTemplateHttpRequestRateLimitUri {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpRequestRateLimitUri, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpRequestRateLimitUri
		oi.EqualCfg = getObjectDdosTemplateHttpRequestRateLimitUriEqualCfg(in["equal_cfg"].([]interface{}))
		oi.ContainsCfg = getObjectDdosTemplateHttpRequestRateLimitUriContainsCfg(in["contains_cfg"].([]interface{}))
		oi.StartsCfg = getObjectDdosTemplateHttpRequestRateLimitUriStartsCfg(in["starts_cfg"].([]interface{}))
		oi.EndsCfg = getObjectDdosTemplateHttpRequestRateLimitUriEndsCfg(in["ends_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateHttpRequestRateLimitUriEqualCfg(d []interface{}) edpt.DdosTemplateHttpRequestRateLimitUriEqualCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpRequestRateLimitUriEqualCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlEquals = in["url_equals"].(string)
		ret.UrlEqualsRate = in["url_equals_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpRequestRateLimitUriContainsCfg(d []interface{}) edpt.DdosTemplateHttpRequestRateLimitUriContainsCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpRequestRateLimitUriContainsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlContains = in["url_contains"].(string)
		ret.UrlContainsRate = in["url_contains_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpRequestRateLimitUriStartsCfg(d []interface{}) edpt.DdosTemplateHttpRequestRateLimitUriStartsCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpRequestRateLimitUriStartsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlStartsWith = in["url_starts_with"].(string)
		ret.UrlStartsWithRate = in["url_starts_with_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpRequestRateLimitUriEndsCfg(d []interface{}) edpt.DdosTemplateHttpRequestRateLimitUriEndsCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpRequestRateLimitUriEndsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlEndsWith = in["url_ends_with"].(string)
		ret.UrlEndsWithRate = in["url_ends_with_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpResponseRateLimit(d []interface{}) edpt.DdosTemplateHttpResponseRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpResponseRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjSize = getObjectDdosTemplateHttpResponseRateLimitObjSize(in["obj_size"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateHttpResponseRateLimitObjSize(d []interface{}) edpt.DdosTemplateHttpResponseRateLimitObjSize {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpResponseRateLimitObjSize
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LessCfg = getSliceDdosTemplateHttpResponseRateLimitObjSizeLessCfg(in["less_cfg"].([]interface{}))
		ret.GreaterCfg = getSliceDdosTemplateHttpResponseRateLimitObjSizeGreaterCfg(in["greater_cfg"].([]interface{}))
		ret.BetweenCfg = getSliceDdosTemplateHttpResponseRateLimitObjSizeBetweenCfg(in["between_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosTemplateHttpResponseRateLimitObjSizeLessCfg(d []interface{}) []edpt.DdosTemplateHttpResponseRateLimitObjSizeLessCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpResponseRateLimitObjSizeLessCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpResponseRateLimitObjSizeLessCfg
		oi.ObjLess = in["obj_less"].(int)
		oi.ObjLessRate = in["obj_less_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpResponseRateLimitObjSizeGreaterCfg(d []interface{}) []edpt.DdosTemplateHttpResponseRateLimitObjSizeGreaterCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpResponseRateLimitObjSizeGreaterCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpResponseRateLimitObjSizeGreaterCfg
		oi.ObjGreater = in["obj_greater"].(int)
		oi.ObjGreaterRate = in["obj_greater_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateHttpResponseRateLimitObjSizeBetweenCfg(d []interface{}) []edpt.DdosTemplateHttpResponseRateLimitObjSizeBetweenCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateHttpResponseRateLimitObjSizeBetweenCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateHttpResponseRateLimitObjSizeBetweenCfg
		oi.ObjBetween1 = in["obj_between1"].(int)
		oi.ObjBetween2 = in["obj_between2"].(int)
		oi.ObjBetweenRate = in["obj_between_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateHttpSlowReadDrop(d []interface{}) edpt.DdosTemplateHttpSlowReadDrop {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpSlowReadDrop
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinWindowSize = in["min_window_size"].(int)
		ret.MinWindowCount = in["min_window_count"].(int)
	}
	return ret
}

func getObjectDdosTemplateHttpUseHdrIpCfg(d []interface{}) edpt.DdosTemplateHttpUseHdrIpCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateHttpUseHdrIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UseHdrIpAsSource = in["use_hdr_ip_as_source"].(int)
		ret.L7HdrName = in["l7_hdr_name"].(string)
	}
	return ret
}

func dataToEndpointDdosTemplateHttp(d *schema.ResourceData) edpt.DdosTemplateHttp {
	var ret edpt.DdosTemplateHttp
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AgentFilter = getObjectDdosTemplateHttpAgentFilter(d.Get("agent_filter").([]interface{}))
	ret.Inst.ChallengeCookieName = d.Get("challenge_cookie_name").(string)
	ret.Inst.ChallengeInterval = d.Get("challenge_interval").(int)
	ret.Inst.ChallengeKeepCookie = d.Get("challenge_keep_cookie").(int)
	ret.Inst.ChallengeMethod = d.Get("challenge_method").(string)
	ret.Inst.ChallengeRedirectCode = d.Get("challenge_redirect_code").(string)
	ret.Inst.ChallengeUriEncode = d.Get("challenge_uri_encode").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DisallowConnectMethod = d.Get("disallow_connect_method").(int)
	ret.Inst.FilterHeaderList = getSliceDdosTemplateHttpFilterHeaderList(d.Get("filter_header_list").([]interface{}))
	ret.Inst.HttpTmplName = d.Get("http_tmpl_name").(string)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.IgnoreZeroPayload = d.Get("ignore_zero_payload").(int)
	ret.Inst.MalformedHttp = getObjectDdosTemplateHttpMalformedHttp(d.Get("malformed_http").([]interface{}))
	ret.Inst.MssCfg = getObjectDdosTemplateHttpMssCfg(d.Get("mss_cfg").([]interface{}))
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosTemplateHttpMultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.NonHttpBypass = d.Get("non_http_bypass").(int)
	ret.Inst.OutOfOrderQueueSize = d.Get("out_of_order_queue_size").(int)
	ret.Inst.OutOfOrderQueueTimeout = d.Get("out_of_order_queue_timeout").(int)
	ret.Inst.PostRateLimit = d.Get("post_rate_limit").(int)
	ret.Inst.RefererFilter = getObjectDdosTemplateHttpRefererFilter(d.Get("referer_filter").([]interface{}))
	ret.Inst.RequestHeader = getObjectDdosTemplateHttpRequestHeader(d.Get("request_header").([]interface{}))
	ret.Inst.RequestRateLimit = getObjectDdosTemplateHttpRequestRateLimit(d.Get("request_rate_limit").([]interface{}))
	ret.Inst.ResponseRateLimit = getObjectDdosTemplateHttpResponseRateLimit(d.Get("response_rate_limit").([]interface{}))
	ret.Inst.SlowReadDrop = getObjectDdosTemplateHttpSlowReadDrop(d.Get("slow_read_drop").([]interface{}))
	ret.Inst.UseHdrIpCfg = getObjectDdosTemplateHttpUseHdrIpCfg(d.Get("use_hdr_ip_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
