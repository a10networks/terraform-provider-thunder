package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureAutomatedCapturesStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_packet_capture_automated_captures_stats`: Statistics for the object automated-captures\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPacketCaptureAutomatedCapturesStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total failures",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityPacketCaptureAutomatedCapturesStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureAutomatedCapturesStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureAutomatedCapturesStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPacketCaptureAutomatedCapturesStatsStats := setObjectVisibilityPacketCaptureAutomatedCapturesStatsStats(res)
		d.Set("stats", VisibilityPacketCaptureAutomatedCapturesStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPacketCaptureAutomatedCapturesStatsStats(ret edpt.DataVisibilityPacketCaptureAutomatedCapturesStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_failure": ret.DtVisibilityPacketCaptureAutomatedCapturesStats.Stats.TotalFailure,
		},
	}
}

func getObjectVisibilityPacketCaptureAutomatedCapturesStatsStats(d []interface{}) edpt.VisibilityPacketCaptureAutomatedCapturesStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureAutomatedCapturesStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalFailure = in["total_failure"].(int)
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureAutomatedCapturesStats(d *schema.ResourceData) edpt.VisibilityPacketCaptureAutomatedCapturesStats {
	var ret edpt.VisibilityPacketCaptureAutomatedCapturesStats

	ret.Stats = getObjectVisibilityPacketCaptureAutomatedCapturesStatsStats(d.Get("stats").([]interface{}))
	return ret
}
