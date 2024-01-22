package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbQuicStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_quic_stats`: Statistics for the object quic\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbQuicStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_conn_attempted": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection attempted",
						},
						"client_conn_handshake": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection handshake",
						},
						"client_conn_created": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection created",
						},
						"client_conn_local_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection local closed",
						},
						"client_conn_remote_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection remote closed",
						},
						"client_conn_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection failed",
						},
						"server_conn_attempted": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection attempted",
						},
						"server_conn_handshake": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection handshake",
						},
						"server_conn_created": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection created",
						},
						"server_conn_local_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection local closed",
						},
						"server_conn_remote_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection remote closed",
						},
						"server_conn_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection failed",
						},
						"q_conn_created": {
							Type: schema.TypeInt, Optional: true, Description: "Q connection created",
						},
						"q_conn_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Q connection freed",
						},
						"local_bi_stream_current": {
							Type: schema.TypeInt, Optional: true, Description: "Current local bi-stream",
						},
						"remote_bi_stream_current": {
							Type: schema.TypeInt, Optional: true, Description: "Current remote bi-stream",
						},
						"local_bi_stream_created": {
							Type: schema.TypeInt, Optional: true, Description: "Local bi-stream created",
						},
						"remote_bi_stream_created": {
							Type: schema.TypeInt, Optional: true, Description: "Remote bi-stream created",
						},
						"local_bi_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Local bi-stream closed",
						},
						"remote_bi_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Remote bi-stream closed",
						},
						"local_uni_stream_current": {
							Type: schema.TypeInt, Optional: true, Description: "Current local uni-stream",
						},
						"remote_uni_stream_current": {
							Type: schema.TypeInt, Optional: true, Description: "Current remote uni-stream",
						},
						"local_uni_stream_created": {
							Type: schema.TypeInt, Optional: true, Description: "Local uni-stream created",
						},
						"remote_uni_stream_created": {
							Type: schema.TypeInt, Optional: true, Description: "Remote uni-stream created",
						},
						"local_uni_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Local uni-stream closed",
						},
						"remote_uni_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Remote uni-stream closed",
						},
						"stream_error": {
							Type: schema.TypeInt, Optional: true, Description: "Stream error",
						},
						"stream_fail_to_insert": {
							Type: schema.TypeInt, Optional: true, Description: "Stream fail to insert",
						},
						"padding_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "padding frame receive",
						},
						"padding_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "padding frame send",
						},
						"ping_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "ping frame receive",
						},
						"ping_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "ping frame send",
						},
						"ack_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "ack frame receive",
						},
						"ack_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "ack frame send",
						},
						"ack_ecn_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "ack enc frame receive",
						},
						"ack_ecn_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "ack enc frame send",
						},
						"stream_rst_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream reset frame receive",
						},
						"stream_rst_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream reset frame send",
						},
						"stream_stop_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream stop frame receive",
						},
						"stream_stop_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream stop frame send",
						},
						"crypto_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "crypto frame receive",
						},
						"crypto_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "crypto frame send",
						},
						"new_token_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "new token frame receive",
						},
						"new_token_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "new token frame send",
						},
						"stream_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream frame receive",
						},
						"stream_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream frame send",
						},
						"stream_09_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 09 frame receive",
						},
						"stream_09_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 09 frame send",
						},
						"stream_0a_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0a frame receive",
						},
						"stream_0a_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0a frame send",
						},
						"stream_0b_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0b frame receive",
						},
						"stream_0b_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0b frame send",
						},
						"stream_0c_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0c frame receive",
						},
						"stream_0c_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0c frame send",
						},
						"stream_0d_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0d frame receive",
						},
						"stream_0d_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0d frame send",
						},
						"stream_0e_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0e frame receive",
						},
						"stream_0e_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0e frame send",
						},
						"stream_0f_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0f frame receive",
						},
						"stream_0f_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream 0f frame send",
						},
						"max_data_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "max data frame receive",
						},
						"max_data_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "max data frame send",
						},
						"max_stream_data_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "max stream data frame receive",
						},
						"max_stream_data_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "max stream data frame send",
						},
						"max_bi_stream_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "max bi stream frame receive",
						},
						"max_bi_stream_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "max bi stream frame send",
						},
						"max_uni_stream_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "max uni stream frame receive",
						},
						"max_uni_stream_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "max uni stream frame send",
						},
						"data_blocked_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "data blocked frame receive",
						},
						"data_blocked_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "data blocked frame send",
						},
						"stream_data_blocked_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "stream data blocked frame receive",
						},
						"stream_data_blocked_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "stream data blocked frame send",
						},
						"bi_stream_data_blocked_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "bi stream data blocked frame receive",
						},
						"bi_stream_data_blocked_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "bi stream data blocked frame send",
						},
						"uni_stream_data_blocked_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "uni stream data blocked frame receive",
						},
						"uni_stream_data_blocked_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "uni stream data blocked frame send",
						},
						"new_conn_id_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "new conn id frame receive",
						},
						"new_conn_id_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "new conn id frame send",
						},
						"retire_conn_id_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "retire conn id frame receive",
						},
						"retire_conn_id_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "retire conn id frame send",
						},
						"path_challenge_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "path challenge frame receive",
						},
						"path_challenge_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "path challenge frame send",
						},
						"path_response_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "path response frame receive",
						},
						"path_response_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "path response frame send",
						},
						"conn_close_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "conn close frame receive",
						},
						"conn_close_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "conn close frame send",
						},
						"app_conn_close_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "app conn close frame receive",
						},
						"app_conn_close_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "app conn close frame send",
						},
						"handshake_done_frame_rx": {
							Type: schema.TypeInt, Optional: true, Description: "handshake done frame receive",
						},
						"handshake_done_frame_tx": {
							Type: schema.TypeInt, Optional: true, Description: "handshake done frame send",
						},
						"unknown_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown frame",
						},
						"stream_fin_receive": {
							Type: schema.TypeInt, Optional: true, Description: "Stream FIN receive",
						},
						"stream_fin_up": {
							Type: schema.TypeInt, Optional: true, Description: "Stream FIN up",
						},
						"stream_fin_down": {
							Type: schema.TypeInt, Optional: true, Description: "Stream FIN down",
						},
						"stream_fin_send": {
							Type: schema.TypeInt, Optional: true, Description: "Stream FIN send",
						},
						"stream_congest": {
							Type: schema.TypeInt, Optional: true, Description: "Stream congest",
						},
						"stream_open": {
							Type: schema.TypeInt, Optional: true, Description: "Stream open",
						},
						"stream_pause_data": {
							Type: schema.TypeInt, Optional: true, Description: "Stream pause data",
						},
						"stream_resume_data": {
							Type: schema.TypeInt, Optional: true, Description: "Stream resume data",
						},
						"stream_not_send": {
							Type: schema.TypeInt, Optional: true, Description: "Stream not send",
						},
						"stream_stop_send": {
							Type: schema.TypeInt, Optional: true, Description: "Stream stop send",
						},
						"stream_created": {
							Type: schema.TypeInt, Optional: true, Description: "Stream created",
						},
						"stream_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Stream freed",
						},
						"initial_rx": {
							Type: schema.TypeInt, Optional: true, Description: "INITIAL receive",
						},
						"initial_tx": {
							Type: schema.TypeInt, Optional: true, Description: "INITIAL send",
						},
						"rtt_0_rx": {
							Type: schema.TypeInt, Optional: true, Description: "RTT_0 receive",
						},
						"rtt_0_tx": {
							Type: schema.TypeInt, Optional: true, Description: "RTT_0 send",
						},
						"handshake_rx": {
							Type: schema.TypeInt, Optional: true, Description: "HANDSHAKE receive",
						},
						"handshake_tx": {
							Type: schema.TypeInt, Optional: true, Description: "HANDSHAKE send",
						},
						"retry_rx": {
							Type: schema.TypeInt, Optional: true, Description: "RETRY receive",
						},
						"retry_tx": {
							Type: schema.TypeInt, Optional: true, Description: "RETRY send",
						},
						"ver_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Version receive",
						},
						"ver_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Version send",
						},
						"rtt_updated": {
							Type: schema.TypeInt, Optional: true, Description: "RTT updated",
						},
						"needs_ack": {
							Type: schema.TypeInt, Optional: true, Description: "Needs ACK",
						},
						"delayed_ack": {
							Type: schema.TypeInt, Optional: true, Description: "Delayed ACK",
						},
						"packet_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Packet receive",
						},
						"packet_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Packet send",
						},
						"packet_tx_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Packet send failed",
						},
						"congest_wnd_inc": {
							Type: schema.TypeInt, Optional: true, Description: "Congestion window increase",
						},
						"congest_wnd_dec": {
							Type: schema.TypeInt, Optional: true, Description: "Congestion window decrease",
						},
						"no_congest_wnd": {
							Type: schema.TypeInt, Optional: true, Description: "No congestion window",
						},
						"burst_limited": {
							Type: schema.TypeInt, Optional: true, Description: "Burst limited",
						},
						"packet_loop_limited": {
							Type: schema.TypeInt, Optional: true, Description: "Packet loop limited",
						},
						"receive_wnd_limited": {
							Type: schema.TypeInt, Optional: true, Description: "Receive window limited",
						},
						"parse_error": {
							Type: schema.TypeInt, Optional: true, Description: "Parse error",
						},
						"error_close": {
							Type: schema.TypeInt, Optional: true, Description: "Conn closed of error",
						},
						"unknown_scid": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown scid",
						},
						"dcid_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Dcid mismatch",
						},
						"packet_too_short": {
							Type: schema.TypeInt, Optional: true, Description: "Packet_too_short",
						},
						"invalid_version": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid version",
						},
						"invalid_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid packet",
						},
						"invalid_conn_match": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid conn match",
						},
						"invalid_session_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid session packet",
						},
						"stateless_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Stateless resert",
						},
						"packet_lost": {
							Type: schema.TypeInt, Optional: true, Description: "Packet lost",
						},
						"packet_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Packet drop",
						},
						"packet_retransmit": {
							Type: schema.TypeInt, Optional: true, Description: "Packet retransmit",
						},
						"packet_out_of_order": {
							Type: schema.TypeInt, Optional: true, Description: "Packet out of order",
						},
						"quic_packet_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Quic packet drop",
						},
						"encode_error": {
							Type: schema.TypeInt, Optional: true, Description: "Encode error",
						},
						"decode_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Decode failed",
						},
						"decode_stream_error": {
							Type: schema.TypeInt, Optional: true, Description: "Decode stream error",
						},
						"exceed_flow_control": {
							Type: schema.TypeInt, Optional: true, Description: "Exceed flow control",
						},
						"crypto_stream_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Crypto stream not found",
						},
						"exceed_max_stream_id": {
							Type: schema.TypeInt, Optional: true, Description: "Exceed_max_stream_id",
						},
						"stream_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Stream_id_mismatch",
						},
						"ack_delay_huge": {
							Type: schema.TypeInt, Optional: true, Description: "Ack_delay_huge",
						},
						"ack_rng_huge_1": {
							Type: schema.TypeInt, Optional: true, Description: "Ack_rng_huge_1",
						},
						"ack_rng_huge_2": {
							Type: schema.TypeInt, Optional: true, Description: "Ack_rng_huge_2",
						},
						"ack_rng_huge_3": {
							Type: schema.TypeInt, Optional: true, Description: "Ack_rng_huge_3",
						},
						"too_noisy_fuzzing": {
							Type: schema.TypeInt, Optional: true, Description: "Too_noisy_fuzzing",
						},
						"max_stream_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Max_stream_too_big",
						},
						"stream_blocked": {
							Type: schema.TypeInt, Optional: true, Description: "Stream_blocked",
						},
						"new_conn_id_len_zero": {
							Type: schema.TypeInt, Optional: true, Description: "New_conn_id_len_zero",
						},
						"new_conn_id_len_non_zero": {
							Type: schema.TypeInt, Optional: true, Description: "New_conn_id_len_non_zero",
						},
						"illegal_stream_len": {
							Type: schema.TypeInt, Optional: true, Description: "Illegal_stream_len",
						},
						"illegal_reason_len": {
							Type: schema.TypeInt, Optional: true, Description: "Illegal_reason_len",
						},
						"illegal_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Illegal_seq",
						},
						"illegal_rpt": {
							Type: schema.TypeInt, Optional: true, Description: "Illegal_rpt",
						},
						"illegal_len": {
							Type: schema.TypeInt, Optional: true, Description: "Illegal_len",
						},
						"illegal_token_len": {
							Type: schema.TypeInt, Optional: true, Description: "Illegal_token_len",
						},
						"cannot_insert_cid": {
							Type: schema.TypeInt, Optional: true, Description: "Cannot_insert_cid",
						},
						"cannot_insert_srt": {
							Type: schema.TypeInt, Optional: true, Description: "Cannot_insert_srt",
						},
						"cannot_retire_cid": {
							Type: schema.TypeInt, Optional: true, Description: "Cannot_retire_cid",
						},
						"no_next_scid": {
							Type: schema.TypeInt, Optional: true, Description: "No_next_scid",
						},
						"token_len_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Token_len_too_long",
						},
						"server_receive_new_token": {
							Type: schema.TypeInt, Optional: true, Description: "Server_receive_new_token",
						},
						"zero_frame_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Zero_frame_packet",
						},
						"err_frame_dec1": {
							Type: schema.TypeInt, Optional: true, Description: "Err_frame_dec1",
						},
						"err_frame_dec": {
							Type: schema.TypeInt, Optional: true, Description: "Err_frame_dec",
						},
						"err_frame_decb": {
							Type: schema.TypeInt, Optional: true, Description: "Err_frame_decb",
						},
						"err_frame_final_size": {
							Type: schema.TypeInt, Optional: true, Description: "Err_frame_final_size",
						},
						"err_flow_control": {
							Type: schema.TypeInt, Optional: true, Description: "Err_flow_control",
						},
						"err_protocol_violation": {
							Type: schema.TypeInt, Optional: true, Description: "Err_protocol_violation",
						},
						"server_rx_handshake_done": {
							Type: schema.TypeInt, Optional: true, Description: "Server_rx_handshake_done",
						},
						"pkt_acked_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt_acked_failed",
						},
						"pn_insert_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Pn insert failed",
						},
						"pn_delete_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Pn delete failed",
						},
						"acked_packet_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Acked packet freed",
						},
						"tx_buffer_enq": {
							Type: schema.TypeInt, Optional: true, Description: "Tx buffer enqueued",
						},
						"tx_buffer_deq": {
							Type: schema.TypeInt, Optional: true, Description: "Tx buffer dequeued",
						},
						"app_buffer_enq": {
							Type: schema.TypeInt, Optional: true, Description: "App buffer enqueued",
						},
						"app_buffer_deq": {
							Type: schema.TypeInt, Optional: true, Description: "App buffer dequeued",
						},
						"app_buffer_queue_full": {
							Type: schema.TypeInt, Optional: true, Description: "App buffer queue full",
						},
						"iov_buffer_bind": {
							Type: schema.TypeInt, Optional: true, Description: "Iov buffer bind",
						},
						"iov_buffer_unbind": {
							Type: schema.TypeInt, Optional: true, Description: "Iov buffer unbind",
						},
						"iov_buffer_dup": {
							Type: schema.TypeInt, Optional: true, Description: "Iov buffer dup",
						},
						"iov_alloc_len": {
							Type: schema.TypeInt, Optional: true, Description: "Iov alloc len",
						},
						"iov_io": {
							Type: schema.TypeInt, Optional: true, Description: "Iov IO",
						},
						"iov_system": {
							Type: schema.TypeInt, Optional: true, Description: "Iov System",
						},
						"no_tx_queue": {
							Type: schema.TypeInt, Optional: true, Description: "No tx queue",
						},
						"wsocket_created": {
							Type: schema.TypeInt, Optional: true, Description: "wsocket created",
						},
						"wsocket_closed": {
							Type: schema.TypeInt, Optional: true, Description: "wsocket closed",
						},
						"a10_socket_created": {
							Type: schema.TypeInt, Optional: true, Description: "a10 socket created",
						},
						"a10_socket_closed": {
							Type: schema.TypeInt, Optional: true, Description: "a10 socket closed",
						},
						"no_a10_socket": {
							Type: schema.TypeInt, Optional: true, Description: "no a10 socket",
						},
						"no_other_side_socket": {
							Type: schema.TypeInt, Optional: true, Description: "no other side socket",
						},
						"no_w_engine": {
							Type: schema.TypeInt, Optional: true, Description: "no w engine",
						},
						"no_w_socket": {
							Type: schema.TypeInt, Optional: true, Description: "no w socket",
						},
						"on_ld_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "lost detection timeout",
						},
						"idle_alarm": {
							Type: schema.TypeInt, Optional: true, Description: "conn idle timeout",
						},
						"ack_alarm": {
							Type: schema.TypeInt, Optional: true, Description: "ack timeout",
						},
						"close_alarm": {
							Type: schema.TypeInt, Optional: true, Description: "close timeout",
						},
						"delay_alarm": {
							Type: schema.TypeInt, Optional: true, Description: "delay timeout",
						},
						"quic_malloc": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC malloc",
						},
						"quic_free": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC free",
						},
						"quic_malloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC malloc failure",
						},
						"quick_malloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "quick malloc failure",
						},
						"quic_lb": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC LB",
						},
						"cid_zero": {
							Type: schema.TypeInt, Optional: true, Description: "CID Zero",
						},
						"cid_cpu_hash": {
							Type: schema.TypeInt, Optional: true, Description: "CID CPU Hash",
						},
						"invalid_cid_sig": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid CID Signature",
						},
						"key_update_rx": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC TLS key update received",
						},
						"key_update_tx": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC TLS key update sent",
						},
					},
				},
			},
		},
	}
}

func resourceSlbQuicStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbQuicStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbQuicStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbQuicStatsStats := setObjectSlbQuicStatsStats(res)
		d.Set("stats", SlbQuicStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbQuicStatsStats(ret edpt.DataSlbQuicStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"client_conn_attempted":            ret.DtSlbQuicStats.Stats.Client_conn_attempted,
			"client_conn_handshake":            ret.DtSlbQuicStats.Stats.Client_conn_handshake,
			"client_conn_created":              ret.DtSlbQuicStats.Stats.Client_conn_created,
			"client_conn_local_closed":         ret.DtSlbQuicStats.Stats.Client_conn_local_closed,
			"client_conn_remote_closed":        ret.DtSlbQuicStats.Stats.Client_conn_remote_closed,
			"client_conn_failed":               ret.DtSlbQuicStats.Stats.Client_conn_failed,
			"server_conn_attempted":            ret.DtSlbQuicStats.Stats.Server_conn_attempted,
			"server_conn_handshake":            ret.DtSlbQuicStats.Stats.Server_conn_handshake,
			"server_conn_created":              ret.DtSlbQuicStats.Stats.Server_conn_created,
			"server_conn_local_closed":         ret.DtSlbQuicStats.Stats.Server_conn_local_closed,
			"server_conn_remote_closed":        ret.DtSlbQuicStats.Stats.Server_conn_remote_closed,
			"server_conn_failed":               ret.DtSlbQuicStats.Stats.Server_conn_failed,
			"q_conn_created":                   ret.DtSlbQuicStats.Stats.Q_conn_created,
			"q_conn_freed":                     ret.DtSlbQuicStats.Stats.Q_conn_freed,
			"local_bi_stream_current":          ret.DtSlbQuicStats.Stats.Local_bi_stream_current,
			"remote_bi_stream_current":         ret.DtSlbQuicStats.Stats.Remote_bi_stream_current,
			"local_bi_stream_created":          ret.DtSlbQuicStats.Stats.Local_bi_stream_created,
			"remote_bi_stream_created":         ret.DtSlbQuicStats.Stats.Remote_bi_stream_created,
			"local_bi_stream_closed":           ret.DtSlbQuicStats.Stats.Local_bi_stream_closed,
			"remote_bi_stream_closed":          ret.DtSlbQuicStats.Stats.Remote_bi_stream_closed,
			"local_uni_stream_current":         ret.DtSlbQuicStats.Stats.Local_uni_stream_current,
			"remote_uni_stream_current":        ret.DtSlbQuicStats.Stats.Remote_uni_stream_current,
			"local_uni_stream_created":         ret.DtSlbQuicStats.Stats.Local_uni_stream_created,
			"remote_uni_stream_created":        ret.DtSlbQuicStats.Stats.Remote_uni_stream_created,
			"local_uni_stream_closed":          ret.DtSlbQuicStats.Stats.Local_uni_stream_closed,
			"remote_uni_stream_closed":         ret.DtSlbQuicStats.Stats.Remote_uni_stream_closed,
			"stream_error":                     ret.DtSlbQuicStats.Stats.Stream_error,
			"stream_fail_to_insert":            ret.DtSlbQuicStats.Stats.Stream_fail_to_insert,
			"padding_frame_rx":                 ret.DtSlbQuicStats.Stats.Padding_frame_rx,
			"padding_frame_tx":                 ret.DtSlbQuicStats.Stats.Padding_frame_tx,
			"ping_frame_rx":                    ret.DtSlbQuicStats.Stats.Ping_frame_rx,
			"ping_frame_tx":                    ret.DtSlbQuicStats.Stats.Ping_frame_tx,
			"ack_frame_rx":                     ret.DtSlbQuicStats.Stats.Ack_frame_rx,
			"ack_frame_tx":                     ret.DtSlbQuicStats.Stats.Ack_frame_tx,
			"ack_ecn_frame_rx":                 ret.DtSlbQuicStats.Stats.Ack_ecn_frame_rx,
			"ack_ecn_frame_tx":                 ret.DtSlbQuicStats.Stats.Ack_ecn_frame_tx,
			"stream_rst_frame_rx":              ret.DtSlbQuicStats.Stats.Stream_rst_frame_rx,
			"stream_rst_frame_tx":              ret.DtSlbQuicStats.Stats.Stream_rst_frame_tx,
			"stream_stop_frame_rx":             ret.DtSlbQuicStats.Stats.Stream_stop_frame_rx,
			"stream_stop_frame_tx":             ret.DtSlbQuicStats.Stats.Stream_stop_frame_tx,
			"crypto_frame_rx":                  ret.DtSlbQuicStats.Stats.Crypto_frame_rx,
			"crypto_frame_tx":                  ret.DtSlbQuicStats.Stats.Crypto_frame_tx,
			"new_token_frame_rx":               ret.DtSlbQuicStats.Stats.New_token_frame_rx,
			"new_token_frame_tx":               ret.DtSlbQuicStats.Stats.New_token_frame_tx,
			"stream_frame_rx":                  ret.DtSlbQuicStats.Stats.Stream_frame_rx,
			"stream_frame_tx":                  ret.DtSlbQuicStats.Stats.Stream_frame_tx,
			"stream_09_frame_rx":               ret.DtSlbQuicStats.Stats.Stream_09_frame_rx,
			"stream_09_frame_tx":               ret.DtSlbQuicStats.Stats.Stream_09_frame_tx,
			"stream_0a_frame_rx":               ret.DtSlbQuicStats.Stats.Stream_0a_frame_rx,
			"stream_0a_frame_tx":               ret.DtSlbQuicStats.Stats.Stream_0a_frame_tx,
			"stream_0b_frame_rx":               ret.DtSlbQuicStats.Stats.Stream_0b_frame_rx,
			"stream_0b_frame_tx":               ret.DtSlbQuicStats.Stats.Stream_0b_frame_tx,
			"stream_0c_frame_rx":               ret.DtSlbQuicStats.Stats.Stream_0c_frame_rx,
			"stream_0c_frame_tx":               ret.DtSlbQuicStats.Stats.Stream_0c_frame_tx,
			"stream_0d_frame_rx":               ret.DtSlbQuicStats.Stats.Stream_0d_frame_rx,
			"stream_0d_frame_tx":               ret.DtSlbQuicStats.Stats.Stream_0d_frame_tx,
			"stream_0e_frame_rx":               ret.DtSlbQuicStats.Stats.Stream_0e_frame_rx,
			"stream_0e_frame_tx":               ret.DtSlbQuicStats.Stats.Stream_0e_frame_tx,
			"stream_0f_frame_rx":               ret.DtSlbQuicStats.Stats.Stream_0f_frame_rx,
			"stream_0f_frame_tx":               ret.DtSlbQuicStats.Stats.Stream_0f_frame_tx,
			"max_data_frame_rx":                ret.DtSlbQuicStats.Stats.Max_data_frame_rx,
			"max_data_frame_tx":                ret.DtSlbQuicStats.Stats.Max_data_frame_tx,
			"max_stream_data_frame_rx":         ret.DtSlbQuicStats.Stats.Max_stream_data_frame_rx,
			"max_stream_data_frame_tx":         ret.DtSlbQuicStats.Stats.Max_stream_data_frame_tx,
			"max_bi_stream_frame_rx":           ret.DtSlbQuicStats.Stats.Max_bi_stream_frame_rx,
			"max_bi_stream_frame_tx":           ret.DtSlbQuicStats.Stats.Max_bi_stream_frame_tx,
			"max_uni_stream_frame_rx":          ret.DtSlbQuicStats.Stats.Max_uni_stream_frame_rx,
			"max_uni_stream_frame_tx":          ret.DtSlbQuicStats.Stats.Max_uni_stream_frame_tx,
			"data_blocked_frame_rx":            ret.DtSlbQuicStats.Stats.Data_blocked_frame_rx,
			"data_blocked_frame_tx":            ret.DtSlbQuicStats.Stats.Data_blocked_frame_tx,
			"stream_data_blocked_frame_rx":     ret.DtSlbQuicStats.Stats.Stream_data_blocked_frame_rx,
			"stream_data_blocked_frame_tx":     ret.DtSlbQuicStats.Stats.Stream_data_blocked_frame_tx,
			"bi_stream_data_blocked_frame_rx":  ret.DtSlbQuicStats.Stats.Bi_stream_data_blocked_frame_rx,
			"bi_stream_data_blocked_frame_tx":  ret.DtSlbQuicStats.Stats.Bi_stream_data_blocked_frame_tx,
			"uni_stream_data_blocked_frame_rx": ret.DtSlbQuicStats.Stats.Uni_stream_data_blocked_frame_rx,
			"uni_stream_data_blocked_frame_tx": ret.DtSlbQuicStats.Stats.Uni_stream_data_blocked_frame_tx,
			"new_conn_id_frame_rx":             ret.DtSlbQuicStats.Stats.New_conn_id_frame_rx,
			"new_conn_id_frame_tx":             ret.DtSlbQuicStats.Stats.New_conn_id_frame_tx,
			"retire_conn_id_frame_rx":          ret.DtSlbQuicStats.Stats.Retire_conn_id_frame_rx,
			"retire_conn_id_frame_tx":          ret.DtSlbQuicStats.Stats.Retire_conn_id_frame_tx,
			"path_challenge_frame_rx":          ret.DtSlbQuicStats.Stats.Path_challenge_frame_rx,
			"path_challenge_frame_tx":          ret.DtSlbQuicStats.Stats.Path_challenge_frame_tx,
			"path_response_frame_rx":           ret.DtSlbQuicStats.Stats.Path_response_frame_rx,
			"path_response_frame_tx":           ret.DtSlbQuicStats.Stats.Path_response_frame_tx,
			"conn_close_frame_rx":              ret.DtSlbQuicStats.Stats.Conn_close_frame_rx,
			"conn_close_frame_tx":              ret.DtSlbQuicStats.Stats.Conn_close_frame_tx,
			"app_conn_close_frame_rx":          ret.DtSlbQuicStats.Stats.App_conn_close_frame_rx,
			"app_conn_close_frame_tx":          ret.DtSlbQuicStats.Stats.App_conn_close_frame_tx,
			"handshake_done_frame_rx":          ret.DtSlbQuicStats.Stats.Handshake_done_frame_rx,
			"handshake_done_frame_tx":          ret.DtSlbQuicStats.Stats.Handshake_done_frame_tx,
			"unknown_frame":                    ret.DtSlbQuicStats.Stats.Unknown_frame,
			"stream_fin_receive":               ret.DtSlbQuicStats.Stats.Stream_fin_receive,
			"stream_fin_up":                    ret.DtSlbQuicStats.Stats.Stream_fin_up,
			"stream_fin_down":                  ret.DtSlbQuicStats.Stats.Stream_fin_down,
			"stream_fin_send":                  ret.DtSlbQuicStats.Stats.Stream_fin_send,
			"stream_congest":                   ret.DtSlbQuicStats.Stats.Stream_congest,
			"stream_open":                      ret.DtSlbQuicStats.Stats.Stream_open,
			"stream_pause_data":                ret.DtSlbQuicStats.Stats.Stream_pause_data,
			"stream_resume_data":               ret.DtSlbQuicStats.Stats.Stream_resume_data,
			"stream_not_send":                  ret.DtSlbQuicStats.Stats.Stream_not_send,
			"stream_stop_send":                 ret.DtSlbQuicStats.Stats.Stream_stop_send,
			"stream_created":                   ret.DtSlbQuicStats.Stats.Stream_created,
			"stream_freed":                     ret.DtSlbQuicStats.Stats.Stream_freed,
			"initial_rx":                       ret.DtSlbQuicStats.Stats.Initial_rx,
			"initial_tx":                       ret.DtSlbQuicStats.Stats.Initial_tx,
			"rtt_0_rx":                         ret.DtSlbQuicStats.Stats.Rtt_0_rx,
			"rtt_0_tx":                         ret.DtSlbQuicStats.Stats.Rtt_0_tx,
			"handshake_rx":                     ret.DtSlbQuicStats.Stats.Handshake_rx,
			"handshake_tx":                     ret.DtSlbQuicStats.Stats.Handshake_tx,
			"retry_rx":                         ret.DtSlbQuicStats.Stats.Retry_rx,
			"retry_tx":                         ret.DtSlbQuicStats.Stats.Retry_tx,
			"ver_rx":                           ret.DtSlbQuicStats.Stats.Ver_rx,
			"ver_tx":                           ret.DtSlbQuicStats.Stats.Ver_tx,
			"rtt_updated":                      ret.DtSlbQuicStats.Stats.Rtt_updated,
			"needs_ack":                        ret.DtSlbQuicStats.Stats.Needs_ack,
			"delayed_ack":                      ret.DtSlbQuicStats.Stats.Delayed_ack,
			"packet_rx":                        ret.DtSlbQuicStats.Stats.Packet_rx,
			"packet_tx":                        ret.DtSlbQuicStats.Stats.Packet_tx,
			"packet_tx_failed":                 ret.DtSlbQuicStats.Stats.Packet_tx_failed,
			"congest_wnd_inc":                  ret.DtSlbQuicStats.Stats.Congest_wnd_inc,
			"congest_wnd_dec":                  ret.DtSlbQuicStats.Stats.Congest_wnd_dec,
			"no_congest_wnd":                   ret.DtSlbQuicStats.Stats.No_congest_wnd,
			"burst_limited":                    ret.DtSlbQuicStats.Stats.Burst_limited,
			"packet_loop_limited":              ret.DtSlbQuicStats.Stats.Packet_loop_limited,
			"receive_wnd_limited":              ret.DtSlbQuicStats.Stats.Receive_wnd_limited,
			"parse_error":                      ret.DtSlbQuicStats.Stats.Parse_error,
			"error_close":                      ret.DtSlbQuicStats.Stats.Error_close,
			"unknown_scid":                     ret.DtSlbQuicStats.Stats.Unknown_scid,
			"dcid_mismatch":                    ret.DtSlbQuicStats.Stats.Dcid_mismatch,
			"packet_too_short":                 ret.DtSlbQuicStats.Stats.Packet_too_short,
			"invalid_version":                  ret.DtSlbQuicStats.Stats.Invalid_version,
			"invalid_packet":                   ret.DtSlbQuicStats.Stats.Invalid_packet,
			"invalid_conn_match":               ret.DtSlbQuicStats.Stats.Invalid_conn_match,
			"invalid_session_packet":           ret.DtSlbQuicStats.Stats.Invalid_session_packet,
			"stateless_reset":                  ret.DtSlbQuicStats.Stats.Stateless_reset,
			"packet_lost":                      ret.DtSlbQuicStats.Stats.Packet_lost,
			"packet_drop":                      ret.DtSlbQuicStats.Stats.Packet_drop,
			"packet_retransmit":                ret.DtSlbQuicStats.Stats.Packet_retransmit,
			"packet_out_of_order":              ret.DtSlbQuicStats.Stats.Packet_out_of_order,
			"quic_packet_drop":                 ret.DtSlbQuicStats.Stats.Quic_packet_drop,
			"encode_error":                     ret.DtSlbQuicStats.Stats.Encode_error,
			"decode_failed":                    ret.DtSlbQuicStats.Stats.Decode_failed,
			"decode_stream_error":              ret.DtSlbQuicStats.Stats.Decode_stream_error,
			"exceed_flow_control":              ret.DtSlbQuicStats.Stats.Exceed_flow_control,
			"crypto_stream_not_found":          ret.DtSlbQuicStats.Stats.Crypto_stream_not_found,
			"exceed_max_stream_id":             ret.DtSlbQuicStats.Stats.Exceed_max_stream_id,
			"stream_id_mismatch":               ret.DtSlbQuicStats.Stats.Stream_id_mismatch,
			"ack_delay_huge":                   ret.DtSlbQuicStats.Stats.Ack_delay_huge,
			"ack_rng_huge_1":                   ret.DtSlbQuicStats.Stats.Ack_rng_huge_1,
			"ack_rng_huge_2":                   ret.DtSlbQuicStats.Stats.Ack_rng_huge_2,
			"ack_rng_huge_3":                   ret.DtSlbQuicStats.Stats.Ack_rng_huge_3,
			"too_noisy_fuzzing":                ret.DtSlbQuicStats.Stats.Too_noisy_fuzzing,
			"max_stream_too_big":               ret.DtSlbQuicStats.Stats.Max_stream_too_big,
			"stream_blocked":                   ret.DtSlbQuicStats.Stats.Stream_blocked,
			"new_conn_id_len_zero":             ret.DtSlbQuicStats.Stats.New_conn_id_len_zero,
			"new_conn_id_len_non_zero":         ret.DtSlbQuicStats.Stats.New_conn_id_len_non_zero,
			"illegal_stream_len":               ret.DtSlbQuicStats.Stats.Illegal_stream_len,
			"illegal_reason_len":               ret.DtSlbQuicStats.Stats.Illegal_reason_len,
			"illegal_seq":                      ret.DtSlbQuicStats.Stats.Illegal_seq,
			"illegal_rpt":                      ret.DtSlbQuicStats.Stats.Illegal_rpt,
			"illegal_len":                      ret.DtSlbQuicStats.Stats.Illegal_len,
			"illegal_token_len":                ret.DtSlbQuicStats.Stats.Illegal_token_len,
			"cannot_insert_cid":                ret.DtSlbQuicStats.Stats.Cannot_insert_cid,
			"cannot_insert_srt":                ret.DtSlbQuicStats.Stats.Cannot_insert_srt,
			"cannot_retire_cid":                ret.DtSlbQuicStats.Stats.Cannot_retire_cid,
			"no_next_scid":                     ret.DtSlbQuicStats.Stats.No_next_scid,
			"token_len_too_long":               ret.DtSlbQuicStats.Stats.Token_len_too_long,
			"server_receive_new_token":         ret.DtSlbQuicStats.Stats.Server_receive_new_token,
			"zero_frame_packet":                ret.DtSlbQuicStats.Stats.Zero_frame_packet,
			"err_frame_dec1":                   ret.DtSlbQuicStats.Stats.Err_frame_dec1,
			"err_frame_dec":                    ret.DtSlbQuicStats.Stats.Err_frame_dec,
			"err_frame_decb":                   ret.DtSlbQuicStats.Stats.Err_frame_decb,
			"err_frame_final_size":             ret.DtSlbQuicStats.Stats.Err_frame_final_size,
			"err_flow_control":                 ret.DtSlbQuicStats.Stats.Err_flow_control,
			"err_protocol_violation":           ret.DtSlbQuicStats.Stats.Err_protocol_violation,
			"server_rx_handshake_done":         ret.DtSlbQuicStats.Stats.Server_rx_handshake_done,
			"pkt_acked_failed":                 ret.DtSlbQuicStats.Stats.Pkt_acked_failed,
			"pn_insert_failed":                 ret.DtSlbQuicStats.Stats.Pn_insert_failed,
			"pn_delete_failed":                 ret.DtSlbQuicStats.Stats.Pn_delete_failed,
			"acked_packet_freed":               ret.DtSlbQuicStats.Stats.Acked_packet_freed,
			"tx_buffer_enq":                    ret.DtSlbQuicStats.Stats.Tx_buffer_enq,
			"tx_buffer_deq":                    ret.DtSlbQuicStats.Stats.Tx_buffer_deq,
			"app_buffer_enq":                   ret.DtSlbQuicStats.Stats.App_buffer_enq,
			"app_buffer_deq":                   ret.DtSlbQuicStats.Stats.App_buffer_deq,
			"app_buffer_queue_full":            ret.DtSlbQuicStats.Stats.App_buffer_queue_full,
			"iov_buffer_bind":                  ret.DtSlbQuicStats.Stats.Iov_buffer_bind,
			"iov_buffer_unbind":                ret.DtSlbQuicStats.Stats.Iov_buffer_unbind,
			"iov_buffer_dup":                   ret.DtSlbQuicStats.Stats.Iov_buffer_dup,
			"iov_alloc_len":                    ret.DtSlbQuicStats.Stats.Iov_alloc_len,
			"iov_io":                           ret.DtSlbQuicStats.Stats.Iov_io,
			"iov_system":                       ret.DtSlbQuicStats.Stats.Iov_system,
			"no_tx_queue":                      ret.DtSlbQuicStats.Stats.No_tx_queue,
			"wsocket_created":                  ret.DtSlbQuicStats.Stats.Wsocket_created,
			"wsocket_closed":                   ret.DtSlbQuicStats.Stats.Wsocket_closed,
			"a10_socket_created":               ret.DtSlbQuicStats.Stats.A10_socket_created,
			"a10_socket_closed":                ret.DtSlbQuicStats.Stats.A10_socket_closed,
			"no_a10_socket":                    ret.DtSlbQuicStats.Stats.No_a10_socket,
			"no_other_side_socket":             ret.DtSlbQuicStats.Stats.No_other_side_socket,
			"no_w_engine":                      ret.DtSlbQuicStats.Stats.No_w_engine,
			"no_w_socket":                      ret.DtSlbQuicStats.Stats.No_w_socket,
			"on_ld_timeout":                    ret.DtSlbQuicStats.Stats.On_ld_timeout,
			"idle_alarm":                       ret.DtSlbQuicStats.Stats.Idle_alarm,
			"ack_alarm":                        ret.DtSlbQuicStats.Stats.Ack_alarm,
			"close_alarm":                      ret.DtSlbQuicStats.Stats.Close_alarm,
			"delay_alarm":                      ret.DtSlbQuicStats.Stats.Delay_alarm,
			"quic_malloc":                      ret.DtSlbQuicStats.Stats.Quic_malloc,
			"quic_free":                        ret.DtSlbQuicStats.Stats.Quic_free,
			"quic_malloc_failure":              ret.DtSlbQuicStats.Stats.Quic_malloc_failure,
			"quick_malloc_failure":             ret.DtSlbQuicStats.Stats.Quick_malloc_failure,
			"quic_lb":                          ret.DtSlbQuicStats.Stats.Quic_lb,
			"cid_zero":                         ret.DtSlbQuicStats.Stats.Cid_zero,
			"cid_cpu_hash":                     ret.DtSlbQuicStats.Stats.Cid_cpu_hash,
			"invalid_cid_sig":                  ret.DtSlbQuicStats.Stats.Invalid_cid_sig,
			"key_update_rx":                    ret.DtSlbQuicStats.Stats.Key_update_rx,
			"key_update_tx":                    ret.DtSlbQuicStats.Stats.Key_update_tx,
		},
	}
}

func getObjectSlbQuicStatsStats(d []interface{}) edpt.SlbQuicStatsStats {

	count1 := len(d)
	var ret edpt.SlbQuicStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Client_conn_attempted = in["client_conn_attempted"].(int)
		ret.Client_conn_handshake = in["client_conn_handshake"].(int)
		ret.Client_conn_created = in["client_conn_created"].(int)
		ret.Client_conn_local_closed = in["client_conn_local_closed"].(int)
		ret.Client_conn_remote_closed = in["client_conn_remote_closed"].(int)
		ret.Client_conn_failed = in["client_conn_failed"].(int)
		ret.Server_conn_attempted = in["server_conn_attempted"].(int)
		ret.Server_conn_handshake = in["server_conn_handshake"].(int)
		ret.Server_conn_created = in["server_conn_created"].(int)
		ret.Server_conn_local_closed = in["server_conn_local_closed"].(int)
		ret.Server_conn_remote_closed = in["server_conn_remote_closed"].(int)
		ret.Server_conn_failed = in["server_conn_failed"].(int)
		ret.Q_conn_created = in["q_conn_created"].(int)
		ret.Q_conn_freed = in["q_conn_freed"].(int)
		ret.Local_bi_stream_current = in["local_bi_stream_current"].(int)
		ret.Remote_bi_stream_current = in["remote_bi_stream_current"].(int)
		ret.Local_bi_stream_created = in["local_bi_stream_created"].(int)
		ret.Remote_bi_stream_created = in["remote_bi_stream_created"].(int)
		ret.Local_bi_stream_closed = in["local_bi_stream_closed"].(int)
		ret.Remote_bi_stream_closed = in["remote_bi_stream_closed"].(int)
		ret.Local_uni_stream_current = in["local_uni_stream_current"].(int)
		ret.Remote_uni_stream_current = in["remote_uni_stream_current"].(int)
		ret.Local_uni_stream_created = in["local_uni_stream_created"].(int)
		ret.Remote_uni_stream_created = in["remote_uni_stream_created"].(int)
		ret.Local_uni_stream_closed = in["local_uni_stream_closed"].(int)
		ret.Remote_uni_stream_closed = in["remote_uni_stream_closed"].(int)
		ret.Stream_error = in["stream_error"].(int)
		ret.Stream_fail_to_insert = in["stream_fail_to_insert"].(int)
		ret.Padding_frame_rx = in["padding_frame_rx"].(int)
		ret.Padding_frame_tx = in["padding_frame_tx"].(int)
		ret.Ping_frame_rx = in["ping_frame_rx"].(int)
		ret.Ping_frame_tx = in["ping_frame_tx"].(int)
		ret.Ack_frame_rx = in["ack_frame_rx"].(int)
		ret.Ack_frame_tx = in["ack_frame_tx"].(int)
		ret.Ack_ecn_frame_rx = in["ack_ecn_frame_rx"].(int)
		ret.Ack_ecn_frame_tx = in["ack_ecn_frame_tx"].(int)
		ret.Stream_rst_frame_rx = in["stream_rst_frame_rx"].(int)
		ret.Stream_rst_frame_tx = in["stream_rst_frame_tx"].(int)
		ret.Stream_stop_frame_rx = in["stream_stop_frame_rx"].(int)
		ret.Stream_stop_frame_tx = in["stream_stop_frame_tx"].(int)
		ret.Crypto_frame_rx = in["crypto_frame_rx"].(int)
		ret.Crypto_frame_tx = in["crypto_frame_tx"].(int)
		ret.New_token_frame_rx = in["new_token_frame_rx"].(int)
		ret.New_token_frame_tx = in["new_token_frame_tx"].(int)
		ret.Stream_frame_rx = in["stream_frame_rx"].(int)
		ret.Stream_frame_tx = in["stream_frame_tx"].(int)
		ret.Stream_09_frame_rx = in["stream_09_frame_rx"].(int)
		ret.Stream_09_frame_tx = in["stream_09_frame_tx"].(int)
		ret.Stream_0a_frame_rx = in["stream_0a_frame_rx"].(int)
		ret.Stream_0a_frame_tx = in["stream_0a_frame_tx"].(int)
		ret.Stream_0b_frame_rx = in["stream_0b_frame_rx"].(int)
		ret.Stream_0b_frame_tx = in["stream_0b_frame_tx"].(int)
		ret.Stream_0c_frame_rx = in["stream_0c_frame_rx"].(int)
		ret.Stream_0c_frame_tx = in["stream_0c_frame_tx"].(int)
		ret.Stream_0d_frame_rx = in["stream_0d_frame_rx"].(int)
		ret.Stream_0d_frame_tx = in["stream_0d_frame_tx"].(int)
		ret.Stream_0e_frame_rx = in["stream_0e_frame_rx"].(int)
		ret.Stream_0e_frame_tx = in["stream_0e_frame_tx"].(int)
		ret.Stream_0f_frame_rx = in["stream_0f_frame_rx"].(int)
		ret.Stream_0f_frame_tx = in["stream_0f_frame_tx"].(int)
		ret.Max_data_frame_rx = in["max_data_frame_rx"].(int)
		ret.Max_data_frame_tx = in["max_data_frame_tx"].(int)
		ret.Max_stream_data_frame_rx = in["max_stream_data_frame_rx"].(int)
		ret.Max_stream_data_frame_tx = in["max_stream_data_frame_tx"].(int)
		ret.Max_bi_stream_frame_rx = in["max_bi_stream_frame_rx"].(int)
		ret.Max_bi_stream_frame_tx = in["max_bi_stream_frame_tx"].(int)
		ret.Max_uni_stream_frame_rx = in["max_uni_stream_frame_rx"].(int)
		ret.Max_uni_stream_frame_tx = in["max_uni_stream_frame_tx"].(int)
		ret.Data_blocked_frame_rx = in["data_blocked_frame_rx"].(int)
		ret.Data_blocked_frame_tx = in["data_blocked_frame_tx"].(int)
		ret.Stream_data_blocked_frame_rx = in["stream_data_blocked_frame_rx"].(int)
		ret.Stream_data_blocked_frame_tx = in["stream_data_blocked_frame_tx"].(int)
		ret.Bi_stream_data_blocked_frame_rx = in["bi_stream_data_blocked_frame_rx"].(int)
		ret.Bi_stream_data_blocked_frame_tx = in["bi_stream_data_blocked_frame_tx"].(int)
		ret.Uni_stream_data_blocked_frame_rx = in["uni_stream_data_blocked_frame_rx"].(int)
		ret.Uni_stream_data_blocked_frame_tx = in["uni_stream_data_blocked_frame_tx"].(int)
		ret.New_conn_id_frame_rx = in["new_conn_id_frame_rx"].(int)
		ret.New_conn_id_frame_tx = in["new_conn_id_frame_tx"].(int)
		ret.Retire_conn_id_frame_rx = in["retire_conn_id_frame_rx"].(int)
		ret.Retire_conn_id_frame_tx = in["retire_conn_id_frame_tx"].(int)
		ret.Path_challenge_frame_rx = in["path_challenge_frame_rx"].(int)
		ret.Path_challenge_frame_tx = in["path_challenge_frame_tx"].(int)
		ret.Path_response_frame_rx = in["path_response_frame_rx"].(int)
		ret.Path_response_frame_tx = in["path_response_frame_tx"].(int)
		ret.Conn_close_frame_rx = in["conn_close_frame_rx"].(int)
		ret.Conn_close_frame_tx = in["conn_close_frame_tx"].(int)
		ret.App_conn_close_frame_rx = in["app_conn_close_frame_rx"].(int)
		ret.App_conn_close_frame_tx = in["app_conn_close_frame_tx"].(int)
		ret.Handshake_done_frame_rx = in["handshake_done_frame_rx"].(int)
		ret.Handshake_done_frame_tx = in["handshake_done_frame_tx"].(int)
		ret.Unknown_frame = in["unknown_frame"].(int)
		ret.Stream_fin_receive = in["stream_fin_receive"].(int)
		ret.Stream_fin_up = in["stream_fin_up"].(int)
		ret.Stream_fin_down = in["stream_fin_down"].(int)
		ret.Stream_fin_send = in["stream_fin_send"].(int)
		ret.Stream_congest = in["stream_congest"].(int)
		ret.Stream_open = in["stream_open"].(int)
		ret.Stream_pause_data = in["stream_pause_data"].(int)
		ret.Stream_resume_data = in["stream_resume_data"].(int)
		ret.Stream_not_send = in["stream_not_send"].(int)
		ret.Stream_stop_send = in["stream_stop_send"].(int)
		ret.Stream_created = in["stream_created"].(int)
		ret.Stream_freed = in["stream_freed"].(int)
		ret.Initial_rx = in["initial_rx"].(int)
		ret.Initial_tx = in["initial_tx"].(int)
		ret.Rtt_0_rx = in["rtt_0_rx"].(int)
		ret.Rtt_0_tx = in["rtt_0_tx"].(int)
		ret.Handshake_rx = in["handshake_rx"].(int)
		ret.Handshake_tx = in["handshake_tx"].(int)
		ret.Retry_rx = in["retry_rx"].(int)
		ret.Retry_tx = in["retry_tx"].(int)
		ret.Ver_rx = in["ver_rx"].(int)
		ret.Ver_tx = in["ver_tx"].(int)
		ret.Rtt_updated = in["rtt_updated"].(int)
		ret.Needs_ack = in["needs_ack"].(int)
		ret.Delayed_ack = in["delayed_ack"].(int)
		ret.Packet_rx = in["packet_rx"].(int)
		ret.Packet_tx = in["packet_tx"].(int)
		ret.Packet_tx_failed = in["packet_tx_failed"].(int)
		ret.Congest_wnd_inc = in["congest_wnd_inc"].(int)
		ret.Congest_wnd_dec = in["congest_wnd_dec"].(int)
		ret.No_congest_wnd = in["no_congest_wnd"].(int)
		ret.Burst_limited = in["burst_limited"].(int)
		ret.Packet_loop_limited = in["packet_loop_limited"].(int)
		ret.Receive_wnd_limited = in["receive_wnd_limited"].(int)
		ret.Parse_error = in["parse_error"].(int)
		ret.Error_close = in["error_close"].(int)
		ret.Unknown_scid = in["unknown_scid"].(int)
		ret.Dcid_mismatch = in["dcid_mismatch"].(int)
		ret.Packet_too_short = in["packet_too_short"].(int)
		ret.Invalid_version = in["invalid_version"].(int)
		ret.Invalid_packet = in["invalid_packet"].(int)
		ret.Invalid_conn_match = in["invalid_conn_match"].(int)
		ret.Invalid_session_packet = in["invalid_session_packet"].(int)
		ret.Stateless_reset = in["stateless_reset"].(int)
		ret.Packet_lost = in["packet_lost"].(int)
		ret.Packet_drop = in["packet_drop"].(int)
		ret.Packet_retransmit = in["packet_retransmit"].(int)
		ret.Packet_out_of_order = in["packet_out_of_order"].(int)
		ret.Quic_packet_drop = in["quic_packet_drop"].(int)
		ret.Encode_error = in["encode_error"].(int)
		ret.Decode_failed = in["decode_failed"].(int)
		ret.Decode_stream_error = in["decode_stream_error"].(int)
		ret.Exceed_flow_control = in["exceed_flow_control"].(int)
		ret.Crypto_stream_not_found = in["crypto_stream_not_found"].(int)
		ret.Exceed_max_stream_id = in["exceed_max_stream_id"].(int)
		ret.Stream_id_mismatch = in["stream_id_mismatch"].(int)
		ret.Ack_delay_huge = in["ack_delay_huge"].(int)
		ret.Ack_rng_huge_1 = in["ack_rng_huge_1"].(int)
		ret.Ack_rng_huge_2 = in["ack_rng_huge_2"].(int)
		ret.Ack_rng_huge_3 = in["ack_rng_huge_3"].(int)
		ret.Too_noisy_fuzzing = in["too_noisy_fuzzing"].(int)
		ret.Max_stream_too_big = in["max_stream_too_big"].(int)
		ret.Stream_blocked = in["stream_blocked"].(int)
		ret.New_conn_id_len_zero = in["new_conn_id_len_zero"].(int)
		ret.New_conn_id_len_non_zero = in["new_conn_id_len_non_zero"].(int)
		ret.Illegal_stream_len = in["illegal_stream_len"].(int)
		ret.Illegal_reason_len = in["illegal_reason_len"].(int)
		ret.Illegal_seq = in["illegal_seq"].(int)
		ret.Illegal_rpt = in["illegal_rpt"].(int)
		ret.Illegal_len = in["illegal_len"].(int)
		ret.Illegal_token_len = in["illegal_token_len"].(int)
		ret.Cannot_insert_cid = in["cannot_insert_cid"].(int)
		ret.Cannot_insert_srt = in["cannot_insert_srt"].(int)
		ret.Cannot_retire_cid = in["cannot_retire_cid"].(int)
		ret.No_next_scid = in["no_next_scid"].(int)
		ret.Token_len_too_long = in["token_len_too_long"].(int)
		ret.Server_receive_new_token = in["server_receive_new_token"].(int)
		ret.Zero_frame_packet = in["zero_frame_packet"].(int)
		ret.Err_frame_dec1 = in["err_frame_dec1"].(int)
		ret.Err_frame_dec = in["err_frame_dec"].(int)
		ret.Err_frame_decb = in["err_frame_decb"].(int)
		ret.Err_frame_final_size = in["err_frame_final_size"].(int)
		ret.Err_flow_control = in["err_flow_control"].(int)
		ret.Err_protocol_violation = in["err_protocol_violation"].(int)
		ret.Server_rx_handshake_done = in["server_rx_handshake_done"].(int)
		ret.Pkt_acked_failed = in["pkt_acked_failed"].(int)
		ret.Pn_insert_failed = in["pn_insert_failed"].(int)
		ret.Pn_delete_failed = in["pn_delete_failed"].(int)
		ret.Acked_packet_freed = in["acked_packet_freed"].(int)
		ret.Tx_buffer_enq = in["tx_buffer_enq"].(int)
		ret.Tx_buffer_deq = in["tx_buffer_deq"].(int)
		ret.App_buffer_enq = in["app_buffer_enq"].(int)
		ret.App_buffer_deq = in["app_buffer_deq"].(int)
		ret.App_buffer_queue_full = in["app_buffer_queue_full"].(int)
		ret.Iov_buffer_bind = in["iov_buffer_bind"].(int)
		ret.Iov_buffer_unbind = in["iov_buffer_unbind"].(int)
		ret.Iov_buffer_dup = in["iov_buffer_dup"].(int)
		ret.Iov_alloc_len = in["iov_alloc_len"].(int)
		ret.Iov_io = in["iov_io"].(int)
		ret.Iov_system = in["iov_system"].(int)
		ret.No_tx_queue = in["no_tx_queue"].(int)
		ret.Wsocket_created = in["wsocket_created"].(int)
		ret.Wsocket_closed = in["wsocket_closed"].(int)
		ret.A10_socket_created = in["a10_socket_created"].(int)
		ret.A10_socket_closed = in["a10_socket_closed"].(int)
		ret.No_a10_socket = in["no_a10_socket"].(int)
		ret.No_other_side_socket = in["no_other_side_socket"].(int)
		ret.No_w_engine = in["no_w_engine"].(int)
		ret.No_w_socket = in["no_w_socket"].(int)
		ret.On_ld_timeout = in["on_ld_timeout"].(int)
		ret.Idle_alarm = in["idle_alarm"].(int)
		ret.Ack_alarm = in["ack_alarm"].(int)
		ret.Close_alarm = in["close_alarm"].(int)
		ret.Delay_alarm = in["delay_alarm"].(int)
		ret.Quic_malloc = in["quic_malloc"].(int)
		ret.Quic_free = in["quic_free"].(int)
		ret.Quic_malloc_failure = in["quic_malloc_failure"].(int)
		ret.Quick_malloc_failure = in["quick_malloc_failure"].(int)
		ret.Quic_lb = in["quic_lb"].(int)
		ret.Cid_zero = in["cid_zero"].(int)
		ret.Cid_cpu_hash = in["cid_cpu_hash"].(int)
		ret.Invalid_cid_sig = in["invalid_cid_sig"].(int)
		ret.Key_update_rx = in["key_update_rx"].(int)
		ret.Key_update_tx = in["key_update_tx"].(int)
	}
	return ret
}

func dataToEndpointSlbQuicStats(d *schema.ResourceData) edpt.SlbQuicStats {
	var ret edpt.SlbQuicStats

	ret.Stats = getObjectSlbQuicStatsStats(d.Get("stats").([]interface{}))
	return ret
}
