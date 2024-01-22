package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamJwtAuthorization() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_jwt_authorization`: AAM JWT authorization related configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamJwtAuthorizationCreate,
		UpdateContext: resourceAamJwtAuthorizationUpdate,
		ReadContext:   resourceAamJwtAuthorizationRead,
		DeleteContext: resourceAamJwtAuthorizationDelete,

		Schema: map[string]*schema.Schema{
			"exp_claim_requried": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the exp claim is required for JWT authorization",
			},
			"jwt_cache_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable caching authorized JWT token and skip verification and authorization for cached tokens",
			},
			"jwt_exp_default": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the default token expiration if exp claim is not available (default 1800)",
			},
			"jwt_forwarding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify JWT token will not be stripped while forwarding client request",
			},
			"log_level": {
				Type: schema.TypeString, Optional: true, Description: "'0': log disable; '1': only log authorzation fail (default); '2': only log authorization success; '3': log all;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify JWT authorization template name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'jwt-request': JWT Request; 'jwt-authorize-success': JWT Authorize Success; 'jwt-authorize-failure': JWT Authorize Failure; 'jwt-missing-token': JWT Missing Token; 'jwt-missing-claim': JWT Missing Claim; 'jwt-token-expired': JWT Token Expired; 'jwt-signature-failure': JWT Signature Failure; 'jwt-other-error': JWT Other Error;",
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
			"verification_cert": {
				Type: schema.TypeString, Optional: true, Description: "Specify the certificate to verify JWT token signature",
			},
			"verification_jwks": {
				Type: schema.TypeString, Optional: true, Description: "Specify the jwks file to verify JWT token signature",
			},
			"verification_secret": {
				Type: schema.TypeString, Optional: true, Description: "Specify secret for verify JWT token signature",
			},
		},
	}
}
func resourceAamJwtAuthorizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamJwtAuthorizationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamJwtAuthorization(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamJwtAuthorizationRead(ctx, d, meta)
	}
	return diags
}

func resourceAamJwtAuthorizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamJwtAuthorizationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamJwtAuthorization(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamJwtAuthorizationRead(ctx, d, meta)
	}
	return diags
}
func resourceAamJwtAuthorizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamJwtAuthorizationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamJwtAuthorization(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamJwtAuthorizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamJwtAuthorizationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamJwtAuthorization(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamJwtAuthorizationSamplingEnable(d []interface{}) []edpt.AamJwtAuthorizationSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamJwtAuthorizationSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamJwtAuthorizationSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamJwtAuthorization(d *schema.ResourceData) edpt.AamJwtAuthorization {
	var ret edpt.AamJwtAuthorization
	//omit encrypted
	ret.Inst.ExpClaimRequried = d.Get("exp_claim_requried").(int)
	ret.Inst.JwtCacheEnable = d.Get("jwt_cache_enable").(int)
	ret.Inst.JwtExpDefault = d.Get("jwt_exp_default").(int)
	ret.Inst.JwtForwarding = d.Get("jwt_forwarding").(int)
	ret.Inst.LogLevel = d.Get("log_level").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceAamJwtAuthorizationSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VerificationCert = d.Get("verification_cert").(string)
	ret.Inst.VerificationJwks = d.Get("verification_jwks").(string)
	ret.Inst.VerificationSecret = d.Get("verification_secret").(string)
	return ret
}
