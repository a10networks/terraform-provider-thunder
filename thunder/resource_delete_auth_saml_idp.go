package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteAuthSamlIdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_auth_saml_idp`: SAML metadata of identity provider\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteAuthSamlIdpCreate,
		UpdateContext: resourceDeleteAuthSamlIdpUpdate,
		ReadContext:   resourceDeleteAuthSamlIdpRead,
		DeleteContext: resourceDeleteAuthSamlIdpDelete,

		Schema: map[string]*schema.Schema{
			"saml_idp_name": {
				Type: schema.TypeString, Optional: true, Description: "Local IDP metadata name",
			},
		},
	}
}
func resourceDeleteAuthSamlIdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthSamlIdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthSamlIdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthSamlIdpRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteAuthSamlIdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthSamlIdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthSamlIdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthSamlIdpRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteAuthSamlIdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthSamlIdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthSamlIdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteAuthSamlIdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthSamlIdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthSamlIdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteAuthSamlIdp(d *schema.ResourceData) edpt.DeleteAuthSamlIdp {
	var ret edpt.DeleteAuthSamlIdp
	ret.Inst.SamlIdpName = d.Get("saml_idp_name").(string)
	return ret
}
