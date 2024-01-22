package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationLogonHttpAuthenticateInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_logon_http_authenticate_instance`: HTTP-authenticate Logon\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationLogonHttpAuthenticateInstanceCreate,
		UpdateContext: resourceAamAuthenticationLogonHttpAuthenticateInstanceUpdate,
		ReadContext:   resourceAamAuthenticationLogonHttpAuthenticateInstanceRead,
		DeleteContext: resourceAamAuthenticationLogonHttpAuthenticateInstanceDelete,

		Schema: map[string]*schema.Schema{
			"account_lock": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lock the account when the failed logon attempts is exceeded",
			},
			"auth_method": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"basic": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"basic_realm": {
										Type: schema.TypeString, Optional: true, Description: "Specify realm for basic logon",
									},
									"challenge_response_form": {
										Type: schema.TypeString, Optional: true, Description: "Specify challenge-response form for RSA-RADIUS authentication",
									},
									"challenge_page": {
										Type: schema.TypeString, Optional: true, Description: "Specify challenge page name for RSA-RADIUS",
									},
									"challenge_variable": {
										Type: schema.TypeString, Optional: true, Description: "Specify challenge variable name",
									},
									"new_pin_page": {
										Type: schema.TypeString, Optional: true, Description: "Specify new PIN page name for RSA-RADIUS",
									},
									"next_token_page": {
										Type: schema.TypeString, Optional: true, Description: "Specify next-token page name for RSA-RADIUS",
									},
									"new_pin_variable": {
										Type: schema.TypeString, Optional: true, Description: "Specify new PIN variable name",
									},
									"next_token_variable": {
										Type: schema.TypeString, Optional: true, Description: "Specify next-token variable name",
									},
									"basic_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Basic logon",
									},
								},
							},
						},
						"ntlm": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ntlm_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable NTLM logon",
									},
								},
							},
						},
						"negotiate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"negotiate_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SPENGO logon",
									},
								},
							},
						},
					},
				},
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 1800, Description: "The time an account remains locked in seconds (default 1800)",
			},
			"hsts_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Set HSTS policy expired timeout in seconds, 0 means to disable HSTS policy",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify HTTP-Authenticate logon name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Maximum number of consecutive failed logon attempts (default 3)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'spn_krb_request': SPN Kerberos Request; 'spn_krb_success': SPN Kerberos Success; 'spn_krb_faiure': SPN Kerberos Failure;",
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
func resourceAamAuthenticationLogonHttpAuthenticateInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonHttpAuthenticateInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonHttpAuthenticateInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationLogonHttpAuthenticateInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationLogonHttpAuthenticateInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonHttpAuthenticateInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonHttpAuthenticateInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationLogonHttpAuthenticateInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationLogonHttpAuthenticateInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonHttpAuthenticateInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonHttpAuthenticateInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationLogonHttpAuthenticateInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonHttpAuthenticateInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonHttpAuthenticateInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethod(d []interface{}) edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethod {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethod
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Basic = getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethodBasic(in["basic"].([]interface{}))
		ret.Ntlm = getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNtlm(in["ntlm"].([]interface{}))
		ret.Negotiate = getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNegotiate(in["negotiate"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethodBasic(d []interface{}) edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodBasic {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodBasic
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BasicRealm = in["basic_realm"].(string)
		ret.ChallengeResponseForm = in["challenge_response_form"].(string)
		ret.ChallengePage = in["challenge_page"].(string)
		ret.ChallengeVariable = in["challenge_variable"].(string)
		ret.NewPinPage = in["new_pin_page"].(string)
		ret.NextTokenPage = in["next_token_page"].(string)
		ret.NewPinVariable = in["new_pin_variable"].(string)
		ret.NextTokenVariable = in["next_token_variable"].(string)
		ret.BasicEnable = in["basic_enable"].(int)
	}
	return ret
}

func getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNtlm(d []interface{}) edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNtlm {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNtlm
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NtlmEnable = in["ntlm_enable"].(int)
	}
	return ret
}

func getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNegotiate(d []interface{}) edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNegotiate {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNegotiate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NegotiateEnable = in["negotiate_enable"].(int)
	}
	return ret
}

func getSliceAamAuthenticationLogonHttpAuthenticateInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationLogonHttpAuthenticateInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationLogonHttpAuthenticateInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationLogonHttpAuthenticateInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationLogonHttpAuthenticateInstance(d *schema.ResourceData) edpt.AamAuthenticationLogonHttpAuthenticateInstance {
	var ret edpt.AamAuthenticationLogonHttpAuthenticateInstance
	ret.Inst.AccountLock = d.Get("account_lock").(int)
	ret.Inst.AuthMethod = getObjectAamAuthenticationLogonHttpAuthenticateInstanceAuthMethod(d.Get("auth_method").([]interface{}))
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.HstsTimeout = d.Get("hsts_timeout").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationLogonHttpAuthenticateInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
