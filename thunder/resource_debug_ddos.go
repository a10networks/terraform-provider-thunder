package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugDdos() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ddos`: Debug DDOS\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugDdosCreate,
		UpdateContext: resourceDebugDdosUpdate,
		ReadContext:   resourceDebugDdosRead,
		DeleteContext: resourceDebugDdosDelete,

		Schema: map[string]*schema.Schema{
			"control_var": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug DDOS Control Var",
			},
			"dns_cache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug DDOS DNS Cache",
			},
			"event": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug DDOS event",
			},
			"event_filter": {
				Type: schema.TypeString, Optional: true, Description: "Set debug DDOS event filter",
			},
			"flow_based_detection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug flow samples based DDOS detection",
			},
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-4)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zbar": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug DDOS zbar event",
			},
		},
	}
}
func resourceDebugDdosCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDdosCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDdos(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDdosRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugDdosUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDdosUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDdos(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDdosRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugDdosDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDdosDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDdos(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugDdosRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDdosRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDdos(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugDdos(d *schema.ResourceData) edpt.DebugDdos {
	var ret edpt.DebugDdos
	ret.Inst.ControlVar = d.Get("control_var").(int)
	ret.Inst.DnsCache = d.Get("dns_cache").(int)
	ret.Inst.Event = d.Get("event").(int)
	ret.Inst.EventFilter = d.Get("event_filter").(string)
	ret.Inst.FlowBasedDetection = d.Get("flow_based_detection").(int)
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	ret.Inst.Zbar = d.Get("zbar").(int)
	return ret
}
