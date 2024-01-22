package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPassthroughStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_passthrough_stats`: Statistics for the object passthrough\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPassthroughStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current connections",
						},
						"total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total connections",
						},
						"total_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Forward bytes",
						},
						"total_fwd_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Forward packets",
						},
						"total_rev_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse bytes",
						},
						"total_rev_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse packets",
						},
						"curr_pconn": {
							Type: schema.TypeInt, Optional: true, Description: "Persistent connections",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPassthroughStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPassthroughStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPassthroughStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPassthroughStatsStats := setObjectSlbPassthroughStatsStats(res)
		d.Set("stats", SlbPassthroughStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPassthroughStatsStats(ret edpt.DataSlbPassthroughStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":         ret.DtSlbPassthroughStats.Stats.Curr_conn,
			"total_conn":        ret.DtSlbPassthroughStats.Stats.Total_conn,
			"total_fwd_bytes":   ret.DtSlbPassthroughStats.Stats.Total_fwd_bytes,
			"total_fwd_packets": ret.DtSlbPassthroughStats.Stats.Total_fwd_packets,
			"total_rev_bytes":   ret.DtSlbPassthroughStats.Stats.Total_rev_bytes,
			"total_rev_packets": ret.DtSlbPassthroughStats.Stats.Total_rev_packets,
			"curr_pconn":        ret.DtSlbPassthroughStats.Stats.Curr_pconn,
		},
	}
}

func getObjectSlbPassthroughStatsStats(d []interface{}) edpt.SlbPassthroughStatsStats {

	count1 := len(d)
	var ret edpt.SlbPassthroughStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_conn = in["curr_conn"].(int)
		ret.Total_conn = in["total_conn"].(int)
		ret.Total_fwd_bytes = in["total_fwd_bytes"].(int)
		ret.Total_fwd_packets = in["total_fwd_packets"].(int)
		ret.Total_rev_bytes = in["total_rev_bytes"].(int)
		ret.Total_rev_packets = in["total_rev_packets"].(int)
		ret.Curr_pconn = in["curr_pconn"].(int)
	}
	return ret
}

func dataToEndpointSlbPassthroughStats(d *schema.ResourceData) edpt.SlbPassthroughStats {
	var ret edpt.SlbPassthroughStats

	ret.Stats = getObjectSlbPassthroughStatsStats(d.Get("stats").([]interface{}))
	return ret
}
