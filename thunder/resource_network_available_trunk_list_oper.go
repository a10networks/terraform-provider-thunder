package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAvailableTrunkListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_available_trunk_list_oper`: Operational Status for the object available-trunk-list\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkAvailableTrunkListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"if_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"if_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"if_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"if_status": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceNetworkAvailableTrunkListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkAvailableTrunkListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkAvailableTrunkListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkAvailableTrunkListOperOper := setObjectNetworkAvailableTrunkListOperOper(res)
		d.Set("oper", NetworkAvailableTrunkListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkAvailableTrunkListOperOper(ret edpt.DataNetworkAvailableTrunkListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"if_list": setSliceNetworkAvailableTrunkListOperOperIfList(ret.DtNetworkAvailableTrunkListOper.Oper.IfList),
		},
	}
}

func setSliceNetworkAvailableTrunkListOperOperIfList(d []edpt.NetworkAvailableTrunkListOperOperIfList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["if_type"] = item.IfType
		in["if_num"] = item.IfNum
		in["if_status"] = item.IfStatus
		result = append(result, in)
	}
	return result
}

func getObjectNetworkAvailableTrunkListOperOper(d []interface{}) edpt.NetworkAvailableTrunkListOperOper {

	count1 := len(d)
	var ret edpt.NetworkAvailableTrunkListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IfList = getSliceNetworkAvailableTrunkListOperOperIfList(in["if_list"].([]interface{}))
	}
	return ret
}

func getSliceNetworkAvailableTrunkListOperOperIfList(d []interface{}) []edpt.NetworkAvailableTrunkListOperOperIfList {

	count1 := len(d)
	ret := make([]edpt.NetworkAvailableTrunkListOperOperIfList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkAvailableTrunkListOperOperIfList
		oi.IfType = in["if_type"].(string)
		oi.IfNum = in["if_num"].(int)
		oi.IfStatus = in["if_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkAvailableTrunkListOper(d *schema.ResourceData) edpt.NetworkAvailableTrunkListOper {
	var ret edpt.NetworkAvailableTrunkListOper

	ret.Oper = getObjectNetworkAvailableTrunkListOperOper(d.Get("oper").([]interface{}))
	return ret
}
