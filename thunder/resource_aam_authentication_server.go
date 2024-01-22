package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server`: Authentication server configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerCreate,
		UpdateContext: resourceAamAuthenticationServerUpdate,
		ReadContext:   resourceAamAuthenticationServerRead,
		DeleteContext: resourceAamAuthenticationServerDelete,

		Schema: map[string]*schema.Schema{
			"ldap": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
					},
				},
			},
			"ocsp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'stapling-certificate-good': Total OCSP Stapling Good Certificate Response; 'stapling-certificate-revoked': Total OCSP Stapling Revoked Certificate Response; 'stapling-certificate-unknown': Total OCSP Stapling Unknown Certificate Response; 'stapling-request-normal': Total OSCP Stapling Normal Request; 'stapling-request-dropped': Total OCSP Stapling Dropped Request; 'stapling-response-success': Total OCSP Stapling Success Response; 'stapling-response-failure': Total OCSP Stapling Failure Response; 'stapling-response-error': Total OCSP Stapling Error Response; 'stapling-response-timeout': Total OCSP Stapling Timeout Response; 'stapling-response-other': Total OCSP Stapling Other Response; 'request-normal': Total OSCP Normal Request; 'request-dropped': Total OCSP Dropped Request; 'response-success': Total OCSP Success Response; 'response-failure': Total OCSP Failure Response; 'response-error': Total OCSP Error Response; 'response-timeout': Total OCSP Timeout Response; 'response-other': Total OCSP Other Response; 'job-start-error': Total OCSP Job Start Error; 'polling-control-error': Total OCSP Polling Control Error;",
									},
								},
							},
						},
						"instance_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Specify OCSP authentication server name",
									},
									"url": {
										Type: schema.TypeString, Optional: true, Description: "Specify the OCSP server's address (Format: http://host[:port]/) (The OCSP server's address(Format: http://host[:port]/))",
									},
									"responder_ca": {
										Type: schema.TypeString, Optional: true, Description: "Specify the trusted OCSP responder's CA cert filename",
									},
									"responder_cert": {
										Type: schema.TypeString, Optional: true, Description: "Specify the trusted OCSP responder's cert filename",
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
									"port_health_check": {
										Type: schema.TypeString, Optional: true, Description: "Check port's health status",
									},
									"port_health_check_disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured port health check configuration",
									},
									"http_version": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set HTTP version (default 1.0)",
									},
									"version_type": {
										Type: schema.TypeString, Optional: true, Description: "'1.1': HTTP version 1.1;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'request': Request; 'certificate-good': Good Certificate Response; 'certificate-revoked': Revoked Certificate Response; 'certificate-unknown': Unknown Certificate Response; 'timeout': Timeout; 'fail': Handle OCSP response failed; 'stapling-request': OCSP Stapling Request Send; 'stapling-certificate-good': OCSP Stapling Good Certificate Response; 'stapling-certificate-revoked': OCSP Stapling Revoked Certificate Response; 'stapling-certificate-unknown': OCSP Stapling Unknown Certificate Response; 'stapling-timeout': OCSP Stapling Timeout; 'stapling-fail': Handle OCSP response failed;",
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
					},
				},
			},
			"radius": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'authen_success': Total Authentication Success; 'authen_failure': Total Authentication Failure; 'authorize_success': Total Authorization Success; 'authorize_failure': Total Authorization Failure; 'access_challenge': Total Access-Challenge Message Receive; 'timeout_error': Total Timeout; 'other_error': Total Other Error; 'request': Total Request; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response; 'job-start-error': Total Job Start Error; 'polling-control-error': Total Polling Control Error; 'accounting-request-sent': Accounting-Request Sent; 'accounting-success': Accounting Success; 'accounting-failure': Accounting Failure;",
									},
								},
							},
						},
						"instance_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Specify RADIUS authentication server name",
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
									"secret": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the RADIUS server's secret",
									},
									"secret_string": {
										Type: schema.TypeString, Optional: true, Description: "The RADIUS server's secret",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Default: 1812, Description: "Specify the RADIUS server's authentication port, default is 1812",
									},
									"port_hm": {
										Type: schema.TypeString, Optional: true, Description: "Check port's health status",
									},
									"port_hm_disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured port health check configuration",
									},
									"interval": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the interval time for resend the request (second), default is 3 seconds (The interval time(second), default is 3 seconds)",
									},
									"retry": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the retry number for resend the request, default is 5 (The retry number, default is 5)",
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
									"accounting_port": {
										Type: schema.TypeInt, Optional: true, Default: 1813, Description: "Specify the RADIUS server's accounting port, default is 1813",
									},
									"acct_port_hm": {
										Type: schema.TypeString, Optional: true, Description: "Specify accounting port health check method",
									},
									"acct_port_hm_disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured accounting port health check configuration",
									},
									"auth_type": {
										Type: schema.TypeString, Optional: true, Description: "'pap': PAP authentication. Default; 'mschapv2': MS-CHAPv2 authentication; 'mschapv2-pap': Use MS-CHAPv2 first. If server doesn't support it, try PAP;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'authen_success': Authentication Success; 'authen_failure': Authentication Failure; 'authorize_success': Authorization Success; 'authorize_failure': Authorization Failure; 'access_challenge': Access-Challenge Message Receive; 'timeout_error': Timeout; 'other_error': Other Error; 'request': Request; 'accounting-request-sent': Accounting-Request Sent; 'accounting-success': Accounting Success; 'accounting-failure': Accounting Failure;",
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
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"windows": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'kerberos-request-send': Total Kerberos Request; 'kerberos-response-get': Total Kerberos Response; 'kerberos-timeout-error': Total Kerberos Timeout; 'kerberos-other-error': Total Kerberos Other Error; 'ntlm-authentication-success': Total NTLM Authentication Success; 'ntlm-authentication-failure': Total NTLM Authentication Failure; 'ntlm-proto-negotiation-success': Total NTLM Protocol Negotiation Success; 'ntlm-proto-negotiation-failure': Total NTLM Protocol Negotiation Failure; 'ntlm-session-setup-success': Total NTLM Session Setup Success; 'ntlm-session-setup-failed': Total NTLM Session Setup Failure; 'kerberos-request-normal': Total Kerberos Normal Request; 'kerberos-request-dropped': Total Kerberos Dropped Request; 'kerberos-response-success': Total Kerberos Success Response; 'kerberos-response-failure': Total Kerberos Failure Response; 'kerberos-response-error': Total Kerberos Error Response; 'kerberos-response-timeout': Total Kerberos Timeout Response; 'kerberos-response-other': Total Kerberos Other Response; 'kerberos-job-start-error': Total Kerberos Job Start Error; 'kerberos-polling-control-error': Total Kerberos Polling Control Error; 'ntlm-prepare-req-success': Total NTLM Prepare Request Success; 'ntlm-prepare-req-failed': Total NTLM Prepare Request Failed; 'ntlm-timeout-error': Total NTLM Timeout; 'ntlm-other-error': Total NTLM Other Error; 'ntlm-request-normal': Total NTLM Normal Request; 'ntlm-request-dropped': Total NTLM Dropped Request; 'ntlm-response-success': Total NTLM Success Response; 'ntlm-response-failure': Total NTLM Failure Response; 'ntlm-response-error': Total NTLM Error Response; 'ntlm-response-timeout': Total NTLM Timeout Response; 'ntlm-response-other': Total NTLM Other Response; 'ntlm-job-start-error': Total NTLM Job Start Error; 'ntlm-polling-control-error': Total NTLM Polling Control Error; 'kerberos-pw-expiry': Total Kerberos password expiry; 'kerberos-pw-change-success': Total Kerberos password change success; 'kerberos-pw-change-failure': Total Kerberos password change failure; 'kerberos-validate-kdc-success': Total Kerberos KDC Validation Success; 'kerberos-validate-kdc-failure': Total Kerberos KDC Validation Failure; 'kerberos-generate-kdc-keytab-success': Total Kerberos KDC Keytab Generation Success; 'kerberos-generate-kdc-keytab-failure': Total Kerberos KDC Keytab Generation Failure; 'kerberos-delete-kdc-keytab-success': Total Kerberos KDC Keytab Deletion Success; 'kerberos-delete-kdc-keytab-failure': Total Kerberos KDC Keytab Deletion Failure; 'kerberos-kdc-keytab-count': Current Kerberos KDC Keytab Count;",
									},
								},
							},
						},
						"instance_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Specify Windows authentication server name",
									},
									"host": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hostip": {
													Type: schema.TypeString, Optional: true, Description: "Specify the Windows server's hostname(Length 1-31) or IP address",
												},
												"hostipv6": {
													Type: schema.TypeString, Optional: true, Description: "Specify the Windows server's IPV6 address",
												},
											},
										},
									},
									"timeout": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify connection timeout to server, default is 10 seconds",
									},
									"auth_protocol": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ntlm_disable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable NTLM authentication protocol",
												},
												"ntlm_version": {
													Type: schema.TypeInt, Optional: true, Default: 2, Description: "Specify NTLM version, default is 2",
												},
												"ntlm_health_check": {
													Type: schema.TypeString, Optional: true, Description: "Check NTLM port's health status",
												},
												"ntlm_health_check_disable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured NTLM port health check configuration",
												},
												"kerberos_disable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Kerberos authentication protocol",
												},
												"kerberos_port": {
													Type: schema.TypeInt, Optional: true, Default: 88, Description: "Specify the Kerberos port, default is 88",
												},
												"kport_hm": {
													Type: schema.TypeString, Optional: true, Description: "Check Kerberos port's health status",
												},
												"kport_hm_disable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured Kerberos port health check configuration",
												},
												"kerberos_password_change_port": {
													Type: schema.TypeInt, Optional: true, Default: 464, Description: "Specify the Kerbros password change port, default is 464",
												},
												"kdc_validate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable KDC validation",
												},
												"kerberos_kdc_validation": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kdc_spn": {
																Type: schema.TypeString, Optional: true, Description: "Specify SPN for KDC validation",
															},
															"kdc_account": {
																Type: schema.TypeString, Optional: true, Description: "Specify account for KDC validation",
															},
															"kdc_password": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify account password",
															},
															"kdc_pwd": {
																Type: schema.TypeString, Optional: true, Description: "Account password",
															},
														},
													},
												},
											},
										},
									},
									"realm": {
										Type: schema.TypeString, Optional: true, Description: "Specify realm of Windows server",
									},
									"support_apacheds_kdc": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable weak cipher (DES CRC/MD5/MD4) and merge AS-REQ in single packet",
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
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'krb_send_req_success': Kerberos Request; 'krb_get_resp_success': Kerberos Response; 'krb_timeout_error': Kerberos Timeout; 'krb_other_error': Kerberos Other Error; 'krb_pw_expiry': Kerberos password expiry; 'krb_pw_change_success': Kerberos password change success; 'krb_pw_change_failure': Kerberos password change failure; 'ntlm_proto_nego_success': NTLM Protocol Negotiation Success; 'ntlm_proto_nego_failure': NTLM Protocol Negotiation Failure; 'ntlm_session_setup_success': NTLM Session Setup Success; 'ntlm_session_setup_failure': NTLM Session Setup Failure; 'ntlm_prepare_req_success': NTLM Prepare Request Success; 'ntlm_prepare_req_error': NTLM Prepare Request Error; 'ntlm_auth_success': NTLM Authentication Success; 'ntlm_auth_failure': NTLM Authentication Failure; 'ntlm_timeout_error': NTLM Timeout; 'ntlm_other_error': NTLM Other Error; 'krb_validate_kdc_success': Kerberos KDC Validation Success; 'krb_validate_kdc_failure': Kerberos KDC Validation Failure;",
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
					},
				},
			},
		},
	}
}
func resourceAamAuthenticationServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationServerLdap24(d []interface{}) edpt.AamAuthenticationServerLdap24 {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdap24
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceAamAuthenticationServerLdapSamplingEnable25(in["sampling_enable"].([]interface{}))
		ret.InstanceList = getSliceAamAuthenticationServerLdapInstanceList26(in["instance_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationServerLdapSamplingEnable25(d []interface{}) []edpt.AamAuthenticationServerLdapSamplingEnable25 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapSamplingEnable25, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapSamplingEnable25
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerLdapInstanceList26(d []interface{}) []edpt.AamAuthenticationServerLdapInstanceList26 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapInstanceList26, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapInstanceList26
		oi.Name = in["name"].(string)
		oi.Host = getObjectAamAuthenticationServerLdapInstanceListHost27(in["host"].([]interface{}))
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
		oi.DeriveBindDn = getObjectAamAuthenticationServerLdapInstanceListDeriveBindDn28(in["derive_bind_dn"].([]interface{}))
		oi.HealthCheck = in["health_check"].(int)
		oi.HealthCheckString = in["health_check_string"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.CaCert = in["ca_cert"].(string)
		oi.LdapsConnReuseIdleTimeout = in["ldaps_conn_reuse_idle_timeout"].(int)
		oi.AuthType = in["auth_type"].(string)
		oi.PromptPwChangeBeforeExp = in["prompt_pw_change_before_exp"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationServerLdapInstanceListSamplingEnable29(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerLdapInstanceListHost27(d []interface{}) edpt.AamAuthenticationServerLdapInstanceListHost27 {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapInstanceListHost27
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getObjectAamAuthenticationServerLdapInstanceListDeriveBindDn28(d []interface{}) edpt.AamAuthenticationServerLdapInstanceListDeriveBindDn28 {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapInstanceListDeriveBindDn28
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UsernameAttr = in["username_attr"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerLdapInstanceListSamplingEnable29(d []interface{}) []edpt.AamAuthenticationServerLdapInstanceListSamplingEnable29 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapInstanceListSamplingEnable29, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapInstanceListSamplingEnable29
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerOcsp30(d []interface{}) edpt.AamAuthenticationServerOcsp30 {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOcsp30
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceAamAuthenticationServerOcspSamplingEnable31(in["sampling_enable"].([]interface{}))
		ret.InstanceList = getSliceAamAuthenticationServerOcspInstanceList32(in["instance_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationServerOcspSamplingEnable31(d []interface{}) []edpt.AamAuthenticationServerOcspSamplingEnable31 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspSamplingEnable31, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspSamplingEnable31
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerOcspInstanceList32(d []interface{}) []edpt.AamAuthenticationServerOcspInstanceList32 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspInstanceList32, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspInstanceList32
		oi.Name = in["name"].(string)
		oi.Url = in["url"].(string)
		oi.ResponderCa = in["responder_ca"].(string)
		oi.ResponderCert = in["responder_cert"].(string)
		oi.HealthCheck = in["health_check"].(int)
		oi.HealthCheckString = in["health_check_string"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		oi.PortHealthCheck = in["port_health_check"].(string)
		oi.PortHealthCheckDisable = in["port_health_check_disable"].(int)
		oi.HttpVersion = in["http_version"].(int)
		oi.VersionType = in["version_type"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationServerOcspInstanceListSamplingEnable33(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerOcspInstanceListSamplingEnable33(d []interface{}) []edpt.AamAuthenticationServerOcspInstanceListSamplingEnable33 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspInstanceListSamplingEnable33, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspInstanceListSamplingEnable33
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerRadius34(d []interface{}) edpt.AamAuthenticationServerRadius34 {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerRadius34
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceAamAuthenticationServerRadiusSamplingEnable35(in["sampling_enable"].([]interface{}))
		ret.InstanceList = getSliceAamAuthenticationServerRadiusInstanceList36(in["instance_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationServerRadiusSamplingEnable35(d []interface{}) []edpt.AamAuthenticationServerRadiusSamplingEnable35 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusSamplingEnable35, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusSamplingEnable35
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerRadiusInstanceList36(d []interface{}) []edpt.AamAuthenticationServerRadiusInstanceList36 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusInstanceList36, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusInstanceList36
		oi.Name = in["name"].(string)
		oi.Host = getObjectAamAuthenticationServerRadiusInstanceListHost37(in["host"].([]interface{}))
		oi.Secret = in["secret"].(int)
		oi.SecretString = in["secret_string"].(string)
		//omit encrypted
		oi.Port = in["port"].(int)
		oi.PortHm = in["port_hm"].(string)
		oi.PortHmDisable = in["port_hm_disable"].(int)
		oi.Interval = in["interval"].(int)
		oi.Retry = in["retry"].(int)
		oi.HealthCheck = in["health_check"].(int)
		oi.HealthCheckString = in["health_check_string"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		oi.AccountingPort = in["accounting_port"].(int)
		oi.AcctPortHm = in["acct_port_hm"].(string)
		oi.AcctPortHmDisable = in["acct_port_hm_disable"].(int)
		oi.AuthType = in["auth_type"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationServerRadiusInstanceListSamplingEnable38(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerRadiusInstanceListHost37(d []interface{}) edpt.AamAuthenticationServerRadiusInstanceListHost37 {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerRadiusInstanceListHost37
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerRadiusInstanceListSamplingEnable38(d []interface{}) []edpt.AamAuthenticationServerRadiusInstanceListSamplingEnable38 {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusInstanceListSamplingEnable38, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusInstanceListSamplingEnable38
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerWindows39(d []interface{}) edpt.AamAuthenticationServerWindows39 {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindows39
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceAamAuthenticationServerWindowsSamplingEnable(in["sampling_enable"].([]interface{}))
		ret.InstanceList = getSliceAamAuthenticationServerWindowsInstanceList(in["instance_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationServerWindowsSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerWindowsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerWindowsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerWindowsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerWindowsInstanceList(d []interface{}) []edpt.AamAuthenticationServerWindowsInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerWindowsInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerWindowsInstanceList
		oi.Name = in["name"].(string)
		oi.Host = getObjectAamAuthenticationServerWindowsInstanceListHost(in["host"].([]interface{}))
		oi.Timeout = in["timeout"].(int)
		oi.AuthProtocol = getObjectAamAuthenticationServerWindowsInstanceListAuthProtocol(in["auth_protocol"].([]interface{}))
		oi.Realm = in["realm"].(string)
		oi.SupportApachedsKdc = in["support_apacheds_kdc"].(int)
		oi.HealthCheck = in["health_check"].(int)
		oi.HealthCheckString = in["health_check_string"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationServerWindowsInstanceListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerWindowsInstanceListHost(d []interface{}) edpt.AamAuthenticationServerWindowsInstanceListHost {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsInstanceListHost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getObjectAamAuthenticationServerWindowsInstanceListAuthProtocol(d []interface{}) edpt.AamAuthenticationServerWindowsInstanceListAuthProtocol {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsInstanceListAuthProtocol
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NtlmDisable = in["ntlm_disable"].(int)
		ret.NtlmVersion = in["ntlm_version"].(int)
		ret.NtlmHealthCheck = in["ntlm_health_check"].(string)
		ret.NtlmHealthCheckDisable = in["ntlm_health_check_disable"].(int)
		ret.KerberosDisable = in["kerberos_disable"].(int)
		ret.KerberosPort = in["kerberos_port"].(int)
		ret.KportHm = in["kport_hm"].(string)
		ret.KportHmDisable = in["kport_hm_disable"].(int)
		ret.KerberosPasswordChangePort = in["kerberos_password_change_port"].(int)
		ret.KdcValidate = in["kdc_validate"].(int)
		ret.KerberosKdcValidation = getObjectAamAuthenticationServerWindowsInstanceListAuthProtocolKerberosKdcValidation(in["kerberos_kdc_validation"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationServerWindowsInstanceListAuthProtocolKerberosKdcValidation(d []interface{}) edpt.AamAuthenticationServerWindowsInstanceListAuthProtocolKerberosKdcValidation {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsInstanceListAuthProtocolKerberosKdcValidation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KdcSpn = in["kdc_spn"].(string)
		ret.KdcAccount = in["kdc_account"].(string)
		ret.KdcPassword = in["kdc_password"].(int)
		ret.KdcPwd = in["kdc_pwd"].(string)
		//omit encrypted
	}
	return ret
}

func getSliceAamAuthenticationServerWindowsInstanceListSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerWindowsInstanceListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerWindowsInstanceListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerWindowsInstanceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServer(d *schema.ResourceData) edpt.AamAuthenticationServer {
	var ret edpt.AamAuthenticationServer
	ret.Inst.Ldap = getObjectAamAuthenticationServerLdap24(d.Get("ldap").([]interface{}))
	ret.Inst.Ocsp = getObjectAamAuthenticationServerOcsp30(d.Get("ocsp").([]interface{}))
	ret.Inst.Radius = getObjectAamAuthenticationServerRadius34(d.Get("radius").([]interface{}))
	//omit uuid
	ret.Inst.Windows = getObjectAamAuthenticationServerWindows39(d.Get("windows").([]interface{}))
	return ret
}
