package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemReset() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_reset`: Reset system to the default configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResetCreate,
		UpdateContext: resourceSystemResetUpdate,
		ReadContext:   resourceSystemResetRead,
		DeleteContext: resourceSystemResetDelete,

		Schema: map[string]*schema.Schema{
			"reboot_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
		},
	}
}
func resourceSystemResetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemReset(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResetRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemReset(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResetRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemReset(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemReset(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemReset(d *schema.ResourceData) edpt.SystemReset {
	var ret edpt.SystemReset
	ret.Inst.RebootFlag = d.Get("reboot_flag").(int)
	return ret
}
