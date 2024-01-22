package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateHttpPolicyMultiMatchRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_http_policy_multi_match_rule`: Multi-match-rule block\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateHttpPolicyMultiMatchRuleCreate,
		UpdateContext: resourceSlbTemplateHttpPolicyMultiMatchRuleUpdate,
		ReadContext:   resourceSlbTemplateHttpPolicyMultiMatchRuleRead,
		DeleteContext: resourceSlbTemplateHttpPolicyMultiMatchRuleDelete,

		Schema: map[string]*schema.Schema{
			"cookie_name_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie value string",
			},
			"cookie_name_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': Cookie name contains string;",
			},
			"cookie_name_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie name string",
			},
			"cookie_name_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': Cookie name ends-with string;",
			},
			"cookie_name_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie name string",
			},
			"cookie_name_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': Cookie name equals to string;",
			},
			"cookie_name_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie name string",
			},
			"cookie_name_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': Cookie name starts-with string;",
			},
			"cookie_value_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie value string",
			},
			"cookie_value_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': Cookie value contains string;",
			},
			"cookie_value_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie value string",
			},
			"cookie_value_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': Cookie value ends-with string;",
			},
			"cookie_value_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie value string",
			},
			"cookie_value_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': Cookie value equals to string;",
			},
			"cookie_value_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Cookie value string",
			},
			"cookie_value_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': Cookie value starts-with string;",
			},
			"header_name_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "Header name string",
			},
			"header_name_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': Header name contains string;",
			},
			"header_name_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Header name string",
			},
			"header_name_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': Header name ends-with string;",
			},
			"header_name_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "Header name string",
			},
			"header_name_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': Header name equals to string;",
			},
			"header_name_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Header name string",
			},
			"header_name_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': Header name starts-with string;",
			},
			"header_value_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "Header value string",
			},
			"header_value_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': Header value contains string;",
			},
			"header_value_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Header value string",
			},
			"header_value_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': Header value ends-with string;",
			},
			"header_value_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "Header value string",
			},
			"header_value_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': Header value equals to string;",
			},
			"header_value_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Header value string",
			},
			"header_value_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': Header value starts-with string;",
			},
			"host_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "Host string",
			},
			"host_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': Host contains string;",
			},
			"host_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Host string",
			},
			"host_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': Host ends-with string;",
			},
			"host_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "Host string",
			},
			"host_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': Host equals to string;",
			},
			"host_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "Host string",
			},
			"host_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': Host starts-with string;",
			},
			"multi_match": {
				Type: schema.TypeString, Required: true, Description: "Specify a multi-match-rule name",
			},
			"query_param_name_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter name string",
			},
			"query_param_name_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': query parameter name contains string;",
			},
			"query_param_name_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter name string",
			},
			"query_param_name_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': query parameter name ends-with string;",
			},
			"query_param_name_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter name string, use \"[no-name]\" for empty query-param-name match",
			},
			"query_param_name_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': query parameter name equals to string;",
			},
			"query_param_name_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter name string",
			},
			"query_param_name_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': query parameter name starts-with string;",
			},
			"query_param_value_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter value string",
			},
			"query_param_value_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': query parameter value contains string;",
			},
			"query_param_value_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter value string",
			},
			"query_param_value_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': query parameter value ends-with string;",
			},
			"query_param_value_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter value string, use \"[no-value]\" for empty query-param-value match",
			},
			"query_param_value_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': query parameter value equals to string;",
			},
			"query_param_value_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "query parameter value string",
			},
			"query_param_value_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': query parameter value starts-with string;",
			},
			"seq_num": {
				Type: schema.TypeInt, Optional: true, Description: "Specify a sequence number",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Service Group to be used (Service Group Name)",
			},
			"url_contains_string": {
				Type: schema.TypeString, Optional: true, Description: "URL string",
			},
			"url_contains_type": {
				Type: schema.TypeString, Optional: true, Description: "'contains': URL contains string;",
			},
			"url_ends_with_string": {
				Type: schema.TypeString, Optional: true, Description: "URL string",
			},
			"url_ends_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'ends-with': URL ends-with string;",
			},
			"url_equals_string": {
				Type: schema.TypeString, Optional: true, Description: "URL string",
			},
			"url_equals_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': URL equals to string;",
			},
			"url_starts_with_string": {
				Type: schema.TypeString, Optional: true, Description: "URL string",
			},
			"url_starts_with_type": {
				Type: schema.TypeString, Optional: true, Description: "'starts-with': URL starts-with string;",
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
func resourceSlbTemplateHttpPolicyMultiMatchRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyMultiMatchRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicyMultiMatchRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpPolicyMultiMatchRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateHttpPolicyMultiMatchRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyMultiMatchRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicyMultiMatchRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpPolicyMultiMatchRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateHttpPolicyMultiMatchRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyMultiMatchRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicyMultiMatchRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateHttpPolicyMultiMatchRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyMultiMatchRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicyMultiMatchRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateHttpPolicyMultiMatchRule(d *schema.ResourceData) edpt.SlbTemplateHttpPolicyMultiMatchRule {
	var ret edpt.SlbTemplateHttpPolicyMultiMatchRule
	ret.Inst.CookieNameContainsString = d.Get("cookie_name_contains_string").(string)
	ret.Inst.CookieNameContainsType = d.Get("cookie_name_contains_type").(string)
	ret.Inst.CookieNameEndsWithString = d.Get("cookie_name_ends_with_string").(string)
	ret.Inst.CookieNameEndsWithType = d.Get("cookie_name_ends_with_type").(string)
	ret.Inst.CookieNameEqualsString = d.Get("cookie_name_equals_string").(string)
	ret.Inst.CookieNameEqualsType = d.Get("cookie_name_equals_type").(string)
	ret.Inst.CookieNameStartsWithString = d.Get("cookie_name_starts_with_string").(string)
	ret.Inst.CookieNameStartsWithType = d.Get("cookie_name_starts_with_type").(string)
	ret.Inst.CookieValueContainsString = d.Get("cookie_value_contains_string").(string)
	ret.Inst.CookieValueContainsType = d.Get("cookie_value_contains_type").(string)
	ret.Inst.CookieValueEndsWithString = d.Get("cookie_value_ends_with_string").(string)
	ret.Inst.CookieValueEndsWithType = d.Get("cookie_value_ends_with_type").(string)
	ret.Inst.CookieValueEqualsString = d.Get("cookie_value_equals_string").(string)
	ret.Inst.CookieValueEqualsType = d.Get("cookie_value_equals_type").(string)
	ret.Inst.CookieValueStartsWithString = d.Get("cookie_value_starts_with_string").(string)
	ret.Inst.CookieValueStartsWithType = d.Get("cookie_value_starts_with_type").(string)
	ret.Inst.HeaderNameContainsString = d.Get("header_name_contains_string").(string)
	ret.Inst.HeaderNameContainsType = d.Get("header_name_contains_type").(string)
	ret.Inst.HeaderNameEndsWithString = d.Get("header_name_ends_with_string").(string)
	ret.Inst.HeaderNameEndsWithType = d.Get("header_name_ends_with_type").(string)
	ret.Inst.HeaderNameEqualsString = d.Get("header_name_equals_string").(string)
	ret.Inst.HeaderNameEqualsType = d.Get("header_name_equals_type").(string)
	ret.Inst.HeaderNameStartsWithString = d.Get("header_name_starts_with_string").(string)
	ret.Inst.HeaderNameStartsWithType = d.Get("header_name_starts_with_type").(string)
	ret.Inst.HeaderValueContainsString = d.Get("header_value_contains_string").(string)
	ret.Inst.HeaderValueContainsType = d.Get("header_value_contains_type").(string)
	ret.Inst.HeaderValueEndsWithString = d.Get("header_value_ends_with_string").(string)
	ret.Inst.HeaderValueEndsWithType = d.Get("header_value_ends_with_type").(string)
	ret.Inst.HeaderValueEqualsString = d.Get("header_value_equals_string").(string)
	ret.Inst.HeaderValueEqualsType = d.Get("header_value_equals_type").(string)
	ret.Inst.HeaderValueStartsWithString = d.Get("header_value_starts_with_string").(string)
	ret.Inst.HeaderValueStartsWithType = d.Get("header_value_starts_with_type").(string)
	ret.Inst.HostContainsString = d.Get("host_contains_string").(string)
	ret.Inst.HostContainsType = d.Get("host_contains_type").(string)
	ret.Inst.HostEndsWithString = d.Get("host_ends_with_string").(string)
	ret.Inst.HostEndsWithType = d.Get("host_ends_with_type").(string)
	ret.Inst.HostEqualsString = d.Get("host_equals_string").(string)
	ret.Inst.HostEqualsType = d.Get("host_equals_type").(string)
	ret.Inst.HostStartsWithString = d.Get("host_starts_with_string").(string)
	ret.Inst.HostStartsWithType = d.Get("host_starts_with_type").(string)
	ret.Inst.MultiMatch = d.Get("multi_match").(string)
	ret.Inst.QueryParamNameContainsString = d.Get("query_param_name_contains_string").(string)
	ret.Inst.QueryParamNameContainsType = d.Get("query_param_name_contains_type").(string)
	ret.Inst.QueryParamNameEndsWithString = d.Get("query_param_name_ends_with_string").(string)
	ret.Inst.QueryParamNameEndsWithType = d.Get("query_param_name_ends_with_type").(string)
	ret.Inst.QueryParamNameEqualsString = d.Get("query_param_name_equals_string").(string)
	ret.Inst.QueryParamNameEqualsType = d.Get("query_param_name_equals_type").(string)
	ret.Inst.QueryParamNameStartsWithString = d.Get("query_param_name_starts_with_string").(string)
	ret.Inst.QueryParamNameStartsWithType = d.Get("query_param_name_starts_with_type").(string)
	ret.Inst.QueryParamValueContainsString = d.Get("query_param_value_contains_string").(string)
	ret.Inst.QueryParamValueContainsType = d.Get("query_param_value_contains_type").(string)
	ret.Inst.QueryParamValueEndsWithString = d.Get("query_param_value_ends_with_string").(string)
	ret.Inst.QueryParamValueEndsWithType = d.Get("query_param_value_ends_with_type").(string)
	ret.Inst.QueryParamValueEqualsString = d.Get("query_param_value_equals_string").(string)
	ret.Inst.QueryParamValueEqualsType = d.Get("query_param_value_equals_type").(string)
	ret.Inst.QueryParamValueStartsWithString = d.Get("query_param_value_starts_with_string").(string)
	ret.Inst.QueryParamValueStartsWithType = d.Get("query_param_value_starts_with_type").(string)
	ret.Inst.SeqNum = d.Get("seq_num").(int)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.UrlContainsString = d.Get("url_contains_string").(string)
	ret.Inst.UrlContainsType = d.Get("url_contains_type").(string)
	ret.Inst.UrlEndsWithString = d.Get("url_ends_with_string").(string)
	ret.Inst.UrlEndsWithType = d.Get("url_ends_with_type").(string)
	ret.Inst.UrlEqualsString = d.Get("url_equals_string").(string)
	ret.Inst.UrlEqualsType = d.Get("url_equals_type").(string)
	ret.Inst.UrlStartsWithString = d.Get("url_starts_with_string").(string)
	ret.Inst.UrlStartsWithType = d.Get("url_starts_with_type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
