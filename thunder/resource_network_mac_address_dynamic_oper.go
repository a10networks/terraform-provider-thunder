package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkMacAddressDynamicOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_mac_address_dynamic_oper`: Operational Status for the object dynamic\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkMacAddressDynamicOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"age_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"macoper": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mac_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vlan": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
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

func resourceNetworkMacAddressDynamicOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAddressDynamicOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAddressDynamicOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkMacAddressDynamicOperOper := setObjectNetworkMacAddressDynamicOperOper(res)
		d.Set("oper", NetworkMacAddressDynamicOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkMacAddressDynamicOperOper(ret edpt.DataNetworkMacAddressDynamicOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"age_time": ret.DtNetworkMacAddressDynamicOper.Oper.AgeTime,
			"macoper":  setSliceNetworkMacAddressDynamicOperOperMacoper(ret.DtNetworkMacAddressDynamicOper.Oper.Macoper),
		},
	}
}

func setSliceNetworkMacAddressDynamicOperOperMacoper(d []edpt.NetworkMacAddressDynamicOperOperMacoper) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["mac_address"] = item.MacAddress
		in["port"] = item.Port
		in["type"] = item.Type
		in["index"] = item.Index
		in["vlan"] = item.Vlan
		in["age"] = item.Age
		result = append(result, in)
	}
	return result
}

func getObjectNetworkMacAddressDynamicOperOper(d []interface{}) edpt.NetworkMacAddressDynamicOperOper {

	count1 := len(d)
	var ret edpt.NetworkMacAddressDynamicOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AgeTime = in["age_time"].(int)
		ret.Macoper = getSliceNetworkMacAddressDynamicOperOperMacoper(in["macoper"].([]interface{}))
	}
	return ret
}

func getSliceNetworkMacAddressDynamicOperOperMacoper(d []interface{}) []edpt.NetworkMacAddressDynamicOperOperMacoper {

	count1 := len(d)
	ret := make([]edpt.NetworkMacAddressDynamicOperOperMacoper, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkMacAddressDynamicOperOperMacoper
		oi.MacAddress = in["mac_address"].(string)
		oi.Port = in["port"].(int)
		oi.Type = in["type"].(string)
		oi.Index = in["index"].(int)
		oi.Vlan = in["vlan"].(int)
		oi.Age = in["age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkMacAddressDynamicOper(d *schema.ResourceData) edpt.NetworkMacAddressDynamicOper {
	var ret edpt.NetworkMacAddressDynamicOper

	ret.Oper = getObjectNetworkMacAddressDynamicOperOper(d.Get("oper").([]interface{}))
	return ret
}
