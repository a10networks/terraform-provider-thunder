package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSctpGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sctp_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceSctpGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sctp_static_nat_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Static NAT Session Created",
						},
						"sctp_static_nat_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Static NAT Session Deleted",
						},
						"sctp_fw_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Firewall Session Created",
						},
						"sctp_fw_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Firewall Session Deleted",
						},
						"pkt_err_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Error Drop",
						},
						"bad_csum": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Checksum",
						},
						"bad_payload_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Payload Drop",
						},
						"bad_alignment_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Alignment Drop",
						},
						"oos_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Out-of-state Packet Drop",
						},
						"max_multi_home_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum Multi-homing IP Addresses Drop",
						},
						"multi_home_remove_ip_skip": {
							Type: schema.TypeInt, Optional: true, Description: "Multi-homing Remove IP Parameter Skip",
						},
						"multi_home_addr_not_found_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Multi-homing IP Address Not Found Drop",
						},
						"static_nat_cfg_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Static NAT Config Not Found Drop",
						},
						"cfg_err_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Configuration Error Drop",
						},
						"vrrp_standby_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NAT Resource VRRP-A Standby Drop",
						},
						"invalid_frag_chunk_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Fragmented Chunks Drop",
						},
						"disallowed_chunk_filtered": {
							Type: schema.TypeInt, Optional: true, Description: "Disallowed Chunk Filtered",
						},
						"disallowed_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Disallowed Packet Drop",
						},
						"rate_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop",
						},
						"sby_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Standby Session Created",
						},
						"sby_session_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Standby Session Create Failed",
						},
						"sby_session_updated": {
							Type: schema.TypeInt, Optional: true, Description: "Standby Session Updated",
						},
						"sby_session_update_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Standby Session Update Failed",
						},
						"sby_static_nat_cfg_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Static NAT Config Not Found on Standby",
						},
						"sctp_chunk_type_init": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type INIT",
						},
						"sctp_chunk_type_init_ack": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type INIT-ACK",
						},
						"sctp_chunk_type_cookie_echo": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type COOKIE-ECHO",
						},
						"sctp_chunk_type_cookie_ack": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type COOKIE-ACK",
						},
						"sctp_chunk_type_sack": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type SACK",
						},
						"sctp_chunk_type_asconf": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type ASCONF",
						},
						"sctp_chunk_type_asconf_ack": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type ASCONF-ACK",
						},
						"sctp_chunk_type_data": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type DATA",
						},
						"sctp_chunk_type_abort": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type ABORT",
						},
						"sctp_chunk_type_shutdown": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type SHUTDOWN",
						},
						"sctp_chunk_type_shutdown_ack": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type SHUTDOWN-ACK",
						},
						"sctp_chunk_type_shutdown_complete": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type SHUTDOWN-COMPLETE",
						},
						"sctp_chunk_type_error_op": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type ERROR-OP",
						},
						"sctp_chunk_type_heartbeat": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type HEARTBEAT",
						},
						"sctp_chunk_type_heartbeat_ack": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type HEARTBEAT-ACK",
						},
						"sctp_chunk_type_other": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type OTHER",
						},
						"sctp_heartbeat_no_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP HEARTBEAT/ACK no tuple found",
						},
						"sctp_data_no_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP DATA chunk no tuple found",
						},
						"sctp_data_no_ext_match": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP DATA no extended match found",
						},
						"sctp_chunk_type_init_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type INIT drop",
						},
						"sctp_chunk_type_init_ack_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type INIT-ACK drop",
						},
						"sctp_chunk_type_shutdown_complete_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type SHUTDOWN-COMPLETE drop",
						},
						"sctp_chunk_type_abort_data_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type with ABORT and DATA drop",
						},
						"sctp_chunk_heart_beat_clubbed": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP HEARTBEAT chunk with other chunk",
						},
						"sctp_retx_init_ack_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type INIT_ACK with retx mismatched vtag drop",
						},
						"sctp_route_err_heartbeat_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP HEARTBEAT ROUTE lookup failed drop",
						},
						"sctp_reroute_failover": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP REROUTE lookup for chunks other than HEARTBEAT",
						},
						"sctp_route_err_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP ROUTE lookup failed for chunks other than HEARTBEAT drop",
						},
						"sctp_no_ext_match": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP no extended match found",
						},
						"sctp_retx_init_ack": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Chunk Type INIT_ACK retransmitted",
						},
						"sctp_retx_init_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Retransmitted INIT drop",
						},
						"sctp_retx_init": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Retransmitted INIT",
						},
						"sctp_asconf_process_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP ASCONF process drop",
						},
						"sctp_init_vtag_zero_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP INIT VTAG ZERO drop",
						},
						"pkt_len_err_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Packet Length Drop",
						},
						"pkt_chunk_len_err_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Chunk Length Drop",
						},
						"pkt_asconf_param_len_err_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Parameter Length Drop",
						},
					},
				},
			},
		},
	}
}

func resourceSctpGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSctpGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSctpGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SctpGlobalStatsStats := setObjectSctpGlobalStatsStats(res)
		d.Set("stats", SctpGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSctpGlobalStatsStats(ret edpt.DataSctpGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sctp_static_nat_session_created":        ret.DtSctpGlobalStats.Stats.SctpStaticNatSessionCreated,
			"sctp_static_nat_session_deleted":        ret.DtSctpGlobalStats.Stats.SctpStaticNatSessionDeleted,
			"sctp_fw_session_created":                ret.DtSctpGlobalStats.Stats.SctpFwSessionCreated,
			"sctp_fw_session_deleted":                ret.DtSctpGlobalStats.Stats.SctpFwSessionDeleted,
			"pkt_err_drop":                           ret.DtSctpGlobalStats.Stats.PktErrDrop,
			"bad_csum":                               ret.DtSctpGlobalStats.Stats.BadCsum,
			"bad_payload_drop":                       ret.DtSctpGlobalStats.Stats.BadPayloadDrop,
			"bad_alignment_drop":                     ret.DtSctpGlobalStats.Stats.BadAlignmentDrop,
			"oos_pkt_drop":                           ret.DtSctpGlobalStats.Stats.OosPktDrop,
			"max_multi_home_drop":                    ret.DtSctpGlobalStats.Stats.MaxMultiHomeDrop,
			"multi_home_remove_ip_skip":              ret.DtSctpGlobalStats.Stats.MultiHomeRemoveIpSkip,
			"multi_home_addr_not_found_drop":         ret.DtSctpGlobalStats.Stats.MultiHomeAddrNotFoundDrop,
			"static_nat_cfg_not_found":               ret.DtSctpGlobalStats.Stats.StaticNatCfgNotFound,
			"cfg_err_drop":                           ret.DtSctpGlobalStats.Stats.CfgErrDrop,
			"vrrp_standby_drop":                      ret.DtSctpGlobalStats.Stats.VrrpStandbyDrop,
			"invalid_frag_chunk_drop":                ret.DtSctpGlobalStats.Stats.InvalidFragChunkDrop,
			"disallowed_chunk_filtered":              ret.DtSctpGlobalStats.Stats.DisallowedChunkFiltered,
			"disallowed_pkt_drop":                    ret.DtSctpGlobalStats.Stats.DisallowedPktDrop,
			"rate_limit_drop":                        ret.DtSctpGlobalStats.Stats.RateLimitDrop,
			"sby_session_created":                    ret.DtSctpGlobalStats.Stats.SbySessionCreated,
			"sby_session_create_fail":                ret.DtSctpGlobalStats.Stats.SbySessionCreateFail,
			"sby_session_updated":                    ret.DtSctpGlobalStats.Stats.SbySessionUpdated,
			"sby_session_update_fail":                ret.DtSctpGlobalStats.Stats.SbySessionUpdateFail,
			"sby_static_nat_cfg_not_found":           ret.DtSctpGlobalStats.Stats.SbyStaticNatCfgNotFound,
			"sctp_chunk_type_init":                   ret.DtSctpGlobalStats.Stats.SctpChunkTypeInit,
			"sctp_chunk_type_init_ack":               ret.DtSctpGlobalStats.Stats.SctpChunkTypeInitAck,
			"sctp_chunk_type_cookie_echo":            ret.DtSctpGlobalStats.Stats.SctpChunkTypeCookieEcho,
			"sctp_chunk_type_cookie_ack":             ret.DtSctpGlobalStats.Stats.SctpChunkTypeCookieAck,
			"sctp_chunk_type_sack":                   ret.DtSctpGlobalStats.Stats.SctpChunkTypeSack,
			"sctp_chunk_type_asconf":                 ret.DtSctpGlobalStats.Stats.SctpChunkTypeAsconf,
			"sctp_chunk_type_asconf_ack":             ret.DtSctpGlobalStats.Stats.SctpChunkTypeAsconfAck,
			"sctp_chunk_type_data":                   ret.DtSctpGlobalStats.Stats.SctpChunkTypeData,
			"sctp_chunk_type_abort":                  ret.DtSctpGlobalStats.Stats.SctpChunkTypeAbort,
			"sctp_chunk_type_shutdown":               ret.DtSctpGlobalStats.Stats.SctpChunkTypeShutdown,
			"sctp_chunk_type_shutdown_ack":           ret.DtSctpGlobalStats.Stats.SctpChunkTypeShutdownAck,
			"sctp_chunk_type_shutdown_complete":      ret.DtSctpGlobalStats.Stats.SctpChunkTypeShutdownComplete,
			"sctp_chunk_type_error_op":               ret.DtSctpGlobalStats.Stats.SctpChunkTypeErrorOp,
			"sctp_chunk_type_heartbeat":              ret.DtSctpGlobalStats.Stats.SctpChunkTypeHeartbeat,
			"sctp_chunk_type_heartbeat_ack":          ret.DtSctpGlobalStats.Stats.SctpChunkTypeHeartbeatAck,
			"sctp_chunk_type_other":                  ret.DtSctpGlobalStats.Stats.SctpChunkTypeOther,
			"sctp_heartbeat_no_tuple":                ret.DtSctpGlobalStats.Stats.SctpHeartbeatNoTuple,
			"sctp_data_no_tuple":                     ret.DtSctpGlobalStats.Stats.SctpDataNoTuple,
			"sctp_data_no_ext_match":                 ret.DtSctpGlobalStats.Stats.SctpDataNoExtMatch,
			"sctp_chunk_type_init_drop":              ret.DtSctpGlobalStats.Stats.SctpChunkTypeInitDrop,
			"sctp_chunk_type_init_ack_drop":          ret.DtSctpGlobalStats.Stats.SctpChunkTypeInitAckDrop,
			"sctp_chunk_type_shutdown_complete_drop": ret.DtSctpGlobalStats.Stats.SctpChunkTypeShutdownCompleteDrop,
			"sctp_chunk_type_abort_data_drop":        ret.DtSctpGlobalStats.Stats.SctpChunkTypeAbortDataDrop,
			"sctp_chunk_heart_beat_clubbed":          ret.DtSctpGlobalStats.Stats.SctpChunkHeartBeatClubbed,
			"sctp_retx_init_ack_drop":                ret.DtSctpGlobalStats.Stats.SctpRetxInitAckDrop,
			"sctp_route_err_heartbeat_drop":          ret.DtSctpGlobalStats.Stats.SctpRouteErrHeartbeatDrop,
			"sctp_reroute_failover":                  ret.DtSctpGlobalStats.Stats.SctpRerouteFailover,
			"sctp_route_err_drop":                    ret.DtSctpGlobalStats.Stats.SctpRouteErrDrop,
			"sctp_no_ext_match":                      ret.DtSctpGlobalStats.Stats.SctpNoExtMatch,
			"sctp_retx_init_ack":                     ret.DtSctpGlobalStats.Stats.SctpRetxInitAck,
			"sctp_retx_init_drop":                    ret.DtSctpGlobalStats.Stats.SctpRetxInitDrop,
			"sctp_retx_init":                         ret.DtSctpGlobalStats.Stats.SctpRetxInit,
			"sctp_asconf_process_drop":               ret.DtSctpGlobalStats.Stats.SctpAsconfProcessDrop,
			"sctp_init_vtag_zero_drop":               ret.DtSctpGlobalStats.Stats.SctpInitVtagZeroDrop,
			"pkt_len_err_drop":                       ret.DtSctpGlobalStats.Stats.PktLenErrDrop,
			"pkt_chunk_len_err_drop":                 ret.DtSctpGlobalStats.Stats.PktChunkLenErrDrop,
			"pkt_asconf_param_len_err_drop":          ret.DtSctpGlobalStats.Stats.PktAsconfParamLenErrDrop,
		},
	}
}

