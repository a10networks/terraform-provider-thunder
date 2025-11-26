package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowCollectorTemplateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_netflow_collector_template_stats`: Statistics for the object template\n\n__PLACEHOLDER__",
		ReadContext: resourceNetflowCollectorTemplateStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v9_templates_created": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Templates Created",
						},
						"v9_templates_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Templates Deleted",
						},
						"v10_templates_created": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v10(IPFIX) Templates Created",
						},
						"v10_templates_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v10(IPFIX) Templates Deleted",
						},
						"template_drop_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Templates Dropped Maximum Limit Reached",
						},
						"template_drop_out_of_memory": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Templates Dropped OOM",
						},
					},
				},
			},
		},
	}
}

func resourceNetflowCollectorTemplateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCollectorTemplateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCollectorTemplateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetflowCollectorTemplateStatsStats := setObjectNetflowCollectorTemplateStatsStats(res)
		d.Set("stats", NetflowCollectorTemplateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetflowCollectorTemplateStatsStats(ret edpt.DataNetflowCollectorTemplateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"v9_templates_created":        ret.DtNetflowCollectorTemplateStats.Stats.V9TemplatesCreated,
			"v9_templates_deleted":        ret.DtNetflowCollectorTemplateStats.Stats.V9TemplatesDeleted,
			"v10_templates_created":       ret.DtNetflowCollectorTemplateStats.Stats.V10TemplatesCreated,
			"v10_templates_deleted":       ret.DtNetflowCollectorTemplateStats.Stats.V10TemplatesDeleted,
			"template_drop_exceeded":      ret.DtNetflowCollectorTemplateStats.Stats.TemplateDropExceeded,
			"template_drop_out_of_memory": ret.DtNetflowCollectorTemplateStats.Stats.TemplateDropOutOfMemory,
		},
	}
}

func getObjectNetflowCollectorTemplateStatsStats(d []interface{}) edpt.NetflowCollectorTemplateStatsStats {

	count1 := len(d)
	var ret edpt.NetflowCollectorTemplateStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V9TemplatesCreated = in["v9_templates_created"].(int)
		ret.V9TemplatesDeleted = in["v9_templates_deleted"].(int)
		ret.V10TemplatesCreated = in["v10_templates_created"].(int)
		ret.V10TemplatesDeleted = in["v10_templates_deleted"].(int)
		ret.TemplateDropExceeded = in["template_drop_exceeded"].(int)
		ret.TemplateDropOutOfMemory = in["template_drop_out_of_memory"].(int)
	}
	return ret
}

func dataToEndpointNetflowCollectorTemplateStats(d *schema.ResourceData) edpt.NetflowCollectorTemplateStats {
	var ret edpt.NetflowCollectorTemplateStats

	ret.Stats = getObjectNetflowCollectorTemplateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
