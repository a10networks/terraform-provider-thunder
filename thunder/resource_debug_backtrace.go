package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugBacktrace() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_backtrace`: add backtrace into rima_log\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugBacktraceCreate,
		UpdateContext: resourceDebugBacktraceUpdate,
		ReadContext:   resourceDebugBacktraceRead,
		DeleteContext: resourceDebugBacktraceDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugBacktraceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBacktraceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBacktrace(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugBacktraceRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugBacktraceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBacktraceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBacktrace(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugBacktraceRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugBacktraceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBacktraceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBacktrace(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugBacktraceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBacktraceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBacktrace(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugBacktrace(d *schema.ResourceData) edpt.DebugBacktrace {
	var ret edpt.DebugBacktrace
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
