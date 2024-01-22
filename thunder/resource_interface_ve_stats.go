package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_ve_stats`: Statistics for the object ve\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceVeStatsRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Virtual ethernet interface number",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Input packets",
						},
						"num_total_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Input bytes",
						},
						"num_unicast_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Received unicasts",
						},
						"num_broadcast_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Received broadcasts",
						},
						"num_multicast_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Received multicasts",
						},
						"num_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted packets",
						},
						"num_total_tx_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted bytes",
						},
						"num_unicast_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted unicasts",
						},
						"num_broadcast_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted broadcasts",
						},
						"num_multicast_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted multicasts",
						},
						"rate_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Packet sent rate packets/sec",
						},
						"rate_byte_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Byte sent rate bits/sec",
						},
						"rate_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Packet received rate packets/sec",
						},
						"rate_byte_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Byte received rate bits/sec",
						},
						"load_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Load Interval",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceVeStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceVeStatsStats := setObjectInterfaceVeStatsStats(res)
		d.Set("stats", InterfaceVeStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceVeStatsStats(ret edpt.DataInterfaceVeStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num_pkts":              ret.DtInterfaceVeStats.Stats.Num_pkts,
			"num_total_bytes":       ret.DtInterfaceVeStats.Stats.Num_total_bytes,
			"num_unicast_pkts":      ret.DtInterfaceVeStats.Stats.Num_unicast_pkts,
			"num_broadcast_pkts":    ret.DtInterfaceVeStats.Stats.Num_broadcast_pkts,
			"num_multicast_pkts":    ret.DtInterfaceVeStats.Stats.Num_multicast_pkts,
			"num_tx_pkts":           ret.DtInterfaceVeStats.Stats.Num_tx_pkts,
			"num_total_tx_bytes":    ret.DtInterfaceVeStats.Stats.Num_total_tx_bytes,
			"num_unicast_tx_pkts":   ret.DtInterfaceVeStats.Stats.Num_unicast_tx_pkts,
			"num_broadcast_tx_pkts": ret.DtInterfaceVeStats.Stats.Num_broadcast_tx_pkts,
			"num_multicast_tx_pkts": ret.DtInterfaceVeStats.Stats.Num_multicast_tx_pkts,
			"rate_pkt_sent":         ret.DtInterfaceVeStats.Stats.Rate_pkt_sent,
			"rate_byte_sent":        ret.DtInterfaceVeStats.Stats.Rate_byte_sent,
			"rate_pkt_rcvd":         ret.DtInterfaceVeStats.Stats.Rate_pkt_rcvd,
			"rate_byte_rcvd":        ret.DtInterfaceVeStats.Stats.Rate_byte_rcvd,
			"load_interval":         ret.DtInterfaceVeStats.Stats.Load_interval,
		},
	}
}

func getObjectInterfaceVeStatsStats(d []interface{}) edpt.InterfaceVeStatsStats {

	count1 := len(d)
	var ret edpt.InterfaceVeStatsStats
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
		ret.Rate_pkt_sent = in["rate_pkt_sent"].(int)
		ret.Rate_byte_sent = in["rate_byte_sent"].(int)
		ret.Rate_pkt_rcvd = in["rate_pkt_rcvd"].(int)
		ret.Rate_byte_rcvd = in["rate_byte_rcvd"].(int)
		ret.Load_interval = in["load_interval"].(int)
	}
	return ret
}

func dataToEndpointInterfaceVeStats(d *schema.ResourceData) edpt.InterfaceVeStats {
	var ret edpt.InterfaceVeStats

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Stats = getObjectInterfaceVeStatsStats(d.Get("stats").([]interface{}))
	return ret
}
