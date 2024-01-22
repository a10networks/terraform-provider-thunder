package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkArpDynamicOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_arp_dynamic_oper`: Operational Status for the object dynamic\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkArpDynamicOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mac_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interface": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vlan": {
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

func resourceNetworkArpDynamicOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpDynamicOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpDynamicOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkArpDynamicOperOper := setObjectNetworkArpDynamicOperOper(res)
		d.Set("oper", NetworkArpDynamicOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkArpDynamicOperOper(ret edpt.DataNetworkArpDynamicOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"arp_list": setSliceNetworkArpDynamicOperOperArpList(ret.DtNetworkArpDynamicOper.Oper.ArpList),
		},
	}
}

func setSliceNetworkArpDynamicOperOperArpList(d []edpt.NetworkArpDynamicOperOperArpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_address"] = item.IpAddress
		in["mac_address"] = item.MacAddress
		in["type"] = item.Type
		in["age"] = item.Age
		in["interface"] = item.Interface
		in["vlan"] = item.Vlan
		result = append(result, in)
	}
	return result
}

func getObjectNetworkArpDynamicOperOper(d []interface{}) edpt.NetworkArpDynamicOperOper {

	count1 := len(d)
	var ret edpt.NetworkArpDynamicOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ArpList = getSliceNetworkArpDynamicOperOperArpList(in["arp_list"].([]interface{}))
	}
	return ret
}

func getSliceNetworkArpDynamicOperOperArpList(d []interface{}) []edpt.NetworkArpDynamicOperOperArpList {

	count1 := len(d)
	ret := make([]edpt.NetworkArpDynamicOperOperArpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkArpDynamicOperOperArpList
		oi.IpAddress = in["ip_address"].(string)
		oi.MacAddress = in["mac_address"].(string)
		oi.Type = in["type"].(string)
		oi.Age = in["age"].(int)
		oi.Interface = in["interface"].(string)
		oi.Vlan = in["vlan"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkArpDynamicOper(d *schema.ResourceData) edpt.NetworkArpDynamicOper {
	var ret edpt.NetworkArpDynamicOper

	ret.Oper = getObjectNetworkArpDynamicOperOper(d.Get("oper").([]interface{}))
	return ret
}
