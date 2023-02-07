package thunder

//Thunder resource SlbTemplateHttpPolicy

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateHttpPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateHttpPolicyCreate,
		UpdateContext: resourceSlbTemplateHttpPolicyUpdate,
		ReadContext:   resourceSlbTemplateHttpPolicyRead,
		DeleteContext: resourceSlbTemplateHttpPolicyDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cookie_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"http_policy_match": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"match_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"geo_location_match": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_template_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"multi_match_rule_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_match": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"host_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_name_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cookie_value_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_name_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_value_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_name_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_equals_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_equals_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_contains_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_contains_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_starts_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_starts_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_ends_with_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_param_value_ends_with_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_waf": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateHttpPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateHttpPolicy (Inside resourceSlbTemplateHttpPolicyCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateHttpPolicy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateHttpPolicy --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateHttpPolicy(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateHttpPolicyRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateHttpPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateHttpPolicy (Inside resourceSlbTemplateHttpPolicyRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateHttpPolicy(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceSlbTemplateHttpPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateHttpPolicy   (Inside resourceSlbTemplateHttpPolicyUpdate) ")
		data := dataToSlbTemplateHttpPolicy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateHttpPolicy ")
		err := go_thunder.PutSlbTemplateHttpPolicy(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateHttpPolicyRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateHttpPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateHttpPolicyDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateHttpPolicy(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateHttpPolicy(d *schema.ResourceData) go_thunder.SlbTemplateHttpPolicy {
	var vc go_thunder.SlbTemplateHttpPolicy
	var c go_thunder.SlbTemplateHTTPPolicyInstance
	c.SlbTemplateHTTPPolicyInstanceName = d.Get("name").(string)
	c.SlbTemplateHTTPPolicyInstanceCookieName = d.Get("cookie_name").(string)

	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchCount := d.Get("http_policy_match.#").(int)
	c.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchType = make([]go_thunder.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatch, 0, SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchCount)

	for i := 0; i < SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchCount; i++ {
		var obj1 go_thunder.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatch
		prefix1 := fmt.Sprintf("http_policy_match.%d.", i)
		obj1.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchType = d.Get(prefix1 + "type").(string)
		obj1.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchMatchType = d.Get(prefix1 + "match_type").(string)
		obj1.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchMatchString = d.Get(prefix1 + "match_string").(string)
		obj1.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchServiceGroup = d.Get(prefix1 + "service_group").(string)
		obj1.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchTemplate = d.Get(prefix1 + "template").(string)
		obj1.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchTemplateName = d.Get(prefix1 + "template_name").(string)
		c.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchType = append(c.SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchType, obj1)
	}

	SlbTemplateHTTPPolicyInstanceGeoLocationMatchCount := d.Get("geo_location_match.#").(int)
	c.SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocation = make([]go_thunder.SlbTemplateHTTPPolicyInstanceGeoLocationMatch, 0, SlbTemplateHTTPPolicyInstanceGeoLocationMatchCount)

	for i := 0; i < SlbTemplateHTTPPolicyInstanceGeoLocationMatchCount; i++ {
		var obj2 go_thunder.SlbTemplateHTTPPolicyInstanceGeoLocationMatch
		prefix2 := fmt.Sprintf("geo_location_match.%d.", i)
		obj2.SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocation = d.Get(prefix2 + "geo_location").(string)
		obj2.SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocationServiceGroup = d.Get(prefix2 + "geo_location_service_group").(string)
		obj2.SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocationTemplate = d.Get(prefix2 + "geo_location_template").(string)
		obj2.SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocationTemplateName = d.Get(prefix2 + "geo_location_template_name").(string)
		c.SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocation = append(c.SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocation, obj2)
	}

	c.SlbTemplateHTTPPolicyInstanceUserTag = d.Get("user_tag").(string)

	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCount := d.Get("multi_match_rule_list.#").(int)
	c.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListMultiMatch = make([]go_thunder.SlbTemplateHTTPPolicyInstanceMultiMatchRuleList, 0, SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCount)

	for i := 0; i < SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCount; i++ {
		var obj3 go_thunder.SlbTemplateHTTPPolicyInstanceMultiMatchRuleList
		prefix3 := fmt.Sprintf("multi_match_rule_list.%d.", i)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListMultiMatch = d.Get(prefix3 + "multi_match").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListSeqNum = d.Get(prefix3 + "seq_num").(int)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEqualsType = d.Get(prefix3 + "host_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEqualsString = d.Get(prefix3 + "host_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostContainsType = d.Get(prefix3 + "host_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostContainsString = d.Get(prefix3 + "host_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostStartsWithType = d.Get(prefix3 + "host_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostStartsWithString = d.Get(prefix3 + "host_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEndsWithType = d.Get(prefix3 + "host_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEndsWithString = d.Get(prefix3 + "host_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEqualsType = d.Get(prefix3 + "cookie_name_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEqualsString = d.Get(prefix3 + "cookie_name_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameContainsType = d.Get(prefix3 + "cookie_name_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameContainsString = d.Get(prefix3 + "cookie_name_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameStartsWithType = d.Get(prefix3 + "cookie_name_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameStartsWithString = d.Get(prefix3 + "cookie_name_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEndsWithType = d.Get(prefix3 + "cookie_name_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEndsWithString = d.Get(prefix3 + "cookie_name_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEqualsType = d.Get(prefix3 + "cookie_value_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEqualsString = d.Get(prefix3 + "cookie_value_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueContainsType = d.Get(prefix3 + "cookie_value_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueContainsString = d.Get(prefix3 + "cookie_value_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueStartsWithType = d.Get(prefix3 + "cookie_value_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueStartsWithString = d.Get(prefix3 + "cookie_value_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEndsWithType = d.Get(prefix3 + "cookie_value_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEndsWithString = d.Get(prefix3 + "cookie_value_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEqualsType = d.Get(prefix3 + "url_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEqualsString = d.Get(prefix3 + "url_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLContainsType = d.Get(prefix3 + "url_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLContainsString = d.Get(prefix3 + "url_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLStartsWithType = d.Get(prefix3 + "url_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLStartsWithString = d.Get(prefix3 + "url_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEndsWithType = d.Get(prefix3 + "url_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEndsWithString = d.Get(prefix3 + "url_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEqualsType = d.Get(prefix3 + "header_name_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEqualsString = d.Get(prefix3 + "header_name_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameContainsType = d.Get(prefix3 + "header_name_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameContainsString = d.Get(prefix3 + "header_name_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameStartsWithType = d.Get(prefix3 + "header_name_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameStartsWithString = d.Get(prefix3 + "header_name_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEndsWithType = d.Get(prefix3 + "header_name_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEndsWithString = d.Get(prefix3 + "header_name_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEqualsType = d.Get(prefix3 + "header_value_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEqualsString = d.Get(prefix3 + "header_value_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueContainsType = d.Get(prefix3 + "header_value_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueContainsString = d.Get(prefix3 + "header_value_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueStartsWithType = d.Get(prefix3 + "header_value_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueStartsWithString = d.Get(prefix3 + "header_value_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEndsWithType = d.Get(prefix3 + "header_value_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEndsWithString = d.Get(prefix3 + "header_value_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEqualsType = d.Get(prefix3 + "query_param_name_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEqualsString = d.Get(prefix3 + "query_param_name_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameContainsType = d.Get(prefix3 + "query_param_name_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameContainsString = d.Get(prefix3 + "query_param_name_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameStartsWithType = d.Get(prefix3 + "query_param_name_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameStartsWithString = d.Get(prefix3 + "query_param_name_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEndsWithType = d.Get(prefix3 + "query_param_name_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEndsWithString = d.Get(prefix3 + "query_param_name_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEqualsType = d.Get(prefix3 + "query_param_value_equals_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEqualsString = d.Get(prefix3 + "query_param_value_equals_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueContainsType = d.Get(prefix3 + "query_param_value_contains_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueContainsString = d.Get(prefix3 + "query_param_value_contains_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueStartsWithType = d.Get(prefix3 + "query_param_value_starts_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueStartsWithString = d.Get(prefix3 + "query_param_value_starts_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEndsWithType = d.Get(prefix3 + "query_param_value_ends_with_type").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEndsWithString = d.Get(prefix3 + "query_param_value_ends_with_string").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListServiceGroup = d.Get(prefix3 + "service_group").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListTemplateWaf = d.Get(prefix3 + "template_waf").(string)
		obj3.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListUserTag = d.Get(prefix3 + "user_tag").(string)
		c.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListMultiMatch = append(c.SlbTemplateHTTPPolicyInstanceMultiMatchRuleListMultiMatch, obj3)
	}

	vc.SlbTemplateHTTPPolicyInstanceName = c
	return vc
}
