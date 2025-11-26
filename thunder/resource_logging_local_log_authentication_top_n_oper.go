package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAuthenticationTopNOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_authentication_top_n_oper`: Operational Status for the object top-n\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogAuthenticationTopNOperRead,

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
						"auth_result": {
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

func resourceLoggingLocalLogAuthenticationTopNOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAuthenticationTopNOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAuthenticationTopNOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogAuthenticationTopNOperOper := setObjectLoggingLocalLogAuthenticationTopNOperOper(res)
		d.Set("oper", LoggingLocalLogAuthenticationTopNOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogAuthenticationTopNOperOper(ret edpt.DataLoggingLocalLogAuthenticationTopNOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"max_entries":       ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.MaxEntries,
			"start_time":        ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.StartTime,
			"interval":          ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.Interval,
			"interval_position": ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.IntervalPosition,
			"top":               ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.Top,
			"auth_result":       ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.AuthResult,
			"total":             ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.Total,
			"log_list":          setSliceLoggingLocalLogAuthenticationTopNOperOperLogList(ret.DtLoggingLocalLogAuthenticationTopNOper.Oper.LogList),
		},
	}
}

func setSliceLoggingLocalLogAuthenticationTopNOperOperLogList(d []edpt.LoggingLocalLogAuthenticationTopNOperOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func getObjectLoggingLocalLogAuthenticationTopNOperOper(d []interface{}) edpt.LoggingLocalLogAuthenticationTopNOperOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAuthenticationTopNOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxEntries = in["max_entries"].(int)
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.Top = in["top"].(string)
		ret.AuthResult = in["auth_result"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAuthenticationTopNOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAuthenticationTopNOperOperLogList(d []interface{}) []edpt.LoggingLocalLogAuthenticationTopNOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAuthenticationTopNOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAuthenticationTopNOperOperLogList
		oi.Name = in["name"].(string)
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogAuthenticationTopNOper(d *schema.ResourceData) edpt.LoggingLocalLogAuthenticationTopNOper {
	var ret edpt.LoggingLocalLogAuthenticationTopNOper

	ret.Oper = getObjectLoggingLocalLogAuthenticationTopNOperOper(d.Get("oper").([]interface{}))
	return ret
}
