package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkTwampResponderIpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_twamp_responder_ip_stats`: Statistics for the object ip\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkTwampResponderIpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Rx IP TWAMP test packets",
						},
						"tx_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Tx IP TWAMP test packets",
						},
						"rx_drop_not_enabled_v4": {
							Type: schema.TypeInt, Optional: true, Description: "Rx IP disabled drop",
						},
						"rx_acl_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Rx IP client-list drop",
						},
						"twamp_hdr_len_err": {
							Type: schema.TypeInt, Optional: true, Description: "Rx TWAMP hdr length error drop",
						},
						"no_route_err": {
							Type: schema.TypeInt, Optional: true, Description: "Tx IP no route error drop",
						},
						"other_err": {
							Type: schema.TypeInt, Optional: true, Description: "IP other error drop",
						},
					},
				},
			},
		},
	}
}

func resourceNetworkTwampResponderIpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkTwampResponderIpStatsStats := setObjectNetworkTwampResponderIpStatsStats(res)
		d.Set("stats", NetworkTwampResponderIpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkTwampResponderIpStatsStats(ret edpt.DataNetworkTwampResponderIpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rx_pkts":                ret.DtNetworkTwampResponderIpStats.Stats.Rx_pkts,
			"tx_pkts":                ret.DtNetworkTwampResponderIpStats.Stats.Tx_pkts,
			"rx_drop_not_enabled_v4": ret.DtNetworkTwampResponderIpStats.Stats.Rx_drop_not_enabled_v4,
			"rx_acl_drop":            ret.DtNetworkTwampResponderIpStats.Stats.Rx_acl_drop,
			"twamp_hdr_len_err":      ret.DtNetworkTwampResponderIpStats.Stats.Twamp_hdr_len_err,
			"no_route_err":           ret.DtNetworkTwampResponderIpStats.Stats.No_route_err,
			"other_err":              ret.DtNetworkTwampResponderIpStats.Stats.Other_err,
		},
	}
}

func getObjectNetworkTwampResponderIpStatsStats(d []interface{}) edpt.NetworkTwampResponderIpStatsStats {

	count1 := len(d)
	var ret edpt.NetworkTwampResponderIpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rx_pkts = in["rx_pkts"].(int)
		ret.Tx_pkts = in["tx_pkts"].(int)
		ret.Rx_drop_not_enabled_v4 = in["rx_drop_not_enabled_v4"].(int)
		ret.Rx_acl_drop = in["rx_acl_drop"].(int)
		ret.Twamp_hdr_len_err = in["twamp_hdr_len_err"].(int)
		ret.No_route_err = in["no_route_err"].(int)
		ret.Other_err = in["other_err"].(int)
	}
	return ret
}

func dataToEndpointNetworkTwampResponderIpStats(d *schema.ResourceData) edpt.NetworkTwampResponderIpStats {
	var ret edpt.NetworkTwampResponderIpStats

	ret.Stats = getObjectNetworkTwampResponderIpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
