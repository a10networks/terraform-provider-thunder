package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpATeMacOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_te_mac_oper`: Operational Status for the object te-mac\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpATeMacOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"invalid_flag": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vrid_mac_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mac_address_inside": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mac_address_outside": {
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

func resourceVrrpATeMacOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpATeMacOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpATeMacOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpATeMacOperOper := setObjectVrrpATeMacOperOper(res)
		d.Set("oper", VrrpATeMacOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpATeMacOperOper(ret edpt.DataVrrpATeMacOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"invalid_flag":  ret.DtVrrpATeMacOper.Oper.InvalidFlag,
			"vrid_mac_list": setSliceVrrpATeMacOperOperVridMacList(ret.DtVrrpATeMacOper.Oper.VridMacList),
		},
	}
}

func setSliceVrrpATeMacOperOperVridMacList(d []edpt.VrrpATeMacOperOperVridMacList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vrid"] = item.Vrid
		in["mac_address_inside"] = item.Mac_address_inside
		in["mac_address_outside"] = item.Mac_address_outside
		result = append(result, in)
	}
	return result
}

func getObjectVrrpATeMacOperOper(d []interface{}) edpt.VrrpATeMacOperOper {

	count1 := len(d)
	var ret edpt.VrrpATeMacOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InvalidFlag = in["invalid_flag"].(int)
		ret.VridMacList = getSliceVrrpATeMacOperOperVridMacList(in["vrid_mac_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpATeMacOperOperVridMacList(d []interface{}) []edpt.VrrpATeMacOperOperVridMacList {

	count1 := len(d)
	ret := make([]edpt.VrrpATeMacOperOperVridMacList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpATeMacOperOperVridMacList
		oi.Vrid = in["vrid"].(int)
		oi.Mac_address_inside = in["mac_address_inside"].(string)
		oi.Mac_address_outside = in["mac_address_outside"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpATeMacOper(d *schema.ResourceData) edpt.VrrpATeMacOper {
	var ret edpt.VrrpATeMacOper

	ret.Oper = getObjectVrrpATeMacOperOper(d.Get("oper").([]interface{}))
	return ret
}
