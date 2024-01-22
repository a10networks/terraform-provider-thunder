package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBootBlockFix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_boot_block_fix`: Fix booting problem\n\n__PLACEHOLDER__",
		CreateContext: resourceBootBlockFixCreate,
		UpdateContext: resourceBootBlockFixUpdate,
		ReadContext:   resourceBootBlockFixRead,
		DeleteContext: resourceBootBlockFixDelete,

		Schema: map[string]*schema.Schema{
			"cf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Compact flash",
			},
			"hd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Hard disk",
			},
		},
	}
}
func resourceBootBlockFixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootBlockFixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootBlockFix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBootBlockFixRead(ctx, d, meta)
	}
	return diags
}

func resourceBootBlockFixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootBlockFixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootBlockFix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBootBlockFixRead(ctx, d, meta)
	}
	return diags
}
func resourceBootBlockFixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootBlockFixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootBlockFix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBootBlockFixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootBlockFixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootBlockFix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointBootBlockFix(d *schema.ResourceData) edpt.BootBlockFix {
	var ret edpt.BootBlockFix
	ret.Inst.Cf = d.Get("cf").(int)
	ret.Inst.Hd = d.Get("hd").(int)
	return ret
}
