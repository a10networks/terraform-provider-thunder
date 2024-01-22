package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAL3InlineModeFlag() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_l3_inline_mode_flag`: Enable OSPF under Layer 3 Inline Hot Standby Mode\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAL3InlineModeFlagCreate,
		UpdateContext: resourceVrrpAL3InlineModeFlagUpdate,
		ReadContext:   resourceVrrpAL3InlineModeFlagRead,
		DeleteContext: resourceVrrpAL3InlineModeFlagDelete,

		Schema: map[string]*schema.Schema{
			"l3_inline_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Layer 3 Inline Hot Standby Mode",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVrrpAL3InlineModeFlagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL3InlineModeFlagCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL3InlineModeFlag(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAL3InlineModeFlagRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAL3InlineModeFlagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL3InlineModeFlagUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL3InlineModeFlag(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAL3InlineModeFlagRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAL3InlineModeFlagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL3InlineModeFlagDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL3InlineModeFlag(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAL3InlineModeFlagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL3InlineModeFlagRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL3InlineModeFlag(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAL3InlineModeFlag(d *schema.ResourceData) edpt.VrrpAL3InlineModeFlag {
	var ret edpt.VrrpAL3InlineModeFlag
	ret.Inst.L3InlineMode = d.Get("l3_inline_mode").(int)
	//omit uuid
	return ret
}
