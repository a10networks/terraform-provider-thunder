package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireGlobalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_virtual_wire_global_oper`: Operational Status for the object virtual-wire-global\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVirtualWireGlobalOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_group": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_member": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"vlan_set": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"set_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_pair": {
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

func resourceNetworkVirtualWireGlobalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireGlobalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireGlobalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVirtualWireGlobalOperOper := setObjectNetworkVirtualWireGlobalOperOper(res)
		d.Set("oper", NetworkVirtualWireGlobalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVirtualWireGlobalOperOper(ret edpt.DataNetworkVirtualWireGlobalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vlan_group": setSliceNetworkVirtualWireGlobalOperOperVlan_group(ret.DtNetworkVirtualWireGlobalOper.Oper.Vlan_group),
			"vlan_set":   setSliceNetworkVirtualWireGlobalOperOperVlan_set(ret.DtNetworkVirtualWireGlobalOper.Oper.Vlan_set),
		},
	}
}

func setSliceNetworkVirtualWireGlobalOperOperVlan_group(d []edpt.NetworkVirtualWireGlobalOperOperVlan_group) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["group_id"] = item.Group_id
		in["active_member"] = item.Active_member
		result = append(result, in)
	}
	return result
}

func setSliceNetworkVirtualWireGlobalOperOperVlan_set(d []edpt.NetworkVirtualWireGlobalOperOperVlan_set) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["set_id"] = item.Set_id
		in["active_pair"] = item.Active_pair
		result = append(result, in)
	}
	return result
}

func getObjectNetworkVirtualWireGlobalOperOper(d []interface{}) edpt.NetworkVirtualWireGlobalOperOper {

	count1 := len(d)
	var ret edpt.NetworkVirtualWireGlobalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vlan_group = getSliceNetworkVirtualWireGlobalOperOperVlan_group(in["vlan_group"].([]interface{}))
		ret.Vlan_set = getSliceNetworkVirtualWireGlobalOperOperVlan_set(in["vlan_set"].([]interface{}))
	}
	return ret
}

func getSliceNetworkVirtualWireGlobalOperOperVlan_group(d []interface{}) []edpt.NetworkVirtualWireGlobalOperOperVlan_group {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireGlobalOperOperVlan_group, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireGlobalOperOperVlan_group
		oi.Group_id = in["group_id"].(int)
		oi.Active_member = in["active_member"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVirtualWireGlobalOperOperVlan_set(d []interface{}) []edpt.NetworkVirtualWireGlobalOperOperVlan_set {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireGlobalOperOperVlan_set, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireGlobalOperOperVlan_set
		oi.Set_id = in["set_id"].(int)
		oi.Active_pair = in["active_pair"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireGlobalOper(d *schema.ResourceData) edpt.NetworkVirtualWireGlobalOper {
	var ret edpt.NetworkVirtualWireGlobalOper

	ret.Oper = getObjectNetworkVirtualWireGlobalOperOper(d.Get("oper").([]interface{}))
	return ret
}
