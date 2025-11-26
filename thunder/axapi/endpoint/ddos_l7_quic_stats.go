package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosL7QuicStats struct {
	Stats DdosL7QuicStatsStats `json:"stats"`
}
type DataDdosL7QuicStats struct {
	DtDdosL7QuicStats DdosL7QuicStats `json:"l7-quic"`
}

type DdosL7QuicStatsStats struct {
	Quic_packet_received                int `json:"quic_packet_received"`
	Quic_initial_received               int `json:"quic_initial_received"`
	Quic_version_negotiation_received   int `json:"quic_version_negotiation_received"`
	Quic_retry_received                 int `json:"quic_retry_received"`
	Quic_0rtt_recevied                  int `json:"quic_0rtt_recevied"`
	Quic_handshake_received             int `json:"quic_handshake_received"`
	Quic_version_match_action_taken     int `json:"quic_version_match_action_taken"`
	Quic_version_match_action_drop      int `json:"quic_version_match_action_drop"`
	Quic_version_match_action_blacklist int `json:"quic_version_match_action_blacklist"`
	Quic_malformed_action_taken         int `json:"quic_malformed_action_taken"`
	Quic_malformed_action_drop          int `json:"quic_malformed_action_drop"`
	Quic_malformed_action_blacklist     int `json:"quic_malformed_action_blacklist"`
	Quic_malformed_dcid_len_max_exceed  int `json:"quic_malformed_dcid_len_max_exceed"`
	Quic_malformed_scid_len_max_exceed  int `json:"quic_malformed_scid_len_max_exceed"`
	Quic_fixed_bit_not_set              int `json:"quic_fixed_bit_not_set"`
	Quic_retry_auth_sent                int `json:"quic_retry_auth_sent"`
	Quic_retry_auth_pass                int `json:"quic_retry_auth_pass"`
	Quic_retry_auth_fail                int `json:"quic_retry_auth_fail"`
	Quic_connection_close_sent          int `json:"quic_connection_close_sent"`
	Quic_invalid_retry_token            int `json:"quic_invalid_retry_token"`
	Quic_short_header_received          int `json:"quic_short_header_received"`
	Quic_short_header_action_drop       int `json:"quic_short_header_action_drop"`
	Quic_encrypt_fail                   int `json:"quic_encrypt_fail"`
	Quic_decrypt_fail                   int `json:"quic_decrypt_fail"`
	Quic_encrypt_success                int `json:"quic_encrypt_success"`
	Quic_decrypt_success                int `json:"quic_decrypt_success"`
	Quic_0rtt_drop                      int `json:"quic_0rtt_drop"`
	Quic_aead_pkt_rate_exceed           int `json:"quic_aead_pkt_rate_exceed"`
	Quic_dcid_pkt_rate_exceed           int `json:"quic_dcid_pkt_rate_exceed"`
	Quic_create_conn_init_only          int `json:"quic_create_conn_init_only"`
	Quic_version_no_match_drop          int `json:"quic_version_no_match_drop"`
}

func (p *DdosL7QuicStats) GetId() string {
	return "1"
}

func (p *DdosL7QuicStats) getPath() string {
	return "ddos/l7-quic/stats"
}

func (p *DdosL7QuicStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL7QuicStats, error) {
	logger.Println("DdosL7QuicStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosL7QuicStats
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
