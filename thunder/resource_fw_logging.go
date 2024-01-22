package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_logging`: Bind a logging template to firewall\n\n__PLACEHOLDER__",
		CreateContext: resourceFwLoggingCreate,
		UpdateContext: resourceFwLoggingUpdate,
		ReadContext:   resourceFwLoggingRead,
		DeleteContext: resourceFwLoggingDelete,

		Schema: map[string]*schema.Schema{
			"gtp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'log_type_gtp_invalid_teid': Log Event Type GTP Invalid TEID; 'log_gtp_type_reserved_ie_present': Log Event Type GTP Reserved IE Present; 'log_type_gtp_mandatory_ie_missing': Log Event Type GTP Mandatory IE Missing; 'log_type_gtp_mandatory_ie_inside_grouped_ie_missing': Log Event Type GTP Mandatory IE Missing Inside Grouped IE; 'log_type_gtp_msisdn_filtering': Log Event Type GTP MSISDN Filtering; 'log_type_gtp_out_of_order_ie': Log Event Type GTP Out of Order IE V1; 'log_type_gtp_out_of_state_ie': Log Event Type GTP Out of State IE; 'log_type_enduser_ip_spoofed': Log Event Type GTP Enduser IP Spoofed; 'log_type_crosslayer_correlation': Log Event GTP Crosslayer Correlation; 'log_type_message_not_supported': Log Event GTP Reserved Message Found; 'log_type_out_of_state': Log Event GTP Out of State Message; 'log_type_max_msg_length': Log Event GTP Message Length Exceeded Max; 'log_type_gtp_message_filtering': Log Event Type GTP Message Filtering; 'log_type_gtp_apn_filtering': Log Event Type GTP Apn Filtering; 'log_type_gtp_rat_type_filtering': Log Event GTP RAT Type Filtering; 'log_type_country_code_mismatch': Log Event GTP Country Code Mismatch; 'log_type_gtp_in_gtp_filtering': Log Event GTP in GTP Filtering; 'log_type_gtp_node_restart': Log Event GTP SGW/PGW restarted; 'log_type_gtp_seq_num_mismatch': Log Event GTP Response Sequence number Mismatch; 'log_type_gtp_rate_limit_periodic': Log Event GTP Rate Limit Periodic; 'log_type_gtp_invalid_message_length': Log Event GTP Invalid message length across layers; 'log_type_gtp_hdr_invalid_protocol_flag': Log Event GTP Protocol flag in header; 'log_type_gtp_hdr_invalid_spare_bits': Log Event GTP invalid spare bits in header; 'log_type_gtp_hdr_invalid_piggy_flag': Log Event GTP invalid piggyback flag in header; 'log_type_gtp_invalid_version': Log Event invalid GTP version; 'log_type_gtp_invalid_ports': Log Event mismatch of GTP message and ports;",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Logging Template Name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'log_message_sent': Log Packet Sent; 'log_type_reset': Log Event Type Reset; 'log_type_deny': Log Event Type Deny; 'log_type_session_closed': Log Event Type Session Close; 'log_type_session_opened': Log Event Type Session Open; 'rule_not_logged': Firewall Rule Not Logged; 'log-dropped': Log Packets Dropped; 'tcp-session-created': TCP Session Created; 'tcp-session-deleted': TCP Session Deleted; 'udp-session-created': UDP Session Created; 'udp-session-deleted': UDP Session Deleted; 'icmp-session-deleted': ICMP Session Deleted; 'icmp-session-created': ICMP Session Created; 'icmpv6-session-deleted': ICMPV6 Session Deleted; 'icmpv6-session-created': ICMPV6 Session Created; 'other-session-deleted': Other Session Deleted; 'other-session-created': Other Session Created; 'http-request-logged': HTTP Request Logged; 'http-logging-invalid-format': HTTP Logging Invalid Format Error; 'dcmsg_permit': Dcmsg Permit; 'alg_override_permit': Alg Override Permit; 'template_error': Template Error; 'ipv4-frag-applied': IPv4 Fragmentation Applied; 'ipv4-frag-failed': IPv4 Fragmentation Failed; 'ipv6-frag-applied': IPv6 Fragmentation Applied; 'ipv6-frag-failed': IPv6 Fragmentation Failed; 'out-of-buffers': Out of Buffers; 'add-msg-failed': Add Message to Buffer Failed; 'tcp-logging-conn-established': TCP Logging Conn Established; 'tcp-logging-conn-create-failed': TCP Logging Conn Create Failed; 'tcp-logging-conn-dropped': TCP Logging Conn Dropped; 'log-message-too-long': Log message too long; 'http-out-of-order-dropped': HTTP out-of-order dropped; 'http-alloc-failed': HTTP Request Info Allocation Failed; 'sctp-session-created': SCTP Session Created; 'sctp-session-deleted': SCTP Session Deleted; 'log_type_sctp_inner_proto_filter': Log Event Type SCTP Inner Proto Filter; 'tcp-logging-port-allocated': TCP Logging Port Allocated; 'tcp-logging-port-freed': TCP Logging Port Freed; 'tcp-logging-port-allocation-failed': TCP Logging Port Allocation Failed; 'iddos-blackhole-entry-create': iDDoS IP Entry Created; 'iddos-blackhole-entry-delete': iDDoS IP Entry Deleted; 'session-limit-exceeded': Session Limit Exceeded;",
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
func resourceFwLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceFwLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceFwLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFwLoggingGtp371(d []interface{}) edpt.FwLoggingGtp371 {

	count1 := len(d)
	var ret edpt.FwLoggingGtp371
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceFwLoggingGtpSamplingEnable372(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceFwLoggingGtpSamplingEnable372(d []interface{}) []edpt.FwLoggingGtpSamplingEnable372 {

	count1 := len(d)
	ret := make([]edpt.FwLoggingGtpSamplingEnable372, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwLoggingGtpSamplingEnable372
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwLoggingSamplingEnable(d []interface{}) []edpt.FwLoggingSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwLoggingSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwLoggingSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwLogging(d *schema.ResourceData) edpt.FwLogging {
	var ret edpt.FwLogging
	ret.Inst.Gtp = getObjectFwLoggingGtp371(d.Get("gtp").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceFwLoggingSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
