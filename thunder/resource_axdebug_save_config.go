package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebugSaveConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug_save_config`: save AXDebug config file\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugSaveConfigCreate,
		UpdateContext: resourceAxdebugSaveConfigUpdate,
		ReadContext:   resourceAxdebugSaveConfigRead,
		DeleteContext: resourceAxdebugSaveConfigDelete,

		Schema: map[string]*schema.Schema{
			"config_file": {
				Type: schema.TypeString, Optional: true, Description: "config file name",
			},
			"default": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "save to default config file",
			},
		},
	}
}
func resourceAxdebugSaveConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugSaveConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugSaveConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugSaveConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugSaveConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugSaveConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugSaveConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugSaveConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugSaveConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugSaveConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugSaveConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugSaveConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugSaveConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugSaveConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAxdebugSaveConfig(d *schema.ResourceData) edpt.AxdebugSaveConfig {
	var ret edpt.AxdebugSaveConfig
	ret.Inst.ConfigFile = d.Get("config_file").(string)
	ret.Inst.Default = d.Get("default").(int)
	return ret
}
