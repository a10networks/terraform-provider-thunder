package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatPoolOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat_pool_oper`: Operational Status for the object pool\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6NatPoolOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"users": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_reserved": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_peak": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_hit_full": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_reserved": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_peak": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_hit_full": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_port_overloaded": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_reserved": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_peak": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_hit_full": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_port_overloaded": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtsp_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"obsoleted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name",
			},
		},
	}
}

func resourceCgnv6NatPoolOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPoolOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6NatPoolOperOper := setObjectCgnv6NatPoolOperOper(res)
		d.Set("oper", Cgnv6NatPoolOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6NatPoolOperOper(ret edpt.DataCgnv6NatPoolOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"nat_ip_list": setSliceCgnv6NatPoolOperOperNatIpList(ret.DtCgnv6NatPoolOper.Oper.NatIpList),
		},
	}
}

func setSliceCgnv6NatPoolOperOperNatIpList(d []edpt.Cgnv6NatPoolOperOperNatIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_address"] = item.IpAddress
		in["users"] = item.Users
		in["icmp_used"] = item.IcmpUsed
		in["icmp_freed"] = item.IcmpFreed
		in["icmp_total"] = item.IcmpTotal
		in["icmp_reserved"] = item.IcmpReserved
		in["icmp_peak"] = item.IcmpPeak
		in["icmp_hit_full"] = item.IcmpHitFull
		in["udp_used"] = item.UdpUsed
		in["udp_freed"] = item.UdpFreed
		in["udp_total"] = item.UdpTotal
		in["udp_reserved"] = item.UdpReserved
		in["udp_peak"] = item.UdpPeak
		in["udp_hit_full"] = item.UdpHitFull
		in["udp_port_overloaded"] = item.UdpPortOverloaded
		in["tcp_used"] = item.TcpUsed
		in["tcp_freed"] = item.TcpFreed
		in["tcp_total"] = item.TcpTotal
		in["tcp_reserved"] = item.TcpReserved
		in["tcp_peak"] = item.TcpPeak
		in["tcp_hit_full"] = item.TcpHitFull
		in["tcp_port_overloaded"] = item.TcpPortOverloaded
		in["rtsp_used"] = item.RtspUsed
		in["obsoleted"] = item.Obsoleted
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6NatPoolOperOper(d []interface{}) edpt.Cgnv6NatPoolOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6NatPoolOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatIpList = getSliceCgnv6NatPoolOperOperNatIpList(in["nat_ip_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6NatPoolOperOperNatIpList(d []interface{}) []edpt.Cgnv6NatPoolOperOperNatIpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatPoolOperOperNatIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatPoolOperOperNatIpList
		oi.IpAddress = in["ip_address"].(string)
		oi.Users = in["users"].(int)
		oi.IcmpUsed = in["icmp_used"].(int)
		oi.IcmpFreed = in["icmp_freed"].(int)
		oi.IcmpTotal = in["icmp_total"].(int)
		oi.IcmpReserved = in["icmp_reserved"].(int)
		oi.IcmpPeak = in["icmp_peak"].(int)
		oi.IcmpHitFull = in["icmp_hit_full"].(int)
		oi.UdpUsed = in["udp_used"].(int)
		oi.UdpFreed = in["udp_freed"].(int)
		oi.UdpTotal = in["udp_total"].(int)
		oi.UdpReserved = in["udp_reserved"].(int)
		oi.UdpPeak = in["udp_peak"].(int)
		oi.UdpHitFull = in["udp_hit_full"].(int)
		oi.UdpPortOverloaded = in["udp_port_overloaded"].(int)
		oi.TcpUsed = in["tcp_used"].(int)
		oi.TcpFreed = in["tcp_freed"].(int)
		oi.TcpTotal = in["tcp_total"].(int)
		oi.TcpReserved = in["tcp_reserved"].(int)
		oi.TcpPeak = in["tcp_peak"].(int)
		oi.TcpHitFull = in["tcp_hit_full"].(int)
		oi.TcpPortOverloaded = in["tcp_port_overloaded"].(int)
		oi.RtspUsed = in["rtsp_used"].(int)
		oi.Obsoleted = in["obsoleted"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6NatPoolOper(d *schema.ResourceData) edpt.Cgnv6NatPoolOper {
	var ret edpt.Cgnv6NatPoolOper

	ret.Oper = getObjectCgnv6NatPoolOperOper(d.Get("oper").([]interface{}))

	ret.PoolName = d.Get("pool_name").(string)
	return ret
}
