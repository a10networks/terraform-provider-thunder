package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProbeInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_controller_probe_info_oper`: Operational Status for the object probe-info\n\n__PLACEHOLDER__",
		ReadContext: resourceControllerProbeInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"probe_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"controller_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"data_showtech_export_log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"showtech_filename": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"showtech_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"showtech_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"data_varlog_export_log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"varlog_filename": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"varlog_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"varlog_timestamp": {
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

func resourceControllerProbeInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProbeInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProbeInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ControllerProbeInfoOperOper := setObjectControllerProbeInfoOperOper(res)
		d.Set("oper", ControllerProbeInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectControllerProbeInfoOperOper(ret edpt.DataControllerProbeInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"probe_status":                  ret.DtControllerProbeInfoOper.Oper.ProbeStatus,
			"controller_status":             ret.DtControllerProbeInfoOper.Oper.ControllerStatus,
			"data_showtech_export_log_list": setSliceControllerProbeInfoOperOperDataShowtechExportLogList(ret.DtControllerProbeInfoOper.Oper.DataShowtechExportLogList),
			"data_varlog_export_log_list":   setSliceControllerProbeInfoOperOperDataVarlogExportLogList(ret.DtControllerProbeInfoOper.Oper.DataVarlogExportLogList),
		},
	}
}

func setSliceControllerProbeInfoOperOperDataShowtechExportLogList(d []edpt.ControllerProbeInfoOperOperDataShowtechExportLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["showtech_filename"] = item.ShowtechFilename
		in["showtech_size"] = item.ShowtechSize
		in["showtech_timestamp"] = item.ShowtechTimestamp
		result = append(result, in)
	}
	return result
}

func setSliceControllerProbeInfoOperOperDataVarlogExportLogList(d []edpt.ControllerProbeInfoOperOperDataVarlogExportLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["varlog_filename"] = item.VarlogFilename
		in["varlog_size"] = item.VarlogSize
		in["varlog_timestamp"] = item.VarlogTimestamp
		result = append(result, in)
	}
	return result
}

func getObjectControllerProbeInfoOperOper(d []interface{}) edpt.ControllerProbeInfoOperOper {

	count1 := len(d)
	var ret edpt.ControllerProbeInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProbeStatus = in["probe_status"].(string)
		ret.ControllerStatus = in["controller_status"].(string)
		ret.DataShowtechExportLogList = getSliceControllerProbeInfoOperOperDataShowtechExportLogList(in["data_showtech_export_log_list"].([]interface{}))
		ret.DataVarlogExportLogList = getSliceControllerProbeInfoOperOperDataVarlogExportLogList(in["data_varlog_export_log_list"].([]interface{}))
	}
	return ret
}

func getSliceControllerProbeInfoOperOperDataShowtechExportLogList(d []interface{}) []edpt.ControllerProbeInfoOperOperDataShowtechExportLogList {

	count1 := len(d)
	ret := make([]edpt.ControllerProbeInfoOperOperDataShowtechExportLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ControllerProbeInfoOperOperDataShowtechExportLogList
		oi.ShowtechFilename = in["showtech_filename"].(string)
		oi.ShowtechSize = in["showtech_size"].(int)
		oi.ShowtechTimestamp = in["showtech_timestamp"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceControllerProbeInfoOperOperDataVarlogExportLogList(d []interface{}) []edpt.ControllerProbeInfoOperOperDataVarlogExportLogList {

	count1 := len(d)
	ret := make([]edpt.ControllerProbeInfoOperOperDataVarlogExportLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ControllerProbeInfoOperOperDataVarlogExportLogList
		oi.VarlogFilename = in["varlog_filename"].(string)
		oi.VarlogSize = in["varlog_size"].(int)
		oi.VarlogTimestamp = in["varlog_timestamp"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointControllerProbeInfoOper(d *schema.ResourceData) edpt.ControllerProbeInfoOper {
	var ret edpt.ControllerProbeInfoOper

	ret.Oper = getObjectControllerProbeInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
