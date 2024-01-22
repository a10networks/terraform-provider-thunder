package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTelemetryLogTopKAppSvcListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_telemetry_log_top_k_app_svc_list_oper`: Operational Status for the object top-k-app-svc-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTelemetryLogTopKAppSvcListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top_k_app_svc_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicator": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_svc_uuid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"app_svc_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicator_value": {
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

func resourceSystemTelemetryLogTopKAppSvcListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTelemetryLogTopKAppSvcListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTelemetryLogTopKAppSvcListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTelemetryLogTopKAppSvcListOperOper := setObjectSystemTelemetryLogTopKAppSvcListOperOper(res)
		d.Set("oper", SystemTelemetryLogTopKAppSvcListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTelemetryLogTopKAppSvcListOperOper(ret edpt.DataSystemTelemetryLogTopKAppSvcListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"top_k_app_svc_list": setSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList(ret.DtSystemTelemetryLogTopKAppSvcListOper.Oper.TopKAppSvcList),
		},
	}
}

func setSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList(d []edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator"] = item.Indicator
		in["indicator_list"] = setSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList(item.IndicatorList)
		result = append(result, in)
	}
	return result
}

func setSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList(d []edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["app_svc_uuid"] = item.AppSvcUuid
		in["app_svc_name"] = item.AppSvcName
		in["indicator_value"] = item.IndicatorValue
		result = append(result, in)
	}
	return result
}

func getObjectSystemTelemetryLogTopKAppSvcListOperOper(d []interface{}) edpt.SystemTelemetryLogTopKAppSvcListOperOper {

	count1 := len(d)
	var ret edpt.SystemTelemetryLogTopKAppSvcListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TopKAppSvcList = getSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList(in["top_k_app_svc_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList(d []interface{}) []edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList {

	count1 := len(d)
	ret := make([]edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList
		oi.Indicator = in["indicator"].(string)
		oi.IndicatorList = getSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList(d []interface{}) []edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList
		oi.AppSvcUuid = in["app_svc_uuid"].(string)
		oi.AppSvcName = in["app_svc_name"].(string)
		oi.IndicatorValue = in["indicator_value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemTelemetryLogTopKAppSvcListOper(d *schema.ResourceData) edpt.SystemTelemetryLogTopKAppSvcListOper {
	var ret edpt.SystemTelemetryLogTopKAppSvcListOper

	ret.Oper = getObjectSystemTelemetryLogTopKAppSvcListOperOper(d.Get("oper").([]interface{}))
	return ret
}
