package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationOauthGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_oauth_global`: Global Oauth statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationOauthGlobalCreate,
		UpdateContext: resourceAamAuthenticationOauthGlobalUpdate,
		ReadContext:   resourceAamAuthenticationOauthGlobalRead,
		DeleteContext: resourceAamAuthenticationOauthGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'auth-req': auth-req; 'auth-succ': auth-succ; 'auth-fail': auth-fail; 'auth-error': auth-error; 'relay-req': relay-req; 'relay-succ': relay-succ; 'relay-fail': relay-fail; 'other-error': other-error;",
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
func resourceAamAuthenticationOauthGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationOauthGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationOauthGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationOauthGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationOauthGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationOauthGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationOauthGlobalSamplingEnable(d []interface{}) []edpt.AamAuthenticationOauthGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationOauthGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationOauthGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationOauthGlobal(d *schema.ResourceData) edpt.AamAuthenticationOauthGlobal {
	var ret edpt.AamAuthenticationOauthGlobal
	ret.Inst.SamplingEnable = getSliceAamAuthenticationOauthGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
