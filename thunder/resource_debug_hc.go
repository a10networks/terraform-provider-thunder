package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugHc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_hc`: Debug Harmony Controller\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugHcCreate,
		UpdateContext: resourceDebugHcUpdate,
		ReadContext:   resourceDebugHcRead,
		DeleteContext: resourceDebugHcDelete,

		Schema: map[string]*schema.Schema{
			"anomaly": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dump per-request in anomaly cases only",
			},
			"app_svc_id": {
				Type: schema.TypeString, Optional: true, Description: "Application service id (virtual-server_port_protocol)",
			},
			"error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for harmony controller (error)",
			},
			"metrics": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for harmony controller (metrics)",
			},
			"object_uuid": {
				Type: schema.TypeString, Optional: true, Description: "UUID of the object to filter",
			},
			"per_connection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for harmony controller (per-connection)",
			},
			"per_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for harmony controller (per-request)",
			},
			"registration": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for harmony controller (registration)",
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
func resourceDebugHcCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHcCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHcRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugHcUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHcUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHcRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugHcDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHcDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugHcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHcRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugHc(d *schema.ResourceData) edpt.DebugHc {
	var ret edpt.DebugHc
	ret.Inst.Anomaly = d.Get("anomaly").(int)
	ret.Inst.AppSvcId = d.Get("app_svc_id").(string)
	ret.Inst.Error = d.Get("error").(int)
	ret.Inst.Metrics = d.Get("metrics").(int)
	ret.Inst.ObjectUuid = d.Get("object_uuid").(string)
	ret.Inst.PerConnection = d.Get("per_connection").(int)
	ret.Inst.PerRequest = d.Get("per_request").(int)
	ret.Inst.Registration = d.Get("registration").(int)
	ret.Inst.Uri = d.Get("uri").(string)
	//omit uuid
	return ret
}
