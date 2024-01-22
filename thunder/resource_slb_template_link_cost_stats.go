package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLinkCostStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_link_cost_stats`: Statistics for the object link-cost\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateLinkCostStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Link-Cost Template Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_total_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Total bytes fwd for link",
						},
						"interval_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Link Cost bytes processed in forward direction per interval",
						},
						"link_total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Link Cost total connection",
						},
						"interval_avg_throughput": {
							Type: schema.TypeInt, Optional: true, Description: "Link Cost average throughput per interval",
						},
						"interval_max_throughput": {
							Type: schema.TypeInt, Optional: true, Description: "Link Cost max throughput per interval",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateLinkCostStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkCostStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkCostStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateLinkCostStatsStats := setObjectSlbTemplateLinkCostStatsStats(res)
		d.Set("stats", SlbTemplateLinkCostStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateLinkCostStatsStats(ret edpt.DataSlbTemplateLinkCostStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"link_total_fwd_bytes":    ret.DtSlbTemplateLinkCostStats.Stats.Link_total_fwd_bytes,
			"interval_fwd_bytes":      ret.DtSlbTemplateLinkCostStats.Stats.Interval_fwd_bytes,
			"link_total_conn":         ret.DtSlbTemplateLinkCostStats.Stats.Link_total_conn,
			"interval_avg_throughput": ret.DtSlbTemplateLinkCostStats.Stats.Interval_avg_throughput,
			"interval_max_throughput": ret.DtSlbTemplateLinkCostStats.Stats.Interval_max_throughput,
		},
	}
}

func getObjectSlbTemplateLinkCostStatsStats(d []interface{}) edpt.SlbTemplateLinkCostStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplateLinkCostStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Link_total_fwd_bytes = in["link_total_fwd_bytes"].(int)
		ret.Interval_fwd_bytes = in["interval_fwd_bytes"].(int)
		ret.Link_total_conn = in["link_total_conn"].(int)
		ret.Interval_avg_throughput = in["interval_avg_throughput"].(int)
		ret.Interval_max_throughput = in["interval_max_throughput"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateLinkCostStats(d *schema.ResourceData) edpt.SlbTemplateLinkCostStats {
	var ret edpt.SlbTemplateLinkCostStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplateLinkCostStatsStats(d.Get("stats").([]interface{}))
	return ret
}
