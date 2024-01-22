package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVlanGlobalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_vlan_global_oper`: Operational Status for the object vlan-global\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVlanGlobalOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_trans_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan": {
										Type: schema.TypeInt, Optional: true, Description: "VLAN id",
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

func resourceNetworkVlanGlobalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGlobalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGlobalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVlanGlobalOperOper := setObjectNetworkVlanGlobalOperOper(res)
		d.Set("oper", NetworkVlanGlobalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVlanGlobalOperOper(ret edpt.DataNetworkVlanGlobalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vlan_trans_list": setObjectNetworkVlanGlobalOperOperVlanTransList(ret.DtNetworkVlanGlobalOper.Oper.VlanTransList),
		},
	}
}

func setObjectNetworkVlanGlobalOperOperVlanTransList(d edpt.NetworkVlanGlobalOperOperVlanTransList) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["vlan"] = d.Vlan
	result = append(result, in)
	return result
}

func getObjectNetworkVlanGlobalOperOper(d []interface{}) edpt.NetworkVlanGlobalOperOper {

	count1 := len(d)
	var ret edpt.NetworkVlanGlobalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VlanTransList = getObjectNetworkVlanGlobalOperOperVlanTransList(in["vlan_trans_list"].([]interface{}))
	}
	return ret
}

func getObjectNetworkVlanGlobalOperOperVlanTransList(d []interface{}) edpt.NetworkVlanGlobalOperOperVlanTransList {

	count1 := len(d)
	var ret edpt.NetworkVlanGlobalOperOperVlanTransList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vlan = in["vlan"].(int)
	}
	return ret
}

func dataToEndpointNetworkVlanGlobalOper(d *schema.ResourceData) edpt.NetworkVlanGlobalOper {
	var ret edpt.NetworkVlanGlobalOper

	ret.Oper = getObjectNetworkVlanGlobalOperOper(d.Get("oper").([]interface{}))
	return ret
}
