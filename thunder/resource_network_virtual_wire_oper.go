package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_virtual_wire_oper`: Operational Status for the object virtual-wire\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVirtualWireOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoints": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"endpoint_intf": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"input_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"input_byte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"output_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"output_byte": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drop_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"virtual_wire_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"virtual_wire_id": {
				Type: schema.TypeInt, Required: true, Description: "virtual wire id",
			},
		},
	}
}

func resourceNetworkVirtualWireOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVirtualWireOperOper := setObjectNetworkVirtualWireOperOper(res)
		d.Set("oper", NetworkVirtualWireOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVirtualWireOperOper(ret edpt.DataNetworkVirtualWireOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"endpoints":           setSliceNetworkVirtualWireOperOperEndpoints(ret.DtNetworkVirtualWireOper.Oper.Endpoints),
			"virtual_wire_status": ret.DtNetworkVirtualWireOper.Oper.Virtual_wire_status,
		},
	}
}

func setSliceNetworkVirtualWireOperOperEndpoints(d []edpt.NetworkVirtualWireOperOperEndpoints) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["endpoint_type"] = item.Endpoint_type
		in["endpoint_intf"] = item.Endpoint_intf
		in["input_packet"] = item.Input_packet
		in["input_byte"] = item.Input_byte
		in["output_packet"] = item.Output_packet
		in["output_byte"] = item.Output_byte
		in["drop_packet"] = item.Drop_packet
		result = append(result, in)
	}
	return result
}

func getObjectNetworkVirtualWireOperOper(d []interface{}) edpt.NetworkVirtualWireOperOper {

	count1 := len(d)
	var ret edpt.NetworkVirtualWireOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Endpoints = getSliceNetworkVirtualWireOperOperEndpoints(in["endpoints"].([]interface{}))
		ret.Virtual_wire_status = in["virtual_wire_status"].(string)
	}
	return ret
}

func getSliceNetworkVirtualWireOperOperEndpoints(d []interface{}) []edpt.NetworkVirtualWireOperOperEndpoints {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireOperOperEndpoints, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireOperOperEndpoints
		oi.Endpoint_type = in["endpoint_type"].(string)
		oi.Endpoint_intf = in["endpoint_intf"].(int)
		oi.Input_packet = in["input_packet"].(int)
		oi.Input_byte = in["input_byte"].(int)
		oi.Output_packet = in["output_packet"].(int)
		oi.Output_byte = in["output_byte"].(int)
		oi.Drop_packet = in["drop_packet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireOper(d *schema.ResourceData) edpt.NetworkVirtualWireOper {
	var ret edpt.NetworkVirtualWireOper

	ret.Oper = getObjectNetworkVirtualWireOperOper(d.Get("oper").([]interface{}))

	ret.VirtualWireId = d.Get("virtual_wire_id").(int)
	return ret
}
