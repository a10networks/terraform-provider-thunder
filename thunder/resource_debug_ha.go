package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugHa() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ha`: Debug HA\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugHaCreate,
		UpdateContext: resourceDebugHaUpdate,
		ReadContext:   resourceDebugHaRead,
		DeleteContext: resourceDebugHaDelete,

		Schema: map[string]*schema.Schema{
			"debug_level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-7)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugHaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHa(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHaRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugHaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHa(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHaRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugHaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHa(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugHaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHa(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugHa(d *schema.ResourceData) edpt.DebugHa {
	var ret edpt.DebugHa
	ret.Inst.DebugLevel = d.Get("debug_level").(int)
	//omit uuid
	return ret
}
