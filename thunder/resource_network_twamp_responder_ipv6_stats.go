package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkTwampResponderIpv6Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_twamp_responder_ipv6_stats`: Statistics for the object ipv6\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkTwampResponderIpv6StatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rx_pkts_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Rx IPv6 TWAMP test packets",
						},
						"tx_pkts_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Tx IPv6 TWAMP test packets",
						},
						"rx_drop_not_enabled_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Rx IPv6 disabled drop",
						},
						"rx_acl_drop_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Rx IPv6 client-list drop",
						},
						"twamp_hdr_len_err_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Rx IPv6 TWAMP hdr length error drop",
						},
						"no_route_err_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Tx IPv6 no route error drop",
						},
						"other_err_v6": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Other error drop",
						},
					},
				},
			},
		},
	}
}

func resourceNetworkTwampResponderIpv6StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpv6StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIpv6Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkTwampResponderIpv6StatsStats := setObjectNetworkTwampResponderIpv6StatsStats(res)
		d.Set("stats", NetworkTwampResponderIpv6StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkTwampResponderIpv6StatsStats(ret edpt.DataNetworkTwampResponderIpv6Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rx_pkts_v6":             ret.DtNetworkTwampResponderIpv6Stats.Stats.Rx_pkts_v6,
			"tx_pkts_v6":             ret.DtNetworkTwampResponderIpv6Stats.Stats.Tx_pkts_v6,
			"rx_drop_not_enabled_v6": ret.DtNetworkTwampResponderIpv6Stats.Stats.Rx_drop_not_enabled_v6,
			"rx_acl_drop_v6":         ret.DtNetworkTwampResponderIpv6Stats.Stats.Rx_acl_drop_v6,
			"twamp_hdr_len_err_v6":   ret.DtNetworkTwampResponderIpv6Stats.Stats.Twamp_hdr_len_err_v6,
			"no_route_err_v6":        ret.DtNetworkTwampResponderIpv6Stats.Stats.No_route_err_v6,
			"other_err_v6":           ret.DtNetworkTwampResponderIpv6Stats.Stats.Other_err_v6,
		},
	}
}

func getObjectNetworkTwampResponderIpv6StatsStats(d []interface{}) edpt.NetworkTwampResponderIpv6StatsStats {

	count1 := len(d)
	var ret edpt.NetworkTwampResponderIpv6StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rx_pkts_v6 = in["rx_pkts_v6"].(int)
		ret.Tx_pkts_v6 = in["tx_pkts_v6"].(int)
		ret.Rx_drop_not_enabled_v6 = in["rx_drop_not_enabled_v6"].(int)
		ret.Rx_acl_drop_v6 = in["rx_acl_drop_v6"].(int)
		ret.Twamp_hdr_len_err_v6 = in["twamp_hdr_len_err_v6"].(int)
		ret.No_route_err_v6 = in["no_route_err_v6"].(int)
		ret.Other_err_v6 = in["other_err_v6"].(int)
	}
	return ret
}

func dataToEndpointNetworkTwampResponderIpv6Stats(d *schema.ResourceData) edpt.NetworkTwampResponderIpv6Stats {
	var ret edpt.NetworkTwampResponderIpv6Stats

	ret.Stats = getObjectNetworkTwampResponderIpv6StatsStats(d.Get("stats").([]interface{}))
	return ret
}
