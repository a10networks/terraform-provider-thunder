package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_saml_global`: Global SAML statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationSamlGlobalCreate,
		UpdateContext: resourceAamAuthenticationSamlGlobalUpdate,
		ReadContext:   resourceAamAuthenticationSamlGlobalRead,
		DeleteContext: resourceAamAuthenticationSamlGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'requests-to-a10saml': Total Request to A10 SAML Service; 'responses-from-a10saml': Total Response from A10 SAML Service; 'sp-metadata-export-req': Total Metadata Export Request; 'sp-metadata-export-success': Toal Metadata Export Success; 'login-auth-req': Total Login Authentication Request; 'login-auth-resp': Total Login Authentication Response; 'acs-req': Total SAML Single-Sign-On Request; 'acs-success': Total SAML Single-Sign-On Success; 'acs-authz-fail': Total SAML Single-Sign-On Authorization Fail; 'acs-error': Total SAML Single-Sign-On Error; 'slo-req': Total Single Logout Request; 'slo-success': Total Single Logout Success; 'slo-error': Total Single Logout Error; 'sp-slo-req': Total SP-initiated Single Logout Request; 'glo-slo-success': Total Global Logout Success; 'loc-slo-success': Total Local Logout Success; 'par-slo-success': Total Partial Logout Success; 'relay-req': relay-req; 'relay-success': relay-success; 'relay-fail': relay-fail; 'relay-error': relay-error; 'other-error': Total Other Error;",
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
func resourceAamAuthenticationSamlGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationSamlGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationSamlGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationSamlGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationSamlGlobalSamplingEnable(d []interface{}) []edpt.AamAuthenticationSamlGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlGlobal(d *schema.ResourceData) edpt.AamAuthenticationSamlGlobal {
	var ret edpt.AamAuthenticationSamlGlobal
	ret.Inst.SamplingEnable = getSliceAamAuthenticationSamlGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
