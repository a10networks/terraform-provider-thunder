package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugRmthttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_rmthttp`: Debug RMTHTTP API\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugRmthttpCreate,
		UpdateContext: resourceDebugRmthttpUpdate,
		ReadContext:   resourceDebugRmthttpRead,
		DeleteContext: resourceDebugRmthttpDelete,

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
func resourceDebugRmthttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRmthttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRmthttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRmthttpRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugRmthttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRmthttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRmthttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRmthttpRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugRmthttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRmthttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRmthttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugRmthttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRmthttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRmthttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugRmthttp(d *schema.ResourceData) edpt.DebugRmthttp {
	var ret edpt.DebugRmthttp
	ret.Inst.Dummy = d.Get("dummy").(int)
	//omit uuid
	return ret
}
