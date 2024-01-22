package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAuthentication() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_authentication`: Configure authentication feature\n\n__PLACEHOLDER__",
		CreateContext: resourceAuthenticationCreate,
		UpdateContext: resourceAuthenticationUpdate,
		ReadContext:   resourceAuthenticationRead,
		DeleteContext: resourceAuthenticationDelete,

		Schema: map[string]*schema.Schema{
			"console": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "The login authentication type",
									},
									"console_type": {
										Type: schema.TypeString, Optional: true, Description: "",
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
			"enable_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_auth_type": {
							Type: schema.TypeString, Optional: true, Default: "local", Description: "The enable-password authentication type",
						},
					},
				},
			},
			"login_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"privilege_mode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure to enter privilege-mode",
						},
						"local": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure local user to enter privilege-mode",
						},
					},
				},
			},
			"mode_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type: schema.TypeString, Optional: true, Default: "single", Description: "'multiple': Multiple authentication mode. If an authentication method rejected, try next one; 'single': Single authentication mode. If an authentication method rejected, don't try next one;",
						},
					},
				},
			},
			"multiple_auth_reject": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Multiple same user login reject",
			},
			"type_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Default: "local", Description: "The login authentication type",
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
func resourceAuthenticationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthentication(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuthenticationRead(ctx, d, meta)
	}
	return diags
}

func resourceAuthenticationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthentication(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuthenticationRead(ctx, d, meta)
	}
	return diags
}
func resourceAuthenticationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthentication(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAuthenticationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthentication(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAuthenticationConsole67(d []interface{}) edpt.AuthenticationConsole67 {

	count1 := len(d)
	var ret edpt.AuthenticationConsole67
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TypeCfg = getObjectAuthenticationConsoleTypeCfg68(in["type_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectAuthenticationConsoleTypeCfg68(d []interface{}) edpt.AuthenticationConsoleTypeCfg68 {

	count1 := len(d)
	var ret edpt.AuthenticationConsoleTypeCfg68
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = in["type"].(int)
		ret.ConsoleType = in["console_type"].(string)
	}
	return ret
}

func getObjectAuthenticationEnableCfg(d []interface{}) edpt.AuthenticationEnableCfg {

	count1 := len(d)
	var ret edpt.AuthenticationEnableCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableAuthType = in["enable_auth_type"].(string)
	}
	return ret
}

func getObjectAuthenticationLoginCfg(d []interface{}) edpt.AuthenticationLoginCfg {

	count1 := len(d)
	var ret edpt.AuthenticationLoginCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrivilegeMode = in["privilege_mode"].(int)
		ret.Local = in["local"].(int)
	}
	return ret
}

func getObjectAuthenticationModeCfg(d []interface{}) edpt.AuthenticationModeCfg {

	count1 := len(d)
	var ret edpt.AuthenticationModeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectAuthenticationTypeCfg(d []interface{}) edpt.AuthenticationTypeCfg {

	count1 := len(d)
	var ret edpt.AuthenticationTypeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = in["type"].(string)
	}
	return ret
}

func dataToEndpointAuthentication(d *schema.ResourceData) edpt.Authentication {
	var ret edpt.Authentication
	ret.Inst.Console = getObjectAuthenticationConsole67(d.Get("console").([]interface{}))
	ret.Inst.EnableCfg = getObjectAuthenticationEnableCfg(d.Get("enable_cfg").([]interface{}))
	ret.Inst.LoginCfg = getObjectAuthenticationLoginCfg(d.Get("login_cfg").([]interface{}))
	ret.Inst.ModeCfg = getObjectAuthenticationModeCfg(d.Get("mode_cfg").([]interface{}))
	ret.Inst.MultipleAuthReject = d.Get("multiple_auth_reject").(int)
	ret.Inst.TypeCfg = getObjectAuthenticationTypeCfg(d.Get("type_cfg").([]interface{}))
	//omit uuid
	return ret
}
