package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugController() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_controller`: Debug Controller\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugControllerCreate,
		UpdateContext: resourceDebugControllerUpdate,
		ReadContext:   resourceDebugControllerRead,
		DeleteContext: resourceDebugControllerDelete,

		Schema: map[string]*schema.Schema{
			"anomaly": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dump per-request in anomaly cases only",
			},
			"app_svc_id": {
				Type: schema.TypeString, Optional: true, Description: "Application service id (virtual-server_port_protocol)",
			},
			"error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for controller (error)",
			},
			"logd_audit_export": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for controller (logd-audit-export)",
			},
			"logd_syslog_export": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for controller (logd-syslog-export)",
			},
			"metrics": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for controller (metrics)",
			},
			"object_uuid": {
				Type: schema.TypeString, Optional: true, Description: "UUID of the object to filter",
			},
			"per_connection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for controller (per-connection)",
			},
			"per_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for controller (per-request)",
			},
			"registration": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for controller (registration)",
			},
			"uri": {
				Type: schema.TypeString, Optional: true, Description: "URI of the object to filter",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugControllerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugControllerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugController(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugControllerRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugControllerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugControllerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugController(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugControllerRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugControllerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugControllerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugController(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugControllerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugControllerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugController(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugController(d *schema.ResourceData) edpt.DebugController {
	var ret edpt.DebugController
	ret.Inst.Anomaly = d.Get("anomaly").(int)
	ret.Inst.AppSvcId = d.Get("app_svc_id").(string)
	ret.Inst.Error = d.Get("error").(int)
	ret.Inst.LogdAuditExport = d.Get("logd_audit_export").(int)
	ret.Inst.LogdSyslogExport = d.Get("logd_syslog_export").(int)
	ret.Inst.Metrics = d.Get("metrics").(int)
	ret.Inst.ObjectUuid = d.Get("object_uuid").(string)
	ret.Inst.PerConnection = d.Get("per_connection").(int)
	ret.Inst.PerRequest = d.Get("per_request").(int)
	ret.Inst.Registration = d.Get("registration").(int)
	ret.Inst.Uri = d.Get("uri").(string)
	//omit uuid
	return ret
}
