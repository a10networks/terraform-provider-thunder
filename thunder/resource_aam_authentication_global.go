package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_global`: Global AAM authentication statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationGlobalCreate,
		UpdateContext: resourceAamAuthenticationGlobalUpdate,
		ReadContext:   resourceAamAuthenticationGlobalRead,
		DeleteContext: resourceAamAuthenticationGlobalDelete,

		Schema: map[string]*schema.Schema{
			"max_auth_resp_size": {
				Type: schema.TypeInt, Optional: true, Default: 65536, Description: "Specify the max auth resp size in bytes(from authd to a10lb), default is 64KB",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'requests': Total Authentication Request; 'responses': Total Authentication Response; 'misses': Total Authentication Request Missed; 'ocsp-stapling-requests-to-a10authd': Total OCSP Stapling Request; 'ocsp-stapling-responses-from-a10authd': Total OCSP Stapling Response; 'opened-socket': Total AAM Socket Opened; 'open-socket-failed': Total AAM Open Socket Failed; 'connect': Total AAM Connection; 'connect-failed': Total AAM Connect Failed; 'created-timer': Total AAM Timer Created; 'create-timer-failed': Total AAM Timer Creation Failed; 'total-request': Total Request Received by A10 Auth Service; 'get-socket-option-failed': Total AAM Get Socket Option Failed; 'aflex-authz-succ': Total Authorization success number in aFleX; 'aflex-authz-fail': Total Authorization failure number in aFleX; 'authn-success': Total Authentication success number; 'authn-failure': Total Authentication failure number; 'authz-success': Total Authorization success number; 'authz-failure': Total Authorization failure number; 'active-session': Total Active Auth-Sessions; 'active-user': Total Active Users; 'dns-resolve-failed': Total AAM DNS resolve failed; 'domain-wlist-match': Total DOMAIN WHITELIST match number; 'domain-wlist-unmatch': Total DOMAIN WHITELIST unmatch number; 'auth_ctx_num': Total Auth Contexts;",
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
func resourceAamAuthenticationGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationGlobalSamplingEnable(d []interface{}) []edpt.AamAuthenticationGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationGlobal(d *schema.ResourceData) edpt.AamAuthenticationGlobal {
	var ret edpt.AamAuthenticationGlobal
	ret.Inst.MaxAuthRespSize = d.Get("max_auth_resp_size").(int)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
