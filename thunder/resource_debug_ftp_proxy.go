package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugFtpProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ftp_proxy`: Debug FTP Proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugFtpProxyCreate,
		UpdateContext: resourceDebugFtpProxyUpdate,
		ReadContext:   resourceDebugFtpProxyRead,
		DeleteContext: resourceDebugFtpProxyDelete,

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
func resourceDebugFtpProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFtpProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFtpProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugFtpProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugFtpProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFtpProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFtpProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugFtpProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugFtpProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFtpProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFtpProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugFtpProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFtpProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFtpProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugFtpProxy(d *schema.ResourceData) edpt.DebugFtpProxy {
	var ret edpt.DebugFtpProxy
	ret.Inst.Dummy = d.Get("dummy").(int)
	//omit uuid
	return ret
}
