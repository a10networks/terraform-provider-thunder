package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateCaptchaInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_captcha_instance`: CAPTCHA template instance\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateCaptchaInstanceCreate,
		UpdateContext: resourceSlbTemplateCaptchaInstanceUpdate,
		ReadContext:   resourceSlbTemplateCaptchaInstanceRead,
		DeleteContext: resourceSlbTemplateCaptchaInstanceDelete,

		Schema: map[string]*schema.Schema{
			"captcha_type": {
				Type: schema.TypeString, Optional: true, Description: "'reCAPTCHAv2-checkbox': Google reCAPTCHAv2 Checkbox; 'reCAPTCHAv2-invisible': Google reCAPTCHAv2 Invisible; 'reCAPTCHAv3': Google reCAPTCHAv3;",
			},
			"client_ip_param_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify client ip parameter name (Set parameter name)",
			},
			"default_captcha": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the default CAPTCHA service",
			},
			"method": {
				Type: schema.TypeString, Optional: true, Default: "POST", Description: "'POST': Uses POST method; 'GET': Uses GET method;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify captcha template name",
			},
			"resp_error_code_field_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify error code field name used in JSON response (Set field name)",
			},
			"resp_result_field_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the result field name used in JSON response (Set field name)",
			},
			"secret_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the secret key",
			},
			"secret_key_param_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify secret key parameter name (Set parameter name)",
			},
			"secret_key_string": {
				Type: schema.TypeString, Optional: true, Description: "Secret key string",
			},
			"send_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send the client IP address if client IP is required in the verification",
			},
			"site_key_string": {
				Type: schema.TypeString, Optional: true, Description: "Site key string",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify the timeout time (Specify timeout value, default is 10 seconds)",
			},
			"token_param_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify CAPTCHA token parameter name (Set parameter name)",
			},
			"url": {
				Type: schema.TypeString, Optional: true, Description: "Specify the server URL for verifying the CAPTCHA token, default scheme is https",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateCaptchaInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptchaInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCaptchaInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateCaptchaInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptchaInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCaptchaInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateCaptchaInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptchaInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateCaptchaInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptchaInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateCaptchaInstance(d *schema.ResourceData) edpt.SlbTemplateCaptchaInstance {
	var ret edpt.SlbTemplateCaptchaInstance
	ret.Inst.CaptchaType = d.Get("captcha_type").(string)
	ret.Inst.ClientIpParamName = d.Get("client_ip_param_name").(string)
	ret.Inst.DefaultCaptcha = d.Get("default_captcha").(int)
	ret.Inst.Method = d.Get("method").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RespErrorCodeFieldName = d.Get("resp_error_code_field_name").(string)
	ret.Inst.RespResultFieldName = d.Get("resp_result_field_name").(string)
	ret.Inst.SecretKey = d.Get("secret_key").(int)
	//omit secret_key_encrypted
	ret.Inst.SecretKeyParamName = d.Get("secret_key_param_name").(string)
	ret.Inst.SecretKeyString = d.Get("secret_key_string").(string)
	ret.Inst.SendClientIp = d.Get("send_client_ip").(int)
	//omit site_key_encrypted
	ret.Inst.SiteKeyString = d.Get("site_key_string").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.TokenParamName = d.Get("token_param_name").(string)
	ret.Inst.Url = d.Get("url").(string)
	//omit uuid
	return ret
}
