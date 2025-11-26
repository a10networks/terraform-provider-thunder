package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnelStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_tunnel_stats`: Statistics for the object tunnel\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceTunnelStatsRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Tunnel interface number",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_rx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "received packets",
						},
						"num_total_rx_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "received bytes",
						},
						"num_tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "sent packets",
						},
						"num_total_tx_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "sent bytes",
						},
						"num_rx_err_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "received error packets",
						},
						"num_tx_err_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "sent error packets",
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
					},
				},
			},
		},
	}
}

func resourceInterfaceTunnelStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceTunnelStatsStats := setObjectInterfaceTunnelStatsStats(res)
		d.Set("stats", InterfaceTunnelStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceTunnelStatsStats(ret edpt.DataInterfaceTunnelStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num_rx_pkts":        ret.DtInterfaceTunnelStats.Stats.NumRxPkts,
			"num_total_rx_bytes": ret.DtInterfaceTunnelStats.Stats.NumTotalRxBytes,
			"num_tx_pkts":        ret.DtInterfaceTunnelStats.Stats.NumTxPkts,
			"num_total_tx_bytes": ret.DtInterfaceTunnelStats.Stats.NumTotalTxBytes,
			"num_rx_err_pkts":    ret.DtInterfaceTunnelStats.Stats.NumRxErrPkts,
			"num_tx_err_pkts":    ret.DtInterfaceTunnelStats.Stats.NumTxErrPkts,
			"rate_pkt_sent":      ret.DtInterfaceTunnelStats.Stats.Rate_pkt_sent,
			"rate_byte_sent":     ret.DtInterfaceTunnelStats.Stats.Rate_byte_sent,
			"rate_pkt_rcvd":      ret.DtInterfaceTunnelStats.Stats.Rate_pkt_rcvd,
			"rate_byte_rcvd":     ret.DtInterfaceTunnelStats.Stats.Rate_byte_rcvd,
		},
	}
}

func getObjectInterfaceTunnelStatsStats(d []interface{}) edpt.InterfaceTunnelStatsStats {

	count1 := len(d)
	var ret edpt.InterfaceTunnelStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumRxPkts = in["num_rx_pkts"].(int)
		ret.NumTotalRxBytes = in["num_total_rx_bytes"].(int)
		ret.NumTxPkts = in["num_tx_pkts"].(int)
		ret.NumTotalTxBytes = in["num_total_tx_bytes"].(int)
		ret.NumRxErrPkts = in["num_rx_err_pkts"].(int)
		ret.NumTxErrPkts = in["num_tx_err_pkts"].(int)
		ret.Rate_pkt_sent = in["rate_pkt_sent"].(int)
		ret.Rate_byte_sent = in["rate_byte_sent"].(int)
		ret.Rate_pkt_rcvd = in["rate_pkt_rcvd"].(int)
		ret.Rate_byte_rcvd = in["rate_byte_rcvd"].(int)
	}
	return ret
}

func dataToEndpointInterfaceTunnelStats(d *schema.ResourceData) edpt.InterfaceTunnelStats {
	var ret edpt.InterfaceTunnelStats

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Stats = getObjectInterfaceTunnelStatsStats(d.Get("stats").([]interface{}))
	return ret
}
