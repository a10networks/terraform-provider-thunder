package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNotificationTemplateApiAuthentication() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_notification_template_api_authentication`: Authernication URI for the host\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNotificationTemplateApiAuthenticationCreate,
		UpdateContext: resourceDdosNotificationTemplateApiAuthenticationUpdate,
		ReadContext:   resourceDdosNotificationTemplateApiAuthenticationRead,
		DeleteContext: resourceDdosNotificationTemplateApiAuthenticationDelete,

		Schema: map[string]*schema.Schema{
			"api_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure api-key as a mode of authentication",
			},
			"api_key_string": {
				Type: schema.TypeString, Optional: true, Description: "Configure api-key as a mode of authentication",
			},
			"auth_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure the authentication user password (Authentication password)",
			},
			"auth_password_val": {
				Type: schema.TypeString, Optional: true, Description: "Configure the authentication user password (Authentication password)",
			},
			"auth_username": {
				Type: schema.TypeString, Optional: true, Description: "Configure the authentication user name",
			},
			"relative_login_uri": {
				Type: schema.TypeString, Optional: true, Description: "Configure the authentication login uri",
			},
			"relative_logoff_uri": {
				Type: schema.TypeString, Optional: true, Description: "Configure the authentication logoff uri",
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
func resourceDdosNotificationTemplateApiAuthenticationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiAuthenticationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApiAuthentication(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateApiAuthenticationRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNotificationTemplateApiAuthenticationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiAuthenticationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApiAuthentication(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateApiAuthenticationRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNotificationTemplateApiAuthenticationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiAuthenticationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApiAuthentication(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNotificationTemplateApiAuthenticationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiAuthenticationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApiAuthentication(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNotificationTemplateApiAuthentication(d *schema.ResourceData) edpt.DdosNotificationTemplateApiAuthentication {
	var ret edpt.DdosNotificationTemplateApiAuthentication
	ret.Inst.ApiKey = d.Get("api_key").(int)
	//omit api_key_encrypted
	ret.Inst.ApiKeyString = d.Get("api_key_string").(string)
	ret.Inst.AuthPassword = d.Get("auth_password").(int)
	ret.Inst.AuthPasswordVal = d.Get("auth_password_val").(string)
	ret.Inst.AuthUsername = d.Get("auth_username").(string)
	//omit encrypted
	ret.Inst.RelativeLoginUri = d.Get("relative_login_uri").(string)
	ret.Inst.RelativeLogoffUri = d.Get("relative_logoff_uri").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
