package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgRtspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_alg_rtsp_stats`: Statistics for the object rtsp\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAlgRtspStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transport_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Transport Created",
						},
						"transport_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Transport Freed",
						},
						"transport_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Transport Alloc Failure",
						},
						"data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Created",
						},
						"data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Freed",
						},
					},
				},
			},
		},
	}
}

func resourceFwAlgRtspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgRtspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgRtspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAlgRtspStatsStats := setObjectFwAlgRtspStatsStats(res)
		d.Set("stats", FwAlgRtspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAlgRtspStatsStats(ret edpt.DataFwAlgRtspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"transport_inserted":      ret.DtFwAlgRtspStats.Stats.TransportInserted,
			"transport_freed":         ret.DtFwAlgRtspStats.Stats.TransportFreed,
			"transport_alloc_failure": ret.DtFwAlgRtspStats.Stats.TransportAllocFailure,
			"data_session_created":    ret.DtFwAlgRtspStats.Stats.DataSessionCreated,
			"data_session_freed":      ret.DtFwAlgRtspStats.Stats.DataSessionFreed,
		},
	}
}

func getObjectFwAlgRtspStatsStats(d []interface{}) edpt.FwAlgRtspStatsStats {

	count1 := len(d)
	var ret edpt.FwAlgRtspStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TransportInserted = in["transport_inserted"].(int)
		ret.TransportFreed = in["transport_freed"].(int)
		ret.TransportAllocFailure = in["transport_alloc_failure"].(int)
		ret.DataSessionCreated = in["data_session_created"].(int)
		ret.DataSessionFreed = in["data_session_freed"].(int)
	}
	return ret
}

func dataToEndpointFwAlgRtspStats(d *schema.ResourceData) edpt.FwAlgRtspStats {
	var ret edpt.FwAlgRtspStats

	ret.Stats = getObjectFwAlgRtspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
