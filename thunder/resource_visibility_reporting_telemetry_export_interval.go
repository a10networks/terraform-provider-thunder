package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityReportingTelemetryExportInterval() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_reporting_telemetry_export_interval`: Configure monitor entity telemetry data export interval\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityReportingTelemetryExportIntervalCreate,
		UpdateContext: resourceVisibilityReportingTelemetryExportIntervalUpdate,
		ReadContext:   resourceVisibilityReportingTelemetryExportIntervalRead,
		DeleteContext: resourceVisibilityReportingTelemetryExportIntervalDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Monitored entity telemetry data export interval in mins (Default 5 mins)",
			},
		},
	}
}
func resourceVisibilityReportingTelemetryExportIntervalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTelemetryExportIntervalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTelemetryExportInterval(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityReportingTelemetryExportIntervalRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityReportingTelemetryExportIntervalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTelemetryExportIntervalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTelemetryExportInterval(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityReportingTelemetryExportIntervalRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityReportingTelemetryExportIntervalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTelemetryExportIntervalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTelemetryExportInterval(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityReportingTelemetryExportIntervalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTelemetryExportIntervalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTelemetryExportInterval(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityReportingTelemetryExportInterval(d *schema.ResourceData) edpt.VisibilityReportingTelemetryExportInterval {
	var ret edpt.VisibilityReportingTelemetryExportInterval
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	return ret
}
