package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTelemetryLogDeviceStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_telemetry_log_device_status_oper`: Operational Status for the object device-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTelemetryLogDeviceStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"memory_usage": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"control_cpu_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cpu_usage_overall": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ratio_session_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ratio_buffer_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_bytes_in": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_bytes_out": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTelemetryLogDeviceStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTelemetryLogDeviceStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTelemetryLogDeviceStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTelemetryLogDeviceStatusOperOper := setObjectSystemTelemetryLogDeviceStatusOperOper(res)
		d.Set("oper", SystemTelemetryLogDeviceStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTelemetryLogDeviceStatusOperOper(ret edpt.DataSystemTelemetryLogDeviceStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"memory_usage":        ret.DtSystemTelemetryLogDeviceStatusOper.Oper.Memory_usage,
			"control_cpu_usage":   ret.DtSystemTelemetryLogDeviceStatusOper.Oper.Control_cpu_usage,
			"cpu_usage_overall":   ret.DtSystemTelemetryLogDeviceStatusOper.Oper.Cpu_usage_overall,
			"ratio_session_count": ret.DtSystemTelemetryLogDeviceStatusOper.Oper.Ratio_session_count,
			"ratio_buffer_count":  ret.DtSystemTelemetryLogDeviceStatusOper.Oper.Ratio_buffer_count,
			"total_bytes_in":      ret.DtSystemTelemetryLogDeviceStatusOper.Oper.Total_bytes_in,
			"total_bytes_out":     ret.DtSystemTelemetryLogDeviceStatusOper.Oper.Total_bytes_out,
		},
	}
}

func getObjectSystemTelemetryLogDeviceStatusOperOper(d []interface{}) edpt.SystemTelemetryLogDeviceStatusOperOper {

	count1 := len(d)
	var ret edpt.SystemTelemetryLogDeviceStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Memory_usage = in["memory_usage"].(string)
		ret.Control_cpu_usage = in["control_cpu_usage"].(int)
		ret.Cpu_usage_overall = in["cpu_usage_overall"].(int)
		ret.Ratio_session_count = in["ratio_session_count"].(string)
		ret.Ratio_buffer_count = in["ratio_buffer_count"].(string)
		ret.Total_bytes_in = in["total_bytes_in"].(int)
		ret.Total_bytes_out = in["total_bytes_out"].(int)
	}
	return ret
}

func dataToEndpointSystemTelemetryLogDeviceStatusOper(d *schema.ResourceData) edpt.SystemTelemetryLogDeviceStatusOper {
	var ret edpt.SystemTelemetryLogDeviceStatusOper

	ret.Oper = getObjectSystemTelemetryLogDeviceStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
