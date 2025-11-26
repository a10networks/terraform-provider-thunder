package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerTelemetryProbe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_telemetry_probe`: controller probe config\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerTelemetryProbeCreate,
		UpdateContext: resourceControllerTelemetryProbeUpdate,
		ReadContext:   resourceControllerTelemetryProbeRead,
		DeleteContext: resourceControllerTelemetryProbeDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the probe functionality; 'disable': Disable the probe functionality;",
			},
			"export_policy": {
				Type: schema.TypeString, Optional: true, Default: "snapshots-new", Description: "'snapshots-all': Export historical/missed snapshots.; 'snapshots-new': Export only new snapshots(default).;",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "snapshot export interval in minute,default is 15.",
			},
			"log_level": {
				Type: schema.TypeString, Optional: true, Default: "ERROR", Description: "'ERROR': show errors only(default).; 'WARNING': show warnings; 'INFO': show info messages; 'DEBUG': show debug logs;",
			},
			"target": {
				Type: schema.TypeString, Optional: true, Default: "remote", Description: "'remote': Export data to remote. This is the default value.; 'local': Export data local.;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceControllerTelemetryProbeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryProbeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetryProbe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerTelemetryProbeRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerTelemetryProbeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryProbeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetryProbe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerTelemetryProbeRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerTelemetryProbeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryProbeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetryProbe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerTelemetryProbeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTelemetryProbeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTelemetryProbe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointControllerTelemetryProbe(d *schema.ResourceData) edpt.ControllerTelemetryProbe {
	var ret edpt.ControllerTelemetryProbe
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ExportPolicy = d.Get("export_policy").(string)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.LogLevel = d.Get("log_level").(string)
	ret.Inst.Target = d.Get("target").(string)
	//omit uuid
	return ret
}
