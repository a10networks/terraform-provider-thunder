package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServicePortOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_port_oper`: Operational Status for the object service-port\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServicePortOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_port_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_port_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"attributes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"active_real_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"current_connections": {
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

func resourceGslbServicePortOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServicePortOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServicePortOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServicePortOperOper := setObjectGslbServicePortOperOper(res)
		d.Set("oper", GslbServicePortOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbServicePortOperOper(ret edpt.DataGslbServicePortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"service_port_list": setSliceGslbServicePortOperOperServicePortList(ret.DtGslbServicePortOper.Oper.ServicePortList),
		},
	}
}

func setSliceGslbServicePortOperOperServicePortList(d []edpt.GslbServicePortOperOperServicePortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["service_port_name"] = item.ServicePortName
		in["attributes"] = item.Attributes
		in["state"] = item.State
		in["active_real_server"] = item.ActiveRealServer
		in["current_connections"] = item.CurrentConnections
		result = append(result, in)
	}
	return result
}

func getObjectGslbServicePortOperOper(d []interface{}) edpt.GslbServicePortOperOper {

	count1 := len(d)
	var ret edpt.GslbServicePortOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServicePortList = getSliceGslbServicePortOperOperServicePortList(in["service_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbServicePortOperOperServicePortList(d []interface{}) []edpt.GslbServicePortOperOperServicePortList {

	count1 := len(d)
	ret := make([]edpt.GslbServicePortOperOperServicePortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServicePortOperOperServicePortList
		oi.ServicePortName = in["service_port_name"].(string)
		oi.Attributes = in["attributes"].(string)
		oi.State = in["state"].(string)
		oi.ActiveRealServer = in["active_real_server"].(int)
		oi.CurrentConnections = in["current_connections"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbServicePortOper(d *schema.ResourceData) edpt.GslbServicePortOper {
	var ret edpt.GslbServicePortOper

	ret.Oper = getObjectGslbServicePortOperOper(d.Get("oper").([]interface{}))
	return ret
}
