package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy`: Forward Policy commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicyCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicyUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicyRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicyDelete,

		Schema: map[string]*schema.Schema{
			"acos_event_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable acos event logging",
			},
			"action_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Action policy name",
						},
						"action1": {
							Type: schema.TypeString, Optional: true, Description: "'forward-to-internet': Forward request to Internet; 'forward-to-service-group': Forward request to service group; 'forward-to-proxy': Forward request to HTTP proxy server; 'drop': Drop request;",
						},
						"fake_sg": {
							Type: schema.TypeString, Optional: true, Description: "service group to forward the packets to Internet",
						},
						"real_sg": {
							Type: schema.TypeString, Optional: true, Description: "service group to forward the packets",
						},
						"forward_snat": {
							Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group",
						},
						"forward_snat_pt_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source port translation only",
						},
						"fall_back": {
							Type: schema.TypeString, Optional: true, Description: "Fallback service group for Internet",
						},
						"fall_back_snat": {
							Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group for fallback server",
						},
						"fall_back_snat_pt_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source port translation only for fallback server",
						},
						"proxy_chaining": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable proxy chaining feature",
						},
						"proxy_chaining_bypass": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward all https packets to upstream proxy",
						},
						"support_cert_fetch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fetch server certificate by upstream proxy",
						},
						"log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable logging",
						},
						"drop_response_code": {
							Type: schema.TypeInt, Optional: true, Description: "Specify response code for drop action",
						},
						"drop_message": {
							Type: schema.TypeString, Optional: true, Description: "drop-message sent to the client as webpage(html tags are included and quotation marks are required for white spaces)",
						},
						"drop_redirect_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify URL to which client request is redirected upon being dropped",
						},
						"http_status_code": {
							Type: schema.TypeString, Optional: true, Default: "302", Description: "'301': Moved permanently; '302': Found;",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests matching this destination rule;",
									},
								},
							},
						},
					},
				},
			},
			"dual_stack_action_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Action name",
						},
						"ipv4": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 service group to forward",
						},
						"ipv4_snat": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 source NAT pool or pool group",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 service group to forward",
						},
						"ipv6_snat": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 source NAT pool or pool group",
						},
						"fall_back": {
							Type: schema.TypeString, Optional: true, Description: "Fallback service group",
						},
						"fall_back_snat": {
							Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group for fallback",
						},
						"log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable logging",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests forward by this action;",
									},
								},
							},
						},
					},
				},
			},
			"enable_adv_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable adv-match rules and deactive all the other kinds of destination rules",
			},
			"filtering": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssli_url_filtering": {
							Type: schema.TypeString, Optional: true, Description: "'bypassed-sni-disable': Disable SNI filtering for bypassed URL's(enabled by default); 'intercepted-sni-enable': Enable SNI filtering for intercepted URL's(disabled by default); 'intercepted-http-disable': Disable HTTP(host/URL) filtering for intercepted URL's(enabled by default); 'no-sni-allow': Allow connection if SNI filtering is enabled and SNI header is not present(Drop by default);",
						},
					},
				},
			},
			"forward_http_connect_to_icap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward HTTP CONNECT request to ICAP server",
			},
			"local_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local logging",
			},
			"no_client_conn_reuse": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inspects only first request of a connection",
			},
			"reqmod_icap": {
				Type: schema.TypeString, Optional: true, Description: "ICAP reqmod template (Reqmod ICAP Template Name)",
			},
			"require_web_category": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wait for web category to be resolved before taking proxy decision",
			},
			"san_filtering": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssli_url_filtering_san": {
							Type: schema.TypeString, Optional: true, Description: "'enable-san': Enable SAN filtering(disabled by default); 'bypassed-san-disable': Disable SAN filtering for bypassed URL's(enabled by default); 'intercepted-san-enable': Enable SAN filtering for intercepted URL's(disabled by default); 'no-san-allow': Allow connection if SAN filtering is enabled and SAN field is not present(Drop by default);",
						},
					},
				},
			},
			"source_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "source destination match rule name",
						},
						"match_class_list": {
							Type: schema.TypeString, Optional: true, Description: "Class List Name",
						},
						"match_any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match any source",
						},
						"match_authorize_policy": {
							Type: schema.TypeString, Optional: true, Description: "Authorize-policy for user and group based policy",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "Priority of the source(higher the number higher the priority, default 0)",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests matching this source rule; 'destination-match-not-found': Number of requests without matching destination rule; 'no-host-info': Failed to parse ip or host information from request;",
									},
								},
							},
						},
						"destination": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"adv_match_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": {
													Type: schema.TypeInt, Required: true, Description: "Rule priority (1000 is highest)",
												},
												"match_host": {
													Type: schema.TypeString, Optional: true, Description: "Match request host (HTTP stage) or SNI/SAN (SSL stage)",
												},
												"match_http_content_encoding": {
													Type: schema.TypeString, Optional: true, Description: "Match the value of HTTP header \"Content-Encoding\"",
												},
												"match_http_content_length_range_begin": {
													Type: schema.TypeInt, Optional: true, Description: "Match the value of HTTP header \"Content-Length\" with an inclusive range",
												},
												"match_http_content_length_range_end": {
													Type: schema.TypeInt, Optional: true, Description: "End of the \"Content-Length\" range",
												},
												"match_http_content_type": {
													Type: schema.TypeString, Optional: true, Description: "Match the value of HTTP header \"Content-Type\"",
												},
												"match_http_header": {
													Type: schema.TypeString, Optional: true, Description: "Matching the name of all request headers",
												},
												"match_http_method_connect": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method CONNECT",
												},
												"match_http_method_delete": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method DELETE",
												},
												"match_http_method_get": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method GET",
												},
												"match_http_method_head": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method HEAD",
												},
												"match_http_method_options": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method OPTIONS",
												},
												"match_http_method_patch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method PATCH",
												},
												"match_http_method_post": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method POST",
												},
												"match_http_method_put": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method PUT",
												},
												"match_http_method_trace": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method TRACE",
												},
												"match_http_request_file_extension": {
													Type: schema.TypeString, Optional: true, Description: "Match file extension of URL in HTTP request line",
												},
												"match_http_url_regex": {
													Type: schema.TypeString, Optional: true, Description: "Match URI in HTTP request line by given regular expression",
												},
												"match_http_url": {
													Type: schema.TypeString, Optional: true, Description: "Match URL in HTTP request line",
												},
												"match_http_user_agent": {
													Type: schema.TypeString, Optional: true, Description: "Matching the value of HTTP header \"User-Agent\"",
												},
												"match_server_address": {
													Type: schema.TypeString, Optional: true, Description: "Match target server IP address",
												},
												"match_server_port": {
													Type: schema.TypeInt, Optional: true, Description: "Match target server port number",
												},
												"match_server_port_range_begin": {
													Type: schema.TypeInt, Optional: true, Description: "Math targer server port range inclusively",
												},
												"match_server_port_range_end": {
													Type: schema.TypeInt, Optional: true, Description: "End of port range",
												},
												"match_time_range": {
													Type: schema.TypeString, Optional: true, Description: "Enable rule in this time-range",
												},
												"match_web_category_list": {
													Type: schema.TypeString, Optional: true, Description: "Match web-category list",
												},
												"match_web_reputation_scope": {
													Type: schema.TypeString, Optional: true, Description: "Match web-reputation scope",
												},
												"disable_reqmod_icap": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable REQMOD ICAP template",
												},
												"disable_respmod_icap": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable RESPMOD ICAP template",
												},
												"notify_page": {
													Type: schema.TypeString, Optional: true, Description: "Send notify-page to client",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "Forwading action of this rule",
												},
												"dual_stack_action": {
													Type: schema.TypeString, Optional: true, Description: "Forwarding action of this rule",
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
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests hit this rule;",
															},
														},
													},
												},
											},
										},
									},
									"class_list_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dest_class_list": {
													Type: schema.TypeString, Required: true, Description: "Destination Class List Name",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "Action to be performed",
												},
												"dual_stack_action": {
													Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
												},
												"type": {
													Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': Match URL; 'ip': Match destination IP address;",
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"web_reputation_scope_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"web_reputation_scope": {
													Type: schema.TypeString, Required: true, Description: "Destination Web Reputation Scope Name",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "Action to be performed",
												},
												"dual_stack_action": {
													Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
												},
												"type": {
													Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': match URL;",
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"web_category_list_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"web_category_list": {
													Type: schema.TypeString, Required: true, Description: "Destination Web Category List Name",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "Action to be performed",
												},
												"dual_stack_action": {
													Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
												},
												"type": {
													Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': match URL;",
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"any": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type: schema.TypeString, Optional: true, Description: "Action to be performed",
												},
												"dual_stack_action": {
													Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"sampling_enable": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"counters1": {
																Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests matching this destination rule;",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplatePolicyForwardPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplatePolicyForwardPolicyActionList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyActionList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyActionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyActionList
		oi.Name = in["name"].(string)
		oi.Action1 = in["action1"].(string)
		oi.FakeSg = in["fake_sg"].(string)
		oi.RealSg = in["real_sg"].(string)
		oi.ForwardSnat = in["forward_snat"].(string)
		oi.ForwardSnatPtOnly = in["forward_snat_pt_only"].(int)
		oi.FallBack = in["fall_back"].(string)
		oi.FallBackSnat = in["fall_back_snat"].(string)
		oi.FallBackSnatPtOnly = in["fall_back_snat_pt_only"].(int)
		oi.ProxyChaining = in["proxy_chaining"].(int)
		oi.ProxyChainingBypass = in["proxy_chaining_bypass"].(int)
		oi.SupportCertFetch = in["support_cert_fetch"].(int)
		oi.Log = in["log"].(int)
		oi.DropResponseCode = in["drop_response_code"].(int)
		oi.DropMessage = in["drop_message"].(string)
		oi.DropRedirectUrl = in["drop_redirect_url"].(string)
		oi.HttpStatusCode = in["http_status_code"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicyActionListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyActionListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyDualStackActionList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyDualStackActionList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyDualStackActionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyDualStackActionList
		oi.Name = in["name"].(string)
		oi.Ipv4 = in["ipv4"].(string)
		oi.Ipv4Snat = in["ipv4_snat"].(string)
		oi.Ipv6 = in["ipv6"].(string)
		oi.Ipv6Snat = in["ipv6_snat"].(string)
		oi.FallBack = in["fall_back"].(string)
		oi.FallBackSnat = in["fall_back_snat"].(string)
		oi.Log = in["log"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyFiltering(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyFiltering {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyFiltering, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyFiltering
		oi.SsliUrlFiltering = in["ssli_url_filtering"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySanFiltering(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySanFiltering {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySanFiltering, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySanFiltering
		oi.SsliUrlFilteringSan = in["ssli_url_filtering_san"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceList
		oi.Name = in["name"].(string)
		oi.MatchClassList = in["match_class_list"].(string)
		oi.MatchAny = in["match_any"].(int)
		oi.MatchAuthorizePolicy = in["match_authorize_policy"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.Destination = getObjectSlbTemplatePolicyForwardPolicySourceListDestination(in["destination"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicySourceListDestination(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceListDestination {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceListDestination
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdvMatchList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList(in["adv_match_list"].([]interface{}))
		ret.ClassListList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListList(in["class_list_list"].([]interface{}))
		ret.WebReputationScopeList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList(in["web_reputation_scope_list"].([]interface{}))
		ret.WebCategoryListList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList(in["web_category_list_list"].([]interface{}))
		ret.Any = getObjectSlbTemplatePolicyForwardPolicySourceListDestinationAny(in["any"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList
		oi.Priority = in["priority"].(int)
		oi.MatchHost = in["match_host"].(string)
		oi.MatchHttpContentEncoding = in["match_http_content_encoding"].(string)
		oi.MatchHttpContentLengthRangeBegin = in["match_http_content_length_range_begin"].(int)
		oi.MatchHttpContentLengthRangeEnd = in["match_http_content_length_range_end"].(int)
		oi.MatchHttpContentType = in["match_http_content_type"].(string)
		oi.MatchHttpHeader = in["match_http_header"].(string)
		oi.MatchHttpMethodConnect = in["match_http_method_connect"].(int)
		oi.MatchHttpMethodDelete = in["match_http_method_delete"].(int)
		oi.MatchHttpMethodGet = in["match_http_method_get"].(int)
		oi.MatchHttpMethodHead = in["match_http_method_head"].(int)
		oi.MatchHttpMethodOptions = in["match_http_method_options"].(int)
		oi.MatchHttpMethodPatch = in["match_http_method_patch"].(int)
		oi.MatchHttpMethodPost = in["match_http_method_post"].(int)
		oi.MatchHttpMethodPut = in["match_http_method_put"].(int)
		oi.MatchHttpMethodTrace = in["match_http_method_trace"].(int)
		oi.MatchHttpRequestFileExtension = in["match_http_request_file_extension"].(string)
		oi.MatchHttpUrlRegex = in["match_http_url_regex"].(string)
		oi.MatchHttpUrl = in["match_http_url"].(string)
		oi.MatchHttpUserAgent = in["match_http_user_agent"].(string)
		oi.MatchServerAddress = in["match_server_address"].(string)
		oi.MatchServerPort = in["match_server_port"].(int)
		oi.MatchServerPortRangeBegin = in["match_server_port_range_begin"].(int)
		oi.MatchServerPortRangeEnd = in["match_server_port_range_end"].(int)
		oi.MatchTimeRange = in["match_time_range"].(string)
		oi.MatchWebCategoryList = in["match_web_category_list"].(string)
		oi.MatchWebReputationScope = in["match_web_reputation_scope"].(string)
		oi.DisableReqmodIcap = in["disable_reqmod_icap"].(int)
		oi.DisableRespmodIcap = in["disable_respmod_icap"].(int)
		oi.NotifyPage = in["notify_page"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList
		oi.DestClassList = in["dest_class_list"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList
		oi.WebReputationScope = in["web_reputation_scope"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList
		oi.WebCategoryList = in["web_category_list"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicySourceListDestinationAny(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAny {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAny
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.DualStackAction = in["dual_stack_action"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicy(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicy {
	var ret edpt.SlbTemplatePolicyForwardPolicy
	ret.Inst.AcosEventLog = d.Get("acos_event_log").(int)
	ret.Inst.ActionList = getSliceSlbTemplatePolicyForwardPolicyActionList(d.Get("action_list").([]interface{}))
	ret.Inst.DualStackActionList = getSliceSlbTemplatePolicyForwardPolicyDualStackActionList(d.Get("dual_stack_action_list").([]interface{}))
	ret.Inst.EnableAdvMatch = d.Get("enable_adv_match").(int)
	ret.Inst.Filtering = getSliceSlbTemplatePolicyForwardPolicyFiltering(d.Get("filtering").([]interface{}))
	ret.Inst.ForwardHttpConnectToIcap = d.Get("forward_http_connect_to_icap").(int)
	ret.Inst.LocalLogging = d.Get("local_logging").(int)
	ret.Inst.NoClientConnReuse = d.Get("no_client_conn_reuse").(int)
	ret.Inst.ReqmodIcap = d.Get("reqmod_icap").(string)
	ret.Inst.RequireWebCategory = d.Get("require_web_category").(int)
	ret.Inst.SanFiltering = getSliceSlbTemplatePolicyForwardPolicySanFiltering(d.Get("san_filtering").([]interface{}))
	ret.Inst.SourceList = getSliceSlbTemplatePolicyForwardPolicySourceList(d.Get("source_list").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
