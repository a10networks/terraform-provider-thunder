package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPowerOnSelfTestOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_power_on_self_test_oper`: Operational Status for the object power-on-self-test\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemPowerOnSelfTestOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"power_on_log": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dlog_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dlog_data_search": {
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

func resourceSystemPowerOnSelfTestOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPowerOnSelfTestOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPowerOnSelfTestOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemPowerOnSelfTestOperOper := setObjectSystemPowerOnSelfTestOperOper(res)
		d.Set("oper", SystemPowerOnSelfTestOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemPowerOnSelfTestOperOper(ret edpt.DataSystemPowerOnSelfTestOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"power_on_log": setSliceSystemPowerOnSelfTestOperOperPowerOnLog(ret.DtSystemPowerOnSelfTestOper.Oper.PowerOnLog),
		},
	}
}

func setSliceSystemPowerOnSelfTestOperOperPowerOnLog(d []edpt.SystemPowerOnSelfTestOperOperPowerOnLog) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dlog_data"] = item.DlogData
		in["dlog_data_search"] = item.DlogDataSearch
		result = append(result, in)
	}
	return result
}

func getObjectSystemPowerOnSelfTestOperOper(d []interface{}) edpt.SystemPowerOnSelfTestOperOper {

	count1 := len(d)
	var ret edpt.SystemPowerOnSelfTestOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PowerOnLog = getSliceSystemPowerOnSelfTestOperOperPowerOnLog(in["power_on_log"].([]interface{}))
	}
	return ret
}

func getSliceSystemPowerOnSelfTestOperOperPowerOnLog(d []interface{}) []edpt.SystemPowerOnSelfTestOperOperPowerOnLog {

	count1 := len(d)
	ret := make([]edpt.SystemPowerOnSelfTestOperOperPowerOnLog, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemPowerOnSelfTestOperOperPowerOnLog
		oi.DlogData = in["dlog_data"].(string)
		oi.DlogDataSearch = in["dlog_data_search"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemPowerOnSelfTestOper(d *schema.ResourceData) edpt.SystemPowerOnSelfTestOper {
	var ret edpt.SystemPowerOnSelfTestOper

	ret.Oper = getObjectSystemPowerOnSelfTestOperOper(d.Get("oper").([]interface{}))
	return ret
}
