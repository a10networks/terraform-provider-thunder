package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbAflowOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_aflow_oper`: Operational Status for the object aflow\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbAflowOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aflow_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pause_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pause_conn_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resume_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"event_resume_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"timer_resume_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"try_to_resume_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retry_resume_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"error_resume_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_new_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reuse_server_idle_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inc_aflow_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbAflowOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAflowOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAflowOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbAflowOperOper := setObjectSlbAflowOperOper(res)
		d.Set("oper", SlbAflowOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbAflowOperOper(ret edpt.DataSlbAflowOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"aflow_cpu_list": setSliceSlbAflowOperOperAflowCpuList(ret.DtSlbAflowOper.Oper.AflowCpuList),
			"cpu_count":      ret.DtSlbAflowOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbAflowOperOperAflowCpuList(d []edpt.SlbAflowOperOperAflowCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pause_conn"] = item.Pause_conn
		in["pause_conn_fail"] = item.Pause_conn_fail
		in["resume_conn"] = item.Resume_conn
		in["event_resume_conn"] = item.Event_resume_conn
		in["timer_resume_conn"] = item.Timer_resume_conn
		in["try_to_resume_conn"] = item.Try_to_resume_conn
		in["retry_resume_conn"] = item.Retry_resume_conn
		in["error_resume_conn"] = item.Error_resume_conn
		in["open_new_server_conn"] = item.Open_new_server_conn
		in["reuse_server_idle_conn"] = item.Reuse_server_idle_conn
		in["inc_aflow_limit"] = item.Inc_aflow_limit
		result = append(result, in)
	}
	return result
}

func getObjectSlbAflowOperOper(d []interface{}) edpt.SlbAflowOperOper {

	count1 := len(d)
	var ret edpt.SlbAflowOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AflowCpuList = getSliceSlbAflowOperOperAflowCpuList(in["aflow_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbAflowOperOperAflowCpuList(d []interface{}) []edpt.SlbAflowOperOperAflowCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbAflowOperOperAflowCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAflowOperOperAflowCpuList
		oi.Pause_conn = in["pause_conn"].(int)
		oi.Pause_conn_fail = in["pause_conn_fail"].(int)
		oi.Resume_conn = in["resume_conn"].(int)
		oi.Event_resume_conn = in["event_resume_conn"].(int)
		oi.Timer_resume_conn = in["timer_resume_conn"].(int)
		oi.Try_to_resume_conn = in["try_to_resume_conn"].(int)
		oi.Retry_resume_conn = in["retry_resume_conn"].(int)
		oi.Error_resume_conn = in["error_resume_conn"].(int)
		oi.Open_new_server_conn = in["open_new_server_conn"].(int)
		oi.Reuse_server_idle_conn = in["reuse_server_idle_conn"].(int)
		oi.Inc_aflow_limit = in["inc_aflow_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbAflowOper(d *schema.ResourceData) edpt.SlbAflowOper {
	var ret edpt.SlbAflowOper

	ret.Oper = getObjectSlbAflowOperOper(d.Get("oper").([]interface{}))
	return ret
}
