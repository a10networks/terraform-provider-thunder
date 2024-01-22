package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSharedPollMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_shared_poll_mode`: Configure shared poll mode\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSharedPollModeCreate,
		UpdateContext: resourceSystemSharedPollModeUpdate,
		ReadContext:   resourceSystemSharedPollModeRead,
		DeleteContext: resourceSystemSharedPollModeDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable shared poll mode",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable shared poll mode",
			},
		},
	}
}
func resourceSystemSharedPollModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSharedPollModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSharedPollMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSharedPollModeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSharedPollModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSharedPollModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSharedPollMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSharedPollModeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSharedPollModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSharedPollModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSharedPollMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSharedPollModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSharedPollModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSharedPollMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSharedPollMode(d *schema.ResourceData) edpt.SystemSharedPollMode {
	var ret edpt.SystemSharedPollMode
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	return ret
}
