package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_lif_stats`: Statistics for the object lif\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceLifStatsRead,

		Schema: map[string]*schema.Schema{
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Lif interface name",
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
						"dropped_rx_pkts_gre_key": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceLifStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceLifStatsStats := setObjectInterfaceLifStatsStats(res)
		d.Set("stats", InterfaceLifStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceLifStatsStats(ret edpt.DataInterfaceLifStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num_pkts":                ret.DtInterfaceLifStats.Stats.Num_pkts,
			"num_total_bytes":         ret.DtInterfaceLifStats.Stats.Num_total_bytes,
			"num_unicast_pkts":        ret.DtInterfaceLifStats.Stats.Num_unicast_pkts,
			"num_broadcast_pkts":      ret.DtInterfaceLifStats.Stats.Num_broadcast_pkts,
			"num_multicast_pkts":      ret.DtInterfaceLifStats.Stats.Num_multicast_pkts,
			"num_tx_pkts":             ret.DtInterfaceLifStats.Stats.Num_tx_pkts,
			"num_total_tx_bytes":      ret.DtInterfaceLifStats.Stats.Num_total_tx_bytes,
			"num_unicast_tx_pkts":     ret.DtInterfaceLifStats.Stats.Num_unicast_tx_pkts,
			"num_broadcast_tx_pkts":   ret.DtInterfaceLifStats.Stats.Num_broadcast_tx_pkts,
			"num_multicast_tx_pkts":   ret.DtInterfaceLifStats.Stats.Num_multicast_tx_pkts,
			"dropped_dis_rx_pkts":     ret.DtInterfaceLifStats.Stats.Dropped_dis_rx_pkts,
			"dropped_rx_pkts":         ret.DtInterfaceLifStats.Stats.Dropped_rx_pkts,
			"dropped_dis_tx_pkts":     ret.DtInterfaceLifStats.Stats.Dropped_dis_tx_pkts,
			"dropped_tx_pkts":         ret.DtInterfaceLifStats.Stats.Dropped_tx_pkts,
			"dropped_rx_pkts_gre_key": ret.DtInterfaceLifStats.Stats.Dropped_rx_pkts_gre_key,
		},
	}
}

func getObjectInterfaceLifStatsStats(d []interface{}) edpt.InterfaceLifStatsStats {

	count1 := len(d)
	var ret edpt.InterfaceLifStatsStats
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
		ret.Dropped_rx_pkts_gre_key = in["dropped_rx_pkts_gre_key"].(int)
	}
	return ret
}

func dataToEndpointInterfaceLifStats(d *schema.ResourceData) edpt.InterfaceLifStats {
	var ret edpt.InterfaceLifStats

	ret.Ifname = d.Get("ifname").(string)

	ret.Stats = getObjectInterfaceLifStatsStats(d.Get("stats").([]interface{}))
	return ret
}
