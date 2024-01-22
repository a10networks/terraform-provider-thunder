package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSessionsExtOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sessions_ext_oper`: Operational Status for the object ext\n\n__PLACEHOLDER__",
		ReadContext: resourceSessionsExtOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_ext_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cpu_round_robin_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_exceed": {
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

func resourceSessionsExtOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsExtOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionsExtOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SessionsExtOperOper := setObjectSessionsExtOperOper(res)
		d.Set("oper", SessionsExtOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSessionsExtOperOper(ret edpt.DataSessionsExtOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_ext_list": setSliceSessionsExtOperOperSessionExtList(ret.DtSessionsExtOper.Oper.SessionExtList),
		},
	}
}

func setSliceSessionsExtOperOperSessionExtList(d []edpt.SessionsExtOperOperSessionExtList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["alloc"] = item.Alloc
		in["free"] = item.Free
		in["fail"] = item.Fail
		in["cpu_round_robin_fail"] = item.CpuRoundRobinFail
		in["alloc_exceed"] = item.AllocExceed
		result = append(result, in)
	}
	return result
}

func getObjectSessionsExtOperOper(d []interface{}) edpt.SessionsExtOperOper {

	count1 := len(d)
	var ret edpt.SessionsExtOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionExtList = getSliceSessionsExtOperOperSessionExtList(in["session_ext_list"].([]interface{}))
	}
	return ret
}

func getSliceSessionsExtOperOperSessionExtList(d []interface{}) []edpt.SessionsExtOperOperSessionExtList {

	count1 := len(d)
	ret := make([]edpt.SessionsExtOperOperSessionExtList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsExtOperOperSessionExtList
		oi.Type = in["type"].(string)
		oi.Alloc = in["alloc"].(int)
		oi.Free = in["free"].(int)
		oi.Fail = in["fail"].(int)
		oi.CpuRoundRobinFail = in["cpu_round_robin_fail"].(int)
		oi.AllocExceed = in["alloc_exceed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSessionsExtOper(d *schema.ResourceData) edpt.SessionsExtOper {
	var ret edpt.SessionsExtOper

	ret.Oper = getObjectSessionsExtOperOper(d.Get("oper").([]interface{}))
	return ret
}
