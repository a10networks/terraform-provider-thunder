package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_virtual_wire_global_stats`: Statistics for the object virtual-wire-global\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVirtualWireGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_update": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN update",
						},
						"mac_update": {
							Type: schema.TypeInt, Optional: true, Description: "MAC update",
						},
						"vlan_pair_update": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN pair update",
						},
					},
				},
			},
		},
	}
}

func resourceNetworkVirtualWireGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVirtualWireGlobalStatsStats := setObjectNetworkVirtualWireGlobalStatsStats(res)
		d.Set("stats", NetworkVirtualWireGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVirtualWireGlobalStatsStats(ret edpt.DataNetworkVirtualWireGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vlan_update":      ret.DtNetworkVirtualWireGlobalStats.Stats.VlanUpdate,
			"mac_update":       ret.DtNetworkVirtualWireGlobalStats.Stats.MacUpdate,
			"vlan_pair_update": ret.DtNetworkVirtualWireGlobalStats.Stats.VlanPairUpdate,
		},
	}
}

func getObjectNetworkVirtualWireGlobalStatsStats(d []interface{}) edpt.NetworkVirtualWireGlobalStatsStats {

	count1 := len(d)
	var ret edpt.NetworkVirtualWireGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VlanUpdate = in["vlan_update"].(int)
		ret.MacUpdate = in["mac_update"].(int)
		ret.VlanPairUpdate = in["vlan_pair_update"].(int)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireGlobalStats(d *schema.ResourceData) edpt.NetworkVirtualWireGlobalStats {
	var ret edpt.NetworkVirtualWireGlobalStats

	ret.Stats = getObjectNetworkVirtualWireGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
