package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayOauth() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_oauth`: Oauth 2.0 Authentication Relay\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayOauthCreate,
		UpdateContext: resourceAamAuthenticationRelayOauthUpdate,
		ReadContext:   resourceAamAuthenticationRelayOauthRead,
		DeleteContext: resourceAamAuthenticationRelayOauthDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All URI can be relay",
			},
			"match_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': URI exactly matches the string; 'contains': URI string contains another sub string; 'starts-with': URI string starts with sub string; 'ends-with': URI string ends with sub string;",
			},
			"match_uri": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify oauth authentication relay name",
			},
			"relay_type": {
				Type: schema.TypeString, Optional: true, Description: "'access-token': Relay access token to backend; 'id-token': Relay JWT to backend;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'relay-req': relay-req; 'relay-succ': relay-succ; 'relay-fail': relay-fail;",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationRelayOauthCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayOauthCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayOauth(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayOauthRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayOauthUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayOauthUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayOauth(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayOauthRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayOauthDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayOauthDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayOauth(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayOauthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayOauthRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayOauth(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelayOauthSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayOauthSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayOauthSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayOauthSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayOauth(d *schema.ResourceData) edpt.AamAuthenticationRelayOauth {
	var ret edpt.AamAuthenticationRelayOauth
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.MatchType = d.Get("match_type").(string)
	ret.Inst.MatchUri = d.Get("match_uri").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RelayType = d.Get("relay_type").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelayOauthSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
