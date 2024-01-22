package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugReachabilityOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_reachability_oper`: Operational Status for the object reachability\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugReachabilityOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"part_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"scaleout_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mac": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vnp_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vlan_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"netmask": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"real_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name": {
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

func resourceScaleoutDebugReachabilityOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugReachabilityOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugReachabilityOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugReachabilityOperOper := setObjectScaleoutDebugReachabilityOperOper(res)
		d.Set("oper", ScaleoutDebugReachabilityOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugReachabilityOperOper(ret edpt.DataScaleoutDebugReachabilityOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"part_name":        ret.DtScaleoutDebugReachabilityOper.Oper.Part_name,
			"scaleout_ip_list": setSliceScaleoutDebugReachabilityOperOperScaleoutIpList(ret.DtScaleoutDebugReachabilityOper.Oper.ScaleoutIpList),
		},
	}
}

func setSliceScaleoutDebugReachabilityOperOperScaleoutIpList(d []edpt.ScaleoutDebugReachabilityOperOperScaleoutIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["node"] = item.Node
		in["ip_addr"] = item.Ip_addr
		in["mac"] = item.Mac
		in["vnp_id"] = item.Vnp_id
		in["vlan_id"] = item.Vlan_id
		in["netmask"] = item.Netmask
		in["real_port"] = item.Real_port
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugReachabilityOperOper(d []interface{}) edpt.ScaleoutDebugReachabilityOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugReachabilityOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Part_name = in["part_name"].(string)
		ret.ScaleoutIpList = getSliceScaleoutDebugReachabilityOperOperScaleoutIpList(in["scaleout_ip_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugReachabilityOperOperScaleoutIpList(d []interface{}) []edpt.ScaleoutDebugReachabilityOperOperScaleoutIpList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugReachabilityOperOperScaleoutIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugReachabilityOperOperScaleoutIpList
		oi.Node = in["node"].(int)
		oi.Ip_addr = in["ip_addr"].(string)
		oi.Mac = in["mac"].(string)
		oi.Vnp_id = in["vnp_id"].(int)
		oi.Vlan_id = in["vlan_id"].(int)
		oi.Netmask = in["netmask"].(int)
		oi.Real_port = in["real_port"].(int)
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugReachabilityOper(d *schema.ResourceData) edpt.ScaleoutDebugReachabilityOper {
	var ret edpt.ScaleoutDebugReachabilityOper

	ret.Oper = getObjectScaleoutDebugReachabilityOperOper(d.Get("oper").([]interface{}))
	return ret
}
