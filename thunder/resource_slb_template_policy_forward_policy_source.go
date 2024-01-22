package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySource() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_source`: proxy source list\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicySourceCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicySourceUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicySourceRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicySourceDelete,

		Schema: map[string]*schema.Schema{
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
			"match_any": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match any source",
			},
			"match_authorize_policy": {
				Type: schema.TypeString, Optional: true, Description: "Authorize-policy for user and group based policy",
			},
			"match_class_list": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "source destination match rule name",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "Priority of the source(higher the number higher the priority, default 0)",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplatePolicyForwardPolicySourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySource(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySource(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicySourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySource(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySource(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplatePolicyForwardPolicySourceDestination1450(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceDestination1450 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestination1450
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdvMatchList = getSliceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchList(in["adv_match_list"].([]interface{}))
		ret.ClassListList = getSliceSlbTemplatePolicyForwardPolicySourceDestinationClassListList(in["class_list_list"].([]interface{}))
		ret.WebReputationScopeList = getSliceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeList(in["web_reputation_scope_list"].([]interface{}))
		ret.WebCategoryListList = getSliceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListList(in["web_category_list_list"].([]interface{}))
		ret.Any = getObjectSlbTemplatePolicyForwardPolicySourceDestinationAny1451(in["any"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchList
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
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationClassListList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassListList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassListList
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

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeList
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

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListList
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

func getObjectSlbTemplatePolicyForwardPolicySourceDestinationAny1451(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceDestinationAny1451 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationAny1451
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.DualStackAction = in["dual_stack_action"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable1452(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable1452(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable1452 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable1452, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable1452
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySource(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySource {
	var ret edpt.SlbTemplatePolicyForwardPolicySource
	ret.Inst.Destination = getObjectSlbTemplatePolicyForwardPolicySourceDestination1450(d.Get("destination").([]interface{}))
	ret.Inst.MatchAny = d.Get("match_any").(int)
	ret.Inst.MatchAuthorizePolicy = d.Get("match_authorize_policy").(string)
	ret.Inst.MatchClassList = d.Get("match_class_list").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
