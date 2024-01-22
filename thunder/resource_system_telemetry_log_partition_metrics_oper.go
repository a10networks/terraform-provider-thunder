package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTelemetryLogPartitionMetricsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_telemetry_log_partition_metrics_oper`: Operational Status for the object partition-metrics\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTelemetryLogPartitionMetricsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_cpu_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTelemetryLogPartitionMetricsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTelemetryLogPartitionMetricsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTelemetryLogPartitionMetricsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTelemetryLogPartitionMetricsOperOper := setObjectSystemTelemetryLogPartitionMetricsOperOper(res)
		d.Set("oper", SystemTelemetryLogPartitionMetricsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTelemetryLogPartitionMetricsOperOper(ret edpt.DataSystemTelemetryLogPartitionMetricsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"data_cpu_usage": ret.DtSystemTelemetryLogPartitionMetricsOper.Oper.Data_cpu_usage,
		},
	}
}

func getObjectSystemTelemetryLogPartitionMetricsOperOper(d []interface{}) edpt.SystemTelemetryLogPartitionMetricsOperOper {

	count1 := len(d)
	var ret edpt.SystemTelemetryLogPartitionMetricsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Data_cpu_usage = in["data_cpu_usage"].(int)
	}
	return ret
}

func dataToEndpointSystemTelemetryLogPartitionMetricsOper(d *schema.ResourceData) edpt.SystemTelemetryLogPartitionMetricsOper {
	var ret edpt.SystemTelemetryLogPartitionMetricsOper

	ret.Oper = getObjectSystemTelemetryLogPartitionMetricsOperOper(d.Get("oper").([]interface{}))
	return ret
}
