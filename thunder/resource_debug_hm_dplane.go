package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugHmDplane() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_hm_dplane`: Debug health monitor in dplane\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugHmDplaneCreate,
		UpdateContext: resourceDebugHmDplaneUpdate,
		ReadContext:   resourceDebugHmDplaneRead,
		DeleteContext: resourceDebugHmDplaneDelete,

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
func resourceDebugHmDplaneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmDplaneCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHmDplane(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHmDplaneRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugHmDplaneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmDplaneUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHmDplane(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHmDplaneRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugHmDplaneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmDplaneDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHmDplane(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugHmDplaneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmDplaneRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHmDplane(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugHmDplane(d *schema.ResourceData) edpt.DebugHmDplane {
	var ret edpt.DebugHmDplane
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
