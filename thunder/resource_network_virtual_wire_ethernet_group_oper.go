package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireEthernetGroupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_virtual_wire_ethernet_group_oper`: Operational Status for the object virtual-wire-ethernet-group\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVirtualWireEthernetGroupOperRead,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eth_member_status": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eth_member_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"eth_member_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"trunk_member_status": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk_member_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"trunk_member_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"lead_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"group_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceNetworkVirtualWireEthernetGroupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireEthernetGroupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireEthernetGroupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVirtualWireEthernetGroupOperOper := setObjectNetworkVirtualWireEthernetGroupOperOper(res)
		d.Set("oper", NetworkVirtualWireEthernetGroupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVirtualWireEthernetGroupOperOper(ret edpt.DataNetworkVirtualWireEthernetGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"eth_member_status":   setSliceNetworkVirtualWireEthernetGroupOperOperEth_member_status(ret.DtNetworkVirtualWireEthernetGroupOper.Oper.Eth_member_status),
			"trunk_member_status": setSliceNetworkVirtualWireEthernetGroupOperOperTrunk_member_status(ret.DtNetworkVirtualWireEthernetGroupOper.Oper.Trunk_member_status),
			"lead_port":           ret.DtNetworkVirtualWireEthernetGroupOper.Oper.Lead_port,
			"group_status":        ret.DtNetworkVirtualWireEthernetGroupOper.Oper.Group_status,
		},
	}
}

func setSliceNetworkVirtualWireEthernetGroupOperOperEth_member_status(d []edpt.NetworkVirtualWireEthernetGroupOperOperEth_member_status) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["eth_member_num"] = item.Eth_member_num
		in["eth_member_status"] = item.Eth_member_status
		result = append(result, in)
	}
	return result
}

func setSliceNetworkVirtualWireEthernetGroupOperOperTrunk_member_status(d []edpt.NetworkVirtualWireEthernetGroupOperOperTrunk_member_status) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["trunk_member_num"] = item.Trunk_member_num
		in["trunk_member_status"] = item.Trunk_member_status
		result = append(result, in)
	}
	return result
}

func getObjectNetworkVirtualWireEthernetGroupOperOper(d []interface{}) edpt.NetworkVirtualWireEthernetGroupOperOper {

	count1 := len(d)
	var ret edpt.NetworkVirtualWireEthernetGroupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Eth_member_status = getSliceNetworkVirtualWireEthernetGroupOperOperEth_member_status(in["eth_member_status"].([]interface{}))
		ret.Trunk_member_status = getSliceNetworkVirtualWireEthernetGroupOperOperTrunk_member_status(in["trunk_member_status"].([]interface{}))
		ret.Lead_port = in["lead_port"].(int)
		ret.Group_status = in["group_status"].(string)
	}
	return ret
}

func getSliceNetworkVirtualWireEthernetGroupOperOperEth_member_status(d []interface{}) []edpt.NetworkVirtualWireEthernetGroupOperOperEth_member_status {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireEthernetGroupOperOperEth_member_status, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireEthernetGroupOperOperEth_member_status
		oi.Eth_member_num = in["eth_member_num"].(int)
		oi.Eth_member_status = in["eth_member_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVirtualWireEthernetGroupOperOperTrunk_member_status(d []interface{}) []edpt.NetworkVirtualWireEthernetGroupOperOperTrunk_member_status {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireEthernetGroupOperOperTrunk_member_status, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireEthernetGroupOperOperTrunk_member_status
		oi.Trunk_member_num = in["trunk_member_num"].(int)
		oi.Trunk_member_status = in["trunk_member_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireEthernetGroupOper(d *schema.ResourceData) edpt.NetworkVirtualWireEthernetGroupOper {
	var ret edpt.NetworkVirtualWireEthernetGroupOper

	ret.GroupId = d.Get("group_id").(int)

	ret.Oper = getObjectNetworkVirtualWireEthernetGroupOperOper(d.Get("oper").([]interface{}))
	return ret
}
