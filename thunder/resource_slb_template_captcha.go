package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateCaptcha() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_captcha`: CAPTCHA template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateCaptchaCreate,
		UpdateContext: resourceSlbTemplateCaptchaUpdate,
		ReadContext:   resourceSlbTemplateCaptchaRead,
		DeleteContext: resourceSlbTemplateCaptchaDelete,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify captcha template name",
						},
						"default_captcha": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the default CAPTCHA service",
						},
						"captcha_type": {
							Type: schema.TypeString, Optional: true, Description: "'reCAPTCHAv2-checkbox': Google reCAPTCHAv2 Checkbox; 'reCAPTCHAv2-invisible': Google reCAPTCHAv2 Invisible; 'reCAPTCHAv3': Google reCAPTCHAv3;",
						},
						"site_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Site key string",
						},
						"secret_key": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the secret key",
						},
						"secret_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Secret key string",
						},
						"url": {
							Type: schema.TypeString, Optional: true, Description: "Specify the server URL for verifying the CAPTCHA token, default scheme is https",
						},
						"method": {
							Type: schema.TypeString, Optional: true, Default: "POST", Description: "'POST': Uses POST method; 'GET': Uses GET method;",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify the timeout time (Specify timeout value, default is 10 seconds)",
						},
						"secret_key_param_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify secret key parameter name (Set parameter name)",
						},
						"token_param_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify CAPTCHA token parameter name (Set parameter name)",
						},
						"resp_result_field_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the result field name used in JSON response (Set field name)",
						},
						"resp_error_code_field_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify error code field name used in JSON response (Set field name)",
						},
						"send_client_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send the client IP address if client IP is required in the verification",
						},
						"client_ip_param_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify client ip parameter name (Set parameter name)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceSlbTemplateCaptchaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptcha(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCaptchaRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateCaptchaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptcha(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCaptchaRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateCaptchaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptcha(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateCaptchaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptcha(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateCaptchaInstanceList(d []interface{}) []edpt.SlbTemplateCaptchaInstanceList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateCaptchaInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateCaptchaInstanceList
		oi.Name = in["name"].(string)
		oi.DefaultCaptcha = in["default_captcha"].(int)
		oi.CaptchaType = in["captcha_type"].(string)
		oi.SiteKeyString = in["site_key_string"].(string)
		//omit site_key_encrypted
		oi.SecretKey = in["secret_key"].(int)
		oi.SecretKeyString = in["secret_key_string"].(string)
		//omit secret_key_encrypted
		oi.Url = in["url"].(string)
		oi.Method = in["method"].(string)
		oi.Timeout = in["timeout"].(int)
		oi.SecretKeyParamName = in["secret_key_param_name"].(string)
		oi.TokenParamName = in["token_param_name"].(string)
		oi.RespResultFieldName = in["resp_result_field_name"].(string)
		oi.RespErrorCodeFieldName = in["resp_error_code_field_name"].(string)
		oi.SendClientIp = in["send_client_ip"].(int)
		oi.ClientIpParamName = in["client_ip_param_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateCaptcha(d *schema.ResourceData) edpt.SlbTemplateCaptcha {
	var ret edpt.SlbTemplateCaptcha
	ret.Inst.InstanceList = getSliceSlbTemplateCaptchaInstanceList(d.Get("instance_list").([]interface{}))
	return ret
}
