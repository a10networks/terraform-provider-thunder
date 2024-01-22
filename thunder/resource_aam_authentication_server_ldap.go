package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerLdap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server_ldap`: LDAP Authentication Server\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerLdapCreate,
		UpdateContext: resourceAamAuthenticationServerLdapUpdate,
		ReadContext:   resourceAamAuthenticationServerLdapRead,
		DeleteContext: resourceAamAuthenticationServerLdapDelete,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify LDAP authentication server name",
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
						"base": {
							Type: schema.TypeString, Optional: true, Description: "Specify the LDAP server's search base",
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
						"pwdmaxage": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the LDAP server's default password expiration time (in seconds) (The LDAP server's default password expiration time (in seconds), default is 0 (no expiration))",
						},
						"admin_dn": {
							Type: schema.TypeString, Optional: true, Description: "The LDAP server's admin DN",
						},
						"admin_secret": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the LDAP server's admin secret password",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "secret password",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify timout for LDAP, default is 10 seconds (The timeout, default is 10 seconds)",
						},
						"dn_attribute": {
							Type: schema.TypeString, Optional: true, Default: "cn", Description: "Specify Distinguished Name attribute, default is CN",
						},
						"default_domain": {
							Type: schema.TypeString, Optional: true, Description: "Specify default domain for LDAP",
						},
						"bind_with_dn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce using DN for LDAP binding(All user input name will be used to create DN)",
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
						"health_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check server's health status",
						},
						"health_check_string": {
							Type: schema.TypeString, Optional: true, Description: "Health monitor name",
						},
						"health_check_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Default: "ldap", Description: "'ldap': Use LDAP (default); 'ldaps': Use LDAP over SSL; 'starttls': Use LDAP StartTLS;",
						},
						"ca_cert": {
							Type: schema.TypeString, Optional: true, Description: "Specify the LDAPS CA cert filename (Trusted LDAPS CA cert filename)",
						},
						"ldaps_conn_reuse_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify LDAPS connection reuse idle timeout value (in seconds) (Specify idle timeout value (in seconds), default is 0 (not reuse LDAPS connection))",
						},
						"auth_type": {
							Type: schema.TypeString, Optional: true, Description: "'ad': Active Directory. Default; 'open-ldap': OpenLDAP;",
						},
						"prompt_pw_change_before_exp": {
							Type: schema.TypeInt, Optional: true, Description: "Prompt user to change password before expiration in N days. This option only takes effect when server type is AD (Prompt user to change password before expiration in N days, default is not to prompt the user)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
						"packet_capture_template": {
							Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'admin-bind-success': Total Admin Bind Success; 'admin-bind-failure': Total Admin Bind Failure; 'bind-success': Total User Bind Success; 'bind-failure': Total User Bind Failure; 'search-success': Total Search Success; 'search-failure': Total Search Failure; 'authorize-success': Total Authorization Success; 'authorize-failure': Total Authorization Failure; 'timeout-error': Total Timeout; 'other-error': Total Other Error; 'request': Total Request; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response; 'job-start-error': Total Job Start Error; 'polling-control-error': Total Polling Control Error; 'ssl-session-created': TLS/SSL Session Created; 'ssl-session-failure': TLS/SSL Session Failure; 'ldaps-idle-conn-num': LDAPS Idle Connection Number; 'ldaps-inuse-conn-num': LDAPS In-use Connection Number; 'pw-expiry': Total Password expiry; 'pw-change-success': Total password change success; 'pw-change-failure': Total password change failure;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServerLdapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerLdapRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerLdapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerLdapRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerLdapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerLdapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationServerLdapInstanceList(d []interface{}) []edpt.AamAuthenticationServerLdapInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapInstanceList
		oi.Name = in["name"].(string)
		oi.Host = getObjectAamAuthenticationServerLdapInstanceListHost(in["host"].([]interface{}))
		oi.Base = in["base"].(string)
		oi.Port = in["port"].(int)
		oi.PortHm = in["port_hm"].(string)
		oi.PortHmDisable = in["port_hm_disable"].(int)
		oi.Pwdmaxage = in["pwdmaxage"].(int)
		oi.AdminDn = in["admin_dn"].(string)
		oi.AdminSecret = in["admin_secret"].(int)
		oi.SecretString = in["secret_string"].(string)
		//omit encrypted
		oi.Timeout = in["timeout"].(int)
		oi.DnAttribute = in["dn_attribute"].(string)
		oi.DefaultDomain = in["default_domain"].(string)
		oi.BindWithDn = in["bind_with_dn"].(int)
		oi.DeriveBindDn = getObjectAamAuthenticationServerLdapInstanceListDeriveBindDn(in["derive_bind_dn"].([]interface{}))
		oi.HealthCheck = in["health_check"].(int)
		oi.HealthCheckString = in["health_check_string"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.CaCert = in["ca_cert"].(string)
		oi.LdapsConnReuseIdleTimeout = in["ldaps_conn_reuse_idle_timeout"].(int)
		oi.AuthType = in["auth_type"].(string)
		oi.PromptPwChangeBeforeExp = in["prompt_pw_change_before_exp"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationServerLdapInstanceListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerLdapInstanceListHost(d []interface{}) edpt.AamAuthenticationServerLdapInstanceListHost {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapInstanceListHost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getObjectAamAuthenticationServerLdapInstanceListDeriveBindDn(d []interface{}) edpt.AamAuthenticationServerLdapInstanceListDeriveBindDn {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapInstanceListDeriveBindDn
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UsernameAttr = in["username_attr"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerLdapInstanceListSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerLdapInstanceListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapInstanceListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapInstanceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerLdapSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerLdapSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerLdap(d *schema.ResourceData) edpt.AamAuthenticationServerLdap {
	var ret edpt.AamAuthenticationServerLdap
	ret.Inst.InstanceList = getSliceAamAuthenticationServerLdapInstanceList(d.Get("instance_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServerLdapSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
