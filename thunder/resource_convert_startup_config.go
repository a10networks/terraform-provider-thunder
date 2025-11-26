package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConvertStartupConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_convert_startup_config`: Convert configuration profile from older version to new version\n\n__PLACEHOLDER__",
		CreateContext: resourceConvertStartupConfigCreate,
		UpdateContext: resourceConvertStartupConfigUpdate,
		ReadContext:   resourceConvertStartupConfigRead,
		DeleteContext: resourceConvertStartupConfigDelete,

		Schema: map[string]*schema.Schema{
			"dest_profile": {
				Type: schema.TypeString, Optional: true, Description: "Local Configuration Profile Name",
			},
			"profile": {
				Type: schema.TypeString, Optional: true, Description: "Choose a config profile to convert",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'3_0': Convert from 3.0 to 3.2; '3_2_5': Convert from 3.2.5 to 5.0.0-TPS; '5_0': Convert from 5.0-TPS to 5.3-TPS;",
			},
		},
	}
}
func resourceConvertStartupConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConvertStartupConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConvertStartupConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConvertStartupConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceConvertStartupConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConvertStartupConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConvertStartupConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConvertStartupConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceConvertStartupConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConvertStartupConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConvertStartupConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceConvertStartupConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConvertStartupConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConvertStartupConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointConvertStartupConfig(d *schema.ResourceData) edpt.ConvertStartupConfig {
	var ret edpt.ConvertStartupConfig
	ret.Inst.DestProfile = d.Get("dest_profile").(string)
	ret.Inst.Profile = d.Get("profile").(string)
	ret.Inst.Type = d.Get("type").(string)
	return ret
}
