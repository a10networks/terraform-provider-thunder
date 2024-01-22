package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPortInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_port_info_oper`: Operational Status for the object port-info\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemPortInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_port_info": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_num": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pci_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dev": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"info": {
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

func resourceSystemPortInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPortInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemPortInfoOperOper := setObjectSystemPortInfoOperOper(res)
		d.Set("oper", SystemPortInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemPortInfoOperOper(ret edpt.DataSystemPortInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"system_port_info": setSliceSystemPortInfoOperOperSystemPortInfo(ret.DtSystemPortInfoOper.Oper.SystemPortInfo),
		},
	}
}

func setSliceSystemPortInfoOperOperSystemPortInfo(d []edpt.SystemPortInfoOperOperSystemPortInfo) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["pci_addr"] = item.PciAddr
		in["dev"] = item.Dev
		in["info"] = item.Info
		result = append(result, in)
	}
	return result
}

func getObjectSystemPortInfoOperOper(d []interface{}) edpt.SystemPortInfoOperOper {

	count1 := len(d)
	var ret edpt.SystemPortInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemPortInfo = getSliceSystemPortInfoOperOperSystemPortInfo(in["system_port_info"].([]interface{}))
	}
	return ret
}

func getSliceSystemPortInfoOperOperSystemPortInfo(d []interface{}) []edpt.SystemPortInfoOperOperSystemPortInfo {

	count1 := len(d)
	ret := make([]edpt.SystemPortInfoOperOperSystemPortInfo, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemPortInfoOperOperSystemPortInfo
		oi.PortNum = in["port_num"].(string)
		oi.PciAddr = in["pci_addr"].(string)
		oi.Dev = in["dev"].(string)
		oi.Info = in["info"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemPortInfoOper(d *schema.ResourceData) edpt.SystemPortInfoOper {
	var ret edpt.SystemPortInfoOper

	ret.Oper = getObjectSystemPortInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
