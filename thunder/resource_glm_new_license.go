package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlmNewLicense() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_glm_new_license`: Create license\n\n__PLACEHOLDER__",
		CreateContext: resourceGlmNewLicenseCreate,
		UpdateContext: resourceGlmNewLicenseUpdate,
		ReadContext:   resourceGlmNewLicenseRead,
		DeleteContext: resourceGlmNewLicenseDelete,

		Schema: map[string]*schema.Schema{
			"account_name": {
				Type: schema.TypeString, Optional: true, Description: "Account Name",
			},
			"country": {
				Type: schema.TypeString, Optional: true, Description: "Country",
			},
			"existing_org": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use existing account with organization ID",
			},
			"existing_user": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use an existing account with email and password",
			},
			"first_name": {
				Type: schema.TypeString, Optional: true, Description: "First Name",
			},
			"glm_email": {
				Type: schema.TypeString, Optional: true, Description: "GLM email",
			},
			"glm_password": {
				Type: schema.TypeString, Optional: true, Description: "GLM password",
			},
			"last_name": {
				Type: schema.TypeString, Optional: true, Description: "Last Name",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "License name (Configure license name)",
			},
			"new_email": {
				Type: schema.TypeString, Optional: true, Description: "GLM email",
			},
			"new_password": {
				Type: schema.TypeString, Optional: true, Description: "GLM password",
			},
			"new_user": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create a new account",
			},
			"org_id": {
				Type: schema.TypeInt, Optional: true, Description: "GLM organization id",
			},
			"phone": {
				Type: schema.TypeString, Optional: true, Description: "Phone",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'webroot': webroot; 'webroot_trial': webroot_trial; 'webroot_ti': webroot_ti; 'webroot_ti_trial': webroot_ti_trial; 'qosmos': qosmos; 'qosmos_trial': qosmos_trial; 'ipsec_vpn': ipsec_vpn;",
			},
		},
	}
}
func resourceGlmNewLicenseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmNewLicenseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmNewLicense(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmNewLicenseRead(ctx, d, meta)
	}
	return diags
}

func resourceGlmNewLicenseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmNewLicenseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmNewLicense(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmNewLicenseRead(ctx, d, meta)
	}
	return diags
}
func resourceGlmNewLicenseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmNewLicenseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmNewLicense(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGlmNewLicenseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmNewLicenseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmNewLicense(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGlmNewLicense(d *schema.ResourceData) edpt.GlmNewLicense {
	var ret edpt.GlmNewLicense
	ret.Inst.AccountName = d.Get("account_name").(string)
	ret.Inst.Country = d.Get("country").(string)
	ret.Inst.ExistingOrg = d.Get("existing_org").(int)
	ret.Inst.ExistingUser = d.Get("existing_user").(int)
	ret.Inst.FirstName = d.Get("first_name").(string)
	ret.Inst.GlmEmail = d.Get("glm_email").(string)
	ret.Inst.GlmPassword = d.Get("glm_password").(string)
	ret.Inst.LastName = d.Get("last_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NewEmail = d.Get("new_email").(string)
	ret.Inst.NewPassword = d.Get("new_password").(string)
	ret.Inst.NewUser = d.Get("new_user").(int)
	ret.Inst.OrgId = d.Get("org_id").(int)
	ret.Inst.Phone = d.Get("phone").(string)
	ret.Inst.Type = d.Get("type").(string)
	return ret
}
