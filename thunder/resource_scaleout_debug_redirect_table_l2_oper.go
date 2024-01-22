package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugRedirectTableL2Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_redirect_table_l2_oper`: Operational Status for the object l2\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugRedirectTableL2OperRead,

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
						"redirect_info_list": {
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
									"ip_addr_str": {
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

func resourceScaleoutDebugRedirectTableL2OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugRedirectTableL2OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugRedirectTableL2Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugRedirectTableL2OperOper := setObjectScaleoutDebugRedirectTableL2OperOper(res)
		d.Set("oper", ScaleoutDebugRedirectTableL2OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugRedirectTableL2OperOper(ret edpt.DataScaleoutDebugRedirectTableL2Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ethernet":           ret.DtScaleoutDebugRedirectTableL2Oper.Oper.Ethernet,
			"ve":                 ret.DtScaleoutDebugRedirectTableL2Oper.Oper.Ve,
			"trunk":              ret.DtScaleoutDebugRedirectTableL2Oper.Oper.Trunk,
			"redirect_info_list": setSliceScaleoutDebugRedirectTableL2OperOperRedirect_info_list(ret.DtScaleoutDebugRedirectTableL2Oper.Oper.Redirect_info_list),
			"link_status":        ret.DtScaleoutDebugRedirectTableL2Oper.Oper.Link_status,
		},
	}
}

func setSliceScaleoutDebugRedirectTableL2OperOperRedirect_info_list(d []edpt.ScaleoutDebugRedirectTableL2OperOperRedirect_info_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["node"] = item.Node
		in["valid"] = item.Valid
		in["reachable"] = item.Reachable
		in["intf_num"] = item.Intf_num
		in["vlan"] = item.Vlan
		in["ip_addr_str"] = item.Ip_addr_str
		in["mac_str"] = item.Mac_str
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugRedirectTableL2OperOper(d []interface{}) edpt.ScaleoutDebugRedirectTableL2OperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugRedirectTableL2OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethernet = in["ethernet"].(int)
		ret.Ve = in["ve"].(int)
		ret.Trunk = in["trunk"].(int)
		ret.Redirect_info_list = getSliceScaleoutDebugRedirectTableL2OperOperRedirect_info_list(in["redirect_info_list"].([]interface{}))
		ret.Link_status = in["link_status"].(int)
	}
	return ret
}

func getSliceScaleoutDebugRedirectTableL2OperOperRedirect_info_list(d []interface{}) []edpt.ScaleoutDebugRedirectTableL2OperOperRedirect_info_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugRedirectTableL2OperOperRedirect_info_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugRedirectTableL2OperOperRedirect_info_list
		oi.Node = in["node"].(int)
		oi.Valid = in["valid"].(int)
		oi.Reachable = in["reachable"].(int)
		oi.Intf_num = in["intf_num"].(int)
		oi.Vlan = in["vlan"].(int)
		oi.Ip_addr_str = in["ip_addr_str"].(string)
		oi.Mac_str = in["mac_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugRedirectTableL2Oper(d *schema.ResourceData) edpt.ScaleoutDebugRedirectTableL2Oper {
	var ret edpt.ScaleoutDebugRedirectTableL2Oper

	ret.Oper = getObjectScaleoutDebugRedirectTableL2OperOper(d.Get("oper").([]interface{}))
	return ret
}
