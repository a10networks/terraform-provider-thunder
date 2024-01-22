package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatch() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_source_destination_adv_match`: Advanced match rule\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "Forwading action of this rule",
			},
			"disable_reqmod_icap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable REQMOD ICAP template",
			},
			"disable_respmod_icap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable RESPMOD ICAP template",
			},
			"dual_stack_action": {
				Type: schema.TypeString, Optional: true, Description: "Forwarding action of this rule",
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
			"match_http_url": {
				Type: schema.TypeString, Optional: true, Description: "Match URL in HTTP request line",
			},
			"match_http_url_regex": {
				Type: schema.TypeString, Optional: true, Description: "Match URI in HTTP request line by given regular expression",
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
			"notify_page": {
				Type: schema.TypeString, Optional: true, Description: "Send notify-page to client",
			},
			"priority": {
				Type: schema.TypeInt, Required: true, Description: "Rule priority (1000 is highest)",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAdvMatch(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAdvMatch(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAdvMatch(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAdvMatch(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAdvMatch(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DisableReqmodIcap = d.Get("disable_reqmod_icap").(int)
	ret.Inst.DisableRespmodIcap = d.Get("disable_respmod_icap").(int)
	ret.Inst.DualStackAction = d.Get("dual_stack_action").(string)
	ret.Inst.MatchHost = d.Get("match_host").(string)
	ret.Inst.MatchHttpContentEncoding = d.Get("match_http_content_encoding").(string)
	ret.Inst.MatchHttpContentLengthRangeBegin = d.Get("match_http_content_length_range_begin").(int)
	ret.Inst.MatchHttpContentLengthRangeEnd = d.Get("match_http_content_length_range_end").(int)
	ret.Inst.MatchHttpContentType = d.Get("match_http_content_type").(string)
	ret.Inst.MatchHttpHeader = d.Get("match_http_header").(string)
	ret.Inst.MatchHttpMethodConnect = d.Get("match_http_method_connect").(int)
	ret.Inst.MatchHttpMethodDelete = d.Get("match_http_method_delete").(int)
	ret.Inst.MatchHttpMethodGet = d.Get("match_http_method_get").(int)
	ret.Inst.MatchHttpMethodHead = d.Get("match_http_method_head").(int)
	ret.Inst.MatchHttpMethodOptions = d.Get("match_http_method_options").(int)
	ret.Inst.MatchHttpMethodPatch = d.Get("match_http_method_patch").(int)
	ret.Inst.MatchHttpMethodPost = d.Get("match_http_method_post").(int)
	ret.Inst.MatchHttpMethodPut = d.Get("match_http_method_put").(int)
	ret.Inst.MatchHttpMethodTrace = d.Get("match_http_method_trace").(int)
	ret.Inst.MatchHttpRequestFileExtension = d.Get("match_http_request_file_extension").(string)
	ret.Inst.MatchHttpUrl = d.Get("match_http_url").(string)
	ret.Inst.MatchHttpUrlRegex = d.Get("match_http_url_regex").(string)
	ret.Inst.MatchHttpUserAgent = d.Get("match_http_user_agent").(string)
	ret.Inst.MatchServerAddress = d.Get("match_server_address").(string)
	ret.Inst.MatchServerPort = d.Get("match_server_port").(int)
	ret.Inst.MatchServerPortRangeBegin = d.Get("match_server_port_range_begin").(int)
	ret.Inst.MatchServerPortRangeEnd = d.Get("match_server_port_range_end").(int)
	ret.Inst.MatchTimeRange = d.Get("match_time_range").(string)
	ret.Inst.MatchWebCategoryList = d.Get("match_web_category_list").(string)
	ret.Inst.MatchWebReputationScope = d.Get("match_web_reputation_scope").(string)
	ret.Inst.NotifyPage = d.Get("notify_page").(string)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
