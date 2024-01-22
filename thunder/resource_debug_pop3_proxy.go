package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugPop3Proxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_pop3_proxy`: Debug POP3 proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugPop3ProxyCreate,
		UpdateContext: resourceDebugPop3ProxyUpdate,
		ReadContext:   resourceDebugPop3ProxyRead,
		DeleteContext: resourceDebugPop3ProxyDelete,

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
func resourceDebugPop3ProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPop3ProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPop3Proxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPop3ProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugPop3ProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPop3ProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPop3Proxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPop3ProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugPop3ProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPop3ProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPop3Proxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugPop3ProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPop3ProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPop3Proxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugPop3Proxy(d *schema.ResourceData) edpt.DebugPop3Proxy {
	var ret edpt.DebugPop3Proxy
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
