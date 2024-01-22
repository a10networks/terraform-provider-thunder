package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRateLimitLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_rate_limit_log_oper`: Operational Status for the object rate-limit-log\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbRateLimitLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit_log_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total_log_times": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_log_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local_log_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"remote_log_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local_log_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"remote_log_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_too_big": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buff_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buff_send_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"free_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_repeat_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local_log_dropped": {
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

func resourceSlbRateLimitLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRateLimitLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRateLimitLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbRateLimitLogOperOper := setObjectSlbRateLimitLogOperOper(res)
		d.Set("oper", SlbRateLimitLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbRateLimitLogOperOper(ret edpt.DataSlbRateLimitLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rate_limit_log_cpu_list": setSliceSlbRateLimitLogOperOperRateLimitLogCpuList(ret.DtSlbRateLimitLogOper.Oper.RateLimitLogCpuList),
			"cpu_count":               ret.DtSlbRateLimitLogOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbRateLimitLogOperOperRateLimitLogCpuList(d []edpt.SlbRateLimitLogOperOperRateLimitLogCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["total_log_times"] = item.Total_log_times
		in["total_log_msg"] = item.Total_log_msg
		in["local_log_msg"] = item.Local_log_msg
		in["remote_log_msg"] = item.Remote_log_msg
		in["local_log_rate"] = item.Local_log_rate
		in["remote_log_rate"] = item.Remote_log_rate
		in["msg_too_big"] = item.Msg_too_big
		in["buff_alloc_fail"] = item.Buff_alloc_fail
		in["no_route"] = item.No_route
		in["buff_send_fail"] = item.Buff_send_fail
		in["alloc_conn"] = item.Alloc_conn
		in["free_conn"] = item.Free_conn
		in["conn_alloc_fail"] = item.Conn_alloc_fail
		in["no_repeat_msg"] = item.No_repeat_msg
		in["local_log_dropped"] = item.Local_log_dropped
		result = append(result, in)
	}
	return result
}

func getObjectSlbRateLimitLogOperOper(d []interface{}) edpt.SlbRateLimitLogOperOper {

	count1 := len(d)
	var ret edpt.SlbRateLimitLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimitLogCpuList = getSliceSlbRateLimitLogOperOperRateLimitLogCpuList(in["rate_limit_log_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbRateLimitLogOperOperRateLimitLogCpuList(d []interface{}) []edpt.SlbRateLimitLogOperOperRateLimitLogCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbRateLimitLogOperOperRateLimitLogCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRateLimitLogOperOperRateLimitLogCpuList
		oi.Total_log_times = in["total_log_times"].(int)
		oi.Total_log_msg = in["total_log_msg"].(int)
		oi.Local_log_msg = in["local_log_msg"].(int)
		oi.Remote_log_msg = in["remote_log_msg"].(int)
		oi.Local_log_rate = in["local_log_rate"].(int)
		oi.Remote_log_rate = in["remote_log_rate"].(int)
		oi.Msg_too_big = in["msg_too_big"].(int)
		oi.Buff_alloc_fail = in["buff_alloc_fail"].(int)
		oi.No_route = in["no_route"].(int)
		oi.Buff_send_fail = in["buff_send_fail"].(int)
		oi.Alloc_conn = in["alloc_conn"].(int)
		oi.Free_conn = in["free_conn"].(int)
		oi.Conn_alloc_fail = in["conn_alloc_fail"].(int)
		oi.No_repeat_msg = in["no_repeat_msg"].(int)
		oi.Local_log_dropped = in["local_log_dropped"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbRateLimitLogOper(d *schema.ResourceData) edpt.SlbRateLimitLogOper {
	var ret edpt.SlbRateLimitLogOper

	ret.Oper = getObjectSlbRateLimitLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
