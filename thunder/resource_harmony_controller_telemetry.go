package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerTelemetry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_harmony_controller_telemetry`: Harmony controller telemetry config\n\n__PLACEHOLDER__",
		CreateContext: resourceHarmonyControllerTelemetryCreate,
		UpdateContext: resourceHarmonyControllerTelemetryUpdate,
		ReadContext:   resourceHarmonyControllerTelemetryRead,
		DeleteContext: resourceHarmonyControllerTelemetryDelete,

		Schema: map[string]*schema.Schema{
			"log_rate": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Max number of session logs sent by the partition per second",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceHarmonyControllerTelemetryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerTelemetryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerTelemetry(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerTelemetryRead(ctx, d, meta)
	}
	return diags
}

func resourceHarmonyControllerTelemetryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerTelemetryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerTelemetry(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerTelemetryRead(ctx, d, meta)
	}
	return diags
}
func resourceHarmonyControllerTelemetryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerTelemetryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerTelemetry(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHarmonyControllerTelemetryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerTelemetryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerTelemetry(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHarmonyControllerTelemetry(d *schema.ResourceData) edpt.HarmonyControllerTelemetry {
	var ret edpt.HarmonyControllerTelemetry
	ret.Inst.LogRate = d.Get("log_rate").(int)
	//omit uuid
	return ret
}
