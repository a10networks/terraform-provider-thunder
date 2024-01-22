package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugIpv6ReachabilityOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_ipv6_reachability_oper`: Operational Status for the object reachability\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugIpv6ReachabilityOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"part_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"scaleout_ipv6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_addr": {
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
									"prefix_len": {
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

func resourceScaleoutDebugIpv6ReachabilityOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugIpv6ReachabilityOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugIpv6ReachabilityOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugIpv6ReachabilityOperOper := setObjectScaleoutDebugIpv6ReachabilityOperOper(res)
		d.Set("oper", ScaleoutDebugIpv6ReachabilityOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugIpv6ReachabilityOperOper(ret edpt.DataScaleoutDebugIpv6ReachabilityOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"part_name":          ret.DtScaleoutDebugIpv6ReachabilityOper.Oper.Part_name,
			"scaleout_ipv6_list": setSliceScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List(ret.DtScaleoutDebugIpv6ReachabilityOper.Oper.ScaleoutIpv6List),
		},
	}
}

func setSliceScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List(d []edpt.ScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["node"] = item.Node
		in["ipv6_addr"] = item.Ipv6_addr
		in["mac"] = item.Mac
		in["vnp_id"] = item.Vnp_id
		in["vlan_id"] = item.Vlan_id
		in["prefix_len"] = item.Prefix_len
		in["real_port"] = item.Real_port
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugIpv6ReachabilityOperOper(d []interface{}) edpt.ScaleoutDebugIpv6ReachabilityOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugIpv6ReachabilityOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Part_name = in["part_name"].(string)
		ret.ScaleoutIpv6List = getSliceScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List(in["scaleout_ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List(d []interface{}) []edpt.ScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List
		oi.Node = in["node"].(int)
		oi.Ipv6_addr = in["ipv6_addr"].(string)
		oi.Mac = in["mac"].(string)
		oi.Vnp_id = in["vnp_id"].(int)
		oi.Vlan_id = in["vlan_id"].(int)
		oi.Prefix_len = in["prefix_len"].(int)
		oi.Real_port = in["real_port"].(int)
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugIpv6ReachabilityOper(d *schema.ResourceData) edpt.ScaleoutDebugIpv6ReachabilityOper {
	var ret edpt.ScaleoutDebugIpv6ReachabilityOper

	ret.Oper = getObjectScaleoutDebugIpv6ReachabilityOperOper(d.Get("oper").([]interface{}))
	return ret
}
