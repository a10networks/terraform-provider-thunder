package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerTelemetry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_telemetry`: A10 Control telemetry config\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerTelemetryCreate,
		UpdateContext: resourceControllerTelemetryUpdate,
		ReadContext:   resourceControllerTelemetryRead,
		DeleteContext: resourceControllerTelemetryDelete,

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
func resourceControllerTelemetryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetry(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerTelemetryRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerTelemetryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetry(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerTelemetryRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerTelemetryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetry(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerTelemetryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetry(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointControllerTelemetry(d *schema.ResourceData) edpt.ControllerTelemetry {
	var ret edpt.ControllerTelemetry
	if v, ok := d.GetOk("log_rate"); ok {
		ret.Telemetry.LogRate = v.(int)
	} else {
		ret.Telemetry.LogRate = 10
	}
	//omit uuid
	return ret
}
