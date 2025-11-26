package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugCgnv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_cgnv6`: cgnv6 packet debugging\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugCgnv6Create,
		UpdateContext: resourceDebugCgnv6Update,
		ReadContext:   resourceDebugCgnv6Read,
		DeleteContext: resourceDebugCgnv6Delete,

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
func resourceDebugCgnv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgnv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCgnv6Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugCgnv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgnv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCgnv6Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugCgnv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgnv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugCgnv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgnv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugCgnv6(d *schema.ResourceData) edpt.DebugCgnv6 {
	var ret edpt.DebugCgnv6
	ret.Inst.Dummy = d.Get("dummy").(int)
	//omit uuid
	return ret
}
