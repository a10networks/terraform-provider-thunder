package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationPortalNotifyChangePassword() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_portal_notify_change_password`: Change password notification page configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationPortalNotifyChangePasswordCreate,
		UpdateContext: resourceAamAuthenticationPortalNotifyChangePasswordUpdate,
		ReadContext:   resourceAamAuthenticationPortalNotifyChangePasswordRead,
		DeleteContext: resourceAamAuthenticationPortalNotifyChangePasswordDelete,

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
			"change_text": {
				Type: schema.TypeString, Optional: true, Description: "Specify change button text in default change password notification page (Default: Change)",
			},
			"change_url": {
				Type: schema.TypeString, Optional: true, Description: "Specify change password action URL in default change password notification page (Default: /notify_change.fo)",
			},
			"confirm_password_var": {
				Type: schema.TypeString, Optional: true, Description: "Specify confirm password variable name in default change password notification page (Default: cp_cfm_pwd)",
			},
			"continue_text": {
				Type: schema.TypeString, Optional: true, Description: "Specify continue button text in default change password notification page (Default: Continue)",
			},
			"continue_url": {
				Type: schema.TypeString, Optional: true, Description: "Specify continue action URL in default change password notification page (Default: /continue.fo)",
			},
			"new_password_var": {
				Type: schema.TypeString, Optional: true, Description: "Specify new password variable name in default change password notification page (Default: cp_new_pwd)",
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
			"old_password_var": {
				Type: schema.TypeString, Optional: true, Description: "Specify old password variable name in default change password notification page (Default: cp_old_pwd)",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceAamAuthenticationPortalNotifyChangePasswordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalNotifyChangePasswordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalNotifyChangePassword(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalNotifyChangePasswordRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationPortalNotifyChangePasswordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalNotifyChangePasswordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalNotifyChangePassword(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalNotifyChangePasswordRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationPortalNotifyChangePasswordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalNotifyChangePasswordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalNotifyChangePassword(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationPortalNotifyChangePasswordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalNotifyChangePasswordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalNotifyChangePassword(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationPortalNotifyChangePasswordBackground(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordBackground {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordBackground
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgfile = in["bgfile"].(string)
		ret.Bgstyle = in["bgstyle"].(string)
		ret.BgcolorName = in["bgcolor_name"].(string)
		ret.BgcolorValue = in["bgcolor_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalNotifyChangePasswordCfmPwdCfg(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordCfmPwdCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordCfmPwdCfg
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

func getObjectAamAuthenticationPortalNotifyChangePasswordNewPwdCfg(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordNewPwdCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordNewPwdCfg
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

func getObjectAamAuthenticationPortalNotifyChangePasswordOldPwdCfg(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordOldPwdCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordOldPwdCfg
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

func getObjectAamAuthenticationPortalNotifyChangePasswordUsernameCfg(d []interface{}) edpt.AamAuthenticationPortalNotifyChangePasswordUsernameCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalNotifyChangePasswordUsernameCfg
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

func dataToEndpointAamAuthenticationPortalNotifyChangePassword(d *schema.ResourceData) edpt.AamAuthenticationPortalNotifyChangePassword {
	var ret edpt.AamAuthenticationPortalNotifyChangePassword
	ret.Inst.Background = getObjectAamAuthenticationPortalNotifyChangePasswordBackground(d.Get("background").([]interface{}))
	ret.Inst.CfmPwdCfg = getObjectAamAuthenticationPortalNotifyChangePasswordCfmPwdCfg(d.Get("cfm_pwd_cfg").([]interface{}))
	ret.Inst.ChangeText = d.Get("change_text").(string)
	ret.Inst.ChangeUrl = d.Get("change_url").(string)
	ret.Inst.ConfirmPasswordVar = d.Get("confirm_password_var").(string)
	ret.Inst.ContinueText = d.Get("continue_text").(string)
	ret.Inst.ContinueUrl = d.Get("continue_url").(string)
	ret.Inst.NewPasswordVar = d.Get("new_password_var").(string)
	ret.Inst.NewPwdCfg = getObjectAamAuthenticationPortalNotifyChangePasswordNewPwdCfg(d.Get("new_pwd_cfg").([]interface{}))
	ret.Inst.OldPasswordVar = d.Get("old_password_var").(string)
	ret.Inst.OldPwdCfg = getObjectAamAuthenticationPortalNotifyChangePasswordOldPwdCfg(d.Get("old_pwd_cfg").([]interface{}))
	ret.Inst.UsernameCfg = getObjectAamAuthenticationPortalNotifyChangePasswordUsernameCfg(d.Get("username_cfg").([]interface{}))
	ret.Inst.UsernameVar = d.Get("username_var").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
