package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSessionsSmpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sessions_smp_oper`: Operational Status for the object smp\n\n__PLACEHOLDER__",
		ReadContext: resourceSessionsSmpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_smp_list": {
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
									"alloc_fail": {
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

func resourceSessionsSmpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsSmpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionsSmpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SessionsSmpOperOper := setObjectSessionsSmpOperOper(res)
		d.Set("oper", SessionsSmpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSessionsSmpOperOper(ret edpt.DataSessionsSmpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_smp_list": setSliceSessionsSmpOperOperSessionSmpList(ret.DtSessionsSmpOper.Oper.SessionSmpList),
		},
	}
}

func setSliceSessionsSmpOperOperSessionSmpList(d []edpt.SessionsSmpOperOperSessionSmpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["alloc"] = item.Alloc
		in["free"] = item.Free
		in["alloc_fail"] = item.AllocFail
		result = append(result, in)
	}
	return result
}

func getObjectSessionsSmpOperOper(d []interface{}) edpt.SessionsSmpOperOper {

	count1 := len(d)
	var ret edpt.SessionsSmpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionSmpList = getSliceSessionsSmpOperOperSessionSmpList(in["session_smp_list"].([]interface{}))
	}
	return ret
}

func getSliceSessionsSmpOperOperSessionSmpList(d []interface{}) []edpt.SessionsSmpOperOperSessionSmpList {

	count1 := len(d)
	ret := make([]edpt.SessionsSmpOperOperSessionSmpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsSmpOperOperSessionSmpList
		oi.Type = in["type"].(string)
		oi.Alloc = in["alloc"].(int)
		oi.Free = in["free"].(int)
		oi.AllocFail = in["alloc_fail"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSessionsSmpOper(d *schema.ResourceData) edpt.SessionsSmpOper {
	var ret edpt.SessionsSmpOper

	ret.Oper = getObjectSessionsSmpOperOper(d.Get("oper").([]interface{}))
	return ret
}
