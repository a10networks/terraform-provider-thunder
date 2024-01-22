package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugIpv6RedirectTableL2Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_ipv6_redirect_table_l2_oper`: Operational Status for the object l2\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugIpv6RedirectTableL2OperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"redirect_info_list_v6": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"valid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reachable": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"intf_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vlan": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_addr_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mac_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"link_status": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceScaleoutDebugIpv6RedirectTableL2OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugIpv6RedirectTableL2OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugIpv6RedirectTableL2Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugIpv6RedirectTableL2OperOper := setObjectScaleoutDebugIpv6RedirectTableL2OperOper(res)
		d.Set("oper", ScaleoutDebugIpv6RedirectTableL2OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugIpv6RedirectTableL2OperOper(ret edpt.DataScaleoutDebugIpv6RedirectTableL2Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ethernet":              ret.DtScaleoutDebugIpv6RedirectTableL2Oper.Oper.Ethernet,
			"ve":                    ret.DtScaleoutDebugIpv6RedirectTableL2Oper.Oper.Ve,
			"trunk":                 ret.DtScaleoutDebugIpv6RedirectTableL2Oper.Oper.Trunk,
			"redirect_info_list_v6": setSliceScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6(ret.DtScaleoutDebugIpv6RedirectTableL2Oper.Oper.Redirect_info_list_v6),
			"link_status":           ret.DtScaleoutDebugIpv6RedirectTableL2Oper.Oper.Link_status,
		},
	}
}

func setSliceScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6(d []edpt.ScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["node"] = item.Node
		in["valid"] = item.Valid
		in["reachable"] = item.Reachable
		in["intf_num"] = item.Intf_num
		in["vlan"] = item.Vlan
		in["ipv6_addr_str"] = item.Ipv6_addr_str
		in["mac_str"] = item.Mac_str
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugIpv6RedirectTableL2OperOper(d []interface{}) edpt.ScaleoutDebugIpv6RedirectTableL2OperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugIpv6RedirectTableL2OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethernet = in["ethernet"].(int)
		ret.Ve = in["ve"].(int)
		ret.Trunk = in["trunk"].(int)
		ret.Redirect_info_list_v6 = getSliceScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6(in["redirect_info_list_v6"].([]interface{}))
		ret.Link_status = in["link_status"].(int)
	}
	return ret
}

func getSliceScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6(d []interface{}) []edpt.ScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6
		oi.Node = in["node"].(int)
		oi.Valid = in["valid"].(int)
		oi.Reachable = in["reachable"].(int)
		oi.Intf_num = in["intf_num"].(int)
		oi.Vlan = in["vlan"].(int)
		oi.Ipv6_addr_str = in["ipv6_addr_str"].(string)
		oi.Mac_str = in["mac_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugIpv6RedirectTableL2Oper(d *schema.ResourceData) edpt.ScaleoutDebugIpv6RedirectTableL2Oper {
	var ret edpt.ScaleoutDebugIpv6RedirectTableL2Oper

	ret.Oper = getObjectScaleoutDebugIpv6RedirectTableL2OperOper(d.Get("oper").([]interface{}))
	return ret
}
