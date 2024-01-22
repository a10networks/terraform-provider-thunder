package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVlanStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_vlan_stats`: Statistics for the object vlan\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVlanStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"broadcast_count": {
							Type: schema.TypeInt, Optional: true, Description: "Broadcast counter",
						},
						"multicast_count": {
							Type: schema.TypeInt, Optional: true, Description: "Multicast counter",
						},
						"ip_multicast_count": {
							Type: schema.TypeInt, Optional: true, Description: "IP Multicast counter",
						},
						"unknown_unicast_count": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Unicast counter",
						},
						"mac_movement_count": {
							Type: schema.TypeInt, Optional: true, Description: "Mac Movement counter",
						},
						"shared_vlan_partition_switched_counter": {
							Type: schema.TypeInt, Optional: true, Description: "SVLAN Partition switched counter",
						},
					},
				},
			},
			"vlan_num": {
				Type: schema.TypeInt, Required: true, Description: "VLAN number",
			},
		},
	}
}

func resourceNetworkVlanStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVlanStatsStats := setObjectNetworkVlanStatsStats(res)
		d.Set("stats", NetworkVlanStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVlanStatsStats(ret edpt.DataNetworkVlanStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"broadcast_count":                        ret.DtNetworkVlanStats.Stats.Broadcast_count,
			"multicast_count":                        ret.DtNetworkVlanStats.Stats.Multicast_count,
			"ip_multicast_count":                     ret.DtNetworkVlanStats.Stats.Ip_multicast_count,
			"unknown_unicast_count":                  ret.DtNetworkVlanStats.Stats.Unknown_unicast_count,
			"mac_movement_count":                     ret.DtNetworkVlanStats.Stats.Mac_movement_count,
			"shared_vlan_partition_switched_counter": ret.DtNetworkVlanStats.Stats.Shared_vlan_partition_switched_counter,
		},
	}
}

func getObjectNetworkVlanStatsStats(d []interface{}) edpt.NetworkVlanStatsStats {

	count1 := len(d)
	var ret edpt.NetworkVlanStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Broadcast_count = in["broadcast_count"].(int)
		ret.Multicast_count = in["multicast_count"].(int)
		ret.Ip_multicast_count = in["ip_multicast_count"].(int)
		ret.Unknown_unicast_count = in["unknown_unicast_count"].(int)
		ret.Mac_movement_count = in["mac_movement_count"].(int)
		ret.Shared_vlan_partition_switched_counter = in["shared_vlan_partition_switched_counter"].(int)
	}
	return ret
}

func dataToEndpointNetworkVlanStats(d *schema.ResourceData) edpt.NetworkVlanStats {
	var ret edpt.NetworkVlanStats

	ret.Stats = getObjectNetworkVlanStatsStats(d.Get("stats").([]interface{}))

	ret.VlanNum = d.Get("vlan_num").(int)
	return ret
}
