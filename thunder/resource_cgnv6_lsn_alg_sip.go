package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_alg_sip`: Change LSN SIP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnAlgSipCreate,
		UpdateContext: resourceCgnv6LsnAlgSipUpdate,
		ReadContext:   resourceCgnv6LsnAlgSipRead,
		DeleteContext: resourceCgnv6LsnAlgSipDelete,

		Schema: map[string]*schema.Schema{
			"rtp_stun_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "RTP/RTCP STUN timeout in minutes (Default is 5 minutes)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'method-register': SIP Method REGISTER; 'method-invite': SIP Method INVITE; 'method-ack': SIP Method ACK; 'method-cancel': SIP Method CANCEL; 'method-bye': SIP Method BYE; 'method-options': SIP Method OPTIONS; 'method-prack': SIP Method PRACK; 'method-subscribe': SIP Method SUBSCRIBE; 'method-notify': SIP Method NOTIFY; 'method-publish': SIP Method PUBLISH; 'method-info': SIP Method INFO; 'method-refer': SIP Method REFER; 'method-message': SIP Method MESSAGE; 'method-update': SIP Method UPDATE; 'method-unknown': SIP Method UNKNOWN; 'parse-error': SIP Message Parse Error; 'req-uri-op-failrue': SIP Operate Request Uri Failure; 'via-hdr-op-failrue': SIP Operate Via Header Failure; 'contact-hdr-op-failrue': SIP Operate Contact Header Failure; 'from-hdr-op-failrue': SIP Operate From Header Failure; 'to-hdr-op-failrue': SIP Operate To Header Failure; 'route-hdr-op-failrue': SIP Operate Route Header Failure; 'record-route-hdr-op-failrue': SIP Operate Record-Route Header Failure; 'content-length-hdr-op-failrue': SIP Operate Content-Length Failure; 'third-party-registration': SIP Third-Party Registration; 'conn-ext-creation-failure': SIP Create Connection Extension Failure; 'alloc-contact-port-failure': SIP Alloc Contact Port Failure; 'outside-contact-port-mismatch': SIP Outside Contact Port Mismatch NAT Port; 'inside-contact-port-mismatch': SIP Inside Contact Port Mismatch; 'third-party-sdp': SIP Third-Party SDP; 'sdp-process-candidate-failure': SIP Operate SDP Media Candidate Attribute Failure; 'sdp-op-failure': SIP Operate SDP Failure; 'sdp-alloc-port-map-success': SIP Alloc SDP Port Map Success; 'sdp-alloc-port-map-failure': SIP Alloc SDP Port Map Failure; 'modify-failure': SIP Message Modify Failure; 'rewrite-failure': SIP Message Rewrite Failure; 'tcp-out-of-order-drop': TCP Out-of-Order Drop; 'smp-conn-alloc-failure': SMP Helper Conn Alloc Failure; 'helper-found': SMP Helper Conn Found; 'helper-created': SMP Helper Conn Created; 'helper-deleted': SMP Helper Conn Already Deleted; 'helper-freed': SMP Helper Conn Freed; 'helper-failure': SMP Helper Failure;",
						},
					},
				},
			},
			"sip_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable SIP ALG for LSN;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnAlgSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgSipRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnAlgSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgSipRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnAlgSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnAlgSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnAlgSipSamplingEnable(d []interface{}) []edpt.Cgnv6LsnAlgSipSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnAlgSipSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnAlgSipSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgSip(d *schema.ResourceData) edpt.Cgnv6LsnAlgSip {
	var ret edpt.Cgnv6LsnAlgSip
	ret.Inst.RtpStunTimeout = d.Get("rtp_stun_timeout").(int)
	ret.Inst.SamplingEnable = getSliceCgnv6LsnAlgSipSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SipValue = d.Get("sip_value").(string)
	//omit uuid
	return ret
}
