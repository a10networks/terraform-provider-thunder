package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLoggingGtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_logging_gtp`: Counters for GTP Logging\n\n__PLACEHOLDER__",
		CreateContext: resourceFwLoggingGtpCreate,
		UpdateContext: resourceFwLoggingGtpUpdate,
		ReadContext:   resourceFwLoggingGtpRead,
		DeleteContext: resourceFwLoggingGtpDelete,

		Schema: map[string]*schema.Schema{
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwLoggingGtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingGtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLoggingGtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwLoggingGtpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwLoggingGtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingGtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLoggingGtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwLoggingGtpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwLoggingGtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingGtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLoggingGtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwLoggingGtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingGtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLoggingGtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwLoggingGtpSamplingEnable(d []interface{}) []edpt.FwLoggingGtpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwLoggingGtpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwLoggingGtpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwLoggingGtp(d *schema.ResourceData) edpt.FwLoggingGtp {
	var ret edpt.FwLoggingGtp
	ret.Inst.SamplingEnable = getSliceFwLoggingGtpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
