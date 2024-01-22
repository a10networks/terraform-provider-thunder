package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugHttp2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_http2`: Debug HTTP/2 processing\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugHttp2Create,
		UpdateContext: resourceDebugHttp2Update,
		ReadContext:   resourceDebugHttp2Read,
		DeleteContext: resourceDebugHttp2Delete,

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
func resourceDebugHttp2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttp2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttp2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHttp2Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugHttp2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttp2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttp2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHttp2Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugHttp2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttp2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttp2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugHttp2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHttp2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHttp2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugHttp2(d *schema.ResourceData) edpt.DebugHttp2 {
	var ret edpt.DebugHttp2
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	return ret
}
