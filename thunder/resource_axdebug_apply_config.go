package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebugApplyConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug_apply_config`: Apply AXDebug config file\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugApplyConfigCreate,
		UpdateContext: resourceAxdebugApplyConfigUpdate,
		ReadContext:   resourceAxdebugApplyConfigRead,
		DeleteContext: resourceAxdebugApplyConfigDelete,

		Schema: map[string]*schema.Schema{
			"config_file": {
				Type: schema.TypeString, Optional: true, Description: "config file name",
			},
		},
	}
}
func resourceAxdebugApplyConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugApplyConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugApplyConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugApplyConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugApplyConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugApplyConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugApplyConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugApplyConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugApplyConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugApplyConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugApplyConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugApplyConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugApplyConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugApplyConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAxdebugApplyConfig(d *schema.ResourceData) edpt.AxdebugApplyConfig {
	var ret edpt.AxdebugApplyConfig
	ret.Inst.ConfigFile = d.Get("config_file").(string)
	return ret
}
