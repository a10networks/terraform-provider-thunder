package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerLdapInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server_ldap_instance`: LDAP Authentication Server\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerLdapInstanceCreate,
		UpdateContext: resourceAamAuthenticationServerLdapInstanceUpdate,
		ReadContext:   resourceAamAuthenticationServerLdapInstanceRead,
		DeleteContext: resourceAamAuthenticationServerLdapInstanceDelete,

		Schema: map[string]*schema.Schema{
			"admin_dn": {
				Type: schema.TypeString, Optional: true, Description: "The LDAP server's admin DN",
			},
			"admin_secret": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the LDAP server's admin secret password",
			},
			"auth_type": {
				Type: schema.TypeString, Optional: true, Description: "'ad': Active Directory. Default; 'open-ldap': OpenLDAP;",
			},
			"base": {
				Type: schema.TypeString, Optional: true, Description: "Specify the LDAP server's search base",
			},
			"bind_with_dn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce using DN for LDAP binding(All user input name will be used to create DN)",
			},
			"ca_cert": {
				Type: schema.TypeString, Optional: true, Description: "Specify the LDAPS CA cert filename (Trusted LDAPS CA cert filename)",
			},
			"default_domain": {
				Type: schema.TypeString, Optional: true, Description: "Specify default domain for LDAP",
			},
			"derive_bind_dn": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username_attr": {
							Type: schema.TypeString, Optional: true, Description: "Specify attribute name of username",
						},
					},
				},
			},
			"dn_attribute": {
				Type: schema.TypeString, Optional: true, Default: "cn", Description: "Specify Distinguished Name attribute, default is CN",
			},
			"health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check server's health status",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
			},
			"health_check_string": {
				Type: schema.TypeString, Optional: true, Description: "Health monitor name",
			},
			"host": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostip": {
							Type: schema.TypeString, Optional: true, Description: "Server's hostname(Length 1-31) or IP address",
						},
						"hostipv6": {
							Type: schema.TypeString, Optional: true, Description: "Server's IPV6 address",
						},
					},
				},
			},
			"ldaps_conn_reuse_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify LDAPS connection reuse idle timeout value (in seconds) (Specify idle timeout value (in seconds), default is 0 (not reuse LDAPS connection))",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify LDAP authentication server name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 389, Description: "Specify the LDAP server's authentication port, default is 389",
			},
			"port_hm": {
				Type: schema.TypeString, Optional: true, Description: "Check port's health status",
			},
			"port_hm_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured port health check configuration",
			},
			"prompt_pw_change_before_exp": {
				Type: schema.TypeInt, Optional: true, Description: "Prompt user to change password before expiration in N days. This option only takes effect when server type is AD (Prompt user to change password before expiration in N days, default is not to prompt the user)",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Default: "ldap", Description: "'ldap': Use LDAP (default); 'ldaps': Use LDAP over SSL; 'starttls': Use LDAP StartTLS;",
			},
			"pwdmaxage": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the LDAP server's default password expiration time (in seconds) (The LDAP server's default password expiration time (in seconds), default is 0 (no expiration))",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'admin-bind-success': Admin Bind Success; 'admin-bind-failure': Admin Bind Failure; 'bind-success': User Bind Success; 'bind-failure': User Bind Failure; 'search-success': Search Success; 'search-failure': Search Failure; 'authorize-success': Authorization Success; 'authorize-failure': Authorization Failure; 'timeout-error': Timeout; 'other-error': Other Error; 'request': Request; 'ssl-session-created': TLS/SSL Session Created; 'ssl-session-failure': TLS/SSL Session Failure; 'pw_expiry': Password expiry; 'pw_change_success': Password change success; 'pw_change_failure': Password change failure;",
						},
					},
				},
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "secret password",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify timout for LDAP, default is 10 seconds (The timeout, default is 10 seconds)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServerLdapInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdapInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerLdapInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerLdapInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdapInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerLdapInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerLdapInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdapInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerLdapInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdapInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationServerLdapInstanceDeriveBindDn(d []interface{}) edpt.AamAuthenticationServerLdapInstanceDeriveBindDn {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapInstanceDeriveBindDn
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UsernameAttr = in["username_attr"].(string)
	}
	return ret
}

func getObjectAamAuthenticationServerLdapInstanceHost(d []interface{}) edpt.AamAuthenticationServerLdapInstanceHost {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapInstanceHost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerLdapInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerLdapInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerLdapInstance(d *schema.ResourceData) edpt.AamAuthenticationServerLdapInstance {
	var ret edpt.AamAuthenticationServerLdapInstance
	ret.Inst.AdminDn = d.Get("admin_dn").(string)
	ret.Inst.AdminSecret = d.Get("admin_secret").(int)
	ret.Inst.AuthType = d.Get("auth_type").(string)
	ret.Inst.Base = d.Get("base").(string)
	ret.Inst.BindWithDn = d.Get("bind_with_dn").(int)
	ret.Inst.CaCert = d.Get("ca_cert").(string)
	ret.Inst.DefaultDomain = d.Get("default_domain").(string)
	ret.Inst.DeriveBindDn = getObjectAamAuthenticationServerLdapInstanceDeriveBindDn(d.Get("derive_bind_dn").([]interface{}))
	ret.Inst.DnAttribute = d.Get("dn_attribute").(string)
	//omit encrypted
	ret.Inst.HealthCheck = d.Get("health_check").(int)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckString = d.Get("health_check_string").(string)
	ret.Inst.Host = getObjectAamAuthenticationServerLdapInstanceHost(d.Get("host").([]interface{}))
	ret.Inst.LdapsConnReuseIdleTimeout = d.Get("ldaps_conn_reuse_idle_timeout").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortHm = d.Get("port_hm").(string)
	ret.Inst.PortHmDisable = d.Get("port_hm_disable").(int)
	ret.Inst.PromptPwChangeBeforeExp = d.Get("prompt_pw_change_before_exp").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Pwdmaxage = d.Get("pwdmaxage").(int)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServerLdapInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
