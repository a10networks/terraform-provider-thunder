package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAccessLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_access_log_oper`: Operational Status for the object log\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogAccessLogOperRead,

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
									"policy": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vip_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vip_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"web_domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"uri": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"web_category": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"web_reputation": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ssl_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"reason": {
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

func resourceLoggingLocalLogAccessLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAccessLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAccessLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogAccessLogOperOper := setObjectLoggingLocalLogAccessLogOperOper(res)
		d.Set("oper", LoggingLocalLogAccessLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogAccessLogOperOper(ret edpt.DataLoggingLocalLogAccessLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"max_entries":       ret.DtLoggingLocalLogAccessLogOper.Oper.MaxEntries,
			"start_time":        ret.DtLoggingLocalLogAccessLogOper.Oper.StartTime,
			"interval":          ret.DtLoggingLocalLogAccessLogOper.Oper.Interval,
			"interval_position": ret.DtLoggingLocalLogAccessLogOper.Oper.IntervalPosition,
			"total":             ret.DtLoggingLocalLogAccessLogOper.Oper.Total,
			"log_list":          setSliceLoggingLocalLogAccessLogOperOperLogList(ret.DtLoggingLocalLogAccessLogOper.Oper.LogList),
		},
	}
}

func setSliceLoggingLocalLogAccessLogOperOperLogList(d []edpt.LoggingLocalLogAccessLogOperOperLogList) []map[string]interface{} {
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
		in["policy"] = item.Policy
		in["action"] = item.Action
		in["vip_name"] = item.VipName
		in["vip_port"] = item.VipPort
		in["web_domain"] = item.WebDomain
		in["uri"] = item.Uri
		in["web_category"] = item.WebCategory
		in["web_reputation"] = item.WebReputation
		in["ssl_status"] = item.SslStatus
		in["reason"] = item.Reason
		result = append(result, in)
	}
	return result
}

func getObjectLoggingLocalLogAccessLogOperOper(d []interface{}) edpt.LoggingLocalLogAccessLogOperOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAccessLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxEntries = in["max_entries"].(int)
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAccessLogOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAccessLogOperOperLogList(d []interface{}) []edpt.LoggingLocalLogAccessLogOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAccessLogOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAccessLogOperOperLogList
		oi.Time = in["time"].(string)
		oi.UserName = in["user_name"].(string)
		oi.UserDomain = in["user_domain"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.DestinationIp = in["destination_ip"].(string)
		oi.SourcePort = in["source_port"].(int)
		oi.DestinationPort = in["destination_port"].(int)
		oi.Policy = in["policy"].(string)
		oi.Action = in["action"].(string)
		oi.VipName = in["vip_name"].(string)
		oi.VipPort = in["vip_port"].(int)
		oi.WebDomain = in["web_domain"].(string)
		oi.Uri = in["uri"].(string)
		oi.WebCategory = in["web_category"].(string)
		oi.WebReputation = in["web_reputation"].(string)
		oi.SslStatus = in["ssl_status"].(string)
		oi.Reason = in["reason"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogAccessLogOper(d *schema.ResourceData) edpt.LoggingLocalLogAccessLogOper {
	var ret edpt.LoggingLocalLogAccessLogOper

	ret.Oper = getObjectLoggingLocalLogAccessLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
