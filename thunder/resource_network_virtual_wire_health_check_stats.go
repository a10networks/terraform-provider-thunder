package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireHealthCheckStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_virtual_wire_health_check_stats`: Statistics for the object virtual-wire-health-check\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVirtualWireHealthCheckStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"act_event": {
							Type: schema.TypeInt, Optional: true, Description: "Active Event Count",
						},
						"sby_event": {
							Type: schema.TypeInt, Optional: true, Description: "Standby Event Count",
						},
						"packet_count": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Count",
						},
					},
				},
			},
			"vlan": {
				Type: schema.TypeInt, Required: true, Description: "VLAN ID, specify 1 for untagged traffic",
			},
		},
	}
}

func resourceNetworkVirtualWireHealthCheckStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireHealthCheckStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireHealthCheckStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVirtualWireHealthCheckStatsStats := setObjectNetworkVirtualWireHealthCheckStatsStats(res)
		d.Set("stats", NetworkVirtualWireHealthCheckStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVirtualWireHealthCheckStatsStats(ret edpt.DataNetworkVirtualWireHealthCheckStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"act_event":    ret.DtNetworkVirtualWireHealthCheckStats.Stats.ActEvent,
			"sby_event":    ret.DtNetworkVirtualWireHealthCheckStats.Stats.SbyEvent,
			"packet_count": ret.DtNetworkVirtualWireHealthCheckStats.Stats.PacketCount,
		},
	}
}

func getObjectNetworkVirtualWireHealthCheckStatsStats(d []interface{}) edpt.NetworkVirtualWireHealthCheckStatsStats {

	count1 := len(d)
	var ret edpt.NetworkVirtualWireHealthCheckStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActEvent = in["act_event"].(int)
		ret.SbyEvent = in["sby_event"].(int)
		ret.PacketCount = in["packet_count"].(int)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireHealthCheckStats(d *schema.ResourceData) edpt.NetworkVirtualWireHealthCheckStats {
	var ret edpt.NetworkVirtualWireHealthCheckStats

	ret.Stats = getObjectNetworkVirtualWireHealthCheckStatsStats(d.Get("stats").([]interface{}))

	ret.Vlan = d.Get("vlan").(int)
	return ret
}
