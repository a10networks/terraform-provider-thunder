package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatHistogramOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_histogram_oper`: Operational Status for the object histogram\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatHistogramOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside_user_v4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_user_v6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_start_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_end_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_start_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_end_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"outside_start_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"outside_end_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_users": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"histogram_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bin_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bin_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_tcp_users": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_udp_users": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatHistogramOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatHistogramOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatHistogramOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatHistogramOperOper := setObjectCgnv6FixedNatHistogramOperOper(res)
		d.Set("oper", Cgnv6FixedNatHistogramOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatHistogramOperOper(ret edpt.DataCgnv6FixedNatHistogramOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"inside_user_v4":    ret.DtCgnv6FixedNatHistogramOper.Oper.InsideUserV4,
			"inside_user_v6":    ret.DtCgnv6FixedNatHistogramOper.Oper.InsideUserV6,
			"nat_address":       ret.DtCgnv6FixedNatHistogramOper.Oper.NatAddress,
			"inside_start_ipv4": ret.DtCgnv6FixedNatHistogramOper.Oper.InsideStartIpv4,
			"inside_end_ipv4":   ret.DtCgnv6FixedNatHistogramOper.Oper.InsideEndIpv4,
			"inside_start_ipv6": ret.DtCgnv6FixedNatHistogramOper.Oper.InsideStartIpv6,
			"inside_end_ipv6":   ret.DtCgnv6FixedNatHistogramOper.Oper.InsideEndIpv6,
			"outside_start_ip":  ret.DtCgnv6FixedNatHistogramOper.Oper.OutsideStartIp,
			"outside_end_ip":    ret.DtCgnv6FixedNatHistogramOper.Oper.OutsideEndIp,
			"total_users":       ret.DtCgnv6FixedNatHistogramOper.Oper.TotalUsers,
			"histogram_list":    setSliceCgnv6FixedNatHistogramOperOperHistogramList(ret.DtCgnv6FixedNatHistogramOper.Oper.HistogramList),
		},
	}
}

func setSliceCgnv6FixedNatHistogramOperOperHistogramList(d []edpt.Cgnv6FixedNatHistogramOperOperHistogramList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["bin_start"] = item.BinStart
		in["bin_end"] = item.BinEnd
		in["active_tcp_users"] = item.ActiveTcpUsers
		in["active_udp_users"] = item.ActiveUdpUsers
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6FixedNatHistogramOperOper(d []interface{}) edpt.Cgnv6FixedNatHistogramOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatHistogramOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InsideUserV4 = in["inside_user_v4"].(string)
		ret.InsideUserV6 = in["inside_user_v6"].(string)
		ret.NatAddress = in["nat_address"].(string)
		ret.InsideStartIpv4 = in["inside_start_ipv4"].(string)
		ret.InsideEndIpv4 = in["inside_end_ipv4"].(string)
		ret.InsideStartIpv6 = in["inside_start_ipv6"].(string)
		ret.InsideEndIpv6 = in["inside_end_ipv6"].(string)
		ret.OutsideStartIp = in["outside_start_ip"].(string)
		ret.OutsideEndIp = in["outside_end_ip"].(string)
		ret.TotalUsers = in["total_users"].(int)
		ret.HistogramList = getSliceCgnv6FixedNatHistogramOperOperHistogramList(in["histogram_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6FixedNatHistogramOperOperHistogramList(d []interface{}) []edpt.Cgnv6FixedNatHistogramOperOperHistogramList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatHistogramOperOperHistogramList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatHistogramOperOperHistogramList
		oi.BinStart = in["bin_start"].(int)
		oi.BinEnd = in["bin_end"].(int)
		oi.ActiveTcpUsers = in["active_tcp_users"].(int)
		oi.ActiveUdpUsers = in["active_udp_users"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatHistogramOper(d *schema.ResourceData) edpt.Cgnv6FixedNatHistogramOper {
	var ret edpt.Cgnv6FixedNatHistogramOper

	ret.Oper = getObjectCgnv6FixedNatHistogramOperOper(d.Get("oper").([]interface{}))
	return ret
}
