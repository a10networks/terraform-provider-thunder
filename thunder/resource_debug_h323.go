package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugH323() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_h323`: Debug H323\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugH323Create,
		UpdateContext: resourceDebugH323Update,
		ReadContext:   resourceDebugH323Read,
		DeleteContext: resourceDebugH323Delete,

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
func resourceDebugH323Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugH323Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugH323(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugH323Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugH323Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugH323Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugH323(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugH323Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugH323Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugH323Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugH323(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugH323Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugH323Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugH323(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugH323(d *schema.ResourceData) edpt.DebugH323 {
	var ret edpt.DebugH323
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
