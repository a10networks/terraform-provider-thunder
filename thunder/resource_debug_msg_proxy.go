package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugMsgProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_msg_proxy`: Debug msg proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugMsgProxyCreate,
		UpdateContext: resourceDebugMsgProxyUpdate,
		ReadContext:   resourceDebugMsgProxyRead,
		DeleteContext: resourceDebugMsgProxyDelete,

		Schema: map[string]*schema.Schema{
			"dummy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugMsgProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMsgProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMsgProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugMsgProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugMsgProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMsgProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMsgProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugMsgProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugMsgProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMsgProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMsgProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugMsgProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMsgProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMsgProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugMsgProxy(d *schema.ResourceData) edpt.DebugMsgProxy {
	var ret edpt.DebugMsgProxy
	ret.Inst.Dummy = d.Get("dummy").(int)
	//omit uuid
	return ret
}
