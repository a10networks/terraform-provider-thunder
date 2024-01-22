package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugVisibility() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_visibility`: Debug visibility\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugVisibilityCreate,
		UpdateContext: resourceDebugVisibilityUpdate,
		ReadContext:   resourceDebugVisibilityRead,
		DeleteContext: resourceDebugVisibilityDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"xflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "debug xflow monitoring",
			},
		},
	}
}
func resourceDebugVisibilityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVisibilityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVisibility(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVisibilityRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugVisibilityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVisibilityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVisibility(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVisibilityRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugVisibilityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVisibilityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVisibility(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugVisibilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVisibilityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVisibility(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugVisibility(d *schema.ResourceData) edpt.DebugVisibility {
	var ret edpt.DebugVisibility
	//omit uuid
	ret.Inst.Xflow = d.Get("xflow").(int)
	return ret
}
