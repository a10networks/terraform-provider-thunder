package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorNetflowTemplateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_flow_collector_netflow_template_stats`: Statistics for the object template\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityFlowCollectorNetflowTemplateStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"templates_added_to_delq": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow templates added to the delete queue",
						},
						"templates_removed_from_delq": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow templates removed from the delete queue",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityFlowCollectorNetflowTemplateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowTemplateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowTemplateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityFlowCollectorNetflowTemplateStatsStats := setObjectVisibilityFlowCollectorNetflowTemplateStatsStats(res)
		d.Set("stats", VisibilityFlowCollectorNetflowTemplateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityFlowCollectorNetflowTemplateStatsStats(ret edpt.DataVisibilityFlowCollectorNetflowTemplateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"templates_added_to_delq":     ret.DtVisibilityFlowCollectorNetflowTemplateStats.Stats.TemplatesAddedToDelq,
			"templates_removed_from_delq": ret.DtVisibilityFlowCollectorNetflowTemplateStats.Stats.TemplatesRemovedFromDelq,
		},
	}
}

func getObjectVisibilityFlowCollectorNetflowTemplateStatsStats(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplateStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowTemplateStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplatesAddedToDelq = in["templates_added_to_delq"].(int)
		ret.TemplatesRemovedFromDelq = in["templates_removed_from_delq"].(int)
	}
	return ret
}

func dataToEndpointVisibilityFlowCollectorNetflowTemplateStats(d *schema.ResourceData) edpt.VisibilityFlowCollectorNetflowTemplateStats {
	var ret edpt.VisibilityFlowCollectorNetflowTemplateStats

	ret.Stats = getObjectVisibilityFlowCollectorNetflowTemplateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
