package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerTelemetry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_telemetry`: Controller telemetry config\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerTelemetryCreate,
		UpdateContext: resourceControllerTelemetryUpdate,
		ReadContext:   resourceControllerTelemetryRead,
		DeleteContext: resourceControllerTelemetryDelete,

		Schema: map[string]*schema.Schema{
			"log_rate": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Max number of session logs sent by the partition per second",
			},
			"probe": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the probe functionality; 'disable': Disable the probe functionality;",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 15, Description: "snapshot export interval in minute,default is 15.",
						},
						"log_level": {
							Type: schema.TypeString, Optional: true, Default: "ERROR", Description: "'ERROR': show errors only(default).; 'WARNING': show warnings; 'INFO': show info messages; 'DEBUG': show debug logs;",
						},
						"export_policy": {
							Type: schema.TypeString, Optional: true, Default: "snapshots-new", Description: "'snapshots-all': Export historical/missed snapshots.; 'snapshots-new': Export only new snapshots(default).;",
						},
						"target": {
							Type: schema.TypeString, Optional: true, Default: "remote", Description: "'remote': Export data to remote. This is the default value.; 'local': Export data local.;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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

func getObjectControllerTelemetryProbe144(d []interface{}) edpt.ControllerTelemetryProbe144 {

	count1 := len(d)
	var ret edpt.ControllerTelemetryProbe144
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.Interval = in["interval"].(int)
		ret.LogLevel = in["log_level"].(string)
		ret.ExportPolicy = in["export_policy"].(string)
		ret.Target = in["target"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointControllerTelemetry(d *schema.ResourceData) edpt.ControllerTelemetry {
	var ret edpt.ControllerTelemetry
	ret.Inst.LogRate = d.Get("log_rate").(int)
	ret.Inst.Probe = getObjectControllerTelemetryProbe144(d.Get("probe").([]interface{}))
	//omit uuid
	return ret
}
