package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListDefaultStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_rule_list_default_stats`: Statistics for the object default\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRuleListDefaultStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceCgnv6LsnRuleListDefaultStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDefaultStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDefaultStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRuleListDefaultStatsStats := setObjectCgnv6LsnRuleListDefaultStatsStats(res)
		d.Set("stats", Cgnv6LsnRuleListDefaultStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRuleListDefaultStatsStats(ret edpt.DataCgnv6LsnRuleListDefaultStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectCgnv6LsnRuleListDefaultStatsStats(d []interface{}) edpt.Cgnv6LsnRuleListDefaultStatsStats {

	var ret edpt.Cgnv6LsnRuleListDefaultStatsStats
	return ret
}

func dataToEndpointCgnv6LsnRuleListDefaultStats(d *schema.ResourceData) edpt.Cgnv6LsnRuleListDefaultStats {
	var ret edpt.Cgnv6LsnRuleListDefaultStats

	ret.Stats = getObjectCgnv6LsnRuleListDefaultStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
