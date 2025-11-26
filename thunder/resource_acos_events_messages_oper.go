package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsMessagesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_acos_events_messages_oper`: Operational Status for the object messages\n\n__PLACEHOLDER__",
		ReadContext: resourceAcosEventsMessagesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_msgs": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"active_msgs": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"debug_msgs": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"message_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"message_id_scope": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_log_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"active_log_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"messages_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"syslog": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"syslog_collector_grps_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"cef": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cef_collector_grps_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"leef": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"leef_collector_grps_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"message_route": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"enabled": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"disabled": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"not_defined": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"collector_group": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"descriptive_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"signature": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"trigger_reason": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"defined_severity": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"configured_severity": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cef_importance": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rate_limiting": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"default_route": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"configured_route": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"log_detail_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"log_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_logging_conf": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_logging_conf": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"collector_groups_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"log_format": {
													Type: schema.TypeString, Optional: true, Description: "",
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

func resourceAcosEventsMessagesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessagesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessagesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AcosEventsMessagesOperOper := setObjectAcosEventsMessagesOperOper(res)
		d.Set("oper", AcosEventsMessagesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAcosEventsMessagesOperOper(ret edpt.DataAcosEventsMessagesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all_msgs":         ret.DtAcosEventsMessagesOper.Oper.AllMsgs,
			"active_msgs":      ret.DtAcosEventsMessagesOper.Oper.ActiveMsgs,
			"debug_msgs":       ret.DtAcosEventsMessagesOper.Oper.DebugMsgs,
			"message_id":       ret.DtAcosEventsMessagesOper.Oper.MessageId,
			"message_id_scope": ret.DtAcosEventsMessagesOper.Oper.MessageIdScope,
			"total_log_count":  ret.DtAcosEventsMessagesOper.Oper.TotalLogCount,
			"active_log_count": ret.DtAcosEventsMessagesOper.Oper.ActiveLogCount,
			"messages_list":    setSliceAcosEventsMessagesOperOperMessagesList(ret.DtAcosEventsMessagesOper.Oper.MessagesList),
		},
	}
}

func setSliceAcosEventsMessagesOperOperMessagesList(d []edpt.AcosEventsMessagesOperOperMessagesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["message_id"] = item.MessageId
		in["syslog"] = item.Syslog
		in["syslog_collector_grps_list"] = setSliceAcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList(item.SyslogCollectorGrpsList)
		in["cef"] = item.Cef
		in["cef_collector_grps_list"] = setSliceAcosEventsMessagesOperOperMessagesListCefCollectorGrpsList(item.CefCollectorGrpsList)
		in["leef"] = item.Leef
		in["leef_collector_grps_list"] = setSliceAcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList(item.LeefCollectorGrpsList)
		in["message_route"] = item.MessageRoute
		in["enabled"] = item.Enabled
		in["disabled"] = item.Disabled
		in["not_defined"] = item.NotDefined
		in["local"] = item.Local
		in["remote"] = item.Remote
		in["collector_group"] = item.CollectorGroup
		in["descriptive_name"] = item.DescriptiveName
		in["signature"] = item.Signature
		in["trigger_reason"] = item.TriggerReason
		in["defined_severity"] = item.DefinedSeverity
		in["configured_severity"] = item.ConfiguredSeverity
		in["cef_importance"] = item.CefImportance
		in["rate_limiting"] = item.RateLimiting
		in["default_route"] = item.DefaultRoute
		in["configured_route"] = item.ConfiguredRoute
		in["log_detail_list"] = setSliceAcosEventsMessagesOperOperMessagesListLogDetailList(item.LogDetailList)
		result = append(result, in)
	}
	return result
}

func setSliceAcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList(d []edpt.AcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func setSliceAcosEventsMessagesOperOperMessagesListCefCollectorGrpsList(d []edpt.AcosEventsMessagesOperOperMessagesListCefCollectorGrpsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func setSliceAcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList(d []edpt.AcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func setSliceAcosEventsMessagesOperOperMessagesListLogDetailList(d []edpt.AcosEventsMessagesOperOperMessagesListLogDetailList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["log_name"] = item.LogName
		in["local_logging_conf"] = item.LocalLoggingConf
		in["remote_logging_conf"] = item.RemoteLoggingConf
		in["collector_groups_list"] = setSliceAcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList(item.CollectorGroupsList)
		in["log_format"] = item.LogFormat
		result = append(result, in)
	}
	return result
}

func setSliceAcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList(d []edpt.AcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func getObjectAcosEventsMessagesOperOper(d []interface{}) edpt.AcosEventsMessagesOperOper {

	count1 := len(d)
	var ret edpt.AcosEventsMessagesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllMsgs = in["all_msgs"].(int)
		ret.ActiveMsgs = in["active_msgs"].(int)
		ret.DebugMsgs = in["debug_msgs"].(int)
		ret.MessageId = in["message_id"].(string)
		ret.MessageIdScope = in["message_id_scope"].(string)
		ret.TotalLogCount = in["total_log_count"].(int)
		ret.ActiveLogCount = in["active_log_count"].(int)
		ret.MessagesList = getSliceAcosEventsMessagesOperOperMessagesList(in["messages_list"].([]interface{}))
	}
	return ret
}

func getSliceAcosEventsMessagesOperOperMessagesList(d []interface{}) []edpt.AcosEventsMessagesOperOperMessagesList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsMessagesOperOperMessagesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsMessagesOperOperMessagesList
		oi.MessageId = in["message_id"].(string)
		oi.Syslog = in["syslog"].(string)
		oi.SyslogCollectorGrpsList = getSliceAcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList(in["syslog_collector_grps_list"].([]interface{}))
		oi.Cef = in["cef"].(string)
		oi.CefCollectorGrpsList = getSliceAcosEventsMessagesOperOperMessagesListCefCollectorGrpsList(in["cef_collector_grps_list"].([]interface{}))
		oi.Leef = in["leef"].(string)
		oi.LeefCollectorGrpsList = getSliceAcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList(in["leef_collector_grps_list"].([]interface{}))
		oi.MessageRoute = in["message_route"].(string)
		oi.Enabled = in["enabled"].(int)
		oi.Disabled = in["disabled"].(int)
		oi.NotDefined = in["not_defined"].(int)
		oi.Local = in["local"].(string)
		oi.Remote = in["remote"].(string)
		oi.CollectorGroup = in["collector_group"].(string)
		oi.DescriptiveName = in["descriptive_name"].(string)
		oi.Signature = in["signature"].(string)
		oi.TriggerReason = in["trigger_reason"].(string)
		oi.DefinedSeverity = in["defined_severity"].(int)
		oi.ConfiguredSeverity = in["configured_severity"].(int)
		oi.CefImportance = in["cef_importance"].(int)
		oi.RateLimiting = in["rate_limiting"].(string)
		oi.DefaultRoute = in["default_route"].(string)
		oi.ConfiguredRoute = in["configured_route"].(string)
		oi.LogDetailList = getSliceAcosEventsMessagesOperOperMessagesListLogDetailList(in["log_detail_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList(d []interface{}) []edpt.AcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsMessagesOperOperMessagesListCefCollectorGrpsList(d []interface{}) []edpt.AcosEventsMessagesOperOperMessagesListCefCollectorGrpsList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsMessagesOperOperMessagesListCefCollectorGrpsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsMessagesOperOperMessagesListCefCollectorGrpsList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList(d []interface{}) []edpt.AcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsMessagesOperOperMessagesListLogDetailList(d []interface{}) []edpt.AcosEventsMessagesOperOperMessagesListLogDetailList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsMessagesOperOperMessagesListLogDetailList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsMessagesOperOperMessagesListLogDetailList
		oi.LogName = in["log_name"].(string)
		oi.LocalLoggingConf = in["local_logging_conf"].(string)
		oi.RemoteLoggingConf = in["remote_logging_conf"].(string)
		oi.CollectorGroupsList = getSliceAcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList(in["collector_groups_list"].([]interface{}))
		oi.LogFormat = in["log_format"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList(d []interface{}) []edpt.AcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsMessagesOper(d *schema.ResourceData) edpt.AcosEventsMessagesOper {
	var ret edpt.AcosEventsMessagesOper

	ret.Oper = getObjectAcosEventsMessagesOperOper(d.Get("oper").([]interface{}))
	return ret
}
