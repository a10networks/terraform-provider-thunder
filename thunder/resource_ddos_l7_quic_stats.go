package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL7QuicStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l7_quic_stats`: Statistics for the object l7-quic\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL7QuicStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quic_packet_received": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Packet Received",
						},
						"quic_initial_received": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Initial Packet Received",
						},
						"quic_version_negotiation_received": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Version Negotiation Received",
						},
						"quic_retry_received": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Retry Received",
						},
						"quic_0rtt_recevied": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC 0RTT Received",
						},
						"quic_handshake_received": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Handshake Received",
						},
						"quic_version_match_action_taken": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Version Match Action Taken",
						},
						"quic_version_match_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Version Match Action Drop",
						},
						"quic_version_match_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Version Match Action Blacklist",
						},
						"quic_malformed_action_taken": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Malformed Action Taken",
						},
						"quic_malformed_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Malformed Action Drop",
						},
						"quic_malformed_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Malformed Action Blacklist",
						},
						"quic_malformed_dcid_len_max_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Malformed DCID Len Max Exceed",
						},
						"quic_malformed_scid_len_max_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Malformed SCID Len Max Exceed",
						},
						"quic_fixed_bit_not_set": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Fixed Bit Not Set",
						},
						"quic_retry_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Retry Auth Sent",
						},
						"quic_retry_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Retry Auth Pass",
						},
						"quic_retry_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Retry Auth Fail",
						},
						"quic_connection_close_sent": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Connection Close Sent",
						},
						"quic_invalid_retry_token": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Invalid Retry Token",
						},
						"quic_short_header_received": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Short Header Received",
						},
						"quic_short_header_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Short Header Drop",
						},
						"quic_encrypt_fail": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Encrypt Fail",
						},
						"quic_decrypt_fail": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Decrypt Fail",
						},
						"quic_encrypt_success": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Encrypt Success",
						},
						"quic_decrypt_success": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Decrypt Success",
						},
						"quic_0rtt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC 0RTT Drop",
						},
						"quic_aead_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC AEAD Packet Rate Exceed",
						},
						"quic_dcid_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC DCID Packet Rate Exceed",
						},
						"quic_create_conn_init_only": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Create Connection on Initial Only",
						},
						"quic_version_no_match_drop": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Version No Match Drop",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL7QuicStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7QuicStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7QuicStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL7QuicStatsStats := setObjectDdosL7QuicStatsStats(res)
		d.Set("stats", DdosL7QuicStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL7QuicStatsStats(ret edpt.DataDdosL7QuicStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"quic_packet_received":                ret.DtDdosL7QuicStats.Stats.Quic_packet_received,
			"quic_initial_received":               ret.DtDdosL7QuicStats.Stats.Quic_initial_received,
			"quic_version_negotiation_received":   ret.DtDdosL7QuicStats.Stats.Quic_version_negotiation_received,
			"quic_retry_received":                 ret.DtDdosL7QuicStats.Stats.Quic_retry_received,
			"quic_0rtt_recevied":                  ret.DtDdosL7QuicStats.Stats.Quic_0rtt_recevied,
			"quic_handshake_received":             ret.DtDdosL7QuicStats.Stats.Quic_handshake_received,
			"quic_version_match_action_taken":     ret.DtDdosL7QuicStats.Stats.Quic_version_match_action_taken,
			"quic_version_match_action_drop":      ret.DtDdosL7QuicStats.Stats.Quic_version_match_action_drop,
			"quic_version_match_action_blacklist": ret.DtDdosL7QuicStats.Stats.Quic_version_match_action_blacklist,
			"quic_malformed_action_taken":         ret.DtDdosL7QuicStats.Stats.Quic_malformed_action_taken,
			"quic_malformed_action_drop":          ret.DtDdosL7QuicStats.Stats.Quic_malformed_action_drop,
			"quic_malformed_action_blacklist":     ret.DtDdosL7QuicStats.Stats.Quic_malformed_action_blacklist,
			"quic_malformed_dcid_len_max_exceed":  ret.DtDdosL7QuicStats.Stats.Quic_malformed_dcid_len_max_exceed,
			"quic_malformed_scid_len_max_exceed":  ret.DtDdosL7QuicStats.Stats.Quic_malformed_scid_len_max_exceed,
			"quic_fixed_bit_not_set":              ret.DtDdosL7QuicStats.Stats.Quic_fixed_bit_not_set,
			"quic_retry_auth_sent":                ret.DtDdosL7QuicStats.Stats.Quic_retry_auth_sent,
			"quic_retry_auth_pass":                ret.DtDdosL7QuicStats.Stats.Quic_retry_auth_pass,
			"quic_retry_auth_fail":                ret.DtDdosL7QuicStats.Stats.Quic_retry_auth_fail,
			"quic_connection_close_sent":          ret.DtDdosL7QuicStats.Stats.Quic_connection_close_sent,
			"quic_invalid_retry_token":            ret.DtDdosL7QuicStats.Stats.Quic_invalid_retry_token,
			"quic_short_header_received":          ret.DtDdosL7QuicStats.Stats.Quic_short_header_received,
			"quic_short_header_action_drop":       ret.DtDdosL7QuicStats.Stats.Quic_short_header_action_drop,
			"quic_encrypt_fail":                   ret.DtDdosL7QuicStats.Stats.Quic_encrypt_fail,
			"quic_decrypt_fail":                   ret.DtDdosL7QuicStats.Stats.Quic_decrypt_fail,
			"quic_encrypt_success":                ret.DtDdosL7QuicStats.Stats.Quic_encrypt_success,
			"quic_decrypt_success":                ret.DtDdosL7QuicStats.Stats.Quic_decrypt_success,
			"quic_0rtt_drop":                      ret.DtDdosL7QuicStats.Stats.Quic_0rtt_drop,
			"quic_aead_pkt_rate_exceed":           ret.DtDdosL7QuicStats.Stats.Quic_aead_pkt_rate_exceed,
			"quic_dcid_pkt_rate_exceed":           ret.DtDdosL7QuicStats.Stats.Quic_dcid_pkt_rate_exceed,
			"quic_create_conn_init_only":          ret.DtDdosL7QuicStats.Stats.Quic_create_conn_init_only,
			"quic_version_no_match_drop":          ret.DtDdosL7QuicStats.Stats.Quic_version_no_match_drop,
		},
	}
}

func getObjectDdosL7QuicStatsStats(d []interface{}) edpt.DdosL7QuicStatsStats {

	count1 := len(d)
	var ret edpt.DdosL7QuicStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic_packet_received = in["quic_packet_received"].(int)
		ret.Quic_initial_received = in["quic_initial_received"].(int)
		ret.Quic_version_negotiation_received = in["quic_version_negotiation_received"].(int)
		ret.Quic_retry_received = in["quic_retry_received"].(int)
		ret.Quic_0rtt_recevied = in["quic_0rtt_recevied"].(int)
		ret.Quic_handshake_received = in["quic_handshake_received"].(int)
		ret.Quic_version_match_action_taken = in["quic_version_match_action_taken"].(int)
		ret.Quic_version_match_action_drop = in["quic_version_match_action_drop"].(int)
		ret.Quic_version_match_action_blacklist = in["quic_version_match_action_blacklist"].(int)
		ret.Quic_malformed_action_taken = in["quic_malformed_action_taken"].(int)
		ret.Quic_malformed_action_drop = in["quic_malformed_action_drop"].(int)
		ret.Quic_malformed_action_blacklist = in["quic_malformed_action_blacklist"].(int)
		ret.Quic_malformed_dcid_len_max_exceed = in["quic_malformed_dcid_len_max_exceed"].(int)
		ret.Quic_malformed_scid_len_max_exceed = in["quic_malformed_scid_len_max_exceed"].(int)
		ret.Quic_fixed_bit_not_set = in["quic_fixed_bit_not_set"].(int)
		ret.Quic_retry_auth_sent = in["quic_retry_auth_sent"].(int)
		ret.Quic_retry_auth_pass = in["quic_retry_auth_pass"].(int)
		ret.Quic_retry_auth_fail = in["quic_retry_auth_fail"].(int)
		ret.Quic_connection_close_sent = in["quic_connection_close_sent"].(int)
		ret.Quic_invalid_retry_token = in["quic_invalid_retry_token"].(int)
		ret.Quic_short_header_received = in["quic_short_header_received"].(int)
		ret.Quic_short_header_action_drop = in["quic_short_header_action_drop"].(int)
		ret.Quic_encrypt_fail = in["quic_encrypt_fail"].(int)
		ret.Quic_decrypt_fail = in["quic_decrypt_fail"].(int)
		ret.Quic_encrypt_success = in["quic_encrypt_success"].(int)
		ret.Quic_decrypt_success = in["quic_decrypt_success"].(int)
		ret.Quic_0rtt_drop = in["quic_0rtt_drop"].(int)
		ret.Quic_aead_pkt_rate_exceed = in["quic_aead_pkt_rate_exceed"].(int)
		ret.Quic_dcid_pkt_rate_exceed = in["quic_dcid_pkt_rate_exceed"].(int)
		ret.Quic_create_conn_init_only = in["quic_create_conn_init_only"].(int)
		ret.Quic_version_no_match_drop = in["quic_version_no_match_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosL7QuicStats(d *schema.ResourceData) edpt.DdosL7QuicStats {
	var ret edpt.DdosL7QuicStats

	ret.Stats = getObjectDdosL7QuicStatsStats(d.Get("stats").([]interface{}))
	return ret
}
