package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugHttpProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_http_proxy`: Debug HTTP proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugHttpProxyCreate,
		UpdateContext: resourceDebugHttpProxyUpdate,
		ReadContext:   resourceDebugHttpProxyRead,
		DeleteContext: resourceDebugHttpProxyDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-4)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugHttpProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttpProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttpProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHttpProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugHttpProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttpProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttpProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHttpProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugHttpProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttpProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttpProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugHttpProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttpProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttpProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugHttpProxy(d *schema.ResourceData) edpt.DebugHttpProxy {
	var ret edpt.DebugHttpProxy
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	return ret
}
