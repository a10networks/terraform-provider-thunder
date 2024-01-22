package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceAvailableEthListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_available_eth_list_oper`: Operational Status for the object available-eth-list\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceAvailableEthListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tot_num_of_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
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
									"state": {
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

func resourceInterfaceAvailableEthListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceAvailableEthListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceAvailableEthListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceAvailableEthListOperOper := setObjectInterfaceAvailableEthListOperOper(res)
		d.Set("oper", InterfaceAvailableEthListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceAvailableEthListOperOper(ret edpt.DataInterfaceAvailableEthListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tot_num_of_ports": ret.DtInterfaceAvailableEthListOper.Oper.Tot_num_of_ports,
			"if_list":          setSliceInterfaceAvailableEthListOperOperIfList(ret.DtInterfaceAvailableEthListOper.Oper.IfList),
		},
	}
}

func setSliceInterfaceAvailableEthListOperOperIfList(d []edpt.InterfaceAvailableEthListOperOperIfList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["if_type"] = item.IfType
		in["if_num"] = item.IfNum
		in["if_status"] = item.IfStatus
		in["state"] = item.State
		result = append(result, in)
	}
	return result
}

func getObjectInterfaceAvailableEthListOperOper(d []interface{}) edpt.InterfaceAvailableEthListOperOper {

	count1 := len(d)
	var ret edpt.InterfaceAvailableEthListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tot_num_of_ports = in["tot_num_of_ports"].(int)
		ret.IfList = getSliceInterfaceAvailableEthListOperOperIfList(in["if_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceAvailableEthListOperOperIfList(d []interface{}) []edpt.InterfaceAvailableEthListOperOperIfList {

	count1 := len(d)
	ret := make([]edpt.InterfaceAvailableEthListOperOperIfList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceAvailableEthListOperOperIfList
		oi.IfType = in["if_type"].(string)
		oi.IfNum = in["if_num"].(int)
		oi.IfStatus = in["if_status"].(string)
		oi.State = in["state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceAvailableEthListOper(d *schema.ResourceData) edpt.InterfaceAvailableEthListOper {
	var ret edpt.InterfaceAvailableEthListOper

	ret.Oper = getObjectInterfaceAvailableEthListOperOper(d.Get("oper").([]interface{}))
	return ret
}
