package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAppFwTopNOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_app_fw_top_n_oper`: Operational Status for the object top-n\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogAppFwTopNOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"start_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"interval": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"interval_position": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"top": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"application_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"category": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"client_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"threat_category": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"threat_category_match": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
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

func resourceLoggingLocalLogAppFwTopNOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAppFwTopNOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAppFwTopNOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogAppFwTopNOperOper := setObjectLoggingLocalLogAppFwTopNOperOper(res)
		d.Set("oper", LoggingLocalLogAppFwTopNOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogAppFwTopNOperOper(ret edpt.DataLoggingLocalLogAppFwTopNOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"max_entries":           ret.DtLoggingLocalLogAppFwTopNOper.Oper.MaxEntries,
			"start_time":            ret.DtLoggingLocalLogAppFwTopNOper.Oper.StartTime,
			"interval":              ret.DtLoggingLocalLogAppFwTopNOper.Oper.Interval,
			"interval_position":     ret.DtLoggingLocalLogAppFwTopNOper.Oper.IntervalPosition,
			"top":                   ret.DtLoggingLocalLogAppFwTopNOper.Oper.Top,
			"application_name":      ret.DtLoggingLocalLogAppFwTopNOper.Oper.ApplicationName,
			"category":              ret.DtLoggingLocalLogAppFwTopNOper.Oper.Category,
			"client_ip":             ret.DtLoggingLocalLogAppFwTopNOper.Oper.ClientIp,
			"threat_category":       ret.DtLoggingLocalLogAppFwTopNOper.Oper.ThreatCategory,
			"threat_category_match": ret.DtLoggingLocalLogAppFwTopNOper.Oper.ThreatCategoryMatch,
			"action":                ret.DtLoggingLocalLogAppFwTopNOper.Oper.Action,
			"total":                 ret.DtLoggingLocalLogAppFwTopNOper.Oper.Total,
			"log_list":              setSliceLoggingLocalLogAppFwTopNOperOperLogList(ret.DtLoggingLocalLogAppFwTopNOper.Oper.LogList),
		},
	}
}

func setSliceLoggingLocalLogAppFwTopNOperOperLogList(d []edpt.LoggingLocalLogAppFwTopNOperOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func getObjectLoggingLocalLogAppFwTopNOperOper(d []interface{}) edpt.LoggingLocalLogAppFwTopNOperOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAppFwTopNOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxEntries = in["max_entries"].(int)
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.Top = in["top"].(string)
		ret.ApplicationName = in["application_name"].(string)
		ret.Category = in["category"].(string)
		ret.ClientIp = in["client_ip"].(string)
		ret.ThreatCategory = in["threat_category"].(string)
		ret.ThreatCategoryMatch = in["threat_category_match"].(string)
		ret.Action = in["action"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAppFwTopNOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAppFwTopNOperOperLogList(d []interface{}) []edpt.LoggingLocalLogAppFwTopNOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAppFwTopNOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAppFwTopNOperOperLogList
		oi.Name = in["name"].(string)
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogAppFwTopNOper(d *schema.ResourceData) edpt.LoggingLocalLogAppFwTopNOper {
	var ret edpt.LoggingLocalLogAppFwTopNOper

	ret.Oper = getObjectLoggingLocalLogAppFwTopNOperOper(d.Get("oper").([]interface{}))
	return ret
}
