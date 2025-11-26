package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReload() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_reload`: Reload the system without rebooting\n\n__PLACEHOLDER__",
		CreateContext: resourceReloadCreate,
		UpdateContext: resourceReloadUpdate,
		ReadContext:   resourceReloadRead,
		DeleteContext: resourceReloadDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reload all devices when VCS is enabled, or only this device itself if VCS is not enabled",
			},
			"device": {
				Type: schema.TypeInt, Optional: true, Description: "Reload a specific device when VCS is enabled (device id)",
			},
		},
	}
}
func resourceReloadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReloadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReload(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReloadRead(ctx, d, meta)
	}
	return diags
}

func resourceReloadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReloadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReload(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReloadRead(ctx, d, meta)
	}
	return diags
}
func resourceReloadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReloadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReload(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceReloadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReloadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReload(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointReload(d *schema.ResourceData) edpt.Reload {
	var ret edpt.Reload
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Device = d.Get("device").(int)
	return ret
}
