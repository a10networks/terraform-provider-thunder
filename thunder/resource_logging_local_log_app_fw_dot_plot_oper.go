package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAppFwDotPlotOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_app_fw_dot_plot_oper`: Operational Status for the object dot-plot\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogAppFwDotPlotOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"interval": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"interval_position": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"application_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"client_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"data": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counter": {
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

func resourceLoggingLocalLogAppFwDotPlotOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAppFwDotPlotOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAppFwDotPlotOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogAppFwDotPlotOperOper := setObjectLoggingLocalLogAppFwDotPlotOperOper(res)
		d.Set("oper", LoggingLocalLogAppFwDotPlotOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogAppFwDotPlotOperOper(ret edpt.DataLoggingLocalLogAppFwDotPlotOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time":        ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.StartTime,
			"interval":          ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.Interval,
			"interval_position": ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.IntervalPosition,
			"application_name":  ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.ApplicationName,
			"client_ip":         ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.ClientIp,
			"data":              ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.Data,
			"total":             ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.Total,
			"log_list":          setSliceLoggingLocalLogAppFwDotPlotOperOperLogList(ret.DtLoggingLocalLogAppFwDotPlotOper.Oper.LogList),
		},
	}
}

func setSliceLoggingLocalLogAppFwDotPlotOperOperLogList(d []edpt.LoggingLocalLogAppFwDotPlotOperOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func getObjectLoggingLocalLogAppFwDotPlotOperOper(d []interface{}) edpt.LoggingLocalLogAppFwDotPlotOperOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAppFwDotPlotOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.ApplicationName = in["application_name"].(string)
		ret.ClientIp = in["client_ip"].(string)
		ret.Data = in["data"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAppFwDotPlotOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAppFwDotPlotOperOperLogList(d []interface{}) []edpt.LoggingLocalLogAppFwDotPlotOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAppFwDotPlotOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAppFwDotPlotOperOperLogList
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogAppFwDotPlotOper(d *schema.ResourceData) edpt.LoggingLocalLogAppFwDotPlotOper {
	var ret edpt.LoggingLocalLogAppFwDotPlotOper

	ret.Oper = getObjectLoggingLocalLogAppFwDotPlotOperOper(d.Get("oper").([]interface{}))
	return ret
}
