package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_saml_identity_provider`: Authentication identity provider\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationSamlIdentityProviderCreate,
		UpdateContext: resourceAamAuthenticationSamlIdentityProviderUpdate,
		ReadContext:   resourceAamAuthenticationSamlIdentityProviderRead,
		DeleteContext: resourceAamAuthenticationSamlIdentityProviderDelete,

		Schema: map[string]*schema.Schema{
			"metadata": {
				Type: schema.TypeString, Optional: true, Description: "URL of SAML identity provider's metadata file",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SAML authentication identity provider name",
			},
			"reload_interval": {
				Type: schema.TypeInt, Optional: true, Default: 28800, Description: "Specify URI metadata reload period (Specify URI metadata reload period in seconds, default is 28800)",
			},
			"reload_metadata": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reload IdP's metadata immediately",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationSamlIdentityProviderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlIdentityProviderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlIdentityProvider(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlIdentityProviderRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationSamlIdentityProviderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlIdentityProviderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlIdentityProvider(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlIdentityProviderRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationSamlIdentityProviderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlIdentityProviderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlIdentityProvider(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationSamlIdentityProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlIdentityProviderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlIdentityProvider(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthenticationSamlIdentityProvider(d *schema.ResourceData) edpt.AamAuthenticationSamlIdentityProvider {
	var ret edpt.AamAuthenticationSamlIdentityProvider
	ret.Inst.Metadata = d.Get("metadata").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ReloadInterval = d.Get("reload_interval").(int)
	ret.Inst.ReloadMetadata = d.Get("reload_metadata").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
