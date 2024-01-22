package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugL7session() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_l7session`: Debug L7 session processing\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugL7sessionCreate,
		UpdateContext: resourceDebugL7sessionUpdate,
		ReadContext:   resourceDebugL7sessionRead,
		DeleteContext: resourceDebugL7sessionDelete,

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
func resourceDebugL7sessionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7sessionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7session(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL7sessionRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugL7sessionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7sessionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7session(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugL7sessionRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugL7sessionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7sessionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7session(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugL7sessionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugL7sessionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugL7session(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugL7session(d *schema.ResourceData) edpt.DebugL7session {
	var ret edpt.DebugL7session
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	return ret
}
