package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_policy`: Configure GTP Policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpPolicyCreate,
		UpdateContext: resourceTemplateGtpPolicyUpdate,
		ReadContext:   resourceTemplateGtpPolicyRead,
		DeleteContext: resourceTemplateGtpPolicyDelete,

		Schema: map[string]*schema.Schema{
			"filtering_policy_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify GTP Filtering Policy",
			},
			"general_policy_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify GTP General Policy",
			},
			"logging_policy_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify GTP Logging Policy",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Configure the GTP Policy Name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"rate_limit_policy_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify Rate Limit Policy",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'gtp-v0-c-tunnel-created': GTPv0-C Tunnel Created; 'gtp-v0-c-tunnel-half-open': GTPv0-C Half open tunnel created; 'gtp-v0-c-tunnel-half-closed': GTPv0-C Tunnel Delete Request; 'gtp-v0-c-tunnel-closed': GTPv0-C Tunnel Marked Deleted; 'gtp-v0-c-tunnel-deleted': GTPv0-C Tunnel Deleted; 'gtp-v0-c-half-open-tunnel-closed': GTPv0-C Half open tunnel closed; 'gtp-v1-c-tunnel-created': GTPv1-C Tunnel Created; 'gtp-v1-c-tunnel-half-open': GTPv1-C Half open tunnel created; 'gtp-v1-c-tunnel-half-closed': GTPv1-C Tunnel Delete Request; 'gtp-v1-c-tunnel-closed': GTPv1-C Tunnel Marked Deleted; 'gtp-v1-c-tunnel-deleted': GTPv1-C Tunnel Deleted; 'gtp-v1-c-half-open-tunnel-closed': GTPv1-C Half open tunnel closed; 'gtp-v2-c-tunnel-created': GTPv2-C Tunnel Created; 'gtp-v2-c-tunnel-half-open': GTPv2-C Half open tunnel created; 'gtp-v2-c-tunnel-half-closed': GTPv2-C Tunnel Delete Request; 'gtp-v2-c-tunnel-closed': GTPv2-C Tunnel Marked Deleted; 'gtp-v2-c-tunnel-deleted': GTPv2-C Tunnel Deleted; 'gtp-v2-c-half-open-tunnel-closed': GTPv2-C Half open tunnel closed; 'gtp-u-tunnel-created': GTP-U Tunnel Created; 'gtp-u-tunnel-deleted': GTP-U Tunnel Deleted; 'gtp-v0-c-update-pdp-resp-unsuccess': GTPv0-C Update PDP Context Response Unsuccessful; 'gtp-v1-c-update-pdp-resp-unsuccess': GTPv1-C Update PDP Context Response Unsuccessful; 'gtp-v2-c-mod_bearer-resp-unsuccess': GTPv2-C Modify Bearer Response Unsuccessful; 'gtp-v0-c-create-pdp-resp-unsuccess': GTPv0-C Create PDP Context Response Unsuccessful; 'gtp-v1-c-create-pdp-resp-unsuccess': GTPv1-C Create PDP Context Response Unsuccessful; 'gtp-v2-c-create-sess-resp-unsuccess': GTPv2-C Create Session Response Unsuccessful; 'gtp-v2-c-piggyback-message': GTPv2-C Piggyback Message; 'gtp-path-management-message': GTP Path Management Messages Received; 'gtp-v0-c-tunnel-deleted-restart': GTPv0-C Tunnel Deleted with Restart/failure; 'gtp-v1-c-tunnel-deleted-restart': GTPv1-C Tunnel Deleted with Restart/failure; 'gtp-v2-c-tunnel-deleted-restart': GTPv2-C Tunnel Deleted with Restart/failure; 'gtp-v0-c-reserved-message-allow': Permit GTPv0-C Reserved Messages; 'gtp-v1-c-reserved-message-allow': Permit GTPv1-C Reserved Messages; 'gtp-v2-c-reserved-message-allow': Permit GTPv2-C Reserved Messages; 'gtp-v2-c-load-contr-info-exceed': GTPv2-C Load Control Info IEs in message exceeded 2; 'gtp-v1-c-pdu-notification-request-forward': GTPv1-C PDU Notification Request Forward; 'gtp-v1-c-pdu-notification-reject-request-forward': GTPv1-C PDU Notification Reject Request Forward; 'gtp-v0-c-pdu-notification-request-forward': GTPv0-C PDU Notification Request Forward; 'gtp-v0-c-pdu-notification-reject-request-forward': GTPv0-C PDU Notification Reject Request Forward; 'gtp-v0-c-message-skipped-apn-filtering-no-imsi': GTPv0-C APN/IMSI Filtering Skipped (No IMSI); 'gtp-v1-c-message-skipped-apn-filtering-no-imsi': GTPv1-C APN/IMSI Filtering Skipped (No IMSI); 'gtp-v2-c-message-skipped-apn-filtering-no-imsi': GTPv2-C APN/IMSI Filtering Skipped (No IMSI); 'gtp-v0-c-message-skipped-msisdn-filtering-no-imsi': GTPv0-C MSISDN Filtering Skipped (No MSISDN); 'gtp-v1-c-message-skipped-msisdn-filtering-no-imsi': GTPv1-C MSISDN Filtering Skipped (No MSISDN); 'gtp-v2-c-message-skipped-msisdn-filtering-no-imsi': GTPv2-C MSISDN Filtering Skipped (No MSISDN); 'gtp-v0-c-packet-dummy-msisdn': GTPv0-C Packet With Dummy MSISDN Forwarded; 'gtp-v1-c-packet-dummy-msisdn': GTPv1-C Packet With Dummy MSISDN Forwarded; 'gtp-v2-c-packet-dummy-msisdn': GTPv2-C Packet With Dummy MSISDN Forwarded; 'drop-vld-sanity-gtp-v2-c-message-with-teid-zero-expected': Validation Drop: GTPv2-C Create Session Request with TEID; 'drop-vld-sanity-gtp-v1-c-message-with-teid-zero-expected': Validation Drop: GTPv1-C PDU Notification Request with TEID; 'drop-vld-sanity-gtp-v0-c-message-with-teid-zero-expected': Validation Drop: GTPv0-C PDU Notification Request with TEID; 'drop-vld-gtp-ie-repeat-count-exceed': Validation Drop: GTP repeated IE count exceeded; 'drop-vld-reserved-field-set': Validation Drop: Reserved Header Field Set; 'drop-vld-tunnel-id-flag': Validation Drop: Tunnel Header Flag Not Set; 'drop-vld-invalid-flow-label-v0': Validation Drop: Invalid Flow Label in GTPv0-C Header; 'drop-vld-invalid-teid': Validation Drop: Invalid TEID Value; 'drop-vld-out-of-state': Validation Drop: Out Of State GTP Message; 'drop-vld-mandatory-information-element': Validation Drop: Mandatory IE Not Present; 'drop-vld-mandatory-ie-in-grouped-ie': Validation Drop: Mandatory IE in Grouped IE Not Present; 'drop-vld-out-of-order-ie': Validation Drop: GTPv1-C Message Out of Order IE; 'drop-vld-out-of-state-ie': Validation Drop: Unexpected IE Present in Message; 'drop-vld-reserved-information-element': Validation Drop: Reserved IE Field Present; 'drop-vld-version-not-supported': Validation Drop: Invalid GTP version; 'drop-vld-message-length': Validation Drop: Message Length Exceeded; 'drop-vld-cross-layer-correlation': Validation Drop: Cross Layer IP Address Mismatch; 'drop-vld-country-code-mismatch': Validation Drop: Country Code Mismatch in IMSI and MSISDN; 'drop-vld-gtp-u-spoofed-source-address': Validation Drop: GTP-U IP Address Spoofed; 'drop-vld-gtp-bearer-count-exceed': Validation Drop: GTP Bearer count exceeded max (11); 'drop-vld-gtp-v2-wrong-lbi-create-bearer-req': Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request; 'gtp-c-handover-in-progress-with-conn': GTP-C matching a conn with Handover In Progress; 'drop-vld-v0-reserved-message-drop': Validation Drop: GTPv0-C Reserved Message Drop; 'drop-vld-v1-reserved-message-drop': Validation Drop: GTPv1-C Reserved Message Drop; 'drop-vld-v2-reserved-message-drop': Validation Drop: GTPv2-C Reserved Message Drop; 'drop-vld-invalid-pkt-len-piggyback': Validation Drop: Piggyback message invalid packet length; 'drop-vld-sanity-failed-piggyback': Validation Drop: piggyback message anomaly failed; 'drop-vld-sequence-num-correlation': Validation Drop: GTP-C Sequence number Mismatch; 'drop-vld-gtpv0-seqnum-buffer-full': Validation Drop: GTPV0-C conn Sequence number Buffer Full; 'drop-vld-gtpv1-seqnum-buffer-full': Validation Drop: GTPV1-C conn Sequence number Buffer Full; 'drop-vld-gtpv2-seqnum-buffer-full': Validation Drop: GTPV2-C conn Sequence number Buffer Full; 'drop-vld-gtp-invalid-imsi-len-drop': Validation Drop: GTP-C Invalid IMSI Length Drop; 'drop-vld-gtp-invalid-apn-len-drop': Validation Drop: GTP-C Invalid APN Length Drop; 'drop-vld-protocol-flag-unset': Validation Drop: Protocol flag in Header Field not Set; 'drop-vld-gtpv0-subscriber-attr-miss': Validation Drop: GTPV0-c Subscriber Attributes Missing; 'drop-vld-gtpv1-subscriber-attr-miss': Validation Drop: GTPV1-c Subscriber Attributes Missing; 'drop-vld-gtpv2-subscriber-attr-miss': Validation Drop: GTPV2-c Subscriber Attributes Missing; 'drop-vld-gtp-v0-c-ie-len-exceed-msg-len': GTPv0-C IE Length Exceeds Message Length; 'drop-vld-gtp-v1-c-ie-len-exceed-msg-len': GTPv1-C IE Length Exceeds Message Length; 'drop-vld-gtp-v2-c-ie-len-exceed-msg-len': GTPv2-C IE Length Exceeds Message Length; 'drop-vld-gtp-v0-c-message-length-mismatch': GTPv0-C Message Length Mismatch Across Layers; 'drop-vld-gtp-v1-c-message-length-mismatch': GTPv1-C Message Length Mismatch Across Layers; 'drop-vld-gtp-v2-c-message-length-mismatch': GTPv2-C Message Length Mismatch Across Layers; 'drop-vld-gtp-v0-c-message-skipped-apn-filtering-no-apn': Validation Drop: GTPv0-C APN/IMSI Filtering Dropped (No APN); 'drop-vld-gtp-v1-c-message-skipped-apn-filtering-no-apn': Validation Drop: GTPv1-C APN/IMSI Filtering Dropped (No APN); 'drop-vld-gtp-v2-c-message-skipped-apn-filtering-no-apn': Validation Drop: GTPv2-C APN/IMSI Filtering Dropped (No APN);",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'drop-flt-message-filtering': Filtering Drop: Message Type Not Permitted on Interface; 'drop-flt-apn-filtering': Filtering Drop: APN IMSI Filtering; 'drop-flt-msisdn-filtering': Filtering Drop: MSISDN Filtering; 'drop-flt-rat-type-filtering': Filtering Drop: RAT Type Filtering; 'drop-flt-gtp-in-gtp': Filtering Drop: GTP in GTP Tunnel Present; 'drop-rl-gtp-v0-c-agg': Rate-limit Drop: Maximum GTPv0-C Message rate; 'drop-rl-gtp-v1-c-agg': Rate-limit Drop: Maximum GTPv1-C Message rate; 'drop-rl-gtp-v2-c-agg': Rate-limit Drop: Maximum GTPv2-C Message rate; 'drop-rl-gtp-v1-c-create-pdp-request': Rate-limit Drop: GTPv1-C Create PDP Request rate; 'drop-rl-gtp-v2-c-create-session-request': Rate-limit Drop: GTPv2-C Create Session Request rate; 'drop-rl-gtp-v1-c-update-pdp-request': Rate-limit Drop: GTPv1-C Update PDP Request rate; 'drop-rl-gtp-v2-c-modify-bearer-request': Rate-limit Drop: GTPv2-C Modify Bearer Request rate; 'drop-rl-gtp-u-tunnel-create': Rate-limit Drop: GTP-U Tunnel Creation rate; 'drop-rl-gtp-u-uplink-byte': Rate-limit Drop: GTP-U Uplink byte rate; 'drop-rl-gtp-u-uplink-packet': Rate-limit Drop: GTP-U Uplink packet rate; 'drop-rl-gtp-u-downlink-byte': Rate-limit Drop: GTP-U Downlink byte rate; 'drop-rl-gtp-u-downlink-packet': Rate-limit Drop: GTP-U Downlink packet rate; 'drop-rl-gtp-u-total-byte': Rate-limit Drop: GTP-U Total byte rate; 'drop-rl-gtp-u-total-packet': Rate-limit Drop: GTP-U Total packet rate; 'drop-rl-gtp-u-max-concurrent-tunnels': Rate-limit Drop: GTP-U Concurrent Tunnels;",
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
			"validation_policy_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify GTP Validation Policy",
			},
		},
	}
}
func resourceTemplateGtpPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceTemplateGtpPolicySamplingEnable(d []interface{}) []edpt.TemplateGtpPolicySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.TemplateGtpPolicySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateGtpPolicySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTemplateGtpPolicy(d *schema.ResourceData) edpt.TemplateGtpPolicy {
	var ret edpt.TemplateGtpPolicy
	ret.Inst.FilteringPolicyName = d.Get("filtering_policy_name").(string)
	ret.Inst.GeneralPolicyName = d.Get("general_policy_name").(string)
	ret.Inst.LoggingPolicyName = d.Get("logging_policy_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.RateLimitPolicyName = d.Get("rate_limit_policy_name").(string)
	ret.Inst.SamplingEnable = getSliceTemplateGtpPolicySamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ValidationPolicyName = d.Get("validation_policy_name").(string)
	return ret
}
