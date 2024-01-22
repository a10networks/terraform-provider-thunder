package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationLogonFormBased() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_logon_form_based`: Form-based Authentication Logon\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationLogonFormBasedCreate,
		UpdateContext: resourceAamAuthenticationLogonFormBasedUpdate,
		ReadContext:   resourceAamAuthenticationLogonFormBasedRead,
		DeleteContext: resourceAamAuthenticationLogonFormBasedDelete,

		Schema: map[string]*schema.Schema{
			"account_lock": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lock the account when the failed logon attempts is exceeded",
			},
			"challenge_variable": {
				Type: schema.TypeString, Optional: true, Description: "Specify challenge variable name in form submission",
			},
			"cp_page_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"changepassword_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify changepassword form submission action url (changepassword action url)",
						},
						"cp_user_enum": {
							Type: schema.TypeString, Optional: true, Description: "'changepassword-username-variable': Specify username variable name in form submission;",
						},
						"cp_user_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify username variable name",
						},
						"cp_old_pwd_enum": {
							Type: schema.TypeString, Optional: true, Description: "'changepassword-old-password-variable': Specify old password variable name in form submission;",
						},
						"cp_old_pwd_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify old password variable name",
						},
						"cp_new_pwd_enum": {
							Type: schema.TypeString, Optional: true, Description: "'changepassword-new-password-variable': Specify new password variable name in form submission;",
						},
						"cp_new_pwd_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify new password variable name",
						},
						"cp_cfm_pwd_enum": {
							Type: schema.TypeString, Optional: true, Description: "'changepassword-password-confirm-variable': Specify password confirm variable name in form submission;",
						},
						"cp_cfm_pwd_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify password confirm variable name",
						},
					},
				},
			},
			"csp_support": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"none": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set CSP frame-ancestors to none (also X-Frame-Options deny)",
						},
						"self": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set CSP frame-ancestors to self (also X-Frame-Options same-origin)",
						},
						"specificuri": {
							Type: schema.TypeString, Optional: true, Description: "Set customized CSP frame-ancestors (maximum 2 URIs can be set)",
						},
						"optional_second_uri": {
							Type: schema.TypeString, Optional: true, Description: "Set optional second customized CSP URI",
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
			"logon_page_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify form submission action url",
						},
						"username_variable": {
							Type: schema.TypeString, Optional: true, Description: "Specify username variable name in form submission",
						},
						"password_variable": {
							Type: schema.TypeString, Optional: true, Description: "Specify password variable name in form submission",
						},
						"passcode_variable": {
							Type: schema.TypeString, Optional: true, Description: "Specify passcode variable name in form submission",
						},
						"captcha_variable": {
							Type: schema.TypeString, Optional: true, Description: "Specify captcha variable name in form submission",
						},
						"login_failure_message": {
							Type: schema.TypeString, Optional: true, Description: "Specify login failure message shown in logon page (Specify error string, default is \"Invalid username or password. Please try again.\")",
						},
						"authz_failure_message": {
							Type: schema.TypeString, Optional: true, Description: "Specify authorization failure message shown in logon page (Specify error string, default is \"Authorization failed. Please contact your system administrator.\")",
						},
						"disable_change_password_link": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't display change password link on logon page forcibly even backend authentication server supports it (LDAP or Kerberos)",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify form-based authentication logon name",
			},
			"new_pin_variable": {
				Type: schema.TypeString, Optional: true, Description: "Specify new-pin variable name in form submission",
			},
			"next_token_variable": {
				Type: schema.TypeString, Optional: true, Description: "Specify next-token variable name in form submission",
			},
			"notify_cp_page_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notifychangepassword_change_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify change password action url for notifychangepassword form",
						},
						"notifychangepassword_continue_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify continue action url for notifychangepassword form",
						},
					},
				},
			},
			"portal": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_portal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use default portal",
						},
						"portal_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify portal name",
						},
						"logon": {
							Type: schema.TypeString, Optional: true, Description: "Specify logon page name",
						},
						"failpage": {
							Type: schema.TypeString, Optional: true, Description: "Specify logon fail page name (portal fail page name)",
						},
						"changepasswordpage": {
							Type: schema.TypeString, Optional: true, Description: "Specify change password page name",
						},
						"notifychangepasswordpage": {
							Type: schema.TypeString, Optional: true, Description: "Specify change password notification page name",
						},
						"challenge_page": {
							Type: schema.TypeString, Optional: true, Description: "Specify challenge page name for RSA-RADIUS",
						},
						"new_pin_page": {
							Type: schema.TypeString, Optional: true, Description: "Specify new PIN page name for RSA-RADIUS",
						},
						"next_token_page": {
							Type: schema.TypeString, Optional: true, Description: "Specify next token page name for RSA-RADIUS",
						},
					},
				},
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Maximum number of consecutive failed logon attempts (default 3)",
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
func resourceAamAuthenticationLogonFormBasedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonFormBasedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonFormBased(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationLogonFormBasedRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationLogonFormBasedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonFormBasedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonFormBased(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationLogonFormBasedRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationLogonFormBasedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonFormBasedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonFormBased(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationLogonFormBasedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonFormBasedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonFormBased(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationLogonFormBasedCpPageCfg(d []interface{}) edpt.AamAuthenticationLogonFormBasedCpPageCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonFormBasedCpPageCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ChangepasswordUrl = in["changepassword_url"].(string)
		ret.CpUserEnum = in["cp_user_enum"].(string)
		ret.CpUserVar = in["cp_user_var"].(string)
		ret.CpOldPwdEnum = in["cp_old_pwd_enum"].(string)
		ret.CpOldPwdVar = in["cp_old_pwd_var"].(string)
		ret.CpNewPwdEnum = in["cp_new_pwd_enum"].(string)
		ret.CpNewPwdVar = in["cp_new_pwd_var"].(string)
		ret.CpCfmPwdEnum = in["cp_cfm_pwd_enum"].(string)
		ret.CpCfmPwdVar = in["cp_cfm_pwd_var"].(string)
	}
	return ret
}

func getObjectAamAuthenticationLogonFormBasedCspSupport(d []interface{}) edpt.AamAuthenticationLogonFormBasedCspSupport {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonFormBasedCspSupport
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.None = in["none"].(int)
		ret.Self = in["self"].(int)
		ret.Specificuri = in["specificuri"].(string)
		ret.OptionalSecondUri = in["optional_second_uri"].(string)
	}
	return ret
}

func getObjectAamAuthenticationLogonFormBasedLogonPageCfg(d []interface{}) edpt.AamAuthenticationLogonFormBasedLogonPageCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonFormBasedLogonPageCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActionUrl = in["action_url"].(string)
		ret.UsernameVariable = in["username_variable"].(string)
		ret.PasswordVariable = in["password_variable"].(string)
		ret.PasscodeVariable = in["passcode_variable"].(string)
		ret.CaptchaVariable = in["captcha_variable"].(string)
		ret.LoginFailureMessage = in["login_failure_message"].(string)
		ret.AuthzFailureMessage = in["authz_failure_message"].(string)
		ret.DisableChangePasswordLink = in["disable_change_password_link"].(int)
	}
	return ret
}

func getObjectAamAuthenticationLogonFormBasedNotifyCpPageCfg(d []interface{}) edpt.AamAuthenticationLogonFormBasedNotifyCpPageCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonFormBasedNotifyCpPageCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NotifychangepasswordChangeUrl = in["notifychangepassword_change_url"].(string)
		ret.NotifychangepasswordContinueUrl = in["notifychangepassword_continue_url"].(string)
	}
	return ret
}

func getObjectAamAuthenticationLogonFormBasedPortal(d []interface{}) edpt.AamAuthenticationLogonFormBasedPortal {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonFormBasedPortal
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortal = in["default_portal"].(int)
		ret.PortalName = in["portal_name"].(string)
		ret.Logon = in["logon"].(string)
		ret.Failpage = in["failpage"].(string)
		ret.Changepasswordpage = in["changepasswordpage"].(string)
		ret.Notifychangepasswordpage = in["notifychangepasswordpage"].(string)
		ret.ChallengePage = in["challenge_page"].(string)
		ret.NewPinPage = in["new_pin_page"].(string)
		ret.NextTokenPage = in["next_token_page"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationLogonFormBased(d *schema.ResourceData) edpt.AamAuthenticationLogonFormBased {
	var ret edpt.AamAuthenticationLogonFormBased
	ret.Inst.AccountLock = d.Get("account_lock").(int)
	ret.Inst.ChallengeVariable = d.Get("challenge_variable").(string)
	ret.Inst.CpPageCfg = getObjectAamAuthenticationLogonFormBasedCpPageCfg(d.Get("cp_page_cfg").([]interface{}))
	ret.Inst.CspSupport = getObjectAamAuthenticationLogonFormBasedCspSupport(d.Get("csp_support").([]interface{}))
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.HstsTimeout = d.Get("hsts_timeout").(int)
	ret.Inst.LogonPageCfg = getObjectAamAuthenticationLogonFormBasedLogonPageCfg(d.Get("logon_page_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NewPinVariable = d.Get("new_pin_variable").(string)
	ret.Inst.NextTokenVariable = d.Get("next_token_variable").(string)
	ret.Inst.NotifyCpPageCfg = getObjectAamAuthenticationLogonFormBasedNotifyCpPageCfg(d.Get("notify_cp_page_cfg").([]interface{}))
	ret.Inst.Portal = getObjectAamAuthenticationLogonFormBasedPortal(d.Get("portal").([]interface{}))
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
