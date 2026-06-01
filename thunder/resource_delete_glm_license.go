package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteGlmLicense() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_glm_license`: Remove glm license\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteGlmLicenseCreate,
		UpdateContext: resourceDeleteGlmLicenseUpdate,
		ReadContext:   resourceDeleteGlmLicenseRead,
		DeleteContext: resourceDeleteGlmLicenseDelete,

		Schema: map[string]*schema.Schema{
			"a10_ti": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove A10 Threat Intel license",
			},
			"ipsec_vpn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove IPSEC VPN license",
			},
			"ngwaf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove NGWAF license",
			},
			"qosmos": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove QOSMOS license",
			},
			"secure_gaming": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove Secure gaming license",
			},
			"threatstop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove ThreatSTOP license",
			},
			"webroot": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove Webroot license",
			},
			"webroot_ti": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "only remove Webroot Threat Intel license",
			},
		},
	}
}
func resourceDeleteGlmLicenseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGlmLicenseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGlmLicense(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteGlmLicenseRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteGlmLicenseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGlmLicenseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGlmLicense(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteGlmLicenseRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteGlmLicenseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGlmLicenseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGlmLicense(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteGlmLicenseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteGlmLicenseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteGlmLicense(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteGlmLicense(d *schema.ResourceData) edpt.DeleteGlmLicense {
	var ret edpt.DeleteGlmLicense
	ret.Inst.A10Ti = d.Get("a10_ti").(int)
	ret.Inst.IpsecVpn = d.Get("ipsec_vpn").(int)
	ret.Inst.Ngwaf = d.Get("ngwaf").(int)
	ret.Inst.Qosmos = d.Get("qosmos").(int)
	ret.Inst.SecureGaming = d.Get("secure_gaming").(int)
	ret.Inst.Threatstop = d.Get("threatstop").(int)
	ret.Inst.Webroot = d.Get("webroot").(int)
	ret.Inst.WebrootTi = d.Get("webroot_ti").(int)
	return ret
}
