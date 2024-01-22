package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayWsFederation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_ws_federation`: WS-Federation Authentication Relay\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayWsFederationCreate,
		UpdateContext: resourceAamAuthenticationRelayWsFederationUpdate,
		ReadContext:   resourceAamAuthenticationRelayWsFederationRead,
		DeleteContext: resourceAamAuthenticationRelayWsFederationDelete,

		Schema: map[string]*schema.Schema{
			"application_server": {
				Type: schema.TypeString, Optional: true, Description: "'sharepoint': Microsoft SharePoint; 'exchange-owa': Microsoft Exchange OWA;",
			},
			"authentication_uri": {
				Type: schema.TypeString, Optional: true, Description: "Specify WS-Federation relay URI, default is /_trust/",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify WS-Federation authentication relay name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'request': Request; 'success': Success; 'failure': Failure;",
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
func resourceAamAuthenticationRelayWsFederationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayWsFederationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayWsFederation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayWsFederationRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayWsFederationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayWsFederationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayWsFederation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayWsFederationRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayWsFederationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayWsFederationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayWsFederation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayWsFederationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayWsFederationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayWsFederation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelayWsFederationSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayWsFederationSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayWsFederationSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayWsFederationSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayWsFederation(d *schema.ResourceData) edpt.AamAuthenticationRelayWsFederation {
	var ret edpt.AamAuthenticationRelayWsFederation
	ret.Inst.ApplicationServer = d.Get("application_server").(string)
	ret.Inst.AuthenticationUri = d.Get("authentication_uri").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelayWsFederationSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
