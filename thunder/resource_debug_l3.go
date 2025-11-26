package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugL3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_l3`: Layer 3\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugL3Create,
		UpdateContext: resourceDebugL3Update,
		ReadContext:   resourceDebugL3Read,
		DeleteContext: resourceDebugL3Delete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug icmp packets",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugL3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL3Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugL3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL3Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugL3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugL3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugL3(d *schema.ResourceData) edpt.DebugL3 {
	var ret edpt.DebugL3
	ret.Inst.Icmp = d.Get("icmp").(int)
	//omit uuid
	return ret
}
