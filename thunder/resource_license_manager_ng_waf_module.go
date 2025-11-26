package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseManagerNgWafModule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_license_manager_ng_waf_module`: Enter WAF license-keys\n\n__PLACEHOLDER__",
		CreateContext: resourceLicenseManagerNgWafModuleCreate,
		UpdateContext: resourceLicenseManagerNgWafModuleUpdate,
		ReadContext:   resourceLicenseManagerNgWafModuleRead,
		DeleteContext: resourceLicenseManagerNgWafModuleDelete,

		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type: schema.TypeString, Optional: true, Description: "access-key",
			},
			"secret_access_key": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
		},
	}
}
func resourceLicenseManagerNgWafModuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerNgWafModuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerNgWafModule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerNgWafModuleRead(ctx, d, meta)
	}
	return diags
}

func resourceLicenseManagerNgWafModuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerNgWafModuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerNgWafModule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerNgWafModuleRead(ctx, d, meta)
	}
	return diags
}
func resourceLicenseManagerNgWafModuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerNgWafModuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerNgWafModule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLicenseManagerNgWafModuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerNgWafModuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManagerNgWafModule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLicenseManagerNgWafModule(d *schema.ResourceData) edpt.LicenseManagerNgWafModule {
	var ret edpt.LicenseManagerNgWafModule
	ret.Inst.AccessKeyId = d.Get("access_key_id").(string)
	ret.Inst.SecretAccessKey = d.Get("secret_access_key").(string)
	return ret
}
