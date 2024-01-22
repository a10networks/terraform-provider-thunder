package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatPortMappingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_port_mapping_oper`: Operational Status for the object port-mapping\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatPortMappingOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mapping_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"tcp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"assigned_to": {
										Type: schema.TypeString, Optional: true, Description: "",
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
						"nat_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatPortMappingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatPortMappingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatPortMappingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatPortMappingOperOper := setObjectCgnv6FixedNatPortMappingOperOper(res)
		d.Set("oper", Cgnv6FixedNatPortMappingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatPortMappingOperOper(ret edpt.DataCgnv6FixedNatPortMappingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"mapping_list":   setSliceCgnv6FixedNatPortMappingOperOperMappingList(ret.DtCgnv6FixedNatPortMappingOper.Oper.MappingList),
			"partition":      ret.DtCgnv6FixedNatPortMappingOper.Oper.Partition,
			"inside_user_v4": ret.DtCgnv6FixedNatPortMappingOper.Oper.InsideUserV4,
			"inside_user_v6": ret.DtCgnv6FixedNatPortMappingOper.Oper.InsideUserV6,
			"nat_ip":         ret.DtCgnv6FixedNatPortMappingOper.Oper.NatIp,
			"nat_port":       ret.DtCgnv6FixedNatPortMappingOper.Oper.NatPort,
		},
	}
}

func setSliceCgnv6FixedNatPortMappingOperOperMappingList(d []edpt.Cgnv6FixedNatPortMappingOperOperMappingList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["nat_address"] = item.NatAddress
		in["tcp_port_start"] = item.TcpPortStart
		in["tcp_port_end"] = item.TcpPortEnd
		in["udp_port_start"] = item.UdpPortStart
		in["udp_port_end"] = item.UdpPortEnd
		in["icmp_port_start"] = item.IcmpPortStart
		in["icmp_port_end"] = item.IcmpPortEnd
		in["assigned_to"] = item.AssignedTo
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6FixedNatPortMappingOperOper(d []interface{}) edpt.Cgnv6FixedNatPortMappingOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatPortMappingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MappingList = getSliceCgnv6FixedNatPortMappingOperOperMappingList(in["mapping_list"].([]interface{}))
		ret.Partition = in["partition"].(string)
		ret.InsideUserV4 = in["inside_user_v4"].(string)
		ret.InsideUserV6 = in["inside_user_v6"].(string)
		ret.NatIp = in["nat_ip"].(string)
		ret.NatPort = in["nat_port"].(int)
	}
	return ret
}

func getSliceCgnv6FixedNatPortMappingOperOperMappingList(d []interface{}) []edpt.Cgnv6FixedNatPortMappingOperOperMappingList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatPortMappingOperOperMappingList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatPortMappingOperOperMappingList
		oi.NatAddress = in["nat_address"].(string)
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		oi.IcmpPortStart = in["icmp_port_start"].(int)
		oi.IcmpPortEnd = in["icmp_port_end"].(int)
		oi.AssignedTo = in["assigned_to"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatPortMappingOper(d *schema.ResourceData) edpt.Cgnv6FixedNatPortMappingOper {
	var ret edpt.Cgnv6FixedNatPortMappingOper

	ret.Oper = getObjectCgnv6FixedNatPortMappingOperOper(d.Get("oper").([]interface{}))
	return ret
}
