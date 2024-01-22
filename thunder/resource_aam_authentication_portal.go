package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationPortal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_portal`: Authentication portal configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationPortalCreate,
		UpdateContext: resourceAamAuthenticationPortalUpdate,
		ReadContext:   resourceAamAuthenticationPortalRead,
		DeleteContext: resourceAamAuthenticationPortalDelete,

		Schema: map[string]*schema.Schema{
			"change_password": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"title_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure title in default change password page",
									},
									"title_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify title (Default: Please Change Your Password)",
									},
									"title_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"title_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"title_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"title_size": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify font size (Default: 5)",
									},
									"title_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"title_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"title_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"action_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify form action URL in default change password page (Default: /change.fo)",
						},
						"username_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure username text in default change password page",
									},
									"user_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify username text (Default: Username)",
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
							Type: schema.TypeString, Optional: true, Description: "Specify username variable name in default change password page (Default: cp_usr)",
						},
						"old_pwd_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"old_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure old password text in default change password page",
									},
									"old_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify old password text (Default: Old Password)",
									},
									"old_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"old_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"old_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"old_size": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
									},
									"old_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"old_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"old_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"old_password_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify old password variable name in default change password page (Default: cp_old_pwd)",
						},
						"new_pwd_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"new_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure new password text in default change password page",
									},
									"new_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify new password text (Default: New Password)",
									},
									"new_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"new_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"new_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"new_size": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
									},
									"new_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"new_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"new_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"new_password_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify new password variable name in default change password page (Default: cp_new_pwd)",
						},
						"cfm_pwd_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"confirm_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure confirm password text in default change password page",
									},
									"cfm_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify confirm password text (Default: Confirm New Password)",
									},
									"cfm_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"cfm_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"cfm_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"cfm_size": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
									},
									"cfm_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"cfm_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"cfm_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"confirm_password_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify confirm password variable name in default change password page (Default: cp_cfm_pwd)",
						},
						"submit_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify submit button text in default change password page (Default: Submit)",
						},
						"reset_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify reset button text in default change password page (Default: Reset)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"logo_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logo": {
							Type: schema.TypeString, Optional: true, Description: "Specify logo image filename",
						},
						"width": {
							Type: schema.TypeInt, Optional: true, Default: 134, Description: "Specify logo image width (Default: 134)",
						},
						"height": {
							Type: schema.TypeInt, Optional: true, Default: 71, Description: "Specify logo image height (Default: 71)",
						},
					},
				},
			},
			"logon": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"action_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify form action URL in default logon page (Default: /logon.fo)",
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
						"enable_passcode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable passcode field in default logon page",
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
						"enable_captcha": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CAPTCHA in deafult logon page",
						},
						"captcha_type": {
							Type: schema.TypeString, Optional: true, Description: "'reCAPTCHAv2-checkbox': Google reCAPTCHAv2 Checkbox; 'reCAPTCHAv2-invisible': Google reCAPTCHAv2 Invisible; 'reCAPTCHAv3': Google reCAPTCHAv3;",
						},
						"site_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Site key string",
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
						"submit_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify submit button text in default logon page (Default: Log In)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"logon_fail": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"title_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure title in default logon fail page",
									},
									"title_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify title (Default: Try Too Many Times)",
									},
									"title_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"title_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"title_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"title_size": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify font size (Default: 5)",
									},
									"title_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"title_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"title_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"fail_msg_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fail_msg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure logon failure message in default logon fail page",
									},
									"fail_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify logon failure message (Default: Login Failed!!)",
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
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
									},
									"fail_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"fail_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"fail_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "'default-portal': Default portal configuration;",
			},
			"notify_change_password": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"continue_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify continue action URL in default change password notification page (Default: /continue.fo)",
						},
						"change_url": {
							Type: schema.TypeString, Optional: true, Description: "Specify change password action URL in default change password notification page (Default: /notify_change.fo)",
						},
						"username_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure username text in default change password notification page",
									},
									"user_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify username text (Default: Username)",
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
							Type: schema.TypeString, Optional: true, Description: "Specify username variable name in default change password notification page (Default: cp_usr)",
						},
						"old_pwd_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"old_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure old password text in default change password notification page",
									},
									"old_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify old password text (Default: Old Password)",
									},
									"old_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"old_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"old_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"old_size": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
									},
									"old_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"old_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"old_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"old_password_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify old password variable name in default change password notification page (Default: cp_old_pwd)",
						},
						"new_pwd_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"new_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure new password text in default change password notification page",
									},
									"new_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify new password text (Default: New Password)",
									},
									"new_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"new_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"new_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"new_size": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
									},
									"new_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"new_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"new_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"new_password_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify new password variable name in default change password notification page (Default: cp_new_pwd)",
						},
						"cfm_pwd_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"confirm_password": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure confirm password text in default change password notification page",
									},
									"cfm_text": {
										Type: schema.TypeString, Optional: true, Description: "Specify confirm password text (Default: Confirm New Password)",
									},
									"cfm_font": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Sepcify font (Default: Arial)",
									},
									"cfm_face": {
										Type: schema.TypeString, Optional: true, Default: "Arial", Description: "'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;",
									},
									"cfm_font_custom": {
										Type: schema.TypeString, Optional: true, Description: "Specify custom font",
									},
									"cfm_size": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify font size (Default: 3)",
									},
									"cfm_color": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify font color (Default: black)",
									},
									"cfm_color_name": {
										Type: schema.TypeString, Optional: true, Default: "black", Description: "'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;",
									},
									"cfm_color_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify 6-digit HEX color value",
									},
								},
							},
						},
						"confirm_password_var": {
							Type: schema.TypeString, Optional: true, Description: "Specify confirm password variable name in default change password notification page (Default: cp_cfm_pwd)",
						},
						"change_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify change button text in default change password notification page (Default: Change)",
						},
						"continue_text": {
							Type: schema.TypeString, Optional: true, Description: "Specify continue button text in default change password notification page (Default: Continue)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
func resourceAamAuthenticationPortalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationPortalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationPortalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationPortalChangePassword0(d []interface{}) edpt.AamAuthenticationPortalChangePassword0 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalChangePassword0
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Background = getObjectAamAuthenticationPortalChangePasswordBackground1(in["background"].([]interface{}))
		ret.TitleCfg = getObjectAamAuthenticationPortalChangePasswordTitleCfg2(in["title_cfg"].([]interface{}))
		ret.ActionUrl = in["action_url"].(string)
		ret.UsernameCfg = getObjectAamAuthenticationPortalChangePasswordUsernameCfg3(in["username_cfg"].([]interface{}))
		ret.UsernameVar = in["username_var"].(string)
		ret.OldPwdCfg = getObjectAamAuthenticationPortalChangePasswordOldPwdCfg4(in["old_pwd_cfg"].([]interface{}))
		ret.OldPasswordVar = in["old_password_var"].(string)
		ret.NewPwdCfg = getObjectAamAuthenticationPortalChangePasswordNewPwdCfg5(in["new_pwd_cfg"].([]interface{}))
		ret.NewPasswordVar = in["new_password_var"].(string)
		ret.CfmPwdCfg = getObjectAamAuthenticationPortalChangePasswordCfmPwdCfg6(in["cfm_pwd_cfg"].([]interface{}))
		ret.ConfirmPasswordVar = in["confirm_password_var"].(string)
		ret.SubmitText = in["submit_text"].(string)
		ret.ResetText = in["reset_text"].(string)
		//omit uuid
	}
	return ret
}

func getObjectAamAuthenticationPortalChangePasswordBackground1(d []interface{}) edpt.AamAuthenticationPortalChangePasswordBackground1 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalChangePasswordBackground1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgfile = in["bgfile"].(string)
		ret.Bgstyle = in["bgstyle"].(string)
		ret.BgcolorName = in["bgcolor_name"].(string)
		ret.BgcolorValue = in["bgcolor_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalChangePasswordTitleCfg2(d []interface{}) edpt.AamAuthenticationPortalChangePasswordTitleCfg2 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalChangePasswordTitleCfg2
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Title = in["title"].(int)
		ret.TitleText = in["title_text"].(string)
		ret.TitleFont = in["title_font"].(int)
		ret.TitleFace = in["title_face"].(string)
		ret.TitleFontCustom = in["title_font_custom"].(string)
		ret.TitleSize = in["title_size"].(int)
		ret.TitleColor = in["title_color"].(int)
		ret.TitleColorName = in["title_color_name"].(string)
		ret.TitleColorValue = in["title_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalChangePasswordUsernameCfg3(d []interface{}) edpt.AamAuthenticationPortalChangePasswordUsernameCfg3 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalChangePasswordUsernameCfg3
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

func getObjectAamAuthenticationPortalChangePasswordOldPwdCfg4(d []interface{}) edpt.AamAuthenticationPortalChangePasswordOldPwdCfg4 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalChangePasswordOldPwdCfg4
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OldPassword = in["old_password"].(int)
		ret.OldText = in["old_text"].(string)
		ret.OldFont = in["old_font"].(int)
		ret.OldFace = in["old_face"].(string)
		ret.OldFontCustom = in["old_font_custom"].(string)
		ret.OldSize = in["old_size"].(int)
		ret.OldColor = in["old_color"].(int)
		ret.OldColorName = in["old_color_name"].(string)
		ret.OldColorValue = in["old_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalChangePasswordNewPwdCfg5(d []interface{}) edpt.AamAuthenticationPortalChangePasswordNewPwdCfg5 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalChangePasswordNewPwdCfg5
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NewPassword = in["new_password"].(int)
		ret.NewText = in["new_text"].(string)
		ret.NewFont = in["new_font"].(int)
		ret.NewFace = in["new_face"].(string)
		ret.NewFontCustom = in["new_font_custom"].(string)
		ret.NewSize = in["new_size"].(int)
		ret.NewColor = in["new_color"].(int)
		ret.NewColorName = in["new_color_name"].(string)
		ret.NewColorValue = in["new_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalChangePasswordCfmPwdCfg6(d []interface{}) edpt.AamAuthenticationPortalChangePasswordCfmPwdCfg6 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalChangePasswordCfmPwdCfg6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConfirmPassword = in["confirm_password"].(int)
		ret.CfmText = in["cfm_text"].(string)
		ret.CfmFont = in["cfm_font"].(int)
		ret.CfmFace = in["cfm_face"].(string)
		ret.CfmFontCustom = in["cfm_font_custom"].(string)
		ret.CfmSize = in["cfm_size"].(int)
		ret.CfmColor = in["cfm_color"].(int)
		ret.CfmColorName = in["cfm_color_name"].(string)
		ret.CfmColorValue = in["cfm_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogoCfg(d []interface{}) edpt.AamAuthenticationPortalLogoCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogoCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logo = in["logo"].(string)
		ret.Width = in["width"].(int)
		ret.Height = in["height"].(int)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogon7(d []interface{}) edpt.AamAuthenticationPortalLogon7 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogon7
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Background = getObjectAamAuthenticationPortalLogonBackground8(in["background"].([]interface{}))
		ret.FailMsgCfg = getObjectAamAuthenticationPortalLogonFailMsgCfg9(in["fail_msg_cfg"].([]interface{}))
		ret.ActionUrl = in["action_url"].(string)
		ret.UsernameCfg = getObjectAamAuthenticationPortalLogonUsernameCfg10(in["username_cfg"].([]interface{}))
		ret.UsernameVar = in["username_var"].(string)
		ret.PasswordCfg = getObjectAamAuthenticationPortalLogonPasswordCfg11(in["password_cfg"].([]interface{}))
		ret.PasswordVar = in["password_var"].(string)
		ret.EnablePasscode = in["enable_passcode"].(int)
		ret.PasscodeCfg = getObjectAamAuthenticationPortalLogonPasscodeCfg12(in["passcode_cfg"].([]interface{}))
		ret.PasscodeVar = in["passcode_var"].(string)
		ret.EnableCaptcha = in["enable_captcha"].(int)
		ret.CaptchaType = in["captcha_type"].(string)
		ret.SiteKeyString = in["site_key_string"].(string)
		//omit encrypted
		ret.RecaptchaCfg = getObjectAamAuthenticationPortalLogonRecaptchaCfg13(in["recaptcha_cfg"].([]interface{}))
		ret.SubmitText = in["submit_text"].(string)
		//omit uuid
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonBackground8(d []interface{}) edpt.AamAuthenticationPortalLogonBackground8 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonBackground8
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgfile = in["bgfile"].(string)
		ret.Bgstyle = in["bgstyle"].(string)
		ret.BgcolorName = in["bgcolor_name"].(string)
		ret.BgcolorValue = in["bgcolor_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonFailMsgCfg9(d []interface{}) edpt.AamAuthenticationPortalLogonFailMsgCfg9 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailMsgCfg9
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

func getObjectAamAuthenticationPortalLogonUsernameCfg10(d []interface{}) edpt.AamAuthenticationPortalLogonUsernameCfg10 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonUsernameCfg10
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

func getObjectAamAuthenticationPortalLogonPasswordCfg11(d []interface{}) edpt.AamAuthenticationPortalLogonPasswordCfg11 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonPasswordCfg11
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

func getObjectAamAuthenticationPortalLogonPasscodeCfg12(d []interface{}) edpt.AamAuthenticationPortalLogonPasscodeCfg12 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonPasscodeCfg12
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

func getObjectAamAuthenticationPortalLogonRecaptchaCfg13(d []interface{}) edpt.AamAuthenticationPortalLogonRecaptchaCfg13 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonRecaptchaCfg13
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RecaptchaTheme = in["recaptcha_theme"].(string)
		ret.RecaptchaSize = in["recaptcha_size"].(string)
		ret.RecaptchaBadge = in["recaptcha_badge"].(string)
		ret.RecaptchaAction = in["recaptcha_action"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonFail14(d []interface{}) edpt.AamAuthenticationPortalLogonFail14 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFail14
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Background = getObjectAamAuthenticationPortalLogonFailBackground15(in["background"].([]interface{}))
		ret.TitleCfg = getObjectAamAuthenticationPortalLogonFailTitleCfg16(in["title_cfg"].([]interface{}))
		ret.FailMsgCfg = getObjectAamAuthenticationPortalLogonFailFailMsgCfg17(in["fail_msg_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonFailBackground15(d []interface{}) edpt.AamAuthenticationPortalLogonFailBackground15 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailBackground15
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgfile = in["bgfile"].(string)
		ret.Bgstyle = in["bgstyle"].(string)
		ret.BgcolorName = in["bgcolor_name"].(string)
		ret.BgcolorValue = in["bgcolor_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonFailTitleCfg16(d []interface{}) edpt.AamAuthenticationPortalLogonFailTitleCfg16 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailTitleCfg16
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Title = in["title"].(int)
		ret.TitleText = in["title_text"].(string)
		ret.TitleFont = in["title_font"].(int)
		ret.TitleFace = in["title_face"].(string)
		ret.TitleFontCustom = in["title_font_custom"].(string)
		ret.TitleSize = in["title_size"].(int)
		ret.TitleColor = in["title_color"].(int)
		ret.TitleColorName = in["title_color_name"].(string)
		ret.TitleColorValue = in["title_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonFailFailMsgCfg17(d []interface{}) edpt.AamAuthenticationPortalLogonFailFailMsgCfg17 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailFailMsgCfg17
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
	}
	return ret
}

func getObjectAamAuthenticationPortalNotifyChangePassword18(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePassword18 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePassword18
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Background = getObjectAamAuthenticationPortalNotifyChangePasswordBackground19(in["background"].([]interface{}))
		ret.ContinueUrl = in["continue_url"].(string)
		ret.ChangeUrl = in["change_url"].(string)
		ret.UsernameCfg = getObjectAamAuthenticationPortalNotifyChangePasswordUsernameCfg20(in["username_cfg"].([]interface{}))
		ret.UsernameVar = in["username_var"].(string)
		ret.OldPwdCfg = getObjectAamAuthenticationPortalNotifyChangePasswordOldPwdCfg21(in["old_pwd_cfg"].([]interface{}))
		ret.OldPasswordVar = in["old_password_var"].(string)
		ret.NewPwdCfg = getObjectAamAuthenticationPortalNotifyChangePasswordNewPwdCfg22(in["new_pwd_cfg"].([]interface{}))
		ret.NewPasswordVar = in["new_password_var"].(string)
		ret.CfmPwdCfg = getObjectAamAuthenticationPortalNotifyChangePasswordCfmPwdCfg23(in["cfm_pwd_cfg"].([]interface{}))
		ret.ConfirmPasswordVar = in["confirm_password_var"].(string)
		ret.ChangeText = in["change_text"].(string)
		ret.ContinueText = in["continue_text"].(string)
		//omit uuid
	}
	return ret
}

func getObjectAamAuthenticationPortalNotifyChangePasswordBackground19(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordBackground19 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordBackground19
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgfile = in["bgfile"].(string)
		ret.Bgstyle = in["bgstyle"].(string)
		ret.BgcolorName = in["bgcolor_name"].(string)
		ret.BgcolorValue = in["bgcolor_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalNotifyChangePasswordUsernameCfg20(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordUsernameCfg20 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordUsernameCfg20
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

func getObjectAamAuthenticationPortalNotifyChangePasswordOldPwdCfg21(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordOldPwdCfg21 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordOldPwdCfg21
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OldPassword = in["old_password"].(int)
		ret.OldText = in["old_text"].(string)
		ret.OldFont = in["old_font"].(int)
		ret.OldFace = in["old_face"].(string)
		ret.OldFontCustom = in["old_font_custom"].(string)
		ret.OldSize = in["old_size"].(int)
		ret.OldColor = in["old_color"].(int)
		ret.OldColorName = in["old_color_name"].(string)
		ret.OldColorValue = in["old_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalNotifyChangePasswordNewPwdCfg22(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordNewPwdCfg22 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordNewPwdCfg22
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NewPassword = in["new_password"].(int)
		ret.NewText = in["new_text"].(string)
		ret.NewFont = in["new_font"].(int)
		ret.NewFace = in["new_face"].(string)
		ret.NewFontCustom = in["new_font_custom"].(string)
		ret.NewSize = in["new_size"].(int)
		ret.NewColor = in["new_color"].(int)
		ret.NewColorName = in["new_color_name"].(string)
		ret.NewColorValue = in["new_color_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalNotifyChangePasswordCfmPwdCfg23(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordCfmPwdCfg23 {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordCfmPwdCfg23
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConfirmPassword = in["confirm_password"].(int)
		ret.CfmText = in["cfm_text"].(string)
		ret.CfmFont = in["cfm_font"].(int)
		ret.CfmFace = in["cfm_face"].(string)
		ret.CfmFontCustom = in["cfm_font_custom"].(string)
		ret.CfmSize = in["cfm_size"].(int)
		ret.CfmColor = in["cfm_color"].(int)
		ret.CfmColorName = in["cfm_color_name"].(string)
		ret.CfmColorValue = in["cfm_color_value"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationPortal(d *schema.ResourceData) edpt.AamAuthenticationPortal {
	var ret edpt.AamAuthenticationPortal
	ret.Inst.ChangePassword = getObjectAamAuthenticationPortalChangePassword0(d.Get("change_password").([]interface{}))
	ret.Inst.LogoCfg = getObjectAamAuthenticationPortalLogoCfg(d.Get("logo_cfg").([]interface{}))
	ret.Inst.Logon = getObjectAamAuthenticationPortalLogon7(d.Get("logon").([]interface{}))
	ret.Inst.LogonFail = getObjectAamAuthenticationPortalLogonFail14(d.Get("logon_fail").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NotifyChangePassword = getObjectAamAuthenticationPortalNotifyChangePassword18(d.Get("notify_change_password").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
