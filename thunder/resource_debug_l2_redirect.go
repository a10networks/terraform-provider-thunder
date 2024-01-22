package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugL2Redirect() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_l2_redirect`: Debug L2 SSLi Redirect\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugL2RedirectCreate,
		UpdateContext: resourceDebugL2RedirectUpdate,
		ReadContext:   resourceDebugL2RedirectRead,
		DeleteContext: resourceDebugL2RedirectDelete,

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
func resourceDebugL2RedirectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2RedirectCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2Redirect(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL2RedirectRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugL2RedirectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2RedirectUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2Redirect(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL2RedirectRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugL2RedirectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2RedirectDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2Redirect(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugL2RedirectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2RedirectRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2Redirect(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugL2Redirect(d *schema.ResourceData) edpt.DebugL2Redirect {
	var ret edpt.DebugL2Redirect
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	return ret
}
