package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_log_oper`: Operational Status for the object log\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpn_log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpn_log_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"vpn_log_offset": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vpn_log_over": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"follow": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"from_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_lines": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVpnLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnLogOperOper := setObjectVpnLogOperOper(res)
		d.Set("oper", VpnLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnLogOperOper(ret edpt.DataVpnLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vpn_log_list":   setSliceVpnLogOperOperVpnLogList(ret.DtVpnLogOper.Oper.VpnLogList),
			"vpn_log_offset": ret.DtVpnLogOper.Oper.VpnLogOffset,
			"vpn_log_over":   ret.DtVpnLogOper.Oper.VpnLogOver,
			"follow":         ret.DtVpnLogOper.Oper.Follow,
			"from_start":     ret.DtVpnLogOper.Oper.FromStart,
			"num_lines":      ret.DtVpnLogOper.Oper.NumLines,
		},
	}
}

func setSliceVpnLogOperOperVpnLogList(d []edpt.VpnLogOperOperVpnLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vpn_log_data"] = item.VpnLogData
		result = append(result, in)
	}
	return result
}

func getObjectVpnLogOperOper(d []interface{}) edpt.VpnLogOperOper {

	count1 := len(d)
	var ret edpt.VpnLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VpnLogList = getSliceVpnLogOperOperVpnLogList(in["vpn_log_list"].([]interface{}))
		ret.VpnLogOffset = in["vpn_log_offset"].(int)
		ret.VpnLogOver = in["vpn_log_over"].(int)
		ret.Follow = in["follow"].(int)
		ret.FromStart = in["from_start"].(int)
		ret.NumLines = in["num_lines"].(int)
	}
	return ret
}

func getSliceVpnLogOperOperVpnLogList(d []interface{}) []edpt.VpnLogOperOperVpnLogList {

	count1 := len(d)
	ret := make([]edpt.VpnLogOperOperVpnLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnLogOperOperVpnLogList
		oi.VpnLogData = in["vpn_log_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnLogOper(d *schema.ResourceData) edpt.VpnLogOper {
	var ret edpt.VpnLogOper

	ret.Oper = getObjectVpnLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
