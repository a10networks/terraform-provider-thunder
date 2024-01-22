package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationOauthAuthorizationServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_oauth_authorization_server`: Authentication 2.0 authorization server\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationOauthAuthorizationServerCreate,
		UpdateContext: resourceAamAuthenticationOauthAuthorizationServerUpdate,
		ReadContext:   resourceAamAuthenticationOauthAuthorizationServerRead,
		DeleteContext: resourceAamAuthenticationOauthAuthorizationServerDelete,

		Schema: map[string]*schema.Schema{
			"authorization_endpoint": {
				Type: schema.TypeString, Optional: true, Description: "Specify URI for authorization",
			},
			"client_method": {
				Type: schema.TypeString, Optional: true, Description: "'ignored': Clients' browser will send data according to server spec (default); 'post': Clients' browser will send data by POST; 'get': Clients' browser will send data by GET;",
			},
			"issuer": {
				Type: schema.TypeString, Optional: true, Description: "Specify openid provider name for authorization",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify authorization server object name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'auth-req': auth-req; 'auth-succ': auth-succ; 'auth-fail': auth-fail; 'auth-error': auth-error; 'other-error': other-error;",
						},
					},
				},
			},
			"server_method": {
				Type: schema.TypeString, Optional: true, Description: "'post': AX will send data to server by POST (default); 'get': AX will send data to server by GET;",
			},
			"token_endpoint": {
				Type: schema.TypeString, Optional: true, Description: "Specify URI for token exchange",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"verification_cert": {
				Type: schema.TypeString, Optional: true, Description: "Specify certificate to verify ID token signature",
			},
			"verification_jwks": {
				Type: schema.TypeString, Optional: true, Description: "Specify jwks file to verify ID token signature",
			},
		},
	}
}
func resourceAamAuthenticationOauthAuthorizationServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthAuthorizationServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthAuthorizationServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationOauthAuthorizationServerRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationOauthAuthorizationServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthAuthorizationServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthAuthorizationServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationOauthAuthorizationServerRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationOauthAuthorizationServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthAuthorizationServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthAuthorizationServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationOauthAuthorizationServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthAuthorizationServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthAuthorizationServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationOauthAuthorizationServerSamplingEnable(d []interface{}) []edpt.AamAuthenticationOauthAuthorizationServerSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationOauthAuthorizationServerSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationOauthAuthorizationServerSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationOauthAuthorizationServer(d *schema.ResourceData) edpt.AamAuthenticationOauthAuthorizationServer {
	var ret edpt.AamAuthenticationOauthAuthorizationServer
	ret.Inst.AuthorizationEndpoint = d.Get("authorization_endpoint").(string)
	ret.Inst.ClientMethod = d.Get("client_method").(string)
	ret.Inst.Issuer = d.Get("issuer").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationOauthAuthorizationServerSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerMethod = d.Get("server_method").(string)
	ret.Inst.TokenEndpoint = d.Get("token_endpoint").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VerificationCert = d.Get("verification_cert").(string)
	ret.Inst.VerificationJwks = d.Get("verification_jwks").(string)
	return ret
}
