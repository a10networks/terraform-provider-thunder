package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAaaPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_aaa_policy`: AAM AAA policy configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAaaPolicyCreate,
		UpdateContext: resourceAamAaaPolicyUpdate,
		ReadContext:   resourceAamAaaPolicyRead,
		DeleteContext: resourceAamAaaPolicyDelete,

		Schema: map[string]*schema.Schema{
			"aaa_rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": {
							Type: schema.TypeInt, Required: true, Description: "Specify AAA rule index",
						},
						"uri": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"match_type": {
										Type: schema.TypeString, Optional: true, Description: "'contains': Match URI if request URI contains specified URI; 'ends-with': Match URI if request URI ends with specified URI; 'equals': Match URI if request URI equals specified URI; 'starts-with': Match URI if request URI starts with specified URI;",
									},
									"uri_str": {
										Type: schema.TypeString, Optional: true, Description: "Specify URI string",
									},
								},
							},
						},
						"host": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host_match_type": {
										Type: schema.TypeString, Optional: true, Description: "'contains': Match HOST if request HTTP HOST header contains specified hostname; 'ends-with': Match HOST if request HTTP HOST header ends with specified hostname; 'equals': Match HOST if request HTTP HOST header equals specified hostname; 'starts-with': Match HOST if request HTTP HOST header starts with specified hostname;",
									},
									"host_str": {
										Type: schema.TypeString, Optional: true, Description: "Specify URI string",
									},
								},
							},
						},
						"domain_whitelist": {
							Type: schema.TypeString, Optional: true, Description: "Specify the AC type class-list for the domain-whitelist",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify port number for aaa-rule, default is 0 for all port numbers",
						},
						"match_encoded_uri": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable URL decoding for URI matching",
						},
						"access_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_id": {
										Type: schema.TypeInt, Optional: true, Description: "ACL id",
									},
									"acl_name": {
										Type: schema.TypeString, Optional: true, Description: "'ip-name': Apply an IP named access list; 'ipv6-name': Apply an IPv6 named access list;",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "Specify Named Access List",
									},
								},
							},
						},
						"domain_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify domain name to bind to the AAA rule (ex: a10networks.com, www.a10networks.com)",
						},
						"user_agent": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_agent_match_type": {
										Type: schema.TypeString, Optional: true, Description: "'contains': Match request User-Agent header if it contains specified string; 'ends-with': Match request User-Agent header if it ends with specified string; 'equals': Match request User-Agent header if it equals specified string; 'starts-with': Match request User-Agent header if it starts with specified string;",
									},
									"user_agent_str": {
										Type: schema.TypeString, Optional: true, Description: "Specify request User-Agent string",
									},
								},
							},
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'allow': Allow traffic that matches this rule; 'deny': Deny traffic that matches this rule;",
						},
						"authentication_template": {
							Type: schema.TypeString, Optional: true, Description: "Specify authentication template name to bind to the AAA rule",
						},
						"authorize_policy": {
							Type: schema.TypeString, Optional: true, Description: "Specify authorization policy to bind to the AAA rule",
						},
						"captcha_authz_policy": {
							Type: schema.TypeString, Optional: true, Description: "Specify authorization policy for CAPTCHA (Authorization policy name)",
						},
						"auth_failure_bypass": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward client request even though authentication has failed",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_count': total_count; 'hit_deny': hit_deny; 'hit_auth': hit_auth; 'hit_bypass': hit_bypass; 'failure_bypass': failure_bypass;",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify AAA policy name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'req': Request; 'req-reject': Request Rejected; 'req-auth': Request Matching Authentication Template; 'req-bypass': Request Bypassed; 'req-skip': Request Skipped; 'error': Error; 'failure-bypass': Auth Failure Bypass;",
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
func resourceAamAaaPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAaaPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAaaPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAaaPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAaaPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAaaPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAaaPolicyAaaRuleList(d []interface{}) []edpt.AamAaaPolicyAaaRuleList {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleList
		oi.Index = in["index"].(int)
		oi.Uri = getSliceAamAaaPolicyAaaRuleListUri(in["uri"].([]interface{}))
		oi.Host = getSliceAamAaaPolicyAaaRuleListHost(in["host"].([]interface{}))
		oi.DomainWhitelist = in["domain_whitelist"].(string)
		oi.Port = in["port"].(int)
		oi.MatchEncodedUri = in["match_encoded_uri"].(int)
		oi.AccessList = getObjectAamAaaPolicyAaaRuleListAccessList(in["access_list"].([]interface{}))
		oi.DomainName = in["domain_name"].(string)
		oi.UserAgent = getSliceAamAaaPolicyAaaRuleListUserAgent(in["user_agent"].([]interface{}))
		oi.Action = in["action"].(string)
		oi.AuthenticationTemplate = in["authentication_template"].(string)
		oi.AuthorizePolicy = in["authorize_policy"].(string)
		oi.CaptchaAuthzPolicy = in["captcha_authz_policy"].(string)
		oi.AuthFailureBypass = in["auth_failure_bypass"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceAamAaaPolicyAaaRuleListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleListUri(d []interface{}) []edpt.AamAaaPolicyAaaRuleListUri {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleListUri, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleListUri
		oi.MatchType = in["match_type"].(string)
		oi.UriStr = in["uri_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleListHost(d []interface{}) []edpt.AamAaaPolicyAaaRuleListHost {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleListHost, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleListHost
		oi.HostMatchType = in["host_match_type"].(string)
		oi.HostStr = in["host_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAaaPolicyAaaRuleListAccessList(d []interface{}) edpt.AamAaaPolicyAaaRuleListAccessList {

	count1 := len(d)
	var ret edpt.AamAaaPolicyAaaRuleListAccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleListUserAgent(d []interface{}) []edpt.AamAaaPolicyAaaRuleListUserAgent {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleListUserAgent, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleListUserAgent
		oi.UserAgentMatchType = in["user_agent_match_type"].(string)
		oi.UserAgentStr = in["user_agent_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleListSamplingEnable(d []interface{}) []edpt.AamAaaPolicyAaaRuleListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAaaPolicySamplingEnable(d []interface{}) []edpt.AamAaaPolicySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAaaPolicy(d *schema.ResourceData) edpt.AamAaaPolicy {
	var ret edpt.AamAaaPolicy
	ret.Inst.AaaRuleList = getSliceAamAaaPolicyAaaRuleList(d.Get("aaa_rule_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceAamAaaPolicySamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
