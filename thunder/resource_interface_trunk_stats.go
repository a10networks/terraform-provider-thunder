package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_trunk_stats`: Statistics for the object trunk\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceTrunkStatsRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Trunk interface number",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_total_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_unicast_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_broadcast_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_multicast_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_total_tx_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_unicast_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_broadcast_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_multicast_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dropped_dis_rx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dropped_rx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dropped_dis_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dropped_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceTrunkStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceTrunkStatsStats := setObjectInterfaceTrunkStatsStats(res)
		d.Set("stats", InterfaceTrunkStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceTrunkStatsStats(ret edpt.DataInterfaceTrunkStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num_pkts":              ret.DtInterfaceTrunkStats.Stats.Num_pkts,
			"num_total_bytes":       ret.DtInterfaceTrunkStats.Stats.Num_total_bytes,
			"num_unicast_pkts":      ret.DtInterfaceTrunkStats.Stats.Num_unicast_pkts,
			"num_broadcast_pkts":    ret.DtInterfaceTrunkStats.Stats.Num_broadcast_pkts,
			"num_multicast_pkts":    ret.DtInterfaceTrunkStats.Stats.Num_multicast_pkts,
			"num_tx_pkts":           ret.DtInterfaceTrunkStats.Stats.Num_tx_pkts,
			"num_total_tx_bytes":    ret.DtInterfaceTrunkStats.Stats.Num_total_tx_bytes,
			"num_unicast_tx_pkts":   ret.DtInterfaceTrunkStats.Stats.Num_unicast_tx_pkts,
			"num_broadcast_tx_pkts": ret.DtInterfaceTrunkStats.Stats.Num_broadcast_tx_pkts,
			"num_multicast_tx_pkts": ret.DtInterfaceTrunkStats.Stats.Num_multicast_tx_pkts,
			"dropped_dis_rx_pkts":   ret.DtInterfaceTrunkStats.Stats.Dropped_dis_rx_pkts,
			"dropped_rx_pkts":       ret.DtInterfaceTrunkStats.Stats.Dropped_rx_pkts,
			"dropped_dis_tx_pkts":   ret.DtInterfaceTrunkStats.Stats.Dropped_dis_tx_pkts,
			"dropped_tx_pkts":       ret.DtInterfaceTrunkStats.Stats.Dropped_tx_pkts,
		},
	}
}

func getObjectInterfaceTrunkStatsStats(d []interface{}) edpt.InterfaceTrunkStatsStats {

	count1 := len(d)
	var ret edpt.InterfaceTrunkStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num_pkts = in["num_pkts"].(int)
		ret.Num_total_bytes = in["num_total_bytes"].(int)
		ret.Num_unicast_pkts = in["num_unicast_pkts"].(int)
		ret.Num_broadcast_pkts = in["num_broadcast_pkts"].(int)
		ret.Num_multicast_pkts = in["num_multicast_pkts"].(int)
		ret.Num_tx_pkts = in["num_tx_pkts"].(int)
		ret.Num_total_tx_bytes = in["num_total_tx_bytes"].(int)
		ret.Num_unicast_tx_pkts = in["num_unicast_tx_pkts"].(int)
		ret.Num_broadcast_tx_pkts = in["num_broadcast_tx_pkts"].(int)
		ret.Num_multicast_tx_pkts = in["num_multicast_tx_pkts"].(int)
		ret.Dropped_dis_rx_pkts = in["dropped_dis_rx_pkts"].(int)
		ret.Dropped_rx_pkts = in["dropped_rx_pkts"].(int)
		ret.Dropped_dis_tx_pkts = in["dropped_dis_tx_pkts"].(int)
		ret.Dropped_tx_pkts = in["dropped_tx_pkts"].(int)
	}
	return ret
}

func dataToEndpointInterfaceTrunkStats(d *schema.ResourceData) edpt.InterfaceTrunkStats {
	var ret edpt.InterfaceTrunkStats

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Stats = getObjectInterfaceTrunkStatsStats(d.Get("stats").([]interface{}))
	return ret
}
