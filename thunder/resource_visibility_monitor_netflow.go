package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorNetflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_netflow`: Configure Netflow parameters for flow based monitoring\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorNetflowCreate,
		UpdateContext: resourceVisibilityMonitorNetflowUpdate,
		ReadContext:   resourceVisibilityMonitorNetflowRead,
		DeleteContext: resourceVisibilityMonitorNetflowDelete,

		Schema: map[string]*schema.Schema{
			"listening_port": {
				Type: schema.TypeInt, Optional: true, Default: 9996, Description: "Netflow port to receive packets (Netflow port number(default 9996))",
			},
			"template_active_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Configure active timeout of the netflow templates received in mins (Template active timeout(mins)(default 30mins))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityMonitorNetflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorNetflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorNetflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorNetflowRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorNetflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorNetflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorNetflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorNetflowRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorNetflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorNetflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorNetflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorNetflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorNetflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorNetflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorNetflow(d *schema.ResourceData) edpt.VisibilityMonitorNetflow {
	var ret edpt.VisibilityMonitorNetflow
	ret.Inst.ListeningPort = d.Get("listening_port").(int)
	ret.Inst.TemplateActiveTimeout = d.Get("template_active_timeout").(int)
	//omit uuid
	return ret
}
