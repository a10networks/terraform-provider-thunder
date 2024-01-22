package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelaySaml() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_saml`: SAML Authentication Relay\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelaySamlCreate,
		UpdateContext: resourceAamAuthenticationRelaySamlUpdate,
		ReadContext:   resourceAamAuthenticationRelaySamlRead,
		DeleteContext: resourceAamAuthenticationRelaySamlDelete,

		Schema: map[string]*schema.Schema{
			"idp_auth_uri": {
				Type: schema.TypeString, Optional: true, Description: "Specify the URI for IDP to handle SAML authentication request",
			},
			"match_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': URI exactly matches the string; 'contains': URI string contains another sub string; 'starts-with': URI string starts with sub string; 'ends-with': URI string ends with sub string;",
			},
			"match_uri": {
				Type: schema.TypeString, Optional: true, Description: "Match URI",
			},
			"method": {
				Type: schema.TypeString, Optional: true, Description: "'get-from-backend': Get RelayState parameter from backend server; 'request-uri': Use the (URL encoded) current request-uri as the RelayState;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify SAML authentication relay name",
			},
			"relay_acs_uri": {
				Type: schema.TypeString, Optional: true, Description: "Specify the backend server assertion consuming service URI",
			},
			"retry_number": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify how many continuous fail for SAML relay will trigger. Default will not retry.",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'request': Request; 'success': Success; 'failure': Failure; 'error': Error;",
						},
					},
				},
			},
			"server_cookie_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the cookie name that used by backend server for authenticated users",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeString, Optional: true, Description: "Use the fixed string as the RelayState",
			},
		},
	}
}
func resourceAamAuthenticationRelaySamlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelaySamlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelaySaml(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelaySamlRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelaySamlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelaySamlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelaySaml(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelaySamlRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelaySamlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelaySamlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelaySaml(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelaySamlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelaySamlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelaySaml(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelaySamlSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelaySamlSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelaySamlSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelaySamlSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelaySaml(d *schema.ResourceData) edpt.AamAuthenticationRelaySaml {
	var ret edpt.AamAuthenticationRelaySaml
	ret.Inst.IdpAuthUri = d.Get("idp_auth_uri").(string)
	ret.Inst.MatchType = d.Get("match_type").(string)
	ret.Inst.MatchUri = d.Get("match_uri").(string)
	ret.Inst.Method = d.Get("method").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RelayAcsUri = d.Get("relay_acs_uri").(string)
	ret.Inst.RetryNumber = d.Get("retry_number").(int)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelaySamlSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerCookieName = d.Get("server_cookie_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	return ret
}
