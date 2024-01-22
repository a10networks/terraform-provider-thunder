package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTelemetryLogTopKSourceListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_telemetry_log_top_k_source_list_oper`: Operational Status for the object top-k-source-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTelemetryLogTopKSourceListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top_k_source_list": {
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
												"ip_address": {
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

func resourceSystemTelemetryLogTopKSourceListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTelemetryLogTopKSourceListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTelemetryLogTopKSourceListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTelemetryLogTopKSourceListOperOper := setObjectSystemTelemetryLogTopKSourceListOperOper(res)
		d.Set("oper", SystemTelemetryLogTopKSourceListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTelemetryLogTopKSourceListOperOper(ret edpt.DataSystemTelemetryLogTopKSourceListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"top_k_source_list": setSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceList(ret.DtSystemTelemetryLogTopKSourceListOper.Oper.TopKSourceList),
		},
	}
}

func setSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceList(d []edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator"] = item.Indicator
		in["indicator_list"] = setSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList(item.IndicatorList)
		result = append(result, in)
	}
	return result
}

func setSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList(d []edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_address"] = item.IpAddress
		in["indicator_value"] = item.IndicatorValue
		result = append(result, in)
	}
	return result
}

func getObjectSystemTelemetryLogTopKSourceListOperOper(d []interface{}) edpt.SystemTelemetryLogTopKSourceListOperOper {

	count1 := len(d)
	var ret edpt.SystemTelemetryLogTopKSourceListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TopKSourceList = getSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceList(in["top_k_source_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceList(d []interface{}) []edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceList {

	count1 := len(d)
	ret := make([]edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceList
		oi.Indicator = in["indicator"].(string)
		oi.IndicatorList = getSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList(d []interface{}) []edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList
		oi.IpAddress = in["ip_address"].(string)
		oi.IndicatorValue = in["indicator_value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemTelemetryLogTopKSourceListOper(d *schema.ResourceData) edpt.SystemTelemetryLogTopKSourceListOper {
	var ret edpt.SystemTelemetryLogTopKSourceListOper

	ret.Oper = getObjectSystemTelemetryLogTopKSourceListOperOper(d.Get("oper").([]interface{}))
	return ret
}
