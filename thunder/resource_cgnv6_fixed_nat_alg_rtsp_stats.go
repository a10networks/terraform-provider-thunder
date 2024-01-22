package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgRtspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_alg_rtsp_stats`: Statistics for the object rtsp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatAlgRtspStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"streams_created": {
							Type: schema.TypeInt, Optional: true, Description: "Streams Created",
						},
						"streams_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Streams Freed",
						},
						"stream_creation_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Stream Creation Failures",
						},
						"ports_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Stream Client Ports Allocated",
						},
						"ports_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Stream Client Ports Freed",
						},
						"port_allocation_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Stream Client Port Allocation Failures",
						},
						"unknown_client_port_from_server": {
							Type: schema.TypeInt, Optional: true, Description: "Server Replies With Unknown Client Ports",
						},
						"data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Created",
						},
						"data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Freed",
						},
						"no_session_mem": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Creation Failures",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatAlgRtspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgRtspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgRtspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatAlgRtspStatsStats := setObjectCgnv6FixedNatAlgRtspStatsStats(res)
		d.Set("stats", Cgnv6FixedNatAlgRtspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatAlgRtspStatsStats(ret edpt.DataCgnv6FixedNatAlgRtspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"streams_created":                 ret.DtCgnv6FixedNatAlgRtspStats.Stats.StreamsCreated,
			"streams_freed":                   ret.DtCgnv6FixedNatAlgRtspStats.Stats.StreamsFreed,
			"stream_creation_failure":         ret.DtCgnv6FixedNatAlgRtspStats.Stats.StreamCreationFailure,
			"ports_allocated":                 ret.DtCgnv6FixedNatAlgRtspStats.Stats.PortsAllocated,
			"ports_freed":                     ret.DtCgnv6FixedNatAlgRtspStats.Stats.PortsFreed,
			"port_allocation_failure":         ret.DtCgnv6FixedNatAlgRtspStats.Stats.PortAllocationFailure,
			"unknown_client_port_from_server": ret.DtCgnv6FixedNatAlgRtspStats.Stats.UnknownClientPortFromServer,
			"data_session_created":            ret.DtCgnv6FixedNatAlgRtspStats.Stats.DataSessionCreated,
			"data_session_freed":              ret.DtCgnv6FixedNatAlgRtspStats.Stats.DataSessionFreed,
			"no_session_mem":                  ret.DtCgnv6FixedNatAlgRtspStats.Stats.NoSessionMem,
		},
	}
}

func getObjectCgnv6FixedNatAlgRtspStatsStats(d []interface{}) edpt.Cgnv6FixedNatAlgRtspStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatAlgRtspStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StreamsCreated = in["streams_created"].(int)
		ret.StreamsFreed = in["streams_freed"].(int)
		ret.StreamCreationFailure = in["stream_creation_failure"].(int)
		ret.PortsAllocated = in["ports_allocated"].(int)
		ret.PortsFreed = in["ports_freed"].(int)
		ret.PortAllocationFailure = in["port_allocation_failure"].(int)
		ret.UnknownClientPortFromServer = in["unknown_client_port_from_server"].(int)
		ret.DataSessionCreated = in["data_session_created"].(int)
		ret.DataSessionFreed = in["data_session_freed"].(int)
		ret.NoSessionMem = in["no_session_mem"].(int)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgRtspStats(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgRtspStats {
	var ret edpt.Cgnv6FixedNatAlgRtspStats

	ret.Stats = getObjectCgnv6FixedNatAlgRtspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
