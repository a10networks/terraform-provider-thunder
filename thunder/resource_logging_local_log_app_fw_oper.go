package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAppFwOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_app_fw_oper`: Operational Status for the object app-fw\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogAppFwOperRead,

		Schema: map[string]*schema.Schema{
			"dot_plot": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
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
						"source_threat_category_match": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"destination_threat_category_match": {
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
									"application_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"category": {
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
									"source_threat_list_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"destination_threat_list_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source_threat_category": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"destination_threat_category": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"policy_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rule_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"top_n": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
		},
	}
}

func resourceLoggingLocalLogAppFwOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAppFwOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAppFwOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogAppFwOperDotPlot := setObjectLoggingLocalLogAppFwOperDotPlot(res)
		d.Set("dot_plot", LoggingLocalLogAppFwOperDotPlot)
		LoggingLocalLogAppFwOperOper := setObjectLoggingLocalLogAppFwOperOper(res)
		d.Set("oper", LoggingLocalLogAppFwOperOper)
		LoggingLocalLogAppFwOperTopN := setObjectLoggingLocalLogAppFwOperTopN(res)
		d.Set("top_n", LoggingLocalLogAppFwOperTopN)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogAppFwOperDotPlot(ret edpt.DataLoggingLocalLogAppFwOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectLoggingLocalLogAppFwOperDotPlotOper(ret.DtLoggingLocalLogAppFwOper.DotPlot.Oper),
		},
	}
}

func setObjectLoggingLocalLogAppFwOperDotPlotOper(d edpt.LoggingLocalLogAppFwOperDotPlotOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["start_time"] = d.StartTime

	in["interval"] = d.Interval

	in["interval_position"] = d.IntervalPosition

	in["application_name"] = d.ApplicationName

	in["client_ip"] = d.ClientIp

	in["data"] = d.Data

	in["total"] = d.Total
	in["log_list"] = setSliceLoggingLocalLogAppFwOperDotPlotOperLogList(d.LogList)
	result = append(result, in)
	return result
}

func setSliceLoggingLocalLogAppFwOperDotPlotOperLogList(d []edpt.LoggingLocalLogAppFwOperDotPlotOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func setObjectLoggingLocalLogAppFwOperOper(ret edpt.DataLoggingLocalLogAppFwOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"max_entries":                       ret.DtLoggingLocalLogAppFwOper.Oper.MaxEntries,
			"start_time":                        ret.DtLoggingLocalLogAppFwOper.Oper.StartTime,
			"interval":                          ret.DtLoggingLocalLogAppFwOper.Oper.Interval,
			"interval_position":                 ret.DtLoggingLocalLogAppFwOper.Oper.IntervalPosition,
			"source_threat_category_match":      ret.DtLoggingLocalLogAppFwOper.Oper.SourceThreatCategoryMatch,
			"destination_threat_category_match": ret.DtLoggingLocalLogAppFwOper.Oper.DestinationThreatCategoryMatch,
			"total":                             ret.DtLoggingLocalLogAppFwOper.Oper.Total,
			"log_list":                          setSliceLoggingLocalLogAppFwOperOperLogList(ret.DtLoggingLocalLogAppFwOper.Oper.LogList),
		},
	}
}

func setSliceLoggingLocalLogAppFwOperOperLogList(d []edpt.LoggingLocalLogAppFwOperOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["application_name"] = item.ApplicationName
		in["category"] = item.Category
		in["client_ip"] = item.ClientIp
		in["destination_ip"] = item.DestinationIp
		in["source_port"] = item.SourcePort
		in["destination_port"] = item.DestinationPort
		in["source_threat_list_name"] = item.SourceThreatListName
		in["destination_threat_list_name"] = item.DestinationThreatListName
		in["source_threat_category"] = item.SourceThreatCategory
		in["destination_threat_category"] = item.DestinationThreatCategory
		in["action"] = item.Action
		in["policy_name"] = item.PolicyName
		in["rule_name"] = item.RuleName
		in["bytes"] = item.Bytes
		result = append(result, in)
	}
	return result
}

func setObjectLoggingLocalLogAppFwOperTopN(ret edpt.DataLoggingLocalLogAppFwOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectLoggingLocalLogAppFwOperTopNOper(ret.DtLoggingLocalLogAppFwOper.TopN.Oper),
		},
	}
}

func setObjectLoggingLocalLogAppFwOperTopNOper(d edpt.LoggingLocalLogAppFwOperTopNOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["max_entries"] = d.MaxEntries

	in["start_time"] = d.StartTime

	in["interval"] = d.Interval

	in["interval_position"] = d.IntervalPosition

	in["top"] = d.Top

	in["application_name"] = d.ApplicationName

	in["category"] = d.Category

	in["client_ip"] = d.ClientIp

	in["threat_category"] = d.ThreatCategory

	in["threat_category_match"] = d.ThreatCategoryMatch

	in["action"] = d.Action

	in["total"] = d.Total
	in["log_list"] = setSliceLoggingLocalLogAppFwOperTopNOperLogList(d.LogList)
	result = append(result, in)
	return result
}

func setSliceLoggingLocalLogAppFwOperTopNOperLogList(d []edpt.LoggingLocalLogAppFwOperTopNOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func getObjectLoggingLocalLogAppFwOperDotPlot(d []interface{}) edpt.LoggingLocalLogAppFwOperDotPlot {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAppFwOperDotPlot
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectLoggingLocalLogAppFwOperDotPlotOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectLoggingLocalLogAppFwOperDotPlotOper(d []interface{}) edpt.LoggingLocalLogAppFwOperDotPlotOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAppFwOperDotPlotOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.ApplicationName = in["application_name"].(string)
		ret.ClientIp = in["client_ip"].(string)
		ret.Data = in["data"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAppFwOperDotPlotOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAppFwOperDotPlotOperLogList(d []interface{}) []edpt.LoggingLocalLogAppFwOperDotPlotOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAppFwOperDotPlotOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAppFwOperDotPlotOperLogList
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectLoggingLocalLogAppFwOperOper(d []interface{}) edpt.LoggingLocalLogAppFwOperOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAppFwOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxEntries = in["max_entries"].(int)
		ret.StartTime = in["start_time"].(string)
		ret.Interval = in["interval"].(string)
		ret.IntervalPosition = in["interval_position"].(string)
		ret.SourceThreatCategoryMatch = in["source_threat_category_match"].(string)
		ret.DestinationThreatCategoryMatch = in["destination_threat_category_match"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceLoggingLocalLogAppFwOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAppFwOperOperLogList(d []interface{}) []edpt.LoggingLocalLogAppFwOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAppFwOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAppFwOperOperLogList
		oi.Time = in["time"].(string)
		oi.ApplicationName = in["application_name"].(string)
		oi.Category = in["category"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.DestinationIp = in["destination_ip"].(string)
		oi.SourcePort = in["source_port"].(int)
		oi.DestinationPort = in["destination_port"].(int)
		oi.SourceThreatListName = in["source_threat_list_name"].(string)
		oi.DestinationThreatListName = in["destination_threat_list_name"].(string)
		oi.SourceThreatCategory = in["source_threat_category"].(string)
		oi.DestinationThreatCategory = in["destination_threat_category"].(string)
		oi.Action = in["action"].(string)
		oi.PolicyName = in["policy_name"].(string)
		oi.RuleName = in["rule_name"].(string)
		oi.Bytes = in["bytes"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectLoggingLocalLogAppFwOperTopN(d []interface{}) edpt.LoggingLocalLogAppFwOperTopN {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAppFwOperTopN
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectLoggingLocalLogAppFwOperTopNOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectLoggingLocalLogAppFwOperTopNOper(d []interface{}) edpt.LoggingLocalLogAppFwOperTopNOper {

	count1 := len(d)
	var ret edpt.LoggingLocalLogAppFwOperTopNOper
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
		ret.LogList = getSliceLoggingLocalLogAppFwOperTopNOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceLoggingLocalLogAppFwOperTopNOperLogList(d []interface{}) []edpt.LoggingLocalLogAppFwOperTopNOperLogList {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogAppFwOperTopNOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogAppFwOperTopNOperLogList
		oi.Name = in["name"].(string)
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogAppFwOper(d *schema.ResourceData) edpt.LoggingLocalLogAppFwOper {
	var ret edpt.LoggingLocalLogAppFwOper

	ret.DotPlot = getObjectLoggingLocalLogAppFwOperDotPlot(d.Get("dot_plot").([]interface{}))

	ret.Oper = getObjectLoggingLocalLogAppFwOperOper(d.Get("oper").([]interface{}))

	ret.TopN = getObjectLoggingLocalLogAppFwOperTopN(d.Get("top_n").([]interface{}))
	return ret
}
