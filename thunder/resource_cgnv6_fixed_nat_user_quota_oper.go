package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatUserQuotaOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_user_quota_oper`: Operational Status for the object user-quota\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatUserQuotaOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quota_usage_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside_user": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"session_quota_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_ports_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_ports_available": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_ports_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_ports_available": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_resources_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_resources_available": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_user_v4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_user_v6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatUserQuotaOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatUserQuotaOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatUserQuotaOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatUserQuotaOperOper := setObjectCgnv6FixedNatUserQuotaOperOper(res)
		d.Set("oper", Cgnv6FixedNatUserQuotaOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatUserQuotaOperOper(ret edpt.DataCgnv6FixedNatUserQuotaOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"quota_usage_list": setSliceCgnv6FixedNatUserQuotaOperOperQuotaUsageList(ret.DtCgnv6FixedNatUserQuotaOper.Oper.QuotaUsageList),
			"partition":        ret.DtCgnv6FixedNatUserQuotaOper.Oper.Partition,
			"inside_user_v4":   ret.DtCgnv6FixedNatUserQuotaOper.Oper.InsideUserV4,
			"inside_user_v6":   ret.DtCgnv6FixedNatUserQuotaOper.Oper.InsideUserV6,
		},
	}
}

func setSliceCgnv6FixedNatUserQuotaOperOperQuotaUsageList(d []edpt.Cgnv6FixedNatUserQuotaOperOperQuotaUsageList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_user"] = item.InsideUser
		in["nat_address"] = item.NatAddress
		in["session_quota_used"] = item.SessionQuotaUsed
		in["tcp_ports_used"] = item.TcpPortsUsed
		in["tcp_ports_available"] = item.TcpPortsAvailable
		in["udp_ports_used"] = item.UdpPortsUsed
		in["udp_ports_available"] = item.UdpPortsAvailable
		in["icmp_resources_used"] = item.IcmpResourcesUsed
		in["icmp_resources_available"] = item.IcmpResourcesAvailable
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6FixedNatUserQuotaOperOper(d []interface{}) edpt.Cgnv6FixedNatUserQuotaOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatUserQuotaOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QuotaUsageList = getSliceCgnv6FixedNatUserQuotaOperOperQuotaUsageList(in["quota_usage_list"].([]interface{}))
		ret.Partition = in["partition"].(string)
		ret.InsideUserV4 = in["inside_user_v4"].(string)
		ret.InsideUserV6 = in["inside_user_v6"].(string)
	}
	return ret
}

func getSliceCgnv6FixedNatUserQuotaOperOperQuotaUsageList(d []interface{}) []edpt.Cgnv6FixedNatUserQuotaOperOperQuotaUsageList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatUserQuotaOperOperQuotaUsageList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatUserQuotaOperOperQuotaUsageList
		oi.InsideUser = in["inside_user"].(string)
		oi.NatAddress = in["nat_address"].(string)
		oi.SessionQuotaUsed = in["session_quota_used"].(int)
		oi.TcpPortsUsed = in["tcp_ports_used"].(int)
		oi.TcpPortsAvailable = in["tcp_ports_available"].(int)
		oi.UdpPortsUsed = in["udp_ports_used"].(int)
		oi.UdpPortsAvailable = in["udp_ports_available"].(int)
		oi.IcmpResourcesUsed = in["icmp_resources_used"].(int)
		oi.IcmpResourcesAvailable = in["icmp_resources_available"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatUserQuotaOper(d *schema.ResourceData) edpt.Cgnv6FixedNatUserQuotaOper {
	var ret edpt.Cgnv6FixedNatUserQuotaOper

	ret.Oper = getObjectCgnv6FixedNatUserQuotaOperOper(d.Get("oper").([]interface{}))
	return ret
}
