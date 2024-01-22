package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationPortalLogonFail() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_portal_logon_fail`: Logon fail page configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationPortalLogonFailCreate,
		UpdateContext: resourceAamAuthenticationPortalLogonFailUpdate,
		ReadContext:   resourceAamAuthenticationPortalLogonFailRead,
		DeleteContext: resourceAamAuthenticationPortalLogonFailDelete,

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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceAamAuthenticationPortalLogonFailCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonFailCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogonFail(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalLogonFailRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationPortalLogonFailUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonFailUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogonFail(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationPortalLogonFailRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationPortalLogonFailDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonFailDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogonFail(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationPortalLogonFailRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPortalLogonFailRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPortalLogonFail(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationPortalLogonFailBackground(d []interface{}) edpt.AamAuthenticationPortalLogonFailBackground {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailBackground
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgfile = in["bgfile"].(string)
		ret.Bgstyle = in["bgstyle"].(string)
		ret.BgcolorName = in["bgcolor_name"].(string)
		ret.BgcolorValue = in["bgcolor_value"].(string)
	}
	return ret
}

func getObjectAamAuthenticationPortalLogonFailFailMsgCfg(d []interface{}) edpt.AamAuthenticationPortalLogonFailFailMsgCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailFailMsgCfg
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

func getObjectAamAuthenticationPortalLogonFailTitleCfg(d []interface{}) edpt.AamAuthenticationPortalLogonFailTitleCfg {

	count1 := len(d)
	var ret edpt.AamAuthenticationPortalLogonFailTitleCfg
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

func dataToEndpointAamAuthenticationPortalLogonFail(d *schema.ResourceData) edpt.AamAuthenticationPortalLogonFail {
	var ret edpt.AamAuthenticationPortalLogonFail
	ret.Inst.Background = getObjectAamAuthenticationPortalLogonFailBackground(d.Get("background").([]interface{}))
	ret.Inst.FailMsgCfg = getObjectAamAuthenticationPortalLogonFailFailMsgCfg(d.Get("fail_msg_cfg").([]interface{}))
	ret.Inst.TitleCfg = getObjectAamAuthenticationPortalLogonFailTitleCfg(d.Get("title_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
