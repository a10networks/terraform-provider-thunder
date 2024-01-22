package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationCaptcha() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_captcha`: CAPTCHA Profile\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationCaptchaCreate,
		UpdateContext: resourceAamAuthenticationCaptchaUpdate,
		ReadContext:   resourceAamAuthenticationCaptchaRead,
		DeleteContext: resourceAamAuthenticationCaptchaDelete,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify captcha profile name",
						},
						"secret_key": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify secret key",
						},
						"secret_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Secret key string",
						},
						"url": {
							Type: schema.TypeString, Optional: true, Description: "Specify verify API URL, default scheme is https",
						},
						"method": {
							Type: schema.TypeString, Optional: true, Default: "POST", Description: "'POST': API uses POST method; 'GET': API uses GET method;",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify timeout for verify API response (Specify timeout value, default is 10 seconds)",
						},
						"secret_key_param_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify secret key parameter name used in API (Set parameter name)",
						},
						"token_param_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify token parameter name used in API (Set parameter name)",
						},
						"resp_result_field_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify result field name used in JSON response (Set field name)",
						},
						"resp_error_code_field_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify error code field name used in JSON response (Set field name)",
						},
						"send_client_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client IP address in API",
						},
						"client_ip_param_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify client ip parameter name used in API (Set parameter name)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceAamAuthenticationCaptchaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptcha(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationCaptchaRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationCaptchaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptcha(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationCaptchaRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationCaptchaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptcha(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationCaptchaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptcha(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationCaptchaInstanceList(d []interface{}) []edpt.AamAuthenticationCaptchaInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationCaptchaInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationCaptchaInstanceList
		oi.Name = in["name"].(string)
		oi.SecretKey = in["secret_key"].(int)
		oi.SecretKeyString = in["secret_key_string"].(string)
		//omit encrypted
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

func dataToEndpointAamAuthenticationCaptcha(d *schema.ResourceData) edpt.AamAuthenticationCaptcha {
	var ret edpt.AamAuthenticationCaptcha
	ret.Inst.InstanceList = getSliceAamAuthenticationCaptchaInstanceList(d.Get("instance_list").([]interface{}))
	//omit uuid
	return ret
}
