package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAuthenticationLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_authentication_log_oper`: Operational Status for the object log\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogAuthenticationLogOperRead,

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
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user_domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"client_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"destination_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"destination_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auth_result": {
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

func resourceLoggingLocalLogAuthenticationLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAuthenticationLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAuthenticationLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogAuthenticationLogOperOper := setObjectLoggingLocalLogAuthenticationLogOperOper(res)
		d.Set("oper", LoggingLocalLogAuthenticationLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogAuthenticationLogOperOper(ret edpt.DataLoggingLocalLogAuthenticationLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"max_entries":       ret.DtLoggingLocalLogAuthenticationLogOper.Oper.MaxEntries,
			"start_time":        ret.DtLoggingLocalLogAuthenticationLogOper.Oper.StartTime,
			"interval":          ret.DtLoggingLocalLogAuthenticationLogOper.Oper.Interval,
			"interval_position": ret.DtLoggingLocalLogAuthenticationLogOper.Oper.IntervalPosition,
			"total":             ret.DtLoggingLocalLogAuthenticationLogOper.Oper.Total,
			"log_list":          setSliceLoggingLocalLogAuthenticationLogOperOperLogList(ret.DtLoggingLocalLogAuthenticationLogOper.Oper.LogList),
		},
	}
}

func setSliceLoggingLocalLogAuthenticationLogOperOperLogList(d []edpt.LoggingLocalLogAuthenticationLogOperOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["user_name"] = item.UserName
		in["user_domain"] = item.UserDomain
		in["client_ip"] = item.ClientIp
		in["destination_ip"] = item.DestinationIp
		in["source_port"] = item.SourcePort
		in["destination_port"] = item.DestinationPort
		in["auth_result"] = item.AuthResult
		result = append(result, in)
	}
	return result
}

func getObjectLoggingLocalLogAuthenticationLogOperOper(d []interface{}) edpt.LoggingLocalLogAuthenticationLogOperOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAuthenticationLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxEntries = in["max_entries"].(int)
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAuthenticationLogOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAuthenticationLogOperOperLogList(d []interface{}) []edpt.LoggingLocalLogAuthenticationLogOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAuthenticationLogOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAuthenticationLogOperOperLogList
		oi.Time = in["time"].(string)
		oi.UserName = in["user_name"].(string)
		oi.UserDomain = in["user_domain"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.DestinationIp = in["destination_ip"].(string)
		oi.SourcePort = in["source_port"].(int)
		oi.DestinationPort = in["destination_port"].(int)
		oi.AuthResult = in["auth_result"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogAuthenticationLogOper(d *schema.ResourceData) edpt.LoggingLocalLogAuthenticationLogOper {
	var ret edpt.LoggingLocalLogAuthenticationLogOper

	ret.Oper = getObjectLoggingLocalLogAuthenticationLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
