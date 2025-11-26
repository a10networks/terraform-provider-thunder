package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugL2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_l2`: Layer 2\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugL2Create,
		UpdateContext: resourceDebugL2Update,
		ReadContext:   resourceDebugL2Read,
		DeleteContext: resourceDebugL2Delete,

		Schema: map[string]*schema.Schema{
			"mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug mac learned packets",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugL2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL2Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugL2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL2Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugL2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugL2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugL2(d *schema.ResourceData) edpt.DebugL2 {
	var ret edpt.DebugL2
	ret.Inst.Mac = d.Get("mac").(int)
	//omit uuid
	return ret
}
