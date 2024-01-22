package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_source_destination_web_category_list_stats`: Statistics for the object web-category-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests matching this destination rule",
						},
					},
				},
			},
			"web_category_list": {
				Type: schema.TypeString, Required: true, Description: "Destination Web Category List Name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats := setObjectSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats.Stats.Hits,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats(d.Get("stats").([]interface{}))

	ret.WebCategoryList = d.Get("web_category_list").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
