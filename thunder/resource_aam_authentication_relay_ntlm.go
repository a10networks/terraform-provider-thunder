package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayNtlm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_ntlm`: NTLM Authentication Relay\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayNtlmCreate,
		UpdateContext: resourceAamAuthenticationRelayNtlmUpdate,
		ReadContext:   resourceAamAuthenticationRelayNtlmRead,
		DeleteContext: resourceAamAuthenticationRelayNtlmDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Specify NTLM domain, default is null",
			},
			"large_request_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable NTLM relay processing for large requests",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify NTLM authentication relay name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'success': Success; 'failure': Failure; 'request': Request; 'response': Response; 'http-code-200': HTTP 200 OK; 'http-code-400': HTTP 400 Bad Request; 'http-code-401': HTTP 401 Unauthorized; 'http-code-403': HTTP 403 Forbidden; 'http-code-404': HTTP 404 Not Found; 'http-code-500': HTTP 500 Internal Server Error; 'http-code-503': HTTP 503 Service Unavailable; 'http-code-other': Other HTTP Response; 'buffer-alloc-fail': Buffer Allocation Failure; 'encoding-fail': Encoding Failure; 'insert-header-fail': Insert Header Failure; 'parse-header-fail': Parse Header Failure; 'internal-error': Internal Error; 'ntlm-auth-skipped': Requests for which NTLM relay is skipped; 'large-request-processing': Requests invoking large request processing; 'large-request-flushed': Large requests sent to server; 'head-negotiate-request-sent': HEAD requests sent with NEGOTIATE header; 'head-auth-request-sent': HEAD requests sent with AUTH header;",
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
			"version": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Specify NTLM version, default is NTLM 2",
			},
		},
	}
}
func resourceAamAuthenticationRelayNtlmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayNtlmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayNtlm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayNtlmRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayNtlmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayNtlmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayNtlm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayNtlmRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayNtlmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayNtlmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayNtlm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayNtlmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayNtlmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayNtlm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelayNtlmSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayNtlmSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayNtlmSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayNtlmSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayNtlm(d *schema.ResourceData) edpt.AamAuthenticationRelayNtlm {
	var ret edpt.AamAuthenticationRelayNtlm
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.LargeRequestDisable = d.Get("large_request_disable").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelayNtlmSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Version = d.Get("version").(int)
	return ret
}
