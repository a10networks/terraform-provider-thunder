package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTelemetryLogEnvironmentOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_telemetry_log_environment_oper`: Operational Status for the object environment\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTelemetryLogEnvironmentOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"physical_system_temperature": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_fan_setting": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"system_voltage": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"power_unit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTelemetryLogEnvironmentOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTelemetryLogEnvironmentOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTelemetryLogEnvironmentOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTelemetryLogEnvironmentOperOper := setObjectSystemTelemetryLogEnvironmentOperOper(res)
		d.Set("oper", SystemTelemetryLogEnvironmentOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTelemetryLogEnvironmentOperOper(ret edpt.DataSystemTelemetryLogEnvironmentOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"physical_system_temperature": ret.DtSystemTelemetryLogEnvironmentOper.Oper.PhysicalSystemTemperature,
			"hw_fan_setting":              ret.DtSystemTelemetryLogEnvironmentOper.Oper.HwFanSetting,
			"fan_report":                  ret.DtSystemTelemetryLogEnvironmentOper.Oper.FanReport,
			"fan_value":                   ret.DtSystemTelemetryLogEnvironmentOper.Oper.FanValue,
			"system_voltage":              ret.DtSystemTelemetryLogEnvironmentOper.Oper.SystemVoltage,
			"power_unit":                  ret.DtSystemTelemetryLogEnvironmentOper.Oper.PowerUnit,
		},
	}
}

func getObjectSystemTelemetryLogEnvironmentOperOper(d []interface{}) edpt.SystemTelemetryLogEnvironmentOperOper {

	count1 := len(d)
	var ret edpt.SystemTelemetryLogEnvironmentOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PhysicalSystemTemperature = in["physical_system_temperature"].(int)
		ret.HwFanSetting = in["hw_fan_setting"].(string)
		ret.FanReport = in["fan_report"].(string)
		ret.FanValue = in["fan_value"].(int)
		ret.SystemVoltage = in["system_voltage"].(string)
		ret.PowerUnit = in["power_unit"].(string)
	}
	return ret
}

func dataToEndpointSystemTelemetryLogEnvironmentOper(d *schema.ResourceData) edpt.SystemTelemetryLogEnvironmentOper {
	var ret edpt.SystemTelemetryLogEnvironmentOper

	ret.Oper = getObjectSystemTelemetryLogEnvironmentOperOper(d.Get("oper").([]interface{}))
	return ret
}
