package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSmtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_smtp`: Configure SMTP Server\n\n__PLACEHOLDER__",
		CreateContext: resourceSmtpCreate,
		UpdateContext: resourceSmtpUpdate,
		ReadContext:   resourceSmtpRead,
		DeleteContext: resourceSmtpDelete,

		Schema: map[string]*schema.Schema{
			"mailfrom": {
				Type: schema.TypeString, Optional: true, Description: "Configure email source address",
			},
			"needauthentication": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure SMTP server need authtication",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Configure SMTP Port (Configure SMTP port, default is 25)",
			},
			"smtp_server": {
				Type: schema.TypeString, Optional: true, Description: "Configure SMTP Server (length:1-254)",
			},
			"smtp_server_v6": {
				Type: schema.TypeString, Optional: true, Description: "Configure SMTP Server IPV6 address",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"username_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type: schema.TypeString, Optional: true, Description: "Configure SMTP login username",
						},
						"password": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"smtp_password": {
										Type: schema.TypeString, Optional: true, Description: "Configure SMTP login password",
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
		},
	}
}
func resourceSmtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSmtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSmtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSmtpRead(ctx, d, meta)
	}
	return diags
}

func resourceSmtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSmtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSmtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSmtpRead(ctx, d, meta)
	}
	return diags
}
func resourceSmtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSmtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSmtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSmtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSmtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSmtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSmtpUsernameCfg(d []interface{}) edpt.SmtpUsernameCfg {

	count1 := len(d)
	var ret edpt.SmtpUsernameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Username = in["username"].(string)
		ret.Password = getObjectSmtpUsernameCfgPassword(in["password"].([]interface{}))
	}
	return ret
}

func getObjectSmtpUsernameCfgPassword(d []interface{}) edpt.SmtpUsernameCfgPassword {

	count1 := len(d)
	var ret edpt.SmtpUsernameCfgPassword
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SmtpPassword = in["smtp_password"].(string)
		//omit encrypted
	}
	return ret
}

func dataToEndpointSmtp(d *schema.ResourceData) edpt.Smtp {
	var ret edpt.Smtp
	ret.Inst.Mailfrom = d.Get("mailfrom").(string)
	ret.Inst.Needauthentication = d.Get("needauthentication").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.SmtpServer = d.Get("smtp_server").(string)
	ret.Inst.SmtpServerV6 = d.Get("smtp_server_v6").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UsernameCfg = getObjectSmtpUsernameCfg(d.Get("username_cfg").([]interface{}))
	//omit uuid
	return ret
}
