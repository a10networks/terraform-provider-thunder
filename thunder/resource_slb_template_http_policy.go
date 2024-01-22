package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateHttpPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_http_policy`: http-policy template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateHttpPolicyCreate,
		UpdateContext: resourceSlbTemplateHttpPolicyUpdate,
		ReadContext:   resourceSlbTemplateHttpPolicyRead,
		DeleteContext: resourceSlbTemplateHttpPolicyDelete,

		Schema: map[string]*schema.Schema{
			"cookie_name": {
				Type: schema.TypeString, Optional: true, Description: "name of cookie to match (Cookie Name)",
			},
			"geo_location_match": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location": {
							Type: schema.TypeString, Optional: true, Description: "Geolocation name",
						},
						"geo_location_service_group": {
							Type: schema.TypeString, Optional: true, Description: "Service Group to be used (Service Group Name)",
						},
					},
				},
			},
			"http_policy_match": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'cookie': cookie value match; 'host': hostname match; 'url': URL match; 'header-name': header name match; 'header-value': header value match; 'query-param-name': query parameter name; 'query-param-value': query parameter value;",
						},
						"match_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Select service group if URL string contains another string; 'ends-with': Select service group if URL string ends with another string; 'equals': Select service group if URL string equals another string; 'starts-with': Select service group if URL string starts with another string;",
						},
						"match_string": {
							Type: schema.TypeString, Optional: true, Description: "URL String, use \"[no-name]\" for empty query-param-name match, use \"[no-value]\" for empty query-param-value match",
						},
						"service_group": {
							Type: schema.TypeString, Optional: true, Description: "Service Group to be used (Service Group Name)",
						},
					},
				},
			},
			"multi_match_rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_match": {
							Type: schema.TypeString, Required: true, Description: "Specify a multi-match-rule name",
						},
						"seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Specify a sequence number",
						},
						"host_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': Host equals to string;",
						},
						"host_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "Host string",
						},
						"host_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Host contains string;",
						},
						"host_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "Host string",
						},
						"host_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': Host starts-with string;",
						},
						"host_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Host string",
						},
						"host_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': Host ends-with string;",
						},
						"host_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Host string",
						},
						"cookie_name_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': Cookie name equals to string;",
						},
						"cookie_name_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie name string",
						},
						"cookie_name_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Cookie name contains string;",
						},
						"cookie_name_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie value string",
						},
						"cookie_name_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': Cookie name starts-with string;",
						},
						"cookie_name_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie name string",
						},
						"cookie_name_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': Cookie name ends-with string;",
						},
						"cookie_name_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie name string",
						},
						"cookie_value_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': Cookie value equals to string;",
						},
						"cookie_value_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie value string",
						},
						"cookie_value_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Cookie value contains string;",
						},
						"cookie_value_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie value string",
						},
						"cookie_value_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': Cookie value starts-with string;",
						},
						"cookie_value_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie value string",
						},
						"cookie_value_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': Cookie value ends-with string;",
						},
						"cookie_value_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Cookie value string",
						},
						"url_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': URL equals to string;",
						},
						"url_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "URL string",
						},
						"url_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': URL contains string;",
						},
						"url_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "URL string",
						},
						"url_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': URL starts-with string;",
						},
						"url_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "URL string",
						},
						"url_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': URL ends-with string;",
						},
						"url_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "URL string",
						},
						"header_name_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': Header name equals to string;",
						},
						"header_name_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "Header name string",
						},
						"header_name_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Header name contains string;",
						},
						"header_name_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "Header name string",
						},
						"header_name_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': Header name starts-with string;",
						},
						"header_name_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Header name string",
						},
						"header_name_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': Header name ends-with string;",
						},
						"header_name_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Header name string",
						},
						"header_value_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': Header value equals to string;",
						},
						"header_value_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "Header value string",
						},
						"header_value_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Header value contains string;",
						},
						"header_value_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "Header value string",
						},
						"header_value_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': Header value starts-with string;",
						},
						"header_value_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Header value string",
						},
						"header_value_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': Header value ends-with string;",
						},
						"header_value_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "Header value string",
						},
						"query_param_name_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': query parameter name equals to string;",
						},
						"query_param_name_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter name string, use \"[no-name]\" for empty query-param-name match",
						},
						"query_param_name_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': query parameter name contains string;",
						},
						"query_param_name_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter name string",
						},
						"query_param_name_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': query parameter name starts-with string;",
						},
						"query_param_name_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter name string",
						},
						"query_param_name_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': query parameter name ends-with string;",
						},
						"query_param_name_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter name string",
						},
						"query_param_value_equals_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': query parameter value equals to string;",
						},
						"query_param_value_equals_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter value string, use \"[no-value]\" for empty query-param-value match",
						},
						"query_param_value_contains_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': query parameter value contains string;",
						},
						"query_param_value_contains_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter value string",
						},
						"query_param_value_starts_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'starts-with': query parameter value starts-with string;",
						},
						"query_param_value_starts_with_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter value string",
						},
						"query_param_value_ends_with_type": {
							Type: schema.TypeString, Optional: true, Description: "'ends-with': query parameter value ends-with string;",
						},
						"query_param_value_ends_with_string": {
							Type: schema.TypeString, Optional: true, Description: "query parameter value string",
						},
						"service_group": {
							Type: schema.TypeString, Optional: true, Description: "Service Group to be used (Service Group Name)",
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
				Type: schema.TypeString, Required: true, Description: "http-policy template name",
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
func resourceSlbTemplateHttpPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateHttpPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateHttpPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateHttpPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateHttpPolicyGeoLocationMatch(d []interface{}) []edpt.SlbTemplateHttpPolicyGeoLocationMatch {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpPolicyGeoLocationMatch, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpPolicyGeoLocationMatch
		oi.GeoLocation = in["geo_location"].(string)
		oi.GeoLocationServiceGroup = in["geo_location_service_group"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpPolicyHttpPolicyMatch(d []interface{}) []edpt.SlbTemplateHttpPolicyHttpPolicyMatch {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpPolicyHttpPolicyMatch, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpPolicyHttpPolicyMatch
		oi.Type = in["type"].(string)
		oi.MatchType = in["match_type"].(string)
		oi.MatchString = in["match_string"].(string)
		oi.ServiceGroup = in["service_group"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpPolicyMultiMatchRuleList(d []interface{}) []edpt.SlbTemplateHttpPolicyMultiMatchRuleList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpPolicyMultiMatchRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpPolicyMultiMatchRuleList
		oi.MultiMatch = in["multi_match"].(string)
		oi.SeqNum = in["seq_num"].(int)
		oi.HostEqualsType = in["host_equals_type"].(string)
		oi.HostEqualsString = in["host_equals_string"].(string)
		oi.HostContainsType = in["host_contains_type"].(string)
		oi.HostContainsString = in["host_contains_string"].(string)
		oi.HostStartsWithType = in["host_starts_with_type"].(string)
		oi.HostStartsWithString = in["host_starts_with_string"].(string)
		oi.HostEndsWithType = in["host_ends_with_type"].(string)
		oi.HostEndsWithString = in["host_ends_with_string"].(string)
		oi.CookieNameEqualsType = in["cookie_name_equals_type"].(string)
		oi.CookieNameEqualsString = in["cookie_name_equals_string"].(string)
		oi.CookieNameContainsType = in["cookie_name_contains_type"].(string)
		oi.CookieNameContainsString = in["cookie_name_contains_string"].(string)
		oi.CookieNameStartsWithType = in["cookie_name_starts_with_type"].(string)
		oi.CookieNameStartsWithString = in["cookie_name_starts_with_string"].(string)
		oi.CookieNameEndsWithType = in["cookie_name_ends_with_type"].(string)
		oi.CookieNameEndsWithString = in["cookie_name_ends_with_string"].(string)
		oi.CookieValueEqualsType = in["cookie_value_equals_type"].(string)
		oi.CookieValueEqualsString = in["cookie_value_equals_string"].(string)
		oi.CookieValueContainsType = in["cookie_value_contains_type"].(string)
		oi.CookieValueContainsString = in["cookie_value_contains_string"].(string)
		oi.CookieValueStartsWithType = in["cookie_value_starts_with_type"].(string)
		oi.CookieValueStartsWithString = in["cookie_value_starts_with_string"].(string)
		oi.CookieValueEndsWithType = in["cookie_value_ends_with_type"].(string)
		oi.CookieValueEndsWithString = in["cookie_value_ends_with_string"].(string)
		oi.UrlEqualsType = in["url_equals_type"].(string)
		oi.UrlEqualsString = in["url_equals_string"].(string)
		oi.UrlContainsType = in["url_contains_type"].(string)
		oi.UrlContainsString = in["url_contains_string"].(string)
		oi.UrlStartsWithType = in["url_starts_with_type"].(string)
		oi.UrlStartsWithString = in["url_starts_with_string"].(string)
		oi.UrlEndsWithType = in["url_ends_with_type"].(string)
		oi.UrlEndsWithString = in["url_ends_with_string"].(string)
		oi.HeaderNameEqualsType = in["header_name_equals_type"].(string)
		oi.HeaderNameEqualsString = in["header_name_equals_string"].(string)
		oi.HeaderNameContainsType = in["header_name_contains_type"].(string)
		oi.HeaderNameContainsString = in["header_name_contains_string"].(string)
		oi.HeaderNameStartsWithType = in["header_name_starts_with_type"].(string)
		oi.HeaderNameStartsWithString = in["header_name_starts_with_string"].(string)
		oi.HeaderNameEndsWithType = in["header_name_ends_with_type"].(string)
		oi.HeaderNameEndsWithString = in["header_name_ends_with_string"].(string)
		oi.HeaderValueEqualsType = in["header_value_equals_type"].(string)
		oi.HeaderValueEqualsString = in["header_value_equals_string"].(string)
		oi.HeaderValueContainsType = in["header_value_contains_type"].(string)
		oi.HeaderValueContainsString = in["header_value_contains_string"].(string)
		oi.HeaderValueStartsWithType = in["header_value_starts_with_type"].(string)
		oi.HeaderValueStartsWithString = in["header_value_starts_with_string"].(string)
		oi.HeaderValueEndsWithType = in["header_value_ends_with_type"].(string)
		oi.HeaderValueEndsWithString = in["header_value_ends_with_string"].(string)
		oi.QueryParamNameEqualsType = in["query_param_name_equals_type"].(string)
		oi.QueryParamNameEqualsString = in["query_param_name_equals_string"].(string)
		oi.QueryParamNameContainsType = in["query_param_name_contains_type"].(string)
		oi.QueryParamNameContainsString = in["query_param_name_contains_string"].(string)
		oi.QueryParamNameStartsWithType = in["query_param_name_starts_with_type"].(string)
		oi.QueryParamNameStartsWithString = in["query_param_name_starts_with_string"].(string)
		oi.QueryParamNameEndsWithType = in["query_param_name_ends_with_type"].(string)
		oi.QueryParamNameEndsWithString = in["query_param_name_ends_with_string"].(string)
		oi.QueryParamValueEqualsType = in["query_param_value_equals_type"].(string)
		oi.QueryParamValueEqualsString = in["query_param_value_equals_string"].(string)
		oi.QueryParamValueContainsType = in["query_param_value_contains_type"].(string)
		oi.QueryParamValueContainsString = in["query_param_value_contains_string"].(string)
		oi.QueryParamValueStartsWithType = in["query_param_value_starts_with_type"].(string)
		oi.QueryParamValueStartsWithString = in["query_param_value_starts_with_string"].(string)
		oi.QueryParamValueEndsWithType = in["query_param_value_ends_with_type"].(string)
		oi.QueryParamValueEndsWithString = in["query_param_value_ends_with_string"].(string)
		oi.ServiceGroup = in["service_group"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateHttpPolicy(d *schema.ResourceData) edpt.SlbTemplateHttpPolicy {
	var ret edpt.SlbTemplateHttpPolicy
	ret.Inst.CookieName = d.Get("cookie_name").(string)
	ret.Inst.GeoLocationMatch = getSliceSlbTemplateHttpPolicyGeoLocationMatch(d.Get("geo_location_match").([]interface{}))
	ret.Inst.HttpPolicyMatch = getSliceSlbTemplateHttpPolicyHttpPolicyMatch(d.Get("http_policy_match").([]interface{}))
	ret.Inst.MultiMatchRuleList = getSliceSlbTemplateHttpPolicyMultiMatchRuleList(d.Get("multi_match_rule_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
