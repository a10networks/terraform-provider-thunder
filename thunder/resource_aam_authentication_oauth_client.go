package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationOauthClient() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_oauth_client`: Authentication 2.0 Oauth client\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationOauthClientCreate,
		UpdateContext: resourceAamAuthenticationOauthClientUpdate,
		ReadContext:   resourceAamAuthenticationOauthClientRead,
		DeleteContext: resourceAamAuthenticationOauthClientDelete,

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type: schema.TypeString, Optional: true, Description: "Specify oauth client-id",
			},
			"client_secret": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"grant_type": {
				Type: schema.TypeString, Optional: true, Description: "'implicit': The authorization server will return access token directly.; 'authorization-code': The authorization server will respond with code which can be exchange for access token.; 'hybrid-code-id-token': The authorization server will respond with both code and id token.; 'hybrid-code-token': The authorization server will respond with both code and access token.; 'hybrid-all': The authorization server will respond with code, access token and id token;",
			},
			"infinity": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auth session never time out whatever value oauth servers' response",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify client object name",
			},
			"no_reply": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "AX will not check the nonce value in response",
			},
			"parameter_nonce_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable nonce parameter for authorization and token request",
			},
			"redirection_endpoint": {
				Type: schema.TypeString, Optional: true, Description: "Oauth client redirection endpoint service URL.",
			},
			"scope": {
				Type: schema.TypeString, Optional: true, Description: "Specify request scope parameters (e.g. profile email address phone)",
			},
			"session_init_ttl": {
				Type: schema.TypeInt, Optional: true, Description: "TTL for Thunder to wait for first response from authorization server",
			},
			"token_lifetime": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'openid-connect': openid-connect;",
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
func resourceAamAuthenticationOauthClientCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthClientCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthClient(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationOauthClientRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationOauthClientUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthClientUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthClient(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationOauthClientRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationOauthClientDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthClientDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthClient(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationOauthClientRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthClientRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthClient(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthenticationOauthClient(d *schema.ResourceData) edpt.AamAuthenticationOauthClient {
	var ret edpt.AamAuthenticationOauthClient
	ret.Inst.ClientId = d.Get("client_id").(string)
	ret.Inst.ClientSecret = d.Get("client_secret").(string)
	//omit encrypted
	ret.Inst.GrantType = d.Get("grant_type").(string)
	ret.Inst.Infinity = d.Get("infinity").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NoReply = d.Get("no_reply").(int)
	ret.Inst.ParameterNonceEnable = d.Get("parameter_nonce_enable").(int)
	ret.Inst.RedirectionEndpoint = d.Get("redirection_endpoint").(string)
	ret.Inst.Scope = d.Get("scope").(string)
	ret.Inst.SessionInitTtl = d.Get("session_init_ttl").(int)
	ret.Inst.TokenLifetime = d.Get("token_lifetime").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
