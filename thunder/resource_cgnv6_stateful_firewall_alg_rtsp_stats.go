package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgRtspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_stateful_firewall_alg_rtsp_stats`: Statistics for the object rtsp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6StatefulFirewallAlgRtspStatsRead,

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

func resourceCgnv6StatefulFirewallAlgRtspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgRtspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgRtspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6StatefulFirewallAlgRtspStatsStats := setObjectCgnv6StatefulFirewallAlgRtspStatsStats(res)
		d.Set("stats", Cgnv6StatefulFirewallAlgRtspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6StatefulFirewallAlgRtspStatsStats(ret edpt.DataCgnv6StatefulFirewallAlgRtspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"transport_inserted":      ret.DtCgnv6StatefulFirewallAlgRtspStats.Stats.TransportInserted,
			"transport_freed":         ret.DtCgnv6StatefulFirewallAlgRtspStats.Stats.TransportFreed,
			"transport_alloc_failure": ret.DtCgnv6StatefulFirewallAlgRtspStats.Stats.TransportAllocFailure,
			"data_session_created":    ret.DtCgnv6StatefulFirewallAlgRtspStats.Stats.DataSessionCreated,
			"data_session_freed":      ret.DtCgnv6StatefulFirewallAlgRtspStats.Stats.DataSessionFreed,
		},
	}
}

func getObjectCgnv6StatefulFirewallAlgRtspStatsStats(d []interface{}) edpt.Cgnv6StatefulFirewallAlgRtspStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6StatefulFirewallAlgRtspStatsStats
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

func dataToEndpointCgnv6StatefulFirewallAlgRtspStats(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgRtspStats {
	var ret edpt.Cgnv6StatefulFirewallAlgRtspStats

	ret.Stats = getObjectCgnv6StatefulFirewallAlgRtspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
