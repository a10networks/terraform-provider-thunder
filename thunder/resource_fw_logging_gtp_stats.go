package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLoggingGtpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_logging_gtp_stats`: Statistics for the object gtp\n\n__PLACEHOLDER__",
		ReadContext: resourceFwLoggingGtpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_type_gtp_invalid_teid": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Invalid TEID",
						},
						"log_gtp_type_reserved_ie_present": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Reserved IE Present",
						},
						"log_type_gtp_mandatory_ie_missing": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Mandatory IE Missing",
						},
						"log_type_gtp_mandatory_ie_inside_grouped_ie_missing": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Mandatory IE Missing Inside Grouped IE",
						},
						"log_type_gtp_msisdn_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP MSISDN Filtering",
						},
						"log_type_gtp_out_of_order_ie": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Out of Order IE V1",
						},
						"log_type_gtp_out_of_state_ie": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Out of State IE",
						},
						"log_type_enduser_ip_spoofed": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Enduser IP Spoofed",
						},
						"log_type_crosslayer_correlation": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Crosslayer Correlation",
						},
						"log_type_message_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Reserved Message Found",
						},
						"log_type_out_of_state": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Out of State Message",
						},
						"log_type_max_msg_length": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Message Length Exceeded Max",
						},
						"log_type_gtp_message_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Message Filtering",
						},
						"log_type_gtp_apn_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Apn Filtering",
						},
						"log_type_gtp_rat_type_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP RAT Type Filtering",
						},
						"log_type_country_code_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Country Code Mismatch",
						},
						"log_type_gtp_in_gtp_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP in GTP Filtering",
						},
						"log_type_gtp_node_restart": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP SGW/PGW restarted",
						},
						"log_type_gtp_seq_num_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Response Sequence number Mismatch",
						},
						"log_type_gtp_rate_limit_periodic": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Rate Limit Periodic",
						},
						"log_type_gtp_invalid_message_length": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Invalid message length across layers",
						},
						"log_type_gtp_hdr_invalid_protocol_flag": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Protocol flag in header",
						},
						"log_type_gtp_hdr_invalid_spare_bits": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP invalid spare bits in header",
						},
						"log_type_gtp_hdr_invalid_piggy_flag": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event GTP invalid piggyback flag in header",
						},
						"log_type_gtp_invalid_version": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event invalid GTP version",
						},
						"log_type_gtp_invalid_ports": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event mismatch of GTP message and ports",
						},
					},
				},
			},
		},
	}
}

func resourceFwLoggingGtpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingGtpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLoggingGtpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwLoggingGtpStatsStats := setObjectFwLoggingGtpStatsStats(res)
		d.Set("stats", FwLoggingGtpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwLoggingGtpStatsStats(ret edpt.DataFwLoggingGtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"log_type_gtp_invalid_teid":                           ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_invalid_teid,
			"log_gtp_type_reserved_ie_present":                    ret.DtFwLoggingGtpStats.Stats.Log_gtp_type_reserved_ie_present,
			"log_type_gtp_mandatory_ie_missing":                   ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_mandatory_ie_missing,
			"log_type_gtp_mandatory_ie_inside_grouped_ie_missing": ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_mandatory_ie_inside_grouped_ie_missing,
			"log_type_gtp_msisdn_filtering":                       ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_msisdn_filtering,
			"log_type_gtp_out_of_order_ie":                        ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_out_of_order_ie,
			"log_type_gtp_out_of_state_ie":                        ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_out_of_state_ie,
			"log_type_enduser_ip_spoofed":                         ret.DtFwLoggingGtpStats.Stats.Log_type_enduser_ip_spoofed,
			"log_type_crosslayer_correlation":                     ret.DtFwLoggingGtpStats.Stats.Log_type_crosslayer_correlation,
			"log_type_message_not_supported":                      ret.DtFwLoggingGtpStats.Stats.Log_type_message_not_supported,
			"log_type_out_of_state":                               ret.DtFwLoggingGtpStats.Stats.Log_type_out_of_state,
			"log_type_max_msg_length":                             ret.DtFwLoggingGtpStats.Stats.Log_type_max_msg_length,
			"log_type_gtp_message_filtering":                      ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_message_filtering,
			"log_type_gtp_apn_filtering":                          ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_apn_filtering,
			"log_type_gtp_rat_type_filtering":                     ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_rat_type_filtering,
			"log_type_country_code_mismatch":                      ret.DtFwLoggingGtpStats.Stats.Log_type_country_code_mismatch,
			"log_type_gtp_in_gtp_filtering":                       ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_in_gtp_filtering,
			"log_type_gtp_node_restart":                           ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_node_restart,
			"log_type_gtp_seq_num_mismatch":                       ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_seq_num_mismatch,
			"log_type_gtp_rate_limit_periodic":                    ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_rate_limit_periodic,
			"log_type_gtp_invalid_message_length":                 ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_invalid_message_length,
			"log_type_gtp_hdr_invalid_protocol_flag":              ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_hdr_invalid_protocol_flag,
			"log_type_gtp_hdr_invalid_spare_bits":                 ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_hdr_invalid_spare_bits,
			"log_type_gtp_hdr_invalid_piggy_flag":                 ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_hdr_invalid_piggy_flag,
			"log_type_gtp_invalid_version":                        ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_invalid_version,
			"log_type_gtp_invalid_ports":                          ret.DtFwLoggingGtpStats.Stats.Log_type_gtp_invalid_ports,
		},
	}
}

func getObjectFwLoggingGtpStatsStats(d []interface{}) edpt.FwLoggingGtpStatsStats {

	count1 := len(d)
	var ret edpt.FwLoggingGtpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Log_type_gtp_invalid_teid = in["log_type_gtp_invalid_teid"].(int)
		ret.Log_gtp_type_reserved_ie_present = in["log_gtp_type_reserved_ie_present"].(int)
		ret.Log_type_gtp_mandatory_ie_missing = in["log_type_gtp_mandatory_ie_missing"].(int)
		ret.Log_type_gtp_mandatory_ie_inside_grouped_ie_missing = in["log_type_gtp_mandatory_ie_inside_grouped_ie_missing"].(int)
		ret.Log_type_gtp_msisdn_filtering = in["log_type_gtp_msisdn_filtering"].(int)
		ret.Log_type_gtp_out_of_order_ie = in["log_type_gtp_out_of_order_ie"].(int)
		ret.Log_type_gtp_out_of_state_ie = in["log_type_gtp_out_of_state_ie"].(int)
		ret.Log_type_enduser_ip_spoofed = in["log_type_enduser_ip_spoofed"].(int)
		ret.Log_type_crosslayer_correlation = in["log_type_crosslayer_correlation"].(int)
		ret.Log_type_message_not_supported = in["log_type_message_not_supported"].(int)
		ret.Log_type_out_of_state = in["log_type_out_of_state"].(int)
		ret.Log_type_max_msg_length = in["log_type_max_msg_length"].(int)
		ret.Log_type_gtp_message_filtering = in["log_type_gtp_message_filtering"].(int)
		ret.Log_type_gtp_apn_filtering = in["log_type_gtp_apn_filtering"].(int)
		ret.Log_type_gtp_rat_type_filtering = in["log_type_gtp_rat_type_filtering"].(int)
		ret.Log_type_country_code_mismatch = in["log_type_country_code_mismatch"].(int)
		ret.Log_type_gtp_in_gtp_filtering = in["log_type_gtp_in_gtp_filtering"].(int)
		ret.Log_type_gtp_node_restart = in["log_type_gtp_node_restart"].(int)
		ret.Log_type_gtp_seq_num_mismatch = in["log_type_gtp_seq_num_mismatch"].(int)
		ret.Log_type_gtp_rate_limit_periodic = in["log_type_gtp_rate_limit_periodic"].(int)
		ret.Log_type_gtp_invalid_message_length = in["log_type_gtp_invalid_message_length"].(int)
		ret.Log_type_gtp_hdr_invalid_protocol_flag = in["log_type_gtp_hdr_invalid_protocol_flag"].(int)
		ret.Log_type_gtp_hdr_invalid_spare_bits = in["log_type_gtp_hdr_invalid_spare_bits"].(int)
		ret.Log_type_gtp_hdr_invalid_piggy_flag = in["log_type_gtp_hdr_invalid_piggy_flag"].(int)
		ret.Log_type_gtp_invalid_version = in["log_type_gtp_invalid_version"].(int)
		ret.Log_type_gtp_invalid_ports = in["log_type_gtp_invalid_ports"].(int)
	}
	return ret
}

func dataToEndpointFwLoggingGtpStats(d *schema.ResourceData) edpt.FwLoggingGtpStats {
	var ret edpt.FwLoggingGtpStats

	ret.Stats = getObjectFwLoggingGtpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
