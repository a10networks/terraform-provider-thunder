package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryReputationScope() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_category_reputation_scope`: Configure the scope of reputation score\n\n__PLACEHOLDER__",
		CreateContext: resourceWebCategoryReputationScopeCreate,
		UpdateContext: resourceWebCategoryReputationScopeUpdate,
		ReadContext:   resourceWebCategoryReputationScopeRead,
		DeleteContext: resourceWebCategoryReputationScopeDelete,

		Schema: map[string]*schema.Schema{
			"greater_than": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"greater_trustworthy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 81",
						},
						"greater_low_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 61",
						},
						"greater_moderate_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 41",
						},
						"greater_suspicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 21",
						},
						"greater_malicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 1",
						},
						"greater_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Reputation score is greater than or equal to the customized score (1-100)",
						},
					},
				},
			},
			"less_than": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"less_trustworthy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 100",
						},
						"less_low_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 80",
						},
						"less_moderate_risk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 60",
						},
						"less_suspicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 40",
						},
						"less_malicious": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 20",
						},
						"less_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Reputation score is less than or equal to a customized value (1-100)",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Reputation Scope name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'trustworthy': Trustworthy level(81-100); 'low-risk': Low-risk level(61-80); 'moderate-risk': Moderate-risk level(41-60); 'suspicious': Suspicious level(21-40); 'malicious': Malicious level(1-20);",
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
func resourceWebCategoryReputationScopeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryReputationScopeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryReputationScope(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryReputationScopeRead(ctx, d, meta)
	}
	return diags
}

func resourceWebCategoryReputationScopeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryReputationScopeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryReputationScope(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryReputationScopeRead(ctx, d, meta)
	}
	return diags
}
func resourceWebCategoryReputationScopeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryReputationScopeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryReputationScope(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebCategoryReputationScopeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryReputationScopeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryReputationScope(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectWebCategoryReputationScopeGreaterThan(d []interface{}) edpt.WebCategoryReputationScopeGreaterThan {

	count1 := len(d)
	var ret edpt.WebCategoryReputationScopeGreaterThan
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreaterTrustworthy = in["greater_trustworthy"].(int)
		ret.GreaterLowRisk = in["greater_low_risk"].(int)
		ret.GreaterModerateRisk = in["greater_moderate_risk"].(int)
		ret.GreaterSuspicious = in["greater_suspicious"].(int)
		ret.GreaterMalicious = in["greater_malicious"].(int)
		ret.GreaterThreshold = in["greater_threshold"].(int)
	}
	return ret
}

func getObjectWebCategoryReputationScopeLessThan(d []interface{}) edpt.WebCategoryReputationScopeLessThan {

	count1 := len(d)
	var ret edpt.WebCategoryReputationScopeLessThan
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LessTrustworthy = in["less_trustworthy"].(int)
		ret.LessLowRisk = in["less_low_risk"].(int)
		ret.LessModerateRisk = in["less_moderate_risk"].(int)
		ret.LessSuspicious = in["less_suspicious"].(int)
		ret.LessMalicious = in["less_malicious"].(int)
		ret.LessThreshold = in["less_threshold"].(int)
	}
	return ret
}

func getSliceWebCategoryReputationScopeSamplingEnable(d []interface{}) []edpt.WebCategoryReputationScopeSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.WebCategoryReputationScopeSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryReputationScopeSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryReputationScope(d *schema.ResourceData) edpt.WebCategoryReputationScope {
	var ret edpt.WebCategoryReputationScope
	ret.Inst.GreaterThan = getObjectWebCategoryReputationScopeGreaterThan(d.Get("greater_than").([]interface{}))
	ret.Inst.LessThan = getObjectWebCategoryReputationScopeLessThan(d.Get("less_than").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceWebCategoryReputationScopeSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
