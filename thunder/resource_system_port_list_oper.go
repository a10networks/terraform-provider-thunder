package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPortListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_port_list_oper`: Operational Status for the object port-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemPortListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_port_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_idx": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_num": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mac_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"speed": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"node": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pci_addr": {
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

func resourceSystemPortListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPortListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemPortListOperOper := setObjectSystemPortListOperOper(res)
		d.Set("oper", SystemPortListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemPortListOperOper(ret edpt.DataSystemPortListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"system_port_list": setSliceSystemPortListOperOperSystemPortList(ret.DtSystemPortListOper.Oper.SystemPortList),
		},
	}
}

func setSliceSystemPortListOperOperSystemPortList(d []edpt.SystemPortListOperOperSystemPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_idx"] = item.PortIdx
		in["port_num"] = item.PortNum
		in["status"] = item.Status
		in["mac_addr"] = item.MacAddr
		in["speed"] = item.Speed
		in["node"] = item.Node
		in["pci_addr"] = item.PciAddr
		result = append(result, in)
	}
	return result
}

func getObjectSystemPortListOperOper(d []interface{}) edpt.SystemPortListOperOper {

	count1 := len(d)
	var ret edpt.SystemPortListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemPortList = getSliceSystemPortListOperOperSystemPortList(in["system_port_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemPortListOperOperSystemPortList(d []interface{}) []edpt.SystemPortListOperOperSystemPortList {

	count1 := len(d)
	ret := make([]edpt.SystemPortListOperOperSystemPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemPortListOperOperSystemPortList
		oi.PortIdx = in["port_idx"].(string)
		oi.PortNum = in["port_num"].(string)
		oi.Status = in["status"].(string)
		oi.MacAddr = in["mac_addr"].(string)
		oi.Speed = in["speed"].(string)
		oi.Node = in["node"].(string)
		oi.PciAddr = in["pci_addr"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemPortListOper(d *schema.ResourceData) edpt.SystemPortListOper {
	var ret edpt.SystemPortListOper

	ret.Oper = getObjectSystemPortListOperOper(d.Get("oper").([]interface{}))
	return ret
}
