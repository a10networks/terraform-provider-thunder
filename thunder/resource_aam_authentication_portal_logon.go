package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationPortalLogon() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_portal_logon`: Logon page configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationPortalLogonCreate,
		UpdateContext: resourceAamAuthenticationPortalLogonUpdate,
		ReadContext:   resourceAamAuthenticationPortalLogonRead,
		DeleteContext: resourceAamAuthenticationPortalLogonDelete,

		Schema: map[string]*schema.Schema{
			"action_url": {
				Type: schema.TypeString, Optional: true, Description: "Specify form action URL in default logon page (Default: /logon.fo)",
			},
			"background": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgfile": {
							Type: schema.TypeString, Optional: true, Description: "Specify background image filename",
						},
						"bgstyle": {
							Type: schema.TypeString, Optional: true, Default: "tile", Description: "'tile': Tile; 'stretch': Stretch; 'fit': Fit;",
						},
						"bgcolor_name": {
							Type: schema.TypeString, Optional: true, Default: "white", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
						},
						"bgcolor_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
						},
					},
				},
			},
			"captcha_type": {
				Type: schema.TypeString, Optional: true, Description: "'reCAPTCHAv2-checkbox': Google reCAPTCHAv2 Checkbox; 'reCAPTCHAv2-invisible': Google reCAPTCHAv2 Invisible; 'reCAPTCHAv3': Google reCAPTCHAv3;",
			},
			"enable_captcha": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CAPTCHA in deafult logon page",
			},
			"enable_passcode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable passcode field in default logon page",
			},
			"fail_msg_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fail_msg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure login failure message in default logon page",
						},
						"fail_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify login failure message (Default: Invalid username or password. Please try again.)",
						},
						"fail_font": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
						},
						"fail_face": {
							Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
						},
						"fail_font_custom": {
							Type: schema.TypeString, Optional: true, Description: "Specify custom font",
						},
						"fail_size": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify font size (Default: 5)",
						},
						"fail_color": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: red)",
						},
						"fail_color_name": {
							Type: schema.TypeString, Optional: true, Default: "red", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
						},
						"fail_color_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
						},
						"authz_fail_msg": {
							Type: schema.TypeString, Optional: true, Description: "Configure authorization failure message in default logon page, its text attributes follow fail-msg's (Specify authorization failure message (Default: Authorization failed. Please contact your system administrator.))",
						},
					},
				},
			},
			"passcode_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"passcode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure passcode text in default logon page",
						},
						"passcode_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify passcode text (Default: Passcode)",
						},
						"passcode_font": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
						},
						"passcode_face": {
							Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
						},
						"passcode_font_custom": {
							Type: schema.TypeString, Optional: true, Description: "Specify custom font",
						},
						"passcode_size": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
						},
						"passcode_color": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
						},
						"passcode_color_name": {
							Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
						},
						"passcode_color_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
						},
					},
				},
			},
			"passcode_var": {
				Type: schema.TypeString, Optional: true, Description: "Specify passcode variable name in default logon page (Default: passcode)",
			},
			"password_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure password text in default logon page",
						},
						"pass_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify password text (Default: Password)",
						},
						"pass_font": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
						},
						"pass_face": {
							Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
						},
						"pass_font_custom": {
							Type: schema.TypeString, Optional: true, Description: "Specify custom font",
						},
						"pass_size": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
						},
						"pass_color": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
						},
						"pass_color_name": {
							Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
						},
						"pass_color_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
						},
					},
				},
			},
			"password_var": {
				Type: schema.TypeString, Optional: true, Description: "Specify password variable name in default logon page (Default: pwd)",
			},
			"recaptcha_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recaptcha_theme": {
							Type: schema.TypeString, Optional: true, Default: "light", Description: "'light': light theme; 'dark': dark theme;",
						},
						"recaptcha_size": {
							Type: schema.TypeString, Optional: true, Default: "normal", Description: "'normal': normal size; 'compact': compact size;",
						},
						"recaptcha_badge": {
							Type: schema.TypeString, Optional: true, Default: "bottom-right", Description: "'bottom-left': bottom left corner; 'bottom-right': bottom right corner;",
						},
						"recaptcha_action": {
							Type: schema.TypeString, Optional: true, Default: "A10_DEFAULT_LOGON", Description: "Specify reCAPTCHA action (Specify action string, only accept alphanumeric, underscore, and slash (Default: A10_DEFAULT_LOGON))",
						},
					},
				},
			},
			"site_key_string": {
				Type: schema.TypeString, Optional: true, Description: "Site key string",
			},
			"submit_text": {
				Type: schema.TypeString, Optional: true, Description: "Specify submit button text in default logon page (Default: Log In)",
			},
			"username_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure username text in default logon page",
						},
						"user_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify username text (Default: User Name)",
						},
						"user_font": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
						},
						"user_face": {
							Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
						},
						"user_font_custom": {
							Type: schema.TypeString, Optional: true, Description: "Specify custom font",
						},
						"user_size": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
						},
						"user_color": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
						},
						"user_color_name": {
							Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
						},
						"user_color_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
						},
					},
				},
			},
			"username_var": {
				Type: schema.TypeString, Optional: true, Description: "Specify username variable name in default logon page (Default: user)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceAamAuthenticationPortalLogonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogon(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalLogonRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationPortalLogonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogon(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalLogonRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationPortalLogonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogon(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationPortalLogonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogon(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationPortalLogonBackground(d []interface{}) edpt.AamAuthenticationPortalLogonBackground {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonBackground
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgfile = in["bgfile"].(string)
		ret.Bgstyle = in["bgstyle"].(string)
		ret.BgcolorName = in["bgcolor_name"].(string)
		ret.BgcolorValue = in["bgcolor_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonFailMsgCfg(d []interface{}) edpt.AamAuthenticationPortalLogonFailMsgCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailMsgCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FailMsg = in["fail_msg"].(int)
		ret.FailText = in["fail_text"].(string)
		ret.FailFont = in["fail_font"].(int)
		ret.FailFace = in["fail_face"].(string)
		ret.FailFontCustom = in["fail_font_custom"].(string)
		ret.FailSize = in["fail_size"].(int)
		ret.FailColor = in["fail_color"].(int)
		ret.FailColorName = in["fail_color_name"].(string)
		ret.FailColorValue = in["fail_color_value"].(string)
		ret.AuthzFailMsg = in["authz_fail_msg"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonPasscodeCfg(d []interface{}) edpt.AamAuthenticationPortalLogonPasscodeCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonPasscodeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Passcode = in["passcode"].(int)
		ret.PasscodeText = in["passcode_text"].(string)
		ret.PasscodeFont = in["passcode_font"].(int)
		ret.PasscodeFace = in["passcode_face"].(string)
		ret.PasscodeFontCustom = in["passcode_font_custom"].(string)
		ret.PasscodeSize = in["passcode_size"].(int)
		ret.PasscodeColor = in["passcode_color"].(int)
		ret.PasscodeColorName = in["passcode_color_name"].(string)
		ret.PasscodeColorValue = in["passcode_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonPasswordCfg(d []interface{}) edpt.AamAuthenticationPortalLogonPasswordCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonPasswordCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Password = in["password"].(int)
		ret.PassText = in["pass_text"].(string)
		ret.PassFont = in["pass_font"].(int)
		ret.PassFace = in["pass_face"].(string)
		ret.PassFontCustom = in["pass_font_custom"].(string)
		ret.PassSize = in["pass_size"].(int)
		ret.PassColor = in["pass_color"].(int)
		ret.PassColorName = in["pass_color_name"].(string)
		ret.PassColorValue = in["pass_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonRecaptchaCfg(d []interface{}) edpt.AamAuthenticationPortalLogonRecaptchaCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonRecaptchaCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RecaptchaTheme = in["recaptcha_theme"].(string)
		ret.RecaptchaSize = in["recaptcha_size"].(string)
		ret.RecaptchaBadge = in["recaptcha_badge"].(string)
		ret.RecaptchaAction = in["recaptcha_action"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonUsernameCfg(d []interface{}) edpt.AamAuthenticationPortalLogonUsernameCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonUsernameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Username = in["username"].(int)
		ret.UserText = in["user_text"].(string)
		ret.UserFont = in["user_font"].(int)
		ret.UserFace = in["user_face"].(string)
		ret.UserFontCustom = in["user_font_custom"].(string)
		ret.UserSize = in["user_size"].(int)
		ret.UserColor = in["user_color"].(int)
		ret.UserColorName = in["user_color_name"].(string)
		ret.UserColorValue = in["user_color_value"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationPortalLogon(d *schema.ResourceData) edpt.AamAuthenticationPortalLogon {
	var ret edpt.AamAuthenticationPortalLogon
	ret.Inst.ActionUrl = d.Get("action_url").(string)
	ret.Inst.Background = getObjectAamAuthenticationPortalLogonBackground(d.Get("background").([]interface{}))
	ret.Inst.CaptchaType = d.Get("captcha_type").(string)
	ret.Inst.EnableCaptcha = d.Get("enable_captcha").(int)
	ret.Inst.EnablePasscode = d.Get("enable_passcode").(int)
	//omit encrypted
	ret.Inst.FailMsgCfg = getObjectAamAuthenticationPortalLogonFailMsgCfg(d.Get("fail_msg_cfg").([]interface{}))
	ret.Inst.PasscodeCfg = getObjectAamAuthenticationPortalLogonPasscodeCfg(d.Get("passcode_cfg").([]interface{}))
	ret.Inst.PasscodeVar = d.Get("passcode_var").(string)
	ret.Inst.PasswordCfg = getObjectAamAuthenticationPortalLogonPasswordCfg(d.Get("password_cfg").([]interface{}))
	ret.Inst.PasswordVar = d.Get("password_var").(string)
	ret.Inst.RecaptchaCfg = getObjectAamAuthenticationPortalLogonRecaptchaCfg(d.Get("recaptcha_cfg").([]interface{}))
	ret.Inst.SiteKeyString = d.Get("site_key_string").(string)
	ret.Inst.SubmitText = d.Get("submit_text").(string)
	ret.Inst.UsernameCfg = getObjectAamAuthenticationPortalLogonUsernameCfg(d.Get("username_cfg").([]interface{}))
	ret.Inst.UsernameVar = d.Get("username_var").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
