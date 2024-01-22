package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatPoolStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat_pool_stats`: Statistics for the object pool\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6NatPoolStatsRead,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type: schema.TypeInt, Optional: true, Description: "Users",
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP",
						},
						"icmp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Freed",
						},
						"icmp_total": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total",
						},
						"icmp_rsvd": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Reserved",
						},
						"icmp_peak": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Peak",
						},
						"icmp_hit_full": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Hit Full",
						},
						"udp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP",
						},
						"udp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Freed",
						},
						"udp_total": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total",
						},
						"udp_rsvd": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Reserved",
						},
						"udp_peak": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Peak",
						},
						"udp_hit_full": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Hit Full",
						},
						"udp_port_overloaded": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloaded",
						},
						"udp_port_overload_create": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloading Session Created",
						},
						"udp_port_overload_free": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloading Session Freed",
						},
						"tcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP",
						},
						"tcp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Freed",
						},
						"tcp_total": {
							Type: schema.TypeInt, Optional: true, Description: "TCP total",
						},
						"tcp_rsvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Reserved",
						},
						"tcp_peak": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Peak",
						},
						"tcp_hit_full": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Hit Full",
						},
						"tcp_port_overloaded": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloaded",
						},
						"tcp_port_overload_create": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloading Session Created",
						},
						"tcp_port_overload_free": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloading Session Freed",
						},
						"ip_used": {
							Type: schema.TypeInt, Optional: true, Description: "IP Used",
						},
						"ip_free": {
							Type: schema.TypeInt, Optional: true, Description: "IP Free",
						},
						"ip_total": {
							Type: schema.TypeInt, Optional: true, Description: "IP Total",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6NatPoolStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6NatPoolStatsStats := setObjectCgnv6NatPoolStatsStats(res)
		d.Set("stats", Cgnv6NatPoolStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6NatPoolStatsStats(ret edpt.DataCgnv6NatPoolStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"users":                    ret.DtCgnv6NatPoolStats.Stats.Users,
			"icmp":                     ret.DtCgnv6NatPoolStats.Stats.Icmp,
			"icmp_freed":               ret.DtCgnv6NatPoolStats.Stats.IcmpFreed,
			"icmp_total":               ret.DtCgnv6NatPoolStats.Stats.IcmpTotal,
			"icmp_rsvd":                ret.DtCgnv6NatPoolStats.Stats.IcmpRsvd,
			"icmp_peak":                ret.DtCgnv6NatPoolStats.Stats.IcmpPeak,
			"icmp_hit_full":            ret.DtCgnv6NatPoolStats.Stats.IcmpHitFull,
			"udp":                      ret.DtCgnv6NatPoolStats.Stats.Udp,
			"udp_freed":                ret.DtCgnv6NatPoolStats.Stats.UdpFreed,
			"udp_total":                ret.DtCgnv6NatPoolStats.Stats.UdpTotal,
			"udp_rsvd":                 ret.DtCgnv6NatPoolStats.Stats.UdpRsvd,
			"udp_peak":                 ret.DtCgnv6NatPoolStats.Stats.UdpPeak,
			"udp_hit_full":             ret.DtCgnv6NatPoolStats.Stats.UdpHitFull,
			"udp_port_overloaded":      ret.DtCgnv6NatPoolStats.Stats.UdpPortOverloaded,
			"udp_port_overload_create": ret.DtCgnv6NatPoolStats.Stats.UdpPortOverloadCreate,
			"udp_port_overload_free":   ret.DtCgnv6NatPoolStats.Stats.UdpPortOverloadFree,
			"tcp":                      ret.DtCgnv6NatPoolStats.Stats.Tcp,
			"tcp_freed":                ret.DtCgnv6NatPoolStats.Stats.TcpFreed,
			"tcp_total":                ret.DtCgnv6NatPoolStats.Stats.TcpTotal,
			"tcp_rsvd":                 ret.DtCgnv6NatPoolStats.Stats.TcpRsvd,
			"tcp_peak":                 ret.DtCgnv6NatPoolStats.Stats.TcpPeak,
			"tcp_hit_full":             ret.DtCgnv6NatPoolStats.Stats.TcpHitFull,
			"tcp_port_overloaded":      ret.DtCgnv6NatPoolStats.Stats.TcpPortOverloaded,
			"tcp_port_overload_create": ret.DtCgnv6NatPoolStats.Stats.TcpPortOverloadCreate,
			"tcp_port_overload_free":   ret.DtCgnv6NatPoolStats.Stats.TcpPortOverloadFree,
			"ip_used":                  ret.DtCgnv6NatPoolStats.Stats.IpUsed,
			"ip_free":                  ret.DtCgnv6NatPoolStats.Stats.IpFree,
			"ip_total":                 ret.DtCgnv6NatPoolStats.Stats.IpTotal,
		},
	}
}

func getObjectCgnv6NatPoolStatsStats(d []interface{}) edpt.Cgnv6NatPoolStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6NatPoolStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Users = in["users"].(int)
		ret.Icmp = in["icmp"].(int)
		ret.IcmpFreed = in["icmp_freed"].(int)
		ret.IcmpTotal = in["icmp_total"].(int)
		ret.IcmpRsvd = in["icmp_rsvd"].(int)
		ret.IcmpPeak = in["icmp_peak"].(int)
		ret.IcmpHitFull = in["icmp_hit_full"].(int)
		ret.Udp = in["udp"].(int)
		ret.UdpFreed = in["udp_freed"].(int)
		ret.UdpTotal = in["udp_total"].(int)
		ret.UdpRsvd = in["udp_rsvd"].(int)
		ret.UdpPeak = in["udp_peak"].(int)
		ret.UdpHitFull = in["udp_hit_full"].(int)
		ret.UdpPortOverloaded = in["udp_port_overloaded"].(int)
		ret.UdpPortOverloadCreate = in["udp_port_overload_create"].(int)
		ret.UdpPortOverloadFree = in["udp_port_overload_free"].(int)
		ret.Tcp = in["tcp"].(int)
		ret.TcpFreed = in["tcp_freed"].(int)
		ret.TcpTotal = in["tcp_total"].(int)
		ret.TcpRsvd = in["tcp_rsvd"].(int)
		ret.TcpPeak = in["tcp_peak"].(int)
		ret.TcpHitFull = in["tcp_hit_full"].(int)
		ret.TcpPortOverloaded = in["tcp_port_overloaded"].(int)
		ret.TcpPortOverloadCreate = in["tcp_port_overload_create"].(int)
		ret.TcpPortOverloadFree = in["tcp_port_overload_free"].(int)
		ret.IpUsed = in["ip_used"].(int)
		ret.IpFree = in["ip_free"].(int)
		ret.IpTotal = in["ip_total"].(int)
	}
	return ret
}

func dataToEndpointCgnv6NatPoolStats(d *schema.ResourceData) edpt.Cgnv6NatPoolStats {
	var ret edpt.Cgnv6NatPoolStats

	ret.PoolName = d.Get("pool_name").(string)

	ret.Stats = getObjectCgnv6NatPoolStatsStats(d.Get("stats").([]interface{}))
	return ret
}
