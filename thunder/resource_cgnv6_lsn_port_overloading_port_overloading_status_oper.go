package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPortOverloadingPortOverloadingStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_port_overloading_port_overloading_status_oper`: Operational Status for the object port-overloading-status\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnPortOverloadingPortOverloadingStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_port_tcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"end_port_tcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_tcp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"udp_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_port_udp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"end_port_udp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_udp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"unique": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnPortOverloadingPortOverloadingStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingPortOverloadingStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingPortOverloadingStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnPortOverloadingPortOverloadingStatusOperOper := setObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOper(res)
		d.Set("oper", Cgnv6LsnPortOverloadingPortOverloadingStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOper(ret edpt.DataCgnv6LsnPortOverloadingPortOverloadingStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_list": setObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList(ret.DtCgnv6LsnPortOverloadingPortOverloadingStatusOper.Oper.TcpList),
			"udp_list": setObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList(ret.DtCgnv6LsnPortOverloadingPortOverloadingStatusOper.Oper.UdpList),
			"unique":   ret.DtCgnv6LsnPortOverloadingPortOverloadingStatusOper.Oper.Unique,
		},
	}
}

func setObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList(d edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["start_port_tcp"] = d.StartPortTcp

	in["end_port_tcp"] = d.EndPortTcp

	in["status_tcp"] = d.StatusTcp
	result = append(result, in)
	return result
}

func setObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList(d edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["start_port_udp"] = d.StartPortUdp

	in["end_port_udp"] = d.EndPortUdp

	in["status_udp"] = d.StatusUdp
	result = append(result, in)
	return result
}

func getObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOper(d []interface{}) edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpList = getObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList(in["tcp_list"].([]interface{}))
		ret.UdpList = getObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList(in["udp_list"].([]interface{}))
		ret.Unique = in["unique"].(string)
	}
	return ret
}

func getObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList(d []interface{}) edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList {

	count1 := len(d)
	var ret edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartPortTcp = in["start_port_tcp"].(int)
		ret.EndPortTcp = in["end_port_tcp"].(int)
		ret.StatusTcp = in["status_tcp"].(string)
	}
	return ret
}

func getObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList(d []interface{}) edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList {

	count1 := len(d)
	var ret edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartPortUdp = in["start_port_udp"].(int)
		ret.EndPortUdp = in["end_port_udp"].(int)
		ret.StatusUdp = in["status_udp"].(string)
	}
	return ret
}

func dataToEndpointCgnv6LsnPortOverloadingPortOverloadingStatusOper(d *schema.ResourceData) edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOper {
	var ret edpt.Cgnv6LsnPortOverloadingPortOverloadingStatusOper

	ret.Oper = getObjectCgnv6LsnPortOverloadingPortOverloadingStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
