package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationCaptchaInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_captcha_instance`: CAPTCHA Profile\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationCaptchaInstanceCreate,
		UpdateContext: resourceAamAuthenticationCaptchaInstanceUpdate,
		ReadContext:   resourceAamAuthenticationCaptchaInstanceRead,
		DeleteContext: resourceAamAuthenticationCaptchaInstanceDelete,

		Schema: map[string]*schema.Schema{
			"client_ip_param_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify client ip parameter name used in API (Set parameter name)",
			},
			"method": {
				Type: schema.TypeString, Optional: true, Default: "POST", Description: "'POST': API uses POST method; 'GET': API uses GET method;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify captcha profile name",
			},
			"resp_error_code_field_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify error code field name used in JSON response (Set field name)",
			},
			"resp_result_field_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify result field name used in JSON response (Set field name)",
			},
			"secret_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify secret key",
			},
			"secret_key_param_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify secret key parameter name used in API (Set parameter name)",
			},
			"secret_key_string": {
				Type: schema.TypeString, Optional: true, Description: "Secret key string",
			},
			"send_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client IP address in API",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify timeout for verify API response (Specify timeout value, default is 10 seconds)",
			},
			"token_param_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify token parameter name used in API (Set parameter name)",
			},
			"url": {
				Type: schema.TypeString, Optional: true, Description: "Specify verify API URL, default scheme is https",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationCaptchaInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptchaInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationCaptchaInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationCaptchaInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptchaInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationCaptchaInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationCaptchaInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptchaInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationCaptchaInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptchaInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthenticationCaptchaInstance(d *schema.ResourceData) edpt.AamAuthenticationCaptchaInstance {
	var ret edpt.AamAuthenticationCaptchaInstance
	ret.Inst.ClientIpParamName = d.Get("client_ip_param_name").(string)
	//omit encrypted
	ret.Inst.Method = d.Get("method").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RespErrorCodeFieldName = d.Get("resp_error_code_field_name").(string)
	ret.Inst.RespResultFieldName = d.Get("resp_result_field_name").(string)
	ret.Inst.SecretKey = d.Get("secret_key").(int)
	ret.Inst.SecretKeyParamName = d.Get("secret_key_param_name").(string)
	ret.Inst.SecretKeyString = d.Get("secret_key_string").(string)
	ret.Inst.SendClientIp = d.Get("send_client_ip").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.TokenParamName = d.Get("token_param_name").(string)
	ret.Inst.Url = d.Get("url").(string)
	//omit uuid
	return ret
}
