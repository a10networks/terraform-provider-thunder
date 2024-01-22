package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewShowProcessOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_show_process_oper`: Operational Status for the object show-process\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewShowProcessOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proc_info": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proc_data": {
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

func resourceSystemViewShowProcessOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewShowProcessOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewShowProcessOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewShowProcessOperOper := setObjectSystemViewShowProcessOperOper(res)
		d.Set("oper", SystemViewShowProcessOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewShowProcessOperOper(ret edpt.DataSystemViewShowProcessOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"proc_info": setSliceSystemViewShowProcessOperOperProcInfo(ret.DtSystemViewShowProcessOper.Oper.ProcInfo),
		},
	}
}

func setSliceSystemViewShowProcessOperOperProcInfo(d []edpt.SystemViewShowProcessOperOperProcInfo) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["proc_data"] = item.ProcData
		result = append(result, in)
	}
	return result
}

func getObjectSystemViewShowProcessOperOper(d []interface{}) edpt.SystemViewShowProcessOperOper {

	count1 := len(d)
	var ret edpt.SystemViewShowProcessOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProcInfo = getSliceSystemViewShowProcessOperOperProcInfo(in["proc_info"].([]interface{}))
	}
	return ret
}

func getSliceSystemViewShowProcessOperOperProcInfo(d []interface{}) []edpt.SystemViewShowProcessOperOperProcInfo {

	count1 := len(d)
	ret := make([]edpt.SystemViewShowProcessOperOperProcInfo, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewShowProcessOperOperProcInfo
		oi.ProcData = in["proc_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemViewShowProcessOper(d *schema.ResourceData) edpt.SystemViewShowProcessOper {
	var ret edpt.SystemViewShowProcessOper

	ret.Oper = getObjectSystemViewShowProcessOperOper(d.Get("oper").([]interface{}))
	return ret
}
