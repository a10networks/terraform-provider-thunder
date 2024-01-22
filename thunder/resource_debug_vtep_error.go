package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugVtepError() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_vtep_error`: Debug Overlay errors\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugVtepErrorCreate,
		UpdateContext: resourceDebugVtepErrorUpdate,
		ReadContext:   resourceDebugVtepErrorRead,
		DeleteContext: resourceDebugVtepErrorDelete,

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
func resourceDebugVtepErrorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepErrorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepError(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVtepErrorRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugVtepErrorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepErrorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepError(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVtepErrorRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugVtepErrorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepErrorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepError(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugVtepErrorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepErrorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepError(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugVtepError(d *schema.ResourceData) edpt.DebugVtepError {
	var ret edpt.DebugVtepError
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
