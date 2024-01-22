package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NeighborDynamicOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_neighbor_dynamic_oper`: Operational Status for the object dynamic\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6NeighborDynamicOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v6neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address": {
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
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceIpv6NeighborDynamicOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NeighborDynamicOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NeighborDynamicOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6NeighborDynamicOperOper := setObjectIpv6NeighborDynamicOperOper(res)
		d.Set("oper", Ipv6NeighborDynamicOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6NeighborDynamicOperOper(ret edpt.DataIpv6NeighborDynamicOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"v6neighbor_list": setSliceIpv6NeighborDynamicOperOperV6neighborList(ret.DtIpv6NeighborDynamicOper.Oper.V6neighborList),
		},
	}
}

func setSliceIpv6NeighborDynamicOperOperV6neighborList(d []edpt.Ipv6NeighborDynamicOperOperV6neighborList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv6_address"] = item.Ipv6Address
		in["mac_address"] = item.MacAddress
		in["type"] = item.Type
		in["age"] = item.Age
		in["state"] = item.State
		in["interface"] = item.Interface
		in["vlan"] = item.Vlan
		result = append(result, in)
	}
	return result
}

func getObjectIpv6NeighborDynamicOperOper(d []interface{}) edpt.Ipv6NeighborDynamicOperOper {

	count1 := len(d)
	var ret edpt.Ipv6NeighborDynamicOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V6neighborList = getSliceIpv6NeighborDynamicOperOperV6neighborList(in["v6neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceIpv6NeighborDynamicOperOperV6neighborList(d []interface{}) []edpt.Ipv6NeighborDynamicOperOperV6neighborList {

	count1 := len(d)
	ret := make([]edpt.Ipv6NeighborDynamicOperOperV6neighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6NeighborDynamicOperOperV6neighborList
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.MacAddress = in["mac_address"].(string)
		oi.Type = in["type"].(string)
		oi.Age = in["age"].(int)
		oi.State = in["state"].(string)
		oi.Interface = in["interface"].(string)
		oi.Vlan = in["vlan"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6NeighborDynamicOper(d *schema.ResourceData) edpt.Ipv6NeighborDynamicOper {
	var ret edpt.Ipv6NeighborDynamicOper

	ret.Oper = getObjectIpv6NeighborDynamicOperOper(d.Get("oper").([]interface{}))
	return ret
}
