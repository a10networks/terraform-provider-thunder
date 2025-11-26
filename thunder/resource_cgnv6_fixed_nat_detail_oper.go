package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatDetailOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_detail_oper`: Operational Status for the object detail\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatDetailOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_nat_config_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside_user": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"index": {
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

func resourceCgnv6FixedNatDetailOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatDetailOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatDetailOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatDetailOperOper := setObjectCgnv6FixedNatDetailOperOper(res)
		d.Set("oper", Cgnv6FixedNatDetailOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatDetailOperOper(ret edpt.DataCgnv6FixedNatDetailOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"fixed_nat_config_list": setSliceCgnv6FixedNatDetailOperOperFixedNatConfigList(ret.DtCgnv6FixedNatDetailOper.Oper.FixedNatConfigList),
		},
	}
}

func setSliceCgnv6FixedNatDetailOperOperFixedNatConfigList(d []edpt.Cgnv6FixedNatDetailOperOperFixedNatConfigList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_user"] = item.InsideUser
		in["index"] = item.Index
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6FixedNatDetailOperOper(d []interface{}) edpt.Cgnv6FixedNatDetailOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatDetailOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixedNatConfigList = getSliceCgnv6FixedNatDetailOperOperFixedNatConfigList(in["fixed_nat_config_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6FixedNatDetailOperOperFixedNatConfigList(d []interface{}) []edpt.Cgnv6FixedNatDetailOperOperFixedNatConfigList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatDetailOperOperFixedNatConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatDetailOperOperFixedNatConfigList
		oi.InsideUser = in["inside_user"].(string)
		oi.Index = in["index"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatDetailOper(d *schema.ResourceData) edpt.Cgnv6FixedNatDetailOper {
	var ret edpt.Cgnv6FixedNatDetailOper

	ret.Oper = getObjectCgnv6FixedNatDetailOperOper(d.Get("oper").([]interface{}))
	return ret
}
