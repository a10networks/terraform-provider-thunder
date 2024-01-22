package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAaaPolicyAaaRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_aaa_policy_aaa_rule`: Rules of AAA policy\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAaaPolicyAaaRuleCreate,
		UpdateContext: resourceAamAaaPolicyAaaRuleUpdate,
		ReadContext:   resourceAamAaaPolicyAaaRuleRead,
		DeleteContext: resourceAamAaaPolicyAaaRuleDelete,

		Schema: map[string]*schema.Schema{
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
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'allow': Allow traffic that matches this rule; 'deny': Deny traffic that matches this rule;",
			},
			"auth_failure_bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward client request even though authentication has failed",
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
			"domain_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify domain name to bind to the AAA rule (ex: a10networks.com, www.a10networks.com)",
			},
			"domain_whitelist": {
				Type: schema.TypeString, Optional: true, Description: "Specify the AC type class-list for the domain-whitelist",
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
			"index": {
				Type: schema.TypeInt, Required: true, Description: "Specify AAA rule index",
			},
			"match_encoded_uri": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable URL decoding for URI matching",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify port number for aaa-rule, default is 0 for all port numbers",
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
func resourceAamAaaPolicyAaaRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyAaaRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicyAaaRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAaaPolicyAaaRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAaaPolicyAaaRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyAaaRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicyAaaRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAaaPolicyAaaRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAaaPolicyAaaRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyAaaRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicyAaaRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAaaPolicyAaaRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyAaaRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicyAaaRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAaaPolicyAaaRuleAccessList(d []interface{}) edpt.AamAaaPolicyAaaRuleAccessList {

	count1 := len(d)
	var ret edpt.AamAaaPolicyAaaRuleAccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleHost(d []interface{}) []edpt.AamAaaPolicyAaaRuleHost {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleHost, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleHost
		oi.HostMatchType = in["host_match_type"].(string)
		oi.HostStr = in["host_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleSamplingEnable(d []interface{}) []edpt.AamAaaPolicyAaaRuleSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleUri(d []interface{}) []edpt.AamAaaPolicyAaaRuleUri {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleUri, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleUri
		oi.MatchType = in["match_type"].(string)
		oi.UriStr = in["uri_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAaaPolicyAaaRuleUserAgent(d []interface{}) []edpt.AamAaaPolicyAaaRuleUserAgent {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyAaaRuleUserAgent, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyAaaRuleUserAgent
		oi.UserAgentMatchType = in["user_agent_match_type"].(string)
		oi.UserAgentStr = in["user_agent_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAaaPolicyAaaRule(d *schema.ResourceData) edpt.AamAaaPolicyAaaRule {
	var ret edpt.AamAaaPolicyAaaRule
	ret.Inst.AccessList = getObjectAamAaaPolicyAaaRuleAccessList(d.Get("access_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AuthFailureBypass = d.Get("auth_failure_bypass").(int)
	ret.Inst.AuthenticationTemplate = d.Get("authentication_template").(string)
	ret.Inst.AuthorizePolicy = d.Get("authorize_policy").(string)
	ret.Inst.CaptchaAuthzPolicy = d.Get("captcha_authz_policy").(string)
	ret.Inst.DomainName = d.Get("domain_name").(string)
	ret.Inst.DomainWhitelist = d.Get("domain_whitelist").(string)
	ret.Inst.Host = getSliceAamAaaPolicyAaaRuleHost(d.Get("host").([]interface{}))
	ret.Inst.Index = d.Get("index").(int)
	ret.Inst.MatchEncodedUri = d.Get("match_encoded_uri").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.SamplingEnable = getSliceAamAaaPolicyAaaRuleSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Uri = getSliceAamAaaPolicyAaaRuleUri(d.Get("uri").([]interface{}))
	ret.Inst.UserAgent = getSliceAamAaaPolicyAaaRuleUserAgent(d.Get("user_agent").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
