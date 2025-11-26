package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAccessTopNOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_access_top_n_oper`: Operational Status for the object top-n\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogAccessTopNOperRead,

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

func resourceLoggingLocalLogAccessTopNOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAccessTopNOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAccessTopNOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogAccessTopNOperOper := setObjectLoggingLocalLogAccessTopNOperOper(res)
		d.Set("oper", LoggingLocalLogAccessTopNOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogAccessTopNOperOper(ret edpt.DataLoggingLocalLogAccessTopNOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"max_entries":       ret.DtLoggingLocalLogAccessTopNOper.Oper.MaxEntries,
			"start_time":        ret.DtLoggingLocalLogAccessTopNOper.Oper.StartTime,
			"interval":          ret.DtLoggingLocalLogAccessTopNOper.Oper.Interval,
			"interval_position": ret.DtLoggingLocalLogAccessTopNOper.Oper.IntervalPosition,
			"top":               ret.DtLoggingLocalLogAccessTopNOper.Oper.Top,
			"action":            ret.DtLoggingLocalLogAccessTopNOper.Oper.Action,
			"total":             ret.DtLoggingLocalLogAccessTopNOper.Oper.Total,
			"log_list":          setSliceLoggingLocalLogAccessTopNOperOperLogList(ret.DtLoggingLocalLogAccessTopNOper.Oper.LogList),
		},
	}
}

func setSliceLoggingLocalLogAccessTopNOperOperLogList(d []edpt.LoggingLocalLogAccessTopNOperOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func getObjectLoggingLocalLogAccessTopNOperOper(d []interface{}) edpt.LoggingLocalLogAccessTopNOperOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAccessTopNOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxEntries = in["max_entries"].(int)
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.Top = in["top"].(string)
		ret.Action = in["action"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAccessTopNOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAccessTopNOperOperLogList(d []interface{}) []edpt.LoggingLocalLogAccessTopNOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAccessTopNOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAccessTopNOperOperLogList
		oi.Name = in["name"].(string)
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogAccessTopNOper(d *schema.ResourceData) edpt.LoggingLocalLogAccessTopNOper {
	var ret edpt.LoggingLocalLogAccessTopNOper

	ret.Oper = getObjectLoggingLocalLogAccessTopNOperOper(d.Get("oper").([]interface{}))
	return ret
}
