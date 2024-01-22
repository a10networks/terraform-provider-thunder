package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerWindowsInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server_windows_instance`: \"Windows Server, using Kerberos or NTLM for authentication\"\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerWindowsInstanceCreate,
		UpdateContext: resourceAamAuthenticationServerWindowsInstanceUpdate,
		ReadContext:   resourceAamAuthenticationServerWindowsInstanceRead,
		DeleteContext: resourceAamAuthenticationServerWindowsInstanceDelete,

		Schema: map[string]*schema.Schema{
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
							Type: schema.TypeString, Optional: true, Description: "Specify the Windows server's hostname(Length 1-31) or IP address",
						},
						"hostipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify the Windows server's IPV6 address",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Windows authentication server name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"realm": {
				Type: schema.TypeString, Optional: true, Description: "Specify realm of Windows server",
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
			"support_apacheds_kdc": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable weak cipher (DES CRC/MD5/MD4) and merge AS-REQ in single packet",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify connection timeout to server, default is 10 seconds",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServerWindowsInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerWindowsInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerWindowsInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerWindowsInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerWindowsInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerWindowsInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerWindowsInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerWindowsInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerWindowsInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerWindowsInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerWindowsInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerWindowsInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerWindowsInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerWindowsInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationServerWindowsInstanceAuthProtocol(d []interface{}) edpt.AamAuthenticationServerWindowsInstanceAuthProtocol {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsInstanceAuthProtocol
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
		ret.KerberosKdcValidation = getObjectAamAuthenticationServerWindowsInstanceAuthProtocolKerberosKdcValidation(in["kerberos_kdc_validation"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationServerWindowsInstanceAuthProtocolKerberosKdcValidation(d []interface{}) edpt.AamAuthenticationServerWindowsInstanceAuthProtocolKerberosKdcValidation {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsInstanceAuthProtocolKerberosKdcValidation
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

func getObjectAamAuthenticationServerWindowsInstanceHost(d []interface{}) edpt.AamAuthenticationServerWindowsInstanceHost {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsInstanceHost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerWindowsInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerWindowsInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerWindowsInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerWindowsInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerWindowsInstance(d *schema.ResourceData) edpt.AamAuthenticationServerWindowsInstance {
	var ret edpt.AamAuthenticationServerWindowsInstance
	ret.Inst.AuthProtocol = getObjectAamAuthenticationServerWindowsInstanceAuthProtocol(d.Get("auth_protocol").([]interface{}))
	ret.Inst.HealthCheck = d.Get("health_check").(int)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckString = d.Get("health_check_string").(string)
	ret.Inst.Host = getObjectAamAuthenticationServerWindowsInstanceHost(d.Get("host").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Realm = d.Get("realm").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServerWindowsInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SupportApachedsKdc = d.Get("support_apacheds_kdc").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
