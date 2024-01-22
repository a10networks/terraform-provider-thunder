package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwApplyChanges() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_apply_changes`: Recompile rule-set immediately\n\n__PLACEHOLDER__",
		CreateContext: resourceFwApplyChangesCreate,
		UpdateContext: resourceFwApplyChangesUpdate,
		ReadContext:   resourceFwApplyChangesRead,
		DeleteContext: resourceFwApplyChangesDelete,

		Schema: map[string]*schema.Schema{
			"forced": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force recompile rule-set",
			},
		},
	}
}
func resourceFwApplyChangesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwApplyChangesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApplyChanges(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwApplyChangesRead(ctx, d, meta)
	}
	return diags
}

func resourceFwApplyChangesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwApplyChangesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApplyChanges(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwApplyChangesRead(ctx, d, meta)
	}
	return diags
}
func resourceFwApplyChangesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwApplyChangesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApplyChanges(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwApplyChangesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwApplyChangesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApplyChanges(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwApplyChanges(d *schema.ResourceData) edpt.FwApplyChanges {
	var ret edpt.FwApplyChanges
	ret.Inst.Forced = d.Get("forced").(int)
	return ret
}
