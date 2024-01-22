package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsVmasterTakeOver() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_vmaster_take_over`: Forcefully take over mastership\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsVmasterTakeOverCreate,
		UpdateContext: resourceVcsVmasterTakeOverUpdate,
		ReadContext:   resourceVcsVmasterTakeOverRead,
		DeleteContext: resourceVcsVmasterTakeOverDelete,

		Schema: map[string]*schema.Schema{
			"vmaster_take_over": {
				Type: schema.TypeInt, Optional: true, Description: "vMaster take over priority",
			},
		},
	}
}
func resourceVcsVmasterTakeOverCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterTakeOverCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterTakeOver(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVmasterTakeOverRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsVmasterTakeOverUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterTakeOverUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterTakeOver(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVmasterTakeOverRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsVmasterTakeOverDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterTakeOverDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterTakeOver(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsVmasterTakeOverRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVmasterTakeOverRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVmasterTakeOver(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVcsVmasterTakeOver(d *schema.ResourceData) edpt.VcsVmasterTakeOver {
	var ret edpt.VcsVmasterTakeOver
	ret.Inst.VmasterTakeOver = d.Get("vmaster_take_over").(int)
	return ret
}
