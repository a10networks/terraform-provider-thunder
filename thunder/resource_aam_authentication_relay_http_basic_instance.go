package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayHttpBasicInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_http_basic_instance`: HTTP Basic Authentication Relay instance\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayHttpBasicInstanceCreate,
		UpdateContext: resourceAamAuthenticationRelayHttpBasicInstanceUpdate,
		ReadContext:   resourceAamAuthenticationRelayHttpBasicInstanceRead,
		DeleteContext: resourceAamAuthenticationRelayHttpBasicInstanceDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Specify user domain, default is null",
			},
			"domain_format": {
				Type: schema.TypeString, Optional: true, Default: "down-level-logon-name", Description: "'user-principal-name': Append domain with User Principal Name format. (e.g. user@domain); 'down-level-logon-name': Append domain with Down-Level Logon Name format. (e.g. domain-user);",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify HTTP basic authentication relay name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'success': Success; 'no-creds': No Credential; 'bad-req': Bad Request; 'unauth': Unauthorized; 'forbidden': Forbidden; 'not-found': Not Found; 'server-error': Internal Server Error; 'unavailable': Service Unavailable;",
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
func resourceAamAuthenticationRelayHttpBasicInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayHttpBasicInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayHttpBasicInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayHttpBasicInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayHttpBasicInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayHttpBasicInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayHttpBasicInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayHttpBasicInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayHttpBasicInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayHttpBasicInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayHttpBasicInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayHttpBasicInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayHttpBasicInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayHttpBasicInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelayHttpBasicInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayHttpBasicInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayHttpBasicInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayHttpBasicInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayHttpBasicInstance(d *schema.ResourceData) edpt.AamAuthenticationRelayHttpBasicInstance {
	var ret edpt.AamAuthenticationRelayHttpBasicInstance
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.DomainFormat = d.Get("domain_format").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelayHttpBasicInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