func getObjectSctpGlobalStatsStats(d []interface{}) edpt.SctpGlobalStatsStats {

	count1 := len(d)
	var ret edpt.SctpGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SctpStaticNatSessionCreated = in["sctp_static_nat_session_created"].(int)
		ret.SctpStaticNatSessionDeleted = in["sctp_static_nat_session_deleted"].(int)
		ret.SctpFwSessionCreated = in["sctp_fw_session_created"].(int)
		ret.SctpFwSessionDeleted = in["sctp_fw_session_deleted"].(int)
		ret.PktErrDrop = in["pkt_err_drop"].(int)
		ret.BadCsum = in["bad_csum"].(int)
		ret.BadPayloadDrop = in["bad_payload_drop"].(int)
		ret.BadAlignmentDrop = in["bad_alignment_drop"].(int)
		ret.OosPktDrop = in["oos_pkt_drop"].(int)
		ret.MaxMultiHomeDrop = in["max_multi_home_drop"].(int)
		ret.MultiHomeRemoveIpSkip = in["multi_home_remove_ip_skip"].(int)
		ret.MultiHomeAddrNotFoundDrop = in["multi_home_addr_not_found_drop"].(int)
		ret.StaticNatCfgNotFound = in["static_nat_cfg_not_found"].(int)
		ret.CfgErrDrop = in["cfg_err_drop"].(int)
		ret.VrrpStandbyDrop = in["vrrp_standby_drop"].(int)
		ret.InvalidFragChunkDrop = in["invalid_frag_chunk_drop"].(int)
		ret.DisallowedChunkFiltered = in["disallowed_chunk_filtered"].(int)
		ret.DisallowedPktDrop = in["disallowed_pkt_drop"].(int)
		ret.RateLimitDrop = in["rate_limit_drop"].(int)
		ret.SbySessionCreated = in["sby_session_created"].(int)
		ret.SbySessionCreateFail = in["sby_session_create_fail"].(int)
		ret.SbySessionUpdated = in["sby_session_updated"].(int)
		ret.SbySessionUpdateFail = in["sby_session_update_fail"].(int)
		ret.SbyStaticNatCfgNotFound = in["sby_static_nat_cfg_not_found"].(int)
		ret.SctpChunkTypeInit = in["sctp_chunk_type_init"].(int)
		ret.SctpChunkTypeInitAck = in["sctp_chunk_type_init_ack"].(int)
		ret.SctpChunkTypeCookieEcho = in["sctp_chunk_type_cookie_echo"].(int)
		ret.SctpChunkTypeCookieAck = in["sctp_chunk_type_cookie_ack"].(int)
		ret.SctpChunkTypeSack = in["sctp_chunk_type_sack"].(int)
		ret.SctpChunkTypeAsconf = in["sctp_chunk_type_asconf"].(int)
		ret.SctpChunkTypeAsconfAck = in["sctp_chunk_type_asconf_ack"].(int)
		ret.SctpChunkTypeData = in["sctp_chunk_type_data"].(int)
		ret.SctpChunkTypeAbort = in["sctp_chunk_type_abort"].(int)
		ret.SctpChunkTypeShutdown = in["sctp_chunk_type_shutdown"].(int)
		ret.SctpChunkTypeShutdownAck = in["sctp_chunk_type_shutdown_ack"].(int)
		ret.SctpChunkTypeShutdownComplete = in["sctp_chunk_type_shutdown_complete"].(int)
		ret.SctpChunkTypeErrorOp = in["sctp_chunk_type_error_op"].(int)
		ret.SctpChunkTypeHeartbeat = in["sctp_chunk_type_heartbeat"].(int)
		ret.SctpChunkTypeHeartbeatAck = in["sctp_chunk_type_heartbeat_ack"].(int)
		ret.SctpChunkTypeOther = in["sctp_chunk_type_other"].(int)
		ret.SctpHeartbeatNoTuple = in["sctp_heartbeat_no_tuple"].(int)
		ret.SctpDataNoTuple = in["sctp_data_no_tuple"].(int)
		ret.SctpDataNoExtMatch = in["sctp_data_no_ext_match"].(int)
		ret.SctpChunkTypeInitDrop = in["sctp_chunk_type_init_drop"].(int)
		ret.SctpChunkTypeInitAckDrop = in["sctp_chunk_type_init_ack_drop"].(int)
		ret.SctpChunkTypeShutdownCompleteDrop = in["sctp_chunk_type_shutdown_complete_drop"].(int)
		ret.SctpChunkTypeAbortDataDrop = in["sctp_chunk_type_abort_data_drop"].(int)
		ret.SctpChunkHeartBeatClubbed = in["sctp_chunk_heart_beat_clubbed"].(int)
		ret.SctpRetxInitAckDrop = in["sctp_retx_init_ack_drop"].(int)
		ret.SctpRouteErrHeartbeatDrop = in["sctp_route_err_heartbeat_drop"].(int)
		ret.SctpRerouteFailover = in["sctp_reroute_failover"].(int)
		ret.SctpRouteErrDrop = in["sctp_route_err_drop"].(int)
		ret.SctpNoExtMatch = in["sctp_no_ext_match"].(int)
		ret.SctpRetxInitAck = in["sctp_retx_init_ack"].(int)
		ret.SctpRetxInitDrop = in["sctp_retx_init_drop"].(int)
		ret.SctpRetxInit = in["sctp_retx_init"].(int)
		ret.SctpAsconfProcessDrop = in["sctp_asconf_process_drop"].(int)
		ret.SctpInitVtagZeroDrop = in["sctp_init_vtag_zero_drop"].(int)
		ret.PktLenErrDrop = in["pkt_len_err_drop"].(int)
		ret.PktChunkLenErrDrop = in["pkt_chunk_len_err_drop"].(int)
		ret.PktAsconfParamLenErrDrop = in["pkt_asconf_param_len_err_drop"].(int)
	}
	return ret
}

func dataToEndpointSctpGlobalStats(d *schema.ResourceData) edpt.SctpGlobalStats {
	var ret edpt.SctpGlobalStats

	ret.Stats = getObjectSctpGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
