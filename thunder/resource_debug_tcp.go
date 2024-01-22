package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_tcp`: Debug TCP\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugTcpCreate,
		UpdateContext: resourceDebugTcpUpdate,
		ReadContext:   resourceDebugTcpRead,
		DeleteContext: resourceDebugTcpDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-4)",
			},
			"stack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug TCP stack",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugTcp(d *schema.ResourceData) edpt.DebugTcp {
	var ret edpt.DebugTcp
	ret.Inst.Level = d.Get("level").(int)
	ret.Inst.Stack = d.Get("stack").(int)
	//omit uuid
	return ret
}
