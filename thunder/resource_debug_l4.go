package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugL4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_l4`: Layer 4\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugL4Create,
		UpdateContext: resourceDebugL4Update,
		ReadContext:   resourceDebugL4Read,
		DeleteContext: resourceDebugL4Delete,

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
func resourceDebugL4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL4Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugL4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL4Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugL4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugL4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugL4(d *schema.ResourceData) edpt.DebugL4 {
	var ret edpt.DebugL4
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
