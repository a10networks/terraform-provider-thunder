package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpWindowCheckStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_tcp_window_check_stats`: Statistics for the object tcp-window-check\n\n__PLACEHOLDER__",
		ReadContext: resourceFwTcpWindowCheckStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside_window": {
							Type: schema.TypeInt, Optional: true, Description: "packet dropped counter for outside of tcp window",
						},
					},
				},
			},
		},
	}
}

func resourceFwTcpWindowCheckStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpWindowCheckStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpWindowCheckStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwTcpWindowCheckStatsStats := setObjectFwTcpWindowCheckStatsStats(res)
		d.Set("stats", FwTcpWindowCheckStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwTcpWindowCheckStatsStats(ret edpt.DataFwTcpWindowCheckStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"outside_window": ret.DtFwTcpWindowCheckStats.Stats.OutsideWindow,
		},
	}
}

func getObjectFwTcpWindowCheckStatsStats(d []interface{}) edpt.FwTcpWindowCheckStatsStats {

	count1 := len(d)
	var ret edpt.FwTcpWindowCheckStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutsideWindow = in["outside_window"].(int)
	}
	return ret
}

func dataToEndpointFwTcpWindowCheckStats(d *schema.ResourceData) edpt.FwTcpWindowCheckStats {
	var ret edpt.FwTcpWindowCheckStats

	ret.Stats = getObjectFwTcpWindowCheckStatsStats(d.Get("stats").([]interface{}))
	return ret
}
