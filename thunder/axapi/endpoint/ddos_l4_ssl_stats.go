

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL4SslStats struct {
    
    Stats DdosL4SslStatsStats `json:"stats"`

}
type DataDdosL4SslStats struct {
    DtDdosL4SslStats DdosL4SslStats `json:"l4-ssl"`
}


type DdosL4SslStatsStats struct {
    Ssl_l4_policy_reset int `json:"ssl_l4_policy_reset"`
    Ssl_l4_policy_drop int `json:"ssl_l4_policy_drop"`
    Ssl_l4_drop_packet int `json:"ssl_l4_drop_packet"`
    Ssl_l4_er_condition int `json:"ssl_l4_er_condition"`
    Ssl_l4_processed int `json:"ssl_l4_processed"`
    Ssl_l4_new_syn int `json:"ssl_l4_new_syn"`
    Ssl_l4_is_ssl3 int `json:"ssl_l4_is_ssl3"`
    Ssl_l4_is_tls1_1 int `json:"ssl_l4_is_tls1_1"`
    Ssl_l4_is_tls1_2 int `json:"ssl_l4_is_tls1_2"`
    Ssl_l4_is_tls1_3 int `json:"ssl_l4_is_tls1_3"`
    Ssl_l4_is_renegotiation int `json:"ssl_l4_is_renegotiation"`
    Ssl_l4_renegotiation_exceed int `json:"ssl_l4_renegotiation_exceed"`
    Ssl_l4_is_dst_req_rate_exceed int `json:"ssl_l4_is_dst_req_rate_exceed"`
    Ssl_l4_is_src_req_rate_exceed int `json:"ssl_l4_is_src_req_rate_exceed"`
    Ssl_l4_do_auth_handshake int `json:"ssl_l4_do_auth_handshake"`
    Ssl_l4_reset_for_handshake int `json:"ssl_l4_reset_for_handshake"`
    Ssl_l4_handshake_timeout int `json:"ssl_l4_handshake_timeout"`
    Ssl_l4_auth_handshake_ok int `json:"ssl_l4_auth_handshake_ok"`
    Ssl_l4_auth_handshake_bl int `json:"ssl_l4_auth_handshake_bl"`
    Ssl_l4_auth_handshake_fail int `json:"ssl_l4_auth_handshake_fail"`
    Ssl_l4_renegotiation_incomplete int `json:"ssl_l4_renegotiation_incomplete"`
    Ssl_l4_auth_drop int `json:"ssl_l4_auth_drop"`
    Ssl_l4_auth_resp int `json:"ssl_l4_auth_resp"`
    Ssl_non_tls int `json:"ssl_non_tls"`
    Ssl_header_invalid_type int `json:"ssl_header_invalid_type"`
    Ssl_header_bad_ver int `json:"ssl_header_bad_ver"`
    Ssl_header_bad_len int `json:"ssl_header_bad_len"`
    Ssl_bad_header_forw int `json:"ssl_bad_header_forw"`
    Ssl_bad_header_drop int `json:"ssl_bad_header_drop"`
}

func (p *DdosL4SslStats) GetId() string{
    return "1"
}

func (p *DdosL4SslStats) getPath() string{
    return "ddos/l4-ssl/stats"
}

func (p *DdosL4SslStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL4SslStats,error) {
logger.Println("DdosL4SslStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL4SslStats
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
