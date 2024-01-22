package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorSflowStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_flow_collector_sflow_stats`: Statistics for the object sflow\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityFlowCollectorSflowStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkts_received": {
							Type: schema.TypeInt, Optional: true, Description: "Total sflow pkts received",
						},
						"frag_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total sflow fragment packets droppped",
						},
						"agent_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "sflow pkts from not configured agents",
						},
						"version_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "sflow version not supported",
						},
						"unknown_dir": {
							Type: schema.TypeInt, Optional: true, Description: "sflow sample direction is unknown",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityFlowCollectorSflowStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorSflowStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorSflowStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityFlowCollectorSflowStatsStats := setObjectVisibilityFlowCollectorSflowStatsStats(res)
		d.Set("stats", VisibilityFlowCollectorSflowStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityFlowCollectorSflowStatsStats(ret edpt.DataVisibilityFlowCollectorSflowStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pkts_received":         ret.DtVisibilityFlowCollectorSflowStats.Stats.PktsReceived,
			"frag_dropped":          ret.DtVisibilityFlowCollectorSflowStats.Stats.FragDropped,
			"agent_not_found":       ret.DtVisibilityFlowCollectorSflowStats.Stats.AgentNotFound,
			"version_not_supported": ret.DtVisibilityFlowCollectorSflowStats.Stats.VersionNotSupported,
			"unknown_dir":           ret.DtVisibilityFlowCollectorSflowStats.Stats.UnknownDir,
		},
	}
}

func getObjectVisibilityFlowCollectorSflowStatsStats(d []interface{}) edpt.VisibilityFlowCollectorSflowStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorSflowStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktsReceived = in["pkts_received"].(int)
		ret.FragDropped = in["frag_dropped"].(int)
		ret.AgentNotFound = in["agent_not_found"].(int)
		ret.VersionNotSupported = in["version_not_supported"].(int)
		ret.UnknownDir = in["unknown_dir"].(int)
	}
	return ret
}

func dataToEndpointVisibilityFlowCollectorSflowStats(d *schema.ResourceData) edpt.VisibilityFlowCollectorSflowStats {
	var ret edpt.VisibilityFlowCollectorSflowStats

	ret.Stats = getObjectVisibilityFlowCollectorSflowStatsStats(d.Get("stats").([]interface{}))
	return ret
}
