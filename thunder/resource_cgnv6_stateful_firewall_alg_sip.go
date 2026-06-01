package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_alg_sip`: Configure SIP ALG for NAT stateful firewall (default: enabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallAlgSipCreate,
		UpdateContext: resourceCgnv6StatefulFirewallAlgSipUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallAlgSipRead,
		DeleteContext: resourceCgnv6StatefulFirewallAlgSipDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'stat-request': Request Received; 'stat-response': Response Received; 'method-register': Method REGISTER; 'method-invite': Method INVITE; 'method-ack': Method ACK; 'method-cancel': Method CANCEL; 'method-bye': Method BYE; 'method-port-config': Method OPTIONS; 'method-prack': Method PRACK; 'method-subscribe': Method SUBSCRIBE; 'method-notify': Method NOTIFY; 'method-publish': Method PUBLISH; 'method-info': Method INFO; 'method-refer': Method REFER; 'method-message': Method MESSAGE; 'method-update': Method UPDATE; 'method-unknown': Method Unknown; 'parse-error': Message Parse Error; 'keep-alive': Keep Alive; 'contact-error': Contact Process Error; 'sdp-error': SDP Process Error; 'rtp-port-no-op': RTP Port No Op; 'rtp-rtcp-port-success': RTP RTCP Port Success; 'rtp-port-failure': RTP Port Failure; 'rtcp-port-failure': RTCP Port Failure; 'contact-port-no-op': Contact Port No Op; 'contact-port-success': Contact Port Success; 'contact-port-failure': Contact Port Failure; 'contact-new': Contact Alloc; 'contact-alloc-failure': Contact Alloc Failure; 'contact-eim': Contact EIM; 'contact-eim-set': Contact EIM Set; 'rtp-new': RTP Alloc; 'rtp-alloc-failure': RTP Alloc Failure; 'rtp-eim': RTP EIM;",
						},
					},
				},
			},
			"sip_value": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallAlgSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgSipRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgSipRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallAlgSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6StatefulFirewallAlgSipSamplingEnable(d []interface{}) []edpt.Cgnv6StatefulFirewallAlgSipSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6StatefulFirewallAlgSipSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6StatefulFirewallAlgSipSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallAlgSip(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgSip {
	var ret edpt.Cgnv6StatefulFirewallAlgSip
	ret.Inst.SamplingEnable = getSliceCgnv6StatefulFirewallAlgSipSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SipValue = d.Get("sip_value").(string)
	//omit uuid
	return ret
}
