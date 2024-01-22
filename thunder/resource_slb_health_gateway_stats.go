package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthGatewayStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_gateway_stats`: Statistics for the object health-gateway\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthGatewayStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Total health-check sent",
						},
						"total_retry_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Total health-check retry sent",
						},
						"total_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Total health-check timeout",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthGatewayStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthGatewayStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthGatewayStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthGatewayStatsStats := setObjectSlbHealthGatewayStatsStats(res)
		d.Set("stats", SlbHealthGatewayStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthGatewayStatsStats(ret edpt.DataSlbHealthGatewayStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_sent":       ret.DtSlbHealthGatewayStats.Stats.Total_sent,
			"total_retry_sent": ret.DtSlbHealthGatewayStats.Stats.Total_retry_sent,
			"total_timeout":    ret.DtSlbHealthGatewayStats.Stats.Total_timeout,
		},
	}
}

func getObjectSlbHealthGatewayStatsStats(d []interface{}) edpt.SlbHealthGatewayStatsStats {

	count1 := len(d)
	var ret edpt.SlbHealthGatewayStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_sent = in["total_sent"].(int)
		ret.Total_retry_sent = in["total_retry_sent"].(int)
		ret.Total_timeout = in["total_timeout"].(int)
	}
	return ret
}

func dataToEndpointSlbHealthGatewayStats(d *schema.ResourceData) edpt.SlbHealthGatewayStats {
	var ret edpt.SlbHealthGatewayStats

	ret.Stats = getObjectSlbHealthGatewayStatsStats(d.Get("stats").([]interface{}))
	return ret
}
