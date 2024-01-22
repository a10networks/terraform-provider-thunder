package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationJwt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_jwt`: JWT issuance configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationJwtCreate,
		UpdateContext: resourceAamAuthenticationJwtUpdate,
		ReadContext:   resourceAamAuthenticationJwtRead,
		DeleteContext: resourceAamAuthenticationJwtDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'redirect': redirect JWT to specific URI; 'relay': relay JWT to back-end;",
			},
			"issuer": {
				Type: schema.TypeString, Optional: true, Description: "Specify JWT issuer claim value",
			},
			"jwt_relay_uri": {
				Type: schema.TypeString, Optional: true, Description: "Specify JWT relay URI (for relay action)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify JWT issuer template name",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "The JWT signature secret",
			},
			"signature_secret": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the JWT signature secret",
			},
			"token_lifetime": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Specify JWT token lifetime (Specify lifetime (in seconds), default is 300.)",
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
func resourceAamAuthenticationJwtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationJwtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationJwt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationJwtRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationJwtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationJwtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationJwt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationJwtRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationJwtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationJwtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationJwt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationJwtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationJwtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationJwt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthenticationJwt(d *schema.ResourceData) edpt.AamAuthenticationJwt {
	var ret edpt.AamAuthenticationJwt
	ret.Inst.Action = d.Get("action").(string)
	//omit encrypted
	ret.Inst.Issuer = d.Get("issuer").(string)
	ret.Inst.JwtRelayUri = d.Get("jwt_relay_uri").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.SignatureSecret = d.Get("signature_secret").(int)
	ret.Inst.TokenLifetime = d.Get("token_lifetime").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
