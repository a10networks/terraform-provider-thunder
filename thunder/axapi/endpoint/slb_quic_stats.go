

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbQuicStats struct {
    
    Stats SlbQuicStatsStats `json:"stats"`

}
type DataSlbQuicStats struct {
    DtSlbQuicStats SlbQuicStats `json:"quic"`
}


type SlbQuicStatsStats struct {
    Client_conn_attempted int `json:"client_conn_attempted"`
    Client_conn_handshake int `json:"client_conn_handshake"`
    Client_conn_created int `json:"client_conn_created"`
    Client_conn_local_closed int `json:"client_conn_local_closed"`
    Client_conn_remote_closed int `json:"client_conn_remote_closed"`
    Client_conn_failed int `json:"client_conn_failed"`
    Server_conn_attempted int `json:"server_conn_attempted"`
    Server_conn_handshake int `json:"server_conn_handshake"`
    Server_conn_created int `json:"server_conn_created"`
    Server_conn_local_closed int `json:"server_conn_local_closed"`
    Server_conn_remote_closed int `json:"server_conn_remote_closed"`
    Server_conn_failed int `json:"server_conn_failed"`
    Q_conn_created int `json:"q_conn_created"`
    Q_conn_freed int `json:"q_conn_freed"`
    Local_bi_stream_current int `json:"local_bi_stream_current"`
    Remote_bi_stream_current int `json:"remote_bi_stream_current"`
    Local_bi_stream_created int `json:"local_bi_stream_created"`
    Remote_bi_stream_created int `json:"remote_bi_stream_created"`
    Local_bi_stream_closed int `json:"local_bi_stream_closed"`
    Remote_bi_stream_closed int `json:"remote_bi_stream_closed"`
    Local_uni_stream_current int `json:"local_uni_stream_current"`
    Remote_uni_stream_current int `json:"remote_uni_stream_current"`
    Local_uni_stream_created int `json:"local_uni_stream_created"`
    Remote_uni_stream_created int `json:"remote_uni_stream_created"`
    Local_uni_stream_closed int `json:"local_uni_stream_closed"`
    Remote_uni_stream_closed int `json:"remote_uni_stream_closed"`
    Stream_error int `json:"stream_error"`
    Stream_fail_to_insert int `json:"stream_fail_to_insert"`
    Padding_frame_rx int `json:"padding_frame_rx"`
    Padding_frame_tx int `json:"padding_frame_tx"`
    Ping_frame_rx int `json:"ping_frame_rx"`
    Ping_frame_tx int `json:"ping_frame_tx"`
    Ack_frame_rx int `json:"ack_frame_rx"`
    Ack_frame_tx int `json:"ack_frame_tx"`
    Ack_ecn_frame_rx int `json:"ack_ecn_frame_rx"`
    Ack_ecn_frame_tx int `json:"ack_ecn_frame_tx"`
    Stream_rst_frame_rx int `json:"stream_rst_frame_rx"`
    Stream_rst_frame_tx int `json:"stream_rst_frame_tx"`
    Stream_stop_frame_rx int `json:"stream_stop_frame_rx"`
    Stream_stop_frame_tx int `json:"stream_stop_frame_tx"`
    Crypto_frame_rx int `json:"crypto_frame_rx"`
    Crypto_frame_tx int `json:"crypto_frame_tx"`
    New_token_frame_rx int `json:"new_token_frame_rx"`
    New_token_frame_tx int `json:"new_token_frame_tx"`
    Stream_frame_rx int `json:"stream_frame_rx"`
    Stream_frame_tx int `json:"stream_frame_tx"`
    Stream_09_frame_rx int `json:"stream_09_frame_rx"`
    Stream_09_frame_tx int `json:"stream_09_frame_tx"`
    Stream_0a_frame_rx int `json:"stream_0a_frame_rx"`
    Stream_0a_frame_tx int `json:"stream_0a_frame_tx"`
    Stream_0b_frame_rx int `json:"stream_0b_frame_rx"`
    Stream_0b_frame_tx int `json:"stream_0b_frame_tx"`
    Stream_0c_frame_rx int `json:"stream_0c_frame_rx"`
    Stream_0c_frame_tx int `json:"stream_0c_frame_tx"`
    Stream_0d_frame_rx int `json:"stream_0d_frame_rx"`
    Stream_0d_frame_tx int `json:"stream_0d_frame_tx"`
    Stream_0e_frame_rx int `json:"stream_0e_frame_rx"`
    Stream_0e_frame_tx int `json:"stream_0e_frame_tx"`
    Stream_0f_frame_rx int `json:"stream_0f_frame_rx"`
    Stream_0f_frame_tx int `json:"stream_0f_frame_tx"`
    Max_data_frame_rx int `json:"max_data_frame_rx"`
    Max_data_frame_tx int `json:"max_data_frame_tx"`
    Max_stream_data_frame_rx int `json:"max_stream_data_frame_rx"`
    Max_stream_data_frame_tx int `json:"max_stream_data_frame_tx"`
    Max_bi_stream_frame_rx int `json:"max_bi_stream_frame_rx"`
    Max_bi_stream_frame_tx int `json:"max_bi_stream_frame_tx"`
    Max_uni_stream_frame_rx int `json:"max_uni_stream_frame_rx"`
    Max_uni_stream_frame_tx int `json:"max_uni_stream_frame_tx"`
    Data_blocked_frame_rx int `json:"data_blocked_frame_rx"`
    Data_blocked_frame_tx int `json:"data_blocked_frame_tx"`
    Stream_data_blocked_frame_rx int `json:"stream_data_blocked_frame_rx"`
    Stream_data_blocked_frame_tx int `json:"stream_data_blocked_frame_tx"`
    Bi_stream_data_blocked_frame_rx int `json:"bi_stream_data_blocked_frame_rx"`
    Bi_stream_data_blocked_frame_tx int `json:"bi_stream_data_blocked_frame_tx"`
    Uni_stream_data_blocked_frame_rx int `json:"uni_stream_data_blocked_frame_rx"`
    Uni_stream_data_blocked_frame_tx int `json:"uni_stream_data_blocked_frame_tx"`
    New_conn_id_frame_rx int `json:"new_conn_id_frame_rx"`
    New_conn_id_frame_tx int `json:"new_conn_id_frame_tx"`
    Retire_conn_id_frame_rx int `json:"retire_conn_id_frame_rx"`
    Retire_conn_id_frame_tx int `json:"retire_conn_id_frame_tx"`
    Path_challenge_frame_rx int `json:"path_challenge_frame_rx"`
    Path_challenge_frame_tx int `json:"path_challenge_frame_tx"`
    Path_response_frame_rx int `json:"path_response_frame_rx"`
    Path_response_frame_tx int `json:"path_response_frame_tx"`
    Conn_close_frame_rx int `json:"conn_close_frame_rx"`
    Conn_close_frame_tx int `json:"conn_close_frame_tx"`
    App_conn_close_frame_rx int `json:"app_conn_close_frame_rx"`
    App_conn_close_frame_tx int `json:"app_conn_close_frame_tx"`
    Handshake_done_frame_rx int `json:"handshake_done_frame_rx"`
    Handshake_done_frame_tx int `json:"handshake_done_frame_tx"`
    Unknown_frame int `json:"unknown_frame"`
    Stream_fin_receive int `json:"stream_fin_receive"`
    Stream_fin_up int `json:"stream_fin_up"`
    Stream_fin_down int `json:"stream_fin_down"`
    Stream_fin_send int `json:"stream_fin_send"`
    Stream_congest int `json:"stream_congest"`
    Stream_open int `json:"stream_open"`
    Stream_pause_data int `json:"stream_pause_data"`
    Stream_resume_data int `json:"stream_resume_data"`
    Stream_not_send int `json:"stream_not_send"`
    Stream_stop_send int `json:"stream_stop_send"`
    Stream_created int `json:"stream_created"`
    Stream_freed int `json:"stream_freed"`
    Initial_rx int `json:"INITIAL_rx"`
    Initial_tx int `json:"INITIAL_tx"`
    Rtt_0_rx int `json:"RTT_0_rx"`
    Rtt_0_tx int `json:"RTT_0_tx"`
    Handshake_rx int `json:"HANDSHAKE_rx"`
    Handshake_tx int `json:"HANDSHAKE_tx"`
    Retry_rx int `json:"RETRY_rx"`
    Retry_tx int `json:"RETRY_tx"`
    Ver_rx int `json:"VER_rx"`
    Ver_tx int `json:"VER_tx"`
    Rtt_updated int `json:"RTT_updated"`
    Needs_ack int `json:"Needs_ack"`
    Delayed_ack int `json:"Delayed_ack"`
    Packet_rx int `json:"Packet_rx"`
    Packet_tx int `json:"Packet_tx"`
    Packet_tx_failed int `json:"Packet_tx_failed"`
    Congest_wnd_inc int `json:"Congest_wnd_inc"`
    Congest_wnd_dec int `json:"Congest_wnd_dec"`
    No_congest_wnd int `json:"No_congest_wnd"`
    Burst_limited int `json:"Burst_limited"`
    Packet_loop_limited int `json:"Packet_loop_limited"`
    Receive_wnd_limited int `json:"Receive_wnd_limited"`
    Parse_error int `json:"Parse_error"`
    Error_close int `json:"Error_close"`
    Unknown_scid int `json:"Unknown_scid"`
    Dcid_mismatch int `json:"Dcid_mismatch"`
    Packet_too_short int `json:"Packet_too_short"`
    Invalid_version int `json:"Invalid_version"`
    Invalid_packet int `json:"Invalid_Packet"`
    Invalid_conn_match int `json:"Invalid_conn_match"`
    Invalid_session_packet int `json:"Invalid_session_packet"`
    Stateless_reset int `json:"Stateless_reset"`
    Packet_lost int `json:"Packet_lost"`
    Packet_drop int `json:"Packet_drop"`
    Packet_retransmit int `json:"Packet_retransmit"`
    Packet_out_of_order int `json:"Packet_out_of_order"`
    Quic_packet_drop int `json:"Quic_packet_drop"`
    Encode_error int `json:"Encode_error"`
    Decode_failed int `json:"Decode_failed"`
    Decode_stream_error int `json:"Decode_stream_error"`
    Exceed_flow_control int `json:"Exceed_flow_control"`
    Crypto_stream_not_found int `json:"Crypto_stream_not_found"`
    Exceed_max_stream_id int `json:"Exceed_max_stream_id"`
    Stream_id_mismatch int `json:"Stream_id_mismatch"`
    Ack_delay_huge int `json:"Ack_delay_huge"`
    Ack_rng_huge_1 int `json:"Ack_rng_huge_1"`
    Ack_rng_huge_2 int `json:"Ack_rng_huge_2"`
    Ack_rng_huge_3 int `json:"Ack_rng_huge_3"`
    Too_noisy_fuzzing int `json:"Too_noisy_fuzzing"`
    Max_stream_too_big int `json:"Max_stream_too_big"`
    Stream_blocked int `json:"Stream_blocked"`
    New_conn_id_len_zero int `json:"New_conn_id_len_zero"`
    New_conn_id_len_non_zero int `json:"New_conn_id_len_non_zero"`
    Illegal_stream_len int `json:"Illegal_stream_len"`
    Illegal_reason_len int `json:"Illegal_reason_len"`
    Illegal_seq int `json:"Illegal_seq"`
    Illegal_rpt int `json:"Illegal_rpt"`
    Illegal_len int `json:"Illegal_len"`
    Illegal_token_len int `json:"Illegal_token_len"`
    Cannot_insert_cid int `json:"Cannot_insert_cid"`
    Cannot_insert_srt int `json:"Cannot_insert_srt"`
    Cannot_retire_cid int `json:"Cannot_retire_cid"`
    No_next_scid int `json:"No_next_scid"`
    Token_len_too_long int `json:"Token_len_too_long"`
    Server_receive_new_token int `json:"Server_receive_new_token"`
    Zero_frame_packet int `json:"Zero_frame_packet"`
    Err_frame_dec1 int `json:"Err_frame_dec1"`
    Err_frame_dec int `json:"Err_frame_dec"`
    Err_frame_decb int `json:"Err_frame_decb"`
    Err_frame_final_size int `json:"Err_frame_final_size"`
    Err_flow_control int `json:"Err_flow_control"`
    Err_protocol_violation int `json:"Err_protocol_violation"`
    Server_rx_handshake_done int `json:"Server_rx_handshake_done"`
    Pkt_acked_failed int `json:"Pkt_acked_failed"`
    Pn_insert_failed int `json:"Pn_insert_failed"`
    Pn_delete_failed int `json:"Pn_delete_failed"`
    Acked_packet_freed int `json:"Acked_packet_freed"`
    Tx_buffer_enq int `json:"Tx_buffer_enq"`
    Tx_buffer_deq int `json:"Tx_buffer_deq"`
    App_buffer_enq int `json:"App_buffer_enq"`
    App_buffer_deq int `json:"App_buffer_deq"`
    App_buffer_queue_full int `json:"App_buffer_queue_full"`
    Iov_buffer_bind int `json:"Iov_buffer_bind"`
    Iov_buffer_unbind int `json:"Iov_buffer_unbind"`
    Iov_buffer_dup int `json:"Iov_buffer_dup"`
    Iov_alloc_len int `json:"Iov_alloc_len"`
    Iov_io int `json:"Iov_IO"`
    Iov_system int `json:"Iov_System"`
    No_tx_queue int `json:"No_tx_queue"`
    Wsocket_created int `json:"wsocket_created"`
    Wsocket_closed int `json:"wsocket_closed"`
    A10_socket_created int `json:"a10_socket_created"`
    A10_socket_closed int `json:"a10_socket_closed"`
    No_a10_socket int `json:"No_a10_socket"`
    No_other_side_socket int `json:"No_other_side_socket"`
    No_w_engine int `json:"No_w_engine"`
    No_w_socket int `json:"No_w_socket"`
    On_ld_timeout int `json:"on_ld_timeout"`
    Idle_alarm int `json:"idle_alarm"`
    Ack_alarm int `json:"ack_alarm"`
    Close_alarm int `json:"close_alarm"`
    Delay_alarm int `json:"delay_alarm"`
    Quic_malloc int `json:"quic_malloc"`
    Quic_free int `json:"quic_free"`
    Quic_malloc_failure int `json:"quic_malloc_failure"`
    Quick_malloc_failure int `json:"quick_malloc_failure"`
    Quic_lb int `json:"quic_lb"`
    Cid_zero int `json:"cid_zero"`
    Cid_cpu_hash int `json:"cid_cpu_hash"`
    Invalid_cid_sig int `json:"invalid_cid_sig"`
    Key_update_rx int `json:"key_update_rx"`
    Key_update_tx int `json:"key_update_tx"`
}

func (p *SlbQuicStats) GetId() string{
    return "1"
}

func (p *SlbQuicStats) getPath() string{
    return "slb/quic/stats"
}

func (p *SlbQuicStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbQuicStats,error) {
logger.Println("SlbQuicStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbQuicStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
