package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlg() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg`: Configure ALG\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgCreate,
		UpdateContext: resourceFwAlgUpdate,
		ReadContext:   resourceFwAlgRead,
		DeleteContext: resourceFwAlgDelete,

		Schema: map[string]*schema.Schema{
			"dns": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_port_disable": {
							Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable DNS ALG default port 53;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"esp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_port_disable": {
							Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable ESP ALG default port 500;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'session-created': ESP Sessions Created; 'helper-created': ESP Helper Sessions Created; 'helper-freed': ESP Helper Sessions Freed; 'helper-freed-used': ESP Helper Sessions freed used; 'helper-freed-unused': ESP Helper Sessions freed unused; 'helper-already-used': ESP Helper Session already used; 'helper-in-rml': ESP Helper Session in Remove List;",
									},
								},
							},
						},
					},
				},
			},
			"ftp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_port_disable": {
							Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable FTP ALG default port 21;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'client-port-request': PORT Requests From Client; 'client-eprt-request': EPRT Requests From Client; 'server-pasv-reply': PASV Replies From Server; 'server-epsv-reply': EPSV Replies From Server; 'port-retransmits': PORT Retransmits; 'pasv-retransmits': PASV Retransmits; 'smp-app-type-mismatch': SMP App Type Mismatch; 'retransmit-sanity-check-failure': Retransmit Sanity Check Failure; 'smp-conn-alloc-failure': SMP Helper Conn Alloc Failure; 'port-helper-created': PORT Helper Created; 'pasv-helper-created': PASV Helper Created; 'port-helper-acquire-in-del-q': PORT Helper Acquire In Del Queue; 'port-helper-acquire-already-used': PORT Helper Acquire Already Used; 'pasv-helper-acquire-in-del-q': PASV Helper Acquire In Del Queue; 'pasv-helper-acquire-already-used': PASV Helper Acquire Already Used; 'port-helper-freed-used': PORT Helper Freed Used; 'port-helper-freed-unused': PORT Helper Freed Unused; 'pasv-helper-freed-used': PASV Helper Freed Used; 'pasv-helper-freed-unused': PASV Helper Freed Unused;",
									},
								},
							},
						},
					},
				},
			},
			"icmp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type: schema.TypeString, Optional: true, Description: "'disable': Disable ICMP ALG which allows ICMP errors to pass the firewall;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pptp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_port_disable": {
							Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable PPTP ALG default port 1723;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'calls-established': Calls Established; 'call-req-pns-call-id-mismatch': Call ID Mismatch on Call Request; 'call-reply-pns-call-id-mismatch': Call ID Mismatch on Call Reply; 'gre-session-created': GRE Session Created; 'gre-session-freed': GRE Session Freed; 'call-req-retransmit': Call Request Retransmit; 'call-req-new': Call Request New; 'call-req-ext-alloc-failure': Call Request Ext Alloc Failure; 'call-reply-call-id-unknown': Call Reply Unknown Client Call ID; 'call-reply-retransmit': Call Reply Retransmit; 'call-reply-ext-ext-alloc-failure': Call Request Ext Alloc Failure; 'smp-app-type-mismatch': SMP App Type Mismatch; 'smp-client-call-id-mismatch': SMP Client Call ID Mismatch; 'smp-sessions-created': SMP Session Created; 'smp-sessions-freed': SMP Session Freed; 'smp-alloc-failure': SMP Session Alloc Failure; 'gre-conn-creation-failure': GRE Conn Alloc Failure; 'gre-conn-ext-creation-failure': GRE Conn Ext Alloc Failure; 'gre-no-fwd-route': GRE No Fwd Route; 'gre-no-rev-route': GRE No Rev Route; 'gre-no-control-conn': GRE No Control Conn; 'gre-conn-already-exists': GRE Conn Already Exists; 'gre-free-no-ext': GRE Free No Ext; 'gre-free-no-smp': GRE Free No SMP; 'gre-free-smp-app-type-mismatch': GRE Free SMP App Type Mismatch; 'control-freed': Control Session Freed; 'control-free-no-ext': Control Free No Ext; 'control-free-no-smp': Control Free No SMP; 'control-free-smp-app-type-mismatch': Control Free SMP App Type Mismatch;",
									},
								},
							},
						},
					},
				},
			},
			"rtsp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_port_disable": {
							Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable RTSP ALG default port 554;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'transport-inserted': Transport Created; 'transport-freed': Transport Freed; 'transport-alloc-failure': Transport Alloc Failure; 'data-session-created': Data Session Created; 'data-session-freed': Data Session Freed; 'ext-creation-failure': Extension Creation Failure; 'transport-add-to-ext': Transport Added to Extension; 'transport-removed-from-ext': Transport Removed from Extension; 'transport-too-many': Too Many Transports for Control Conn; 'transport-already-in-ext': Transport Already in Extension; 'transport-exists': Transport Already Exists; 'transport-link-ext-failure-control': Transport Link to Extension Failure Control; 'transport-link-ext-data': Transport Link to Extension Data; 'transport-link-ext-failure-data': Transport Link to Extension Failure Data; 'transport-inserted-shadow': Transport Inserted Shadow; 'transport-creation-race': Transport Create Race; 'transport-alloc-failure-shadow': Transport Alloc Failure Shadow; 'transport-put-in-del-q': Transport Put in Delete Queue; 'transport-freed-shadow': Transport Freed Shadow; 'transport-acquired-from-control': Transport Acquired Control; 'transport-found-from-prev-control': Transport Found From Prev Control; 'transport-acquire-failure-from-control': Transport Acquire Failure Control; 'transport-released-from-control': Transport Released Control; 'transport-double-release-from-control': Transport Double Release Control; 'transport-acquired-from-data': Transport Acquired Data; 'transport-acquire-failure-from-data': Transport Acquire Failure Data; 'transport-released-from-data': Transport Released Data; 'transport-double-release-from-data': Transport Double Release Data; 'transport-retry-lookup-on-data-free': Transport Retry Lookup Data; 'transport-not-found-on-data-free': Transport Not Found Data; 'data-session-created-shadow': Data Session Created Shadow; 'data-session-freed-shadow': Data Session Freed Shadow; 'ha-control-ext-creation-failure': HA Control Extension Creation Failure; 'ha-control-session-created': HA Control Session Created; 'ha-data-session-created': HA Data Session Created;",
									},
								},
							},
						},
					},
				},
			},
			"sctp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'disable': disable; 'enable': enable;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_port_disable": {
							Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable SIP ALG default port 5060;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'stat-request': Request Received; 'stat-response': Response Received; 'method-register': Method REGISTER; 'method-invite': Method INVITE; 'method-ack': Method ACK; 'method-cancel': Method CANCEL; 'method-bye': Method BYE; 'method-options': Method OPTIONS; 'method-prack': Method PRACK; 'method-subscribe': Method SUBSCRIBE; 'method-notify': Method NOTIFY; 'method-publish': Method PUBLISH; 'method-info': Method INFO; 'method-refer': Method REFER; 'method-message': Method MESSAGE; 'method-update': Method UPDATE; 'method-unknown': Method Unknown; 'parse-error': Message Parse Error; 'keep-alive': Keep Alive; 'contact-error': Contact Process Error; 'sdp-error': SDP Process Error; 'rtp-port-no-op': RTP Port No Op; 'rtp-rtcp-port-success': RTP RTCP Port Success; 'rtp-port-failure': RTP Port Failure; 'rtcp-port-failure': RTCP Port Failure; 'contact-port-no-op': Contact Port No Op; 'contact-port-success': Contact Port Success; 'contact-port-failure': Contact Port Failure; 'contact-new': Contact Alloc; 'contact-alloc-failure': Contact Alloc Failure; 'contact-eim': Contact EIM; 'contact-eim-set': Contact EIM Set; 'rtp-new': RTP Alloc; 'rtp-alloc-failure': RTP Alloc Failure; 'rtp-eim': RTP EIM; 'helper-found': SMP Helper Conn Found; 'helper-created': SMP Helper Conn Created; 'helper-deleted': SMP Helper Conn Already Deleted; 'helper-freed': SMP Helper Conn Freed; 'helper-failure': SMP Helper Failure;",
									},
								},
							},
						},
					},
				},
			},
			"tftp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_port_disable": {
							Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable TFTP ALG default port 69;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'session-created': TFTP Client Sessions Created; 'helper-created': TFTP Helper Sessions created; 'helper-freed': TFTP Helper Sessions freed; 'helper-freed-used': TFTP Helper Sessions freed used; 'helper-freed-unused': TFTP Helper Sessions freed unused; 'helper-already-used': TFTP Helper Session already used; 'helper-in-rml': TFTP Helper Session in Remove List;",
									},
								},
							},
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
func resourceFwAlgCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlg(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlg(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlg(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlg(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFwAlgDns426(d []interface{}) edpt.FwAlgDns426 {

	count1 := len(d)
	var ret edpt.FwAlgDns426
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortDisable = in["default_port_disable"].(string)
		//omit uuid
	}
	return ret
}

func getObjectFwAlgEsp427(d []interface{}) edpt.FwAlgEsp427 {

	count1 := len(d)
	var ret edpt.FwAlgEsp427
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortDisable = in["default_port_disable"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceFwAlgEspSamplingEnable428(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceFwAlgEspSamplingEnable428(d []interface{}) []edpt.FwAlgEspSamplingEnable428 {

	count1 := len(d)
	ret := make([]edpt.FwAlgEspSamplingEnable428, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgEspSamplingEnable428
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwAlgFtp429(d []interface{}) edpt.FwAlgFtp429 {

	count1 := len(d)
	var ret edpt.FwAlgFtp429
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortDisable = in["default_port_disable"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceFwAlgFtpSamplingEnable430(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceFwAlgFtpSamplingEnable430(d []interface{}) []edpt.FwAlgFtpSamplingEnable430 {

	count1 := len(d)
	ret := make([]edpt.FwAlgFtpSamplingEnable430, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgFtpSamplingEnable430
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwAlgIcmp431(d []interface{}) edpt.FwAlgIcmp431 {

	count1 := len(d)
	var ret edpt.FwAlgIcmp431
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(string)
		//omit uuid
	}
	return ret
}

func getObjectFwAlgPptp432(d []interface{}) edpt.FwAlgPptp432 {

	count1 := len(d)
	var ret edpt.FwAlgPptp432
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortDisable = in["default_port_disable"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceFwAlgPptpSamplingEnable433(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceFwAlgPptpSamplingEnable433(d []interface{}) []edpt.FwAlgPptpSamplingEnable433 {

	count1 := len(d)
	ret := make([]edpt.FwAlgPptpSamplingEnable433, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgPptpSamplingEnable433
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwAlgRtsp434(d []interface{}) edpt.FwAlgRtsp434 {

	count1 := len(d)
	var ret edpt.FwAlgRtsp434
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortDisable = in["default_port_disable"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceFwAlgRtspSamplingEnable435(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceFwAlgRtspSamplingEnable435(d []interface{}) []edpt.FwAlgRtspSamplingEnable435 {

	count1 := len(d)
	ret := make([]edpt.FwAlgRtspSamplingEnable435, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgRtspSamplingEnable435
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwAlgSctp436(d []interface{}) edpt.FwAlgSctp436 {

	count1 := len(d)
	var ret edpt.FwAlgSctp436
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectFwAlgSip437(d []interface{}) edpt.FwAlgSip437 {

	count1 := len(d)
	var ret edpt.FwAlgSip437
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortDisable = in["default_port_disable"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceFwAlgSipSamplingEnable438(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceFwAlgSipSamplingEnable438(d []interface{}) []edpt.FwAlgSipSamplingEnable438 {

	count1 := len(d)
	ret := make([]edpt.FwAlgSipSamplingEnable438, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgSipSamplingEnable438
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwAlgTftp439(d []interface{}) edpt.FwAlgTftp439 {

	count1 := len(d)
	var ret edpt.FwAlgTftp439
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultPortDisable = in["default_port_disable"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceFwAlgTftpSamplingEnable440(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceFwAlgTftpSamplingEnable440(d []interface{}) []edpt.FwAlgTftpSamplingEnable440 {

	count1 := len(d)
	ret := make([]edpt.FwAlgTftpSamplingEnable440, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgTftpSamplingEnable440
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwAlg(d *schema.ResourceData) edpt.FwAlg {
	var ret edpt.FwAlg
	ret.Inst.Dns = getObjectFwAlgDns426(d.Get("dns").([]interface{}))
	ret.Inst.Esp = getObjectFwAlgEsp427(d.Get("esp").([]interface{}))
	ret.Inst.Ftp = getObjectFwAlgFtp429(d.Get("ftp").([]interface{}))
	ret.Inst.Icmp = getObjectFwAlgIcmp431(d.Get("icmp").([]interface{}))
	ret.Inst.Pptp = getObjectFwAlgPptp432(d.Get("pptp").([]interface{}))
	ret.Inst.Rtsp = getObjectFwAlgRtsp434(d.Get("rtsp").([]interface{}))
	ret.Inst.Sctp = getObjectFwAlgSctp436(d.Get("sctp").([]interface{}))
	ret.Inst.Sip = getObjectFwAlgSip437(d.Get("sip").([]interface{}))
	ret.Inst.Tftp = getObjectFwAlgTftp439(d.Get("tftp").([]interface{}))
	//omit uuid
	return ret
}
