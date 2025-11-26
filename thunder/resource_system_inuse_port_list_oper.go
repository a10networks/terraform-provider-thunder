package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemInusePortListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_inuse_port_list_oper`: Operational Status for the object inuse-port-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemInusePortListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_inuse_port_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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

func resourceSystemInusePortListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemInusePortListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemInusePortListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemInusePortListOperOper := setObjectSystemInusePortListOperOper(res)
		d.Set("oper", SystemInusePortListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemInusePortListOperOper(ret edpt.DataSystemInusePortListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"system_inuse_port_list": setSliceSystemInusePortListOperOperSystemInusePortList(ret.DtSystemInusePortListOper.Oper.SystemInusePortList),
		},
	}
}

func setSliceSystemInusePortListOperOperSystemInusePortList(d []edpt.SystemInusePortListOperOperSystemInusePortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
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

func getObjectSystemInusePortListOperOper(d []interface{}) edpt.SystemInusePortListOperOper {

	count1 := len(d)
	var ret edpt.SystemInusePortListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemInusePortList = getSliceSystemInusePortListOperOperSystemInusePortList(in["system_inuse_port_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemInusePortListOperOperSystemInusePortList(d []interface{}) []edpt.SystemInusePortListOperOperSystemInusePortList {

	count1 := len(d)
	ret := make([]edpt.SystemInusePortListOperOperSystemInusePortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemInusePortListOperOperSystemInusePortList
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

func dataToEndpointSystemInusePortListOper(d *schema.ResourceData) edpt.SystemInusePortListOper {
	var ret edpt.SystemInusePortListOper

	ret.Oper = getObjectSystemInusePortListOperOper(d.Get("oper").([]interface{}))
	return ret
}
