package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_template`: Authentication template\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationTemplateCreate,
		UpdateContext: resourceAamAuthenticationTemplateUpdate,
		ReadContext:   resourceAamAuthenticationTemplateRead,
		DeleteContext: resourceAamAuthenticationTemplateDelete,

		Schema: map[string]*schema.Schema{
			"account": {
				Type: schema.TypeString, Optional: true, Description: "Specify AD domain account",
			},
			"accounting_server": {
				Type: schema.TypeString, Optional: true, Description: "Specify a RADIUS accounting server",
			},
			"accounting_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Specify an authentication service group for RADIUS accounting",
			},
			"auth_sess_mode": {
				Type: schema.TypeString, Optional: true, Description: "'cookie-based': Track auth-session by cookie (default); 'ip-based': Track auth-session by client IP;",
			},
			"captcha": {
				Type: schema.TypeString, Optional: true, Description: "Specify captcha profile (Specify captcha proflie name)",
			},
			"chain": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chain_server": {
							Type: schema.TypeString, Optional: true, Description: "Specify authentication server (Specify authentication server template name)",
						},
						"chain_server_priority": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set server priority, higher the number higher the priority. Default is 3. (Chain server priority, higher the number higher the priority. Default is 3.)",
						},
						"chain_sg": {
							Type: schema.TypeString, Optional: true, Description: "Bind an authentication service group to this template (Specify authentication service group name)",
						},
						"chain_sg_priority": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set service-group priority, higher the number higher the priority. Default is 3. (Chain service-group priority, higher the number higher the priority. Default is 3.)",
						},
					},
				},
			},
			"cookie_domain": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cookie_dmn": {
							Type: schema.TypeString, Optional: true, Description: "Specify domain scope for the authentication (ex: .a10networks.com)",
						},
					},
				},
			},
			"cookie_domain_group": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cookie_dmngrp": {
							Type: schema.TypeInt, Optional: true, Description: "Specify group id to join in the cookie-domain",
						},
					},
				},
			},
			"cookie_httponly_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable httponly attribute for AAM cookies",
			},
			"cookie_max_age": {
				Type: schema.TypeInt, Optional: true, Default: 604800, Description: "Configure Max-Age for authentication session cookie (Configure Max-Age in seconds, 0 for no Max-Age/Expires attributes. Default is 604800 (1 week).)",
			},
			"cookie_samesite": {
				Type: schema.TypeString, Optional: true, Description: "'strict': Specify SameSite attribute as Strict for AAM cookie; 'lax': Specify SameSite attribute as Lax for AAM cookie; 'none': Specify SameSite attribute as None for AAM cookie;",
			},
			"cookie_secure_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable secure attribute for AAM cookies",
			},
			"forward_logout_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable forward logout request to backend application server. The config-field logout-url must be configured first",
			},
			"jwt": {
				Type: schema.TypeString, Optional: true, Description: "Specify authentication jwt template",
			},
			"local_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local logging",
			},
			"log": {
				Type: schema.TypeString, Optional: true, Default: "use-partition-level-config", Description: "'use-partition-level-config': Use configuration of authentication-log enable command; 'enable': Enable authentication logs for this template; 'disable': Disable authentication logs for this template;",
			},
			"logon": {
				Type: schema.TypeString, Optional: true, Description: "Specify authentication logon (Specify authentication logon template name)",
			},
			"logout_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Specify idle logout time (Specify idle timeout in seconds, default is 300)",
			},
			"logout_url": {
				Type: schema.TypeString, Optional: true, Description: "Specify logout url (Specify logout url string)",
			},
			"max_session_time": {
				Type: schema.TypeInt, Optional: true, Description: "Specify default SAML token lifetime (Specify lifetime (in seconds) of SAML token when it not provided by token attributes, default is 28800. (0 for indefinite))",
			},
			"modify_content_security_policy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Put redirect-uri or service-principal-name into CSP header to avoid CPS break authentication process",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Authentication template name",
			},
			"oauth_authorization_server": {
				Type: schema.TypeString, Optional: true, Description: "Specify OAUTH authorization server",
			},
			"oauth_client": {
				Type: schema.TypeString, Optional: true, Description: "Specify OAUTH client",
			},
			"redirect_hostname": {
				Type: schema.TypeString, Optional: true, Description: "Hostname(Length 1-31) for transparent-proxy authentication",
			},
			"relay": {
				Type: schema.TypeString, Optional: true, Description: "Specify authentication relay (Specify authentication relay template name)",
			},
			"saml_idp": {
				Type: schema.TypeString, Optional: true, Description: "Specify SAML identity provider",
			},
			"saml_sp": {
				Type: schema.TypeString, Optional: true, Description: "Specify SAML service provider",
			},
			"server": {
				Type: schema.TypeString, Optional: true, Description: "Specify authentication server (Specify authentication server template name)",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind an authentication service group to this template (Specify authentication service group name)",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'saml': SAML authentication template; 'standard': Standard authentication template; 'oauth': Oauth 2.0 authentication template;",
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
func resourceAamAuthenticationTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationTemplateChain(d []interface{}) []edpt.AamAuthenticationTemplateChain {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationTemplateChain, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationTemplateChain
		oi.ChainServer = in["chain_server"].(string)
		oi.ChainServerPriority = in["chain_server_priority"].(int)
		oi.ChainSg = in["chain_sg"].(string)
		oi.ChainSgPriority = in["chain_sg_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationTemplateCookieDomain(d []interface{}) []edpt.AamAuthenticationTemplateCookieDomain {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationTemplateCookieDomain, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationTemplateCookieDomain
		oi.CookieDmn = in["cookie_dmn"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationTemplateCookieDomainGroup(d []interface{}) []edpt.AamAuthenticationTemplateCookieDomainGroup {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationTemplateCookieDomainGroup, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationTemplateCookieDomainGroup
		oi.CookieDmngrp = in["cookie_dmngrp"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationTemplate(d *schema.ResourceData) edpt.AamAuthenticationTemplate {
	var ret edpt.AamAuthenticationTemplate
	ret.Inst.Account = d.Get("account").(string)
	ret.Inst.AccountingServer = d.Get("accounting_server").(string)
	ret.Inst.AccountingServiceGroup = d.Get("accounting_service_group").(string)
	ret.Inst.AuthSessMode = d.Get("auth_sess_mode").(string)
	ret.Inst.Captcha = d.Get("captcha").(string)
	ret.Inst.Chain = getSliceAamAuthenticationTemplateChain(d.Get("chain").([]interface{}))
	ret.Inst.CookieDomain = getSliceAamAuthenticationTemplateCookieDomain(d.Get("cookie_domain").([]interface{}))
	ret.Inst.CookieDomainGroup = getSliceAamAuthenticationTemplateCookieDomainGroup(d.Get("cookie_domain_group").([]interface{}))
	ret.Inst.CookieHttponlyEnable = d.Get("cookie_httponly_enable").(int)
	ret.Inst.CookieMaxAge = d.Get("cookie_max_age").(int)
	ret.Inst.CookieSamesite = d.Get("cookie_samesite").(string)
	ret.Inst.CookieSecureEnable = d.Get("cookie_secure_enable").(int)
	ret.Inst.ForwardLogoutDisable = d.Get("forward_logout_disable").(int)
	ret.Inst.Jwt = d.Get("jwt").(string)
	ret.Inst.LocalLogging = d.Get("local_logging").(int)
	ret.Inst.Log = d.Get("log").(string)
	ret.Inst.Logon = d.Get("logon").(string)
	ret.Inst.LogoutIdleTimeout = d.Get("logout_idle_timeout").(int)
	ret.Inst.LogoutUrl = d.Get("logout_url").(string)
	ret.Inst.MaxSessionTime = d.Get("max_session_time").(int)
	ret.Inst.ModifyContentSecurityPolicy = d.Get("modify_content_security_policy").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OauthAuthorizationServer = d.Get("oauth_authorization_server").(string)
	ret.Inst.OauthClient = d.Get("oauth_client").(string)
	ret.Inst.RedirectHostname = d.Get("redirect_hostname").(string)
	ret.Inst.Relay = d.Get("relay").(string)
	ret.Inst.SamlIdp = d.Get("saml_idp").(string)
	ret.Inst.SamlSp = d.Get("saml_sp").(string)
	ret.Inst.Server = d.Get("server").(string)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
