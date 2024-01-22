package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugL7Http3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_l7_http3`: HTTP/3 processing\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugL7Http3Create,
		UpdateContext: resourceDebugL7Http3Update,
		ReadContext:   resourceDebugL7Http3Read,
		DeleteContext: resourceDebugL7Http3Delete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-5)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugL7Http3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7Http3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7Http3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL7Http3Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugL7Http3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7Http3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7Http3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL7Http3Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugL7Http3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7Http3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7Http3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugL7Http3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7Http3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7Http3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugL7Http3(d *schema.ResourceData) edpt.DebugL7Http3 {
	var ret edpt.DebugL7Http3
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	return ret
}
