package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugDiameter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_diameter`: Debug diameter proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugDiameterCreate,
		UpdateContext: resourceDebugDiameterUpdate,
		ReadContext:   resourceDebugDiameterRead,
		DeleteContext: resourceDebugDiameterDelete,

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
func resourceDebugDiameterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDiameterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDiameter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDiameterRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugDiameterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDiameterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDiameter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDiameterRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugDiameterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDiameterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDiameter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugDiameterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDiameterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDiameter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugDiameter(d *schema.ResourceData) edpt.DebugDiameter {
	var ret edpt.DebugDiameter
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
