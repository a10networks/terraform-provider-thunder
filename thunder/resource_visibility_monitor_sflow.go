package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorSflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_sflow`: Configure sFlow parameters for flow based monitoring\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorSflowCreate,
		UpdateContext: resourceVisibilityMonitorSflowUpdate,
		ReadContext:   resourceVisibilityMonitorSflowRead,
		DeleteContext: resourceVisibilityMonitorSflowDelete,

		Schema: map[string]*schema.Schema{
			"listening_port": {
				Type: schema.TypeInt, Optional: true, Default: 6343, Description: "sFlow port to receive packets (sFlow port number(default 6343))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityMonitorSflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSflowRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorSflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSflowRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorSflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorSflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorSflow(d *schema.ResourceData) edpt.VisibilityMonitorSflow {
	var ret edpt.VisibilityMonitorSflow
	ret.Inst.ListeningPort = d.Get("listening_port").(int)
	//omit uuid
	return ret
}
