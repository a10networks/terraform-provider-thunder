package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAMacOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_mac_oper`: Operational Status for the object mac\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAMacOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid_mac_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mac_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mac_with_syn_cookie": {
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

func resourceVrrpAMacOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAMacOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAMacOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAMacOperOper := setObjectVrrpAMacOperOper(res)
		d.Set("oper", VrrpAMacOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAMacOperOper(ret edpt.DataVrrpAMacOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vrid_mac_list": setSliceVrrpAMacOperOperVridMacList(ret.DtVrrpAMacOper.Oper.VridMacList),
		},
	}
}

func setSliceVrrpAMacOperOperVridMacList(d []edpt.VrrpAMacOperOperVridMacList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vrid"] = item.Vrid
		in["mac_address"] = item.Mac_address
		in["mac_with_syn_cookie"] = item.Mac_with_syn_cookie
		result = append(result, in)
	}
	return result
}

func getObjectVrrpAMacOperOper(d []interface{}) edpt.VrrpAMacOperOper {

	count1 := len(d)
	var ret edpt.VrrpAMacOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VridMacList = getSliceVrrpAMacOperOperVridMacList(in["vrid_mac_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAMacOperOperVridMacList(d []interface{}) []edpt.VrrpAMacOperOperVridMacList {

	count1 := len(d)
	ret := make([]edpt.VrrpAMacOperOperVridMacList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAMacOperOperVridMacList
		oi.Vrid = in["vrid"].(int)
		oi.Mac_address = in["mac_address"].(string)
		oi.Mac_with_syn_cookie = in["mac_with_syn_cookie"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAMacOper(d *schema.ResourceData) edpt.VrrpAMacOper {
	var ret edpt.VrrpAMacOper

	ret.Oper = getObjectVrrpAMacOperOper(d.Get("oper").([]interface{}))
	return ret
}
