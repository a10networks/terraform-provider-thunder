package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTunnelStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_tunnel_stats`: Statistics for the object tunnel\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosTunnelStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Received",
						},
						"ip_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Rate Limit Inner Pkts",
						},
						"ip_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Encap",
						},
						"ip_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Encap Failed",
						},
						"ip_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Decap",
						},
						"ip_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Decap Failed",
						},
						"ip6_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Received",
						},
						"ip6_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Rate Limit Inner Pkts",
						},
						"ip6_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Encap",
						},
						"ip6_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Encap Failed",
						},
						"ip6_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Decap",
						},
						"ip6_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Decap Failed",
						},
						"ip_gre_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Received",
						},
						"ip_gre_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Rate Limit Inner Pkts",
						},
						"ip_gre_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Encap",
						},
						"ip_gre_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Encap Failed",
						},
						"ip_gre_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap",
						},
						"ip_gre_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap Failed",
						},
						"ip_gre_tunnel_encap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Encap W/ Key",
						},
						"ip_gre_tunnel_decap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap W/ Key",
						},
						"ip_gre_tunnel_decap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap Key Mismatch Dropped",
						},
						"ip6_gre_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Received",
						},
						"ip6_gre_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Rate Limit Inner Pkts",
						},
						"ip6_gre_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Encap",
						},
						"ip6_gre_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Encap Received Failed",
						},
						"ip6_gre_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap",
						},
						"ip6_gre_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap Failed",
						},
						"ip6_gre_tunnel_encap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Encap W/ Key",
						},
						"ip6_gre_tunnel_decap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap W/ Key",
						},
						"ip6_gre_tunnel_decap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap No Key Dropped",
						},
						"ip_esp_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Tunnel Received",
						},
						"ip_esp_tunnel_inspect": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Tunnel Inspect",
						},
						"ip_esp_tunnel_inspect_fail": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Tunnel Inspect Fail",
						},
						"ip6_esp_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ESP V6 Tunnel Received",
						},
						"ip6_esp_tunnel_inspect": {
							Type: schema.TypeInt, Optional: true, Description: "ESP V6 Tunnel Inspect",
						},
						"ip6_esp_tunnel_inspect_fail": {
							Type: schema.TypeInt, Optional: true, Description: "ESP V6 Tunnel Inspect Fail",
						},
						"ip_vxlan_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Received",
						},
						"ip_vxlan_tunnel_invalid_vni": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Invalid VNI",
						},
						"ip_vxlan_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Decap",
						},
						"ip_vxlan_tunnel_decap_err": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Decap Error",
						},
						"ip_gre_tunnel_keepalive_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Keep-Alive Received",
						},
						"jumbo_in_tunnel_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Packet in Tunnel Drop",
						},
					},
				},
			},
		},
	}
}

func resourceDdosTunnelStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTunnelStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTunnelStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosTunnelStatsStats := setObjectDdosTunnelStatsStats(res)
		d.Set("stats", DdosTunnelStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosTunnelStatsStats(ret edpt.DataDdosTunnelStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_tunnel_rcvd":                  ret.DtDdosTunnelStats.Stats.Ip_tunnel_rcvd,
			"ip_tunnel_rate_limit_inner":      ret.DtDdosTunnelStats.Stats.Ip_tunnel_rate_limit_inner,
			"ip_tunnel_encap":                 ret.DtDdosTunnelStats.Stats.Ip_tunnel_encap,
			"ip_tunnel_encap_fail":            ret.DtDdosTunnelStats.Stats.Ip_tunnel_encap_fail,
			"ip_tunnel_decap":                 ret.DtDdosTunnelStats.Stats.Ip_tunnel_decap,
			"ip_tunnel_decap_fail":            ret.DtDdosTunnelStats.Stats.Ip_tunnel_decap_fail,
			"ip6_tunnel_rcvd":                 ret.DtDdosTunnelStats.Stats.Ip6_tunnel_rcvd,
			"ip6_tunnel_rate_limit_inner":     ret.DtDdosTunnelStats.Stats.Ip6_tunnel_rate_limit_inner,
			"ip6_tunnel_encap":                ret.DtDdosTunnelStats.Stats.Ip6_tunnel_encap,
			"ip6_tunnel_encap_fail":           ret.DtDdosTunnelStats.Stats.Ip6_tunnel_encap_fail,
			"ip6_tunnel_decap":                ret.DtDdosTunnelStats.Stats.Ip6_tunnel_decap,
			"ip6_tunnel_decap_fail":           ret.DtDdosTunnelStats.Stats.Ip6_tunnel_decap_fail,
			"ip_gre_tunnel_rcvd":              ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_rcvd,
			"ip_gre_tunnel_rate_limit_inner":  ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_rate_limit_inner,
			"ip_gre_tunnel_encap":             ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_encap,
			"ip_gre_tunnel_encap_fail":        ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_encap_fail,
			"ip_gre_tunnel_decap":             ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_decap,
			"ip_gre_tunnel_decap_fail":        ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_decap_fail,
			"ip_gre_tunnel_encap_key":         ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_encap_key,
			"ip_gre_tunnel_decap_key":         ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_decap_key,
			"ip_gre_tunnel_decap_drop":        ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_decap_drop,
			"ip6_gre_tunnel_rcvd":             ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_rcvd,
			"ip6_gre_tunnel_rate_limit_inner": ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_rate_limit_inner,
			"ip6_gre_tunnel_encap":            ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_encap,
			"ip6_gre_tunnel_encap_fail":       ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_encap_fail,
			"ip6_gre_tunnel_decap":            ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_decap,
			"ip6_gre_tunnel_decap_fail":       ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_decap_fail,
			"ip6_gre_tunnel_encap_key":        ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_encap_key,
			"ip6_gre_tunnel_decap_key":        ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_decap_key,
			"ip6_gre_tunnel_decap_drop":       ret.DtDdosTunnelStats.Stats.Ip6_gre_tunnel_decap_drop,
			"ip_esp_tunnel_rcvd":              ret.DtDdosTunnelStats.Stats.Ip_esp_tunnel_rcvd,
			"ip_esp_tunnel_inspect":           ret.DtDdosTunnelStats.Stats.Ip_esp_tunnel_inspect,
			"ip_esp_tunnel_inspect_fail":      ret.DtDdosTunnelStats.Stats.Ip_esp_tunnel_inspect_fail,
			"ip6_esp_tunnel_rcvd":             ret.DtDdosTunnelStats.Stats.Ip6_esp_tunnel_rcvd,
			"ip6_esp_tunnel_inspect":          ret.DtDdosTunnelStats.Stats.Ip6_esp_tunnel_inspect,
			"ip6_esp_tunnel_inspect_fail":     ret.DtDdosTunnelStats.Stats.Ip6_esp_tunnel_inspect_fail,
			"ip_vxlan_tunnel_rcvd":            ret.DtDdosTunnelStats.Stats.Ip_vxlan_tunnel_rcvd,
			"ip_vxlan_tunnel_invalid_vni":     ret.DtDdosTunnelStats.Stats.Ip_vxlan_tunnel_invalid_vni,
			"ip_vxlan_tunnel_decap":           ret.DtDdosTunnelStats.Stats.Ip_vxlan_tunnel_decap,
			"ip_vxlan_tunnel_decap_err":       ret.DtDdosTunnelStats.Stats.Ip_vxlan_tunnel_decap_err,
			"ip_gre_tunnel_keepalive_rcvd":    ret.DtDdosTunnelStats.Stats.Ip_gre_tunnel_keepalive_rcvd,
			"jumbo_in_tunnel_drop":            ret.DtDdosTunnelStats.Stats.Jumbo_in_tunnel_drop,
		},
	}
}

func getObjectDdosTunnelStatsStats(d []interface{}) edpt.DdosTunnelStatsStats {

	count1 := len(d)
	var ret edpt.DdosTunnelStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip_tunnel_rcvd = in["ip_tunnel_rcvd"].(int)
		ret.Ip_tunnel_rate_limit_inner = in["ip_tunnel_rate_limit_inner"].(int)
		ret.Ip_tunnel_encap = in["ip_tunnel_encap"].(int)
		ret.Ip_tunnel_encap_fail = in["ip_tunnel_encap_fail"].(int)
		ret.Ip_tunnel_decap = in["ip_tunnel_decap"].(int)
		ret.Ip_tunnel_decap_fail = in["ip_tunnel_decap_fail"].(int)
		ret.Ip6_tunnel_rcvd = in["ip6_tunnel_rcvd"].(int)
		ret.Ip6_tunnel_rate_limit_inner = in["ip6_tunnel_rate_limit_inner"].(int)
		ret.Ip6_tunnel_encap = in["ip6_tunnel_encap"].(int)
		ret.Ip6_tunnel_encap_fail = in["ip6_tunnel_encap_fail"].(int)
		ret.Ip6_tunnel_decap = in["ip6_tunnel_decap"].(int)
		ret.Ip6_tunnel_decap_fail = in["ip6_tunnel_decap_fail"].(int)
		ret.Ip_gre_tunnel_rcvd = in["ip_gre_tunnel_rcvd"].(int)
		ret.Ip_gre_tunnel_rate_limit_inner = in["ip_gre_tunnel_rate_limit_inner"].(int)
		ret.Ip_gre_tunnel_encap = in["ip_gre_tunnel_encap"].(int)
		ret.Ip_gre_tunnel_encap_fail = in["ip_gre_tunnel_encap_fail"].(int)
		ret.Ip_gre_tunnel_decap = in["ip_gre_tunnel_decap"].(int)
		ret.Ip_gre_tunnel_decap_fail = in["ip_gre_tunnel_decap_fail"].(int)
		ret.Ip_gre_tunnel_encap_key = in["ip_gre_tunnel_encap_key"].(int)
		ret.Ip_gre_tunnel_decap_key = in["ip_gre_tunnel_decap_key"].(int)
		ret.Ip_gre_tunnel_decap_drop = in["ip_gre_tunnel_decap_drop"].(int)
		ret.Ip6_gre_tunnel_rcvd = in["ip6_gre_tunnel_rcvd"].(int)
		ret.Ip6_gre_tunnel_rate_limit_inner = in["ip6_gre_tunnel_rate_limit_inner"].(int)
		ret.Ip6_gre_tunnel_encap = in["ip6_gre_tunnel_encap"].(int)
		ret.Ip6_gre_tunnel_encap_fail = in["ip6_gre_tunnel_encap_fail"].(int)
		ret.Ip6_gre_tunnel_decap = in["ip6_gre_tunnel_decap"].(int)
		ret.Ip6_gre_tunnel_decap_fail = in["ip6_gre_tunnel_decap_fail"].(int)
		ret.Ip6_gre_tunnel_encap_key = in["ip6_gre_tunnel_encap_key"].(int)
		ret.Ip6_gre_tunnel_decap_key = in["ip6_gre_tunnel_decap_key"].(int)
		ret.Ip6_gre_tunnel_decap_drop = in["ip6_gre_tunnel_decap_drop"].(int)
		ret.Ip_esp_tunnel_rcvd = in["ip_esp_tunnel_rcvd"].(int)
		ret.Ip_esp_tunnel_inspect = in["ip_esp_tunnel_inspect"].(int)
		ret.Ip_esp_tunnel_inspect_fail = in["ip_esp_tunnel_inspect_fail"].(int)
		ret.Ip6_esp_tunnel_rcvd = in["ip6_esp_tunnel_rcvd"].(int)
		ret.Ip6_esp_tunnel_inspect = in["ip6_esp_tunnel_inspect"].(int)
		ret.Ip6_esp_tunnel_inspect_fail = in["ip6_esp_tunnel_inspect_fail"].(int)
		ret.Ip_vxlan_tunnel_rcvd = in["ip_vxlan_tunnel_rcvd"].(int)
		ret.Ip_vxlan_tunnel_invalid_vni = in["ip_vxlan_tunnel_invalid_vni"].(int)
		ret.Ip_vxlan_tunnel_decap = in["ip_vxlan_tunnel_decap"].(int)
		ret.Ip_vxlan_tunnel_decap_err = in["ip_vxlan_tunnel_decap_err"].(int)
		ret.Ip_gre_tunnel_keepalive_rcvd = in["ip_gre_tunnel_keepalive_rcvd"].(int)
		ret.Jumbo_in_tunnel_drop = in["jumbo_in_tunnel_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosTunnelStats(d *schema.ResourceData) edpt.DdosTunnelStats {
	var ret edpt.DdosTunnelStats

	ret.Stats = getObjectDdosTunnelStatsStats(d.Get("stats").([]interface{}))
	return ret
}
