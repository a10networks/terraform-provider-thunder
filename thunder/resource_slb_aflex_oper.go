package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbAflexOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_aflex_oper`: Operational Status for the object aflex\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbAflexOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter_event": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter_debug": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_substring": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_file_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"syntax": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vport_bound": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vport_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vserver": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"events": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"event_type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"total_executions": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"failures": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"aborts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"thread_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"thread": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_msg": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"error_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"event_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"cmd_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"err_msg": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"line_number": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"file_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"aflex_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbAflexOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAflexOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAflexOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbAflexOperOper := setObjectSlbAflexOperOper(res)
		d.Set("oper", SlbAflexOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbAflexOperOper(ret edpt.DataSlbAflexOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter_entry":        ret.DtSlbAflexOper.Oper.Filter_entry,
			"filter_event":        ret.DtSlbAflexOper.Oper.Filter_event,
			"filter_debug":        ret.DtSlbAflexOper.Oper.Filter_debug,
			"filter_substring":    ret.DtSlbAflexOper.Oper.Filter_substring,
			"aflex_file_size_max": ret.DtSlbAflexOper.Oper.AflexFileSizeMax,
			"file_list":           setSliceSlbAflexOperOperFileList(ret.DtSlbAflexOper.Oper.FileList),
			"thread_list":         setSliceSlbAflexOperOperThreadList(ret.DtSlbAflexOper.Oper.ThreadList),
			"aflex_count":         ret.DtSlbAflexOper.Oper.AflexCount,
		},
	}
}

func setSliceSlbAflexOperOperFileList(d []edpt.SlbAflexOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["syntax"] = item.Syntax
		in["vport_bound"] = item.Vport_bound
		in["vport_list"] = setSliceSlbAflexOperOperFileListVportList(item.VportList)
		in["events"] = setSliceSlbAflexOperOperFileListEvents(item.Events)
		result = append(result, in)
	}
	return result
}

func setSliceSlbAflexOperOperFileListVportList(d []edpt.SlbAflexOperOperFileListVportList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vserver"] = item.Vserver
		in["port"] = item.Port
		result = append(result, in)
	}
	return result
}

func setSliceSlbAflexOperOperFileListEvents(d []edpt.SlbAflexOperOperFileListEvents) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["event_type"] = item.EventType
		in["total_executions"] = item.TotalExecutions
		in["failures"] = item.Failures
		in["aborts"] = item.Aborts
		result = append(result, in)
	}
	return result
}

func setSliceSlbAflexOperOperThreadList(d []edpt.SlbAflexOperOperThreadList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["thread"] = item.Thread
		in["err_line"] = item.ErrLine
		in["last_msg"] = item.LastMsg
		in["error_list"] = setSliceSlbAflexOperOperThreadListErrorList(item.ErrorList)
		result = append(result, in)
	}
	return result
}

func setSliceSlbAflexOperOperThreadListErrorList(d []edpt.SlbAflexOperOperThreadListErrorList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["event_name"] = item.EventName
		in["cmd_name"] = item.CmdName
		in["err_msg"] = item.ErrMsg
		in["line_number"] = item.LineNumber
		in["file_name"] = item.FileName
		result = append(result, in)
	}
	return result
}

func getObjectSlbAflexOperOper(d []interface{}) edpt.SlbAflexOperOper {

	count1 := len(d)
	var ret edpt.SlbAflexOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Filter_entry = in["filter_entry"].(string)
		ret.Filter_event = in["filter_event"].(string)
		ret.Filter_debug = in["filter_debug"].(int)
		ret.Filter_substring = in["filter_substring"].(int)
		ret.AflexFileSizeMax = in["aflex_file_size_max"].(int)
		ret.FileList = getSliceSlbAflexOperOperFileList(in["file_list"].([]interface{}))
		ret.ThreadList = getSliceSlbAflexOperOperThreadList(in["thread_list"].([]interface{}))
		ret.AflexCount = in["aflex_count"].(int)
	}
	return ret
}

func getSliceSlbAflexOperOperFileList(d []interface{}) []edpt.SlbAflexOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.SlbAflexOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAflexOperOperFileList
		oi.File = in["file"].(string)
		oi.Syntax = in["syntax"].(string)
		oi.Vport_bound = in["vport_bound"].(string)
		oi.VportList = getSliceSlbAflexOperOperFileListVportList(in["vport_list"].([]interface{}))
		oi.Events = getSliceSlbAflexOperOperFileListEvents(in["events"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbAflexOperOperFileListVportList(d []interface{}) []edpt.SlbAflexOperOperFileListVportList {

	count1 := len(d)
	ret := make([]edpt.SlbAflexOperOperFileListVportList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAflexOperOperFileListVportList
		oi.Vserver = in["vserver"].(string)
		oi.Port = in["port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbAflexOperOperFileListEvents(d []interface{}) []edpt.SlbAflexOperOperFileListEvents {

	count1 := len(d)
	ret := make([]edpt.SlbAflexOperOperFileListEvents, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAflexOperOperFileListEvents
		oi.EventType = in["event_type"].(string)
		oi.TotalExecutions = in["total_executions"].(int)
		oi.Failures = in["failures"].(int)
		oi.Aborts = in["aborts"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbAflexOperOperThreadList(d []interface{}) []edpt.SlbAflexOperOperThreadList {

	count1 := len(d)
	ret := make([]edpt.SlbAflexOperOperThreadList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAflexOperOperThreadList
		oi.Thread = in["thread"].(int)
		oi.ErrLine = in["err_line"].(int)
		oi.LastMsg = in["last_msg"].(string)
		oi.ErrorList = getSliceSlbAflexOperOperThreadListErrorList(in["error_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbAflexOperOperThreadListErrorList(d []interface{}) []edpt.SlbAflexOperOperThreadListErrorList {

	count1 := len(d)
	ret := make([]edpt.SlbAflexOperOperThreadListErrorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAflexOperOperThreadListErrorList
		oi.EventName = in["event_name"].(string)
		oi.CmdName = in["cmd_name"].(string)
		oi.ErrMsg = in["err_msg"].(string)
		oi.LineNumber = in["line_number"].(int)
		oi.FileName = in["file_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbAflexOper(d *schema.ResourceData) edpt.SlbAflexOper {
	var ret edpt.SlbAflexOper

	ret.Oper = getObjectSlbAflexOperOper(d.Get("oper").([]interface{}))
	return ret
}
