package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceSwitchDropReasonOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_switch_drop_reason_oper`: Operational Status for the object switch-drop-reason\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceSwitchDropReasonOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interfaces": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"if_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinconddrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinvlandrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinbitmapzerodrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinfilterdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinportdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinipv4drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinipv6drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinmcdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinqueuedrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinerrordrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifineventdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinundersizedrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinmtuexceeddrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinfragdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinjabberdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifinfcsdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifoutipv4drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifoutipv6drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifouttnldrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifoutmcdrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifoutvlandrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifoutconddrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifoutparitydrop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ifoutstgdrop": {
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

func resourceInterfaceSwitchDropReasonOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceSwitchDropReasonOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceSwitchDropReasonOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceSwitchDropReasonOperOper := setObjectInterfaceSwitchDropReasonOperOper(res)
		d.Set("oper", InterfaceSwitchDropReasonOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceSwitchDropReasonOperOper(ret edpt.DataInterfaceSwitchDropReasonOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"interfaces": setSliceInterfaceSwitchDropReasonOperOperInterfaces(ret.DtInterfaceSwitchDropReasonOper.Oper.Interfaces),
		},
	}
}

func setSliceInterfaceSwitchDropReasonOperOperInterfaces(d []edpt.InterfaceSwitchDropReasonOperOperInterfaces) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["if_type"] = item.IfType
		in["port_num"] = item.Port_num
		in["ifinconddrop"] = item.Ifinconddrop
		in["ifinvlandrop"] = item.Ifinvlandrop
		in["ifinbitmapzerodrop"] = item.Ifinbitmapzerodrop
		in["ifinfilterdrop"] = item.Ifinfilterdrop
		in["ifinportdrop"] = item.Ifinportdrop
		in["ifinipv4drop"] = item.Ifinipv4drop
		in["ifinipv6drop"] = item.Ifinipv6drop
		in["ifinmcdrop"] = item.Ifinmcdrop
		in["ifinqueuedrop"] = item.Ifinqueuedrop
		in["ifinerrordrop"] = item.Ifinerrordrop
		in["ifineventdrop"] = item.Ifineventdrop
		in["ifinundersizedrop"] = item.Ifinundersizedrop
		in["ifinmtuexceeddrop"] = item.Ifinmtuexceeddrop
		in["ifinfragdrop"] = item.Ifinfragdrop
		in["ifinjabberdrop"] = item.Ifinjabberdrop
		in["ifinfcsdrop"] = item.Ifinfcsdrop
		in["ifoutipv4drop"] = item.Ifoutipv4drop
		in["ifoutipv6drop"] = item.Ifoutipv6drop
		in["ifouttnldrop"] = item.Ifouttnldrop
		in["ifoutmcdrop"] = item.Ifoutmcdrop
		in["ifoutvlandrop"] = item.Ifoutvlandrop
		in["ifoutconddrop"] = item.Ifoutconddrop
		in["ifoutparitydrop"] = item.Ifoutparitydrop
		in["ifoutstgdrop"] = item.Ifoutstgdrop
		result = append(result, in)
	}
	return result
}

func getObjectInterfaceSwitchDropReasonOperOper(d []interface{}) edpt.InterfaceSwitchDropReasonOperOper {

	count1 := len(d)
	var ret edpt.InterfaceSwitchDropReasonOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interfaces = getSliceInterfaceSwitchDropReasonOperOperInterfaces(in["interfaces"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceSwitchDropReasonOperOperInterfaces(d []interface{}) []edpt.InterfaceSwitchDropReasonOperOperInterfaces {

	count1 := len(d)
	ret := make([]edpt.InterfaceSwitchDropReasonOperOperInterfaces, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceSwitchDropReasonOperOperInterfaces
		oi.IfType = in["if_type"].(string)
		oi.Port_num = in["port_num"].(int)
		oi.Ifinconddrop = in["ifinconddrop"].(int)
		oi.Ifinvlandrop = in["ifinvlandrop"].(int)
		oi.Ifinbitmapzerodrop = in["ifinbitmapzerodrop"].(int)
		oi.Ifinfilterdrop = in["ifinfilterdrop"].(int)
		oi.Ifinportdrop = in["ifinportdrop"].(int)
		oi.Ifinipv4drop = in["ifinipv4drop"].(int)
		oi.Ifinipv6drop = in["ifinipv6drop"].(int)
		oi.Ifinmcdrop = in["ifinmcdrop"].(int)
		oi.Ifinqueuedrop = in["ifinqueuedrop"].(int)
		oi.Ifinerrordrop = in["ifinerrordrop"].(int)
		oi.Ifineventdrop = in["ifineventdrop"].(int)
		oi.Ifinundersizedrop = in["ifinundersizedrop"].(int)
		oi.Ifinmtuexceeddrop = in["ifinmtuexceeddrop"].(int)
		oi.Ifinfragdrop = in["ifinfragdrop"].(int)
		oi.Ifinjabberdrop = in["ifinjabberdrop"].(int)
		oi.Ifinfcsdrop = in["ifinfcsdrop"].(int)
		oi.Ifoutipv4drop = in["ifoutipv4drop"].(int)
		oi.Ifoutipv6drop = in["ifoutipv6drop"].(int)
		oi.Ifouttnldrop = in["ifouttnldrop"].(int)
		oi.Ifoutmcdrop = in["ifoutmcdrop"].(int)
		oi.Ifoutvlandrop = in["ifoutvlandrop"].(int)
		oi.Ifoutconddrop = in["ifoutconddrop"].(int)
		oi.Ifoutparitydrop = in["ifoutparitydrop"].(int)
		oi.Ifoutstgdrop = in["ifoutstgdrop"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceSwitchDropReasonOper(d *schema.ResourceData) edpt.InterfaceSwitchDropReasonOper {
	var ret edpt.InterfaceSwitchDropReasonOper

	ret.Oper = getObjectInterfaceSwitchDropReasonOperOper(d.Get("oper").([]interface{}))
	return ret
}
